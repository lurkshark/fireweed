---
bases:
    AST-008: 9ce0d11f
module: assets
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

### new:help-purpose
Help MUST briefly describe Fireweed as a CLI for spec-driven development in Git.
It MUST explain that coding agents write evals and implementation in separate
steps that Fireweed verifies and records in Git.

### new:help-setup-sequence
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
