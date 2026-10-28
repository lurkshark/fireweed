package evals_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/contracts"
)

const authoringSpec = "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\nevals: []\n---\n# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST keep its current text.\n"

func authoringFixture(t *testing.T) string {
	t.Helper()
	root := appProjectFixture(t)
	appPut(t, root, "packages/alpha/SPEC.md", authoringSpec)
	return root
}

func authoringChange(t *testing.T, root, file, motivation, added, modified string) {
	t.Helper()
	appPut(t, root, "changes/"+file, "---\nmodule: alpha\n---\n\n## Motivation\n"+motivation+"\n## Interface\n\n## Added\n"+added+"\n## Modified\n"+modified+"\n## Removed\n\n## Decisions\n")
}

func commandStringData(t *testing.T, output string) string {
	t.Helper()
	envelope := decodeEnvelope(t, output)
	var data string
	if err := json.Unmarshal(envelope["data"], &data); err != nil {
		t.Fatalf("data=%s; want string: %v", envelope["data"], err)
	}
	var diagnostics []json.RawMessage
	if err := json.Unmarshal(envelope["diagnostics"], &diagnostics); err != nil || diagnostics == nil || len(diagnostics) != 0 {
		t.Errorf("diagnostics=%s; want empty array", envelope["diagnostics"])
	}
	return data
}

func commandFailure(t *testing.T, status, want int, stdout, stderr string, jsonMode bool) string {
	t.Helper()
	if status != want {
		t.Errorf("status=%d; want %d; stdout=%q stderr=%q", status, want, stdout, stderr)
	}
	if jsonMode {
		if stderr != "" {
			t.Errorf("JSON stderr=%q; want empty", stderr)
		}
		message := assertFailedJSON(t, stdout)
		if strings.Contains(strings.ToLower(message), "under construction") || strings.Contains(strings.ToLower(message), "not implemented") {
			t.Errorf("command returned an unimplemented diagnostic: %q", message)
		}
		return message
	}
	if !strings.Contains(stdout+stderr, ":") || strings.TrimSpace(stdout+stderr) == "" {
		t.Errorf("plain diagnostic=%q; want code and message", stdout+stderr)
	}
	if strings.Contains(strings.ToLower(stdout+stderr), "under construction") || strings.Contains(strings.ToLower(stdout+stderr), "not implemented") {
		t.Errorf("command returned an unimplemented diagnostic: %q", stdout+stderr)
	}
	return stdout + stderr
}

func Test_APP_006_APP_022_StampAliasesAndUnassignedHandles(t *testing.T) {
	bin := buildAppBinary(t)
	for _, alias := range []string{"revise", "0001-revise.md", "changes/0001-revise.md"} {
		t.Run(alias, func(t *testing.T) {
			root := authoringFixture(t)
			authoringChange(t, root, "0001-revise.md", "Users need a revision.\n", "### new:local\nLocal MUST stay unassigned.\n", "### ALP-001\nAlpha MUST use revised text.\n")
			status, stdout, stderr := runAppBinary(t, bin, appNestedDir(t, root), "stamp", alias)
			if status != 0 || stderr != "" || strings.TrimSpace(stdout) == "" || !strings.Contains(stdout, "revise") {
				t.Errorf("stamp %q: status=%d stdout=%q stderr=%q; want confirmation", alias, status, stdout, stderr)
			}
			body, err := os.ReadFile(filepath.Join(root, "changes/0001-revise.md"))
			if err != nil {
				t.Fatal(err)
			}
			doc, err := contracts.ParseDocument(string(body))
			if err != nil {
				t.Fatal(err)
			}
			bases, ok := doc.Frontmatter["bases"].(map[string]any)
			want := contracts.SpecHash(contracts.Item{ID: "ALP-001", Body: "Alpha MUST keep its current text.\n"})
			if !ok || bases["ALP-001"] != want || len(bases) != 1 {
				t.Errorf("bases=%#v; want only ALP-001=%s", doc.Frontmatter["bases"], want)
			}
			if !strings.Contains(doc.Body, "### new:local\n") || strings.Contains(doc.Body, "### ALP-002\n") {
				t.Errorf("stamp assigned Added handle: %s", doc.Body)
			}
		})
	}
	root := authoringFixture(t)
	authoringChange(t, root, "0001-added.md", "Users need an addition.\n", "### new:only\nOnly MUST work.\n", "")
	status, stdout, stderr := runAppBinary(t, bin, appNestedDir(t, root), "stamp", "added", "--json")
	if status != 0 || stderr != "" || !strings.Contains(commandStringData(t, stdout), "added") {
		t.Errorf("Added-only JSON stamp: status=%d stdout=%q stderr=%q", status, stdout, stderr)
	}
	body, err := os.ReadFile(filepath.Join(root, "changes/0001-added.md"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(body), "### new:only\n") || strings.Contains(string(body), "### ALP-002\n") {
		t.Errorf("Added-only stamp assigned an ID: %s", body)
	}
}

func Test_APP_021_APP_022_StampErrorExits(t *testing.T) {
	bin := buildAppBinary(t)
	for _, tc := range []struct {
		name string
		args []string
		want int
		edit func(*testing.T, string)
	}{
		{"missing-argument", []string{"stamp"}, 2, nil},
		{"extra-argument", []string{"stamp", "revise", "extra"}, 2, nil},
		{"invalid-config", []string{"stamp", "revise"}, 2, func(t *testing.T, root string) { appPut(t, root, "fireweed.yaml", "modules: [broken\n") }},
		{"missing-change", []string{"stamp", "absent"}, 1, nil},
		{"absent-modified-ID", []string{"stamp", "revise"}, 1, func(t *testing.T, root string) {
			authoringChange(t, root, "0001-revise.md", "Users need a revision.\n", "", "### ALP-999\nMissing MUST change.\n")
		}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := authoringFixture(t)
			if tc.edit != nil {
				tc.edit(t, root)
			}
			for _, jsonMode := range []bool{false, true} {
				args := append([]string{}, tc.args...)
				if jsonMode {
					args = append(args, "--json")
				}
				status, stdout, stderr := runAppBinary(t, bin, appNestedDir(t, root), args...)
				commandFailure(t, status, tc.want, stdout, stderr, jsonMode)
			}
		})
	}
}

func Test_APP_007_APP_024_CheckValidAndAllDiagnostics(t *testing.T) {
	bin := buildAppBinary(t)
	root := authoringFixture(t)
	authoringChange(t, root, "0001-valid.md", "Users need this addition.\n", "### new:valid\nValid MUST work.\n", "")
	for _, jsonMode := range []bool{false, true} {
		args := []string{"check"}
		if jsonMode {
			args = append(args, "--json")
		}
		status, stdout, stderr := runAppBinary(t, bin, appNestedDir(t, root), args...)
		if status != 0 || stderr != "" {
			t.Errorf("valid check: status=%d stdout=%q stderr=%q", status, stdout, stderr)
		}
		if jsonMode {
			envelope := decodeEnvelope(t, stdout)
			var summary string
			if err := json.Unmarshal(envelope["data"], &summary); err != nil || strings.TrimSpace(summary) == "" {
				t.Error("JSON success needs a summary")
			}
			var diagnostics []struct {
				Code    string `json:"code"`
				Message string `json:"message"`
			}
			if err := json.Unmarshal(envelope["diagnostics"], &diagnostics); err != nil || diagnostics == nil {
				t.Errorf("diagnostics=%s; want array: %v", envelope["diagnostics"], err)
			}
			for _, diagnostic := range diagnostics {
				if diagnostic.Code == "" || diagnostic.Message == "" {
					t.Errorf("diagnostic needs code and message: %+v", diagnostic)
				}
			}
		} else if strings.TrimSpace(stdout) == "" {
			t.Error("plain success needs a summary")
		}
	}

	bad := authoringFixture(t)
	authoringChange(t, bad, "0001-bad.md", "", "### new:bad\nA statement without a normative verb.\n", "")
	status, plainOut, plainErr := runAppBinary(t, bin, appNestedDir(t, bad), "check")
	commandFailure(t, status, 1, plainOut, plainErr, false)
	status, jsonOut, jsonErr := runAppBinary(t, bin, appNestedDir(t, bad), "check", "--json")
	commandFailure(t, status, 1, jsonOut, jsonErr, true)
	envelope := decodeEnvelope(t, jsonOut)
	var diagnostics []struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}
	if err := json.Unmarshal(envelope["diagnostics"], &diagnostics); err != nil || len(diagnostics) < 2 {
		t.Fatalf("diagnostics=%s; want both independent errors: %v", envelope["diagnostics"], err)
	}
	for _, diagnostic := range diagnostics {
		if diagnostic.Code == "" || diagnostic.Message == "" || !strings.Contains(plainOut+plainErr, diagnostic.Message) {
			t.Errorf("diagnostic missing code/message or plain equivalent: %+v plain=%q", diagnostic, plainOut+plainErr)
		}
	}
}

func Test_APP_023_APP_024_CheckUsageAndConfiguration(t *testing.T) {
	bin := buildAppBinary(t)
	for _, tc := range []struct {
		name string
		args []string
		edit func(*testing.T, string)
	}{
		{"extra-argument", []string{"check", "extra"}, nil},
		{"invalid-config", []string{"check"}, func(t *testing.T, root string) { appPut(t, root, "fireweed.yaml", "modules: [broken\n") }},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := authoringFixture(t)
			if tc.edit != nil {
				tc.edit(t, root)
			}
			for _, jsonMode := range []bool{false, true} {
				args := append([]string{}, tc.args...)
				if jsonMode {
					args = append(args, "--json")
				}
				status, stdout, stderr := runAppBinary(t, bin, appNestedDir(t, root), args...)
				commandFailure(t, status, 2, stdout, stderr, jsonMode)
			}
		})
	}
}

func Test_APP_008_APP_025_ShowCurrentProjectedAndJSONParity(t *testing.T) {
	bin := buildAppBinary(t)
	root := authoringFixture(t)
	authoringChange(t, root, "0001-first.md", "Users need the first item.\n", "### new:first\nFirst MUST work.\n", "")
	authoringChange(t, root, "0002-second.md", "Users need the second item.\n", "### new:second\nSecond MUST work.\n", "### ALP-001\nAlpha MUST have projected text.\n")
	for _, tc := range []struct {
		name     string
		args     []string
		contains []string
		absent   []string
	}{
		{"current-full", []string{"show", "alpha"}, []string{"# Alpha", "### ALP-001", "Alpha MUST keep its current text."}, []string{"new:first", "new:second", "projected text"}},
		{"current-item", []string{"show", "alpha", "ALP-001"}, []string{"### ALP-001", "Alpha MUST keep its current text."}, []string{"# Alpha", "new:first"}},
		{"projected-full", []string{"show", "alpha", "--projected"}, []string{"# Alpha", "### new:first", "### new:second", "Alpha MUST have projected text."}, []string{"Alpha MUST keep its current text."}},
		{"projected-item", []string{"show", "alpha", "new:second", "--projected"}, []string{"### new:second", "Second MUST work."}, []string{"# Alpha", "new:first"}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			dir := appNestedDir(t, root)
			status, plain, stderr := runAppBinary(t, bin, dir, tc.args...)
			if status != 0 || stderr != "" {
				t.Fatalf("show: status=%d stdout=%q stderr=%q", status, plain, stderr)
			}
			for _, token := range tc.contains {
				if !strings.Contains(plain, token) {
					t.Errorf("show omitted %q: %q", token, plain)
				}
			}
			for _, token := range tc.absent {
				if strings.Contains(plain, token) {
					t.Errorf("show unexpectedly included %q: %q", token, plain)
				}
			}
			args := append(append([]string{}, tc.args...), "--json")
			status, output, stderr := runAppBinary(t, bin, dir, args...)
			if status != 0 || stderr != "" {
				t.Fatalf("show JSON: status=%d stdout=%q stderr=%q", status, output, stderr)
			}
			if data := commandStringData(t, output); data != plain {
				t.Errorf("JSON text=%q; want plain text %q", data, plain)
			}
		})
	}
}

func Test_APP_026_ShowNamedErrors(t *testing.T) {
	bin := buildAppBinary(t)
	for _, tc := range []struct {
		name string
		args []string
		edit func(*testing.T, string)
	}{
		{"unknown-module", []string{"show", "absent"}, nil},
		{"unknown-item", []string{"show", "alpha", "ALP-999"}, nil},
		{"ambiguous-item", []string{"show", "alpha", "new:shared", "--projected"}, func(t *testing.T, root string) {
			authoringChange(t, root, "0001-first.md", "First need.\n", "### new:shared\nFirst MUST work.\n", "")
			authoringChange(t, root, "0002-second.md", "Second need.\n", "### new:shared\nSecond MUST work.\n", "")
		}},
		{"ambiguous-module", []string{"show", "alpha"}, func(t *testing.T, root string) {
			appPut(t, root, "packages/beta/SPEC.md", "---\nname: alpha\nid_prefix: BET\ndepends_on: []\n---\n# Beta\n\n## Specs\n")
		}},
		{"missing-module", []string{"show"}, nil},
		{"extra-argument", []string{"show", "alpha", "ALP-001", "extra"}, nil},
		{"invalid-config", []string{"show", "alpha"}, func(t *testing.T, root string) { appPut(t, root, "fireweed.yaml", "modules: [broken\n") }},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := authoringFixture(t)
			if tc.edit != nil {
				tc.edit(t, root)
			}
			for _, jsonMode := range []bool{false, true} {
				args := append([]string{}, tc.args...)
				if jsonMode {
					args = append(args, "--json")
				}
				status, stdout, stderr := runAppBinary(t, bin, appNestedDir(t, root), args...)
				message := commandFailure(t, status, 2, stdout, stderr, jsonMode)
				if tc.name == "unknown-module" && !strings.Contains(message, "absent") ||
					tc.name == "unknown-item" && !strings.Contains(message, "ALP-999") ||
					tc.name == "ambiguous-item" && !strings.Contains(message, "new:shared") {
					t.Errorf("diagnostic %q does not identify %s", message, tc.name)
				}
			}
		})
	}
}
