# Rebuild progress

## Current handoff — 2026-10-29

NEW now proves manual lifecycle, setup, bounded next/run, retries/blocking,
independent continuation, saved-attempt loading, agent-commit undo, usage limits
and total-time stops with owned resume. Latest actual Go proof: run --for2s
interrupts an agent in2.09s with original claim/run/phase/IDs/work and no commit;
next completes that same run in5.84s, later queue unchanged, clean/unclaimed.
See [0097 proof](reports/run-time-cli-0097.md). Full design/ACCEPTANCE remain open.

**Applied since review: 1 (0099).** Separate Sol/high0094–0098 checkpoint PASS
at clean9949dfe with build/vet/race/static/history (33 commits, zero history
diagnostics). [Checkpoint](reports/checkpoint-0094-0098.md). Authorized push
advanced origin/main28ced38→7e0da27. No pending change or claim; ledger in
[HISTORY](HISTORY.md). Full design/ACCEPTANCE remain incomplete.

0099 eval-fix lifecycle archived at23ad9de with361 full passes, isolated
four-ID audit PASS and89 own passes. Actual NEW Go next corrected a genuinely
wrong ALP-001 eval in4.74s, passed evals/audit/no-op impl with three checked
commits, preserved contracts/implementation and released its claim cleanly.
Fixture removed. See [0099 proof](reports/eval-fix-cli-0099.md). Maintenance,
time/usage stop and earlier recovery proofs remain valid.

## Next slice — eval-fix worker guidance

The current embedded impl playbook directs workers to disposable code, while
eval-fix impl permits only the change file. Add scoped conditional eval-fix
instructions to evals/audit/impl playbooks: correct evals only, retain listed
coverage, all module evals pass, audit assertions, impl changes only the change.
Expose the resulting guidance through actual NEW playbook commands. No new
backend/API, behavior matrix, drill or unrelated prose cleanup. Then continue
remaining lifecycle kinds/bounds and final design inventory.

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
23ad9de2a67c22be063ac860bd48175c74c85ceb, Go1.27.2 darwin/arm64, CGO disabled,
vcs.modified=false. This build completed without a warning.
Last checkpoint build/vet/race/static/history source is9949dfe;0099 native
full checks and actual proof pass. No fresh final race/vet/CI/doctor/checkout/
drill/self-hosting claim follows from0099.

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
