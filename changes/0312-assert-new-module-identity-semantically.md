---
eval_fix:
    - APP-050
    - APP-101
module: app
---

## Motivation

The November 17 assets drill reported one dependent failure in the app
`new-module` eval. The generated `SPEC.md` declared `empty2-3-mod` with prefix
`E2M`. The eval expected the identity keys on adjacent lines and failed when
`depends_on: []` appeared between them. APP-101 requires the requested
identity and a module that loads in the project graph.

The neighboring `new-module` case also checks adjacent keys even though it
validates the loaded graph. The Go and local `init` cases look for identity
text anywhere in `SPEC.md`. APP-050 requires `init` to create the default
module with the specified identity and no dependencies.

## Interface

## Added

## Modified

## Removed

## Decisions

## Notes

- In `Test_APP_101_APP_104_APP_159_NewModuleEmptyGoModule`, load the project
  with `project.Load` and assert that the discovered module has name
  `empty2-3-mod`, ID prefix `E2M`, and directory `packages/empty2-3-mod`.
  Exercise valid frontmatter with `depends_on` between the identity keys.
  Keep the CLI result and empty `go_module` substitution assertions.
- In `Test_APP_100_APP_101_APP_104_NewModuleFromNearestProjectTemplates`, use
  the existing loaded module and graph assertions for the identity check.
  Keep the `Project-owned SPEC marker` assertion for template selection.
- In the `GoInit` and `LocalInit` cases in `packages/app/evals/init_test.go`,
  parse the generated `SPEC.md` with `contracts.ParseDocument`. Check that
  frontmatter declares name `app` and ID prefix `APP`. In `GoInit`, also
  check that `depends_on` is an empty sequence. Keep the existing template
  substitution and layout assertions.
