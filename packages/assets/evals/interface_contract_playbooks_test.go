package evals_test

import (
	"regexp"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/assets"
)

func Test_AST_074_AuthorInputContracts(t *testing.T) {
	blocks := playbookProse(t, "author-change")
	const specify = `\b(?:specify|state|define|describe|document|explain|spell out|make clear|say)\b`
	requirePlaybookPassage(t, blocks, "specify which inputs govern when several sources can supply them",
		specify,
		`\b(?:multiple|several|different|competing|alternative|two or more|more than one)\b[^.!?]*\b(?:sources?|inputs?|settings?)\b|\b(?:sources?|inputs?)\b[^.!?]*\b(?:compete|overlap)\b`,
		`\b(?:governs?|controls?|wins?|authoritative|precedence|takes? priority|overrides?|takes? effect)\b|\bwhich (?:inputs?|sources?|values?)\b[^.!?]*\b(?:use\w*|read\w*|appl\w*|select\w*|consult\w*)\b`)
	requirePlaybookPassage(t, blocks, "state applicable fallback, merge and preservation behavior that callers rely on",
		specify,
		`\b(?:fallbacks?|falls? back|defaults?|in the absence of|when\b[^.!?]*\b(?:absent|missing|unset|omitted))\b`,
		`\b(?:merg\w*|combin\w*|overlay\w*)\b`,
		`\b(?:preserv\w*|retain\w*|kept|surviv\w*|stays? the same|remain\w* unchanged|keep\w*\s+(?:existing|other))\b`,
		`\b(?:callers?|users?|consumers?)\b[^.!?]*\b(?:depend|rely|need|expect)\w*\b|\b(?:affects?|matters? to|relied on by)\b[^.!?]*\b(?:callers?|users?|consumers?)\b`)
}

func Test_AST_075_EvalInputSelection(t *testing.T) {
	blocks := playbookProse(t, "evals")
	requirePlaybookPassage(t, blocks, "test source selection through the supplying public interface with a conflicting or invalid competing input when it may be ignored",
		`\b(?:test|exercise|check|verify|assert)\w*\b`,
		`\b(?:(?:input|source) selection|precedence|authoritative|takes? effect|wins?)\b|\bwhich (?:input|source|setting|value)\b[^.!?]*\b(?:chosen|selected|used|read|applied)\b`,
		`\b(?:suppl\w*|produc\w*|provid\w*|own\w*)\b[^.!?]*\bpublic (?:interface|api|entry point)\b|\bpublic (?:interface|api|entry point)\b[^.!?]*\b(?:suppl\w*|produc\w*|provid\w*|own\w*|selects? the (?:source|input))\b`,
		`\b(?:competing|other|alternate|alternative|inactive|unused|lower[- ]priority|unselected|ignored)\b[^.!?]*\b(?:input|source|config\w*|setting|value)s?\b`,
		`\b(?:mak(?:e|ing)|giv(?:e|ing)|us(?:e|ing)|suppl(?:y|ying)|set\w*|provid\w*|put\w*|introduc\w*|add\w*|leav\w*|pass\w*)\b[^.!?]*\b(?:conflict\w*|contradict\w*|invalid|malformed|disagree\w*)\b`,
		`\b(?:when|if|where|provided)\b[^.!?]*(?:\b(?:contract|spec\w*|documented)\b[^.!?]*\b(?:ignor\w*|skip\w*|unused|unread|does not (?:read|use))\b|\b(?:ignor\w*|unused|unread)\b[^.!?]*\b(?:by design|under (?:the )?(?:contract|spec\w*))\b)`)
}

func Test_AST_076_EvalStructuredResults(t *testing.T) {
	blocks := playbookProse(t, "evals")
	const assert = `\b(?:assert|check|verify|compare|inspect|test)\w*\b`
	const spec = `\b(?:contract\w*|spec\w*|promis\w*|guarantee\w*)\b`
	const result = `\b(?:meanings?|values?|fields?|semantics?|results?|behaviou?r|content|relationships?)\b`
	const exactRepresentation = `\b(?:exact|literal|byte[- ]for[- ]byte|verbatim|raw)\b[^.!?]*\b(?:representation|format\w*|order\w*|layout|text|seriali[sz]\w*|bytes?)\b`
	const required = `\b(?:specifi(?:ed|es)|contracted|promis(?:ed|es)|requir(?:ed|es)|mandat(?:ed|es)|guarantee(?:d|s))\b`
	const whenRequired = `\b(?:only (?:if|when|where)|unless)\b[^.!?]*` + required
	requirePlaybookPassage(t, blocks, "assert the contracted result of structured outputs",
		assert,
		`\b(?:structured (?:outputs?|results?|data|responses?)|decoded (?:outputs?|results?|data|responses?))\b`,
		result+`[^.!?]*`+spec+`|`+spec+`[^.!?]*`+result)
	requirePlaybookPassage(t, blocks, "assert exact representation only when the spec requires it",
		assert+`|\breserve\w*\b`,
		exactRepresentation+`[^.!?]*`+whenRequired+`|`+
			whenRequired+`[^.!?]*`+exactRepresentation+`|`+
			`\breserve\w*\b[^.!?]*`+exactRepresentation+`[^.!?]*\bfor\b[^.!?]*`+required)
	requirePlaybookPassage(t, blocks, "use a second valid representation to exercise an incidental format or order regression",
		`\b(?:regression|bug|breakage|failure)s?\b`,
		`\b(?:incidental|uncontracted|unspecified|irrelevant|equivalent|cosmetic)\b[^.!?]*\b(?:format\w*|order\w*|layout|representation)\b|\b(?:reorder\w*|(?:different|changed|reversed) (?:field |key )?order)\b`,
		`\b(?:use|try|run|exercise|test|check|include|add|rerun)\w*\b[^.!?]*\b(?:second|another|alternative|two|different)\b[^.!?]*\b(?:valid|equivalent)\b[^.!?]*\b(?:representation|encoding|seriali[sz]\w*|format|forms?|versions?|inputs?)\b|\b(?:rerun|repeat|test|exercise)\w*\b[^.!?]*\b(?:same|identical)\b[^.!?]*\b(?:fields?|keys?)\b[^.!?]*\b(?:different|changed|reversed) order\b`)
}

func playbookProse(t *testing.T, name string) []string {
	t.Helper()
	body, err := assets.Playbook(name)
	if err != nil {
		t.Fatalf("Playbook(%q): %v", name, err)
	}
	blocks := regexp.MustCompile(`\r?\n[ \t]*\r?\n`).Split(body, -1)
	for i, block := range blocks {
		blocks[i] = strings.Join(strings.Fields(strings.NewReplacer("`", "", "*", "", "_", "").Replace(block)), " ")
	}
	return blocks
}

// Match the required ideas within one passage so unrelated advice cannot
// satisfy a clause. Allow rewording, wrapping and multiple sentences.
func requirePlaybookPassage(t *testing.T, blocks []string, description string, terms ...string) {
	t.Helper()
	patterns := make([]*regexp.Regexp, len(terms))
	for i, term := range terms {
		patterns[i] = regexp.MustCompile(`(?i)` + term)
	}
	for _, block := range blocks {
		matched := true
		for _, pattern := range patterns {
			if !pattern.MatchString(block) {
				matched = false
				break
			}
		}
		if matched {
			return
		}
	}
	t.Errorf("playbook must guide readers to %s", description)
}
