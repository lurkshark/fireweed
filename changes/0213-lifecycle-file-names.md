---
module: workflow
---

## Motivation

The evals playbook tells workers that stubs are temporary behavior that
implementation replaces (AST-043, AST-044). In workflow's lifecycle
package, eleven files named `stub.go` or `*_stub.go` hold about 2,800 lines
of the real implementation, including Start, Current, Verify, Done and
Context. A worker following the playbook can take them for disposable
placeholders, and reviewers cannot tell placeholder code from finished
code by file name.

## Interface

None. File names in the disposable zone only.

## Added

## Modified

## Removed

## Decisions

## Notes

In `packages/workflow/internal/lifecycle/`, rename each `*_stub.go` file to
the name without `_stub` (for example `abandon_stub.go` to `abandon.go`,
`status_stub.go` to `status.go`) and rename `stub.go` to `lifecycle.go`.
Splitting `lifecycle.go` by operation is allowed if it stays a pure move.
Change no identifiers or behavior, and keep the `new eval-step stub` label
in the PR body, which names a category of eval-step files rather than these
files.

No spec, eval or other module refers to these file names, so all evals
must pass unchanged and the workflow drill proof stays valid.
