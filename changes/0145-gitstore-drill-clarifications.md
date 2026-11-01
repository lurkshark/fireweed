---
already_satisfied:
    - GIT-003
    - GIT-007
bases:
    GIT-003: d38fe764
    GIT-007: b5f091aa
module: gitstore
---

## Motivation

The independent gitstore reconstruction reported two ambiguities, and the owner
approved its choices: report each changed path once while preserving a rename
reported by either tracked diff, and commit selected worktree bytes when the
index differs. The design's Regeneration section requires owner-reviewed gaps
to become agreed changes. Its dependent evals also demonstrated that Commit
must accept newly created selected files, as lifecycle commits create eval and
archive paths. This change makes those existing local Git behaviors explicit.

Focused public probes of current main already satisfy both interpretations,
including an unstaged rename exposed with intent-to-add, selected new files,
and unrelated staged content remaining uncommitted. Both IDs legitimately use
already_satisfied. Preserve all existing assertions and implementation; do not
manufacture failures by regressing correct behavior.

## Interface

No public signature, type, module metadata or configuration changes.

## Added

## Modified

### GIT-003
Changes MUST report repository-relative staged and unstaged tracked edits and untracked paths against the given base, once per reported Path. Renames reported by either tracked diff MUST retain both names.

### GIT-007
Commit MUST commit the worktree contents of explicitly selected paths, including new files, with the given subject and return the new HEAD. It MUST leave unrelated working and index changes uncommitted.

## Removed

## Decisions

## Notes

The owner approved the two reported interpretations after the interrupted drill
checkpoint. Use this approved prose and retained public contracts/evals as
authority, not old rebuilt code, source history, saved attempts or bootstrap.
No drill implementation is imported, and no real drill runs in evals.

Preserve every existing gitstore testcase, assertion and ID. Add focused public
assertions using existing temporary Git fixtures: one reported Path for a file
with staged and unstaged modifications; old/new rename names when a staged
rename also has an unstaged edit; and a rename exposed by the unstaged tracked
diff (ordinary intent-to-add suffices). Assert the count as well as rename data
so a map cannot conceal duplicate entries. Do not choose an unspecified status
precedence or build a speculative Git-state matrix.

For Commit, select a tracked file whose index bytes differ from its worktree
and a new file. Assert selected worktree bytes and exact committed paths, the
subject and returned HEAD, and preservation of unrelated staged and working
content. Reuse existing fixtures/assertions rather than duplicate them. All
these cases should pass current main; implementation and retained interfaces
remain unchanged. Keep normal restricted independent audit, immutable phase
zones, approval and configured bounds. Workers read context before each step,
never commit, and leave phase verification/commits to the controller.

This correction does not resolve the interrupted drill's missing native report
or reset its attempts. Another real invocation requires an explicit recovery
decision; assets remains pending until gitstore proof and evidence close.

## Audit

## Provenance
