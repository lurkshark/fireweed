---
module: contracts
phase: audit
run: 29ab73909c7b96c1b7b1ff0c
---

## Motivation

The new product must read and later project module specs and queued changes without losing human sections, confusing fenced examples with structure, or treating omissions as explicit empty overrides.

## Interface

Add retained Document {Metadata map[string]any; Preamble string; Sections []Section}, Section {Name, Body string}, Module {Name, IDPrefix, Dir string; DependsOn []string; Zones; Hooks; BoundariesOptOut string; Specs []Spec; Decisions string; Document Document}, Change {Module, Path, Slug, Kind, Run, Phase, Blocked string; Number int; Set map[string]any; Bases map[string]string; Fixes, AlreadySatisfied []string; Regenerate bool; Added, Modified, Removed []Spec; Document Document}. Move the unchanged Spec shape to model with snake_case wire tags and alias it through api.go. Expose ParseMarkdown(text string) (Document,error), RenderMarkdown(Document) (string,error), ParseModule(text,dir string,defaults Zones,hooks Hooks) (Module,error), and ParseChange(text,path,prefix string) (Change,error). Interfaces are values; parsing performs no I/O.

## Added

### CON-012
Markdown document parsing MUST require an initial closed YAML frontmatter mapping, reject malformed or duplicate metadata keys and duplicate live level-two section names, and retain preamble, section order and bodies including fenced examples. Render followed by parse MUST preserve metadata values and human content. CRLF MUST be accepted. Fenced headings MUST NOT become sections.

### CON-013
Module parsing MUST require valid name and id_prefix, exactly one Specs section, valid normative spec items, and typed known metadata fields. Omitted zone lists and hooks MUST inherit supplied project defaults; explicit empty lists and explicitly empty hook strings MUST override them. Depends_on, boundaries_opt_out and Decisions MUST be retained. Parsing MUST NOT mutate or share inherited slices with the caller or other parsed modules.

### CON-014
Change parsing MUST require typed known metadata, a module, nonempty Motivation and a numbered lowercase slug filename under changes/ or changes/archive/. It MUST retain run, phase, blocked, bases, fixes, already_satisfied and regenerate; phase MUST be empty, evals, audit, impl or applied. Kind MUST be regenerate when requested, otherwise spec when any spec edits exist, fix when fixes exist, and maintenance otherwise. Invalid filenames and phases MUST be rejected.

### CON-015
Added and Modified MUST contain full normative spec items; Removed MUST contain headings only. Local handles MUST be unique across Added, MUST NOT appear as Modified or Removed IDs, and every handle reference in the change MUST refer to one declared in its Added section. An ID MUST NOT appear in more than one edit section. Empty spec-edit sections MUST be accepted for maintenance. All author prose, Decisions, Removed decisions, Questions, Audit and Provenance MUST remain available in the parsed document.

### CON-016
Change set metadata MUST allow only depends_on, interface, evals, disposable, preserve, append_only, boundaries_opt_out and per-module hooks. It MUST reject name/id_prefix identity edits, unknown fields and incorrectly typed values, and retain explicit empty lists or hook strings. Bases MUST contain eight lowercase hexadecimal hashes keyed by real IDs; fixes and already_satisfied MUST contain real IDs rather than handles.

## Modified

## Removed

## Decisions

Module names and change slugs start with a lowercase letter or digit and continue with lowercase letters, digits or hyphens; prefixes start with an uppercase letter and continue with uppercase letters or digits. Filenames use at least four decimal digits and are parsed from repository-relative changes paths, not shorthand arguments. Required metadata may not be null. Module depends_on defaults to an empty list. Explicit module hooks override only named hook values. The document representation normalizes line endings and insignificant outer blank lines; it does not preserve byte-level YAML formatting. Queued Added may contain handles or already assigned IDs; the workflow, not this parser, enforces late assignment. Regenerate classification does not suppress spec edits. Cross-change stale bases, queue ordering and ID allocation are validated by project/workflow. Unknown human section names are retained; duplicate names are rejected to avoid ambiguous projections.

## Rejected alternatives

A generic mutable global document model would risk leaking inherited slices. Parsing Removed through the normative body rule would reject its designed heading-only grammar.

## Provenance

- evals | 2026-10-22T06:49:43Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 20 | failing: 5 | passing: 15
- audit | 2026-10-22T06:53:09Z | agent: manual | fireweed: 0.4.0 | playbook: a9c982d24b476621316049b527e6c7bbc319f06b12268af150fa7d8c6df92230 | tests: 0 | failing: 0 | passing: 0
- evals | 2026-10-22T06:53:17Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 25 | failing: 9 | passing: 16

## Audit

- CON-012: "Render followed by parse MUST preserve metadata values and human content." — round trip starts with parser output; independently supplied nested metadata values are not asserted.
- CON-013: "valid name and id_prefix" — malformed string prefix such as lowercase auth is not rejected by an assertion.
- CON-013: "typed known metadata fields" — non-string list elements are not asserted invalid.
- CON-014: "a module" — explicitly empty or null required module values are not asserted invalid.
- CON-014: "otherwise spec when any spec edits exist, fix when fixes exist" — spec edits taking precedence over fixes is not asserted.
- CON-014: "Kind MUST be regenerate when requested" — retaining parsed spec edits when regenerating is not asserted.
- CON-015: "Added and Modified MUST contain full normative spec items" — a nonempty Modified body without normative language is not asserted invalid.
- CON-015: "every handle reference in the change MUST refer to one declared in its Added section" — a dangling reference within a spec body is not asserted invalid.
- CON-015: "All author prose" — change preamble and prose before spec items are not asserted retained.
- CON-016: "retain explicit empty lists or hook strings" — set.depends_on empty list retention is not asserted.
- CON-016: "eight lowercase hexadecimal hashes" — eight lowercase nonhexadecimal characters are not asserted invalid.
- CON-016: "keyed by real IDs; fixes and already_satisfied MUST contain real IDs rather than handles" — malformed non-handle IDs are not asserted invalid.
