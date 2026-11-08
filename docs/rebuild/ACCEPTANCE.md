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

## Planned milestone — Workflow contracts and regeneration proof (WF1–WF4)

**Residual fixes/controller handoff and single retry complete; native verification failed.**
Normal audited [0181](../../changes/archive/0181-recorded-regenerate-preservation.md)
and [0182](../../changes/archive/0182-recovery-history-assertions.md) are applied.
The [corrected-source validation and verified pin handoff](../../.fireweed/drills/residual-fixes-controller-handoff.md)
passed before the one owner-authorized native workflow invocation. The
[current native report](../../.fireweed/drills/workflow.md) and
[complete review](../../.fireweed/drills/workflow-corrected-controller-retry-20261107T232508/review.md)
retain the actual failures, coverage, gap limitations, failed ref, full output and
cleanup. The focused corrected cases passed, but workflow/runner still fail.
No passing workflow proof or WF4 closure is claimed.

The owner's sequence approval is consumed. Stop for review with no active run or
claim. Further correction, diagnostic replay, drill, controller promotion,
reconstructed-code application or push requires new authorization. Original
79cf33a, prior evidence/refs, recoveries, bounds/counters and unrelated worktree
are preserved. Worker notes authorize no new change. Full-completion gates remain.

**Entry.** Follow SYSTEM and the pin/deadline in PROGRESS. Read AGENTS,
DECISIONS, the relevant unchanged design clauses and projected workflow
contracts. Inspect live status, queue, runs and shared claims. Preserve prior
evidence/refs, configured bounds/counters and unrelated worktrees. Stop at
the recorded deadline or a SYSTEM stop. Protected author commits occur
between active runs.

1. **WF1 — Focused workflow contract review.** Review workflow's Specs,
   Decisions, retained interfaces and public evals against its existing design
   obligations. Cover claim ownership and approved state, projected zones,
   fail-first/audit/impl transitions, verification error causes, full Done
   verification, atomic apply/archive/provenance, blocked recovery and release,
   context, eval_fix and regenerate behavior. Include the existing FLOW-055
   and FLOW-062 corrections. Reuse the completed checks review and passing
   independent proof. Record the conclusion in committed milestone evidence.
   A demonstrated defect, uncovered behavioral choice or proposed product/eval
   correction is an owner checkpoint before authoring. Any approved correction
   follows the normal audited lifecycle with frozen obligations and counters.

2. **WF2 — Identified clean candidate and approval handoff.** After WF1 closes,
   use a fresh clean checkout of the selected committed source and build a
   separate subject executable. Pass applicable normal/race Go suites and vet,
   then that candidate's check, doctor --no-agent, explicit-range check-commits
   and CI with hook doctor enabled. Preserve history baseline
   `be35552b706aa3000ff8f11306d129ff1d701b85`; resolve and record the exact
   nonempty range and candidate identity in committed evidence. Exercise actual
   candidate CLI paths affected by any approved correction. Reuse valid
   prefixes and unaffected evidence, while running required phase verification.
   Keep the subject separate from the pinned controller and confirm setup
   copies the selected invoker. Explain existing warnings without scope growth.
   Verify clean checkouts, cleanup, no live run/claim and preserved pin/bounds.
   Append the exact preparation evidence to the existing reviewed checks
   report, clearly labeled workflow preparation and preserving its native
   result. Update PROGRESS with the handoff state and use archived change
   evidence for any approved correction. **Stop before invoking drill and request owner
   approval for exactly the identified candidate.**

3. **WF3 — Native workflow drill after approval.** Only a separate owner
   authorization permits one native workflow invocation under the existing
   attempt/retry bounds. Use an isolated clone of the validated source with
   the unchanged pinned controller and rehearsed live isolation gate. Start
   the inspector alongside the controller. Before release, inspect actual
   workflow spec/Decisions/interfaces/evals and approved retained inputs, plus
   only declared contracts/gitstore/assets/project/checks public interfaces,
   including model interfaces. Keep dependency/dependent implementations,
   manifests and hooks in the separate verifier. Exclude original disposable
   code, old history, prior reports/reconstructions, binaries and bootstrap from
   worker inputs. Input Git exposes no old objects, remotes or alternates.
   Review every command, full script and incremental patch. Before a second
   configured attempt, review the first full trace/native feedback and repeat
   the live isolation check. Monitor the same invocation to its terminal result.
   Enforce workflow and all transitive dependents, currently runner and app,
   with retained inputs unchanged. A failed invocation or owner decision stops
   the milestone with native evidence retained. Further drills and gap fixes
   require separate authorization.

4. **WF4 — Reviewed proof and clean handoff.** Completion requires WF1/WF2
   and a passing independent workflow proof with no unresolved owner checkpoint.
   Review all native fields/cases, coverage, Decisions, gaps and independence.
   Retain the reviewed `.fireweed/drills/workflow.md` and native proof ref,
   keeping older reports accessible by commit and carrying forward the
   preparation evidence. Commit protected evidence with
   no active run/claim. Confirm scratch cleanup, clean checkouts and preserved
   pin/bounds/counters, prior proofs and unrelated worktrees. Update PROGRESS
   with results, exact candidate/range, scoped change commits or why none were
   needed, live unapproved count and any remaining checkpoint. Stop at WF4.

The requirement/command inventories and all full-completion gates remain.
Real regenerate changes, rebuilt-code import, further controller promotion, later
module drills, unrelated cleanup, publication and push remain outside this milestone.

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
2. **Process:** the recorded controller owns fail-first/audit/impl/archive commits (bootstrap before EC6, pinned NEW after it) for behavioral work; immutable evals/interfaces during impl; separate protected setup/author commits; no pending/blocked rebuild changes or active claims.
3. **Go:** with `GOCACHE=/tmp/fireweed-rebuild-cache`, a fresh checkout builds the new binary and passes `go test ./...`, `go test -race ./...`, `go vet ./...`. Permanent tests run with local fixtures/fake agents and without bootstrap/agent authentication. Cross-compilation is not platform runtime testing; report limits accurately.
4. **Controller checks:** the recorded controller runs `check`, `doctor`, `check-commits <setup-sha>..HEAD`, and `ci <setup-sha>..HEAD` with doctor enabled. Use the explicit baseline in WF2; retain the exact nonempty range and per-module coverage in committed evidence. Resolve errors. For warnings, make a scoped correction or record the reason/limitation; zero warnings is not a separate gate. An empty default range proves no history.
5. **Product checks:** NEW binary in fresh repositories performs init/doctor, author/stamp/check, documented base approval, start/failing-evals/audit/impl/done, coverage/history/PR/CI. Exercise fixes, eval_fix, maintenance, blocked recovery and regenerate. Normal unattended approval rejection is tested without --allow-unmerged. Pin subject paths so bootstrap cannot be used accidentally.
6. **Reference comparison (smoke):** one equivalent fixture lifecycle run with both binaries, comparing exit codes, JSON envelope shape and resulting commits/refs. Record intentional differences briefly. Reference bugs do not define acceptance, and per-edge-case equivalence is not required.
7. **Real regeneration:** successful controller-driven drills for every rebuilt disposable module, dependency-first. Real agents reconstruct from retained specs/Decisions/interfaces/evals and permitted dependency interfaces, without old implementation/history, saved attempts, old binaries or bootstrap. Own/transitive dependent evals pass without preserved-contract edits. Commit reviewed reports/retain refs and clean scratch worktrees. Repeat only drills invalidated by relevant changes or unresolved failures. Gaps require human agreement before authored fixes; agent suggestions alone authorize none. Missing real-agent access blocks this gate.
8. **Self-hosting:** in an isolated clone/fixture, the NEW binary governs a small real spec change and a real drill, with checked fail-first/audit/impl commits, archive/provenance, CI and scratch cleanup. Keep the original repository governed by its recorded controller; EC6 may already have promoted NEW. The real-drill portion remains required for full completion, separately from the earlier core self-hosting proof.
9. **Delivery:** README build/test/use instructions, locally validated distribution/installer checksum behavior, clean committed target tree, concise restartable handoff and independent product executable. No product runtime read/import/command requires bootstrap. Report exact checks/commits/reports and limitations. Push only on explicit owner instruction; do not publish or tag.

