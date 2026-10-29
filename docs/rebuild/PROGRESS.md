# Rebuild progress

## Current handoff — 2026-10-29

NEW now proves manual lifecycle, Go/none/local init and doctor/guidance, bounded
next/run, verification retry, saved blocking, independent continuation,
saved-attempt loading, agent-commit undo, and clean usage-budget stop with owned
resume. The latest actual Go proof passes: Alpha waits/retries/applies, Beta
retains its claim/work/run/ID at the shared budget, then next completes the same
Beta run; Gamma stays queued, final tree/claims clean. See
[0095 proof](reports/usage-stop-cli-0095.md). Full design/ACCEPTANCE remain open.

**Applied since review: 3 (0094–0096).** Separate Sol/high 0089–0093 checkpoint
passed at clean06a94e9 with build/vet/race/static and40-commit history extension
(zero history diagnostics). Authorized push advanced origin/main82ef71d→28ced38.
See [checkpoint](reports/checkpoint-0089-0093.md). No pending change or claim.

0094 runner usage control applied at a720cf6: RUN-029–035 and qualifiedRUN-015,
247 full passes after one native RUN-033 approved-policy assertion correction
and isolated re-audit PASS. Fresh implementation corrected one attempt-count
error and passed69 own cases. [Report](reports/usage-limit-runner-0094.md).
0095 app caller applied at dbda392: APP-076/077/079/083/084/086,186 full passes,
one scoped naming correction/re-audit, clean exit3/data/reason/guidance and C01
config exit2. Actual NEW run/next proof passed6.25s/4.62s with eight agent calls
and six phase commits. Fixture removed. [Report](reports/usage-stop-cli-0095.md).
Earlier capabilities/evidence and the approved0093 abandon/eval_fix recovery
remain in [HISTORY](HISTORY.md) and linked reports; do not repeat valid proofs.

## Next slice — run --for time limit

0096-run-time-limit archived at2f6870f: RUN-036–040 plus qualifiedRUN-016/030,
262 full passes and isolated seven-ID audit PASS;76 own passes in194.04s.
Incomplete pre-freeze timeout was superseded by complete native evidence.
Only For/IsTimeStop were added publicly; total deadline, interruption, wait
preflight and mixed-failure control are private. See
[runner report](reports/run-time-runner-0096.md). No pending change/claim.

Next: wire app run --for duration parsing and clean time-stop output immediately,
then build NEW and run the prepared actual Go fixture: stop an agent while
retaining exact claim/run/phase/IDs/work, then complete the same run through
next. The time fixture has not run yet; installed NEW is still dbda392.
Remaining kinds/bounds, command inventory, full design and final gates stay open.

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
dbda392145c7c23f41dbb6f645deda048b8c4284, Go1.27.2 darwin/arm64, CGO disabled,
vcs.modified=false. Optional stat-cache sandbox write warning did not fail build.
No new final race/vet/CI/doctor/fresh-checkout/drill/self-hosting claim follows
from these two changes; use the last checkpoint's exact source/range honestly.

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
