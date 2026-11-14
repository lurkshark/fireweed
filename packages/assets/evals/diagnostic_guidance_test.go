package evals_test

import (
	"regexp"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/assets"
)

func Test_AST_066_C07ChangeStructure(t *testing.T) {
	explanation := diagnosticGuidance(t, "C07")
	const field = `\b(?:fixes|eval_fix|already_satisfied)\b`
	const list = `\blists?\b`
	const invalid = `\b(?:invalid|malformed|incorrect|must|require\w*)\b`
	const invalidTarget = invalid + `|\b(?:unknown|foreign|missing|nonexistent|valid|existing)\b`
	const identities = `\b(?:IDs?|identit(?:y|ies)|targets?|specs?|items?)\b`
	// Tie the rule to each field so unrelated list guidance cannot satisfy it.
	for _, name := range []string{"fixes", "eval_fix", "already_satisfied"} {
		requireGuidanceTogether(t, explanation, "list requirements for "+name, `\b`+name+`\b`, list, invalid)
		requireGuidanceTogether(t, explanation, "target validity for "+name, `\b`+name+`\b`, identities, invalidTarget)
	}
	requirePattern(t, explanation,
		`(?i)\bRemoved\b[^.!?]*\b(?:only|limited to)\b[^.!?]*\bheadings?\b|\bRemoved\b[^.!?]*\bheadings?\b[^.!?]*\bonly\b|`+
			`\b(?:only|invalid|reject\w*)\b[^.!?]*\b(?:headings?|non[- ]heading|prose|body text|commentary)\b[^.!?]*\bRemoved\b|`+
			`\bRemoved\b[^.!?]*\b(?:reject\w*|disallow\w*|cannot contain)\b[^.!?]*\b(?:non[- ]heading|prose|body text|commentary)\b|`+
			`\b(?:non[- ]heading|prose|body text|commentary)\b[^.!?]*(?:\bRemoved\b[^.!?]*\b(?:invalid|forbidden|reject\w*)\b|\b(?:invalid|forbidden|reject\w*)\b[^.!?]*\bRemoved\b)`,
		"Removed permits heading-only content")
	requireGuidanceTogether(t, explanation, "correct the listed identities",
		`\b(?:correct|fix|update|replace|remove|use)\b`, identities, field+`|\b(?:listed|entries|spec IDs?)\b`)
	requirePattern(t, explanation,
		`(?i)\b(?:move|put|place|relocate)\b[^.!?]*\b(?:commentary|prose|text|explanations?)\b[^.!?]*\b(?:Motivation|Notes)\b`,
		"move removal commentary to Motivation or Notes")
}

func Test_AST_067_P05BaseTipWarning(t *testing.T) {
	explanation := diagnosticGuidance(t, "P05")
	const baseTip = `(?:\bbase(?: branch| ref)?(?:'s)? (?:tip|head|commit)\b|\b(?:tip|head|commit) of [^.!?]*\bbase\b)`
	const lacksTip = `\bHEAD\b[^.!?]*\b(?:lacks?|(?:does not|doesn't) (?:contain|include|descend from)|is missing|is behind)\b[^.!?]*` + baseTip +
		`|` + baseTip + `[^.!?]*\b(?:absent from|missing from|(?:not|isn't) (?:in|contained in|an ancestor of))\s+HEAD\b`
	requireGuidanceTogether(t, explanation, "warn when HEAD lacks the locally resolved base tip",
		`\bwarn\w*\b`, `\blocal(?:ly)?\b`, lacksTip)
	requireGuidanceTogether(t, explanation, "reject a stale spec hash",
		`\b(?:stale|mismatch\w*|out[- ]of[- ]date)\b`, `\bspec(?:ification)?\b[^.!?]*\bhash\w*\b`,
		`\b(?:reject\w*|block\w*|error|fail\w*)\b`)
	requirePattern(t, explanation,
		`(?i)\bqueue\b[^.!?]*\b(?:may|might|could) be stale\b|\bpotentially stale\s+queue\b`,
		"staleness warning applies to the queue")
	requirePattern(t, explanation, `(?i)\b(?:review|check|inspect)\b[^.!?]*\bbranch\b`, "review the branch")
	const update = `\b(?:update|sync\w*|bring)\b[^.!?]*\bbranch\b`
	const continueQueue = `\bbefore\b[^.!?]*\b(?:continu\w*|proceed\w*|resum\w*|run\w*)\b`
	requirePattern(t, explanation, `(?i)`+update+`[^.!?]*`+continueQueue+`|`+continueQueue+`[^.!?]*`+update,
		"update the branch before continuing")
}

func Test_AST_068_P02DuplicateRunRecovery(t *testing.T) {
	explanation := diagnosticGuidance(t, "P02")
	requireGuidanceTogether(t, explanation, "duplicate runs of the change on different machines",
		`\b(?:duplicate runs?|two runs|same change|run twice)\b`,
		`\b(?:different|separate|two) (?:machines|computers|hosts)\b`)
	requireGuidanceTogether(t, explanation, "merge conflict in the archived change's run identity",
		`\bconflict\w*\b`, `\barchiv\w*\b`, `\brun[- ](?:IDs?|identity|identities|identifiers?)\b`)

	const abort = `\b(?:abort|cancel|stop)\b[^.!?]*\bmerge\b|\bgit merge --abort\b`
	const commits = `\bfireweed commits\s+<change>`
	const revert = `\brevert\b[^.!?]*\b(?:commits?|them)\b`
	const retry = `\b(?:re[- ]?merge|merge\b[^.!?]*\bagain|(?:retry|rerun|repeat)\b[^.!?]*\bmerge)\b`
	requirePattern(t, explanation, `(?i)(?:`+abort+`).*?`+commits+`.*?`+revert+`.*?`+retry,
		"abort the merge, look up commits, revert them, and merge again in that order")
	requireGuidanceTogether(t, explanation, "use fireweed commits <change> to identify the incoming duplicate's commits",
		commits, `\b(?:list|find|identify|show|locate)\w*\b`, `\b(?:duplicate|incoming)\b`)
	requireGuidanceTogether(t, explanation, "revert the duplicate commits on the incoming branch",
		revert, `\bincoming branch\b`)
	requirePattern(t, explanation,
		`(?i)\b(?:choos|pick|tak|accept|keep|select)\w*\b[^.!?]*\b(?:one|either)\b[^.!?]*\b(?:archiv\w*|side|version)\b[^.!?]*`+
			`\b(?:can|could|may|might)\b[^.!?]*\b(?:retain|keep|leave|preserve)\w*\b[^.!?]*\bboth implementations\b`,
		"choosing one archive side can retain both implementations")
}

func diagnosticGuidance(t *testing.T, code string) string {
	t.Helper()
	explanation, err := assets.Explain(code)
	if err != nil {
		t.Fatalf("Explain(%s): %v", code, err)
	}
	return strings.Join(strings.Fields(strings.ReplaceAll(explanation, "`", "")), " ")
}

func requireGuidanceTogether(t *testing.T, explanation, description string, terms ...string) {
	t.Helper()
	patterns := make([]*regexp.Regexp, len(terms))
	for i, term := range terms {
		patterns[i] = regexp.MustCompile(`(?i)` + term)
	}
	for _, sentence := range strings.FieldsFunc(explanation, func(r rune) bool {
		return r == '.' || r == '!' || r == '?'
	}) {
		matches := true
		for _, pattern := range patterns {
			if !pattern.MatchString(sentence) {
				matches = false
				break
			}
		}
		if matches {
			return
		}
	}
	t.Errorf("missing guidance to %s in %q", description, explanation)
}
