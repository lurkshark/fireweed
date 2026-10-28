---
audit_rounds: 0
bases: {}
eval_fix:
    - CHK-055
module: checks
phase: impl
run: 20261028T230031-7c37d81c38953bf7362a2c42
---



## Motivation

The supplemental restricted review of 0076 found that CHK-055 baseline-failure
assertions require C14 but do not assert that canary planting stops. The owner
explicitly approved this normal eval_fix to close that assertion gap before the
remaining runner correction and scoped checkpoint re-review. Preserve the
existing requirement and behavior; this change edits evals only.

## Interface



## Added



## Modified



## Removed



## Decisions



## Rejected alternatives



## Notes

Strengthen existing CHK-055 baseline-failure evals in evals/doctor_test.go so a
Doctor that reports a failed baseline but continues to plant a canary cannot
pass. Reuse ordinary local fixtures/hooks and their existing observable probe
logs or bounded fake-command results. Preserve all prior assertions, coverage,
IDs, cleanup controls, recovery cases and the three-minute CHK-065 regression.
No new process observer, instrumentation, helper framework or case matrix.
All own evals must pass; eval_fix requires no fail-first evidence. Product code,
interfaces and specs must stay unchanged in every phase; implementation has no
source work. Keep own Verify under five minutes and full checks under fifteen.
Use focused edited-case tests while working; native full phase checks already
include the expensive default-timeout case, so do not duplicate them with an
unnecessary whole-suite quick pass.

The coordinator controls native start/context/done and all phase commits. Do not
call start/done/next/run, commit, alter phase/counters, edit protected docs, or
change any implementation. The separate auditor receives an isolated snapshot
of this change, projected specs, evals and retained public interfaces only.
Audit only CHK-055's assertions; prior passes in other IDs stay passed. Further
unresolved issues or blocks stop for human review under the approved sequence.

## Provenance

- evals 2026-10-28T23:08:29.023366Z agent=manual fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=298 failing=0 passing=298
- audit 2026-10-28T23:10:52.732036Z agent=manual fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0

## Audit

No findings. Checked CHK-055. The isolated auditor confirmed baseline-stop
assertions and retained recovery assertions using permitted inputs only.
