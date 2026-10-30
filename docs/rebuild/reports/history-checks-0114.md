# History checker dependency — 0114

CHK-070–078 phasesedfeee9/ba4aeb1/76e2ef5/628dd1b/04a65d7: first524 cases20
failures/504 passes, corrected525 cases21 failures/504 passes, final525 passes
at04a65d782819ffe4fbbfbaf3c292a7beb446c6d1. Full impl about728s, under15m.
Real fixed Sol/high workers;21 focused named events pass. Public History facade
and two report types remain thin; behavior is one disposable checks file, using
project.LoadAt and local Git reads. No hooks/agents/temp checkout/ref or dependency.

Independent audit found two CHK-070 gaps: branching topological order and default
non-main divergent-base selection. One new test and one adapted fixture preserve
all prior assertions. Both exact one-line quotes were natively persisted under
same run20261030T025913-b1814486063e4bf2bf4245fa, IDs, round1. Scoped CHK-070-only
re-audit PASS; complete audit command inspection found only allowed public snapshot
reads, no implementation/SYSTEM/history/outside/code execution/writes. Snapshots
removed. No phase-write exception. Prefreeze fixture repairs corrected an absent
archive directory, squash setup range and an invalid protected set to a public-
validated module-relative interface projection; no false frozen assertion retained.

Review caught normal-history failures after focused passes: the implementation
rejected valid queued documents without phase and treated the controller's SPEC
application as H01. Public API calls on recorded inputs4c41880..57f4ed4 and
eed7cc4..ea83a82 reproduced each. A scoped implementation-only correction permits
unassigned phases without classifying authors as workers, and separates intrinsic
SPEC from explicit protected zones. Same inputs then return clean2-commit and
1-commit reports; frozen tests/interface unchanged. No extra probe matrix.

Local logs /tmp/fireweed-0114-{evals-done.log,evals-correction-done.log,
audit-worker.log,reaudit-worker.log,impl-done.log,impl-focused.json,
queued-phase-repro.log,spec-application-repro.log,queued-phase-corrected.log,
spec-application-corrected.log}. Existing B02 CHK-05516/APP-10317/APP-10419 remain
explained. Native archive releases claim and leaves main tree clean.

Count3 since passing/pushed0106–0111 gate. This completes the second consecutive
history dependency; stop before the third history change under SYSTEM.md42.
Concrete proposed remaining assets explanation/app caller/actual NEW proof and
fifth-review scope: /tmp/fireweed-history-continuation-checkpoint.md. No new push,
CLI completion, real drill, final acceptance or goal completion claim.
