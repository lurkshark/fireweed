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

## Decisions
