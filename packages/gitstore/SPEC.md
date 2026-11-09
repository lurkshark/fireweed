---
depends_on:
    - contracts
id_prefix: GIT
interface:
    - api.go
    - model/**
name: gitstore
---

# gitstore

## Specs

### GIT-001
Inspect MUST return canonical absolute worktree/common-Git directories plus HEAD and branch, from a directory inside a primary or linked worktree. The branch MUST be empty for detached HEAD, and non-repositories MUST return an error.

### GIT-002
BaseRef MUST prefer a locally known remote/base ref, falling back to the local base branch and rejecting a missing base. An omitted remote MUST mean origin. An omitted base MUST use the local remote HEAD's branch when known, otherwise main.

### GIT-003
Changes MUST report repository-relative staged and unstaged tracked edits and untracked paths against the given base, once per reported Path. Renames reported by either tracked diff MUST retain both names.

### GIT-004
WriteBlob MUST store a byte payload in local Git objects and return its object ID.

### GIT-005
Run MUST launch the complete argv in Args, whose first element names the executable, without prepending any command. It MUST use the requested directory and supplied environment/stdin and return captured stdout, stderr and exit code, returning an error with that output on nonzero exit. Timeout or context cancellation MUST terminate the launched process group and return an error.

### GIT-006
ReadAt MUST return a revision's file bytes or an error for absent paths.

### GIT-007
Commit MUST commit the worktree contents of explicitly selected paths, including new files, with the given subject and return the new HEAD. It MUST leave unrelated working and index changes uncommitted.

### GIT-008
ReadBlob MUST return the bytes of a local Git blob object by its object ID.

### GIT-009
Ref MUST return the named local ref's object ID and existence.

### GIT-010
CASRef MUST create, update or delete a ref only when its exact current object matches the expected old value, returning an error without overwriting a newer value on mismatch.

## Decisions

[0010-manual-lifecycle-git] Change.Status uses the leading letter of Git's name-status code for tracked changes; untracked paths use `?`.
