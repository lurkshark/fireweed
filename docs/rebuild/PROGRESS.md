# Rebuild progress

Current handoff only. Per-change evidence remains in archived Provenance
(`changes/archive/`); no separate report inventories or hash manifests.

## State — 2026-11-01

**SP1–SP4 passed.** EC1–EC6, DR1–DR4 and FD1–FD4 remain complete and preserved.
The owner-selected stale-change authoring milestone is finished: both CLI options
are implemented, normal lifecycle/candidate gates passed, and project now has a
reviewed passing independent proof. Full product completion remains open; every
later full-completion gate is still required. No further execution slice is selected.

**SP1/SP2 — product and lifecycle.** 0146 applied at `21d5dde` with project
StaleSpecs/StampMissing APIs and PRJ-060–066. Genuine fail-first (`9e08d86`:
126 cases, 17 failing), restricted audit (`a313e6d`) and full implementation
verification (798 passing cases) are preserved in
[0146's archive](../../changes/archive/0146-stale-change-authoring.md).
0147 applied at `78c334f`, exposing show --stale and missing-bases-only check --fix.
Its 13 new cases failed first while all 291 existing cases passed; restricted
audit and full verification passed all 304 final app cases. Evidence remains in
[0147's archive](../../changes/archive/0147-stale-authoring-cli.md).
Actual NEW CLI fixtures cover stamped history recovery, unavailable/mismatched
fallback, mixed missing/valid/stale bases, file/index/ref effects, JSON and
applicable diagnostic/usage/lookup exits. Inspectable fixtures:
`/private/tmp/fireweed-sp-cli-y4uu9x98`. Runtime behavior is unchanged by recovery,
so those fixtures remain valid.

Owner-approved recovery changes are also applied and archived. 0148 (`a11ace0`)
corrected RUN-040's timing fixture under eval_fix, preserving its deadline and
all assertions; 132 own and 436 full cases passed. The owner explicitly approved
repair/restart of its initially mistaken author metadata. 0149 (`c4d34a2`)
clarified existing diagnostic contracts with already_satisfied, passing 127 own
and 812 full cases. 0150 (`b1a6015`) repaired app's corrective-word helper and
malformed-history fixture under eval_fix, passing all 304 cases. Their restricted
audits passed with no findings. Archives hold the per-change Provenance; no
runtime source, public interface or configured bound changed in 0148–0150.
P03 rejected 0150's first premature worker phase edit before a worker commit;
the normal remaining attempt restored approved state. Final phase provenance is
controller-owned, and explicit worker phase control is retained in SYSTEM.
A bounded doctor probe after 0149's first invocation failure passed both exact
configured worker/auditor fixed replies; the same claim then resumed normally.

**Verified candidate.** Inspectable clone:
`/private/tmp/fireweed-sp-project-8ko73j3j/candidate`. Executable source/aligned
base `3f95497a164b045bf374da4a1498011bb0bc14c2`, version 0.1.0,
`vcs.modified=false`, SHA-256
`4e6f5adf84ca0bc3dcb28001d9623b51bb68bebc5a4d53d93dbbb6c479fd927f`.
The candidate retains its input-only SYSTEM/AGENTS policy. Runtime source,
interfaces, configuration, setup and hooks are byte-identical to prior verified
`fed193a`; recovery changes project specs/evals and app evals only. Valid full
Go normal/race/vet evidence is reused for unaffected modules. Affected project
normal/race passed in 3.681/5.394 seconds, app in 268.242/300.839 seconds, and
affected vet passed. Earlier affected runner normal/race passed in
269.433/283.520 seconds. NEW check and all ten doctor --no-agent canaries passed.

The standing history baseline remains `be35552b706aa3000ff8f11306d129ff1d701b85`.
NEW check-commits passed without diagnostics over
`be35552b706aa3000ff8f11306d129ff1d701b85..5bacb4ab5817e8eaea38218e6a8bcf9354dce42c`,
`5bacb4ab5817e8eaea38218e6a8bcf9354dce42c..f22a327811f0c188cbbbe239a544e27fda80ce5b`,
`f22a327811f0c188cbbbe239a544e27fda80ce5b..fed193a3f8d2c86dab33aef69972543e41d45e9f`,
and the 14-commit recovery append
`fed193a3f8d2c86dab33aef69972543e41d45e9f..3f95497a164b045bf374da4a1498011bb0bc14c2`.
Native CI with hook doctor enabled passed without diagnostics over
`be35552b706aa3000ff8f11306d129ff1d701b85..fed193a3f8d2c86dab33aef69972543e41d45e9f`
in 1065.946 seconds and the same explicit recovery append in 505.103 seconds.
The valid prefix is reused; the append CI checked all current module evals,
coverage, boundaries and doctor canaries, without skipping any required gate.

**SP3 — independent project proof passed.** The owner-authorized fresh drill
used that clean candidate, exact configured gpt-6-sol/high worker and unchanged
bounds. Live inspection confirmed 19 byte-identical retained inputs, fresh
one-commit Git metadata, no remotes/alternates, no old report/history, saved
attempts, binaries, manifests or other-module internals. The separate verifier's
old target disposable files were absent and its controller matched the invoker.
Review covered all 51 shell commands and 22 patch operations, including complete
script bodies and incremental changes behind repeated cumulative patch output.
Reads stayed within supplied inputs/new files. Attempts to read absent app evals
returned no files; no dependent implementation or outside oracle was accessed.

The native result is pass after its configured retry, in 2216.237 seconds:
**all 812 enforced cases passed** — project 127, app 304, checks 101, runner 132
and workflow 148 — with passing coverage and all retained inputs byte-unchanged.
No contradicted Decision or unresolved behavioral gap remains. The agent's
no-manifest note is the intentional input-only testing limitation, resolved by
the separate verifier's passing own and all-transitive-dependent evals.
Native report: [.fireweed/drills/project.md](../../.fireweed/drills/project.md),
preserved verbatim. Passing ref, retained in main and the inspectable clone:
`refs/fireweed/drills/project/20261102T031647.104747000` at
`481bc00c1da360d3e8bf76cd5ba04053d2dccaeb`, based on the identified candidate.
The drill left invoking HEAD/executable unchanged. Candidate report-only author
commit `3e40231` then retained the reviewed report with a clean checkout.
Rebuilt implementation remains in the proof ref; main's implementation was not replaced.

The earlier failed project proof remains inspectable at
`refs/fireweed/drills/project/20261102T000116.032120000` (`dc7a84f1ae92d7f0bab3b91cc293882427591fad`),
with its verbatim report in main commit `86c512c` and candidate commit `41bdb56`.
Its 16 failed dependent entries led to the owner-approved 0149/0150 corrections;
no retry bound was reset. Completed contracts/gitstore/assets reports, refs,
archives and inspectable clones remain preserved.

**SP4 — handoff.** Native input and verifier scratch worktrees were removed;
main and candidate are clean and unclaimed, with no active worker/run. Main's
separate NEW pin and fireweed.yaml bounds are unchanged. The unrelated worktree
at `/Users/joel/.codex/worktrees/stub-handoff-guidance/fireweed` is preserved.
The owner permanently raised SYSTEM's full implementation Verify process
cost guardrail from 15 to 25 minutes on 2026-11-01; own Verify stays under five
minutes and configured controller bounds were never changed. Only the earlier
push through `dae9b28` was authorized; subsequent work remains local. No controller
promotion, main regeneration replacement, other-module real drill or release occurred.

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
verbatim, including its final blank line warning. Evidence was committed before
assets started.

**Assets: FD1/FD3 passed.** Inspectable clone:
`/private/tmp/fireweed-fd-assets-s53ujqr9/candidate`, aligned base
`87261cd10e45a243d450ec419a5a927bd98156e0`. It uses the same verified clean
`1602083` executable/source/hash above; runtime source and assets retained
contracts/evals are unchanged. Unaffected DR2 and subsequent checks were reused;
candidate `check` and `check-commits fc12ffc..87261cd` passed without diagnostics.
Public zone classification confirmed every old bundle file disposable, including
the nested recipe SPEC. Preflight and live inspection found exactly 17 retained
files (target spec/Decisions/api/model/evals, contracts interfaces and isolation
SYSTEM), one fresh input commit and no remotes/alternates. No old bundle data,
history, binaries, attempts or other-module internals entered the worker inputs.
The separate verifier deleted all old assets disposable files and copied the
actual invoker with matching hash. Agents used the exact configured argv/model/
effort, prescribed cache and current-tree PATH; Go verification used GOPROXY=off.

Native report: [.fireweed/drills/assets.md](../../.fireweed/drills/assets.md).
The first reconstruction invocation passed all **863 cases**: 82 own and 781
across app, checks, project, runner and workflow. No contradicted Decisions were
identified. The report's Gaps entry describes inability to run the full retained
suite in the input-only workspace, where manifests and dependency implementation
were deliberately absent. This is the required isolation arrangement, already
explained in supplied SYSTEM; the separate native verifier ran and passed that
full suite. The note is resolved by this verification and requires no behavioral
choice or owner-agreed fix; it stays verbatim in the native report.

Review of all 26 executed shell commands (including complete Python bodies)
found reads confined to supplied inputs, newly written code and owned in-input
fixtures. Recipes/playbooks/help/explanations were reconstructed without querying
the invoker. A rejected combined cleanup command was safely replaced with file
unlink/directory removal; temporary fixtures were removed. Retained inputs are
byte-unchanged. The rebuilt ref changes only assets disposable paths:
`refs/fireweed/drills/assets/20261101T173119.320562000`
(commit `53ed45d3ee36dcac05fc5b736fa461152dd011d4`). The invoking implementation
and HEAD stayed unchanged during the drill. Native input/verifier cleanup passed;
no live run/shared claim remains. Trailing whitespace and final blank-line
warnings inside captured native observations are retained verbatim as evidence
formatting, not a product change.

**FD4 passed.** Both native reports are committed as protected author evidence;
refs and inspectable clones are retained. Main's original source/hash pin and
all configuration/bounds are unchanged. Main implementation was not replaced or
imported; no real regenerate change, controller promotion, other-module drill,
deferred product work or push occurred. Completed contracts evidence and the
unrelated worktree remain intact. Full-completion gates remain required.

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
- FD1–FD4 and SP1–SP4 are completed evidence: gitstore/assets proofs, stale
  authoring in project/app, then the passing project proof. Project's dependency
  inputs are contracts/gitstore/assets interfaces only; dependency internals and
  invoking candidate content are excluded. Only ordinary feature implementation
  belongs on main; independently rebuilt project code stays in the proof ref.
  Real regenerate changes, controller promotion, other-module drills, unrelated
  deferred work and push remain outside SP1–SP4. Later gates remain required.

## Next

SP1–SP4 is complete. No further slice is selected. Preserve the pinned controller,
configured bounds, completed proofs and inspectable clones. Await owner selection
before further implementation, real drills, controller promotion or deferred work.
Full product completion and every later full-completion gate remain open.
