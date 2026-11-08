from pathlib import Path
import json,shutil,subprocess,hashlib
R=Path(__file__).resolve().parent;ROOT=Path('/Users/joel/Workspace/fireweed');NAME='workflow-corrected-controller-retry-20261107T232508';E=ROOT/'.fireweed/drills'/NAME;S='ccd3d6b7732d613e80a027ba7969b65adc58bc0e';H='20278504a8ea89fcd5f23d1de9a6ec849389fd55707054992c59967eec65e2b1';ref=json.loads((R/'retention.json').read_text());assert (R/'cleanup.json').exists() and not (R/'caller').exists();assert hashlib.sha256((ROOT/'.fireweed/bin/fireweed').read_bytes()).hexdigest()==H;assert subprocess.check_output(['git','-C',str(ROOT),'for-each-ref','--format=%(refname)','refs/fireweed/claims'],text=True).strip()=='';assert not list((ROOT/'changes').glob('*.md'));shutil.copytree(R,E)
for n in ['entry.json','prepare.stdout','prepare.stderr']:shutil.copy2(R.parent/n,E/n)
review=f'''# Corrected-source workflow drill: terminal review

The owner's requested sequence is complete: two normal audited fixes, clean-source
validation, an idle verified controller update, then exactly one native workflow
invocation on corrected subject `{S}`. It **failed native verification** after its
two configured workers. No reconstruction was applied and no additional correction,
diagnostic replay, drill or push followed. The approval is consumed; stop here.

## Subject, controller and preparation

[0181](../../changes/archive/0181-recorded-regenerate-preservation.md) applied at
fd7ae49 and [0182](../../changes/archive/0182-recovery-history-assertions.md) applied
at ccd3d6b. Their independent audits, phase reports, actual pause/rejection history,
cost limitation, clean normal/race/vet/check/doctor/history/CI and compiled CLI
validation are in [the verified handoff](../residual-fixes-controller-handoff.md),
committed separately at 3853f17. The original 79cf33a remains a preserved baseline.
The initial CHK-083 normal-suite timing failure, passing focused replay and single
passing ordinary retry are retained there; no cause or bound change was invented.

The current controller source is that corrected subject and SHA-256 is `{H}`.
It stayed unchanged throughout this native invocation. Its retiring 12c571c copy
and every older recovery executable remain preserved. The outside inspector's
candidate/pin metadata was updated only after both normal claims were released.
Original worker command/model/effort, worker SYSTEM instructions, all configuration
bounds/counters and the existing 2026-11-10 22:51:05 UTC deadline stayed intact.

The clean isolated setup child is `{ref['setup']}`, parented directly on the
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

Native failed ref `{ref['ref']}` points to `{ref['oid']}`. It was retained locally
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
'''
(E/'review.md').write_text(review)
owner=f'''\n\n## Owner review — corrected-source retry\n\nThe preceding native report is retained byte-for-byte. The two scoped fixes are\napplied and independently audited, the clean corrected source/controller handoff\nis complete, and exactly one approved workflow invocation used that new pin.\nNative verification failed: 646 passing and 4 failing cases; app passed. All\nenforced IDs are represented. The focused FLOW-081/FLOW-083 and RUN-027/RUN-028\ncases pass, but there is no passing independent workflow proof.\n\n[Complete evidence and review]({NAME}/review.md) retains actual failure text,\ncoverage, full feedback/trace, gap limitations, durable raw streams, exact native\nforwarding, failed ref, preservation and cleanup. Worker gap claims are not\nauthorization for retained-input corrections. Native normal failure publication\nwas observed; rejection/interruption/emergency reporting was not exercised.\n\nThe corrected subject is `{S}`. The failed ref is\n`{ref['ref']}`; its code was not applied.\nThe previous complete report is preserved at\n[pre-retry evidence](workflow-before-corrected-controller-retry-3853f17.md).\nThe existing 2026-11-10 22:51:05 UTC deadline, bounds/counters, pin, prior refs and\nunrelated worktree are preserved. Approval consumed: stop with no active run or\nclaim. No further correction, drill, controller change or push is authorized.\n'''
(ROOT/'.fireweed/drills/workflow.md').write_bytes((R/'native-report.md').read_bytes()+owner.encode())
(ROOT/'docs/rebuild/GOAL.md').write_text(f'''# Stopped after the corrected-source workflow retry

The owner's full sequence is complete: normal independently audited fixes
[0181](../../changes/archive/0181-recorded-regenerate-preservation.md) and
[0182](../../changes/archive/0182-recovery-history-assertions.md),
[clean-source validation and verified controller update](../../.fireweed/drills/residual-fixes-controller-handoff.md),
then exactly one native workflow invocation on corrected subject `{S}`.
It failed native verification; app passed, workflow and runner have remaining
failures. The focused recorded-phase/Context and recovery-history cases pass.

The [exact native report and owner review](../../.fireweed/drills/workflow.md) and
[complete evidence](../../.fireweed/drills/{NAME}/review.md)
retain the failed ref, actual assertions, coverage, worker gap limits, durable
streams, exact forwarding and cleanup. No reconstruction was applied. Worker
claims about retained eval conflicts are not accepted fixes or permission to
weaken tests. WF4 and full completion remain open.

**Stop for owner review. This approval is consumed.** No further correction,
diagnostic replay, drill, controller promotion, source application or push is
authorized. Any later scoped action requires a new owner instruction. Read
AGENTS, SYSTEM, PROGRESS and ACCEPTANCE before that work; preserve all evidence,
refs, bounds/counters, recoveries and unrelated worktrees. Do not infer a new
execution goal from the remaining failures or worker notes.

No active normal/native/validation run or shared claim remains; the queue is
empty and owned scratch is cleaned up. The pin in PROGRESS stays unchanged.
Original 79cf33a remains a preserved baseline. The existing deadline is still
**2026-11-10 22:51:05 UTC**, with no further extension.
''')
(ROOT/'docs/rebuild/PROGRESS.md').write_text(f'''# Rebuild handoff

The owner-requested sequence is complete and stopped. The two scoped fixes are
applied and independently audited in [0181](../../changes/archive/0181-recorded-regenerate-preservation.md)
and [0182](../../changes/archive/0182-recovery-history-assertions.md). No runtime
source, public interface, hook or configuration correction was needed. The old
pin governed both normal changes; the
[verified corrected-source handoff](../../.fireweed/drills/residual-fixes-controller-handoff.md)
records phase verification, clean-source normal/race/vet/check/doctor/history/CI,
compiled CLI regressions, actual pause/rejection history and cost/timing limits.

Corrected subject/controller source is `{S}`. Exactly one approved native
workflow invocation then used that new pin and prepared outside inspector.
It **failed native verification**. App passed; workflow and runner have remaining
failures. The focused FLOW-081/FLOW-083 and RUN-027/RUN-028 cases passed. There is
no passing independent workflow proof. WF4 and full completion remain open.

The [current workflow report](../../.fireweed/drills/workflow.md) retains the exact
native prefix followed by owner review. Its
[complete evidence and review](../../.fireweed/drills/{NAME}/review.md)
records every case, actual assertion text, coverage, full first feedback and both
traces, worker gap limitations, durable streams/exact forwarding, retained failed
ref and cleanup. Gap claims are not approved corrections. Normal native failure
publication was observed; rejection/interruption/emergency reporting was not used.
Neither attempt reported a compiler error. No extra diagnostic replay followed.

Failed ref: `{ref['ref']}`. It was retained locally; reconstructed code was not
applied. The previous complete workflow report remains
[byte-preserved](../../.fireweed/drills/workflow-before-corrected-controller-retry-3853f17.md).
Original `79cf33a7690082f28c06bf41a72719d282bc00fe`, all prior archives/reports/proof/
attempt refs, old phase counters, recoveries and unrelated worktree are preserved.
Final live counter/warning details are in evidence; no limits were raised or reset.
The valid checks proof and WF1/WF2 history prefixes remain valid.

**Stop for owner review. The approval is consumed.** No further correction,
diagnostic replay, drill, source application, controller update or push is
currently authorized. Follow [GOAL](GOAL.md) and [ACCEPTANCE](ACCEPTANCE.md).
The owning native controller, inspector and both workers exited; input/verifier/
caller scratch is cleaned up. The queue/shared claims are empty and no owned
normal/native/validation run remains active. Protected final evidence is committed
separately after terminal cleanup.

The owner added exactly 72 hours to the original 2026-11-07 22:51:05 UTC deadline.
The existing deadline remains **2026-11-10 22:51:05 UTC** (14:51:05 PST), with no
further extension. All configured bounds, independent audits and isolation gates
remain unchanged. This deadline grants no extra work after the completed invocation.

**Controller pin** `/Users/joel/Workspace/fireweed/.fireweed/bin/fireweed`, version
0.1.0, source `{S}`, SHA-256
`{H}`. Built with Go 1.27.2 and `-trimpath` from clean source on darwin/arm64.
Its owner-authorized verified handoff is recorded above; it stayed unchanged
throughout the native retry.

Recovery executables remain at `.fireweed/bin/fireweed-12c571c`,
`.fireweed/bin/fireweed-1936346`, `.fireweed/bin/fireweed-99a1d6d`,
`.fireweed/bin/fireweed-pre-ck-handoff` and `.fireweed/bin/fireweed-bootstrap`.
Use them only for explicit owner-directed recovery with no automatic fallback or
cross-controller resume. Previous push approvals are historical; no push is
currently authorized.
''')
p=ROOT/'docs/rebuild/ACCEPTANCE.md';t=p.read_text();a=t.index('**Residual fixes and verified controller handoff complete;');b=t.index('**Entry.**',a);t=t[:a]+f'''**Residual fixes/controller handoff and single retry complete; native verification failed.**
Normal audited [0181](../../changes/archive/0181-recorded-regenerate-preservation.md)
and [0182](../../changes/archive/0182-recovery-history-assertions.md) are applied.
The [corrected-source validation and verified pin handoff](../../.fireweed/drills/residual-fixes-controller-handoff.md)
passed before the one owner-authorized native workflow invocation. The
[current native report](../../.fireweed/drills/workflow.md) and
[complete review](../../.fireweed/drills/{NAME}/review.md)
retain the actual failures, coverage, gap limitations, failed ref, full output and
cleanup. The focused corrected cases passed, but workflow/runner still fail.
No passing workflow proof or WF4 closure is claimed.

The owner's sequence approval is consumed. Stop for review with no active run or
claim. Further correction, diagnostic replay, drill, controller promotion,
reconstructed-code application or push requires new authorization. Original
79cf33a, prior evidence/refs, recoveries, bounds/counters and unrelated worktree
are preserved. Worker notes authorize no new change. Full-completion gates remain.

'''+t[b:];p.write_text(t)
p=ROOT/'docs/rebuild/workflow-drill/README.md';t=p.read_text();a=t.index('`guard.py`');t=f'''# Prepared workflow wrapper and inspector — stopped

The two normal audited residual fixes and verified controller update are complete.
Exactly one approved native workflow retry on corrected subject `{S}` then failed
verification. PROGRESS records the current unchanged pin;
[complete retry evidence](../../../.fireweed/drills/{NAME}/review.md)
retains output, actual failures, failed ref and cleanup. The focused corrected
cases passed; WF4 remains open. Original 79cf33a and all prior evidence are preserved.

**Approval consumed. Do not launch another drill.** The commands below describe a
restart only after a new owner instruction, a new outside root and fresh live
gates. No further correction, source application, controller update or push is
currently authorized. The existing deadline remains **2026-11-10 22:51:05 UTC**.
Supervisor behavior, exact worker command/model/effort, input policy and every
isolation/output/stop gate are unchanged; earlier fake rehearsals remain valid
for that behavior, not proof of native reporting on rejection/interruption.

'''+t[a:];t=t.replace('## Execute the single owner-authorized retry','## Restart only after new owner authorization').replace('**This command is for the currently approved single retry only:**','**This command requires a new explicit owner approval:**');p.write_text(t)
print('Native evidence and protected stopped handoff written; final integrity/status/author checks and commit remain.')
