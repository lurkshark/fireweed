---
module: runner
phase: applied
run: 20261113T191602-028a559aba4e0132fb3ba856
---

## Motivation

The regenerate-runner correction added a parent-directory permission precheck
so a cleanup eval's teardown could still find Git's worktree registration.
RUN-062 requires cleanup failures to be returned, not preservation of that
registration. Following the owner-approved fixture repair, assess removing
this fixture-driven implementation constraint while preserving actual
cleanup and failure reporting.

## Interface

## Added

## Modified

## Removed

## Decisions

## Notes

No retained interface or behavior changes. Maintenance is limited to runner's
disposable cleanup implementation.

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
## Provenance

- impl 2026-11-13T19:31:09.438314Z agent=codex fireweed=0.1.0 playbook=5f856d6bcab308b9ff1a8a7c03b9d2547744112aba81dd30dea8f61ec8062b03 tests=539 failing=0 passing=539

