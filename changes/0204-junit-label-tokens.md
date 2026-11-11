---
already_satisfied:
    - CHK-028
bases:
    CHK-028: 31f5ab46
module: checks
phase: impl
run: 20261111T014410-f7ffd23af86d40e5326ae401
---

## Motivation

The design's ID matching rules decide which testcases cover a spec, and
RUN-044 depends on failure details surviving a testcase named like
`TST-001_case_00`. CHK-028 says "unknown spec IDs" without saying which text
counts as an ID. The checks drill's first reconstruction counted a numbered
label as an unknown ID, rejected the whole report with C11, and passed only
after the retry showed it the runner failure. This states the rule the
current code already follows.

## Interface

No API changes. The current implementation already satisfies the item.

## Added

## Modified

### CHK-028
ParseJUnit MUST reject, with a C11-prefixed error, a testcase that references no supplied spec ID or that contains a token with a supplied spec's prefix, matched under the ID rules, naming no supplied spec. Text without such a prefix, including numbered labels, MUST NOT count as an ID reference.

## Removed

## Decisions

## Notes

Evals: a passing and a failing testcase named `ALP-001_case_00` (with
`ALP-001` supplied) must parse with ID `ALP-001` and, for the failing one,
its complete Case.Failure text. `ALP-001_ALP-009` with only `ALP-001`
supplied must still return C11.
## Provenance
- evals 2026-11-11T01:54:31.755042Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=129 failing=0 passing=129
- audit 2026-11-11T01:54:54.750226Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-astra -c model_reasoning_effort="medium" -s workspace-write {prompt} fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
## Audit

No findings; checked CHK-028
