---
eval_fix:
    - APP-038
module: app
phase: audit
run: 20261113T172337-69d40369829ae74d6220dd79
---

## Motivation

During regenerate-project, app's lookup eval rejected a U02 message that
named the requested selector and valid pending slugs because it said
"unknown pending change" instead of "not found". APP-038 assigns the code;
PRJ-004 specifies the selector and valid-slug information, not that phrase.
The owner approved correcting this assertion so reconstruction is checked
against the diagnostic contract rather than incidental prose.

## Interface

## Added

## Modified

## Removed

## Decisions

## Notes

Audited eval_fix only. Edit only the app evals zone and this change.

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

Eval verification passed with `fireweed verify lookup-diagnostic-evals`.
The controller reported B02 case-count warnings for APP-103 and APP-104 and
a B03 eval-to-disposable-lines warning for app.
## Provenance

- evals 2026-11-13T17:32:23.65404Z agent=codex fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=369 failing=0 passing=369

