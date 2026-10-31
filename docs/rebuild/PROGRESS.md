# Rebuild progress

Current handoff only. Keep this file short: state, next work, standing rules.
Per-change evidence lives in each archived change's Provenance
(`changes/archive/`); older reports and the former HISTORY.md ledger are in Git
history (last present at `c536833`).

## State — 2026-10-30

Owner selected early Go core-loop cutover, with 0134–0137 required rather than
optional. This author branch, codex/stub-handoff-guidance, contains those queued
changes plus 0138's restricted audit-workspace fix. No EC gate is claimed passed
and no controller promotion has occurred.

At the primary-checkout inspection, main had applied through 0131 (0109 was
abandoned), resumed 0132 after an interrupted full check, and still had 0133
queued. Re-read its live run/queue/claims before integration. This branch forks
from 0131's audit commit; its copied pending 0131 phase is stale and must not be
executed as the primary checkout's current state. Integrate only the author
commits onto current main after its run returns clean and unclaimed.

The current controller remains bootstrap 0.1.1 at the pin below. Prior approval
to push the 2026-10-30 bootstrap repin is historical, not authorization to push
this branch or the future NEW promotion.

**Working through NEW (`bin/fireweed`), each proven by an actual CLI run:**

- Setup: `init` (Go recipe), `doctor`, `new-module`, `new-change`, `stamp`.
- Lifecycle: `start`, `verify [--quick]`, `done`, `context`; spec, fix,
  eval_fix, maintenance and regenerate kinds; audit rounds; approval and P07.
- Unattended: `next`, `run [--max] [--for]`, retries, saved attempts and
  `--from-attempt`, projected blocking with independent continuation,
  agent-commit undo, usage/time stops with exit 3.
- Management: `release`, `commits`, `abandon`.
- Reading and checks: `list`, `status`, `modules`, `show [ID] [--projected]
  [--history]`, `explain`, `playbook`, `check`, `coverage`, `check-commits`,
  `ci [--no-doctor]`, `pr-body`, `agents-md`, `version`.

**Deferred beyond early cutover:**

- Real drills/regeneration validation; gc/migrate and stale/fix conveniences.
- Additional recipes, version-range/file-format migration, C17–C19 warnings.
- Packaging and the remaining full-product acceptance/delivery obligations.

## Next

1. Let the primary run finish its existing queue; stop at any new block.
2. At an idle clean integration point, apply this branch's author commits and
   complete 0134–0138 with bootstrap and normal phase checks.
3. Prove ACCEPTANCE EC1–EC5: focused core review, fresh-checkout checks,
   reference smoke and isolated real-worker self-hosting without a drill.
4. Perform EC6's separate controller pin/instruction handoff and record its
   source/hash and actual gate evidence here. Until then use bootstrap.

Gate status: EC1–EC6 pending. Further prerequisite changes need a demonstrated
core failure or unmet core clause; no speculative cleanup or deferred features.

## Standing rules

- Stop on any blocked change. There is no applied-change count or checkpoint
  review (owner, 2026-10-30). Audit rounds and shared-base P07 still apply.
- Never record the unapproved count or its limit here. Before starting work,
  read the live count from controller `status --json` (`unapproved`) and the
  limit from `authoring.max_unapproved` in fireweed.yaml.
- Push only on explicit owner instruction.
- Drills: only the contracts drill has run. Its native result passed (373
  cases), but the agent searched other modules' `internal/` code, so it does
  not meet the independence gate. Rebuilt code is on
  `refs/fireweed/drills/contracts/20261028T103005.180922000`. No further drill
  or gap fix without owner agreement.
- Known warnings with explanations: C01 (embedded recipe discovery), C18,
  B02 (CHK-055 has 16 cases), and Go's nonfatal module stat-cache write warning.
  Zero warnings is not a gate.

## Fixed execution context

- Current controller (pre-EC6): `.fireweed/bin/fireweed-bootstrap` 0.1.1, source
  `78242c8251ff5c2ed7007502dc759c414dca5d75`, SHA-256
  `ba14dfa072bfa72756d41e1fa0b13fec0df654a7ee629e654cbeb954df8d4683`. Repin
  or rebuild only on explicit owner instruction. EC6 will record a separate NEW
  controller pin; mutable bin/fireweed is never that pin.
- History baseline: `be35552b706aa3000ff8f11306d129ff1d701b85`. Always pass an
  explicit range; the default range on `main` can be empty.
- `GOCACHE=/tmp/fireweed-rebuild-cache`.
- Workers, auditors and regeneration: `codex exec --ephemeral
  --ignore-user-config -m gpt-6-sol -c 'model_reasoning_effort="high"'
  -s workspace-write {prompt}`.
- Audits run through the controller's audit step with the configured
  `audit.command`. Automatic audits use a temporary workspace containing only
  the change, projected spec and evals, plus empty Git metadata. Manual isolated
  audit workspaces are no longer required.

After EC6, replace the current-controller entry with the verified NEW path,
source commit and binary hash, retain bootstrap's recovery pin, and record
which evidence closed each EC gate. Do not prefill promotion evidence.
