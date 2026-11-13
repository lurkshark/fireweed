---
module: app
eval_fix: [APP-147, APP-148]
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
