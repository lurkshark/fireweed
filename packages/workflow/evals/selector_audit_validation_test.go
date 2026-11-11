package evals_test

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/workflow"
)

func Test_FLOW_059_AuditVerifyChecksSubmissionWithoutAdvancing(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	toAudit(t, root)
	base := head(t, root)
	current, err := workflow.Current(root, "manual")
	if err != nil {
		t.Fatal(err)
	}
	for _, tc := range []struct{ body, want string }{
		{"No findings\n", "P09: audit requires Checked IDs and No findings"},
		{"Checked IDs: ALP-002\nNo findings\n", "P09: audit did not check ALP-001"},
	} {
		audit(t, root, tc.body)
		before := readFile(t, root, pendingCommitChange)
		for _, quick := range []bool{true, false} {
			result, err := verify(root, quick)
			assertErrorCode(t, "audit Verify", err, "P09")
			if err == nil || !strings.HasPrefix(err.Error(), tc.want) {
				t.Errorf("audit Verify(quick=%t, %q) = phase %q, error %v; want %s", quick, tc.body, result.State.Phase, err, tc.want)
			}
		}
		if got := readFile(t, root, pendingCommitChange); got != before {
			t.Errorf("rejected Verify changed audit submission: %q", got)
		}
		if head(t, root) != base {
			t.Error("rejected Verify committed")
		}
		if now, err := workflow.Current(root, "manual"); err != nil || now.ClaimOID != current.ClaimOID || now.Phase != "audit" || now.PhaseBase != base {
			t.Errorf("rejected Verify changed claim or phase: %#v, %v", now, err)
		}
	}

	// Findings that Done can process must also pass Verify without being filtered or written.
	audit(t, root, "- ALP-001: \""+newClause+"\" — missing assertion\n- ALP-999: \"Unknown MUST work.\" — invalid finding\n")
	beforeFinding := readFile(t, root, pendingCommitChange)
	if result, err := verify(root, true); err != nil || result.State.Phase != "audit" {
		t.Errorf("finding audit Verify = phase %q, error %v", result.State.Phase, err)
	}
	if got := readFile(t, root, pendingCommitChange); got != beforeFinding || head(t, root) != base {
		t.Error("finding audit Verify normalized submission or committed")
	}

	// A complete submission is accepted by Verify without Done's normalization.
	audit(t, root, "Checked IDs: ALP-001\nNo findings\n")
	before := readFile(t, root, pendingCommitChange)
	if result, err := verify(root, true); err != nil || result.State.Phase != "audit" {
		t.Errorf("complete audit Verify = phase %q, error %v", result.State.Phase, err)
	}
	if got := readFile(t, root, pendingCommitChange); got != before || head(t, root) != base {
		t.Error("successful audit Verify wrote the submission or committed")
	}
	if now, err := workflow.Current(root, "manual"); err != nil || now.ClaimOID != current.ClaimOID || now.Phase != "audit" || now.PhaseBase != base {
		t.Errorf("successful Verify changed claim or phase: %#v, %v", now, err)
	}
	if result, err := done(root); err != nil || result.State.Phase != "impl" {
		t.Errorf("complete audit Done = phase %q, error %v", result.State.Phase, err)
	}
}

func Test_FLOW_002_AuditRejectsForbiddenEditBeforeHooks(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	marker := hookMarker(t, root)
	auditState := toAudit(t, root)
	if auditState.Phase != "audit" {
		t.Fatalf("held phase = %q, want audit", auditState.Phase)
	}
	if err := os.Remove(marker); err != nil {
		t.Fatalf("clear evals Done hook marker: %v", err)
	}
	audit(t, root, "Checked IDs: ALP-001\nNo findings\n")
	submission := readFile(t, root, pendingCommitChange)
	forbidden := "packages/alpha/evals/audit_zone_test.go"
	write(t, root, forbidden, "package alpha_test\n")
	base := head(t, root)

	for _, operation := range []struct {
		name string
		call func() (workflow.Result, error)
	}{
		{"quick Verify", func() (workflow.Result, error) { return verify(root, true) }},
		{"full Verify", func() (workflow.Result, error) { return verify(root, false) }},
		{"Done", func() (workflow.Result, error) { return done(root) }},
	} {
		result, err := operation.call()
		if err == nil || !strings.HasPrefix(err.Error(), "Z01:") || !strings.Contains(err.Error(), forbidden) {
			t.Errorf("%s = phase %q, error %v; want Z01 naming %s", operation.name, result.State.Phase, err, forbidden)
		}
		if got := head(t, root); got != base {
			t.Errorf("%s advanced HEAD: %s -> %s", operation.name, base, got)
		}
		if got := readFile(t, root, pendingCommitChange); got != submission {
			t.Errorf("%s changed the valid audit submission", operation.name)
		}
		current, currentErr := workflow.Current(root, "manual")
		if currentErr != nil || current.Phase != "audit" || current.PhaseBase != base || current.ClaimOID != auditState.ClaimOID {
			t.Errorf("%s changed held audit phase or claim: %#v, %v", operation.name, current, currentErr)
		}
		if _, markerErr := os.Stat(marker); !os.IsNotExist(markerErr) {
			t.Errorf("%s ran the eval hook before zone rejection: %v", operation.name, markerErr)
		}
	}
}

func Test_FLOW_061_UnknownLifecycleSelectorsNameValidSlugs(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	request := workflow.Request{Root: root, Change: "ghost"}
	operations := []struct {
		name string
		call func(string) error
	}{
		{"Start", func(s string) error { _, err := workflow.Start(root, s); return err }},
		{"StartFromAttempt", func(s string) error { _, err := workflow.StartFromAttempt(root, s, 1); return err }},
		{"Current", func(s string) error { _, err := workflow.Current(root, s); return err }},
		{"Context", func(s string) error { _, err := workflow.Context(root, s, false); return err }},
		{"Verify", func(s string) error {
			request.Change = s
			_, err := workflow.Verify(context.Background(), request)
			return err
		}},
		{"Done", func(s string) error {
			request.Change = s
			_, err := workflow.Done(context.Background(), request)
			return err
		}},
		{"Release", func(s string) error { return workflow.Release(root, s) }},
		{"Commits", func(s string) error { _, err := workflow.Commits(root, s); return err }},
		{"Abandon", func(s string) error { _, err := workflow.Abandon(root, s); return err }},
	}
	before := head(t, root)
	for _, selector := range []string{"ghost", "0009-ghost.md"} {
		for _, op := range operations {
			err := op.call(selector)
			if err == nil || !strings.HasPrefix(err.Error(), "U02:") || !strings.Contains(err.Error(), selector) || !strings.Contains(err.Error(), "manual") {
				t.Errorf("%s(%q) = %v; want U02 with rejected selector and valid slug manual", op.name, selector, err)
			}
		}
	}
	if head(t, root) != before || git(t, root, "status", "--porcelain=v1", "--untracked-files=all") != "" {
		t.Error("unknown selector changed HEAD or working tree")
	}

	// Commits and Abandon also look up archived documents.
	archive := filepath.Join(root, archivedCommitChange)
	if err := os.MkdirAll(filepath.Dir(archive), 0755); err != nil {
		t.Fatal(err)
	}
	git(t, root, "mv", pendingCommitChange, archivedCommitChange)
	commit(t, root, "archive fixture change")
	for _, op := range operations[len(operations)-2:] {
		err := op.call("ghost")
		if err == nil || !strings.HasPrefix(err.Error(), "U02:") || !strings.Contains(err.Error(), "ghost") || !strings.Contains(err.Error(), "manual") {
			t.Errorf("archived %s(ghost) = %v; want U02 with archive slug manual", op.name, err)
		}
	}
}
