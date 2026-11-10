---
bases:
    AST-041: dcc48e02
module: assets
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
