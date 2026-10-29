---
audit_rounds: 0
bases: {}
eval_fix:
    - RUN-016
module: runner
phase: impl
run: 20261029T141235-0e6c1ce6b922f1a559353659
---



## Motivation

The design's usage limits repeat the same phase attempt and retain invocation
results. During 0093-usage-limit-control, the new shared-budget scenario applied
Alpha with four captured invocations, but its reused assertRunApplied helper
required exactly three. The owner approved native abandonment and this
fixture-only prerequisite before re-authoring usage handling. RUN-016 returns
applied NextResults in execution order; its existing ordinary-success tests
must keep the exact three-step assertion. Factor an exact-count helper so the
later usage scenario can explicitly require four without discarding results.
The observable slice remains NEW usage wait/success/budget stop/resume, after
this fifth-change quality checkpoint.

## Interface

No public interface, model, product behavior or Spec change.

## Added



## Modified



## Removed



## Decisions



## Notes

Only packages/runner/evals/run_test.go may change. Keep assertRunApplied's
signature and make it a wrapper supplying exactly three to a factored helper,
assertRunAppliedSteps(t, f, applied, stepCount, want...). The factored helper
keeps every existing count, state, slug, root, run, archive/audit and genuine
fail-first/pass report assertion; substitute stepCount only for its literal 3.
No test/case/coverage-ID deletion, assertion weakening, new product feature or
new test matrix. Ordinary callers remain unchanged. Do not add the future usage
scenario in this eval_fix. All own evals must pass; use the normal independent
isolated RUN-016 assertion audit and full native phases. Impl edits nothing.

Workers never commit or run governing workflow commands. No implementation,
interface, SPEC, other module, protected infrastructure/doc, controller runtime,
claim, ref, audit counter or bound edits. No bootstrap source/tests, rebuild-v1,
deleted or abandoned implementation/worker logs as oracle, phase-write exception
or drill. The owner explicitly approved consecutive runner recovery changes.
After this fifth application, separate quality/goal review and required checks
must PASS before the authorized push and continuation; findings stop for owner.

## Provenance

- evals 2026-10-29T14:18:35.723119Z agent=manual fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=231 failing=0 passing=231
- audit 2026-10-29T14:20:28.661567Z agent=manual fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0

## Audit

No findings
Checked IDs: RUN-016
