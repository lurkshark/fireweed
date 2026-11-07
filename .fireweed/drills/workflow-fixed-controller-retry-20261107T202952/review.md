# Fixed-controller workflow retry review

One owner-approved real native workflow invocation used frozen subject
`79cf33a7690082f28c06bf41a72719d282bc00fe`, the fixed pinned controller and the
prepared outside inspector. It completed normally with native exit 1 and result
**fail**, using exactly the two configured reconstruction workers. No rejection,
interruption or emergency stop was needed. No further native invocation ran.

| Module | Passing cases | Failing cases |
| --- | ---: | ---: |
| workflow | 166 | 1 |
| app | 323 | 0 |
| runner | 158 | 1 |
| Total | 647 | 2 |

The remaining published failures are
`Test_FLOW_083_RecordedRegenerateImplContextUsesRegeneratePlaybook` (FLOW-083) and
`Test_RUN_027_RUN_028_RecoversAgentCommitsAcrossLifecycle` (RUN-027, RUN-028).
All 649 native case rows were reviewed: identities are unique per module, IDs are
known, and all enforced workflow/app/runner IDs are represented. This is coverage,
not passing verification. The actual report includes module stdout/stderr summaries
and E01 testcase identities; no unpublished final assertion text is claimed.
Neither attempt produced a compilation failure, so this invocation does not
establish compiler-error-path behavior; the earlier CHK-034 fixture is separate.

The exact [native report](native-report.md), [CLI result](native.stdout),
[terminal record](terminal.json) and [case review](case-review.json) are retained.
The failed reconstruction ref is
`refs/fireweed/drills/workflow/20261107T202952.649365000`, object
`50bd33c5526fb495d541e3290868d198ed010b78`. It is retained locally, not applied to
main. Its native base is setup child `d207f68fd4581e49d063906e961ba28921e2372a`,
whose parent is the frozen subject and whose only changes are SYSTEM and
agent.wrap. The fixed controller source remains
`12c571cc53c7e03b051811477adc83e5518f157d`; it is not a replacement subject.

## Independence and output review

Live isolation passed before both releases and again immediately before each
release. First readiness contained exactly 31 permitted immutable input files;
second readiness added only the reviewed first fresh code and notes. Input Git
had one input-only commit, no remote, alternates, symlinks or unreachable old
objects. The separate verifier retained its 394 expected files and excluded the
16 original workflow disposable paths before reconstruction. All three controller
copies matched the fixed pin. Dependency/dependent implementations, manifests,
hooks, old history and prior attempts/reports stayed outside worker inputs.

The outside reviewer inspected actual retained workflow spec/Decisions, API/model,
eval identities/imports/test inventory and relevant fixture assertions, plus only
the declared dependency public interfaces. Every observed command and full script
was reviewed: 33 in attempt one, 37 in attempt two. The raw cumulative diffs were
projected into incremental reviews covering 147 and 93 trace snapshots, with
formatting normalized only in the outside review display. Final fresh bytes
matched those traces plus the logged formatting commands. The raw logs were not
modified. All immutable input hashes remained identical.

Before release two, the complete first trace, final fresh output and native
feedback were reviewed. Its prompt was exactly the original prompt plus native
feedback. No manual hint, import inventory, extra input, command interception,
filesystem/process instrumentation, model/effort change or worker-side build,
test or module setup was introduced. Attempt one queried Go environment metadata;
it did not compile or set up a module. Attempt two named two relative app eval
paths absent from the supplied tree; rg reported both missing, no app tree was
provided and no outside source access was observed. These commands and their
errors remain in the trace. No independence violation was observed.

Both workers exited zero; the wrapper recorded EOF and complete forwarding of
both raw streams. [Forwarding receipts](forwarding-1.json) and
[second receipt](forwarding-2.json) compare each durable stream byte-for-byte with
the corresponding native worker-log payload. Stdout sizes are 575 and 492 bytes;
stderr sizes are 2,078,941 and 593,781 bytes. Both raw streams and native logs are
retained, and the native report embeds them exactly. This invocation establishes
normal native failure-report/ref publication only; it establishes no native
reporting behavior after rejection/interruption or emergency termination.

## Gaps, preservation and stop

The native Contradicted Decisions section repeats the worker's “None identified”
submission. Its gap notes describe fresh implementation limitations for FLOW-001,
FLOW-073 and FLOW-117. The FLOW-117 note overlaps the retained historical-category
Decision and remains an unresolved review observation; the worker's submission is
not an independent finding that every Decision is satisfied. The missing go.mod
note describes the intentional worker/verifier separation, not an established
input defect. These notes authorize no correction. The two failing native cases
already prevent WF4 closure. No retained contract/eval defect or scoped follow-up
fix is established or authored by this retry review.

The owner added exactly 72 hours to the original 2026-11-07 22:51:05 UTC deadline;
the current deadline is **2026-11-10 22:51:05 UTC**. Authorization/deadline commit
`05f4b4f` changed only protected planning/deadline metadata and preserved a
byte-identical prior report at workflow-before-fixed-controller-retry-48a619b.md.
Worker command/model/effort, hold/stop/drain/reporting bounds, configured attempts,
timeouts, limits, archive counters and isolation gates were unchanged.

Native cleanup removed the input and verifier worktree. The inspector and native
controller exited; both worker groups and the controller group were absent.
The owned caller was removed only after retaining the report, full outside
artifacts and native reconstruction ref. No queued change or shared claim remains.
[Cleanup](cleanup.json), [entry snapshot](entry.json), final status and preservation
receipts retain the checks. Final status retains the existing B02 testcase-count
warnings for CHK-055 (16), APP-103 (17) and APP-104 (19), against limit 15.
Their tests and bounds are unchanged; no unrelated correction was undertaken.
All earlier reports, archives, proof/attempt refs,
controller/recovery executables and the unrelated linked worktree are preserved.
The current report has this exact native report as an unchanged prefix, followed
by a separately labeled owner review. No reconstructed source was applied, no
controller was promoted and no push occurred.

Stop after protected evidence/handoff commit. This single approval is consumed.
WF4 and full product completion remain open. Any further drill or correction needs
separate owner authorization, with fresh live gates and the current deadline.
