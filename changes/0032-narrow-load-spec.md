---
already_satisfied:
    - PRJ-001
    - new:load-zone-defaults
    - new:load-invalid-documents
bases:
    PRJ-001: 47df5144
module: project
---



## Motivation

The approved project spec-split slice applies SYSTEM's one-behavior spec shape
to PRJ-001. Preserve the design's Modules and zone map defaults and
Configuration requirements while separating discovery, zone default resolution
and document rejection. Existing public evals already assert these behaviors.

## Interface



## Added

### new:load-zone-defaults
Load MUST inherit an omitted module zone field from module_defaults while retaining an explicitly empty module zone list as empty.

### new:load-invalid-documents
Load MUST return an error for missing or malformed config or module documents.

## Modified

### PRJ-001
Load MUST read fireweed.yaml and module SPEC.md files only in configured modules globs, returning module names, prefixes, dependencies, documents and zones.

## Removed



## Decisions



## Notes

Retarget existing tests only; retain every assertion and 53 reported project
cases. LoadConfiguredModulesAndDefaults retains PRJ-001 and gains the defaults
ID. LoadInvalidDocuments and its existing four cases use the invalid-documents
ID rather than PRJ-001. No new Test functions, t.Run cases, fixture instances,
helpers, product code or interface edits are needed. Implementation may edit
only this change file. Keep all other tests and IDs unchanged.
