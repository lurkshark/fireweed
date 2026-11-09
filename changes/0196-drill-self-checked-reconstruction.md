---
bases:
    RUN-053: fe3af3c1
    RUN-054: 533417ee
    RUN-055: dca82c6c
module: runner
---

## Motivation

The design's drill has the agent rebuild a module's disposable files "from
the spec, Decisions, interface and evals" in a scratch worktree and then
runs its own and its dependents' evals. A regenerate change does the same
work and lets the agent run its evals while it works. The current drill
instead gives the worker only interface files from its dependencies and no
root manifests or hooks, so nothing it writes can compile or be tested
before the single outside verification. Repeated workflow drills failed on
compile errors and simple slips that one local eval run would have caught,
which says nothing about whether the spec is sufficient.

Make the input safe by construction instead of by restricting the worker:
leave out the old implementation, its history and the dependents whose
evals are the drill's independent signal, and let the worker run its own
eval hook as often as it needs.

## Interface

No API or configuration changes. The input workspace's contents, the
reconstruction prompt and submission validation change as specified.

## Added

## Modified

### RUN-053
Drill MUST invoke reconstruction in a separate input workspace holding the invoking HEAD's tracked files except the target's disposable files, every direct and transitive dependent module's files, the changes directory and drill reports. Its Git metadata MUST hold one fresh commit and none of the invoking repository's history, refs or saved attempts. Drill MUST run the configured setup hook there before agent invocation.

### RUN-054
The reconstruction prompt MUST include the target spec and Decisions, identify its retained interfaces and evals, and require rebuilding only disposable paths without reading outside the input workspace. It MUST give a command that runs the target's eval hook in that workspace. It MUST request separate Contradicted Decisions and Gaps sections in DRILL_NOTES.md, each gap naming the nearest spec ID or "no spec" and a target module.

### RUN-055
Drill MUST reject an attempt that modified or deleted any input file present after setup, before importing code. It MUST import only the target's rebuilt disposable files into the verification worktree, ignoring other files the worker or its eval runs created, without changing retained files or the invoking implementation.

## Removed

## Decisions

## Notes

Dependency implementations stay in the input so the target compiles and
its evals run; they are not what the drill tests, and the boundaries hook
already limits imports to declared interfaces. Dependents are removed from
the input, so a worker cannot fit its code to their evals; the verification
worktree still runs every dependent as RUN-059 requires. Sibling modules
that are neither dependencies nor dependents may stay.

Build the input by copying the verification worktree's tracked files after
disposable removal, minus the excluded paths, then create the single fresh
commit. Take the retained-file inventory after the setup hook runs so setup
output is part of the baseline. In the prompt, expand the target's eval
hook with `{module}`, `{module_dir}` and a `{report}` path inside the input
workspace, for example `.fireweed/runs/drill/<module>.xml`; files like that
report are ignored on import. Keep the existing usage-limit, retry, report,
ref and cleanup behavior.

Update the existing restricted-input and submission fixtures. The input
now contains root files such as `fireweed.yaml` and the hooks, plus
dependency implementation, and must not contain the target's old
implementation, any dependent module, `changes/` or `.fireweed/drills/`, or
the invoking history. Have one fake worker run the advertised eval command
inside the input and use its result. A stray new file outside the
disposable zone is ignored and not imported; a modified or deleted input
file is still rejected.
