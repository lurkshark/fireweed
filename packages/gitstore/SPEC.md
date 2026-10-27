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
Changes MUST report repository-relative staged/unstaged tracked edits and untracked paths against the given base, retaining both names for staged renames. ReadAt MUST return a revision's file bytes or an error for absent paths. Commit MUST commit only explicitly selected paths with the given subject, return the new HEAD, and leave unrelated working/index changes uncommitted.

### GIT-004
WriteBlob and ReadBlob MUST round-trip byte payloads through local Git objects. Ref MUST return the named local ref's object ID and existence. CASRef MUST create, update or delete a ref only when its exact current object matches the expected old value, returning an error without overwriting a newer value on mismatch.

### GIT-005
Run MUST launch argv in the requested directory with the supplied environment/stdin and return captured stdout, stderr and exit code. A nonzero exit MUST return an error while preserving output. Timeout or context cancellation MUST terminate the launched process group and return an error; processes escaping that group are outside the supported contract.

## Decisions

[0010-manual-lifecycle-git] Change.Status uses the leading letter of Git's name-status code for tracked changes; untracked paths use `?`.
