---
audit_rounds: 0
bases: {}
module: workflow
phase: impl
run: 20261029T233903-c9fd8b6bf7425b01233b2e82
---



## Motivation

The design requires abandon <change> to undo unapplied work in one commit,
release owned claims, retain attempts and free IDs. Owner explicitly approved
native cancellation of0109 after its extra frozen Current assertion contradicted
FLOW-061. Re-author the same eight obligations from retained public contracts
and evals, correcting only that lookup expectation P02 to U02. Do not change
Current or reuse the discarded private implementation. Immediately return to
the app caller and actual NEW proof, then the fifth-application review.

## Interface

Add thin public Abandon(root, change string) (string, error), returning the new
commit OID. Behavior belongs under workflow/internal using existing dependencies;
no new public model or process primitive.

## Added

### FLOW-089
Abandon MUST undo the unapplied change's phase-commit effects identified by Commits and delete its pending change file in one commit, returning that commit's full object ID.

### FLOW-090
Abandon MUST reject a dirty working tree with P01 before changing files, HEAD or refs.

### FLOW-091
Abandon MUST allow an unclaimed change or a claim owned by the calling worktree.
It MUST reject a foreign claim with P02 before changing files, HEAD or refs.
On success it MUST release the selected owned claim.

### FLOW-092
Abandon MUST reject an applied archived change with P01.
After git revert returns that change to the pending queue, Abandon MUST allow dropping it under the same preconditions.

### FLOW-093
Abandon MUST retain saved attempt refs unchanged.

### FLOW-094
The Abandon commit MUST not assign a new phase value.

### FLOW-095
After successful Abandon, its assigned but unapplied spec IDs MUST be available for subsequent Start allocation.

### FLOW-096
Abandon MUST accept the same unambiguous repository-relative path, filename and slug selectors as Commits, rejecting missing or ambiguous selectors with U02.

## Modified



## Removed



## Decisions



## Notes

Owner reply: "Approve native recovery". The canceled0109 is not an application;
three applied changes remain since review. This recovery and immediate app
caller are explicitly authorized consecutive abandon work. Keep normal audit
bounds, full verification and fifth-change checkpoint; no phase-write exception.
Use the retained corrected PUBLIC eval file and declared contract only, creating
a fresh compiling stub. Fresh real Sol/high implementation must not read/reuse
abandoned private code, saved implementation or old worker logs; bootstrap
source/tests remain forbidden. Existing Current U02 behavior is retained.
Use public workflow/project APIs, temporary Git identities and real public Block
for attempt retention/ID reuse, representative ownership/dirty/applied/alias
cases. Preserve all eight obligations and every assertion except the explicitly
corrected extra lookup prefix. No private serialization, fake Git proxies,
matrices, per-spec convention instrumentation, other-module/protected edits,
new process primitives or unrelated workflow reorganization. Workers never
commit or advance phases. Interface/evals freeze before impl, separate restricted
audit and native full checks; own under5m/full under15m with unchanged bounds.

## Provenance

- evals 2026-10-29T23:42:51.29465Z agent=manual fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=402 failing=12 passing=390
- audit 2026-10-29T23:44:54.446633Z agent=manual fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0

## Audit

No findings
Checked IDs: FLOW-089, FLOW-090, FLOW-091, FLOW-092, FLOW-093, FLOW-094, FLOW-095, FLOW-096
