# Rebuild progress

## Current handoff — 2026-10-29

NEW now proves manual lifecycle, setup, bounded next/run, retries/blocking,
independent continuation, saved-attempt loading, agent-commit undo, usage limits
and total-time stops with owned resume. Latest actual Go proof: run --for2s
interrupts an agent in2.09s with original claim/run/phase/IDs/work and no commit;
next completes that same run in5.84s, later queue unchanged, clean/unclaimed.
See [0097 proof](reports/run-time-cli-0097.md). Full design/ACCEPTANCE remain open.

**Applied since review: 4 (0094–0097).** Separate Sol/high0089–0093 checkpoint
passed at clean06a94e9 with build/vet/race/static/history. Last authorized push
advanced origin/main82ef71d→28ced38. See [checkpoint](reports/checkpoint-0089-0093.md).
No pending change or claim. [HISTORY](HISTORY.md) owns the applied ledger.

0096 runner time dependency archived at2f6870f with262 passes;0097 app --for
archived atab7b317 with196 passes after one scoped APP-089 audit assertion
correction/re-audit. Actual NEW stop/resume proof passed; fixture removed.
Usage limit and earlier proofs remain valid; do not repeat them.

## Next slice — maintenance through NEW

The design sends maintenance changes (no spec edits/fixes/eval_fix) directly to
implementation. NEW Start currently defaults every fresh change to evals.
Establish that missing phase selection, then run actual NEW next on a small Go
refactor: one implementation agent, no evals/audit phase, full checks and one
archive commit, retained specs/interfaces/evals and clean claim release.
This will be the fifth application: separate quality/goal subagent review plus
appropriate fresh checks is required; PASS authorizes push and continuation.
Remaining kinds/bounds, full command inventory and final gates stay open.

## Authority and deferred work

- Fifth application since review requires a separate subagent quality/goal
  checkpoint plus appropriate checks. PASS authorizes push to origin and resume;
  any finding requires owner stop. Direct trusted approval explicitly covers
  this process. Blocked-change stops, consecutive module/behavior checks,
  native audit rounds and shared-base P07 remain enforced. Push alone resets
  no count; never change refs/bounds/specs/evals/counts to evade review.
- The owner approved scoped0093 discard/native abandon, fixture-only RUN-016
  eval_fix, fresh usage reauthoring and consecutive runner recovery. It is now
  complete. Former P08 runner warning is gone after app caller; existing C01
  embedded-recipe discovery and B02 CHK-055/16-case warnings remain explained.
- No new phase-write exception. Both0085 scoped exceptions are consumed.
  Historical deviations remain historical, not retroactively valid evidence.
- Only the contracts drill was authorized and has unresolved independence
  limitations, including dependent-internal reads. No gap fix, repeat or other
  module drill is authorized. See [drill](reports/contracts-drill.md).
- No workflow stub.go reorganization without demonstrated feature need.
  No numeric affected-spec cap; keep coherent single-module changes, atomic
  requirements and normal enforcement. Completed global splits/D006 migration
  need no rewrite. Add diagnostic explanations only for justified emitted paths.

## Fixed execution context

Controller: /Users/joel/Workspace/fireweed/.fireweed/bin/fireweed-bootstrap,
0.1.0 source7f925190878e38fc38f30d16bb1cf47aca098cf7, SHA-256
f1516b5686435c23b54b71c5ac554aba6859b5454f072e0c5eb031dfd6f0fd4f; unchanged.
Original history baseline be35552b706aa3000ff8f11306d129ff1d701b85; extend
recorded valid explicit nonempty ranges. GOCACHE=/tmp/fireweed-rebuild-cache.

NEW: /Users/joel/Workspace/fireweed/bin/fireweed, built at clean
ab7b3173871fa6552a4aaadb75ef755a7c0c9b60, Go1.27.2 darwin/arm64, CGO disabled,
vcs.modified=false. Optional stat-cache sandbox write warning did not fail build.
No new final race/vet/CI/doctor/fresh-checkout/drill/self-hosting claim follows
from this slice; use the last checkpoint's exact source/range honestly.

Workers and separate assertion auditors use codex exec --ephemeral
--ignore-user-config -m gpt-6-sol -c 'model_reasoning_effort="high"'
-s workspace-write {prompt}. Context before every step; workers never commit.
Auditors run in isolated snapshots with change/projected specs/evals/public
interfaces/rules, no SYSTEM/implementation/history/outside access or code
execution/writes. Inspect complete command logs and remove snapshots. Native
coordinator provenance is manual with real CLI logs retained. Do not use native
automatic audit dispatch after the historical input/phase deviations. Never use
bootstrap source/tests, rebuild-v1, abandoned/deleted/saved implementation or old
worker logs as oracles. Regeneration/drill bounds remain unchanged.
