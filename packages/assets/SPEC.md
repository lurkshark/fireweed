---
depends_on:
    - contracts
id_prefix: AST
interface:
    - api.go
    - model/**
name: assets
---

# assets

## Specs

### AST-001
Version MUST return 0.1.0.

### AST-002
Playbook MUST return nonempty instructions for author-change, evals, audit, impl and regenerate, and reject unknown names.

### AST-003
Explain MUST describe the cause and corrective action for C01 (invalid configuration), C04 (non-normative text), C05 (foreign or duplicate spec/edit identities), C06 (invalid tags), C07 (missing Motivation or invalid change module target), C08 (missing/stale edit bases or absent targets), C09 (undeclared handles), X01 (boundary violations), X02 (missing/incomplete boundary hooks), X03 (reasoned boundary opt-out), X04 (unnecessary boundary enforcement), U01 (invalid arguments/options or unsupported commands), U02 (unknown or ambiguous requested names) and P2-01 (missing genuine fail-first evidence). Explain MUST reject unknown codes.

### AST-004
BuiltinRecipe MUST return the embedded Go recipe for go and reject unknown names.

### AST-005
The Go eval hook MUST run the requested module's Go tests and write JUnit testcase names, classnames, pass/fail/skip and elapsed seconds while preserving test failure exit status.

### AST-006
The Go boundaries hook MUST query FIREWEED_BIN modules --json and use its module dirs, depends_on and zone map with the workspace module path.

### AST-007
The Go bundle MUST provide module SPEC/api.go templates and version-1 baseline, undeclared, internal_eval, failing_eval and missing_id canary maps.

### AST-008
Help MUST identify fireweed and list every command in the design's CLI table, including version, explain and playbook.

### AST-009
The audit playbook MUST restrict findings to exact spec clauses without assertions.

### AST-010
The audit playbook MUST restrict re-audits to IDs with prior findings.

### AST-011
The author playbook MUST require a user-visible or design basis for Motivation.

### AST-012
The author playbook MUST recommend narrowing costly specs with already_satisfied.

### AST-013
The Go bundle MUST contain instructions, api.go/evals/**/internal/** defaults, eval/boundaries/setup commands, go.sum worker_writable and project-relative recipe, hook, template and canary files.

### AST-014
Go recipe instructions MUST explain Go module and workspace setup.

### AST-015
Go recipe instructions MUST explain deriving the graph from FIREWEED_BIN modules --json.

### AST-016
The Go eval hook MUST retain started-but-unfinished tests and package or compile failures as visible failures rather than passing or empty evidence.

### AST-017
The Go boundaries hook MUST report module-prefixed violations for imports of undeclared modules.

### AST-018
The Go boundaries hook MUST report module-prefixed violations for eval imports reaching beyond retained interface packages, including own internal implementation.

### AST-019
Clean module fixtures MUST pass the Go boundaries hook.

### AST-020
The Go bundle templates and canary snippets MUST be substitutable with name, prefix, package, go_module, module_dir and other_dir.

### AST-021
Rendered Go baseline modules MUST compile and pass.

### AST-022
Each supplied Go canary MUST expose its named violation through the corresponding Go hook.

### AST-023
Go hooks MUST pass cleanly after the supplied canary is removed.

### AST-024
Explain MUST describe the cause and a corrective action for C02 (invalid module graph), C03 (unreadable or malformed documents/items), C12 (invalid metadata or zones) and C13 (queue identity errors).

### AST-025
Explain MUST describe the cause and a corrective action for C10 (missing enforced coverage), C11 (eval-hook/JUnit contract failures) and E01 (failed enforced testcases).

### AST-026
Explain MUST describe the cause and a corrective action for P01 (operational preconditions, including lifecycle, new-change and init), P02 (claims), P03 (unapproved state alteration), P04 (approval warning), P05 (stale bases), P06 (blocked change), P07 (unapproved-work bound), P09 (audit submissions) and Z01 (forbidden phase-zone edits).

### AST-027
Explain MUST describe the cause and a corrective action for C14 (doctor hook/canary setup, detection or recovery failures), C15 (agent probe failures or usage limits, and the shared-auditor-command warning) and C16 (missing/stale managed guidance or playbook embedded_base metadata).

### AST-028
The audit playbook MUST show a passing ## Audit example with No findings on one
line and Checked IDs: plus at least two illustrative IDs on the next line. It
MUST instruct the auditor to replace those with the actual checked IDs.

### AST-029
The evals playbook MUST instruct eval_fix workers to correct only evals, retain
coverage of listed IDs and require all module evals to pass instead of fail first.

### AST-030
The audit playbook MUST instruct eval_fix auditors to check that corrected tests
still assert the listed specs.

### AST-031
The impl playbook MUST instruct eval_fix workers to edit only the current change
file, leaving implementation unchanged.

### AST-032
Explain MUST describe P10 as a non-blocking regenerate warning that no passing
module drill is newer than its latest applied implementation. It MUST advise
running a module drill and reviewing and committing its report.

### AST-033
Explain MUST describe C07 as including eval_fix combined with fixes or spec
edits. It MUST advise separating eval-fix work from fix or spec changes.

### AST-034
Explain MUST describe H01 as a worker phase commit changing a protected path or project playbook override.
It MUST advise restoring that edit and separating human-owned setup or override work into an author commit outside worker phases.

### AST-035
Explain MUST describe H02 as an impl-to-applied commit changing evals or interface files.
It MUST advise making those edits in the evals step before implementation and reworking the affected phase history before rerunning history checks.

### AST-036
Explain MUST describe H03 as modifying, deleting or renaming an existing append-only path.
It MUST advise restoring the original path/content and adding a new file for new data.

### AST-037
Explain MUST describe H04 as modifying an archived change file, while allowing archived-file deletion for a revert.
It MUST advise restoring the archive and authoring follow-up work as a pending change.

### AST-038
Explain MUST describe B01 as a budget warning when summed JUnit testcase seconds for evaluated modules exceed budgets.verify_time, whose default is 15m.
It MUST advise reviewing expensive or redundant eval work while preserving required assertions.

### AST-039
Explain MUST describe B02 as a budget warning when a spec's testcase count exceeds budgets.cases_per_spec, whose default is 15.
It MUST advise consolidating redundant cases while preserving required behavioral assertions.

### AST-040
Explain MUST describe B03 as a budget warning when eval lines divided by disposable lines exceed budgets.eval_ratio, whose default is 4, including positive eval lines with zero disposable lines.
It MUST advise simplifying redundant eval or fixture code while preserving required assertions.

### AST-041
Explain MUST describe P08 as a warning when at least three consecutive applied changes target the same module.
It MUST advise reviewing that sequence's scope with the project owner before continuing.

### AST-042
Explain MUST describe B04 as a warning for binary files or files larger than 256 KB newly added to a disposable zone, reported by pr-body.
It MUST advise reviewing these opaque additions and removing or replacing unnecessary generated or vendored artifacts.

## Decisions

[0014-go-setup-bundle] Recipe files use .fireweed/recipes/go.md, .fireweed/hooks.go,
.fireweed/templates/module/SPEC.md and api.go.tmpl, and
.fireweed/canaries/canaries.yaml. Setup runs go mod download after the project has
a Go module; the recipe instructions explain that prerequisite.
