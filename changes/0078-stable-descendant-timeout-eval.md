---
bases: {}
eval_fix:
    - RUN-007
    - RUN-009
module: runner
---



## Motivation

The 0071–0075 full race check failed the descendant-timeout fixture when its
500 ms bound expired before readiness under load. A focused race rerun passed.
The checkpoint subagent classified this as a timing-dependent frozen eval and
also observed that readiness precedes expected stdout/stderr writes, allowing
cancellation to race those writes. The owner approved this normal eval_fix,
with product behavior and all existing coverage/assertions preserved.

## Interface



## Added



## Modified



## Removed



## Decisions



## Rejected alternatives



## Notes

Change only Test_RUN_007_RUN_009_TimeoutAndCancellationKillOrdinaryDescendant in
packages/runner/evals/invoke_test.go. Give the ordinary fake command a reasonable
bounded startup window, keep descendant completion later than its timeout, and
adjust observation windows consistently. Emit expected stdout/stderr before
publishing readiness. Preserve readiness validation, nonnil timeout/cancellation
errors, output-retention equality, prompt return, descendant completion-marker
absence, both timeout/cancel modes and all ID coverage. Do not add tests, skip
conditions, retries, implementation changes, observer frameworks or platform
branches. No product timeout, controller timeout, verification limit, interface
or spec changes. All own evals must pass; eval_fix requires no fail-first.

Run the edited parent case with -race first, then leave required full controller
phase checks to the coordinator; final full-repository race follows both fixes.
Keep the fixture modest and own Verify under five minutes. This is the observed
failure correction, not a broader timeout/calibration project.

The coordinator controls native start/context/done and all commits. Do not call
start/done/next/run, commit, edit phase/counters or protected docs. Implementation
has no source work. The separate auditor receives only an isolated snapshot of
change, projected specs, evals, public interfaces and governing rules. Review
only RUN-007/RUN-009 and preserve prior passes. Any new block or unresolved issue
stops for human review; no bound or phase-write exception is authorized.
