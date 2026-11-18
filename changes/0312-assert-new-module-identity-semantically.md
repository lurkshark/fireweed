---
eval_fix:
    - APP-101
module: app
---

## Motivation

The November 17 assets drill reported one dependent failure in the app
`new-module` eval. The generated `SPEC.md` declared `empty2-3-mod` with prefix
`E2M`. The eval expected the identity keys on adjacent lines and failed when
`depends_on: []` appeared between them. APP-101 requires the requested
identity and a module that loads in the project graph.

## Interface

## Added

## Modified

## Removed

## Decisions

## Notes

In `Test_APP_101_APP_104_APP_159_NewModuleEmptyGoModule`, load the generated
module with `project.Load` and assert that the discovered module has name
`empty2-3-mod`, ID prefix `E2M`, and directory `packages/empty2-3-mod`. Exercise
the valid frontmatter layout with `depends_on` between the identity keys to
cover the layout seen in the drill. Keep the assertions for the CLI result
and empty `go_module` substitution.
