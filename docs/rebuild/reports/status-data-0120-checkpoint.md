# Status data — frozen-eval checkpoint

0120 is unapplied, phase impl, run
`20261030T144326-24f855074c2173810865ffc2`, audit_rounds1. Its owned claim was
released natively at the checkpoint. Implementation remains uncommitted in
workflow/internal/lifecycle/eligible_stub.go and status_stub.go; no files were
discarded and no phase was changed by hand. Applied count remains3 (0117–0119).

Author `2b7956a` allocated FLOW-101–113 through the controller. Evals
`5644ae0bc92bfeb1d8b0c7a09de42bc4ea38305e` recorded457 cases:7 intended
failures/450 passes. A restricted independent Sol/high audit found FLOW-113
did not assert beta's presence in latest costs; the controller persisted it in
`e3e0d75063aef539918b75cbc42afbe45fb5579d`, round1/pending113. The scoped
seven-line expected-module correction preserved existing assertions; native
evals `ea3a01c4a3a71117baacd7ca6999e84d871ea799` recorded the same counts.
Restricted FLOW-113-only re-audit PASS was persisted in
`2426c60611fdcf45d33861472f488680eec9fa25`, advancing to impl. Both auditor
command logs stayed within supplied change/projected SPEC/evals/public inputs;
no SYSTEM/private implementation/history/outside access or code execution.
Snapshots removed. Own verification stayed below5m and full below15m.

A fresh fixed Sol/high implementation worker passed6/7 focused cases. The
failure is a wrong frozen FLOW-101 assertion: eligibilityBase authors a change
with no spec edits, so existing FLOW-077 Start records impl. Status correctly
returns phase impl. The test instead expects evals. The worker stopped without
altering frozen tests, misreporting phase, committing or advancing workflow.
Full impl verification/Done was not run; no application or archive followed.
This is an implementation checkpoint, not a native audit block.

SYSTEM.md40/42 and the approved slice require an owner decision. Public inputs
and the exact proposed one-token evals→impl correction are retained in
`/tmp/fireweed-0120-public-inputs`; no implementation is in that snapshot.
`/tmp/fireweed-0120-frozen-phase-checkpoint.md` proposes native cancellation and
fresh re-authoring, with no discarded implementation reuse or phase exception.
It also offers the necessary B01–B03 Explain change before the fifth review,
then immediate app status/NEW proof after PASS. No recovery is approved yet.

Logs remain under `/tmp/fireweed-0120-`: evals-worker, evals-refinement-worker,
evals-done, audit-worker, audit-done, evals-round1-worker, evals-round1-done,
reaudit-worker, reaudit-done, impl-worker and checkpoint-release; focused JSON
results use the same prefix. Actual NEW status proof, app wiring, fifth quality
review/checks and push are unrun. B01–B03 Explain remains an integration gap.
