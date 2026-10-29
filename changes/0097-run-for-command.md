---
audit_pending:
    - APP-089
audit_rounds: 1
bases:
    APP-084: 070b95dc
module: app
phase: evals
run: 20261029T165024-3fef7348df22114b38bedf67
---



## Motivation

The design exposes run --for6h and clean exit3 for time limits. The runner now
provides a single elapsed-time bound and IsTimeStop, but the CLI rejects --for
and only formats usage stops. Wire the caller immediately. Close this slice
through actual NEW in a Go fixture: run --for2s stops an agent after it writes
work, retains the original claim/run/phase/IDs without a commit, then next
completes that same run with real fail-first/pass evidence and later queue
unchanged. Preserve the already working usage, Max, approval and recovery paths.

## Interface

No public Go facade/model changes. Parse --for as a positive Go duration for
run only, pass it as runner.RunRequest.For, and use runner.IsTimeStop for clean
classification. Reuse existing run JSON/plain result fields and data.reason.
No new diagnostic allocation; invalid flags use corrective U01. Existing usage
reason and ordinary failure shapes/codes stay unchanged.

## Added

### APP-089
run --for D MUST require one positive Go duration and pass it as RunRequest.For,
using zero when absent. It MUST reject missing, malformed, nonpositive or repeated
values and use by other commands with corrective U01 and exit2 before mutation,
showing valid run usage.

### APP-090
On a clean runner time stop, run MUST return exit3 with retained applied summaries
and a nonempty reason naming the time limit. Plain output MUST include the
applied count and resume guidance. JSON MUST use the existing schema1 run object
with reason and empty/limited flags, leaving stderr empty.

## Modified

### APP-084
run MUST return exit0 for successful queue completion, Max termination or an
empty queue, exit1 for execution or lifecycle failure including any earlier
blocked change, exit2 for usage or configuration errors, and exit3 for a clean
time or usage stop.

## Removed



## Decisions



## Notes

Use both clean classifiers rather than matching error wording; mixed blocked
failures remain exit1/null/cause/ref+stop reason. Preserve data.reason for usage,
normal success/empty/Max output and warnings. --for must compose with existing
--max/--allow-unmerged, not change their rules. Do not add a next time flag,
new clock/model fields, general scheduler, diagnostics, other commands or claim/
phase mutation. Current run/next continuation APIs handle recovery.

Evals-only step adds representative real public CLI cases with bounded fake
agents and genuine local Go hooks. Cover JSON time stop retaining earlier Alpha
application with Beta held, then same-run Beta resume/work/ID preservation;
plain time-stop reason/count/guidance; and prior-block/later-time-stop exit1
precedence. Use generous short duration and a longer one-time fixture sleep;
do not reuse wait/budget-specific usage reason checks for time errors, or
success helpers requiring clean/unclaimed state at a stop. Capture original
public claim/run identity before interruption, or through ordinary fixture agent
Git/document logs, and compare after. Keep actor selection from picking a
blocked document in the precedence fixture; existing permanent helpers stay
unchanged. No process instrumentation, fake clock, timing matrix or redundant
confidence tests. Invalid option cases should assert U01/exit2, valid run usage,
unchanged files/HEAD/claims and no agent calls; representative other-command
rejection suffices. Existing assertions/callers remain intact. Own Verify<5m,
full<15m; do not raise limits.

Workers never commit or run governing workflow commands. No interface/model,
SPEC/other-module/protected docs/config/infrastructure, governing author/runtime/
phase/Audit/Provenance/ref/claim/counter/bound edits, delegation or drill. No
bootstrap/rebuild-v1/deleted/abandoned/saved implementation or old worker logs as
oracles. Coordinator owns normal phases and then builds clean NEW and runs the
prepared actual time fixture. Three applications since review precede this one;
normal fifth-change checkpoint remains in force.

## Provenance

- evals 2026-10-29T17:05:15.902012Z agent=manual fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=196 failing=10 passing=186
- audit 2026-10-29T17:08:27.147018Z agent=manual fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0

## Audit

- APP-089: "showing valid run usage" — checks only run and --for tokens, without usable command syntax or a duration value/placeholder.
