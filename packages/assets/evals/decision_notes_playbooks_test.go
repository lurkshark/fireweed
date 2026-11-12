package evals_test

import (
	"regexp"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/assets"
)

func Test_AST_045_DecisionsAndNotesGuidance(t *testing.T) {
	for _, name := range []string{"evals", "impl", "regenerate"} {
		t.Run(name, func(t *testing.T) {
			playbook, err := assets.Playbook(name)
			if err != nil {
				t.Fatalf("Playbook(%q): %v", name, err)
			}
			requireGuidanceSentence(t, playbook, "Decisions contain only previously unspecified behavioral choices",
				`\bDecisions\b`, `\bonly\b`, `\b(?:previously unspecified|not (?:already )?specified|absent from (?:the )?spec)\b`, `\bbehavio(?:u?r|u?ral)\w*\b`)
			requireGuidanceSentence(t, playbook, "one or two sentences per choice",
				`\b(?:one or two|1\s*(?:or|to|-)\s*2)\s+sentences?\b`, `\b(?:per|each|every)\b`, `\b(?:choice|decision)\b`)
			requireGuidanceSentence(t, playbook, "verification, debugging and recovery records belong in Notes",
				`\bNotes\b`, `\bverification\b`, `\bdebugging\b`, `\brecovery\b`)
		})
	}
}

func requireGuidanceSentence(t *testing.T, playbook, description string, terms ...string) {
	t.Helper()
	for _, sentence := range strings.Split(playbook, ".") {
		found := true
		for _, term := range terms {
			if !regexp.MustCompile(`(?i)` + term).MatchString(sentence) {
				found = false
				break
			}
		}
		if found {
			return
		}
	}
	t.Errorf("AST-045: playbook lacks guidance that %s", description)
}

func Test_AST_046_DecisionAndNoteExamples(t *testing.T) {
	for _, name := range []string{"evals", "impl", "regenerate"} {
		t.Run(name, func(t *testing.T) {
			playbook, err := assets.Playbook(name)
			if err != nil {
				t.Fatalf("Playbook(%q): %v", name, err)
			}
			decision := exampleLine(t, playbook, "Decision")
			note := exampleLine(t, playbook, "Note")
			if !regexp.MustCompile(`(?i)\b(?:when|if|default|choose|use|retain|preserve|return|treat|store|allow|reject)\b`).MatchString(decision) {
				t.Errorf("Decision example must describe a lasting behavioral choice, got %q", decision)
			}
			if !regexp.MustCompile(`(?i)\b(?:verif\w*|test\w*|check\w*|debug\w*|recover\w*|restor\w*|reran|failed|passed)\b`).MatchString(note) {
				t.Errorf("Note example must describe verification or recovery, got %q", note)
			}
		})
	}
}

func exampleLine(t *testing.T, playbook, label string) string {
	t.Helper()
	line := regexp.MustCompile(`(?im)^\s*(?:[-*]\s*)?(?:(?:example|e\.g\.)\s+)?` + label + `(?:s)?(?:\s+example)?\s*(?::|—|-)\s*([^\n]+)`).FindStringSubmatch(playbook)
	if len(line) < 2 || strings.TrimSpace(line[1]) == "" {
		t.Errorf("AST-046: playbook needs a labeled %s example", label)
		return ""
	}
	if len([]rune(line[1])) > 180 {
		t.Errorf("AST-046: %s example should be short, got %d characters", label, len([]rune(line[1])))
	}
	return strings.TrimSpace(line[1])
}
