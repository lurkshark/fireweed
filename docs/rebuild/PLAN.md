# Fireweed rebuild plan

Prepared 2026-10-21. Status: planning only. Reinspect current state when the goal starts.

## Authority and scope

Build a new implementation of the complete v1 design with permanent behavioral specs and black-box evals. Use the bootstrap executable as development controller. The final product must work independently of the bootstrap checkout and executable. This is a rewrite with reference access, not a strict clean-room exercise.

Read the repository-local [fireweed-design-v1.md](../../fireweed-design-v1.md) as the product authority, then [ACCEPTANCE.md](ACCEPTANCE.md). Consult the bootstrap README, source, tests, embedded recipes and playbooks for operational examples. Where they disagree, record source locations, the discrepancy, the decision and a design-derived test. The explicit design wins; consistent reference behavior can fill unspecified details. Deferred/rejected ideas are not requirements. Ask only about consequential ambiguities the sources cannot resolve.

Keep this design copy as an unchanged reference snapshot and protect it from workers during setup. Record clarifications and intentional departures separately in author decisions and change files so the original requirements remain auditable.

These documents contain planning labels and proposed architecture, not permanent spec IDs. Author specs through change files; only the controller applies module Specs and Decisions.

## Two executables

| Role | Suggested ignored location | Purpose |
| --- | --- | --- |
| Bootstrap controller | `.fireweed/bin/fireweed-bootstrap` | Pinned tool owning the rebuild's claims, phase checks, commits, spec application and archives |
| New product | `bin/fireweed`, built with `go build -o bin/fireweed .` | Subject of product acceptance and isolated self-hosting rehearsals |

Use explicit executable paths. Child workers whose playbooks invoke `fireweed` must resolve the controller through their environment; scratch worktrees need the same access. Keep an explicit separate product path in eval helpers. The Go boundary hook reads `FIREWEED_BIN` to query the graph; the bootstrap supplies that variable when invoking hooks. Do not redirect this accidentally to the product.

Keep the bootstrap in control through acceptance. Rehearse control by the new binary in an isolated fixture/clone; do not replace the controller midway through the evidence trail. The product must never delegate its commands to the bootstrap.

## Stage 0: setup as author

Before starting any worker change:

1. Inspect repository instructions, both working trees, Git author/base refs, toolchains and available agent commands. Preserve unrelated work. Prefer the current checkout; use `codex/` if a new branch is needed.
2. Locate a compiled bootstrap or build from the reference checkout into the dedicated ignored controller path. Verify/pin source commit and record reported version, build command/toolchain and checksum. If the source differs from the inspection record below, review and record the difference. Do not overwrite either source tree or the future product binary.
3. Read controller `help`, `version --json` and `playbook author-change`. Run `init --recipe go --agent <available-agent> --audit-agent <available-auditor>`, retaining app. `--agent none` supports manual lifecycle driving, but unattended execution and real drills require a working configured agent. Do not assume a different vendor's CLI is installed.
4. Complete SYSTEM.md with product scope, module ownership, conventions, authority, executable roles and local author/worker policy. Preserve the managed AGENTS block; put controller-path guidance outside it.
5. Choose the Go module path, normally `github.com/lurkshark/fireweed`; update generated canaries' `example.com/project` imports accordingly. Create minimal module scaffolds with `new-module`. Templates must also work for doctor's randomly named modules.
6. Establish build, JUnit eval reporting, boundary checking and scratch setup. Protect planning/reference documents, CI, root dependency/build configuration, entry-point glue and enforcement code. Keep shared writable paths narrow, initially the recipe's `go.sum`. Add root go.mod dependencies through separate author commits before changes that need them. Ignore build/cache/log outputs.
7. Add minimal root main.go forwarding to the app public entry point and only a compiling placeholder for that entry point. Public wrappers/types are interfaces; business behavior remains disposable. This setup allowance must not implement the product in protected or preserved files.
8. Commit setup and pass controller `check` and `doctor`, including all four canaries and recovery after their removal. Record the full setup commit SHA as the history baseline. Always use its explicit range for final verification; the default range on main can be empty.

Later protected/tooling/documentation edits are separate author commits, with justification and fresh canary checks if enforcement changed. Make them between applied changes, with no active worker claim. Do not smuggle author edits into worker commits or broaden permissions to pass checks.

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
| 1. Contracts and inspection | Grammar/config/version defaults, IDs/hashes/tags, zones/glob overlap, Git/path primitives, assets, discovery/projection/validation, basic read commands and JSON/exits | Valid/invalid fixture projects and diagnostic assertions against the new binary |
| 2. Checked manual lifecycle | JUnit/fail-first, hooks/boundaries, authoring/context, start/claims, eval/audit/impl transitions, apply/archive/Decisions/provenance | Real fresh-project change with failing-before/passing-after evidence and rejected zone violations |
| 3. Queue and recovery | Approval content, dependency blocking, retries/limits, attempts/resume/release/gc, agent-created commits, process cancellation | Fake-agent success/failure/limits/recovery; two worktrees contending on one claim |
| 4. Regeneration and checks | Regenerate/drills/reports/refs/drift, history H01–H04, merges/squashes/reverts, CI/PR/history queries, doctor/overrides | Real canaries, own/dependent evals after regeneration, adversarial Git histories |
| 5. Complete product | Every recipe/local recipe/none, migration/pins, AGENTS refresh, distribution/installer/checksums, usable README, full comparison | All acceptance rows traced; independent build and final gates |

Consult reference tests for edge cases, but write new evals against public interfaces and isolated fixtures. Do not transplant private-function tests. Root packaging/CI changes are author operations; observable product behavior still needs owned specs/evals.

## Per-change loop

1. As author, read SYSTEM, projected specs, relevant design clauses and the authoring playbook. Use `new-change`; target one module with focused observable specs under Added using new-handles, or Modified/Removed/fixes metadata. Describe interface edits. `stamp` Modified/Removed, `check`, then commit the authored change and prerequisites.
2. Execute under the goal's delegated local authority. A human merge to base remains the product's normal approval. Record manual start's unmerged warning; do not claim a human merge happened. If using run/next, the goal explicitly permits `--allow-unmerged` for planned local rebuild changes. Do not use that flag to test normal approval enforcement.
3. Controller `start <slug>`, then `context <slug>`. Write permanent black-box evals, the described interface, and allowed new stubs/manifests. Each changed/fixes ID needs a genuinely failing testcase. Compile errors, unrelated failures, skips and empty reports do not suffice. already_satisfied is only for justified existing behavior.
4. Controller `verify`, inspect fail-first and zone evidence, then `done` to commit evals before implementation.
5. Keep audit enabled. Use the audit playbook and a separate auditor context where available, supplying only approved change/spec/interfaces/evals and permitted context. Record exact ID/clause findings and return through evals if needed. A passing audit lists IDs. Do not fabricate independent review; label any same-context limitation and resolve it before claiming independence.
6. Implement only allowed files. Do not edit evals/interface/SPEC, other modules/changes, protected paths or append-only contents. Controller `verify` and `done` must apply/archive/commit. Contract mistakes require a newly authored change, not weaker tests in impl. Maintenance is for behavior-preserving work; behavioral fixes require a failing regression eval.
7. After application and claim release, update acceptance/progress evidence and commit it separately as author. Record assigned IDs, test names, archive/phase commits, commands/results and the next change.

## Verification and recovery

Permanent tests use temporary Git repositories, explicit fixture identities, local hooks/fake agents, bounded timeouts and deterministic failure injection. They must run without bootstrap access, authenticated agents or network access. Fake-agent protocol tests do not prove real regeneration.

During real drills, the rebuilding agent must not consult bootstrap source, deleted implementation from Git history, saved attempts or old binaries as an implementation oracle. Give it the retained contracts and allowed dependency interfaces. Otherwise a passing drill proves copying works rather than that the specs/evals are sufficient. Keep ordinary reference comparison separate from this reconstruction exercise.

Compare bootstrap and product in separate equivalent fixtures: documented exit codes, stable JSON/diagnostics, file changes, commits, refs and phase transitions. Normalize only volatile dates/run IDs/temp paths/incidental commit hashes. Preserve order, spec IDs and meaningful content. Resolve mismatches against the design rather than requiring bug-for-bug equivalence. Controller CI verifies the rebuild process; separate tests must prove the product's own CI, doctor, lifecycle and runner.

Create docs/rebuild/PROGRESS.md during execution. At safe author boundaries record controller identity, baseline SHA, graph decisions, completed/next work, actual requirement mappings, results, discrepancies and blockers. On restart inspect Git status, controller status, changes and claims before editing. Resume the current phase; do not reset work or release another owner's claim to get a clean tree. Preserve useful attempts/drill refs.

Complete every gate in ACCEPTANCE.md. Commit reviewed reports and final evidence as author. Successful completion leaves no pending/blocked rebuild changes or active claims. Report executable path, checked SHA range, evidence and limitations. Publishing is outside scope.

## Reference inspection record

- Bootstrap clean source commit: `583b3c53740cf2341e0b0a79afee5b7f646ec1d4`.
- Local design snapshot: `fireweed-design-v1.md` at the repository root, copied byte-for-byte from `/Users/joel/Workspace/fireweed-bootstrap/fireweed-design-v1.md` at that commit on 2026-10-21.
- Design SHA-256: `1c17d0ab9496b1b15ad5ca18eed4f4ca623d01cff596bc557bf4c97400d00789`.
- Target initial commit: `d543462579d3fe182460c07476846476d38c6dc5` on main. Choose a later setup commit as the execution baseline.
- No fireweed binary was on PATH or at the checked common install locations. A temporary build reported `0.4.0`, using `go1.27.2 darwin/arm64` and a temporary writable GOCACHE.
- In a separate temporary repository, init with the Go recipe, agent none and audit-agent codex, followed by committed setup and doctor, passed boundaries/interface/failing-eval/missing-id with no diagnostics. This verifies plumbing, not the full bootstrap test suite or an actual agent/auditor execution.
- The temporary executable is not a durable input: rebuild/pin when the goal starts. Neither workspace repository was initialized or implemented by the probe.
