---
depends_on:
    - contracts
id_prefix: AST
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

## Decisions
