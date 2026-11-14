package evals_test

import (
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
)

func Test_PRJ_078_CheckRemovedSectionContainsOnlyHeadings(t *testing.T) {
	const slug = "remove-obsolete-specs"
	const filename = "0001-" + slug + ".md"
	const changePath = "changes/" + filename
	const first = "### ALP-001\n"
	const second = "### ALP-002\n"
	for _, tc := range []struct {
		name, removed string
		targets       []string
		wantC07       bool
		wantItem      string
	}{
		{"item_body", first + "This removal has an explanation.\n", []string{"ALP-001"}, true, "ALP-001"},
		{"leading_prose", "Remove the obsolete spec.\n\n" + first, []string{"ALP-001"}, true, ""},
		{"prose_only", "Remove the obsolete spec.\n", nil, true, ""},
		{"second_item_body", first + "\n" + second + "This explains the second removal.\n", []string{"ALP-001", "ALP-002"}, true, "ALP-002"},
		{"empty", "", nil, false, ""},
		{"headings_with_whitespace", "\n\t\n" + first + "\n \t\n" + second + "\t\n", []string{"ALP-001", "ALP-002"}, false, ""},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := staleFixture(t)
			var front strings.Builder
			if len(tc.targets) > 0 {
				front.WriteString("bases:\n")
				for _, id := range tc.targets {
					front.WriteString("  " + id + ": " + contracts.SpecHash(staleBaseItem(t, root, id)) + "\n")
				}
			}
			queued(t, root, filename, front.String(), "Retire obsolete behavior.\n", "", "", tc.removed)

			diagnostics := project.Check(root)
			found := false
			for _, d := range diagnostics {
				if d.Warning {
					continue
				}
				if !tc.wantC07 || d.Code != "C07" {
					t.Errorf("unexpected error diagnostic: %#v", d)
					continue
				}
				if (d.Path == changePath || strings.Contains(d.Message, slug)) && strings.Contains(d.Message, tc.wantItem) {
					found = true
				}
			}
			if tc.wantC07 && !found {
				t.Errorf("Check diagnostics = %#v; want non-warning C07 identifying %q and item %q when specified", diagnostics, changePath, tc.wantItem)
			}
		})
	}
}

func Test_PRJ_078_CheckRemovedPreservesHeadingIdentityAndBaseDiagnostics(t *testing.T) {
	const filename = "0001-remove.md"
	for _, tc := range []struct {
		name, removed, front, wantCode, wantMessage string
	}{
		{"malformed_heading", "### ALP-001 [broken\n", "", "C03", "ALP-001"},
		{"duplicate_identity", "### ALP-001\n\n### ALP-001\n", "", "C05", "ALP-001"},
		{"absent_target", "### ALP-999\n", "", "C08", "ALP-999"},
		{"stale_base", "### ALP-001\n", "bases:\n  ALP-001: deadbeef\n", "C08", "stale base for ALP-001"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := staleFixture(t)
			queued(t, root, filename, tc.front, "Retire obsolete behavior.\n", "", "", tc.removed)
			diagnostics := project.Check(root)
			found := false
			for _, d := range diagnostics {
				if d.Warning {
					continue
				}
				if d.Code != tc.wantCode {
					t.Errorf("unexpected error diagnostic: %#v", d)
				}
				if d.Code == tc.wantCode && strings.Contains(d.Message, tc.wantMessage) {
					found = true
				}
			}
			if !found {
				t.Errorf("Check diagnostics = %#v; want %s naming %q", diagnostics, tc.wantCode, tc.wantMessage)
			}
		})
	}
}
