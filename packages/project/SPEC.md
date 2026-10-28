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
Project MUST apply pending changes through the named change inclusively, or the whole queue when through is empty, preserving Added/Modified full text and removing Removed items. It MUST apply set values for depends_on, interface, evals, disposable, preserve, append_only, boundaries_opt_out and hooks, rejecting identity/unknown fields and an invalid projected graph. Added handles MUST remain local to their declaring change; different changes may reuse handle text.

### PRJ-006
Stamp MUST write the canonical SpecHash of each Modified/Removed item from the projection preceding the target change into its bases. It MUST reject IDs absent from that projection and preserve other frontmatter and human sections. Stamping an Added-only change MUST succeed without assigning IDs.

### PRJ-007
Check MUST return no error diagnostics for valid module specs and changes. Each failure diagnostic MUST have a nonempty code and message.

### PRJ-008
Check MUST reject a pending change without a nonempty Motivation section.

### PRJ-009
Check MUST reject a module spec or Added/Modified spec whose body contains no normative keyword.

### PRJ-010
Check MUST reject spec tags other than review or quarantine with a nonempty reason.

### PRJ-011
Check MUST reject duplicate spec IDs within a module's Specs section.

### PRJ-012
Check MUST reject module spec IDs and change edit IDs whose prefix belongs to another module.

### PRJ-013
Check MUST reject a change that uses the same edit identity more than once across Added, Modified and Removed.

### PRJ-014
Check MUST reject a pending change that references a handle absent from that change's Added entries.

### PRJ-015
Check MUST reject Modified or Removed entries whose bases hash is missing or differs from the preceding projected spec's canonical hash.

### PRJ-016
Check MUST reject Modified or Removed IDs absent from the projection preceding the change.

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

## Decisions
