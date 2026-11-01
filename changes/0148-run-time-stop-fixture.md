---
bases: {}
eval_fix:
    - RUN-040
module: runner
phase: audit
run: 20261101T232415-5d4da9457a9f8ef9cf7c9ba8
---

## Motivation

RUN-040 requires an earlier blocking cause and saved attempt ref to survive a
later time stop. Owner-authorized diagnosis reproduced candidate CI's failure:
the fake gamma agent's six-second sleep can finish before the unchanged
eight-second run budget, advancing to audit while the test expects held evals.
Keep that agent waiting for cancellation so the existing case observes the
intended interruption. Preserve every case, assertion, covered ID, runtime
implementation, public interface and configured bound.

## Interface

## Added

## Modified

## Removed

## Decisions

## Notes

The unchanged protected runner hook reproduced 132 cases with only RUN-040
failing at time_limit_test.go's held-phase assertion (audit instead of evals).
All blocking-cause assertions passed. Correct only the fake-agent wait in
Test_RUN_040_EarlierBlockRetainsCauseBeforeTimeStop; do not increase its deadline
or weaken any assertion. Use normal eval_fix verification and restricted audit.

The owner approved metadata repair and lifecycle restart on 2026-11-01,
including restart of attempt accounting affected by the authoring error. The
prior run was stopped and released before this repair. Reuse the preserved
one-line correction. For worker feedback use fireweed verify --quick; the
controller performs required full checks and owns every phase commit.
## Provenance
- evals 2026-11-01T23:30:42.628942Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=132 failing=0 passing=132
