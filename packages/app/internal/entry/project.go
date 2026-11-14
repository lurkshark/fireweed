package entry

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow"
)

func loadGraph(root string) (project.Snapshot, error) {
	snapshot, err := project.Load(root)
	if err == nil {
		err = validateGraph(snapshot.Modules)
	}
	return snapshot, err
}

func validateGraph(modules []project.Module) error {
	if err := project.ValidateGraph(modules); err != nil {
		return coded("C02", "", err)
	}
	return nil
}

func approvedConfigFailure(err error, slug string) response {
	return failed(2, fmt.Errorf("approved configuration for %q: %w", slug, err), "C01",
		"Use a compatible Fireweed binary, or correct fireweed.yaml and obtain a new approved snapshot.")
}

// Check for an owned snapshot at the Git worktree root even if fireweed.yaml
// has been removed. An enclosing Git worktree cannot override a nested project.
func discoverApproved(dir string) (string, *workflow.ApprovedSnapshot, error) {
	for root := dir; ; root = filepath.Dir(root) {
		config, err := exists(filepath.Join(root, "fireweed.yaml"))
		if err != nil {
			return "", nil, err
		}
		git, err := exists(filepath.Join(root, ".git"))
		if err != nil {
			return "", nil, err
		}
		if git {
			approved, held, err := workflow.Approved(root)
			if held {
				return root, &approved, nil
			}
			if err != nil && cause(err, "") != "" {
				return "", nil, err
			}
		}
		if config {
			return root, nil, nil
		}
		if filepath.Dir(root) == root {
			return "", nil, fmt.Errorf("fireweed.yaml not found from %s", dir)
		}
	}
}

func moduleByName(snapshot project.Snapshot, name string) (project.Module, bool) {
	for _, module := range snapshot.Modules {
		if module.Name == name {
			return module, true
		}
	}
	return project.Module{}, false
}

func specItems(module project.Module) ([]contracts.Item, error) {
	section, _ := contracts.Section(module.Document.Body, "Specs")
	items, err := contracts.ParseItems(section)
	if err != nil {
		file := path.Join(module.Dir, "SPEC.md")
		return nil, coded(cause(err, "C03"), file, fmt.Errorf("%s: %w", file, err))
	}
	return items, nil
}

func (c command) projectCommand(root string) response {
	switch c.name {
	case "new-module":
		return c.newModule(root)
	case "new-change":
		file, err := project.NewChange(root, c.arg(0), c.arg(1))
		if err != nil {
			if cause(err, "") == "C01" {
				return configFailure(err)
			}
			return failed(2, err, "P01", "Use a valid module and an unused slug, and check the changes directory.")
		}
		return message(file)
	case "stamp":
		if _, err := project.Load(root); err != nil && cause(err, "") == "C01" {
			return configFailure(err)
		}
		if err := project.Stamp(root, c.arg(0)); err != nil {
			return configuredOperationFailure(err, "Use fireweed list to check the selector and correct the change's edits.")
		}
		return message("Stamped " + c.arg(0) + ".")
	case "list":
		if _, err := project.Load(root); err != nil {
			return configFailure(err)
		}
		entries, err := project.Queue(root)
		if err != nil {
			return operationFailure(err, "Correct the pending change documents and retry.")
		}
		var plain strings.Builder
		rows := make([]queueRow, 0, len(entries))
		for _, entry := range entries {
			front := entry.Document.Frontmatter
			phase, _ := front["phase"].(string)
			blocked, _ := front["blocked"].(string)
			rounds, _ := front["audit_rounds"].(int)
			row := makeQueueRow(entry.Slug, entry.Path, entry.Module, phase, blocked, rounds)
			fmt.Fprintln(&plain, row.text())
			rows = append(rows, row)
		}
		return output(plain.String(), rows)
	case "show":
		return c.show(root)
	case "agents-md":
		if _, err := project.Load(root); err != nil && cause(err, "") == "C01" {
			return configFailure(err)
		}
		if err := refreshAgents(root); err != nil {
			return failed(2, err, "P01", "Correct AGENTS.md and use fireweed agents-md --write again.")
		}
		return message("Updated AGENTS.md.")
	case "migrate":
		snapshot, err := project.Load(root)
		if err != nil {
			return configFailure(err)
		}
		if snapshot.Config["version"] != 1 {
			return configFailure(fmt.Errorf("fireweed.yaml: unsupported project format %v (expected 1)", snapshot.Config["version"]))
		}
		return checked(verbatim("Project format 1 is already current.\n"), project.Check(root), "Correct the reported static issue and run fireweed check.")
	}
	panic("unexpected project command: " + c.name)
}

func (c command) show(root string) response {
	switch {
	case c.has("--stale"):
		return showStale(root, c.arg(0))
	case c.has("--history"):
		return showHistory(root, c.arg(0), c.arg(1))
	}

	var snapshot project.Snapshot
	var err error
	projected := c.has("--projected")
	if projected {
		snapshot, err = project.Project(root, "")
	} else {
		snapshot, err = loadGraph(root)
	}
	if err != nil {
		return configFailure(err)
	}
	name, id := c.arg(0), c.arg(1)
	module, found := moduleByName(snapshot, name)
	if !found {
		return failure(2, "U02", "", fmt.Sprintf("unknown module %q. Use fireweed modules to list modules.", name))
	}
	if id == "" && !projected {
		file := path.Join(module.Dir, "SPEC.md")
		data, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(file)))
		if err != nil {
			return configFailure(coded("C03", file, err))
		}
		return verbatim(string(data))
	}
	if id == "" {
		body, err := contracts.RenderDocument(module.Document)
		if err != nil {
			return configFailure(err)
		}
		return verbatim(body)
	}
	items, err := specItems(module)
	if projected && cause(err, "") == "C05" {
		// A projection can repeat a local handle across changes. Parse the
		// blocks separately so that unrelated handles remain selectable.
		section, _ := contracts.Section(module.Document.Body, "Specs")
		items, err = projectedItems(section)
	}
	if err != nil {
		return configFailure(err)
	}
	var match *contracts.Item
	for i := range items {
		item := &items[i]
		if item.ID != id {
			continue
		}
		if match != nil {
			return failure(2, "U02", "", fmt.Sprintf("duplicate item ID: %s. Use fireweed show %s --projected to inspect the full projection.", id, name))
		}
		match = item
	}
	if match != nil {
		return verbatim(contracts.RenderItem(*match))
	}
	return failure(2, "U02", "", fmt.Sprintf("unknown item %q in %q. Use fireweed show %s to list its items.", id, name, name))
}

func projectedItems(section string) ([]contracts.Item, error) {
	var items []contracts.Item
	fixed := make(map[string]bool)
	lines := strings.SplitAfter(section, "\n")
	for start := 0; start < len(lines); {
		end := start + 1
		for end < len(lines) && !strings.HasPrefix(lines[end], "### ") {
			end++
		}
		block, err := contracts.ParseItems(strings.Join(lines[start:end], ""))
		if err != nil {
			return nil, err
		}
		for _, item := range block {
			if !strings.HasPrefix(item.ID, "new:") {
				if fixed[item.ID] {
					return nil, coded("C05", item.ID, fmt.Errorf("duplicate item ID: %s", item.ID))
				}
				fixed[item.ID] = true
			}
		}
		items = append(items, block...)
		start = end
	}
	return items, nil
}

type historyRow struct {
	Path               string   `json:"path"`
	Module             string   `json:"module"`
	IDs                []string `json:"ids"`
	Motivation         string   `json:"motivation"`
	Decisions          string   `json:"decisions"`
	RemovedDecisions   string   `json:"removed_decisions"`
	Provenance         string   `json:"provenance"`
	ImplementationDate string   `json:"implementation_date"`
}

func showHistory(root, module, id string) response {
	entries, err := project.SpecHistory(root, module, id)
	if err != nil {
		switch cause(err, "") {
		case "C01":
			return configFailure(err)
		case "U02":
			return failed(2, err, "U02", "Use fireweed modules or fireweed show to find the module or spec ID.")
		default:
			return operationFailure(err, "Correct the project documents and retry.")
		}
	}
	rows := make([]historyRow, 0, len(entries))
	var plain strings.Builder
	if len(entries) == 0 {
		fmt.Fprintln(&plain, "No history for the selected specs.")
	}
	for _, entry := range entries {
		date, label := "", "undated"
		if !entry.ImplementationDate.IsZero() {
			date = entry.ImplementationDate.Format(time.RFC3339)
			label = date
		}
		rows = append(rows, historyRow{
			Path: entry.Path, Module: entry.Module, IDs: entry.IDs, Motivation: entry.Motivation,
			Decisions: entry.Decisions, RemovedDecisions: entry.RemovedDecisions,
			Provenance: entry.Provenance, ImplementationDate: date,
		})
		fmt.Fprintf(&plain, "%s\nModule: %s\nIDs: %s\nImplementation date: %s\n", entry.Path, entry.Module, strings.Join(entry.IDs, ", "), label)
		fmt.Fprintf(&plain, "Motivation\n%s\nDecisions\n%s\nRemoved decisions\n%s\nProvenance\n%s\n", entry.Motivation, entry.Decisions, entry.RemovedDecisions, entry.Provenance)
	}
	return output(plain.String(), rows)
}

func showStale(root, selector string) response {
	if _, err := loadGraph(root); err != nil {
		return configFailure(err)
	}
	entries, err := project.StaleSpecs(root, selector)
	if err != nil {
		status := 1
		switch cause(err, "") {
		case "U02", "C13", "C01":
			status = 2
		}
		return failed(status, fmt.Errorf("%s: %w", selector, err), "P01", "Use fireweed list to check the selector and correct the queued change.")
	}
	var plain strings.Builder
	if len(entries) == 0 {
		fmt.Fprintln(&plain, "No stale specs.")
	}
	for _, entry := range entries {
		fmt.Fprintln(&plain, entry.ID)
		if entry.Recovered {
			lineDiff(&plain, entry.Before, entry.Current)
		} else {
			fmt.Fprintf(&plain, "Recorded hash: %s (stamped text unavailable)\nCurrent\n%s\n", entry.RecordedHash, entry.Current)
		}
	}
	return verbatim(plain.String())
}

func lineDiff(out *strings.Builder, before, current string) {
	oldLines := strings.Split(strings.TrimSuffix(before, "\n"), "\n")
	newLines := strings.Split(strings.TrimSuffix(current, "\n"), "\n")
	// Longest common subsequence, computed from the end for a forward diff.
	common := make([][]int, len(oldLines)+1)
	for i := len(oldLines); i >= 0; i-- {
		common[i] = make([]int, len(newLines)+1)
		for j := len(newLines) - 1; i < len(oldLines) && j >= 0; j-- {
			if oldLines[i] == newLines[j] {
				common[i][j] = 1 + common[i+1][j+1]
			} else {
				common[i][j] = max(common[i+1][j], common[i][j+1])
			}
		}
	}
	for i, j := 0, 0; i < len(oldLines) || j < len(newLines); {
		switch {
		case i < len(oldLines) && j < len(newLines) && oldLines[i] == newLines[j]:
			fmt.Fprintln(out, " "+oldLines[i])
			i++
			j++
		case j == len(newLines) || i < len(oldLines) && common[i+1][j] >= common[i][j+1]:
			fmt.Fprintln(out, "-"+oldLines[i])
			i++
		default:
			fmt.Fprintln(out, "+"+newLines[j])
			j++
		}
	}
}
