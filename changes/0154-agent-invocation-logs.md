---
bases: {}
module: runner
---

## Motivation

ACCEPTANCE C2 selects the design's run identifier that also locates its logs,
and its gitignored .fireweed/runs layout. Persist every runner-owned agent
invocation's stdout and stderr so worker, audit and reconstruction failures
can be diagnosed after the subprocess returns. Keep existing isolation,
output, usage, cancellation and bounded retry behavior intact.

## Interface

Extend the retained AgentRequest with optional LogRoot, Run, Phase and Attempt
fields describing log context. Extend AgentResult with LogPath. Existing
fields and public function signatures remain. Next and Drill supply their
invoking project root and stable run context internally; standalone callers
may supply context or use the defaults in new:standalone-log-context.

## Added

### new:invocation-log
InvokeAgent MUST persist the invocation's complete stdout and stderr, including partial output on command failure, timeout or cancellation, at .fireweed/runs/<run>/<phase>-<attempt>.log under its log project root. Repeated invocations at the same path MUST append without discarding earlier records. AgentResult.LogPath MUST identify that file by an absolute path.

### new:next-log-context
Next MUST assign each invocation the held change's run, current phase and one-based phase attempt, writing its log under the invoking project's .fireweed/runs directory. Audit logs MUST stay in that invoking project, outside the restricted audit input workspace.

### new:drill-log-context
Drill MUST allocate one log run identity per invocation and reuse it across reconstruction retries and usage waits. Its logs MUST use phase drill, the current one-based reconstruction attempt and the invoking project's root, outside scratch and reconstruction input workspaces.

### new:standalone-log-context
InvokeAgent MUST accept a log project root, run, phase and positive one-based attempt in AgentRequest. Unset fields MUST default to Snapshot.Root, a fresh run identifier, request Role and attempt 1, respectively.

## Modified

## Removed

## Decisions

## Notes

All agent invocations owned by runner are in scope: Next workers/audits, Drill
reconstruction and direct InvokeAgent calls. Doctor probe logging follows in a separate checks-owned C2 change;
this runner portion does not complete C2 or authorize a real module drill. Use bounded fake agents
and temporary repositories for permanent public-interface evals. Prefer reuse
of representative fixtures/assertions and keep every pre-existing assertion.
New Added IDs need genuine failing testcase evidence, not compilation errors.

Keep existing result streams, exit codes, usage classification, timeouts and
retry/round accounting intact. Retain output on errors. Logs must remain under
the chosen log project's .fireweed/runs directory and must not be written into
restricted audit or reconstruction input workspaces. Do not relax input
inventories, zone checks, frozen phases or any existing enforcement to admit
logs. A record format may label streams while retaining their complete bytes;
record any necessary behavioral choice briefly in Decisions.

Evals/impl workers run pinned fireweed context before each step, use fireweed
verify --quick for feedback and return; poll live session handles until terminal.
Only the outside owning controller advances phases, commits, applies or archives,
and it performs normal full verification. Never edit run/phase/Provenance.

The separate restricted auditor edits only the existing Audit section body,
retaining all other bytes, including whitespace. Findings use the design's
one-line form: - ID: "exact clause" — missing assertion. A passing audit writes
No findings and Checked IDs. Re-audits inspect only IDs with previous findings;
retain earlier clearances without repeating their review.

## Audit

