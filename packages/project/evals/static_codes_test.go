package evals_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/project"
)

func Test_PRJ_040_CheckConfigurationCauses(t *testing.T) {
	for _, tc := range []struct {
		name string
		edit func(*testing.T, string)
	}{
		{"missing", func(t *testing.T, root string) {
			if err := os.Remove(filepath.Join(root, "fireweed.yaml")); err != nil {
				t.Fatal(err)
			}
		}},
		{"malformed", func(t *testing.T, root string) {
			put(t, root, "fireweed.yaml", "modules: [broken\n")
		}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := fixture(t)
			tc.edit(t, root)
			_, cause := project.Load(root)
			if cause == nil {
				t.Fatal("Load accepted invalid configuration")
			}
			for _, diagnostic := range project.Check(root) {
				if !diagnostic.Warning && diagnostic.Code == "C01" && strings.Contains(diagnostic.Message, cause.Error()) {
					return
				}
			}
			t.Errorf("Check did not report C01 for configuration cause %q", cause)
		})
	}
}

func Test_PRJ_041_CheckGraphCauses(t *testing.T) {
	t.Run("current", func(t *testing.T) {
		root := fixture(t)
		put(t, root, "packages/beta/SPEC.md", "---\nname: beta\nid_prefix: BET\ndepends_on: [missing]\n---\n# Beta\n\n## Specs\n")
		snapshot, err := project.Load(root)
		if err != nil {
			t.Fatal(err)
		}
		cause := project.ValidateGraph(snapshot.Modules)
		if cause == nil {
			t.Fatal("ValidateGraph accepted missing dependency")
		}
		for _, diagnostic := range project.Check(root) {
			if !diagnostic.Warning && diagnostic.Code == "C02" && strings.Contains(diagnostic.Message, cause.Error()) {
				return
			}
		}
		t.Errorf("Check did not report C02 for current graph cause %q", cause)
	})
	t.Run("projected", func(t *testing.T) {
		root := fixture(t)
		queued(t, root, "0001-invalid.md", "set:\n  depends_on: [missing]\n", "Needed.\n", "", "", "")
		_, cause := project.Project(root, "")
		if cause == nil {
			t.Fatal("Project accepted invalid projected graph")
		}
		for _, diagnostic := range project.Check(root) {
			if !diagnostic.Warning && diagnostic.Code == "C02" && strings.Contains(diagnostic.Message, cause.Error()) {
				return
			}
		}
		t.Errorf("Check did not report C02 for projected graph cause %q", cause)
	})
}

func Test_PRJ_042_CheckDocumentAndItemCauses(t *testing.T) {
	for _, tc := range []struct {
		name, source string
		edit         func(*testing.T, string)
	}{
		{"missing_module", "load", func(t *testing.T, root string) {
			if err := os.Remove(filepath.Join(root, "packages/alpha/SPEC.md")); err != nil {
				t.Fatal(err)
			}
		}},
		{"malformed_module", "load", func(t *testing.T, root string) {
			put(t, root, "packages/alpha/SPEC.md", "---\nname: [broken\n---\n# Alpha\n")
		}},
		{"malformed_change", "queue", func(t *testing.T, root string) {
			put(t, root, "changes/0001-bad.md", "---\nmodule: [broken\n---\n")
		}},
		{"malformed_item", "queue", func(t *testing.T, root string) {
			queued(t, root, "0001-bad-item.md", "", "Needed.\n", "### new:item [broken\nItem MUST work.\n", "", "")
		}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := fixture(t)
			tc.edit(t, root)
			var cause error
			if tc.source == "load" {
				_, cause = project.Load(root)
			} else {
				_, cause = project.Queue(root)
			}
			if cause == nil {
				t.Fatal("public operation accepted invalid document or item")
			}
			for _, diagnostic := range project.Check(root) {
				if !diagnostic.Warning && diagnostic.Code == "C03" && strings.Contains(diagnostic.Message, cause.Error()) {
					return
				}
			}
			t.Errorf("Check did not report C03 for %s cause %q", tc.source, cause)
		})
	}
}

func Test_PRJ_043_CheckMetadataAndSetCauses(t *testing.T) {
	for _, tc := range []struct {
		name, source string
		edit         func(*testing.T, string)
	}{
		{"metadata_type", "load", func(t *testing.T, root string) {
			put(t, root, "packages/alpha/SPEC.md", "---\nname: [alpha]\nid_prefix: ALP\ndepends_on: []\n---\n# Alpha\n\n## Specs\n")
		}},
		{"zone_type", "load", func(t *testing.T, root string) {
			put(t, root, "packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\nevals: 7\n---\n# Alpha\n\n## Specs\n")
		}},
		{"identity_set_key", "project", func(t *testing.T, root string) {
			queued(t, root, "0001-invalid.md", "set:\n  name: renamed\n", "Needed.\n", "", "", "")
		}},
		{"unknown_set_key", "project", func(t *testing.T, root string) {
			queued(t, root, "0001-invalid.md", "set:\n  surprise: true\n", "Needed.\n", "", "", "")
		}},
		{"set_value_type", "project", func(t *testing.T, root string) {
			queued(t, root, "0001-invalid.md", "set:\n  interface: 7\n", "Needed.\n", "", "", "")
		}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := fixture(t)
			tc.edit(t, root)
			var cause error
			if tc.source == "load" {
				_, cause = project.Load(root)
			} else {
				_, cause = project.Project(root, "")
			}
			if cause == nil {
				t.Fatal("public operation accepted invalid metadata or set")
			}
			for _, diagnostic := range project.Check(root) {
				if !diagnostic.Warning && diagnostic.Code == "C12" && strings.Contains(diagnostic.Message, cause.Error()) {
					return
				}
			}
			t.Errorf("Check did not report C12 for %s cause %q", tc.source, cause)
		})
	}
}

func Test_PRJ_044_CheckQueueIdentityCause(t *testing.T) {
	root := fixture(t)
	queued(t, root, "0001-repeat.md", "", "Needed.\n", "", "", "")
	queued(t, root, "0002-repeat.md", "", "Needed.\n", "", "", "")
	_, cause := project.Queue(root)
	if cause == nil {
		t.Fatal("Queue accepted duplicate slug")
	}
	for _, diagnostic := range project.Check(root) {
		if !diagnostic.Warning && diagnostic.Code == "C13" && strings.Contains(diagnostic.Message, cause.Error()) {
			return
		}
	}
	t.Errorf("Check did not report C13 for queue identity cause %q", cause)
}

func Test_PRJ_045_CheckPendingModuleTargetCause(t *testing.T) {
	root := fixture(t)
	put(t, root, "changes/0001-bad.md", "---\nmodule: missing\n---\n## Motivation\nNeeded.\n")
	_, cause := project.Queue(root)
	if cause == nil {
		t.Fatal("Queue accepted unknown change module")
	}
	for _, diagnostic := range project.Check(root) {
		if !diagnostic.Warning && diagnostic.Code == "C07" && strings.Contains(diagnostic.Message, cause.Error()) {
			return
		}
	}
	t.Errorf("Check did not report C07 for pending module cause %q", cause)
}

func Test_PRJ_013_CheckDuplicateWithinEditSection(t *testing.T) {
	root := fixture(t)
	queued(t, root, "0001-duplicate.md", "", "Needed.\n", "### new:local\nLocal MUST work.\n\n### new:local\nDuplicate MUST fail.\n", "", "")
	_, cause := project.Queue(root)
	if cause == nil {
		t.Fatal("Queue accepted duplicate Added identity")
	}
	for _, diagnostic := range project.Check(root) {
		if !diagnostic.Warning && diagnostic.Code == "C05" && strings.Contains(diagnostic.Message, cause.Error()) {
			return
		}
	}
	t.Errorf("Check did not report C05 for duplicate Added identity cause %q", cause)
}
