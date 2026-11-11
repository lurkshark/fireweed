---
bases:
    AST-041: dcc48e02
module: assets
phase: audit
run: 20261111T091304-6f7aa87891652f11b09c0272
---

## Motivation

`fireweed explain P08` describes the warning as three consecutive applied
changes on one module (AST-041). 0217 changes P08 to warn on a chain of
applied changes that rework the same spec IDs, which is the regression
chain the design names as the warning's purpose. The explanation has to
describe the new trigger.

## Interface

None. Only the P08 explanation text changes.

## Added

## Modified

### AST-041
Explain MUST describe P08 as a warning that recently applied changes form a chain in which each one reworks spec IDs the previous one added or modified.
It MUST advise reviewing that chain with the project owner before continuing.

## Removed

## Decisions

## Notes

Keep the explanation to two or three sentences and avoid the window size
and chain length, which belong to workflow's FLOW-111.

`go test ./packages/assets/evals -run '^Test_AST_041_P08Explanation$' -count=1`
failed because P08 still describes consecutive changes to one module and
advises reviewing their scope. `fireweed verify rework-chain-explanation`
accepted the AST-041 fail-first evidence. The other 124 testcases passed.

The P08 explanation lives in the implementation-writable file
`packages/assets/internal/content/content.go`.
## Provenance
- evals 2026-11-11T09:16:09.100411Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=125 failing=1 passing=124
