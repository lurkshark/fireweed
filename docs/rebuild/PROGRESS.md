# Rebuild progress

## Current handoff — 2026-10-28

Manual lifecycle, Go/none/local init, doctor/guidance CLI and bounded agent
invocation work through NEW. NEW next now completes a checked spec-change lifecycle; bounded Run now also passes its actual Go CLI fixture. The full design and all final ACCEPTANCE gates remain open.

The approved 0076–0078 corrections passed isolated audits, full phase checks,
fresh build/vet/race/static/history, the NEW doctor smoke and the same checkpoint
subagent's scoped re-review. The owner-authorized push synchronized main with
origin at `510cc91`. No pending change or claim remained at that boundary.
**Applied since review: 3 after 0084–0086; the 0079–0083 checkpoint passed.** Earlier audit/phase deviations remain recorded;
the passing correction does not retroactively validate them. See the
[correction report](reports/checkpoint-corrections-0076.md) and
[checkpoint report](reports/checkpoint-0071-0075.md).

## Active slice — independent continuation through NEW run

The reviewed 0079–0083 batch pushed to origin/main at `8aff1e4` after separate
checkpoint PASS and fresh build/vet/race/static/history. Reuse the actual
[Next lifecycle](reports/next-0080-block.md), [bounded queue](reports/bounded-run-cli.md)
and [checkpoint](reports/checkpoint-0079-0083.md) evidence.

0084 verification retries passed 214 full phase cases and the actual NEW next
retry fixture; see [the report](reports/verification-retries-0084.md). It completed in 6.53s, preserving work and exact verification feedback,
with four fake-agent calls, three phase commits and clean claim/tree cleanup.

0085 saved-attempt blocking applied at `526f90e` with 268 full phase passes. The
sole FLOW-066 finding was corrected in native round 1, and its isolated scoped
re-audit passed. Same run/IDs were retained through two explicitly approved
author revert recoveries after ignored finding submissions. Those mechanical
passes remain invalid audit evidence; both phase-write exceptions are consumed.
The quote diagnosis, capacity retry and full phase history are in
[the dependency report](reports/saved-attempt-blocking-0085.md). No pending
change or claim remains, and no push after `8aff1e4` occurred.

0086 applied at c797264 with 213 full phase passes and isolated audit PASS.
The actual NEW next --json exhaustion fixture exited 1 in 5.61s with its E01
cause/ref, saved the final failed edits, restored a clean checkout, committed
only block state and released the claim. SPEC stayed unapplied and no archive
was created. Four fake calls, three branch commits and the saved commit were
inspected; fixture removed. See [the CLI proof](reports/exhausted-next-0086.md).
No governing change/claim remains and no push after `8aff1e4` occurred.

Next: implement projected blocking eligibility in the lifecycle's read/Start
selection so waiting changes can be skipped; return immediately to runner
Next/Run and prove independent continuation through actual NEW run. Status
views, saved-attempt resume, nonzero-agent retries, agent-commit undo, time/usage
stops, remaining kinds and persistent bounds follow. Full design and all final
ACCEPTANCE gates remain open; no new drill is authorized.

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
`c7972645f5079fd93c1267fbf5aee4eab7f21f5e`, CGO disabled. Reuse the
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
