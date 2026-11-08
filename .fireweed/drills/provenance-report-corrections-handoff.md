# Provenance and report contracts — audited, applied and stopped

The owner authorized creating the changes from the proposed provenance-prefix and
dependent-report resolutions and running the change process. Both selected changes
are complete. No reconstruction defect was fixed, no drill was run and no controller
was repinned. WF4 and full completion remain open.

Authoring began at `ba58a0a4e751cdb72bd314706169346b9abfa67d`. Scope and queue were
committed separately at `b9c1e29` and `063f081`. Applied source is
`9f4770abab40f1fe9bd8f5a0c62f44e8484dba1b`.

## Process and results

The unchanged pinned controller ran `fireweed run --max 2 --json` with its directory
first on PATH and GOCACHE=/tmp/fireweed-rebuild-cache. Normal base approval applied;
no --allow-unmerged, bound change, counter reset, replacement change or retry-budget
reset was used. Configured ephemeral workers used gpt-6.1-sol/high, and independent
restricted auditors used gpt-6-astra/medium, with --ignore-user-config and
workspace-write. Workers used context and quick Verify, returned without committing
or invoking lifecycle mutators, and made no impl edits. The outside controller
verified every phase, made all six phase commits, applied Specs, archived changes
and released claims. The run exited 0: two applied, empty true, limited false, no
diagnostics. [Exact terminal result](provenance-report-corrections/run.json).

| Change | Result | Evals / audit / apply commits | Evals cases | Full impl cases |
| --- | --- | --- | ---: | ---: |
| [0188](../../changes/archive/0188-provenance-and-complete-reports.md) | FLOW-051 clarification and new FLOW-134 complete-report contract; legitimate already_satisfied established against unchanged runtime. | `811c2ed` / `3caf7ba` / `bdf43d4` | 175 | 657 |
| [0189](../../changes/archive/0189-provenance-preservation-eval.md) | Audited RUN-047/RUN-048 eval_fix; contracts unchanged. | `933c87b` / `b626257` / `9f4770a` | 159 | 482 |

Both audits returned No findings for all changed IDs, with no findings round trip.
The second auditor attempted relative local Verify and received exit 127 because
the restricted input workspace intentionally contains no controller or project
configuration. This was no executed product verification; its assertion review and
Audit submission succeeded, and the owning controller validated the audit before
advancing. Both temporary audit workspaces were observed removed.

0188 tests prior valid provenance entries separated by blank lines, preserving each
nonblank line byte for byte and appending exactly one entry after each successful
Done. Original phase/date/manual/configured-agent/version/hash/count, context and
retained-report assertions remain. Four full impl cases cover own fail/direct,
own fail/all, own fail/none and own pass/all. They require every selected mapped
report and its failure sentinel, exclude unselected invocations, and preserve HEAD
and the held impl state. Valid failed reports use ordinary nonzero hook exits.

0189 retains insertion-only and exact author-section checks. A first permitted
Audit edit exits 9, then a fresh invocation captures the imported input. Raw
Provenance equals the original before any successful audit Done. Four ordered
agent results and their exit codes are asserted. Only the final archive comparison
ignores blank lines, preserving nonblank line text and order exactly. All other
audit evals and retry_failure_details_test.go remain unchanged, including diagnostic
identity/order, pass/skip exclusion, Unicode bounds and truncation markers.

Only workflow SPEC.md and the two selected public eval files changed in packages.
Runtime, interfaces/models, runner specs (including RUN-043/RUN-044/RUN-047/RUN-048),
FLOW-025/FLOW-026, Decisions, hooks/configuration, SYSTEM and inspector tooling are
unchanged. Existing quick/cancel/timeout/hook/report-error and rejection evals remain.

## Verification and limits

[Static check](provenance-report-corrections/check.json) passed without diagnostics.
[History check](provenance-report-corrections/history.json) covers exactly the six
phase commits in `063f081..9f4770a`, with no diagnostics or squashed entries.
`git diff --check ba58a0a..9f4770a` passed. Required full phase checks and boundary
verification passed; valid Go cached results were reused normally.

Committed final reports retain all cases and hook output:
[workflow](provenance-report-corrections/workflow-impl.json),
[runner](provenance-report-corrections/runner-impl.json), and
[app](provenance-report-corrections/app-impl.json).
The [verification summary](provenance-report-corrections/verification-summary.json)
records coverage and observed worker quick-check durations. Workflow has 175 passing
rows covering 134 current IDs, runner 159 covering 67, and app 323 covering 155.
No missing/unknown IDs, nonpassing rows or duplicate classname/name pairs were found.
Case rows include ordinary parent/subtest rows rather than independent behaviors.

Worker quick Verify took 176.838s / 179.394s for workflow evals/impl and
293.596s / 303.486s for runner evals/impl. Runner impl exceeds the five-minute
own-Verify target by 3.486s. Earlier timing exceptions remain unresolved; this
sequence does not close them. No limit was raised, assertion removed, extra benchmark
run or performance change attempted. Each complete impl phase, including its worker
and all controller checks, lasted under seven minutes, bounding every full Verify
below the unchanged 25-minute limit.

[Final status](provenance-report-corrections/status.json) has an empty queue, two
unapproved applied changes against the unchanged cap of ten, and three phase commits
for each new archive. Existing B02 warnings remain CHK-055 16, APP-103 17 and APP-104
19 against the unchanged 15-case threshold. These warning sources remain outside
scope. Derived workflow/runner drift is now 1/27; no passing drill follows.

## Preservation and stop

The owning run exited and no shared claim remains. Idle context returned U02 (no
valid slugs) and idle Verify P02 (worktree holds no change); these are state checks,
not product verification. The working tree was clean before protected handoff
writing. No linked worktree was added; the unrelated stub-handoff-guidance worktree
remains at `b59e9cc15a7d6c82f6436ef72fe5c2c3ccded211`.

The controller pin remains source `a26f7b9909f6c311416a0706d1b16749d16b9146`, SHA-256
`962149ff80fb5f6e396dbafda4f44fb47bb28840d96d37b3c7c9718444c7bf5a`.
No fallback, recovery promotion, source import or push occurred. Prior evidence,
archives, attempt/proof refs, recoveries and bounds/counters were preserved.
The failed workflow report and retained evidence are unchanged; native ref
`refs/fireweed/drills/workflow/20261108T090949.204125000` still identifies
`e0dc010ce1b5a06cdb4f4d87a290b7bbb5016475`.

The two-change scope is consumed. Stop with no active run, queue or claim; no further
fix, replay, drill, repin, later milestone or push is authorized. WF4/full completion
remain open. Deadline remains **2026-11-10 22:51:05 UTC**. Follow GOAL/PROGRESS.
