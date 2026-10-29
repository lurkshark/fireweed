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

## Active slice — approved 0080 recovery

0079's approved recovery applied at 1ddea1c with 191 full phase passes. The
owner now approved native abandon/re-author of 0080 and its one-line diagnostic
fixture correction. Native abandon committed c014190; earlier block/history
and attempt 1 remain retained. Only the authored change and retained evals were
saved for reuse, with no implementation oracle. Applied-since-review stays 1.

Reauthored 0080 keeps the same six CLI specs and no API changes. Restore the
retained evals and make only the omit-new hook exit consistently with its passing
report. Validate that this public fixture now returns C10/ALP-002 before
freezing it; then run normal fail-first, isolated audit and implementation.
See [the historical blocker and approved recovery](reports/next-0080-block.md).

After application, build NEW and run /tmp/fireweed-next-cli-smoke.py run using
/tmp/fireweed-next-cli-smoke-state.json. Its clean Go fixture remains retained
for that proof. Then prove bounded NEW run --max, later continuation and empty
queue. Retry/recovery and remaining kinds follow that runnable path. No NEW
unattended real-agent execution before the bounded fake-agent run passes. All
final gates and drill deferrals remain open; no new push or phase-write exception.

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
