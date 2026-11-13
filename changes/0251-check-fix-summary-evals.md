---
eval_fix:
    - APP-147
    - APP-148
module: app
phase: impl
run: 20261113T175834-8a5aa5595c0fd3b923f98df5
---

## Motivation

After regenerate-app correctly stamped missing bases, its remaining
check --fix failure compared plain "Check passed." with a terminal newline
against the decoded JSON summary without that newline. APP-024 requires a
string summary and APP-148 preserves each output shape; neither requires
cross-format byte equality. The owner approved separating output validation
from the exact file-idempotence assertion.

## Interface

## Added

## Modified

## Removed

## Decisions

## Notes

Audited eval_fix only. Edit only the app evals zone and this change.

Correct Test_APP_147_APP_148_CheckFixMissingOnlySuccessAndNoOp in
packages/app/evals/check_fix_cli_test.go. Validate each plain/JSON success
shape independently, or normalize only the terminal line ending when
comparing their summary text. Keep a meaningful success-summary assertion.

Keep exact before/after file-byte checks for repeat invocation, expected
canonical bases, preservation of human content and unrelated files, and
all stale-base, rejection, exit-code and diagnostic assertions. Do not
normalize file contents or other JSON fields.

Do not change product output or specs. Drafted and queued only. No lifecycle
execution is authorized by the drafting request.
## Audit

No findings; checked APP-147, APP-148

## Provenance

- evals 2026-11-13T18:06:00.743213Z agent=codex fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=369 failing=0 passing=369
- audit 2026-11-13T18:06:19.712762Z agent=codex fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0

