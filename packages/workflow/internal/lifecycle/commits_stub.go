package lifecycle

import (
	"context"
	"fmt"
	"path"
	"regexp"
	"strings"
	"time"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
)

var commitChangeName = regexp.MustCompile(`^[0-9]{4,}-[a-z0-9-]+\.md$`)

func commitsGit(root string, args ...string) (string, error) {
	result, err := gitstore.Run(context.Background(), gitstore.Request{
		Args:    append([]string{"git", "-C", root, "-c", "core.hooksPath=/dev/null"}, args...),
		Timeout: 30 * time.Second,
		Env: map[string]string{
			"GIT_NO_LAZY_FETCH": "1", "GIT_TERMINAL_PROMPT": "0", "GIT_OPTIONAL_LOCKS": "0",
		},
	})
	if err != nil {
		return "", fmt.Errorf("git %s: %w: %s", strings.Join(args, " "), err, result.Stderr)
	}
	return result.Stdout, nil
}

func commitChangePaths(root, revision string) ([]string, error) {
	out, err := commitsGit(root, "ls-tree", "-r", "--name-only", "-z", revision, "--", "changes")
	if err != nil {
		return nil, err
	}
	var paths []string
	for _, candidate := range strings.Split(out, "\x00") {
		if candidate == "" {
			continue
		}
		dir := path.Dir(candidate)
		if (dir == "changes" || dir == "changes/archive") && commitChangeName.MatchString(path.Base(candidate)) {
			paths = append(paths, candidate)
		}
	}
	return paths, nil
}

func selectCommitChange(paths []string, selector string) (string, error) {
	var matches []string
	for _, candidate := range paths {
		base := path.Base(candidate)
		slug := strings.TrimSuffix(base[strings.IndexByte(base, '-')+1:], ".md")
		if selector == candidate {
			return base, nil
		}
		if selector == base || selector == slug {
			matches = append(matches, base)
		}
	}
	if len(matches) == 1 {
		return matches[0], nil
	}
	if len(matches) == 0 {
		return "", fmt.Errorf("U02: change %q not found", selector)
	}
	return "", fmt.Errorf("U02: change %q is ambiguous", selector)
}

func commitPhaseAt(root, revision, basename string) (string, error) {
	if revision == "" {
		return "", nil
	}
	pending := "changes/" + basename
	archived := "changes/archive/" + basename
	out, err := commitsGit(root, "ls-tree", "-r", "--name-only", "-z", revision, "--", pending, archived)
	if err != nil {
		return "", err
	}
	var found string
	for _, candidate := range strings.Split(out, "\x00") {
		if candidate != pending && candidate != archived {
			continue
		}
		if found != "" {
			return "", fmt.Errorf("change %s has pending and archived copies at %s", basename, revision)
		}
		found = candidate
	}
	if found == "" {
		return "", nil
	}
	data, err := gitstore.ReadAt(root, revision, found)
	if err != nil {
		return "", fmt.Errorf("read %s at %s: %w", found, revision, err)
	}
	doc, err := contracts.ParseDocument(string(data))
	if err != nil {
		return "", fmt.Errorf("parse %s at %s: %w", found, revision, err)
	}
	phase, _ := doc.Frontmatter["phase"].(string)
	return phase, nil
}

// Commits returns reachable commits that changed this document's phase.
func Commits(root, change string) ([]string, error) {
	repo, err := gitstore.Inspect(root)
	if err != nil {
		return nil, err
	}
	root = repo.Root
	paths, err := commitChangePaths(root, repo.Head)
	if err != nil {
		return nil, err
	}
	basename, err := selectCommitChange(paths, change)
	if err != nil {
		return nil, err
	}
	out, err := commitsGit(root, "rev-list", "--topo-order", "--parents", repo.Head)
	if err != nil {
		return nil, err
	}
	cache := map[string]string{}
	phaseAt := func(oid string) (string, error) {
		if phase, ok := cache[oid]; ok {
			return phase, nil
		}
		phase, err := commitPhaseAt(root, oid, basename)
		if err == nil {
			cache[oid] = phase
		}
		return phase, err
	}
	var commits []string
	for _, line := range strings.Split(strings.TrimSpace(out), "\n") {
		fields := strings.Fields(line)
		if len(fields) == 0 {
			continue
		}
		phase, err := phaseAt(fields[0])
		if err != nil {
			return nil, err
		}
		parent := ""
		if len(fields) > 1 {
			parent = fields[1]
		}
		previous, err := phaseAt(parent)
		if err != nil {
			return nil, err
		}
		if phase != previous {
			commits = append(commits, fields[0])
		}
	}
	return commits, nil
}
