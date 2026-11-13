---
module: project
phase: impl
run: 20261113T230327-76aae603f900d8b331584964
---

## Motivation

The design's Modules and zone map invariants prohibit overlap between evals
and disposable zones (R02). Main 3b3b64d accepts evals: [evals/**] with both
disposable: [evals/**] and disposable: ["**/*_test.go"], even without files.
Precedence protects existing evals but does not satisfy this invariant.

## Interface

No public signature or type changes.

## Added

### PRJ-075
Check MUST reject with non-warning C12 when any module's resolved evals and disposable globs match a common possible module-relative file path, whether or not that file exists. The diagnostic MUST identify the module and conflicting globs.

### PRJ-076
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

Evals recorded fail-first for PRJ-075 and PRJ-076. The targeted Go run found
missing C12 diagnostics for identical and intersecting globs, an inherited
current overlap, and a queued overlap both before and after a later edit
cleared it. The disjoint and explicit-empty cases passed. Controller
`fireweed verify disjoint-eval-disposable-zones` passed in the evals phase.
Both checks can be implemented in `packages/project/internal/queue/`, where
`Check` reads the current modules and applies each queued change.
## Audit

No findings; checked PRJ-075, PRJ-076

## Provenance

- evals 2026-11-13T23:07:21.603372Z agent=codex fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=168 failing=9 passing=159
- audit 2026-11-13T23:07:38.255016Z agent=codex fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0

