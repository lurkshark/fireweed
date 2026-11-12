package queue

import (
	"fmt"
	"maps"
	"regexp"
	"slices"
	"strings"
	"unicode"

	"github.com/lurkshark/fireweed/packages/project/model"
)

// Require whitespace or the end of the body after a terminator so that
// dotted identifiers and decimals stay within their sentence.
var sentenceEnd = regexp.MustCompile(`[.!?]+["'”’)\]]*(?:[[:space:]\p{Z}\x{85}]+|$)`)

func repeatedSentenceWarnings(p *projection) []model.Diagnostic {
	type sentenceUses struct {
		text  string
		items []string
	}
	uses := make(map[string]sentenceUses)
	for _, module := range p.snapshot.Modules {
		for _, entry := range p.items[module.Name] {
			name := entry.item.ID
			if entry.owner != "" {
				name += " (" + entry.owner + ")"
			}
			seen := make(map[string]bool)
			for _, sentence := range sentenceEnd.Split(entry.item.Body, -1) {
				key := strings.Map(func(r rune) rune {
					if unicode.IsPunct(r) || unicode.IsSpace(r) {
						return -1
					}
					return unicode.ToLower(r)
				}, sentence)
				if key == "" || seen[key] {
					continue
				}
				seen[key] = true
				use := uses[key]
				if len(use.items) == 0 {
					use.text = strings.Join(strings.Fields(sentence), " ")
				}
				use.items = append(use.items, name)
				uses[key] = use
			}
		}
	}
	var warnings []model.Diagnostic
	for _, key := range slices.Sorted(maps.Keys(uses)) {
		use := uses[key]
		if len(use.items) >= 3 {
			warnings = append(warnings, model.Diagnostic{
				Code: "C18", Warning: true,
				Message: fmt.Sprintf("sentence %q is repeated in %s", use.text, strings.Join(use.items, ", ")),
			})
		}
	}
	return warnings
}
