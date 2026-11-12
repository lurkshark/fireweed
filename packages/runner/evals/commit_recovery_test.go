package evals_test

import (
	"slices"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow"
)

func assertRecoveryWarnings(t *testing.T, diagnostics []project.Diagnostic, phases ...string) {
	t.Helper()
	var warnings []project.Diagnostic
	for _, diagnostic := range diagnostics {
		if diagnostic.Code == "P05" {
			warnings = append(warnings, diagnostic)
		}
	}
	if len(warnings) != len(phases) {
		t.Errorf("P05 recovery warnings = %#v, want one per phase %v", warnings, phases)
		return
	}
	for i, phase := range phases {
		message := strings.ToLower(warnings[i].Message)
		if !warnings[i].Warning || !strings.Contains(message, phase) || !strings.Contains(message, "reset") || !strings.Contains(message, "retain") {
			t.Errorf("%s recovery warning = %#v", phase, warnings[i])
		}
	}
}

func Test_RUN_027_RUN_028_RecoversAgentCommitsAcrossLifecycle(t *testing.T) {
	f := newNextFixture(t, true, true, "commit-success")
	base := nextGit(t, f.root, "rev-parse", "HEAD")
	r, err := f.run(false)
	if err != nil {
		t.Fatalf("Next after fixture agent commits: %v; state=%#v; steps=%#v", err, r.State, r.Steps)
	}
	if r.Empty || r.State.Phase != "applied" || r.State.Run == "" || r.State.ClaimOID != "" || len(r.Steps) != 3 {
		t.Fatalf("incomplete lifecycle after agent commits: %#v", r)
	}
	for i, want := range []struct{ phase, role, stdout string }{{"evals", "worker", "worker-evals"}, {"audit", "audit", "auditor"}, {"impl", "worker", "worker-impl"}} {
		step := r.Steps[i]
		if step.Phase != want.phase || step.Role != want.role || step.Agent.Stdout != want.stdout || step.Agent.ExitCode != 0 {
			t.Errorf("step %d = %#v", i, step)
		}
	}
	for _, phase := range []string{"evals", "impl"} {
		found := false
		for _, diagnostic := range r.State.Diagnostics {
			if diagnostic.Code == "P05" && diagnostic.Warning && strings.Contains(strings.ToLower(diagnostic.Message), phase) && strings.Contains(strings.ToLower(diagnostic.Message), "retain") {
				found = true
			}
		}
		if !found {
			t.Errorf("missing %s worker recovery warning: %#v", phase, r.State.Diagnostics)
		}
	}
	if got := strings.TrimSpace(nextRead(t, f.logDir, "evals-reset-head")); got != base {
		t.Errorf("evals verification HEAD = %s, want held base %s", got, base)
	}
	for _, observation := range []struct{ file, path string }{
		{"evals-staged", "packages/alpha/evals/staged_test.go"},
		{"evals-staged", "packages/alpha/evals/second_test.go"},
		{"evals-unstaged", "packages/alpha/evals/next_test.go"},
		{"evals-untracked", "packages/alpha/evals/untracked_test.go"},
	} {
		if !strings.Contains(nextRead(t, f.logDir, observation.file), observation.path) {
			t.Errorf("%s omitted retained %s", observation.file, observation.path)
		}
	}
	if count := nextGit(t, f.root, "rev-list", "--count", base+"..HEAD"); count != "3" {
		t.Errorf("branch has %s commits after three checked phases, want 3", count)
	}
	commits := strings.Fields(nextGit(t, f.root, "rev-list", "--reverse", base+"..HEAD"))
	phaseCommits, commitsErr := workflow.Commits(f.root, "next")
	if commitsErr != nil || len(commits) != 3 || !slices.Equal(phaseCommits, []string{commits[2], commits[1], commits[0]}) {
		t.Fatalf("checked phase commits = %v; branch commits = %v; error = %v", phaseCommits, commits, commitsErr)
	}
	for i, want := range []struct{ path, phase string }{
		{nextChange, "audit"},
		{nextChange, "impl"},
		{"changes/archive/0001-next.md", "applied"},
	} {
		document, parseErr := contracts.ParseDocument(nextGit(t, f.root, "show", commits[i]+":"+want.path))
		if parseErr != nil || document.Frontmatter["phase"] != want.phase {
			t.Errorf("phase commit %s phase = %v, parse error = %v; want %s", commits[i], document.Frontmatter["phase"], parseErr, want.phase)
		}
	}
	fixtureCommits := strings.Split(strings.TrimSpace(nextRead(t, f.logDir, "fixture-commits")), "\n")
	reachable := strings.Fields(nextGit(t, f.root, "rev-list", "HEAD"))
	for i, label := range []string{"evals-first", "evals-second", "audit", "impl-first"} {
		if i >= len(fixtureCommits) {
			t.Errorf("fixture agent commit %s was not recorded", label)
			continue
		}
		parts := strings.Fields(fixtureCommits[i])
		if len(parts) != 2 || parts[0] != label || len(parts[1]) != 40 {
			t.Errorf("fixture agent commit %d = %q, want %s and a full OID", i, fixtureCommits[i], label)
			continue
		}
		if slices.Contains(reachable, parts[1]) {
			t.Errorf("fixture agent commit %s (%s) is reachable from final HEAD", label, parts[1])
		}
	}
	if len(fixtureCommits) != 4 {
		t.Errorf("fixture agent commits = %q, want four", fixtureCommits)
	}
	archive, parseErr := contracts.ParseDocument(nextRead(t, f.root, "changes/archive/0001-next.md"))
	if parseErr != nil || archive.Frontmatter["phase"] != "applied" {
		t.Errorf("archive phase = %v, parse error = %v; want applied", archive.Frontmatter["phase"], parseErr)
	}
	for _, marker := range []string{"No findings", "- evals ", "- audit ", "- impl "} {
		if !strings.Contains(nextRead(t, f.root, "changes/archive/0001-next.md"), marker) {
			t.Errorf("archive omitted %q", marker)
		}
	}
	for _, marker := range []string{"EVAL_CONTENT_MARKER", "TRACKED_EVAL_WORK"} {
		if !strings.Contains(nextRead(t, f.root, "packages/alpha/evals/next_test.go"), marker) {
			t.Errorf("retained eval file omitted %q", marker)
		}
	}
	if !strings.Contains(nextRead(t, f.root, "packages/alpha/evals/second_test.go"), "SECOND_AGENT_COMMIT") || !strings.Contains(nextRead(t, f.root, "packages/alpha/evals/staged_test.go"), "STAGED_EVAL_WORK") || !strings.Contains(nextRead(t, f.root, "packages/alpha/evals/untracked_test.go"), "UNTRACKED_EVAL_WORK") {
		t.Error("archive lost eval work after reset")
	}
	if !strings.Contains(nextRead(t, f.root, "packages/alpha/internal/value.go"), `return "new"`) || !strings.Contains(nextRead(t, f.root, "packages/alpha/SPEC.md"), "Alpha MUST return the new value.") {
		t.Error("checked implementation or projected spec missing")
	}
	if refs := nextGit(t, f.root, "for-each-ref", "--format=%(refname)", "refs/fireweed/claims"); refs != "" {
		t.Errorf("claim retained after archive: %s", refs)
	}
	if _, err := workflow.Current(f.root, "next"); err == nil {
		t.Error("archived change remains current")
	}
}

func Test_RUN_027_RUN_028_RecoversCommitBeforeInvocationFailure(t *testing.T) {
	f := newNextFixture(t, true, true, "commit-exit")
	base := nextGit(t, f.root, "rev-parse", "HEAD")
	r, err := f.run(false)
	if err == nil || len(r.Steps) != 1 {
		t.Fatalf("agent failure = %v, steps=%#v", err, r.Steps)
	}
	step := r.Steps[0]
	if step.Phase != "evals" || step.Role != "worker" || step.Agent.Stdout != "partial worker output" || step.Agent.Stderr != "worker failure" || step.Agent.ExitCode != 9 {
		t.Errorf("captured failing agent = %#v", step)
	}
	if head := nextGit(t, f.root, "rev-parse", "HEAD"); head != base {
		t.Errorf("failure HEAD = %s, want held base %s", head, base)
	}
	state, currentErr := workflow.Current(f.root, "next")
	if currentErr != nil || state.Phase != "evals" || r.State.Phase != "evals" || state.Run == "" || r.State.Run != state.Run || state.ClaimOID == "" {
		t.Errorf("failure did not retain held phase: current=%#v, result=%#v, err=%v", state, r.State, currentErr)
	}
	if claim := nextGit(t, f.root, "for-each-ref", "--format=%(refname)", "refs/fireweed/claims"); claim != "refs/fireweed/claims/next" {
		t.Errorf("held claim = %q", claim)
	}
	assertRecoveryWarnings(t, r.State.Diagnostics, "evals")
	if staged := nextGit(t, f.root, "diff", "--cached", "--name-only"); !strings.Contains(staged, "packages/alpha/evals/next_test.go") || !strings.Contains(staged, "packages/alpha/evals/staged_test.go") {
		t.Errorf("reset did not retain index work: %q", staged)
	}
	if unstaged := nextGit(t, f.root, "diff", "--name-only"); !strings.Contains(unstaged, "packages/alpha/evals/next_test.go") {
		t.Errorf("reset did not retain tracked work: %q", unstaged)
	}
	if untracked := nextGit(t, f.root, "ls-files", "--others", "--exclude-standard"); !strings.Contains(untracked, "packages/alpha/evals/untracked_test.go") {
		t.Errorf("reset did not retain untracked work: %q", untracked)
	}
	if len(nextInvocations(t, f)) != 1 {
		t.Error("failed invocation retried or advanced")
	}
}
