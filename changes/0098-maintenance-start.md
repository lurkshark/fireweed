---
bases: {}
module: workflow
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

### new:maintenance-phase
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
