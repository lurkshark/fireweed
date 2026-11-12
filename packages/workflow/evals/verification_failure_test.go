package evals_test

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/workflow"
)

func Test_FLOW_136_DoneMarksFullVerificationFailureAndPreservesResult(t *testing.T) {
	root := fixture(t, "audit:\n  enabled: false\n", "", "", "", "")
	hook := filepath.Join(root, ".fireweed", "missing-boundary-hook")
	config := readFile(t, root, "fireweed.yaml")
	write(t, root, "fireweed.yaml", strings.Replace(config,
		"boundaries: [sh, .fireweed/boundaries.sh]", fmt.Sprintf("boundaries: [%q]", hook), 1))
	commit(t, root, "approve missing boundary executable fixture")
	start(t, root)
	if result, err := done(root); err != nil || result.State.Phase != "impl" {
		t.Fatalf("fixture evals Done = phase %q, error %v", result.State.Phase, err)
	}
	write(t, root, "packages/alpha/internal/ready", "ready\n")

	verified, verifyErr := verify(root, false)
	if verifyErr == nil || len(verified.Reports["alpha"].Cases) == 0 ||
		!containsCode(verified.Diagnostics, "X02") {
		t.Fatalf("fixture full Verify = reports %v, diagnostics %v, error %v", verified.Reports, verified.Diagnostics, verifyErr)
	}
	assertErrorCode(t, "full Verify", verifyErr, "X02")
	if workflow.IsVerificationFailure(verifyErr) {
		t.Errorf("direct Verify error was marked: %v", verifyErr)
	}

	result, err := done(root)
	if err == nil {
		t.Fatal("Done accepted the failed boundary hook")
	}
	if !workflow.IsVerificationFailure(err) {
		t.Errorf("Done's full verification error was not marked: %v", err)
	}
	if wrapped := fmt.Errorf("finish attempt: %w", err); !workflow.IsVerificationFailure(wrapped) {
		t.Errorf("wrapped Done verification error was not recognized: %v", wrapped)
	}
	assertErrorCode(t, "Done", err, "X02")
	_, wantSubject, _ := contracts.ErrorCode(verifyErr)
	if _, subject, _ := contracts.ErrorCode(err); subject != wantSubject {
		t.Errorf("Done error subject = %q, want %q", subject, wantSubject)
	}
	if err.Error() != verifyErr.Error() {
		t.Errorf("Done error text = %q, want %q", err.Error(), verifyErr.Error())
	}
	var cause *os.PathError
	if !errors.Is(err, os.ErrNotExist) || !errors.As(err, &cause) || cause.Path != hook {
		t.Errorf("Done error %v lost the missing executable cause for %q", err, hook)
	}
	if !reflect.DeepEqual(result, verified) {
		t.Errorf("Done verification result = %#v, want %#v", result, verified)
	}
}

func Test_FLOW_136_DoneDistinguishesSameCodeByOrigin(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	hook := readFile(t, root, ".fireweed/eval.sh")
	write(t, root, ".fireweed/eval.sh", "printf 'protected edit\\n' >> SYSTEM.md\n"+hook)
	commit(t, root, "approve hook that edits a protected file")

	_, err := done(root)
	assertErrorCode(t, "unclaimed Done", err, "P02")
	if workflow.IsVerificationFailure(err) {
		t.Errorf("Done precondition error was marked: %v", err)
	}

	start(t, root)
	result, err := done(root)
	assertErrorCode(t, "Done after hook edits SYSTEM.md", err, "Z01")
	if len(result.Reports["alpha"].Cases) == 0 || !containsCode(result.Diagnostics, "Z01") {
		t.Fatalf("expected a post-verification zone rejection with a completed report, got %v and %v", result.Reports, result.Diagnostics)
	}
	if workflow.IsVerificationFailure(err) {
		t.Errorf("Done phase-transition error was marked: %v", err)
	}

	// The edit now exists before the next Verify, so that call rejects it before the hook runs.
	result, err = done(root)
	assertErrorCode(t, "Done with existing SYSTEM.md edit", err, "Z01")
	if len(result.Reports) != 0 || !containsCode(result.Diagnostics, "Z01") {
		t.Fatalf("expected a verification zone rejection before the hook, got %v and %v", result.Reports, result.Diagnostics)
	}
	if !workflow.IsVerificationFailure(err) {
		t.Errorf("Done verification error with the same code was not marked: %v", err)
	}
}
