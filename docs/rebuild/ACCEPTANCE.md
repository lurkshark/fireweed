# Acceptance

The ordered milestone plan and acceptance gates, with the full inventory of
design obligations. Completed criteria are historical evidence, not a queue to
repeat. Execute only the separately owner-authorized milestone.
R01–R29 are planning labels, not another spec layer. The design defines the
obligations; SYSTEM.md bounds how they are implemented and verified. Evidence is
committed archived changes with Provenance and native drill reports in
.fireweed/drills/; link it here only when a gate closes. Current status is in
[PROGRESS.md](PROGRESS.md).

## Requirement inventory

| Ref | Design section / proposed owner | Required behavior and representative evidence |
| --- | --- | --- |
| R01 | Repository layout; Distribution and init / app, assets, project | Init creates config, SYSTEM, managed AGENTS, queue/archive, protected recipe copy and optional app. Exercise no-app, invalid input, existing files and safe non-overwrite behavior. |
| R02 | Modules and zone map / contracts, project | Omitted defaults differ from explicit empty lists; unique ownership; strict zone precedence; interface/append-only survive regeneration; eval/disposable overlap rejected even for future paths; unique prefixes and acyclic existing dependencies; warn about out-of-glob specs. |
| R03 | Specs / contracts, project | Single Specs section, valid headings/normative text, review/quarantine semantics and quarantine reason; precise ID matching in name/classname across case/separators/token boundaries; removed IDs never reused; review/quarantine visibility. |
| R04 | Change files / contracts, project | Number/slug ordering including ties; unambiguous slugs; unique local handles and cross-change rejection; late IDs and all-occurrence rewriting; normalized SHA-256 bases including tags; stale/missing bases rejected; allowed set fields/immutable identity; spec/fix/maintenance/regenerate classification. |
| R05 | Specs and Change files / project, workflow | Projection includes preceding queue changes and current zones; Added/Modified full text, Removed headings only; apply Decisions with slug, remove through Removed decisions; archive-backed history/provenance; archived changes immutable. |
| R06 | Lifecycle / workflow, gitstore | Clean tree including untracked files; queue eligibility; atomic owner/worktree claim; late IDs/run/phase, start does not commit; stale-base/unapproved warnings; resume without new IDs; verify never commits; done advances/applies/archives with recovery on failure. |
| R07 | Evals step / checks, workflow | Only allowed eval/interface/new-stub/manifest/shared edits; no existing impl mutation. Every Added/Modified/fixes ID needs a failing testcase unless legitimately already_satisfied. No compile/collection/skipped/unrelated-failure loophole. Remove references to removed IDs. Test acceptance and rejection. |
| R08 | Audit step / workflow, runner | On by default/configurable; real ID and exact clause validation; findings return to evals/count attempts; untestable findings block with Questions; passing audit lists checked IDs; audit-only edits; command fallback and doctor warning. |
| R09 | Implementation step / workflow, checks | Immutable eval/interface/SPEC/protected/other-module files; only current change editable; append-only additions only. Own evals, configured dependents (direct by default) and boundaries pass; exercise all/direct/none and CI always-all. Apply metadata/specs/Decisions and archive intact. Negative tests for forbidden zones. |
| R10 | Language support / checks | JUnit testsuite/testsuites; every testcase mapped to current IDs; unknown/missing IDs and coverage rejected; skips do not cover; malformed/missing reports and hook errors fail; only actual failed/errored cases prove fail-first; reviews/quarantines respected. |
| R11 | Hooks/recipes/doctor / checks, assets | Shell placeholder quoting versus argv expansion; hook environment/module overrides; module-attributed boundary output; exact baseline handling; disposable modules require boundaries or visible reasoned opt-out; graph derived from module metadata. All four canaries detected and recovery verified, including each hook override. |
| R12 | Retries/blocking / runner | Attempts and timeouts; retries retain work/verification feedback; exhausted attempts saved as refs; block commit includes only state/Questions; checkout restored/claim released; projected dependency blocking with independent continuation; from-attempt resume; agent commits soft-reset before checking. |
| R13 | Limits / runner, gitstore | Usage patterns inspect final 20 lines; usage waits differ from failed attempts; wait budget/--for/--max; exit 3 keeps state/work recoverable; resume works; cancellation kills the hook or agent process group on macOS and Linux (processes that leave the group are a documented limitation). Tests use bounded durations, not real 30-minute waits. |
| R14 | Git behavior / gitstore, workflow | Locally known remote/base defaults/fallback; no product fetch/push/branch creation; linked-worktree claims; retained attempts/drill refs; release/gc; common-Git-dir scratch paths/setup/cleanup; duplicate-run conflict and commits/recovery guidance. |
| R15 | History / checks | H01 protected worker edits; H02 impl eval/interface edits; H03 append-only edits by everyone including combined merge resolutions; H04 archive mutation with deletion/revert allowed. Before/after zones, merge/rebase/squash limits and full revert to queued state. Parse state, not commit subjects. |
| R16 | Regeneration / runner, workflow | Maintenance regenerate deletes only disposable code and passes normal checks. Real-agent drill uses spec/Decisions/interfaces/evals in scratch, runs own/transitive dependents, reports contradicted Decisions/targeted gaps, retains ref, cleans up and reports drift. No automatic change authoring from gaps. |
| R17 | CLI / app, contracts, assets | Full inventory below, noninteractive argument/usage behavior, stable diagnostic codes/explain, versioned JSON envelope/diagnostics, documented exits 0/1/2/3 including empty queue; no hidden bootstrap delegation/network behavior. |
| R18 | CI/PR body / checks, app | CI runs static/history/all coverage/boundaries/doctor with documented skip; explicit ranges work. PR body uses range endpoint, covers changes, interface/stub/manifest/shared/preserved edits, review specs, audits, Decisions, quarantines, opt-outs and blocked changes. |
| R19 | Recipes / assets, app | Go/TypeScript-pnpm/Python-uv, none, local recipes/defaults/templates/canaries; accurate ready-made versus user-built glue instructions; substitutions work for real/random modules; copied recipe becomes project-owned and never auto-upgrades. |
| R20 | Versioning/playbooks/AGENTS / contracts, assets, app | Binary version-range enforcement; distinct format version with current/previous support/migration; embedded playbook hash, override embedded-base hash/stale detection; managed AGENTS refresh preserves surrounding content; explain/playbooks match checks. |
| R21 | Packaging/principles / app and author infrastructure | Single dependency-light Go binary; no runtime networking/update/telemetry; build/install docs; Linux/macOS amd64/arm64 packaging/checksums, explicit-version installer, Go install/Homebrew/release configuration. Validate without publishing or requiring a tap token. |
| R22 | Audit step; Change lifecycle (revised) / workflow, runner | `audit_rounds` state field incremented by audit-findings commits; re-audits keep only findings for IDs that had findings last round; after `audit.max_rounds` (default 2) the change blocks with remaining findings in Questions, whether driven by run, an agent calling done, or a person. |
| R23 | Kinds of change (revised) / contracts, project, workflow | `eval_fix` kind: evals-zone-only edits, listed IDs stay covered, all module evals pass instead of fail-first, audit confirms assertions, impl changes nothing but the change file and archives. |
| R24 | Approval (revised) / workflow, runner, app | Unapproved archived changes on HEAD counted against the base; start/next/run refuse more with P07 at `authoring.max_unapproved` (default 5) even with --allow-unmerged; status shows the count and P08 for three consecutive applied changes to one module. |
| R25 | Specs; Cost is measured; config (revised) / project, checks, app | C17 spec-shape and C18 repeated-clause warnings with `spec_limits`; per-module eval time, cases per spec and eval/disposable line ratio recorded and shown by status/pr-body with B01–B03 budget warnings; pr-body lists binary or >256 KB files added to disposable zones (B04); `verify_dependents` defaults to direct; playbooks carry the Motivation and audit-scope rules. |
| R26 | Approval; Abandoning a change (usability revision) / workflow, app | A change file on either `origin/<base>` or local `<base>` is approved and starts without P04; the unapproved-work bound counts against `origin/<base>` when it exists. `abandon` refuses applied, foreign-claimed or dirty-tree cases, reverts the change's commits and deletes the change file in one non-worker commit, releases the claim, and frees its never-applied IDs. |
| R27 | Driving the state machine; CLI (usability revision) / workflow, app | `verify --quick` runs only the target module's evals plus zone/state checks; `done` always runs the full checks. `context` returns full change, SPEC, interface, SYSTEM and playbook text plus paths and sizes of other step files, and `--full` adds contents. `status` names, for each waiting change, the change holding it and the blocking rule. |
| R28 | Doctor; stale changes; Decisions (usability revision) / checks, project, app | `doctor` runs the agent and audit commands with a fixed-reply prompt, reports failures with output and limit-pattern matches as usage limits, `--no-agent` skips it and `ci` never runs it. `show --stale` diffs each stale spec from its stamped text (or shows current text and hash when not found); `check --fix` stamps only missing bases. `check` warns C19 when module Decisions exceed `spec_limits.max_decision_words`. |
| R29 | Merge strategies; drills (merge-strategy revision) / checks, runner, app | `check-commits` recognizes a squashed change commit (sets `applied` when the parent did not have the change at `impl`), skips H01/H02 for it, reports it as squashed in JSON without a warning, and still applies H03/H04, including several changes in one squash. Drill drift and the regenerate warning compare implementation Provenance dates with the drill report date, so they work after squash and rebase merges. |

## Required command inventory for full completion

Full completion must exercise NEW beyond public function tests: every design
command and documented option, read-command JSON, and representative invalid-input
and exit behavior. Combine related checks in existing fixtures; no Cartesian
product of commands, options, errors, recipes or platforms is required unless the
design specifies the interaction. Old-controller syntax adds no requirements.

```text
help, version
init --recipe <go|typescript-pnpm|python-uv|none|local-path> --agent <agent> [--audit-agent <agent>] [--no-app]
doctor [--no-agent]
new-module <name> --prefix <P>
new-change <module> <slug>
stamp <change>
list, status, modules
show <module> [ID] [--projected | --history], show --stale <change>
context <change> [--full]
explain <code>, playbook <name>
check [--fix]
coverage <module> [--with-change <change>]
check-commits <range>
ci [<range>] [--no-doctor]
pr-body <range>
start <change> [--from-attempt <n>]
verify [<change>] [--quick], done [<change>]
next [--allow-unmerged]
run [--max N] [--for <duration>] [--allow-unmerged]
release <change>, commits <change>, abandon <change>
drill <module>
gc, migrate, agents-md --write
```

## Current milestone — Stabilization and cleanup

Owner-selected on **2026-11-02**. Execute A → B → C1–C8 in order; one normal
product change at a time. Execution state and the retained budget deadline are
in PROGRESS. CK1–CK4 are deferred. The requirement inventory,
command inventory and every full-completion gate remain required.

**Budget:** the owner selected **three days from execution start**. Record the
start and resulting deadline in PROGRESS when execution begins; stop for owner
review at that deadline or after C8, whichever comes first. Module boundaries
can require multiple changes per stage. Keep normal approval/P07 and all other
bounds active; the time budget does not permit changing them.

**Stops:** stop at the first blocked change, failed verification unresolved under
SYSTEM's correction limit, access/independence failure or decision not covered
by this plan. Do not reset attempts/rounds, start replacement changes to evade
bounds, repin independently or push. No module drills, CK execution or other
follow-up is included after C8.

### A — Restore and validate the identical controller (author work)

With no active change/run/claim, build the controller from the source pinned in
PROGRESS in a clean temporary clone. Use `GOCACHE=/tmp/fireweed-rebuild-cache`
and matching `-trimpath` settings if discoverable. Compare the full SHA-256
against the single controller pin in PROGRESS. Install at its pinned path only
if identical; on mismatch, stop and report expected and actual hashes to the
owner, without repinning or recording another hash in planning documents.

The supplied outside review reported the pin, subject bin directory and fireweed
refs absent. Planning-time observation found the pinned executable present and
matching. This does not complete A's reproducible build or validation; reconcile
that snapshot during A without re-investigating the supplied fixture diagnosis.
Never select bootstrap automatically or resume another controller's claim.

With the restored controller, run `status --json`, `check` and
`doctor --no-agent`, and report results before proceeding. Worker PATH starts
with the current tree's `.fireweed/bin`; scratch setup copies the invoking
controller. Mutable `bin/fireweed` remains the product under test.

### B — Protected process cleanup (one author commit)

With no active change/run/claim, review and finish the prepared documentation:

1. **B1 — SYSTEM.** Use the exact committed-evidence rule, remove completed
   milestone-code references, shorten Scope guardrails to roughly half their
   former length while retaining applicable rules, and require a milestone
   budget in changes or days followed by owner review when reached.
2. **B2 — PROGRESS.** One screen or less: current state, single controller
   path/source/SHA-256, next work in C and open problems. Remove the historical
   narrative; git history retains it. No temporary paths, evidence inventories,
   case counts or additional binary hashes in planning documents.
3. **B3 — ACCEPTANCE and GOAL.** Make this milestone current and defer CK1–CK4;
   retain the full-completion requirements and a restartable ordered plan.
4. **B4 — Commit.** Keep B1–B3 in one protected author commit, separately from
   product changes. On 2026-11-02 the owner authorized committing/pushing the
   initial planning update before A. After A, review the committed B edits and
   commit any remaining protected cleanup separately before C. This exception
   starts no product work and authorizes no product-work push.

### C — Product changes through the normal lifecycle

Author each change with the pinned controller's `author-change` playbook.
Read projected specs, stamp/check/show the authored change, and run `context`
before each step. Keep normal fail-first (or legitimate already_satisfied),
audited eval_fix, independent restricted audits, frozen phases and full
verification. Workers use `fireweed verify` and never commit; fireweed owns
phase/application commits and archives. Preserve every existing assertion.
Batch related work within module boundaries; finish each stage before the next.

1. **C1 — Fixture eval_fix (runner, app and any other affected fixture owner).**
   Accept the supplied diagnosis: macOS sh kills scripts receiving long prompt
   arguments; it is not a timeout. Change fake agents and every fixture passing
   `{prompt}` to a script to use `bash`, or supply the prompt through stdin.
   Keep all assertions. Completion requires `go test ./...` passing on this Mac;
   do not weaken coverage, raise timeouts or investigate the OS failure again.
2. **C2 — Agent logs (runner spec change).** Persist every agent invocation's
   stdout and stderr at `.fireweed/runs/<run>/<phase>-<attempt>.log`, implementing
   the design's run identifier that also locates its logs. Keep output inside
   the existing gitignored runs directory.
3. **C3 — Prompt file (runner/workflow spec change).** Add `{prompt_file}` to
   `agent.command` and `audit.command`: write the prompt to a temporary file and
   substitute its path. Keep `{prompt}` behavior unchanged. Coordinate any
   necessary shared config/interface work through its owning module.
4. **C4 — Small fixes, in as few coherent module changes as possible.**
   Unknown change names or filenames supplied to `start`, `verify`, `done` and
   other change selectors receive an unknown-change diagnostic listing valid
   slugs, instead of P01 not-eligible-queue-head. Accepting `NNNN-slug` aliases
   is also allowed. In audit, `verify` applies the same submission checks as
   `done`; incomplete submissions fail with P09. The Go recipe and module
   template explain the preceding-letter boundary in eval names: use
   `Test_GRT_001_Name`; `TestGRT001` does not match. `init` writes
   `>=MAJOR.MINOR, <MAJOR.(MINOR+1)` rather than `=X.Y.Z`.
5. **C5 — gc.** Implement the design's stale-claim pruning, attempt-ref pruning
   for archived changes and drill-ref pruning beyond a configured age. Add the
   config key and document its default through normal owning-module changes.
6. **C6 — migrate.** Implement the design's migration command. For format 1,
   a verified no-op reporting `already current` is sufficient; include JSON
   output and diagnostics.
7. **C7 — Additional recipes.** Add `typescript-pnpm` and `python-uv`, each
   including its recipe document, module templates, canaries and default hooks,
   following the Go recipe. Each change's proof runs `doctor --no-agent` in a
   fresh scratch project made by `init --recipe <name>`. If pnpm or uv is not
   installed, stop and report; do not skip the recipe or its proof.
8. **C8 — Real gitstore regeneration.** Run exactly one real `regenerate: true`
   gitstore change through the normal lifecycle and report the result. Supply
   retained contracts/evals and approved dependency interfaces; never read
   bootstrap, deleted implementation, old history/binaries or saved attempts
   as reconstruction inputs. This is a regenerate change, not a module drill.

**Final report after execution:** applied change slug and commit for each change;
anything blocked/skipped with its reason; live `status --json` unapproved count;
and the `go test ./...` result on this machine. Stop after C8 even if other
full-completion gates remain open. Do not push.

## Completed milestones

Core cutover, drill support/contracts proof, foundation proofs and stale-change
authoring/project proof remain completed. Their detailed historical handoffs and
criteria remain in git history; do not repeat them unless relevant changes or
unresolved failures invalidate their proof. Committed per-change evidence is in
[changes/archive](../../changes/archive/), including Provenance. Native reports:
[contracts](../../.fireweed/drills/contracts.md),
[gitstore](../../.fireweed/drills/gitstore.md),
[assets](../../.fireweed/drills/assets.md),
[project](../../.fireweed/drills/project.md).
Preserve their committed records and required proof refs; temporary artifacts
are not evidence. Full product completion remains open.

## Deferred milestone — Checks contracts and proof (CK1–CK4)

**Deferred by the stabilization milestone selected on 2026-11-02.** All CK gates
remain pending and need separate owner selection, including a budget in changes
or days. They cover checks-owned obligations from the unchanged design
(R07/R09–R11/R15/R18/R25/R28/R29) and its independent proof (R16/full gate 7).
They add no product scope and are not part of the current A–C plan. Existing committed
archives/reports remain available; only relevant invalidation calls for new proof.

1. **CK1 — Focused checks contract review and necessary corrections.** Review
   checks' current spec/Decisions/interfaces/evals once against its existing
   JUnit/evidence/coverage, zone, hook/boundary, doctor and history obligations.
   Record concise conclusions and close only concrete unmet clauses or
   demonstrated failures within checks. No implementation is presumed needed.
   Each necessary change uses normal fail-first unless legitimately already
   satisfied, separate restricted independent audit and full phase checks;
   preserve existing cases, assertions, ID coverage and behavior. Wrong frozen
   evals use eval_fix. Review fixes only against affected requirements. Ambiguous
   behavior or required changes outside checks, including dependency interfaces,
   stop for owner agreement before authoring; no speculative hardening or
   repository-wide spec rewrite. Finish and archive scoped changes before CK3.
2. **CK2 — Identified clean NEW candidate.** Pass applicable Go normal/race/vet
   checks, NEW check, doctor --no-agent, explicit-baseline check-commits and CI
   with hook doctor enabled. Exercise actual NEW CLI paths affected by any
   correction. Keep exact executable identity and explicit nonempty history
   ranges in the applicable committed evidence, with only the controller pin
   in planning. Preserve the standing baseline and reuse valid prefixes and
   unaffected evidence without skipping required phase checks. Explain warnings
   without unrelated cleanup. No live scoped change or claim remains before CK3.
3. **CK3 — One independent checks drill.** In an isolated verified NEW candidate
   clone, align instructions/setup/PATH and confirm setup copies the invoker.
   Inspect independence before the real invocation and in the live worker
   input/verifier. Supply retained checks spec/Decisions/interfaces/evals and
   approved retained files, plus only declared contracts/gitstore/assets/project
   interfaces, including model interfaces. Keep actual dependency/dependent
   implementations, full manifests and hooks outside reconstruction inputs in
   the separate verifier. Exclude old target disposable code/data, history,
   saved attempts, binaries, bootstrap and other-module internals; never query
   invoker content as an oracle. Input Git must expose no old history,
   remotes or alternates. Use the exact configured real agent/model/
   effort and unchanged bounds. Checks and every transitive dependent
   (workflow, runner, app) must pass enforced evals and coverage with retained
   inputs byte-unchanged. Review every command, full script, incremental patch
   and native report fields/cases, contradicted Decisions and gaps.
4. **CK4 — Reviewed evidence and clean handoff.** Stop on any blocked change,
   independence/access failure or gap/contradicted Decision needing owner
   agreement. No automatic gap fixes, fresh drill after exhausted attempts,
   counter resets or expanded scope. Owner-approved recovery follows the normal
   audited lifecycle, preserves frozen obligations and reassesses affected
   evidence; a fresh drill needs owner authorization. Retain the reviewed native
   `.fireweed/drills/checks.md` and `refs/fireweed/drills/checks/<date>`; commit
   the report as protected author evidence with no live run/claim.
   Update PROGRESS with current state and next work; keep per-change evidence
   in archived Provenance and drill results in the committed native report.
   Verify clean checkouts, scratch cleanup, no active run/claim, unchanged pin/bounds and
   preserved prior proofs/unrelated worktrees. No supplemental reports, testcase
   inventories or hash manifests. Completion requires CK1/CK2 and a passing
   independent checks proof with no unresolved owner checkpoint.

No rebuilt-code import/replacement on main, real regenerate change, controller
promotion, workflow/runner/app real drill, gc/migrate, additional recipes,
compatibility/file-format migration, unrelated warning cleanup, delivery or push
is included in CK1–CK4. Later drill order is workflow → runner → app, each owner-selected.
Remaining drills alone do not complete the product; all full gates below remain.
Stabilization now selects real regenerate validation, gc/migrate and additional
recipes before these deferred gates. Further version/file-format work and delivery
remain later work. Reassess these gates when selected, and repeat only evidence
invalidated by relevant changes or unresolved failures.

## Full-completion gates (still required)

Use one design/inventory review, then check fixes against affected requirements. Reuse valid evidence under SYSTEM's rules; required phase checks still run. Every gate below must pass for full product completion; the completed core cutover did not waive them. Passing applicable gates ends verification unless later changes invalidate evidence. Rebuild-planning edits alone do not require product tests/drills.

1. **Traceability:** map design requirements to specs, referencing evals and archives; SYSTEM conventions use prescribed review/eval evidence. Verify every R row within that scope. Legitimate already_satisfied changes reuse existing assertions; review, quarantine or maintenance cannot waive design obligations. No extra spec layer or testcase inventory is required.
2. **Process:** the recorded controller owns fail-first/audit/impl/archive commits (bootstrap before EC6, pinned NEW after it) for behavioral work; immutable evals/interfaces during impl; separate protected setup/author commits; no pending/blocked rebuild changes or active claims.
3. **Go:** with `GOCACHE=/tmp/fireweed-rebuild-cache`, a fresh checkout builds the new binary and passes `go test ./...`, `go test -race ./...`, `go vet ./...`. Permanent tests run with local fixtures/fake agents and without bootstrap/agent authentication. Cross-compilation is not platform runtime testing; report limits accurately.
4. **Controller checks:** the recorded controller runs `check`, `doctor`, `check-commits <setup-sha>..HEAD`, and `ci <setup-sha>..HEAD` with doctor enabled. Use the stage-2 baseline in GOAL; retain the exact nonempty range and per-module coverage in committed evidence. Resolve errors. For warnings, make a scoped correction or record the reason/limitation; zero warnings is not a separate gate. An empty default range proves no history.
5. **Product checks:** NEW binary in fresh repositories performs init/doctor, author/stamp/check, documented base approval, start/failing-evals/audit/impl/done, coverage/history/PR/CI. Exercise fixes, eval_fix, maintenance, blocked recovery and regenerate. Normal unattended approval rejection is tested without --allow-unmerged. Pin subject paths so bootstrap cannot be used accidentally.
6. **Reference comparison (smoke):** one equivalent fixture lifecycle run with both binaries, comparing exit codes, JSON envelope shape and resulting commits/refs. Record intentional differences briefly. Reference bugs do not define acceptance, and per-edge-case equivalence is not required.
7. **Real regeneration:** successful controller-driven drills for every rebuilt disposable module, dependency-first. Real agents reconstruct from retained specs/Decisions/interfaces/evals and permitted dependency interfaces, without old implementation/history, saved attempts, old binaries or bootstrap. Own/transitive dependent evals pass without preserved-contract edits. Commit reviewed reports/retain refs and clean scratch worktrees. Repeat only drills invalidated by relevant changes or unresolved failures. Gaps require human agreement before authored fixes; agent suggestions alone authorize none. Missing real-agent access blocks this gate.
8. **Self-hosting:** in an isolated clone/fixture, the NEW binary governs a small real spec change and a real drill, with checked fail-first/audit/impl commits, archive/provenance, CI and scratch cleanup. Keep the original repository governed by its recorded controller; EC6 may already have promoted NEW. The real-drill portion remains required for full completion, separately from the earlier core self-hosting proof.
9. **Delivery:** README build/test/use instructions, locally validated distribution/installer checksum behavior, clean committed target tree, concise restartable handoff and independent product executable. No product runtime read/import/command requires bootstrap. Report exact checks/commits/reports and limitations. Push only on explicit owner instruction; do not publish or tag.

