---
already_satisfied:
    - PRJ-001
bases:
    PRJ-001: e686ba9c
module: project
phase: audit
run: 20261116T220110-f6e6a9cec4c80a19fa611903
---

## Motivation

APP-100, APP-101 and APP-103 require new-module to load and validate rendered
module templates before creation. The CLI uses a configuration with a dot
module glob and a SPEC.md at its root for that validation. The retained
project build skipped the root even when explicitly selected, producing an
empty snapshot and a panic in all three affected template cases.

## Interface

## Added

## Modified

### PRJ-001
Load MUST read fireweed.yaml and module SPEC.md files only in configured
modules globs, returning module names, prefixes, dependencies, documents and
zones. An explicit `.` entry MUST select the project root as a module
directory and return it as `.` when the root SPEC.md is valid.

## Removed

## Decisions

## Notes

Use a public Load fixture with modules: ["."] and a valid root SPEC.md.
Assert the root module and relative dir ".". Keep existing discovery and
malformed-module cases. The APP-100/101/104 custom template, APP-103/104
malformed-template subcase and APP-101/104/159 empty-go_module case all reach
the same temporary-root load. Current Load handles this input.
## Provenance

- evals 2026-11-16T22:02:00.536644Z agent=codex fireweed=0.2.1 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=318 failing=0 passing=318

