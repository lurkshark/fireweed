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

Prerequisite EC1 was proven at `ff54e38`: all prerequisites applied, execution checkout clean, queue empty,
no shared claims or active controller/worker processes. EC2's initial focused core
review covered: current public evals and
implementation cover fail-first, immutable metadata/zones, restricted audit
inputs and Audit-only import/cleanup, approval/P07, saved attempts, worker-commit
recovery and clean time/usage stops. The retry fixture above proves the
0135/0136 feedback slice. Product runtime source has no bootstrap or rebuild-v1
reference. Intentional controller differences remain in DECISIONS D005/D010.

The clean `ff54e38` candidate built and passed ordinary `go test ./...`.
EC4's two fixture lifecycles completed with exit 0, three phase commits,
matching held refs, restricted inputs and audit cleanup. Bootstrap returns
string next JSON data; NEW's object follows APP-079 (record at handoff).

EC5's real NEW run reached CUT-001 fail-first evidence, then rejected
the audit submission. A focused exact-candidate CLI reproduction confirms a
valid first Audit section inserted before Provenance is rejected despite every
original byte remaining intact. Normal runner fix `0139-audit-section-import`
is queued for RUN-047/048. The stopped fixture's NEW-owned claim was released;
retain `/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-ec5-y5nxlf50/candidate`
for inspection. No change is blocked. Return to affected gates after the fix.
The controller remains bootstrap at the pin below. EC3–EC6 are not complete;
no promotion occurred. The owner authorized only the earlier push through the
core-cutover authoring integration at `dae9b28`; subsequent work remains local.

## Next

1. Complete 0139 through the normal fail-first/audit/implementation loop.
   Then use the corrected clean candidate for EC3: build, ordinary/race Go tests, vet, and
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
