---
already_satisfied:
    - RUN-061
bases:
    RUN-061: 16edb595
module: runner
---

## Motivation

The design's Regeneration workflow keeps rebuilt code for inspection. In the
fresh runner drill report in commit `b25dc33`, APP-142 stopped on a usage
limit during a retry and returned an empty ref even though the first
reconstruction had been accepted and evaluated.

## Interface

The existing `DrillResult.Ref` identifies the latest accepted rebuild when a
retry stops on the usage budget.

## Added

## Modified

### RUN-061
After accepted reconstruction, Drill MUST retain the rebuilt verification tree at `refs/fireweed/drills/<module>/<date>` and return that ref without moving the invoking HEAD, including on a usage stop during a later attempt. A passing result MUST require accepted reconstruction and passing own and all dependent evals. An ordinary failure MUST return a nonnil error with Result fail.

## Removed

## Decisions

## Notes

Add a public `runner.Drill` fixture with two attempts and a zero usage budget.
Have the first agent invocation produce an accepted rebuild that fails its
own eval. On the retry, have the agent write a different disposable value
and exit with the configured usage-limit marker. Assert `runner.IsUsageStop`,
Result `stopped`, and a `Ref` that resolves to the first rebuild's value.
Check that invoking HEAD is unchanged and scratch workspaces are removed.
Use the existing usage-stop fixture before any acceptance to assert an empty
ref and no retained drill refs in that fixture.

The current runner writes the ref when it accepts a reconstruction, so it
already returns that ref on a later usage stop.
