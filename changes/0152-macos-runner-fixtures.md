---
bases: {}
eval_fix:
    - RUN-001
    - RUN-002
    - RUN-003
    - RUN-004
    - RUN-005
    - RUN-006
    - RUN-007
    - RUN-008
    - RUN-009
    - RUN-010
    - RUN-011
    - RUN-012
    - RUN-013
    - RUN-014
    - RUN-015
    - RUN-016
    - RUN-017
    - RUN-018
    - RUN-019
    - RUN-020
    - RUN-021
    - RUN-022
    - RUN-023
    - RUN-024
    - RUN-025
    - RUN-026
    - RUN-027
    - RUN-028
    - RUN-029
    - RUN-030
    - RUN-031
    - RUN-032
    - RUN-033
    - RUN-034
    - RUN-035
    - RUN-036
    - RUN-037
    - RUN-038
    - RUN-039
    - RUN-040
    - RUN-041
    - RUN-042
    - RUN-043
    - RUN-044
    - RUN-045
    - RUN-046
    - RUN-047
    - RUN-048
    - RUN-049
    - RUN-050
    - RUN-051
    - RUN-052
    - RUN-053
    - RUN-054
    - RUN-055
    - RUN-056
    - RUN-057
    - RUN-058
    - RUN-059
    - RUN-060
    - RUN-061
    - RUN-062
module: runner
---

## Motivation

ACCEPTANCE C1 requires passing local Go evals without changing product behavior.
Accept the owner's supplied diagnosis that macOS sh kills fake-agent scripts
receiving long prompt arguments. Correct the runner's shared fake worker,
auditor, wrapper and reconstruction fixtures to use bash, or provide prompts
through stdin. Preserve every assertion, spec obligation and configured bound;
do not re-investigate the OS failure.

## Interface

None. This eval_fix changes no specs or interfaces and needs no new handles.

## Added

## Modified

## Removed

## Decisions

## Notes

The shared next/run/drill/invocation fixtures support the listed runner evals,
including retries, audit isolation, usage/time stops and reconstruction. Review
all prompt-bearing scripts, including direct executable shebangs and wrappers,
and change only their fixture interpreter or prompt transport. Keep every
assertion, testcase name, timeout, retry/round allowance, coverage requirement,
fixture semantics and product implementation intact. Do not alter hook scripts
that receive no prompt merely for consistency.

Only runner eval files may change. Use normal fireweed verify and return;
workers never invoke lifecycle mutators, commit, edit run/phase/Provenance or
advance themselves. The owning pinned controller performs normal phase checks,
separate restricted audit, full dependent verification and archive commits.
