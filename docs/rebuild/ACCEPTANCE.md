# Acceptance

The full inventory of design obligations, the early core-cutover gates, and
the later gates that end the rebuild.
R01–R29 are planning labels, not another spec layer. The design defines the
obligations; SYSTEM.md bounds how they are implemented and verified. Evidence is
the archived changes, their evals and actual NEW runs; link it here only when a
gate closes. Current status is in [PROGRESS.md](PROGRESS.md).

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

The early milestone uses the scoped command list below. Full completion must
exercise NEW beyond public function tests: every design command and documented option, read-command JSON, and representative invalid-input and exit behavior. Combine related checks in existing fixtures; no Cartesian product of commands, options, errors, recipes or platforms is required unless the design specifies the interaction. Old-controller syntax adds no requirements.

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

## Early Go core-loop cutover — EC1–EC6

The owner selected this milestone on 2026-10-30. It permits promotion before
real drills/regeneration, remaining recipes, gc/migrate, show --stale/check --fix,
version-range/file-format migration, C17–C19 warning completeness and packaging.
Those obligations remain open for full completion. Existing deterministic Go
evals still run, including fixture tests of already-built deferred behavior.
No real drill or regeneration run is authorized by this milestone.

The adoption scope is Go/current-format authoring and normal changes: init and
doctor; new-module/new-change/stamp/check; context/start/verify/done; spec, fix,
eval_fix and maintenance; next/run with existing max/time/usage bounds; approval,
independent audits, saved attempts and blocked recovery; release/commits/abandon;
list/status/modules/show, coverage, check-commits, CI and PR body. Approval/P07,
phase zones, frozen evals/interfaces, fail-first, audit rounds and history rules
remain mandatory. Permanent fixtures may supply negative and bounded-limit
proof; there is no new Cartesian product of scenarios to run.

EC labels are execution gates, not product specs. All are pending until their
evidence is recorded. A focused review can discover a further concrete core
blocker; author a scoped change for that failure and recheck affected evidence.
Do not turn the review into speculative hardening or deferred feature work.

1. **EC1 — Complete prerequisites.** Finish the preceding 0132/0133 diagnostic
   queue, integrate this branch's authored files onto current main while its run
   is stopped, then apply 0134–0138 through bootstrap with normal lifecycle
   evidence. All five are required, including the guidance improvements.
   Resolve every blocked cutover change; require a clean execution checkout and
   no shared claims before handoff. A brief gap between claims in a live run
   is not an integration/promotion window.
2. **EC2 — Focused core review.** Review the scoped commands and safeguards once
   against the design, current specs and existing CLI/public evals. Check fail-first,
   immutable metadata/zones, independent audit input isolation and Audit-only
   import, approval/P07, saved-attempt recovery, agent commit recovery and
   clean time/usage stops. Confirm 0135/0136 carry JUnit-only dependent assertions
   into bounded retry feedback. Record concrete unresolved findings and known
   intentional bootstrap/NEW differences briefly, then review only affected
   fixes. This planning change does not claim that review is complete.
3. **EC3 — Fresh-checkout checks.** At the exact candidate commit, build NEW and
   pass `go test ./...`, `go test -race ./...` and `go vet ./...` using
   GOCACHE=/tmp/fireweed-rebuild-cache. Run controller and NEW `check`,
   `doctor --no-agent`, `check-commits <baseline>..HEAD` and `ci <baseline>..HEAD`
   over a recorded explicit nonempty range using PROGRESS's history baseline.
   CI includes hook doctor checks; the real-worker gate proves agent access.
   Resolve errors and explain warnings without treating zero warnings as a gate.
   These checks must use clean candidate trees, not another worker's live files.
4. **EC4 — One reference lifecycle smoke.** Run an equivalent ordinary spec-change
   fixture through bootstrap and NEW. Compare exit codes, JSON envelope shape,
   phase/commit outcomes and resulting ref shape, allowing generated IDs/times
   and commit hashes to differ. Reuse existing deterministic hook/fake-agent
   fixtures. Record intentional differences in DECISIONS; reference bugs do not
   define acceptance. Per-edge-case equivalence is not required.
5. **EC5 — Real-worker self-hosting without a drill.** In an isolated clone of
   the candidate, configure instructions, setup and worker PATH explicitly for
   NEW and make NEW govern one small real spec change through failing evals,
   a separate restricted audit, implementation, full verification and archive.
   Use the configured Codex model/effort and normal approval; inspect phase
   commits, Provenance, NEW CI and temporary-audit cleanup. Pin executable paths
   and identify the binary actually invoked so bootstrap cannot supply a false
   pass. Retain inspectable clone evidence until the concise handoff is recorded.
   Missing real-worker/auditor access blocks this gate. No drill is required.
6. **EC6 — Promote separately from the subject.** Only after EC1–EC5 pass and
   every repository run has returned with no shared claims, install a verified
   copy of NEW as the pinned controller (intended path
   `/Users/joel/Workspace/fireweed/.fireweed/bin/fireweed`). Keep mutable
   `bin/fireweed` as the product subject and retain the existing bootstrap
   executable as an explicit recovery option. In one protected author handoff
   commit, align AGENTS.md, SYSTEM.md, PROGRESS, recipe/setup instructions and
   worker PATH; ensure scratch hooks copy the invoking controller. Record the
   NEW source commit, binary hash, gate results and deferred scope in PROGRESS.
   Verify the pin and smoke-check `status --json`, `check` and
   `doctor --no-agent` with it. No claim may be resumed by a different controller,
   and no automatic fallback may mix them. This is core cutover, not a release
   or a declaration that the full design is complete.

Reuse valid gate evidence; repeat only what later changes invalidate. Keep
per-change evidence in archived Provenance and concise gate outcomes in
PROGRESS. No separate report files, case inventories or hash manifests.

## Full-completion gates (after early cutover)

Use one design/inventory review, then check fixes against affected requirements. Reuse valid evidence under SYSTEM's rules; required phase checks still run. Every gate below must pass for full product completion; EC1–EC6 separately authorize early cutover. Passing applicable gates ends verification unless later changes invalidate evidence. Rebuild-planning edits alone do not require product tests/drills.

1. **Traceability:** map design requirements to specs, referencing evals and archives; SYSTEM conventions use prescribed review/eval evidence. Verify every R row within that scope. Legitimate already_satisfied changes reuse existing assertions; review, quarantine or maintenance cannot waive design obligations. No extra spec layer or testcase inventory is required.
2. **Process:** the recorded controller owns fail-first/audit/impl/archive commits (bootstrap before EC6, pinned NEW after it) for behavioral work; immutable evals/interfaces during impl; separate protected setup/author commits; no pending/blocked rebuild changes or active claims.
3. **Go:** with `GOCACHE=/tmp/fireweed-rebuild-cache`, a fresh checkout builds the new binary and passes `go test ./...`, `go test -race ./...`, `go vet ./...`. Permanent tests run with local fixtures/fake agents and without bootstrap/agent authentication. Cross-compilation is not platform runtime testing; report limits accurately.
4. **Controller checks:** the recorded controller runs `check`, `doctor`, `check-commits <setup-sha>..HEAD`, and `ci <setup-sha>..HEAD` with doctor enabled. Use the recorded stage-2 baseline; record the exact nonempty range and link per-module coverage. Resolve errors. For warnings, make a scoped correction or record the reason/limitation; zero warnings is not a separate gate. An empty default range proves no history.
5. **Product checks:** NEW binary in fresh repositories performs init/doctor, author/stamp/check, documented base approval, start/failing-evals/audit/impl/done, coverage/history/PR/CI. Exercise fixes, eval_fix, maintenance, blocked recovery and regenerate. Normal unattended approval rejection is tested without --allow-unmerged. Pin subject paths so bootstrap cannot be used accidentally.
6. **Reference comparison (smoke):** one equivalent fixture lifecycle run with both binaries, comparing exit codes, JSON envelope shape and resulting commits/refs. Record intentional differences briefly. Reference bugs do not define acceptance, and per-edge-case equivalence is not required.
7. **Real regeneration:** successful controller-driven drills for every rebuilt disposable module, dependency-first. Real agents reconstruct from retained specs/Decisions/interfaces/evals and permitted dependency interfaces, without old implementation/history, saved attempts, old binaries or bootstrap. Own/transitive dependent evals pass without preserved-contract edits. Commit reviewed reports/retain refs and clean scratch worktrees. Repeat only drills invalidated by relevant changes or unresolved failures. Gaps require human agreement before authored fixes; agent suggestions alone authorize none. Missing real-agent access blocks this gate.
8. **Self-hosting:** in an isolated clone/fixture, the NEW binary governs a small real spec change and a real drill, with checked fail-first/audit/impl commits, archive/provenance, CI and scratch cleanup. Keep the original repository governed by its recorded controller; EC6 may already have promoted NEW. The real-drill portion remains required for full completion, separately from the earlier core self-hosting proof.
9. **Delivery:** README build/test/use instructions, locally validated distribution/installer checksum behavior, clean committed target tree, restartable evidence record and independent product executable. No product runtime read/import/command requires bootstrap. Report exact checks/commits/reports and limitations. Push only on explicit owner instruction; do not publish or tag.

