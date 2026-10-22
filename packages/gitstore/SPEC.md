---
depends_on:
    - contracts
id_prefix: GIT
name: gitstore
---

# gitstore

## Specs

### GIT-001
Repository discovery MUST work from a checkout or its nested directory and return the checkout root, worktree Git directory and common Git directory as canonical absolute paths. Linked worktrees MUST have distinct roots and Git directories but share their common Git directory. Non-repositories MUST return an error. Inspection MUST NOT create branches, change refs, fetch or push.

### GIT-002
Base resolution MUST use the configured remote (default origin) and configured branch; when the branch is omitted it MUST use that remote's locally known symbolic default, falling back to main. It MUST prefer an existing refs/remotes/<remote>/<base> over refs/heads/<base>, otherwise return the local ref name even when absent. It MUST NOT contact a remote. Branch names containing slashes MUST be retained.

### GIT-003
Clean-tree inspection MUST detect tracked staged or unstaged edits, deletions and untracked files, while ignoring ignored files. Worktree diffs against an explicit local commit MUST include staged and unstaged additions/modifications/deletions, untracked files and renames with old and new paths. Spaces, tabs and newlines in filenames MUST be preserved. Results MUST be deterministic and invalid Git roots or revision arguments MUST return errors rather than empty success.

### GIT-004
Safe path resolution MUST accept only nonempty repository-relative forward-slash paths without absolute roots, drive prefixes, dot/traversal components, repeated/trailing slashes, backslashes or NUL. It MUST return an absolute path inside the canonical root, allow nonexistent descendants and internal symlinks, and reject existing or ancestor symlinks escaping the root, including sibling paths sharing its string prefix. It MUST NOT create or modify files.

## Decisions

[local-repository]
Git operations use argument vectors, not shell interpolation. Edit statuses are A/M/D/R (or another first status character reported by Git); rename Path is the new path and OldPath the old path. Untracked files are A. The diff is relative to the requested commit, not merely the index. Identical entries are deduplicated and results sorted by Path then OldPath then Status. SafePath rejects dangling symlinks and roots that do not exist. Root canonicalization follows symlinks. Base resolution reads only local refs; names are validated with git check-ref-format to prevent revision/options injection. Invalid base/remote names return errors. No temporary branches or network commands are used by inspection.
