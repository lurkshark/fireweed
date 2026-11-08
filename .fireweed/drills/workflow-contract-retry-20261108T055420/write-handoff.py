from pathlib import Path
import json,subprocess,hashlib,shutil
O=Path(__file__).resolve().parent;R=Path('/Users/joel/Workspace/fireweed');S=O/'run';A=json.loads((O/'native-analysis.json').read_text());assert (O/'cleanup.json').exists() and (O/'process-cleanup.json').exists();assert not subprocess.check_output(['git','-C',str(R),'for-each-ref','--format=%(refname)','refs/fireweed/claims'],text=True).strip();assert not list((R/'changes').glob('*.md'));D='.fireweed/drills/workflow-contract-retry-20261108T055420';E=R/D;E.mkdir(exist_ok=False)
review='''# Contract-diagnostics controller retry — rejected and stopped

The single owner-approved native workflow invocation on validated source
`9a206c558e8e5625c0ae55bb57c42ae8130b4f15` is complete and failed. The second
reconstruction worker invoked prohibited local `go test`; the inspector rejected
it and used the prepared worker-first stop. No passing independent workflow proof
or WF4 closure is claimed. No correction, extra invocation, source application,
controller update or push followed. The approval/budget is consumed.

## Scope, source and retained result

The unchanged primary controller is `/Users/joel/Workspace/fireweed/.fireweed/bin/fireweed`,
source 9a206c5, SHA-256 `cc22f47c7aa607c96a13273d29d5bf4ade4f912453840f621331d73dc385f112`.
Reuse [clean-source validation](../workflow-contract-diagnostics-handoff.md) and
unchanged earlier fake supervisor rehearsals. The current invocation used a fresh
no-hardlinks caller at that source and setup child
`dee40749e8a0c803f1b699944bc99f968150011f`, changing only worker SYSTEM and
agent.wrap. Removing the wrap line recovers exact source configuration. Worker
command/model/effort, attempts, timeout, limits, audit configuration, all isolation
and output/stop gates, and **2026-11-10 22:51:05 UTC** deadline stayed unchanged.

The controller published the exact [native report](native-workflow.md) and
`refs/fireweed/drills/workflow/20261108T055420.712886000` at
`8859d4433b16bf721907009e58b269937b8af4fe`. The native CLI returned exit 1,
P01, data null and cause `reconstruction agent: exit status 130`, naming both
retained artifacts. Its generic diagnostic advice to correct/retry grants no
owner authorization. The native report frontmatter says fail, agent codex,
model gpt-6-sol, date 2026-11-08T05:54:20.712886Z, base_commit equal to the setup
child. The ref's parent is that setup child, and its six disposable files match
all first-attempt verified source hashes. Rejected second partial code is absent
from that proof. Required evidence objects/ref were retained locally without
fetch, push, checkout, merge or applying reconstruction to the product.

The root [workflow report](../workflow.md) preserves this complete native prefix
before owner review. Its preceding report is
[byte-preserved](../workflow-before-contract-diagnostics-retry-618ecde.md).
Original 79cf33a and prior corrected ccd3d6b/9a206c5 objects, all older archives/
reports/proof/attempt refs, recoveries, bounds/counters and unrelated worktree
are preserved. No product/spec/eval correction or worker-input expansion.

## Actual inputs and first worker

Both live preflights passed, including the repeated check immediately before each
release. The actual 31 retained files (241,234 bytes) were read in full: unchanged
worker SYSTEM, workflow SPEC/Decisions, facade/model/evals and only declared
contracts/gitstore/assets/project/checks public api/model interfaces. The input
Git had one commit, no remotes/alternates/unreachable old objects and no symlinks.
Build/setup/manifests/hooks and all dependency/dependent implementations stayed
in the separate verifier. All three controller copies matched. Attempt two had
only those same retained files plus seven reviewed first-output files; its prompt
was exactly the original plus the full 16,412-byte native verification feedback.
There were no supplied import inventories, implementation hints, command shims,
interceptors or process/filesystem instrumentation.

First worker: 47 full shell commands and 24 applied patch records reviewed.
Its shell scripts and incremental changes were confined to supplied input/fresh
code, with only workflow/internal and DRILL_NOTES written. All retained hashes
remained unchanged at checks. It formatted its new code and returned without a
build/test/module setup or commit. A relative absent checks/evals path was named
by rg; no dependency eval content was provided. One failed apply_patch context
lookup produced a visible diagnostic and no applied edit. The received applied
diff records, full scripts and hashed snapshots are retained; repeated cumulative
diffs were reviewed by exact content identity rather than repeatedly printed.
First exit was 0. Full first trace/feedback and fresh-output hashes were reviewed
before the final configured release; no additional feedback was injected.

## Verification cases and limitations

All reported cases are first-attempt verification, retained when the second
reconstruction failed. The second partial edits were **not reverified**.

| Module | Pass | Fail | Total |
| --- | ---: | ---: | ---: |
| workflow | 155 | 15 | 170 |
| app | 317 | 6 | 323 |
| runner | 158 | 1 | 159 |
| Total | 630 | 22 | 652 |

All 652 case rows/names/ID mappings/results were reviewed in the
[case catalog](case-catalog.txt); all failure bodies match the reviewed first
[native feedback](run/feedback-1.txt) and are retained in [failures.json](failures.json).
Coverage includes all 355 current IDs (workflow 133, app 155, runner 67), with no
unknown/missing IDs, duplicate case names, skips or error-result cases. Failure
counts include parent/subtest rows and are not counts of distinct defects.
Native hook stdout records Go suite times 118.949s workflow, 274.329s app and
293.781s runner. No verifier compiler error occurred. The failed worker-local Go
command below is a missing-module diagnostic, not a verifier compiler failure.

The focused new FLOW-007 reservation and FLOW-023 review exclusion tests passed;
the existing review audit-convergence and impl-application cases also passed.
Focused recorded-regenerate FLOW-081/FLOW-083, RUN-027/RUN-028 and RUN-037
verification-deadline cases passed. This does not establish a RUN-037 timing fix
or cause, and the combined FLOW-080/081/083 integration case still failed because
project-protected global.go was deleted.

Other first-attempt failures cover staged Abandon effects/cleanup, ambiguous
selector U02 handling, eval_fix/maintenance SPEC byte preservation, stale/static
cause precedence, fixes fail-first, failed Done phase retention, boundary cause
text, empty PRBody wording, numbered blocker identity, B04 path metadata,
unavailable Status base and Context headings (RUN-013/RUN-014). These are observed
reconstruction failures, not approved corrections to the retained product/evals.
The full assertion text is preserved rather than reduced to nearest spec IDs.

Native Contradicted Decisions says none identified. Its two Gaps concern FLOW-089
conflicting reverts/manual recovery and FLOW-117 phase-specific historical zones/
evals stubs. They are worker claims/limitations, not accepted missing requirements
or authorized changes. FLOW-117's representative cases passed despite the note;
passing tests do not independently establish every stated gap's completeness.
No replay, correction, eval weakening or bootstrap/reconstructed-source import
was used to resolve a gap after this invocation.

## Observed rejection, output and publication

Second worker: 23 shell commands and three applied patch records reviewed through
termination. It read fresh/retained code and named unavailable runner/app eval
paths inside the input tree; those files were absent, so no dependent eval content
was exposed. Command 18 invoked:

```text
gofmt -w packages/workflow/internal/lifecycle && go test ./packages/workflow/internal/lifecycle ./packages/workflow
```

The immutable SYSTEM explicitly says not to run local builds, tests or module
setup. Go returned `cannot find main module` and suggested go mod init. This
command was prohibited regardless of its failure; no module setup was observed.
The inspector rejection was posted on observing the captured command. A few more
already-observed/queued commands and one applied patch had arrived before the
worker stopped; all received records were reviewed. No command interception was
added. Second partial code was unverified and discarded with owned native input;
its last captured source snapshot is not certified as a complete cancelled final
source tree. Native report/ref retain the verified first source only.

The wrapper terminated only its owned worker group first. Worker 2 exited by
SIGTERM (-15); wrapper returned 130, both stream EOFs arrived, forwarding completed
and errors was empty. Abort latched and no further readiness/release appeared.
The native controller was left running and actually published report/ref/P01
before exiting 1. Inspector terminal says rejected true, emergency false. This
establishes native publication after **this running-worker rejection**, not all
interruption/storage/broken-pipe/whole-run emergency outcomes.

| Attempt | stdout received/forwarded | stderr received/forwarded |
| --- | ---: | ---: |
| 1 | 573 / 573 | 700530 / 700530 |
| 2 | 0 / 0 | 192056 / 192056 |

Both wrappers report complete forwarding and EOF. Every nonempty raw stream is
byte-identical to the stream retained in the native report; native empty stdout
for attempt 2 is consistent with its zero count. Logs were continuously fsynced
outside reconstruction before forwarding unchanged. Complete forwarding means all
**received** bytes were preserved; it does not claim the terminated worker
finished its intended output or commands.

## Cleanup and stopped handoff

[Process cleanup](process-cleanup.json) checks the known inspector/controller/
worker groups exited. Native cleanup removed the input and verifier. Only the
owned idle caller clone was removed after claim/worktree checks and durable
report/ref retention. Escaped/daemonized/reparented processes retain SYSTEM's
stated limitation; no global process instrumentation was introduced.
[Cleanup](cleanup.json), [ref retention](retained-ref.json), [native analysis](native-analysis.json),
all gate/ready/review/release/stop/finished/terminal records, preflight logs, raw
streams and review ledgers are retained. Counters/configuration/older artifacts
and the unrelated worktree are checked in preservation.json. Normal prior cost
warnings remain recorded in live status; no bounds or counters were raised/reset.
Expected U02 context/P02 idle Verify are state checks, not product verification.

The approved single invocation and evidence handoff are complete. **Stop with no
active run or claim.** No passing independent proof/WF4 closure. Further correction,
worker-policy/tooling change, diagnostic replay, pin update, retry or push requires
new owner selection/authorization. Follow current GOAL/PROGRESS. The remaining
runtime failures and the worker instruction breach are separate owner checkpoints;
no automatic action follows from a native diagnostic or worker Gap.
'''
(E/'review.md').write_text(review)
old=O/'workflow-before-retry-618ecde.md';assert hashlib.sha256(old.read_bytes()).hexdigest()==json.loads((O/'entry.json').read_text())['files']['.fireweed/drills/workflow.md'];(R/'.fireweed/drills/workflow-before-contract-diagnostics-retry-618ecde.md').write_bytes(old.read_bytes());native=(O/'native-workflow.md').read_bytes();owner=('\n\n# Owner review — rejected single retry\n\nThe approved invocation failed. Its second worker ran prohibited local go test;\nthe inspector stopped it, drained/forwarded all received output and the native\ncontroller published this failure report and retained ref. The 652 reported cases\nare first-attempt verification (630 pass, 22 fail); second partial edits were not\nreverified. No passing independent workflow proof or WF4 closure.\n\n[Complete review and evidence](workflow-contract-retry-20261108T055420/review.md)\nrecords exact inputs, cases, actual assertion text, trace/rejection/publication,\nref identity, limitations and cleanup. The preceding report is\n[byte-preserved](workflow-before-contract-diagnostics-retry-618ecde.md).\nNo correction, extra invocation, source application, repin or push followed.\nApproval consumed; stop for owner review with no active run or claim.\n').encode();(R/'.fireweed/drills/workflow.md').write_bytes(native+owner)
pin='''**Controller pin** `/Users/joel/Workspace/fireweed/.fireweed/bin/fireweed`, version
0.1.0, source `9a206c558e8e5625c0ae55bb57c42ae8130b4f15`, SHA-256
`cc22f47c7aa607c96a13273d29d5bf4ade4f912453840f621331d73dc385f112`.
Built with Go 1.27.2 and -trimpath from clean darwin/arm64 source. It stayed
unchanged throughout the retry. No fallback or cross-controller resume.

Recovery executables remain `.fireweed/bin/fireweed-ccd3d6b`,
`.fireweed/bin/fireweed-12c571c`, `.fireweed/bin/fireweed-1936346`,
`.fireweed/bin/fireweed-99a1d6d`, `.fireweed/bin/fireweed-pre-ck-handoff` and
`.fireweed/bin/fireweed-bootstrap`. Use only for explicit owner-directed recovery.
'''
progress='''# Rebuild handoff — rejected retry, stopped

The one owner-approved supervised workflow retry on validated source 9a206c5 is
complete and failed. The second worker invoked local go test despite the supplied
SYSTEM prohibition. The prepared inspector rejected it, stopped the worker first,
drained/forwarded all received output and gave the native controller its bounded
publication opportunity. The controller actually published failure evidence/ref
and exited without emergency escalation. This observation establishes that
reporting outcome for this running-worker rejection, not every interruption or
emergency path. No further invocation occurred.

The [current native report](../../.fireweed/drills/workflow.md) retains its exact
native prefix followed by owner review. The
[complete evidence/review](../../.fireweed/drills/workflow-contract-retry-20261108T055420/review.md)
records first-attempt verification failures, all cases/coverage and assertions,
actual inputs, exact first feedback, both traces, rejected local command, durable
unchanged stream forwarding, published ref, process/scratch cleanup and limitations.
Second partial edits were not reverified. No passing independent workflow proof;
WF4 and full completion remain open. No compiler error occurred in native verification.
The focused RUN-037 case passed; its preceding native timing cause remains unresolved.

Retained failed ref: `refs/fireweed/drills/workflow/20261108T055420.712886000`.
It identifies first verified reconstruction only; no reconstruction was applied
to the product. The preceding workflow report is
[byte-preserved](../../.fireweed/drills/workflow-before-contract-diagnostics-retry-618ecde.md).
All older reports/archives/proof/attempt refs, original 79cf33a and prior corrected
ccd3d6b/9a206c5 objects, bounds/counters, recoveries and the unrelated worktree remain.
Live counts/warnings are retained in evidence, not reset. The valid checks proof
and composed WF1/WF2/history prefixes remain valid.

The [two audited changes and verified preparation](../../.fireweed/drills/workflow-contract-diagnostics-handoff.md)
remain complete: FLOW-007/FLOW-023 clarification in 0183, diagnostic-only RUN-037
in 0184, with no runtime/interface/hook/configuration correction. Prior residual
fixes/handoff and the [remaining-failure investigation](../../.fireweed/drills/workflow-failure-investigation-20261108/findings.md)
remain historical evidence. No worker Gap or native diagnostic authorizes a fix.

**Approval consumed. Stop for owner review.** No active owning process, queued
change or shared claim remains. Native input/verifier and owned caller are cleaned
up; the required native ref and protected evidence are retained. Further correction,
worker-policy/tooling change, replay, retry, controller promotion, source application,
later work or push requires new owner authorization. Follow GOAL.

The owner previously added exactly 72 hours to the original 2026-11-07 22:51:05 UTC
deadline. The existing deadline remains **2026-11-10 22:51:05 UTC** (14:51:05 PST),
without another extension. Configured bounds and isolation/audit requirements stay
unchanged. Remaining time grants no further work after this completed invocation.

'''+pin;(R/'docs/rebuild/PROGRESS.md').write_text(progress)
goal='''# Stop after the rejected workflow retry

The owner-approved single native workflow invocation on validated source 9a206c5
with the pinned controller is complete. The second worker invoked prohibited
local go test, so the inspector rejected it using the worker-first stop/drain
path. All received output was preserved/forwarded; the native controller published
a failure report/ref and exited. No emergency escalation or additional invocation.

Read AGENTS, SYSTEM, PROGRESS and ACCEPTANCE before any newly authorized work.
The [complete reviewed evidence](../../.fireweed/drills/workflow-contract-retry-20261108T055420/review.md)
contains exact source/pin/setup, actual retained-input review, all native cases/
assertions/coverage, full first feedback, traces, rejection/publication, retained
ref, limitations and cleanup. The [workflow report](../../.fireweed/drills/workflow.md)
preserves the full native prefix before owner review. Its preceding report and
all earlier evidence/refs remain accessible. No passing independent proof or
WF4/full-completion claim. Cases retained in that report are first verification;
rejected second partial code was not reverified. RUN-037's preceding native timing
cause remains unresolved, although this invocation's focused case passed.

**Stop for owner review with no active run or claim. Approval/budget consumed.**
No extra correction, diagnostic replay, worker-policy/tooling change, native retry,
source application, pin update, later milestone or push is authorized. Native
advice to correct/retry and worker Gaps grant no permission. Further action needs
owner selection of a concrete scope. Preserve original 79cf33a, corrected ccd3d6b/
9a206c5 objects, all reports/archives/proof/attempt refs, exact controller/recoveries,
configuration/counters/bounds and unrelated worktrees. Existing deadline remains
**2026-11-10 22:51:05 UTC**, with no further extension.

Reuse completed normal audited fixes/clarifications and clean-source validation
where relevant. No product/spec/eval correction or reconstruction import occurred
after this native failure. The required ref retains first verified reconstruction
as evidence only. Owned native input/verifier/caller are removed and known owning
groups exited. No release followed rejection.

If a later owner instruction authorizes another invocation, first follow the
[inspector README](workflow-drill/README.md) and recheck live pin/source, queue,
claims, exact command/model/effort, deadline and every isolation/review/output/stop
gate using a new outside root. Never reuse stale release attestations, provide
old reconstructions, expand inputs, add hints/interceptors/instrumentation or
switch controllers during a live run/claim. No such later invocation is approved now.
''';(R/'docs/rebuild/GOAL.md').write_text(goal)
p=R/'docs/rebuild/ACCEPTANCE.md';s=p.read_text();start=s.index('The owner separately approved exactly one supervised native workflow retry');end=s.index('**Residual fixes/controller handoff',start);stop='''The subsequently owner-approved single native retry on validated source 9a206c5
is complete and failed. A prohibited worker-local go test caused running-worker
rejection. The native controller published its failure report/ref after the
prepared worker-first stop/drain. The
[reviewed terminal evidence](../../.fireweed/drills/workflow-contract-retry-20261108T055420/review.md)
records all first-attempt cases/failures, coverage, actual input/trace review,
rejection/publication and cleanup; second partial code was not reverified. No
passing independent proof, WF4 closure, correction, extra invocation, source
application, repin or push. Approval consumed; stop with no active run/claim.
Further work requires owner selection; all prior gates/evidence and preservation
requirements remain. Follow current GOAL/PROGRESS.

''';s=s[:start]+stop+s[end:];s=s.replace('This budget is exhausted; stop for separate\nowner approval of one retry on the identified source/pin.','That preparation budget is exhausted. The separately approved retry is now\ncomplete and stopped as recorded above.');p.write_text(s)
p=R/'docs/rebuild/workflow-drill/README.md';s=p.read_text();i=s.index('`guard.py` is an');top='''# Workflow inspector — rejected retry, stopped

The one approved native workflow invocation on validated source 9a206c5 is complete.
Its second worker invoked forbidden local go test. The inspector rejected it,
stopped the worker first, drained/forwarded all received output and allowed the
native controller to publish failure evidence/ref. It did so and exited without
emergency escalation. [Complete review/evidence](../../../.fireweed/drills/workflow-contract-retry-20261108T055420/review.md)
records the actual result and cleanup. No passing independent workflow proof;
WF4 stays open. Second partial code was not reverified or applied.

**Approval consumed. Do not launch another drill.** The restart commands below
require new explicit owner authorization, a fresh outside root and fresh live
review/isolation gates. No correction, policy/tooling change, source application,
controller update, extra invocation or push is authorized now. Existing deadline
remains **2026-11-10 22:51:05 UTC**. PROGRESS records the unchanged pin. Original
79cf33a, corrected ccd3d6b/9a206c5, all earlier reports/refs, recovery copies,
configuration/bounds/counters and unrelated worktrees remain preserved.

Supervisor/inspector behavior, worker SYSTEM and exact command/model/effort are
unchanged. Existing fake rehearsals still establish only their wrapper outcomes.
This actual invocation additionally observed native publication after its
running-worker rejection; it establishes no generic native reporting guarantee
for other interruption/storage/broken-pipe/emergency outcomes.

''';p.write_text(top+s[i:]);print('Wrote protected stopped handoff and native-prefix review; no product files changed.')
