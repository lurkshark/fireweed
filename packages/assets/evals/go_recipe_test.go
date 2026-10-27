package evals_test

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/lurkshark/fireweed/packages/assets"
	"github.com/lurkshark/fireweed/packages/contracts"
	"gopkg.in/yaml.v3"
)

func goRecipe(t *testing.T) assets.Recipe {
	t.Helper()
	r, err := assets.BuiltinRecipe("go")
	if err != nil {
		t.Fatalf("BuiltinRecipe(go): %v", err)
	}
	return r
}

func Test_AST_004_GoRecipe(t *testing.T) {
	r := goRecipe(t)
	if r.Name != "go" {
		t.Errorf("recipe name = %q, want go", r.Name)
	}
	if _, err := assets.BuiltinRecipe("unknown-recipe"); err == nil {
		t.Error("unknown recipe name was accepted")
	}
	for label, patterns := range map[string][]string{
		"interface":  r.ModuleDefaults.Interface,
		"evals":      r.ModuleDefaults.Evals,
		"disposable": r.ModuleDefaults.Disposable,
	} {
		want := map[string]string{"interface": "api.go", "evals": "evals/**", "disposable": "internal/**"}[label]
		if !containsString(patterns, want) {
			t.Errorf("module default %s = %q, missing %q", label, patterns, want)
		}
	}
	for _, key := range []string{"eval", "boundaries", "setup"} {
		if s, ok := r.Hooks[key].(string); !ok || strings.TrimSpace(s) == "" {
			t.Errorf("missing %s hook command: %v", key, r.Hooks[key])
		}
	}
	if !containsString(r.WorkerWritable, "go.sum") {
		t.Errorf("go.sum absent from worker_writable: %q", r.WorkerWritable)
	}
	for _, path := range []string{
		".fireweed/recipes/go.md", ".fireweed/hooks.go",
		".fireweed/templates/module/SPEC.md", ".fireweed/templates/module/api.go.tmpl",
		".fireweed/canaries/canaries.yaml",
	} {
		if strings.TrimSpace(r.Files[path]) == "" {
			t.Errorf("missing project-relative recipe file %q", path)
		}
	}
	text := strings.ToLower(r.Instructions + "\n" + r.Files[".fireweed/recipes/go.md"])
	for _, term := range []string{"go.mod", "module", "workspace", "fireweed_bin", "modules --json"} {
		if !strings.Contains(text, term) {
			t.Errorf("Go instructions omit %q", term)
		}
	}
	for _, command := range []string{"go mod init", "go work init", "go work use"} {
		if !strings.Contains(text, command) {
			t.Errorf("Go instructions do not explain module/workspace setup with %q", command)
		}
	}
	graphProcedure := false
	for _, paragraph := range strings.Split(text, "\n\n") {
		if strings.Contains(paragraph, "fireweed_bin modules --json") &&
			strings.Contains(paragraph, "graph") &&
			strings.Contains(paragraph, "depends_on") &&
			strings.Contains(paragraph, "zones") &&
			(strings.Contains(paragraph, "derive") || strings.Contains(paragraph, "build") || strings.Contains(paragraph, "generate")) {
			graphProcedure = true
		}
	}
	if !graphProcedure {
		t.Error("Go instructions must explain deriving the graph's dependencies and zones from FIREWEED_BIN modules --json")
	}
}

func Test_AST_005_EvalHookResults(t *testing.T) {
	r := goRecipe(t)
	w := newGoRecipeWorkspace(t, r)
	w.write("packages/alpha/evals/results_test.go", `package evals
import "testing"
func Test_ALC001Pass(t *testing.T) { t.Log("passed") }
func Test_ALC001Fail(t *testing.T) { t.Fatal("deliberate failure") }
func Test_ALC001Skip(t *testing.T) { t.Skip("deliberate skip") }
`)
	report, output, err := w.eval("alpha")
	if err == nil {
		t.Fatalf("failing Go test returned success; output: %s", output)
	}
	for _, item := range []struct {
		name, state string
	}{
		{"Test_ALC001Pass", "pass"},
		{"Test_ALC001Fail", "fail"},
		{"Test_ALC001Skip", "skip"},
	} {
		c := report.caseNamed(item.name)
		if c == nil {
			t.Errorf("JUnit report lacks %s: %+v", item.name, report.Cases)
			continue
		}
		if !strings.Contains(c.Classname, "alpha/evals") {
			t.Errorf("%s classname = %q", item.name, c.Classname)
		}
		if c.Time == "" {
			t.Errorf("%s lacks elapsed seconds", item.name)
		} else if seconds, parseErr := strconv.ParseFloat(c.Time, 64); parseErr != nil || seconds < 0 {
			t.Errorf("%s has invalid elapsed seconds %q", item.name, c.Time)
		}
		if got := c.state(); got != item.state {
			t.Errorf("%s state = %s, want %s", item.name, got, item.state)
		}
	}
}

func Test_AST_005_UnfinishedAndCompileFailure(t *testing.T) {
	r := goRecipe(t)
	w := newGoRecipeWorkspace(t, r)
	w.write("packages/alpha/evals/unfinished_test.go", `package evals
import ("os"; "testing")
func Test_ALC001Unfinished(t *testing.T) { os.Exit(3) }
`)
	report, output, err := w.eval("alpha")
	if err == nil {
		t.Fatalf("aborted Go test returned success; output: %s", output)
	}
	if c := report.caseNamed("Test_ALC001Unfinished"); c == nil || c.state() != "fail" {
		t.Errorf("started test has no visible failure: %+v", report.Cases)
	}
	if err := os.Remove(filepath.Join(w.root, "packages/alpha/evals/unfinished_test.go")); err != nil {
		t.Fatal(err)
	}
	w.write("packages/alpha/evals/broken_test.go", "package evals\nfunc broken( {\n")
	report, output, err = w.eval("alpha")
	if err == nil {
		t.Fatalf("compile failure returned success; output: %s", output)
	}
	if len(report.Cases) == 0 || !report.hasFailure() {
		t.Errorf("compile failure left empty/passing JUnit evidence: %+v", report.Cases)
	}
}

func Test_AST_005_EvalOnlyRequestedModule(t *testing.T) {
	w := newGoRecipeWorkspace(t, goRecipe(t))
	w.write("packages/alpha/evals/alpha_test.go", `package evals
import "testing"
func Test_ALC001Selected(t *testing.T) {}
`)
	w.write("packages/beta/evals/beta_test.go", `package evals
import "testing"
func Test_BET001OtherModuleFails(t *testing.T) { t.Fatal("beta must be excluded from alpha eval") }
`)
	report, output, err := w.eval("alpha")
	if err != nil || report.caseNamed("Test_ALC001Selected") == nil {
		t.Fatalf("alpha eval did not run alpha tests: err=%v output=%q cases=%+v", err, output, report.Cases)
	}
	if c := report.caseNamed("Test_BET001OtherModuleFails"); c != nil {
		t.Errorf("alpha eval included beta testcase: %+v", *c)
	}
	if report, output, err := w.eval("beta"); err == nil || report.caseNamed("Test_BET001OtherModuleFails") == nil {
		t.Errorf("beta fixture did not fail when requested: err=%v output=%q cases=%+v", err, output, report.Cases)
	}
}

func Test_AST_006_BoundaryGraph(t *testing.T) {
	r := goRecipe(t)
	w := newGoRecipeWorkspace(t, r)
	w.setGraph(nil, "api.go", "evals/**")
	if output, err := w.boundaries(); err != nil {
		t.Fatalf("clean graph rejected: %v: %s", err, output)
	}
	if query, err := os.ReadFile(filepath.Join(w.root, "modules-query.log")); err != nil || !strings.Contains(string(query), "modules --json") {
		t.Fatalf("boundary hook did not query FIREWEED_BIN modules --json: %q, %v", query, err)
	}
	w.write("packages/alpha/foreign.go", `package alpha
import _ "example.com/recipefixture/packages/beta"
`)
	output, err := w.boundaries()
	if err == nil || !strings.Contains(output, "alpha:") || !strings.Contains(output, "beta") {
		t.Errorf("undeclared import must produce module-prefixed violation: err=%v output=%q", err, output)
	}
	w.remove("packages/alpha/foreign.go")
	w.setGraph([]string{"beta"}, "api.go", "evals/**")
	w.write("packages/alpha/foreign.go", `package alpha
import _ "example.com/recipefixture/packages/beta"
`)
	if output, err := w.boundaries(); err != nil {
		t.Errorf("declared public dependency rejected: %v: %s", err, output)
	}
}

func Test_AST_006_EvalInterfaceZones(t *testing.T) {
	r := goRecipe(t)
	w := newGoRecipeWorkspace(t, r)
	w.setGraph([]string{"beta"}, "public/**", "checks/**")
	w.write("packages/beta/public/public.go", "package public\n")
	w.write("packages/beta/details/details.go", "package details\n")
	w.write("packages/beta/internal/secret/secret.go", "package secret\n")
	w.write("packages/alpha/internal/secret/secret.go", "package secret\n")
	w.write("packages/alpha/checks/canary_test.go", `package checks
import _ "example.com/recipefixture/packages/beta/public"
`)
	if output, err := w.boundaries(); err != nil {
		t.Errorf("eval import of declared retained interface rejected: %v: %s", err, output)
	}
	for _, tc := range []struct{ name, path string }{
		{"dependency outside interface", "example.com/recipefixture/packages/beta/details"},
		{"dependency implementation", "example.com/recipefixture/packages/beta/internal/secret"},
		{"own implementation", "example.com/recipefixture/packages/alpha/internal/secret"},
	} {
		w.write("packages/alpha/checks/canary_test.go", "package checks\nimport _ \""+tc.path+"\"\n")
		output, err := w.boundaries()
		if err == nil || !strings.Contains(output, "alpha:") || !strings.Contains(output, tc.path) {
			t.Errorf("%s eval import escaped interface: err=%v output=%q", tc.name, err, output)
		}
	}
	w.remove("packages/alpha/checks/canary_test.go")
	if output, err := w.boundaries(); err != nil {
		t.Errorf("clean graph did not recover: %v: %s", err, output)
	}
}

func Test_AST_006_UsesModuleDirsAndWorkspacePath(t *testing.T) {
	w := newGoRecipeWorkspace(t, goRecipe(t))
	w.write("go.mod", "module example.net/customworkspace\n\ngo 1.23\n")
	w.remove("packages/alpha/api.go")
	w.remove("packages/beta/api.go")
	const alphaDir = "src/engine"
	const betaDir = "libraries/ledger"
	const betaImport = "example.net/customworkspace/libraries/ledger/public"
	w.write(alphaDir+"/api.go", "package engine\n")
	w.write(betaDir+"/public/api.go", "package public\n")
	w.write(betaDir+"/details/details.go", "package details\n")
	graph := func(deps []string) {
		modules := []map[string]any{
			{"name": "alpha", "dir": alphaDir, "depends_on": deps,
				"zones": map[string]any{"interface": []string{"api.go"}, "evals": []string{"checks/**"}, "disposable": []string{"internal/**"}}},
			{"name": "beta", "dir": betaDir, "depends_on": []string{},
				"zones": map[string]any{"interface": []string{"public/**"}, "evals": []string{"checks/**"}, "disposable": []string{"internal/**"}}},
		}
		b, err := json.Marshal(map[string]any{"data": modules})
		if err != nil {
			t.Fatal(err)
		}
		w.write("modules.json", string(b))
	}
	graph(nil)
	if output, err := w.boundaries(); err != nil {
		t.Fatalf("clean relocated modules rejected: %v: %s", err, output)
	}
	w.write(alphaDir+"/foreign.go", "package engine\nimport _ \""+betaImport+"\"\n")
	if output, err := w.boundaries(); err == nil || !strings.Contains(output, "alpha:") || !strings.Contains(output, "beta") {
		t.Errorf("undeclared relocated module import not reported: %v: %s", err, output)
	}
	graph([]string{"beta"})
	if output, err := w.boundaries(); err != nil {
		t.Errorf("declared relocated module import rejected: %v: %s", err, output)
	}
	w.write(alphaDir+"/checks/canary_test.go", "package checks\nimport _ \""+betaImport+"\"\n")
	if output, err := w.boundaries(); err != nil {
		t.Errorf("eval import of relocated retained interface rejected: %v: %s", err, output)
	}
	const detailsImport = "example.net/customworkspace/libraries/ledger/details"
	w.write(alphaDir+"/checks/canary_test.go", "package checks\nimport _ \""+detailsImport+"\"\n")
	if output, err := w.boundaries(); err == nil || !strings.Contains(output, "alpha:") || !strings.Contains(output, detailsImport) {
		t.Errorf("eval import outside relocated interface not reported: %v: %s", err, output)
	}
	w.remove(alphaDir + "/checks/canary_test.go")
	if output, err := w.boundaries(); err != nil {
		t.Errorf("relocated graph did not recover: %v: %s", err, output)
	}
}

func Test_AST_007_TemplatesAndCanaries(t *testing.T) {
	r := goRecipe(t)
	w := newGoRecipeWorkspace(t, r)
	for _, path := range []string{".fireweed/templates/module/SPEC.md", ".fireweed/templates/module/api.go.tmpl"} {
		content := r.Files[path]
		for _, field := range []string{"name", "prefix", "package"} {
			if path == ".fireweed/templates/module/SPEC.md" && field == "package" {
				continue
			}
			if path == ".fireweed/templates/module/api.go.tmpl" && field != "package" {
				continue
			}
			if !strings.Contains(content, "{{"+field+"}}") {
				t.Errorf("%s omits %s placeholder", path, field)
			}
		}
	}
	var canaries struct {
		Version      int               `yaml:"version"`
		Baseline     map[string]string `yaml:"baseline"`
		Undeclared   map[string]string `yaml:"undeclared"`
		InternalEval map[string]string `yaml:"internal_eval"`
		FailingEval  map[string]string `yaml:"failing_eval"`
		MissingID    map[string]string `yaml:"missing_id"`
	}
	if err := yaml.Unmarshal([]byte(r.Files[".fireweed/canaries/canaries.yaml"]), &canaries); err != nil {
		t.Fatalf("canary YAML: %v", err)
	}
	if canaries.Version != 1 {
		t.Errorf("canary map version = %d, want 1", canaries.Version)
	}
	for name, files := range map[string]map[string]string{
		"baseline": canaries.Baseline, "undeclared": canaries.Undeclared,
		"internal_eval": canaries.InternalEval, "failing_eval": canaries.FailingEval,
		"missing_id": canaries.MissingID,
	} {
		if len(files) == 0 {
			t.Errorf("%s canary has no files", name)
		}
	}
	if t.Failed() {
		return
	}
	for _, module := range []string{"alpha", "beta"} {
		w.applyCanary(canaries.Baseline, module)
	}
	for _, module := range []string{"alpha", "beta"} {
		report, output, err := w.eval(module)
		if err != nil || !report.hasPass() {
			t.Errorf("rendered %s baseline must compile and pass: %v: %s; cases=%+v", module, err, output, report.Cases)
		}
	}
	for path, content := range map[string]string{
		"packages/beta/SPEC.md": r.Files[".fireweed/templates/module/SPEC.md"],
		"packages/beta/api.go":  r.Files[".fireweed/templates/module/api.go.tmpl"],
	} {
		w.write(path, renderCanaryFields(content, "beta"))
	}
	if spec, err := os.ReadFile(filepath.Join(w.root, "packages/beta/SPEC.md")); err != nil ||
		!strings.Contains(string(spec), "name: beta") || !strings.Contains(string(spec), "id_prefix: BET") {
		t.Errorf("rendered SPEC template missing beta identity: %v: %s", err, spec)
	}
	if report, output, err := w.eval("beta"); err != nil || !report.hasPass() {
		t.Errorf("rendered api.go template must compile: %v: %s; cases=%+v", err, output, report.Cases)
	}
	w.setGraph(nil, "api.go", "evals/**")
	if output, err := w.boundaries(); err != nil {
		t.Errorf("rendered baseline boundaries: %v: %s", err, output)
	}
	for _, tc := range []struct {
		name  string
		files map[string]string
		hook  string
	}{
		{"undeclared", canaries.Undeclared, "boundaries"},
		{"internal_eval", canaries.InternalEval, "boundaries"},
		{"failing_eval", canaries.FailingEval, "eval"},
		{"missing_id", canaries.MissingID, "eval"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			paths := w.applyCanary(tc.files, "alpha")
			if tc.hook == "boundaries" {
				output, err := w.boundaries()
				if err == nil || !strings.Contains(output, "alpha:") {
					t.Errorf("%s canary not reported as alpha violation: %v: %s", tc.name, err, output)
				}
			} else {
				report, output, err := w.eval("alpha")
				if tc.name == "failing_eval" {
					if err == nil || !report.hasFailure() {
						t.Errorf("failing_eval canary not reported: %v: %s; cases=%+v", err, output, report.Cases)
					}
				} else {
					if err != nil || !report.hasUnlinked("ALC-001") {
						t.Errorf("missing_id canary not exposed as unlinked JUnit case: %v: %s; cases=%+v", err, output, report.Cases)
					}
				}
			}
			for _, path := range paths {
				w.remove(path)
			}
			if tc.hook == "boundaries" {
				if output, err := w.boundaries(); err != nil {
					t.Errorf("%s boundary recovery failed: %v: %s", tc.name, err, output)
				}
			} else if report, output, err := w.eval("alpha"); err != nil || report.hasFailure() || report.hasUnlinked("ALC-001") {
				t.Errorf("%s eval recovery failed: %v: %s; cases=%+v", tc.name, err, output, report.Cases)
			}
		})
	}
}

type goRecipeWorkspace struct {
	t      *testing.T
	root   string
	recipe assets.Recipe
}

func newGoRecipeWorkspace(t *testing.T, recipe assets.Recipe) *goRecipeWorkspace {
	t.Helper()
	w := &goRecipeWorkspace{t: t, root: t.TempDir(), recipe: recipe}
	w.write("go.mod", "module example.com/recipefixture\n\ngo 1.23\n")
	for path, content := range recipe.Files {
		w.write(path, content)
	}
	w.write("packages/alpha/api.go", "package alpha\n")
	w.write("packages/beta/api.go", "package beta\n")
	w.setGraph(nil, "api.go", "evals/**")
	return w
}

func (w *goRecipeWorkspace) write(path, data string) {
	w.t.Helper()
	abs := filepath.Join(w.root, path)
	if err := os.MkdirAll(filepath.Dir(abs), 0755); err != nil {
		w.t.Fatal(err)
	}
	if err := os.WriteFile(abs, []byte(data), 0644); err != nil {
		w.t.Fatal(err)
	}
}

func (w *goRecipeWorkspace) remove(path string) {
	w.t.Helper()
	if err := os.Remove(filepath.Join(w.root, path)); err != nil {
		w.t.Fatal(err)
	}
}

func (w *goRecipeWorkspace) setGraph(alphaDeps []string, betaInterface, evalZone string) {
	w.t.Helper()
	if alphaDeps == nil {
		alphaDeps = []string{}
	}
	modules := []map[string]any{
		{"name": "alpha", "dir": "packages/alpha", "depends_on": alphaDeps,
			"zones": map[string]any{"interface": []string{"api.go"}, "evals": []string{evalZone}, "disposable": []string{"internal/**"}}},
		{"name": "beta", "dir": "packages/beta", "depends_on": []string{},
			"zones": map[string]any{"interface": []string{betaInterface}, "evals": []string{evalZone}, "disposable": []string{"internal/**"}}},
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

func (w *goRecipeWorkspace) runHook(name, module, report string) (string, error) {
	w.t.Helper()
	command, ok := w.recipe.Hooks[name].(string)
	if !ok || strings.TrimSpace(command) == "" {
		w.t.Fatalf("recipe has no %s hook command", name)
	}
	for key, value := range map[string]string{
		"{module}": module, "{module_dir}": "packages/" + module, "{report}": report,
	} {
		command = strings.ReplaceAll(command, key, value)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, "sh", "-c", command)
	cmd.Dir = w.root
	cmd.Env = append(os.Environ(),
		"FIREWEED_BIN="+filepath.Join(w.root, "fake-fireweed"),
		"FIREWEED_MODULES_JSON="+filepath.Join(w.root, "modules.json"),
		"FIREWEED_QUERY_LOG="+filepath.Join(w.root, "modules-query.log"),
		"GOCACHE=/tmp/fireweed-rebuild-cache", "GOWORK=off", "GOPROXY=off", "GOSUMDB=off")
	output, err := cmd.CombinedOutput()
	if ctx.Err() != nil {
		w.t.Fatalf("%s hook exceeded 30 seconds: %v; output: %s", name, ctx.Err(), output)
	}
	return string(output), err
}

func (w *goRecipeWorkspace) eval(module string) (junitSuite, string, error) {
	w.t.Helper()
	report := filepath.Join(w.root, fmt.Sprintf("%s-report.xml", module))
	_ = os.Remove(report)
	output, err := w.runHook("eval", module, report)
	b, readErr := os.ReadFile(report)
	if readErr != nil {
		w.t.Fatalf("eval hook did not write JUnit report: %v; exit=%v; output=%s", readErr, err, output)
	}
	var suite junitSuite
	if parseErr := xml.Unmarshal(b, &suite); parseErr != nil {
		w.t.Fatalf("invalid JUnit XML: %v; report=%s", parseErr, b)
	}
	return suite, output, err
}

func (w *goRecipeWorkspace) boundaries() (string, error) {
	return w.runHook("boundaries", "alpha", "")
}

func (w *goRecipeWorkspace) applyCanary(files map[string]string, module string) []string {
	w.t.Helper()
	paths := make([]string, 0, len(files))
	for path, content := range files {
		content = renderCanaryFields(content, module)
		if strings.Contains(content, "{{") {
			w.t.Errorf("unrendered canary placeholder in %s: %s", path, content)
		}
		full := filepath.Join("packages", module, path)
		w.write(full, content)
		paths = append(paths, full)
	}
	return paths
}

func renderCanaryFields(content, module string) string {
	fields := map[string]string{
		"name": module, "prefix": map[string]string{"alpha": "ALC", "beta": "BET"}[module],
		"package": module, "go_module": "example.com/recipefixture",
		"module_dir": "packages/" + module, "other_dir": "packages/beta",
	}
	for key, value := range fields {
		content = strings.ReplaceAll(content, "{{"+key+"}}", value)
	}
	return content
}

type junitSuite struct {
	Cases []junitCase `xml:"testcase"`
}
type junitCase struct {
	Name      string    `xml:"name,attr"`
	Classname string    `xml:"classname,attr"`
	Time      string    `xml:"time,attr"`
	Failure   *struct{} `xml:"failure"`
	Skipped   *struct{} `xml:"skipped"`
}

func (s junitSuite) caseNamed(name string) *junitCase {
	for i := range s.Cases {
		if s.Cases[i].Name == name {
			return &s.Cases[i]
		}
	}
	return nil
}
func (s junitSuite) hasFailure() bool {
	for _, c := range s.Cases {
		if c.state() == "fail" {
			return true
		}
	}
	return false
}
func (s junitSuite) hasPass() bool {
	for _, c := range s.Cases {
		if c.state() == "pass" {
			return true
		}
	}
	return false
}
func (s junitSuite) hasUnlinked(id string) bool {
	for _, c := range s.Cases {
		if c.state() == "pass" && !contracts.MatchesID(id, c.Name, c.Classname) {
			return true
		}
	}
	return false
}
func (c junitCase) state() string {
	if c.Failure != nil {
		return "fail"
	}
	if c.Skipped != nil {
		return "skip"
	}
	return "pass"
}
func containsString(items []string, target string) bool {
	for _, item := range items {
		if item == target {
			return true
		}
	}
	return false
}
