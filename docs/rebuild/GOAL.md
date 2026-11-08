# Complete three audited eval corrections, then stop

Read AGENTS.md, SYSTEM.md, PROGRESS.md and the author-change playbook. The owner
requests the three eval/contract mismatch fixes from review of rejected attempt
d63c16e, through normal audited changes preserving behavioral coverage.

Budget: exactly three eval_fix changes, 0185–0187, one each for workflow, runner
and app. Workflow accepts explicit empty-selection wording including "No selected
changes." without losing range/artifact assertions. Runner asserts full context
contents rather than unspecified heading spelling, retaining audit isolation and
lifecycle checks. App accepts optional workflow B04 Path while verifying every
returned diagnostic field, and accepts the compliant empty-selection wording.
No spec, Decision, runtime, interface/model, hook/configuration, worker-policy or
tooling correction. Workers obey context/zones and never commit or transition;
the unchanged pinned controller owns phase verification, independent restricted
audits, commits and archives. Preserve normal approval and enforcement bounds.

Require all own evals in eval_fix, independent audit and full implementation-phase
verification. Review coverage and resulting diffs, run affected static/history
checks, record the handoff idle and stop. Preserve all prior evidence/refs and the
failed workflow proof. No drill, reconstruction import, controller promotion,
expanded milestone, later work or push. Another drill requires new owner approval.
The existing deadline remains 2026-11-10 22:51:05 UTC; no extension.
