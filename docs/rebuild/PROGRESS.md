# Rebuild progress

## Current handoff — 2026-10-30

NEW proves manual lifecycle/setup, bounded next/run and recovery/resume,
agent-commit undo, time/usage stops, maintenance/eval-fix/ordinary-fix/regenerate,
and release/commits/abandon. Latest [actual Go abandon proof](reports/abandon-cli-0111.md)
retains a real saved attempt, restores phase work/deletes pending change in one
commit, releases its claim and frees ALP-002 for Start reuse; fixture removed.
Full design/ACCEPTANCE remain incomplete.

**Applied since review: 3 (0117–0119); scoped 0120 recovery approved.**
Prior fixed Sol/high quality/goal review of0112–0116 PASS and fresh checks passed
at929e32b; authorized origin push e5287d5→1ca04f3 completed. [Gate](reports/checkpoint-0112-0116.md).
0117 CI remains local with247 native passes and actual NEW Go proof PASS.

Corrected0118-list-queue-command now exposes list through project.Queue. Native
250 passes at clean a47a7fd, fresh restricted independent audit PASS, fresh
implementation worker and actual NEW queued/held/empty read proof PASS; fixture
removed. [List evidence](reports/list-cli-0118.md). Original eval-only0118-list-command
was natively blocked for the unassertable fallback and plain-output gap, then
owner-approved native abandon78705e1 removed it. The corrected contract resolves
only the fallback condition; fresh public evals carry the empty-reason assertion.
No saved implementation existed or was reused. Original findings remain historical.

0119 FLOW-097–100 applied at00a4a30: recorded Verify/Done cost snapshots,
B01–B03 warnings and public Costs reader;450 native passes and isolated audit
PASS. [Evidence](reports/verification-costs-0119.md). A scoped implementation
review correction added root protected-pattern precedence before application;
interrupted verification is not passing evidence. Main tree clean/unclaimed.
0120 status data reached impl after457 native cases (7 intended failures/450
passes), one persisted FLOW-113 correction and restricted re-audit PASS. Fresh
implementation passes6/7 focused cases: FLOW-101's frozen fixture expects evals
although its maintenance Start records impl. The worker stopped rather than
misreporting phase or editing frozen evals. [Checkpoint](reports/status-data-0120-checkpoint.md).
Its own claim was natively released; only two implementation files remain
uncommitted, phase/run/IDs/round unchanged. No archive/full impl Done, app caller,
NEW status proof, fifth review or push ran before the checkpoint.

Owner: "Approve status slice and fifth review". This supersedes the completed
recovery's stop only for `/tmp/fireweed-status-continuation-checkpoint.md`:
workflow cost recording/B01–B03, workflow status data with projected waiting
reasons, then immediate app status wiring/actual NEW proof. Three applications
reach the fifth-change quality/goal subagent review and appropriate fresh checks;
push only on PASS, then stop and report. Stop earlier on any block. No phase
exception, drill, deferred gap fix or bound/count/ref change. Native status
confirmed unapproved2/P08 before this slice. Cost recording is complete;0120
is pending at the wrong frozen assertion. The reviewable public correction and
native recovery proposal are in `/tmp/fireweed-0120-frozen-phase-checkpoint.md`;
owner now replied "Approve recovery and revised slice". Discard only the two
uncommitted0120 implementation files, natively abandon/re-author with the exact
public evals→impl assertion correction, and use fresh workers without discarded
implementation/log reuse. Corrected workflow status becomes count4; scoped
B01–B03 Explain becomes count5. Mandatory separate Sol/high quality/goal review
and appropriate checks follow; stop on findings and push only on PASS. Then
return immediately to app status/actual NEW proof and stop with that caller
local (first application after review). No phase exception, drill, direct ref
edit or bound/count evasion is authorized. B01–B03 Explain remains open until
that scoped assets change; no passing gate/push before resolution. Full design/ACCEPTANCE,
remaining commands/recipes/delivery/final gates and drill restrictions remain.

## Approved correction gate

Direct owner reply: "Approve push and correction gate". Push the recorded batch
and finding to free shared-base P07, then make exactly the scoped project mixed-
kind validation and C07 explanation changes with independent audits/full checks.
Use the same checkpoint reviewer for a scoped re-review. Push corrections and
continue only on PASS; retain the five prior applications until the gate passes.
This approval covers those correction applications beyond the normal five-
change checkpoint. It changes no refs artificially, bounds, phase ownership or
drill authority. After PASS, return immediately to app release wiring/NEW proof.

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
- Owner explicitly approved scoped0107 recovery: preserve the12-line eval patch, restore only its uncommitted edit, natively release the claim, one-time revert of invalid audit9ec2165, then resume the same run/IDs and persist the real finding before correction/re-audit. Release and revert 4188612 completed; this exception is consumed. Both0085 scoped exceptions remain consumed. A separately approved tested0107 recovery now released/reverted onlyb2616c9 as f58bef7; that second0107 exception is also consumed. See [recovery](reports/recovery-0107.md).
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
a47a7fd96ce08a777a341405a1f091d1d5b46c70, Go1.27.2 darwin/arm64, CGO disabled,
vcs.modified=false. This clean build passed without a warning. The prior checkpoint build
nonfatal cache-metadata warning remains recorded in its report.
Build/vet/race/static/history and independent quality review passed at929e32b;
0117/0118 subsequently passed required native app verification and actual NEW CI/list proofs;
actual NEW rejects the recorded invalid mixed-kind fixture before mutation. No final CI/doctor/fresh-checkout/drill/self-hosting
or overall acceptance claim follows.

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

Controller audit submission: the proven native finding quote fits exactly one
physical spec line; describe the full missing assertion after the dash. Full
multiline and flattened sentences were silently discarded in isolated probes.
After every Done inspect phase, audit_rounds and pending IDs before dispatch;
mechanical No findings never substitutes for the separate restricted audit.
This is pinned-controller process guidance, not a product requirement.
