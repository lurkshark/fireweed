---
audit_pending:
    - PRJ-002
    - PRJ-003
audit_rounds: 1
bases: {}
module: project
phase: impl
run: 20261026T235149-fb041978fac8086ee1d5472c
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

### PRJ-001
Load MUST read fireweed.yaml and module SPEC.md files only in configured modules globs, returning module names, prefixes, dependencies, documents and zones. Omitted module zone fields MUST inherit module_defaults while explicit empty lists remain empty. Missing or malformed config/module documents MUST return an error.

### PRJ-002
ValidateGraph MUST reject empty or duplicate module names/prefixes and dependencies naming absent modules, self-dependencies or cycles. It MUST accept an acyclic graph with unique identities and return an explanatory error on rejection.

### PRJ-003
NewChange MUST create a queued Markdown skeleton for an existing module with Motivation, Interface, Added, Modified, Removed and Decisions sections, containing no assigned spec IDs. It MUST use the next number after the largest pending or archived change number and reject an existing pending/archived slug without overwriting files. Unknown modules and slugs containing anything besides lowercase letters, digits and hyphens MUST be rejected.

## Modified



## Removed



## Decisions



## Provenance

- evals 2026-10-26T23:54:26.040665Z agent=manual fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=51 failing=24 passing=27
- audit 2026-10-26T23:55:46.613665Z agent=manual fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0
- evals 2026-10-26T23:57:51.611438Z agent=manual fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=51 failing=24 passing=27
- audit 2026-10-26T23:58:58.726249Z agent=manual fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0

## Audit

No findings. Checked PRJ-002, PRJ-003.
