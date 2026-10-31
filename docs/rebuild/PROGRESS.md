# Rebuild progress

Current handoff only. Keep this file short: state, next work, standing rules.
Per-change evidence lives in each archived change's Provenance
(`changes/archive/`); older reports and the former HISTORY.md ledger are in Git
history (last present at `c536833`).

## State — 2026-10-31

136 changes applied (0001–0137; 0109 was abandoned). All six typed-diagnostic
changes (0128–0133) have passing implementation Provenance; spec changes passed
independent audits and maintenance preserved frozen evals/interfaces. 0132
resumed from saved attempt 1 after the owner's interruption and continuation.
NEW was rebuilt after 0133; actual CLI smokes confirmed an empty queue, doctor
C05 with a module SPEC.md path, Stamp C08, missing-change U02 and preserved C12
behavior for doctor/agents-md. The 0133 in-step review resolved a direct project-
load classification regression before application. Evidence remains archived.

The owner selected early Go core-loop cutover, with 0134–0137 required. The
author commits from codex/stub-handoff-guidance are now integrated on main,
including 0138's restricted audit-workspace fix and the revised milestone plan.
Integration followed completion of the primary run with a clean unclaimed
checkout. 0134 applied at `b42370c` after one corrected audit finding, with
passing implementation Provenance (789 cases). NEW was rebuilt and its actual
`playbook evals` output confirmed the replaceable-stub guidance.

0135 applied at `4308d9c` with passing implementation Provenance (610 cases)
and an independent audit with no findings. On 2026-10-31 the owner authorized
resuming the lead's interruption block from saved attempt 1 and correcting the
older whole-Case equality test before freeze. The configured real worker made
that one-line correction; bootstrap's manual `done` verified/committed evals.
Bootstrap `next` required releasing this checkout's own claim before reclaiming
the committed audit phase; no foreign claim, IDs, run, bounds or requirements
changed.

0136 applied at `8aef907` after one corrected audit finding, with passing
implementation Provenance (364 cases). Rebuilt `bin/fireweed next` completed
an audit-enabled deterministic retry fixture: its saved prompt contains the
dependent module's JUnit-only assertion, a 4,096-code-point failure excerpt and
a 4,225-code-point section including the heading. Inspectable fixture:
`/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-0136-smoke-agyg5ph1`.
0137 applied at `6b26ab7`, with passing implementation Provenance (803 cases)
and an independent audit with no findings. Rebuilt NEW's actual evals, impl
and regenerate playbook outputs confirmed the Decisions/Notes guidance and
examples while preserving the stub handoff instructions.

0138 passed audit at `62485a1` with all seven IDs checked, after the owner-
authorized audit restart and one findings round. Its corrected evals committed
at `359eb86` with fail-first evidence (380 cases, 16 failing). Implementation
is now blocked at `98cc19a`: three legacy queue tests share a frozen helper
that still requires audit cwd to equal the held checkout, contradicting the
new audit isolation. Both implementation attempts failed on that assertion;
the focused new workspace tests pass. The implementation is saved at
`refs/fireweed/attempts/0138-isolated-audit-workspaces/2`; attempt 1 retains the
earlier partial audit. No worker eval edit was retained.

The owner authorized the proposed one-time recovery exception on 2026-10-31.
A restricted independent review of the helper patch reported No findings for
RUN-016–019; its only edit was the review change's Audit section. This author
commit corrects the helper and clears blocked while unclaimed, preserving
module/phase/Bin checks and held worker cwd while requiring a separate,
cleaned-up audit cwd. Resume saved implementation attempt 2 under bootstrap.
IDs, run, audit rounds, configured bounds and product requirements remain
unchanged. EC1–EC6 remain pending and no promotion occurred.

The current controller remains bootstrap 0.1.1 at the verified pin below.
The owner authorized pushing current main to origin/main on 2026-10-30,
including the completed diagnostic queue and the core-cutover authoring
integration at dae9b28.

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

1. Resume 0138 implementation attempt 2 with bootstrap as authorized and run
   normal verification against the independently reviewed helper correction.
   Stop at any new block. 0132–0137 are complete, including the JUnit retry-
   feedback CLI slice and worker guidance. Acceptance gates remain required.
2. Prove ACCEPTANCE EC1–EC5: focused core review, fresh-checkout checks,
   reference smoke and isolated real-worker self-hosting without a drill.
3. Perform EC6's separate controller pin/instruction handoff and record its
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
  B02 (APP-103/104 and CHK-055 case counts), and Go's nonfatal module stat-cache write warning.
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
