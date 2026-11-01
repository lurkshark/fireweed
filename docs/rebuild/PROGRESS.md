# Rebuild progress

Current handoff only. Per-change evidence remains in archived Provenance
(`changes/archive/`); no separate report inventories or hash manifests.

## State — 2026-11-01

**EC1–EC6 and DR1–DR4 passed.** Early Go core-loop cutover, native drill support
and the independent contracts proof are complete. Full product completion is open.
The owner selected **independent gitstore, then assets drills, one at a time**;
PLAN orders execution, ACCEPTANCE FD1–FD4 defines completion, and GOAL is the
restart prompt. Both targets depend only on contracts.

**FD1/FD2 passed for gitstore; assets is next.** The owner-approved recovery
proof passed on its first reconstruction invocation. Native gitstore evidence
is reviewed and retained below. Assets has not started. Main and the proof
clone are unclaimed, with no live drill/worker and scratch removed. Main's pin
and configured bounds remain unchanged. Preserve the completed contracts proof
and unrelated worktrees. Only the earlier push through `dae9b28` was authorized;
later work remains local. Full product completion stays open.

## Foundation proofs — FD1–FD4

**Gitstore: FD1/FD2 passed.** Inspectable clone:
`/private/tmp/fireweed-fd-gitstore-recovery-f73u9gx9/candidate`, aligned base
`89d7dbfb23b788d3efc98ba8c7709291420d8438`. It reused the clean NEW executable
built from `1602083a955df97910348be76f99cc3816729d93`, version 0.1.0,
`vcs.modified=false`, SHA-256
`5c52a4d949d9e39bfb8560bd349da5cde884533dea8e976647656e0ded52b74f`.
Runtime source is unchanged since that build. DR2's unaffected checks and the
contracts proof were reused; affected gitstore race tests and vet, candidate
`check`, and `check-commits 0eb1353..89d7dbf` passed without diagnostics.

The approved recovery invocation used exact configured Codex argv/model/effort,
cache, PATH and unchanged bounds. Preflight and live inspection found exactly
eight retained input files: isolation SYSTEM, target SPEC/Decisions/api/model/
evals (including empty `.gitkeep`), and contracts interfaces only. Input Git
had one fresh commit, no remotes/alternates or old history. Old disposable code,
binaries, saved attempts and dependent internals were excluded; the separate
verifier's controller copy matched the invoker. Review of all 21 shell commands
and six patch operations found reads confined to supplied inputs, newly written
code and temporary eval fixtures. Its temporary in-input test layout was removed
with unlink/rmdir after a rejected rm cleanup command.

Native report: [.fireweed/drills/gitstore.md](../../.fireweed/drills/gitstore.md).
All **790 cases passed**: 9 own, 781 across app, checks, project, runner and
workflow. No gaps or contradicted Decisions were identified. Retained inputs
are byte-unchanged; the rebuilt ref changes only target disposable paths:
`refs/fireweed/drills/gitstore/20261101T170411.628421000`
(commit `3cdc945cea3923dd22877e56e9a570f7a3caf346`). Invoking implementation and
HEAD stayed unchanged during the drill. Both input/verifier scratch directories
were removed, and no live run/shared claim remains. The native report is retained
verbatim, including its final blank line warning. Evidence is committed before
starting assets. FD3 and final FD4 remain pending.

The earlier interrupted gitstore invocation remains incomplete, with no native
report/ref, at `/private/tmp/fireweed-fd-gitstore-ierzwr1i/candidate`. Its second
attempt was interrupted with SIGINT; the CLI exited 130 without native cleanup.
The worker exited and only that run's scratch was manually removed. No rebuilt
code was reused. The owner's 2026-11-01 `Approve` accepted its reported choices
and recovery through the fresh invocation above. Normal change
[0145](../../changes/archive/0145-gitstore-drill-clarifications.md) applied at
`0eb1353` with legitimate already_satisfied entries: public probes already
satisfied both choices, including unstaged tracked-diff rename retention and
selected new files. It preserved all prior assertions, implementation and
interfaces. Its 9 eval cases and restricted independent audit of both IDs passed,
then full controller verification passed 790 cases. The first eval worker's
premature phase advance was rejected with P03; the remaining configured retry
restored runtime fields and removed premature Provenance before normal phase
commits. Per-change evidence remains in the archive.

## Completed drill evidence — DR1–DR4

0141 implemented runner drill support at `74dca97`; 0142 added the native CLI at
`d8fb93a`; 0143 fixed approved project protection at `8d6f51b`. Their fail-first,
audit, verification and recovery details stay in archived Provenance/Notes.
0141's owner-authorized frozen-eval exception was specific to that recovery,
preserved coverage/pin/bounds and grants no standing exception for later work.
Do not restart completed changes or expose saved attempts as reconstruction input.

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

The owner approved all three reported interpretations. Normal change
[0144](../../changes/archive/0144-contracts-drill-clarifications.md) applied at
`b5fc1f54e2be9638c8ed686e817c9468dbb0554a`: CON-008 rejects nil frontmatter and
bounds type identity to YAML representations, CON-010 defines assigned-ID
syntax, and CON-002/009 select the first duplicate section. All existing
assertions remain. Its evals had three genuine failures; restricted independent
audit passed all four IDs, followed by full controller verification with all
935 cases passing. Evals/interfaces stayed frozen through implementation.
Pin `check` and `check-commits 4bffbd0..b5fc1f5` passed without diagnostics.

DR4 reassessment exported the retained rebuilt-code ref into a separate temporary
verifier and supplied the newly frozen contracts spec/evals, without changing
the retained ref, proof clone or native report. All 64 contracts cases passed
under `go test -race`, and vet passed; main passed the same affected Go checks.
The reported interpretations already hold in that independent rebuild, so the
original real proof remains valid and was reused. Other modules' code/evals and
their native passing evidence are unchanged. The temporary verifier was removed;
no second real drill or main implementation import occurred. The native report's
historical Gaps section stays verbatim; this owner-approved change resolves it.
DR1–DR4 are complete. The native proof remains valid after 0144; full product
completion remains open.

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
  workspace. Reconstruction inputs must separately meet FD1 independence.

## Standing rules and remaining scope

- SYSTEM governs independence, immutable phase zones, approval/P07, retries,
  audit rounds and owner checkpoints. Stop on any blocked change, failed
  independence/access or gap/contradicted Decision needing owner agreement.
  Corrections require agreement before authoring; preserve cases, assertions
  and ID coverage through normal fail-first and restricted independent audit.
  Never record live unapproved counts or configured limits here.
- Native reports/refs and inspectable clones are the drill evidence. No target
  implementation, embedded bundle content, source history, saved attempt,
  product binary or bootstrap may become reconstruction input. For assets,
  this also forbids querying the candidate's built-in content as an oracle.
- Known warnings: bootstrap C01 identifies embedded recipe specs outside the
  graph; C18 flags repeated generic requirement sentences. B02 reports existing
  APP-103/104 and CHK-055 case counts. Go's module stat-cache write warning is
  nonfatal. The native contracts report retains its extra final blank line.
  Zero warnings is not a gate; no speculative warning cleanup is selected.
- The earlier bootstrap contracts ref
  `refs/fireweed/drills/contracts/20261028T103005.180922000` is not independent
  because its worker searched other modules' internals. Use the reviewed NEW
  proof above. Preserve completed archives and evidence; repeat only checks or
  proofs invalidated by relevant changes, within the selected scope.
- FD1–FD4 cover gitstore and assets only. Main implementation import/replacement,
  real regenerate changes, controller promotion, other modules' real drills,
  deferred product work and push remain outside this milestone. All later
  full-completion gates remain in ACCEPTANCE.

## Next

Prepare and independently inspect assets inputs, excluding every disposable
embedded bundle file and all invoker content as reconstruction sources. Run only
assets next, with exact configured agents and unchanged main pin/bounds. Stop on
new gaps, contradicted Decisions, access/independence failure or a blocked change.
Commit reviewed native evidence and complete FD4 before marking this goal done.
