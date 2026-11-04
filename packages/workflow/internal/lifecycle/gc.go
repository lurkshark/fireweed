package lifecycle

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow/model"
)

const gcDrillDate = "20060102T150405.000000000"

type gcRef struct{ name, oid string }

func gcGit(ctx context.Context, root string, args ...string) (string, error) {
	r, err := gitstore.Run(ctx, gitstore.Request{
		Args:    append([]string{"git", "-C", root, "-c", "core.hooksPath=/dev/null"}, args...),
		Timeout: 30 * time.Second,
		Env:     map[string]string{"GIT_NO_LAZY_FETCH": "1", "GIT_TERMINAL_PROMPT": "0", "GIT_OPTIONAL_LOCKS": "0"},
	})
	if err != nil {
		return "", fmt.Errorf("git %s: %w: %s", strings.Join(args, " "), err, r.Stderr)
	}
	return r.Stdout, nil
}

func gcMaxAge(config map[string]any) (time.Duration, error) {
	raw, exists := config["gc"]
	if !exists {
		return 720 * time.Hour, nil
	}
	settings, ok := raw.(map[string]any)
	if !ok {
		return 0, fmt.Errorf("C01: gc must be a mapping")
	}
	raw, exists = settings["drill_max_age"]
	if !exists {
		return 720 * time.Hour, nil
	}
	value, ok := raw.(string)
	if !ok {
		return 0, fmt.Errorf("C01: gc.drill_max_age must be a positive Go duration")
	}
	age, err := time.ParseDuration(value)
	if err != nil || age <= 0 {
		return 0, fmt.Errorf("C01: gc.drill_max_age must be a positive Go duration: %q", value)
	}
	return age, nil
}

func gcRefs(ctx context.Context, root string) ([]gcRef, error) {
	out, err := gcGit(ctx, root, "for-each-ref", "--format=%(refname) %(objectname)",
		"refs/fireweed/claims", "refs/fireweed/attempts", "refs/fireweed/drills")
	if err != nil {
		return nil, err
	}
	var refs []gcRef
	for _, line := range strings.Split(strings.TrimSpace(out), "\n") {
		fields := strings.Fields(line)
		if len(fields) == 2 {
			refs = append(refs, gcRef{fields[0], fields[1]})
		}
	}
	return refs, nil
}

func gcRegistered(ctx context.Context, root string) (map[string]bool, error) {
	out, err := gcGit(ctx, root, "worktree", "list", "--porcelain")
	if err != nil {
		return nil, err
	}
	registered := map[string]bool{}
	for _, line := range strings.Split(out, "\n") {
		if strings.HasPrefix(line, "worktree ") {
			registered[strings.TrimPrefix(line, "worktree ")] = true
		}
	}
	return registered, nil
}

func gcArchivedSlugs(root string) (map[string]bool, error) {
	archived := map[string]bool{}
	for _, directory := range []string{"changes/archive", "changes"} {
		entries, err := os.ReadDir(filepath.Join(root, directory))
		if os.IsNotExist(err) {
			continue
		}
		if err != nil {
			return nil, err
		}
		for _, entry := range entries {
			name := entry.Name()
			if entry.IsDir() || !commitChangeName.MatchString(name) {
				continue
			}
			slug := strings.TrimSuffix(name[strings.IndexByte(name, '-')+1:], ".md")
			if directory == "changes/archive" {
				archived[slug] = true
			} else {
				delete(archived, slug)
			}
		}
	}
	return archived, nil
}

func gcClaimStale(root string, entry gcRef, registered map[string]bool, common string) bool {
	data, err := gitstore.ReadBlob(root, entry.oid)
	if err != nil {
		return false
	}
	var claim struct {
		Worktree string `json:"worktree"`
	}
	if json.Unmarshal(data, &claim) != nil || !filepath.IsAbs(claim.Worktree) {
		return false
	}
	worktree, err := filepath.EvalSymlinks(claim.Worktree)
	if os.IsNotExist(err) {
		return true
	}
	if err != nil {
		return false
	}
	if !registered[worktree] {
		return true
	}
	owner, err := gitstore.Inspect(worktree)
	return err != nil || owner.Root != worktree || owner.CommonDir != common
}

func gcArchivedAttempt(name string, archived map[string]bool) bool {
	remainder, ok := strings.CutPrefix(name, "refs/fireweed/attempts/")
	if !ok {
		return false
	}
	parts := strings.Split(remainder, "/")
	if len(parts) != 2 || !archived[parts[0]] {
		return false
	}
	n, err := strconv.ParseUint(parts[1], 10, 64)
	return err == nil && n > 0 && strconv.FormatUint(n, 10) == parts[1]
}

func gcExpiredDrill(name string, cutoff time.Time) bool {
	remainder, ok := strings.CutPrefix(name, "refs/fireweed/drills/")
	if !ok {
		return false
	}
	parts := strings.Split(remainder, "/")
	if len(parts) != 2 || parts[0] == "" {
		return false
	}
	date, err := time.ParseInLocation(gcDrillDate, parts[1], time.UTC)
	return err == nil && date.Format(gcDrillDate) == parts[1] && date.Before(cutoff)
}

// GC removes only recognized, expired local bookkeeping refs.
func GC(ctx context.Context, request model.GCRequest) (model.GCResult, error) {
	var result model.GCResult
	repo, err := gitstore.Inspect(request.Root)
	if err != nil {
		return result, err
	}
	snapshot, err := project.Load(repo.Root)
	if err != nil {
		return result, err
	}
	age, err := gcMaxAge(snapshot.Config)
	if err != nil {
		return result, err
	}
	now := request.Now
	if now.IsZero() {
		now = time.Now().UTC()
	}
	cutoff := now.UTC().Add(-age)
	refs, err := gcRefs(ctx, repo.Root)
	if err != nil {
		return result, err
	}
	registered, err := gcRegistered(ctx, repo.Root)
	if err != nil {
		return result, err
	}
	archived, err := gcArchivedSlugs(repo.Root)
	if err != nil {
		return result, err
	}
	for _, entry := range refs {
		if err := ctx.Err(); err != nil {
			return result, err
		}
		var prune bool
		switch {
		case strings.HasPrefix(entry.name, "refs/fireweed/claims/"):
			prune = gcClaimStale(repo.Root, entry, registered, repo.CommonDir)
		case strings.HasPrefix(entry.name, "refs/fireweed/attempts/"):
			prune = gcArchivedAttempt(entry.name, archived)
		case strings.HasPrefix(entry.name, "refs/fireweed/drills/"):
			prune = gcExpiredDrill(entry.name, cutoff)
		}
		if !prune {
			continue
		}
		if err := gitstore.CASRef(repo.Root, entry.name, entry.oid, ""); err != nil {
			sort.Strings(result.Pruned)
			return result, fmt.Errorf("delete %s: %w", entry.name, err)
		}
		result.Pruned = append(result.Pruned, entry.name)
	}
	sort.Strings(result.Pruned)
	return result, nil
}
