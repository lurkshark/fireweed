package evals_test

import (
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
)

const fixtureConfig = `fireweed: ">=0.1, <0.2"
version: 1
base: main
modules: ["packages/*"]
module_defaults:
  interface: [api.go]
  evals: [evals/**]
  disposable: [internal/**]
  preserve: [data/**]
`

func put(t *testing.T, root, name, body string) {
	t.Helper()
	path := filepath.Join(root, filepath.FromSlash(name))
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(body), 0644); err != nil {
		t.Fatal(err)
	}
}

func fixture(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	put(t, root, "fireweed.yaml", fixtureConfig)
	put(t, root, "packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\nevals: []\n---\n# Alpha\n\n## Specs\n\n### ALP-001\nAlpha behavior.\n")
	put(t, root, "packages/beta/SPEC.md", "---\nname: beta\nid_prefix: BET\ndepends_on: [alpha]\n---\n# Beta\n\n## Specs\n")
	return root
}

func Test_PRJ_001_LoadConfiguredModulesAndDefaults(t *testing.T) {
	root := fixture(t)
	put(t, root, "other/ignored/SPEC.md", "not a document")
	got, err := project.Load(root)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if got.Root != root {
		t.Errorf("Root = %q; want %q", got.Root, root)
	}
	if got.Config["version"] != 1 {
		t.Errorf("decoded Config version = %#v", got.Config["version"])
	}
	byName := map[string]project.Module{}
	for _, m := range got.Modules {
		byName[m.Name] = m
	}
	if len(got.Modules) != 2 || len(byName) != 2 {
		t.Fatalf("modules = %#v; want exactly alpha and beta", got.Modules)
	}
	alpha, beta := byName["alpha"], byName["beta"]
	if alpha.IDPrefix != "ALP" || alpha.Dir != "packages/alpha" || len(alpha.DependsOn) != 0 {
		t.Errorf("alpha identity = %#v", alpha)
	}
	if beta.IDPrefix != "BET" || beta.Dir != "packages/beta" || !reflect.DeepEqual(beta.DependsOn, []string{"alpha"}) {
		t.Errorf("beta identity/dependencies = %#v", beta)
	}
	if alpha.Document.Frontmatter["name"] != "alpha" || !strings.Contains(alpha.Document.Body, "### ALP-001\nAlpha behavior.") {
		t.Errorf("alpha Document = %#v", alpha.Document)
	}
	if !reflect.DeepEqual(alpha.Zones.Interface, []string{"api.go"}) || !reflect.DeepEqual(alpha.Zones.Disposable, []string{"internal/**"}) || !reflect.DeepEqual(alpha.Zones.Preserve, []string{"data/**"}) {
		t.Errorf("inherited zones = %#v", alpha.Zones)
	}
	if alpha.Zones.Evals == nil || len(alpha.Zones.Evals) != 0 {
		t.Errorf("explicit empty evals zone lost: %#v", alpha.Zones.Evals)
	}
	if !reflect.DeepEqual(beta.Zones.Evals, []string{"evals/**"}) {
		t.Errorf("default evals zone = %#v", beta.Zones.Evals)
	}
}

func Test_PRJ_001_LoadInvalidDocuments(t *testing.T) {
	for _, tc := range []struct {
		name string
		edit func(*testing.T, string)
	}{
		{"missing_config", func(t *testing.T, root string) {
			if err := os.Remove(filepath.Join(root, "fireweed.yaml")); err != nil {
				t.Fatal(err)
			}
		}},
		{"malformed_config", func(t *testing.T, root string) { put(t, root, "fireweed.yaml", "modules: [broken\n") }},
		{"missing_module_document", func(t *testing.T, root string) {
			if err := os.Remove(filepath.Join(root, "packages/alpha/SPEC.md")); err != nil {
				t.Fatal(err)
			}
		}},
		{"malformed_module_document", func(t *testing.T, root string) {
			put(t, root, "packages/alpha/SPEC.md", "---\nname: [broken\n---\n# Alpha\n")
		}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := fixture(t)
			tc.edit(t, root)
			if _, err := project.Load(root); err == nil {
				t.Error("Load accepted missing or malformed project document")
			}
		})
	}
}

func Test_PRJ_002_Graph(t *testing.T) {
	module := func(name, prefix string, deps ...string) project.Module {
		return project.Module{Name: name, IDPrefix: prefix, DependsOn: deps}
	}
	for _, tc := range []struct {
		name    string
		modules []project.Module
		valid   bool
	}{
		{"acyclic", []project.Module{module("alpha", "ALP"), module("beta", "BET", "alpha"), module("gamma", "GAM", "beta")}, true},
		{"empty_name", []project.Module{module("", "ALP")}, false},
		{"empty_prefix", []project.Module{module("alpha", "")}, false},
		{"duplicate_name", []project.Module{module("alpha", "ALP"), module("alpha", "BET")}, false},
		{"duplicate_prefix", []project.Module{module("alpha", "ALP"), module("beta", "ALP")}, false},
		{"missing_dependency", []project.Module{module("alpha", "ALP", "missing")}, false},
		{"self_dependency", []project.Module{module("alpha", "ALP", "alpha")}, false},
		{"cycle", []project.Module{module("alpha", "ALP", "beta"), module("beta", "BET", "alpha")}, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			err := project.ValidateGraph(tc.modules)
			if tc.valid && err != nil {
				t.Errorf("valid graph rejected: %v", err)
			}
			if !tc.valid && (err == nil || strings.TrimSpace(err.Error()) == "") {
				t.Errorf("invalid graph needs an explanatory error; got %v", err)
			}
		})
	}
}

func Test_PRJ_003_NewChangeSkeletonAndNumbering(t *testing.T) {
	root := fixture(t)
	put(t, root, "changes/0003-waiting.md", "pending marker")
	put(t, root, "changes/archive/0007-finished.md", "archived marker")
	rel, err := project.NewChange(root, "alpha", "new-feature")
	if err != nil {
		t.Fatalf("NewChange: %v", err)
	}
	if rel != "changes/0008-new-feature.md" {
		t.Errorf("relative path = %q; want changes/0008-new-feature.md", rel)
	}
	data, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(rel)))
	if err != nil {
		t.Fatalf("read new change: %v", err)
	}
	doc, err := contracts.ParseDocument(string(data))
	if err != nil {
		t.Fatalf("new change document: %v", err)
	}
	if doc.Frontmatter["module"] != "alpha" {
		t.Errorf("new change module = %#v", doc.Frontmatter["module"])
	}
	for _, heading := range []string{"Motivation", "Interface", "Added", "Modified", "Removed", "Decisions"} {
		if _, ok := contracts.Section(doc.Body, heading); !ok {
			t.Errorf("missing %s section", heading)
		}
	}
	if regexp.MustCompile(`(?m)^### [A-Z]+-[0-9]+\b`).Match(data) {
		t.Errorf("skeleton assigned a spec ID: %s", data)
	}
}

func Test_PRJ_003_NewChangeRejectsDuplicatesWithoutOverwrite(t *testing.T) {
	for _, archived := range []bool{false, true} {
		name := "pending"
		path := "changes/0002-repeat.md"
		if archived {
			name = "archived"
			path = "changes/archive/0002-repeat.md"
		}
		t.Run(name, func(t *testing.T) {
			root := fixture(t)
			put(t, root, path, "original marker")
			if _, err := project.NewChange(root, "alpha", "repeat"); err == nil {
				t.Error("duplicate slug accepted")
			}
			data, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(path)))
			if err != nil || string(data) != "original marker" {
				t.Errorf("existing change overwritten: %q, %v", data, err)
			}
		})
	}
}

func Test_PRJ_003_NewChangeInvalidArguments(t *testing.T) {
	for _, tc := range []struct{ name, module, slug string }{
		{"unknown_module", "missing", "valid"},
		{"uppercase_slug", "alpha", "Bad"},
		{"underscore_slug", "alpha", "bad_name"},
		{"slash_slug", "alpha", "../bad"},
		{"empty_slug", "alpha", ""},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := fixture(t)
			if _, err := project.NewChange(root, tc.module, tc.slug); err == nil {
				t.Errorf("accepted module %q and slug %q", tc.module, tc.slug)
			}
		})
	}
}
