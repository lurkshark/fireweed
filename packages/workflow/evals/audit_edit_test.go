package evals_test

import (
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/workflow"
)

const auditNoFindings = "## Audit\nChecked IDs: ALP-001\nNo findings\n"

func auditBeforeProvenance(t *testing.T, approved string) string {
	t.Helper()
	const provenanceHeading = "## Provenance\n"
	prefix, provenance, found := strings.Cut(approved, "\n"+provenanceHeading)
	if !found {
		t.Fatal("fixture is missing the approved Provenance section")
	}
	return prefix + "\n\n" + auditNoFindings + provenanceHeading + provenance
}

func Test_FLOW_147_ApplyAuditEditPreservesApprovedBodyAndSeparators(t *testing.T) {
	const prefix = "## Motivation\nFirst paragraph.\n\nSecond paragraph.  \n"
	const provenance = "## Provenance\n- Earlier entry.\n\n- Later entry.\n"
	const priorAudit = "## Audit\n- A retained finding.\n"
	for _, tc := range []struct{ name, original, submitted, want string }{
		{
			name:      "first Audit before Provenance",
			original:  prefix + provenance,
			submitted: prefix + "\n" + auditNoFindings + "\n" + provenance,
			want:      prefix + auditNoFindings + provenance,
		},
		{
			name:      "replacement with changed adjacent separators",
			original:  prefix + "\n" + priorAudit + "\n\n" + provenance,
			submitted: prefix + auditNoFindings + "\n\n\n" + provenance,
			want:      prefix + "\n" + auditNoFindings + "\n\n" + provenance,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got, err := workflow.ApplyAuditEdit(tc.original, tc.submitted)
			if err != nil || got != tc.want {
				t.Errorf("ApplyAuditEdit = %q, %v; want %q", got, err, tc.want)
			}
		})
	}
}

func Test_FLOW_147_ApplyAuditEditRejectsProtectedEditsAndDuplicateAudit(t *testing.T) {
	const prefix = "## Motivation\nFirst paragraph.\n\nSecond paragraph.\n"
	const provenance = "## Provenance\n- Approved entry.\n"
	const original = prefix + provenance
	const valid = prefix + "\n" + auditNoFindings + "\n" + provenance
	for _, tc := range []struct{ name, submitted string }{
		{"blank line inside Motivation", strings.Replace(valid, "First paragraph.\n\n", "First paragraph.\n", 1)},
		{"edited Provenance", strings.Replace(valid, "Approved entry.", "Edited entry.", 1)},
		{"two Audit sections", valid + "\n## Audit ##\nChecked IDs: ALP-001\nNo findings\n"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got, err := workflow.ApplyAuditEdit(original, tc.submitted); err == nil {
				t.Errorf("ApplyAuditEdit accepted %s: %q", tc.name, got)
			}
		})
	}
}
