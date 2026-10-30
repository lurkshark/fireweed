package evals_test

import (
	"testing"

	"github.com/lurkshark/fireweed/packages/assets"
)

func Test_AST_034_H01Explanation(t *testing.T) {
	explanation, err := assets.Explain("H01")
	if err != nil {
		t.Fatalf("Explain(H01): %v", err)
	}

	requirePattern(t, explanation, `(?i)worker[^.\n]*phase[^.\n]*commit[^.\n]*(chang|edit|touch)[^.\n]*protected[^.\n]*path`, "worker phase commit changing a protected path")
	requirePattern(t, explanation, `(?i)worker[^.\n]*phase[^.\n]*commit[^.\n]*(chang|edit|touch)[^.\n]*project[^.\n]*playbook[^.\n]*override`, "worker phase commit changing a project playbook override")
	requirePattern(t, explanation, `(?i)(restore|revert)[^.\n]*(edit|change|path|override)`, "restore the worker edit")
	requirePattern(t, explanation, `(?i)(separat|move|keep|place)[^.\n]*(human|author)[^.\n]*setup[^.\n]*author commit`, "separate human-owned setup work into an author commit")
	requirePattern(t, explanation, `(?i)(separat|move|keep|place)[^.\n]*(human|author)[^.\n]*override[^.\n]*author commit`, "separate human-owned override work into an author commit")
	requirePattern(t, explanation, `(?i)author commit[^.\n]*(outside|apart from|separate from)[^.\n]*worker phase`, "author commit outside worker phases")
}

func Test_AST_035_H02Explanation(t *testing.T) {
	explanation, err := assets.Explain("H02")
	if err != nil {
		t.Fatalf("Explain(H02): %v", err)
	}

	requirePattern(t, explanation, `(?i)impl.to.applied[^.\n]*commit[^.\n]*(chang|edit|touch)[^.\n]*eval`, "impl-to-applied commit changing evals")
	requirePattern(t, explanation, `(?i)impl.to.applied[^.\n]*commit[^.\n]*(chang|edit|touch)[^.\n]*interface file`, "impl-to-applied commit changing interface files")
	requirePattern(t, explanation, `(?i)(make|move|put|do)[^.\n]*(eval[^.\n]*edit|edit[^.\n]*eval)[^.\n]*evals step[^.\n]*before[^.\n]*implementation`, "make eval edits in the evals step before implementation")
	requirePattern(t, explanation, `(?i)(make|move|put|do)[^.\n]*(interface[^.\n]*edit|edit[^.\n]*interface)[^.\n]*evals step[^.\n]*before[^.\n]*implementation`, "make interface edits in the evals step before implementation")
	requirePattern(t, explanation, `(?i)(rework|repair|rewrite|correct)[^.\n]*phase history[^.\n]*before[^.\n]*(rerun|repeat)[^.\n]*history check`, "rework phase history before rerunning history checks")
}

func Test_AST_036_H03Explanation(t *testing.T) {
	explanation, err := assets.Explain("H03")
	if err != nil {
		t.Fatalf("Explain(H03): %v", err)
	}

	for _, operation := range []string{"modif", "delet", "renam"} {
		requirePattern(t, explanation, `(?i)`+operation+`[^.\n]*existing[^.\n]*append.only[^.\n]*path`, operation+" of an existing append-only path")
	}
	requirePattern(t, explanation, `(?i)(restore|reinstate)[^.\n]*original[^.\n]*path`, "restore the original path")
	requirePattern(t, explanation, `(?i)(restore|reinstate)[^.\n]*original[^.\n]*content`, "restore the original content")
	requirePattern(t, explanation, `(?i)(?:^|[.!?]\s+)(?:add|create) (?:a )?new file for new (?:data|content)(?:[.!?]|$)`, "add a new file for new data")
}

func Test_AST_037_H04Explanation(t *testing.T) {
	explanation, err := assets.Explain("H04")
	if err != nil {
		t.Fatalf("Explain(H04): %v", err)
	}

	requirePattern(t, explanation, `(?i)modif[^.\n]*archiv[^.\n]*change file|archiv[^.\n]*change file[^.\n]*modif`, "modification of an archived change file")
	requirePattern(t, explanation, `(?i)(?:^|[.!?]\s+)(?:deleting|deletion of) (?:an? |the )?archived (?:change )?file (?:is|remains) (?:allowed|permitted) for (?:a )?revert(?:[.!?]|$)`, "archived file deletion allowed for a revert")
	requirePattern(t, explanation, `(?i)(restore|revert)[^.\n]*archiv`, "restore the archive")
	requirePattern(t, explanation, `(?i)(author|create|open|add)[^.\n]*follow.up[^.\n]*pending change`, "author follow-up work as a pending change")
}
