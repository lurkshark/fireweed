package evals_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/workflow"
)

func commitsAppPhase(t *testing.T, root, path, old, next string) string {
	t.Helper()
	data, err := os.ReadFile(filepath.Join(root, path))
	if err != nil {
		t.Fatal(err)
	}
	before := string(data)
	var after string
	if old == "" {
		after = strings.Replace(before, "module: alpha\n", "module: alpha\nphase: "+next+"\n", 1)
	} else {
		after = strings.Replace(before, "phase: "+old+"\n", "phase: "+next+"\n", 1)
	}
	if after == before {
		t.Fatalf("fixture lacks phase %q in %s", old, path)
	}
	appPut(t, root, path, after)
	cliCommit(t, root, "fixture phase")
	return cliGit(t, root, "rev-parse", "HEAD")
}

func commitsAppSuccess(t *testing.T, root, selector string, jsonMode bool, want []string) {
	t.Helper()
	args := []string{"commits", selector}
	if jsonMode {
		args = append(args, "--json")
	}
	status, out, errOut := runSavedAttemptApp(t, appNestedDir(t, root), args...)
	if status != 0 || errOut != "" {
		t.Fatalf("commits %q: exit=%d stdout=%q stderr=%q", selector, status, out, errOut)
	}
	if !jsonMode {
		wantPlain := strings.Join(want, "\n")
		if len(want) > 0 {
			wantPlain += "\n"
		}
		if got := out; got != wantPlain {
			t.Errorf("plain commits %q=%q; want ordered IDs %v", selector, got, want)
		}
		return
	}
	env := decodeEnvelope(t, out)
	var got []string
	if err := json.Unmarshal(env["data"], &got); err != nil || got == nil {
		t.Fatalf("JSON commits data=%s; want array: %v", env["data"], err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("JSON commits %q=%v; want ordered IDs %v", selector, got, want)
	}
	var diagnostics []json.RawMessage
	if err := json.Unmarshal(env["diagnostics"], &diagnostics); err != nil || diagnostics == nil || len(diagnostics) != 0 {
		t.Errorf("success diagnostics=%s; want empty array", env["diagnostics"])
	}
}

func Test_APP_094_APP_096_CommitsPendingAndArchivedOrder(t *testing.T) {
	root := lifecycleCLIProject(t, "fixture-start-does-not-run-hooks", false)
	path := "changes/0001-manual.md"
	commitsAppSuccess(t, root, "manual", false, []string{})
	commitsAppSuccess(t, root, path, true, []string{})
	evals := commitsAppPhase(t, root, path, "", "evals")
	audit := commitsAppPhase(t, root, path, "evals", "audit")
	impl := commitsAppPhase(t, root, path, "audit", "impl")
	wantPending := []string{impl, audit, evals}
	commitsAppSuccess(t, root, "manual", false, wantPending)
	commitsAppSuccess(t, root, "0001-manual.md", true, wantPending)
	archived := "changes/archive/0001-manual.md"
	if err := os.MkdirAll(filepath.Join(root, "changes", "archive"), 0755); err != nil {
		t.Fatal(err)
	}
	cliGit(t, root, "mv", path, archived)
	applied := commitsAppPhase(t, root, archived, "impl", "applied")
	wantArchived := []string{applied, impl, audit, evals}
	commitsAppSuccess(t, root, archived, false, wantArchived)
	commitsAppSuccess(t, root, "manual", true, wantArchived)
}

func commitsAppFailure(t *testing.T, root string, args []string, wantStatus int, code string) {
	t.Helper()
	jsonMode := len(args) > 0 && args[len(args)-1] == "--json"
	status, out, errOut := runSavedAttemptApp(t, appNestedDir(t, root), args...)
	message := commandFailure(t, status, wantStatus, out, errOut, jsonMode)
	if jsonMode {
		if code == "" {
			message = assertFailedJSON(t, out)
		} else {
			message = assertFailedJSON(t, out, code)
		}
	} else if code != "" && !strings.HasPrefix(strings.TrimSpace(message), code+":") {
		t.Errorf("diagnostic=%q; want %s cause", message, code)
	}
	if !corrective(message) {
		t.Errorf("diagnostic %q must suggest a correction", message)
	}
}

func Test_APP_095_APP_096_CommitsFailures(t *testing.T) {
	root := lifecycleCLIProject(t, "fixture-start-does-not-run-hooks", false)
	for _, tc := range []struct {
		name string
		args []string
		code string
	}{
		{"missing", []string{"commits"}, "U01"},
		{"extra-json", []string{"commits", "manual", "extra", "--json"}, "U01"},
		{"unsupported-option", []string{"commits", "manual", "--bogus"}, "U01"},
		{"unknown-selector", []string{"commits", "missing", "--json"}, "U02"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			commitsAppFailure(t, root, tc.args, 2, tc.code)
		})
	}
	appPut(t, root, "fireweed.yaml", "modules: [broken\n")
	commitsAppFailure(t, root, []string{"commits", "manual", "--json"}, 2, "C01")

	operational := lifecycleCLIProject(t, "fixture-start-does-not-run-hooks", false)
	if err := os.RemoveAll(filepath.Join(operational, ".git")); err != nil {
		t.Fatal(err)
	}
	_, dependencyErr := workflow.Commits(operational, "manual")
	if dependencyErr == nil {
		t.Fatal("fixture must cause workflow.Commits to fail")
	}
	commitsAppFailure(t, operational, []string{"commits", "manual", "--json"}, 1, "P01")
}
