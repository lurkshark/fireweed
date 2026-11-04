package evals_test

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/project"
)

const migrateConfirmation = "Project format 1 is already current.\n"

type migrateState struct {
	files map[string][]byte
	index []byte
	head  string
	refs  string
}

func migrateProject(t *testing.T) string {
	t.Helper()
	root := authoringFixture(t)
	appPut(t, root, "notes/tracked.txt", "original\n")
	cliGit(t, root, "init", "-q", "-b", "main")
	cliCommit(t, root, "format 1 fixture")
	return root
}

func migrateCapture(t *testing.T, root string) migrateState {
	t.Helper()
	files := map[string][]byte{}
	err := filepath.WalkDir(root, func(path string, entry os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() {
			if entry.Name() == ".git" {
				return filepath.SkipDir
			}
			return nil
		}
		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		files[filepath.ToSlash(rel)], err = os.ReadFile(path)
		return err
	})
	if err != nil {
		t.Fatal(err)
	}
	index, err := os.ReadFile(filepath.Join(root, ".git", "index"))
	if err != nil {
		t.Fatal(err)
	}
	return migrateState{
		files: files, index: index,
		head: cliGit(t, root, "rev-parse", "HEAD"),
		refs: cliGit(t, root, "for-each-ref", "--format=%(refname) %(objectname)"),
	}
}

func migrateUnchanged(t *testing.T, root string, before migrateState) {
	t.Helper()
	after := migrateCapture(t, root)
	if !reflect.DeepEqual(after.files, before.files) {
		t.Errorf("migrate changed working files: before=%v after=%v", before.files, after.files)
	}
	if !bytes.Equal(after.index, before.index) || after.head != before.head || after.refs != before.refs {
		t.Errorf("migrate changed index, HEAD or refs: before=(%s, %s) after=(%s, %s)", before.head, before.refs, after.head, after.refs)
	}
}

func migrateDiagnostics(t *testing.T, output string) []project.Diagnostic {
	t.Helper()
	env := decodeEnvelope(t, output)
	var raw []map[string]json.RawMessage
	if err := json.Unmarshal(env["diagnostics"], &raw); err != nil || raw == nil {
		t.Fatalf("diagnostics=%s; want array: %v", env["diagnostics"], err)
	}
	var diagnostics []project.Diagnostic
	if err := json.Unmarshal(env["diagnostics"], &diagnostics); err != nil {
		t.Fatal(err)
	}
	for _, item := range raw {
		for _, field := range []string{"code", "message", "warning"} {
			if _, ok := item[field]; !ok {
				t.Errorf("diagnostic %v lacks %s", item, field)
			}
		}
	}
	return diagnostics
}

func Test_APP_153_APP_154_MigrateNearestRootAndRepeatedNoOp(t *testing.T) {
	bin := buildAppBinary(t)
	outer := migrateProject(t)
	appPut(t, outer, "packages/alpha/SPEC.md", strings.Replace(authoringSpec, "Alpha MUST keep its current text.", "Invalid outer text.", 1))
	inner := filepath.Join(outer, "inner-project")
	appPut(t, inner, "fireweed.yaml", appFixtureConfig)
	appPut(t, inner, "packages/alpha/SPEC.md", authoringSpec)
	appPut(t, inner, "notes/tracked.txt", "original\n")
	cliGit(t, inner, "init", "-q", "-b", "main")
	cliCommit(t, inner, "inner format 1 fixture")
	cliGit(t, inner, "update-ref", "refs/fireweed/fixture", cliGit(t, inner, "rev-parse", "HEAD"))
	appPut(t, inner, "notes/tracked.txt", "staged\n")
	cliGit(t, inner, "add", "notes/tracked.txt")
	appPut(t, inner, "notes/tracked.txt", "staged and dirty\n")
	appPut(t, inner, "notes/untracked.txt", "untracked\n")
	beforeInner, beforeOuter := migrateCapture(t, inner), migrateCapture(t, outer)
	nested := appNestedDir(t, inner)
	for i := 0; i < 2; i++ {
		status, plain, stderr := runAppBinary(t, bin, nested, "migrate")
		if status != 0 || stderr != "" || plain != migrateConfirmation {
			t.Errorf("plain migrate %d: status=%d stdout=%q stderr=%q; want exact confirmation", i, status, plain, stderr)
		}
		status, output, stderr := runAppBinary(t, bin, nested, "migrate", "--json")
		if status != 0 || stderr != "" {
			t.Fatalf("JSON migrate %d: status=%d stdout=%q stderr=%q", i, status, output, stderr)
		}
		if got := commandStringData(t, output); got != plain {
			t.Errorf("JSON confirmation=%q; want plain %q", got, plain)
		}
		migrateUnchanged(t, inner, beforeInner)
		migrateUnchanged(t, outer, beforeOuter)
	}
}

func Test_APP_153_APP_154_MigrateEmptyCurrentProject(t *testing.T) {
	bin := buildAppBinary(t)
	root := t.TempDir()
	appPut(t, root, "fireweed.yaml", "version: 1\nmodules: []\n")
	cliGit(t, root, "init", "-q", "-b", "main")
	cliCommit(t, root, "empty format 1 project")
	before := migrateCapture(t, root)
	status, plain, stderr := runAppBinary(t, bin, root, "migrate")
	if status != 0 || stderr != "" || plain != migrateConfirmation {
		t.Errorf("empty plain migrate: status=%d stdout=%q stderr=%q", status, plain, stderr)
	}
	status, output, stderr := runAppBinary(t, bin, root, "migrate", "--json")
	if status != 0 || stderr != "" || commandStringData(t, output) != plain {
		t.Errorf("empty JSON migrate: status=%d stdout=%q stderr=%q", status, output, stderr)
	}
	migrateUnchanged(t, root, before)
}

func Test_APP_154_APP_155_MigrateStaticDiagnosticsAndFailures(t *testing.T) {
	bin := buildAppBinary(t)
	for _, tc := range []struct {
		name, code string
		status     int
		edit       func(*testing.T, string)
	}{
		{"static-check", "C04", 1, func(t *testing.T, root string) {
			appPut(t, root, "packages/alpha/SPEC.md", strings.Replace(authoringSpec, "Alpha MUST keep its current text.", "Alpha has no normative verb.", 1))
		}},
		{"malformed-config", "C01", 2, func(t *testing.T, root string) {
			appPut(t, root, "fireweed.yaml", "modules: [broken\n")
		}},
		{"malformed-module", "C03", 2, func(t *testing.T, root string) {
			appPut(t, root, "packages/alpha/SPEC.md", strings.Replace(authoringSpec, "### ALP-001", "### invalid heading", 1))
		}},
		{"unsupported-version", "C01", 2, func(t *testing.T, root string) {
			appPut(t, root, "fireweed.yaml", strings.Replace(appFixtureConfig, "version: 1", "version: 2", 1))
		}},
		{"missing-version", "C01", 2, func(t *testing.T, root string) {
			appPut(t, root, "fireweed.yaml", strings.Replace(appFixtureConfig, "version: 1\n", "", 1))
		}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := migrateProject(t)
			tc.edit(t, root)
			before := migrateCapture(t, root)
			var expected []project.Diagnostic
			if tc.name == "static-check" {
				expected = project.Check(root)
				if len(expected) == 0 || expected[0].Code != tc.code || expected[0].Warning {
					t.Fatalf("static fixture did not produce %s: %+v", tc.code, expected)
				}
			}
			for _, jsonMode := range []bool{false, true} {
				args := []string{"migrate"}
				if jsonMode {
					args = append(args, "--json")
				}
				status, out, stderr := runAppBinary(t, bin, appNestedDir(t, root), args...)
				if status != tc.status {
					t.Errorf("%s JSON=%v: status=%d stdout=%q stderr=%q; want %d", tc.name, jsonMode, status, out, stderr, tc.status)
				}
				if jsonMode {
					if stderr != "" {
						t.Errorf("JSON stderr=%q", stderr)
					}
					if string(decodeEnvelope(t, out)["data"]) != "null" {
						t.Errorf("failure data=%s; want null", decodeEnvelope(t, out)["data"])
					}
					got := migrateDiagnostics(t, out)
					if len(got) == 0 || got[0].Code != tc.code {
						t.Errorf("JSON diagnostics=%+v; want %s", got, tc.code)
					}
					for i, want := range expected {
						if i >= len(got) || got[i].Code != want.Code || !strings.Contains(got[i].Message, want.Message) || got[i].Warning != want.Warning || got[i].Path != want.Path {
							t.Errorf("static diagnostic %d=%+v; want cause %+v", i, got, want)
						}
					}
					for _, d := range got {
						if !corrective(d.Message) {
							t.Errorf("diagnostic lacks corrective guidance: %+v", d)
						}
					}
				} else if out != "" || !strings.Contains(stderr, tc.code+":") || !corrective(stderr) {
					t.Errorf("plain failure stdout=%q stderr=%q; want corrective %s", out, stderr, tc.code)
				}
				migrateUnchanged(t, root, before)
			}
		})
	}
	for _, jsonMode := range []bool{false, true} {
		args := []string{"migrate"}
		if jsonMode {
			args = append(args, "--json")
		}
		status, out, stderr := runAppBinary(t, bin, t.TempDir(), args...)
		message := commandFailure(t, status, 2, out, stderr, jsonMode)
		if jsonMode {
			message = assertFailedJSON(t, out, "C01")
		} else if !strings.HasPrefix(strings.TrimSpace(stderr), "C01:") {
			t.Errorf("discovery stderr=%q; want C01", stderr)
		}
		if !corrective(message) {
			t.Errorf("discovery diagnostic lacks guidance: %q", message)
		}
	}
}

func Test_APP_155_MigrateRejectsArgumentsBeforeMutation(t *testing.T) {
	bin := buildAppBinary(t)
	root := migrateProject(t)
	before := migrateCapture(t, root)
	for _, args := range [][]string{{"migrate", "extra"}, {"migrate", "--bogus"}, {"migrate", "--max", "1"}} {
		for _, jsonMode := range []bool{false, true} {
			call := append([]string(nil), args...)
			if jsonMode {
				call = append(call, "--json")
			}
			status, out, stderr := runAppBinary(t, bin, appNestedDir(t, root), call...)
			message := commandFailure(t, status, 2, out, stderr, jsonMode)
			if jsonMode {
				message = assertFailedJSON(t, out, "U01")
			} else if out != "" || !strings.HasPrefix(strings.TrimSpace(stderr), "U01:") {
				t.Errorf("plain invalid request: stdout=%q stderr=%q", out, stderr)
			}
			if !corrective(message) {
				t.Errorf("invalid request lacks guidance: %q", message)
			}
			migrateUnchanged(t, root, before)
		}
	}
}
