package queue

import (
	"fmt"
	"regexp"
	"slices"
	"strings"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
	"github.com/lurkshark/fireweed/packages/project/model"
)

var handleReference = regexp.MustCompile(`new:[a-z0-9-]+`)
var normative = regexp.MustCompile(`\b(?:MUST|SHOULD)\b`)

func Check(root string) []model.Diagnostic {
	var diagnostics []model.Diagnostic
	addCode := func(code, path, message string) {
		diagnostics = append(diagnostics, model.Diagnostic{Code: code, Path: path, Message: message})
	}
	snapshot, err := fixture.Load(root)
	if err != nil {
		addCode(errorCode(err, "C01"), fixture.LoadPath(err), err.Error())
		return diagnostics
	}
	graphErr := fixture.ValidateGraph(snapshot.Modules)
	if graphErr != nil {
		addCode("C02", "fireweed.yaml", graphErr.Error())
	}
	limits := snapshot.Config["spec_limits"].(map[string]any)
	maxClauses := limits["max_clauses"].(int)
	maxWords := limits["max_words"].(int)
	maxDecisionWords := limits["max_decision_words"].(int)
	diagnostics = append(diagnostics, specsOutsideModules(snapshot)...)
	p, err := newProjection(snapshot)
	if err != nil {
		addCode(errorCode(err, "C03"), "SPEC.md", err.Error())
		return diagnostics
	}
	for _, module := range snapshot.Modules {
		path := module.Dir + "/SPEC.md"
		decisions, _ := contracts.Section(module.Document.Body, "Decisions")
		if words := len(strings.Fields(decisions)); words > maxDecisionWords {
			diagnostics = append(diagnostics, model.Diagnostic{
				Code: "C19", Path: path, Warning: true,
				Message: fmt.Sprintf("module %s has %d Decisions words (limit %d)", module.Name, words, maxDecisionWords),
			})
		}
		for _, entry := range p.items[module.Name] {
			item := entry.item
			if !strings.HasPrefix(item.ID, module.IDPrefix+"-") {
				addCode("C05", path, "foreign spec ID "+item.ID)
			}
			diagnostics = append(diagnostics, checkItem(path, item, true, maxClauses, maxWords)...)
		}
	}
	changes, err := readQueue(root, snapshot)
	if err != nil {
		addCode(errorCode(err, "C03"), "changes", err.Error())
		return diagnostics
	}
	projectable := graphErr == nil
	for _, change := range changes {
		path := change.Path
		motivation, exists := contracts.Section(change.Document.Body, "Motivation")
		if !exists || strings.TrimSpace(motivation) == "" {
			addCode("C07", path, "missing Motivation")
		}
		if _, hasEvalFix := change.Document.Frontmatter["eval_fix"]; hasEvalFix {
			_, hasFixes := change.Document.Frontmatter["fixes"]
			if hasFixes || len(change.Added) > 0 || len(change.Modified) > 0 || len(change.Removed) > 0 {
				addCode("C07", path, "eval_fix cannot be combined with fixes or spec edits")
			}
		}
		declared, seen := map[string]bool{}, map[string]bool{}
		for _, item := range change.Added {
			if strings.HasPrefix(item.ID, "new:") {
				declared[item.ID] = true
			}
		}
		for _, group := range []struct {
			items     []contracts.Item
			normative bool
		}{
			{change.Added, true}, {change.Modified, true}, {change.Removed, false},
		} {
			for _, item := range group.items {
				if seen[item.ID] {
					addCode("C05", path, "duplicate edit identity "+item.ID)
				}
				seen[item.ID] = true
				if !strings.HasPrefix(item.ID, "new:") && !strings.HasPrefix(item.ID, modulePrefix(snapshot, change.Module)+"-") {
					addCode("C05", path, "foreign module ID "+item.ID)
				}
				diagnostics = append(diagnostics, checkItem(path, item, group.normative, maxClauses, maxWords)...)
			}
		}
		for _, ref := range handleReference.FindAllString(change.Document.Body, -1) {
			if !declared[ref] {
				addCode("C09", path, "reference to undeclared handle "+ref)
			}
		}
		var bases map[string]any
		if raw, exists := change.Document.Frontmatter["bases"]; exists {
			bases, _ = raw.(map[string]any)
		}
		for _, item := range append(append([]contracts.Item{}, change.Modified...), change.Removed...) {
			base, ok := p.find(change.Module, item.ID)
			if !ok {
				addCode("C08", path, "ID "+item.ID+" absent from preceding projection")
				continue
			}
			got, ok := bases[item.ID].(string)
			if !ok || got == "" {
				addCode("C08", path, "missing base for "+item.ID)
			} else if got != contracts.SpecHash(base) {
				addCode("C08", path, "stale base for "+item.ID)
			}
		}
		if err := p.apply(change); err != nil {
			addCode(errorCode(err, "C02"), path, err.Error())
			projectable = false
		}
	}
	if projectable {
		diagnostics = append(diagnostics, repeatedSentenceWarnings(p)...)
	}
	return diagnostics
}

func specsOutsideModules(snapshot model.Snapshot) []model.Diagnostic {
	tracked, err := readGit(snapshot.Root, "ls-files", "--cached", "-z", "--", ":(glob)**/SPEC.md")
	if err != nil {
		// Check also validates project documents outside a Git repository.
		return nil
	}
	var warnings []model.Diagnostic
	seen := make(map[string]bool)
	for _, path := range strings.Split(string(tracked), "\x00") {
		if path == "" || strings.HasPrefix(path, ".fireweed/") || seen[path] {
			continue
		}
		if slices.ContainsFunc(snapshot.Modules, func(module model.Module) bool {
			return strings.HasPrefix(path, module.Dir+"/")
		}) {
			continue
		}
		seen[path] = true
		warnings = append(warnings, model.Diagnostic{
			Code:    "C20",
			Path:    path,
			Message: path + " is outside the configured module globs",
			Warning: true,
		})
	}
	return warnings
}

func modulePrefix(snapshot model.Snapshot, name string) string {
	for _, module := range snapshot.Modules {
		if module.Name == name {
			return module.IDPrefix
		}
	}
	return ""
}

func checkItem(path string, item contracts.Item, requireNormative bool, maxClauses, maxWords int) []model.Diagnostic {
	var diagnostics []model.Diagnostic
	if tag := item.Tag; tag != "" && tag != "review" && !(strings.HasPrefix(tag, "quarantine: ") && strings.TrimSpace(strings.TrimPrefix(tag, "quarantine: ")) != "") {
		diagnostics = append(diagnostics, model.Diagnostic{Code: "C06", Path: path, Message: "invalid review/quarantine tag " + tag})
	}
	if !requireNormative {
		return diagnostics
	}
	// Matching MUST once also counts MUST NOT as a single clause.
	clauses := len(normative.FindAllStringIndex(item.Body, -1))
	if clauses == 0 {
		diagnostics = append(diagnostics, model.Diagnostic{Code: "C04", Path: path, Message: "non-normative spec body for " + item.ID})
	}
	words := len(strings.Fields(item.Body))
	if clauses > maxClauses || words > maxWords {
		diagnostics = append(diagnostics, model.Diagnostic{
			Code: "C17", Path: path, Warning: true,
			Message: fmt.Sprintf("spec item %s has %d normative clauses (limit %d) and %d words (limit %d)", item.ID, clauses, maxClauses, words, maxWords),
		})
	}
	return diagnostics
}
