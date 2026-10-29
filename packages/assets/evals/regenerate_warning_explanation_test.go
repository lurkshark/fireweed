package evals_test

import (
	"testing"

	"github.com/lurkshark/fireweed/packages/assets"
)

func Test_AST_032_RegenerateWarningExplanation(t *testing.T) {
	explanation, err := assets.Explain("P10")
	if err != nil {
		t.Fatalf("Explain(P10): %v", err)
	}

	requirePattern(t, explanation, `(?i)non.blocking[^.\n]*warn|warn[^.\n]*non.blocking`, "non-blocking warning severity")
	requirePattern(t, explanation, `(?i)regenerat`, "regenerate context")
	requirePattern(t, explanation, `(?i)(no|not|without|missing)[^.\n]*(passing|passed|successful)[^.\n]*module drill[^.\n]*(newer|after|since)[^.\n]*(latest|most recent)[^.\n]*applied implementation`, "absence of a passing module drill newer than the latest applied implementation")
	requirePattern(t, explanation, `(?i)(run|perform)[^.\n]*module drill`, "action to run a module drill")
	requirePattern(t, explanation, `(?i)review[^.\n]*report`, "action to review the drill report")
	requirePattern(t, explanation, `(?i)commit[^.\n]*report`, "action to commit the drill report")
}
