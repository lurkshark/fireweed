---
already_satisfied:
    - CHK-028
bases:
    CHK-028: 31f5ab46
module: checks
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
