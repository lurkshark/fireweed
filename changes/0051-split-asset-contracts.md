---
already_satisfied:
    - AST-001
    - AST-002
    - AST-004
    - AST-005
    - AST-006
    - AST-007
    - new:help-catalog
    - new:audit-clause-scope
    - new:audit-repeated-ids
    - new:author-motivation
    - new:author-cost-guidance
    - new:go-bundle-payload
    - new:go-setup-instructions
    - new:go-graph-instructions
    - new:go-incomplete-evals
    - new:go-undeclared-imports
    - new:go-eval-interfaces
    - new:go-clean-boundaries
    - new:go-template-substitution
    - new:go-baseline-templates
    - new:go-canary-detection
    - new:go-canary-recovery
bases:
    AST-001: 5416a9ec
    AST-002: 1a5bc52f
    AST-004: cea60c2f
    AST-005: aef8c7f5
    AST-006: "15183709"
    AST-007: 1d17def2
module: assets
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

### new:help-catalog
Help MUST identify fireweed and list every command in the design's CLI table, including version, explain and playbook.

### new:audit-clause-scope
The audit playbook MUST restrict findings to exact spec clauses without assertions.

### new:audit-repeated-ids
The audit playbook MUST restrict re-audits to IDs with prior findings.

### new:author-motivation
The author playbook MUST require a user-visible or design basis for Motivation.

### new:author-cost-guidance
The author playbook MUST recommend narrowing costly specs with already_satisfied.

### new:go-bundle-payload
The Go bundle MUST contain instructions, api.go/evals/**/internal/** defaults, eval/boundaries/setup commands, go.sum worker_writable and project-relative recipe, hook, template and canary files.

### new:go-setup-instructions
Go recipe instructions MUST explain Go module and workspace setup.

### new:go-graph-instructions
Go recipe instructions MUST explain deriving the graph from FIREWEED_BIN modules --json.

### new:go-incomplete-evals
The Go eval hook MUST retain started-but-unfinished tests and package or compile failures as visible failures rather than passing or empty evidence.

### new:go-undeclared-imports
The Go boundaries hook MUST report module-prefixed violations for imports of undeclared modules.

### new:go-eval-interfaces
The Go boundaries hook MUST report module-prefixed violations for eval imports reaching beyond retained interface packages, including own internal implementation.

### new:go-clean-boundaries
Clean module fixtures MUST pass the Go boundaries hook.

### new:go-template-substitution
The Go bundle templates and canary snippets MUST be substitutable with name, prefix, package, go_module, module_dir and other_dir.

### new:go-baseline-templates
Rendered Go baseline modules MUST compile and pass.

### new:go-canary-detection
Each supplied Go canary MUST expose its named violation through the corresponding Go hook.

### new:go-canary-recovery
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
