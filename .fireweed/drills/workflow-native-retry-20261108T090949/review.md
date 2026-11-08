# Supervised native workflow retry — failed, stopped

The owner's instruction authorized exactly one retry despite the stopped handoff.
One pinned native `drill workflow --json` invocation ran with two configured
reconstruction attempts. It exited 1, published the failure report and ref, and
cleaned its native input/verifier. The inspector recorded neither rejection nor
emergency escalation. No reconstruction was applied to the product, and no
correction, replay, third attempt, controller change or push followed.

## Candidate and preparation

The validated candidate was `a26f7b9909f6c311416a0706d1b16749d16b9146`, matching
the unchanged controller source recorded in [the pin handoff](../controller-pin-a26f7b9.md).
The isolated no-hardlinks caller's setup child was
`2e515d11bf83a721f576f788df3d44a1255e1ed4`; its only differences from the candidate
were the existing reconstruction SYSTEM policy and outside `agent.wrap` prefix.
Product/spec/eval/interfaces/manifests/hooks were unchanged.

Protected preparation commit `368bafb` updated only the inspector's candidate,
expected pin and exact worker command to the current SYSTEM values. Supervisor,
worker policy, hold/stop/drain/publication windows, two-attempt cap, configured
bounds and original extended deadline were unchanged. The host fake-process
[rehearsal](rehearsal-results.json) passed; it proves wrapper fixture outcomes,
not a native pass or generic native reporting guarantee.

The completed WF1 review, prior WF2 normal/race/vet/CLI/history/CI preparation,
and audited 0185–0187 phase checks were reused through the verified pin handoff.
The history baseline remains `be35552b706aa3000ff8f11306d129ff1d701b85`; the candidate
range is `be35552b706aa3000ff8f11306d129ff1d701b85..a26f7b9909f6c311416a0706d1b16749d16b9146`.
No new product verification claim replaces those records or their timing exceptions.

## Independence and attempts

Before first release, the outside reviewer read all 31 actual retained files:
workflow SPEC/Decisions, facade/model and public evals, reconstruction SYSTEM,
and only contracts/gitstore/assets/project/checks public facades/model interfaces.
Live preflight checked exact input bytes, no symlinks, one input-only Git commit,
no remotes/alternates/unreachable objects, separate verifier with 536 retained
files unchanged and 16 original disposable files deleted, and all three controller
copies matching the pin. Original/deleted code, history, reports, binaries,
bootstrap and dependency/dependent internals were excluded from worker inputs.

Both workers used exactly the configured `codex exec --ephemeral
--ignore-user-config -m gpt-6.1-sol -c 'model_reasoning_effort="high"'
-s workspace-write {prompt}` command. All emitted commands, complete scripts and
incremental patches were reviewed during execution. Reads stayed within supplied
inputs and fresh output; writes stayed within workflow disposable files and
DRILL_NOTES. No local builds/tests/module setup, outside reads, retained-input
edits, worker commits or lifecycle mutations were observed. The inspector supplied
no implementation hints.

Attempt one returned 11 fresh Go files and notes. Native compilation failed on
`prbody.go:303:29: undefined: contracts`; each module's PackageFailure/no-coverage
feedback was a compilation outcome, not behavioral evidence. The full first
feedback and trace were reviewed before the final configured release. The second
prompt was exactly the original plus native feedback, and the same live isolation
check passed again. Final first-output bytes matched the reviewed emitted patches
after formatting. Attempt two changed only the missing import and notes, then
returned without running builds/tests. [First](first-trace-review.json) and
[second](second-trace-review.json) reviews retain these checks.

Both workers exited 0, and every received stdout/stderr byte reached EOF and was
forwarded unchanged to its corresponding native pipe. The [stream/gate receipt](stream-and-gate-receipt.json)
retains all gate records and exact byte offsets, sizes and digests for all four
raw streams inside the single native report. Readiness prompts are also exactly
recoverable from that prefix; reconstructed readiness JSON hashes were checked.
No duplicate transcript, copied tooling or source snapshot is committed.

## Native result review

The [current report](../workflow.md) preserves its exact native prefix, followed
by outside review. [Summary](verification-summary.json) records native fields,
case counts, coverage and the actual terminal diagnostic. Counts include parent
and subtest rows exactly as emitted, without inventing independent test counts.

| Module | Rows | Pass | Fail | Enforced IDs referenced | Suite seconds |
| --- | ---: | ---: | ---: | ---: | ---: |
| workflow | 170 | 163 | 7 | 133/133 | 151.578 |
| app | 323 | 316 | 7 | 155/155 | 261.980 |
| runner | 159 | 156 | 3 | 67/67 | 262.793 |

All 652 rows were accounted for: 635 pass, 17 fail, no skipped/error rows or
missing enforced IDs. Reference coverage includes failures and is not passing
proof. Second-attempt compilation succeeded. All final failures were reviewed:

- Six workflow GC rows (FLOW-128–132) reject reconstructed `git worktree list --porcelain -z` because the installed Git rejects `-z`; FLOW-062 loses the boundary detail, returning `X01: 1 boundary violation(s)`.
- Five app GC rows (APP-014, APP-150/151) report the same Git-option failure. Two audit lifecycle rows (APP-010/029/030/031) report approved-author rejection or absent committed blocking state.
- Runner's RUN-047/048 row reports altered retained Provenance contents; RUN-043 and RUN-044 omit or misorder dependent failure details. RUN-037 passed in this invocation; that does not resolve its earlier timing cause.

These are observations from this reconstruction and native report, without a
follow-up diagnostic replay or product/eval correction. The worker identified no
contradicted Decisions. Its FLOW-119 note describes the import correction, and
FLOW-003 expresses its then-unverified state; the outside verifier subsequently
compiled and failed the cases above. Those notes establish no agreed spec gap and
authorize no fix. No passing independent workflow proof, WF4 closure or full
completion is claimed. Earlier own-verification timing exceptions remain open.

## Retention, cleanup and stop

Retained native ref: `refs/fireweed/drills/workflow/20261108T090949.204125000`,
OID `e0dc010ce1b5a06cdb4f4d87a290b7bbb5016475`. [Proof review](proof-review.json)
checked all 536 retained files unchanged and final reconstructed Go bytes equal
to the reviewed worker output. Only that exact evidence ref was copied locally;
no code was checked out or applied. The preceding report remains accessible with
`git show 368bafb:.fireweed/drills/workflow.md`; all prior refs/reports/archives remain.

[Terminal](terminal.json) and [cleanup](cleanup.json) confirm no rejection/emergency,
absent owned process groups, removed native input/verifier and caller, empty
queue/claims, unchanged pin/configuration and prior refs, and untouched unrelated
worktree. Final pinned static and explicit preparation-history checks passed.
Context U02 and Verify P02 are expected idle-state checks, not product verification.
Author-written documents pass `git diff --check`; its remaining whitespace findings
are exact native transcript/diff bytes deliberately retained without normalization.
The final status remains zero unapproved with the existing B02 warnings; no bound,
counter, recovery executable or enforcement policy was reset or weakened.

Approval is consumed. Stop with no active run, queue or shared claim. Further
correction, investigation/replay, retry, source application, repin, later work or
push requires another owner-selected scope. The deadline remains
2026-11-10 22:51:05 UTC.
