# Corrected-source workflow drill: terminal review

The owner's requested sequence is complete: two normal audited fixes, clean-source
validation, an idle verified controller update, then exactly one native workflow
invocation on corrected subject `ccd3d6b7732d613e80a027ba7969b65adc58bc0e`. It **failed native verification** after its
two configured workers. No reconstruction was applied and no additional correction,
diagnostic replay, drill or push followed. The approval is consumed; stop here.

## Subject, controller and preparation

[0181](../../../changes/archive/0181-recorded-regenerate-preservation.md) applied at
fd7ae49 and [0182](../../../changes/archive/0182-recovery-history-assertions.md) applied
at ccd3d6b. Their independent audits, phase reports, actual pause/rejection history,
cost limitation, clean normal/race/vet/check/doctor/history/CI and compiled CLI
validation are in [the verified handoff](../residual-fixes-controller-handoff.md),
committed separately at 3853f17. The original 79cf33a remains a preserved baseline.
The initial CHK-083 normal-suite timing failure, passing focused replay and single
passing ordinary retry are retained there; no cause or bound change was invented.

The current controller source is that corrected subject and SHA-256 is `20278504a8ea89fcd5f23d1de9a6ec849389fd55707054992c59967eec65e2b1`.
It stayed unchanged throughout this native invocation. Its retiring 12c571c copy
and every older recovery executable remain preserved. The outside inspector's
candidate/pin metadata was updated only after both normal claims were released.
Original worker command/model/effort, worker SYSTEM instructions, all configuration
bounds/counters and the existing 2026-11-10 22:51:05 UTC deadline stayed intact.

The clean isolated setup child is `616d24f282692d5eb151a8e8fd45b91f6bcf42a9`, parented directly on the
corrected subject. It changes only SYSTEM and agent.wrap; no product, retained
spec/eval, dependency interface, hook or manifest correction is hidden in setup.
The invoker and both scratch controller copies matched the current pin.

## Native result and case review

The [exact native report](native-report.md) records date
2026-11-07T23:25:08.80161Z, that setup base, agent codex, model gpt-6-sol and result
fail. The native controller exited 1 normally, with rejection/emergency false.
[Case review](native-case-review.json) checks every case identity/ID/result and
full published failure text/module output. There are 650 unique cases: 646 pass,
4 fail, no skips or unknown/duplicate entries. Workflow is 165 pass/3 fail, app
323 pass, runner 158 pass/1 fail. All 355 enforced IDs are represented: workflow
133, app 155, runner 67. Coverage does not turn failed cases into a passing proof.

The focused recorded-phase FLOW-081 test, separate FLOW-083 Context test and
RUN-027/RUN-028 recovery history test all pass in the final native report.
The four actual failures are:

| Module | Native testcase | Published assertion/cause |
| --- | --- | --- |
| workflow | Test_FLOW_035_FLOW_037_FLOW_038_FLOW_039_FLOW_042_FLOW_043_FLOW_046_AuditFindingsConvergeAndPersistTwoRoundBound | lifecycle_test.go:735: evals Done rejects missing fail-first for review-tagged ALP-003. |
| workflow | Test_FLOW_004_FLOW_028_FLOW_044_FLOW_045_FLOW_046_FLOW_047_FLOW_050_ImplDoneAppliesSpecsSetAndDecisionsAndArchives | lifecycle_test.go:835: evals transition rejects missing fail-first for ALP-002; published State shows that ID both newly Added with review and Removed. |
| workflow | Test_FLOW_062_LifecyclePreservesEvalZoneAndBoundaryCauses | lifecycle_test.go:1253: boundary cause is `X01: 1 boundary violation(s)`, losing the expected underlying forbidden-dependency text. |
| runner | Test_RUN_037_VerificationDeadlineRetainsHeldWork | time_limit_test.go:182: held `packages/alpha/evals/time-work.txt` cannot be read because it is missing. No timeout root cause is inferred from that message alone. |

The first native verification failed in all three modules and its full
[feedback](native-feedback-1.txt) was reviewed before the second release. The
second worker received exactly the original prompt plus that controller-produced
feedback. No outside finding, proposed repair or implementation hint was added.
Final stdout/stderr and actual assertion text are retained in the report;
[non-case fields](native-noncase-fields.txt) make those fields easy to inspect.
Both attempts compiled and produced actual case reports; no compiler error was
reported, so this invocation adds no compiler-error-path claim to earlier CHK-034
validation.

## Decisions and gap limits

The native Contradicted Decisions field says None identified. Preserve that as
the worker's statement, not independent certification. Final notes contain five
gap claims for FLOW-022, FLOW-023, FLOW-115, FLOW-116 and FLOW-117.

Do not accept the first two as demonstrated retained-eval defects. The published
failure for the audit fixture requests fail-first for a review-tagged ALP-003;
[design tags](../../../fireweed-design-v1.md) state that review specs need no evals,
and CHK-002 explicitly excludes review from coverage. The implementation fixture's
published State instead shows a newly allocated ALP-002 colliding with the Removed
ALP-002. These observations do not establish the claimed conflict in retained
inputs. No fixture/assertion was removed, weakened or corrected on this basis.

The remaining notes describe rename identity, shared file attribution and use of
endpoint zones for historical file categories. FLOW-117's retained Decision
explicitly calls for each phase's declared metadata, so the admitted endpoint-only
classification is an unresolved reconstruction limitation rather than an agreed
new spec gap. The other claims remain unvalidated observations. The unchanged
source passed clean normal/race/CI, but that alone is not a proof about every
uncovered case. No further investigation or correction was authorized by these
notes. WF4 remains open both because native verification failed and because these
claims/limitations are not resolved by an owner decision.

## Isolation, trace and durable output

Actual first inputs were fully reviewed: 31 files, comprising workflow retained
spec/Decisions/interface/model/evals, SYSTEM and only the declared dependency
public api.go/model files. Native input Git contained one input-only commit, no
remotes/alternates/unreachable old objects, and no symlinks. Dependency/dependent
implementations, manifests and hooks stayed in the separate verifier. No import
inventory or expanded worker inputs were supplied. [Input review](input-review.json),
ready records, expected hashes and all repeated live preflights are retained.

First release followed actual input review. Second release followed the complete
first trace, fresh-output and native-feedback review plus repeated isolation.
The first trace has 77 full commands/scripts and 190 cumulative patch snapshots;
the second has 24 commands and 25 snapshots. The outside reviewer displayed every
incremental change while suppressing repeated content and normalizing Go format
for comparison. [First](output-review-1.json) and [second](output-review-2.json)
receipts match all actual fresh files, including notes, to that complete trace and
formatting. The final proof's four Go files match the second receipt. Reviewer
helper `stderr_bytes` display counts decoded characters; authoritative byte counts
are the raw stream/finished/forwarding records, not that display label.

Observed commands stayed in the supplied tree/fresh files. The second worker tried
relative runner/app eval paths that were absent; stderr confirms both reads failed.
It did not read their contents or receive them afterward. No observed outside
read, implementation oracle, network use, other-agent call, local build/test/module
setup, retained-input edit, Git ref/config mutation or commit occurred. Only fresh
code was formatted. No command interception or filesystem/process instrumentation
was introduced. This is an observation of the reviewed commands/gates, not a claim
about unobservable activity or escaped processes.

Both raw streams were continuously fsynced outside reconstruction before identical
bytes were forwarded. Both workers exited zero, finished with EOF on both streams,
no stop/error/truncation, and complete forwarding. Exact native log comparisons:

| Attempt | stdout bytes | stderr bytes | Native log bytes | Match |
| --- | ---: | ---: | ---: | --- |
| 1 | 543 | 2,991,269 | 2,991,852 | exact |
| 2 | 568 | 258,022 | 258,629 | exact |

[Forwarding 1](forwarding-1.json), [forwarding 2](forwarding-2.json), raw gate streams
and native-run-logs retain the bytes. Native normal failure publication was observed.
No rejection, interruption or emergency was used here; native reporting on those
paths remains unexercised. Earlier bounded fake rehearsals remain relevant only
to unchanged supervisor behavior. The superseded pre-release stop established no
running-worker drain behavior.

## Ref, cleanup and preservation

Native failed ref `refs/fireweed/drills/workflow/20261107T232508.801610000` points to `21039de866b8f03d9803b6fe7f80e744e8ad55c9`. It was retained locally
before caller cleanup, without checkout, reconstructed-source application or push.
Its parent/setup/candidate ancestry and disposable-only diff are checked in
[retention](retention.json). The current main implementation remains the validated
source, not this failed reconstruction.

[Cleanup](cleanup.json) records observed native/worker groups exited, native input
and verifier removed, no caller claim or extra worktree, and caller removal after
exact report/log/ref retention. Escaped/reparented processes remain the documented
limit. The previous complete workflow report is preserved byte-for-byte at
[pre-retry report](../workflow-before-corrected-controller-retry-3853f17.md).
The current root workflow report retains this exact native report as its prefix,
followed only by a labeled owner review. Prior archives/reports/proof/attempt refs,
old phase counters, recoveries and the unrelated worktree are preserved; the two
new archives retain exactly their three normal phase commits. Final status records
live unapproved 3 and the existing B02 warnings, with no reset or raised bound.

Stop after the protected evidence commit. The queue/shared claims and all owned
native/validation runs are idle. No further correction, drill, source application,
controller change, publication or push is approved. WF4 and full completion remain
open; a later scoped action needs a new owner instruction.
