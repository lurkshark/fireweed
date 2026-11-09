---
already_satisfied:
    - GIT-005
bases:
    GIT-005: cbffa60a
module: gitstore
phase: audit
run: 20261109T203414-bfe0e9ad44fd0c74b5dd3e2d
---

## Motivation

The design's drill treats a reconstruction failure as a sign that the spec
or interface leaves behavior unstated. The last workflow drill rebuilt a
helper that passed Git subcommands such as `ls-tree` to `gitstore.Run` as the
executable. The public interface never says that `Request.Args` is a full
argv whose first element names the program, and Run's name and module
suggest it runs Git. State the existing behavior at the interface.

## Interface

Add a doc comment to `Request.Args` in `packages/gitstore/model` saying it is
the complete argv: `Args[0]` is the executable (resolved through PATH when it
has no slash) and Run never prepends `git`. Add a matching one-line comment to
`Run` in `api.go`. No signature or behavior changes.

## Added

## Modified

### GIT-005
Run MUST launch the complete argv in Args, whose first element names the executable, without prepending any command. It MUST use the requested directory and supplied environment/stdin and return captured stdout, stderr and exit code, returning an error with that output on nonzero exit. Timeout or context cancellation MUST terminate the launched process group and return an error.

## Removed

## Decisions

## Notes

The existing GIT-005 eval already passes `/bin/sh` as `Args[0]`, so the
behavior is covered and stays already satisfied. Add one assertion that an
argv whose first element is a Git subcommand (for example `ls-tree`) fails to
launch rather than running Git. The sentence about processes escaping the
group moves out of the spec because SYSTEM.md already records it as a
documented limit.
## Provenance
- evals 2026-11-09T20:36:00.777265Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=10 failing=0 passing=10
