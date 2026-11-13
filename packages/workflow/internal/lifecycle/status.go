package lifecycle

import (
	"cmp"
	"fmt"
	"path"
	"slices"
	"strings"
	"time"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow/model"
)

func Status(root string) (model.StatusReport, error) {
	result := model.StatusReport{
		DriftSinceDrill: make(map[string]int),
		PhaseCommits:    make(map[string]int),
	}
	repo, err := gitstore.Inspect(root)
	if err != nil {
		return result, err
	}
	root = repo.Root
	snapshot, err := project.Load(root)
	if err != nil {
		return result, err
	}
	result.Modules = snapshot.Modules
	queue, err := project.Queue(root)
	if err != nil {
		return result, err
	}
	result.Queue, err = queuedStatus(root, queue)
	if err != nil {
		return result, err
	}
	result.Quarantines, result.BoundaryOptOuts, err = exceptions(snapshot)
	if err != nil {
		return result, err
	}
	_, shared, err := baseRefs(snapshot)
	if err != nil {
		return result, err
	}
	result.Unapproved, err = unapprovedCount(root, repo.Head, shared)
	if err != nil {
		return result, err
	}
	applied, err := appliedChanges(root, "")
	if err != nil {
		return result, err
	}
	dates := make(map[string]time.Time)
	for _, module := range snapshot.Modules {
		dates[module.Name], _ = drillDate(root, module.Name)
		result.DriftSinceDrill[module.Name] = 0
	}
	for _, entry := range applied {
		name := entry.change.Module
		if dates[name].IsZero() || entry.date.After(dates[name]) {
			result.DriftSinceDrill[name]++
		}
		commits, err := phaseCommits(root, repo.Head, entry.change)
		if err != nil {
			return result, err
		}
		result.PhaseCommits[numberedName(entry.change)] = len(commits)
	}
	result.Costs, result.Diagnostics, err = latestCosts(root, snapshot.Modules, false)
	if d, ok := reworkWarning(applied); ok {
		result.Diagnostics = append(result.Diagnostics, d)
	}
	return result, err
}

func exceptions(snapshot project.Snapshot) ([]model.Quarantine, []model.BoundaryOptOut, error) {
	var quarantines []model.Quarantine
	var optOuts []model.BoundaryOptOut
	for _, module := range snapshot.Modules {
		specs, err := items(module.Document, "Specs")
		if err != nil {
			return nil, nil, err
		}
		for _, spec := range specs {
			if tag, reason, found := strings.Cut(spec.Tag, ":"); found && strings.EqualFold(strings.TrimSpace(tag), "quarantine") {
				quarantines = append(quarantines, model.Quarantine{Module: module.Name, ID: spec.ID, Reason: strings.TrimSpace(reason)})
			}
		}
		if reason := str(module.Document.Frontmatter["boundaries_opt_out"]); reason != "" {
			optOuts = append(optOuts, model.BoundaryOptOut{Module: module.Name, Reason: reason})
		}
	}
	return quarantines, optOuts, nil
}

type appliedChange struct {
	change project.Change
	date   time.Time
}

func appliedChanges(root, revision string) ([]appliedChange, error) {
	paths, err := changePaths(root, revision, archiveDir)
	if err != nil {
		return nil, err
	}
	var result []appliedChange
	for _, identity := range paths {
		change, err := readChange(root, revision, identity)
		if err != nil {
			return nil, err
		}
		if str(change.Document.Frontmatter["phase"]) == "applied" {
			result = append(result, appliedChange{change, provenanceDate(change.Document)})
		}
	}
	slices.SortFunc(result, func(a, b appliedChange) int {
		if byDate := b.date.Compare(a.date); byDate != 0 {
			return byDate
		}
		return cmp.Compare(a.change.Path, b.change.Path)
	})
	return result, nil
}

func latestImplementation(root, module string) (time.Time, error) {
	applied, err := appliedChanges(root, "")
	if err != nil {
		return time.Time{}, err
	}
	for _, entry := range applied {
		if entry.change.Module == module {
			return entry.date, nil
		}
	}
	return time.Time{}, nil
}

func unapprovedCount(root, head, shared string) (int, error) {
	applied, err := appliedChanges(root, head)
	if err != nil {
		return 0, err
	}
	base, err := changePaths(root, shared, archiveDir)
	if err != nil {
		return 0, err
	}
	pending, err := changePaths(root, shared, pendingDir)
	if err != nil {
		return 0, err
	}
	approved := make(map[string]bool)
	for _, change := range append(base, pending...) {
		approved[path.Base(change.Path)] = true
	}
	count := 0
	for _, entry := range applied {
		if !approved[path.Base(entry.change.Path)] {
			count++
		}
	}
	return count, nil
}

func reworkWarning(applied []appliedChange) (project.Diagnostic, bool) {
	if len(applied) < 3 {
		return project.Diagnostic{}, false
	}
	chain := []string{path.Base(applied[0].change.Path)}
	var shared []string
	for i := 0; i < len(applied); {
		current := applied[i].change
		touched := itemIDs(slices.Concat(current.Modified, current.Removed))
		touched = append(touched, stringsOf(current.Document.Frontmatter["fixes"])...)
		touched = append(touched, stringsOf(current.Document.Frontmatter["eval_fix"])...)
		found := false
		for j := i + 1; j < len(applied) && j <= i+10; j++ {
			previous := applied[j].change
			var common []string
			for _, id := range itemIDs(slices.Concat(previous.Added, previous.Modified)) {
				if slices.Contains(touched, id) {
					common = append(common, id)
				}
			}
			if len(common) > 0 {
				chain = append(chain, path.Base(previous.Path))
				shared = append(shared, common...)
				i, found = j, true
				break
			}
		}
		if !found {
			break
		}
	}
	if len(chain) < 3 {
		return project.Diagnostic{}, false
	}
	slices.Sort(shared)
	return warning("P08", applied[0].change.Path, "rework chain %s shares %s",
		strings.Join(chain, " <- "), strings.Join(slices.Compact(shared), ", ")), true
}

func itemIDs(items []contracts.Item) []string {
	var result []string
	for _, item := range items {
		result = append(result, item.ID)
	}
	return result
}

func writeExceptions(body *strings.Builder, quarantines []model.Quarantine, optOuts []model.BoundaryOptOut, blocked []project.Change) {
	body.WriteString("\n## Project exceptions\n")
	for _, q := range quarantines {
		fmt.Fprintf(body, "- Quarantine %s %s: %s\n", q.Module, q.ID, q.Reason)
	}
	for _, opt := range optOuts {
		fmt.Fprintf(body, "- Boundary opt-out %s: %s\n", opt.Module, opt.Reason)
	}
	for _, change := range blocked {
		if reason := str(change.Document.Frontmatter["blocked"]); reason != "" {
			fmt.Fprintf(body, "- Blocked %s: %s\n", change.Path, reason)
		}
	}
}
