# Contract-diagnostics controller retry — rejected and stopped

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
Live status differs from entry only in derived workflow drift (1 to 0 after
retaining the new failed report); this is not a passing-proof claim or a manual
counter reset. Unapproved and phase-commit counts and all configured limits are
unchanged. Expected U02 context/P02 idle Verify are state checks, not product verification.

The approved single invocation and evidence handoff are complete. **Stop with no
active run or claim.** No passing independent proof/WF4 closure. Further correction,
worker-policy/tooling change, diagnostic replay, pin update, retry or push requires
new owner selection/authorization. Follow current GOAL/PROGRESS. The remaining
runtime failures and the worker instruction breach are separate owner checkpoints;
no automatic action follows from a native diagnostic or worker Gap.
