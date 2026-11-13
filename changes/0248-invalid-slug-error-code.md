---
already_satisfied:
    - PRJ-026
bases:
    PRJ-026: 8c0db16f
module: project
phase: audit
run: 20261113T165543-e4e620ae1922a184531749ac
---

## Motivation

The owner-approved review of regenerate-project found that NewChange rejected
an invalid slug while the dependent CLI returned P01 instead of APP-019's
required U01. PRJ-026 specifies rejection but not its code, so project's own
evals accepted behavior that broke the CLI. Make the project API's existing
U01 behavior explicit under SYSTEM.md's coded-error convention.

## Interface

## Added

## Modified

### PRJ-026
NewChange MUST reject an empty slug or a slug containing characters other than lowercase letters, digits and hyphens with an error whose contracts.ErrorCode is U01.

## Removed

## Decisions

## Notes

No signature or type changes. Clarify the code returned by NewChange for an
invalid slug.

The current implementation already returns U01 for this condition, so
PRJ-026 is legitimately already_satisfied.

Strengthen the existing PRJ-026 cases in
packages/project/evals/project_test.go to assert contracts.ErrorCode as well
as rejection. Keep the existing invalid-slug inputs and PRJ-025 assertions.
Do not prescribe diagnostic wording or change valid-slug grammar.

Drafted and queued only. No lifecycle execution is authorized by the drafting
request.
## Provenance

- evals 2026-11-13T16:56:33.209669Z agent=codex fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=156 failing=0 passing=156

