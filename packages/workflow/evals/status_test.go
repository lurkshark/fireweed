package evals_test

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow"
)

func status(t *testing.T, root string) workflow.StatusReport {
	t.Helper()
	r, err := workflow.Status(root)
	if err != nil {
		t.Fatalf("Status: %v", err)
	}
	return r
}

func statusEntry(t *testing.T, r workflow.StatusReport, numberedSlug string) workflow.StatusEntry {
	t.Helper()
	for _, entry := range r.Queue {
		if strings.TrimSuffix(filepath.Base(entry.Path), ".md") == numberedSlug {
			return entry
		}
	}
	t.Fatalf("Status queue has no %s: %#v", numberedSlug, r.Queue)
	return workflow.StatusEntry{}
}

func waiting(t *testing.T, r workflow.StatusReport, slug, on, rule string) {
	t.Helper()
	e := statusEntry(t, r, slug)
	if e.WaitingOn != on || e.WaitingRule != rule {
		t.Errorf("%s waiting = %q/%q, want %q/%q", slug, e.WaitingOn, e.WaitingRule, on, rule)
	}
}

func Test_FLOW_101_FLOW_102_FLOW_103_StatusQueueAndForwardWaiting(t *testing.T) {
	root := eligibilityBase(t, "alpha")
	eligibilityChange(t, root, 2, "same", "alpha", "", "")
	eligibilityChange(t, root, 3, "direct", "beta", "", "")
	eligibilityChange(t, root, 4, "transitive", "gamma", "", "")
	eligibilityChange(t, root, 5, "free", "delta", "", "")
	eligibilityChange(t, root, 6, "also-blocked", "beta", "", "")
	readyEligibilityFixture(t, root)
	freeFile := readFile(t, root, "changes/0005-free.md")
	write(t, root, "changes/0005-free.md", strings.Replace(freeFile, "module: delta\n", "module: delta\nphase: audit\naudit_rounds: 2\n", 1))
	blockedFile := readFile(t, root, "changes/0006-also-blocked.md")
	write(t, root, "changes/0006-also-blocked.md", strings.Replace(blockedFile, "module: beta\n", "module: beta\nblocked: separately blocked\n", 1))
	r := status(t, root)
	q, err := project.Queue(root)
	if err != nil {
		t.Fatal(err)
	}
	if len(r.Queue) != len(q) {
		t.Fatalf("queue length = %d, want %d", len(r.Queue), len(q))
	}
	for i, ch := range q {
		e := r.Queue[i]
		if e.Slug != ch.Slug || e.Path != ch.Path || e.Module != ch.Module {
			t.Errorf("queue[%d] identity = %#v, want %#v", i, e, ch)
		}
	}
	blocked := statusEntry(t, r, "0001-manual")
	if blocked.Phase != "impl" || blocked.Blocked != "fixture blocked" || blocked.AuditRounds != 0 {
		t.Errorf("blocked state = %#v", blocked)
	}
	free := statusEntry(t, r, "0005-free")
	if free.Phase != "audit" || free.Blocked != "" || free.AuditRounds != 2 {
		t.Errorf("authored state = %#v", free)
	}
	if same := statusEntry(t, r, "0002-same"); same.Phase != "queued" || same.Blocked != "" || same.AuditRounds != 0 {
		t.Errorf("default state = %#v", same)
	}
	waiting(t, r, "0001-manual", "", "")
	waiting(t, r, "0002-same", "0001-manual", "same_module")
	waiting(t, r, "0003-direct", "0001-manual", "depends_on_blocked")
	waiting(t, r, "0004-transitive", "0001-manual", "depends_on_blocked")
	waiting(t, r, "0005-free", "", "")
	waiting(t, r, "0006-also-blocked", "", "")
}

func Test_FLOW_104_FLOW_105_StatusInterfaceAndProjectedWaiting(t *testing.T) {
	root := eligibilityBase(t, "gamma")
	eligibilityChange(t, root, 2, "interface-direct", "beta", "", "Beta API changes.\n")
	eligibilityChange(t, root, 3, "internal-direct", "beta", "", "")
	eligibilityChange(t, root, 4, "interface-transitive", "alpha", "", "Alpha API changes.\n")
	eligibilityChange(t, root, 5, "internal-transitive", "alpha", "", "")
	readyEligibilityFixture(t, root)
	r := status(t, root)
	waiting(t, r, "0002-interface-direct", "0001-manual", "interface_dependency")
	waiting(t, r, "0003-internal-direct", "", "")
	waiting(t, r, "0004-interface-transitive", "0001-manual", "interface_dependency")
	waiting(t, r, "0005-internal-transitive", "", "")

	projected := eligibilityBase(t, "alpha")
	eligibilityChange(t, projected, 2, "own-add", "delta", "set:\n  depends_on: [alpha]\n", "")
	eligibilityChange(t, projected, 3, "detach-beta", "beta", "set:\n  depends_on: []\n", "")
	eligibilityChange(t, projected, 4, "after-detach", "gamma", "", "")
	eligibilityChange(t, projected, 5, "own-remove", "delta", "set:\n  depends_on: []\n", "")
	readyEligibilityFixture(t, projected)
	r = status(t, projected)
	waiting(t, r, "0002-own-add", "0001-manual", "depends_on_blocked")
	waiting(t, r, "0003-detach-beta", "", "")
	waiting(t, r, "0004-after-detach", "", "")
	waiting(t, r, "0005-own-remove", "", "")
}

func Test_FLOW_106_FLOW_107_FLOW_108_StatusModuleMetadata(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	alpha := readFile(t, root, "packages/alpha/SPEC.md")
	alpha = strings.Replace(alpha, "depends_on: []\n", "depends_on: []\nboundaries_opt_out: approved fixture exception\n", 1)
	alpha = strings.Replace(alpha, "### ALP-002\n", "### ALP-002 [quarantine: intermittent fixture]\n", 1)
	write(t, root, "packages/alpha/SPEC.md", alpha)
	write(t, root, "packages/delta/SPEC.md", "---\nname: delta\nid_prefix: DEL\ndepends_on: []\n---\n# Delta\n\n## Specs\n\n### DEL-001 [quarantine: unavailable service]\nDelta MUST work.\n")
	commit(t, root, "module metadata")
	snapshot, err := project.Load(root)
	if err != nil {
		t.Fatal(err)
	}
	r := status(t, root)
	if !reflect.DeepEqual(r.Modules, snapshot.Modules) {
		t.Errorf("modules = %#v, want discovered %#v", r.Modules, snapshot.Modules)
	}
	byName := map[string]project.Module{}
	for _, m := range r.Modules {
		byName[m.Name] = m
	}
	if len(byName) != 4 || byName["alpha"].IDPrefix != "ALP" || byName["alpha"].Dir != "packages/alpha" || !reflect.DeepEqual(byName["gamma"].DependsOn, []string{"beta"}) || !reflect.DeepEqual(byName["beta"].Zones.Interface, []string{"api.go"}) || !reflect.DeepEqual(byName["alpha"].Zones.Disposable, []string{"internal/**"}) {
		t.Errorf("discovered identity/dependencies/zones = %#v", byName)
	}
	wantQ := []workflow.Quarantine{{Module: "alpha", ID: "ALP-002", Reason: "intermittent fixture"}, {Module: "delta", ID: "DEL-001", Reason: "unavailable service"}}
	if !reflect.DeepEqual(r.Quarantines, wantQ) {
		t.Errorf("quarantines = %#v, want %#v", r.Quarantines, wantQ)
	}
	if !reflect.DeepEqual(r.BoundaryOptOuts, []workflow.BoundaryOptOut{{Module: "alpha", Reason: "approved fixture exception"}}) {
		t.Errorf("opt-outs = %#v", r.BoundaryOptOuts)
	}
}

func Test_FLOW_109_StatusProvenanceDrift(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	write(t, root, "packages/delta/SPEC.md", "---\nname: delta\nid_prefix: DEL\ndepends_on: []\n---\n# Delta\n\n## Specs\n")
	archivedImpl(t, root, "0002-alpha-old", "alpha", "2026-06-01T12:00:00Z")
	archivedImpl(t, root, "0003-beta-old", "beta", "2026-06-02T12:00:00Z")
	archivedImpl(t, root, "0003-gamma-old", "gamma", "2026-06-03T12:00:00Z")
	archivedImpl(t, root, "0004-alpha-new", "alpha", "2026-06-11T12:00:00Z")
	archivedImpl(t, root, "0005-alpha-newer", "alpha", "2026-06-12T12:00:00Z")
	archivedImpl(t, root, "0006-alpha-newest", "alpha", "2026-06-13T12:00:00Z")
	write(t, root, ".fireweed/drills/alpha.md", fmt.Sprintf("---\ndate: 2026-06-10T12:00:00Z\nbase_commit: %s\nagent: fixture\nmodel: fixture\nresult: fail\n---\n", head(t, root)))
	write(t, root, ".fireweed/drills/beta.md", fmt.Sprintf("---\ndate: 2026-06-14T12:00:00Z\nbase_commit: %s\nagent: fixture\nmodel: fixture\nresult: fail\n---\n", head(t, root)))
	commit(t, root, "authored archives and failed drill")
	r := status(t, root)
	for name, want := range map[string]int{"alpha": 3, "beta": 0, "gamma": 1, "delta": 0} {
		if got, ok := r.DriftSinceDrill[name]; !ok || got != want {
			t.Errorf("drift %s = %d, present %t; want %d", name, got, ok, want)
		}
	}
}

func Test_FLOW_111_StatusReworkChainWarning(t *testing.T) {
	const oldest = "0010-original"
	const first = "0009-add"
	const middle = "0008-modify"
	const last = "0002-rework"
	type edit struct{ section, body string }
	archive := func(t *testing.T, root, name string, day int, front string, edits ...edit) {
		t.Helper()
		archivedImpl(t, root, name, "alpha", fmt.Sprintf("2026-06-%02dT12:00:00Z", day))
		path := "changes/archive/" + name + ".md"
		doc := strings.Replace(readFile(t, root, path), "phase: applied\n", "phase: applied\n"+front, 1)
		for _, edit := range edits {
			doc = contracts.ReplaceSection(doc, edit.section, "\n"+edit.body+"\n")
		}
		write(t, root, path, doc)
	}
	for _, tc := range []struct {
		name        string
		gap         int
		front       string
		lastEdits   []edit
		latestMatch bool
		wantIDs     []string
	}{
		{name: "three linked changes", front: "fixes: [ALP-001]\n", lastEdits: []edit{{"Removed", "### ALP-002"}}, wantIDs: []string{"ALP-001", "ALP-002", "ALP-003"}},
		{name: "unrelated changes before tenth predecessor", gap: 9, front: "eval_fix: [ALP-001]\n", wantIDs: []string{"ALP-001", "ALP-003", "ALP-005"}},
		{name: "predecessor more than ten applied changes back", gap: 10, front: "eval_fix: [ALP-001]\n"},
		{name: "latest matching predecessor wins", front: "eval_fix: [ALP-001, ALP-004]\n", latestMatch: true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := fixture(t, "", "", "", "", "")
			wantFiles := []string{first + ".md", middle + ".md", last + ".md"}
			firstEdits := []edit{{"Added", "### ALP-003\nAlpha MUST expose a third value."}}
			if tc.gap > 0 {
				archive(t, root, oldest, 1, "", edit{"Added", "### ALP-005\nAlpha MUST expose a fifth value."})
				firstEdits = append(firstEdits, edit{"Modified", "### ALP-005\nAlpha MUST expose a nonempty fifth value."})
				wantFiles = append(wantFiles, oldest+".md")
			}
			archive(t, root, first, 2, "", firstEdits...)
			archive(t, root, middle, 3, "", edit{"Modified", "### ALP-003\nAlpha MUST expose a nonempty third value.\n\n" +
				"### ALP-001\n" + newClause + "\n\n### ALP-002\nAlpha MUST keep the second value stable."})
			for i := 0; i < tc.gap; i++ {
				archivedImpl(t, root, fmt.Sprintf("%04d-unrelated", 11+i), "beta", fmt.Sprintf("2026-06-%02dT12:00:00Z", 4+i))
			}
			if tc.latestMatch {
				archive(t, root, "0007-newer-match", 4, "", edit{"Added", "### ALP-004\nAlpha MUST expose a fourth value."})
			}
			archive(t, root, last, 5+tc.gap, tc.front, tc.lastEdits...)
			commit(t, root, "applied rework fixture")
			diagnostics := status(t, root).Diagnostics
			if len(tc.wantIDs) == 0 {
				if containsCode(diagnostics, "P08") {
					t.Errorf("unexpected P08: %#v", diagnostics)
				}
				return
			}
			d, ok := warningFor(diagnostics, "P08")
			if !ok || !d.Warning {
				t.Fatalf("expected a P08 warning: %#v", diagnostics)
			}
			for _, term := range append(wantFiles, tc.wantIDs...) {
				if !strings.Contains(d.Message, term) {
					t.Errorf("P08 message %q omits %q", d.Message, term)
				}
			}
			if strings.Contains(d.Message, "-unrelated.md") {
				t.Errorf("P08 message %q includes a change outside the chain", d.Message)
			}
			archivedImpl(t, root, "0020-unrelated-latest", "beta", "2026-06-20T12:00:00Z")
			commit(t, root, "apply unrelated change after the chain")
			if diagnostics := status(t, root).Diagnostics; containsCode(diagnostics, "P08") {
				t.Errorf("P08 warned for an older chain: %#v", diagnostics)
			}
		})
	}
	t.Run("same module without shared IDs", func(t *testing.T) {
		root := fixture(t, "", "", "", "", "")
		for i := 1; i <= 3; i++ {
			archive(t, root, fmt.Sprintf("%04d-independent", i+1), i, "", edit{"Added", fmt.Sprintf("### ALP-%03d\nAlpha MUST expose fixture value %d.", i+2, i)})
		}
		commit(t, root, "three independent alpha changes")
		if diagnostics := status(t, root).Diagnostics; containsCode(diagnostics, "P08") {
			t.Errorf("unexpected P08 for independent alpha changes: %#v", diagnostics)
		}
	})
}

func Test_FLOW_110_StatusSharedBaseUsesSameArchivedDocumentIdentity(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	pendingBase := head(t, root)
	git(t, root, "checkout", "-b", "feature")
	git(t, root, "checkout", "main")
	if err := os.MkdirAll(filepath.Join(root, "changes/archive"), 0755); err != nil {
		t.Fatal(err)
	}
	git(t, root, "mv", "changes/0001-manual.md", "changes/archive/0001-manual.md")
	mainArchive := readFile(t, root, "changes/archive/0001-manual.md")
	write(t, root, "changes/archive/0001-manual.md", strings.Replace(mainArchive, "module: alpha\n", "module: alpha\nphase: applied\n", 1))
	commit(t, root, "base has applied manual document")
	approved := head(t, root)
	git(t, root, "checkout", "feature")
	if err := os.MkdirAll(filepath.Join(root, "changes/archive"), 0755); err != nil {
		t.Fatal(err)
	}
	git(t, root, "mv", "changes/0001-manual.md", "changes/archive/0001-manual.md")
	change := readFile(t, root, "changes/archive/0001-manual.md")
	write(t, root, "changes/archive/0001-manual.md", strings.Replace(change, "module: alpha\n", "module: alpha\nphase: applied\n", 1))
	archivedImpl(t, root, "0002-extra", "beta", "2026-06-10T12:00:00Z")
	commit(t, root, "one approved and one new archive")
	if got := status(t, root).Unapproved; got != 1 {
		t.Errorf("local base unapproved = %d, want 1", got)
	}
	git(t, root, "update-ref", "refs/remotes/origin/main", head(t, root))
	if got := status(t, root).Unapproved; got != 0 {
		t.Errorf("remote base unapproved = %d, want 0", got)
	}
	git(t, root, "update-ref", "refs/remotes/origin/main", approved)
	if got := status(t, root).Unapproved; got != 1 {
		t.Errorf("shared base unapproved = %d, want 1", got)
	}
	git(t, root, "update-ref", "refs/remotes/origin/main", pendingBase)
	if got := status(t, root).Unapproved; got != 1 {
		t.Errorf("pending base identity unapproved = %d, want 1", got)
	}
	config := readFile(t, root, "fireweed.yaml")
	write(t, root, "fireweed.yaml", strings.Replace(config, "base: main\n", "base: stable\nremote: upstream\n", 1))
	git(t, root, "update-ref", "refs/heads/stable", head(t, root))
	git(t, root, "update-ref", "refs/remotes/upstream/stable", approved)
	if got := status(t, root).Unapproved; got != 1 {
		t.Errorf("configured remote/base unapproved = %d, want 1", got)
	}
	git(t, root, "update-ref", "-d", "refs/remotes/upstream/stable")
	if got := status(t, root).Unapproved; got != 0 {
		t.Errorf("configured local fallback unapproved = %d, want 0", got)
	}
	write(t, root, "fireweed.yaml", strings.Replace(config, "base: main\n", "", 1))
	git(t, root, "update-ref", "refs/remotes/origin/stable", approved)
	git(t, root, "symbolic-ref", "refs/remotes/origin/HEAD", "refs/remotes/origin/stable")
	if got := status(t, root).Unapproved; got != 1 {
		t.Errorf("default base via origin/HEAD unapproved = %d, want 1", got)
	}
}

func Test_FLOW_112_StatusPhaseCommitsMatchRealLifecycle(t *testing.T) {
	root := fixture(t, "audit:\n  enabled: false\n", "", "", "", "")
	start(t, root)
	if _, err := done(root); err != nil {
		t.Fatalf("evals Done: %v", err)
	}
	write(t, root, "packages/alpha/internal/ready", "ready\n")
	if _, err := done(root); err != nil {
		t.Fatalf("impl Done: %v", err)
	}
	write(t, root, "author-note.txt", "no phase transition\n")
	commit(t, root, "author note after application")
	commits, err := workflow.Commits(root, "0001-manual.md")
	if err != nil {
		t.Fatal(err)
	}
	if len(commits) != 2 {
		t.Fatalf("real phase commits = %d, want 2", len(commits))
	}
	r := status(t, root)
	if got := r.PhaseCommits["0001-manual"]; got != len(commits) {
		t.Errorf("phase commits = %#v, want %d for 0001-manual", r.PhaseCommits, len(commits))
	}
}

func Test_FLOW_113_StatusLatestRecordedCostsAndSavedWarnings(t *testing.T) {
	root := fixture(t, "audit:\n  enabled: false\nbudgets:\n  eval_ratio: 1\n", "", "", "", "")
	write(t, root, "packages/alpha/evals/weight.txt", "one\ntwo\n")
	commit(t, root, "cost fixture")
	start(t, root)
	if _, err := done(root); err != nil {
		t.Fatal(err)
	}
	write(t, root, "packages/alpha/internal/ready", "ready\n")
	if _, err := verify(root, false); err != nil {
		t.Fatal(err)
	}
	snapshots, err := workflow.Costs(root)
	if err != nil {
		t.Fatal(err)
	}
	latest := map[string]workflow.ModuleCost{}
	var latestAlphaB03, olderAlphaB03 project.Diagnostic
	var hasLatestAlphaB03, hasOlderAlphaB03 bool
	for _, s := range snapshots {
		for _, c := range s.Modules {
			latest[c.Module] = c
			if c.Module == "alpha" {
				olderAlphaB03, hasOlderAlphaB03 = latestAlphaB03, hasLatestAlphaB03
				latestAlphaB03, hasLatestAlphaB03 = warningFor(s.Diagnostics, "B03")
			}
		}
	}
	if !hasOlderAlphaB03 || !hasLatestAlphaB03 || olderAlphaB03 == latestAlphaB03 {
		t.Fatalf("cost snapshots need distinct older and latest alpha B03 warnings: %#v", snapshots)
	}
	config := readFile(t, root, "fireweed.yaml")
	write(t, root, "fireweed.yaml", strings.Replace(config, "eval_ratio: 1", "eval_ratio: 100", 1))
	r := status(t, root)
	if len(r.Costs) != 2 {
		t.Errorf("cost modules = %#v, want only measured alpha and beta", r.Costs)
	}
	for _, c := range r.Costs {
		if !reflect.DeepEqual(c, latest[c.Module]) {
			t.Errorf("latest %s cost = %#v, want %#v", c.Module, c, latest[c.Module])
		}
	}
	for _, module := range []string{"alpha", "beta"} {
		want := latest[module]
		got, ok := costFor(r.Costs, module)
		if !ok || !reflect.DeepEqual(got, want) {
			t.Errorf("recorded %s cost = %#v (present %t), want %#v", module, got, ok, want)
		}
	}
	if _, ok := costFor(r.Costs, "gamma"); ok {
		t.Errorf("unmeasured gamma has cost: %#v", r.Costs)
	}
	if d, ok := warningFor(r.Diagnostics, "B03"); !ok || !d.Warning || d != latestAlphaB03 {
		t.Errorf("selected B03 warning = %#v, want %#v", r.Diagnostics, latestAlphaB03)
	}
	if containsDiagnostic(r.Diagnostics, olderAlphaB03) {
		t.Errorf("older B03 warning returned: %#v", r.Diagnostics)
	}
	count := 0
	for _, d := range r.Diagnostics {
		if d.Code == "B03" {
			count++
		}
	}
	if count != 1 {
		t.Errorf("selected B03 warning count = %d, want 1: %#v", count, r.Diagnostics)
	}
}

func containsDiagnostic(ds []project.Diagnostic, want project.Diagnostic) bool {
	for _, d := range ds {
		if d == want {
			return true
		}
	}
	return false
}
