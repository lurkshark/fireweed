package lifecycle

import (
	"sort"
	"strings"

	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow/model"
)

func statusCosts(root string, report *model.StatusReport) error {
	snapshots, err := Costs(root)
	if err != nil {
		return err
	}
	selected := map[string]int{}
	for i, snapshot := range snapshots {
		for _, cost := range snapshot.Modules {
			selected[cost.Module] = i
		}
	}
	names := make([]string, 0, len(selected))
	for name := range selected {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		for _, cost := range snapshots[selected[name]].Modules {
			if cost.Module == name {
				report.Costs = append(report.Costs, cost)
				break
			}
		}
	}
	for i, snapshot := range snapshots {
		for _, warning := range snapshot.Diagnostics {
			if !selectedCostWarning(i, snapshot, selected, warning) {
				continue
			}
			found := false
			for _, existing := range report.Diagnostics {
				if existing == warning {
					found = true
					break
				}
			}
			if !found {
				report.Diagnostics = append(report.Diagnostics, warning)
			}
		}
	}
	return nil
}

func selectedCostWarning(index int, snapshot model.CostSnapshot, selected map[string]int, warning project.Diagnostic) bool {
	switch warning.Code {
	case "B01":
		if len(snapshot.Modules) == 0 {
			return false
		}
		for _, cost := range snapshot.Modules {
			if selected[cost.Module] != index {
				return false
			}
		}
		return true
	case "B02":
		for _, cost := range snapshot.Modules {
			if selected[cost.Module] != index {
				continue
			}
			for id := range cost.CasesPerSpec {
				if strings.HasPrefix(warning.Message, id+":") {
					return true
				}
			}
		}
	case "B03":
		for _, cost := range snapshot.Modules {
			if selected[cost.Module] == index && strings.HasPrefix(warning.Message, cost.Module+":") {
				return true
			}
		}
	}
	return false
}
