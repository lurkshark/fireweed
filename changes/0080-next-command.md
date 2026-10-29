---
audit_rounds: 0
bases: {}
module: app
phase: impl
run: 20261029T000645-7551044ce491f3dfa364881d
---



## Motivation

The design's Driving the state machine and CLI sections require next to drive
one change through its checked phases. 0079 provides the minimal runner loop,
but NEW still rejects next. Wire it now and close this vertical slice with an
actual NEW next invocation in an ordinary Go fixture using separate bounded
fake worker/auditor scripts. Inspect fail-first, audit, implementation, phase
commits, SPEC/archive/provenance and claim release. No broad recovery or run
implementation is a prerequisite for this first observable lifecycle.

## Interface

No public interface change. Compose the existing app.Run and runner.Next APIs.
Keep parsing, diagnostics and presentation in app internal implementation.

## Added

### APP-074
next MUST discover the nearest ancestor fireweed.yaml and invoke runner.Next
for that project, supplying the actual invoking executable as Bin.

### APP-075
next MUST accept --allow-unmerged, passing true only when present. It MUST
reject extra positional arguments and unsupported options with corrective U01
and exit 2; --allow-unmerged MUST be rejected for unrelated implemented commands.

### APP-076
Plain next MUST identify the applied change or empty queue on stdout after
successful runner completion.

### APP-077
next MUST return exit 0 for an applied change or empty queue, exit 1 for runner
execution or lifecycle failure, and exit 2 for usage or configuration-loading
failure.

### APP-078
next failure diagnostics MUST retain existing cause codes from wrapped
dependency errors and include a corrective action.

### APP-079
next --json MUST emit schema 1 with object data containing change, run, phase
and empty on success, or null data with code/message diagnostics on failure.
It MUST leave stderr empty.

## Modified



## Removed



## Decisions



## Rejected alternatives



## Notes

This closes the current CLI slice, not the full execution surface. Do not
implement run, attempts/retries/recovery, drills, new options, agent commit undo,
whole-module completion or speculative cleanup. Existing runtime behavior,
manual workflow APIs and old assertions remain. No public interface edits or
new disposable stub should be needed for runtime fail-first: NEW currently
rejects next. Keep all other unsupported commands explicit errors.

Use existing nearest-project loading, output envelope and CLI helpers. Let
runner.Next and workflow retain eligibility/approval/claim/P07 enforcement.
Request Timeout stays unset so existing configured defaults apply. Plain
failures use stderr; success may retain dependency warnings. For unprefixed
runner errors use existing P01 precondition cause with specific corrective
text; do not allocate speculative codes. Inspect wrapped errors to preserve
an existing cause such as C10 or P07 despite runner phase context.

Black-box evals build/run the actual NEW executable with public interfaces,
ordinary temporary Git repositories, explicit identities, local hooks and
bounded separate fake worker/auditor scripts. Reuse existing helpers where
possible; no bootstrap, auth/network, private imports or observers. Exercise
one full lifecycle from a nested directory, actual FIREWEED_BIN in hooks/agents,
checked commit phase state (not assumed subjects), SPEC/archive/provenance,
claim release, empty queue, normal approval rejection and --allow-unmerged,
representative usage/config/failure diagnostics and JSON success/failure.
Keep cases representative and own Verify below five minutes. No fixture
helper redesign or Cartesian matrix. Run full native phase checks via Done.
After application build bin/fireweed and run the prepared /tmp CLI smoke.

The coordinator runs controller start/context/done and owns commits. Workers
never commit or call controller start/done/next/run, edit protected docs or
phase/counters. Separate auditors receive isolated permitted snapshots and
review exact unasserted spec clauses only. Normal bounds and checkpoint stops
remain; no phase-write exception is authorized.

## Provenance

- evals 2026-10-29T00:13:50.625193Z agent=manual fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=159 failing=6 passing=153
- audit 2026-10-29T00:15:38.705451Z agent=manual fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0

## Audit

No findings. Checked IDs: APP-074, APP-075, APP-076, APP-077, APP-078, APP-079.
The separate Sol/high auditor reviewed only supplied isolated snapshot inputs.
Complete command-log inspection confirms read-only permitted files, without
implementation/history access, execution or writes.
