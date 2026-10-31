package evals_test

import (
	"errors"
	"testing"

	"github.com/lurkshark/fireweed/packages/runner"
)

func Test_RUN_041_UnapprovedChangeRequiresApproval(t *testing.T) {
	f := newNextFixture(t, true, false, "normal")
	_, err := f.run(false)
	if err == nil {
		t.Fatal("Next accepted a change absent from the local and shared base")
	}
	if !runner.IsApprovalRequired(err) {
		t.Errorf("IsApprovalRequired(%v) = false", err)
	}
	f.noAgent(t)
}

func Test_RUN_042_OtherNextErrorsDoNotRequireApproval(t *testing.T) {
	if runner.IsApprovalRequired(nil) {
		t.Error("nil error requires approval")
	}
	if runner.IsApprovalRequired(errors.New("requires approval")) {
		t.Error("unrelated error with approval wording requires approval")
	}

	t.Run("start_failure", func(t *testing.T) {
		f := newNextFixture(t, true, true, "normal")
		nextWrite(t, f.root, "stray.txt", "dirty\n")
		_, err := f.run(false)
		if err == nil {
			t.Fatal("expected dirty-worktree start failure")
		}
		if runner.IsApprovalRequired(err) {
			t.Errorf("start failure requires approval: %v", err)
		}
		f.noAgent(t)
	})

	t.Run("agent_failure", func(t *testing.T) {
		f := newNextFixture(t, true, true, "exit")
		result, err := f.run(false)
		if err == nil || len(result.Steps) != 1 || result.Steps[0].Agent.ExitCode != 9 {
			t.Fatalf("expected agent exit 9, got %#v, %v", result, err)
		}
		if runner.IsApprovalRequired(err) {
			t.Errorf("agent failure requires approval: %v", err)
		}
	})
}
