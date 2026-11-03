---
bases: {}
module: checks
---

## Motivation

ACCEPTANCE C2 requires persistence of every agent invocation's stdout and
stderr. Runner logging is applied; Doctor executes worker and audit probes
independently in checks. Complete C2 for that execution path while preserving
all canary, reply, timeout, usage, diagnostics and cleanup behavior. This is
selected agent logging, not the deferred CK contracts/proof milestone.

## Interface

No new public fields or function signatures are required. Each executed probe's
existing DoctorCheck.Detail includes log: <absolute path>, retaining its existing
outcome information. Normal doctor CLI detail output already exposes that field.

## Added

### new:agent-probe-logs
Doctor MUST persist complete stdout and stderr from every executed worker and audit probe, including failed, timed-out and cancelled probes, under the invoking project's .fireweed/runs/<run>/<phase>-1.log. Each Doctor invocation MUST use a fresh shared run identity, with phases doctor-worker and doctor-audit. Each executed probe's Detail MUST identify its log file by absolute path.

## Modified

## Removed

## Decisions

## Notes

Keep log files under the original approved project root, outside the scratch
worktree so normal cleanup retains them. Use ignored runtime bookkeeping only;
never alter tracked caller files. NoAgent continues to skip probe execution.
Preserve captured bytes and existing outcomes, final-20-line diagnostics,
usage classification, fixed-reply assertions, timeouts and scratch cleanup.
Do not change the agent argv, weaken hook/coverage checks or grow their bounds.

Use representative bounded fake-agent public API evals. Existing fixture setup
may add .fireweed/runs/ to its gitignore, matching init's valid project layout;
retain every existing caller-cleanliness assertion. New Added ID evidence must
be genuine failing testcases, not compile failures. Do not add a second real
three-minute timeout wait; retain the existing timeout proof and use bounded
request timeouts for new logging cases. Avoid redundant expensive fixture runs.

Evals/impl workers run pinned fireweed context, use fireweed verify --quick for
feedback and return, polling live handles until terminal. Only the outside
owning controller advances phases, commits, applies or archives and runs normal
full verification. Never edit run/phase/Provenance or frozen evals/interfaces.

The separate restricted auditor edits only the existing Audit section body,
retaining every other byte. Findings use - ID: "exact clause" — missing assertion.
A passing audit writes No findings and Checked IDs. Re-audits inspect only IDs
with prior findings; retain prior clearances without repeating their review.

## Audit

