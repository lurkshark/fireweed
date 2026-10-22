---
depends_on:
    - contracts
    - gitstore
    - assets
id_prefix: PRJ
name: project
---

# project

## Specs

### PRJ-001
Load MUST discover the canonical Git root from a nested directory, parse fireweed.yaml using the running embedded version, resolve the local base ref, and discover SPEC.md only in module directories matching configured modules globs. It MUST retain resolved module defaults, explicit empty overrides, dependencies, hooks, specs, Decisions and original documents. Modules MUST be sorted by name then directory. It MUST NOT follow directory symlinks or read files through escaping symlinks, run hooks/agents, use the network or mutate repository files/refs. Missing/invalid configuration or invalid configured globs MUST return an error mentioning C01; a nonrepository MUST return an error.

### PRJ-002
Load MUST continue past malformed module/change documents and retain path-attributed diagnostics: C02 warning for SPEC.md outside module globs, C03 error for invalid module metadata, C04 error for invalid Specs grammar, and C05 error for malformed changes or unknown target modules, except that undeclared live handle references MUST produce C09 error diagnostics. Queued handle references anywhere in the change, including human sections and metadata, MUST be local to that change; a declaration in a different change MUST NOT satisfy them. Archive human sections MUST preserve literal handle examples without treating them as live references; archive edit sections and metadata MUST still validate their live references. It MUST ignore .git and .fireweed trees during module discovery. Queue and archive MUST include immediate Markdown files only, parsed with each target's prefix, retaining lifecycle/provenance/author sections; each list MUST sort by numeric filename then slug. Duplicate slugs across either list MUST produce C05. Symlink escapes MUST produce an error diagnostic without reading their contents.

### PRJ-003
Check MUST diagnose duplicate module names or prefixes with C03 and overlapping module directory ownership with C06, including future paths. Dependencies MUST name existing modules, exclude self and have no cycles; violations MUST produce C06. Valid acyclic graphs, including shared transitive dependencies, MUST pass this graph check.

### PRJ-004
Check MUST reject invalid module-relative zone patterns with C03 and eval/disposable intersections with C06, including intersections at nonexistent paths. Nonintersecting zones MUST pass; interface/preserve/append-only overlap with disposable MUST remain permitted under strict precedence. It MUST NOT require files to exist or execute hooks.

### PRJ-005
Check MUST require a nonblank resolved eval hook (C07 error). Modules with disposable globs MUST have a nonblank boundaries hook or a nonblank boundaries_opt_out reason (C07 error otherwise). Every explicit opt-out MUST emit a D04 warning containing its reason; a module without disposable globs and without boundaries MUST emit a D04 warning. Boundary opt-outs MUST NOT waive the eval hook requirement.

### PRJ-006
Check MUST reject any current spec ID listed as Removed in the archive with C08, preserving removed-ID reservations. It MUST include all Load diagnostics and return a non-null deterministically sorted diagnostic list by path, code, severity and message, with nonempty corrective messages and path attribution for module/change issues. A valid snapshot with no warnings MUST return an empty list. Check MUST NOT mutate the supplied snapshot.

### PRJ-007
ApplyChange MUST replace Modified specs in their existing positions, remove Removed specs, and append Added specs in declared order, retaining complete bodies and tags. It MUST retain local Added handles for preview without allocating IDs. It MUST reject a mismatched target module, missing Modified/Removed targets, repeated edit IDs, existing-ID additions, wrong-prefix IDs, malformed normative edits and invalid Removed bodies with an error rather than returning a partial result. It MUST update the document's Specs section to reflect the resulting ordered items while retaining the module preamble and unrelated sections.

### PRJ-008
ApplyChange MUST apply only allowed set fields: depends_on, interface, evals, disposable, preserve, append_only, boundaries_opt_out and named eval/boundaries/setup hook overrides. Explicit empty lists/strings MUST replace resolved values, omitted fields MUST retain them, and partial hooks MUST retain other resolved hooks. Identity edits, unknown fields and wrong field types MUST fail. The returned module's fields and document metadata MUST agree, retaining original valid metadata and unrelated document sections; name, prefix and directory MUST remain unchanged. Projection MUST NOT read a configuration file or re-inherit project defaults over explicit values.

### PRJ-009
Project MUST apply changes in numeric Number then Slug order across their target modules, independent of input slice order. A nonempty before selector MUST return the state after preceding changes only; empty before MUST apply the entire queue. Same-number changes MUST follow slug order. Unknown/ambiguous selectors, unknown target modules and invalid application MUST return errors. Current archive entries MUST NOT be replayed as pending changes. Project MUST preserve unrelated snapshot data, and neither successful nor failed projection/application MAY mutate inputs or return mutable aliases of any input module, document, configuration, change, archive, metadata map or slice. These APIs MUST NOT read/write files, run hooks/agents/Git or use networking.

### PRJ-010
BaseHashes MUST compute a fresh non-null map containing exactly the Modified and Removed IDs of the selected queued change, hashing the target specs as projected immediately before that change. Hashes MUST use the public normalized eight-character SHA-256 item contract, including tags; existing bases MUST be replaced in the returned map without changing input state. Added-only or maintenance changes MUST return an empty map. Missing/ambiguous change, missing target module or missing edited spec MUST return an error. The function MUST remain pure and independent of files and Git, and changing a returned map MUST NOT affect inputs or later calls.

### PRJ-011
CheckQueue MUST process changes in numeric Number then Slug order against preceding structurally valid intent. Every Modified and Removed ID MUST exist immediately before its change and have a matching normalized public SpecHash in bases; missing targets, missing hashes and stale hashes MUST emit C10 error diagnostics at that change's path. Tag changes MUST affect comparison while normalized line endings/trailing whitespace MUST NOT cause false staleness. Stale or missing hashes MUST NOT stop projecting structurally valid intent for later base checks.

### PRJ-012
CheckQueue MUST apply each structurally valid change's set values before validating graph, zones, hook requirements and archive reservations with the public static contract. Newly introduced C03, C06, C07, C08 and D04 diagnostics MUST be attributed to the introducing change path and retain corrective detail, including opt-out reasons. It MUST NOT repeat unchanged preexisting current/static/discovery diagnostics, and a valid later change that removes an introduced problem MUST pass. Every intermediate projected state MUST be checked, even when a later change repairs it; future-path intersections MUST be enforced without requiring files to exist.

### PRJ-013
Unknown target modules and invalid typed changes/applications MUST emit C05 errors; undeclared or cross-change-only handles anywhere in a change MUST emit C09 errors. Structurally invalid applications MUST be skipped while later independent changes are still checked. Real-ID Added items without run and phase MUST be rejected as author-chosen IDs with C05; a removed archive ID MUST be rejected with C08 even when the change also lacks assigned state. Assigned-state additions with unused valid IDs and local author handles MUST remain accepted. fixes IDs MUST exist before the change and remain present after it; already_satisfied IDs MUST exist in the resulting module. Missing or removed references MUST emit C05.

### PRJ-014
CheckQueue MUST return a fresh non-null diagnostic slice sorted by path, code, severity and message, with nonempty corrective messages and error/warning severity. It MUST NOT modify or return mutable aliases of input snapshots, modules, changes, archives, configuration, diagnostics or nested documents/maps/slices, on valid or invalid input. It MUST NOT read/write files, execute hooks/agents/Git or use networking. An empty queue MUST return an empty slice regardless of unrelated input discovery/static diagnostics. A valid queue MUST return an empty slice.

## Decisions

[discovery-validation]
This slice validates the current static graph. Queued projection, set application, hash bases and late ID assignment are separate changes. Discovery scans all repository directories except .git and .fireweed to report out-of-glob specs; WalkDir does not traverse directory symlinks. Symlink file paths are checked by gitstore.SafePath before reads. Configured globs include module discovery, defaults, protected and worker_writable patterns. Module metadata is validated separately from its Specs body so malformed normative content receives C04. Unknown archived target modules are diagnosed C05, rather than silently dropping history. Hook strings remain opaque and are never evaluated by static checks. Module path attribution is <dir>/SPEC.md; queue/archive attribution is its repository-relative path. A root module directory is represented by '.'.

Variable-width numeric-order fixtures use at least four digits, as required by the retained document contract. Original documents retain all valid known metadata and arbitrary human sections; unknown metadata is rejected by contracts. Adjacent stars inside a segment are ordinary wildcards; invalid-pattern tests use malformed classes, absolute/traversal paths and repeated separators. The superseded first cycle was fully reverted after these fixture mistakes surfaced (author decision D008).

[queued-projection]
Projection is an intent preview: it does not enforce existing bases or allocate real IDs. Later queue validation reports missing/stale bases, invalid graph/zones after set, and archive reservations; lifecycle allocation rewrites handles only at start. ApplyChange uses the typed Module and Change fields as authoritative, updating retained documents accordingly. Local handles may appear in Added and references, with the same uniqueness/locality rules as the public change parser; file-backed current module specs and Modified/Removed targets use real IDs. Projected intermediate modules may contain placeholders from preceding changes. Field/type and spec-edit validation is required even for callers constructing typed values directly. An empty module Document is synthesized from typed fields; valid original documents retain preamble and human sections. Project keeps the original pending changes/archive lists and diagnostics as independent copies, preserving their original ordering; only the order of application is sorted. Diagnostics, canonical module ordering and graph validation remain separate existing Check responsibilities. Decisions application/removal will be a separately specified part of archival application; this preview retains existing module Decisions.

ApplyChange errors return the zero Module, never a partial result. Existing-ID addition collisions mean real allocated IDs; handle placeholders are local declarations, so two queued changes may each declare new:local and their preview retains both items without treating that shared spelling as an allocated-ID collision. References in each change must still be declared locally. A preview document can therefore contain repeated placeholder headings from different changes, and is never written as a current SPEC.md until start has allocated real IDs. Invalid excluded target/tail applications are not executed or validated by Project(before); ambiguity of the selector itself is still an error. Base hashes follow typed spec fields, including normalized whitespace, rather than stale retained document text.

The public diagnostic conversion keeps the established non-null slice convention; diagnostic contents and ordering are preserved. Projection otherwise preserves nil versus explicit empty configuration/module lists through independent copies.

[queue-validation]
CheckQueue complements current-state Check and never repeats its input diagnostics. It checks transitions incrementally: newly introduced static diagnostics are compared with the previous projected state's diagnostic signatures; each new problem is attributed to its change. Stale-base errors do not erase the author's intent. Failed structural application skips that change's intent and continues with other queued changes. Numeric/slug application ordering follows typed fields. Assigned state means both nonempty run and phase; normal workflow/history enforcement proves ownership later. Metadata/body validation still uses the public change grammar, including its current real-ID-only fixes/already_satisfied fields. Decisions application, late allocation and all-occurrence rewriting remain separate upcoming changes.

Archived non-edit human sections are historical prose, not live handle bindings. Validate archive metadata and Added/Modified/Removed with public change grammar using a validation-only view whose other bodies are neutral, then retain the original complete parsed document. Queued documents are validated in full. This permits worker Decisions to explain handle examples after allocation without modifying or losing immutable archive history. Author decision D010 records the live regression.
