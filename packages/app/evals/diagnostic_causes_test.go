package evals_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func appCause(t *testing.T, bin, root string, args []string, wantStatus int, code, cause string) {
	t.Helper()
	for _, jsonMode := range []bool{false, true} {
		call := append([]string(nil), args...)
		if jsonMode {
			call = append(call, "--json")
		}
		status, stdout, stderr := runAppBinary(t, bin, root, call...)
		commandFailure(t, status, wantStatus, stdout, stderr, jsonMode)
		if jsonMode {
			if stderr != "" {
				t.Errorf("JSON stderr=%q; want empty", stderr)
			}
			envelope := decodeEnvelope(t, stdout)
			if string(envelope["data"]) != "null" {
				t.Errorf("failure data=%s; want null", envelope["data"])
			}
			var diagnostics []struct {
				Code    string `json:"code"`
				Message string `json:"message"`
			}
			if err := json.Unmarshal(envelope["diagnostics"], &diagnostics); err != nil || len(diagnostics) == 0 {
				t.Fatalf("diagnostics=%s: %v", envelope["diagnostics"], err)
			}
			found := false
			for _, diagnostic := range diagnostics {
				if diagnostic.Code == code && strings.Contains(diagnostic.Message, cause) {
					found = true
				}
			}
			if !found {
				t.Errorf("missing %s cause %q in %s", code, cause, stdout)
			}
		} else if !strings.Contains(stdout+stderr, code+": ") || !strings.Contains(stdout+stderr, cause) {
			t.Errorf("missing %s cause %q in plain output %q", code, cause, stdout+stderr)
		} else if strings.Contains(stdout+stderr, code+": "+code+":") {
			t.Errorf("duplicate %s prefix in %q", code, stdout+stderr)
		}
	}
}

func Test_APP_037_StaticCauses(t *testing.T) {
	bin := buildAppBinary(t)
	for _, tc := range []struct {
		name, code, cause string
		args              []string
		status            int
		edit              func(*testing.T, string)
	}{
		{"configuration", "C01", "fireweed.yaml", []string{"modules"}, 2, func(t *testing.T, root string) {
			appPut(t, root, "fireweed.yaml", "modules: [broken\n")
		}},
		{"graph", "C02", "missing module", []string{"modules"}, 2, func(t *testing.T, root string) {
			appPut(t, root, "packages/beta/SPEC.md", "---\nname: beta\nid_prefix: BET\ndepends_on: [missing]\n---\n# Beta\n\n## Specs\n")
		}},
		{"syntax", "C03", "malformed item heading", []string{"coverage", "alpha"}, 2, func(t *testing.T, root string) {
			appPut(t, root, "packages/alpha/SPEC.md", strings.Replace(authoringSpec, "### ALP-001", "### broken item", 1))
		}},
		{"normative", "C04", "non-normative", []string{"check"}, 1, func(t *testing.T, root string) {
			authoringChange(t, root, "0001-first.md", "Users need a statement.\n", "### new:first\nA statement without a normative verb.\n", "")
		}},
		{"identity", "C05", "duplicate item ID", []string{"check"}, 1, func(t *testing.T, root string) {
			authoringChange(t, root, "0001-first.md", "Users need two items.\n", "### new:one\nOne MUST work.\n### new:one\nOne MUST remain unique.\n", "")
		}},
		{"tag", "C06", "invalid review/quarantine tag", []string{"check"}, 1, func(t *testing.T, root string) {
			authoringChange(t, root, "0001-first.md", "Users need an item.\n", "### new:one [invalid]\nOne MUST work.\n", "")
		}},
		{"structure", "C07", "missing Motivation", []string{"check"}, 1, func(t *testing.T, root string) {
			authoringChange(t, root, "0001-first.md", "", "### new:one\nOne MUST work.\n", "")
		}},
		{"base", "C08", "absent from preceding projection", []string{"stamp", "revise"}, 1, func(t *testing.T, root string) {
			authoringChange(t, root, "0001-revise.md", "Users need a revision.\n", "", "### ALP-999\nMissing MUST change.\n")
		}},
		{"handle", "C09", "undeclared handle", []string{"check"}, 1, func(t *testing.T, root string) {
			authoringChange(t, root, "0001-first.md", "Users need new:missing.\n", "### new:one\nOne MUST work.\n", "")
		}},
		{"metadata", "C12", "set must be a mapping", []string{"show", "alpha", "--projected"}, 2, func(t *testing.T, root string) {
			authoringChange(t, root, "0001-first.md", "Users need a change.\n", "", "")
			path := filepath.Join(root, "changes/0001-first.md")
			body, err := os.ReadFile(path)
			if err != nil {
				t.Fatal(err)
			}
			appPut(t, root, "changes/0001-first.md", strings.Replace(string(body), "module: alpha\n", "module: alpha\nset: invalid\n", 1))
		}},
		{"queue-identity", "C13", "ambiguous change slug", []string{"check"}, 1, func(t *testing.T, root string) {
			authoringChange(t, root, "0001-repeat.md", "Users need one item.\n", "### new:one\nOne MUST work.\n", "")
			authoringChange(t, root, "0002-repeat.md", "Users need another item.\n", "### new:two\nTwo MUST work.\n", "")
		}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := authoringFixture(t)
			tc.edit(t, root)
			appCause(t, bin, root, tc.args, tc.status, tc.code, tc.cause)
		})
	}
}

func Test_APP_037_APP_038_ShowStaticCauses(t *testing.T) {
	bin := buildAppBinary(t)
	for _, tc := range []struct {
		name, code, cause string
		args              []string
		edit              func(*testing.T, string)
	}{
		{"duplicate-graph-name", "C02", "duplicate module name", []string{"show", "alpha"}, func(t *testing.T, root string) {
			appPut(t, root, "packages/beta/SPEC.md", "---\nname: alpha\nid_prefix: BET\ndepends_on: []\n---\n# Beta\n\n## Specs\n")
		}},
		{"duplicate-fixed-spec-id", "C05", "duplicate item ID", []string{"show", "alpha", "ALP-002"}, func(t *testing.T, root string) {
			appPut(t, root, "packages/alpha/SPEC.md", authoringSpec+"\n### ALP-001\nDuplicate MUST fail.\n\n### ALP-002\nDistinct MUST remain selectable.\n")
		}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := authoringFixture(t)
			tc.edit(t, root)
			appCause(t, bin, root, tc.args, 2, tc.code, tc.cause)
		})
	}
}

func Test_APP_038_LookupCauses(t *testing.T) {
	bin := buildAppBinary(t)
	for _, tc := range []struct {
		name, cause string
		args        []string
		status      int
		edit        func(*testing.T, string)
	}{
		{"author-module", "unknown module", []string{"new-change", "unknown", "valid"}, 2, nil},
		{"stamp-change", "not found", []string{"stamp", "absent"}, 1, nil},
		{"show-module", "unknown module", []string{"show", "absent"}, 2, nil},
		{"show-item", "unknown item", []string{"show", "alpha", "ALP-999"}, 2, nil},
		{"show-ambiguous-item", "duplicate item ID", []string{"show", "alpha", "new:shared", "--projected"}, 2, func(t *testing.T, root string) {
			authoringChange(t, root, "0001-first.md", "First need.\n", "### new:shared\nFirst MUST work.\n", "")
			authoringChange(t, root, "0002-second.md", "Second need.\n", "### new:shared\nSecond MUST work.\n", "")
		}},
		{"coverage-module", "Unknown module", []string{"coverage", "missing"}, 2, nil},
		{"coverage-change", "not found", []string{"coverage", "alpha", "--with-change", "missing"}, 2, nil},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := authoringFixture(t)
			if tc.edit != nil {
				tc.edit(t, root)
			}
			appCause(t, bin, root, tc.args, tc.status, "U02", tc.cause)
		})
	}
}

func Test_APP_039_BlockedNewChangeOutputPath(t *testing.T) {
	bin := buildAppBinary(t)
	for _, jsonMode := range []bool{false, true} {
		root := appProjectFixture(t)
		appPut(t, root, "changes/0002-keep.md", "unchanged marker")
		blocked := filepath.Join(root, "changes/0003-blocked.md")
		if err := os.Mkdir(blocked, 0755); err != nil {
			t.Fatal(err)
		}
		before := changeFiles(t, root)
		args := []string{"new-change", "alpha", "blocked"}
		if jsonMode {
			args = append(args, "--json")
		}
		status, stdout, stderr := runAppBinary(t, bin, root, args...)
		message := commandFailure(t, status, 2, stdout, stderr, jsonMode)
		if jsonMode {
			assertFailedJSON(t, stdout, "P01")
		} else if !strings.HasPrefix(strings.TrimSpace(stdout+stderr), "P01:") {
			t.Errorf("plain diagnostic=%q; want P01", stdout+stderr)
		}
		if !strings.Contains(message, "create changes/0003-blocked.md") || !corrective(message) {
			t.Errorf("lost create-file cause or correction: %q", message)
		}
		if after := changeFiles(t, root); !reflect.DeepEqual(after, before) {
			t.Errorf("changes modified: before=%v after=%v", before, after)
		}
		if info, err := os.Stat(blocked); err != nil || info == nil || !info.IsDir() {
			t.Errorf("blocked directory changed: %v", err)
		}
	}
}

func Test_APP_040_ForwardedCheckCode(t *testing.T) {
	bin := buildAppBinary(t)
	root := authoringFixture(t)
	appCause(t, bin, root, []string{"coverage", "alpha"}, 1, "C11", "missing eval hook")
}
