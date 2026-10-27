---
bases: {}
module: assets
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

### new:go-bundle
BuiltinRecipe MUST return the embedded Go recipe for go and reject unknown names. The bundle MUST contain instructions, api.go/evals/**/internal/** defaults, eval/boundaries/setup commands, go.sum worker_writable and project-relative recipe/hook/template/canary files. Its instructions MUST explain Go module/workspace setup and deriving the graph from FIREWEED_BIN modules --json.

### new:go-eval-hook
The Go eval hook MUST run the requested module's Go tests and write JUnit testcase names/classnames, pass/fail/skip and elapsed seconds while preserving test failure exit status. It MUST retain started-but-unfinished tests and package/compile failures as visible failures rather than passing or empty evidence.

### new:go-boundary-hook
The Go boundaries hook MUST query FIREWEED_BIN modules --json and use its module dirs, depends_on and zone map with the workspace module path. It MUST report module-prefixed violations for imports of undeclared modules and eval imports reaching beyond retained interface packages, including own internal implementation. Clean module fixtures MUST pass.

### new:go-templates-canaries
The bundle MUST provide module SPEC/api.go templates and version-1 baseline, undeclared, internal_eval, failing_eval and missing_id canary maps, substitutable with name, prefix, package, go_module, module_dir and other_dir. Rendered baseline modules MUST compile/pass. Each supplied canary MUST expose its named violation through the corresponding Go hook, with clean recovery after removal.

## Modified



## Removed



## Decisions

Recipe files use .fireweed/recipes/go.md, .fireweed/hooks.go,
.fireweed/templates/module/SPEC.md and api.go.tmpl, and
.fireweed/canaries/canaries.yaml. Setup runs go mod download after the project has
a Go module; the recipe instructions explain that prerequisite.
