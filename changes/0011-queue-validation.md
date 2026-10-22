---
bases:
    PRJ-002: 785bd51b
module: project
phase: audit
run: f6f7a58489768db3ab6b9ece
---

## Motivation

Queue previews now exist, but check must detect stale authored intent and enforce the projected zones and graph from the first step. Preserve independent progress when malformed changes occur and distinguish handle-scope errors from general grammar errors.

## Interface

Add CheckQueue(snapshot Snapshot) []Diagnostic as a thin public facade to disposable pure queue validation. It returns queued issues only; callers combine existing Check results for current/static/discovery diagnostics. Extend Load's malformed-change categorization for undeclared handles and preserve archived human prose as historical text. Preserve existing projection interfaces.

## Added

### PRJ-011
CheckQueue MUST process changes in numeric Number then Slug order against preceding structurally valid intent. Every Modified and Removed ID MUST exist immediately before its change and have a matching normalized public SpecHash in bases; missing targets, missing hashes and stale hashes MUST emit C10 error diagnostics at that change's path. Tag changes MUST affect comparison while normalized line endings/trailing whitespace MUST NOT cause false staleness. Stale or missing hashes MUST NOT stop projecting structurally valid intent for later base checks.

### PRJ-012
CheckQueue MUST apply each structurally valid change's set values before validating graph, zones, hook requirements and archive reservations with the public static contract. Newly introduced C03, C06, C07, C08 and D04 diagnostics MUST be attributed to the introducing change path and retain corrective detail, including opt-out reasons. It MUST NOT repeat unchanged preexisting current/static/discovery diagnostics, and a valid later change that removes an introduced problem MUST pass. Every intermediate projected state MUST be checked, even when a later change repairs it; future-path intersections MUST be enforced without requiring files to exist.

### PRJ-013
Unknown target modules and invalid typed changes/applications MUST emit C05 errors; undeclared or cross-change-only handles anywhere in a change MUST emit C09 errors. Structurally invalid applications MUST be skipped while later independent changes are still checked. Real-ID Added items without run and phase MUST be rejected as author-chosen IDs with C05; a removed archive ID MUST be rejected with C08 even when the change also lacks assigned state. Assigned-state additions with unused valid IDs and local author handles MUST remain accepted. fixes IDs MUST exist before the change and remain present after it; already_satisfied IDs MUST exist in the resulting module. Missing or removed references MUST emit C05.

### PRJ-014
CheckQueue MUST return a fresh non-null diagnostic slice sorted by path, code, severity and message, with nonempty corrective messages and error/warning severity. It MUST NOT modify or return mutable aliases of input snapshots, modules, changes, archives, configuration, diagnostics or nested documents/maps/slices, on valid or invalid input. It MUST NOT read/write files, execute hooks/agents/Git or use networking. An empty queue MUST return an empty slice regardless of unrelated input discovery/static diagnostics. A valid queue MUST return an empty slice.

## Modified

### PRJ-002
Load MUST continue past malformed module/change documents and retain path-attributed diagnostics: C02 warning for SPEC.md outside module globs, C03 error for invalid module metadata, C04 error for invalid Specs grammar, and C05 error for malformed changes or unknown target modules, except that undeclared live handle references MUST produce C09 error diagnostics. Queued handle references anywhere in the change, including human sections and metadata, MUST be local to that change; a declaration in a different change MUST NOT satisfy them. Archive human sections MUST preserve literal handle examples without treating them as live references; archive edit sections and metadata MUST still validate their live references. It MUST ignore .git and .fireweed trees during module discovery. Queue and archive MUST include immediate Markdown files only, parsed with each target's prefix, retaining lifecycle/provenance/author sections; each list MUST sort by numeric filename then slug. Duplicate slugs across either list MUST produce C05. Symlink escapes MUST produce an error diagnostic without reading their contents.

## Removed

## Decisions

CheckQueue complements current-state Check and never repeats its input diagnostics. It checks transitions incrementally: newly introduced static diagnostics are compared with the previous projected state's diagnostic signatures; each new problem is attributed to its change. Stale-base errors do not erase the author's intent. Failed structural application skips that change's intent and continues with other queued changes. Numeric/slug application ordering follows typed fields. Assigned state means both nonempty run and phase; normal workflow/history enforcement proves ownership later. Metadata/body validation still uses the public change grammar, including its current real-ID-only fixes/already_satisfied fields. Decisions application, late allocation and all-occurrence rewriting remain separate upcoming changes.

Archived non-edit human sections are historical prose, not live handle bindings. Validate archive metadata and Added/Modified/Removed with public change grammar using a validation-only view whose other bodies are neutral, then retain the original complete parsed document. Queued documents are validated in full. This permits worker Decisions to explain handle examples after allocation without modifying or losing immutable archive history. Author decision D010 records the live regression.

## Rejected alternatives

Validating only the final state would allow a change to weaken enforcement until another queued change repairs it. Stopping at the first malformed change would hide errors in independent modules. Treating another change's identically named handle as a declaration would break the design's local namespace.

## Provenance

- evals | 2026-10-22T14:12:18Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 51 | failing: 11 | passing: 40
- audit | 2026-10-22T14:18:05Z | agent: manual | fireweed: 0.4.0 | playbook: a9c982d24b476621316049b527e6c7bbc319f06b12268af150fa7d8c6df92230 | tests: 0 | failing: 0 | passing: 0
- evals | 2026-10-22T14:22:31Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 57 | failing: 16 | passing: 41

## Audit

- PRJ-011: "Every Modified and Removed ID MUST exist immediately before its change" — missing target fixtures also omit their bases; provide hashes and separately test stale Removed hashes.
- PRJ-011: "Stale or missing hashes MUST NOT stop projecting structurally valid intent" — observe removal intent in a later change, not only modification bodies.
- PRJ-012: "Newly introduced C03, C06, C07, C08 and D04 diagnostics" — test persistence through an intervening change, repair/reintroduction, and a new same-code problem in another module.
- PRJ-013: "invalid typed changes/applications MUST emit C05" — add invalid set types/fields, phase/hash format, and valid native Motivation documents.
- PRJ-013: "Structurally invalid applications MUST be skipped" — include valid modification/set plus invalid edit and check later same-module bases against the original state.
- PRJ-013: "already_satisfied IDs MUST exist in the resulting module" — add acceptance of newly added assigned-state IDs.
- PRJ-013: "undeclared or cross-change-only handles anywhere in a change MUST emit C09" — add nested metadata local acceptance versus undeclared references.
- PRJ-002: "Archive human sections MUST preserve literal handle examples" — add historical preamble/Motivation/Provenance literals and full retained document assertions.
- PRJ-014: "fresh non-null diagnostic slice" — poison an error slice then repeat; compare valid populated snapshots before/after and seed multiple unrelated discovery codes.
