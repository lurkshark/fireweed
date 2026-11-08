# Investigation of the four corrected-source drill failures

> Retention update: this review describes the evidence at `d63c16e`.
> Superseded attachments are recoverable from [Git history](../README.md#historical-evidence);
> current module reports and selected outcome receipts remain in this checkout.

The owner authorized investigation only. The subject and controller remain
ccd3d6b7732d613e80a027ba7969b65adc58bc0e; original 79cf33a remains preserved.
The failed reconstruction remains at
refs/fireweed/drills/workflow/20261107T232508.801610000, object
21039de866b8f03d9803b6fe7f80e744e8ad55c9. No product, spec, retained eval, hook,
controller, worker input or configuration correction was authored. No repository
native workflow CLI drill or real reconstruction agent was invoked. Existing Go
runner-suite tests do exercise Drill in temporary fixtures with bounded fake
workers; that is not a new workflow reconstruction/proof invocation.

## Findings

Three native workflow failures reproduce on the saved reconstruction and pass on
the unchanged candidate. The runner failure did not reproduce, including in one
full unchanged runner-suite replay on the failed ref. Native failures and the
original report remain valid evidence; passing replays do not rewrite that result.

| Native failure | Finding | Classification |
| --- | --- | --- |
| AuditFindingsConvergeAndPersistTwoRoundBound (FLOW-035/037/038/039/042/043/046) | The reconstruction requests fail-first for review-tagged Added ALP-003. Its generic changedIDs includes every Added/Modified item and passes those strings directly to RequireFailFirst. | Confirmed reconstruction defect in FLOW-023 enforcement; the audit behavior is never reached. |
| ImplDoneAppliesSpecsSetAndDecisionsAndArchives (FLOW-004/028/044/045/046/047/050) | In addition to that review-ID error, the reconstruction allocates ALP-002 to the new handle while removing the existing ALP-002 in the same change. It gathers used IDs from the already-projected spec, after the removal, and does not reserve Removed IDs. | Confirmed reconstruction allocation defect (FLOW-007), plus review fail-first defect. The impl application behavior is never reached. |
| LifecyclePreservesEvalZoneAndBoundaryCauses (FLOW-062) | checks.Boundaries returns aggregate error `1 boundary violation(s)` and detailed X01 diagnostic `alpha: forbidden dependency`. The reconstruction chooses the aggregate error first and merely prefixes X01, dropping the diagnostic text. | Confirmed reconstruction cause-loss defect; existing FLOW-062 requires preserving the cause. |
| VerificationDeadlineRetainsHeldWork (RUN-037) | Original native output reports only a missing hook-created time-work.txt. Focused replays on both sources and the full failed-ref runner suite pass. | Unreproduced native failure; no demonstrated retention-loss or runner implementation defect. A timing/phase-entry precondition is plausible, not established as the native cause. |

These are not compiler or import failures. Both original attempts compiled and
produced case-level results. Supplying module imports would not repair the
observed allocation, requested-ID or error-selection behavior.

## Bounded comparison and public probes

[Ordered replay commands/results](results.json) retain exact commits, fixed bounds,
timestamps, exits and complete stdout/stderr. Only the three named workflow tests
and the named RUN-037 testcase were selected first. Candidate workflow passes in
15.164s; the failed-ref workflow reproduces all three failures in 7.262s.
The original RUN-037 case passes on candidate (2.913s package time) and failed
ref (2.868s). No retained assertion was edited for these comparisons.

Because RUN-037 originally failed within a package run, one full failed-ref
runner replay followed, bounded by Go's 8m timeout and an external 600s limit.
It passed all 159 cases, no failures/skips, in 263.121s. Its
[result](../README.md#historical-evidence) (`d63c16e:.fireweed/drills/workflow-failure-investigation-20261108/full-runner-result.json`) and complete output are retained. No further
full-suite repetition or stress/load simulation was performed.

Temporary public-API probes reused the existing ordinary fixtures and fake
hooks/workers; they imported no private implementation package. Their complete
sources are [workflow-probe.txt](../README.md#historical-evidence) (`d63c16e:.fireweed/drills/workflow-failure-investigation-20261108/workflow-probe.txt`) and
[runner-probe.txt](../README.md#historical-evidence) (`d63c16e:.fireweed/drills/workflow-failure-investigation-20261108/runner-probe.txt`), with [commands](probe-results.json) and
[structured observations](observations.json). Probe files were added only to the
owned temporary clones, then removed. Original tracked tests/source stayed intact.
Every held probe claim was released through public Release.

The workflow probe separates allocation from downstream assertions:

- Without removal, both allocate ALP-003 with tag review. Direct public coverage
  succeeds and fail-first for the non-review ALP-001 succeeds on both. Candidate
  Verify succeeds; reconstructed Verify rejects missing fail-first for ALP-003.
- With removal of ALP-002, candidate allocates ALP-003, reconstruction reuses
  ALP-002. Both projected Added items are review-tagged, and direct coverage plus
  non-review fail-first still succeed. Reconstructed Verify rejects ALP-002.
- Direct boundary error/diagnostics are identical on both sources. Candidate
  workflow returns `X01: alpha: forbidden dependency: 1 boundary violation(s)`;
  reconstruction returns `X01: 1 boundary violation(s)`.

Thus the worker's FLOW-022/FLOW-023 gap claims are not supported as retained-eval
conflicts. The [design](../../../fireweed-design-v1.md) says review specs need no
evals and removed IDs are never reused. [CHK-002](../../../packages/checks/SPEC.md)
explicitly excludes review from coverage, while RequireFailFirst receives only an
ID list and cannot infer tags from those strings. Workflow must select the right
requested IDs. In the implementation fixture the extra review item is not
supposed to reuse the removed enforced item's identity. The original assertions
should be preserved; changing the checker or adding fake failing cases for review
items would hide the reconstruction defects.

## What RUN-037 establishes and what remains unknown

The unchanged testcase starts the held change before Run, then gives the whole
Run invocation a two-second For budget. The verification hook creates time-work.txt
only after context preparation, agent invocation and verification setup have
finished. It then sleeps, so the test expects interruption during that stage.
The native failure occurs in a fatal read of that file before the later assertion
for the agent's added_test.go executes. The native output therefore does not
establish whether the hook was reached or whether the agent's work was missing.
Earlier held-state checks and clean time-stop classification reported no error.

The public diagnostic using the original two-second limit reached verification
on both sources, retained both agent work and hook marker, kept the held identity,
returned a clean two-second time stop and applied nothing. A separate **1ns
counterexample**, also on both sources, returned a valid clean time stop before
any agent/hook ran: no marker existed, no work had been created and the held
identity was unchanged. That smaller diagnostic limit was not a change to the
retained test, configured bounds or native invocation, and it does not reproduce
the original two-second failure. It establishes only that a missing hook marker
can mean the budget expired before work existed rather than work being lost.

No filesystem/process instrumentation, command proxy, clock injection or extra
worker input was used. We did not force scheduler load, raise the test deadline,
skip coverage or change production behavior. The original native failure remains
unexplained at that level; calling it a proven flaky test or a proven data-loss
bug would go beyond the evidence.

## Scoped follow-ups for owner consideration

No changes are applied or queued by this investigation.

1. If improving the next reconstruction inputs is desired, one normal workflow
   already_satisfied clarification could make FLOW-007's reserved-ID rule and
   FLOW-023's review-item fail-first exclusion explicit. Separate focused public
   assertions under those IDs should reach allocation/enforcement independently,
   preserving the existing audit/apply assertions. Candidate already satisfies
   these behaviors. This would be contract/eval authoring through normal audit,
   not a patch imported from the failed reconstruction or extra worker inputs.
2. Keep FLOW-062 and its cause-text assertion. They correctly catch the observed
   reconstruction defect; no main-product correction is demonstrated here.
3. Do not raise bounds or propose a RUN-037 runtime fix from this evidence.
   A narrow audited diagnostic improvement could avoid the early fatal read so a
   future failure reports the existing agent invocation log, agent work, hook
   marker and held identity separately, retaining the two-second budget and all
   assertions. Deterministic verification-stage targeting would need a concrete
   design/proposal before authoring; it is not silently solved by this probe.

These are proposals only. No native retry, controller update, retained-input
correction or reconstructed-source application is authorized by the investigation.

## Preservation and stop

The [entry](../README.md#historical-evidence) (`d63c16e:.fireweed/drills/workflow-failure-investigation-20261108/entry.json`), [cleanup](cleanup.json), integrity receipt and author checks
retain the source/ref/pin/worktree identities and bounded commands. The candidate,
primary pin and recoveries, all prior reports/archive/proof/attempt refs,
configuration/bounds/counters, existing deadline and unrelated worktree are
unchanged. Old native report and review were not edited. Owned clones were clean
and claim-free before removal; all top-level replay/probe processes returned.
Only this protected investigation evidence and restart state are committed.
No normal change is queued or claimed. Stop for owner review; no fix, new drill,
repin, reconstruction import, publication or push followed. WF4 remains open.
