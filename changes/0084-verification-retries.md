---
bases:
    RUN-015: 5d5d135e
module: runner
---



## Motivation

The design's Retries section requires bounded phase retries that retain work
and add the last verification output to the next prompt. NEW next and run now
complete successful Go changes, but a single failed phase verification stops
them with no retry. Extend that existing runnable path with verification
retries, then demonstrate an actual NEW next whose first implementation is
wrong and whose second fixes it after receiving feedback. This slice does not
claim saved-attempt/block recovery, nonzero-agent retries or usage/time stops;
those remain subsequent focused steps.

## Interface

No public signature or retained type change. Next and its Run caller read the
existing attempts configuration; invocation results remain ordered in Steps.

## Added

### new:phase-attempt-bound
Next MUST permit at most the configured attempts count of agent invocations
for a phase after verification failures, defaulting to two when unset. Each
phase MUST have its own attempt allowance.

### new:verification-feedback
When phase verification fails and an attempt remains, Next MUST retain the
earlier work and retry that phase. The retry prompt MUST include the last
verification output.

## Modified

### RUN-015
When invocation fails, or phase verification still fails after its attempt
allowance, Next MUST return an error and captured invocation results without
advancing that phase.

## Removed



## Decisions



## Rejected alternatives



## Notes

Keep this change limited to verification retry and its existing Next/Run
composition. Do not prebuild a recovery subsystem or change approval, audit,
claims, agent selection, timeout defaults, diagnostic classification or APIs.
Do not treat usage limits as ordinary verification failures. Native Done
failures retain their existing stop behavior; saved work/blocking and full
recovery will follow in their own changes. Preserve audit round bounds.

Use public interfaces, ordinary temporary repositories and bounded fake agents.
Cover the default two-attempt retry success and exhaustion, one configured
alternative bound, retained edits and actual verification feedback. Reuse the
existing Go fixture where practical; preserve all old behavioral assertions,
including phase retention and captured output, while adapting RUN-015's
invocation count to its changed requirement. No test outside the target module
may be changed. Keep own Verify below five minutes and full below fifteen.

Workers never commit or invoke controller lifecycle commands, alter runtime
state/bounds, edit SPEC/protected docs or other modules. The coordinator reads
context and owns native phase operations. Separate auditors receive isolated
change/projected specs/evals/permitted public interfaces only.
