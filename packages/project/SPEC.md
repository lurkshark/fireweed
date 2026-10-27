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
Load MUST read fireweed.yaml and module SPEC.md files only in configured modules globs, returning module names, prefixes, dependencies, documents and zones. Omitted module zone fields MUST inherit module_defaults while explicit empty lists remain empty. Missing or malformed config/module documents MUST return an error.

### PRJ-002
ValidateGraph MUST reject empty or duplicate module names/prefixes and dependencies naming absent modules, self-dependencies or cycles. It MUST accept an acyclic graph with unique identities and return an explanatory error on rejection.

### PRJ-003
NewChange MUST create a queued Markdown skeleton for an existing module with Motivation, Interface, Added, Modified, Removed and Decisions sections, containing no assigned spec IDs. It MUST use the next number after the largest pending or archived change number and reject an existing pending/archived slug without overwriting files. Unknown modules and slugs containing anything besides lowercase letters, digits and hyphens MUST be rejected.

### PRJ-004
Queue MUST read pending NNNN-slug.md documents in numeric-then-slug order, allowing tied numbers and excluding archives. It MUST return each change's module, frontmatter and parsed Added/Modified/Removed items, rejecting malformed documents, unknown modules and ambiguous slugs. A change argument MUST resolve its unique slug, queued basename or repository-relative queued path, rejecting missing or ambiguous targets.

### PRJ-005
Project MUST apply pending changes through the named change inclusively, or the whole queue when through is empty, preserving Added/Modified full text and removing Removed items. It MUST apply set values for depends_on, interface, evals, disposable, preserve, append_only, boundaries_opt_out and hooks, rejecting identity/unknown fields and an invalid projected graph. Added handles MUST remain local to their declaring change; different changes may reuse handle text.

### PRJ-006
Stamp MUST write the canonical SpecHash of each Modified/Removed item from the projection preceding the target change into its bases. It MUST reject IDs absent from that projection and preserve other frontmatter and human sections. Stamping an Added-only change MUST succeed without assigning IDs.

### PRJ-007
Check MUST reject missing Motivation, non-normative spec bodies, invalid review/quarantine tags, duplicate or foreign module IDs, duplicate edit identities within a change and references to undeclared handles. It MUST reject missing/stale bases and Modified/Removed IDs absent from the preceding projection. Check results MUST have no error diagnostics for valid module specs/changes, and nonempty C01 code/message diagnostics for failures.

## Decisions
