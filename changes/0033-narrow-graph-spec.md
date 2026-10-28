---
already_satisfied:
    - PRJ-002
    - new:graph-module-names
    - new:graph-prefixes
    - new:graph-dependency-targets
    - new:graph-cycles
bases:
    PRJ-002: 42ff7021
module: project
---



## Motivation

Continue the explicitly authorized project splits under SYSTEM's one-behavior
spec-shape rule. Separate PRJ-002's graph constraints from its validation result
contract, retaining the design's Modules and zone map invariants. All existing
Graph table assertions and behavior already satisfy the split.

## Interface



## Added

### new:graph-module-names
ValidateGraph MUST reject module names that are empty or duplicated in the graph.

### new:graph-prefixes
ValidateGraph MUST reject module prefixes that are empty or duplicated in the graph.

### new:graph-dependency-targets
ValidateGraph MUST reject a dependency that does not name a different existing module.

### new:graph-cycles
ValidateGraph MUST reject a graph containing a dependency cycle.

## Modified

### PRJ-002
ValidateGraph MUST return nil for a valid graph and an explanatory error identifying the rejection cause for an invalid graph.

## Removed



## Decisions



## Notes

Retarget the existing Graph table without adding cases. Keep PRJ-002 on the
validation-result/error assertions, and label existing rows with their Added
IDs: empty_name/duplicate_name, empty_prefix/duplicate_prefix,
missing_dependency/self_dependency, and cycle respectively. Preserve the
acyclic acceptance and every explanatory-error assertion. Keep the unchanged
53 reported project cases; no new fixtures, Test functions, t.Run cases or
helpers. No product code or interface changes are needed. Implementation may
edit only this change file. Retain every other test and ID.
