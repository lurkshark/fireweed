package lifecycle

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"time"

	"github.com/lurkshark/fireweed/packages/checks"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow/model"
)

func measureModuleCost(snapshot project.Snapshot, name string, report checks.Report) (model.ModuleCost, error) {
	m := module(snapshot, name)
	cost := model.ModuleCost{Module: name, Seconds: report.Seconds, CasesPerSpec: map[string]int{}}
	for _, testcase := range report.Cases {
		seen := map[string]bool{}
		for _, id := range testcase.IDs {
			if !seen[id] {
				cost.CasesPerSpec[id]++
				seen[id] = true
			}
		}
	}
	moduleDir := filepath.Join(snapshot.Root, filepath.FromSlash(m.Dir))
	protected := list(snapshot.Config["protected"])
	err := filepath.WalkDir(moduleDir, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if !entry.Type().IsRegular() {
			return nil
		}
		rel, err := filepath.Rel(moduleDir, path)
		if err != nil {
			return err
		}
		rootRel, err := filepath.Rel(snapshot.Root, path)
		if err != nil {
			return err
		}
		for _, pattern := range protected {
			if contracts.MatchGlob(pattern, filepath.ToSlash(rootRel)) {
				return nil
			}
		}
		zone := contracts.ResolveZone(filepath.ToSlash(rel), m.Zones)
		if zone != "evals" && zone != "disposable" {
			return nil
		}
		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		lines := bytes.Count(content, []byte{'\n'})
		if len(content) > 0 && content[len(content)-1] != '\n' {
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

func costWarnings(costs []model.ModuleCost, config map[string]any) []project.Diagnostic {
	budgets := mapval(config["budgets"])
	verifyLimit := 15 * time.Minute
	if configured, err := time.ParseDuration(str(budgets["verify_time"])); err == nil {
		verifyLimit = configured
	}
	caseLimit := num(budgets["cases_per_spec"], 15)
	ratioLimit := 4.0
	switch value := budgets["eval_ratio"].(type) {
	case int:
		ratioLimit = float64(value)
	case float64:
		ratioLimit = value
	case string:
		if parsed, err := strconv.ParseFloat(value, 64); err == nil {
			ratioLimit = parsed
		}
	}
	var total float64
	var warnings []project.Diagnostic
	for _, cost := range costs {
		total += cost.Seconds
		ids := make([]string, 0, len(cost.CasesPerSpec))
		for id := range cost.CasesPerSpec {
			ids = append(ids, id)
		}
		sort.Strings(ids)
		for _, id := range ids {
			if count := cost.CasesPerSpec[id]; count > caseLimit {
				warnings = append(warnings, diag("B02", fmt.Sprintf("%s: %d testcases exceeds cases_per_spec limit %d", id, count, caseLimit), true))
			}
		}
		if cost.EvalLines > 0 && (cost.DisposableLines == 0 || float64(cost.EvalLines) > ratioLimit*float64(cost.DisposableLines)) {
			warnings = append(warnings, diag("B03", fmt.Sprintf("%s: eval/disposable ratio %d/%d exceeds eval_ratio limit %g", cost.Module, cost.EvalLines, cost.DisposableLines, ratioLimit), true))
		}
	}
	if total > verifyLimit.Seconds() {
		warnings = append([]project.Diagnostic{diag("B01", fmt.Sprintf("evaluated modules: %g seconds exceeds verify_time limit %s (%g seconds)", total, verifyLimit, verifyLimit.Seconds()), true)}, warnings...)
	}
	return warnings
}

func costPath(root, run, phase string) string {
	return filepath.Join(root, ".fireweed", "runs", run, phase, "costs.json")
}

func saveCosts(state model.State, phase string, costs []model.ModuleCost, warnings []project.Diagnostic) error {
	snapshot := model.CostSnapshot{Run: state.Run, Phase: phase, RecordedAt: time.Now().UTC(), Modules: costs, Diagnostics: warnings}
	data, err := json.MarshalIndent(snapshot, "", "  ")
	if err != nil {
		return err
	}
	path := costPath(state.Snapshot.Root, state.Run, phase)
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

func Costs(root string) ([]model.CostSnapshot, error) {
	runsDir := filepath.Join(root, ".fireweed", "runs")
	runs, err := os.ReadDir(runsDir)
	if os.IsNotExist(err) {
		return []model.CostSnapshot{}, nil
	}
	if err != nil {
		return nil, err
	}
	snapshots := []model.CostSnapshot{}
	for _, run := range runs {
		if !run.IsDir() {
			continue
		}
		phases, err := os.ReadDir(filepath.Join(runsDir, run.Name()))
		if err != nil {
			return nil, err
		}
		for _, phase := range phases {
			if !phase.IsDir() {
				continue
			}
			data, err := os.ReadFile(costPath(root, run.Name(), phase.Name()))
			if os.IsNotExist(err) {
				continue
			}
			if err != nil {
				return nil, err
			}
			var snapshot model.CostSnapshot
			if err := json.Unmarshal(data, &snapshot); err != nil {
				return nil, fmt.Errorf("%s: %w", costPath(root, run.Name(), phase.Name()), err)
			}
			snapshots = append(snapshots, snapshot)
		}
	}
	sort.Slice(snapshots, func(i, j int) bool {
		if !snapshots[i].RecordedAt.Equal(snapshots[j].RecordedAt) {
			return snapshots[i].RecordedAt.Before(snapshots[j].RecordedAt)
		}
		if snapshots[i].Run != snapshots[j].Run {
			return snapshots[i].Run < snapshots[j].Run
		}
		return snapshots[i].Phase < snapshots[j].Phase
	})
	return snapshots, nil
}
