---
bases:
    RUN-053: fc0ba001
module: runner
phase: applied
run: 20261116T032105-5578b407f7bf48b643ab538b
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

### RUN-075
Drill MUST reject a malformed `drill.exclude` setting with C01 pointing to the setting before running setup or invoking the worker. Each entry MUST be a nonempty repository-relative pattern without `.` or `..` path components.

### RUN-076
If `drill.exclude` matches the target's spec or a retained target interface or eval file, Drill MUST return C01 identifying the conflicting path before running setup or invoking the worker.

### RUN-077
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

Against the current runner, the RUN-053 eval found matching files in both
the reconstruction tree and its seed commit. The omitted and empty lists
passed. The RUN-075 and RUN-076 evals observed setup and worker invocations
for invalid inputs. The RUN-077 eval observed a worker invocation after
setup recreated an excluded path.

`fireweed verify exclude-reconstruction-paths` passed in the evals phase.
Its report recorded fail-first cases for all four affected IDs.

Implementation can use `packages/runner/internal/drill/drill.go` for preflight
and post-setup checks and `packages/runner/internal/drill/input.go` for tracked
input filtering before the seed commit. Both files are disposable.
## Audit

No findings; checked RUN-053, RUN-075, RUN-076, RUN-077

## Provenance

- evals 2026-11-16T03:33:05.031379Z agent=codex fireweed=0.2.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=217 failing=16 passing=201
- audit 2026-11-16T03:33:31.906184Z agent=codex fireweed=0.2.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
- impl 2026-11-16T03:55:15.510708Z agent=codex fireweed=0.2.0 playbook=5f856d6bcab308b9ff1a8a7c03b9d2547744112aba81dd30dea8f61ec8062b03 tests=620 failing=0 passing=620

