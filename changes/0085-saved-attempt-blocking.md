---
bases: {}
module: workflow
---



## Motivation

The design's Retries and Blocking sections require exhausted work to be saved
on an attempt ref, the checkout restored, and a state-only block committed with
claim release. NEW next now retries verification, but exhaustion still leaves
the work held and dirty. Add only the lifecycle operation needed for the
runner to finish that outcome. The next caller change will wire it immediately
and close the slice with actual NEW next exhaustion; independent continuation,
from-attempt resume and other recovery mechanisms remain subsequent scenarios.

## Interface

Add Block(BlockRequest) (BlockResult, error) to the public workflow facade.
Retained model types: BlockRequest has Root, Change and Reason strings;
BlockResult has State, AttemptRef and AttemptOID. State uses the existing
workflow State; the ref and OID identify the saved attempt. Keep the facade
thin and all behavior disposable. No existing signatures change.

## Added

### new:block-ownership
Block MUST require the calling worktree's current claim with unchanged approved
metadata and phase. It MUST reject a changed HEAD rather than blocking a stale
phase base.

### new:save-step-attempt
Block MUST save the held step's tracked and untracked work as a commit at the
next unused positive refs/fireweed/attempts/<slug>/<n>, preserving prior attempt
refs. BlockResult MUST identify that ref and its commit OID.

### new:restore-step-checkout
After saving the attempt, Block MUST restore the step's tracked files, untracked
files and index to the phase base, removing the step's edits from the checkout.

### new:block-state-commit
Block MUST commit only the held change's assigned IDs, run, current phase and
requested blocked reason over the phase base. The block commit MUST preserve
the approved author and already-committed audit content and exclude step work.

### new:release-blocked-claim
On successful blocking, Block MUST release the current claim and return the
blocked State with no held ClaimOID. The resulting working tree MUST be clean.

## Modified



## Removed



## Decisions



## Rejected alternatives



## Notes

This is the dependency for the next concrete NEW next exhaustion scenario, not
a general recovery subsystem. Do not wire runner/app, add from-attempt resume,
agent-authored Questions, agent-commit undo, projected eligibility, new change
kinds, usage waiting, or cleanup unrelated to this behavior. Existing audit
blocking and public signatures stay intact. No bootstrap/source/history oracle.

Use public workflow APIs, ordinary temporary Git repositories, explicit fixture
identities and bounded hooks. One representative held evals step should include
tracked modification/deletion, staged and untracked work, assigned IDs/run and
an existing attempt ref so save/restore/state-only commit/release are observable.
A representative impl-phase case should preserve earlier committed eval/audit
content. Cover ownership/approved-state and stale-HEAD rejection without special
instrumentation or speculative failure injection. Reuse existing fixture helpers;
retain every old assertion. Own Verify below five minutes, full below fifteen.

Workers never commit in this checkout, invoke controller lifecycle operations,
change phase/counters/refs in this checkout, or edit protected docs/SPEC/other
modules. Only the controller owns governing commits and phase advancement.
Separate auditors receive isolated change/projected specs/evals/interfaces.
