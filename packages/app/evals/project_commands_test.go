package evals_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/contracts"
)

const appFixtureConfig = `fireweed: ">=0.1, <0.2"
version: 1
base: main
modules: ["packages/*"]
module_defaults:
  interface: [api.go]
  evals: [evals/**]
  disposable: [internal/**]
  preserve: [data/**]
`

func appPut(t *testing.T, root, name, body string) {
	t.Helper()
	path := filepath.Join(root, filepath.FromSlash(name))
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(body), 0644); err != nil {
		t.Fatal(err)
	}
}

func appProjectFixture(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	appPut(t, root, "fireweed.yaml", appFixtureConfig)
	appPut(t, root, "packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\nevals: []\n---\n# Alpha\n\n## Specs\n")
	appPut(t, root, "packages/beta/SPEC.md", "---\nname: beta\nid_prefix: BET\ndepends_on: [alpha]\n---\n# Beta\n\n## Specs\n")
	return root
}

func appNestedDir(t *testing.T, root string) string {
	t.Helper()
	dir := filepath.Join(root, "packages", "alpha", "deep")
	if err := os.MkdirAll(dir, 0755); err != nil {
		t.Fatal(err)
	}
	return dir
}

func runAppBinary(t *testing.T, bin, dir string, args ...string) (int, string, string) {
	t.Helper()
	cmd := exec.Command(bin, args...)
	cmd.Dir = dir
	var stdout, stderr bytes.Buffer
	cmd.Stdout, cmd.Stderr = &stdout, &stderr
	err := cmd.Run()
	if err == nil {
		return 0, stdout.String(), stderr.String()
	}
	var exit *exec.ExitError
	if !errors.As(err, &exit) {
		t.Fatalf("run NEW binary %v: %v", args, err)
	}
	return exit.ExitCode(), stdout.String(), stderr.String()
}

func assertModuleJSON(t *testing.T, output string, names []string) {
	t.Helper()
	envelope := decodeEnvelope(t, output)
	var diagnostics []json.RawMessage
	if err := json.Unmarshal(envelope["diagnostics"], &diagnostics); err != nil || diagnostics == nil || len(diagnostics) != 0 {
		t.Errorf("diagnostics=%s; want empty array", envelope["diagnostics"])
	}
	var modules []struct {
		Name      string   `json:"name"`
		IDPrefix  string   `json:"id_prefix"`
		Dir       string   `json:"dir"`
		DependsOn []string `json:"depends_on"`
		Zones     struct {
			Interface  []string `json:"interface"`
			Evals      []string `json:"evals"`
			Disposable []string `json:"disposable"`
			Preserve   []string `json:"preserve"`
		} `json:"zones"`
	}
	if err := json.Unmarshal(envelope["data"], &modules); err != nil || modules == nil {
		t.Fatalf("data=%s; want module array: %v", envelope["data"], err)
	}
	var rawModules []map[string]json.RawMessage
	if err := json.Unmarshal(envelope["data"], &rawModules); err != nil {
		t.Fatalf("decode module fields: %v", err)
	}
	got := map[string]bool{}
	for i, module := range modules {
		for _, key := range []string{"name", "id_prefix", "dir", "depends_on", "zones"} {
			if _, ok := rawModules[i][key]; !ok {
				t.Errorf("module %q missing %s field", module.Name, key)
			}
		}
		var zoneFields map[string]json.RawMessage
		if err := json.Unmarshal(rawModules[i]["zones"], &zoneFields); err != nil {
			t.Errorf("module %q zones=%s: %v", module.Name, rawModules[i]["zones"], err)
		}
		for _, key := range []string{"protected", "evals", "interface", "append_only", "preserve", "disposable"} {
			if _, ok := zoneFields[key]; !ok {
				t.Errorf("module %q missing resolved %s zone", module.Name, key)
			}
		}
		got[module.Name] = true
		switch module.Name {
		case "alpha":
			if module.IDPrefix != "ALP" || module.Dir != "packages/alpha" || len(module.DependsOn) != 0 || module.Zones.Evals == nil || len(module.Zones.Evals) != 0 {
				t.Errorf("alpha fields/zones = %+v", module)
			}
		case "beta":
			if module.IDPrefix != "BET" || module.Dir != "packages/beta" || !reflect.DeepEqual(module.DependsOn, []string{"alpha"}) || !reflect.DeepEqual(module.Zones.Evals, []string{"evals/**"}) {
				t.Errorf("beta fields/zones = %+v", module)
			}
		case "gamma":
			if module.IDPrefix != "GAM" || module.Dir != "packages/gamma" || len(module.DependsOn) != 0 {
				t.Errorf("gamma fields = %+v", module)
			}
		default:
			t.Errorf("unexpected module %+v", module)
		}
		if !reflect.DeepEqual(module.Zones.Interface, []string{"api.go"}) || !reflect.DeepEqual(module.Zones.Disposable, []string{"internal/**"}) || !reflect.DeepEqual(module.Zones.Preserve, []string{"data/**"}) {
			t.Errorf("unresolved default zones for %s: %+v", module.Name, module.Zones)
		}
	}
	if len(modules) != len(names) {
		t.Errorf("modules=%d; want %d", len(modules), len(names))
	}
	for _, name := range names {
		if !got[name] {
			t.Errorf("missing module %q: %s", name, envelope["data"])
		}
	}
}

func Test_APP_004_APP_016_APP_017_APP_018_ModulesGraphAndNearestAncestor(t *testing.T) {
	bin := buildAppBinary(t)
	root := appProjectFixture(t)
	nested := appNestedDir(t, root)
	t.Run("plain", func(t *testing.T) {
		status, stdout, stderr := runAppBinary(t, bin, nested, "modules")
		if status != 0 || stderr != "" {
			t.Errorf("status=%d stderr=%q; want success", status, stderr)
		}
		for _, value := range []string{"alpha", "ALP", "packages/alpha", "beta", "BET", "packages/beta", "depends", "api.go", "evals/**", "internal/**", "data/**"} {
			if !strings.Contains(stdout, value) {
				t.Errorf("plain graph %q omits %q", stdout, value)
			}
		}
	})
	t.Run("JSON", func(t *testing.T) {
		status, stdout, stderr := runAppBinary(t, bin, nested, "modules", "--json")
		if status != 0 || stderr != "" {
			t.Errorf("status=%d stderr=%q; want success", status, stderr)
		}
		assertModuleJSON(t, stdout, []string{"alpha", "beta"})
	})
	inner := filepath.Join(root, "workspace", "inner")
	appPut(t, inner, "fireweed.yaml", appFixtureConfig)
	appPut(t, inner, "packages/gamma/SPEC.md", "---\nname: gamma\nid_prefix: GAM\ndepends_on: []\n---\n# Gamma\n\n## Specs\n")
	deep := filepath.Join(inner, "packages", "gamma", "deep")
	if err := os.MkdirAll(deep, 0755); err != nil {
		t.Fatal(err)
	}
	t.Run("nearest-ancestor", func(t *testing.T) {
		status, stdout, stderr := runAppBinary(t, bin, deep, "modules")
		if status != 0 || stderr != "" || !strings.Contains(stdout, "gamma") || strings.Contains(stdout, "alpha") || strings.Contains(stdout, "beta") {
			t.Errorf("nearest config: status=%d stdout=%q stderr=%q", status, stdout, stderr)
		}
		status, stdout, stderr = runAppBinary(t, bin, deep, "modules", "--json")
		if status != 0 || stderr != "" {
			t.Errorf("nearest JSON: status=%d stderr=%q", status, stderr)
		}
		assertModuleJSON(t, stdout, []string{"gamma"})
	})
}

func Test_APP_018_ModulesConfigurationErrors(t *testing.T) {
	bin := buildAppBinary(t)
	for _, tc := range []struct {
		name string
		edit func(*testing.T, string)
	}{
		{"missing", func(t *testing.T, root string) {
			if err := os.Remove(filepath.Join(root, "fireweed.yaml")); err != nil {
				t.Fatal(err)
			}
		}},
		{"malformed", func(t *testing.T, root string) { appPut(t, root, "fireweed.yaml", "modules: [broken\n") }},
		{"invalid-graph", func(t *testing.T, root string) {
			appPut(t, root, "packages/beta/SPEC.md", "---\nname: beta\nid_prefix: BET\ndepends_on: [missing]\n---\n# Beta\n\n## Specs\n")
		}},
	} {
		for _, jsonMode := range []bool{false, true} {
			format := "plain"
			if jsonMode {
				format = "JSON"
			}
			t.Run(tc.name+"/"+format, func(t *testing.T) {
				root := appProjectFixture(t)
				tc.edit(t, root)
				args := []string{"modules"}
				if jsonMode {
					args = append(args, "--json")
				}
				status, stdout, stderr := runAppBinary(t, bin, appNestedDir(t, root), args...)
				if status != 2 {
					t.Errorf("status=%d; want 2", status)
				}
				if jsonMode {
					if stderr != "" {
						t.Errorf("JSON stderr=%q; want empty", stderr)
					}
					diagnostic := assertFailedJSON(t, stdout)
					var result struct {
						Diagnostics []struct {
							Code string `json:"code"`
						} `json:"diagnostics"`
					}
					if err := json.Unmarshal([]byte(stdout), &result); err != nil || len(result.Diagnostics) == 0 || result.Diagnostics[0].Code != "C01" {
						t.Errorf("JSON diagnostic needs C01: %s", stdout)
					}
					if !corrective(diagnostic) {
						t.Errorf("diagnostic %q needs correction", diagnostic)
					}
				} else if !strings.Contains(stdout+stderr, "C01") || !corrective(stdout+stderr) {
					t.Errorf("plain diagnostic=%q; want corrective C01", stdout+stderr)
				}
			})
		}
	}
}

func assertChangeSkeleton(t *testing.T, root, rel, module string) {
	t.Helper()
	if filepath.IsAbs(rel) || !strings.HasPrefix(rel, "changes/") {
		t.Errorf("path %q is not repository-relative change path", rel)
	}
	data, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(rel)))
	if err != nil {
		t.Fatalf("read queued change %q: %v", rel, err)
	}
	doc, err := contracts.ParseDocument(string(data))
	if err != nil {
		t.Fatalf("parse queued change: %v", err)
	}
	if doc.Frontmatter["module"] != module {
		t.Errorf("module=%#v; want %q", doc.Frontmatter["module"], module)
	}
	for _, key := range []string{"phase", "run", "blocked"} {
		if value, present := doc.Frontmatter[key]; present {
			t.Errorf("queued change has %s=%#v", key, value)
		}
	}
	for _, heading := range []string{"Motivation", "Interface", "Added", "Modified", "Removed", "Decisions"} {
		if _, ok := contracts.Section(doc.Body, heading); !ok {
			t.Errorf("missing %s section", heading)
		}
	}
	if regexp.MustCompile(`(?m)^### [A-Z]+-[0-9]+\b`).Match(data) {
		t.Errorf("queued skeleton assigned a spec ID: %s", data)
	}
}

func Test_APP_005_APP_020_NewChangePlainAndJSON(t *testing.T) {
	bin := buildAppBinary(t)
	for _, jsonMode := range []bool{false, true} {
		{
			root := appProjectFixture(t)
			appPut(t, root, "changes/0003-waiting.md", "pending marker")
			appPut(t, root, "changes/archive/0007-finished.md", "archived marker")
			args := []string{"new-change", "alpha", "new-feature"}
			if jsonMode {
				args = append(args, "--json")
			}
			status, stdout, stderr := runAppBinary(t, bin, appNestedDir(t, root), args...)
			if status != 0 || stderr != "" {
				t.Errorf("status=%d stderr=%q; want success", status, stderr)
			}
			var rel string
			if jsonMode {
				envelope := decodeEnvelope(t, stdout)
				if err := json.Unmarshal(envelope["data"], &rel); err != nil {
					t.Fatalf("path data=%s: %v", envelope["data"], err)
				}
				var diagnostics []json.RawMessage
				if err := json.Unmarshal(envelope["diagnostics"], &diagnostics); err != nil || diagnostics == nil || len(diagnostics) != 0 {
					t.Errorf("diagnostics=%s; want empty array", envelope["diagnostics"])
				}
			} else {
				rel = strings.TrimSpace(stdout)
			}
			if rel != "changes/0008-new-feature.md" {
				t.Errorf("path=%q; want changes/0008-new-feature.md", rel)
			}
			assertChangeSkeleton(t, root, rel, "alpha")
		}
	}
}

func changeFiles(t *testing.T, root string) map[string]string {
	t.Helper()
	files := map[string]string{}
	changes := filepath.Join(root, "changes")
	err := filepath.WalkDir(changes, func(path string, entry os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() {
			return nil
		}
		body, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(changes, path)
		if err != nil {
			return err
		}
		files[filepath.ToSlash(rel)] = string(body)
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	return files
}

func Test_APP_019_APP_020_NewChangeErrorsPreserveFiles(t *testing.T) {
	bin := buildAppBinary(t)
	for _, tc := range []struct {
		name string
		args []string
		seed string
	}{
		{"missing-argument", []string{"new-change", "alpha"}, ""},
		{"extra-argument", []string{"new-change", "alpha", "valid", "extra"}, ""},
		{"invalid-module", []string{"new-change", "unknown", "valid"}, ""},
		{"invalid-slug", []string{"new-change", "alpha", "Bad"}, ""},
		{"duplicate-pending", []string{"new-change", "alpha", "repeat"}, "changes/0003-repeat.md"},
		{"duplicate-archived", []string{"new-change", "alpha", "repeat"}, "changes/archive/0003-repeat.md"},
	} {
		for _, jsonMode := range []bool{false, true} {
			format := "plain"
			if jsonMode {
				format = "JSON"
			}
			t.Run(tc.name+"/"+format, func(t *testing.T) {
				root := appProjectFixture(t)
				appPut(t, root, "changes/0002-keep.md", "unchanged marker")
				if tc.seed != "" {
					appPut(t, root, tc.seed, "duplicate marker")
				}
				before := changeFiles(t, root)
				args := append([]string(nil), tc.args...)
				if jsonMode {
					args = append(args, "--json")
				}
				status, stdout, stderr := runAppBinary(t, bin, appNestedDir(t, root), args...)
				if status != 2 {
					t.Errorf("status=%d; want 2", status)
				}
				if jsonMode {
					if stderr != "" {
						t.Errorf("JSON stderr=%q; want empty", stderr)
					}
					if diagnostic := assertFailedJSON(t, stdout); !corrective(diagnostic) {
						t.Errorf("diagnostic %q needs correction", diagnostic)
					}
				} else if !corrective(stdout + stderr) {
					t.Errorf("diagnostic=%q; want correction", stdout+stderr)
				}
				if after := changeFiles(t, root); !reflect.DeepEqual(after, before) {
					t.Errorf("changes modified: before=%v after=%v", before, after)
				}
			})
		}
	}
}
