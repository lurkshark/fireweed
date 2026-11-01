package evals_test

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
)

func staleCLIText(t *testing.T, root, path string) []byte {
	t.Helper()
	data, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(path)))
	if err != nil {
		t.Fatal(err)
	}
	return data
}

func staleCLIProject(t *testing.T) string {
	t.Helper()
	root := authoringFixture(t)
	appPut(t, root, "packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\n---\n# Alpha\n\n## Specs\n\n### ALP-001\nOne MUST work.\n\n### ALP-002\nTwo MUST work.\n")
	return root
}

func staleCLIChange(t *testing.T, root, path, bases, motivation, modified, removed string) {
	t.Helper()
	appPut(t, root, path, "---\nmodule: alpha\n"+bases+"---\n\n## Motivation\n"+motivation+"\n## Interface\n\n## Added\n\n## Modified\n"+modified+"\n## Removed\n"+removed+"\n## Decisions\n")
}

func staleCLISuccess(t *testing.T, bin, dir, selector string) string {
	t.Helper()
	status, plain, stderr := runAppBinary(t, bin, dir, "show", "--stale", selector)
	if status != 0 || stderr != "" {
		t.Fatalf("show --stale %q: exit=%d stdout=%q stderr=%q", selector, status, plain, stderr)
	}
	status, output, stderr := runAppBinary(t, bin, dir, "show", "--stale", selector, "--json")
	if status != 0 || stderr != "" || commandStringData(t, output) != plain {
		t.Errorf("JSON stale view %q: exit=%d stdout=%q stderr=%q; plain=%q", selector, status, output, stderr, plain)
	}
	return plain
}

func Test_APP_143_APP_144_StaleRecoveryDiffAndReadOnly(t *testing.T) {
	bin := buildAppBinary(t)
	root := staleCLIProject(t)
	cliGit(t, root, "init", "-q", "-b", "main")
	staleCLIChange(t, root, "changes/0001-earlier.md", "", "Earlier work.\n", "### ALP-001 [review]\nOne MUST use old text.  \n\n### ALP-002 [review]\nTwo MUST use old text.\n", "")
	staleCLIChange(t, root, "changes/0002-target.md", "", "Target work.\n", "### ALP-001\nTarget MUST propose different text.\n", "### ALP-002\n")
	cliCommit(t, root, "unstamped queue")
	if err := project.Stamp(root, "target"); err != nil {
		t.Fatal(err)
	}
	path := "changes/0002-target.md"
	stamped := string(staleCLIText(t, root, path))
	doc, err := contracts.ParseDocument(stamped)
	if err != nil {
		t.Fatal(err)
	}
	bases := doc.Frontmatter["bases"].(map[string]any)
	if bases["ALP-001"] == "" || bases["ALP-002"] == "" {
		t.Fatal("fixture lacks both stamped bases")
	}
	bases["ALP-002"] = "deadbeef"
	changed, err := contracts.RenderDocument(doc)
	if err != nil {
		t.Fatal(err)
	}
	appPut(t, root, path, changed)
	cliCommit(t, root, "last bases change")
	staleCLIChange(t, root, "changes/0001-earlier.md", "", "Earlier work.\n", "### ALP-001 [quarantine: changed]\nOne MUST use current text.\n\n### ALP-002\nTwo MUST use current text.\n", "")
	appPut(t, root, path, strings.Replace(string(staleCLIText(t, root, path)), "Target work.", "Later unrelated prose.", 1))
	cliCommit(t, root, "later edit without bases change")
	entries, err := project.StaleSpecs(root, "target")
	if err != nil || len(entries) != 2 || !entries[0].Recovered || entries[1].Recovered {
		t.Fatalf("fixture stale entries=%+v err=%v", entries, err)
	}
	beforeFiles := map[string][]byte{}
	for _, file := range []string{"fireweed.yaml", "packages/alpha/SPEC.md", "changes/0001-earlier.md", path, ".git/index"} {
		beforeFiles[file] = staleCLIText(t, root, file)
	}
	head, refs := cliGit(t, root, "rev-parse", "HEAD"), cliGit(t, root, "for-each-ref", "--format=%(refname) %(objectname)")
	dir := appNestedDir(t, root)
	plain := staleCLISuccess(t, bin, dir, "target")
	for _, alias := range []string{"0002-target.md", path} {
		if got := staleCLISuccess(t, bin, dir, alias); got != plain {
			t.Errorf("alias %q differs from slug output", alias)
		}
	}
	first, second := strings.Index(plain, "ALP-001"), strings.Index(plain, "ALP-002")
	if first < 0 || second <= first {
		t.Fatalf("stale entries are missing or out of order: %q", plain)
	}
	recovered, fallback := plain[first:second], plain[second:]
	for _, marker := range []string{"-### ALP-001 [review]", "-One MUST use old text.  ", "+### ALP-001 [quarantine: changed]", "+One MUST use current text."} {
		if !strings.Contains(recovered, marker) {
			t.Errorf("recovered diff lacks %q: %q", marker, recovered)
		}
	}
	if strings.Contains(recovered, "Target MUST propose different text") || strings.Contains(recovered, "deadbeef") {
		t.Errorf("recovered diff used proposed edit or another item's hash: %q", recovered)
	}
	for _, marker := range []string{"deadbeef", "Two MUST use current text.", "unavailable"} {
		if !strings.Contains(strings.ToLower(fallback), strings.ToLower(marker)) {
			t.Errorf("unrecovered entry lacks %q: %q", marker, fallback)
		}
	}
	if strings.Contains(fallback, "Two MUST use old text") || strings.Contains(fallback, "-Two MUST") || strings.Contains(fallback, "+Two MUST") {
		t.Errorf("unverified history appeared as a diff: %q", fallback)
	}
	for file, before := range beforeFiles {
		if !bytes.Equal(before, staleCLIText(t, root, file)) {
			t.Errorf("show changed %s", file)
		}
	}
	if cliGit(t, root, "rev-parse", "HEAD") != head || cliGit(t, root, "for-each-ref", "--format=%(refname) %(objectname)") != refs {
		t.Error("show changed HEAD or refs")
	}
}

func Test_APP_144_StaleUnavailableHistoryAndEmptyView(t *testing.T) {
	bin := buildAppBinary(t)
	root := staleCLIProject(t)
	staleCLIChange(t, root, "changes/0001-target.md", "bases:\n  ALP-001: deadbeef\n", "Needed.\n", "### ALP-001\nProposed MUST differ.\n", "")
	plain := staleCLISuccess(t, bin, appNestedDir(t, root), "target")
	for _, marker := range []string{"ALP-001", "deadbeef", "One MUST work.", "unavailable"} {
		if !strings.Contains(strings.ToLower(plain), strings.ToLower(marker)) {
			t.Errorf("unavailable-history output lacks %q: %q", marker, plain)
		}
	}
	if strings.Contains(plain, "-One MUST") || strings.Contains(plain, "+One MUST") || strings.Contains(plain, "Proposed MUST differ") {
		t.Errorf("unavailable history invented a recovered diff: %q", plain)
	}
	// A separate valid missing-base change has no stale entries.
	staleCLIChange(t, root, "changes/0001-target.md", "", "Needed.\n", "### ALP-001\nProposed MUST differ.\n", "")
	staleCLISuccess(t, bin, appNestedDir(t, root), "target")
}

func Test_APP_143_StaleViewChoosesNearestProject(t *testing.T) {
	bin := buildAppBinary(t)
	outer := staleCLIProject(t)
	staleCLIChange(t, outer, "changes/0001-target.md", "bases:\n  ALP-001: outerhash\n", "Outer.\n", "### ALP-001\nOuter MUST differ.\n", "")
	inner := filepath.Join(outer, "nested-project")
	appPut(t, inner, "fireweed.yaml", appFixtureConfig)
	appPut(t, inner, "packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\n---\n# Inner\n\n## Specs\n\n### ALP-001\nInner MUST work.\n")
	staleCLIChange(t, inner, "changes/0001-target.md", "bases:\n  ALP-001: innerhash\n", "Inner.\n", "### ALP-001\nInner MUST differ.\n", "")
	plain := staleCLISuccess(t, bin, appNestedDir(t, inner), "target")
	if !strings.Contains(plain, "innerhash") || !strings.Contains(plain, "Inner MUST work.") || strings.Contains(plain, "outerhash") || strings.Contains(plain, "One MUST work.") {
		t.Errorf("stale view used wrong project: %q", plain)
	}
}

func Test_APP_145_APP_146_StaleUsageLookupAndQueryFailures(t *testing.T) {
	bin := buildAppBinary(t)
	root := staleCLIProject(t)
	staleCLIChange(t, root, "changes/0001-target.md", "bases:\n  ALP-001: deadbeef\n", "Needed.\n", "### ALP-001\nProposed MUST differ.\n", "")
	dir := appNestedDir(t, root)
	for _, args := range [][]string{
		{"show", "--stale"}, {"show", "--stale", "target", "extra"},
		{"show", "--stale", "target", "--history"}, {"show", "--stale", "target", "--projected"},
		{"modules", "--stale"}, {"check", "--stale", "--json"},
	} {
		for _, jsonMode := range []bool{false, true} {
			call := append([]string(nil), args...)
			if jsonMode && !strings.Contains(strings.Join(call, " "), "--json") {
				call = append(call, "--json")
			}
			status, out, stderr := runAppBinary(t, bin, dir, call...)
			assertAppHistoryFailure(t, status, 2, out, stderr, strings.Contains(strings.Join(call, " "), "--json"), "U01", "show --stale")
		}
	}
	for _, tc := range []struct {
		name, selector, code string
		exit                 int
	}{
		{"unknown", "absent", "U02", 2},
		{"ambiguous", "target", "C13", 2},
		{"bad-bases", "changes/0001-target.md", "C08", 1},
	} {
		t.Run(tc.name, func(t *testing.T) {
			fixture := staleCLIProject(t)
			staleCLIChange(t, fixture, "changes/0001-target.md", "bases:\n  ALP-001: deadbeef\n", "Needed.\n", "### ALP-001\nProposed MUST differ.\n", "")
			if tc.name == "ambiguous" {
				staleCLIChange(t, fixture, "changes/0002-target.md", "", "Another.\n", "", "")
			}
			if tc.name == "bad-bases" {
				staleCLIChange(t, fixture, "changes/0001-target.md", "bases: [invalid]\n", "Needed.\n", "### ALP-001\nProposed MUST differ.\n", "")
			}
			for _, jsonMode := range []bool{false, true} {
				args := []string{"show", "--stale", tc.selector}
				if jsonMode {
					args = append(args, "--json")
				}
				status, out, stderr := runAppBinary(t, bin, appNestedDir(t, fixture), args...)
				assertAppHistoryFailure(t, status, tc.exit, out, stderr, jsonMode, tc.code, tc.selector)
			}
		})
	}
	for _, bad := range []struct {
		name, code string
		edit       func(*testing.T, string)
	}{
		{"missing-config", "C01", func(t *testing.T, root string) {
			if err := os.Remove(filepath.Join(root, "fireweed.yaml")); err != nil {
				t.Fatal(err)
			}
		}},
		{"malformed-config", "C01", func(t *testing.T, root string) { appPut(t, root, "fireweed.yaml", "modules: [broken\n") }},
		{"invalid-graph", "C02", func(t *testing.T, root string) {
			appPut(t, root, "packages/beta/SPEC.md", "---\nname: beta\nid_prefix: BET\ndepends_on: [missing]\n---\n# Beta\n\n## Specs\n")
		}},
	} {
		t.Run(bad.name, func(t *testing.T) {
			fixture := staleCLIProject(t)
			bad.edit(t, fixture)
			status, out, stderr := runAppBinary(t, bin, appNestedDir(t, fixture), "show", "--stale", "target", "--json")
			assertAppHistoryFailure(t, status, 2, out, stderr, true, bad.code, "")
		})
	}
}
