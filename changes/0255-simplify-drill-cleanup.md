---
module: runner
---

## Motivation

The regenerate-runner correction added a parent-directory permission precheck
so a cleanup eval's teardown could still find Git's worktree registration.
RUN-062 requires cleanup failures to be returned, not preservation of that
registration. Following the owner-approved fixture repair, assess removing
this fixture-driven implementation constraint while preserving actual
cleanup and failure reporting.

## Interface

No retained interface or behavior changes. Maintenance is limited to runner's
disposable cleanup implementation.

## Added

## Modified

## Removed

## Decisions

## Notes

Run only after drill-cleanup-fixture has applied through its audited eval_fix
lifecycle. Keep this change separate because an eval_fix's impl phase cannot
edit implementation.

Assess packages/runner/internal/drill/drill.go's removeWorktree permission
precheck. Remove it and any newly unused import if ordinary Git worktree
removal still satisfies RUN-062 with the corrected fixture. Preserve returned
cleanup errors, nonpassing results on failure, reconstruction-input cleanup,
native report handling and cause-code behavior.

Use the existing module and required dependent verification through Fireweed.
Do not alter evals, specs, timeouts or enforcement, add another preflight
mechanism, or run a standalone reconstruction drill. If removal reveals a
substantive contract failure, retain the required behavior and record the
concrete finding rather than expanding this maintenance change.

Drafted and queued only. No lifecycle execution is authorized by the drafting
request.
