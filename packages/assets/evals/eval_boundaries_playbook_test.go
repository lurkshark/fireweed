package evals_test

import (
	"regexp"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/assets"
)

func Test_AST_071_PublicBehaviorGuidance(t *testing.T) {
	playbook, err := assets.Playbook("evals")
	if err != nil {
		t.Fatalf("Playbook(evals): %v", err)
	}
	// Normalize wrapping within prose blocks and keep unrelated guidance apart.
	blocks := regexp.MustCompile(`\r?\n[ \t]*\r?\n`).Split(playbook, -1)
	for i, block := range blocks {
		blocks[i] = strings.Join(strings.Fields(strings.NewReplacer("`", "", "*", "", "_", "").Replace(block)), " ")
	}
	prose := strings.Join(blocks, "\n")

	const consumers = `\b(?:consumers?|callers?|dependents?|dependent modules?)\b`
	const publicBehavior = `\bpublic (?:behaviou?r|contracts?|interface|API)\b`
	const reliedOn = consumers + `\s+(?:depend on|rely on|use|need)\b|\b(?:used|needed|relied on) by\s+` + consumers
	requirePattern(t, prose, `(?i)\b(?:exercise|test|cover)\b.*`+publicBehavior, "instruction to exercise the module's public behavior")
	requirePattern(t, prose, `(?i)`+publicBehavior+`.*(?:`+reliedOn+`)|(?:`+reliedOn+`).*`+publicBehavior,
		"public behavior that consumers depend on")

	const dependencySetup = `\b(?:build|create|establish|set up|prepare|seed)\b.*\b(?:dependenc(?:y|ies)\s+(?:state|fixtures?)|(?:state|fixtures?)\b.*\bdependenc(?:y|ies)|(?:the |your )?dependencies)\b`
	const publicOperations = `\bpublic\s+(?:operations?|APIs?|interfaces?|methods?)\b`
	requirePattern(t, prose,
		`(?i)`+dependencySetup+`.*\b(?:using|through|via|with)\b.*`+publicOperations+`|\buse\b.*`+publicOperations+`.*\bto\s+`+dependencySetup,
		"instruction to build ordinary dependency state through public operations")

	const codesAndSubjects = `\b(?:codes?\s+and\s+subjects?|subjects?\s+and\s+codes?)\b`
	const diagnosticIdentity = `\bstructured diagnostics?\s+` + codesAndSubjects + `|` + codesAndSubjects + `\s+(?:(?:from|of|in|on)\s+(?:the\s+)?structured diagnostics?\b|(?:diagnostic )?fields\b)|\bdiagnostics?\s+` + codesAndSubjects + `\s+fields\b`
	const contracted = `\b(?:when|if|where)\b.*\b(?:contract\w*|specif\w*|promis\w*)\b`
	requirePattern(t, prose, `(?i)\b(?:assert|check|compare|verify)\b.*(?:`+diagnosticIdentity+`)`,
		"assertions on structured diagnostic codes and subjects")
	requirePattern(t, prose, `(?i)(?:`+diagnosticIdentity+`).*(?:`+contracted+`)|(?:`+contracted+`).*(?:`+diagnosticIdentity+`)`,
		"diagnostic code and subject assertions when contracted")

	const typedValues = `\b(?:(?:typed|native|original|returned|raw) values?|(?:value|concrete) types?|types? of (?:the )?values?)\b`
	requirePattern(t, prose, `(?i)\b(?:inspect|check|assert|compare|verify)\b.*`+typedValues+`.*\bbefore\s+(?:encoding|serializ\w*|marshal\w*|a\s+JSON\s+round.trip)\b`,
		"instruction to inspect typed values before encoding")
	const callerImpact = `\btypes?\b.*\b(?:matters? to|affects?|is observable to)\s+(?:the\s+)?` + consumers +
		`|` + consumers + `\s+(?:distinguish|depend on|rely on|can observe)\b.*\btypes?\b`
	requirePattern(t, prose, `(?i)\b(?:when|if|where)\b.*(?:`+callerImpact+`)`,
		"typed assertions when type affects callers")
}
