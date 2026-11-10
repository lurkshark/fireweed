---
already_satisfied:
    - RUN-056
bases:
    RUN-056: f6dd7bed
module: runner
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

### new:dependent-signal
Drill retry prompts MUST NOT include dependent modules' eval output. When the target's own evals pass and only dependent evals fail, Drill MUST NOT retry reconstruction and MUST return Result fail with those dependent failures in the report.

### new:observation-summary
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
