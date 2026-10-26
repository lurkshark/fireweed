# Fireweed rebuild plan

Prepared 2026-10-21; aligned with the fresh bootstrap handoff on 2026-10-26 (D004). Status: preparation only. Reinspect state when the goal starts.

## Authority and scope

Build a new implementation of the complete v1 design with permanent behavioral specs and black-box evals. Use the bootstrap executable as development controller. The final product must work independently of the bootstrap checkout and executable. Generate the implementation from the design and newly authored contracts; do not use bootstrap source/tests, deleted code or rebuild-v1 as an implementation oracle.

Read the repository-local [fireweed-design-v1.md](../../fireweed-design-v1.md) as the product authority, then [ACCEPTANCE.md](ACCEPTANCE.md). Read the bootstrap HANDOFF.md and use the pinned binary's help, embedded playbooks, generated recipe fixture and one operational CLI comparison. Record design/controller discrepancies briefly with a design-derived eval; the design wins. Resolve unspecified details with a small documented contract rather than copying reference internals. Deferred/rejected ideas are not requirements. Ask only about consequential ambiguities the sources cannot resolve.

Keep this design copy protected from workers. It was an unchanged reference snapshot until 2026-10-26, when the owner revised it for bounded effort (audit rounds, eval_fix, unapproved-change limit, spec-shape and budget warnings, verify_dependents direct); see LESSONS.md. The fresh bootstrap 0.1.0 covers the revised design; pin and validate it during setup. Its bounds apply throughout development, including before the NEW binary implements them. Human review checkpoints in SYSTEM.md remain in force. Record clarifications and intentional departures separately in author decisions and change files so the original requirements remain auditable.

These documents contain planning labels and proposed architecture, not permanent spec IDs. Author specs through change files; only the controller applies module Specs and Decisions.

## Two executables

| Role | Suggested ignored location | Purpose |
| --- | --- | --- |
| Bootstrap controller | `.fireweed/bin/fireweed-bootstrap` | Pinned tool owning the rebuild's claims, phase checks, commits, spec application and archives |
| New product | `bin/fireweed`, built with `go build -o bin/fireweed .` | Subject of product acceptance and isolated self-hosting rehearsals |

Use explicit executable paths. Child workers whose playbooks invoke `fireweed` must resolve the controller through their environment; scratch worktrees need the same access. Keep an explicit separate product path in eval helpers. The Go boundary hook reads `FIREWEED_BIN` to query the graph of its invoking executable. Rebuild checks use the controller; product fixtures must receive the NEW binary. Never redirect product acceptance to the controller.

Keep the bootstrap in control through acceptance. Rehearse control by the new binary in an isolated fixture/clone; do not replace the controller midway through the evidence trail. The product must never delegate its commands to the bootstrap.

## Stage 0: setup as author

Before starting any worker change:

1. Inspect this repository's instructions/state, Git author/base refs, toolchains, the bootstrap handoff and available Codex commands. Preserve unrelated work. Prefer the current checkout; use `codex/` if a new branch is needed.
2. Pin the existing tested `/Users/joel/Workspace/fireweed-bootstrap/bin/fireweed` at `.fireweed/bin/fireweed-bootstrap`; verify 0.1.0 and embedded source SHA `7f925190878e38fc38f30d16bb1cf47aca098cf7` from HANDOFF.md. Do not rebuild from bootstrap source or use the old ignored controller. Record the pin and keep it unchanged throughout acceptance.
3. Read controller `help`, `version --json` and `playbook author-change` through explicit paths. In a disposable repository, use `init --recipe go --agent codex --audit-agent codex` to inspect its current recipe/canary contract. Adapt existing protected setup without reinitializing/overwriting this checkout. Its current >=0.4,<0.5 pin must be reconciled with controller 0.1.0. Configure worker and auditor argv for Codex `gpt-6-sol`, high reasoning effort, as specified in GOAL.md; validate both with bounded doctor probes.
4. Complete SYSTEM.md with product scope, module ownership, conventions, authority, executable roles and local author/worker policy. Preserve the managed AGENTS block; put controller-path guidance outside it.
5. Keep the Go module path `github.com/lurkshark/fireweed`; adapt recipe substitutions to that module. Retain existing minimal scaffolds and use `new-module` only when a new module is needed. Templates must also work for doctor's randomly named modules.
6. Establish build, JUnit eval reporting, boundary checking and scratch setup. Protect planning/reference documents, CI, root dependency/build configuration, entry-point glue and enforcement code. Keep shared writable paths narrow, initially the recipe's `go.sum`. Add root go.mod dependencies through separate author commits before changes that need them. Ignore build/cache/log outputs.
7. Add minimal root main.go forwarding to the app public entry point and only a compiling placeholder for that entry point. Public wrappers/types are interfaces; business behavior remains disposable. This setup allowance must not implement the product in protected or preserved files.
8. Commit setup and pass controller `check` and `doctor`, including all four canaries and recovery after their removal. Record the full setup commit SHA as the history baseline. Always use its explicit range for final verification; the default range on main can be empty.

Later protected/tooling/documentation edits are separate author commits, with justification and fresh canary checks if enforcement changed. Make them between applied changes, with no active worker claim. Do not smuggle author edits into worker commits or broaden permissions to pass checks.

## Restart (2026-10-26)

The first run (now the `rebuild-v1` branch) archived 51 changes but reached only 11 of 30 commands while spending most of its effort on hardening beyond the design. Local `main` was reset to the setup baseline 14cec73, and the revised design, guardrails and these documents were reapplied on top. [LESSONS.md](LESSONS.md) records what to keep from the first run. The SYSTEM.md scope guardrails override any broader reading of this plan.

The historical reset above is not a new reset instruction. Follow the vertical-slice milestones below. The following list describes responsibilities to draw from as each observable step needs them, not layers to finish in sequence:

1. contracts: spec and change grammar, IDs and hashes, zones and globs, configuration.
2. gitstore: discovery, base ref, diffs and clean tree, process execution with process groups and timeouts.
3. assets: version, playbooks, diagnostic catalog.
4. project: load and static check, queue projection and bases, ID allocation, finalize.
5. checks: JUnit parsing and coverage, eval and boundaries hooks.
6. workflow: claims, start, verify, done (apply, archive, provenance, blocking).
7. app: immediate CLI composition as operations become available; standalone commands early, then the manual lifecycle slice and init/doctor.
8. runner: next and run with attempts, limits and blocking; then ci, check-commits, pr-body and drill.
9. Cross-cutting R22–R29 coverage lands with the corresponding features; recipes/distribution and real drills/self-hosting complete acceptance.

Keep each change to at most five specs, and stop for human approval after every five applied changes. Record the new history baseline in PROGRESS.md; acceptance ranges start there, not at 14cec73.

## Proposed module graph

Start with the Go recipe's api.go, evals/ and internal/ layout. Validate the graph before bulk authoring. Its actual source of truth is SPEC.md frontmatter, queried through `modules --json`; this table is a proposal.

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

Check Go type ownership early: internal implementations cannot import their own public parent without cycles. Use shared public contracts or conversion at thin facades. Evals and dependent modules import public packages only. Avoid one giant mutable object coupling all modules. Adapt this proposed split through documented author decisions when a simpler acyclic design works better.

## Implementation milestones

Each milestone contains multiple single-module changes, in dependency order. Add the corresponding app command slice as soon as its dependencies work. Do not implement entire layers before their specs and evals.

| Stage | Changes to author | Exit evidence |
| --- | --- | --- |
| 1. Manual lifecycle vertical slice | Only the grammar/config/IDs/zones/Git/assets/project/JUnit/hooks/claims needed by one scenario; wire each operation through app promptly. Build standalone help/version/explain/playbook early, then author/stamp/check/start/verify/done/apply/archive. Use a tiny hand-authored product fixture until init is implemented. | NEW binary completes a real failing-before/audit/passing-after lifecycle; coverage, specs, archive, phase commits and provenance are inspectable; forbidden-zone and false fail-first cases rejected. |
| 2. Setup and bounded unattended execution | Complete init/doctor with Go canaries, next/run, all change kinds including eval_fix, projected blocking, claims/resume, attempts, limits, recovery and persistent audit/approval bounds. | Product-created fixture and canary recovery; bounded fake-agent success/failure/usage/time/recovery; linked-worktree contention and manual as well as unattended bounds. |
| 3. Complete surface | Finish grammar and views/options, budgets/warnings, history/CI/PR, merge strategies/reverts, abandon/release/gc, regeneration/drill/drift, recipes, migration/pins/AGENTS and local packaging. | Every design command/option and all R rows have meaningful evidence; the original lifecycle fixture keeps passing. |
| 4. Acceptance and delivery | Complete ACCEPTANCE gates, dependency-first real module drills and the isolated NEW self-hosting rehearsal. | Fresh-source Go checks, controller/product checks, reports/refs and cleanup, independent executable, clean tree and honest limitations. |

Derive new evals from the design and public interfaces in isolated fixtures. Use LESSONS.md only as a bounded historical checklist; do not consult or transplant old/reference tests. Root packaging/CI changes are author operations; observable product behavior still needs owned specs/evals.

## Per-change loop

1. As author, read SYSTEM, projected specs, relevant design clauses and the authoring playbook. Use `new-change`; target one module with focused observable specs under Added using new-handles, or Modified/Removed/fixes metadata. Describe interface edits. `stamp` Modified/Removed, `check`, then commit the authored change and prerequisites.
2. Execute under the goal's delegated local authority. A human merge to base remains the product's normal approval. Record P04 if emitted; a local-base change may start without it under the revised design. Do not claim a human merge happened or bypass the shared-base P07 limit. If using run/next, the goal explicitly permits `--allow-unmerged` for planned local rebuild changes. Do not use that flag to test normal approval enforcement.
3. Controller `start <slug>`, then `context <slug>`. Write permanent black-box evals, the described interface, and allowed new stubs/manifests. Each changed/fixes ID needs a genuinely failing testcase. Compile errors, unrelated failures, skips and empty reports do not suffice. already_satisfied is only for justified existing behavior.
4. Controller `verify`, inspect fail-first and zone evidence, then `done` to commit evals before implementation.
5. Keep audit enabled. Use the audit playbook and a separate restricted Codex auditor invocation, supplying only the change, projected specs and evals, without implementation or SYSTEM conventions. Record exact ID/clause findings and return through evals if needed. A passing audit lists IDs. Do not fabricate independent review; label any same-context limitation and resolve it before claiming independence.
6. Implement only allowed files. Do not edit evals/interface/SPEC, other modules/changes, protected paths or append-only contents. Controller `verify` and `done` must apply/archive/commit. Contract mistakes require a newly authored change, not weaker tests in impl. Maintenance is for behavior-preserving work; behavioral fixes require a failing regression eval.
7. After application and claim release, update acceptance/progress evidence and commit it separately as author. Record assigned IDs, test names, archive/phase commits, commands/results and the next change.

## Verification and recovery

Permanent tests use temporary Git repositories, explicit fixture identities, local hooks/fake agents, bounded timeouts and deterministic failure injection. They must run without bootstrap access, authenticated agents or network access. Fake-agent protocol tests do not prove real regeneration.

During real drills, the rebuilding agent must not consult bootstrap source, deleted implementation from Git history, saved attempts or old binaries as an implementation oracle. Give it the retained contracts and allowed dependency interfaces. Otherwise a passing drill proves copying works rather than that the specs/evals are sufficient. Keep the limited binary CLI comparison separate from this reconstruction exercise; bootstrap source/tests remain off limits throughout.

Compare bootstrap and product in separate equivalent fixtures: documented exit codes, stable JSON/diagnostics, file changes, commits, refs and phase transitions. Normalize only volatile dates/run IDs/temp paths/incidental commit hashes. Preserve order, spec IDs and meaningful content. Resolve mismatches against the design rather than requiring bug-for-bug equivalence. Controller CI verifies the rebuild process; separate tests must prove the product's own CI, doctor, lifecycle and runner.

Update docs/rebuild/PROGRESS.md during execution. At safe author boundaries record controller identity, baseline SHA, graph decisions, completed/next work, actual requirement mappings, results, discrepancies and blockers. On restart inspect Git status, controller status, changes and claims before editing. Resume the current phase; do not reset work or release another owner's claim to get a clean tree. Preserve useful attempts/drill refs.

Complete every gate in ACCEPTANCE.md. Commit reviewed reports and final evidence as author. Successful completion leaves no pending/blocked rebuild changes or active claims. Report executable path, checked SHA range, evidence and limitations. Publishing is outside scope.

## Historical reference inspection record — superseded by D004

The entries below describe the retired controller and the original preparation. They are not current pin, agent-access or readiness evidence. The current candidate is the 0.1.0 binary and source revision in HANDOFF.md; installation and fresh setup checks remain pending.

- Bootstrap clean source commit: `583b3c53740cf2341e0b0a79afee5b7f646ec1d4`.
- Local design snapshot: `fireweed-design-v1.md` at the repository root, copied byte-for-byte from `/Users/joel/Workspace/fireweed-bootstrap/fireweed-design-v1.md` at that commit on 2026-10-21.
- Design SHA-256: `1c17d0ab9496b1b15ad5ca18eed4f4ca623d01cff596bc557bf4c97400d00789`. Revised 2026-10-26: `a1fe8a2a2335c04443d4e7d3e6aa3a138e221cd5bd8302c6c66bfd947761d919`; usability revision 2026-10-26: `5d7cfcc9e3987fb594d16cce94f5ffa057c583f17f985114fe1de946aef5cf9d`; merge-strategy revision 2026-10-26: `52623b89d8184680f535519e82ceda8a49490b3d4dec000347cbd2cde429b21a`.
- Target initial commit: `d543462579d3fe182460c07476846476d38c6dc5` on main. Choose a later setup commit as the execution baseline.
- No fireweed binary was on PATH or at the checked common install locations. A temporary build reported `0.4.0`, using `go1.27.2 darwin/arm64` and a temporary writable GOCACHE.
- In a separate temporary repository, init with the Go recipe, agent none and audit-agent codex, followed by committed setup and doctor, passed boundaries/interface/failing-eval/missing-id with no diagnostics. This verifies plumbing, not the full bootstrap test suite or an actual agent/auditor execution.
- The temporary executable is not a durable input: rebuild/pin when the goal starts. Neither workspace repository was initialized or implemented by the probe.
