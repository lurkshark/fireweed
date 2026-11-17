---
fixes:
    - CHK-065
    - CHK-066
module: checks
phase: impl
run: 20261116T235732-2fe35230e92cf36479efcd08
---

## Motivation

Doctor should diagnose the same usage-limit signal users encounter during a
run. On current main a failed command that emits RATE LIMIT on stdout,
followed by 19 stderr lines, is classified as an ordinary failure by Doctor.
InvokeAgent classifies the same output as a usage limit. Doctor inserts a
blank line at the stream boundary and drops the oldest actual line from its
20-line window.

## Interface

## Added

## Modified

## Removed

## Decisions

## Notes

Add focused Doctor cases with one matched stdout line followed by 19 and
20 stderr lines. Use a nonzero exit for both. Assert C15 usage-limit
classification when the match is twentieth from the end, then ordinary
failure and exclusion when the match is twenty-first from the end. Keep the
captured streams and full probe logs intact. Count real emitted lines,
including deliberate blank lines. Use the existing stdout-then-stderr order
for independently captured streams.

The focused eval failed against the current implementation for both IDs.
At twentieth, Doctor dropped the match and reported an ordinary failure.
At twenty-first, the ordinary failure excluded the match, but the tail
showed a blank at the stream boundary instead of the emitted trailing blank.
Both probe logs retained the complete stdout and stderr. `fireweed verify`
passed the evals phase with this fail-first evidence.

The implementation can be updated in `packages/checks/internal/engine/doctor.go`.
## Audit

No findings; checked CHK-065, CHK-066

## Provenance

- evals 2026-11-17T00:07:30.337208Z agent=codex fireweed=0.2.1 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=138 failing=3 passing=135
- audit 2026-11-17T00:07:49.867697Z agent=codex fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0

