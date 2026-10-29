---
audit_rounds: 0
bases: {}
module: workflow
phase: impl
run: 20261029T173024-225998e74522a3f430076281
---



## Motivation

The design's Kinds of change and lifecycle send maintenance directly to
implementation. NEW currently sets evals for every fresh change, so next asks
for tests and an audit even when the author requests only a refactor. Establish
the missing workflow selection, then close the slice with actual NEW next in a
disposable Go project: one implementation agent, no evals/audit phase, full
checks and one archive commit while existing specs/interfaces/evals survive.

## Interface

No public API/model changes. Existing workflow.Start and normal Verify/Done
compose this path. Regenerate deletion and eval_fix semantics are separate
remaining design work; do not implement them in this change.

## Added

### FLOW-077
For an unstarted change with no Added, Modified or Removed specs, fixes or
eval_fix, and regenerate not true, Start MUST set phase to impl.

## Modified



## Removed



## Decisions



## Rejected alternatives



## Notes

Use a representative public workflow fixture with an actual maintenance body
and no fake spec edits. Assert returned and persisted impl phase, unchanged
HEAD, held claim/run and no hook execution at Start. Drive ordinary Verify/Done
with passing current coverage and show a single impl archive/provenance commit,
no evals/audit transition and retained specs/interfaces/evals. Existing universal
Start/Verify/Done requirements already govern these lifecycle properties;
reference them rather than restating them in new specs. Preserve existing
assertions, public interfaces and spec/fix phase behavior. No code organization,
eval_fix/regenerate feature, new diagnostic, command, instrumentation, drill,
matrix or unrelated cleanup. One new file may contain focused public evals;
permanent fixtures use bounded deterministic hooks. Own check under5m, full
under15m. Workers never commit or advance phases; controller owns transitions.

## Provenance

- evals 2026-10-29T17:36:30.115891Z agent=manual fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=353 failing=1 passing=352
- audit 2026-10-29T17:37:29.494998Z agent=manual fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0

## Audit

No findings
Checked IDs: FLOW-077
