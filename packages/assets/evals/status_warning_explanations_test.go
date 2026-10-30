package evals_test

import (
	"testing"

	"github.com/lurkshark/fireweed/packages/assets"
)

func Test_AST_038_B01Explanation(t *testing.T) {
	explanation, err := assets.Explain("B01")
	if err != nil {
		t.Fatalf("Explain(B01): %v", err)
	}

	requirePattern(t, explanation, `(?i)budget[^.\n]*warn|warn[^.\n]*budget`, "budget warning")
	requirePattern(t, explanation, `(?i)(sum|total|aggregat)`, "summed duration")
	requirePattern(t, explanation, `(?i)JUnit`, "JUnit timing source")
	requirePattern(t, explanation, `(?i)(testcase|test case)[^.\n]*(second|duration|time)|(second|duration|time)[^.\n]*(testcase|test case)`, "testcase seconds")
	requirePattern(t, explanation, `(?i)evaluat[^.\n]*modul|modul[^.\n]*evaluat`, "evaluated modules")
	requirePattern(t, explanation, `(?i)(exceed|over|above)[^\n]*verify_time|verify_time[^\n]*(exceed|over|above)`, "verify_time threshold exceeded")
	requirePattern(t, explanation, `(?i)default[^.\n]*(15m|15 min)|(?:15m|15 min)[^.\n]*default`, "15m default")
	requirePattern(t, explanation, `(?i)review[^.\n]*(expensive|costly|redundan)|review[^.\n]*eval[^.\n]*(expensive|costly|redundan)`, "review expensive or redundant eval work")
	requirePattern(t, explanation, `(?i)(preserv|retain|keep)[^.\n]*(required|necessary)[^.\n]*assertion`, "preserve required assertions")
}

func Test_AST_039_B02Explanation(t *testing.T) {
	explanation, err := assets.Explain("B02")
	if err != nil {
		t.Fatalf("Explain(B02): %v", err)
	}

	requirePattern(t, explanation, `(?i)budget[^.\n]*warn|warn[^.\n]*budget`, "budget warning")
	requirePattern(t, explanation, `(?i)(spec|specification)[^.\n]*(testcase|test case|case)|(testcase|test case|case)[^.\n]*(spec|specification)`, "testcase count per spec")
	requirePattern(t, explanation, `(?i)(exceed|over|above)[^\n]*cases_per_spec|cases_per_spec[^\n]*(exceed|over|above)`, "cases_per_spec threshold exceeded")
	requirePattern(t, explanation, `(?i)default[^.\n]*15\b|15\b[^.\n]*default`, "default of 15")
	requirePattern(t, explanation, `(?i)consolidat[^.\n]*redundan[^.\n]*case|consolidat[^.\n]*case[^.\n]*redundan`, "consolidate redundant cases")
	requirePattern(t, explanation, `(?i)(preserv|retain|keep)[^.\n]*(required|necessary)[^.\n]*behavioral[^.\n]*assertion`, "preserve required behavioral assertions")
}

func Test_AST_040_B03Explanation(t *testing.T) {
	explanation, err := assets.Explain("B03")
	if err != nil {
		t.Fatalf("Explain(B03): %v", err)
	}

	requirePattern(t, explanation, `(?i)budget[^.\n]*warn|warn[^.\n]*budget`, "budget warning")
	requirePattern(t, explanation, `(?i)(ratio|divid|per)[^.\n]*eval[^.\n]*disposable|eval[^.\n]*(divid|/|per|to)[^.\n]*disposable`, "eval lines divided by disposable lines")
	requirePattern(t, explanation, `(?i)eval[^.\n]*lines?`, "eval line count")
	requirePattern(t, explanation, `(?i)disposable[^.\n]*lines?`, "disposable line count")
	requirePattern(t, explanation, `(?i)(exceed|over|above)[^\n]*eval_ratio|eval_ratio[^\n]*(exceed|over|above)`, "eval_ratio threshold exceeded")
	requirePattern(t, explanation, `(?i)default[^.\n]*4\b|4\b[^.\n]*default`, "default ratio of 4")
	requirePattern(t, explanation, `(?i)(positive|non.?zero|greater than zero)[^.\n]*eval|eval[^.\n]*(positive|non.?zero|greater than zero)`, "positive eval lines")
	requirePattern(t, explanation, `(?i)(zero|no)[^.\n]*disposable|disposable[^.\n]*(zero|no)`, "zero disposable lines")
	requirePattern(t, explanation, `(?i)simplif[^.\n]*(redundan|duplicat)[^.\n]*(eval|fixture)|simplif[^.\n]*(eval|fixture)[^.\n]*(redundan|duplicat)`, "simplify redundant eval or fixture code")
	requirePattern(t, explanation, `(?i)(preserv|retain|keep)[^.\n]*(required|necessary)[^.\n]*assertion`, "preserve required assertions")
}

func Test_AST_041_P08Explanation(t *testing.T) {
	explanation, err := assets.Explain("P08")
	if err != nil {
		t.Fatalf("Explain(P08): %v", err)
	}

	requirePattern(t, explanation, `(?i)warn`, "warning")
	requirePattern(t, explanation, `(?i)(at least\s+(three|3)|(three|3)\s+or\s+more|>=\s*3|3\+)[^.\n]*(consecutive|in a row)|(consecutive|in a row)[^.\n]*(at least\s+(three|3)|(three|3)\s+or\s+more|>=\s*3|3\+)`, "at least three consecutive changes")
	requirePattern(t, explanation, `(?i)applied[^.\n]*chang|chang[^.\n]*applied`, "applied changes")
	requirePattern(t, explanation, `(?i)same[^.\n]*modul|modul[^.\n]*same`, "same module")
	requirePattern(t, explanation, `(?i)review[^.\n]*scope`, "review sequence scope")
	requirePattern(t, explanation, `(?i)project owner`, "project owner")
	requirePattern(t, explanation, `(?i)(before|prior to)[^.\n]*continu|continu[^.\n]*(after|once|following)[^.\n]*review`, "review before continuing")
}
