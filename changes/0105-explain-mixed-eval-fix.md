---
audit_pending:
    - AST-033
audit_rounds: 1
bases: {}
module: assets
phase: evals
run: 20261029T203110-4257e03113ee7e76c73e09e4
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

### AST-033
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

## Provenance

- evals 2026-10-29T20:33:27.690785Z agent=manual fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=592 failing=1 passing=591
- audit 2026-10-29T20:35:22.806392Z agent=manual fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0

## Audit

- AST-033: "It MUST advise separating eval-fix work from fix or spec changes." — missing assertion.
