---
eval_fix:
    - APP-125
module: app
---

## Motivation

APP-125 requires `status` to keep the P08 warning workflow returns. Its
eval gets that warning by archiving three changes to module `alpha` that
edit no specs, which pins workflow's current trigger rather than APP-125.
0217 changes the trigger to chains of changes that rework the same spec
IDs, so this fixture would stop producing P08 and the workflow change could
not pass its dependents' evals. Build a fixture that triggers P08 under
both rules.

## Interface

Audited eval_fix only. Edit only the app evals zone and this change.

## Added

## Modified

## Removed

## Decisions

## Notes

In `Test_APP_118_APP_119_APP_120_APP_121_APP_122_APP_123_APP_124_APP_125_StatusRichReport`,
give the three `alpha` archives (0005 to 0007) spec edits that form a chain:
0005 adds ALP-001, 0006 modifies ALP-001, and 0007 modifies ALP-001 again,
using an ID the fixture's alpha spec already has. Keep them the three most
recent applied changes on `alpha`, so the current same-module rule also
warns. Extend `statusArchive` with optional Added and Modified items rather
than writing archives by hand. Keep every other assertion.

Run this before 0217.
