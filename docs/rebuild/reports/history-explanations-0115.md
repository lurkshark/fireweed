# History explanations — 0115

AST-034–037 phases e2c8d3b/233d592/131ef12/3a7b556/784db27/9f5feb1/f98b3b3.
Initial and corrected native evals:680 cases,4 expected unknown-H-code failures,
676 passes; final implementation680 passes atf98b3b30f381533cc631ce08f4da1d04ad72679c.
Full implementation about239s, under15m. Four focused public Explain tests pass.
Only four assets/internal lookup strings added; interfaces/models/evals frozen.

Independent restricted audit identified AST-035's both-edit-types action and
AST-036/037's new-data/file-permission relationships. First correction retained
all assertions; scoped re-audit passed035 but retained036/037. Second, simpler
sentence matchers assert affirmative new-file-for-new-data and archived-file-
deletion permission. Final isolated036/037-only re-audit PASS. Two findings
round trips persisted normally, same run20261030T044031-868ada5b81566b0bb0c4bb59,
IDs and audit_rounds2. Complete auditor commands inspected: only allowed snapshot
public-input reads, no implementation/SYSTEM/history/outside/code/tests/writes.
All snapshots removed. No raw phase write or bound adjustment.

Local logs /tmp/fireweed-0115-{evals-done.log,audit-worker.log,
evals-correction-done.log,reaudit-worker.log,round2-evals-done.log,
final-reaudit-worker.log,final-audit-done.log,impl-worker.log,impl-focused.json,
impl-done.log}. Native archive released claim; clean tree before caller authoring.
Existing B02 CHK-05516/APP-10317/APP-10419 remain explained. This is application4,
not CLI slice completion or overall acceptance. Owner-approved caller and actual
NEW proof follow, then fifth-change quality/goal review/checks and push only PASS.
