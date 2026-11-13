package queue

import (
	"fmt"
	"path"
	"regexp"
	"strings"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
	"github.com/lurkshark/fireweed/packages/project/model"
)

type staticCheck struct {
	diagnostics []model.Diagnostic
	maxClauses  int
	maxWords    int
}

func Check(root string) []model.Diagnostic {
	check := staticCheck{}
	snapshot, err := fixture.Load(root)
	if err != nil {
		check.report(err)
		return check.diagnostics
	}
	p, err := newProjection(snapshot)
	if err != nil {
		check.report(err)
		return check.diagnostics
	}
	limits := snapshot.Config["spec_limits"].(map[string]any)
	check.maxClauses = limits["max_clauses"].(int)
	check.maxWords = limits["max_words"].(int)
	for i, module := range snapshot.Modules {
		file := path.Join(module.Dir, "SPEC.md")
		for _, item := range p.items[i] {
			if !ownsID(module, item.ID) {
				check.error("C05", file, "foreign spec ID "+item.ID)
			}
			check.validateItem(file, item)
		}
		decisions, _ := contracts.Section(module.Document.Body, "Decisions")
		if count := len(strings.Fields(decisions)); count > limits["max_decision_words"].(int) {
			check.warn("C19", file, fmt.Sprintf("module %s has %d words in Decisions (limit %d)", module.Name, count, limits["max_decision_words"]))
		}
	}
	check.outsideSpecs(snapshot)

	changes, err := workingQueue(snapshot)
	if err != nil {
		check.report(err)
		return check.diagnostics
	}
	for _, change := range changes {
		check.validateChange(p, change)
		if err := p.apply(change); err != nil {
			// Base validation already reported each missing preceding item.
			if code, _, _ := contracts.ErrorCode(err); code != "C08" {
				check.report(err)
			}
			return check.diagnostics
		}
	}
	check.repeatedSentences(p.items)
	return check.diagnostics
}

func (check *staticCheck) report(err error) {
	code, subject, _ := contracts.ErrorCode(err)
	check.error(code, subject, err.Error())
}

func (check *staticCheck) error(code, path, message string) {
	check.diagnostics = append(check.diagnostics, model.Diagnostic{Code: code, Path: path, Message: message})
}

func (check *staticCheck) warn(code, path, message string) {
	check.diagnostics = append(check.diagnostics, model.Diagnostic{Code: code, Path: path, Message: message, Warning: true})
}

var normative = regexp.MustCompile(`\b(?:MUST|SHOULD)\b`)
var handleReference = regexp.MustCompile(`new:[a-z0-9-]+`)

func ownsID(module model.Module, id string) bool {
	prefix, _, _ := strings.Cut(id, "-")
	return prefix == module.IDPrefix
}

func (check *staticCheck) validateTag(file string, item contracts.Item) {
	kind, reason, _ := strings.Cut(item.Tag, ":")
	if item.Tag != "" && item.Tag != "review" &&
		(kind != "quarantine" || strings.TrimSpace(reason) == "") {
		check.error("C06", file, "invalid review/quarantine tag on "+item.ID)
	}
}

func (check *staticCheck) validateItem(file string, item contracts.Item) {
	check.validateTag(file, item)
	clauses := len(normative.FindAllStringIndex(item.Body, -1))
	if clauses == 0 {
		check.error("C04", file, "non-normative spec body for "+item.ID)
	}
	words := len(strings.Fields(item.Body))
	if clauses > check.maxClauses || words > check.maxWords {
		check.warn("C17", file, fmt.Sprintf("%s has %d normative keywords and %d words (limits %d and %d)",
			item.ID, clauses, words, check.maxClauses, check.maxWords))
	}
}

func (check *staticCheck) validateChange(p *projection, change model.Change) {
	motivation, _ := contracts.Section(change.Document.Body, "Motivation")
	if strings.TrimSpace(motivation) == "" {
		check.error("C07", change.Path, "missing Motivation")
	}
	fields := change.Document.Frontmatter
	_, evalFix := fields["eval_fix"]
	_, fixes := fields["fixes"]
	if evalFix && (fixes || len(change.Added)+len(change.Modified)+len(change.Removed) > 0) {
		check.error("C07", change.Path, "eval_fix cannot be combined with fixes or spec edits")
	}
	index := p.modules[change.Module]
	module := p.snapshot.Modules[index]
	seen := make(map[string]bool)
	handles := make(map[string]bool)
	for _, item := range change.Added {
		if strings.HasPrefix(item.ID, "new:") {
			handles[item.ID] = true
		}
	}
	for _, edits := range []struct {
		items   []contracts.Item
		added   bool
		removed bool
	}{
		{items: change.Added, added: true}, {items: change.Modified}, {items: change.Removed, removed: true},
	} {
		for _, item := range edits.items {
			if seen[item.ID] {
				check.error("C05", change.Path, "duplicate edit identity "+item.ID)
			}
			seen[item.ID] = true
			if !ownsID(module, item.ID) && !(edits.added && handles[item.ID]) {
				check.error("C05", change.Path, "foreign module ID "+item.ID)
			}
			if edits.removed {
				check.validateTag(change.Path, item)
			} else {
				check.validateItem(change.Path, item)
			}
		}
	}
	text, err := contracts.RenderDocument(change.Document)
	if err != nil {
		check.error("C03", change.Path, err.Error())
		return
	}
	for _, handle := range handleReference.FindAllString(text, -1) {
		if !handles[handle] {
			check.error("C09", change.Path, "reference to undeclared handle "+handle)
			handles[handle] = true
		}
	}
	check.validateBases(change, assignedItems(p.items[index]))
}

func (check *staticCheck) validateBases(change model.Change, previous map[string]contracts.Item) {
	mapping, err := bases(change)
	if err != nil {
		check.report(err)
	}
	for _, edit := range baseEdits(change) {
		item, exists := previous[edit.ID]
		if !exists {
			check.report(missingSpec(change, edit.ID))
			continue
		}
		if err != nil {
			continue
		}
		value := mapping[edit.ID]
		switch {
		case value == nil || value == "":
			check.error("C08", change.Path, "missing base for "+edit.ID)
		case recordedHash(value) != contracts.SpecHash(item):
			check.error("C08", change.Path, "stale base for "+edit.ID)
		}
	}
}
