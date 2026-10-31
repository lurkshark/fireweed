# Rebuild progress

Current handoff only. Keep this file short: state, next work, standing rules.
Per-change evidence lives in each archived change's Provenance
(`changes/archive/`); older reports and the former HISTORY.md ledger are in Git
history (last present at `c536833`).

## State — 2026-10-30

127 changes applied (0001–0128; 0109 was abandoned). The owner explicitly
approved running all queued changes on 2026-10-30. The author repaired 0128's
ErrorCode facade to delegate to disposable code, preserving its frozen evals;
resumed implementation and full controller verification passed.
0129 passed evals and independent audit, then blocked after both implementation
attempts failed Test_APP_065_APP_066_APP_067_DoctorFailures. A focused run
confirmed C05 and the duplicate-ID message are correct, but diagnostic Path is
fireweed.yaml instead of SPEC.md: Load now detects malformed items earlier,
while project.Check labels every load failure fireweed.yaml. No claim remains;
0130–0133 wait on project. The owner approved the scoped recovery below.
Queued predicates 0130/0131 state the facade/stub wiring requirement before
their evals phases.
The owner authorized the controller upgrade to 0.1.1 on 2026-10-30; version,
executable hash and clean source revision
were verified after installation. Controller `check` passed with the existing
C01/C18 warnings.
The owner authorized pushing this controller repin to `origin/main` on 2026-10-30.

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

**Not built yet:**

- Commands: `drill`, `gc`, `migrate`, `show --stale`, `check --fix`.
- Recipes: `typescript-pnpm`, `python-uv`, `none`, local-path recipes.
- `fireweed:` version-range enforcement and file-format migration.
- Packaging (goreleaser, installer, Homebrew config).
- Every ACCEPTANCE completion gate: fresh-checkout Go checks, controller
  CI over the full baseline range, reference smoke, real drills of every
  disposable module, isolated self-hosting, delivery.

## Next

Owner approved resuming 0129 with the project-local diagnostic-path correction
on 2026-10-30, preserving frozen evals and existing run/phase, then continuing
0130–0133. The new recovery guidance is in 0129 Notes. Stop on any new block.
The queue completes typed diagnostics and addresses the contracts drill gap
(CON-011/012). Suggested later order:

1. `drill` in NEW, then real drills of mid-graph modules (`project`, `checks`)
   with strict input isolation.
2. Remaining commands, recipes, versioning and packaging, then the final gates.

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

- Controller: `.fireweed/bin/fireweed-bootstrap` 0.1.1, source
  `78242c8251ff5c2ed7007502dc759c414dca5d75`, SHA-256
  `ba14dfa072bfa72756d41e1fa0b13fec0df654a7ee629e654cbeb954df8d4683`. Repin
  or rebuild only on explicit owner instruction.
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
