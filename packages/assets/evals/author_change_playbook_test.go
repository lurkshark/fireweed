package evals_test

import (
	"regexp"
	"slices"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/assets"
)

func Test_AST_064_AuthorChangeProtocol(t *testing.T) {
	playbook, err := assets.Playbook("author-change")
	if err != nil {
		t.Fatalf("Playbook(author-change): %v", err)
	}
	paragraphs := regexp.MustCompile(`\r?\n[ \t]*\r?\n`).Split(playbook, -1)
	for i, paragraph := range paragraphs {
		paragraphs[i] = strings.Join(strings.Fields(strings.ReplaceAll(paragraph, "`", "")), " ")
	}

	const stamp = `\bfireweed stamp\s+(?:<change>|[a-z0-9][a-z0-9_./-]*)`
	const stampInstruction = `\b(?:run|invoke|execute|use|with)\b[^.!?]*` + stamp +
		`|` + stamp + `[^.!?]*\b(?:must|should|is required)\b`
	for _, guidance := range []struct {
		description string
		terms       []string
	}{
		{"use new:<slug> handles for Added specs", []string{
			`\b(?:Added|add(?:ing)? (?:new )?specs?|new spec(?:ification)?s?)\b`,
			`\b(?:use|assign|give|name|label|declare|write|start|take|require)\w*\b[^.!?]*\bnew:(?:<slug>|[a-z0-9][a-z0-9-]*\b)`,
		}},
		{"keep handles unique within their change", []string{
			`\bhandles?\b|\bnew:`,
			`\b(?:unique|distinct|different|no duplicates?)\b`,
			`\b(?:change[- ]local|per change|(?:within|inside|in|to) (?:the|a|each|one|that|this|its) (?:same )?change)\b`,
		}},
		{"combine changes when one needs another's handle", []string{
			`\b(?:combine|merge|fold|consolidate)\b[^.!?]*\b(?:changes|them|together|into (?:one|a single) change)\b`,
			`\b(?:need\w*|depend\w* on|refer\w* to|reference\w*|use\w*|rely on|requires?)\b[^.!?]*\b(?:handles?\b|new:)`,
			`\b(?:another|different|other|separate) (?:pending |queued )?(?:change|draft)\b|\b(?:in|from|by) another\b|\b(?:across changes|cross[- ]change)\b`,
		}},
		{"run fireweed stamp for Modified specs", []string{
			`\b(?:Modified|modif(?:y|ies|ying) (?:a |any |existing )?specs?)\b`,
			stampInstruction,
		}},
		{"run fireweed stamp for Removed specs", []string{
			`\b(?:Removed|remov(?:e|es|ing) (?:a |any |existing )?specs?)\b`,
			stampInstruction,
		}},
		{"resolve stale conflicts before replacing a base", []string{
			`\b(?:stale|out.of.date)\b`,
			`\b(?:resolve|reconcile|address|fix|rework|rewrite|revise|rebase)\b[^.!?]*\b(?:conflicts?|edits?|changes?|specs?|intent)\b`,
			`\b(?:never|do not|don['’]t|avoid|without|rather than|instead of)\b[^.!?]*\b(?:overwrit|replac|refresh|updat)\w*\b[^.!?]*\b(?:stale|bases?|hashes)\b` +
				`|\b(?:before|only then|then)\b[^.!?]*\b(?:re-?stamp\w*|stamp\w* again|fireweed stamp)\b` +
				`|\bbefore\b[^.!?]*\b(?:overwrit|replac|refresh|updat)\w*\b[^.!?]*\b(?:bases?|hashes)\b`,
		}},
	} {
		patterns := make([]*regexp.Regexp, len(guidance.terms))
		for i, term := range guidance.terms {
			patterns[i] = regexp.MustCompile(`(?i)` + term)
		}
		found := slices.ContainsFunc(paragraphs, func(paragraph string) bool {
			for _, pattern := range patterns {
				if !pattern.MatchString(paragraph) {
					return false
				}
			}
			return true
		})
		if !found {
			t.Errorf("author-change playbook must instruct authors to %s", guidance.description)
		}
	}
}
