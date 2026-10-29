---
bases:
    FLOW-011: b6ee6f47
module: workflow
---



## Motivation

The design's Blocking section says later changes wait for the blocked module,
its dependents, and dependency interface changes, while independent work runs.
NEW next now saves and blocks exhausted work, but lifecycle selection still
uses the first merely unblocked change. Supply canonical projected eligibility
for manual Start and its runner caller, then immediately return to Next/Run
and close with actual NEW run applying independent work while dependent work
waits. This is that runnable continuation dependency, not a status-view or
whole recovery-library implementation.

## Interface

Add Eligible(root string) ([]project.Change, error) to the thin public workflow
facade, using the existing public project.Change type. No retained model or
existing signatures change. Start uses the same selection; all behavior stays
in disposable lifecycle code. No app/runner wiring in this dependency change.

## Added

### new:eligible-queue-order
Eligible MUST return unblocked queued changes not held back by earlier blocked
changes, in queue order. Empty or entirely waiting queues MUST return an empty
result without an error.

### new:same-module-wait
Eligible MUST omit a later change targeting the same module as an earlier
blocked change.

### new:dependent-wait
Eligible MUST omit a later change whose module depends directly or transitively
on an earlier blocked change's module.

### new:interface-dependency-wait
When a later change has a nonempty Interface section, Eligible MUST omit it if
an earlier blocked change's module depends directly or transitively on its
module. With an empty Interface, that dependency change MUST remain eligible
unless another blocking rule applies.

### new:projected-blocking-graph
For each candidate, Eligible MUST apply the blocking rules to the graph
projected through that candidate, including preceding queued set edits and its
own set edits.

## Modified

### FLOW-011
Start MUST reject a requested change unless it is the first change returned by
Eligible. With no requested target, Start MUST start that first eligible
change. When Eligible returns no change, Start MUST reject with P01.

## Removed



## Decisions



## Rejected alternatives



## Notes

Keep this coherent selection behavior within workflow. The owner removed numeric
spec-count caps; each rule remains atomic and normal cost/checkpoint bounds
apply. Preserve existing Start ownership, clean-tree, approval/P07, snapshot,
late-ID/run, resume, zone and lifecycle behavior. Do not reorganize stub.go;
change only the selection required by this demonstrated feature gap.

Use representative public-API temporary repositories and bounded hooks. Cover
same-module wait, direct/transitive downstream wait, upstream Interface blank
versus nonempty, independent order, empty/all-waiting results, and candidate
projection including preceding/own set edits. Include manual Start skipping a
waiting unblocked change and rejecting a non-head/ineligible target. Validate
fixtures and retain old assertions; no failure matrix, special instrumentation
or observer/proxy helpers. No tests outside workflow may change. New failing
cases must map only to the changed/new IDs, not unrelated existing specs.

Only needed interface additions/compiling Eligible stub and evals in evals;
real selection changes follow the independent assertion audit. Workers never
commit or invoke governing lifecycle commands, alter runtime/counters/refs in
this checkout, or edit SPEC/protected docs/other modules. Fixture-local public
workflows/Git are allowed with explicit identity. No bootstrap source/tests,
deleted code, rebuild-v1 or saved implementation attempts as oracle. Own Verify
below five minutes, full below fifteen. No new phase-write exception or drill.

Runner wiring and actual CLI continuation immediately follow application;
status explanations, approval skipping, from-attempt resume, Questions,
agent-commit undo, usage/time stops and remaining kinds/bounds remain deferred.

Re-authored after the archived fixture correction and requested progress review.
Reuse only the retained author requirements, public facade and public evals;
recreate the implementation fresh. The next applied change reaches the fifth
change checkpoint; do not wire runner before that required review.
