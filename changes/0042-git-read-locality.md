---
fixes:
    - GIT-002
module: gitstore
---

## Motivation

The unchanged design requires local Git inspection without fetches. Independent public marker-only discovery on installed Git2.30.1 confirms BaseRef attempts transport when an existing preferred tracking ref names an unavailable commit, then incorrectly returns the local fallback. Edits attempts transport before rejecting an unavailable selected commit or root tree. The same discovery proves a simple modification with an unavailable old blob and a comparison with an irrelevant missing parent can succeed without transport. Preserve those available results rather than requiring a complete object graph.

The environment hint GIT_NO_LAZY_FETCH=1 is ineffective on this installed supported Git. Direct experiments also show remote.promisor=false overrides do not suppress its fetch attempts. A useful error after an attempted fetch is still a violation. This change fixes the existing base-selection contract and specifies the local-only behavior of edit inspection without changing its API.

## Interface

No public API/type changes. Keep Open, BaseRef, Clean, Edits, SafePath, Run and Capture signatures and all existing contracts. Add permanent public-API regressions for GIT-002 and new:local-edits, plus independent real-Git controls proving marker-only transport and Trace2 fetch observations in allowed and denied protocol modes. Exercise ordinary/linked SHA-1/SHA-256 repositories supported by installed Git, required and irrelevant missing objects, valid single commit/tag/ancestor inputs, revision-set rejection and available built-in diff/index/worktree configuration and rename behavior. Compare successful available results against independently executed real Git rather than an implementation-specific command allowlist. Preserve all existing evals and their genuine success controls.

## Added

### new:local-edits
Edits MUST inspect the requested local baseline without attempting fetch, push, branch creation or network commands, including partially populated repositories on supported Git versions. It MUST return a useful non-nil error and an empty edit result when objects required to resolve the selected commit or determine its edits are unavailable; it MUST NOT return partial successful edits. Missing unrelated ancestors, trees or blob bodies that are not needed for the requested comparison MUST NOT prevent otherwise available results. Available ordinary and linked SHA-1/SHA-256 repositories MUST retain the existing GIT-003 staged, unstaged, untracked, filename and rename semantics, including built-in Git configuration affecting index/worktree comparisons. Full commit IDs and valid single commit/tag/ancestor expressions MUST remain usable; revision sets MUST be rejected. Successful and rejected calls MUST remain noninteractive, use literal arguments, preserve HEAD, refs, original index bytes, working/ignored files and parent cwd/environment, and clean owned temporary storage. No bootstrap or authenticated agent MAY be required or invoked.

## Modified

## Removed

## Decisions

BaseRef chooses an existing exact tracking ref by local ref identity rather than requiring that ref's target commit body. An unavailable target is not permission to fetch or silently substitute a different base. Preserve configured/default remote/branch selection, slash-containing names, symbolic remote defaults and absent-tracking fallback under GIT-002. Existing edit ordering, rename source/destination and filename-byte contracts remain authoritative.

Availability is scoped to the requested operation. Do not validate every reachable commit/blob or decode unrelated content as a precondition. Prove the unavailable-object transport observer with real local Git controls, including a denied protocol where a fetch child may still be attempted. Clear controls before API calls and snapshot all state after fixture construction. Native Git output on available fixtures supplies the comparison result; a reference command that fails is not a positive control.

A temporary minimal Git metadata view avoided transport in an exploratory four-case experiment, but does not establish a production design or configuration compatibility. Any chosen implementation must preserve required built-in comparison configuration, linked-worktree/index/object-format and rename behavior, retain useful errors, and contain/clean only its own temporary resources. Neither original repository configuration nor user environment may be rewritten to suppress the bug. Evaluate implementation only after complete fail-first and an independent restricted audit.

The user's full rebuild goal authorizes local authored execution and commits with the controller. Manual start's normal P04 warning does not assert a human merge or waive approval policy. Keep the pinned bootstrap controller separate from the NEW subject; no pushes or publication. Subsequent workflow Verify composition must also fix its independently confirmed approval/exact-HEAD precondition fetches through genuine regressions, not through this module's worker edits.

## Rejected alternatives

An environment hint or failed network request after the fact does not prevent fetch. Requiring all repository objects would reject the demonstrated valid comparisons. Returning a local branch after a broken preferred tracking target changes the configured base. Weakening frozen dependent/native evals, losing built-in comparison behavior, editing another module or calling the bootstrap to implement the product would hide rather than repair the problem.
