package evals_test

import (
	"testing"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
)

func Test_PRJ_046_CheckEvalFixCannotMixWithFixesOrSpecEdits(t *testing.T) {
	const added = "### new:feature\nThe feature MUST work.\n"
	baseHash := contracts.SpecHash(contracts.Item{ID: "ALP-001", Body: "Alpha MUST work.\n"})
	for _, tc := range []struct {
		name, front, added, modified, removed string
		wantC07            bool
	}{
		{"eval_fix_with_fixes", "eval_fix: [ALP-001]\nfixes: [ALP-001]\n", "", "", "", true},
		{"eval_fix_with_added", "eval_fix: [ALP-001]\n", added, "", "", true},
		{"eval_fix_with_modified", "eval_fix: [ALP-001]\nbases:\n  ALP-001: " + baseHash + "\n", "", "### ALP-001\nAlpha MUST change.\n", "", true},
		{"eval_fix_with_removed", "eval_fix: [ALP-001]\nbases:\n  ALP-001: " + baseHash + "\n", "", "", "### ALP-001\n", true},
		{"eval_fix_only", "eval_fix: [ALP-001]\n", "", "", "", false},
		{"fixes_only", "fixes: [ALP-001]\n", "", "", "", false},
		{"added_only", "", added, "", "", false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := fixture(t)
			put(t, root, "packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\n---\n# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST work.\n")
			queued(t, root, "0001-check.md", tc.front, "Users need the change.\n", tc.added, tc.modified, tc.removed)

			var errors []project.Diagnostic
			for _, diagnostic := range project.Check(root) {
				if !diagnostic.Warning {
					errors = append(errors, diagnostic)
				}
			}
			if tc.wantC07 {
				if len(errors) != 1 || errors[0].Code != "C07" {
					t.Errorf("mixed change error diagnostics = %#v; want only error C07", errors)
				}
			} else if len(errors) != 0 {
				t.Errorf("valid change error diagnostics = %#v; want none", errors)
			}
		})
	}
}
