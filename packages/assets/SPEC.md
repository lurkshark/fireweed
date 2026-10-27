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
Version MUST return `0.1.0`. Help MUST identify fireweed and list every command
in the design's CLI table, including version, explain and playbook.

### AST-002
Playbook MUST return nonempty instructions for author-change, evals, audit, impl
and regenerate, and reject unknown names. The audit instructions MUST restrict
findings to exact spec clauses without assertions and limit re-audits to prior
finding IDs. The author instructions MUST require a user-visible or design
Motivation and recommend narrowing costly specs with already_satisfied.

### AST-003
Explain MUST describe the cause and a corrective action for C01 (invalid project
configuration) and P2-01 (missing genuine fail-first evidence). Explain MUST
reject unknown codes.

### AST-004
BuiltinRecipe MUST return the embedded Go recipe for go and reject unknown names. The bundle MUST contain instructions, api.go/evals/**/internal/** defaults, eval/boundaries/setup commands, go.sum worker_writable and project-relative recipe/hook/template/canary files. Its instructions MUST explain Go module/workspace setup and deriving the graph from FIREWEED_BIN modules --json.

### AST-005
The Go eval hook MUST run the requested module's Go tests and write JUnit testcase names/classnames, pass/fail/skip and elapsed seconds while preserving test failure exit status. It MUST retain started-but-unfinished tests and package/compile failures as visible failures rather than passing or empty evidence.

### AST-006
The Go boundaries hook MUST query FIREWEED_BIN modules --json and use its module dirs, depends_on and zone map with the workspace module path. It MUST report module-prefixed violations for imports of undeclared modules and eval imports reaching beyond retained interface packages, including own internal implementation. Clean module fixtures MUST pass.

### AST-007
The bundle MUST provide module SPEC/api.go templates and version-1 baseline, undeclared, internal_eval, failing_eval and missing_id canary maps, substitutable with name, prefix, package, go_module, module_dir and other_dir. Rendered baseline modules MUST compile/pass. Each supplied canary MUST expose its named violation through the corresponding Go hook, with clean recovery after removal.

## Decisions

[0014-go-setup-bundle] Recipe files use .fireweed/recipes/go.md, .fireweed/hooks.go,
.fireweed/templates/module/SPEC.md and api.go.tmpl, and
.fireweed/canaries/canaries.yaml. Setup runs go mod download after the project has
a Go module; the recipe instructions explain that prerequisite.
