package evals_test

import (
	"regexp"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/assets"
)

func Test_AST_029_EvalFixEvalsGuidance(t *testing.T) {
	playbook, guidance := evalFixPlaybook(t, "evals")
	requirePattern(t, guidance, `(?i)\b(correct|repair|fix|edit|change)\b[^.]*\b(evals?|tests?)\b[^.]*\bonly\b|\bonly\b[^.]*\b(evals?|tests?)\b`, "eval_fix edits only evals")
	requirePattern(t, guidance, `(?i)\b(retain|preserve|keep|maintain)\b[^.]*\bcoverage\b[^.]*\b(listed|assigned|affected)\b[^.]*\bIDs?\b|\bcoverage\b[^.]*\b(listed|assigned|affected)\b[^.]*\bIDs?\b[^.]*\b(retain|preserve|keep|maintain)\b`, "coverage of listed IDs retained")
	requirePattern(t, guidance, `(?i)\ball\b[^.]*\bmodule\b[^.]*\b(evals?|tests?)\b[^.]*\bpass\b|\b(evals?|tests?)\b[^.]*\bmodule\b[^.]*\ball\b[^.]*\bpass\b`, "all module evals pass")
	requirePattern(t, guidance, `(?i)\bfail[ -]?first\b[^.]*\b(not required|does not apply|unnecessary|waived|exempt|instead)\b|\b(no|without|rather than|instead of)\b[^.]*\bfail[ -]?first\b`, "eval_fix does not require fail-first")
	requirePattern(t, playbook, `(?i)\bfail[ -]?first\b`, "ordinary fail-first guidance retained")
}

func Test_AST_030_EvalFixAuditGuidance(t *testing.T) {
	_, guidance := evalFixPlaybook(t, "audit")
	requirePattern(t, guidance, `(?i)\b(check|verify|confirm|ensure)\b[^.]*\b(corrected|repaired|fixed|edited)\b[^.]*\b(evals?|tests?)\b[^.]*\b(assert|cover|test)\b[^.]*\b(listed|assigned|affected)\b[^.]*\b(specs?|IDs?|requirements?)\b`, "corrected tests still assert listed specs")
}

func Test_AST_031_EvalFixImplGuidance(t *testing.T) {
	playbook, guidance := evalFixPlaybook(t, "impl")
	requirePattern(t, guidance, `(?i)\b(edit|change|modify|update)\b[^.]*\bonly\b[^.]*\b(current|active)\b[^.]*\bchange\b[^.]*\bfile\b|\b(edit|change|modify|update)\b[^.]*\b(current|active)\b[^.]*\bchange\b[^.]*\bfile\b[^.]*\bonly\b`, "only the current change file may be edited")
	requirePattern(t, guidance, `(?i)\b(leave|keep|preserve)\b[^.]*\bimplementation\b[^.]*\b(unchanged|untouched)|\b(no|without)\b[^.]*\bimplementation\b[^.]*\b(edits?|changes?)\b`, "implementation remains unchanged")
	requirePattern(t, playbook, `(?i)\bimplement\b[^.]*\b(disposable|module)\b`, "ordinary implementation guidance retained")
}

func evalFixPlaybook(t *testing.T, name string) (string, string) {
	t.Helper()
	playbook, err := assets.Playbook(name)
	if err != nil {
		t.Fatalf("Playbook(%q): %v", name, err)
	}
	for _, paragraph := range regexp.MustCompile(`\n\s*\n`).Split(playbook, -1) {
		if strings.Contains(strings.ToLower(paragraph), "eval_fix") {
			requirePattern(t, paragraph, `(?i)\b(if|when|for|in)\b[^.]*\beval_fix\b`, "guidance conditional on eval_fix")
			requirePattern(t, playbook, `(?i)\bdo not commit\b`, "worker no-commit guidance retained")
			return playbook, paragraph
		}
	}
	t.Fatalf("Playbook(%q) lacks eval_fix guidance", name)
	return "", ""
}
