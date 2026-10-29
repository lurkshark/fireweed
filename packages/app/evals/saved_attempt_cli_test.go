package evals_test

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/app"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/workflow"
)

// The public app entry point discovers its project from the process directory.
// These tests do not run in parallel because this helper changes that directory.
func runSavedAttemptApp(t *testing.T, dir string, args ...string) (int, string, string) {
	t.Helper()
	previous, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(dir); err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := os.Chdir(previous); err != nil {
			t.Fatal(err)
		}
	}()
	var out, errOut bytes.Buffer
	status := app.Run(args, &out, &errOut)
	return status, out.String(), errOut.String()
}

func savedAttemptAppFixture(t *testing.T, attempts int) (string, workflow.State) {
	t.Helper()
	root := lifecycleCLIProject(t, "fixture-start-does-not-run-hooks", false)
	cliGit(t, root, "config", "user.name", "Fixture")
	cliGit(t, root, "config", "user.email", "fixture@example.invalid")
	appPut(t, root, "packages/alpha/evals/current_test.go", "original eval\n")
	appPut(t, root, "packages/alpha/evals/retired_test.go", "retired eval\n")
	cliCommit(t, root, "add baseline eval")

	started, err := workflow.Start(root, "manual")
	if err != nil || started.Phase != "evals" || started.Run == "" || len(started.Change.Added) != 1 {
		t.Fatalf("fixture Start: state=%#v err=%v", started, err)
	}
	for n := 1; n <= attempts; n++ {
		appPut(t, root, "packages/alpha/evals/current_test.go", "saved attempt "+string(rune('0'+n))+"\n")
		appPut(t, root, "packages/alpha/evals/new_test.go", "new attempt "+string(rune('0'+n))+"\n")
		if err := os.Remove(filepath.Join(root, "packages/alpha/evals/retired_test.go")); err != nil && !os.IsNotExist(err) {
			t.Fatal(err)
		}
		blocked, err := workflow.Block(workflow.BlockRequest{Root: root, Change: "manual", Reason: "eval attempts exhausted"})
		if err != nil || blocked.AttemptRef != "refs/fireweed/attempts/manual/"+string(rune('0'+n)) || blocked.AttemptOID == "" {
			t.Fatalf("fixture Block %d: result=%#v err=%v", n, blocked, err)
		}
		if claim := cliGit(t, root, "for-each-ref", "--format=%(objectname)", "refs/fireweed/claims/manual"); claim != "" {
			t.Fatalf("Block retained claim: %s", claim)
		}
		if dirty := cliGit(t, root, "status", "--porcelain", "-uall"); dirty != "" {
			t.Fatalf("Block left dirty checkout: %s", dirty)
		}
		path := filepath.Join(root, "changes/0001-manual.md")
		source, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		doc, err := contracts.ParseDocument(string(source))
		if err != nil {
			t.Fatal(err)
		}
		if _, blocked := doc.Frontmatter["blocked"]; !blocked {
			t.Fatal("Block did not mark change blocked")
		}
		delete(doc.Frontmatter, "blocked")
		approved, err := contracts.RenderDocument(doc)
		if err != nil {
			t.Fatal(err)
		}
		appPut(t, root, "changes/0001-manual.md", approved)
		cliCommit(t, root, "author unblocks saved attempt")
		if n < attempts {
			continued, err := workflow.Start(root, "manual")
			if err != nil || continued.Run != started.Run || continued.Phase != started.Phase || len(continued.Change.Added) != 1 || continued.Change.Added[0].ID != started.Change.Added[0].ID {
				t.Fatalf("fixture Start after unblock: state=%#v err=%v", continued, err)
			}
		}
	}
	return root, started
}

func assertSavedAttemptUnchanged(t *testing.T, root, head, change string) {
	t.Helper()
	if got := cliGit(t, root, "rev-parse", "HEAD"); got != head {
		t.Errorf("rejected start moved HEAD: %s -> %s", head, got)
	}
	body, err := os.ReadFile(filepath.Join(root, "changes/0001-manual.md"))
	if err != nil || string(body) != change {
		t.Errorf("rejected start changed document: %v", err)
	}
	if claim := cliGit(t, root, "for-each-ref", "--format=%(objectname)", "refs/fireweed/claims/manual"); claim != "" {
		t.Errorf("rejected start claimed change: %s", claim)
	}
	if dirty := cliGit(t, root, "status", "--porcelain", "-uall"); dirty != "" {
		t.Errorf("rejected start dirtied checkout: %s", dirty)
	}
}

func Test_APP_087_StartFromRequestedAttemptRestoresWorkAndContext(t *testing.T) {
	for _, tc := range []struct {
		name     string
		attempts int
		json     bool
	}{
		{"plain-second-ref", 2, false},
		{"json-first-ref", 1, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root, started := savedAttemptAppFixture(t, tc.attempts)
			beforeHead := cliGit(t, root, "rev-parse", "HEAD")
			beforeDoc, err := os.ReadFile(filepath.Join(root, "changes/0001-manual.md"))
			if err != nil {
				t.Fatal(err)
			}
			args := []string{"start", "manual", "--from-attempt", string(rune('0' + tc.attempts))}
			if tc.json {
				args = append(args, "--json")
			}
			status, out, errOut := runSavedAttemptApp(t, appNestedDir(t, root), args...)
			if status != 0 || errOut != "" {
				t.Fatalf("resume: exit=%d stdout=%q stderr=%q", status, out, errOut)
			}
			context, err := workflow.Context(root, "manual", false)
			if err != nil {
				t.Fatal(err)
			}
			if tc.json {
				if got := commandStringData(t, out); got != context {
					t.Errorf("JSON context differs from workflow.Context")
				}
			} else if out != context {
				t.Errorf("plain context differs from workflow.Context")
			}
			current, err := workflow.Current(root, "manual")
			if err != nil || current.ClaimOID == "" || current.Run != started.Run || current.Phase != started.Phase || len(current.Change.Added) != 1 || current.Change.Added[0].ID != started.Change.Added[0].ID {
				t.Fatalf("resume lost claim, run, phase or ID: state=%#v err=%v", current, err)
			}
			if head := cliGit(t, root, "rev-parse", "HEAD"); head != beforeHead || current.Head != beforeHead || current.PhaseBase != beforeHead {
				t.Errorf("resume changed HEAD or phase base: HEAD=%s state=%#v", head, current)
			}
			for path, want := range map[string]string{
				"packages/alpha/evals/current_test.go": "saved attempt " + string(rune('0'+tc.attempts)) + "\n",
				"packages/alpha/evals/new_test.go":     "new attempt " + string(rune('0'+tc.attempts)) + "\n",
			} {
				body, err := os.ReadFile(filepath.Join(root, path))
				if err != nil || string(body) != want {
					t.Errorf("restored %s=%q, want %q: %v", path, body, want, err)
				}
			}
			if _, err := os.Stat(filepath.Join(root, "packages/alpha/evals/retired_test.go")); !os.IsNotExist(err) {
				t.Errorf("saved deletion was not restored: %v", err)
			}
			afterDoc, err := os.ReadFile(filepath.Join(root, "changes/0001-manual.md"))
			if err != nil || string(afterDoc) != string(beforeDoc) {
				t.Errorf("resume changed author-approved document: %v", err)
			}
		})
	}
}

func Test_APP_087_MissingSavedAttemptRetainsLifecycleFailureShape(t *testing.T) {
	root, _ := savedAttemptAppFixture(t, 1)
	head := cliGit(t, root, "rev-parse", "HEAD")
	change, err := os.ReadFile(filepath.Join(root, "changes/0001-manual.md"))
	if err != nil {
		t.Fatal(err)
	}
	status, out, errOut := runSavedAttemptApp(t, root, "start", "manual", "--from-attempt", "9", "--json")
	if status != 1 || errOut != "" {
		t.Errorf("missing saved ref: exit=%d stdout=%q stderr=%q", status, out, errOut)
	}
	if message := assertFailedJSON(t, out, "P01"); !corrective(message) {
		t.Errorf("missing-ref diagnostic is not corrective: %q", message)
	}
	assertSavedAttemptUnchanged(t, root, head, string(change))
}

func Test_APP_088_FromAttemptUsageRejectsBeforeMutation(t *testing.T) {
	root := lifecycleCLIProject(t, "fixture-start-does-not-run-hooks", false)
	head := cliGit(t, root, "rev-parse", "HEAD")
	change, err := os.ReadFile(filepath.Join(root, "changes/0001-manual.md"))
	if err != nil {
		t.Fatal(err)
	}
	for _, tc := range []struct {
		name  string
		args  []string
		value bool
	}{
		{"missing", []string{"start", "manual", "--from-attempt", "--json"}, true},
		{"noninteger", []string{"start", "manual", "--from-attempt", "abc", "--json"}, true},
		{"zero", []string{"start", "manual", "--from-attempt", "0", "--json"}, true},
		{"negative", []string{"start", "manual", "--from-attempt", "-1", "--json"}, true},
		{"repeated", []string{"start", "manual", "--from-attempt", "1", "--from-attempt", "2", "--json"}, true},
		{"context", []string{"context", "manual", "--from-attempt", "1", "--json"}, false},
		{"help", []string{"help", "--from-attempt", "1", "--json"}, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			status, out, errOut := runSavedAttemptApp(t, root, tc.args...)
			if status != 2 || errOut != "" {
				t.Errorf("usage: exit=%d stdout=%q stderr=%q", status, out, errOut)
			}
			message := assertFailedJSON(t, out, "U01")
			if !corrective(message) {
				t.Errorf("usage diagnostic is not corrective: %q", message)
			}
			if tc.value && (!strings.Contains(message, "start") || !strings.Contains(message, "--from-attempt") || !strings.Contains(message, "<change>")) {
				t.Errorf("value diagnostic omits valid start usage: %q", message)
			}
			assertSavedAttemptUnchanged(t, root, head, string(change))
		})
	}
}
