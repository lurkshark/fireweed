---
module: app
eval_fix:
    - APP-009
    - APP-037
    - APP-116
    - APP-117
---

## Motivation

The design makes codes, not wording, the stable diagnostic contract.
APP-037 requires the CLI to keep each static failure's D006 code, but its
eval also requires phrases from the lower modules' messages, such as
`missing module` and `ambiguous change slug`. The APP-116/APP-117 eval
asserts `ambiguous change slug` against `project.Queue` directly. The
APP-009 evals look for the literal playbook heading `# Write evals`. None of
that wording is in any spec. The project and assets drills each needed a
retry fed by these app failures before they passed, and both reports list
the wording as a gap.

## Interface

Audited eval_fix only. Edit only the app evals zone and this change.

## Added

## Modified

## Removed

## Decisions

## Notes

- `Test_APP_037_StaticCauses` and `Test_APP_037_APP_038_ShowStaticCauses`
  (diagnostic_causes_test.go): keep the code, exit status, corrective-action
  and no-duplicate-prefix checks. Replace each wording phrase with the input
  the fixture broke when the CLI must name it to be correctable (the
  `fireweed.yaml` path, the missing module `missing`, the slug `repeat`, the
  ID `ALP-999`, the undeclared handle the C09 fixture references).
  Otherwise drop the phrase.
- `Test_APP_116_APP_117_ListUsageConfigurationAndQueueErrors`
  (list_cli_test.go): drop the `ambiguous change slug` check on
  `project.Queue`'s cause. Keep the requirement that the CLI reports C13
  and reproduces the public cause text.
- `Test_APP_009_APP_027_APP_028_StartContextAssignmentAndRejection` and
  `Test_APP_009_APP_027_APP_028_APP_040_StartPreservesApprovalWarning`
  (lifecycle_cli_test.go): compare with `assets.Playbook("evals")` instead
  of the literal `# Write evals`.

Do not change any other app eval. Each listed ID keeps its existing tests.
