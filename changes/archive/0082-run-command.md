---
audit_rounds: 0
bases: {}
module: app
phase: applied
run: 20261029T022255-defa616cfd7f19df2cabf3fe
---



## Motivation

The design's Driving the state machine and CLI sections require run to process
queued changes until empty or a limit. 0081 now provides the small bounded loop;
NEW next already passed an actual Go fixture lifecycle. Wire run promptly and
prove --max, a later continuation and empty queue through NEW with separate
bounded fake agents. Do not build a broad recovery subsystem before this
observable queue path.

## Interface

No public interface change. Compose app.Run with the existing runner.Run API;
keep parsing, output and diagnostics in app internal implementation.

## Added

### APP-080
run MUST discover the nearest ancestor fireweed.yaml and invoke runner.Run for
that project with the actual invoking executable as Bin.

### APP-081
run --max N MUST accept a positive integer and supply it as Max, using zero
when --max is absent. It MUST reject missing, noninteger or nonpositive values,
extra positional arguments and unsupported options with corrective U01 and
exit 2; --max MUST be rejected for unrelated implemented commands.

### APP-082
run MUST accept --allow-unmerged and supply AllowUnmerged true only when the
flag is present, using false by default.

### APP-083
After successful runner completion, plain run MUST identify applied changes
and their count on stdout, including whether the queue is empty or Max left
queued work.

### APP-084
run MUST return exit 0 for successful queue completion, Max termination or an
empty queue, exit 1 for runner execution or lifecycle failure, and exit 2 for
usage or configuration-loading failure.

### APP-085
run failure diagnostics MUST retain existing cause codes from wrapped
dependency errors and include a corrective action.

### APP-086
run --json MUST emit schema 1 with object data containing an applied array of
change/run/phase summaries, empty and limited flags on success, or null data
with code/message diagnostics on failure. It MUST leave stderr empty.

## Modified



## Removed



## Decisions



## Rejected alternatives



## Notes

This closes the bounded successful queue slice, not the complete run surface.
Do not implement --for, usage waits, retry/recovery, saved/block attempts,
approval skipping or released-claim independent continuation, other change
kinds, agent-commit undo, drills or unrelated cleanup here. Keep unsupported
options explicit errors. Future scenarios retain every design obligation.

Reuse nearest-project loading, schema/diagnostic helpers and current runner.Run.
Leave request Timeout unset for existing configured behavior. Extend the shared
--allow-unmerged parsing to run while preserving rejection for unrelated commands.
Use representative valid/invalid --max values; do not freeze expectations about
planned named options that later changes will implement. Applied JSON is an
array even when empty; summaries identify each change, run and applied phase.
Successful output may retain dependency warnings. Unprefixed runner failures
use existing P01; preserve a wrapped cause such as P07 and add corrective text.

Black-box evals run actual NEW binaries in ordinary temporary Go repositories
with explicit identities, local hooks and separate bounded fake agents. Reuse
existing helpers where possible; no bootstrap, agent auth, network, private
imports, process observers or test matrices. Validate representative fixture
setup through CURRENT public runner.Run before freezing CLI evals, using a
small temporary harness removed before controller verification. The prior
C11/C10 fixture mistake and inline-audit mismatch must not recur: use consistent
hook exit/report pairs and canonical separate-line No findings/Checked IDs.

Exercise a small independent queue from a nested directory with --max, inspect
applied/remaining changes, count/flags and actual FIREWEED_BIN, then continue and
verify no reapplication or claim remains; empty invocation must make no calls or
commits. Include approval flag/default, representative usage/config errors,
wrapped diagnostics and JSON success/failure. Use committed phase state, never
assumed commit subjects. Keep earlier assertions/interfaces frozen and own
Verify below five minutes. Full native Done is authoritative. After application
build NEW and execute the prepared /tmp/fireweed-run-cli-smoke.py scenario.

Coordinator owns native start/context/done and all commits. Workers never
commit or call controller start/done/next/run, edit protected docs or phase/
counters. Separate isolated Sol/high auditors check unasserted clauses only.
Normal checkpoint/audit/P07 bounds remain; no new exception is authorized.

## Provenance

- evals 2026-10-29T02:34:11.280553Z agent=manual fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=166 failing=7 passing=159
- audit 2026-10-29T02:36:42.300036Z agent=manual fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0
- impl 2026-10-29T02:44:00.207836Z agent=manual fireweed=0.1.0 playbook=07554a3249d8fca0ba70df7504135f297491465caa63458734a4b8351734d0e5 tests=166 failing=0 passing=166

## Audit

No findings
Checked IDs: APP-080, APP-081, APP-082, APP-083, APP-084, APP-085, APP-086

The separate Sol/high auditor used only permitted isolated snapshot inputs.
Complete command-log inspection confirms read-only supplied files, without
implementation/history access, execution or writes.
