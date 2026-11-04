package evals_test

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/lurkshark/fireweed/packages/assets"
	"github.com/lurkshark/fireweed/packages/contracts"
	"gopkg.in/yaml.v3"
)

var additionalRecipeNames = []string{"typescript-pnpm", "python-uv"}

func additionalRecipe(t *testing.T, name string) assets.Recipe {
	t.Helper()
	r, err := assets.BuiltinRecipe(name)
	if err != nil {
		t.Fatalf("BuiltinRecipe(%q): %v", name, err)
	}
	return r
}

func Test_AST_004_AdditionalBuiltinNames(t *testing.T) {
	if goRecipe(t).Name != "go" {
		t.Error("Go recipe changed name")
	}
	for _, name := range additionalRecipeNames {
		if got := additionalRecipe(t, name).Name; got != name {
			t.Errorf("BuiltinRecipe(%q).Name = %q", name, got)
		}
	}
	if _, err := assets.BuiltinRecipe("not-a-recipe"); err == nil {
		t.Error("unknown recipe accepted")
	}
}

func Test_AST_049_BundleMetadata(t *testing.T) {
	for _, name := range additionalRecipeNames {
		t.Run(name, func(t *testing.T) {
			r := additionalRecipe(t, name)
			if strings.TrimSpace(r.Instructions) == "" {
				t.Error("empty instructions")
			}
			for label, paths := range map[string][]string{
				"interface": r.ModuleDefaults.Interface, "evals": r.ModuleDefaults.Evals,
				"disposable": r.ModuleDefaults.Disposable,
			} {
				if len(paths) == 0 {
					t.Errorf("empty %s defaults", label)
				}
			}
			for _, hook := range []string{"eval", "boundaries", "setup"} {
				if cmd, ok := r.Hooks[hook].(string); !ok || strings.TrimSpace(cmd) == "" {
					t.Errorf("empty %s hook command", hook)
				}
			}
			for _, path := range []string{".fireweed/**", "fireweed.yaml"} {
				if !containsString(r.Protected, path) {
					t.Errorf("missing protected glue %q", path)
				}
			}
			lock := map[string]string{"typescript-pnpm": "pnpm-lock.yaml", "python-uv": "uv.lock"}[name]
			if !containsString(r.WorkerWritable, lock) {
				t.Errorf("missing worker-writable %s", lock)
			}
			for _, path := range []string{
				".fireweed/recipes/" + name + ".md", ".fireweed/canaries/canaries.yaml",
				".fireweed/templates/module/SPEC.md",
			} {
				if strings.TrimSpace(r.Files[path]) == "" {
					t.Errorf("missing project file %s", path)
				}
			}
			hasHook, hasTemplate := false, false
			for path, body := range r.Files {
				if !strings.HasPrefix(path, ".fireweed/") || strings.TrimSpace(body) == "" {
					t.Errorf("invalid bundle file %q", path)
				}
				for _, hook := range []string{"eval", "boundaries", "setup"} {
					if command, ok := r.Hooks[hook].(string); ok && strings.Contains(command, path) {
						hasHook = true
					}
				}
				if strings.HasPrefix(path, ".fireweed/templates/module/") && path != ".fireweed/templates/module/SPEC.md" {
					hasTemplate = true
				}
			}
			if !hasHook || !hasTemplate {
				t.Errorf("missing hook or language module template: hook=%v template=%v", hasHook, hasTemplate)
			}
		})
	}
}

func Test_AST_050_WorkspaceInstructions(t *testing.T) {
	for _, name := range additionalRecipeNames {
		t.Run(name, func(t *testing.T) {
			r := additionalRecipe(t, name)
			guide := strings.ToLower(r.Instructions + "\n" + r.Files[".fireweed/recipes/"+name+".md"])
			for _, word := range []string{"prerequisit", "workspace", "fireweed_bin modules --json", "depends_on", "zones", "dir"} {
				if !strings.Contains(guide, word) {
					t.Errorf("instructions omit %q", word)
				}
			}
			for _, word := range map[string][]string{
				"typescript-pnpm": {"pnpm", "pnpm-workspace.yaml", "package.json", "exports"},
				"python-uv":       {"uv", "pyproject.toml", "import linter"},
			}[name] {
				if !strings.Contains(guide, word) {
					t.Errorf("instructions omit %q", word)
				}
			}
			if !regexp.MustCompile(`(?s)(derive|generate|build).*fireweed_bin modules --json|fireweed_bin modules --json.*(derive|generate|build)`).MatchString(guide) {
				t.Error("instructions do not explain graph derivation")
			}
		})
	}
}

type extraCanaries struct {
	Version      int               `yaml:"version"`
	Baseline     map[string]string `yaml:"baseline"`
	Undeclared   map[string]string `yaml:"undeclared"`
	InternalEval map[string]string `yaml:"internal_eval"`
	FailingEval  map[string]string `yaml:"failing_eval"`
	MissingID    map[string]string `yaml:"missing_id"`
}

func readExtraCanaries(t *testing.T, r assets.Recipe) extraCanaries {
	t.Helper()
	var c extraCanaries
	if err := yaml.Unmarshal([]byte(r.Files[".fireweed/canaries/canaries.yaml"]), &c); err != nil {
		t.Fatal(err)
	}
	return c
}

func Test_AST_051_TemplateSubstitutions(t *testing.T) {
	for _, name := range additionalRecipeNames {
		t.Run(name, func(t *testing.T) {
			r := additionalRecipe(t, name)
			c := readExtraCanaries(t, r)
			for _, field := range []string{"name", "prefix", "package", "module_dir", "other_dir"} {
				found := false
				for path, body := range r.Files {
					if strings.HasPrefix(path, ".fireweed/templates/module/") && strings.Contains(body, "{{"+field+"}}") {
						found = true
					}
				}
				for _, files := range []map[string]string{c.Baseline, c.Undeclared, c.InternalEval, c.FailingEval, c.MissingID} {
					for path, body := range files {
						if strings.Contains(path+body, "{{"+field+"}}") {
							found = true
						}
					}
				}
				if !found {
					t.Errorf("doctor field %s unused", field)
				}
			}
			for path, body := range r.Files {
				if !strings.HasPrefix(path, ".fireweed/templates/module/") {
					continue
				}
				if unrendered := regexp.MustCompile(`\{\{[^}]+\}\}`).FindString(renderExtra(body, "alpha")); unrendered != "" {
					t.Errorf("%s has unsupported placeholder %q", path, unrendered)
				}
			}
		})
	}
}

func Test_AST_052_CanaryMap(t *testing.T) {
	for _, name := range additionalRecipeNames {
		t.Run(name, func(t *testing.T) {
			c := readExtraCanaries(t, additionalRecipe(t, name))
			if c.Version != 1 {
				t.Errorf("version = %d", c.Version)
			}
			for label, files := range map[string]map[string]string{"baseline": c.Baseline, "undeclared": c.Undeclared, "internal_eval": c.InternalEval, "failing_eval": c.FailingEval, "missing_id": c.MissingID} {
				if len(files) == 0 {
					t.Errorf("%s empty", label)
				}
				for path, body := range files {
					if path == "" || filepath.IsAbs(path) || strings.HasPrefix(path, "../") || body == "" {
						t.Errorf("%s has invalid entry %q", label, path)
					}
				}
			}
		})
	}
}

func renderExtra(s, module string) string {
	other := "beta"
	if module == "beta" {
		other = "alpha"
	}
	fields := map[string]string{"name": module, "prefix": map[string]string{"alpha": "ALC", "beta": "BET"}[module], "package": module, "module_dir": "packages/" + module, "other_dir": "packages/" + other, "go_module": "example.com/fixture"}
	for k, v := range fields {
		s = strings.ReplaceAll(s, "{{"+k+"}}", v)
	}
	return s
}

type extraWorkspace struct {
	t        *testing.T
	root     string
	recipe   assets.Recipe
	canaries extraCanaries
}

func newExtraWorkspace(t *testing.T, name string) *extraWorkspace {
	t.Helper()
	r := additionalRecipe(t, name)
	w := &extraWorkspace{t: t, root: t.TempDir(), recipe: r, canaries: readExtraCanaries(t, r)}
	if name == "typescript-pnpm" {
		w.write("package.json", `{"private":true,"type":"module","packageManager":"pnpm@10.15.1"}`)
		w.write("pnpm-workspace.yaml", "packages:\n  - 'packages/*'\n")
	} else {
		w.write("pyproject.toml", "[project]\nname = 'recipe-fixture'\nversion = '0.1.0'\nrequires-python = '>=3.11'\n[tool.uv.workspace]\nmembers = ['packages/*']\n")
	}
	for path, body := range r.Files {
		w.write(path, body)
	}
	for _, module := range []string{"alpha", "beta"} {
		for path, body := range r.Files {
			if !strings.HasPrefix(path, ".fireweed/templates/module/") {
				continue
			}
			w.write(filepath.Join("packages", module, strings.TrimSuffix(strings.TrimPrefix(path, ".fireweed/templates/module/"), ".tmpl")), renderExtra(body, module))
		}
		w.apply(w.canaries.Baseline, module)
	}
	w.graph(nil, "evals/**", r.ModuleDefaults.Interface)
	return w
}

func (w *extraWorkspace) write(path, body string) {
	w.t.Helper()
	abs := filepath.Join(w.root, path)
	if err := os.MkdirAll(filepath.Dir(abs), 0755); err != nil {
		w.t.Fatal(err)
	}
	if err := os.WriteFile(abs, []byte(body), 0644); err != nil {
		w.t.Fatal(err)
	}
}
func (w *extraWorkspace) apply(files map[string]string, module string) []string {
	w.t.Helper()
	var paths []string
	for path, body := range files {
		if unrendered := regexp.MustCompile(`\{\{[^}]+\}\}`).FindString(renderExtra(path+body, module)); unrendered != "" {
			w.t.Fatalf("unrendered %s", unrendered)
		}
		full := filepath.Join("packages", module, strings.TrimSuffix(renderExtra(path, module), ".tmpl"))
		w.write(full, renderExtra(body, module))
		paths = append(paths, full)
	}
	return paths
}
func (w *extraWorkspace) remove(paths []string) {
	for _, path := range paths {
		if err := os.Remove(filepath.Join(w.root, path)); err != nil {
			w.t.Fatal(err)
		}
	}
}
func (w *extraWorkspace) graph(deps []string, eval string, betaInterface []string) {
	if deps == nil {
		deps = []string{}
	}
	modules := []map[string]any{
		{"name": "alpha", "dir": "packages/alpha", "depends_on": deps, "zones": map[string]any{"interface": w.recipe.ModuleDefaults.Interface, "evals": []string{eval}, "disposable": w.recipe.ModuleDefaults.Disposable}},
		{"name": "beta", "dir": "packages/beta", "depends_on": []string{}, "zones": map[string]any{"interface": betaInterface, "evals": []string{eval}, "disposable": w.recipe.ModuleDefaults.Disposable}},
	}
	b, err := json.Marshal(map[string]any{"data": modules})
	if err != nil {
		w.t.Fatal(err)
	}
	w.write("modules.json", string(b))
	w.write("fake-fireweed", "#!/bin/sh\nif [ \"$1\" != modules ] || [ \"$2\" != --json ]; then exit 2; fi\nprintf '%s\\n' \"$*\" >> \"$FIREWEED_QUERY_LOG\"\ncat \"$FIREWEED_MODULES_JSON\"\n")
	if err := os.Chmod(filepath.Join(w.root, "fake-fireweed"), 0755); err != nil {
		w.t.Fatal(err)
	}
}
func (w *extraWorkspace) hook(name, module string) (junitSuite, string, error) {
	w.t.Helper()
	cmdText, ok := w.recipe.Hooks[name].(string)
	if !ok || cmdText == "" {
		w.t.Fatalf("missing %s hook", name)
	}
	report := filepath.Join(w.root, module+"-report.xml")
	_ = os.Remove(report)
	for k, v := range map[string]string{"{module}": module, "{module_dir}": "packages/" + module, "{report}": report} {
		cmdText = strings.ReplaceAll(cmdText, k, v)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 45*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, "sh", "-c", cmdText)
	cmd.Dir = w.root
	cmd.Env = append(os.Environ(), "FIREWEED_BIN="+filepath.Join(w.root, "fake-fireweed"), "FIREWEED_MODULES_JSON="+filepath.Join(w.root, "modules.json"), "FIREWEED_QUERY_LOG="+filepath.Join(w.root, "modules-query.log"), "UV_CACHE_DIR="+filepath.Join(w.root, "uv-cache"), "UV_OFFLINE=1", "COREPACK_ENABLE_NETWORK=0")
	out, err := cmd.CombinedOutput()
	if ctx.Err() != nil {
		w.t.Fatalf("%s timed out: %s", name, out)
	}
	var suite junitSuite
	if name == "eval" {
		data, readErr := os.ReadFile(report)
		if readErr != nil {
			w.t.Fatalf("eval report missing: %v; hook=%v output=%s", readErr, err, out)
		}
		if xmlErr := xml.Unmarshal(data, &suite); xmlErr != nil {
			w.t.Fatalf("bad JUnit: %v: %s", xmlErr, data)
		}
		if len(suite.Cases) == 0 {
			var nested struct {
				Suites []junitSuite `xml:"testsuite"`
			}
			if xmlErr := xml.Unmarshal(data, &nested); xmlErr == nil {
				for _, part := range nested.Suites {
					suite.Cases = append(suite.Cases, part.Cases...)
				}
			}
		}
	}
	return suite, string(out), err
}

func Test_AST_053_BaselineHooks(t *testing.T) {
	for _, name := range additionalRecipeNames {
		t.Run(name, func(t *testing.T) {
			w := newExtraWorkspace(t, name)
			for _, module := range []string{"alpha", "beta"} {
				s, out, err := w.hook("eval", module)
				if err != nil || !s.hasPass() {
					t.Errorf("%s baseline eval: %v %s %+v", module, err, out, s.Cases)
				}
			}
			_, out, err := w.hook("boundaries", "alpha")
			if err != nil {
				t.Errorf("baseline boundaries: %v %s", err, out)
			}
		})
	}
}

func Test_AST_054_EvalJUnitAndExit(t *testing.T) {
	for _, name := range additionalRecipeNames {
		t.Run(name, func(t *testing.T) {
			w := newExtraWorkspace(t, name)
			baseline, out, err := w.hook("eval", "alpha")
			if err != nil || !baseline.hasPass() {
				t.Fatalf("baseline: %v %s %+v", err, out, baseline.Cases)
			}
			skip := map[string][2]string{
				"typescript-pnpm": {"packages/alpha/evals/skip.test.ts", "import { test } from 'node:test'; test.skip('ALC001Skip', () => {});\n"},
				"python-uv":       {"packages/alpha/evals/test_skip.py", "import unittest\nclass TestSkip(unittest.TestCase):\n    @unittest.skip('doctor fixture')\n    def test_ALC001Skip(self):\n        pass\n"},
			}[name]
			w.write(skip[0], skip[1])
			skipped, out, err := w.hook("eval", "alpha")
			foundSkip := false
			for _, c := range skipped.Cases {
				if strings.Contains(c.Name, "ALC001Skip") && c.state() == "skip" {
					foundSkip = true
				}
			}
			if err != nil || !foundSkip {
				t.Errorf("skip not reported: %v %s %+v", err, out, skipped.Cases)
			}
			w.remove([]string{skip[0]})
			otherPaths := w.apply(w.canaries.FailingEval, "beta")
			selected, out, err := w.hook("eval", "alpha")
			if err != nil || selected.hasFailure() {
				t.Errorf("alpha eval ran beta failure: %v %s %+v", err, out, selected.Cases)
			}
			w.remove(otherPaths)
			paths := w.apply(w.canaries.FailingEval, "alpha")
			s, out, err := w.hook("eval", "alpha")
			if err == nil || !s.hasFailure() {
				t.Errorf("failure exit/report: %v %s %+v", err, out, s.Cases)
			}
			for _, c := range s.Cases {
				if c.Name == "" || c.Classname == "" {
					t.Errorf("missing testcase identity: %+v", c)
				}
				seconds, parseErr := strconv.ParseFloat(c.Time, 64)
				if parseErr != nil || seconds < 0 {
					t.Errorf("invalid elapsed seconds: %+v", c)
				}
			}
			w.remove(paths)
			if s, out, err = w.hook("eval", "alpha"); err != nil || s.hasFailure() {
				t.Errorf("eval did not recover: %v %s %+v", err, out, s.Cases)
			}
		})
	}
}

func Test_AST_055_IncompleteAndCompileEvidence(t *testing.T) {
	for _, name := range additionalRecipeNames {
		t.Run(name, func(t *testing.T) {
			w := newExtraWorkspace(t, name)
			incomplete := map[string][2]string{
				"typescript-pnpm": {"packages/alpha/evals/incomplete.test.ts", "import { test } from 'node:test'; test('ALC001Unfinished', () => process.exit(3));\n"},
				"python-uv":       {"packages/alpha/evals/test_incomplete.py", "import os\nimport unittest\nclass TestIncomplete(unittest.TestCase):\n    def test_ALC001Unfinished(self):\n        os._exit(3)\n"},
			}[name]
			w.write(incomplete[0], incomplete[1])
			s, out, err := w.hook("eval", "alpha")
			unfinishedFailed := false
			for _, c := range s.Cases {
				if strings.Contains(c.Name, "ALC001Unfinished") && c.state() == "fail" {
					unfinishedFailed = true
				}
			}
			if err == nil || !unfinishedFailed {
				t.Errorf("unfinished test hidden: %v %s %+v", err, out, s.Cases)
			}
			w.remove([]string{incomplete[0]})
			broken := map[string][2]string{"typescript-pnpm": {"packages/alpha/evals/broken.test.ts", "this is not valid TypeScript !!!\n"}, "python-uv": {"packages/alpha/evals/test_broken.py", "def broken(:\n"}}[name]
			w.write(broken[0], broken[1])
			s, out, err = w.hook("eval", "alpha")
			if err == nil || !s.hasFailure() {
				t.Errorf("compile/collection failure hidden: %v %s %+v", err, out, s.Cases)
			}
		})
	}
}

func Test_AST_056_GraphDerivedBoundaries(t *testing.T) {
	for _, name := range additionalRecipeNames {
		t.Run(name, func(t *testing.T) {
			w := newExtraWorkspace(t, name)
			_, out, err := w.hook("boundaries", "alpha")
			if err != nil {
				t.Fatalf("clean boundaries: %v %s", err, out)
			}
			query, err := os.ReadFile(filepath.Join(w.root, "modules-query.log"))
			if err != nil || !strings.Contains(string(query), "modules --json") {
				t.Errorf("no FIREWEED_BIN modules --json query: %s %v", query, err)
			}
			paths := w.apply(w.canaries.Undeclared, "alpha")
			_, out, err = w.hook("boundaries", "alpha")
			if err == nil || !strings.Contains(out, "alpha:") {
				t.Errorf("undeclared dependency not derived: %v %s", err, out)
			}
			w.graph([]string{"beta"}, "evals/**", w.recipe.ModuleDefaults.Interface)
			_, out, err = w.hook("boundaries", "alpha")
			if err != nil {
				t.Errorf("declared dependency rejected: %v %s", err, out)
			}
			w.remove(paths)
		})
	}
}

func Test_AST_057_EvalInterfaceAndOwnImplementation(t *testing.T) {
	for _, name := range additionalRecipeNames {
		t.Run(name, func(t *testing.T) {
			w := newExtraWorkspace(t, name)
			w.graph([]string{"beta"}, "evals/**", w.recipe.ModuleDefaults.Interface)
			paths := w.apply(w.canaries.InternalEval, "alpha")
			_, out, err := w.hook("boundaries", "alpha")
			if err == nil || !strings.Contains(out, "alpha:") {
				t.Errorf("own implementation import not reported: %v %s", err, out)
			}
			w.remove(paths)
			_, out, err = w.hook("boundaries", "alpha")
			if err != nil {
				t.Errorf("clean interface rejected: %v %s", err, out)
			}
			// Move the canary implementation to a declared dependency and import
			// it from alpha's eval zone. Its path remains outside beta's interface.
			var internalPath, internalBody string
			for path, body := range w.canaries.InternalEval {
				if !strings.HasPrefix(path, "evals/") && strings.Contains(path, "internal") {
					internalPath, internalBody = path, body
					break
				}
			}
			if internalPath == "" {
				t.Fatal("internal_eval canary has no implementation file")
			}
			foreignSource := w.apply(map[string]string{internalPath: internalBody}, "beta")
			var foreignEval string
			if name == "typescript-pnpm" {
				foreignEval = "packages/alpha/evals/foreign.test.ts"
				w.write(foreignEval, "import '../../beta/"+strings.TrimSuffix(internalPath, ".tmpl")+"';\n")
			} else {
				foreignEval = "packages/alpha/evals/test_foreign.py"
				importName := strings.TrimSuffix(strings.TrimPrefix(strings.TrimSuffix(internalPath, ".tmpl"), "src/"), ".py")
				importName = strings.TrimPrefix(importName, "{{package}}/")
				importName = strings.ReplaceAll(importName, "/", ".")
				if !strings.HasPrefix(importName, "beta.") {
					importName = "beta." + importName
				}
				w.write(foreignEval, "import "+importName+"\n")
			}
			_, out, err = w.hook("boundaries", "alpha")
			if err == nil || !strings.Contains(out, "alpha:") {
				t.Errorf("declared dependency implementation import not reported: %v %s", err, out)
			}
			w.remove(append(foreignSource, foreignEval))
			_, out, err = w.hook("boundaries", "alpha")
			if err != nil {
				t.Errorf("dependency interface recovery: %v %s", err, out)
			}
		})
	}
}

func Test_AST_058_CanaryDetectionAndRecovery(t *testing.T) {
	for _, name := range additionalRecipeNames {
		t.Run(name, func(t *testing.T) {
			w := newExtraWorkspace(t, name)
			for _, tc := range []struct {
				name, hook string
				files      map[string]string
			}{
				{"undeclared", "boundaries", w.canaries.Undeclared}, {"internal_eval", "boundaries", w.canaries.InternalEval},
				{"failing_eval", "eval", w.canaries.FailingEval}, {"missing_id", "eval", w.canaries.MissingID},
			} {
				t.Run(tc.name, func(t *testing.T) {
					paths := w.apply(tc.files, "alpha")
					s, out, err := w.hook(tc.hook, "alpha")
					switch tc.name {
					case "undeclared", "internal_eval":
						if err == nil || !strings.Contains(out, "alpha:") {
							t.Errorf("canary undetected: %v %s", err, out)
						}
					case "failing_eval":
						referencedFailure := false
						for _, c := range s.Cases {
							if c.state() == "fail" && contracts.MatchesID("ALC-001", c.Name, c.Classname) {
								referencedFailure = true
							}
						}
						if err == nil || !referencedFailure {
							t.Errorf("canary undetected: %v %s %+v", err, out, s.Cases)
						}
					case "missing_id":
						if err != nil || !s.hasUnlinked("ALC-001") {
							t.Errorf("missing ID undetected: %v %s %+v", err, out, s.Cases)
						}
					}
					w.remove(paths)
					s, out, err = w.hook(tc.hook, "alpha")
					if err != nil || (tc.hook == "eval" && (s.hasFailure() || s.hasUnlinked("ALC-001"))) {
						t.Errorf("recovery failed: %v %s %+v", err, out, s.Cases)
					}
				})
			}
		})
	}
}
