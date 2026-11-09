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
| R17 | CLI / app, contracts, assets | Full inventory below, noninteractive argument/usage behavior, stable diagnostic codes/explain, versioned JSON envelope/diagnostics, documented exits 0/1/2/3 including empty queue; no network behavior. |
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

Full completion must exercise the product binary beyond public function tests: every design
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

## Completed milestone — Workflow drill (WF1–WF4)

The earlier workflow contract review and candidate validation are in git
history (before `515516a`); their evidence stays valid where nothing it
covered has changed.

1. **WF1 — Simplified drill.** Run the queued
   [0195](../../changes/0195-run-argv-contract.md) and
   [0196](../../changes/0196-drill-self-checked-reconstruction.md) through the
   normal lifecycle. 0196 gives the reconstruction worker a buildable input
   workspace and its own eval hook, without the old implementation, history
   or dependents.
2. **WF2 — Controller.** With no run or claim active, rebuild
   `.fireweed/bin/fireweed` from clean main and record its commit in PROGRESS.
   `go test ./...`, `go vet ./...` and `fireweed ci` pass on that commit.
3. **WF3 — Workflow drill.** Run `fireweed drill workflow` once with the
   configured worker. Don't add wrappers, supervisors or extra prompts.
4. **WF4 — Report reviewed and gaps triaged.** Commit
   `.fireweed/drills/workflow.md`. Review its own and dependent eval results,
   contradicted Decisions and gaps with the owner. Turn each agreed gap into
   a queued change, or record why it is not one. A failing drill closes WF4
   once its gaps are triaged; repeat the drill only after changes that
   address its gaps have been applied.

WF1–WF4 are complete. The [workflow drill report](../../.fireweed/drills/workflow.md)
passed, and [PROGRESS.md](PROGRESS.md) records the owner-approved gap dispositions.

The requirement/command inventories and all full-completion gates remain.

## Completed milestones

Checks contracts and proof (CK1–CK4) are complete. The reviewed
[checks report](../../.fireweed/drills/checks.md) contains the contract review,
candidate validation and passing independent proof. The original CK procedure
is retained in commit `c34c3cb`.

Core cutover, drill support/contracts proof, foundation proofs, stale-change
projection/project proof, and stabilization A → B → C1–C8 are complete. The
completed stabilization procedure and former budget remain in git history;
they are not a queue to execute again. Its additional recipes passed fresh-init
doctor proofs, and its single real gitstore regenerate passed normal full
verification. That change does not substitute for the remaining module drills.
Per-change evidence remains in [archived changes](../../changes/archive/),
including [recipes](../../changes/archive/0168-additional-language-recipes.md)
and [gitstore regeneration](../../changes/archive/0169-real-gitstore-regeneration.md).
Retain the native [contracts](../../.fireweed/drills/contracts.md),
[gitstore](../../.fireweed/drills/gitstore.md),
[assets](../../.fireweed/drills/assets.md) and
[project](../../.fireweed/drills/project.md) reports and their required proof refs.
Reuse valid evidence; repeat only when relevant changes or unresolved failures
invalidate it and the selected scope authorizes repetition.

## Full-completion gates (still required)

Use one design/inventory review, then check fixes against affected requirements. Reuse valid evidence under SYSTEM's rules; required phase checks still run. Every gate below must pass for full product completion; the completed core cutover did not waive them. Passing applicable gates ends verification unless later changes invalidate evidence. Rebuild-planning edits alone do not require product tests/drills.

1. **Traceability:** map design requirements to specs, referencing evals and archives; SYSTEM conventions use prescribed review/eval evidence. Verify every R row within that scope. Legitimate already_satisfied changes reuse existing assertions; review, quarantine or maintenance cannot waive design obligations. No extra spec layer or testcase inventory is required.
2. **Process:** the recorded controller owns fail-first/audit/impl/archive commits for behavioral work; immutable evals/interfaces during impl; separate protected setup/author commits; no pending/blocked rebuild changes or active claims.
3. **Go:** with `GOCACHE=/tmp/fireweed-rebuild-cache`, a fresh checkout builds the new binary and passes `go test ./...`, `go test -race ./...`, `go vet ./...`. Permanent tests run with local fixtures/fake agents and without agent authentication. Cross-compilation is not platform runtime testing; report limits accurately.
4. **Controller checks:** the recorded controller runs `check`, `doctor`, `check-commits <setup-sha>..HEAD`, and `ci <setup-sha>..HEAD` with doctor enabled. Use the explicit history baseline `be35552b706aa3000ff8f11306d129ff1d701b85`; retain the exact nonempty range and per-module coverage in committed evidence. Resolve errors. For warnings, make a scoped correction or record the reason/limitation; zero warnings is not a separate gate. An empty default range proves no history.
5. **Product checks:** the product binary in fresh repositories performs init/doctor, author/stamp/check, documented base approval, start/failing-evals/audit/impl/done, coverage/history/PR/CI. Exercise fixes, eval_fix, maintenance, blocked recovery and regenerate. Normal unattended approval rejection is tested without --allow-unmerged.
6. **Real regeneration:** successful controller-driven drills for every rebuilt disposable module, dependency-first. Real agents reconstruct in the input workspace the product's drill builds (RUN-053), without old implementation/history, saved attempts or old binaries. Own/transitive dependent evals pass without preserved-contract edits. Commit reviewed reports/retain refs and clean scratch worktrees. Repeat only drills invalidated by relevant changes or unresolved failures. Gaps require human agreement before authored fixes; agent suggestions alone authorize none. Missing real-agent access blocks this gate.
7. **Self-hosting:** in an isolated clone/fixture, the product binary governs a small real spec change and a real drill, with checked fail-first/audit/impl commits, archive/provenance, CI and scratch cleanup. Keep the original repository governed by its recorded controller. The real-drill portion remains required for full completion, separately from the earlier core self-hosting proof.
8. **Delivery:** README build/test/use instructions, locally validated distribution/installer checksum behavior, clean committed target tree, concise restartable handoff and independent product executable. Report exact checks/commits/reports and limitations. Push only on explicit owner instruction; do not publish or tag.

