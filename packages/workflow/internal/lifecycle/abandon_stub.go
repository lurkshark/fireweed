package lifecycle

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/lurkshark/fireweed/packages/gitstore"
)

// Abandon removes a pending change and reverses its committed phase work.
func Abandon(root, change string) (string, error) {
	repo, err := gitstore.Inspect(root)
	if err != nil {
		return "", err
	}
	root = repo.Root
	dirty, err := gitstore.Changes(root, repo.Head)
	if err != nil {
		return "", err
	}
	if len(dirty) != 0 {
		return "", errors.New("P01: abandon requires clean worktree")
	}
	paths, err := commitChangePaths(root, repo.Head)
	if err != nil {
		return "", err
	}
	basename, err := selectCommitChange(paths, change)
	if err != nil {
		return "", err
	}
	pending := "changes/" + basename
	archived := "changes/archive/" + basename
	foundPending := false
	for _, candidate := range paths {
		if candidate == pending {
			foundPending = true
		}
		if candidate == archived {
			return "", errors.New("P01: applied archived change cannot be abandoned")
		}
	}
	if !foundPending {
		return "", fmt.Errorf("U02: change %q not found", change)
	}
	slug := strings.TrimSuffix(basename[strings.IndexByte(basename, '-')+1:], ".md")
	claimOID, claimed, err := gitstore.Ref(root, ref(slug))
	if err != nil {
		return "", fmt.Errorf("P02: read claim: %w", err)
	}
	if claimed {
		owner, _, err := existing(root, slug)
		if err != nil {
			return "", err
		}
		if owner.Worktree != root || owner.Change != slug {
			return "", errors.New("P02: claim belongs to another worktree")
		}
	}
	commits, err := Commits(root, pending)
	if err != nil {
		return "", err
	}
	patches := make([]string, 0, len(commits))
	for _, oid := range commits {
		line, err := commitsGit(root, "rev-list", "--parents", "-n", "1", oid)
		if err != nil {
			return "", err
		}
		fields := strings.Fields(line)
		if len(fields) < 2 {
			return "", fmt.Errorf("P01: phase commit %s has no parent", oid)
		}
		parent := fields[1]
		phase, err := commitPhaseAt(root, oid, basename)
		if err != nil {
			return "", err
		}
		previous, err := commitPhaseAt(root, parent, basename)
		if err != nil {
			return "", err
		}
		if phase == "" || phase == "applied" || previous == "applied" {
			continue
		}
		patch, err := commitsGit(root, "diff", "--binary", "--no-ext-diff", parent, oid,
			"--", ".", ":(exclude)"+pending, ":(exclude)"+archived)
		if err != nil {
			return "", err
		}
		if patch != "" {
			patches = append(patches, patch)
		}
	}
	changed := false
	rollback := func(cause error) (string, error) {
		if changed {
			_, _ = blockGit(root, nil, "reset", "--hard", repo.Head)
			_, _ = blockGit(root, nil, "clean", "-fd")
		}
		return "", cause
	}
	for _, patch := range patches {
		changed = true
		if _, err := blockGit(root, []byte(patch), "apply", "--reverse", "--3way", "--index"); err != nil {
			return rollback(err)
		}
	}
	changed = true
	if err := os.Remove(file(root, pending)); err != nil {
		return rollback(err)
	}
	if _, err := blockGit(root, nil, "add", "-A"); err != nil {
		return rollback(err)
	}
	if _, err := blockGit(root, nil, "commit", "-m", "chore: abandon "+path.Base(pending)); err != nil {
		return rollback(err)
	}
	oid, err := blockGit(root, nil, "rev-parse", "--verify", "HEAD")
	if err != nil {
		return "", err
	}
	if claimed {
		if err := gitstore.CASRef(root, ref(slug), claimOID, ""); err != nil {
			return "", fmt.Errorf("P02: release claim: %w", err)
		}
	}
	return oid, nil
}
