package evals_test

import (
	"regexp"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/assets"
)

func Test_AST_043_EvalsStubHandoff(t *testing.T) {
	playbook, err := assets.Playbook("evals")
	if err != nil {
		t.Fatalf("Playbook(evals): %v", err)
	}
	requirePattern(t, playbook, `(?i)temporar(y|ily)|stub`, "temporary stub behavior")
	requirePattern(t, playbook, `(?i)\b(replace|rewrite|remove|discard|delete)\w*\b[^.\n]*(stub|temporar)|\b(stub|temporar)\w*\b[^.\n]*\b(replace|rewrite|remove|discard|delete)\w*\b`, "temporary behavior can be replaced during implementation")
	requirePattern(t, playbook, `(?i)\b(implement|implementation)\b[^.\n]*\b(frozen|retained|interface)\b|\b(frozen|retained|interface)\b[^.\n]*\b(implement|implementation)\b`, "implementation does not require editing frozen interface files")
	requirePattern(t, playbook, `(?i)\b(identify|name|list|record)\b[^.\n]*\b(implementation.writable|writable|disposable)\b[^.\n]*\bfiles?\b`, "identify implementation-writable files")
	requirePattern(t, playbook, `(?i)\b(missing|unimplemented|remaining)\b[^.\n]*\b(behavior|functionality)\b`, "identify files for missing behavior")
	requirePattern(t, playbook, `(?i)\b(before|prior to)\b[^.\n]*\b(complet|finish|end)\w*\b[^.\n]*\bevals?\b|\b(complet|finish|end)\w*\b[^.\n]*\bevals?\b[^.\n]*\b(before|prior to)\b`, "identify writable files before completing evals")
}

func Test_AST_044_GoRecipeStubHandoff(t *testing.T) {
	r := goRecipe(t)
	for name, instructions := range map[string]string{
		"Instructions":  r.Instructions,
		"bundled go.md": r.Files[".fireweed/recipes/go.md"],
	} {
		t.Run(name, func(t *testing.T) {
			requirePattern(t, instructions, `(?i)api\.go`, "retained api.go example")
			requirePattern(t, instructions, `(?i)internal/[^\s`+"`"+`]+\.go`, "new internal Go stub file")
			// The example needs executable-looking Go declarations and a facade call,
			// rather than a prose-only description of the handoff.
			requirePattern(t, instructions, `(?s)\bfunc\s+[A-Za-z_]\w*\s*\([^)]*\)[^{\n]*\{[^}]*\b[A-Za-z_]\w*\.[A-Za-z_]\w*\s*\(`, "api.go function delegates to an internal function")
			if len(regexp.MustCompile(`\bfunc\s+[A-Za-z_]\w*\s*\(`).FindAllString(instructions, -1)) < 2 {
				t.Error("example must show both the facade and temporary stub functions")
			}
			requirePattern(t, instructions, `(?i)\b(replace|rewrite|remove|discard|delete)\w*\b[^.\n]*\b(stub|temporar)\w*\b|\b(stub|temporar)\w*\b[^.\n]*\b(replace|rewrite|remove|discard|delete)\w*\b`, "implementation replaces the disposable stub")
			requirePattern(t, instructions, `(?i)\b(api\.go|facade)\b[^.\n]*\b(unchanged|untouched|same)\b|\b(unchanged|untouched|same)\b[^.\n]*\b(api\.go|facade)\b`, "retained facade remains unchanged")
			if !strings.Contains(strings.ToLower(instructions), "disposable") {
				t.Error("example must identify the stub as disposable")
			}
		})
	}
}
