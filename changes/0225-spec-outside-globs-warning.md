---
module: project
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

### new:spec-outside-globs
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
