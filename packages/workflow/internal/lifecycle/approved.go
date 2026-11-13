package lifecycle

import (
	"context"
	"encoding/hex"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/workflow/model"
)

// Approved reads the held snapshot and queued path from Git bookkeeping.
func Approved(root string) (model.ApprovedSnapshot, bool, error) {
	var approved model.ApprovedSnapshot
	top, err := commitsGit(root, "rev-parse", "--show-toplevel")
	if err != nil {
		return approved, false, err
	}
	root, err = filepath.EvalSymlinks(strings.TrimSuffix(top, "\n"))
	if err != nil {
		return approved, false, err
	}
	prefix := ref("")
	refs, err := gitstore.Run(context.Background(), gitstore.Request{
		Args: []string{"git", "-C", root, "-c", "core.hooksPath=/dev/null", "for-each-ref",
			"--sort=refname", "--format=%(refname) %(objectname)", prefix},
		Timeout: 30 * time.Second,
		Env:     map[string]string{"GIT_NO_LAZY_FETCH": "1", "GIT_TERMINAL_PROMPT": "0", "GIT_OPTIONAL_LOCKS": "0"},
	})
	if err != nil {
		return approved, false, codedErrorf("P02", prefix, "list claims: %w: %s", err, strings.TrimSpace(refs.Stderr))
	}
	// Git can warn and omit a broken claim ref while still exiting successfully.
	if warning := strings.TrimSpace(refs.Stderr); warning != "" {
		return approved, false, codedErrorf("P02", prefix, "list claims: %s", warning)
	}
	var owned claim
	for _, line := range strings.Split(strings.TrimSuffix(refs.Stdout, "\n"), "\n") {
		if line == "" {
			continue
		}
		name, oid, _ := strings.Cut(line, " ")
		slug := strings.TrimPrefix(name, prefix)
		c, err := readClaim(root, slug, oid)
		if err != nil {
			return approved, false, err
		}
		if base, err := hex.DecodeString(c.PhaseBase); err != nil || len(base) != 20 ||
			!filepath.IsAbs(c.Worktree) || c.Change != slug || !commitChangeName.MatchString("0000-"+slug+".md") ||
			c.Snapshot.Root != c.Worktree || c.Snapshot.Config == nil || len(c.Snapshot.Modules) == 0 {
			return approved, false, codedErrorf("P02", slug, "malformed claim data")
		}
		if c.Worktree == root && owned.Change == "" {
			owned = c
		}
	}
	if owned.Change == "" {
		return approved, false, nil
	}

	// Existing claims record the phase base and slug. Find the pending path
	// in that commit even if the checkout no longer contains the change.
	paths, err := commitChangePaths(root, owned.PhaseBase+"^{commit}")
	if err != nil {
		return approved, false, codedErrorf("P02", owned.Change, "read claim phase base: %w", err)
	}
	pending := paths[:0]
	for _, candidate := range paths {
		if path.Dir(candidate) == "changes" {
			pending = append(pending, candidate)
		}
	}
	basename, err := selectCommitChange(pending, owned.Change)
	if err != nil {
		return approved, false, codedErrorf("P02", owned.Change, "resolve queued path at claim phase base: %w", err)
	}
	owned.Snapshot.Root = root
	return model.ApprovedSnapshot{Snapshot: owned.Snapshot, Slug: owned.Change, Path: "changes/" + basename}, true, nil
}
