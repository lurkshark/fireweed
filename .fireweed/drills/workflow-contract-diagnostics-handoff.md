# Workflow contracts and stop diagnostics — verified preparation

The owner-approved two-change sequence is complete and stopped. Source/candidate
`9a206c558e8e5625c0ae55bb57c42ae8130b4f15` is clean, validated and pinned. No repository native workflow invocation
or reconstruction worker was launched in this sequence. The preceding
[native failure](workflow.md) remains current; WF4/full completion remain open.
Original `79cf33a7690082f28c06bf41a72719d282bc00fe` and prior corrected `ccd3d6b7732d613e80a027ba7969b65adc58bc0e`
remain unchanged objects. No reconstructed code was imported or applied.

## Normal audited changes

[0183](../../changes/archive/0183-reserved-ids-review-fail-first.md) clarifies
FLOW-007 current/removed ID reservation and FLOW-023 review-item fail-first
exclusion as legitimately already_satisfied. Two focused public cases assert
returned/recorded allocation with current and archived removals and all-occurrence
handle replacement, and successful evals Verify with genuine enforced fail-first
but no fabricated review case. All original integration, negative fail-first,
quarantine, already_satisfied and abandoned-ID assertions remain. FLOW-062 and its
boundary cause-text test are unchanged. Evals commit 620017b passed 170 cases;
independent audit ac9e140 recorded No findings for FLOW-007/FLOW-023; apply/archive
595dbd2 passed 652 workflow/runner/app cases. No runtime/interface edit.

[0184](../../changes/archive/0184-verification-stop-diagnostics.md) is a diagnostic-only
RUN-037 eval_fix. Only the named verification-deadline case and its import change:
ordinary fixture reads now report agent invocations/work, hook marker, verification
report, clean-time-stop result, held state/assigned IDs/claim and HEAD before
assertions. The fatal marker read becomes a failing nonfatal assertion, so missing
marker evidence cannot suppress agent-work/no-commit/identity checks. All original
obligations remain, including marker content, agent work, no applied change,
no commit, phase/run/claim/ID/HEAD and root identity. The exact 2s For, 4s hook sleep,
request timeout and fake-worker/hook setup remain. Shared helpers and other tests
are unchanged. Evals df745fe passed 159 cases; independent audit b0bfbb9 recorded
No findings for RUN-037; apply/archive 9a206c5 passed 482 runner/app cases.
This improves failure evidence, not an established timing root cause or runtime
fix. The original native missing-marker cause remains unresolved.

The unchanged old pin governed both complete normal lifecycles. Fireweed owned
all six phase commits and both applications/archives; configured real author
workers and restricted auditors used the exact command/model/effort. There was
no manual phase advance, frozen-eval edit, rejected submission, counter reset or
replacement change. Raw [phase reports/logs](workflow-contract-diagnostics-handoff/runs/)
and [summary](workflow-contract-diagnostics-handoff/phase-summary.json) retain
actual invocations and results. Workflow quick Verify took about 149/152 seconds;
runner about 313/317 seconds, exceeding the five-minute own-verify target. This
cost limit is recorded, not waived or claimed passed; no bound was raised and
no unrelated cost correction was added.

## Identified clean source and idle promotion

A fresh no-hardlinks clone at the exact source above built with Go 1.27.2,
`-trimpath`, darwin/arm64. Build info records that revision and `vcs.modified=false`.
[Ordered validation results](workflow-contract-diagnostics-handoff/validation/results.json)
retain actual commands, fixed bounds, timestamps, outputs and exits for normal
and race Go suites, vet, static check, hook doctor (`--no-agent`), explicit history
and CI with hook doctor enabled. All gates passed. No validation retry or test/
runtime relaxation was needed. The delta is `ccd3d6b7732d613e80a027ba7969b65adc58bc0e..9a206c558e8e5625c0ae55bb57c42ae8130b4f15`;
valid committed WF2/CHK-034/residual-fix prefixes are reused as composed evidence.
This is not a fresh standalone whole-history run; history runtime is unchanged.

Two actual compiled CLI fixtures separately verify current-change removal and
an earlier archived removal reserve IDs. They also verify review-tagged additions
need no fail-first case while enforced ALP-001 supplies genuine failing evidence.
Start assigned ALP-003/004 and rewrote the handle; evals Verify succeeded without
a fake review case. Both fixture claims were released. Deterministic local fake
hooks ran; no reconstruction agent or repository native workflow invocation.
Ordinary Go suites retain their existing bounded fake-agent drill API/CLI fixtures;
those do not constitute a new owner workflow regeneration drill.

Existing B02 warnings remain for CHK-055 (16), APP-103 (17), APP-104 (19), against
15. These representative tests were not changed. Final live unapproved count is
5 (entry 3); exactly three phase commits were added for each new archive,
and all old counters/bounds remain. Protected author commits assert no human merge
approval. [Preservation](workflow-contract-diagnostics-handoff/preservation.json)
checks every entry tracked file outside the named authorized edits, prior refs,
all old reports/archives/proofs/attempts/configuration, recoveries and the unrelated
worktree. Native workflow report/ref bytes and identities are unchanged.

Only after both owning lifecycles and clean validation terminated without queue/
shared claim was the pin atomically replaced. New source is `9a206c558e8e5625c0ae55bb57c42ae8130b4f15` and SHA-256 is
`cc22f47c7aa607c96a13273d29d5bf4ade4f912453840f621331d73dc385f112`. Retiring source ccd3d6b bytes remain at `.fireweed/bin/fireweed-ccd3d6b`, SHA-256
`20278504a8ea89fcd5f23d1de9a6ec849389fd55707054992c59967eec65e2b1`; every older recovery is unchanged. The
[promotion receipt](workflow-contract-diagnostics-handoff/promotion.json) records
idle conditions and the expected P02 idle Verify. U02 idle Context/P02 Verify are
idle-state checks, not product validation. An initial context call without its
required change argument returned U01; the corrected named lookup returned U02.
Both read-only receipts are retained; neither is counted as a validation gate.

## Prepared retry and stop

Outside tooling changes only candidate/expected-pin metadata. Exact configured
worker command/model/effort, worker SYSTEM, wrapper/inspector behavior, retained
input policy, live immutable-input/verifier separation checks, review-before-
release, two-attempt cap, stop/drain/publication windows and whole-run emergency
stop remain unchanged. No import inventory, hints, interception, instrumentation
or expanded inputs. Existing deadline remains **2026-11-10 22:51:05 UTC**.
Earlier bounded fake rehearsals remain applicable to unchanged supervisor behavior;
this sequence establishes no new native reporting on rejection/interruption/
emergency stop. Normal native failure publication belongs only to prior actual
reports. No publication outcome is inferred from preparation.

Only inspector `prepare` ran. It made a clean caller at the identified source,
committed setup child `644f51f1090f9c8e5669e133bb98a7248a68ef5d` changing only SYSTEM/agent.wrap,
and copied the exact new pin. [Preparation](README.md#historical-evidence) (`d63c16e:.fireweed/drills/workflow-contract-diagnostics-handoff/preparation.json`),
[expected hashes](README.md#historical-evidence) (`d63c16e:.fireweed/drills/workflow-contract-diagnostics-handoff/prepared/expected.json`) and
[setup patch](README.md#historical-evidence) (`d63c16e:.fireweed/drills/workflow-contract-diagnostics-handoff/setup-child.patch`) retain the
result. Primary/caller hashes and retained bytes matched; removing only the wrap
line recovers exact source configuration. No gate, verifier, worker or native
drill invocation was created. The verifier's third pin copy and actual input
isolation remain future live checks, not facts established here. All owned
validation/fixture/preparation clones were removed after claim/worktree checks;
[cleanup](workflow-contract-diagnostics-handoff/cleanup.json) is retained. Test/
CLI commands and owning drivers are terminal; no queue/shared claim remains.

**Stop for separate owner approval of one workflow retry on the source/pin above.**
Follow [GOAL](../../docs/rebuild/GOAL.md) and the
[inspector README](../../docs/rebuild/workflow-drill/README.md), prepare a new outside
root and repeat fresh live gates. Review actual retained inputs before first
release, every command/script/incremental patch, full native feedback and fresh
output identity before any configured second release. Keep received logs durable
outside reconstruction and forward identical bytes; on rejection/interruption use
the bounded worker-first stop and give the controller only its configured bounded
publication opportunity. Retain actual evidence/ref or record its absence honestly,
clean only owned scratch, commit protected evidence idle and stop. Failure/gaps
authorize no additional correction or invocation. No further pin update, later
work, source import, push or release is authorized now.
