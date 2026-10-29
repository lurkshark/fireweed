---
audit_pending:
    - FLOW-074
    - FLOW-075
audit_rounds: 1
bases: {}
module: workflow
phase: impl
run: 20261029T085817-dd0bc520b24864b5d616184a
---



## Motivation

The design's Resuming section requires start --from-attempt <n> to load saved
work while keeping the existing phase and run. The actual NEW run proof now
leaves a clean blocked Alpha with a real saved impl attempt, but no public
lifecycle operation loads it. Add only that workflow dependency, then promptly
wire the app option and close with actual NEW start, correction and archive in
the retained Go fixture. Do not expand the recovery library before that caller.

## Interface

Add StartFromAttempt(root, change string, attempt int) (State, error) to the thin
workflow facade. Keep Start and every existing signature/model unchanged.
All loading behavior belongs under disposable lifecycle code. The caller must
first remove blocked and commit the author unblock, as ordinary Start requires.

## Added

### FLOW-073
StartFromAttempt MUST resume a committed phase with the same clean-worktree,
eligibility and claim-ownership preconditions as Start. It MUST preserve the
existing run, phase and assigned IDs and leave HEAD unchanged.

### FLOW-074
StartFromAttempt MUST restore the saved step's modified, deleted and newly
added files from refs/fireweed/attempts/<canonical-slug>/<attempt> into the
working tree without reverting later independent changes.

### FLOW-075
While restoring saved worker-editable body sections for the resumed phase,
StartFromAttempt MUST preserve the current change's approved author fields,
runtime state and controller-written provenance.

### FLOW-076
StartFromAttempt MUST reject a non-positive attempt number or missing saved
attempt ref with P01 before claiming or writing the change.

## Modified



## Removed



## Decisions



## Rejected alternatives



## Notes

Use current public Git/process/document primitives; no gitstore, app or runner
edits. Preserve ordinary Start approval/P07, projected snapshot, ownership and
claim behavior. This operation resumes existing committed phases; it does not
automatically unblock, create a phase commit, advance phases or allocate a new
run/ID. Return a held State appropriate for normal Verify/Done and Context.

Load saved step deltas, not an old whole repository tree. Preserve unrelated
later commits and current approved author/runtime/provenance content; restore
saved worker-owned Decisions/Questions and other permitted phase body work.
Audit body work belongs only to an audit resume; impl must retain the committed
audit. A failed restore must report its actual cause; do not silently discard
work or hard-reset later commits. Do not add a conflict/platform/failure matrix
or invent old-attempt migration requirements beyond this supported scenario.

Use representative public-API fixtures with real Block saved refs, then a
committed author unblock. Cover retained run/phase/assigned IDs/HEAD and owned
claim; tracked edits/deletion/new files; saved worker body work with current
author/provenance retained; later independent work preserved; ordinary dirty,
waiting/blocked or competing-claim rejection; and non-positive/missing ref
preconditions without mutation. Reuse existing fixture hooks and assertions,
not real agents, observers or instrumentation. Own Verify below five minutes,
full below fifteen. Do not weaken any existing assertion or alter other modules.

Only the facade addition, new compiling stub and public evals in evals; fresh
real loading code follows isolated assertion audit. Workers never commit or
run governing workflow commands, edit protected docs/SPEC/other modules, alter
claims/counters/runtime, or read bootstrap source/tests, deleted implementation,
saved rebuild implementations or rebuild-v1. No drill or phase-write exception.
After application immediately return to the app flag and actual NEW proof.

## Audit

No findings
Checked IDs: FLOW-074, FLOW-075

## Provenance

- evals 2026-10-29T09:10:08.724224Z agent=manual fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=295 failing=8 passing=287
- audit 2026-10-29T09:14:20.974825Z agent=manual fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0
- evals 2026-10-29T09:18:48.448129Z agent=manual fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=297 failing=10 passing=287
- audit 2026-10-29T09:22:57.290214Z agent=manual fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0
