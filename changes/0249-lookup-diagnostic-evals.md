---
module: app
eval_fix: [APP-038]
---

## Motivation

During regenerate-project, app's lookup eval rejected a U02 message that
named the requested selector and valid pending slugs because it said
"unknown pending change" instead of "not found". APP-038 assigns the code;
PRJ-004 specifies the selector and valid-slug information, not that phrase.
The owner approved correcting this assertion so reconstruction is checked
against the diagnostic contract rather than incidental prose.

## Interface

Audited eval_fix only. Edit only the app evals zone and this change.

## Added

## Modified

## Removed

## Decisions

## Notes

In Test_APP_038_LookupCauses in
packages/app/evals/diagnostic_causes_test.go, replace the "not found"
requirements for stamp-change and coverage-change with checks for U02, the
rejected selector and the valid pending-slug information required by PRJ-004.
Use a valid nonempty queue fixture if needed to make the slug-list check
meaningful.

Retain exit status, plain/JSON failure shapes and the other lookup and static
cause assertions. Do not remove explicitly specified wording requirements
from unrelated tests or alter product messages, specs or implementation.

Drafted and queued only. No lifecycle execution is authorized by the drafting
request.
