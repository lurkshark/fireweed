# Rebuild progress

## Current handoff — 2026-10-28

Manual lifecycle, Go/none/local init, doctor/guidance CLI and bounded agent
invocation work through NEW. NEW next now completes a checked spec-change lifecycle; bounded Run now also passes its actual Go CLI fixture. The full design and all final ACCEPTANCE gates remain open.

The approved 0076–0078 corrections passed isolated audits, full phase checks,
fresh build/vet/race/static/history, the NEW doctor smoke and the same checkpoint
subagent's scoped re-review. The owner-authorized push synchronized main with
origin at `510cc91`. No pending change or claim remained at that boundary.
**Applied since review: 0; the 0079–0083 checkpoint passed.** Earlier audit/phase deviations remain recorded;
the passing correction does not retroactively validate them. See the
[correction report](reports/checkpoint-corrections-0076.md) and
[checkpoint report](reports/checkpoint-0071-0075.md).

## Active slice — NEW verification retries

The authorized single Sol/high audit retry passed. 0080 applied at a11d713 with
159 full phase cases passing; no pending change or claim remains at that boundary.
Fresh NEW build source is a11d713e13922087f2172ab5665570a80db1ec0c, CGO disabled
and vcs.modified=false. The actual NEW next Go fixture passed in 5.73 seconds:
one genuine Go eval failure became a pass, three checked phase commits applied
SPEC/archive/provenance, separate fake agents ran exactly three times, and claims
released with a clean tree. See [the phase and CLI proof](reports/next-0080-block.md).
The completed disposable fixture was removed after inspection.

0082 applied at 58351b2 with 166 full phase passes and an isolated audit pass.
The actual NEW run --max 1 applied Alpha in 6.35s, a later run applied Beta and
Gamma in 9.16s, and an empty run added no calls or commits. All exited 0 with
empty stderr; nine checked phase commits, real Go fail-first/pass, provenance
and claim cleanup were inspected. The owned fixture was removed. See
[the bounded-run proof](reports/bounded-run-cli.md).

0083 applied at 6ce22e9: the passing audit example now shows No findings on
one line and a populated Checked IDs field on the next. Native first findings
returned to evals with audit_rounds 1; a separate author clarification preserved
the run/round, corrected the assertion, and AST-028-only isolated re-audit passed.
All 461 full phase cases pass; no pending change or claim remains. See
[the audit-example record](reports/audit-example-0083.md).

The separate subagent checkpoint review returned PASS for 0079–0083. Fresh
build/vet/race/static/history checks also passed at clean 6ce22e9; history covers
38 new commits with zero diagnostics. The NEW binary has CGO disabled and
vcs.modified=false. The standing push and continuation are authorized; see
[the checkpoint report](reports/checkpoint-0079-0083.md). Full acceptance remains
open.
Retries/saved attempts, blocking/independent continuation, resume, time/usage
stops and remaining kinds follow the runnable queue path. No NEW unattended
real-agent execution before the bounded fake-agent run passes. Full design and
final ACCEPTANCE gates remain open; drill deferrals and normal bounds remain.
The reviewed batch pushed successfully to origin/main at 8aff1e4. The next
change is 0084-verification-retries: bounded retries after failed verification,
with retained work and last-output feedback, through the existing NEW next/run
path. Close this slice with an actual NEW next whose first implementation is
wrong and second succeeds. Saved attempts/blocking, nonzero-agent retries and
usage/time stops remain subsequent work.

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
`6ce22e9e9d618e7c5c4451f48b3f1b522ab4cf08`, CGO disabled. Reuse the
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
