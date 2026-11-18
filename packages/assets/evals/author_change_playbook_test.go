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

func Test_AST_072_AuthorChangeFailureGuidance(t *testing.T) {
	playbook, err := assets.Playbook("author-change")
	if err != nil {
		t.Fatalf("Playbook(author-change): %v", err)
	}

	const failure = `\b(?:fail\w*|errors?)\b`
	const distinction = `\b(?:distinguish\w*|tell\b[^.!?]*\bapart|differentiat\w*|discriminat\w*|(?:observable|distinct) signals?)\b`
	interfaceGuidance := regexp.MustCompile(`(?i)\bobservable\b.*\b(?:behaviou?r|outcomes?)\b.*\b(?:interface|API)\b|\b(?:interface|API)\b.*\bobservable\b.*\b(?:behaviou?r|outcomes?)\b`)
	failureGuidance := regexp.MustCompile(`(?i)` + failure + `.*` + distinction + `|` + distinction + `.*` + failure)
	paragraphs := regexp.MustCompile(`\r?\n[ \t]*\r?\n`).Split(playbook, -1)
	for i, paragraph := range paragraphs {
		paragraphs[i] = strings.Join(strings.Fields(strings.NewReplacer("`", "", "*", "").Replace(paragraph)), " ")
	}
	intro := slices.IndexFunc(paragraphs, interfaceGuidance.MatchString)
	if intro < 0 {
		t.Fatal("author-change playbook lacks guidance on observable behavior at the module interface")
	}
	guidance := slices.IndexFunc(paragraphs[intro+1:], failureGuidance.MatchString)
	if guidance < 0 {
		t.Fatal("author-change playbook must add guidance on distinguishing failures after the observable interface guidance")
	}
	paragraph := paragraphs[intro+1+guidance]
	const differentActions = `\b(?:(?:different|distinct) (?:actions?|responses?|(?:recovery )?steps?|remedies)|(?:respond|act|recover) differently|different(?:ly)? depending on)\b`
	const nextAction = `\b(?:next (?:actions?|steps?|moves?)|what (?:to do|they do|can be done) next|how to (?:respond|recover))\b`
	const iface = `\b(?:interface|API|entry point)\b`
	const produces = `\b(?:produc\w*|originat\w*|return\w*|rais\w*|report\w*|emit\w*|expos\w*|fails?)\b`
	requirePattern(t, paragraph, `(?i)\b(?:specify|state|describe|document|define|spell out|make clear)\b`, "instruction to specify how failures are distinguished")
	requirePattern(t, paragraph, `(?i)\b(?:callers?|users?|consumers?)\b`, "caller or user who distinguishes failures")
	requireGuidanceTogether(t, paragraph, "connect different failure outcomes to different responses", failure, differentActions)
	requireGuidanceTogether(t, paragraph, "use the interface where the failure originates", iface, produces)
	requireGuidanceTogether(t, paragraph, "make the distinction testable through the producing interface",
		`\b(?:test\w*|verif\w*|assert\w*|check\w*|observable)\b`, iface)
	requireGuidanceTogether(t, paragraph, "specify post-failure changes or availability when they affect the next action",
		`\b(?:specify|state|describe|document|define|spell out|capture|record|include|what)\b`,
		failure, `\b(?:chang\w*|remain\w*|still available|preserv\w*|retain\w*|post[- ]failure state|state after)\b`,
		`\b(?:affect\w*|determin\w*|influenc\w*|matter\w*|shap\w*|guid\w*|relevant|rely on|depend\w* on)\b`, nextAction)

	const wording = `\b(?:(?:exact|literal|verbatim) (?:wording|text|messages?|phrasing|language)|word[- ]for[- ]word)\b`
	const promise = `\b(?:interface|API|contract|spec\w*)\b[^.!?]*\b(?:promis\w*|guarantee\w*|specif\w*|requir\w*|contract\w*)\b|\b(?:promis\w*|guarantee\w*|specif\w*|requir\w*|contract\w*)\b[^.!?]*\b(?:by|in|part of)\b[^.!?]*\b(?:interface|API|contract|spec\w*)\b`
	const requireWording = `\b(?:requir\w*|assert\w*|specif\w*|insist\w*|reserv\w*|demand\w*)\b`
	const onlyWhenPromised = `\bonly\s+(?:` + requireWording + `[^.!?]*)?(?:if|when|where|for)\b|\breserve\b[^.!?]*\bfor\b|\b(?:avoid|never|do not|don['’]t)\b[^.!?]*\bunless\b`
	requireGuidanceTogether(t, paragraph, "require exact wording only when the interface promises it",
		wording, promise, requireWording, onlyWhenPromised)
}
