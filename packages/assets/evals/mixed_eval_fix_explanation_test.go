package evals_test

import (
	"testing"

	"github.com/lurkshark/fireweed/packages/assets"
)

func Test_AST_033_MixedEvalFixExplanation(t *testing.T) {
	explanation, err := assets.Explain("C07")
	if err != nil {
		t.Fatalf("Explain(C07): %v", err)
	}

	requirePattern(t, explanation, `(?i)\beval_fix\b`, "eval_fix change kind")
	requirePattern(t, explanation, `(?i)(mix|combin|alongside|together|same change|coexist)[^.\n]*\bfix(?:es)?\b`, "cause: eval_fix combined with fixes")
	requirePattern(t, explanation, `(?i)(mix|combin|alongside|together|same change|coexist)[^.\n]*\bspec(?:ification)?[ -](?:edit|change)s?\b`, "cause: eval_fix combined with spec edits")
	requirePattern(t, explanation, `(?i)(separat|split|isolat|keep)[^.\n]*\beval[ _-]?fix\b|\beval[ _-]?fix\b[^.\n]*(separat|split|isolat)`, "action: separate eval-fix work")
}
