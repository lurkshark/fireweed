package lifecycle

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/fs"
	"maps"
	"math"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"time"

	"github.com/lurkshark/fireweed/packages/checks"
	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow/model"
)

func measure(snapshot project.Snapshot, module project.Module, report checks.Report) (model.ModuleCost, error) {
	cost := model.ModuleCost{Module: module.Name, Seconds: report.Seconds, CasesPerSpec: make(map[string]int)}
	for _, c := range report.Cases {
		seen := make(map[string]bool)
		for _, id := range c.IDs {
			if !seen[id] {
				cost.CasesPerSpec[id]++
				seen[id] = true
			}
		}
	}
	err := walkModule(snapshot, module, func(file string, entry fs.DirEntry) error {
		zone := fileZone(snapshot, module, file)
		if !entry.Type().IsRegular() || (zone != "evals" && zone != "disposable") {
			return nil
		}
		data, err := os.ReadFile(disk(snapshot.Root, file))
		if err != nil {
			return err
		}
		lines := bytes.Count(data, []byte{'\n'})
		if len(data) != 0 && data[len(data)-1] != '\n' {
			lines++
		}
		if zone == "evals" {
			cost.EvalLines += lines
		} else {
			cost.DisposableLines += lines
		}
		return nil
	})
	return cost, err
}

func budgetWarnings(snapshot project.Snapshot, costs []model.ModuleCost) []project.Diagnostic {
	config := mapping(snapshot.Config["budgets"])
	timeLimit := 15 * time.Minute
	if duration, err := time.ParseDuration(str(config["verify_time"])); err == nil && duration > 0 {
		timeLimit = duration
	}
	caseLimit := integer(config["cases_per_spec"], 15)
	ratioLimit := float64(integer(config["eval_ratio"], 4))
	if ratio, ok := config["eval_ratio"].(float64); ok {
		ratioLimit = ratio
	}
	var warnings []project.Diagnostic
	var seconds float64
	for _, cost := range costs {
		seconds += cost.Seconds
	}
	if seconds > timeLimit.Seconds() {
		warnings = append(warnings, warning("B01", "", "verification took %gs (limit %gs)", seconds, timeLimit.Seconds()))
	}
	for _, cost := range costs {
		module := moduleNamed(snapshot, cost.Module)
		for _, id := range slices.Sorted(maps.Keys(cost.CasesPerSpec)) {
			count := cost.CasesPerSpec[id]
			if count > caseLimit {
				warnings = append(warnings, warning("B02", module.Dir, "%s: %d cases (limit %d)", id, count, caseLimit))
			}
		}
		var ratio float64
		if cost.DisposableLines != 0 {
			ratio = float64(cost.EvalLines) / float64(cost.DisposableLines)
		} else if cost.EvalLines > 0 {
			ratio = math.Inf(1)
		}
		if ratio > ratioLimit {
			warnings = append(warnings, warning("B03", module.Dir, "%s: eval/disposable ratio %g (%d/%d lines, limit %g)",
				cost.Module, ratio, cost.EvalLines, cost.DisposableLines, ratioLimit))
		}
	}
	return warnings
}

func retainReport(state model.State, module string, report checks.Report) error {
	dir := runPath(state.Snapshot.Root, state.Run, state.Phase)
	if strings.EqualFold(module, "costs") {
		dir = filepath.Join(dir, "reports")
	}
	return writeJSON(filepath.Join(dir, module+".json"), report)
}

func retainCosts(state model.State, costs []model.ModuleCost, diagnostics []project.Diagnostic) error {
	snapshot := model.CostSnapshot{Run: state.Run, Phase: state.Phase, RecordedAt: time.Now().UTC(), Modules: costs, Diagnostics: diagnostics}
	return writeJSON(filepath.Join(runPath(state.Snapshot.Root, state.Run, state.Phase), "costs.json"), snapshot)
}

func writeJSON(file string, value any) error {
	dir := filepath.Dir(file)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	temp, err := os.CreateTemp(dir, ".report-*")
	if err != nil {
		return err
	}
	defer os.Remove(temp.Name())
	err = json.NewEncoder(temp).Encode(value)
	closeErr := temp.Close()
	if err != nil {
		return err
	}
	if closeErr != nil {
		return closeErr
	}
	return os.Rename(temp.Name(), file)
}

func Costs(root string) ([]model.CostSnapshot, error) {
	pattern := filepath.Join(disk(root, runsDir), "*", "*", "costs.json")
	files, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}
	result := make([]model.CostSnapshot, 0, len(files))
	for _, file := range files {
		data, err := os.ReadFile(file)
		if err != nil {
			return nil, err
		}
		var snapshot model.CostSnapshot
		if err := json.Unmarshal(data, &snapshot); err != nil {
			return nil, fmt.Errorf("read cost snapshot %s: %w", file, err)
		}
		result = append(result, snapshot)
	}
	slices.SortStableFunc(result, func(a, b model.CostSnapshot) int { return a.RecordedAt.Compare(b.RecordedAt) })
	return result, nil
}

func latestCosts(root string, modules []project.Module, filter bool) ([]model.ModuleCost, []project.Diagnostic, error) {
	snapshots, err := Costs(root)
	if err != nil {
		return nil, nil, err
	}
	eligible := make(map[string]bool)
	for _, module := range modules {
		eligible[module.Name] = true
	}
	type selectedCost struct {
		index int
		cost  model.ModuleCost
	}
	selected := make(map[string]selectedCost)
	for i, snapshot := range snapshots {
		for _, cost := range snapshot.Modules {
			if !filter || eligible[cost.Module] {
				selected[cost.Module] = selectedCost{i, cost}
			}
		}
	}

	var costs []model.ModuleCost
	bySnapshot := make(map[int][]model.ModuleCost)
	for _, name := range slices.Sorted(maps.Keys(selected)) {
		value := selected[name]
		costs = append(costs, value.cost)
		bySnapshot[value.index] = append(bySnapshot[value.index], value.cost)
	}
	var diagnostics []project.Diagnostic
	seen := make(map[project.Diagnostic]bool)
	for _, index := range slices.Sorted(maps.Keys(bySnapshot)) {
		for _, d := range snapshots[index].Diagnostics {
			if !seen[d] && measurementWarning(d, bySnapshot[index], modules) {
				diagnostics = append(diagnostics, d)
				seen[d] = true
			}
		}
	}
	return costs, diagnostics, nil
}

func measurementWarning(d project.Diagnostic, selected []model.ModuleCost, modules []project.Module) bool {
	switch d.Code {
	case "B01":
		return true
	case "B02", "B03":
		for _, cost := range selected {
			if d.Code == "B02" {
				for id := range cost.CasesPerSpec {
					if strings.Contains(d.Message, id) {
						return true
					}
				}
			} else if strings.HasPrefix(d.Message, cost.Module+":") {
				return true
			}
			for _, module := range modules {
				if module.Name == cost.Module && d.Path != "" && d.Path == module.Dir {
					return true
				}
			}
		}
	}
	return false
}
