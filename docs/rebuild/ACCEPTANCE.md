# Acceptance and traceability

Prepared 2026-10-21, reset 2026-10-26 when the rebuild restarted. All rows are **unverified**. R01–R29 are planning labels, not fireweed spec IDs. Expand them into atomic observable specs through changes; this inventory does not replace reading the full design.

The authoritative design snapshot is [fireweed-design-v1.md](../../fireweed-design-v1.md) in this repository; its origin and checksum are recorded in [PLAN.md](PLAN.md).

The first evidence milestone is the NEW binary's manual lifecycle slice in GOAL.md/PLAN.md; the rows below are eventual coverage, not a layer-by-layer implementation order. The current controller candidate is 0.1.0 from HANDOFF.md; old controller assertions do not establish readiness.

At safe author boundaries map each row to actual assigned spec IDs, permanent eval names, archived changes/phase commits and results. Record exceptions and design/reference discrepancies explicitly.

**Scope (revised 2026-10-26, see LESSONS.md).** Rows describe design behavior, not hardening beyond it. The SYSTEM.md scope guardrails decide how far each row goes: supported environment, process-group cancellation, conventions checked by review and one app PATH-shim test rather than per-spec locality evals, eval cost limits and audit scope. Do not expand a row with edge cases the design does not name. R22–R25 cover the bounded-effort revision of the design R26–R28 the usability revision and R29 the merge-strategy revision; they replace nothing above.

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

Exercise the new executable, beyond public function tests. Cover applicable documented options, JSON output, invalid input and exit semantics. JSON is required for design read commands. Additional old-controller syntax is not a product requirement.

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

1. **Traceability:** all design requirements mapped to assigned specs, black-box evals referencing them (or honest review requirements), and archived changes. All R rows verified. Do not evade acceptance through review/quarantine/already_satisfied/maintenance.
2. **Process:** bootstrap-managed fail-first/audit/impl/archive commits for behavioral work; immutable evals/interfaces during impl; separate protected setup/author commits; no pending/blocked rebuild changes or active claims.
3. **Go:** with `GOCACHE=/tmp/fireweed-rebuild-cache`, a fresh checkout builds the new binary and passes `go test ./...`, `go test -race ./...`, `go vet ./...`. Permanent tests run with local fixtures/fake agents and without bootstrap/agent authentication. Cross-compilation is not platform runtime testing; report limits accurately.
4. **Bootstrap checks:** controller `check`, `doctor`, `check-commits <setup-sha>..HEAD`, and `ci <setup-sha>..HEAD` with doctor enabled. Record the exact nonempty SHA range and per-module coverage. Resolve errors and account for warnings. An empty default range on main proves no history.
5. **Product checks:** NEW binary in fresh repositories performs init/doctor, author/stamp/check, documented base approval, start/failing-evals/audit/impl/done, coverage/history/PR/CI. Exercise fixes, eval_fix, maintenance, blocked recovery and regenerate. Normal unattended approval rejection is tested without --allow-unmerged. Pin subject paths so bootstrap cannot be used accidentally.
6. **Reference comparison (smoke):** one equivalent fixture lifecycle run with both binaries, comparing exit codes, JSON envelope shape and resulting commits/refs. Record intentional differences briefly. Reference bugs do not define acceptance, and per-edge-case equivalence is not required.
7. **Real regeneration:** successful bootstrap-driven drills for every rebuilt module declaring disposable code, using a working agent reconstructing from specs/Decisions/interfaces/evals without consulting old implementation/history/bootstrap source as an oracle. Run dependency-first; repeat only when relevant changes stale evidence. Own and transitive dependent evals pass without preserved-contract edits. Commit reviewed reports/retain refs. Fix agreed gaps through changes; never automatically approve all agent suggestions. Missing agent access blocks this gate; a fake agent cannot satisfy it.
8. **Self-hosting:** in an isolated clone/fixture, the NEW binary governs a small real spec change and a real drill, with checked fail-first/audit/impl commits, archive/provenance, CI and scratch cleanup. Keep the original repository governed by the pinned bootstrap throughout acceptance.
9. **Delivery:** README build/test/use instructions, locally validated distribution/installer checksum behavior, clean committed target tree, restartable evidence record and independent product executable. No product runtime read/import/command requires bootstrap. Report exact checks/commits/reports and limitations. Do not push, publish or tag.

## Evidence record shape

Create execution evidence when work runs; never prefill predicted successes.

```text
Requirement: R07 (specific clause)
Module / actual spec ID:
Change / archive path:
Permanent eval names and asserted edge cases:
Fail-first command, testcase IDs, and evidence location:
Evals / audit / implementation commit SHAs:
Final verification command, exit status, and evidence location:
Reference discrepancy / decision, if any:
Status: unverified | in progress | verified | blocked
```

Bulky logs may remain ignored; commit concise findings/counts and durable commit/test references. Independent audits, platform execution, real regeneration and self-hosting claims must match what actually ran.


## Partial evidence (rows remain unverified)

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
