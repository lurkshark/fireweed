package evals_test

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
)

func queued(t *testing.T, root, filename, front, motivation, added, modified, removed string) {
	t.Helper()
	put(t, root, "changes/"+filename, "---\nmodule: alpha\n"+front+"---\n\n## Motivation\n"+motivation+"\n## Interface\n\n## Added\n"+added+"\n## Modified\n"+modified+"\n## Removed\n"+removed+"\n## Decisions\n\n## Notes\nHuman note.\n")
}

func moduleByName(t *testing.T, snapshot project.Snapshot, name string) project.Module {
	t.Helper()
	for _, module := range snapshot.Modules {
		if module.Name == name {
			return module
		}
	}
	t.Fatalf("module %q absent from projection", name)
	return project.Module{}
}

func specItems(t *testing.T, module project.Module) []contracts.Item {
	t.Helper()
	section, ok := contracts.Section(module.Document.Body, "Specs")
	if !ok {
		t.Fatalf("Specs section absent from %s", module.Name)
	}
	items, err := contracts.ParseItems(section)
	if err != nil {
		t.Fatalf("parse projected %s specs: %v", module.Name, err)
	}
	return items
}

func findItem(t *testing.T, items []contracts.Item, id string) contracts.Item {
	t.Helper()
	for _, item := range items {
		if item.ID == id {
			return item
		}
	}
	t.Fatalf("spec %s absent from %#v", id, items)
	return contracts.Item{}
}

func Test_PRJ_027_PRJ_028_QueueOrderContentAndArchive(t *testing.T) {
	root := fixture(t)
	queued(t, root, "0002-zed.md", "set:\n  interface: [next.go]\n", "A queue item is needed.\n", "### new:feature\nFeature MUST work.\n", "### ALP-001\nAlpha MUST change.\n", "### ALP-002\n")
	queued(t, root, "0002-ace.md", "", "Another queue item is needed.\n", "", "", "")
	queued(t, root, "0001-first.md", "", "The first queue item is needed.\n", "", "", "")
	put(t, root, "changes/archive/0000-old.md", "archived material is not a pending document")
	changes, err := project.Queue(root)
	if err != nil {
		t.Fatalf("Queue: %v", err)
	}
	if len(changes) != 3 {
		t.Fatalf("queue = %#v; want three pending changes", changes)
	}
	for i, want := range []struct {
		number     int
		slug, path string
	}{
		{1, "first", "changes/0001-first.md"},
		{2, "ace", "changes/0002-ace.md"},
		{2, "zed", "changes/0002-zed.md"},
	} {
		got := changes[i]
		if got.Number != want.number || got.Slug != want.slug || got.Path != want.path || got.Module != "alpha" {
			t.Errorf("change %d = %#v; want %#v with module alpha", i, got, want)
		}
	}
	last := changes[2]
	if !strings.Contains(last.Document.Body, "A queue item is needed.") || last.Document.Frontmatter["module"] != "alpha" || last.Document.Frontmatter["set"] == nil {
		t.Errorf("change document/frontmatter lost: %#v", last.Document)
	}
	if len(last.Added) != 1 || last.Added[0].ID != "new:feature" || !strings.Contains(last.Added[0].Body, "Feature MUST work.") || len(last.Modified) != 1 || last.Modified[0].ID != "ALP-001" || !strings.Contains(last.Modified[0].Body, "Alpha MUST change.") || len(last.Removed) != 1 || last.Removed[0].ID != "ALP-002" {
		t.Errorf("parsed edits = added %#v, modified %#v, removed %#v", last.Added, last.Modified, last.Removed)
	}
}

func Test_PRJ_029_PRJ_030_QueueRejectsInvalidDocuments(t *testing.T) {
	for _, tc := range []struct{ id, name, body string }{
		{"PRJ-029", "malformed", "---\nmodule: [broken\n---\n"},
		{"PRJ-030", "unknown_module", "---\nmodule: missing\n---\n## Motivation\nNeeded.\n"},
	} {
		t.Run(tc.id+"_"+tc.name, func(t *testing.T) {
			root := fixture(t)
			put(t, root, "changes/0001-bad.md", tc.body)
			if _, err := project.Queue(root); err == nil {
				t.Error("Queue accepted invalid change")
			}
		})
	}
}

func Test_PRJ_004_PRJ_031_PRJ_032_TargetAliasesAndAmbiguousSlug(t *testing.T) {
	root := fixture(t)
	queued(t, root, "0001-first.md", "", "First.\n", "### new:first\nFirst MUST work.\n", "", "")
	queued(t, root, "0002-second.md", "", "Second.\n", "### new:second\nSecond MUST work.\n", "", "")
	for _, target := range []string{"first", "0001-first.md", "changes/0001-first.md"} {
		projected, err := project.Project(root, target)
		if err != nil {
			t.Errorf("Project(%q): %v", target, err)
			continue
		}
		items := specItems(t, moduleByName(t, projected, "alpha"))
		if len(items) != 2 || findItem(t, items, "new:first").ID == "" {
			t.Errorf("Project(%q) did not stop at first change: %#v", target, items)
		}
		if err := project.Stamp(root, target); err != nil {
			t.Errorf("Stamp(%q): %v", target, err)
		}
	}
	if _, err := project.Project(root, "missing"); err == nil {
		t.Error("Project accepted missing target")
	}
	if err := project.Stamp(root, "missing"); err == nil {
		t.Error("Stamp accepted missing target")
	}
	put(t, root, "changes/archive/0000-archived.md", "archived document")
	if _, err := project.Project(root, "changes/archive/0000-archived.md"); err == nil {
		t.Error("Project accepted archived target")
	}
	queued(t, root, "0003-first.md", "", "Another first.\n", "", "", "")
	if _, err := project.Queue(root); err == nil {
		t.Error("Queue accepted ambiguous slug")
	}
	if _, err := project.Project(root, "first"); err == nil {
		t.Error("Project accepted ambiguous slug")
	}
	if err := project.Stamp(root, "first"); err == nil {
		t.Error("Stamp accepted ambiguous slug")
	}
}

func Test_PRJ_005_PRJ_032_PRJ_036_ProjectEditsAndLocalHandles(t *testing.T) {
	root := fixture(t)
	queued(t, root, "0001-first.md", "", "First.\n", "### new:shared\nFirst change MUST retain its own text.\n", "### ALP-001 [review]\nAlpha MUST have revised full text.\n", "")
	queued(t, root, "0002-second.md", "", "Second.\n", "### new:shared\nSecond change MUST retain different text.\n", "", "### ALP-001\n")
	first, err := project.Project(root, "first")
	if err != nil {
		t.Fatalf("Project first: %v", err)
	}
	firstItems := specItems(t, moduleByName(t, first, "alpha"))
	if len(firstItems) != 2 || findItem(t, firstItems, "ALP-001").Tag != "review" || findItem(t, firstItems, "ALP-001").Body != "Alpha MUST have revised full text.\n\n" || findItem(t, firstItems, "new:shared").Body != "First change MUST retain its own text.\n\n" {
		t.Errorf("first projection = %#v", firstItems)
	}
	all, err := project.Project(root, "")
	if err != nil {
		t.Fatalf("Project whole queue: %v", err)
	}
	section, ok := contracts.Section(moduleByName(t, all, "alpha").Document.Body, "Specs")
	if !ok {
		t.Fatal("Specs section absent from whole-queue projection")
	}
	if strings.Count(section, "### new:shared\n") != 2 ||
		!strings.Contains(section, "### new:shared\nFirst change MUST retain its own text.\n") ||
		!strings.Contains(section, "### new:shared\nSecond change MUST retain different text.\n") ||
		strings.Contains(section, "ALP-001") {
		t.Errorf("whole-queue projection lost local handles or retained removed ID: %q", section)
	}
}

func Test_PRJ_033_PRJ_034_PRJ_035_SetFieldsAndRejections(t *testing.T) {
	root := fixture(t)
	put(t, root, "packages/beta/SPEC.md", "---\nname: beta\nid_prefix: BET\ndepends_on: []\n---\n# Beta\n\n## Specs\n")
	queued(t, root, "0001-settings.md", "set:\n  depends_on: [beta]\n  interface: [next.go]\n  evals: [checks/**]\n  disposable: [generated/**]\n  preserve: [records/**]\n  append_only: [migrations/**]\n  boundaries_opt_out: reviewed boundary\n  hooks:\n    test:\n      command: [go, test, ./...]\n", "Settings are needed.\n", "", "", "")
	got, err := project.Project(root, "")
	if err != nil {
		t.Fatalf("Project settings: %v", err)
	}
	alpha := moduleByName(t, got, "alpha")
	wantLists := map[string][]string{
		"depends_on": {"beta"}, "interface": {"next.go"}, "evals": {"checks/**"},
		"disposable": {"generated/**"}, "preserve": {"records/**"}, "append_only": {"migrations/**"},
	}
	for key, want := range wantLists {
		value, ok := alpha.Document.Frontmatter[key].([]any)
		if !ok || len(value) != 1 || value[0] != want[0] {
			t.Errorf("set %s = %#v; want %q", key, alpha.Document.Frontmatter[key], want)
		}
	}
	if !reflect.DeepEqual(alpha.DependsOn, []string{"beta"}) || !reflect.DeepEqual(alpha.Zones.Interface, []string{"next.go"}) || !reflect.DeepEqual(alpha.Zones.Evals, []string{"checks/**"}) || !reflect.DeepEqual(alpha.Zones.Disposable, []string{"generated/**"}) || !reflect.DeepEqual(alpha.Zones.Preserve, []string{"records/**"}) || !reflect.DeepEqual(alpha.Zones.AppendOnly, []string{"migrations/**"}) {
		t.Errorf("projected module fields not updated: %#v", alpha)
	}
	wantHooks := map[string]any{"test": map[string]any{"command": []any{"go", "test", "./..."}}}
	if alpha.Document.Frontmatter["boundaries_opt_out"] != "reviewed boundary" || !reflect.DeepEqual(alpha.Document.Frontmatter["hooks"], wantHooks) {
		t.Errorf("projected scalar/map settings lost: %#v", alpha.Document.Frontmatter)
	}
	for _, tc := range []struct{ id, name, settings string }{
		{"PRJ-034", "identity_name", "name: renamed"},
		{"PRJ-034", "identity_prefix", "id_prefix: NEW"},
		{"PRJ-034", "unknown_field", "surprise: true"},
		{"PRJ-035", "invalid_graph", "depends_on: [missing]"},
	} {
		t.Run(tc.id+"_"+tc.name, func(t *testing.T) {
			badRoot := fixture(t)
			queued(t, badRoot, "0001-invalid.md", "set:\n  "+tc.settings+"\n", "Needed.\n", "", "", "")
			if _, err := project.Project(badRoot, ""); err == nil {
				t.Errorf("Project accepted %s", tc.name)
			}
		})
	}
}

func Test_PRJ_006_PRJ_038_StampPrecedingProjectionAndPreservation(t *testing.T) {
	root := fixture(t)
	queued(t, root, "0001-revise.md", "", "Earlier change.\n", "", "### ALP-001 [review]\nAlpha MUST reflect the earlier revision.\n", "")
	queued(t, root, "0002-target.md", "custom: author-value\n", "Keep this human motivation.\n", "", "### ALP-001\nAlpha MUST reflect the later revision.\n", "")
	wantHash := contracts.SpecHash(contracts.Item{ID: "ALP-001", Tag: "review", Body: "Alpha MUST reflect the earlier revision.\n"})
	if err := project.Stamp(root, "changes/0002-target.md"); err != nil {
		t.Fatalf("Stamp: %v", err)
	}
	data, err := os.ReadFile(filepath.Join(root, "changes/0002-target.md"))
	if err != nil {
		t.Fatal(err)
	}
	doc, err := contracts.ParseDocument(string(data))
	if err != nil {
		t.Fatalf("stamped document: %v", err)
	}
	bases, ok := doc.Frontmatter["bases"].(map[string]any)
	if !ok || bases["ALP-001"] != wantHash {
		t.Errorf("bases = %#v; want preceding hash %s", doc.Frontmatter["bases"], wantHash)
	}
	if doc.Frontmatter["custom"] != "author-value" || !strings.Contains(doc.Body, "Keep this human motivation.") || !strings.Contains(doc.Body, "Alpha MUST reflect the later revision.") || !strings.Contains(doc.Body, "Human note.") {
		t.Errorf("Stamp lost author content: %#v", doc)
	}
	queued(t, root, "0003-remove.md", "", "Remove the item.\n", "", "", "### ALP-001\n")
	if err := project.Stamp(root, "0003-remove.md"); err != nil {
		t.Fatalf("Stamp Removed: %v", err)
	}
	removedBytes, err := os.ReadFile(filepath.Join(root, "changes/0003-remove.md"))
	if err != nil {
		t.Fatal(err)
	}
	removedDoc, err := contracts.ParseDocument(string(removedBytes))
	if err != nil {
		t.Fatal(err)
	}
	removedWant := contracts.SpecHash(contracts.Item{ID: "ALP-001", Body: "Alpha MUST reflect the later revision.\n"})
	removedBases, ok := removedDoc.Frontmatter["bases"].(map[string]any)
	if !ok || removedBases["ALP-001"] != removedWant {
		t.Errorf("Removed base = %#v; want %s", removedDoc.Frontmatter["bases"], removedWant)
	}
}

func Test_PRJ_037_PRJ_039_StampAddedOnlyAndAbsentID(t *testing.T) {
	root := fixture(t)
	queued(t, root, "0001-added.md", "", "Added item.\n", "### new:thing\nThing MUST work.\n", "", "")
	if err := project.Stamp(root, "added"); err != nil {
		t.Fatalf("Stamp Added-only: %v", err)
	}
	bytes, err := os.ReadFile(filepath.Join(root, "changes/0001-added.md"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(bytes), "### new:thing") || strings.Contains(string(bytes), "### ALP-002") {
		t.Errorf("Stamp assigned an ID to Added-only change: %s", bytes)
	}
	queued(t, root, "0002-invalid.md", "", "Invalid edit.\n", "", "### ALP-999\nMissing item MUST change.\n", "")
	if err := project.Stamp(root, "invalid"); err == nil {
		t.Error("Stamp accepted Modified ID absent from preceding projection")
	}
	removeRoot := fixture(t)
	queued(t, removeRoot, "0001-invalid-remove.md", "", "Invalid removal.\n", "", "", "### ALP-998\n")
	if err := project.Stamp(removeRoot, "invalid-remove"); err == nil {
		t.Error("Stamp accepted Removed ID absent from preceding projection")
	}
}

func Test_PRJ_007_CheckValidAndDiagnostics(t *testing.T) {
	validRoot := fixture(t)
	put(t, validRoot, "packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\n---\n# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST work.\n")
	queued(t, validRoot, "0001-valid.md", "bases:\n  ALP-001: "+contracts.SpecHash(contracts.Item{ID: "ALP-001", Body: "Alpha MUST work.\n"})+"\n", "A user needs the change.\n", "### new:local [review]\nA local addition MUST work.\n", "### ALP-001 [quarantine: flaky fixture]\nAlpha MUST work with revision.\n", "")
	queued(t, validRoot, "0002-also-valid.md", "", "Another user need.\n", "### new:local\nA second local addition MUST work.\n", "", "")
	for _, diagnostic := range project.Check(validRoot) {
		if !diagnostic.Warning {
			t.Errorf("valid spec/change received error diagnostic: %#v", diagnostic)
		}
	}
}

func Test_PRJ_014_PRJ_015_PRJ_016_CheckDiagnostics(t *testing.T) {
	for _, tc := range []struct {
		id, name, spec, front, motivation, added, modified, removed string
		code, cause                                                 string
	}{
		{"PRJ-008", "missing_motivation", "", "", "\n", "", "", "", "C07", "missing Motivation"},
		{"PRJ-009", "nonnormative_module", "### ALP-001\nAlpha works.\n", "", "Needed.\n", "", "", "", "C04", "non-normative spec body"},
		{"PRJ-009", "nonnormative_change", "", "", "Needed.\n", "### new:item\nAn observation.\n", "", "", "C04", "non-normative spec body"},
		{"PRJ-010", "invalid_review_tag", "", "", "Needed.\n", "### new:item [review: reason]\nItem MUST work.\n", "", "", "C06", "invalid review/quarantine tag"},
		{"PRJ-010", "invalid_quarantine_tag", "", "", "Needed.\n", "### new:item [quarantine]\nItem MUST work.\n", "", "", "C06", "invalid review/quarantine tag"},
		{"PRJ-011", "duplicate_module_id", "### ALP-001\nAlpha MUST work.\n\n### ALP-001\nDuplicate MUST fail.\n", "", "Needed.\n", "", "", "", "C05", "duplicate item ID"},
		{"PRJ-012", "foreign_module_id", "### BET-001\nForeign item MUST fail.\n", "", "Needed.\n", "", "", "", "C05", "foreign spec ID"},
		{"PRJ-012", "foreign_change_id", "", "", "Needed.\n", "", "### BET-001\nForeign edit MUST fail.\n", "", "C05", "foreign module ID"},
		{"PRJ-013", "duplicate_edit_identity", "", "", "Needed.\n", "", "### ALP-001\nAlpha MUST change.\n", "### ALP-001\n", "C05", "duplicate edit identity ALP-001"},
		{"PRJ-014", "undeclared_handle", "", "", "Needed.\n", "### new:local\nLocal MUST work.\n", "### ALP-001\nAlpha MUST refer to new:foreign.\n", "", "C09", "reference to undeclared handle new:foreign"},
		{"PRJ-015", "missing_base", "", "", "Needed.\n", "", "### ALP-001\nAlpha MUST change.\n", "", "C08", "missing base for ALP-001"},
		{"PRJ-015", "stale_base", "", "bases:\n  ALP-001: deadbeef\n", "Needed.\n", "", "### ALP-001\nAlpha MUST change.\n", "", "C08", "stale base for ALP-001"},
		{"PRJ-016", "absent_modified", "", "bases:\n  ALP-999: deadbeef\n", "Needed.\n", "", "### ALP-999\nAbsent MUST change.\n", "", "C08", "ID ALP-999 absent from preceding projection"},
		{"PRJ-016", "absent_removed", "", "bases:\n  ALP-999: deadbeef\n", "Needed.\n", "", "", "### ALP-999\n", "C08", "ID ALP-999 absent from preceding projection"},
	} {
		t.Run(tc.id+"_"+tc.name, func(t *testing.T) {
			root := fixture(t)
			front := tc.front
			baseHash := contracts.SpecHash(contracts.Item{ID: "ALP-001", Body: "Alpha MUST work.\n"})
			if tc.name == "duplicate_edit_identity" || tc.name == "undeclared_handle" {
				front = "bases:\n  ALP-001: " + baseHash + "\n"
			}
			if tc.name == "foreign_change_id" {
				put(t, root, "packages/beta/SPEC.md", "---\nname: beta\nid_prefix: BET\ndepends_on: [alpha]\n---\n# Beta\n\n## Specs\n\n### BET-001\nBeta MUST work.\n")
				front = "bases:\n  BET-001: " + contracts.SpecHash(contracts.Item{ID: "BET-001", Body: "Beta MUST work.\n"}) + "\n"
			}
			put(t, root, "packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\n---\n# Alpha\n\n## Specs\n\n"+func() string {
				if tc.spec != "" {
					return tc.spec
				}
				return "### ALP-001\nAlpha MUST work.\n"
			}())
			if tc.name == "undeclared_handle" {
				queued(t, root, "0000-other.md", "", "Earlier local handle.\n", "### new:foreign\nEarlier item MUST work.\n", "", "")
			}
			queued(t, root, "0001-check.md", front, tc.motivation, tc.added, tc.modified, tc.removed)
			diagnostics := project.Check(root)
			found := false
			foundCause := false
			foundApplyCause := false
			for _, diagnostic := range diagnostics {
				if !diagnostic.Warning && (strings.TrimSpace(diagnostic.Code) == "" || strings.TrimSpace(diagnostic.Message) == "") {
					t.Errorf("Check returned error diagnostic without code and message: %#v", diagnostic)
				}
				if !diagnostic.Warning && strings.TrimSpace(diagnostic.Code) != "" && strings.TrimSpace(diagnostic.Message) != "" {
					found = true
				}
				matchesCause := strings.Contains(diagnostic.Message, tc.cause)
				if tc.id == "PRJ-016" {
					matchesCause = diagnostic.Message == tc.cause
				}
				if tc.code != "" && !diagnostic.Warning && diagnostic.Code == tc.code && matchesCause {
					foundCause = true
				}
				if tc.name == "absent_modified" && !diagnostic.Warning && diagnostic.Code == "C08" && strings.Contains(diagnostic.Message, "Modified ID ALP-999 absent from preceding projection") {
					foundApplyCause = true
				}
				if tc.name == "absent_removed" && !diagnostic.Warning && diagnostic.Code == "C08" && strings.Contains(diagnostic.Message, "Removed ID ALP-999 absent from preceding projection") {
					foundApplyCause = true
				}
			}
			if !found {
				t.Errorf("Check did not report error diagnostic with nonempty code and message: %#v", diagnostics)
			}
			if tc.code != "" && !foundCause {
				t.Errorf("Check did not report %s for %s: %#v", tc.code, tc.cause, diagnostics)
			}
			if tc.id == "PRJ-016" && !foundApplyCause {
				t.Errorf("Check did not report C08 for absent %s during projection application: %#v", tc.name, diagnostics)
			}
			if tc.name == "missing_base" || tc.name == "stale_base" {
				queued(t, root, "0001-check.md", front, "Needed.\n", "", "", "### ALP-001\n")
				removedDiagnostics := project.Check(root)
				found = false
				for _, diagnostic := range removedDiagnostics {
					if !diagnostic.Warning && (strings.TrimSpace(diagnostic.Code) == "" || strings.TrimSpace(diagnostic.Message) == "") {
						t.Errorf("Check returned error diagnostic without code and message for Removed entry: %#v", diagnostic)
					}
					if !diagnostic.Warning && diagnostic.Code == "C08" && strings.Contains(diagnostic.Message, tc.cause) {
						found = true
					}
				}
				if !found {
					t.Errorf("Check did not report C08 for Removed entry with %s (%s): %#v", tc.name, tc.cause, removedDiagnostics)
				}
			}
			if tc.name == "nonnormative_change" || tc.name == "invalid_review_tag" {
				front, added, modified, want := "", "", "", ""
				code := "C04"
				if tc.name == "nonnormative_change" {
					front = "bases:\n  ALP-001: " + baseHash + "\n"
					modified = "### ALP-001\nAn observation.\n"
					want = "non-normative spec body"
				} else {
					code = "C06"
					added = "### new:item [other: reason]\nItem MUST work.\n"
					want = "invalid review/quarantine tag"
				}
				queued(t, root, "0001-check.md", front, "Needed.\n", added, modified, "")
				found = false
				for _, diagnostic := range project.Check(root) {
					if !diagnostic.Warning && (strings.TrimSpace(diagnostic.Code) == "" || strings.TrimSpace(diagnostic.Message) == "") {
						t.Errorf("Check returned error diagnostic without code and message: %#v", diagnostic)
					}
					if !diagnostic.Warning && diagnostic.Code == code && strings.Contains(diagnostic.Message, want) {
						found = true
					}
				}
				if !found {
					t.Errorf("Check did not reject %s: wanted error diagnostic %s with message containing %q", tc.name, code, want)
				}
			}
		})
	}
}
