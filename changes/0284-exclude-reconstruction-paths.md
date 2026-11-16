---
bases:
    RUN-053: fc0ba001
module: runner
---

## Motivation

The `contracts` drill worker read `docs/rebuild/PROGRESS.md` and
`docs/rebuild/ACCEPTANCE.md`, which contain results from earlier runs. Access
to those results compromises the blind reconstruction that the drill is meant
to measure. Project owners need to keep project-specific history and other
operator documentation out of the worker's reconstruction inputs.

## Interface

Add optional `drill.exclude` to `fireweed.yaml`. It is a YAML list of
repository-relative file paths and globs using Fireweed's existing `/`, `*`
and `**` matching rules. Use `docs/**` for an entire directory tree. The list
applies to the drill worker's reconstruction workspace. An omitted or empty
list adds no exclusions.

```yaml
drill:
  exclude:
    - docs/**
    - MAINTAINERS.md
```

## Added

### new:drill-exclusion-config
Drill MUST reject a malformed `drill.exclude` setting with C01 pointing to the setting before running setup or invoking the worker. Each entry MUST be a nonempty repository-relative pattern without `.` or `..` path components.

### new:required-reconstruction-inputs
If `drill.exclude` matches the target's spec or a retained target interface or eval file, Drill MUST return C01 identifying the conflicting path before running setup or invoking the worker.

### new:reconstruction-exclusions-after-setup
If the reconstruction workspace contains a path matching `drill.exclude` after setup, Drill MUST reject the input before invoking the worker and identify the matching path.

## Modified

### RUN-053
Drill MUST invoke reconstruction in a separate workspace containing the invoking HEAD's tracked files except the target's disposable files, direct and transitive dependent modules' files, the changes directory, drill reports and paths matching `drill.exclude`. Its Git metadata MUST hold one fresh commit of the filtered inputs and none of the invoking repository's history, refs or saved attempts. Drill MUST run the configured setup hook there before agent invocation.

## Removed

## Decisions

## Notes

Use a bounded fake agent to assert that a matching root file and a nested
file under `docs/` are absent from the input tree and its seed commit. Exercise
an omitted list and an unrelated file that should still be available. Make the
verification hook check that the excluded files remain available in the
verification tree. Cover malformed settings, a match against a required
target input, and a setup hook that recreates an excluded path. The invoking
checkout and retained reconstruction ref should still contain the tracked
documentation.
