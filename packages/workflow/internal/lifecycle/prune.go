package lifecycle

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow/model"
)

func GC(ctx context.Context, request model.GCRequest) (model.GCResult, error) {
	var result model.GCResult
	repo, err := gitstore.Inspect(request.Root)
	if err != nil {
		return result, err
	}
	root := repo.Root
	snapshot, err := project.Load(root)
	if err != nil {
		return result, err
	}
	const nativeDate = "20060102T150405.000000000"
	age := 720 * time.Hour
	if config, ok := snapshot.Config["gc"]; ok {
		fields, ok := config.(map[string]any)
		if !ok {
			return result, failure("C01", "fireweed.yaml", "gc must be a mapping")
		}
		if value, exists := fields["drill_max_age"]; exists {
			age, err = time.ParseDuration(str(value))
			if err != nil || age <= 0 {
				return result, failure("C01", "fireweed.yaml", "gc.drill_max_age must be a positive Go duration: %v", value)
			}
		}
	}
	now := request.Now
	if now.IsZero() {
		now = time.Now().UTC()
	}
	cutoff := now.Add(-age)
	archived, err := changePaths(root, "", archiveDir)
	if err != nil {
		return result, err
	}
	archivedSlugs := make(map[string]bool)
	for _, change := range archived {
		archivedSlugs[change.Slug] = true
	}
	all, err := refs(root, "refs/fireweed/")
	if err != nil {
		return result, err
	}
	for _, r := range all {
		if err := ctx.Err(); err != nil {
			return result, err
		}
		var expired bool
		switch {
		case strings.HasPrefix(r.name, claimsPrefix):
			data, err := gitstore.ReadBlob(root, r.oid)
			if err != nil {
				continue
			}
			var owner struct {
				Worktree string `json:"worktree"`
			}
			if json.Unmarshal(data, &owner) == nil && owner.Worktree != "" {
				working, err := filepath.EvalSymlinks(owner.Worktree)
				expired = os.IsNotExist(err) || err == nil && !registeredWorktree(working, repo.CommonDir)
			}
		case strings.HasPrefix(r.name, attemptsPrefix):
			slug, value, _ := strings.Cut(strings.TrimPrefix(r.name, attemptsPrefix), "/")
			n, err := strconv.Atoi(value)
			expired = err == nil && n > 0 && strconv.Itoa(n) == value && archivedSlugs[slug]
		case strings.HasPrefix(r.name, drillsPrefix):
			_, value, _ := strings.Cut(strings.TrimPrefix(r.name, drillsPrefix), "/")
			date, err := time.Parse(nativeDate, value)
			expired = err == nil && date.Before(cutoff)
		}
		if expired {
			if err := gitstore.CASRef(root, r.name, r.oid, ""); err == nil {
				result.Pruned = append(result.Pruned, r.name)
			}
		}
	}
	slices.Sort(result.Pruned)
	return result, nil
}

func registeredWorktree(worktree, commonDir string) bool {
	gitfile := filepath.Join(worktree, ".git")
	out, err := git(worktree, "rev-parse", "--resolve-git-dir", gitfile)
	if err != nil {
		return false
	}
	admin, err := filepath.EvalSymlinks(strings.TrimSuffix(out, "\n"))
	if err != nil {
		return false
	}
	if admin == commonDir {
		return gitfile == commonDir
	}
	if filepath.Dir(admin) != filepath.Join(commonDir, "worktrees") {
		return false
	}

	// Match the linked registration's back pointer to the actual worktree.
	// Legacy porcelain can contain literal newlines in worktree paths.
	backlink, err := os.ReadFile(filepath.Join(admin, "gitdir"))
	if err != nil {
		return false
	}
	registered := strings.TrimSuffix(string(backlink), "\n")
	if !filepath.IsAbs(registered) {
		registered = filepath.Join(admin, registered)
	}
	parent, err := filepath.EvalSymlinks(filepath.Dir(registered))
	return err == nil && parent == worktree && filepath.Base(registered) == ".git"
}
