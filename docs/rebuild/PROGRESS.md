# Rebuild progress

Current handoff only. Keep this file short: state, next work, standing rules.
Per-change evidence lives in each archived change's Provenance
(`changes/archive/`); older reports and the former HISTORY.md ledger are in Git
history (last present at `c536833`).

## State — 2026-10-30

126 changes applied (0001–0127; 0109 was abandoned). No pending, blocked or
claimed change. `main` matches `origin/main` at `c536833`.

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

Product work is paused until the owner agrees the next slice. Suggested order:

1. Typed diagnostics: replace the 14 places in `packages/app/internal/entry`
   that recover codes by matching error text. The contracts drill gap
   (CON-011/012) is this problem.
2. `drill` in NEW, then real drills of mid-graph modules (`project`, `checks`)
   with strict input isolation.
3. Remaining commands, recipes, versioning and packaging, then the final gates.

## Standing rules

- Stop on any blocked change. There is no applied-change count or checkpoint
  review (owner, 2026-10-30). Audit rounds and shared-base P07 still apply.
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

- Controller: `.fireweed/bin/fireweed-bootstrap` 0.1.0, source
  `7f925190878e38fc38f30d16bb1cf47aca098cf7`, SHA-256
  `f1516b5686435c23b54b71c5ac554aba6859b5454f072e0c5eb031dfd6f0fd4f`. Do not
  repin or rebuild.
- History baseline: `be35552b706aa3000ff8f11306d129ff1d701b85`. Always pass an
  explicit range; the default range on `main` can be empty.
- `GOCACHE=/tmp/fireweed-rebuild-cache`.
- Workers, auditors and regeneration: `codex exec --ephemeral
  --ignore-user-config -m gpt-6-sol -c 'model_reasoning_effort="high"'
  -s workspace-write {prompt}`.
- Audits run through the controller's audit step with the configured
  `audit.command`.
