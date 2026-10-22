// Package changeview builds validation views without changing retained documents.
package changeview

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/lurkshark/fireweed/packages/contracts"
)

// Text makes typed values authoritative over the retained author document.
// Rendering only reads metadata values, so a fresh outer map and section slice
// suffice: no caller-owned nested value is ever edited.
func Text(ch contracts.Change) (string, error) {
	d := ch.Document
	d.Metadata = make(map[string]any, len(ch.Document.Metadata)+9)
	for k, v := range ch.Document.Metadata {
		d.Metadata[k] = v
	}
	d.Sections = append([]contracts.Section{}, ch.Document.Sections...)
	d.Metadata["module"] = ch.Module
	for k, v := range map[string]any{"set": ch.Set, "bases": ch.Bases, "fixes": ch.Fixes, "already_satisfied": ch.AlreadySatisfied, "regenerate": ch.Regenerate, "run": ch.Run, "phase": ch.Phase, "blocked": ch.Blocked} {
		if v == nil || reflect.ValueOf(v).IsZero() {
			delete(d.Metadata, k)
		} else {
			d.Metadata[k] = v
		}
	}
	if len(d.Sections) == 0 {
		d.Sections = []contracts.Section{{Name: "Motivation", Body: "Project the typed change."}}
	}
	for _, edit := range []struct {
		name  string
		items []contracts.Spec
	}{{"Added", ch.Added}, {"Modified", ch.Modified}, {"Removed", ch.Removed}} {
		body := Items(edit.items)
		found := false
		for i := range d.Sections {
			if d.Sections[i].Name == edit.name {
				d.Sections[i].Body = body
				found = true
			}
		}
		if !found {
			d.Sections = append(d.Sections, contracts.Section{Name: edit.name, Body: body})
		}
	}
	return contracts.RenderMarkdown(d)
}

func Path(ch contracts.Change) string {
	if ch.Path != "" {
		return ch.Path
	}
	slug := ch.Slug
	if slug == "" {
		slug = "projection"
	}
	return fmt.Sprintf("changes/%04d-%s.md", ch.Number, slug)
}

func Items(specs []contracts.Spec) string {
	var out strings.Builder
	for i, item := range specs {
		if i > 0 {
			out.WriteString("\n\n")
		}
		out.WriteString("### " + item.ID)
		if item.Tag != "" {
			out.WriteString(" " + item.Tag)
		}
		if item.Body != "" {
			out.WriteString("\n\n" + item.Body)
		}
	}
	return out.String()
}

// ArchiveText keeps metadata and normative edits live, with historical prose
// neutralized only for validation. The caller retains the original document.
func ArchiveText(original contracts.Document) (string, error) {
	d := original
	d.Preamble = ""
	d.Sections = append([]contracts.Section{}, original.Sections...)
	for i, s := range d.Sections {
		switch s.Name {
		case "Added", "Modified", "Removed":
		case "Motivation":
			// Preserve the required-section validation for absent/empty motivation.
			if strings.TrimSpace(s.Body) != "" {
				d.Sections[i].Body = "Historical motivation."
			}
		default:
			d.Sections[i].Body = ""
		}
	}
	return contracts.RenderMarkdown(d)
}

func ErrorCode(err error) string {
	if strings.Contains(err.Error(), "undeclared handle ") {
		return "C09"
	}
	return "C05"
}
