package evals_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/assets"
	"github.com/lurkshark/fireweed/packages/project"
)

func newModuleTemplates(t *testing.T, root string) {
	t.Helper()
	recipe, err := assets.BuiltinRecipe("go")
	if err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"SPEC.md", "api.go.tmpl"} {
		path := ".fireweed/templates/module/" + name
		appPut(t, root, path, recipe.Files[path])
	}
}

func newModuleAbsent(t *testing.T, root, relative string) {
	t.Helper()
	if _, err := os.Stat(filepath.Join(root, filepath.FromSlash(relative))); !os.IsNotExist(err) {
		t.Errorf("rejected request wrote %s or stat failed: %v", relative, err)
	}
}

func newModuleFailure(t *testing.T, bin, dir string, args []string, code string) {
	t.Helper()
	for _, jsonMode := range []bool{false, true} {
		call := append([]string(nil), args...)
		if jsonMode {
			call = append(call, "--json")
		}
		status, stdout, stderr := runAppBinary(t, bin, dir, call...)
		if status != 2 {
			t.Errorf("%v: status=%d, want 2; stdout=%q stderr=%q", call, status, stdout, stderr)
		}
		if jsonMode {
			if stderr != "" {
				t.Errorf("%v: JSON stderr=%q, want empty", call, stderr)
			}
			message := assertFailedJSON(t, stdout, code)
			if !corrective(message) {
				t.Errorf("%v: diagnostic %q needs corrective action", call, message)
			}
		} else if !strings.HasPrefix(strings.TrimSpace(stdout+stderr), code+":") || !corrective(stdout+stderr) {
			t.Errorf("%v: diagnostic=%q, want corrective %s", call, stdout+stderr, code)
		}
	}
}

func Test_APP_100_APP_101_APP_104_NewModuleFromNearestProjectTemplates(t *testing.T) {
	bin := buildAppBinary(t)
	outer := appProjectFixture(t)
	newModuleTemplates(t, outer)
	appPut(t, outer, ".fireweed/templates/module/README.txt.tmpl", "outer template\n")

	inner := filepath.Join(outer, "nested-project")
	appPut(t, inner, "fireweed.yaml", strings.Replace(appFixtureConfig, "packages/*", "components/*", 1))
	appPut(t, inner, "components/base/SPEC.md", "---\nname: base\nid_prefix: BAS\ndepends_on: []\n---\n# base\n\n## Specs\n")
	appPut(t, inner, "go.mod", "module example.test/nearest\n\ngo 1.23\n")
	newModuleTemplates(t, inner)
	appPut(t, inner, ".fireweed/templates/module/SPEC.md", "---\nname: {{name}}\nid_prefix: {{prefix}}\ndepends_on: []\n---\n# {{name}}\n\n## Specs\n\n## Decisions\n\nProject-owned SPEC marker.\n")
	appPut(t, inner, ".fireweed/templates/module/api.go.tmpl", "package {{package}}\n// Project-owned API: {{name}} {{prefix}}\n")
	appPut(t, inner, ".fireweed/templates/module/README.txt.tmpl", "name={{name}}\nprefix={{prefix}}\npackage={{package}}\nmodule_dir={{module_dir}}\ngo_module={{go_module}}\n")
	from := filepath.Join(inner, "components", "base", "deep")
	if err := os.MkdirAll(from, 0755); err != nil {
		t.Fatal(err)
	}
	status, stdout, stderr := runAppBinary(t, bin, from, "new-module", "quartz-core", "--prefix", "QTZ")
	if status != 0 || stderr != "" || strings.TrimSpace(stdout) != "components/quartz-core" {
		t.Errorf("plain creation: status=%d stdout=%q stderr=%q", status, stdout, stderr)
	}
	newModuleAbsent(t, outer, "packages/quartz-core")
	if got := readInitFile(t, inner, "components/quartz-core/SPEC.md"); !strings.Contains(got, "name: quartz-core\nid_prefix: QTZ") || !strings.Contains(got, "Project-owned SPEC marker.") {
		t.Errorf("rendered SPEC identity/project template: %q", got)
	}
	if got := readInitFile(t, inner, "components/quartz-core/api.go"); got != "package quartz_core\n// Project-owned API: quartz-core QTZ\n" {
		t.Errorf("rendered API=%q", got)
	}
	if got := readInitFile(t, inner, "components/quartz-core/README.txt"); got != "name=quartz-core\nprefix=QTZ\npackage=quartz_core\nmodule_dir=components/quartz-core\ngo_module=example.test/nearest\n" {
		t.Errorf("rendered substitutions=%q", got)
	}
	newModuleAbsent(t, inner, "components/quartz-core/api.go.tmpl")
	newModuleAbsent(t, inner, "components/quartz-core/README.txt.tmpl")
	snapshot, err := project.Load(inner)
	if err != nil {
		t.Fatalf("created module is not loadable: %v", err)
	}
	found := false
	for _, module := range snapshot.Modules {
		if module.Name == "quartz-core" {
			found = true
			if module.IDPrefix != "QTZ" || module.Dir != "components/quartz-core" {
				t.Errorf("created graph module=%+v", module)
			}
		}
	}
	if !found {
		t.Errorf("created module absent from project graph: %+v", snapshot.Modules)
	}
	if err := project.ValidateGraph(snapshot.Modules); err != nil {
		t.Errorf("created graph invalid: %v", err)
	}
	status, stdout, stderr = runAppBinary(t, bin, from, "modules", "--json")
	if status != 0 || stderr != "" || !strings.Contains(stdout, `"name":"quartz-core"`) {
		t.Errorf("CLI cannot discover created module: status=%d stdout=%q stderr=%q", status, stdout, stderr)
	}

	status, stdout, stderr = runAppBinary(t, bin, from, "new-module", "second", "--prefix", "SEC", "--json")
	if status != 0 || stderr != "" {
		t.Errorf("JSON creation: status=%d stdout=%q stderr=%q", status, stdout, stderr)
	}
	if got := commandStringData(t, stdout); got != "components/second" {
		t.Errorf("JSON data=%q, want repository-relative module directory", got)
	}
	if got := readInitFile(t, inner, "components/second/README.txt"); !strings.Contains(got, "module_dir=components/second\n") || !strings.Contains(got, "go_module=example.test/nearest\n") {
		t.Errorf("second module substitutions=%q", got)
	}
}

func Test_APP_102_APP_103_APP_104_NewModuleConflictsPreserveContents(t *testing.T) {
	bin := buildAppBinary(t)
	for _, tc := range []struct {
		name, module, prefix, target string
		prepare                      func(*testing.T, string)
	}{
		{"existing-name", "alpha", "NEW", "components/alpha", func(t *testing.T, root string) {}},
		{"existing-prefix", "fresh", "ALP", "components/fresh", func(t *testing.T, root string) {}},
		{"occupied-directory", "fresh", "FRS", "components/fresh", func(t *testing.T, root string) {
			appPut(t, root, "components/fresh/SPEC.md", "---\nname: occupied\nid_prefix: OCC\ndepends_on: []\n---\n# occupied\n\n## Specs\n")
			appPut(t, root, "components/fresh/keep.txt", "keep occupied directory\n")
		}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := t.TempDir()
			appPut(t, root, "fireweed.yaml", strings.Replace(appFixtureConfig, `"packages/*"`, `"components/*", "packages/*"`, 1))
			appPut(t, root, "packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\n---\n# alpha\n\n## Specs\n")
			newModuleTemplates(t, root)
			tc.prepare(t, root)
			if _, err := project.Load(root); err != nil {
				t.Fatalf("conflict fixture must have a valid existing graph: %v", err)
			}
			before := readInitFile(t, root, "packages/alpha/SPEC.md")
			occupiedBefore := ""
			if tc.name == "occupied-directory" {
				occupiedBefore = readInitFile(t, root, "components/fresh/SPEC.md")
			}
			newModuleFailure(t, bin, root, []string{"new-module", tc.module, "--prefix", tc.prefix}, "P01")
			if got := readInitFile(t, root, "packages/alpha/SPEC.md"); got != before {
				t.Errorf("existing module changed: %q", got)
			}
			if tc.name == "occupied-directory" {
				if got := readInitFile(t, root, "components/fresh/SPEC.md"); got != occupiedBefore {
					t.Errorf("occupied module changed: %q", got)
				}
				if got := readInitFile(t, root, "components/fresh/keep.txt"); got != "keep occupied directory\n" {
					t.Errorf("occupied contents changed: %q", got)
				}
				newModuleAbsent(t, root, "components/fresh/api.go")
			} else {
				newModuleAbsent(t, root, tc.target)
			}
		})
	}
}

func Test_APP_103_APP_104_NewModuleUsageTemplatesAndProjectErrors(t *testing.T) {
	bin := buildAppBinary(t)
	for _, tc := range []struct {
		name, code, target string
		args               []string
		prepare            func(*testing.T, string)
	}{
		{"missing-name", "U01", "packages/new", []string{"new-module", "--prefix", "NEW"}, nil},
		{"missing-prefix", "U01", "packages/new", []string{"new-module", "new"}, nil},
		{"invalid-name", "U01", "packages/new", []string{"new-module", "bad/name", "--prefix", "NEW"}, nil},
		{"invalid-prefix", "U01", "packages/new", []string{"new-module", "new", "--prefix", "bad-prefix"}, nil},
		{"repeated-prefix", "U01", "packages/new", []string{"new-module", "new", "--prefix", "NEW", "--prefix", "AGAIN"}, nil},
		{"extra-argument", "U01", "packages/new", []string{"new-module", "new", "--prefix", "NEW", "extra"}, nil},
		{"unknown-option", "U01", "packages/new", []string{"new-module", "new", "--prefix", "NEW", "--bogus"}, nil},
		{"missing-template", "C03", "packages/new", []string{"new-module", "new", "--prefix", "NEW"}, func(t *testing.T, root string) {
			if err := os.RemoveAll(filepath.Join(root, ".fireweed/templates/module")); err != nil {
				t.Fatal(err)
			}
		}},
		{"malformed-template", "C03", "packages/new", []string{"new-module", "new", "--prefix", "NEW"}, func(t *testing.T, root string) {
			appPut(t, root, ".fireweed/templates/module/SPEC.md", "---\nname: [broken\n---\n")
		}},
		{"invalid-configuration", "C01", "packages/new", []string{"new-module", "new", "--prefix", "NEW"}, func(t *testing.T, root string) {
			appPut(t, root, "fireweed.yaml", "modules: [broken\n")
		}},
		{"invalid-existing-graph", "C02", "packages/new", []string{"new-module", "new", "--prefix", "NEW"}, func(t *testing.T, root string) {
			appPut(t, root, "packages/beta/SPEC.md", "---\nname: beta\nid_prefix: BET\ndepends_on: [missing]\n---\n# Beta\n\n## Specs\n")
		}},
		{"no-configured-destination", "C01", "packages/new", []string{"new-module", "new", "--prefix", "NEW"}, func(t *testing.T, root string) {
			appPut(t, root, "fireweed.yaml", strings.Replace(appFixtureConfig, "packages/*", "packages/*/internal", 1))
			appPut(t, root, "packages/alpha/internal/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\n---\n# alpha\n\n## Specs\n")
		}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := appProjectFixture(t)
			newModuleTemplates(t, root)
			if tc.prepare != nil {
				tc.prepare(t, root)
			}
			newModuleFailure(t, bin, root, tc.args, tc.code)
			newModuleAbsent(t, root, tc.target)
		})
	}
}

func Test_APP_101_APP_104_NewModuleEmptyGoModule(t *testing.T) {
	bin := buildAppBinary(t)
	root := appProjectFixture(t)
	newModuleTemplates(t, root)
	appPut(t, root, ".fireweed/templates/module/README.tmpl", "go_module={{go_module}}\nmodule_dir={{module_dir}}\n")
	status, stdout, stderr := runAppBinary(t, bin, root, "new-module", "empty-mod", "--prefix", "EMP", "--json")
	if status != 0 || stderr != "" || commandStringData(t, stdout) != "packages/empty-mod" {
		t.Errorf("creation without go.mod: status=%d stdout=%q stderr=%q", status, stdout, stderr)
	}
	if got := readInitFile(t, root, "packages/empty-mod/README"); got != "go_module=\nmodule_dir=packages/empty-mod\n" {
		t.Errorf("empty go_module substitution=%q", got)
	}
}
