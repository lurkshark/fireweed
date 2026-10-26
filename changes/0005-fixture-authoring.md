---
bases: {}
module: project
set:
    interface:
        - api.go
        - model/**
---



## Motivation

The manual lifecycle fixture needs to load its declared modules and create an
author change before IDs are assigned. Implement this narrow project entry point
now; projection, stamping, additional static checks and lifecycle follow as the
NEW CLI is connected. Do not complete the entire project subsystem first.

## Interface

Expose retained data-only Snapshot (Root, Config, Modules) and Module (Name,
IDPrefix, Dir, DependsOn, Zones, Document) types through public project aliases.
Config is the decoded string-keyed YAML mapping, Zones and Document use public
contracts types, and module Dir is repository-relative. Expose Load(root string)
(Snapshot, error), ValidateGraph(modules []Module) error and NewChange(root,
module, slug string) (relativePath string, err error). Keep signatures/data only
under model/ and thin wrappers in api.go; behavior belongs under internal/.

## Added

### new:load-fixture
Load MUST read fireweed.yaml and module SPEC.md files only in configured modules globs, returning module names, prefixes, dependencies, documents and zones. Omitted module zone fields MUST inherit module_defaults while explicit empty lists remain empty. Missing or malformed config/module documents MUST return an error.

### new:graph-validation
ValidateGraph MUST reject empty or duplicate module names/prefixes and dependencies naming absent modules, self-dependencies or cycles. It MUST accept an acyclic graph with unique identities and return an explanatory error on rejection.

### new:author-change
NewChange MUST create a queued Markdown skeleton for an existing module with Motivation, Interface, Added, Modified, Removed and Decisions sections, containing no assigned spec IDs. It MUST use the next number after the largest pending or archived change number and reject an existing pending/archived slug without overwriting files. Unknown modules and slugs containing anything besides lowercase letters, digits and hyphens MUST be rejected.

## Modified



## Removed



## Decisions
