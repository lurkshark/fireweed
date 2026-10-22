---
module: project
phase: impl
run: 5bac8caec4836214c46ee6c9
---

## Motivation

Authors need a faithful preview of preceding queued changes before stamping bases, and workers need projected metadata/zones from the first step. Keep deterministic projection separate from file writing and lifecycle state.

## Interface

Add ApplyChange(module contracts.Module, change contracts.Change) (contracts.Module,error), Project(snapshot Snapshot, before string) (Snapshot,error), and BaseHashes(snapshot Snapshot, slug string) (map[string]string,error). All are pure retained facades to disposable implementation. Project's empty selector applies all pending changes; a nonempty selector excludes that change and everything after it. Extend permanent public evals and the existing native public probe to cover new APIs without I/O.

## Added

### PRJ-007
ApplyChange MUST replace Modified specs in their existing positions, remove Removed specs, and append Added specs in declared order, retaining complete bodies and tags. It MUST retain local Added handles for preview without allocating IDs. It MUST reject a mismatched target module, missing Modified/Removed targets, repeated edit IDs, existing-ID additions, wrong-prefix IDs, malformed normative edits and invalid Removed bodies with an error rather than returning a partial result. It MUST update the document's Specs section to reflect the resulting ordered items while retaining the module preamble and unrelated sections.

### PRJ-008
ApplyChange MUST apply only allowed set fields: depends_on, interface, evals, disposable, preserve, append_only, boundaries_opt_out and named eval/boundaries/setup hook overrides. Explicit empty lists/strings MUST replace resolved values, omitted fields MUST retain them, and partial hooks MUST retain other resolved hooks. Identity edits, unknown fields and wrong field types MUST fail. The returned module's fields and document metadata MUST agree, retaining original valid metadata and unrelated document sections; name, prefix and directory MUST remain unchanged. Projection MUST NOT read a configuration file or re-inherit project defaults over explicit values.

### PRJ-009
Project MUST apply changes in numeric Number then Slug order across their target modules, independent of input slice order. A nonempty before selector MUST return the state after preceding changes only; empty before MUST apply the entire queue. Same-number changes MUST follow slug order. Unknown/ambiguous selectors, unknown target modules and invalid application MUST return errors. Current archive entries MUST NOT be replayed as pending changes. Project MUST preserve unrelated snapshot data, and neither successful nor failed projection/application MAY mutate inputs or return mutable aliases of any input module, document, configuration, change, archive, metadata map or slice. These APIs MUST NOT read/write files, run hooks/agents/Git or use networking.

### PRJ-010
BaseHashes MUST compute a fresh non-null map containing exactly the Modified and Removed IDs of the selected queued change, hashing the target specs as projected immediately before that change. Hashes MUST use the public normalized eight-character SHA-256 item contract, including tags; existing bases MUST be replaced in the returned map without changing input state. Added-only or maintenance changes MUST return an empty map. Missing/ambiguous change, missing target module or missing edited spec MUST return an error. The function MUST remain pure and independent of files and Git, and changing a returned map MUST NOT affect inputs or later calls.

## Modified

## Removed

## Decisions

Projection is an intent preview: it does not enforce existing bases or allocate real IDs. Later queue validation reports missing/stale bases, invalid graph/zones after set, and archive reservations; lifecycle allocation rewrites handles only at start. ApplyChange uses the typed Module and Change fields as authoritative, updating retained documents accordingly. Local handles may appear in Added and references, with the same uniqueness/locality rules as the public change parser; file-backed current module specs and Modified/Removed targets use real IDs. Projected intermediate modules may contain placeholders from preceding changes. Field/type and spec-edit validation is required even for callers constructing typed values directly. An empty module Document is synthesized from typed fields; valid original documents retain preamble and human sections. Project keeps the original pending changes/archive lists and diagnostics as independent copies, preserving their original ordering; only the order of application is sorted. Diagnostics, canonical module ordering and graph validation remain separate existing Check responsibilities. Decisions application/removal will be a separately specified part of archival application; this preview retains existing module Decisions.

ApplyChange errors return the zero Module, never a partial result. Existing-ID addition collisions mean real allocated IDs; handle placeholders are local declarations, so two queued changes may each declare new:local and their preview retains both items without treating that shared spelling as an allocated-ID collision. References in each change must still be declared locally. A preview document can therefore contain repeated placeholder headings from different changes, and is never written as a current SPEC.md until start has allocated real IDs. Invalid excluded target/tail applications are not executed or validated by Project(before); ambiguity of the selector itself is still an error. Base hashes follow typed spec fields, including normalized whitespace, rather than stale retained document text.

## Rejected alternatives

Reading working files during projection would break determinism for callers holding a snapshot. Allocating IDs while previewing would reintroduce parallel author collisions. Applying queue input order would make same-number changes depend on filesystem iteration.

## Provenance

- evals | 2026-10-22T08:46:12Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 34 | failing: 7 | passing: 27
- audit | 2026-10-22T08:55:26Z | agent: manual | fireweed: 0.4.0 | playbook: a9c982d24b476621316049b527e6c7bbc319f06b12268af150fa7d8c6df92230 | tests: 0 | failing: 0 | passing: 0
- evals | 2026-10-22T08:58:10Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 40 | failing: 13 | passing: 27
- audit | 2026-10-22T09:00:52Z | agent: manual | fireweed: 0.4.0 | playbook: a9c982d24b476621316049b527e6c7bbc319f06b12268af150fa7d8c6df92230 | tests: 0 | failing: 0 | passing: 0

## Audit

No findings. Independent auditor /root/audit_discovery_validation checked PRJ-007, PRJ-008, PRJ-009 and PRJ-010. Previous position, metadata, handle scope, hash and isolation gaps are resolved. Fixture premises conform to public grammar; retained-document staleness in direct fixtures follows the typed-fields-authoritative Decision. No implementation/history/bootstrap inspection or test execution by auditor.
