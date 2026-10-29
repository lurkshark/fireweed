# Rebuild plan

Revised 2026-10-28. Build the complete [v1 design](../../fireweed-design-v1.md)
through runnable CLI slices. Resume from [PROGRESS.md](PROGRESS.md), which owns
live state and authorizations. [SYSTEM.md](../../SYSTEM.md) owns guardrails;
[ACCEPTANCE.md](ACCEPTANCE.md) owns the full inventory and final gates.
[HISTORY.md](HISTORY.md) preserves completed work and historical evidence.

## Current route to a usable product

The manual lifecycle passed. Init/doctor and bounded agent invocation now exist;
NEW next completes one checked change; bounded run does not yet exist. The approved checkpoint corrections and scoped re-review passed and pushed.
Build the next CLI outcome,
using the existing lifecycle and invocation primitive.

| Slice | User-visible outcome through NEW | Exit evidence |
| --- | --- | --- |
| Manual lifecycle — passed | Author, check, start, fail evals, audit, implement, apply and archive one change. | Reuse the [manual lifecycle report](reports/manual-lifecycle.md); do not rebuild this milestone. |
| Setup — passed | Create a Go fixture with init and verify hooks/guidance with doctor. | Existing [setup smoke](reports/doctor-setup-cli.md) plus approved doctor corrections and scoped checkpoint PASS. Setup is not a reason to delay next after that correction gate. |
| One-change execution — passed | In a disposable Go fixture, next drives one eligible spec change through evals, separate audit, impl and archive using bounded deterministic fake agents. | NEW invocation succeeds; tests genuinely fail before implementation and pass afterward; assigned IDs, phase commits, SPEC/archive/provenance and claim release are inspectable. Library invocation tests alone do not close this slice. |
| Bounded queue | run processes eligible changes, honors --max, and a later invocation continues the remaining queue; an empty queue exits 0. | NEW runs a small fixture queue; inspect the exact applied/remaining changes, commits, exit codes and claims. Implement this before broad recovery or whole-module completion. |
| Recoverable execution | A failed verification feeds a bounded retry; exhausted work is saved and blocked, independent work can continue, and a saved attempt can resume. Time/usage stops preserve recoverable state. | Focused NEW fixture scenarios for retries/refs/state-only block commits, projected blocking, resume, agent-commit undo and exit 3. Split by observable scenario as needed; do not build a failure matrix. |
| Remaining lifecycle kinds and bounds | The same CLI path handles fixes, eval_fix, maintenance/regenerate and enforces audit/approval bounds for manual and unattended use. | Relevant NEW scenarios, including ordinary approval rejection without --allow-unmerged, plus linked-worktree claims and persisted bounds. Reuse existing assertions where valid. |
| Remaining surface and delivery | Complete design commands/options and all final gates. | ACCEPTANCE inventory and completion gates; one focused design review, scoped fixes and valid evidence. |

The first two execution slices are the immediate implementation focus after the
completed corrections. Later rows retain Milestone 2 and final obligations;
they are not prerequisites for wiring the first usable next/run path. No NEW
unattended real-agent execution before the bounded fake-agent run passes.

## Choosing and closing a change

Name the CLI scenario and missing behavior in the existing change Motivation.
Target one module per change and implement only what advances that scenario.
A dependency can land separately, but return to its caller as soon as the needed
interface works. In particular, wire app next after the minimum runner/workflow
path works; do not finish every runner recovery mechanism first.

Every slice ends with an actual NEW command and inspectable result. Reuse
fixtures and valid evidence; add no separate scenario registry, proof checklist
or generic equivalence harness. Required phase checks and audits still run.
A checkpoint may interrupt a slice; record the next observable step and resume
it afterward. Passing a library change is not evidence of CLI completion.

Batch related work within one module with no numeric cap on affected spec items.
Keep changes coherent and reviewable, preserve atomic specs, and use normal
changes/eval_fix for observed defects. Completed spec splitting and D006's
existing-path migration are historical. Keep explain accurate as justified
features emit codes; do not reopen global rewrites or prebuild unused diagnostics.
Workflow file organization and deferred drill gaps remain deferred in PROGRESS.

## Controller and product

| Role | Path | Use |
| --- | --- | --- |
| Pinned bootstrap controller | .fireweed/bin/fireweed-bootstrap | Rebuild claims, checks, commits, application and archives |
| NEW | bin/fireweed, built with go build -o bin/fireweed . | Product CLI fixtures and isolated self-hosting |

Use explicit paths. Worker commands named fireweed must resolve to the
controller; product fixture helpers must select NEW. FIREWEED_BIN identifies the
invoking executable so boundary hooks query its graph. Never substitute the
controller for product acceptance. Keep it in control of this checkout through
acceptance; NEW self-hosting uses an isolated fixture/clone.

Bootstrap help/playbooks, the initial generated recipe fixture and one CLI
lifecycle comparison are operational references. They add no product scope.
The design wins discrepancies, recorded briefly in DECISIONS. Never use bootstrap
source/tests, deleted code or rebuild-v1 as implementation inputs.

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

## Per-change loop

1. As author, read SYSTEM, projected specs, relevant design clauses and the authoring playbook. Use `new-change`; target one module with focused observable specs under Added using new-handles, or Modified/Removed/fixes metadata. Describe interface edits. `stamp` Modified/Removed, `check`, then commit the authored change and prerequisites.
2. Execute under the goal's local authority, preserving product approval from local/shared base and the shared-base P07 limit. Record P04 when emitted; do not claim a human merge occurred. `--allow-unmerged` permits planned local run/next work, never bypassing checkpoints or testing normal approval rejection.
3. Controller `start <slug>`; read `context <slug>` before each worker step. Write permanent black-box evals, the described interface and allowed stubs/manifests. Changed/fixes IDs need genuinely failing cases unless legitimately already_satisfied. Compile errors, unrelated failures, skips and empty reports do not suffice.
4. Iterate with `verify --quick`; full `verify`/`done` supply fail-first and zone evidence. `done` commits evals before implementation and always runs required full checks. Avoid extra full runs solely to duplicate reports.
5. Keep audit enabled. Use the audit playbook and a separate restricted Codex auditor with only the change, projected specs and evals. Findings are exact spec clauses with no assertions; exclude helper/style/convention critiques and unnamed edge cases. Re-audit only IDs with previous findings; a pass lists checked IDs. At most two findings round trips, persisted by the controller; further or untestable findings block with Questions and a human checkpoint. Same-context review cannot establish independence.
6. Implement only allowed files. Evals/interface/SPEC and protected/other-module files stay frozen; append-only permits additions only. Run full checks; only controller `done` applies, archives and commits. Correct frozen evals through eval_fix; correct behavior contracts through authored changes. Maintenance preserves behavior; behavioral fixes need regression evidence.
7. After application and claim release, commit one concise PROGRESS entry as author with assigned IDs, phase commits, results and next work. Link acceptance mappings to durable evals/archives/reports instead of duplicating logs or testcase lists.

## Verification, regeneration and delivery

Permanent evals use public APIs, temporary repositories, explicit Git identities,
local hooks and bounded fake agents. They need neither bootstrap, authenticated
agents nor network. Follow SYSTEM's cost, instrumentation and process limits.
Fake-agent runs prove orchestration; they do not prove real regeneration.

Run only the drill currently authorized in PROGRESS. Final acceptance retains
real dependency-first drills for every disposable module and the isolated NEW
self-hosting lifecycle, drill and CI. Reconstruction receives retained contracts,
evals and permitted dependency interfaces, without old implementation/history,
saved attempts, old binaries or bootstrap. Retain reports/refs, clean scratch
worktrees, and seek human agreement before any gap-driven follow-up.

The single controller/NEW lifecycle smoke compares exit codes, JSON envelope
shape and resulting commits/refs; record design-resolved differences. Avoid
per-edge-case comparisons. Controller CI proves rebuild process; NEW fixtures
prove product behavior.

At safe author boundaries, update PROGRESS's current slice and next step. Put
one concise entry per applied change in HISTORY with IDs, phase commits, actual
counts and limitations; link detailed reports and acceptance evidence. Preserve
controller identity, the original nonempty baseline and approval/checkpoint
counts. Never treat a push alone as a passing review.

Use one focused design/inventory review, then inspect fixes only for affected
requirements. Reuse valid checks/drills; repeat only invalidated or unresolved
evidence. Pass every ACCEPTANCE gate, account for warnings and stop verification
unless later changes invalidate it. Deliver the independent executable and clean
committed tree without pending/blocked rebuild changes or claims, with exact
source/controller revisions, history range and honest limitations. Push only
under explicit owner authorization; publication and tags remain outside scope.
