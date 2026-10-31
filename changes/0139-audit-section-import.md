---
bases: {}
fixes:
    - RUN-047
    - RUN-048
module: runner
---



## Motivation

EC5's real-worker NEW lifecycle reached audit after genuine fail-first evidence,
then stopped with "audit workspace modified change outside Audit". A focused
NEW CLI reproduction demonstrates the same error for a valid Audit section
inserted before existing Provenance, with every original byte preserved. The
current validator accepts a first Audit section only at the end of the file.
That positional restriction is absent from RUN-047/048 and prevents the design's
ordinary Audit submission from reaching the held change. Accept a valid Audit
section in normal Markdown section order while still rejecting every edit
outside it. This is a demonstrated core gate failure, not deferred surface work.

## Interface

No public API or spec changes.

## Added



## Modified



## Removed



## Decisions



## Rejected alternatives

Telling real auditors to append at a special position would evade the existing
Audit-only contract. Weakening non-Audit byte validation would violate it.

## Notes

Add a public regression that supplies an otherwise unchanged current change
and inserts its first Audit section before existing Provenance. Keep every
original byte intact, require successful import/lifecycle completion, and
confirm the held non-Audit contents survive. Preserve existing accepted tail
submissions, partial-work/error/stop behavior, invalid outside-edit rejection,
independent workspace/history isolation and cleanup. Reuse the current fixture
and assertions; add no new requirements, interfaces, instrumentation or real
agents to permanent evals. Only the active controller commits and applies.

Failed EC5 clone remains inspectable at
/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-ec5-y5nxlf50/candidate
with fail-first eval commit a455bb8. Its owning NEW pin released its stopped
claim; no controller was switched or claim resumed across controllers. The
focused diagnostic fixtures also released only their own claims. Resume the
acceptance gates with the corrected candidate after normal verification passes.
