---
bases: {}
module: assets
---



## Motivation

The design's CLI and Distribution sections require embedded help, version,
playbooks and diagnostic explanations. Supply the first executable checkpoint
with project-independent information before wiring the manual lifecycle slice.

## Interface

Expose Version() string, Help() string, Playbook(name string) (string, error),
and Explain(code string) (string, error) through thin assets/api.go wrappers.
All text and lookup behavior belong under internal/.

## Added

### new:version-help
Version MUST return `0.1.0`. Help MUST identify fireweed and list every command
in the design's CLI table, including version, explain and playbook.

### new:workflow-playbooks
Playbook MUST return nonempty instructions for author-change, evals, audit, impl
and regenerate, and reject unknown names. The audit instructions MUST restrict
findings to exact spec clauses without assertions and limit re-audits to prior
finding IDs. The author instructions MUST require a user-visible or design
Motivation and recommend narrowing costly specs with already_satisfied.

### new:diagnostic-explanations
Explain MUST describe the cause and a corrective action for C01 (invalid project
configuration) and P2-01 (missing genuine fail-first evidence). Explain MUST
reject unknown codes.

## Modified



## Removed



## Decisions



## Rejected alternatives
