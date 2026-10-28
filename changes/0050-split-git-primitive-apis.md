---
already_satisfied:
    - GIT-003
    - GIT-004
    - new:read-revision-file
    - new:selected-path-commit
    - new:read-blob
    - new:read-ref
    - new:compare-and-swap-ref
bases:
    GIT-003: a1d9c098
    GIT-004: 6a41ac8b
module: gitstore
---



## Motivation

Batch the previously identified GIT-003/004 API bundles under SYSTEM's
30-item spec-only cap. Separate revision reads, selected commits, blob reads,
ref reads and compare-and-swap from working-tree diffs and blob writes.
Preserve the design's local Git bookkeeping and claim safety obligations,
all assertions and public interfaces. Seven affected items; behavior already
exists and implementation stays unchanged.

## Interface



## Added

### new:read-revision-file
ReadAt MUST return a revision's file bytes or an error for absent paths.

### new:selected-path-commit
Commit MUST commit only explicitly selected paths with the given subject, return the new HEAD, and leave unrelated working and index changes uncommitted.

### new:read-blob
ReadBlob MUST return the bytes of a local Git blob object by its object ID.

### new:read-ref
Ref MUST return the named local ref's object ID and existence.

### new:compare-and-swap-ref
CASRef MUST create, update or delete a ref only when its exact current object matches the expected old value, returning an error without overwriting a newer value on mismatch.

## Modified

### GIT-003
Changes MUST report repository-relative staged and unstaged tracked edits and untracked paths against the given base, retaining both names for staged renames.

### GIT-004
WriteBlob MUST store a byte payload in local Git objects and return its object ID.

## Removed



## Decisions



## Notes

Retarget existing gitstore_test.go evals only. ChangesReadAtAndSelectedCommit
keeps GIT-003 and gains ReadAt/Commit IDs; BlobsAndRefCAS keeps GIT-004 and
gains ReadBlob/Ref/CAS IDs, including the existing stale-update subcase.
Preserve all byte round-trip, selected-path, subject, unrelated index/worktree,
absent/create/update/delete and newer-ref assertions. Keep all eight reported
Git cases. No new Test functions, t.Run cases, fixture instances, helpers,
implementation or interface edits. Do not change literal fixture paths, refs,
object IDs or Git commands. Keep GIT-001/002/005 and the existing behavioral
Decision unchanged. Implementation may edit only this change file.
