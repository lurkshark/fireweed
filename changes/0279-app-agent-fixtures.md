---
eval_fix:
    - APP-063
    - APP-074
    - APP-080
    - APP-136
module: app
phase: impl
run: 20261115T014840-e413ed507a5c4fa629a33204
---

## Motivation

The CLI integration fixtures encode agent commands inline. Runner and checks
will switch to named commands in separate changes, and their implementation
verification runs the app evals. Give those fixtures equivalent named entries
before either dependency switches its command lookup.

## Interface

## Added

## Modified

## Removed

## Decisions

## Notes

In the doctor, next, run, and drill CLI fixtures, define named entries for the
same fake commands and retain the inline entries until both dependencies have
switched. For doctor, prepare these entries in the fixture config callback.
Keep the existing CLI assertions, fake commands, and fixture bounds. The
app CLI change removes the temporary inline fixture entries.
## Audit

No findings; checked APP-063, APP-074, APP-080, APP-136

## Provenance

- evals 2026-11-15T01:57:05.172063Z agent=codex fireweed=0.2.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=395 failing=0 passing=395
- audit 2026-11-15T01:57:23.276715Z agent=codex fireweed=0.2.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0

