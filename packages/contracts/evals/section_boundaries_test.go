package evals_test

import (
	"testing"

	"github.com/lurkshark/fireweed/packages/contracts"
)

func Test_CON_002_CON_009_FencedSectionHeadings(t *testing.T) {
	for _, tc := range []struct {
		name, fence string
	}{
		{"backticks", "```"},
		{"tildes", "~~~"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			const prefix = "# Document\n\n## Notes\n"
			const suffix = "## Next\nKeep this section.\n"
			const replacement = "Updated notes.\n"
			body := "Before the example.\n\n" + tc.fence + "markdown\n" +
				"## Example\n# Example title\nExample text.\n" + tc.fence + "\n\nAfter the example.\n\n"
			source := prefix + body + suffix

			if got, ok := contracts.Section(source, "Notes"); !ok || got != body {
				t.Errorf("Section(Notes) = %q, %v; want %q, true", got, ok, body)
			}
			if _, ok := contracts.Section(source, "Example"); ok {
				t.Error("Section(Example) reported a heading inside fenced code as present")
			}
			if got, want := contracts.ReplaceSection(source, "Notes", replacement), prefix+replacement+suffix; got != want {
				t.Errorf("ReplaceSection(Notes) = %q; want %q", got, want)
			}
			if got, want := contracts.ReplaceSection(source, "Example", replacement), source+"## Example\n"+replacement; got != want {
				t.Errorf("ReplaceSection(Example) = %q; want %q", got, want)
			}
		})
	}
}

func Test_CON_002_CON_009_LevelOneSectionBoundary(t *testing.T) {
	const prefix = "# Document\n\n## Notes\n"
	const body = "Original notes.\n\n"
	const suffix = "# Another title\n\nKeep this introduction.\n\n## Next\nKeep this section.\n"
	const source = prefix + body + suffix
	const replacement = "Updated notes.\n"

	if got, ok := contracts.Section(source, "Notes"); !ok || got != body {
		t.Errorf("Section(Notes) = %q, %v; want %q, true", got, ok, body)
	}
	if got, want := contracts.ReplaceSection(source, "Notes", replacement), prefix+replacement+suffix; got != want {
		t.Errorf("ReplaceSection(Notes) = %q; want %q", got, want)
	}
}

func Test_CON_002_IndentedSectionHeadings(t *testing.T) {
	for _, tc := range []struct{ name, indent string }{
		{"one_space", " "},
		{"two_spaces", "  "},
		{"three_spaces", "   "},
	} {
		t.Run(tc.name, func(t *testing.T) {
			const specs = "Original specs.\n\n    # Example title\n    ## Example section\n\nMore specs.\n\n"
			const decisions = "Project decisions.\n\n"
			source := "# Module\n\n    ## Specs\n    Example specs.\n\n" +
				tc.indent + "## Specs\n" + specs +
				tc.indent + "## Decisions\n" + decisions +
				tc.indent + "# Appendix\n\nAppendix text.\n"

			if got, ok := contracts.Section(source, "Specs"); !ok || got != specs {
				t.Errorf("Section(Specs) = %q, %v; want %q, true", got, ok, specs)
			}
			if got, ok := contracts.Section(source, "Decisions"); !ok || got != decisions {
				t.Errorf("Section(Decisions) = %q, %v; want %q, true", got, ok, decisions)
			}
		})
	}
}
