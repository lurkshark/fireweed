# Doctor 0071 — blocked for human review

Controller stopped at `697ddb2` with `evals attempts exhausted`.
0071 remains in evals with audit_rounds 2 and audit_pending CHK-065. No change
was applied, so the post-review application count remains zero. The working
tree is clean and no claim remains. Saved attempt:
`refs/fireweed/attempts/0071-doctor-probes/1` at
`170f87aaa6131e55856a926145483e7e6494b9c0`.

The owner-approved C12/U02 correction and its checkpoint re-review passed and
were pushed at e083c5f before this separate doctor feature began. D007 and the
0071 author change are local commits after that push. No further push occurred.

## Phase evidence and unresolved finding

Evals f733d10 reported 248 total cases: 22 runtime failures, 226 passes.
Audit 86c25f2 found missing assertions for CHK-054/057/060/065. Corrected evals
07db967 reported 254 total cases: 28 runtime failures, 226 passes (65 own checks
cases, with all 37 existing cases passing). Scoped re-audit 75f09c8 retained
only CHK-065; its other prior findings were resolved. The controller then made
the state-only block commit 697ddb2. No implementation phase or archive exists.
The Doctor API/model types and compiling stub remain from the eval phase.

CHK-065 requires request Timeout, or three minutes when unset, to bound each
agent probe. The added unset-timeout test only verifies two 0.2-second commands
can finish. It does not prove the three-minute upper bound. Keep the requirement;
do not narrow it or reset counters to escape the finding.

## Audit input and phase deviation

The first auditor misread its phase, read implementation files and constructed a
provisional Doctor implementation before attempting verification. In the ignored
auxiliary log
`.fireweed/runs/20261028T143328-b8168faa7b2578efc49bb800/audit-1.log`, commands at
lines 3830 and 4740 read checks and dependency implementations; line 5512 writes
doctor.go. Verification rejected the audit-phase disposable edit at line 7958;
the auditor restored the committed stub at line 7963 and committed only findings
through the controller. No provisional implementation survives in the working
tree or phase commits. Do not consult that generated implementation as an oracle.

The source restoration fixes the zone violation, but those input reads exceed
the goal's requirement to supply the separate auditor only the change, projected
specs and evals. Do not claim that first invocation as restricted audit evidence.
The second audit reviewed its four pending IDs and read only the change/context;
it does not independently establish the other seventeen IDs' assertions.

The coordinator observed the temporary diff and inspected the log. An attempted
interrupt found the controller process already gone; no signal was delivered.
The native block is the authoritative stopping state. No ref, attempt bound,
audit bound or approved change body was altered by the coordinator.

## Proposed recovery — not authorized yet

Approve one bounded recovery of this same change through the controller's saved
attempt mechanism, retaining the audit-round history and saved ref. Add a real
default-timeout eval: with request Timeout zero, one local fake command runs
longer than three minutes and the other returns immediately; assert C15 timeout
near the three-minute bound. Keep the own suite under five minutes, without clock
injection, new testing-only API, instrumentation or real-agent authentication.

Because the first audit exceeded its input boundary, also obtain one fresh
restricted audit of all CHK-049–069 from only the change, projected specs, evals,
retained interface and governing audit rules. This is an explicit human process
repair, not an implicit scope or round reset. Any further unresolved finding
stops for human review. Preserve the original logs, findings and counters.
Do not continue implementation or push before this recovery is approved and its
required checks/audit pass. Doctor diagnostic explanations and app wiring remain
planned; no drill or deferred contracts-gap change is authorized.
