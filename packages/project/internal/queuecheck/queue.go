// Package queuecheck validates each transition without touching repository state.
package queuecheck

import (
	"fmt"
	"sort"
	"strings"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project/internal/changeview"
	"github.com/lurkshark/fireweed/packages/project/internal/inspection"
	"github.com/lurkshark/fireweed/packages/project/internal/projection"
)

func Check(s inspection.Data) []inspection.Diagnostic {
	out := []inspection.Diagnostic{}
	add := func(code, where, message string) {
		out = append(out, inspection.Diagnostic{Code: code, Severity: "error", Path: where, Message: message})
	}
	queue := append([]contracts.Change{}, s.Changes...)
	sort.SliceStable(queue, func(i, j int) bool {
		if queue[i].Number != queue[j].Number {
			return queue[i].Number < queue[j].Number
		}
		return queue[i].Slug < queue[j].Slug
	})
	// Only this outer slice is replaced. Apply returns an independently owned
	// module, and static inspection reads its inputs without mutating them.
	state := s
	state.Modules = append([]contracts.Module{}, s.Modules...)
	state.Diagnostics = nil
	previous := signatures(inspection.Check(state))
	reserved := map[string]string{}
	for _, archived := range s.Archive {
		for _, item := range archived.Removed {
			reserved[item.ID] = archived.Path
		}
	}
	for _, ch := range queue {
		where := changeview.Path(ch)
		target := -1
		for i, m := range state.Modules {
			if m.Name == ch.Module {
				if target >= 0 {
					target = -2
					break
				}
				target = i
			}
		}
		if target < 0 {
			add("C05", where, fmt.Sprintf("Select one existing unambiguous target module instead of %q.", ch.Module))
			continue
		}
		m := state.Modules[target]
		text, err := changeview.Text(ch)
		if err == nil {
			_, err = contracts.ParseChange(text, where, m.IDPrefix)
		}
		if err != nil {
			add(changeview.ErrorCode(err), where, "Correct the typed change and declare every live handle locally: "+err.Error())
			continue
		}
		before := specMap(m)
		for _, group := range [][]contracts.Spec{ch.Modified, ch.Removed} {
			for _, item := range group {
				base, exists := before[item.ID]
				hash, stamped := ch.Bases[item.ID]
				if !exists {
					add("C10", where, fmt.Sprintf("Target %s is missing before this change; correct its edits and stamp again.", item.ID))
				} else if !stamped || hash != contracts.SpecHash(base) {
					add("C10", where, fmt.Sprintf("Base for %s is missing or stale; review preceding intent and stamp this change again.", item.ID))
				}
			}
		}
		for _, item := range ch.Added {
			if !strings.HasPrefix(item.ID, "new:") && (ch.Run == "" || ch.Phase == "") {
				add("C05", where, fmt.Sprintf("Added ID %s is author-chosen; use a local new: handle until start assigns run and phase.", item.ID))
			}
		}
		next, err := projection.Apply(m, ch)
		if err != nil {
			add("C05", where, "Correct this change's application before projecting its intent: "+err.Error())
			for _, item := range ch.Added {
				if archive, ok := reserved[item.ID]; ok {
					add("C08", where, fmt.Sprintf("ID %s was removed in %s; allocate a new ID and retain its reservation.", item.ID, archive))
				}
			}
			continue
		}
		after := specMap(next)
		for _, id := range ch.Fixes {
			_, existed := before[id]
			_, remains := after[id]
			if !existed || !remains {
				add("C05", where, fmt.Sprintf("fixes ID %s must exist before and after this change; correct the reference or edits.", id))
			}
		}
		for _, id := range ch.AlreadySatisfied {
			if _, exists := after[id]; !exists {
				add("C05", where, fmt.Sprintf("already_satisfied ID %s is absent from the result; correct the reference or edits.", id))
			}
		}
		state.Modules[target] = next
		current := inspection.Check(state)
		for _, d := range current {
			if !previous[d] {
				out = append(out, inspection.Diagnostic{Code: d.Code, Severity: d.Severity, Path: where, Message: "Projected " + d.Path + ": " + d.Message})
			}
		}
		previous = signatures(current)
	}
	sort.SliceStable(out, func(i, j int) bool {
		a, b := out[i], out[j]
		if a.Path != b.Path {
			return a.Path < b.Path
		}
		if a.Code != b.Code {
			return a.Code < b.Code
		}
		if a.Severity != b.Severity {
			return a.Severity < b.Severity
		}
		return a.Message < b.Message
	})
	return out
}

func signatures(ds []inspection.Diagnostic) map[inspection.Diagnostic]bool {
	out := make(map[inspection.Diagnostic]bool, len(ds))
	for _, d := range ds {
		out[d] = true
	}
	return out
}
func specMap(m contracts.Module) map[string]contracts.Spec {
	out := make(map[string]contracts.Spec, len(m.Specs))
	for _, item := range m.Specs {
		out[item.ID] = item
	}
	return out
}
