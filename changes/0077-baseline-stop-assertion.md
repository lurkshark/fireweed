---
bases: {}
eval_fix:
    - CHK-055
module: checks
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
