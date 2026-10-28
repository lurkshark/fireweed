---
audit_rounds: 0
bases:
    AST-003: 63d2d29a
module: assets
phase: audit
run: 20261028T080146-819e3aba6c430abbfef92056
---



## Motivation

The design requires stable diagnostics with useful explain output. Applied 0057
and 0058 emit C08/C09 and X01–X04 on existing paths, and broaden C05 to duplicate
edit identities. D006 requires explanations accurate for actually emitted codes.
Complete their explanations without adding unused allocations or features.

## Interface

No interface change.

## Added



## Modified

### AST-003
Explain MUST describe the cause and a corrective action for C01 (invalid configuration), C04 (non-normative text), C05 (foreign or duplicate spec/edit identities), C06 (invalid tags), C07 (missing Motivation), C08 (missing/stale edit bases or absent targets), C09 (undeclared handles), X01 (boundary violations), X02 (missing/incomplete boundary hooks), X03 (reasoned boundary opt-out), X04 (unnecessary boundary enforcement) and P2-01 (missing genuine fail-first evidence). Explain MUST reject unknown codes.

## Removed



## Decisions



## Notes

Preserve existing explanation and unknown-code assertions. In the existing
Explanations test, add one representative case for each of C08, C09, X01, X02,
X03 and X04, using requirePattern to check the specific cause and a useful
corrective action. Strengthen existing C05 assertions/text for duplicate edit
identities: removing or combining a repeated edit is a valid repair, not merely
assigning a new spec ID. C08 covers missing/stale bases and absent preceding
targets; X02 covers required missing hooks and failures to complete. Warnings
X03/X04 must be explained as such, with an action appropriate to their reason.
No new fixture/helper infrastructure, instrumentation or unused explanations.
The newly supported codes must genuinely fail before implementation. Add only
explanation text and the scoped C05 clarification; preserve recipes/playbooks
and all other product behavior. Controller owns specs, commits and archive.

## Provenance

- evals 2026-10-28T08:04:31.344119Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=230 failing=8 passing=222
