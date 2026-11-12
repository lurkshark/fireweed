---
audit_rounds: 0
bases: {}
module: assets
phase: applied
run: 20261026T201612-66519d7778785672f6878ec0
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

## Modified



## Removed



## Decisions



## Rejected alternatives



## Provenance

- evals 2026-10-26T20:19:35.671397Z agent=manual fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=8 failing=8 passing=0
- audit 2026-10-26T20:21:29.908957Z agent=manual fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0
- impl 2026-10-26T20:22:36.707164Z agent=manual fireweed=0.1.0 playbook=07554a3249d8fca0ba70df7504135f297491465caa63458734a4b8351734d0e5 tests=8 failing=0 passing=8

## Audit

No findings. Checked IDs: AST-001, AST-002, AST-003.
