---
audit_rounds: 0
bases: {}
eval_fix:
    - RUN-007
    - RUN-009
module: runner
phase: applied
run: 20261028T231208-cfb868c63baddc1bfdad0dda
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

## Provenance

- evals 2026-10-28T23:15:35.96126Z agent=manual fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=178 failing=0 passing=178
- audit 2026-10-28T23:17:21.243427Z agent=manual fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0
- impl 2026-10-28T23:17:24.084547Z agent=manual fireweed=0.1.0 playbook=07554a3249d8fca0ba70df7504135f297491465caa63458734a4b8351734d0e5 tests=178 failing=0 passing=178

## Audit

No findings. Checked IDs: RUN-007, RUN-009. The isolated auditor reviewed the
existing fixture correction and retained assertions using permitted inputs only.
