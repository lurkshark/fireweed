# Rebuild progress

## Current handoff — 2026-10-28

Manual lifecycle, Go/none/local init, doctor/guidance CLI and bounded agent
invocation work through NEW. Queue orchestration through NEW next/run remains
unimplemented. The full design and all final ACCEPTANCE gates remain open.

The approved 0076–0078 corrections passed isolated audits, full phase checks,
fresh build/vet/race/static/history, the NEW doctor smoke and the same checkpoint
subagent's scoped re-review. The owner-authorized push synchronized main with
origin at `510cc91`. No pending change or claim remained at that boundary.
**Applied since review: 1 after 0079.** Earlier audit/phase deviations remain recorded;
the passing correction does not retroactively validate them. See the
[correction report](reports/checkpoint-corrections-0076.md) and
[checkpoint report](reports/checkpoint-0071-0075.md).

## Waiting prerequisite — fixed-profile auditor capacity

0079's approved recovery remains applied at 1ddea1c. 0080's owner-approved
native abandon/re-author completed, and its one-line fixture correction was
validated through public runner.Next: C10 naming ALP-002, with no C11. The
temporary validation test was removed; every permanent assertion/ID is retained.
Native evals commit 5e6d0bb reports 159 cases, six genuine runtime failures and
153 passes. **Applied since review: 1.**

The separate Sol/high auditor terminated with Selected model is at capacity
and no audit result. It only listed the isolated snapshot entry point before
failure; no passing audit or implementation is claimed. The snapshot was
removed and actual logs remain under run 20261029T012539-cf09ff80f620517ce580ef2f.
Native release dropped this worktree's own claim; 0080 remains committed at
phase audit with its run/IDs intact and no blocked field. The tree is clean.
See [the recovery/capacity record](reports/next-0080-block.md).

The goal explicitly requires stopping on unavailable fixed-profile access.
Resume with native start/context at the existing audit phase, then regenerate
the isolated snapshot and retry the same Sol/high auditor when authorized and
available. Do not switch models, fabricate an audit, rewrite phase state or
restart the completed evals. After a passing audit, reconstruct app wiring from
contracts, run full Done, build NEW and execute the prepared Go CLI smoke.
Its script/state remain /tmp/fireweed-next-cli-smoke.py and
/tmp/fireweed-next-cli-smoke-state.json. NEW next is not yet an applied CLI feature.

Bounded NEW run --max/continuation/empty queue follows that proof, then recovery
and remaining kinds. No NEW unattended real-agent execution before the fake-agent
run passes. Final gates/drill deferrals remain open. No new push occurred.

## Authority and deferred work

- Standing owner policy: after five applied changes, obtain a separate subagent
  checkpoint review; PASS authorizes pushing to origin and continuing, while an
  issue requires a human stop. Blocked-change stops, consecutive-module review,
  phase audit bounds and shared-base P07 remain enforced. Pushes do not reset
  review counts. Explicit exceptions remain scoped to their recorded use.
- Owner approval removed numeric spec-count caps for all change kinds. Keep
  related work within one module, atomic specs and normal checks/audits. The
  completed spec splits and D006 existing-path migration through 0068 need no
  new global rewrite. Explain real paths as justified features add them.
- Workflow stub.go reorganization is deferred unless a feature demonstrates a
  problem. Do not replace it with speculative cleanup.
- The only authorized contracts drill completed with 34 own/339 dependent cases
  and retained ref/cleanup. Its error-prefix gap and agent reads of dependent
  internals mean strict independence acceptance is unverified. No gap fix,
  repeat drill or other module drill is authorized; see
  [the report](reports/contracts-drill.md).
- Historical phase commits, counts, authorizations and acceptance mappings are
  preserved in [HISTORY.md](HISTORY.md). Old pending-work statements there do
  not restart completed work. Detailed reports preserve deviations and limits.

## Fixed execution context

Controller: `/Users/joel/Workspace/fireweed/.fireweed/bin/fireweed-bootstrap`
0.1.0, source `7f925190878e38fc38f30d16bb1cf47aca098cf7`, SHA-256
`f1516b5686435c23b54b71c5ac554aba6859b5454f072e0c5eb031dfd6f0fd4f`;
keep unchanged. History baseline: `be35552b706aa3000ff8f11306d129ff1d701b85`;
use explicit nonempty ranges. Set `GOCACHE=/tmp/fireweed-rebuild-cache`.

NEW: `/Users/joel/Workspace/fireweed/bin/fireweed`, built at clean source
`a34a32ac97c3c1a3ff571641117466963f338ca3`, CGO disabled. Subsequent commits
through 510cc91 changed only evidence. Reuse the
[manual lifecycle](reports/manual-lifecycle.md) and corrected
[doctor smoke](reports/doctor-setup-cli.md) evidence where valid.

Real workers and separate auditors use `codex exec --ephemeral
--ignore-user-config -m gpt-6-sol -c 'model_reasoning_effort="high"'
-s workspace-write {prompt}`. Coordinate native controller start/context/done;
workers never commit. Supply auditors isolated snapshots with only change,
projected specs, evals and permitted interfaces/rules; inspect their full command
logs. Record coordinator phase advancement honestly as agent=manual, retaining
actual CLI logs. Do not use native automatic audit dispatch after the documented
input/phase deviations. No new controller-owned phase-write exception exists.
