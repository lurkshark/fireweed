---
already_satisfied:
    - GIT-012
module: gitstore
phase: audit
run: 20261118T213424-4a1356b2421b1b6c62088a18
---

## Motivation

The November 18 gitstore reconstruction treated Request.Env as a replacement
environment, removing PATH from invoked hooks. Applying it as overrides restored all
52 dependent cases. Hook and agent callers supply settings such as FIREWEED_BIN and
rely on the launched process retaining their toolchain environment.
[Drill reports](https://github.com/lurkshark/fireweed/pull/25)

## Interface

Specify the existing gitstore.Run semantics for Request.Env in the launched process.

## Added

### GIT-012
Run MUST inherit the caller's environment and apply Request.Env as per-variable
overrides to the launched process. Omitted keys MUST retain their inherited values.

## Modified

## Removed

## Decisions

## Notes

Use t.Setenv and fixture-local values. Launch an absolute fixture executable so
the eval observes the child environment independently of executable lookup. Assert
that the overridden value wins and an unrelated inherited sentinel survives. With
nil and empty override maps, assert that both parent sentinel values survive. The
current launcher already meets these cases.
## Provenance

- evals 2026-11-18T21:35:24.379885Z agent=codex fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=19 failing=0 passing=19

