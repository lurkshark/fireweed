package evals_test

import (
	"errors"
	"regexp"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/workflow"
)

const remoteMain = "refs/remotes/origin/main"
const localMain = "refs/heads/main"

func newerBaseTip(t *testing.T, root string) string {
	t.Helper()
	return git(t, root, "commit-tree", "HEAD^{tree}", "-p", "HEAD", "-m", "base-only commit")
}

func Test_FLOW_141_StartWarnsWhenBaseTipIsOutsideHead(t *testing.T) {
	for _, resume := range []bool{false, true} {
		t.Run(acquisitionName(resume)+"/behind remote", func(t *testing.T) {
			root, run := baseTipFixture(t, resume)
			git(t, root, "update-ref", remoteMain, newerBaseTip(t, root))
			assertBaseTipWarning(t, acquireAtBaseTip(t, root, run), remoteMain)
		})
		t.Run(acquisitionName(resume)+"/diverged from local fallback", func(t *testing.T) {
			root, run := baseTipFixture(t, resume)
			git(t, root, "update-ref", localMain, newerBaseTip(t, root))
			git(t, root, "commit", "--allow-empty", "-m", "feature-only commit")
			assertBaseTipWarning(t, acquireAtBaseTip(t, root, run), localMain)
		})
	}
}

func Test_FLOW_142_StartOmitsWarningWhenHeadContainsBaseTip(t *testing.T) {
	for _, resume := range []bool{false, true} {
		t.Run(acquisitionName(resume)+"/equal to remote", func(t *testing.T) {
			root, run := baseTipFixture(t, resume)
			git(t, root, "update-ref", remoteMain, head(t, root))
			git(t, root, "update-ref", localMain, newerBaseTip(t, root))
			state := acquireAtBaseTip(t, root, run)
			if diagnostic, found := warningFor(state.Diagnostics, "P05"); found {
				t.Errorf("P05 with HEAD equal to %s: %#v", remoteMain, diagnostic)
			}
		})
		t.Run(acquisitionName(resume)+"/ahead of local fallback", func(t *testing.T) {
			root, run := baseTipFixture(t, resume)
			git(t, root, "commit", "--allow-empty", "-m", "feature-only commit")
			state := acquireAtBaseTip(t, root, run)
			if diagnostic, found := warningFor(state.Diagnostics, "P05"); found {
				t.Errorf("P05 with HEAD ahead of %s: %#v", localMain, diagnostic)
			}
		})
	}
}

func acquisitionName(resume bool) string {
	if resume {
		return "StartFromAttempt"
	}
	return "Start"
}

func baseTipFixture(t *testing.T, resume bool) (root, run string) {
	t.Helper()
	if resume {
		var started workflow.State
		root, started, _ = savedEvalsAttempt(t)
		unblockAttempt(t, root)
		run = started.Run
	} else {
		root = fixture(t, "", "", "### new:extra\nExtra MUST be checked.\n", "", "")
	}
	git(t, root, "checkout", "-b", "feature")
	return root, run
}

func acquireAtBaseTip(t *testing.T, root, existingRun string) workflow.State {
	t.Helper()
	before := head(t, root)
	var state workflow.State
	if existingRun == "" {
		state = start(t, root)
	} else {
		var err error
		state, err = workflow.StartFromAttempt(root, "manual", 1)
		if err != nil {
			t.Fatalf("StartFromAttempt: %v", err)
		}
		if got := readFile(t, root, "packages/alpha/evals/new_test.go"); got != "saved new eval\n" {
			t.Errorf("restored eval = %q, want saved new eval", got)
		}
	}
	if containsCode(state.Diagnostics, "P04") {
		t.Errorf("approved change received P04: %#v", state.Diagnostics)
	}
	if state.Phase != "evals" || state.Head != before || state.PhaseBase != before || head(t, root) != before {
		t.Errorf("acquired phase/head/base = %q/%s/%s; want evals at unchanged HEAD %s", state.Phase, state.Head, state.PhaseBase, before)
	}
	if state.Run == "" || existingRun != "" && state.Run != existingRun {
		t.Errorf("acquired run = %q, previous run = %q", state.Run, existingRun)
	}
	if len(state.Change.Added) != 1 || state.Change.Added[0].ID != "ALP-003" {
		t.Errorf("acquired Added = %#v; want ALP-003", state.Change.Added)
	}
	current, err := workflow.Current(root, "manual")
	if err != nil || state.ClaimOID == "" || current.ClaimOID != state.ClaimOID {
		t.Errorf("Current claim = %q, error %v; want owned claim %q", current.ClaimOID, err, state.ClaimOID)
	}
	return state
}

func assertBaseTipWarning(t *testing.T, state workflow.State, baseRef string) {
	t.Helper()
	diagnostic, found := warningFor(state.Diagnostics, "P05")
	message := strings.ToLower(diagnostic.Message)
	stale := strings.Contains(message, "stale") || strings.Contains(message, "outdated") || strings.Contains(message, "out of date")
	if !found || !diagnostic.Warning || !strings.Contains(diagnostic.Message, baseRef) || !strings.Contains(message, "queue") || !stale {
		t.Errorf("want P05 warning naming %s and explaining the queue may be stale, got %#v", baseRef, state.Diagnostics)
	}
}

func Test_FLOW_058_StartRejectsStaleSpecHashWithNewerBaseTip(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	change := readFile(t, root, pendingCommitChange)
	stale := regexp.MustCompile(`(?m)^  ALP-001: [0-9a-f]+$`).ReplaceAllString(change, "  ALP-001: deadbeef")
	write(t, root, pendingCommitChange, stale)
	commit(t, root, "stale approved spec hash")
	git(t, root, "update-ref", remoteMain, newerBaseTip(t, root))
	before := head(t, root)
	_, err := workflow.Start(root, "manual")
	assertErrorCode(t, "Start with stale spec hash and newer base tip", err, "P05")
	code, _, _ := contracts.ErrorCode(errors.Unwrap(err))
	if code != "C08" || !strings.Contains(err.Error(), "ALP-001") {
		t.Errorf("Start error = %v, want P05 retaining the C08 cause for ALP-001", err)
	}
	assertNoResumeMutation(t, root, before, stale)
}
