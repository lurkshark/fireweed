package engine

import (
	"bytes"
	"context"
	"fmt"
	"path"
	"sort"
	"strings"
	"time"

	"github.com/lurkshark/fireweed/packages/checks/model"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project"
)

type historyChange struct {
	path, phase, module string
	set                 map[string]any
}
type historyState struct {
	snapshot project.Snapshot
	changes  map[string]historyChange
}
type historyDiff struct{ status, old, path string }

func historyGit(root string, args ...string) ([]byte, error) {
	argv := append([]string{"git", "-C", root, "-c", "core.hooksPath=/dev/null", "--no-optional-locks"}, args...)
	result, err := gitstore.Run(context.Background(), gitstore.Request{Args: argv, Timeout: 30 * time.Second,
		Env: map[string]string{"GIT_NO_LAZY_FETCH": "1", "GIT_TERMINAL_PROMPT": "0"}})
	if err != nil {
		return nil, fmt.Errorf("git %s: %w: %s", strings.Join(args, " "), err, result.Stderr)
	}
	return []byte(result.Stdout), nil
}
func historyLines(data []byte) []string { return strings.Fields(string(data)) }
func historyParents(root, commit string) ([]string, error) {
	data, err := historyGit(root, "rev-list", "--parents", "-n", "1", commit)
	if err != nil {
		return nil, err
	}
	fields := historyLines(data)
	if len(fields) == 0 {
		return nil, fmt.Errorf("missing commit %s", commit)
	}
	return fields[1:], nil
}
func historyRange(root, revisionRange string) ([]string, error) {
	args := []string{"rev-list", "--reverse", "--topo-order", "--end-of-options"}
	if revisionRange != "" {
		args = append(args, revisionRange)
	} else {
		head, err := project.LoadAt(root, "HEAD")
		if err != nil {
			return nil, err
		}
		base, _ := head.Config["base"].(string)
		remote, _ := head.Config["remote"].(string)
		ref, err := gitstore.BaseRef(root, remote, base)
		if err != nil {
			return nil, err
		}
		data, err := historyGit(root, "merge-base", "HEAD", ref)
		if err != nil {
			return nil, err
		}
		mergeBase := strings.TrimSpace(string(data))
		if mergeBase == "" {
			return nil, fmt.Errorf("HEAD and %s have no merge base", ref)
		}
		args = append(args, mergeBase+"..HEAD")
	}
	data, err := historyGit(root, args...)
	if err != nil {
		return nil, err
	}
	return historyLines(data), nil
}
func historyLoad(root, commit string) (historyState, error) {
	snapshot, err := project.LoadAt(root, commit)
	if err != nil {
		return historyState{}, err
	}
	data, err := historyGit(root, "ls-tree", "-r", "--name-only", "-z", commit, "--", "changes")
	if err != nil {
		return historyState{}, err
	}
	state := historyState{snapshot: snapshot, changes: map[string]historyChange{}}
	for _, entry := range bytes.Split(data, []byte{0}) {
		file := string(entry)
		if file == "" || !strings.HasPrefix(file, "changes/") || !strings.HasSuffix(file, ".md") {
			continue
		}
		body, err := gitstore.ReadAt(root, commit, file)
		if err != nil {
			return historyState{}, fmt.Errorf("read %s at %s: %w", file, commit, err)
		}
		doc, err := contracts.ParseDocument(string(body))
		if err != nil {
			return historyState{}, fmt.Errorf("parse %s at %s: %w", file, commit, err)
		}
		phase, ok := doc.Frontmatter["phase"].(string)
		if _, exists := doc.Frontmatter["phase"]; exists && !ok {
			return historyState{}, fmt.Errorf("%s at %s: invalid phase", file, commit)
		}
		module, ok := doc.Frontmatter["module"].(string)
		if !ok || module == "" {
			return historyState{}, fmt.Errorf("%s at %s: missing module", file, commit)
		}
		var set map[string]any
		if raw, exists := doc.Frontmatter["set"]; exists {
			set, ok = raw.(map[string]any)
			if !ok {
				return historyState{}, fmt.Errorf("%s at %s: invalid set", file, commit)
			}
			for key, value := range set {
				switch key {
				case "depends_on", "interface", "evals", "disposable", "preserve", "append_only":
					list, ok := value.([]any)
					if !ok {
						return historyState{}, fmt.Errorf("%s at %s: set %s must be a string list", file, commit, key)
					}
					for _, item := range list {
						if _, ok := item.(string); !ok {
							return historyState{}, fmt.Errorf("%s at %s: set %s must be a string list", file, commit, key)
						}
					}
				case "hooks", "boundaries_opt_out":
				default:
					return historyState{}, fmt.Errorf("%s at %s: set cannot change %q", file, commit, key)
				}
			}
		}
		key := path.Base(file)
		if prior, exists := state.changes[key]; exists {
			return historyState{}, fmt.Errorf("duplicate changes %s and %s at %s", prior.path, file, commit)
		}
		state.changes[key] = historyChange{path: file, phase: phase, module: module, set: set}
	}
	return state, nil
}
func historyDiffs(root, before, after string) ([]historyDiff, error) {
	data, err := historyGit(root, "diff", "--no-ext-diff", "--no-textconv", "--name-status", "-z", "-M", before, after, "--")
	if err != nil {
		return nil, err
	}
	fields := bytes.Split(data, []byte{0})
	var result []historyDiff
	for i := 0; i < len(fields)-1; {
		status := string(fields[i])
		i++
		if status == "" || i >= len(fields)-1 {
			return nil, fmt.Errorf("malformed diff %s..%s", before, after)
		}
		item := historyDiff{status: status[:1], path: string(fields[i])}
		i++
		if item.status == "R" || item.status == "C" {
			if i >= len(fields)-1 {
				return nil, fmt.Errorf("malformed rename %s..%s", before, after)
			}
			item.old, item.path = item.path, string(fields[i])
			i++
		}
		result = append(result, item)
	}
	return result, nil
}
func historyCombined(root, commit string) (map[string]bool, error) {
	data, err := historyGit(root, "diff-tree", "--cc", "--name-only", "-z", "--no-commit-id", commit)
	if err != nil {
		return nil, err
	}
	paths := map[string]bool{}
	for _, field := range bytes.Split(data, []byte{0}) {
		if len(field) > 0 {
			paths[string(field)] = true
		}
	}
	return paths, nil
}
func historyZone(snapshot project.Snapshot, file string, set map[string]any, moduleName string) string {
	var owner *project.Module
	for i := range snapshot.Modules {
		module := &snapshot.Modules[i]
		prefix := strings.TrimSuffix(module.Dir, "/") + "/"
		if strings.HasPrefix(file, prefix) && (owner == nil || len(module.Dir) > len(owner.Dir)) {
			owner = module
		}
	}
	if owner == nil {
		return ""
	}
	zones := owner.Zones
	if owner.Name == moduleName {
		for key, raw := range set {
			values := patterns(raw)
			switch key {
			case "evals":
				zones.Evals = values
			case "interface":
				zones.Interface = values
			case "append_only":
				zones.AppendOnly = values
			case "disposable":
				zones.Disposable = values
			case "preserve":
				zones.Preserve = values
			}
		}
	}
	return contracts.ResolveZone(strings.TrimPrefix(file, strings.TrimSuffix(owner.Dir, "/")+"/"), zones)
}
func historyProtected(snapshot project.Snapshot, file string) bool {
	if matches([]string{".git", ".git/**", ".fireweed", ".fireweed/**", "SYSTEM.md", "fireweed.yaml"}, file) ||
		matches(patterns(snapshot.Config["protected"]), file) || path.Base(file) == "AGENTS.md" || path.Base(file) == "CLAUDE.md" {
		return true
	}
	zone := historyZone(snapshot, file, nil, "")
	return zone == "protected"
}
func historyWorker(after historyState, parents []historyState) []historyChange {
	var assigned []historyChange
	keys := make([]string, 0, len(after.changes))
	for key := range after.changes {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		change := after.changes[key]
		if change.phase == "" {
			continue
		}
		newPhase := true
		for _, parent := range parents {
			if prior, ok := parent.changes[key]; ok && prior.phase == change.phase {
				newPhase = false
				break
			}
		}
		if newPhase {
			assigned = append(assigned, change)
		}
	}
	return assigned
}
func History(root, revisionRange string) (model.HistoryReport, error) {
	var report model.HistoryReport
	commits, err := historyRange(root, revisionRange)
	if err != nil {
		return report, err
	}
	report.Commits = commits
	cache := map[string]historyState{}
	load := func(commit string) (historyState, error) {
		if state, ok := cache[commit]; ok {
			return state, nil
		}
		state, err := historyLoad(root, commit)
		if err == nil {
			cache[commit] = state
		}
		return state, err
	}
	for _, commit := range commits {
		after, err := load(commit)
		if err != nil {
			return model.HistoryReport{}, err
		}
		parentIDs, err := historyParents(root, commit)
		if err != nil {
			return model.HistoryReport{}, err
		}
		parents := make([]historyState, len(parentIDs))
		var diffs []historyDiff
		for i, parentID := range parentIDs {
			parents[i], err = load(parentID)
			if err != nil {
				return model.HistoryReport{}, err
			}
			part, err := historyDiffs(root, parentID, commit)
			if err != nil {
				return model.HistoryReport{}, err
			}
			diffs = append(diffs, part...)
		}
		if len(parentIDs) == 0 {
			diffs, err = historyDiffs(root, "4b825dc642cb6eb9a060e54bf8d69288fbee4904", commit)
			if err != nil {
				return model.HistoryReport{}, err
			}
		}
		if len(parentIDs) > 1 {
			combined, err := historyCombined(root, commit)
			if err != nil {
				return model.HistoryReport{}, err
			}
			filtered := diffs[:0]
			for _, diff := range diffs {
				if combined[diff.path] || combined[diff.old] {
					filtered = append(filtered, diff)
				}
			}
			diffs = filtered
		}
		workers := historyWorker(after, parents)
		var active []historyChange
		for _, change := range workers {
			squashed := change.phase == "applied"
			if squashed {
				for _, parent := range parents {
					if prior, ok := parent.changes[path.Base(change.path)]; ok && prior.phase == "impl" {
						squashed = false
						break
					}
				}
			}
			if squashed {
				report.Squashed = append(report.Squashed, model.SquashedChange{Commit: commit, Change: change.path})
			} else {
				active = append(active, change)
			}
		}
		seen := map[string]bool{}
		add := func(code, file string) {
			key := code + "\x00" + file
			if seen[key] {
				return
			}
			seen[key] = true
			report.Diagnostics = append(report.Diagnostics, project.Diagnostic{Code: code, Path: file, Message: fmt.Sprintf("%s at commit %s", code, commit)})
		}
		for _, diff := range diffs {
			files := []string{diff.path}
			if diff.old != "" {
				files = append(files, diff.old)
			}
			for _, file := range files {
				for _, worker := range active {
					if historyProtected(after.snapshot, file) {
						add("H01", file)
					} else {
						for _, parent := range parents {
							if historyProtected(parent.snapshot, file) {
								add("H01", file)
								break
							}
						}
					}
					if worker.phase == "applied" {
						implParent := false
						for _, parent := range parents {
							if prior, ok := parent.changes[path.Base(worker.path)]; ok && prior.phase == "impl" {
								implParent = true
								break
							}
						}
						if implParent {
							zone := historyZone(after.snapshot, file, worker.set, worker.module)
							if zone == "evals" || zone == "interface" {
								add("H02", file)
							}
							for _, parent := range parents {
								zone := historyZone(parent.snapshot, file, worker.set, worker.module)
								if zone == "evals" || zone == "interface" {
									add("H02", file)
									break
								}
							}
						}
					}
				}
				if diff.status != "A" && diff.status != "C" && (diff.old == "" || file == diff.old) {
					if historyZone(after.snapshot, file, nil, "") == "append_only" {
						add("H03", file)
					} else {
						for _, parent := range parents {
							if historyZone(parent.snapshot, file, nil, "") == "append_only" {
								add("H03", file)
								break
							}
						}
					}
					for _, worker := range workers {
						if historyZone(after.snapshot, file, worker.set, worker.module) == "append_only" {
							add("H03", file)
						}
						for _, parent := range parents {
							if historyZone(parent.snapshot, file, worker.set, worker.module) == "append_only" {
								add("H03", file)
							}
						}
					}
				}
				if strings.HasPrefix(file, "changes/archive/") && diff.status != "A" && diff.status != "D" && (diff.old == "" || file == diff.old) {
					add("H04", file)
				}
			}
		}
	}
	return report, nil
}
