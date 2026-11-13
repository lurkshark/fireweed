---
module: project
---

## Motivation

The design's Modules and zone map invariants prohibit overlap between evals
and disposable zones (R02). Main 3b3b64d accepts evals: [evals/**] with both
disposable: [evals/**] and disposable: ["**/*_test.go"], even without files.
Precedence protects existing evals but does not satisfy this invariant.

## Interface

No public signature or type changes.

## Added

### new:disjoint-eval-disposable-globs
Check MUST reject with non-warning C12 when any module's resolved evals and disposable globs match a common possible module-relative file path, whether or not that file exists. The diagnostic MUST identify the module and conflicting globs.

### new:projected-zone-overlap
Check MUST apply the eval/disposable overlap rule to the current graph and each preceding-through-current queued projection, including inherited defaults, explicit empty lists and set edits.

## Modified



## Removed

## Decisions

## Notes

Use CON-005's existing * and ** glob language. Check possible paths, not only
tracked files or files present on disk. Cover an identical glob, a cross-pattern
overlap such as evals/** with **/*_test.go, disjoint siblings, explicit empty
overrides, and a queue set edit that introduces overlap. Retain permitted
interface/disposable and preserve/disposable overlap and zone precedence.
Prefer a small bounded set of public fixtures; no filesystem enumeration or
arbitrary search-depth bound can establish disjointness for future paths.
Return C12 rather than repurposing dependency-graph C02.

Drafted at the owner's request after the R01–R29 traceability review.
No lifecycle execution, agent invocation, drill or release is authorized by
this drafting request. Workers must obtain context before their step and use
fireweed verify; only the controller may advance or commit lifecycle work.
