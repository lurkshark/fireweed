---
module: gitstore
phase: applied
run: 28ec83301f9665518f908c01
---

## Motivation

The product needs safe local Git discovery and worktree inspection before static validation and lifecycle state can depend on them. Linked worktrees must share the same common Git directory while preserving their own checkout identity.

## Interface

Expose retained Repository {Root, GitDir, CommonDir string} and Edit {Status, Path, OldPath string}, plus Open(start string) (Repository,error), BaseRef(root,base,remote string) (string,error), Clean(root string) (bool,error), Edits(root,base string) ([]Edit,error), SafePath(root,relative string) (string,error). Internal implementations return primitive fields or internal DTOs converted by thin facade wrappers. Paths in edits are repository-relative forward slashes; Repository directories are absolute canonical paths. No Git mutation API is added in this change.

## Added

### GIT-001
Repository discovery MUST work from a checkout or its nested directory and return the checkout root, worktree Git directory and common Git directory as canonical absolute paths. Linked worktrees MUST have distinct roots and Git directories but share their common Git directory. Non-repositories MUST return an error. Inspection MUST NOT create branches, change refs, fetch or push.

### GIT-002
Base resolution MUST use the configured remote (default origin) and configured branch; when the branch is omitted it MUST use that remote's locally known symbolic default, falling back to main. It MUST prefer an existing refs/remotes/<remote>/<base> over refs/heads/<base>, otherwise return the local ref name even when absent. It MUST NOT contact a remote. Branch names containing slashes MUST be retained.

### GIT-003
Clean-tree inspection MUST detect tracked staged or unstaged edits, deletions and untracked files, while ignoring ignored files. Worktree diffs against an explicit local commit MUST include staged and unstaged additions/modifications/deletions, untracked files and renames with old and new paths. Spaces, tabs and newlines in filenames MUST be preserved. Results MUST be deterministic and invalid Git roots or revision arguments MUST return errors rather than empty success.

### GIT-004
Safe path resolution MUST accept only nonempty repository-relative forward-slash paths without absolute roots, drive prefixes, dot/traversal components, repeated/trailing slashes, backslashes or NUL. It MUST return an absolute path inside the canonical root, allow nonexistent descendants and internal symlinks, and reject existing or ancestor symlinks escaping the root, including sibling paths sharing its string prefix. It MUST NOT create or modify files.

## Modified

## Removed

## Decisions

Git operations use argument vectors, not shell interpolation. Edit statuses are A/M/D/R (or another first status character reported by Git); rename Path is the new path and OldPath the old path. Untracked files are A. The diff is relative to the requested commit, not merely the index. Identical entries are deduplicated and results sorted by Path then OldPath then Status. SafePath rejects dangling symlinks and roots that do not exist. Root canonicalization follows symlinks. Base resolution reads only local refs; names are validated with git check-ref-format to prevent revision/options injection. Invalid base/remote names return errors. No temporary branches or network commands are used by inspection.

## Rejected alternatives

String-prefix path containment would allow sibling-directory escapes. Line-based Git output parsing would corrupt legal filenames containing newlines. Index-only diffs would miss the fail-first worker's unstaged changes.

## Provenance

- evals | 2026-10-22T07:00:43Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 4 | failing: 4 | passing: 0
- audit | 2026-10-22T07:03:32Z | agent: manual | fireweed: 0.4.0 | playbook: a9c982d24b476621316049b527e6c7bbc319f06b12268af150fa7d8c6df92230 | tests: 0 | failing: 0 | passing: 0
- evals | 2026-10-22T07:03:56Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 8 | failing: 8 | passing: 0
- audit | 2026-10-22T07:05:31Z | agent: manual | fireweed: 0.4.0 | playbook: a9c982d24b476621316049b527e6c7bbc319f06b12268af150fa7d8c6df92230 | tests: 0 | failing: 0 | passing: 0
- impl | 2026-10-22T07:06:48Z | agent: manual | fireweed: 0.4.0 | playbook: 653912b194316c62181f62d0a69dc213f7668481ea8b1b9926ef889dd48a9e08 | tests: 8 | failing: 0 | passing: 8

## Audit

No findings. Independent auditor /root/audit_local_repository checked GIT-001, GIT-002, GIT-003 and GIT-004 against concrete assertions and Decisions. All eight prior findings are resolved in the committed evals; findings and retry remain in phase history.
