package evals_test

import "testing"

func Test_AST_047_GoRecipeEvalNameBoundary(t *testing.T) {
	recipe := goRecipe(t)
	for _, guide := range []struct {
		name, text string
	}{
		{"recipe instructions", recipe.Instructions},
		{"module SPEC template", recipe.Files[".fireweed/templates/module/SPEC.md"]},
	} {
		t.Run(guide.name, func(t *testing.T) {
			requirePattern(t, guide.text,
				`(?i)\bletter\b[^.\n]*\b(immediately|directly)\b[^.\n]*\b(before|preceding)\b[^.\n]*\b(spec\s+)?prefix\b[^.\n]*\b(prevent\w*|block\w*|stop\w*)\b[^.\n]*\b(match\w*|mapping)\b`,
				"letter immediately before a spec prefix prevents ID matching")
			requirePattern(t, guide.text,
				`(?i)Test_GRT_001_Name[^.\n]*\b(matches|matching|maps)\b[^.\n]*GRT-001`,
				"Test_GRT_001_Name matches GRT-001")
			requirePattern(t, guide.text,
				`(?i)TestGRT001[^.\n]*(does not match|doesn't match|fails to match|cannot match)[^.\n]*GRT-001`,
				"TestGRT001 does not match GRT-001")
		})
	}
}
