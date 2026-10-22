---
module: project
---

## Motivation

Authors need a faithful preview of preceding queued changes before stamping bases, and workers need projected metadata/zones from the first step. Keep deterministic projection separate from file writing and lifecycle state.

## Interface

Add ApplyChange(module contracts.Module, change contracts.Change) (contracts.Module,error), Project(snapshot Snapshot, before string) (Snapshot,error), and BaseHashes(snapshot Snapshot, slug string) (map[string]string,error). All are pure retained facades to disposable implementation. Project's empty selector applies all pending changes; a nonempty selector excludes that change and everything after it. Extend permanent public evals and the existing native public probe to cover new APIs without I/O.

## Added

### new:spec-application
ApplyChange MUST replace Modified specs in their existing positions, remove Removed specs, and append Added specs in declared order, retaining complete bodies and tags. It MUST retain local Added handles for preview without allocating IDs. It MUST reject a mismatched target module, missing Modified/Removed targets, repeated edit IDs, existing-ID additions, wrong-prefix IDs, malformed normative edits and invalid Removed bodies with an error rather than returning a partial result. It MUST update the document's Specs section to reflect the resulting ordered items while retaining the module preamble and unrelated sections.

### new:metadata-projection
ApplyChange MUST apply only allowed set fields: depends_on, interface, evals, disposable, preserve, append_only, boundaries_opt_out and named eval/boundaries/setup hook overrides. Explicit empty lists/strings MUST replace resolved values, omitted fields MUST retain them, and partial hooks MUST retain other resolved hooks. Identity edits, unknown fields and wrong field types MUST fail. The returned module's fields and document metadata MUST agree, retaining original valid metadata and unrelated document sections; name, prefix and directory MUST remain unchanged. Projection MUST NOT read a configuration file or re-inherit project defaults over explicit values.

### new:queue-projection
Project MUST apply changes in numeric Number then Slug order across their target modules, independent of input slice order. A nonempty before selector MUST return the state after preceding changes only; empty before MUST apply the entire queue. Same-number changes MUST follow slug order. Unknown/ambiguous selectors, unknown target modules and invalid application MUST return errors. Current archive entries MUST NOT be replayed as pending changes. Project MUST preserve unrelated snapshot data, and neither successful nor failed projection/application MAY mutate inputs or return mutable aliases of any input module, document, configuration, change, archive, metadata map or slice. These APIs MUST NOT read/write files, run hooks/agents/Git or use networking.

### new:base-hashes
BaseHashes MUST compute a fresh non-null map containing exactly the Modified and Removed IDs of the selected queued change, hashing the target specs as projected immediately before that change. Hashes MUST use the public normalized eight-character SHA-256 item contract, including tags; existing bases MUST be replaced in the returned map without changing input state. Added-only or maintenance changes MUST return an empty map. Missing/ambiguous change, missing target module or missing edited spec MUST return an error. The function MUST remain pure and independent of files and Git, and changing a returned map MUST NOT affect inputs or later calls.

## Modified

## Removed

## Decisions

Projection is an intent preview: it does not enforce existing bases or allocate real IDs. Later queue validation reports missing/stale bases, invalid graph/zones after set, and archive reservations; lifecycle allocation rewrites handles only at start. ApplyChange uses the typed Module and Change fields as authoritative, updating retained documents accordingly. Local handles may appear in Added and references, with the same uniqueness/locality rules as the public change parser; current module specs and Modified/Removed targets use real IDs. Field/type and spec-edit validation is required even for callers constructing typed values directly. An empty module Document is synthesized from typed fields; valid original documents retain preamble and human sections. Project keeps the original pending changes/archive lists and diagnostics as independent copies, preserving their original ordering; only the order of application is sorted. Diagnostics, canonical module ordering and graph validation remain separate existing Check responsibilities. Decisions application/removal will be a separately specified part of archival application; this preview retains existing module Decisions.

## Rejected alternatives

Reading working files during projection would break determinism for callers holding a snapshot. Allocating IDs while previewing would reintroduce parallel author collisions. Applying queue input order would make same-number changes depend on filesystem iteration.
