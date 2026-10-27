---
audit_pending:
    - AST-004
    - AST-005
    - AST-006
audit_rounds: 1
bases: {}
module: assets
phase: impl
run: 20261027T145815-f26a804e7d062856d3d8e07b
set:
    interface:
        - api.go
        - model/**
---



## Motivation

The design's First run and Recipes sections require init to copy working Go
setup into a project. The manual slice is verified; package the Go instructions,
hooks, module templates and doctor canaries next so init can use them. Other
recipes and setup command composition remain later work.

## Interface

Retain data-only Recipe with Name/Instructions strings, ModuleDefaults
contracts.Zones, Hooks map[string]any, Protected/WorkerWritable []string and
Files map[string]string keyed by project-relative output path. Add thin
BuiltinRecipe(name string) (Recipe,error). Behavior/resources belong under
internal; shared data lives under retained model/.

## Added

### AST-004
BuiltinRecipe MUST return the embedded Go recipe for go and reject unknown names. The bundle MUST contain instructions, api.go/evals/**/internal/** defaults, eval/boundaries/setup commands, go.sum worker_writable and project-relative recipe/hook/template/canary files. Its instructions MUST explain Go module/workspace setup and deriving the graph from FIREWEED_BIN modules --json.

### AST-005
The Go eval hook MUST run the requested module's Go tests and write JUnit testcase names/classnames, pass/fail/skip and elapsed seconds while preserving test failure exit status. It MUST retain started-but-unfinished tests and package/compile failures as visible failures rather than passing or empty evidence.

### AST-006
The Go boundaries hook MUST query FIREWEED_BIN modules --json and use its module dirs, depends_on and zone map with the workspace module path. It MUST report module-prefixed violations for imports of undeclared modules and eval imports reaching beyond retained interface packages, including own internal implementation. Clean module fixtures MUST pass.

### AST-007
The bundle MUST provide module SPEC/api.go templates and version-1 baseline, undeclared, internal_eval, failing_eval and missing_id canary maps, substitutable with name, prefix, package, go_module, module_dir and other_dir. Rendered baseline modules MUST compile/pass. Each supplied canary MUST expose its named violation through the corresponding Go hook, with clean recovery after removal.

## Modified



## Removed



## Decisions

Recipe files use .fireweed/recipes/go.md, .fireweed/hooks.go,
.fireweed/templates/module/SPEC.md and api.go.tmpl, and
.fireweed/canaries/canaries.yaml. Setup runs go mod download after the project has
a Go module; the recipe instructions explain that prerequisite.

## Provenance

- evals 2026-10-27T18:58:30.86087Z agent=manual fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=213 failing=6 passing=207
- audit 2026-10-27T19:00:02.64701Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0
- evals 2026-10-27T19:02:27.313221Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=215 failing=8 passing=207
- audit 2026-10-27T19:02:52.177804Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0

## Audit

No findings. Checked IDs: AST-004, AST-005, AST-006.
