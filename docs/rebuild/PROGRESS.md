# Rebuild progress

## Current handoff — 2026-10-28

Manual lifecycle, Go/none/local init, doctor/guidance CLI and bounded agent
invocation work through NEW. NEW next now completes a checked spec-change lifecycle; bounded Run now also passes its actual Go CLI fixture. The full design and all final ACCEPTANCE gates remain open.

The approved 0076–0078 corrections passed isolated audits, full phase checks,
fresh build/vet/race/static/history, the NEW doctor smoke and the same checkpoint
subagent's scoped re-review. The owner-authorized push synchronized main with
origin at `510cc91`. No pending change or claim remained at that boundary.
**Applied since review: 1 after 0084; the 0079–0083 checkpoint passed.** Earlier audit/phase deviations remain recorded;
the passing correction does not retroactively validate them. See the
[correction report](reports/checkpoint-corrections-0076.md) and
[checkpoint report](reports/checkpoint-0071-0075.md).

## Active slice — saved-attempt recovery through NEW

The reviewed 0079–0083 batch pushed to origin/main at 8aff1e4 after separate
checkpoint PASS and fresh build/vet/race/static/history. Reuse the actual
[Next lifecycle](reports/next-0080-block.md), [bounded queue](reports/bounded-run-cli.md)
and [checkpoint](reports/checkpoint-0079-0083.md) evidence.

0084-verification-retries applied at 667e419: 214 full phase cases pass and a
valid isolated assertion audit passed RUN-015/021/022. Both capacity-stopped
workers were resumed by separately approved one-time same Sol/high retries.
The mistaken initial audit working directory is retained as invalid evidence;
the fresh correctly isolated audit is the accepted one. No pending change or
claim remains. See [the retry report](reports/verification-retries-0084.md).

Fresh NEW at clean 667e419, CGO disabled and vcs.modified=false, passed the actual
Go next retry fixture in 6.53s. One genuine impl failure was corrected on its
second attempt, preserving work and exact recorded verification feedback.
Four fake-agent calls, three phase commits, SPEC/archive/provenance, claim
release and clean tree were inspected; the fixture was removed. The external
verifier's over-specific output assertion was corrected against the same
completed fixture, without a product change or another execution.

Next change 0085-saved-attempt-blocking: exhausted verification work must be
saved on an attempt ref, checkout
restored, and a state-only block committed with claim release. Add only the
needed lifecycle primitive, return promptly to its runner caller and close
with an actual NEW next exhaustion fixture. Then projected blocking with
independent continuation, saved-attempt resume, nonzero-agent retries,
time/usage stops, remaining kinds and persistent bounds follow. Bounded fake
Next/Run proof has passed; full design and final ACCEPTANCE gates remain open.
No new push or drill occurred.

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
`667e419fc75c337650d4befaf1954a1826e8d21e`, CGO disabled. Reuse the
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
