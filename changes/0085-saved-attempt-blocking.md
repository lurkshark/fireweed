---
audit_pending:
    - FLOW-066
audit_rounds: 1
bases: {}
module: workflow
phase: audit
run: 20261029T043714-a528702cc3674dd30fe43e62
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

### FLOW-063
Block MUST require the calling worktree's current claim with unchanged approved
metadata and phase. It MUST reject a changed HEAD rather than blocking a stale
phase base.

### FLOW-064
Block MUST save the held step's tracked and untracked work as a commit at the
next unused positive refs/fireweed/attempts/<slug>/<n>, preserving prior attempt
refs. BlockResult MUST identify that ref and its commit OID.

### FLOW-065
After saving the attempt, Block MUST restore the step's tracked files, untracked
files and index to the phase base, removing the step's edits from the checkout.

### FLOW-066
Block MUST commit only the held change's assigned IDs, run, current phase and
requested blocked reason over the phase base. The block commit MUST preserve
the approved author and already-committed audit content and exclude step work.

### FLOW-067
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

## Provenance

- evals 2026-10-29T05:08:52.313761Z agent=manual fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=268 failing=2 passing=266
- audit 2026-10-29T05:27:25.032085Z agent=manual fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0
- evals 2026-10-29T05:32:20.734796Z agent=manual fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=268 failing=2 passing=266

## Audit

- FLOW-066: "Block MUST commit only the held change's assigned IDs" — missing assertion that the other change fields remain unchanged.
