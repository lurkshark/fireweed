---
bases:
    APP-076: 3733d712
    APP-077: 41f0abad
    APP-079: c66f97d6
    APP-083: b91f86e4
    APP-084: 357a4318
    APP-086: def1e432
module: app
---



## Motivation

The design requires exit3 for clean usage/time limits, exit2 for configuration
errors, and recoverable state after a usage stop. The runner dependency now
waits/retries, shares wait budget, preserves applied results and held work, and
classifies clean usage stops, but next/run still report every nonnil runner
error as failure1/null data. Wire this caller promptly. Close the slice through
actual NEW in a Go fixture: Alpha waits/retries/applies, Beta hits the shared
budget with its claim/work/run/ID retained, then a later next resumes and
archives Beta while Gamma stays untouched. Reuse valid normal-path assertions.
This change handles usage stops only; --for follows in its own observable slice.

## Interface

No public Go facade/model change. Extend private CLI JSON success-data shapes
with an optional reason string for clean stops. Stop is a non-error outcome:
retain warnings, use data/plain reason and resume guidance, and do not invent
an error diagnostic code for clean budget exhaustion. Ordinary failure shapes
and wrapped cause codes remain; C01 usage-policy errors are config exit2.

## Added



## Modified

### APP-076
Plain next MUST identify the applied change or empty queue after successful
runner completion. On a clean usage stop it MUST name the held change and show
the wait/budget reason with instructions for resuming it.

### APP-077
next MUST return exit0 for an applied change or empty queue, exit1 for ordinary
runner execution or lifecycle failure, exit2 for usage or configuration errors,
and exit3 for a clean usage stop.

### APP-079
next --json MUST emit schema1 and leave stderr empty. It MUST use object data
containing change, run, phase and empty on success or a clean usage stop, with
a nonempty reason reporting completed wait, next wait and budget on a stop.
On failure it MUST use null data with code/message diagnostics.

### APP-083
After successful runner completion, plain run MUST identify applied changes
and their count, including whether the queue is empty or Max left queued work.
On a clean usage stop it MUST retain the applied names/count and show the
wait/budget reason with instructions for resuming.

### APP-084
run MUST return exit0 for successful queue completion, Max termination or an
empty queue, exit1 for execution or lifecycle failure including any earlier
blocked change, exit2 for usage or configuration errors, and exit3 for a clean
usage stop.

### APP-086
run --json MUST emit schema1 and leave stderr empty. It MUST use object data
containing applied change/run/phase summaries, empty and limited on success or
a clean usage stop, with a nonempty reason reporting completed wait, next wait
and budget on a stop. On failure it MUST use null data with code/message
diagnostics.

## Removed



## Decisions



## Notes

Normal claims, protected metadata, audit/bounds, requests/Max, earlier-block
failure precedence, captured work and phase ownership remain in runner/workflow.
App classifies runner.IsUsageStop before ordinary error handling, returns the
available held Next state or Run applied results, gives resume guidance, and
maps genuine runner C01 policy errors to configuration exit2. Do not advance,
release, block or correct a claim from app. JSON clean-stop reason is data.reason;
plain output goes to stdout. Do not label a clean stop an error; warnings remain.
Existing malformed config/arguments, ordinary failures/codes/null JSON and
success/empty/Max behavior remain asserted. No --for, new flags/framework,
assets diagnostic, other command or unused general control behavior.

Evals-only step adds representative public CLI cases using current bounded
fixture agents and genuine Go hooks. Cover JSON next stop and later same-owned
run/phase/ID/work resume; plain next reason/guidance; JSON and plain Run stop
with earlier Alpha applied and Beta held; and invalid usage policy exit2/C01
before mutation/agent for both commands. Use short waits/attempts1 and exactly
captured invocations (usage exit adds one); preserve claims/work on clean stops.
Never reuse ordinary success/failure helpers that require three invocations,
clean/unclaimed state or null data at a clean stop. Assert the appropriate new
outcome directly, keeping every old helper/assertion/caller unchanged. Avoid
new test matrices/instrumentation; own Verify <5m, full <15m. Interface and
implementation edits are not allowed in evals; current code compiles.

Workers never commit or run governing workflow commands, change other modules,
SPEC/protected docs/infrastructure/interfaces, runtime/phase/Audit/Provenance,
claims/refs/counters/bounds or fixed profiles. Coordinator owns normal phases.
No bootstrap source/tests, rebuild-v1, deleted/abandoned implementation, old
worker logs or saved attempts as oracles. No phase exception, delegation or
drill. After archive, build NEW and run the prepared actual usage CLI fixture;
normal checkpoint/P07 policy continues (one applied since review before this).
