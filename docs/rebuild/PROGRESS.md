# Rebuild progress

Current handoff only. Per-change evidence remains in archived Provenance
(`changes/archive/`); no separate report inventories or hash manifests.

## State — 2026-10-31

**EC1–EC6 passed; the early Go core-loop cutover is complete.** EC6 installed
the separate NEW pin, aligned protected instructions/setup and passed all
required pin/status/check/doctor smokes at handoff commit `c1de8f4`.
Full product completion remains open.

The owner selected the next bounded milestone on 2026-10-31: implement drill
support and prove one independent contracts drill in an isolated clone, with
main's NEW pin unchanged. PLAN orders execution, ACCEPTANCE DR1–DR4 defines
completion, and GOAL is the restart prompt. **DR1–DR3 passed; DR4 is at an owner
checkpoint for reported specification gaps.** 0141 applied
at 74dca97 after owner-authorized frozen-eval recovery; 0142 applied the native
app command at d8fb93a; 0143 applied project protection at 8d6f51b. The clean
candidate passed all required checks and actual CLI fixtures below. The single
independent contracts drill passed all own/transitive evals; its reviewed native
report and rebuilt-code ref are retained. DR4 cannot close until the owner
resolves the reported gaps; no automatic fixes or additional drills are authorized.
Main's NEW pin remains unchanged; unrelated worktrees and cutover evidence
are preserved.
Only the earlier push through `dae9b28` was owner-authorized; later work is local.

0131–0138, the two app audit-cwd eval fixes, and demonstrated follow-ups 0139
(Audit insertion) and 0140 (nested doctor templates) are applied. Their phase
checks/audits and 0138's owner-authorized recovery exception remain in archived
Provenance/Notes. Preserve these archives; the completed queue is not restart work.

## Drill slice

0141's runner API/lifecycle applied at 74dca97 with genuine fail-first,
restricted independent audit and final full verification: 128 runner and 282 app
cases passed. Its archived Provenance/Notes hold the recovery and manual audit
submission history. The original implementation attempt remains at
refs/fireweed/attempts/drill-lifecycle/1.

The owner explicitly approved a one-time exception after the pin's normal
eval_fix path proved unavailable for unapplied spec IDs. Author e19a9d6 changes
only the three faulty fake-agent retry branches, requiring their earlier work
before success. A bounded restricted review/correction/re-audit passed RUN-056
and RUN-057, and targeted tests passed against the saved implementation. Test
names, IDs, assertions, coverage, interfaces, original run/audit state, the pin
and bounds were preserved. Native report date/model integration was corrected
within implementation; temporary recovery workspaces are removed.

A focused public DR1 fixture demonstrated that project-level protected paths
inside disposable globs were deleted despite a passing drill. 0143 corrected
that deletion/input/import obligation through normal fail-first and restricted
audit; final implementation verification passed all 423 runner/app cases.
0142 passed all 291 app cases. Its first audit submission was rejected without
import for edits outside Audit; a fresh restricted audit passed after pre-seeding
that section. No verdict, frozen eval or audit counter was manually changed.
No automatic reconstruction-gap fix or unrelated prerequisite is included.

DR2's clean candidate source is
`8d6f51be53e5f87e01ded3d8232352a1249a6ac7`, version 0.1.0, built with
`vcs.modified=false`; executable `.fireweed/bin/fireweed` SHA-256 is
`9325d9ff69b3a27b81e3f24f010bf18057ce43644df1273ec1c14f25cce9ff21`.
Inspectable clone: `/private/tmp/fireweed-dr-candidate-edppuez8/candidate`.
Build, `go test ./...`, `go test -race ./...`, `go vet ./...`, NEW `check`,
`doctor --no-agent`, `check-commits` and `ci` with hook doctor enabled passed.
The exact nonempty history range was
`be35552b706aa3000ff8f11306d129ff1d701b85..8d6f51be53e5f87e01ded3d8232352a1249a6ac7`;
all four NEW check envelopes had no diagnostics. GOCACHE was the prescribed
`/tmp/fireweed-rebuild-cache`. Actual candidate CLI fixtures passed success
plain/JSON, usage/lookup exits, exhausted verification failure, retained native
report/ref and scratch cleanup; inspectable deterministic fixture:
`/tmp/fireweed-dr2-cli-81n68sm9/project`.

The single real contracts drill returned exit 0 on its first reconstruction
attempt: **923 cases passed** (52 contracts, 871 across all seven transitive
dependents). It used the exact tested candidate above and configured
`gpt-6-sol` / high effort. Clone-only fixture alignment is base
`0b1a36a551073c660e4093ab5aa73729c49a7fbb`; only instructions changed from the
tested source, with configuration, setup logic and bounds unchanged.
Native evidence: [.fireweed/drills/contracts.md](../../.fireweed/drills/contracts.md).
Rebuilt-code ref in the inspectable clone:
`refs/fireweed/drills/contracts/20261101T041658.861289000`
(commit `518caf16bb6f21d86b3986bb1ef52899abad5376`).

Live inspection confirmed one fresh input-only commit containing the five
retained contracts files plus isolation policy, no remotes/alternates, and no
old implementation/history, attempts, binaries or other modules' internals.
The scratch controller hash matched the candidate. Review of every worker tool
command found reads confined to those inputs and newly written code. Retained
spec/Decisions/interfaces/evals are byte-unchanged; the ref changes only target
disposable code. The original clone implementation remains untouched. Input
and verifier scratch directories were removed; no live run/shared claim remains.
The native report is retained verbatim, including its extra final blank line
reported by `git diff --check`; this evidence-format warning is not a product fix.

**Required owner checkpoint:** the report identifies no contradicted Decision,
but records three contracts gaps needing agreement: CON-008's round-trip domain
(nil maps and Go concrete types), CON-010's assigned-ID heading grammar, and
CON-002's handling of duplicate section names (also affecting ReplaceSection).
Passing evals do not waive these unresolved choices. No spec/Decision/eval or
main contracts implementation was changed. The report is reviewed and committed;
DR4 and this goal remain incomplete until the owner decides how to handle the
gaps. Do not author fixes, repeat the drill or expand scope without that decision.

## Completed cutover evidence

- **EC1:** all prerequisites and demonstrated follow-ups applied; clean main,
  empty queue and no shared claims before handoff.
- **EC2:** focused core review and affected fixes complete, with no unresolved
  finding: fail-first, immutable metadata/zones, independent audit inputs/history,
  Audit-only import/cleanup, approval/P07, saved attempts, worker-commit recovery
  and bounded time/usage stops. NEW has no runtime bootstrap dependency.
  Actual retry CLI fixture:
  `/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-0136-smoke-agyg5ph1`.
- **EC3:** clean candidate `281ae9498c0fa20b44dc525494427ff123331ab3` passed
  build, `go test ./...`, `go test -race ./...` and `go vet ./...` with the
  specified GOCACHE. Both bootstrap and NEW passed `check`, `doctor --no-agent`,
  `check-commits` and `ci` with doctor enabled over the explicit nonempty range
  `be35552b706aa3000ff8f11306d129ff1d701b85..281ae9498c0fa20b44dc525494427ff123331ab3`.
  A CHK-005 startup timeout and EC5's RUN-036 budget failure each passed a focused
  unchanged rerun and the unchanged full command; no bounds or requirements changed.
- **EC4:** equivalent bootstrap/NEW lifecycle smoke passed: schema-1 envelopes,
  three phase commits, matching held refs, restricted audit inputs and cleanup.
  Intentional differences are DECISIONS D005/D010/D011. Reused the valid
  `66f0c67` lifecycle proof after the doctor-only fix; fixture:
  `/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-ec4-t1_zlt95`.
- **EC5:** the exact final NEW binary governed a real spec change with the
  configured workers and separate auditor, normal approval and bounds, genuine
  fail-first (`e4635b0`), audit (`8202abf`) and application/archive (`da08353`).
  NEW CI passed over `17439f926aba8d0e5ffa6b26f6b8499bf56dc4ff..da08353`.
  Frozen interfaces/evals and phase commits/Provenance were inspected. The
  observed audit workspace contained only change/spec/evals, lacked held history
  and implementation, and was removed; the clone is clean and unclaimed with
  no bootstrap executable installed. Retained evidence:
  `/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-ec5-f78ffohx/candidate`.
- **EC6:** verified the separate regular NEW pin's hash/source and distinct
  subject file, with bootstrap's recovery hash unchanged. Protected AGENTS,
  SYSTEM, setup/recipe and worker PATH are aligned in this author handoff.
  The promoted pin passed `status --json`, `check` and `doctor --no-agent`
  with no diagnostics; PATH resolved `fireweed` to the recorded NEW pin.
  Queue is empty, shared claims absent and scratch worktrees cleaned up.

## Controller pin and execution

- Current NEW controller:
  `/Users/joel/Workspace/fireweed/.fireweed/bin/fireweed`, version **0.1.0**,
  source `281ae9498c0fa20b44dc525494427ff123331ab3`, SHA-256
  `bf0dbc59b7fa7ce2ff957a6f21f34903655b6b054a10695cfeef5ece67fa3575`.
  Built from a clean candidate (`vcs.modified=false`). It is a separate copy
  from mutable `bin/fireweed`; worker builds cannot replace this pin.
- Explicit recovery only:
  `/Users/joel/Workspace/fireweed/.fireweed/bin/fireweed-bootstrap`, version 0.1.1,
  source `78242c8251ff5c2ed7007502dc759c414dca5d75`, unchanged SHA-256
  `ba14dfa072bfa72756d41e1fa0b13fec0df654a7ee629e654cbeb954df8d4683`.
  No automatic fallback, cross-controller claim resume or repin during a live run.
  Recovery repinning requires explicit owner instruction.
- Worker PATH starts with the current tree's `.fireweed/bin`; on main use
  `export PATH=/Users/joel/Workspace/fireweed/.fireweed/bin:$PATH`.
  Hooks receive the actual invoking controller as FIREWEED_BIN. Protected setup
  copies that executable atomically into scratch `.fireweed/bin/fireweed`.
- Use `GOCACHE=/tmp/fireweed-rebuild-cache`. History baseline remains
  `be35552b706aa3000ff8f11306d129ff1d701b85`; always pass an explicit range.
- GOAL records the exact worker/auditor/drill argv; no model/effort substitution.
  Audits use the configured controller audit step and temporary input-only
  workspace. Reconstruction inputs must separately meet DR3 independence.

## Standing rules and remaining scope

- Stop on any blocked change, failed independence/access or drill gap/contradicted
  Decision needing owner agreement. No automatic gap fixes or extra module drills.
  Audit rounds and shared-base P07 still apply.
  Never record the unapproved count or its limit here; read live controller
  `status --json` and `authoring.max_unapproved` before starting work.
- Push only on explicit owner instruction. Protected author edits stay separate
  from controller-owned worker phases, Specs/Decisions and archives.
- Known warnings: bootstrap C01 identifies embedded recipe specs outside the
  module graph; C18 flags repeated generic requirement sentences. B02 reports
  existing APP-103/104 and CHK-055 case counts. Go's module stat-cache write
  warning is nonfatal. NEW's C16 guidance warning was resolved by EC6.
  Zero warnings is not a gate.
- The earlier contracts drill's native result passed, but the agent searched
  other modules' internal code, so it does not meet independence. Its retained
  ref is `refs/fireweed/drills/contracts/20261028T103005.180922000`. It was a
  bootstrap run and cannot substitute for DR3's independent NEW proof.
- Beyond the current contracts proof, full completion still needs remaining
  modules' independent drills, actual regenerate validation,
  gc/migrate and stale/fix conveniences, additional recipes, version/file-format
  migration, remaining warnings, packaging and every later full-completion gate.
  No additional real module drill, main contracts replacement, real regenerate
  change, main controller promotion or deferred product work is included here.

## Next

Stop at DR4's owner gap checkpoint. Await a decision on the three reported
contracts gaps before any follow-up authoring or drill. Preserve the tested
candidate clone/ref, native report, clean main and unchanged controller pin.
No push is authorized; full product completion remains open.
