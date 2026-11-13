---
module: app
eval_fix: [APP-076, APP-079, APP-083, APP-086]
---

## Motivation

During regenerate-runner, usage-stop output contained completed duration,
next wait and budget but failed app evals because it used "waited" instead
of the literal word "completed". RUN-031 and APP-076/079/083/086 require the
wait/budget information, not that word. The owner approved correcting the
evals while retaining all observable stop behavior.

## Interface

Audited eval_fix only. Edit only the app evals zone and this change.

## Added

## Modified

## Removed

## Decisions

## Notes

Revise cliStopReason in packages/app/evals/usage_cli_test.go and only the
call sites needed for this correction. Check that completed wait, next wait
and budget are present with the expected values and unambiguous roles.
Accept equivalent labels such as "waited" and "completed wait"; merely finding
unassociated durations is insufficient.

Keep existing checks for exit codes, held change identity/work/claims,
already-applied results, resume guidance, JSON schema/data/diagnostics and
stderr behavior. Preserve the distinction between clean stops and failures
preceded by a blocked change. Keep coverage of every listed eval_fix ID and
all existing timeout and usage-budget settings.

Do not change runner or app implementation, specs or enforcement limits.
Drafted and queued only. No lifecycle execution is authorized by the drafting
request.
