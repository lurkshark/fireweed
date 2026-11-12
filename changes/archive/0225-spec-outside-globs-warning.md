---
module: project
phase: applied
run: 20261111T201023-bc783aa29bfe494ba1189352
---

## Motivation

The design says fireweed reads SPEC.md files only under the `modules`
globs, "and `check` warns about any `SPEC.md` outside them". Nothing
implements that warning (project drill gap for PRJ-001). A module whose
directory a glob misses, for example after a typo or a move, silently drops
out of check, coverage and boundaries.

## Interface

None. Check returns one more warning diagnostic.

## Added

### PRJ-070
Check MUST return a C20 warning naming each tracked file named SPEC.md that is neither a configured module's SPEC.md nor inside a configured module directory or `.fireweed/`.

## Modified

## Removed

## Decisions

## Notes

"Tracked" means in the Git index, so ignored and untracked directories such
as `node_modules` are skipped. The exclusions keep project templates
(`.fireweed/templates/module/SPEC.md`) and files inside a module, such as
the recipe templates under `packages/assets/internal/recipe/files/`, from
warning. This repository has no other SPEC.md outside its globs, so its own
`check` stays clean.

Run after 0224, which adds the C20 explanation.
## Provenance
- evals 2026-11-11T20:13:16.612468Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=135 failing=1 passing=134
- audit 2026-11-11T20:13:34.34656Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-astra -c model_reasoning_effort="medium" -s workspace-write {prompt} fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
- impl 2026-11-11T20:51:50.361125Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=5f856d6bcab308b9ff1a8a7c03b9d2547744112aba81dd30dea8f61ec8062b03 tests=941 failing=0 passing=941
## Audit

No findings; checked PRJ-070
