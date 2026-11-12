---
already_satisfied:
    - RUN-056
bases:
    RUN-056: f6dd7bed
module: runner
phase: applied
run: 20261111T002924-87fad14e6cf2106f6ce6245d
---

## Motivation

The design's drill step 3 calls dependent evals "the strongest signal": a
dependent that passes against the old code and fails against the rebuilt
code relies on behavior the spec never states. The current drill feeds
dependent failures back into the retry prompt, so the worker fits its code
to them and the report still says `result: pass`. Five of the eight
regeneration proofs passed only on that retry (assets, project, workflow,
checks, runner), so their pass results do not show that the spec was
sufficient. The project worker's retry summary says it changed wording to
"match the graph and queue wording expected by app".

The design also asks for a report the developer reviews. Each attempt's
Agent observations currently embeds the agent's full stdout and stderr. For
Codex, stderr is the whole session transcript, including the prompt and
every file the worker read. That made the committed reports 3k to 38k lines
long, and they include temporary paths and session IDs.

## Interface

No API or configuration changes. Retry prompts and the report's Agent
observations section change as specified.

## Added

### RUN-068
Drill retry prompts MUST NOT include dependent modules' eval output. When the target's own evals pass and only dependent evals fail, Drill MUST NOT retry reconstruction and MUST return Result fail with those dependent failures in the report.

### RUN-069
Each attempt's Agent observations entry MUST contain that attempt's agent stdout and the repository-relative path of its invocation log, and MUST NOT contain the attempt's stderr.

## Modified

### RUN-056
Drill MUST use the configured worker command, wrapper and timeout through InvokeAgent, retaining ordered AgentResults in Steps. Non-usage failures and verification retries MUST share the configured attempts bound, retain rebuilt work and include the preceding reconstruction or own-eval failure in retry prompts.

## Removed

## Decisions

## Rejected alternatives

Reporting the first attempt's result as the headline and keeping the retry:
the retry still trains the rebuilt code on its dependents, so the retained
reconstruction stops being an independent rebuild from the spec.

Dropping verification retries entirely: a retry on the target's own eval
failures is still useful, and those evals are inputs the worker already has.

## Notes

Narrowing RUN-056 to own-eval feedback keeps current behavior for the
target's own failures, so it is listed in already_satisfied.

Evals: one fake worker whose rebuild passes its own evals but fails a
dependent must yield one agent step, Result fail, and a report listing the
dependent failure. One whose first rebuild fails its own evals must get a
retry prompt that contains the own failure and none of the dependent's
testcase names or output. A fake agent that writes distinct markers to
stdout and stderr must leave only the stdout marker and a
`.fireweed/runs/...` log path in Agent observations; the full log file
still holds both streams (RUN-063).

After this applies, earlier passing drills are not invalidated, but the next
drill of any module may fail where its dependents rely on unstated behavior.
That failure is the measurement; turn agreed gaps into ordinary changes.

On the current implementation, the focused public Drill tests observed two
invocations after dependent-only failure and found both dependent testcase
names and hook output in an own-eval retry prompt (RUN-068). Both Agent
observations entries included stderr and lacked their relative log paths,
while the invocation logs retained both streams (RUN-069). The existing
RUN-056 retry test passed.

The implementation-writable file is `packages/runner/internal/drill/drill.go`.
The retry loop, eval failure collection and report writer are the affected
paths. No stub is needed.
## Provenance
- evals 2026-11-11T00:46:49.43497Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=165 failing=3 passing=162
- audit 2026-11-11T00:47:13.586977Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-astra -c model_reasoning_effort="medium" -s workspace-write {prompt} fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
- impl 2026-11-11T01:07:26.063872Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=5f856d6bcab308b9ff1a8a7c03b9d2547744112aba81dd30dea8f61ec8062b03 tests=488 failing=0 passing=488
## Audit

No findings; checked RUN-056, RUN-068, RUN-069
