# Fireweed rebuild plan

Prepared 2026-10-21; aligned with D004 on 2026-10-26 and tightened 2026-10-27. This is the execution plan, not a current status report. Resume from [PROGRESS.md](PROGRESS.md), honoring pending human checkpoints and reusing completed setup and milestone evidence.

## Authority and scope

Build a new implementation of the complete v1 design with permanent behavioral specs and black-box evals. Use the bootstrap executable as development controller. The final product must work independently of the bootstrap checkout and executable. Generate the implementation from the design and newly authored contracts; do not use bootstrap source/tests, deleted code or rebuild-v1 as an implementation oracle.

Read the local [design](../../fireweed-design-v1.md) initially, then consult relevant clauses and [ACCEPTANCE.md](ACCEPTANCE.md) as each slice needs them. Bootstrap HANDOFF.md, binary help/playbooks and a generated recipe fixture support initial setup; one CLI smoke comparison supports acceptance. Record relevant discrepancies briefly; add evals only for uncovered design behavior. The design wins. Resolve unspecified details with a small contract. Deferred ideas and historical observations add no requirements.

Keep the design protected from workers. Bootstrap 0.1.0 covers the owner's October 26 revisions; its bounds and SYSTEM's checkpoints apply before NEW implements them. Record clarifications and intentional departures in author decisions/change files. SYSTEM's effort, audit and evidence rules govern this plan.

These documents contain planning labels and proposed architecture, not permanent spec IDs. Author specs through change files; only the controller applies module Specs and Decisions.

## Two executables

| Role | Suggested ignored location | Purpose |
| --- | --- | --- |
| Bootstrap controller | `.fireweed/bin/fireweed-bootstrap` | Pinned tool owning the rebuild's claims, phase checks, commits, spec application and archives |
| New product | `bin/fireweed`, built with `go build -o bin/fireweed .` | Subject of product acceptance and isolated self-hosting rehearsals |

Use explicit executable paths. Child workers whose playbooks invoke `fireweed` must resolve the controller through their environment; scratch worktrees need the same access. Keep an explicit separate product path in eval helpers. The Go boundary hook reads `FIREWEED_BIN` to query the graph of its invoking executable. Rebuild checks use the controller; product fixtures must receive the NEW binary. Never redirect product acceptance to the controller.

Keep the bootstrap in control through acceptance. Rehearse control by the new binary in an isolated fixture/clone; do not replace the controller midway through the evidence trail. The product must never delegate its commands to the bootstrap.

## Stage 0: setup as author

Perform this setup once. On resume, inspect status/claims and reuse the setup, controller pin and baseline recorded in PROGRESS.md. Do not reinitialize scaffolds or repeat passed probes unless relevant changes or unresolved failures invalidate them.

1. Inspect this repository's instructions/state, Git author/base refs, toolchains, the bootstrap handoff and available Codex commands. Preserve unrelated work. Prefer the current checkout; use `codex/` if a new branch is needed.
2. Pin the existing tested `/Users/joel/Workspace/fireweed-bootstrap/bin/fireweed` at `.fireweed/bin/fireweed-bootstrap`; verify 0.1.0 and embedded source SHA `7f925190878e38fc38f30d16bb1cf47aca098cf7` from HANDOFF.md. Do not rebuild from bootstrap source or use the old ignored controller. Record the pin and keep it unchanged throughout acceptance.
3. Read controller `help`, `version --json` and `playbook author-change` through explicit paths. Inspect `init --recipe go --agent codex --audit-agent codex` in a disposable repository; reconcile incompatible configuration/hooks without init over this checkout. Configure workers and auditors for Codex `gpt-6-sol`, high effort, using GOAL.md's argv; validate both with bounded doctor probes. Stop on unavailable access rather than switching models or retrying indefinitely.
4. Complete SYSTEM.md with product scope, module ownership, conventions, authority, executable roles and local author/worker policy. Preserve the managed AGENTS block; put controller-path guidance outside it.
5. Keep the Go module path `github.com/lurkshark/fireweed`; adapt recipe substitutions to that module. Retain existing minimal scaffolds and use `new-module` only when a new module is needed. Templates must also work for doctor's randomly named modules.
6. Establish build, JUnit eval reporting, boundary checking and scratch setup. Protect planning/reference documents, CI, root dependency/build configuration, entry-point glue and enforcement code. Keep shared writable paths narrow, initially the recipe's `go.sum`. Add root go.mod dependencies through separate author commits before changes that need them. Ignore build/cache/log outputs.
7. Add minimal root main.go forwarding to the app public entry point and only a compiling placeholder for that entry point. Public wrappers/types are interfaces; business behavior remains disposable. This setup allowance must not implement the product in protected or preserved files.
8. Commit setup and pass controller `check` and `doctor`, including all four canaries and recovery after their removal. Record the full setup commit SHA as the history baseline. Always use its explicit range for final verification; the default range on main can be empty.

Later protected/tooling/documentation edits are separate author commits between active changes, with no worker claim. Run affected canaries when enforcement changes; documentation alone does not invalidate canary evidence. Do not broaden permissions to pass checks.

## Restart (2026-10-26)

The first run (now the `rebuild-v1` branch) archived 51 changes but reached only 11 of 30 commands while spending most of its effort on hardening beyond the design. Local `main` was reset to the setup baseline 14cec73, and the revised design, guardrails and these documents were reapplied on top. [LESSONS.md](LESSONS.md) records what to keep from the first run. The SYSTEM.md scope guardrails override any broader reading of this plan.

This historical reset is not an instruction. Follow the milestones below, drawing from the module responsibilities as each observable step needs them. Cross-cutting R22–R29 coverage lands with its features, not as a separate hardening pass.

Batch related spec-only splits and eval retargeting within one module, with at most 30 affected spec items per change. Preserve every behavioral obligation and assertion, keep resulting specs atomic, and leave product behavior, public interfaces, controller verification and independent audits unchanged. Other changes remain limited to five focused specs. Follow SYSTEM's checkpoints and the owner-authorized recurring subagent checkpoint review policy in PROGRESS.md. Acceptance ranges start at the recorded stage-2 baseline, not the historical reset commit; resuming does not create a new baseline.

## Proposed module graph

Use the Go recipe's api.go, evals/ and internal/ layout. SPEC.md frontmatter, queried through `modules --json`, is the graph authority; this table is a proposal, not a reason to reorganize working modules.

| Module / prefix | Responsibility | Proposed direct dependencies |
| --- | --- | --- |
| contracts / CON | Public data, document grammar, IDs/hashes, glob/zone primitives | none |
| gitstore / GIT | Git/ref/worktree operations, safe paths, diff/process adapters | contracts |
| assets / AST | Embedded recipes/templates/canaries/playbooks and version metadata | contracts |
| project / PRJ | Discovery/defaults/graph, projection/validation, authoring helpers | contracts, gitstore, assets |
| checks / CHK | JUnit/coverage, hook/boundary/history checks, doctor | contracts, gitstore, assets, project |
| workflow / FLOW | Lifecycle/claims, phase enforcement, apply/archive/provenance, context | contracts, gitstore, assets, project, checks |
| runner / RUN | Agents, queue/retries/limits/recovery, drills | contracts, gitstore, assets, project, checks, workflow |
| app / APP | CLI parsing/output/dispatch, init/migration/maintenance composition, end-to-end contract | contracts, gitstore, assets, project, checks, workflow, runner |

Internal implementations cannot import their own public parent without cycles. Use shared public contracts or thin-facade conversions; evals and dependents import public packages only. Keep the existing graph unless a concrete feature needs a documented acyclic adjustment. Architectural preferences alone do not justify cleanup changes.

## Implementation milestones

Each milestone contains single-module changes ordered by the next observable step. Wire app as soon as dependencies work. Reuse completed exits; preserve the authorized next work in PROGRESS.md. No NEW unattended real-agent execution before the Milestone 2 fake-agent run passes.

| Stage | Changes to author | Exit evidence |
| --- | --- | --- |
| 1. Manual lifecycle vertical slice | Only the grammar/config/IDs/zones/Git/assets/project/JUnit/hooks/claims needed by one scenario; wire each operation through app promptly. Build standalone help/version/explain/playbook early, then author/stamp/check/start/verify/done/apply/archive. Use a tiny hand-authored product fixture until init is implemented. | NEW binary completes a real failing-before/audit/passing-after lifecycle; coverage, specs, archive, phase commits and provenance are inspectable; forbidden-zone and false fail-first cases rejected. |
| 2. Setup and bounded unattended execution | Complete init/doctor with Go canaries, next/run, all change kinds including eval_fix, projected blocking, claims/resume, attempts, limits, recovery and persistent audit/approval bounds. | Product-created fixture and canary recovery; bounded fake-agent success/failure/usage/time/recovery; linked-worktree contention and manual as well as unattended bounds. |
| 3. Complete surface | Finish grammar and views/options, budgets/warnings, history/CI/PR, merge strategies/reverts, abandon/release/gc, regeneration/drill/drift, recipes, migration/pins/AGENTS and local packaging. | Every design command/option and all R rows have meaningful evidence; the original lifecycle fixture keeps passing. |
| 4. Acceptance and delivery | Complete ACCEPTANCE gates, dependency-first real module drills and the isolated NEW self-hosting rehearsal. | Fresh-source Go checks, controller/product checks, reports/refs and cleanup, independent executable, clean tree and honest limitations. |

Derive new evals from the design and public interfaces in isolated fixtures. Use LESSONS.md only as a bounded historical checklist; do not consult or transplant old/reference tests. Root packaging/CI changes are author operations; observable product behavior still needs owned specs/evals.

## Per-change loop

1. As author, read SYSTEM, projected specs, relevant design clauses and the authoring playbook. Use `new-change`; target one module with focused observable specs under Added using new-handles, or Modified/Removed/fixes metadata. Describe interface edits. `stamp` Modified/Removed, `check`, then commit the authored change and prerequisites.
2. Execute under the goal's local authority, preserving product approval from local/shared base and the shared-base P07 limit. Record P04 when emitted; do not claim a human merge occurred. `--allow-unmerged` permits planned local run/next work, never bypassing checkpoints or testing normal approval rejection.
3. Controller `start <slug>`; read `context <slug>` before each worker step. Write permanent black-box evals, the described interface and allowed stubs/manifests. Changed/fixes IDs need genuinely failing cases unless legitimately already_satisfied. Compile errors, unrelated failures, skips and empty reports do not suffice.
4. Iterate with `verify --quick`; full `verify`/`done` supply fail-first and zone evidence. `done` commits evals before implementation and always runs required full checks. Avoid extra full runs solely to duplicate reports.
5. Keep audit enabled. Use the audit playbook and a separate restricted Codex auditor with only the change, projected specs and evals. Findings are exact spec clauses with no assertions; exclude helper/style/convention critiques and unnamed edge cases. Re-audit only IDs with previous findings; a pass lists checked IDs. At most two findings round trips, persisted by the controller; further or untestable findings block with Questions and a human checkpoint. Same-context review cannot establish independence.
6. Implement only allowed files. Evals/interface/SPEC and protected/other-module files stay frozen; append-only permits additions only. Run full checks; only controller `done` applies, archives and commits. Correct frozen evals through eval_fix; correct behavior contracts through authored changes. Maintenance preserves behavior; behavioral fixes need regression evidence.
7. After application and claim release, commit one concise PROGRESS entry as author with assigned IDs, phase commits, results and next work. Link acceptance mappings to durable evals/archives/reports instead of duplicating logs or testcase lists.

## Verification and recovery

Permanent tests use temporary Git repositories, explicit fixture identities, local hooks/fake agents, bounded timeouts and deterministic failure injection. They must run without bootstrap access, authenticated agents or network access. Fake-agent protocol tests do not prove real regeneration.

Follow the currently authorized drill in PROGRESS.md. For final acceptance, run real-agent drills dependency-first for all disposable modules. Give reconstruction only retained specs/Decisions/interfaces/evals and permitted dependency interfaces, without old implementation/history, saved attempts, old binaries or bootstrap. Retain reports/refs and clean scratch worktrees. Gaps require human agreement before follow-up changes; a report alone authorizes none.

Run one equivalent lifecycle fixture with controller and NEW; compare exit codes, JSON envelope shape and resulting commits/refs. Record design-resolved differences briefly. Do not build a general equivalence harness or compare edge cases. Controller CI checks the rebuild process; NEW fixtures establish product behavior.

Update docs/rebuild/PROGRESS.md during execution. At safe author boundaries record controller identity, baseline SHA, graph decisions, completed/next work, actual requirement mappings, results, discrepancies and blockers. On restart inspect Git status, controller status, changes and claims before editing. Resume the current phase; do not reset work or release another owner's claim to get a clean tree. Preserve useful attempts/drill refs.

Follow SYSTEM's correction limit and checkpoints. Reuse valid checks/drills; required controller phase checks still run. Use one design/inventory review, inspect fixes for affected requirements, and account for warnings with scoped fixes or explanations.

Pass every existing ACCEPTANCE gate, then stop verification unless later changes invalidate evidence. Commit reviewed reports/final evidence as author. Completion leaves no pending/blocked rebuild changes or claims. Report executable path, checked SHA range, evidence and limitations. Publishing is outside scope.

## Historical reference inspection record — superseded by D004

These entries describe the retired controller and original preparation. They supply historical provenance only, not startup instructions or current readiness. PROGRESS.md records the installed stage-2 controller, completed setup and baseline.

- Bootstrap clean source commit: `583b3c53740cf2341e0b0a79afee5b7f646ec1d4`.
- Local design snapshot: `fireweed-design-v1.md` at the repository root, copied byte-for-byte from `/Users/joel/Workspace/fireweed-bootstrap/fireweed-design-v1.md` at that commit on 2026-10-21.
- Design SHA-256: `1c17d0ab9496b1b15ad5ca18eed4f4ca623d01cff596bc557bf4c97400d00789`. Revised 2026-10-26: `a1fe8a2a2335c04443d4e7d3e6aa3a138e221cd5bd8302c6c66bfd947761d919`; usability revision 2026-10-26: `5d7cfcc9e3987fb594d16cce94f5ffa057c583f17f985114fe1de946aef5cf9d`; merge-strategy revision 2026-10-26: `52623b89d8184680f535519e82ceda8a49490b3d4dec000347cbd2cde429b21a`.
- Target initial commit: `d543462579d3fe182460c07476846476d38c6dc5` on main; superseded by the stage-2 baseline in PROGRESS.md.
- The original temporary 0.4.0 build used `go1.27.2 darwin/arm64`. A disposable Go-recipe init/doctor probe passed four canaries without real agent/auditor execution. Neither workspace was initialized by it. The temporary binary is not a durable input or a reason to rebuild/repin on resume.
