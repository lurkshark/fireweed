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
	requirePattern(t, explanation, `(?i)(?:\beval_fix\b[^.\n]*\b(?:mix\w*|combin\w*|alongside|together|coexist\w*)\b|\b(?:mix\w*|combin\w*|alongside|together|coexist\w*)\b[^.\n]*\beval_fix\b)[^.\n]*\bfix(?:es)?\b`, "cause: eval_fix mixed with fixes")
	requirePattern(t, explanation, `(?i)(?:\beval_fix\b[^.\n]*\b(?:mix\w*|combin\w*|alongside|together|coexist\w*)\b|\b(?:mix\w*|combin\w*|alongside|together|coexist\w*)\b[^.\n]*\beval_fix\b)[^.\n]*\bspec(?:ification)?[ -](?:edit|change)s?\b`, "cause: eval_fix mixed with spec edits")
	requirePattern(t, explanation, `(?i)\b(?:separat|split|isolat|keep)\w*[^.\n]*\beval[ _-]?fix\b[^.\n]*\bfrom\b[^.\n]*\bfix(?:es)?\b`, "action: separate eval-fix work from fixes")
	requirePattern(t, explanation, `(?i)\b(?:separat|split|isolat|keep)\w*[^.\n]*\beval[ _-]?fix\b[^.\n]*\bfrom\b[^.\n]*\bspec(?:ification)?\b`, "action: separate eval-fix work from spec changes")
}
