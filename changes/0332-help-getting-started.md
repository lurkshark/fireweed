---
bases:
    AST-008: 9ce0d11f
module: assets
phase: impl
run: 20261120T204331-fef9419b1f49f900fe2dfbbb
---

## Motivation

When someone installs Fireweed and asks a coding agent to set it up in an empty
Git repository, the top-level help should tell the agent what Fireweed does and
how to finish setup. The current overview lists `init` and `doctor`, but leaves
the sequence implicit. In a fresh Go repository, `init` and `check` succeeded,
but `doctor --no-agent` could not create its scratch worktree until the first
commit existed. The README already gives the setup sequence.

## Interface

Extend the overview returned by `Help()`.

## Added

### AST-077
Help MUST briefly describe Fireweed as a CLI for spec-driven development in Git.
It MUST explain that coding agents write evals and implementation in separate
steps that Fireweed verifies and records in Git.

### AST-078
Help MUST direct users starting a Git project to `fireweed help init` to choose
a recipe and worker. It MUST instruct them to run `fireweed init`, follow the
printed recipe instructions, and describe the project in `SYSTEM.md`. It MUST
direct them to commit the setup before running `fireweed doctor`.

## Modified

### AST-008
Help MUST identify fireweed and list every command in the design's CLI table,
including help, version, explain and playbook, grouped by purpose. It MUST
advertise help <command> and <command> --help, using `fireweed help init` as
the concrete example. It MUST keep command-specific option and configuration
descriptions on the command pages rather than the overview.

## Removed

## Decisions

## Notes

The targeted evals fail on the current overview for AST-008, AST-077 and AST-078.
The overview lives in the implementation-writable `packages/assets/internal/content/content.go`.
## Audit

No findings; checked AST-008, AST-077, AST-078

## Provenance

- evals 2026-11-20T20:48:07.964843Z agent=codex fireweed=0.2.4 playbook=5e05bacdae954379b468564c11429c5532016e26a7afbcd418f9bdb9441e151c tests=162 failing=3 passing=159
- audit 2026-11-20T20:49:35.818791Z agent=codex fireweed=0.2.4 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0

