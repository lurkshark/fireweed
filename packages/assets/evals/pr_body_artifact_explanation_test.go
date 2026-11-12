package evals_test

import (
	"testing"

	"github.com/lurkshark/fireweed/packages/assets"
)

func Test_AST_042_B04Explanation(t *testing.T) {
	explanation, err := assets.Explain("B04")
	if err != nil {
		t.Fatalf("Explain(B04): %v", err)
	}

	requirePattern(t, explanation, `(?i)warn`, "warning severity")
	requirePattern(t, explanation, `(?i)binary`, "binary files")
	requirePattern(t, explanation, `(?i)(>|larger than|greater than|more than|over|exceed\w*)[^.\n]*256\s*(K(?:i)?B|kilobytes?)`, "files strictly larger than 256 KB")
	requirePattern(t, explanation, `(?i)\bor\b`, "binary or oversized files")
	requirePattern(t, explanation, `(?i)(new|newly|added|addition)[^.\n]*disposable|disposable[^.\n]*(new|newly|added|addition)`, "new additions to disposable zones")
	requirePattern(t, explanation, `(?i)pr[- ]body[^.\n]*(report|show|flag|warn)|(report|show|flag|warn)[^.\n]*pr[- ]body`, "PR-body reporting")
	requirePattern(t, explanation, `(?i)opaque`, "opaque additions")
	requirePattern(t, explanation, `(?i)review[^.\n]*(addition|file|artifact|opaque)`, "review opaque additions")
	requirePattern(t, explanation, `(?i)remov|replac`, "remove or replace artifacts")
	requirePattern(t, explanation, `(?i)unnecess|unneeded|not needed|redundan`, "unnecessary artifacts")
	requirePattern(t, explanation, `(?i)generat`, "generated artifacts")
	requirePattern(t, explanation, `(?i)vendor`, "vendored artifacts")
}
