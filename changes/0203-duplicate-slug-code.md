---
already_satisfied:
    - PRJ-024
bases:
    PRJ-024: 09d40177
module: project
---

## Motivation

The design requires stable diagnostic codes, and APP-019 requires
`new-change` to report C13 for a duplicate slug. PRJ-024 requires the
rejection but names no code, so the project drill's first reconstruction
chose a different code and passed only after the retry showed it the app
failure. The code is part of the project boundary and belongs in its spec.

## Interface

No API changes. The current implementation already returns C13.

## Added

## Modified

### PRJ-024
NewChange MUST reject a slug already present in the pending queue or archive with diagnostic C13, without overwriting existing files.

## Removed

## Decisions

## Notes

The existing PRJ-024 evals should also check the C13 code through
`contracts.ErrorCode` for both the pending and the archived duplicate.
Do not assert the error's wording.
