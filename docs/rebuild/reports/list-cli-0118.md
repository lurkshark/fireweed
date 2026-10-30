# Queue list CLI — corrected 0118

APP-114–117 phases921db95e39d0a937e14f935fa7d8d1d060d7126f /
932dd1e97b8808074a830448060acf136bf850c0 /
a47a7fd96ce08a777a341405a1f091d1d5b46c70. Native evals250 cases3 expected
unsupported-list failures/247 passes (about275s); implementation250 passes
(about245s), both under the unchanged own5m/full15m bounds. Three focused cases
pass. Existing APP-10317/APP-10419 warnings remain explained.

Owner approved /tmp/fireweed-0118-audit-checkpoint.md. Controller abandon78705e1
removed the old blocked eval-only attempt; no implementation had been written.
Corrected author contract drops only the currently unassertable P01 conditional,
preserving actual Queue cause/exit/action requirements. The fresh eval worker
restores all saved PUBLIC assertions, including C03/canonical-path and C13 cases,
and adds a row-specific plain blocked="" assertion. The existing single read-
command convention test gets the two list calls; no new observer or convention spec.
The canceled attempt and its Questions remain historical, not passing evidence.
See [original checkpoint](list-audit-0118.md).

Fresh restricted Sol/high audit PASS on all four IDs, audit_rounds0, run
20261030T063511-92982ff89d0d2ea4ed17a021. Complete commands inspected: only supplied
public change/projectedSPEC/evals/interfaces inside the snapshot; no implementation,
SYSTEM/history/outside/code/tests/writes. Snapshot removed. Fresh implementation
worker used current code/frozen evals; no saved inputs/implementation or old logs.
It changed only the entry dispatcher and internal list implementation, retaining frozen
interfaces/evals. No manual phase/ref/bound writes, new error path/API or allocation.

Clean NEW sourcea47a7fd96ce08a777a341405a1f091d1d5b46c70, Go1.27.2 darwin/arm64,
CGO0/vcs.modified=false. Build passed without warning. Actual isolated NEW proof
PASS from a tracked nested module cwd in a none-recipe project. Fixture baseline
ed25d0e1f35b51d388ea4c9fe6555e4e2499c4dd; actual held run20261030T065026-8b516a56d7c1d41563a75fc8.
Ordered first/second entries show queued/default-empty reason/round0 and an authored
blocked reason/round1. Public NEW start assigns ALP-002 and phaseevals; list shows
that live held state while retaining claim/run/document. Public release clears the
claim without changing the document. Fixture-author removal of pending documents
then yields plain empty output and schema1 JSON []/empty diagnostics. JSON stderr
stays empty. Every list read retains HEAD/refs/pending bytes/status; no hooks or
agents run. Unclaimed fixture/state removed. No lifecycle/drill was repeated.

The fixture replaces the none recipe's empty hooks mapping before execution;
no runtime harness repair or rerun was needed. Authored blocked input is a read
fixture, not a new blocked-lifecycle proof. Earlier lifecycle evidence is reused.

Local /tmp/fireweed-list-recovery-{evals-worker.log,evals-focused.json,
evals-done.log,audit-worker.log,audit-done.log,impl-worker.log,
impl-focused.json,impl-done.log}, /tmp/fireweed-0118-abandon.log and
/tmp/fireweed-list-cli-{result.json,commands.json,run.log}. Public saved inputs
remain /tmp/fireweed-0118-public-inputs for recovery traceability; no implementation.

Applied count is2 (0117/0118) since the passing0112–0116 review. The abandoned
attempt is not an application. Main archive clean/unclaimed before this evidence
commit; no pending rebuild change. Stop here under the approved single-change
consecutive-app recovery scope. No new push, publication, phase exception or drill.
Full design/ACCEPTANCE, fresh-source whole-goal gates and remaining surface stay open.
