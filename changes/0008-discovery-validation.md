---
module: project
---

## Motivation

Static inspection needs a deterministic repository snapshot and diagnostics before checked authoring and lifecycle commands can be composed. Discovery must use the same module metadata and zone contracts as execution, without running hooks.

## Interface

Add Diagnostic {Code, Severity, Message, Path string}, Snapshot {Root, BaseRef string; Config contracts.Config; Modules []contracts.Module; Changes, Archive []contracts.Change; Diagnostics []Diagnostic}, Load(start string) (Snapshot,error) and Check(Snapshot) []Diagnostic. Keep retained data/facades thin; discovery and validation live in internal/. Load uses public contracts, gitstore and assets interfaces. Check includes discovery diagnostics and has no I/O.

## Added

### new:discovery
Load MUST discover the canonical Git root from a nested directory, parse fireweed.yaml using the running embedded version, resolve the local base ref, and discover SPEC.md only in module directories matching configured modules globs. It MUST retain resolved module defaults, explicit empty overrides, dependencies, hooks, specs, Decisions and original documents. Modules MUST be sorted by name then directory. It MUST NOT follow directory symlinks or read files through escaping symlinks, run hooks/agents, use the network or mutate repository files/refs. Missing/invalid configuration or invalid configured globs MUST return an error mentioning C01; a nonrepository MUST return an error.

### new:documents
Load MUST continue past malformed module/change documents and retain path-attributed diagnostics: C02 warning for SPEC.md outside module globs, C03 error for invalid module metadata, C04 error for invalid Specs grammar, and C05 error for malformed changes or unknown target modules. It MUST ignore .git and .fireweed trees during module discovery. Queue and archive MUST include immediate Markdown files only, parsed with each target's prefix, retaining lifecycle/provenance/author sections; each list MUST sort by numeric filename then slug. Duplicate slugs across either list MUST produce C05. Symlink escapes MUST produce an error diagnostic without reading their contents.

### new:graph
Check MUST diagnose duplicate module names or prefixes with C03 and overlapping module directory ownership with C06, including future paths. Dependencies MUST name existing modules, exclude self and have no cycles; violations MUST produce C06. Valid acyclic graphs, including shared transitive dependencies, MUST pass this graph check.

### new:zones
Check MUST reject invalid module-relative zone patterns with C03 and eval/disposable intersections with C06, including intersections at nonexistent paths. Nonintersecting zones MUST pass; interface/preserve/append-only overlap with disposable MUST remain permitted under strict precedence. It MUST NOT require files to exist or execute hooks.

### new:hooks
Check MUST require a nonblank resolved eval hook (C07 error). Modules with disposable globs MUST have a nonblank boundaries hook or a nonblank boundaries_opt_out reason (C07 error otherwise). Every explicit opt-out MUST emit a D04 warning containing its reason; a module without disposable globs and without boundaries MUST emit a D04 warning. Boundary opt-outs MUST NOT waive the eval hook requirement.

### new:reservations
Check MUST reject any current spec ID listed as Removed in the archive with C08, preserving removed-ID reservations. It MUST include all Load diagnostics and return a non-null deterministically sorted diagnostic list by path, code, severity and message, with nonempty corrective messages and path attribution for module/change issues. A valid snapshot with no warnings MUST return an empty list. Check MUST NOT mutate the supplied snapshot.

## Modified

## Removed

## Decisions

This slice validates the current static graph. Queued projection, set application, hash bases and late ID assignment are separate changes. Discovery scans all repository directories except .git and .fireweed to report out-of-glob specs; WalkDir does not traverse directory symlinks. Symlink file paths are checked by gitstore.SafePath before reads. Configured globs include module discovery, defaults, protected and worker_writable patterns. Module metadata is validated separately from its Specs body so malformed normative content receives C04. Unknown archived target modules are diagnosed C05, rather than silently dropping history. Hook strings remain opaque and are never evaluated by static checks. Module path attribution is <dir>/SPEC.md; queue/archive attribution is its repository-relative path. A root module directory is represented by '.'.

Variable-width numeric-order fixtures use at least four digits, as required by the retained document contract. Original documents retain all valid known metadata and arbitrary human sections; unknown metadata is rejected by contracts. Adjacent stars inside a segment are ordinary wildcards; invalid-pattern tests use malformed classes, absolute/traversal paths and repeated separators. The superseded first cycle was fully reverted after these fixture mistakes surfaced (author decision D008).

## Rejected alternatives

Running hooks during discovery would make read commands stateful and conflate static checks with enforcement. A second graph in language-specific tooling would drift from SPEC.md metadata. Ignoring malformed files would hide loss of ownership or archive reservations.
