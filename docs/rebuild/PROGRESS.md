# Rebuild progress

Current handoff only. Per-change evidence remains in archived Provenance
(`changes/archive/`); no separate report inventories or hash manifests.

## State — 2026-11-01

**EC1–EC6, DR1–DR4 and FD1–FD4 passed.** The independent contracts, gitstore
and assets proofs remain complete and preserved below. Full product completion
remains open, with every later gate required.

The owner selected the next planned milestone: **stale-change authoring support,
then one independent project drill**. PLAN orders execution, ACCEPTANCE SP1–SP4
defines gates and GOAL is the restart prompt. Scope is `show --stale <change>`,
missing-bases-only `check --fix`, their required project/app integration, and
project's proof after those changes close. Project depends only on contracts,
gitstore and assets; the completed foundation evidence is reused where valid.

**0146 applied at `21d5dde`**: project's StaleSpecs/StampMissing APIs and
PRJ-060–066 are implemented. Genuine fail-first (`9e08d86`: 126 cases, 17
failing), restricted independent audit with no findings (`a313e6d`), and full
controller application/archive passed. Final implementation verification passed
all **798 cases**; retained interfaces/evals stayed byte-unchanged through impl.
Per-change evidence is [0146's archived Provenance](../../changes/archive/0146-stale-change-authoring.md).

**0147 applied at `78c334f`**: app exposes both options with existing JSON and
diagnostic conventions. Its 13 genuine failing cases cover all seven new IDs;
all 291 existing cases passed. Restricted audit passed all eight affected IDs,
then full controller application passed all **304 cases**. Existing evals and
interfaces stayed frozen. Evidence remains in [0147's archived Provenance](../../changes/archive/0147-stale-authoring-cli.md).
SP1 behavior and SP2's feature lifecycle are complete; candidate CI recovery
is proceeding below, and the independent project proof is pending.

**Candidate checks — recovery.** Inspectable clone:
`/private/tmp/fireweed-sp-project-8ko73j3j/candidate`. Current clean executable
source/aligned base `fed193a3f8d2c86dab33aef69972543e41d45e9f`, version 0.1.0,
`vcs.modified=false`, SHA-256
`58d032f11b42868d8ff25d9d93d44e4b996269f06c693bb70a928089d19fa3ea`.
This merges archived 0148 into the isolated candidate, retaining its input-only
SYSTEM/AGENTS policy. Runtime source, interfaces, configuration, setup and hooks
are unchanged from the earlier clean build at
`5bacb4ab5817e8eaea38218e6a8bcf9354dce42c`; only RUN-040's fake-agent wait changed.
Reuse passing full Go normal/race suites and vet for unaffected modules. Affected
runner normal/race suites passed in 269.433/283.520 seconds; runner vet passed.
The rebuilt candidate passed NEW check and all ten doctor --no-agent canaries.
Actual candidate CLI fixtures remain valid for unchanged behavior: history
recovery, fallback, mixed bases, file/index/ref effects, JSON and applicable exits;
inspectable fixtures: `/private/tmp/fireweed-sp-cli-y4uu9x98`.

NEW check-commits passed without diagnostics over
`be35552b706aa3000ff8f11306d129ff1d701b85..5bacb4ab5817e8eaea38218e6a8bcf9354dce42c`
(996 commits), the policy-only append
`5bacb4ab5817e8eaea38218e6a8bcf9354dce42c..f22a327811f0c188cbbbe239a544e27fda80ce5b`,
and the 11-commit correction/alignment append
`f22a327811f0c188cbbbe239a544e27fda80ce5b..fed193a3f8d2c86dab33aef69972543e41d45e9f`.
Native CI with doctor enabled passed without diagnostics in 1065.946 seconds over
the complete explicit range
`be35552b706aa3000ff8f11306d129ff1d701b85..fed193a3f8d2c86dab33aef69972543e41d45e9f`.
The following earlier failure is resolved by 0148 and this passing CI gate.
CI with doctor enabled over the combined baseline-to-alignment range first
reported E01 for existing `Test_RUN_040_EarlierBlockRetainsCauseBeforeTimeStop`.
That case passed normal and race suites, then an unchanged focused replay passed
in 8.38 seconds. The identical full CI replay also failed the same case. A
focused JSON replay with the CI hook's invoking-bin/PATH context passed in 8.35
seconds. Full CI attempts took 1612.351 and 996.942 seconds; both reported only
that E01 cause. No source, frozen eval, pin or configured bound was changed.
An unchanged protected-hook replay recovered the assertion: gamma advanced to
audit before the eight-second deadline, while the test expected held evals.
Its six-second fake-worker sleep can finish before that deadline. The replay
ran 132 cases in 234.914 seconds; only this held-phase assertion failed.

The owner authorized scoped RUN-040 recovery on 2026-11-01: diagnose and make any
necessary correction through the normal pinned-controller lifecycle, using
eval_fix for frozen eval corrections and preserving cases, assertions, ID coverage,
pin and configured bounds. Change 0148 was authored at `90c2875`, but the author
mistakenly used `kind: eval_fix` with `fixes` instead of the documented
`eval_fix: [RUN-040]`. Its one-line worker correction waits for cancellation
instead of sleeping six seconds; every assertion and deadline is unchanged.
The controller's runner hook passed all 132 cases in 284.436 seconds, then
correctly required fail-first for the wrongly authored kind. The owning next
was interrupted (exit 130) before further verification; its worker stopped and
the claim was released. No phase commit, audit, archive or block occurred.
Run `20261101T220404-446c231a1bbfdb04e783d235` and its evals work remain preserved.

**Owner-approved author metadata repair.** On 2026-11-01 the owner approved
correcting 0148 to `eval_fix: [RUN-040]` and restarting its lifecycle, explicitly
allowing the attempt-accounting restart caused by the authoring error. The
stopped run's one-line eval correction is preserved for the restarted worker;
author repair commits contain no eval edits. Resume the normal pinned-controller
lifecycle with the same cases, assertions, ID coverage, pin and bounds.
**0148 applied at `a11ace0`** after that repair. The normal eval_fix lifecycle
passed all 132 own cases (`4e76986`), its restricted independent audit had no
findings for RUN-040 (`376b45f`), and full implementation verification passed
all 436 cases. Only the one fake-agent wait changed; every existing case,
assertion, covered ID, runtime source, interface, spec and configured bound was
preserved. Evidence remains in
[0148's archived Provenance](../../changes/archive/0148-run-time-stop-fixture.md).
Main is clean and unclaimed. The candidate includes this correction; affected
runner checks and full CI passed. Reuse unaffected Go, CLI and
foundation evidence. The real project drill result and owner checkpoint follow.

**SP3 — independent project drill failed; owner checkpoint.** The one authorized
native drill ran from clean `fed193a3f8d2c86dab33aef69972543e41d45e9f` with the
recorded candidate executable, exact configured real worker and unchanged bounds.
Live inspection confirmed 19 byte-identical retained inputs, fresh one-commit Git
metadata with no remotes/alternates, no supplied old code/binaries/manifests or
dependent internals, an empty disposable target in the separate verifier, and an
exact copied invoking controller. Review covered all 50 shell commands, their
complete script bodies and recorded patch operations across both configured
attempts. Reads stayed within supplied inputs/new files; no old implementation,
outside checkout, saved attempt, binary or invoker-content oracle was used.

The native result is **fail** after its configured retry, in 2023.229 seconds.
It retains **795 passing / 16 failing testcase entries**: project 126/0,
app 289/15, checks 101/0, runner 132/0 and workflow 147/1. Parent/subtest failures
are separate native entries. Project's own evals pass, but dependent diagnostic
contracts do not: cause wording/corrective guidance, show-history propagation of
U02/C01/C03 and exit status, and a blank base being classified stale instead of
missing. The CLI names only the first failing case per affected module; the
native report retains every failure. No contradicted Decision was identified.
The agent's no-manifest note is the intended input-only verification limitation;
the actual failed dependent evals, not that limitation, prevent a passing proof.

Retain the native report verbatim at
[.fireweed/drills/project.md](../../.fireweed/drills/project.md), with rebuilt ref
`refs/fireweed/drills/project/20261102T000116.032120000` pointing to
`dc7a84f1ae92d7f0bab3b91cc293882427591fad`, in main and the inspectable candidate.
The ref's 19 retained inputs are byte-identical; changes are confined to project
disposable code and DRILL_NOTES.md. Main's implementation was not replaced.
The drill left its invoking HEAD/executable unchanged. Candidate report-only
author commit `41bdb56` then retained the reviewed evidence with a clean checkout;
executable source remains `fed193a`. Native input and verifier worktrees were
removed, and claims are absent. Preserve the inspectable clone.
The owner authorized scoped diagnosis and normal audited corrections for these
project/app/workflow diagnostic contracts, then one fresh project drill after
affected checks pass, on 2026-11-01. Preserve cases, assertions, ID coverage,
main's pin and configured bounds; use eval_fix for wrong frozen evals.
The failed report/ref and inspectable clone remain preserved unchanged.

Scoped diagnosis selected two coherent changes. 0149 clarifies existing project
compatibility contracts with already_satisfied: missing versus stale bases,
unknown-module/create-file causes, history code prefixes, and module diagnostic
paths. Main already supplies these; its runtime implementation is retained.
0150 is an app eval_fix: recognize meaningful correct/remove/rename guidance,
and validate the malformed-history fixture's C03/path without requiring the word
"parse" or one error representation. Every behavioral assertion remains required.
No workflow change or runtime replacement is currently needed.

**0149 applied at `c4d34a2`**: all 127 own cases passed (`96a5351`), restricted
audit passed all six items (`6596896`), and full implementation verification
passed all 812 cases. Runtime source and interfaces were unchanged. Existing
cases/assertions were retained or strengthened; null/empty-base variants reuse
an existing case. Evidence remains in
[0149's archived Provenance](../../changes/archive/0149-diagnostic-contracts.md).
Its first worker invocation exited before edits with no detailed CLI cause;
one bounded doctor probe passed both configured worker/auditor fixed replies.
The same owning controller and held run then resumed successfully, with no
model/effort change, fallback, verification-attempt reset or block.
**0150 applied at `b1a6015`**: its eval_fix passed all 304 app cases (`a0fd015`),
restricted audit passed APP-107/108/112/113 (`7eb3e4e`), and full implementation
verification passed all 304 cases. Runtime source, interfaces and specs were
unchanged. Evidence remains in
[0150's archived Provenance](../../changes/archive/0150-diagnostic-eval-contracts.md).
P03 rejected the first eval worker's premature phase edit before any worker
commit. The normal remaining attempt restored the approved evals state and
removed temporary manual provenance; final phase commits/provenance belong to
the configured outside controller. Explicit worker phase control was added to
Notes and aligned in SYSTEM between changes; no attempts or bounds were reset.

Both scoped recovery changes are archived, with main clean and unclaimed.
Update the isolated candidate, preserving its failed report/ref/history. Recheck
affected project/app normal/race tests and vet, NEW check and doctor, then NEW
history/CI on the explicit append after `fed193a`, reusing the passing original
`be35552..fed193a` prefix. This retains the standing history baseline and covers
the whole history without repeating unaffected traversal. CI still checks all
current module evals/coverage, boundaries and hook doctor. Identify the clean
updated executable, then use fresh independent inputs for the one newly
authorized project drill. Stop at the next owner checkpoint.

**Owner checkpoint — verification cost.** The initial worker full verify was
interrupted at the 15-minute guardrail (outer next exit 130), preserving work
and releasing its claim. The owner then approved scoped diagnosis with pin,
configured bounds and frozen tests unchanged. Diagnosis found project took
2.354 seconds; app, checks and runner took 258.332, 232.562 and 286.791 seconds
in that interrupted pass. An unchanged app probe passed in 262.355 seconds;
an identical cache probe took 1.901 seconds. Build probes did not justify a
VCS-stamping change. No performance source, hook or frozen-test fix was made.

Recovery kept the same run/impl phase and exact configured Codex worker; it used
quick worker feedback and temporary `GODEBUG=gocachetest=1` diagnostics. The
controller's uncached precheck completed at approximately **15m06s**, slightly
over SYSTEM's cost guardrail, before its ordinary done step reran full checks
from valid Go cache in approximately **12 seconds** and archived automatically.
All 126 project, 291 app, 101 checks, 132 runner and 148 workflow cases passed.
Cache logs show unchanged test binary IDs but different runtime-input IDs between
standalone probes and controller hooks. The protected hook omits JUnit durations,
so native Seconds fields are zero; stdout supplies uncached package timings.
That measurement limitation remains recorded, with no automatic fix.

The owner accepted this timing overrun and permanently raised SYSTEM's full
implementation Verify process guardrail from 15 to **25 minutes** on 2026-11-01.
Own Verify remains under 5 minutes. This changes no configured controller bound,
pin, frozen test, enforcement or evidence requirement. Continue SP1–SP4 and stop
at the next owner checkpoint.
SP1/SP2 are complete; SP3 failed and SP4 completion remains open during
owner-authorized diagnostic-contract recovery. Main's controller source/hash and
fireweed.yaml are unchanged; completed reports/refs/inspectable clones and the
unrelated worktree remain preserved. Only the earlier push through `dae9b28`
was authorized; later work remains local.

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
- FD1–FD4 are completed gitstore/assets evidence. SP1–SP4 is the next planned
  scope: stale authoring in project/app, then one project drill. Its dependency
  inputs are contracts/gitstore/assets interfaces only; dependency internals and
  invoking candidate content are excluded. Only ordinary feature implementation
  belongs on main; independently rebuilt project code stays in the proof ref.
  Real regenerate changes, controller promotion, other-module drills, unrelated
  deferred work and push remain outside SP1–SP4. Later gates remain required.

## Next

0146–0150 are applied and archived. Reassess affected candidate checks,
then run the one newly authorized fresh independent project drill.
Keep main's pin,
configured bounds, frozen tests and completed proofs unchanged; SYSTEM's
full-verification process guardrail is 25 minutes. Stop at owner checkpoints.
SP1–SP4 remains the selected incomplete milestone.
