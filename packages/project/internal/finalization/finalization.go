package finalization

import (
	"fmt"
	"strings"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project/internal/changeview"
	"github.com/lurkshark/fireweed/packages/project/internal/projection"
)

func Finalize(module contracts.Module, change contracts.Change) (contracts.Module, error) {
	fail := func(err error) (contracts.Module, error) { return contracts.Module{}, err }
	text, err := changeview.Text(change)
	if err != nil {
		return fail(err)
	}
	checked, err := contracts.ParseChange(text, change.Path, module.IDPrefix)
	if err != nil {
		return fail(fmt.Errorf("invalid final change: %w", err))
	}
	if change.Slug == "" || change.Slug != checked.Slug {
		return fail(fmt.Errorf("change slug %q must match its valid filename slug %q", change.Slug, checked.Slug))
	}
	result, err := projection.Apply(module, change)
	if err != nil {
		return fail(err)
	}
	// A preview may contain local IDs. Parsing as a current module both rejects
	// those placeholders and establishes one normalized owned representation.
	result, err = canonical(result)
	if err != nil {
		return fail(fmt.Errorf("invalid final module: %w", err))
	}
	remaining := result.Decisions
	removalRequested := false
	for _, line := range strings.Split(section(checked.Document, "Removed decisions"), "\n") {
		quote := strings.TrimSpace(line)
		quote = strings.TrimSpace(strings.TrimPrefix(quote, ">"))
		if quote != "" {
			removalRequested = true
			remaining = strings.ReplaceAll(remaining, quote, "")
		}
	}
	if removalRequested {
		setDecisions(&result, remaining)
		// Removing text can leave outer blank lines. Normalize those before
		// choosing the separator for a newly appended tagged body.
		result, err = canonical(result)
		if err != nil {
			return fail(err)
		}
	}
	addition := section(checked.Document, "Decisions")
	if strings.TrimSpace(addition) != "" {
		body := "[" + change.Slug + "]\n" + addition
		if strings.TrimSpace(result.Decisions) != "" {
			body = result.Decisions + "\n\n" + body
		}
		setDecisions(&result, body)
		result, err = canonical(result)
		if err != nil {
			return fail(err)
		}
	}
	return result, nil
}

func canonical(m contracts.Module) (contracts.Module, error) {
	text, err := contracts.RenderMarkdown(m.Document)
	if err != nil {
		return contracts.Module{}, err
	}
	return contracts.ParseModule(text, m.Dir, contracts.Zones{}, contracts.Hooks{})
}
func section(d contracts.Document, name string) string {
	for _, s := range d.Sections {
		if s.Name == name {
			return s.Body
		}
	}
	return ""
}
func setDecisions(m *contracts.Module, body string) {
	m.Decisions = body
	for i := range m.Document.Sections {
		if m.Document.Sections[i].Name == "Decisions" {
			m.Document.Sections[i].Body = body
			return
		}
	}
	m.Document.Sections = append(m.Document.Sections, contracts.Section{Name: "Decisions", Body: body})
}
