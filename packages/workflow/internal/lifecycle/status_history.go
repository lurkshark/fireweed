package lifecycle

import (
	"path"
	"strings"
)

// One traversal serves every applied archive. Only commits touching a selected
// document can change its phase relative to the first parent.
func phaseCommitCounts(root, head string, archives []appliedChange) (map[string]int, error) {
	counts := make(map[string]int, len(archives))
	wanted := make(map[string]bool, len(archives))
	for _, change := range archives {
		wanted[change.name] = true
		counts[strings.TrimSuffix(change.name, ".md")] = 0
	}
	if len(wanted) == 0 {
		return counts, nil
	}
	out, err := commitsGit(root, "rev-list", "--topo-order", "--parents", head)
	if err != nil {
		return nil, err
	}
	cache := map[string]string{}
	phaseAt := func(oid, name string) (string, error) {
		if oid == "" {
			return "", nil
		}
		key := oid + ":" + name
		if phase, ok := cache[key]; ok {
			return phase, nil
		}
		phase, err := commitPhaseAt(root, oid, name)
		if err == nil {
			cache[key] = phase
		}
		return phase, err
	}
	for _, line := range strings.Split(strings.TrimSpace(out), "\n") {
		fields := strings.Fields(line)
		if len(fields) == 0 {
			continue
		}
		parent := ""
		var args []string
		if len(fields) > 1 {
			parent = fields[1]
			args = []string{"diff-tree", "--name-only", "-r", "-z", parent, fields[0], "--", "changes"}
		} else {
			args = []string{"diff-tree", "--root", "--no-commit-id", "--name-only", "-r", "-z", fields[0], "--", "changes"}
		}
		changed, err := commitsGit(root, args...)
		if err != nil {
			return nil, err
		}
		seen := map[string]bool{}
		for _, p := range strings.Split(changed, "\x00") {
			name := path.Base(p)
			if !wanted[name] || seen[name] {
				continue
			}
			seen[name] = true
			phase, err := phaseAt(fields[0], name)
			if err != nil {
				return nil, err
			}
			previous, err := phaseAt(parent, name)
			if err != nil {
				return nil, err
			}
			if phase != previous {
				counts[strings.TrimSuffix(name, ".md")]++
			}
		}
	}
	return counts, nil
}
