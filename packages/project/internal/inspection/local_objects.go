package inspection

import (
	"fmt"
	"path"
	"strings"
)

// Missing objects must never turn a local metadata read into a promisor fetch.
// --batch-all-objects inventories names without requesting missing objects.
func baselineLocalObjects(root string) (map[string]bool, error) {
	raw, err := baselineGit(root, "cat-file", "--batch-all-objects", "--batch-check=%(objectname)")
	if err != nil {
		return nil, err
	}
	objects := map[string]bool{}
	for _, oid := range strings.Fields(raw) {
		if !fullObjectID(oid) {
			return nil, fmt.Errorf("invalid local object name")
		}
		objects[oid] = true
	}
	return objects, nil
}

// Ranges and parent sets are not single commit expressions. Colons and braces
// also occur inside valid path, search and reflog expressions, so leave their
// contents to Git rather than interpreting them as revision-set operators.
func baselineRevisionSet(revision string) bool {
	for i := 0; i < len(revision); i++ {
		if revision[i] == ':' {
			return false
		}
		if i+1 < len(revision) && (revision[i] == '^' || revision[i] == '@') && revision[i+1] == '{' {
			end := strings.IndexByte(revision[i+2:], '}')
			if end < 0 {
				return false
			}
			i += end + 2
			continue
		}
		if strings.HasPrefix(revision[i:], "..") || strings.HasPrefix(revision[i:], "^@") || strings.HasPrefix(revision[i:], "^!") || strings.HasPrefix(revision[i:], "^-") {
			return true
		}
	}
	return false
}
func baselineCommit(root, revision string, objects map[string]bool) (string, error) {
	if baselineRevisionSet(revision) {
		return "", fmt.Errorf("select one commit rather than a revision set")
	}
	// --missing disables lazy retrieval before Git resolves tags/ancestors too;
	// the environment flag alone is insufficient with older Git versions.
	raw, err := baselineGit(root, "rev-list", "--no-walk", "--missing=print", "--end-of-options", revision)
	if err != nil {
		return "", err
	}
	fields := strings.Fields(raw)
	if len(fields) != 1 || !fullObjectID(fields[0]) {
		return "", fmt.Errorf("Git did not return one full commit ID")
	}
	if !objects[fields[0]] {
		return "", fmt.Errorf("selected commit is unavailable locally")
	}
	return fields[0], nil
}
func baselineTreeMetadata(root, commit string, objects map[string]bool) (string, error) {
	raw, err := baselineGit(root, "cat-file", "commit", commit)
	if err != nil {
		return "", err
	}
	header, _, _ := strings.Cut(raw, "\n")
	tree, ok := strings.CutPrefix(header, "tree ")
	if !ok || !fullObjectID(tree) {
		return "", fmt.Errorf("Git did not return a full selected tree ID")
	}
	var result strings.Builder
	active := map[string]bool{}
	var walk func(string, string) error
	walk = func(oid, prefix string) error {
		if !objects[oid] {
			return fmt.Errorf("selected tree object %s is unavailable locally", oid)
		}
		if active[oid] {
			return fmt.Errorf("cyclic selected tree metadata")
		}
		active[oid] = true
		defer delete(active, oid)
		raw, err := baselineGit(root, "ls-tree", "-z", "--full-tree", oid)
		if err != nil {
			return err
		}
		if _, err := parseCommittedTree(root, raw); err != nil {
			return err
		}
		for _, record := range strings.Split(raw, "\x00") {
			if record == "" {
				continue
			}
			header, name, _ := strings.Cut(record, "\t")
			fields := strings.Fields(header)
			relative := path.Join(prefix, name)
			result.WriteString(header + "\t" + relative + "\x00")
			if fields[1] == "tree" {
				if err := walk(fields[2], relative); err != nil {
					return err
				}
			}
		}
		return nil
	}
	if err := walk(tree, ""); err != nil {
		return "", err
	}
	return result.String(), nil
}
