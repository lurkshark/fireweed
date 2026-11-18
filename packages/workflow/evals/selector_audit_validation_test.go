package evals_test

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"slices"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow"
)

func Test_FLOW_059_FLOW_145_AuditVerifyChecksSubmissionWithoutAdvancing(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	toAudit(t, root)
	base := head(t, root)
	approved := readFile(t, root, pendingCommitChange)
	current, err := workflow.Current(root, "manual")
	if err != nil {
		t.Fatal(err)
	}
	for _, tc := range []struct{ body, want string }{
		{"No findings\n", "audit requires Checked IDs and No findings"},
		{"Checked IDs: ALP-002\nNo findings\n", "audit did not check ALP-001"},
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

	// Append the first Audit section to the approved change with a new blank separator.
	const auditAppend = "\n## Audit\nChecked IDs: ALP-001\nNo findings\n"
	before := approved + auditAppend
	write(t, root, pendingCommitChange, before)
	if result, err := verify(root, true); err != nil || result.State.Phase != "audit" {
		t.Errorf("complete audit Verify = phase %q, error %v", result.State.Phase, err)
	}
	if got := readFile(t, root, pendingCommitChange); got != before || head(t, root) != base {
		t.Error("successful audit Verify wrote the submission or committed")
	}
	if now, err := workflow.Current(root, "manual"); err != nil || now.ClaimOID != current.ClaimOID || now.Phase != "audit" || now.PhaseBase != base {
		t.Errorf("successful Verify changed claim or phase: %#v, %v", now, err)
	}

	prefix, provenance, found := strings.Cut(approved, "\n## Provenance\n")
	edited := strings.Replace(provenance, " agent=manual ", " agent=other ", 1)
	if !found || edited == provenance {
		t.Fatal("fixture is missing the approved Provenance entry")
	}
	write(t, root, pendingCommitChange, prefix+"\n## Provenance\n"+edited+auditAppend)
	_, err = verify(root, true)
	assertErrorCode(t, "audit Verify with edited Provenance", err, "P03")
	_, err = done(root)
	assertErrorCode(t, "audit Done with edited Provenance", err, "P03")
	if now, err := workflow.Current(root, "manual"); head(t, root) != base || err != nil || now.ClaimOID != current.ClaimOID || now.Phase != "audit" || now.PhaseBase != base {
		t.Fatalf("rejected Provenance edit changed HEAD, claim or phase: %#v, %v", now, err)
	}

	write(t, root, pendingCommitChange, before)
	if result, err := done(root); err != nil || result.State.Phase != "impl" {
		t.Errorf("complete audit Done = phase %q, error %v", result.State.Phase, err)
	}
}

func Test_FLOW_002_FLOW_136_AuditRejectsForbiddenEditAndDoneMarksFailure(t *testing.T) {
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

	assertRejection := func(operation string, err error) {
		t.Helper()
		code, subject, ok := contracts.ErrorCode(err)
		if !ok || code != "Z01" || subject != forbidden || !strings.Contains(err.Error(), forbidden) {
			t.Errorf("%s error = %v (code %q, subject %q); want Z01 naming %s", operation, err, code, subject, forbidden)
		}
		if got := head(t, root); got != base {
			t.Errorf("%s advanced HEAD: %s -> %s", operation, base, got)
		}
		if got := readFile(t, root, pendingCommitChange); got != submission {
			t.Errorf("%s changed the valid audit submission", operation)
		}
		current, currentErr := workflow.Current(root, "manual")
		if currentErr != nil || current.Phase != "audit" || current.PhaseBase != base || current.ClaimOID != auditState.ClaimOID {
			t.Errorf("%s changed held audit phase or claim: %#v, %v", operation, current, currentErr)
		}
		if _, markerErr := os.Stat(marker); !os.IsNotExist(markerErr) {
			t.Errorf("%s ran the eval hook before zone rejection: %v", operation, markerErr)
		}
	}

	_, quickErr := verify(root, true)
	assertRejection("quick Verify", quickErr)
	verified, verifyErr := verify(root, false)
	assertRejection("full Verify", verifyErr)
	if !slices.ContainsFunc(verified.Diagnostics, func(d project.Diagnostic) bool {
		return d.Code == "Z01" && d.Path == forbidden && !d.Warning
	}) {
		t.Errorf("full Verify diagnostics = %v, want Z01 for %s", verified.Diagnostics, forbidden)
	}

	result, doneErr := done(root)
	assertRejection("Done", doneErr)
	if doneErr == nil || verifyErr == nil {
		t.FailNow()
	}
	if doneErr.Error() != verifyErr.Error() {
		t.Errorf("Done error text = %q, want %q", doneErr.Error(), verifyErr.Error())
	}
	if !reflect.DeepEqual(result, verified) {
		t.Errorf("Done verification result = %#v, want %#v", result, verified)
	}
	if wrapped := fmt.Errorf("finish audit: %w", doneErr); !workflow.IsVerificationFailure(wrapped) {
		t.Errorf("wrapped Done audit verification error was not recognized: %v", wrapped)
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
			assertErrorCode(t, op.name+"("+selector+")", err, "U02")
			if err == nil || !strings.Contains(err.Error(), selector) || !strings.Contains(err.Error(), "manual") {
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
		assertErrorCode(t, "archived "+op.name, err, "U02")
		if err == nil || !strings.Contains(err.Error(), "ghost") || !strings.Contains(err.Error(), "manual") {
			t.Errorf("archived %s(ghost) = %v; want U02 with archive slug manual", op.name, err)
		}
	}
}
