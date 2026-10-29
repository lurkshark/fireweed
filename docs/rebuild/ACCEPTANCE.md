# Acceptance and traceability

Revised 2026-10-28. Full acceptance remains **unverified**; partial evidence is reusable. R01–R29 are planning labels, not another spec layer. Map design behaviors to atomic specs through changes as each slice needs them, following authorized work in PROGRESS.md.

The authoritative design snapshot is [fireweed-design-v1.md](../../fireweed-design-v1.md) in this repository; its historical origin and checksums are recorded in [HISTORY.md](HISTORY.md#historical-reference-inspection-record--superseded-by-d004).

GOAL.md/PLAN.md sequence runnable CLI slices; these rows describe eventual coverage, not implementation order. PROGRESS.md records the current slice, pinned controller and checkpoint authority. Prior evidence is preserved in HISTORY.md; historical pending-work statements do not restart completed work.

At safe author boundaries link rows to assigned IDs and durable eval/archive/report evidence. Shared evidence can satisfy overlapping rows; do not duplicate specs, tests or reports for each label. Record relevant exceptions and discrepancies briefly.

**Scope.** The design defines obligations; SYSTEM.md bounds their implementation and verification, including supported environments, process-group cancellation, eval cost and audit scope. Conventions use review and one app PATH-shim test, even where several rows mention them. Inventory examples add no requirements or unnamed edge cases. R22–R25 cover bounded effort, R26–R28 usability and R29 merge strategies; overlapping obligations need evidence once.

D006's planned existing-path migration is complete through 0068. Keep diagnostics
and explain accurate as justified features add real paths; unused allocations
wait for those features. Workflow file reorganization is not a prerequisite.
Neither statement waives eventual design obligations.

A slice closes with observable NEW CLI evidence under PLAN; passing module evals
alone does not prove it. Reuse that evidence across relevant rows rather than
adding a separate acceptance campaign for every inventory entry.

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

## Required command inventory

Exercise NEW beyond public function tests: every design command and documented option, read-command JSON, and representative invalid-input and exit behavior. Combine related checks in existing fixtures; no Cartesian product of commands, options, errors, recipes or platforms is required unless the design specifies the interaction. Old-controller syntax adds no requirements.

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

## Completion gates

Use one design/inventory review, then check fixes against affected requirements. Reuse valid evidence under SYSTEM's rules; required phase checks still run. All gates must pass; passing gates end verification unless later changes invalidate evidence. Rebuild-planning edits alone do not require product tests/drills.

1. **Traceability:** map design requirements to specs, referencing evals and archives; SYSTEM conventions use prescribed review/eval evidence. Verify every R row within that scope. Legitimate already_satisfied changes reuse existing assertions; review, quarantine or maintenance cannot waive design obligations. No extra spec layer or testcase inventory is required.
2. **Process:** bootstrap-managed fail-first/audit/impl/archive commits for behavioral work; immutable evals/interfaces during impl; separate protected setup/author commits; no pending/blocked rebuild changes or active claims.
3. **Go:** with `GOCACHE=/tmp/fireweed-rebuild-cache`, a fresh checkout builds the new binary and passes `go test ./...`, `go test -race ./...`, `go vet ./...`. Permanent tests run with local fixtures/fake agents and without bootstrap/agent authentication. Cross-compilation is not platform runtime testing; report limits accurately.
4. **Bootstrap checks:** controller `check`, `doctor`, `check-commits <setup-sha>..HEAD`, and `ci <setup-sha>..HEAD` with doctor enabled. Use the recorded stage-2 baseline; record the exact nonempty range and link per-module coverage. Resolve errors. For warnings, make a scoped correction or record the reason/limitation; zero warnings is not a separate gate. An empty default range proves no history.
5. **Product checks:** NEW binary in fresh repositories performs init/doctor, author/stamp/check, documented base approval, start/failing-evals/audit/impl/done, coverage/history/PR/CI. Exercise fixes, eval_fix, maintenance, blocked recovery and regenerate. Normal unattended approval rejection is tested without --allow-unmerged. Pin subject paths so bootstrap cannot be used accidentally.
6. **Reference comparison (smoke):** one equivalent fixture lifecycle run with both binaries, comparing exit codes, JSON envelope shape and resulting commits/refs. Record intentional differences briefly. Reference bugs do not define acceptance, and per-edge-case equivalence is not required.
7. **Real regeneration:** successful controller-driven drills for every rebuilt disposable module, dependency-first. Real agents reconstruct from retained specs/Decisions/interfaces/evals and permitted dependency interfaces, without old implementation/history, saved attempts, old binaries or bootstrap. Own/transitive dependent evals pass without preserved-contract edits. Commit reviewed reports/retain refs and clean scratch worktrees. Repeat only drills invalidated by relevant changes or unresolved failures. Gaps require human agreement before authored fixes; agent suggestions alone authorize none. Missing real-agent access blocks this gate.
8. **Self-hosting:** in an isolated clone/fixture, the NEW binary governs a small real spec change and a real drill, with checked fail-first/audit/impl commits, archive/provenance, CI and scratch cleanup. Keep the original repository governed by the pinned bootstrap throughout acceptance.
9. **Delivery:** README build/test/use instructions, locally validated distribution/installer checksum behavior, clean committed target tree, restartable evidence record and independent product executable. No product runtime read/import/command requires bootstrap. Report exact checks/commits/reports and limitations. Push only under explicit owner authorization recorded in PROGRESS; do not publish or tag.

## Evidence record shape

Link evidence produced when work runs; never prefill successes. This optional guide needs no separate report if change/progress records already supply the facts.

```text
Requirement(s) / assigned spec IDs:
Change/archive and phase commits:
Eval/report link; actual fail/pass counts and command/result:
Discrepancy, limitation or unrun work, if any:
Status: unverified | in progress | verified | blocked
```

Bulky logs may remain ignored. Commit concise findings/references once, without hash manifests or testcase inventories. Audit, platform, regeneration and self-hosting claims must match what ran.

## Evidence navigation

[PROGRESS.md](PROGRESS.md) records current capability status and the next slice.
[HISTORY.md](HISTORY.md) preserves the former progress ledger, historical partial
requirement mappings and reference inspection record. Linked reports and change
archives carry detailed phase/checkpoint evidence. Keep results once and link
them here when they establish coverage; do not infer current readiness from an
older entry or turn its historical next-work text into a new instruction.

The first NEW next lifecycle now has [actual Go CLI evidence](reports/next-0080-block.md)
for the successful-path portions of R06–R09/R17: genuine fail-first, separate
bounded fake audit, checked commits, SPEC/archive/provenance and claim release.
This is partial coverage; recovery, remaining kinds and all final gates
stay open. Reuse the result without repeating it for each inventory row.

The [bounded NEW run proof](reports/bounded-run-cli.md) adds partial R12/R13/R17
evidence for successful queue order, --max, later continuation and empty queue.
Retries/recovery, time/usage limits and remaining obligations remain open.

The [NEW verification-retry proof](reports/verification-retries-0084.md) adds
partial R12/R17 evidence for phase attempt bounds, retained work and latest
verification feedback. Saved/blocking recovery, independent continuation,
resume, nonzero-agent retries, usage/time stops and final gates remain open.

The [NEW next exhaustion proof](reports/exhausted-next-0086.md) adds partial R12
and R17 evidence: exhausted work/ref retained, phase-base checkout restored,
state-only block committed, original cause reported and claim released. Saved
attempt resume, projected waiting/independent continuation and other recovery
obligations remain open.

The [NEW independent-continuation proof](reports/independent-run-0089.md) adds
partial R12/R17 evidence: projected waiting, independent application after a
clean saved block, and retained exit1/cause/ref despite later successful work.
Saved-attempt loading and remaining recovery/final obligations remain open.

The [NEW saved-attempt resume proof](reports/saved-resume-cli-0091.md) adds
partial R06/R12/R17 evidence: loading real saved work through start --from-attempt,
retained lifecycle state/HEAD, correction and normal checked archive without
repeating earlier successful phases. The fixture was removed after PASS.
Agent-commit undo, usage/time stops and remaining final obligations stay open.

The [NEW agent-commit recovery proof](reports/agent-commits-0092.md) adds partial
R12/R17 evidence: soft-reset phase-agent commits, retained work and visible
recovery warnings, then normal checked archive. Usage/time stops, remaining kinds/
bounds and full acceptance remain open.

The [NEW usage-stop/resume proof](reports/usage-stop-cli-0095.md) adds partial
R13/R17 evidence: same-attempt usage retry, shared wait budget across applied
changes, exit3 with retained result/state/work, and later same-run completion.
Actual Go reports prove fail-first/pass; fixture removed. Time --for and all
unclosed inventory/final gates remain open.

The [NEW time-stop/resume proof](reports/run-time-cli-0097.md) adds partial
R13/R17 evidence: run --for interrupts an active agent, returns exit3 with
retained result/state/work, and next completes the same run with actual Go
fail-first/pass reports. Fixture removed. Remaining kinds/bounds and final gates
remain open.

The [NEW maintenance proof](reports/maintenance-cli-0098.md) adds partial
R06/R09/R17/R23 evidence: maintenance starts in impl, runs one agent, passes
normal Go checks and archives in one commit with retained contracts and claim
release. Eval-fix/regenerate kinds and other final obligations remain open.

The [NEW eval-fix proof](reports/eval-fix-cli-0099.md) adds partial R08/R09/R17/R23
evidence: corrected evals pass with listed coverage, separate audit includes
the ID, impl is change-file-only and normal archive releases the claim. Actual
Go baseline failure/correction and unchanged contracts/implementation are
recorded. Worker guidance, regenerate and remaining full gates stay open.
