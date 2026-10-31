# Rebuild progress

Current handoff only. Per-change evidence remains in archived Provenance
(`changes/archive/`); do not create report inventories or hash manifests.

## State — 2026-10-31

0131–0138 are applied on main, with their archived evidence preserved.
0134 (`b42370c`) supplies replaceable-stub guidance; 0135 (`4308d9c`)
retains JUnit failure text; 0136 (`8aef907`) puts bounded own/dependent
assertions into retry feedback; 0137 (`6b26ab7`) distinguishes Decisions
from process Notes. Actual rebuilt NEW invocations confirmed the guidance and
an audit-enabled retry fixture confirmed the JUnit-only dependent assertion
and the 4,096-code-point excerpt bound. Retry fixture evidence remains at
`/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-0136-smoke-agyg5ph1`.

0138 applied at `f92f97a` after full bootstrap verification: 380 cases, zero
failures. Its independent audit checked RUN-013, RUN-027 and RUN-045–049.
The owner authorized interruption recovery and a one-time independently audited
runner helper author correction; saved implementation attempt 2 was restored
without changing IDs, run, audit round or configured bounds. The real configured
workers supplied the implementation; final manual bootstrap done owns its
verification, application and archive. Recovery details remain in its Notes.

Two normal app eval_fix prerequisites removed demonstrated obsolete audit-cwd
assumptions while preserving CLI assertions: `0138-audit-cwd-cli-evals` applied
at `4c794a6` and batched `0138-audit-resume-cli-evals` at `20a8117`. Each passed
all 282 app cases and an independent audit. No product requirements or public
interfaces changed. The final runner verification passed with both corrections.

EC1 is complete, including demonstrated follow-up fix 0139. All changes are
applied, the checkout is clean, queue empty and no shared claim or active
controller/worker remains. EC2's focused review covered fail-first, immutable
metadata/zones, restricted Audit inputs and Audit-only import/cleanup,
approval/P07, saved attempts, worker-commit recovery and clean time/usage stops.
The retry fixture proves the 0135/0136 feedback slice. Product runtime source
has no bootstrap or rebuild-v1 reference.

The first candidate passed ordinary and race Go suites and the reference
lifecycle comparison, then EC5 exposed a valid Audit insertion rejection.
0139 fixed that concrete core failure at `f79a7be`: genuine fail-first evidence,
independent no-findings audit for RUN-047/048, and 381 passing full cases.
The affected review confirms exact non-Audit byte preservation while accepting
a first Audit section before Provenance. No core finding remains unresolved.
Per-change evidence and the stopped, unclaimed first EC5 clone are in 0139's
archive. Intentional comparison differences are in DECISIONS D005/D010/D011.

The corrected `66f0c67` candidate built and passed ordinary/race Go tests and
vet, plus the affected Audit-insertion and reference lifecycle smokes. Both
controllers passed check; explicit full-range history checks remain running.
A real NEW spec lifecycle in the retained EC5 clone passed fail-first, restricted
audit, full implementation verification and archive (`cafd070` / `a2b9e66` /
`e2c0b25`). The observed audit input contained only
change/spec/evals, could not read the held commit, and was removed. Its checkout
is clean and unclaimed; CI remains pending.

EC3 NEW doctor exposed a separate concrete core defect: valid project-owned
nested module template directories are rejected with C14. Bootstrap doctor
passes the unchanged Go setup. Normal fix `0140-doctor-module-templates` is
queued for CHK-053, without changing templates or weakening checks. The C16
managed-block warning reflects the current bootstrap guidance and is addressed
by EC6's planned instruction alignment. Return to the affected gates after 0140.
The controller remains bootstrap. EC3–EC6 are not complete; no promotion or
new push occurred.

## Next

1. Complete 0140 through the normal lifecycle, then use its clean candidate
   for EC3: build, ordinary/race Go tests, vet, and
   both controller/NEW check, doctor --no-agent, explicit-range history and CI.
2. Pass EC4's one deterministic reference comparison and EC5's isolated real-
   worker NEW lifecycle, restricted audit, archive, CI and cleanup. No drill.
3. Only then perform EC6's separate NEW controller pin and protected instruction/
   setup/PATH handoff; record source/hash and concise gate results here.

Full product completion remains open: real drills/regeneration validation,
gc/migrate and stale/fix conveniences, additional recipes, version/file-format
migration, remaining C17–C19 warnings, packaging and later full-completion gates.
No real drill or regeneration work is authorized by this milestone.

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

After EC3–EC5 pass, EC6 replaces the current-controller entry with the verified
NEW path, source commit and executable hash, retains bootstrap's recovery pin,
and records actual gate results. Never prefill promotion evidence.
