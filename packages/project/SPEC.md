---
depends_on:
    - contracts
    - gitstore
    - assets
id_prefix: PRJ
interface:
    - api.go
    - model/**
name: project
---

# project

## Specs

### PRJ-001
Load MUST read fireweed.yaml and module SPEC.md files only in configured modules globs, returning module names, prefixes, dependencies, documents and zones.

### PRJ-002
ValidateGraph MUST return nil for a valid graph and an explanatory error identifying the rejection cause for an invalid graph.

### PRJ-003
NewChange MUST create a queued Markdown skeleton for an existing module with Motivation, Interface, Added, Modified, Removed and Decisions sections, containing no assigned spec IDs.

### PRJ-004
Project and Stamp MUST resolve a change argument from its unique slug, queued basename or repository-relative queued path, rejecting missing or ambiguous targets.

### PRJ-005
Project MUST preserve complete Added and Modified item text and remove Removed items when applying pending spec edits.

### PRJ-006
Stamp MUST write each Modified or Removed item's canonical SpecHash from the projection preceding the target change into its bases.

### PRJ-007
Check MUST return no error diagnostics for valid module specs and changes. Each failure diagnostic MUST have a nonempty code and message.

### PRJ-008
Check MUST reject a pending change without a nonempty Motivation section with diagnostic C07.

### PRJ-009
Check MUST reject a module spec or Added/Modified spec whose body contains no normative keyword with diagnostic C04.

### PRJ-010
Check MUST reject spec tags other than review or quarantine with a nonempty reason with diagnostic C06.

### PRJ-011
Check MUST reject duplicate spec IDs within a module's Specs section with diagnostic C05.

### PRJ-012
Check MUST reject module spec IDs and change edit IDs whose prefix belongs to another module with diagnostic C05.

### PRJ-013
Check MUST reject a change that uses the same edit identity more than once across Added, Modified and Removed with diagnostic C05.

### PRJ-014
Check MUST reject a pending change that references a handle absent from that change's Added entries with diagnostic C09.

### PRJ-015
Check MUST reject Modified or Removed entries whose bases hash is missing or differs from the preceding projected spec's canonical hash with diagnostic C08.

### PRJ-016
Check MUST reject Modified or Removed IDs absent from the projection preceding the change with diagnostic C08.

### PRJ-017
Load MUST inherit an omitted module zone field from module_defaults while retaining an explicitly empty module zone list as empty.

### PRJ-018
Load MUST return an error for missing or malformed config or module documents.

### PRJ-019
ValidateGraph MUST reject module names that are empty or duplicated in the graph.

### PRJ-020
ValidateGraph MUST reject module prefixes that are empty or duplicated in the graph.

### PRJ-021
ValidateGraph MUST reject a dependency that does not name a different existing module.

### PRJ-022
ValidateGraph MUST reject a graph containing a dependency cycle.

### PRJ-023
NewChange MUST use the next number after the largest pending or archived change number.

### PRJ-024
NewChange MUST reject a slug already present in the pending queue or archive without overwriting existing files.

### PRJ-025
NewChange MUST reject an unknown target module.

### PRJ-026
NewChange MUST reject an empty slug or a slug containing characters other than lowercase letters, digits and hyphens.

### PRJ-027
Queue MUST return pending NNNN-slug.md changes in numeric-then-slug order, allowing tied numbers and excluding archives.

### PRJ-028
Queue MUST return each pending change's module, frontmatter and parsed Added, Modified and Removed items.

### PRJ-029
Queue MUST reject a malformed pending change document.

### PRJ-030
Queue MUST reject a pending change naming an unknown module.

### PRJ-031
Queue MUST reject pending changes that share a slug.

### PRJ-032
Project MUST apply the queue through the named change inclusively, or the whole queue when through is empty.

### PRJ-033
Project MUST apply set values for depends_on, interface, evals, disposable, preserve, append_only, boundaries_opt_out and hooks to the projected module.

### PRJ-034
Project MUST reject identity or unknown fields in set.

### PRJ-035
Project MUST reject an invalid projected module graph.

### PRJ-036
Project MUST keep Added handles local to their declaring change, allowing different changes to reuse handle text.

### PRJ-037
Stamp MUST reject Modified or Removed IDs absent from the projection preceding the target change.

### PRJ-038
Stamp MUST retain other frontmatter values and human section contents.

### PRJ-039
Stamp MUST succeed on an Added-only change without assigning spec IDs.

### PRJ-040
Check MUST report configuration discovery or decoding failures with diagnostic C01.

### PRJ-041
Check MUST report invalid current or projected module graphs with diagnostic C02.

### PRJ-042
Check MUST report unreadable or malformed module/change documents and malformed item syntax with diagnostic C03, retaining C05 for duplicate identities.

### PRJ-043
Check MUST report invalid module metadata or zone field types, and invalid set keys or value types, with diagnostic C12.

### PRJ-044
Check MUST report pending queue identity errors with diagnostic C13.

### PRJ-045
Check MUST report invalid pending change module targets with diagnostic C07.

### PRJ-046
Check MUST report error C07 for a pending change containing eval_fix together
with fixes or any Added, Modified or Removed spec edits.

## Decisions
