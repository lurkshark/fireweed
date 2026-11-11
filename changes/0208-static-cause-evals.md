---
eval_fix:
    - APP-009
    - APP-037
    - APP-116
    - APP-117
module: app
phase: impl
run: 20261111T040654-c17f59f3929080416a3a2c83
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
## Provenance
- evals 2026-11-11T04:19:01.197674Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=323 failing=0 passing=323
- audit 2026-11-11T04:19:27.481418Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-astra -c model_reasoning_effort="medium" -s workspace-write {prompt} fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
## Audit

No findings; checked APP-009, APP-037, APP-116, APP-117
