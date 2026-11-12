package evals_test

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/runner"
	"github.com/lurkshark/fireweed/packages/workflow"
)

func auditGitObject(root, oid string) (string, error) {
	command := exec.Command("git", "cat-file", "-t", oid)
	command.Dir = root
	out, err := command.CombinedOutput()
	return string(out), err
}

// auditAction runs after the ordinary fixture auditor writes its Audit section.
// The script lives outside the fixture repository so it is never an audit input.
func auditAction(t *testing.T, f nextFixture, action string) {
	t.Helper()
	script := nextRead(t, f.logDir, "auditor.sh")
	anchor := "printf 'auditor'\n"
	if !strings.Contains(script, anchor) {
		t.Fatal("auditor insertion point missing")
	}
	nextWrite(t, f.logDir, "auditor.sh", strings.Replace(script, anchor, action+anchor, 1))
}

func auditBeforeEdit(t *testing.T, f nextFixture, action string) {
	t.Helper()
	script := nextRead(t, f.logDir, "auditor.sh")
	anchor := "cat >> changes/0001-next.md <<'AUDIT'\n"
	if !strings.Contains(script, anchor) {
		t.Fatal("auditor input inspection point missing")
	}
	nextWrite(t, f.logDir, "auditor.sh", strings.Replace(script, anchor, action+anchor, 1))
}

func auditWorkspaceGone(t *testing.T, f nextFixture) {
	t.Helper()
	cwd := nextRead(t, f.logDir, "audit.cwd")
	if cwd == f.root || cwd == "" {
		t.Errorf("audit cwd = %q, held root = %q", cwd, f.root)
		return
	}
	if _, err := os.Stat(cwd); !os.IsNotExist(err) {
		t.Errorf("audit workspace %q remains: %v", cwd, err)
	}
}

func Test_RUN_013_RUN_045_RUN_046_AuditUsesOnlyProjectedInputsAndIndependentGit(t *testing.T) {
	f := newNextFixture(t, true, true, "normal")
	heldOID := nextGit(t, f.root, "rev-parse", "HEAD")
	nextGit(t, f.root, "update-ref", "refs/fireweed/attempts/sentinel", heldOID)
	auditBeforeEdit(t, f, `find . -path './.git' -prune -o -type f -print > "$1/audit.files"
git rev-parse --is-inside-work-tree > "$1/audit.is-git"
git for-each-ref --format='%(refname)' > "$1/audit.refs"
if git cat-file -e '`+heldOID+`^{commit}' 2>/dev/null; then printf present > "$1/audit.held-history"; else printf absent > "$1/audit.held-history"; fi
cp changes/0001-next.md "$1/audit.change"
cp packages/alpha/SPEC.md "$1/audit.spec"
cp packages/alpha/evals/next_test.go "$1/audit.eval"
`)
	r, err := f.run(false)
	if err != nil || r.State.Phase != "applied" {
		t.Fatalf("isolated audit lifecycle = %#v, %v", r, err)
	}
	cwd := nextRead(t, f.logDir, "audit.cwd")
	if cwd == f.root || !filepath.IsAbs(cwd) || nextRead(t, f.logDir, "audit.is-git") != "true\n" {
		t.Errorf("audit root = %q; git status = %q", cwd, nextRead(t, f.logDir, "audit.is-git"))
	}
	files := strings.Fields(nextRead(t, f.logDir, "audit.files"))
	sort.Strings(files)
	want := []string{"./changes/0001-next.md", "./packages/alpha/SPEC.md", "./packages/alpha/evals/existing_test.go", "./packages/alpha/evals/next_test.go"}
	if strings.Join(files, "\n") != strings.Join(want, "\n") {
		t.Errorf("audit workspace files = %q, want %q", files, want)
	}
	if got := nextRead(t, f.logDir, "audit.held-history"); got != "absent" {
		t.Errorf("audit Git can read held commit: %q", got)
	}
	if refs := nextRead(t, f.logDir, "audit.refs"); strings.Contains(refs, "sentinel") || strings.Contains(refs, "fireweed/") {
		t.Errorf("audit Git exposes held refs: %q", refs)
	}
	for _, item := range []struct{ file, marker string }{{"audit.change", "phase: audit"}, {"audit.change", "### ALP-002"}, {"audit.spec", "### ALP-002"}, {"audit.eval", "EVAL_CONTENT_MARKER"}} {
		if !strings.Contains(nextRead(t, f.logDir, item.file), item.marker) {
			t.Errorf("%s omitted %q", item.file, item.marker)
		}
	}
	if strings.Contains(nextRead(t, f.logDir, "audit.change"), "No findings") {
		t.Error("audit input already contained auditor result")
	}
	auditWorkspaceGone(t, f)
}

func Test_RUN_013_AuditRetryUsesApprovedAgentConfiguration(t *testing.T) {
	f := newNextFixture(t, true, true, "normal")
	config := nextRead(t, f.root, "fireweed.yaml") + "limit_wait: 1ms\nlimit_budget: 50ms\nlimit_patterns: [RATE LIMIT]\n"
	nextWrite(t, f.root, "fireweed.yaml", config)
	nextCommit(t, f.root, "approved audit retry policy")
	alternate := filepath.Join(f.logDir, "unapproved-auditor.sh")
	unapproved := strings.Replace(config, filepath.Join(f.logDir, "auditor.sh"), alternate, 1)
	if unapproved == config {
		t.Fatal("fixture did not change the audit command")
	}
	nextWrite(t, f.logDir, "unapproved.yaml", unapproved)
	nextWrite(t, f.logDir, "held-root", f.root)
	nextWrite(t, f.logDir, "unapproved-auditor.sh", "#!/bin/sh\nprintf unapproved > \"$1/unapproved-invoked\"\nexit 9\n")
	auditBeforeEdit(t, f, `count=0
[ ! -f "$1/audit-attempts" ] || count=$(cat "$1/audit-attempts")
count=$((count + 1))
printf '%s' "$count" > "$1/audit-attempts"
if [ "$count" -eq 1 ]; then
  printf '%s' "$PWD" > "$1/first-audit.cwd"
  held=$(cat "$1/held-root")
  cp "$1/unapproved.yaml" "$held/fireweed.yaml"
  printf 'RATE LIMIT' >&2
  exit 9
fi
printf approved > "$1/approved-reused"
exit 9
`)
	r, err := f.run(false)
	if err == nil || runner.IsUsageStop(err) || len(r.Steps) != 3 || !r.Steps[1].Agent.UsageLimit || r.Steps[2].Phase != "audit" {
		t.Fatalf("approved audit retry = %#v, %v", r, err)
	}
	if got := nextRead(t, f.logDir, "approved-reused"); got != "approved" {
		t.Errorf("approved auditor was not reused: %q", got)
	}
	if _, statErr := os.Stat(filepath.Join(f.logDir, "unapproved-invoked")); !os.IsNotExist(statErr) {
		t.Errorf("unapproved auditor ran: %v", statErr)
	}
	if got := nextRead(t, f.root, "fireweed.yaml"); got != unapproved {
		t.Error("fixture did not leave the on-disk unapproved config in place")
	}
	if cwd := nextRead(t, f.logDir, "first-audit.cwd"); cwd == f.root || cwd == nextRead(t, f.logDir, "audit.cwd") {
		t.Errorf("audit retry did not use a fresh restricted workspace: first=%q last=%q", cwd, nextRead(t, f.logDir, "audit.cwd"))
	} else if _, statErr := os.Stat(cwd); !os.IsNotExist(statErr) {
		t.Errorf("first audit workspace remains: %v", statErr)
	}
	auditWorkspaceGone(t, f)
}

func Test_RUN_048_PermittedAuditEditReachesHeldChangeBeforeVerificationRetry(t *testing.T) {
	f := newNextFixture(t, true, true, "normal")
	nextWrite(t, f.logDir, "held-root", f.root)
	nextWrite(t, f.logDir, "auditor.sh", `#!/bin/bash
printf 'audit|%s|%s\n' "$PWD" "$FIREWEED_BIN" >> "$1/invocations"
printf '%s' "$PWD" > "$1/audit.cwd"
count=0
[ ! -f "$1/audit-attempts" ] || count=$(cat "$1/audit-attempts")
count=$((count + 1))
printf '%s' "$count" > "$1/audit-attempts"
if [ "$count" -eq 1 ]; then
  cat >> changes/0001-next.md <<'AUDIT'

## Audit
No findings
Checked IDs: ALP-002
AUDIT
  held=$(cat "$1/held-root")
  printf 'package internal\n' > "$held/packages/alpha/internal/audit-zone-violation.go"
else
  held=$(cat "$1/held-root")
  cp "$held/changes/0001-next.md" "$1/held-before-retry"
  cp changes/0001-next.md "$1/workspace-before-retry"
  rm "$held/packages/alpha/internal/audit-zone-violation.go"
fi
printf auditor
`)
	r, err := f.run(false)
	if err != nil || r.State.Phase != "applied" || len(r.Steps) != 4 || r.Steps[1].Phase != "audit" || r.Steps[2].Phase != "audit" {
		t.Fatalf("audit verification retry = %#v, %v", r, err)
	}
	for _, path := range []string{"held-before-retry", "workspace-before-retry"} {
		got := nextRead(t, f.logDir, path)
		if !strings.Contains(got, "No findings\nChecked IDs: ALP-002") || !strings.Contains(got, "One new value is needed.") {
			t.Errorf("first permitted Audit edit absent before retry in %s: %q", path, got)
		}
	}
	archived := nextRead(t, f.root, "changes/archive/0001-next.md")
	if !strings.Contains(archived, "No findings; checked ALP-002") {
		t.Errorf("retry did not retain Audit: %q", archived)
	}
	auditWorkspaceGone(t, f)
}

func Test_RUN_046_AuditCommitDoesNotEnterHeldGit(t *testing.T) {
	f := newNextFixture(t, true, true, "commit-success")
	auditAction(t, f, `git rev-parse HEAD > "$1/audit.commit"
`)
	base := nextGit(t, f.root, "rev-parse", "HEAD")
	r, err := f.run(false)
	if err != nil || r.State.Phase != "applied" {
		t.Fatalf("audit commit lifecycle = %#v, %v", r, err)
	}
	auditOID := strings.TrimSpace(nextRead(t, f.logDir, "audit.commit"))
	if auditOID == "" {
		t.Fatal("auditor did not commit")
	}
	if nextGit(t, f.root, "rev-list", "--count", base+"..HEAD") != "3" {
		t.Error("auditor commit entered held branch")
	}
	if out := nextGit(t, f.root, "for-each-ref", "--format=%(objectname)"); strings.Contains(out, auditOID) {
		t.Errorf("held refs contain audit commit %s", auditOID)
	}
	if out, err := auditGitObject(f.root, auditOID); err == nil {
		t.Errorf("held repository contains audit commit %s: %s", auditOID, out)
	}
	auditWorkspaceGone(t, f)
}

func Test_RUN_047_RUN_048_FirstAuditBeforeProvenanceIsImported(t *testing.T) {
	f := newNextFixture(t, true, true, "normal")
	const inserted = "## Audit\nNo findings\nChecked IDs: ALP-002\n\n"
	auditBeforeEdit(t, f, `if [ -f "$1/audit.original" ]; then
  cp changes/0001-next.md "$1/audit.retry-input"
  printf auditor
  exit 0
fi
cp changes/0001-next.md "$1/audit.original"
cp "$1/audit.cwd" "$1/first-audit.cwd"
awk '
  /^## Provenance$/ && !inserted {
    printf "## Audit\nNo findings\nChecked IDs: ALP-002\n\n"
    inserted = 1
  }
  { print }
  END { if (!inserted) exit 8 }
' "$1/audit.original" > changes/0001-next.md || exit 8
cp changes/0001-next.md "$1/audit.candidate"
printf auditor
exit 9
`)
	first, firstErr := f.run(false)
	if firstErr == nil || first.State.Phase != "audit" || first.State.ClaimOID == "" || len(first.Steps) != 2 {
		t.Fatalf("first audit failure = %#v, %v", first, firstErr)
	}
	auditWorkspaceGone(t, f)
	r, err := f.run(false)
	steps := append(first.Steps, r.Steps...)
	if err != nil || r.State.Phase != "applied" || len(steps) != 4 {
		t.Fatalf("audit before Provenance lifecycle = %#v, %v", r, err)
	}
	for i, want := range []struct {
		phase, role string
		exit        int
	}{{"evals", "worker", 0}, {"audit", "audit", 9}, {"audit", "audit", 0}, {"impl", "worker", 0}} {
		if step := steps[i]; step.Phase != want.phase || step.Role != want.role || step.Agent.ExitCode != want.exit {
			t.Errorf("step %d = %#v, want %#v", i, step, want)
		}
	}
	if cwd := nextRead(t, f.logDir, "first-audit.cwd"); cwd == nextRead(t, f.logDir, "audit.cwd") {
		t.Errorf("audit retry reused workspace %q", cwd)
	}
	original := nextRead(t, f.logDir, "audit.original")
	candidate := nextRead(t, f.logDir, "audit.candidate")
	if !strings.Contains(original, "## Provenance\n") || strings.Count(candidate, inserted) != 1 ||
		strings.Replace(candidate, inserted, "", 1) != original ||
		strings.Index(candidate, "## Audit\n") > strings.Index(candidate, "## Provenance\n") {
		t.Errorf("auditor did not insert only the first Audit before Provenance")
	}
	beforeProvenance, hasOriginalProvenance := contracts.Section(original, "Provenance")
	retryInput := nextRead(t, f.logDir, "audit.retry-input")
	retryProvenance, hasRetryProvenance := contracts.Section(retryInput, "Provenance")
	if !hasOriginalProvenance || !hasRetryProvenance || retryProvenance != beforeProvenance {
		t.Errorf("Audit copying changed raw Provenance before successful Done: got %q, want %q", retryProvenance, beforeProvenance)
	}
	if audit, ok := contracts.Section(retryInput, "Audit"); !ok || !strings.Contains(audit, "No findings\nChecked IDs: ALP-002") {
		t.Errorf("failed auditor's Audit was not imported before retry: %q", audit)
	}
	archived := nextRead(t, f.root, "changes/archive/0001-next.md")
	for _, section := range []string{"Motivation", "Interface", "Added", "Modified", "Removed", "Decisions"} {
		want, _ := contracts.Section(original, section)
		got, _ := contracts.Section(archived, section)
		if got != want {
			t.Errorf("held %s changed outside Audit: got %q, want %q", section, got, want)
		}
	}
	afterProvenance, hasProvenance := contracts.Section(archived, "Provenance")
	nonblankLines := func(section string) []string {
		var lines []string
		for _, line := range strings.Split(section, "\n") {
			if strings.TrimSpace(line) != "" {
				lines = append(lines, line)
			}
		}
		return lines
	}
	beforeLines, afterLines := nonblankLines(beforeProvenance), nonblankLines(afterProvenance)
	preserved := hasProvenance && len(afterLines) >= len(beforeLines)
	for i, line := range beforeLines {
		if i >= len(afterLines) || afterLines[i] != line {
			preserved = false
			break
		}
	}
	if !preserved {
		t.Errorf("held Provenance lost existing contents: %q", afterProvenance)
	}
	if audit, ok := contracts.Section(archived, "Audit"); !ok || !strings.Contains(audit, "No findings; checked ALP-002") {
		t.Errorf("inserted Audit was not imported: %q", audit)
	}
	auditWorkspaceGone(t, f)
}

func Test_RUN_045_RUN_049_SeparateAuditsUseFreshWorkspaces(t *testing.T) {
	f := newRunFixture(t, 2, false, false)
	timeFastHook(t, f)
	r, err := runner.Run(context.Background(), f.request(0, false))
	if err != nil || len(r.Applied) != 2 {
		t.Fatalf("two-change audit lifecycle: applied=%d error=%v", len(r.Applied), err)
	}
	var cwds []string
	for _, line := range runInvocations(t, f) {
		parts := strings.Split(line, "|")
		if len(parts) == 4 && parts[1] == "audit" {
			cwds = append(cwds, parts[2])
		}
	}
	if len(cwds) != 2 {
		t.Fatalf("audit invocation roots = %#v", cwds)
	}
	if cwds[0] == cwds[1] || cwds[0] == f.root || cwds[1] == f.root {
		t.Errorf("audits shared a workspace or used held checkout: %#v", cwds)
	}
	for _, cwd := range cwds {
		if _, err := os.Stat(cwd); !os.IsNotExist(err) {
			t.Errorf("audit workspace %q remains: %v", cwd, err)
		}
	}
}

func Test_RUN_047_InvalidAuditWorkspaceEditsAreRejectedAtomically(t *testing.T) {
	for _, tc := range []struct{ name, action string }{
		{"change outside Audit", `sed 's/One new value is needed./Altered motivation./' changes/0001-next.md > "$1/altered-change"; cp "$1/altered-change" changes/0001-next.md`},
		{"projected spec", `printf '\nAltered spec\n' >> packages/alpha/SPEC.md`},
		{"eval", `printf '\n// altered eval\n' >> packages/alpha/evals/next_test.go`},
		{"added file", `printf 'extra\n' > unexpected.txt`},
		{"deleted file", `rm packages/alpha/evals/existing_test.go`},
		{"committed invalid edit", `printf '\nAltered spec\n' >> packages/alpha/SPEC.md; git add -A; git -c user.name=Fixture -c user.email=fixture@example.invalid commit -q -m invalid-audit-edit`},
	} {
		t.Run(tc.name, func(t *testing.T) {
			f := newNextFixture(t, true, true, "normal")
			nextAttempts(t, f, 1)
			spec := nextRead(t, f.root, "packages/alpha/SPEC.md")
			existing := nextRead(t, f.root, "packages/alpha/evals/existing_test.go")
			auditAction(t, f, tc.action+"\n")
			r, err := f.run(false)
			if err == nil || len(r.Steps) != 2 || r.State.Phase != "audit" || r.State.ClaimOID == "" {
				t.Errorf("invalid audit was not rejected while retaining claim: phase=%q claim=%q steps=%d error=%v", r.State.Phase, r.State.ClaimOID, len(r.Steps), err)
			}
			heldBytes, readErr := os.ReadFile(filepath.Join(f.root, nextChange))
			if readErr != nil {
				t.Fatalf("rejected edit removed held change: %v", readErr)
			}
			held := string(heldBytes)
			if strings.Contains(held, "No findings") || strings.Contains(held, "Altered motivation") || !strings.Contains(held, "One new value is needed.") {
				t.Errorf("rejected edit reached held change: %q", held)
			}
			if got := nextRead(t, f.root, "packages/alpha/SPEC.md"); got != spec {
				t.Error("rejected audit edit reached held spec")
			}
			if got := nextRead(t, f.root, "packages/alpha/evals/existing_test.go"); got != existing {
				t.Error("rejected audit edit reached held eval")
			}
			if strings.Contains(nextRead(t, f.root, "packages/alpha/evals/next_test.go"), "altered eval") {
				t.Error("rejected audit edit reached held new eval")
			}
			if _, statErr := os.Stat(filepath.Join(f.root, "unexpected.txt")); !os.IsNotExist(statErr) {
				t.Errorf("rejected file reached held checkout: %v", statErr)
			}
			if got := nextInvocations(t, f); len(got) != 2 {
				t.Errorf("invalid audit advanced phase: %#v", got)
			}
			auditWorkspaceGone(t, f)
		})
	}
}

func Test_RUN_048_RUN_049_PermittedAuditWorkSurvivesFailureAndStops(t *testing.T) {
	for _, tc := range []struct {
		name, action, policy string
		timeStop             bool
	}{
		{"agent failure", "printf 'partial audit'; exit 9\n", "", false},
		{"usage stop", "printf 'RATE LIMIT' >&2; exit 9\n", "limit_patterns: [RATE LIMIT]\nlimit_budget: 0s\n", false},
		{"time stop", "sleep 6\n", "", true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			f := newNextFixture(t, true, true, "normal")
			if tc.policy != "" {
				nextWrite(t, f.root, "fireweed.yaml", nextRead(t, f.root, "fireweed.yaml")+tc.policy)
				nextCommit(t, f.root, "audit limit policy")
			}
			auditAction(t, f, tc.action)
			base := nextGit(t, f.root, "rev-parse", "HEAD")
			var r runner.NextResult
			var err error
			if tc.timeStop {
				request := runner.RunRequest{Root: f.root, Bin: "/fixture/next-bin", For: 3 * time.Second, Timeout: 10 * time.Second}
				prepareHeldAudit(t, request)
				result, runErr := runner.Run(context.Background(), request)
				err = runErr
				if len(result.Applied) != 0 {
					t.Errorf("time stop applied %d changes", len(result.Applied))
				}
			} else {
				r, err = f.run(false)
			}
			if err == nil || tc.timeStop && !runner.IsTimeStop(err) || tc.name == "usage stop" && !runner.IsUsageStop(err) {
				t.Errorf("audit failure/stop = %v", err)
			}
			if !tc.timeStop && (len(r.Steps) != 2 || r.State.Phase != "audit") {
				t.Errorf("audit failure result = %#v", r)
			}
			held := nextRead(t, f.root, nextChange)
			if !strings.Contains(held, "No findings") || !strings.Contains(held, "One new value is needed.") || !strings.Contains(held, "phase: audit") {
				t.Errorf("partial Audit not retained in held change: %q", held)
			}
			if got := nextGit(t, f.root, "rev-list", "--count", base+"..HEAD"); got != "1" {
				t.Errorf("audit error advanced phases %s times, want only evals", got)
			}
			if state, stateErr := workflow.Current(f.root, "next"); stateErr != nil || state.Phase != "audit" || state.ClaimOID == "" {
				t.Errorf("audit error released or advanced phase: %#v, %v", state, stateErr)
			}
			if got := nextInvocations(t, f); len(got) != 2 {
				t.Errorf("audit error advanced phase: %#v", got)
			}
			auditWorkspaceGone(t, f)
		})
	}
}

// Complete evals before starting the clock so Run can exercise an interrupted auditor.
func prepareHeldAudit(t *testing.T, request runner.RunRequest) {
	t.Helper()
	ctx := context.Background()
	state, err := workflow.Start(request.Root, nextChange)
	if err != nil {
		t.Fatalf("start fixture evals: %v", err)
	}
	prompt, err := workflow.Context(request.Root, nextChange, true)
	if err != nil {
		t.Fatalf("prepare fixture evals context: %v", err)
	}
	agent, err := runner.InvokeAgent(ctx, runner.AgentRequest{
		Snapshot: state.Snapshot, Role: "worker", Prompt: prompt,
		Bin: request.Bin, Timeout: request.Timeout, Run: state.Run, Phase: state.Phase,
	})
	if err != nil || agent.ExitCode != 0 {
		t.Fatalf("invoke fixture evals worker: %#v, %v", agent, err)
	}
	done, err := workflow.Done(ctx, workflow.Request{
		Root: request.Root, Change: nextChange, Bin: request.Bin, Agent: "bash", Timeout: request.Timeout,
	})
	if err != nil || done.State.Phase != "audit" || done.State.ClaimOID == "" {
		t.Fatalf("prepare held audit: phase=%q claim=%q error=%v", done.State.Phase, done.State.ClaimOID, err)
	}
}
