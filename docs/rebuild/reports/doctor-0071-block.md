# Doctor 0071 — blocked for human review

Current result: the approved recovery finished, but the fresh restricted audit
reported eight further assertion gaps. Controller block a81b723 leaves 0071 in
audit with audit_rounds 2. No implementation was applied; tree clean and claim
released. The original block and its deviation remain below as history.

Latest state: the owner approved the eight-ID correction, but native saved-attempt
recovery retained audit phase. The agreed stop condition triggered before any
test edit. Native block 4c8439b leaves the tree clean and releases the claim.
An explicit one-time phase-only repair decision is needed; the correction's
scope remains approved.

The original controller run stopped at `697ddb2` with `evals attempts exhausted`.
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

## Approved bounded recovery — 2026-10-28

The owner approved the following plan. Perform one bounded recovery of this same change through the controller's saved
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
Do not continue implementation or push before the required checks/audit pass.
Use separate Sol/high CLI worker and auditor invocations; coordinate phase
advancement with the pinned controller, which alone owns commits. Supply the
fresh auditor an isolated directory containing only its permitted inputs and
inspect its invocation log before accepting the result. No global audit bounds,
attempt limits or product requirements are changed by this recovery.
Doctor diagnostic explanations and app wiring remain
planned; no drill or deferred contracts-gap change is authorized.

## Recovery outcome — further human review required

Author approval commit c3c895e records the owner's decision. Native start with
from-attempt 1 initially refused the existing blocked marker. The design's
Resuming/state-fields clauses explicitly require removing that marker; author
commit 58cf943 removes only that line. Native start with from-attempt 1 then
succeeded, preserving run, IDs, audit_rounds 2, findings and the original ref.

The separate Sol/high eval worker added one 22-line CHK-065 regression. With
Timeout zero, its worker returns immediately and its audit command sleeps 210
seconds; assertions require a C15 failure naming that command/output and elapsed
time between 165 and 200 seconds. Existing assertions remain. The stub fails at
runtime, and coordinator-driven native done committed evals d78d618: 255 total
cases, 29 failures, 226 passes (66 own cases, 37 existing passes). Provenance
honestly says agent=manual for coordinator advancement; the actual Sol/high CLI
worker is recorded in recovery-evals-1.log under the existing run directory.
The first CLI launch failed during sandbox initialization before worker execution;
the same command succeeded with the controller's parent execution permissions.

The fresh Sol/high auditor received a separate temporary Git input directory
with only the current change, projected checks SPEC, checks evals, retained
public interfaces and governing audit rules. Its complete invocation log
`recovery-fresh-audit-1.log` shows only reads of those supplied files, with no
implementation/history reads, file edits or code/test execution. This is a
valid restricted input review; unlike the original invocation, it did not
construct an implementation. It checked CHK-049–069 and no longer reports
CHK-065. This closes the assertion finding, not proof of implemented timeout
behavior; implementation verification has not run.

The fresh review reports these remaining assertions missing in doctor_test.go:

- CHK-049 (line 238): retaining warning diagnostics and suppressing boundary/audit probes after a static error.
- CHK-050 (line 294): placing scratch below the actual common Git directory, beyond a path-substring check.
- CHK-051 (line 300): deleting the physical scratch directory, beyond registration cleanup.
- CHK-054 (line 325): naming the particular missing template or baseline map in C14.
- CHK-055 (lines 383/398): failed baseline eval/boundary and recovery eval/coverage paths.
- CHK-063 (lines 521/553): the audit role's fixed-reply prompt and wrong-reply rejection.
- CHK-066 (line 543): rejecting a limit-pattern match outside the final 20 output lines.
- CHK-068 (line 620): accepting current managed guidance with surrounding text.

The coordinator directly confirmed the common-directory, physical-cleanup and
surrounding-text gaps against their explicit clauses. Some other findings ask
for additional combinations; human scoping should retain SYSTEM's representative
assertion and cost rules, without treating this as permission for a Cartesian
matrix or unbounded audit loop.

Per the approved plan's stop condition, these findings became Questions and a
blocked reason in the current change. Native done made state-only block a81b723
and released the claim; no audit scope field/counter was reset. Original attempt
ref 1 remains, and native block retained ref 2 at
`8413bab317def40d1770682067fb25fe4bdd0323`. Do not inspect saved implementation
attempts as oracles. No further correction, implementation, feature or push ran.

## Decision needed

Any further recovery needs explicit approval: scope the missing assertions under
the existing guardrails, resume this same change through the controller, retain
all prior passes and the completed default-timeout test, and re-audit only the
outstanding approved IDs from a restricted input snapshot. Preserve counters,
requirements and history. If the controller cannot restore an eval step without
changing bounds, stop before editing frozen evals. No such further recovery is
authorized by the approval already consumed above.

## Further bounded correction approved — 2026-10-28

The owner approved one correction for the eight outstanding IDs listed above.
Preserve all prior passes and the completed CHK-065 default-timeout regression;
re-review only CHK-049/050/051/054/055/063/066/068 from restricted inputs. Retain
audit_rounds 2, IDs/run, requirements and both saved refs. No replacement change,
bound change or push is authorized. Use the controller's saved eval attempt to
restore the correct step; if it cannot restore evals without changing bounds,
stop before editing frozen tests. Any further unresolved finding requires human
review rather than another automatic correction round.

## Native phase-restoration limit — human decision needed

Approval was recorded in 9cbe777. Author commit 3d592f5 removes the block and
answered Questions without changing phase, counters, specs, run/IDs or prior
passes. The saved attempt's change metadata names phase evals, but native
`start doctor-probes --from-attempt 1` retained the current committed audit phase.
It loaded an attempt into the working tree without rewinding the workflow step.
No file diff remained and the completed default-timeout regression was retained.

Because evals are frozen during audit, the coordinator stopped before editing any
test/interface and launched no worker or auditor. Native done committed the
blocked reason and Questions at 4c8439b, released the claim, and retained another
attempt at `refs/fireweed/attempts/0071-doctor-probes/3`:
`2ead5ed35ab48091db840ccdefd5ab49ed26a327`. Original refs 1/2 remain. audit_rounds
is still 2 and the native pending field remains CHK-065. No bound, count, approved
requirement or saved ref was reset, and no implementation or push ran.

The scoped correction itself is approved. The additional decision required is a
one-time explicit exception to controller-owned phase writes: allow the author
to change only phase audit to evals for this same change, in a separately recorded
repair commit, then resume normal controller ownership. Preserve run/IDs,
audit_rounds 2, audit_pending, requirements, findings/history, prior passes and the
real default-timeout test. No global configuration/bound changes or replacement
change are proposed. Only after that decision may the already scoped eight-ID
correction and restricted re-review run; further unresolved findings still stop
for human review. This exception has not been performed or inferred from the
earlier approval.
