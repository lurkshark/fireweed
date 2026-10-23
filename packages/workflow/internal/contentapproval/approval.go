package contentapproval

import (
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/lurkshark/fireweed/packages/contracts"
)

func problem(format string, args ...any) error {
	return fmt.Errorf("C05: %s. Fix the current change before comparing its approved content.", fmt.Sprintf(format, args...))
}

func Matches(module contracts.Module, baseText, currentText, path string) (bool, error) {
	current, err := contracts.ParseChange(currentText, path, module.IDPrefix)
	if err != nil {
		return false, problem("cannot parse current change: %v", err)
	}
	if current.Module != module.Name {
		return false, problem("current target %q differs from supplied module %q", current.Module, module.Name)
	}
	base, err := contracts.ParseChange(baseText, path, module.IDPrefix)
	if err != nil || base.Module != module.Name {
		return false, nil
	}
	if len(base.Added) != len(current.Added) {
		return false, nil
	}

	// Align the base's local bindings with this particular started (or still
	// unstarted) document, without allocating anything from module state.
	bindings := make(map[string]string)
	for i, item := range base.Added {
		if strings.HasPrefix(item.ID, "new:") {
			bindings[item.ID] = current.Added[i].ID
		}
	}
	if len(bindings) != 0 {
		names := make([]string, 0, len(bindings))
		for name := range bindings {
			names = append(names, name)
		}
		sort.Slice(names, func(i, j int) bool {
			if len(names[i]) != len(names[j]) {
				return len(names[i]) > len(names[j])
			}
			return names[i] < names[j]
		})
		pairs := make([]string, 0, 2*len(names))
		for _, name := range names {
			pairs = append(pairs, name, bindings[name])
		}
		canonical, err := contracts.RenderMarkdown(base.Document)
		if err != nil {
			return false, nil
		}
		// NewReplacer visits the original text once: an emitted handle is not fed
		// through another binding. Longest names take precedence over their prefix.
		canonical = strings.NewReplacer(pairs...).Replace(canonical)
		base, err = contracts.ParseChange(canonical, path, module.IDPrefix)
		if err != nil {
			return false, nil
		}
	}
	if !reflect.DeepEqual(authored(base), authored(current)) || base.Document.Preamble != current.Document.Preamble {
		return false, nil
	}
	return reflect.DeepEqual(sections(base), sections(current)), nil
}

type authorMetadata struct {
	Module                  string
	Set                     map[string]any
	Bases                   map[string]string
	Fixes, AlreadySatisfied []string
	Regenerate              bool
}

func authored(change contracts.Change) authorMetadata {
	meta := authorMetadata{Module: change.Module, Set: change.Set, Bases: change.Bases, Fixes: change.Fixes, AlreadySatisfied: change.AlreadySatisfied, Regenerate: change.Regenerate}
	// Only empty top-level collections are equivalent to omission. Empty members
	// within Set remain part of the author's requested override.
	if len(meta.Set) == 0 {
		meta.Set = nil
	}
	if len(meta.Bases) == 0 {
		meta.Bases = nil
	}
	if len(meta.Fixes) == 0 {
		meta.Fixes = nil
	}
	if len(meta.AlreadySatisfied) == 0 {
		meta.AlreadySatisfied = nil
	}
	return meta
}
func sections(change contracts.Change) []contracts.Section {
	var retained []contracts.Section
	for _, section := range change.Document.Sections {
		switch section.Name {
		case "Decisions", "Questions", "Audit", "Provenance", "Notes":
			continue
		}
		switch section.Name {
		case "Added":
			section.Body = specLayout(section.Body, change.Added)
		case "Modified":
			section.Body = specLayout(section.Body, change.Modified)
		case "Removed":
			section.Body = specLayout(section.Body, change.Removed)
		}
		retained = append(retained, section)
	}
	return retained
}

// Item bodies have insignificant leading blank lines. Allocate renders one
// after each heading; comparing that separator as author text would revoke
// approval when IDs are assigned. Keep all actual body lines, including their
// trailing whitespace, and never reinterpret headings inside code fences.
func specLayout(body string, items []contracts.Spec) string {
	headings := make(map[string]contracts.Spec, len(items))
	for _, item := range items {
		headings[item.ID] = item
	}
	lines := strings.Split(body, "\n")
	result := make([]string, 0, len(lines))
	var fence byte
	depth := 0
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		left := strings.TrimLeft(line, " ")
		indented := len(line)-len(left) > 3
		if !indented && len(left) != 0 && (left[0] == '`' || left[0] == '~') {
			marker := left[0]
			count := 0
			for count < len(left) && left[count] == marker {
				count++
			}
			if count >= 3 {
				if fence == 0 {
					fence, depth = marker, count
				} else if fence == marker && count >= depth && strings.TrimSpace(left[count:]) == "" {
					fence, depth = 0, 0
				}
				result = append(result, line)
				continue
			}
		}
		fields := strings.Fields(left)
		if fence == 0 && !indented && len(fields) >= 2 && fields[0] == "###" {
			if item, present := headings[fields[1]]; present {
				heading := "### " + item.ID
				if item.Tag != "" {
					heading += " " + item.Tag
				}
				result = append(result, heading)
				for i+1 < len(lines) && strings.TrimSpace(lines[i+1]) == "" {
					i++
				}
				continue
			}
		}
		result = append(result, line)
	}
	return strings.Join(result, "\n")
}
