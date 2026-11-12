package evals_test

import (
	"bytes"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
)

func staleSpec(t *testing.T, root, through, id string) contracts.Item {
	t.Helper()
	snapshot, err := project.Project(root, through)
	if err != nil {
		t.Fatal(err)
	}
	return findItem(t, specItems(t, moduleByName(t, snapshot, "alpha")), id)
}

func staleBaseItem(t *testing.T, root, id string) contracts.Item {
	t.Helper()
	snapshot, err := project.Load(root)
	if err != nil {
		t.Fatal(err)
	}
	return findItem(t, specItems(t, moduleByName(t, snapshot, "alpha")), id)
}

func staleRead(t *testing.T, root, path string) []byte {
	t.Helper()
	data, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(path)))
	if err != nil {
		t.Fatal(err)
	}
	return data
}

func staleBases(t *testing.T, root, path string) (map[string]any, contracts.Document) {
	t.Helper()
	doc, err := contracts.ParseDocument(string(staleRead(t, root, path)))
	if err != nil {
		t.Fatal(err)
	}
	bases, ok := doc.Frontmatter["bases"].(map[string]any)
	if !ok {
		t.Fatalf("bases = %#v, want mapping", doc.Frontmatter["bases"])
	}
	return bases, doc
}

func staleFixture(t *testing.T) string {
	t.Helper()
	root := fixture(t)
	put(t, root, "packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\n---\n# Alpha\n\n## Specs\n\n### ALP-001\nOne MUST work.\n\n### ALP-002\nTwo MUST work.\n\n### ALP-003\nThree MUST work.\n\n### ALP-004\nFour MUST work.\n")
	return root
}

func Test_PRJ_060_StaleSpecsSelectsRecordedModifiedAndRemovedInEditOrder(t *testing.T) {
	root := staleFixture(t)
	oldOne := contracts.SpecHash(staleSpec(t, root, "", "ALP-001"))
	oldTwo := contracts.SpecHash(staleSpec(t, root, "", "ALP-002"))
	oldThree := contracts.SpecHash(staleSpec(t, root, "", "ALP-003"))
	queued(t, root, "0001-earlier.md", "", "Earlier edits.\n", "", "### ALP-001 [review]\nOne MUST change.\n\n### ALP-002 [review]\nTwo MUST change.\n\n### ALP-003 [review]\nThree MUST change.\n", "")
	queued(t, root, "0002-target.md", "bases:\n  ALP-002: "+oldTwo+"\n  ALP-001: "+oldOne+"\n  ALP-003: "+oldThree+"\n", "Later edits.\n", "", "### ALP-002\nTwo MUST change again.\n\n### ALP-001\nOne MUST change again.\n\n### ALP-004\nFour MUST change without a stamp.\n", "### ALP-003\n")
	got, err := project.StaleSpecs(root, "target")
	if err != nil {
		t.Fatal(err)
	}
	want := []struct{ id, hash, current string }{
		{"ALP-002", oldTwo, contracts.RenderItem(staleSpec(t, root, "earlier", "ALP-002"))},
		{"ALP-001", oldOne, contracts.RenderItem(staleSpec(t, root, "earlier", "ALP-001"))},
		{"ALP-003", oldThree, contracts.RenderItem(staleSpec(t, root, "earlier", "ALP-003"))},
	}
	if len(got) != len(want) {
		t.Fatalf("stale entries = %#v; want %d recorded entries", got, len(want))
	}
	for i, entry := range got {
		if entry.ID != want[i].id || entry.RecordedHash != want[i].hash || entry.Current != want[i].current {
			t.Errorf("entry %d = %#v; want ID %s, hash %s, Current %q", i, entry, want[i].id, want[i].hash, want[i].current)
		}
	}
}

func Test_PRJ_061_PRJ_062_HistoryUsesLastBasesChangeAndValidatesEachHash(t *testing.T) {
	root := staleFixture(t)
	revisionGit(t, root, "init", "-q")
	queued(t, root, "0001-earlier.md", "", "Earlier.\n", "", "### ALP-001 [review]\nOne MUST use old text.  \n\n### ALP-002 [review]\nTwo MUST use old text.\n", "")
	queued(t, root, "0002-target.md", "", "Target.\n", "", "### ALP-001\nOne MUST be edited.\n", "### ALP-002\n")
	revisionCommit(t, root, "initial unstamped queue")
	before := staleSpec(t, root, "earlier", "ALP-001")
	if before.Tag != "review" || !strings.Contains(before.Body, "old text.  ") {
		t.Fatalf("fixture lost tag or trailing spaces: %#v", before)
	}
	if err := project.Stamp(root, "target"); err != nil {
		t.Fatal(err)
	}
	path := "changes/0002-target.md"
	bases, doc := staleBases(t, root, path)
	if bases["ALP-001"] != contracts.SpecHash(before) {
		t.Fatalf("stamped hash %v differs from normalized tagged hash %s", bases["ALP-001"], contracts.SpecHash(before))
	}
	// A second recorded hash is intentionally wrong. Recovery must be per item.
	bases["ALP-002"] = "deadbeef"
	text, err := contracts.RenderDocument(doc)
	if err != nil {
		t.Fatal(err)
	}
	put(t, root, path, text)
	revisionCommit(t, root, "last bases mapping change")
	queued(t, root, "0001-earlier.md", "", "Earlier.\n", "", "### ALP-001 [quarantine: changed]\nOne MUST use current text.\n\n### ALP-002\nTwo MUST use current text.\n", "")
	// This later target-file commit changes only prose, not the bases mapping.
	put(t, root, path, strings.Replace(string(staleRead(t, root, path)), "Target.", "Target with later prose.", 1))
	revisionCommit(t, root, "later unrelated change to target file")
	currentOne := staleSpec(t, root, "earlier", "ALP-001")
	currentTwo := staleSpec(t, root, "earlier", "ALP-002")
	got, err := project.StaleSpecs(root, "target")
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 2 {
		t.Fatalf("stale entries = %#v; want two", got)
	}
	if got[0].ID != "ALP-001" || got[0].RecordedHash != contracts.SpecHash(before) || !got[0].Recovered || got[0].Before != contracts.RenderItem(before) || got[0].Current != contracts.RenderItem(currentOne) {
		t.Errorf("recovered tagged item = %#v; want historical pre-target item %#v and current %#v", got[0], before, currentOne)
	}
	if got[1].ID != "ALP-002" || got[1].RecordedHash != "deadbeef" || got[1].Recovered || got[1].Before != "" || got[1].Current != contracts.RenderItem(currentTwo) {
		t.Errorf("mismatched historical hash = %#v; want current text without Before", got[1])
	}
}

func Test_PRJ_062_UnavailableHistoryStillReportsEveryStaleItem(t *testing.T) {
	root := staleFixture(t) // No Git repository or recorded target history.
	queued(t, root, "0001-target.md", "bases:\n  ALP-001: deadbeef\n  ALP-002: 12345678\n", "Needed.\n", "", "### ALP-001\nOne MUST change.\n", "### ALP-002\n")
	got, err := project.StaleSpecs(root, "target")
	if err != nil || len(got) != 2 {
		t.Fatalf("unavailable history = %#v, %v; want both stale items", got, err)
	}
	for i, want := range []struct{ id, hash string }{{"ALP-001", "deadbeef"}, {"ALP-002", "12345678"}} {
		if got[i].ID != want.id || got[i].RecordedHash != want.hash || got[i].Recovered || got[i].Before != "" || got[i].Current != contracts.RenderItem(staleBaseItem(t, root, want.id)) {
			t.Errorf("entry %d = %#v; want unrecovered %s", i, got[i], want.id)
		}
	}
}

func Test_PRJ_063_StaleSpecsEmptyAndReadOnly(t *testing.T) {
	root := staleFixture(t)
	revisionGit(t, root, "init", "-q")
	queued(t, root, "0001-target.md", "", "Needed.\n", "", "### ALP-001\nOne MUST change.\n", "")
	if err := project.Stamp(root, "target"); err != nil {
		t.Fatal(err)
	}
	revisionCommit(t, root, "stamped target")
	path := "changes/0001-target.md"
	readState := func() (target, index []byte, head, refs string) {
		target = staleRead(t, root, path)
		index = staleRead(t, root, ".git/index")
		head = revisionGit(t, root, "rev-parse", "HEAD")
		refs = revisionGit(t, root, "show-ref")
		return
	}
	target, index, head, refs := readState()
	got, err := project.StaleSpecs(root, "target")
	if err != nil || len(got) != 0 {
		t.Fatalf("matching base = %#v, %v; want empty slice", got, err)
	}
	put(t, root, "packages/alpha/SPEC.md", strings.Replace(string(staleRead(t, root, "packages/alpha/SPEC.md")), "One MUST work.", "One MUST now work differently.", 1))
	spec := staleRead(t, root, "packages/alpha/SPEC.md")
	got, err = project.StaleSpecs(root, "target")
	if err != nil || len(got) != 1 {
		t.Fatalf("stale base = %#v, %v; want one", got, err)
	}
	afterTarget, afterIndex, afterHead, afterRefs := readState()
	if !bytes.Equal(target, afterTarget) || !bytes.Equal(spec, staleRead(t, root, "packages/alpha/SPEC.md")) || !bytes.Equal(index, afterIndex) || head != afterHead || refs != afterRefs {
		t.Error("StaleSpecs changed a worktree file, index or ref")
	}
}

func Test_PRJ_064_StampMissingPreservesExistingBasesAndHumanContent(t *testing.T) {
	root := staleFixture(t)
	queued(t, root, "0001-target.md", "custom: keep-me\nbases:\n  ALP-001: deadbeef\n  ALP-002: "+contracts.SpecHash(staleBaseItem(t, root, "ALP-002"))+"\n  ALP-999: extra\n", "Keep this motivation.\n", "", "### ALP-001\nOne MUST change.\n\n### ALP-002\nTwo MUST change.\n\n### ALP-003\nThree MUST change.\n", "### ALP-004\n")
	path := "changes/0001-target.md"
	_, original := staleBases(t, root, path)
	beforeCheck := project.Check(root)
	if err := project.StampMissing(root, "target"); err != nil {
		t.Fatal(err)
	}
	bases, after := staleBases(t, root, path)
	want := map[string]any{
		"ALP-001": "deadbeef", "ALP-002": contracts.SpecHash(staleBaseItem(t, root, "ALP-002")),
		"ALP-003": contracts.SpecHash(staleBaseItem(t, root, "ALP-003")),
		"ALP-004": contracts.SpecHash(staleBaseItem(t, root, "ALP-004")), "ALP-999": "extra",
	}
	if !reflect.DeepEqual(bases, want) || after.Frontmatter["custom"] != "keep-me" || after.Body != original.Body {
		t.Errorf("StampMissing lost bases/frontmatter/human sections: bases=%#v, body changed=%v", bases, after.Body != original.Body)
	}
	if len(beforeCheck) == 0 {
		t.Fatal("fixture lacked Check diagnostics before stamping")
	}
	staleFound := false
	for _, diagnostic := range project.Check(root) {
		if diagnostic.Code == "C08" && strings.Contains(diagnostic.Message, "stale base for ALP-001") {
			staleFound = true
		}
	}
	if !staleFound {
		t.Error("ordinary Check stopped reporting the preserved stale base")
	}
	if err := project.Stamp(root, "target"); err != nil {
		t.Fatal(err)
	}
	stamped, _ := staleBases(t, root, path)
	if stamped["ALP-001"] != contracts.SpecHash(staleBaseItem(t, root, "ALP-001")) {
		t.Errorf("ordinary Stamp stopped replacing a stale base: %#v", stamped)
	}
}

func Test_PRJ_065_StampMissingRejectsInvalidBasesAndAbsentEditsAtomically(t *testing.T) {
	for _, tc := range []struct{ name, front, modified, removed string }{
		{"nonmapping", "bases: [ALP-001]\n", "### ALP-001\nOne MUST change.\n", ""},
		{"absent_modified_after_valid", "", "### ALP-001\nOne MUST change.\n\n### ALP-999\nMissing MUST change.\n", ""},
		{"absent_removed_after_valid", "", "### ALP-001\nOne MUST change.\n", "### ALP-998\n"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := staleFixture(t)
			path := "changes/0001-target.md"
			queued(t, root, "0001-target.md", tc.front, "Needed.\n", "", tc.modified, tc.removed)
			before := staleRead(t, root, path)
			err := project.StampMissing(root, "target")
			if code, _, ok := contracts.ErrorCode(err); !ok || code != "C08" {
				t.Errorf("StampMissing error = %v; want C08", err)
			}
			if !bytes.Equal(before, staleRead(t, root, path)) {
				t.Error("StampMissing rewrote target after rejection")
			}
		})
	}
}

func Test_PRJ_066_StaleQueriesResolveAliasesAndCodedLookupFailures(t *testing.T) {
	for _, alias := range []string{"target", "0001-target.md", "changes/0001-target.md"} {
		t.Run(alias, func(t *testing.T) {
			root := staleFixture(t)
			queued(t, root, "0001-target.md", "bases:\n  ALP-001: deadbeef\n", "Needed.\n", "", "### ALP-001\nOne MUST change.\n", "### ALP-002\n")
			got, err := project.StaleSpecs(root, alias)
			if err != nil || len(got) != 1 || got[0].ID != "ALP-001" {
				t.Errorf("StaleSpecs(%q) = %#v, %v", alias, got, err)
			}
			if err := project.StampMissing(root, alias); err != nil {
				t.Errorf("StampMissing(%q): %v", alias, err)
			}
			bases, _ := staleBases(t, root, "changes/0001-target.md")
			if bases["ALP-002"] != contracts.SpecHash(staleBaseItem(t, root, "ALP-002")) || bases["ALP-001"] != "deadbeef" {
				t.Errorf("StampMissing(%q) bases = %#v", alias, bases)
			}
		})
	}
	root := staleFixture(t)
	for _, operation := range []struct {
		name string
		call func(string) error
	}{
		{"StaleSpecs", func(target string) error { _, err := project.StaleSpecs(root, target); return err }},
		{"StampMissing", func(target string) error { return project.StampMissing(root, target) }},
	} {
		t.Run(operation.name+"_unknown", func(t *testing.T) {
			const target = "changes/0099-absent.md"
			err := operation.call(target)
			if code, subject, ok := contracts.ErrorCode(err); !ok || code != "U02" || subject != target {
				t.Errorf("unknown selector = %v; want U02 with subject %q", err, target)
			}
		})
	}
	queued(t, root, "0001-target.md", "", "First.\n", "", "", "")
	queued(t, root, "0002-target.md", "", "Second.\n", "", "", "")
	for _, operation := range []struct {
		name string
		call func(string) error
	}{
		{"StaleSpecs", func(target string) error { _, err := project.StaleSpecs(root, target); return err }},
		{"StampMissing", func(target string) error { return project.StampMissing(root, target) }},
	} {
		t.Run(operation.name+"_ambiguous", func(t *testing.T) {
			err := operation.call("target")
			if code, _, ok := contracts.ErrorCode(err); !ok || code != "C13" {
				t.Errorf("ambiguous slug = %v; want C13", err)
			}
		})
	}
}
