---
module: workflow
phase: applied
run: 20261111T080643-a3e51a12ee3e891cbddc2b19
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
## Provenance
- impl 2026-11-11T08:35:02.069937Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=5f856d6bcab308b9ff1a8a7c03b9d2547744112aba81dd30dea8f61ec8062b03 tests=672 failing=0 passing=672
