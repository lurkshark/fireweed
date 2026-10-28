package evals_test

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/runner"
	"github.com/lurkshark/fireweed/packages/workflow"
)

const nextChange = "changes/0001-next.md"

func nextWrite(t *testing.T, root, name, body string) {
	t.Helper()
	p := filepath.Join(root, filepath.FromSlash(name))
	if err := os.MkdirAll(filepath.Dir(p), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(p, []byte(body), 0644); err != nil {
		t.Fatal(err)
	}
}

func nextGit(t *testing.T, root string, args ...string) string {
	t.Helper()
	c := exec.Command("git", args...)
	c.Dir = root
	c.Env = append(os.Environ(), "GIT_CONFIG_NOSYSTEM=1", "GIT_AUTHOR_NAME=Fixture", "GIT_AUTHOR_EMAIL=fixture@example.invalid", "GIT_COMMITTER_NAME=Fixture", "GIT_COMMITTER_EMAIL=fixture@example.invalid")
	out, err := c.CombinedOutput()
	if err != nil {
		t.Fatalf("git %v: %v: %s", args, err, out)
	}
	return strings.TrimSpace(string(out))
}

func nextCommit(t *testing.T, root, message string) {
	t.Helper()
	nextGit(t, root, "add", "-A")
	nextGit(t, root, "commit", "-q", "-m", message)
}

func nextRead(t *testing.T, root, path string) string {
	t.Helper()
	b, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(path)))
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}

type nextFixture struct{ root, logDir string }

func newNextFixture(t *testing.T, queued, approved bool, workerMode string) nextFixture {
	t.Helper()
	root, err := filepath.EvalSymlinks(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	logDir, err := filepath.EvalSymlinks(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	worker := filepath.Join(logDir, "worker.sh")
	auditor := filepath.Join(logDir, "auditor.sh")
	// The scripts are fixture agents. Only the runner may call workflow transitions.
	nextWrite(t, logDir, "worker.sh", `#!/bin/sh
phase=$(sed -n 's/^phase: //p' changes/0001-next.md | head -1)
printf '%s' "$3" > "$1/worker-$phase.prompt"
printf '%s|%s|%s\n' "$phase" "$PWD" "$FIREWEED_BIN" >> "$1/invocations"
if [ "$2" = exit ]; then printf 'partial worker output'; printf 'worker failure' >&2; exit 9; fi
if [ "$2" = noedit ]; then printf 'no edit'; exit 0; fi
sleep 0.15
case "$phase" in
  evals)
    cat > packages/alpha/evals/next_test.go <<'GO'
package evals
import (
    "testing"
    "example.invalid/next-fixture/packages/alpha"
)
func Test_ALP_002_NewValue(t *testing.T) { // EVAL_CONTENT_MARKER
    if alpha.Value() != "new" { t.Fatalf("ALP-002: got %q", alpha.Value()) }
}
GO
    ;;
  impl)
    cat > packages/alpha/internal/value.go <<'GO'
package internal
func Value() string { return "new" }
GO
    ;;
  *) exit 8 ;;
esac
printf 'worker-%s' "$phase"
`)
	nextWrite(t, logDir, "auditor.sh", `#!/bin/sh
printf '%s' "$2" > "$1/audit.prompt"
printf 'audit|%s|%s\n' "$PWD" "$FIREWEED_BIN" >> "$1/invocations"
cat >> changes/0001-next.md <<'AUDIT'

## Audit
No findings
Checked IDs: ALP-002
AUDIT
printf 'auditor'
`)
	nextWrite(t, root, ".gitignore", ".fireweed/runs/\n.fireweed/go-test.log\n")
	nextWrite(t, root, "go.mod", "module example.invalid/next-fixture\n\ngo 1.22\n")
	nextWrite(t, root, "SYSTEM.md", "# Fixture\nSYSTEM_CONTEXT_MARKER\n")
	nextWrite(t, root, "fireweed.yaml", fmt.Sprintf(`fireweed: ">=0.1, <0.2"
version: 1
base: main
modules: ["packages/*"]
module_defaults:
  interface: [api.go]
  evals: [evals/**]
  disposable: [internal/**]
hooks:
  eval: [sh, .fireweed/eval.sh, "{module}", "{report}"]
  boundaries: [sh, .fireweed/boundaries.sh]
agent:
  command: [sh, %q, %q, %q, "{prompt}"]
audit:
  command: [sh, %q, %q, "{prompt}"]
timeout: 50ms
`, worker, logDir, workerMode, auditor, logDir))
	nextWrite(t, root, ".fireweed/eval.sh", `#!/bin/sh
go test ./packages/alpha/evals > .fireweed/go-test.log 2>&1
code=$?
if [ "$code" -eq 0 ]; then result=''; else result='<failure message="Go eval failed"/>'; fi
printf '<testsuite><testcase name="ALP-001" classname="fixture"/><testcase name="ALP-002" classname="fixture">%s</testcase></testsuite>' "$result" > "$2"
exit "$code"
`)
	nextWrite(t, root, ".fireweed/boundaries.sh", "#!/bin/sh\nexit 0\n")
	nextWrite(t, root, "packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\n---\n# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n## Decisions\n")
	nextWrite(t, root, "packages/alpha/api.go", "package alpha\nimport \"example.invalid/next-fixture/packages/alpha/internal\"\nfunc Value() string { return internal.Value() }\n")
	nextWrite(t, root, "packages/alpha/internal/value.go", "package internal\n// FULL_INTERNAL_SECRET\nfunc Value() string { return \"old\" }\n")
	nextWrite(t, root, "packages/alpha/evals/existing_test.go", "package evals\nimport (\"testing\"; \"example.invalid/next-fixture/packages/alpha\")\nfunc Test_ALP_001_Value(t *testing.T) { if alpha.Value() == \"\" { t.Fatal(\"empty\") } }\n")
	nextGit(t, root, "init", "-q")
	nextGit(t, root, "checkout", "-q", "-b", "main")
	nextGit(t, root, "config", "user.name", "Fixture")
	nextGit(t, root, "config", "user.email", "fixture@example.invalid")
	if queued && !approved {
		nextCommit(t, root, "base fixture")
		nextGit(t, root, "checkout", "-q", "-b", "feature")
	}
	if queued {
		nextWrite(t, root, nextChange, "---\nmodule: alpha\nbases: {}\n---\n\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### new:value\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n")
	}
	nextCommit(t, root, "approved fixture")
	for _, d := range project.Check(root) {
		if !d.Warning {
			t.Fatalf("bad fixture: %#v", d)
		}
	}
	return nextFixture{root, logDir}
}

func (f nextFixture) run(allow bool) (runner.NextResult, error) {
	return runner.Next(context.Background(), runner.NextRequest{Root: f.root, Bin: "/fixture/next-bin", AllowUnmerged: allow, Timeout: 20 * time.Second})
}

func (f nextFixture) noAgent(t *testing.T) {
	t.Helper()
	if _, err := os.Stat(filepath.Join(f.logDir, "invocations")); !os.IsNotExist(err) {
		t.Errorf("agent invoked: %v", err)
	}
}

func Test_RUN_010_EmptyQueue(t *testing.T) {
	f := newNextFixture(t, false, true, "normal")
	before := nextGit(t, f.root, "rev-parse", "HEAD")
	r, err := f.run(false)
	if err != nil || !r.Empty || len(r.Steps) != 0 || nextGit(t, f.root, "rev-parse", "HEAD") != before {
		t.Errorf("empty Next = %#v, %v", r, err)
	}
	f.noAgent(t)
}

func Test_RUN_011_RUN_012_StartAndApprovalPrecedeAgents(t *testing.T) {
	t.Run("unmerged", func(t *testing.T) {
		f := newNextFixture(t, true, false, "normal")
		before := nextGit(t, f.root, "rev-parse", "HEAD")
		_, err := f.run(false)
		if err == nil || !strings.Contains(strings.ToLower(err.Error()), "approv") {
			t.Errorf("unapproved Next error = %v", err)
		}
		if nextGit(t, f.root, "rev-parse", "HEAD") != before || strings.Contains(nextRead(t, f.root, nextChange), "run:") {
			t.Error("unapproved change was claimed or assigned")
		}
		if out := nextGit(t, f.root, "for-each-ref", "refs/fireweed/claims"); out != "" {
			t.Errorf("unapproved claim = %s", out)
		}
		f.noAgent(t)
	})
	t.Run("start_failure", func(t *testing.T) {
		f := newNextFixture(t, true, true, "normal")
		nextWrite(t, f.root, "stray.txt", "dirty\n")
		_, err := f.run(false)
		if err == nil || !strings.Contains(err.Error(), "P01") {
			t.Errorf("dirty start error = %v", err)
		}
		f.noAgent(t)
	})
	t.Run("first_unblocked", func(t *testing.T) {
		f := newNextFixture(t, true, true, "exit")
		nextWrite(t, f.root, "changes/0000-blocked.md", "---\nmodule: alpha\nblocked: human review\n---\n\n## Motivation\nBlocked work.\n\n## Interface\n\n## Added\n\n## Modified\n\n## Removed\n\n## Decisions\n")
		nextCommit(t, f.root, "blocked first change")
		_, err := f.run(false)
		if err == nil {
			t.Error("expected worker failure after eligible start")
		}
		state, currentErr := workflow.Current(f.root, "next")
		if currentErr != nil || state.Change.Slug != "next" || state.Phase != "evals" {
			t.Errorf("selected change = %#v, %v", state, currentErr)
		}
	})
	t.Run("all_blocked", func(t *testing.T) {
		f := newNextFixture(t, true, true, "normal")
		change := nextRead(t, f.root, nextChange)
		nextWrite(t, f.root, nextChange, strings.Replace(change, "module: alpha\n", "module: alpha\nblocked: human review\n", 1))
		nextCommit(t, f.root, "block queue")
		if _, err := f.run(false); err == nil {
			t.Error("all-blocked queue accepted")
		}
		if refs := nextGit(t, f.root, "for-each-ref", "refs/fireweed/claims"); refs != "" {
			t.Errorf("all-blocked queue claimed change: %s", refs)
		}
		f.noAgent(t)
	})
	t.Run("unapproved_archive_bound", func(t *testing.T) {
		f := newNextFixture(t, true, true, "normal")
		nextGit(t, f.root, "update-ref", "refs/remotes/origin/main", "HEAD")
		nextWrite(t, f.root, "changes/archive/0000-prior.md", "---\nmodule: alpha\nphase: applied\nrun: prior\nbases: {}\n---\n\n## Motivation\nPrior work.\n\n## Interface\n\n## Added\n\n## Modified\n\n## Removed\n\n## Decisions\n")
		nextWrite(t, f.root, "fireweed.yaml", nextRead(t, f.root, "fireweed.yaml")+"authoring:\n  max_unapproved: 1\n")
		nextCommit(t, f.root, "unapproved prior archive")
		for _, d := range project.Check(f.root) {
			if !d.Warning {
				t.Fatalf("bad bound fixture: %#v", d)
			}
		}
		_, err := f.run(false)
		if err == nil || !strings.Contains(err.Error(), "P07") {
			t.Errorf("unapproved archive bound error = %v", err)
		}
		if refs := nextGit(t, f.root, "for-each-ref", "refs/fireweed/claims"); refs != "" {
			t.Errorf("P07 claimed change: %s", refs)
		}
		f.noAgent(t)
	})
	t.Run("shared_base_approval", func(t *testing.T) {
		f := newNextFixture(t, true, false, "exit")
		nextGit(t, f.root, "update-ref", "refs/remotes/origin/main", "HEAD")
		_, err := f.run(false)
		if err == nil {
			t.Error("expected worker failure after shared-base approval")
		}
		state, currentErr := workflow.Current(f.root, "next")
		if currentErr != nil || state.Phase != "evals" {
			t.Errorf("shared-base approval = %#v, %v", state, currentErr)
		}
	})
	t.Run("allow_unmerged", func(t *testing.T) {
		f := newNextFixture(t, true, false, "exit")
		_, err := f.run(true)
		if err == nil {
			t.Error("expected worker failure after allowed start")
		}
		state, currentErr := workflow.Current(f.root, "next")
		if currentErr != nil || state.Phase != "evals" || state.Run == "" || !strings.Contains(nextRead(t, f.root, nextChange), "ALP-002") {
			t.Errorf("allowed start = %#v, %v", state, currentErr)
		}
	})
}

func Test_RUN_013_RUN_014_CompleteSpecChange(t *testing.T) {
	f := newNextFixture(t, true, true, "normal")
	before := nextGit(t, f.root, "rev-parse", "HEAD")
	r, err := f.run(false)
	if err != nil {
		t.Fatalf("Next: %#v, %v", r, err)
	}
	if r.Empty || r.State.Phase != "applied" || r.State.Change.Slug != "next" || r.State.Run == "" || r.State.ClaimOID != "" || r.State.Snapshot.Root != f.root {
		t.Errorf("applied result = %#v", r)
	}
	if len(r.Steps) != 3 {
		t.Fatalf("steps = %#v", r.Steps)
	}
	for i, want := range []struct{ phase, role, stdout string }{{"evals", "worker", "worker-evals"}, {"audit", "audit", "auditor"}, {"impl", "worker", "worker-impl"}} {
		got := r.Steps[i]
		if got.Phase != want.phase || got.Role != want.role || got.Agent.Stdout != want.stdout || got.Agent.ExitCode != 0 {
			t.Errorf("step %d = %#v", i, got)
		}
	}
	lines := strings.Split(strings.TrimSpace(nextRead(t, f.logDir, "invocations")), "\n")
	if len(lines) != 3 || lines[0] != "evals|"+f.root+"|/fixture/next-bin" || lines[1] != "audit|"+f.root+"|/fixture/next-bin" || lines[2] != "impl|"+f.root+"|/fixture/next-bin" {
		t.Errorf("invocations = %#v", lines)
	}
	for _, phase := range []string{"evals", "impl"} {
		prompt := nextRead(t, f.logDir, "worker-"+phase+".prompt")
		for _, marker := range []string{"# Change", "# Projected spec", "# Interface", "# SYSTEM.md", "# Playbook", "# Other relevant files", "FULL_INTERNAL_SECRET", "SYSTEM_CONTEXT_MARKER", "Alpha MUST return the new value."} {
			if !strings.Contains(prompt, marker) {
				t.Errorf("%s worker context omitted %q", phase, marker)
			}
		}
	}
	audit := nextRead(t, f.logDir, "audit.prompt")
	for _, marker := range []string{nextChange, "Alpha MUST return the new value.", "### ALP-002", "EVAL_CONTENT_MARKER", "# Audit evals"} {
		if !strings.Contains(audit, marker) {
			t.Errorf("audit context omitted %q", marker)
		}
	}
	for _, secret := range []string{"FULL_INTERNAL_SECRET", "SYSTEM_CONTEXT_MARKER"} {
		if strings.Contains(audit, secret) {
			t.Errorf("audit context exposed %q", secret)
		}
	}
	if _, err := os.Stat(filepath.Join(f.root, nextChange)); !os.IsNotExist(err) {
		t.Errorf("pending change remains: %v", err)
	}
	archived := nextRead(t, f.root, "changes/archive/0001-next.md")
	for _, marker := range []string{"phase: applied", "### ALP-002", "No findings", "- evals ", "- audit ", "- impl ", "agent=sh"} {
		if !strings.Contains(archived, marker) {
			t.Errorf("archive omitted %q", marker)
		}
	}
	if !strings.Contains(nextRead(t, f.root, "packages/alpha/SPEC.md"), "Alpha MUST return the new value.") {
		t.Error("SPEC not applied")
	}
	if refs := nextGit(t, f.root, "for-each-ref", "refs/fireweed/claims"); refs != "" {
		t.Errorf("claim retained: %s", refs)
	}
	if count := nextGit(t, f.root, "rev-list", "--count", before+"..HEAD"); count != "3" {
		t.Errorf("phase commit count = %s", count)
	}
	if old := nextGit(t, f.root, "show", "HEAD~2:packages/alpha/internal/value.go"); !strings.Contains(old, `return "old"`) {
		t.Errorf("evals commit changed implementation: %q", old)
	}
	if first := nextGit(t, f.root, "show", "--format=%s", "--no-patch", "HEAD~2"); !strings.Contains(first, "evals") {
		t.Errorf("first phase commit = %q", first)
	}
	if second := nextGit(t, f.root, "show", "--format=%s", "--no-patch", "HEAD~1"); !strings.Contains(second, "audit") {
		t.Errorf("second phase commit = %q", second)
	}
	if third := nextGit(t, f.root, "show", "--format=%s", "--no-patch", "HEAD"); !strings.Contains(third, "impl") {
		t.Errorf("third phase commit = %q", third)
	}
	failReport := nextRead(t, f.root, ".fireweed/runs/"+r.State.Run+"/evals/alpha.json")
	passReport := nextRead(t, f.root, ".fireweed/runs/"+r.State.Run+"/impl/alpha.json")
	if !strings.Contains(failReport, "ALP-002") || !strings.Contains(failReport, "fail") || !strings.Contains(passReport, "ALP-002") || !strings.Contains(passReport, "pass") {
		t.Errorf("phase reports omitted genuine fail-first/pass: evals=%q impl=%q", failReport, passReport)
	}
	if _, err := workflow.Current(f.root, "next"); err == nil {
		t.Error("applied change remains current")
	}
}

func Test_RUN_015_InvocationAndVerificationFailureRetainPhase(t *testing.T) {
	for _, mode := range []string{"exit", "noedit"} {
		t.Run(mode, func(t *testing.T) {
			f := newNextFixture(t, true, true, mode)
			before := nextGit(t, f.root, "rev-parse", "HEAD")
			r, err := f.run(false)
			if err == nil || len(r.Steps) != 1 || r.Steps[0].Phase != "evals" || r.Steps[0].Role != "worker" {
				t.Errorf("failure = %#v, %v", r, err)
			}
			want := "partial worker output"
			if mode == "noedit" {
				want = "no edit"
			}
			if len(r.Steps) == 1 && (r.Steps[0].Agent.Stdout != want || (mode == "exit" && (r.Steps[0].Agent.ExitCode != 9 || r.Steps[0].Agent.Stderr != "worker failure"))) {
				t.Errorf("captured invocation = %#v", r.Steps[0])
			}
			state, currentErr := workflow.Current(f.root, "next")
			if currentErr != nil || state.Phase != "evals" || state.Run == "" || nextGit(t, f.root, "rev-parse", "HEAD") != before {
				t.Errorf("failure advanced phase: %#v, %v", state, currentErr)
			}
		})
	}
}
