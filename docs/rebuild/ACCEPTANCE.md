# Acceptance and traceability

Prepared 2026-10-21, reset 2026-10-26; tightened 2026-10-27. Full acceptance remains **unverified**; partial evidence is reusable. R01–R29 are planning labels, not another spec layer. Map design behaviors to atomic specs through changes as each slice needs them, following authorized work in PROGRESS.md.

The authoritative design snapshot is [fireweed-design-v1.md](../../fireweed-design-v1.md) in this repository; its origin and checksum are recorded in [PLAN.md](PLAN.md).

GOAL.md/PLAN.md sequence the work; these rows describe eventual coverage, not implementation order. PROGRESS.md records completed milestones, the pinned 0.1.0 controller and current checkpoints. Historical pending-work statements below do not restart completed work.

At safe author boundaries link rows to assigned IDs and durable eval/archive/report evidence. Shared evidence can satisfy overlapping rows; do not duplicate specs, tests or reports for each label. Record relevant exceptions and discrepancies briefly.

**Scope.** The design defines obligations; SYSTEM.md bounds their implementation and verification, including supported environments, process-group cancellation, eval cost and audit scope. Conventions use review and one app PATH-shim test, even where several rows mention them. Inventory examples add no requirements or unnamed edge cases. R22–R25 cover bounded effort, R26–R28 usability and R29 merge strategies; overlapping obligations need evidence once.

Current D006 migration changes diagnostics only on existing error paths. Unused
allocations, their tests and explain entries wait for corresponding justified
feature implementation; this does not waive eventual diagnostic requirements.
Workflow file reorganization is not an acceptance or Milestone 2 prerequisite.

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
9. **Delivery:** README build/test/use instructions, locally validated distribution/installer checksum behavior, clean committed target tree, restartable evidence record and independent product executable. No product runtime read/import/command requires bootstrap. Report exact checks/commits/reports and limitations. Do not push, publish or tag.

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

## Historical partial evidence (not current work instructions)

The entries below record results and outstanding work at their cited revisions. Later entries may resolve earlier gaps; use PROGRESS.md and linked current reports for status. Reuse valid evidence without treating an incomplete row as a requirement to redo its passing portions.

- R17/R20: AST-001–AST-003 cover embedded version/help, initial diagnostic lookups
  and five playbooks through public assets evals `Test_AST_001_*`,
  `Test_AST_002_*`, `Test_AST_003_Explanations`. Change
  `changes/archive/0001-standalone-information.md`; phase commits
  `a4ceb17` / `826b9e9` / `2248b8e`; 8 assertion failures before and 8 passes after.
  CLI wiring, full diagnostics/versioning/playbooks/AGENTS remain unverified.

- R17: APP-001–APP-003 cover standalone NEW-binary invocation outside a project
  and with invalid config, schema-1 information JSON, usage errors and explicit
  unsupported-command failures; 30 assertion failures before, 30 passes after.
  Archive `0002-standalone-cli.md`; evals `1298acf`, controller audit `7de7d30`.
  Independent audit found APP-003 JSON gaps discarded by the controller because
  quotes folded line breaks. Correction through eval_fix remains pending.

- R17: `json-error-assertions` eval_fix repairs the APP-003 JSON gaps; phase SHAs
  `eb4704d` / `e184bb6` / `8796d37`. Separate restricted Sol/high audit passed
  APP-003; all 26 app cases pass, with unchanged implementation. This is rebuild
  process evidence, not evidence that the NEW product implements eval_fix.

- R02/R03/R04 partial: CON-001–CON-005 provide document/section/item parsing,
  canonical hashes, ID matching and zone precedence/globs. Archive
  `0004-lifecycle-primitives.md`; evals `11d0180` (18 failures, 9 passes; all changed IDs fail), audit
  `6e341a4` (independent no findings), impl `8e37dd7` (27 passes, existing app/assets
  and direct dependents/boundaries pass). Full grammar, project validation and
  lifecycle remain unverified.

- R02/R04 partial: PRJ-001–PRJ-003 load the minimal fixture, validate module graph
  identities/dependencies and author queued skeletons without assigned IDs. Archive
  `0005-fixture-authoring.md`; initial evals `1bcba65`, audit findings `5c647a2`,
  corrected evals `bfee6b8` (24 failures, 1 pass), passing independent re-audit
  `06f17c6`, impl `91f036c` (25 own and 26 app cases pass). One audit round;
  projection, stamping, other checks and CLI wiring remain unverified.

## First human checkpoint — five applied changes

Current-source Go build/test/race/vet and controller check/history/CI passed at
`91f036cbb33981a5fb72798da3edec9bdc90b077`; exact range starts at
`be35552b706aa3000ff8f11306d129ff1d701b85`. CI included doctor and all 86 reported
cases (including parent cases), with no diagnostics. See PROGRESS for logs and
limits. This is not fresh-checkout evidence or completion of any full acceptance
row. The manual NEW lifecycle, remaining surface, real drills/self-hosting and
delivery gates remain incomplete. Human review is required before change six.

- R04/R17 partial: APP-004–APP-005 wire modules and new-change through the NEW
  executable, including ancestor config discovery, graph/zone JSON and author
  error semantics. Archive `0006-project-commands.md`; phases `da5dba8` /
  `98e5852` / `5f09706`; 12 failures/39 passes before, all 51 app cases pass after.
  NEW smoke created changes/0001-answer.md in the manual slice fixture; stamping
  and lifecycle remain unverified.

- R04 partial: CON-006 extends ParseItems to the design-required new: handles
  without a parallel parser. Archive `0007-change-handles.md`; phases `a22fcd4` /
  `a194629` / `7580897`; one genuine failing case, then all 34 contracts cases pass
  with dependents/boundaries. Allocation, cross-change scope and rewriting remain
  for project/workflow evidence.

- R04/R05 partial: PRJ-004–PRJ-007 cover ordered queue parsing, projected edits/set
  metadata, per-change handle ownership, preceding-projection stamp bases and
  initial static diagnostics. Archive `0008-queue-authoring.md`; phases `6e3ca5f` /
  `776b4d0` / `2ef37df`; 24 failures/25 passes before, all 53 project cases pass after
  plus 51 app cases and boundaries. Full validation/views/lifecycle remain open.

- R04/R05/R17 partial: APP-006–APP-008 expose stamp/check and current/projected
  show through the NEW binary. Archive `0009-authoring-commands.md`; phases
  `5720dce` / `b2aaa14` / `de65319`; 27 failures/51 passes before, all 78 app cases
  pass after. Manual fixture stamp/check/projected and selected new:answer show
  exited 0 with no assigned IDs. Product lifecycle remains unbuilt/unverified.

## Shared-base approval gate after review 1

At clean product source `de65319d646fd0e39e9810415a181dd28911762c`, build/test/race/
vet, controller check and CI on explicit setup baseline `be35552...` passed. CI
includes doctor canaries/recovery, boundaries, history and all 173 reported cases,
with no diagnostics. This is current-source intermediate evidence, not final
acceptance or a fresh-checkout/drill/self-hosting claim. Controller unapproved
count is 5 against origin/main `6e341a4`; P07 requires human shared-base approval
before another change. Four changes were applied since the explicit review at
`3311f67`; no bound/ref was altered.

- R06/R13/R14 partial: GIT-001–GIT-005 supply canonical primary/linked repository
  metadata, local base selection, working/index/committed diffs, selected commits,
  blobs/CAS and bounded process groups. Archive `0010-manual-lifecycle-git.md`;
  evals `3e82adc` / `93ba9a6` / `5d00cac`, audits `98a799f` / `f2ddd0f` / `5911fab`,
  impl `92f999a`; 7 fail-first cases, then 8 permanent passes plus 2 temporary
  public smoke probes and dependent/boundary passes. Escaped process groups remain
  unsupported. Claims/lifecycle and full Git surface remain unverified.

- R07/R09/R10/R11 partial: CHK-001–CHK-005 cover public JUnit/results, hooks, zones and boundaries; archive `0011-lifecycle-checks.md`, phases `69933f7` / `eff6823` / `3c5b40b`. 25 failures before, 28 permanent cases and 4 public smoke probes pass after. Full dependents/boundaries pass. Product lifecycle and doctor integration remain unverified.

- R05–R09/R14/R17 partial: FLOW-001–FLOW-005 provide manual claims/start/verify/done/audit/apply/archive/provenance/context through public APIs. Archive `0012-manual-lifecycle.md`, evals `c5c02de` / `ca11056` / `aec7154`, findings `0b5aeef` / `25e4775`, final audit `8627014`, impl `563d2ee`: 28 initial failures, 32 final permanent cases and 3 public smoke probes pass; full dependents/boundaries pass. Sol reset/retry recorded in ignored evidence. NEW CLI lifecycle, remaining kinds/recovery/history and full acceptance remain unverified.

- Milestone 1 / R05–R11/R17 partial: APP-004, APP-009–APP-012 connect the manual
  lifecycle through NEW source `89b7028`; original phases `678f7bb` / `7db8c47` /
  `a9d195c` / `804cb84` / `89b7028`. Eight genuine initial failures, 86 final app
  cases and 2 public coverage probes pass. Real Go fixture TOY-001 at
  `/tmp/fireweed-stage2-manual-slice/project` completed native evals `f04e9f7`,
  separate audit `bfb84d0`, apply `1323a31`, with Z01 and P2-01 rejections,
  verify/no commit, frozen eval/facade, passing coverage, inspected provenance,
  clean tree and released claim. All workers/auditor were real Sol/high CLI calls.
  This verifies the manual milestone, not init/doctor/run or full acceptance.

- R03/R04/R05/R17 refinement: PRJ-007's Check contract is now split into
  PRJ-007–016 by archives 0030/0031. Existing validation assertions are retained
  and linked to their atomic IDs; cause-specific diagnostic migration remains
  pending. See [review 6 checkpoint](reports/review6-checkpoint.md) and PROGRESS
  for phase commits, audit corrections and verification evidence. This does
  not establish the remaining static-validation or CLI requirements.

- R02/R04 refinement: archives 0032–0034 split PRJ-001–003 into focused
  discovery, graph and authoring contracts, with PRJ-017–026 retaining their
  separated rules. Existing assertions and cases are unchanged. See
  [review 7 checkpoint](reports/review7-checkpoint.md) and PROGRESS for evidence
  and the P07 shared-base gate; PRJ-004–006 and full validation remain open.

- R04/R05 refinement: archives 0035–0039 complete the PRJ-004–006 split, with
  PRJ-027–039 retaining queue, projection and Stamp rules. Existing public
  assertions remain; 0038 strengthens complete-body checks after its audit.
  See [the passing checkpoint review](reports/checkpoint-0035-0039.md) and
  PROGRESS for phase evidence. Remaining module splits and D006 migration
  still precede Milestone 2; full acceptance remains unverified.

- R07/R10/R11 refinement: 0040–0044 split CHK-001–003 and retain CHK-005's
  suppression rules under CHK-039–043 before its final narrowing. Assertions,
  fixtures and implementation remain unchanged. See the
  [passing checkpoint review](reports/checkpoint-0040-0044.md) and PROGRESS
  for phase evidence; remaining splits and diagnostic migration are pending.

- R02/R03/R05/R06/R09/R11/R14/R27 refinement: 0045–0049 finish CHK-005,
  FLOW-004/005 and the identified contracts primitive bundles. Requirements
  and assertions are retained; existing lifecycle fixtures gain direct
  metadata/phase-base/shared-commit/provenance checks. See the
  [passing checkpoint review](reports/checkpoint-0045-0049.md) and PROGRESS
  for actual phase counts and approved execution-policy changes. Remaining
  gitstore/assets/app splits and D006 migration still precede the contracts drill.

- R02/R04/R05/R06/R17 refinement: 0050–0053 split gitstore/assets/app contracts
  with existing obligations and assertions retained. 0054 repairs PRJ-007's
  stale blanket C01 eval comparisons; exact cause migration remains pending.
  See the [passing checkpoint review](reports/checkpoint-0050-0054.md) and
  PROGRESS for phases and counts. This establishes no remaining acceptance gate.

- R04/R07/R11/R17 diagnostic refinement: 0055–0059 assign existing validation
  causes C04–C09 and boundary causes X01–X04, with exact cause assertions and
  accurate explanations. All phase audits passed; current-source build/race/vet
  and NEW JSON explanation smoke checks passed. See the
  [checkpoint review](reports/checkpoint-0055-0059.md) and PROGRESS for scope and
  phase counts. Remaining diagnostic migration and full acceptance are unverified.
