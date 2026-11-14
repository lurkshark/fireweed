package evals_test

import (
	"testing"

	"github.com/lurkshark/fireweed/packages/project"
)

func Test_PRJ_077_CheckRequiresExactlyOneSpecsSectionPerModule(t *testing.T) {
	const path = "packages/alpha/SPEC.md"
	const header = "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\n---\n# Alpha\n\n"
	const populated = "## Specs\n\n### ALP-001\nAlpha MUST work.\n"
	for _, tc := range []struct {
		name, body string
		wantC03    bool
	}{
		{"absent", "## Decisions\n\nPreserve existing records.\n", true},
		{"duplicate_with_repeated_id", populated + "\n## Specs\n\n### ALP-001\nAlpha MUST also retry.\n", true},
		{"duplicate_with_atx_closing_hashes", populated + "\n# Appendix\n\n   ##\tSpecs ###\n\n### ALP-002\nAlpha MUST retry.\n", true},
		{"single_empty", "## Specs\n\n## Decisions\n", false},
		{"single_populated", populated, false},
		{"single_atx_heading_with_crlf", "   ##\tSpecs ###\r\n\r\n### ALP-001\r\nAlpha MUST work.\r\n", false},
		{"single_with_backtick_example", populated + "\n```markdown\n## Specs\n```\n", false},
		{"single_with_tilde_example", "~~~markdown\n## Specs\n~~~\n\n" + populated, false},
		{"fenced_example_only", "````markdown\n```\n## Specs\n````\n", true},
		{"nonmatching_headings_only", "# Specs\n\n### Specs\n\n##Specs\n\n    ## Specs\n", true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := fixture(t)
			put(t, root, path, header+tc.body)

			diagnostics := project.Check(root)
			for _, diagnostic := range diagnostics {
				if diagnostic.Warning {
					continue
				}
				if !tc.wantC03 {
					t.Errorf("valid module received error diagnostic: %#v", diagnostic)
				} else if diagnostic.Code == "C03" && diagnostic.Path == path {
					return
				}
			}
			if tc.wantC03 {
				t.Errorf("Check diagnostics = %#v; want a non-warning C03 at %q", diagnostics, path)
			}
		})
	}

	t.Run("pending_change_without_specs", func(t *testing.T) {
		root := fixture(t)
		put(t, root, path, header+populated)
		put(t, root, "changes/0001-motivation-only.md", "---\nmodule: alpha\n---\n## Motivation\n\nRecord the reason for the change.\n")
		for _, diagnostic := range project.Check(root) {
			if !diagnostic.Warning {
				t.Errorf("valid pending change received error diagnostic: %#v", diagnostic)
			}
		}
	})
}
