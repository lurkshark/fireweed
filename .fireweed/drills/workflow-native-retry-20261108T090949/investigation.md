# Investigation of the 17 native failure rows

The owner authorized investigation and retained findings only. This is a static
review of the [native report](../workflow.md), its exact streams/receipts, the
retained reconstructed workflow, and matching public contracts/evals. No build,
eval replay, new agent, drill, correction, source application or repin was run.
The original native failure and WF4/full-completion stop remain unchanged.

The source inspected is exclusively the reconstruction at
`refs/fireweed/drills/workflow/20261108T090949.204125000`, OID
`e0dc010ce1b5a06cdb4f4d87a290b7bbb5016475`. Its candidate is a26f7b9 and setup
parent is 2e515d1. References below to `reconstructed <file>:<line>` mean that
ref's `packages/workflow/internal/lifecycle/<file>`, readable with `git show`;
they do not refer to the product implementation. No bootstrap, rebuild-v1,
deleted original implementation or dependent internal implementation was read.

**Conclusion:** 14 failed rows trace to three reconstruction defects. Three
runner rows expose two contract/eval gaps: an end-to-end whitespace-prefix
assertion, and an unstated assumption that dependents are evaluated after own
checks already fail. They do not establish three additional runner defects.
These classifications are findings for owner review, not permission to weaken
assertions or change contracts. No native failure is waived.

## Complete accounting

The native report contains 17 failing rows, including parent/subtest rows.
The table accounts for each exactly once; it is not a count of independent bugs.
Names below follow the native Go testcase spelling. Report line numbers refer
to the unchanged current report, whose full original bytes are preserved.

| Report line | Module / testcase | Group | Classification |
| ---: | --- | --- | --- |
| 69 | workflow: Test_FLOW_128_GCPrunesMissingAndUnregisteredClaimsOnly | A | Reconstruction defect |
| 82 | workflow: Test_FLOW_129_GCPrunesOnlyArchivedChangeAttempts | A | Reconstruction defect |
| 95 | workflow: Test_FLOW_130_FLOW_131_GCDrillAgeUsesConfiguredDurationAndNativeUTCDate | A | Reconstruction defect |
| 108 | workflow: Test_FLOW_130_FLOW_131_GCDrillAgeUsesConfiguredDurationAndNativeUTCDate/configured | A | Reconstruction defect |
| 121 | workflow: Test_FLOW_130_FLOW_131_GCDrillAgeUsesConfiguredDurationAndNativeUTCDate/default | A | Reconstruction defect |
| 138 | workflow: Test_FLOW_132_GCReportsOnlyDeletedRefsAndPreservesRepositoryState | A | Reconstruction defect |
| 198 | workflow: Test_FLOW_062_LifecyclePreservesEvalZoneAndBoundaryCauses | B | Reconstruction defect |
| 302 | app: Test_APP_014_MaintenanceCommandsDispatch | A | Reconstruction consequence / parent row |
| 305 | app: Test_APP_014_MaintenanceCommandsDispatch/gc | A | Reconstruction consequence |
| 403 | app: Test_APP_150_APP_151_GCNearestProjectSortedAndEmpty | A | Reconstruction consequence / parent row |
| 406 | app: Test_APP_150_APP_151_GCNearestProjectSortedAndEmpty/plain | A | Reconstruction consequence |
| 410 | app: Test_APP_150_APP_151_GCNearestProjectSortedAndEmpty/JSON | A | Reconstruction consequence |
| 445 | app: Test_APP_010_APP_029_APP_030_APP_031_VerifyDoneAuditAndFullChecks | C | Reconstruction defect |
| 450 | app: Test_APP_010_APP_029_APP_031_DoneReportsBlockedAudit | C | Reconstruction consequence |
| 632 | runner: Test_RUN_047_RUN_048_FirstAuditBeforeProvenanceIsImported | D | Eval overconstraint / attribution gap |
| 742 | runner: Test_RUN_043_RetryIncludesJUnitOnlyCasesFromOwnAndDependentModules | E | Cross-module spec/eval assumption gap |
| 750 | runner: Test_RUN_044_RetryBoundsUnicodeFailureExcerptsAndMarksOmissions | E | Same assumption gap |

## A — GC assumes an unsupported Git option (11 rows)

Every direct failure in this group reports the same concrete cause: installed
Git rejects `worktree list --porcelain -z` with exit 129, `unknown switch 'z'`.
Reconstructed `gc.go:58–60` unconditionally executes that command and returns
its failure before the pruning loops. Lines 62–66 then assume NUL-delimited
output. This is the reconstruction's command choice, not a command imposed by
any public input or eval.

[FLOW-128–132](../../../packages/workflow/SPEC.md) state claim/attempt/drill
pruning and preservation behavior. [APP-014/150/151](../../../packages/app/SPEC.md)
require maintenance dispatch and successful GC output for valid fixture projects.
The [GC evals](../../../packages/workflow/evals/gc_test.go) create ordinary local
repos/refs and invoke public GC; the [CLI evals](../../../packages/app/evals/gc_cli_test.go)
do likewise through the composed executable. The fixture itself uses supported
`git worktree list --porcelain` at gc_test.go:92. No special clone format,
network, injected Git error or timing assumption accounts for this failure.
SYSTEM requires the supplied macOS/Linux environment; no stated product
prerequisite authorizes rejecting this Git installation.

Classification: **one reconstruction compatibility defect**, cascading through
six workflow rows and five app rows. The app diagnostics are carrying the
workflow cleanup failure; there is no evidence here of a separate CLI dispatch
or JSON defect. Pruning/preservation assertions after the failed call were not
reached, so their eventual behavior is not established. No Git-version spec or
eval correction is justified by these observations.

## B — Boundary violation detail is lost from the returned cause (1 row)

The native FLOW-062 result is `X01: 1 boundary violation(s)`; the assertion at
[lifecycle_test.go:1316–1317](../../../packages/workflow/evals/lifecycle_test.go)
requires X01 and the fixture's `forbidden dependency` detail. The fixture boundary
hook emits the module-attributed violation `alpha: forbidden dependency`.

Reconstructed `verify.go:222–233` retains structured boundary diagnostics but
chooses the aggregate returned error as its sole message. If coded, it returns
that error immediately; otherwise it prefixes the first hard diagnostic's Code
onto the same aggregate, without its violation Message. Thus code/severity
survive, while the specific violation disappears from the returned error.

[FLOW-062](../../../packages/workflow/SPEC.md) requires preservation of assigned
boundary causes. The already-reviewed X01-detail interpretation is retained in
the earlier [failure investigation](../workflow-failure-investigation-20261108/findings.md),
and [0175](../../../changes/archive/0175-preserve-x02-boundary-cause.md) expressly
preserves X01 details while distinguishing X02's sole hook-error cause. That
historical conclusion is reused, not re-executed or used as implementation input.

Classification: **reconstruction cause-composition defect**, not evidence that
the existing FLOW-062 eval should drop its detail assertion. The structured
Diagnostics field is not alleged to be lost. No checks-module correction or
new diagnostic allocation is established.

## C — A permitted first Audit section is rejected as author editing (2 rows)

Both CLI fixtures start without an Audit heading. At
[lifecycle_cli_test.go:431 and :530](../../../packages/app/evals/lifecycle_cli_test.go),
they insert a first Audit section before Decisions while preserving the existing
fields and sections. One supplies checked No findings; the other supplies a
valid untestable finding. The first fails with `P03: approved author metadata
changed`. The second leaves the pending document at audit without a committed
blocked reason; its expected blocking transition was never reached.

Reconstructed `core.go:398–403` compares the entire approved and current bodies
after `ReplaceSection(name, "")` for Audit, Decisions and Questions. This is not
removal of section structure: the retained [CON-009](../../../packages/contracts/SPEC.md)
contract says ReplaceSection appends a missing section. On the approved document
it therefore appends an empty Audit heading at the tail; on the submitted
document it keeps an empty Audit heading before Decisions. The resulting section
order differs, so whole-body comparison rejects an otherwise permitted edit.
This is a static explanation from the exact code and public section semantics;
no fresh fixture execution was needed.

The design explicitly allows the auditor to write an Audit section.
[FLOW-036/040/041](../../../packages/workflow/SPEC.md) require the checked passing
and untestable blocking outcomes; FLOW-020/057 reject *unapproved author* edits,
not the controller's permitted Audit submission. A special heading position is
not a stated precondition. [0139](../../../changes/archive/0139-audit-section-import.md)
also records the prior explicit rejection of a special-position requirement for
audit import, although the present fault is workflow's comparison, not runner's
restricted workspace import.

Classification: **one reconstruction Audit-body validation defect**, with two
CLI consequences. The second failure does not establish an additional defect
in round limits, blocking commits, Questions or claim release, because rejection
precedes that code. Neither eval needs a new behavior or extra worker input to
make its valid submission legitimate.

## D — End-to-end raw Provenance prefix is stricter than audit-copy semantics (1 row)

The native RUN-047/048 testcase completed the lifecycle successfully with three
steps. Its only recorded failure is
[audit_workspace_test.go:254–257](../../../packages/runner/evals/audit_workspace_test.go):
`strings.HasPrefix(afterProvenance, beforeProvenance)` on the **final archive**.
The reported after value contains evals, audit and impl provenance lines; no
non-Audit author-section, Audit-import or lifecycle-completion assertion failed.

Reconstructed `done.go:169` writes each Provenance body using
`"\n" + strings.TrimSpace(oldBody) + "\n" + newLine + "\n"`.
For the initial empty body that supplies two leading newlines before the evals
line; the later nonempty case supplies one. The exact-body [CON-002/009 public
evals](../../../packages/contracts/evals/contracts_test.go) make that raw-body
comparison meaningful. The native after value starts with one newline. This
explains the prefix failure through Done's whitespace normalization while
preserving the existing nonempty provenance entry. It does not establish that
runner's Audit copy deleted a provenance entry.

[RUN-047/048](../../../packages/runner/SPEC.md) constrain what audit copying
itself may edit. The assertion runs after two subsequent successful Done calls,
which are required to append provenance by [FLOW-051](../../../packages/workflow/SPEC.md).
The design specifies appended record lines and their fields, and the retained
Decision specifies line syntax; neither specifies a blank-line byte-prefix
invariant across those Done calls. Unchanged runner provenance-copy behavior
cannot be rejected solely because this final formatting prefix differs.

Classification: **eval overconstraint and attribution gap**, not demonstrated
RUN-047/048 content loss. There is an observable reconstruction formatting
difference, but the observed failure does not prove loss of any required record
or a violation of the line-level provenance contract. If the owner intends raw
Provenance byte-prefix preservation as part of “append,” that is a contract
interpretation checkpoint; this investigation neither decides it by changing
specs nor relaxes the current assertion. The original native failure remains.

## E — Retry evals assume reports that Verify need not have produced (2 rows)

The [retry fixtures](../../../packages/runner/evals/retry_failure_details_test.go)
make the first implementation's alpha and aardvark hooks both *capable* of failing.
Both tests require aardvark's failed case in the retry prompt even though alpha's
own evals already fail. The native RUN-043 section contains the two failed alpha
cases in report order and no aardvark case. RUN-044's only recorded assertion
failure is the missing lexically-first dependent case; the native evidence does
not show a Unicode-limit, excerpt-size or omission-marker violation.

Reconstructed `verify.go:183–185` returns immediately after own RequirePass
fails. Dependent selection/evaluation is below that return at lines 189–218.
Consequently this invocation reports alpha only; runner cannot render an
unproduced aardvark report. This branch plus the native prompt explains both
failures. There is no need to inspect runner internals or infer a sort bug.

[RUN-043](../../../packages/runner/SPEC.md) says **every reported module**,
including dependents, and RUN-044 bounds that section. They do not require Next
to evaluate a module that Verify never reported. [FLOW-025/026](../../../packages/workflow/SPEC.md)
select dependent modules and require rejection of a selected dependent's failed
case; they do not explicitly require continuing evaluation after an own-module
failure, returning all selected reports on such a failure, or a particular
failure-collection order. The design requires all relevant checks to pass for
success, which does not resolve this failure-path continuation policy.
The [0136 author record](../../../changes/archive/0136-retry-failure-details.md)
likewise describes consuming failures already produced and adds no workflow
contract or implementation change.

Classification: **cross-module spec/eval assumption gap**. The fixtures rely on
failure-path collection behavior not clearly promised by the retained workflow
contract, while the runner clauses are conditional on reported data. This does
not prove a RUN-043/044 formatter defect, nor does it certify the fail-fast choice
as the intended product policy. Whether full verification must collect selected
dependent reports after own failure is an owner decision before any clarification
or eval change. Coverage and assertions must not be silently removed to force a
passing proof.

## Preservation and stop

[Investigation validation](investigation-validation.json) records exact native
prefix/stream integrity, unchanged proof ref, matching public-file identities,
idle status and final read/static checks. No runtime/source/spec/eval/interface,
hook, configuration, inspector, binary or recovery file was changed. No temporary
clone or reconstruction workspace was created; no agents or evals were invoked.
Only this evidence and the restart handoff are authored and committed.

The 17 failed rows remain the authoritative native outcome. These findings make
no claim of a successful workflow proof or full completion and do not resolve
earlier timing exceptions. No fix, new change, replay, drill, repin or push is
authorized by the findings. Investigation scope is complete; stop for the owner.
