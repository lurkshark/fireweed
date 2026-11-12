package evals_test

import (
	"testing"

	"github.com/lurkshark/fireweed/packages/assets"
)

func Test_AST_060_C17Explanation(t *testing.T) {
	explanation, err := assets.Explain("C17")
	if err != nil {
		t.Fatalf("Explain(C17): %v", err)
	}

	requirePattern(t, explanation, `(?i)\bwarning\b`, "warning severity")
	requirePattern(t, explanation, `(?i)\bspec(?:ification)? item\b`, "per-item limit")
	requirePattern(t, explanation, `(?i)\b(?:exceeds?|over|above|more than)\b[^\n]*\bspec_limits\.max_clauses\b[^\n]*\bnormative keywords\b|\bnormative keywords\b[^\n]*\b(?:exceeds?|over|above|more than)\b[^\n]*\bspec_limits\.max_clauses\b`, "warning above the normative keyword limit")
	requirePattern(t, explanation, `(?i)\b(?:or|also|independently)\b[^\n]*(?:\bspec_limits\.max_words\b[^\n]*\bwords\b|\bwords\b[^\n]*\bspec_limits\.max_words\b)`, "warning independently above the word limit")
	requirePattern(t, explanation, `(?i)\bspec_limits\.max_clauses\b[^.;\n]*\bdefault\b[^.;\n]*\b3\b`, "default of 3 normative keywords")
	requirePattern(t, explanation, `(?i)\bspec_limits\.max_words\b[^.;\n]*\bdefault\b[^.;\n]*\b80\b`, "default of 80 words")
	requirePattern(t, explanation, `(?i)\b(?:split|narrow)\b[^.\n]*\b(?:item|spec)\b`, "advice to split or narrow the oversized item")
}

func Test_AST_060_C18Explanation(t *testing.T) {
	explanation, err := assets.Explain("C18")
	if err != nil {
		t.Fatalf("Explain(C18): %v", err)
	}

	requirePattern(t, explanation, `(?i)\bwarning\b`, "warning severity")
	requirePattern(t, explanation, `(?i)\b(?:same|repeated|duplicate)\b[^.\n]*\bsentence\b|\bsentence\b[^.\n]*\b(?:repeats?|repeated|duplicated)\b`, "repeated sentence")
	requirePattern(t, explanation, `(?i)\b(?:(?:three|3) or more|at least (?:three|3))\s+(?:distinct\s+)?spec(?:ification)? items\b`, "repetition in three or more spec items")
	requirePattern(t, explanation, `(?i)\bmove\b[^\n]*\bsentence\b[^\n]*\bSYSTEM\.md\b[^\n]*\bconvention\b|\bmove\b[^\n]*\bsentence\b[^\n]*\bconvention\b[^\n]*\bSYSTEM\.md\b`, "advice to move the repeated sentence into SYSTEM.md as a convention")
}

func Test_AST_060_C19Explanation(t *testing.T) {
	explanation, err := assets.Explain("C19")
	if err != nil {
		t.Fatalf("Explain(C19): %v", err)
	}

	requirePattern(t, explanation, `(?i)\bwarning\b`, "warning severity")
	requirePattern(t, explanation, `(?i)\bmodule\b[^.\n]*\bDecisions\b[^.\n]*\bsection\b|\bDecisions\b[^.\n]*\bsection\b[^.\n]*\bmodule\b`, "module Decisions section")
	requirePattern(t, explanation, `(?i)\b(?:exceeds?|over|above|more than)\b[^\n]*\bspec_limits\.max_decision_words\b|\bspec_limits\.max_decision_words\b[^\n]*\b(?:exceeded|over|above)\b`, "warning above the Decisions word limit")
	requirePattern(t, explanation, `(?i)\bspec_limits\.max_decision_words\b[^.;\n]*\bdefault\b[^.;\n]*\b300\b`, "default of 300 words")
	requirePattern(t, explanation, `(?i)\b(?:drop|remove)\b[^.\n]*\bstale\b[^.\n]*\bdecisions\b[^.\n]*\bRemoved decisions\b`, "advice to drop stale decisions through Removed decisions")
	requirePattern(t, explanation, `(?i)\b(?:turn|convert|promote)\b[^.\n]*\blasting\b[^.\n]*\bspecs?\b`, "advice to turn lasting decisions into specs")
}
