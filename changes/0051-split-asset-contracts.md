---
already_satisfied:
    - AST-001
    - AST-002
    - AST-004
    - AST-005
    - AST-006
    - AST-007
    - AST-008
    - AST-009
    - AST-010
    - AST-011
    - AST-012
    - AST-013
    - AST-014
    - AST-015
    - AST-016
    - AST-017
    - AST-018
    - AST-019
    - AST-020
    - AST-021
    - AST-022
    - AST-023
audit_rounds: 0
bases:
    AST-001: 5416a9ec
    AST-002: 1a5bc52f
    AST-004: cea60c2f
    AST-005: aef8c7f5
    AST-006: "15183709"
    AST-007: 1d17def2
module: assets
phase: impl
run: 20261028T064006-41b309e9b491a89365abb050
---



## Motivation

Batch the previously identified assets bundles under SYSTEM's 30-item
spec-only cap. Separate standalone text APIs, playbook policy and Go bundle,
hook and template behaviors while preserving the design's Distribution,
Playbooks and Language support obligations. This is 22 affected items; all
behavior exists. Do not add recipes, commands, error paths or explanations.
No product implementation or public interface changes are authorized.

## Interface



## Added

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

## Modified

### AST-001
Version MUST return 0.1.0.

### AST-002
Playbook MUST return nonempty instructions for author-change, evals, audit, impl and regenerate, and reject unknown names.

### AST-004
BuiltinRecipe MUST return the embedded Go recipe for go and reject unknown names.

### AST-005
The Go eval hook MUST run the requested module's Go tests and write JUnit testcase names, classnames, pass/fail/skip and elapsed seconds while preserving test failure exit status.

### AST-006
The Go boundaries hook MUST query FIREWEED_BIN modules --json and use its module dirs, depends_on and zone map with the workspace module path.

### AST-007
The Go bundle MUST provide module SPEC/api.go templates and version-1 baseline, undeclared, internal_eval, failing_eval and missing_id canary maps.

## Removed



## Decisions



## Notes

Retarget existing assets_test.go and go_recipe_test.go labels only, preserving
every assertion and all 20 reported assets cases. Help uses its new ID;
AuditRules/AuthorRules gain their policy IDs instead of AST-002. Playbooks
retains AST-002. GoRecipe keeps AST-004 and gains payload/setup/graph IDs.
UnfinishedAndCompileFailure uses its new ID instead of AST-005; other Go eval
hook tests retain AST-005. Existing boundary tests retain AST-006 and gain
violation/clean-boundary IDs matching their assertions. TemplatesAndCanaries
retains AST-007 and gains substitution/baseline/detection/recovery IDs.
Keep AST-003 and its emitted-code explanations unchanged. Do not change raw
Go source strings, XML, template inputs, hook code, fixture paths, commands,
helper behavior or Decisions. No new Test functions, t.Run cases, fixture
instances, product code or interfaces. Implementation may edit only this
change file. No obligation or assertion is removed.

## Provenance

- evals 2026-10-28T06:41:22.871566Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=219 failing=0 passing=219
- audit 2026-10-28T06:42:42.545475Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0

## Audit

No findings. Checked AST-001, AST-002, AST-004, AST-005, AST-006, AST-007,
AST-008, AST-009, AST-010, AST-011, AST-012, AST-013, AST-014, AST-015,
AST-016, AST-017, AST-018, AST-019, AST-020, AST-021, AST-022, AST-023.
