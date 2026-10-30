package evals_test

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/checks"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project"
)

const historyConfig = `version: 1
base: main
modules: ["packages/*"]
protected: [locked/**]
module_defaults:
  evals: [evals/**]
  interface: [api.go]
  append_only: [ledger/**]
`

type historyFixture struct {
	t    *testing.T
	root string
	base string
}

func (f historyFixture) git(args ...string) string {
	f.t.Helper()
	result, err := gitstore.Run(context.Background(), gitstore.Request{
		Args: append([]string{"git", "-C", f.root}, args...),
		Dir:  f.root,
		Env: map[string]string{
			"GIT_CONFIG_GLOBAL": "/dev/null", "GIT_CONFIG_NOSYSTEM": "1",
			"GIT_AUTHOR_NAME": "History Fixture", "GIT_AUTHOR_EMAIL": "fixture@example.invalid",
			"GIT_COMMITTER_NAME": "History Fixture", "GIT_COMMITTER_EMAIL": "fixture@example.invalid",
		},
	})
	if err != nil {
		f.t.Fatalf("git %v: %v: %s", args, err, result.Stderr)
	}
	return strings.TrimSpace(result.Stdout)
}

func (f historyFixture) put(path, body string) {
	f.t.Helper()
	full := filepath.Join(f.root, filepath.FromSlash(path))
	if err := os.MkdirAll(filepath.Dir(full), 0755); err != nil {
		f.t.Fatal(err)
	}
	if err := os.WriteFile(full, []byte(body), 0644); err != nil {
		f.t.Fatal(err)
	}
}

func (f historyFixture) remove(path string) {
	f.t.Helper()
	if err := os.Remove(filepath.Join(f.root, filepath.FromSlash(path))); err != nil {
		f.t.Fatal(err)
	}
}

func (f historyFixture) move(oldPath, newPath string) {
	f.t.Helper()
	if err := os.MkdirAll(filepath.Dir(filepath.Join(f.root, filepath.FromSlash(newPath))), 0755); err != nil {
		f.t.Fatal(err)
	}
	f.git("mv", oldPath, newPath)
}

func (f historyFixture) commit(subject string) string {
	f.t.Helper()
	f.git("add", "-A")
	f.git("commit", "-qm", subject)
	return f.git("rev-parse", "HEAD")
}

func (f historyFixture) change(path, phase, set string) {
	f.t.Helper()
	body := fmt.Sprintf("---\nmodule: alpha\nphase: %s\n%s---\n\n## Motivation\nHistory fixture.\n\n## Interface\n\n## Added\n\n## Modified\n\n## Removed\n\n## Decisions\n\n## Notes\n", phase, set)
	if _, err := contracts.ParseDocument(body); err != nil {
		f.t.Fatalf("fixture change: %v", err)
	}
	f.put(path, body)
}

func newHistoryFixture(t *testing.T) historyFixture {
	t.Helper()
	f := historyFixture{t: t, root: t.TempDir()}
	f.git("init", "-q", "-b", "main")
	f.put("fireweed.yaml", historyConfig)
	f.put("packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\n---\n# Alpha\n\n## Specs\n\n### ALP-001\nAlpha behavior.\n")
	f.put("packages/alpha/api.go", "package alpha\n")
	f.put("packages/alpha/evals/existing_test.go", "package evals\n")
	f.put("packages/alpha/ledger/entry.txt", "initial\n")
	f.put("packages/alpha/newledger/entry.txt", "initial\n")
	f.put("locked/config.txt", "initial\n")
	f.base = f.commit("base")
	if _, err := project.LoadAt(f.root, f.base); err != nil {
		t.Fatalf("recorded base project: %v", err)
	}
	return f
}

func (f historyFixture) report(revisionRange string) checks.HistoryReport {
	f.t.Helper()
	report, err := checks.History(f.root, revisionRange)
	if err != nil {
		f.t.Fatalf("History(%q): %v", revisionRange, err)
	}
	return report
}

type historyFinding struct{ code, path, commit string }

func expectHistoryFindings(t *testing.T, report checks.HistoryReport, want ...historyFinding) {
	t.Helper()
	if len(report.Diagnostics) != len(want) {
		t.Fatalf("diagnostics = %#v; want %d findings", report.Diagnostics, len(want))
	}
	used := make([]bool, len(report.Diagnostics))
	for _, expected := range want {
		found := false
		for i, d := range report.Diagnostics {
			if !used[i] && d.Code == expected.code && d.Path == expected.path && strings.Contains(d.Message, expected.commit) {
				used[i], found = true, true
				break
			}
		}
		if !found {
			t.Errorf("missing %s on %s at %s in %#v", expected.code, expected.path, expected.commit, report.Diagnostics)
		}
	}
}

func Test_CHK_070_HistoryExplicitRangeOrderAndFullOIDs(t *testing.T) {
	f := newHistoryFixture(t)
	f.put("notes/one.txt", "one")
	one := f.commit("first")
	f.put("notes/two.txt", "two")
	two := f.commit("second")
	got := f.report(f.base + ".." + two)
	if !reflect.DeepEqual(got.Commits, []string{one, two}) {
		t.Errorf("commits = %#v; want [%s %s]", got.Commits, one, two)
	}
	expectHistoryFindings(t, got)
}

func Test_CHK_070_HistoryDefaultRangeUsesLocalConfiguredBase(t *testing.T) {
	f := newHistoryFixture(t)
	f.git("branch", "feature")
	f.git("checkout", "-q", "feature")
	f.put("notes/one.txt", "one")
	one := f.commit("first")
	f.put("notes/two.txt", "two")
	two := f.commit("second")
	got := f.report("")
	if !reflect.DeepEqual(got.Commits, []string{one, two}) {
		t.Errorf("default commits = %#v; want [%s %s]", got.Commits, one, two)
	}
}

func Test_CHK_071_CHK_072_HistoryPhaseContentsAndDeletionControl(t *testing.T) {
	f := newHistoryFixture(t)
	path := "changes/0001-work.md"
	f.change(path, "evals", "")
	f.commit("add queued change")
	f.change(path, "audit", "")
	f.put("locked/config.txt", "worker edit\n")
	worker := f.commit("ordinary subject")
	f.remove(path)
	f.put("locked/config.txt", "author edit\n")
	f.commit("fireweed impl 0001-work")
	got := f.report(f.base + "..HEAD")
	expectHistoryFindings(t, got, historyFinding{"H01", "locked/config.txt", worker})
}

func Test_CHK_072_CHK_073_CHK_077_HistoryPlaybookOverrideAndProjectedInterface(t *testing.T) {
	f := newHistoryFixture(t)
	pending := "changes/0001-work.md"
	archive := "changes/archive/0001-work.md"
	set := "set:\n  interface: [newguard/**]\n"
	f.change(pending, "impl", set)
	f.commit("ready")
	projected, err := project.Project(f.root, "")
	if err != nil {
		t.Fatalf("Project with interface set: %v", err)
	}
	if len(projected.Modules) != 1 || projected.Modules[0].Name != "alpha" || !reflect.DeepEqual(projected.Modules[0].Zones.Interface, []string{"newguard/**"}) {
		t.Fatalf("projected alpha interface = %#v; want [newguard/**]", projected.Modules)
	}
	f.move(pending, archive)
	f.change(archive, "applied", set)
	f.put("packages/alpha/newguard/api.go", "package newguard\n")
	f.put(".fireweed/playbooks/evals.md", "local override\n")
	applied := f.commit("phase assigned")
	got := f.report(f.base + "..HEAD")
	expectHistoryFindings(t, got,
		historyFinding{"H02", "packages/alpha/newguard/api.go", applied},
		historyFinding{"H01", ".fireweed/playbooks/evals.md", applied})
}

func Test_CHK_071_CHK_073_HistoryAppliedImplTouchesEvalsAndInterface(t *testing.T) {
	f := newHistoryFixture(t)
	pending := "changes/0001-work.md"
	archive := "changes/archive/0001-work.md"
	f.change(pending, "impl", "")
	f.commit("ready for implementation")
	f.move(pending, archive)
	f.change(archive, "applied", "")
	f.put("packages/alpha/evals/existing_test.go", "changed test\n")
	f.put("packages/alpha/api.go", "changed api\n")
	applied := f.commit("unrelated subject")
	got := f.report(f.base + "..HEAD")
	expectHistoryFindings(t, got,
		historyFinding{"H02", "packages/alpha/evals/existing_test.go", applied},
		historyFinding{"H02", "packages/alpha/api.go", applied})
}

func Test_CHK_074_HistoryAuthorAppendOnlyModifyDeleteRename(t *testing.T) {
	for _, action := range []string{"modify", "delete", "rename"} {
		t.Run(action, func(t *testing.T) {
			f := newHistoryFixture(t)
			path := "packages/alpha/ledger/entry.txt"
			switch action {
			case "modify":
				f.put(path, "modified\n")
			case "delete":
				f.remove(path)
			case "rename":
				f.move(path, "packages/alpha/ledger/renamed.txt")
			}
			commit := f.commit("human maintenance")
			got := f.report(f.base + "..HEAD")
			if len(got.Diagnostics) == 0 {
				t.Errorf("%s diagnostics = %#v; want H03", action, got.Diagnostics)
			}
			for _, d := range got.Diagnostics {
				if d.Code != "H03" || !strings.Contains(d.Message, commit) || (d.Path != path && d.Path != "packages/alpha/ledger/renamed.txt") {
					t.Errorf("%s unexpected diagnostic %#v", action, d)
				}
			}
		})
	}
}

func Test_CHK_074_HistoryMergeResolutionInMergeOnlyRange(t *testing.T) {
	f := newHistoryFixture(t)
	f.git("checkout", "-qb", "side")
	f.put("packages/alpha/ledger/entry.txt", "side\n")
	f.commit("side edits ledger")
	f.git("checkout", "-q", "main")
	f.put("packages/alpha/ledger/entry.txt", "main\n")
	f.commit("main edits ledger")
	result, err := gitstore.Run(context.Background(), gitstore.Request{Args: []string{"git", "-C", f.root, "merge", "--no-ff", "--no-commit", "side"}, Dir: f.root})
	if err == nil || !strings.Contains(result.Stdout+result.Stderr, "CONFLICT") {
		t.Fatalf("expected ordinary merge conflict, got %v: %s %s", err, result.Stdout, result.Stderr)
	}
	f.put("packages/alpha/ledger/entry.txt", "resolved uniquely\n")
	merge := f.commit("resolve conflict")
	got := f.report(merge + "^!")
	if !reflect.DeepEqual(got.Commits, []string{merge}) {
		t.Errorf("merge-only selection = %#v; want [%s]", got.Commits, merge)
	}
	expectHistoryFindings(t, got, historyFinding{"H03", "packages/alpha/ledger/entry.txt", merge})
}

func Test_CHK_075_HistoryArchiveModifyAndDeletionControl(t *testing.T) {
	f := newHistoryFixture(t)
	path := "changes/archive/0001-old.md"
	f.change(path, "applied", "")
	f.commit("archived baseline")
	f.put(path, "---\nmodule: alpha\nphase: applied\n---\n\nUpdated archive.\n")
	modified := f.commit("edit archive")
	f.remove(path)
	f.commit("delete archive")
	got := f.report(f.base + "..HEAD")
	expectHistoryFindings(t, got, historyFinding{"H04", path, modified})
}

func Test_CHK_076_HistoryMultipleSquashedChangesKeepH03H04(t *testing.T) {
	f := newHistoryFixture(t)
	archiveOld := "changes/archive/0000-old.md"
	f.change(archiveOld, "applied", "")
	archiveBaseline := f.commit("existing archive")
	for _, path := range []string{"changes/archive/0001-one.md", "changes/archive/0002-two.md"} {
		f.change(path, "applied", "")
	}
	f.put("locked/config.txt", "squash protection\n")
	f.put("packages/alpha/evals/existing_test.go", "squash eval\n")
	f.put("packages/alpha/ledger/entry.txt", "squash ledger\n")
	f.put(archiveOld, "---\nmodule: alpha\nphase: applied\n---\n\nChanged old archive.\n")
	squash := f.commit("squashed changes")
	got := f.report(archiveBaseline + "..HEAD")
	wantSquashed := map[checks.SquashedChange]bool{
		{Commit: squash, Change: "changes/archive/0001-one.md"}: true,
		{Commit: squash, Change: "changes/archive/0002-two.md"}: true,
	}
	if len(got.Squashed) != len(wantSquashed) {
		t.Errorf("squashed = %#v; want %#v", got.Squashed, wantSquashed)
	} else {
		for _, entry := range got.Squashed {
			if !wantSquashed[entry] {
				t.Errorf("unexpected squashed entry %#v; want %#v", entry, wantSquashed)
			}
			delete(wantSquashed, entry)
		}
		if len(wantSquashed) != 0 {
			t.Errorf("missing squashed entries %#v", wantSquashed)
		}
	}
	expectHistoryFindings(t, got,
		historyFinding{"H03", "packages/alpha/ledger/entry.txt", squash},
		historyFinding{"H04", archiveOld, squash})
}

func Test_CHK_077_HistoryUsesRecordedBeforeAfterZones(t *testing.T) {
	f := newHistoryFixture(t)
	pending := "changes/0001-work.md"
	f.change(pending, "impl", "")
	f.commit("ready")
	f.move(pending, "changes/archive/0001-work.md")
	f.change("changes/archive/0001-work.md", "applied", "")
	f.put("packages/alpha/evals/new_test.go", "new eval\n")
	applied := f.commit("apply")
	f.put("fireweed.yaml", strings.Replace(historyConfig, "evals: [evals/**]", "evals: [later/**]", 1))
	f.commit("later config")
	got := f.report(f.base + ".." + applied)
	expectHistoryFindings(t, got, historyFinding{"H02", "packages/alpha/evals/new_test.go", applied})
}

func Test_CHK_077_HistoryChecksBothRecordedZoneMaps(t *testing.T) {
	f := newHistoryFixture(t)
	f.put("fireweed.yaml", strings.Replace(historyConfig, "append_only: [ledger/**]", "append_only: [newledger/**]", 1))
	f.put("packages/alpha/ledger/entry.txt", "changed under before zone\n")
	f.put("packages/alpha/newledger/entry.txt", "changed under after zone\n")
	commit := f.commit("author changes zone and both existing files")
	got := f.report(f.base + "..HEAD")
	expectHistoryFindings(t, got,
		historyFinding{"H03", "packages/alpha/ledger/entry.txt", commit},
		historyFinding{"H03", "packages/alpha/newledger/entry.txt", commit})
}

func Test_CHK_078_HistoryRejectsUnreadableRangeAndRecordedDocuments(t *testing.T) {
	for _, tc := range []struct {
		name       string
		edit       func(historyFixture)
		rangeValue string
	}{
		{"invalid_range", nil, "missing-ref..HEAD"},
		{"missing_config", func(f historyFixture) { f.remove("fireweed.yaml") }, ""},
		{"malformed_config", func(f historyFixture) { f.put("fireweed.yaml", "modules: [broken\n") }, ""},
		{"missing_SPEC", func(f historyFixture) { f.remove("packages/alpha/SPEC.md") }, ""},
		{"malformed_change", func(f historyFixture) { f.put("changes/0001-work.md", "---\nmodule: [broken\n---\n") }, ""},
	} {
		t.Run(tc.name, func(t *testing.T) {
			f := newHistoryFixture(t)
			if tc.edit != nil {
				tc.edit(f)
				f.commit("record invalid input")
			}
			revisionRange := tc.rangeValue
			if revisionRange == "" {
				revisionRange = f.base + "..HEAD"
			}
			if _, err := checks.History(f.root, revisionRange); err == nil {
				t.Errorf("History(%q) accepted %s", revisionRange, tc.name)
			}
		})
	}
}
