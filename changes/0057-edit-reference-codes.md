---
bases:
    PRJ-013: 403da758
    PRJ-014: f84a2ce6
    PRJ-015: 76e385f9
    PRJ-016: 940ee23a
module: project
---



## Motivation

The design requires stable diagnostic causes. D006 assigns C05 to duplicate
identities, C09 to handle references and C08 to bases. Check already rejects
these errors using C01. Migrate those existing paths, treating an absent edit
target as a failure to resolve its preceding base, without adding rejections.

## Interface

No public interface change.

## Added



## Modified

### PRJ-013
Check MUST reject a change that uses the same edit identity more than once across Added, Modified and Removed with diagnostic C05.

### PRJ-014
Check MUST reject a pending change that references a handle absent from that change's Added entries with diagnostic C09.

### PRJ-015
Check MUST reject Modified or Removed entries whose bases hash is missing or differs from the preceding projected spec's canonical hash with diagnostic C08.

### PRJ-016
Check MUST reject Modified or Removed IDs absent from the projection preceding the change with diagnostic C08.

## Removed



## Decisions



## Notes

Reuse the existing CheckDiagnostics table and its Removed-base follow-ups. Add
exact assigned-code assertions matched to each existing rejection's reason;
retain all current inputs, valid cases, nonempty code/message checks and reason
assertions, including missing/stale Modified and Removed bases and both absent
target cases. Preserve earlier exact-code assertions. Do not add fixtures,
scenarios or helpers; reporting labels may change as needed for correct ID
attribution. Changed IDs must genuinely fail before implementation.
Migrate only the existing duplicate edit, undeclared handle, missing/stale base
and absent preceding target paths, including existing projection-application
errors for those absent targets. Preserve messages and error behavior, and leave
other diagnostic causes unchanged. Do not implement missing features or create
speculative errors. Controller owns spec application, commits and archive.
