# Workflow retry failure investigation

The final reconstruction compiled and reported 22 failing rows out of 657.
The rows reduce to three causes. This investigation used the retained report
and proof, the current contracts and evals, and two small scratch probes.
[Native report](../workflow.md), [terminal review](review.md).

| Cause | Workflow rows | App rows | Runner rows |
| --- | ---: | ---: | ---: |
| Retained report location | 0 | 5 | 12 |
| Abandon uses unsupported Git options | 2 | 2 | 0 |
| Returned X01 error omits violation detail | 1 | 0 | 0 |

## 1. Make retained report discovery a workflow contract

The rebuilt Verify writes `<phase>/<module>/report.json`. The current product
writes `.fireweed/runs/<run>/<phase>/<module>.json`, which the app and runner
fixtures open directly. Runner decodes the saved file as `checks.Report` and
checks the phase-specific case result and spec ID. The missing-file failures
occurred in those readers. The workflow eval passed because it walks the
entire run tree and checks for the expected case ID and hook output anywhere
under it. FLOW-052 specifies retention under the run directory. The existing
verification-cost Decision mentions neighboring module reports and
replacement, without naming the report format or location.
[FLOW-052](../../../packages/workflow/SPEC.md#flow-052),
[workflow eval](../../../packages/workflow/evals/lifecycle_test.go#L1266),
[workflow writer](../../../packages/workflow/internal/lifecycle/stub.go#L553),
[app reader](../../../packages/app/evals/next_cli_test.go#L211),
[runner reader](../../../packages/runner/evals/run_test.go#L199).

**Proposed change.** Make the stable run/phase/module lookup and the readable
JSON report shape explicit under workflow ownership. Keep the established
flat filenames for modules that have them. In the existing workflow retention
eval, open the file for the expected run, phase and module and check the mapped
fail-first or passing case and the retained hook output. Keep the app and
runner assertions. This brings the same discovery contract into the worker's
retained inputs and gives the owner module a direct assertion.

**Naming decision.** A pinned static check accepted a scratch module named
`costs`. Its flat report path is the same as the documented phase
`costs.json` snapshot. The current code writes the report first and the
snapshot afterward at that path, so making the flat pattern universal would
codify the overwrite. Preserve `costs.json` as the snapshot and select a
deterministic, documented report location for colliding module filenames
before approving the contract. One compatible option uses
`<phase>/reports/<module>.json` only for names that collide with `costs.json`
on the filesystem. A collision fix would need its own fail-first assertion
for coexistence. The scratch probe checked that the name is accepted. The
overwrite follows from the two writers' paths and ordering.
[Cost snapshot Decision](../../../packages/workflow/SPEC.md#decisions),
[cost snapshot writer](../../../packages/workflow/internal/lifecycle/costs_stub.go#L118),
[write order](../../../packages/workflow/internal/lifecycle/stub.go#L598).

## 2. Specify the X01 detail that callers receive

The checks hook returned an X01 diagnostic carrying `alpha: forbidden dependency`
and an aggregate error carrying a count. The reconstruction kept the
structured diagnostics and returned only `X01: 1 boundary violation(s)` as
the error text. FLOW-062 says lifecycle operations preserve assigned causes.
Its existing eval explicitly requires the X01 prefix and the concrete
violation text. The previous native retry had the same failure. The current
product incorporates the concrete violation into its error. Its separate
X02 behavior preserves the hook failure cause with one prefix and retains
structured diagnostics. [Native failure](../workflow.md#own-evals),
[FLOW-062](../../../packages/workflow/SPEC.md#flow-062),
[checks diagnostics and aggregate error](../../../packages/checks/internal/engine/boundaries.go#L99),
[workflow eval](../../../packages/workflow/evals/lifecycle_test.go#L1447),
[previous investigation](../workflow-native-retry-20261108T090949/investigation.md#b--boundary-violation-detail-is-lost-from-the-returned-cause-1-row),
[archived X02 correction](../../../changes/archive/0175-preserve-x02-boundary-cause.md).

**Proposed change.** Clarify FLOW-062 so a returned X01 boundary failure includes
the selected unsuppressed module-attributed violation message in the returned
error. Preserve existing error precedence and the X02 behavior. Reuse the
current X01 assertion through the normal audited lifecycle. No new error
allocation or exact error-string template is needed.

## 3. Carry the supported Git environment into the worker policy

The rebuilt Abandon applies a reverse patch with `--cached --3way`. The two
workflow cases and app's owned JSON case hit the same option rejection. The
app suite also emitted its failed parent row. A scratch Git-only probe on
Apple Git 2.30.1 reproduced exit 128. In the same fixture
`--reverse --3way --index`, as used by the current product, reversed the
patch in the index and worktree while leaving HEAD at the original commit.
This probe tests option compatibility. The existing Abandon evals verify
the product behavior. The same rejected combination is visible in the
reconstruction's attempt-restore fallback, but the final native resume
cases did not demonstrate a failure there. [FLOW-089 and FLOW-092](../../../packages/workflow/SPEC.md#flow-089),
[current Abandon](../../../packages/workflow/internal/lifecycle/abandon_stub.go#L104),
[public Abandon evals](../../../packages/workflow/evals/abandon_test.go#L88).

The Abandon behavior is already specified and asserted. This is a
reconstruction compatibility defect. The first attempt also used
`worktree list --porcelain -z`, which the same Git rejected before the worker
added a fallback in attempt two. If the owner approves a preparation change,
carry the existing supported-host policy into the retained worker SYSTEM
and identify the tested stock Git version of the selected verifier. Keep that
guidance about the supported environment and keep the existing Abandon/GC
assertions. The worker can choose compatible Git operations without access to
the previous implementations, reports or attempts. [Supported environment](../../../SYSTEM.md#scope-guardrails),
[worker policy](../../../docs/rebuild/workflow-drill/worker-system.md).

## Scope of the next decision

Start with the two workflow contract changes if approved. The report-discovery
decision accounts for 17 rows, and the X01 detail has failed in consecutive
drills. Treat the worker Git environment statement as a separate preparation
choice. The report-path change needs a filename collision decision before
authoring. Approval of any change would precede the audited workflow and
affected candidate verification. A fresh native proof is still a separate
approval gate.

The final run passed the direct RUN-037 verification-deadline case. Its
RUN-036/RUN-037 cross-change case failed when it opened the retained report,
which does not explain the earlier timing exception. The corrected FLOW-134,
RUN-043/044 and RUN-047/048 cases passed in this invocation. Final GC and the
app B03 status case passed as well. These outcomes give no reason to reopen
those changes for this result. [Native cases](../workflow.md#dependent-evals).

Both scratch probes were removed. Pinned static check passed, and idle Context
U02 and Verify P02 confirmed there was no held change. The stopped handoff
remains at the owner decision. [PROGRESS](../../../docs/rebuild/PROGRESS.md).
