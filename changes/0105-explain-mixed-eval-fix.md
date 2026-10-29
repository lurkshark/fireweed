---
bases: {}
module: assets
---



## Motivation

Owner approved the checkpoint0099–0103 correction gate. PRJ-046 now emits C07
for eval_fix mixed with fixes or spec edits, closing the design kind-exclusion
bypass. D006 requires accurate explanation for emitted causes. Extend C07 with
this cause and corrective action before rebuilding NEW for the prepared actual
CLI rejection proof and the same checkpoint reviewer's scoped re-review.

## Interface

No API/model changes. Extend existing Explain C07 data only.

## Added

### new:mixed-eval-fix-explanation
Explain MUST describe C07 as including eval_fix combined with fixes or spec
edits. It MUST advise separating eval-fix work from fix or spec changes.

## Modified



## Removed



## Decisions



## Notes

One focused public Explain regression. Preserve C07's existing Motivation and
module causes/actions, all existing assertions and unknown-code rejection.
No new codes, APIs, recipes, hooks, versions, unrelated prose changes or other
module/protected edits. Evals freeze before impl; workers never commit or
advance phases. Coordinator owns native verification, isolated audit/archive,
actual NEW proof and scoped checkpoint re-review. Own under5m/full under15m;
no raised bounds. Five prior changes plus0104 remain counted; this is the
explicitly approved second correction application, not a checkpoint reset.
