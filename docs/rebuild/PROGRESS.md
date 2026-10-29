# Rebuild progress

## Current handoff — 2026-10-29

NEW proves manual lifecycle and setup, bounded next/run, retries/blocking,
independent continuation, saved-attempt loading, agent-commit undo, time/usage
stops with owned resume, maintenance and eval-fix. Latest actual Go proof:
next corrects a wrong eval under the audited eval-fix path, preserving contracts
and implementation, with three checked commits and clean claim release.
See [0099 proof](reports/eval-fix-cli-0099.md). Full acceptance remains open.

**Applied since review: 2 (0099–0100).** Separate Sol/high0094–0098 checkpoint
PASS at clean9949dfe with build/vet/race/static/history. [Checkpoint](reports/checkpoint-0094-0098.md).
Last authorized push origin/main28ced38→7e0da27. No pending change or claim;
applied ledger in [HISTORY](HISTORY.md).

0100 eval-fix playbooks archived atf418f88:569 full passes within15m,
isolated three-ID audit PASS,59 own passes. Actual NEW evals/audit/impl
playbook JSON commands return the conditional guidance, correct envelope and
no diagnostics/stderr. See [0100 proof](reports/eval-fix-playbooks-0100.md).
This closes the two-change eval-fix slice; do not extend the same behavior
for speculative validation. Earlier CLI proofs remain valid.

## Next slice — regenerate lifecycle

The design sends regenerate:true straight to impl with disposable files removed,
uses retained contracts for reconstruction and warns without a sufficiently
recent passing drill. NEW currently sends regenerate to evals and leaves code
in place. Establish this missing path through normal workflow changes, return
to its existing CLI caller and prove a bounded fake-agent Go reconstruction.
No actual drill is authorized here. Keep warning explanations accurate when
a justified diagnostic is emitted. Remaining kinds/bounds, full command
inventory and final gates remain open.

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
f418f880a9adc2577ffbb8555147c31a64ab94cc, Go1.27.2 darwin/arm64, CGO disabled,
vcs.modified=false. This build completed without a warning.
Last checkpoint build/vet/race/static/history source is9949dfe;0099–0100 native
full checks and actual CLI proofs pass. No fresh final race/vet/CI/doctor/
checkout/drill/self-hosting claim follows from these changes.

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
