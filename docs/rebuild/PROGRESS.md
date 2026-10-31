# Rebuild progress

Current handoff only. Per-change evidence remains in archived Provenance
(`changes/archive/`); no separate report inventories or hash manifests.

## State — 2026-10-31

0131–0138 are applied, with earlier archived evidence intact. 0134 supplies
replaceable-stub guidance; 0135 retains JUnit failures; 0136 supplies bounded
own/dependent retry assertions; 0137 separates Decisions from process Notes;
0138 isolates audit inputs, validates Audit-only edits and cleans up workspaces.
Actual NEW invocations confirmed the guidance and bounded retry feedback.
0138 applied at `f92f97a` with all 380 full cases passing after the owner's
recorded recovery exception. Details remain in its archived Notes.

The two normal app eval fixes (`4c794a6`, `20a8117`) removed demonstrated
obsolete audit-cwd assertions, each with all 282 cases passing and separate
no-findings audits. No public interface or product requirement changed.
EC5 then demonstrated a valid Audit insertion rejection; normal fix 0139
applied at `f79a7be` with genuine fail-first, an independent no-findings audit
for RUN-047/048 and all 381 full cases passing.

EC3 NEW doctor exposed valid nested-template rejection. Normal fix 0140 applied
at `9822408`: one corrected audit round for baseline rendering, no findings on
re-audit for CHK-053, and all 630 full cases passing. Protected Go templates,
setup and canary configuration are unchanged. The affected review confirms
recursive template rendering and preserved malformed-input/canary checks.

EC1 is complete: prerequisites and demonstrated follow-ups applied, clean main,
empty queue, no shared claims or active controller/worker processes. EC2's
focused review and affected fixes are complete, with no unresolved core finding:
fail-first, immutable metadata/zones, independent Audit input/history isolation,
Audit-only import/cleanup, approval/P07, saved attempts, worker-commit recovery
and bounded time/usage stops remain enabled. The 0135/0136 retry slice was
proved by the retained NEW CLI fixture at
`/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-0136-smoke-agyg5ph1`.
NEW has no product runtime bootstrap dependency.

The preceding clean `66f0c67` candidate passed build, ordinary/race Go tests,
vet and both controllers' check and explicit 934-commit history checks.
Bootstrap doctor passed; NEW's doctor failure is now fixed by 0140. EC3 must
finish on the corrected candidate, including both doctor/CI runs.
EC4's comparison passed with three lifecycle commits, matching held refs,
restricted audit inputs and cleanup. Its ordinary lifecycle code is unchanged
by 0140; reuse that evidence. Intentional differences are in DECISIONS
D005/D010/D011. Reference fixture:
`/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-ec4-t1_zlt95`.

The real NEW lifecycle in the retained EC5 clone passed fail-first (`cafd070`),
restricted audit (`a2b9e66`) and application/archive (`e2c0b25`). The observed
workspace contained only change/spec/evals, could not read the held commit and
was removed. Clone is clean and unclaimed:
`/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-ec5-qgu1qc4l/candidate`.
Its CI proof was pending the doctor fix. Complete the gate with the corrected
candidate before pinning any main controller.

The controller remains bootstrap at the pin below. EC3, EC5 and EC6 are not
complete; no promotion occurred. Only the earlier push through `dae9b28` was
owner-authorized; later work remains local.

## Next

1. Build the corrected clean candidate and run actual NEW doctor first.
2. Complete exact-candidate EC3 Go checks and both controller/NEW checks,
   doctor --no-agent, explicit-baseline history and CI with doctor enabled.
3. Complete EC5 on the corrected candidate with an explicit separate NEW pin,
   real configured workers/auditor, archive, CI and cleanup. No drill.
4. Only after EC1–EC5 pass and all claims/runs return, perform EC6's protected
   controller pin/instruction/setup/PATH handoff and record actual results.

Full product completion remains open: real drills/regeneration validation,
gc/migrate and stale/fix conveniences, additional recipes, version/file-format
migration, remaining warnings, packaging and later full-completion gates.
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

After EC3–EC5 pass, EC6 records the verified NEW path, source commit and
executable hash, retains bootstrap's recovery pin, and records actual handoff
smokes. Never prefill promotion evidence.
