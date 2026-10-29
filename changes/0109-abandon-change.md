---
bases: {}
module: workflow
---



## Motivation

Design Abandoning a change and the Lifecycle CLI require abandon <change> to
drop unapplied work using the change's commits, one deletion/revert commit,
local-claim ownership, retained attempts and free unapplied IDs. NEW still
rejects abandon as unimplemented. Provide only the needed workflow dependency,
then immediately wire app and prove actual NEW abandons a clean unapplied
change, retaining attempts and allowing ID reuse. Fifth-change review follows
that caller; no whole-module completion or unrelated history cleanup.

## Interface

Add thin public Abandon(root, change string) (string, error), returning the new
commit OID. All behavior stays under workflow/internal using existing Git and
claim helpers. No new public model or process primitive.

## Added

### new:abandon-effects
Abandon MUST undo the unapplied change's phase-commit effects identified by Commits and delete its pending change file in one commit, returning that commit's full object ID.

### new:abandon-clean-tree
Abandon MUST reject a dirty working tree with P01 before changing files, HEAD or refs.

### new:abandon-claim-ownership
Abandon MUST allow an unclaimed change or a claim owned by the calling worktree.
It MUST reject a foreign claim with P02 before changing files, HEAD or refs.
On success it MUST release the selected owned claim.

### new:abandon-applied-eligibility
Abandon MUST reject an applied archived change with P01.
After git revert returns that change to the pending queue, Abandon MUST allow dropping it under the same preconditions.

### new:abandon-attempts
Abandon MUST retain saved attempt refs unchanged.

### new:abandon-no-phase-assignment
The Abandon commit MUST not assign a new phase value.

### new:abandon-id-reuse
After successful Abandon, its assigned but unapplied spec IDs MUST be available for subsequent Start allocation.

### new:abandon-selectors
Abandon MUST accept the same unambiguous repository-relative path, filename and slug selectors as Commits, rejecting missing or ambiguous selectors with U02.

## Modified



## Removed



## Decisions



## Notes

Use public workflow/project APIs, explicit temporary Git identities and existing
lifecycle fixtures. Representative queued/unclaimed and clean owned stepped
changes cover effect reversal, one-commit/OID/file/phase results and ownership;
use real public Block to produce an attempt ref, then check retention and ID
reuse. Cover dirty/foreign/applied refusal without mutation and normal applied
revert-to-pending recovery. Keep aliases representative, no merge/history or
platform matrix. No private claim/attempt serialization or fake Git proxies.
Reuse Commits and exact captured-claim CAS semantics; saved refs are not GC'd.
Use ordinary bounded Git operations, handle normal errors without unrelated
file/ref changes or a generic recovery framework. No hooks/agents, other module
or protected edits, unrelated workflow reorganization, commits or phase commands
by workers. Preserve existing assertions. Interface/evals freeze before impl;
fixed real Sol/high workers and separate restricted audit, native full checks.
Own under5m/full under15m, no raised bounds. No numeric affected-spec cap: these
atomic items describe one coherent abandon behavior. Fourth application after
passing checkpoint; immediately return to app caller, then fifth review.
