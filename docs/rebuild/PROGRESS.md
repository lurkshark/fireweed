# Rebuild progress

## Current handoff — 2026-10-29

NEW now proves manual lifecycle, setup, bounded next/run, retries/blocking,
independent continuation, saved-attempt loading, agent-commit undo, usage limits
and total-time stops with owned resume. Latest actual Go proof: run --for2s
interrupts an agent in2.09s with original claim/run/phase/IDs/work and no commit;
next completes that same run in5.84s, later queue unchanged, clean/unclaimed.
See [0097 proof](reports/run-time-cli-0097.md). Full design/ACCEPTANCE remain open.

**Applied since review: 5 (0094–0098).** Mandatory separate Sol/high quality/goal
checkpoint plus fresh checks is pending. Last passing review0089–0093 was at
06a94e9; last authorized push origin/main82ef71d→28ced38. No next change or push
before this checkpoint passes. [HISTORY](HISTORY.md) owns the applied ledger.

0098 maintenance-start archived atd3272a3 with353 full passes and isolated
FLOW-077 audit PASS. Actual NEW next completed a Go maintenance refactor in3.26s
with one impl agent/one archive commit, retained SPEC/interface/evals and clean
claim release. Fixture removed. See [0098 proof](reports/maintenance-cli-0098.md).
0097 time stop and earlier usage/recovery proofs remain valid.

## Current checkpoint and next slice

Review0094–0098 quality and adherence to the full goal using a separate Sol/high
subagent; run fresh build/vet/race/static/history on the clean reviewed source.
PASS authorizes push to origin and continuation; any finding requires owner
stop. No pending governing change or claim. Full design/ACCEPTANCE remain open.
After PASS, continue remaining lifecycle kinds/bounds through NEW, starting with
its missing eval_fix semantics. No new drill is authorized by this checkpoint.

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
d3272a385b1eb63fec5fe14df37a889505a5f3ea, Go1.27.2 darwin/arm64, CGO disabled,
vcs.modified=false. This build completed without a warning.
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
