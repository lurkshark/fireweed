---
module: workflow
bases: {}
eval_fix:
  - FLOW-002
---

## Motivation

The owner approved the two focused follow-ups to the failed WF3 workflow drill.
FLOW-002 requires Verify to reject forbidden edits under the approved projected
phase zones before hooks. The reconstruction returned from audit submission
checking before checking file zones. Current workflow already calls Zones
before its audit branch, but retained workflow zone evals assert evals/impl
rather than the audit integration. Add the missing public assertion under the
normal audited eval_fix lifecycle, preserving the existing obligation.

## Interface

No public signature, runtime source, spec, Decision, hook or configuration edit.
Only workflow public evals and this change document may change during evals;
implementation changes nothing except this change document.

## Added

## Modified

## Removed

## Decisions

## Notes

Add one focused public FLOW-002 test to the existing lifecycle/audit evals.
Use the established fixture and toAudit helper with a valid complete No findings
submission. Introduce a representative forbidden edit outside the active change
file (for example a target-module eval file) after the audit phase base.
Require quick and full workflow Verify and Done to reject it with Z01 naming
the edited path. Confirm the active audit phase/claim and HEAD do not advance,
and no eval hook runs. Reuse the existing fixture hook marker; its evals Done
setup invocation may run, so clear that prior marker before the audit checks.
Preserve the valid current-change Audit submission and every existing test,
assertion and covered ID. One representative forbidden path is enough to
assert workflow's audit call to Zones; reuse checks' valid independent proof
for its zone classification instead of duplicating a zone matrix.

All current workflow evals must pass, with FLOW-002 still covered. This is an
eval_fix of missing assertions for already-correct behavior; do not manufacture
fail-first by changing runtime code, add a shadow implementation, inspect the
failed reconstruction, add instrumentation or weaken existing tests. Do not
read bootstrap/rebuild-v1 or touch other modules/protected infrastructure.

Workers use pinned fireweed context before each held step, verify with pinned
fireweed verify --quick, poll to terminal and return. They never commit,
invoke lifecycle mutators, edit run/phase/Provenance, or edit protected planning.
Only the outside controller advances, commits, applies and archives with full
verification. Preserve frozen evals in impl, independent restricted auditing,
all attempt/audit counters, pin, bounds and the original deadline.

The independent restricted auditor edits the Audit body only. A passing audit
says No findings and Checked FLOW-002. Findings quote an exact normalized
FLOW-002 clause without an asserting test. Re-audit only IDs with previous
findings; preserve audit rounds and the existing configured limits.

This follow-up is limited to this eval_fix and the separate empty-Current
contract change. After both apply, validate the clean identified candidate,
commit evidence and stop at the exact-candidate WF2 approval gate before
another native workflow invocation. Preserve all prior reports/proof refs and
unrelated worktrees. No source import, controller promotion or push.
