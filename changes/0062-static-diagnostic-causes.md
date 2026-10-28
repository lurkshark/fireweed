---
already_satisfied:
    - new:check-config-codes
bases: {}
fixes:
    - PRJ-013
module: project
---



## Motivation

The design requires stable diagnostic causes; D006 allocates the remaining
static-check causes. Check already rejects these inputs but still labels them
C01, including graph, document, metadata and queue errors. Complete its existing
cause classification as one cohesive project change. A duplicate identity in
one edit section also reaches the queue parser before PRJ-013's later check;
retain that existing rejection while correcting its code to C05.

## Interface

No public interface change.

## Added

### new:check-config-codes
Check MUST report configuration discovery or decoding failures with diagnostic C01.

### new:check-graph-codes
Check MUST report invalid current or projected module graphs with diagnostic C02.

### new:check-syntax-codes
Check MUST report unreadable or malformed module/change documents and malformed item syntax with diagnostic C03, retaining C05 for duplicate identities.

### new:check-metadata-codes
Check MUST report invalid module metadata or zone field types, and invalid set keys or value types, with diagnostic C12.

### new:check-queue-codes
Check MUST report pending queue identity errors with diagnostic C13.

### new:check-change-module-codes
Check MUST report invalid pending change module targets with diagnostic C07.

## Modified



## Removed



## Decisions



## Notes

Reuse relevant existing Load/Queue/Project rejection fixtures by adding Check
assertions for the assigned code and the same cause. Preserve all current API,
valid-input, message, exact-code, Added/Modified/Removed and ID assertions. Add
only a representative fixture where a named cause has no reusable one; no
matrix expansion, helper infrastructure or instrumentation. In the existing
PRJ-013 fixture, exercise a duplicate identity within one edit section as the
regression case, alongside the retained cross-section case. It must fail on
C01 before implementation. Other changed cause IDs must genuinely fail; C01
configuration reading/decoding is legitimately already satisfied.

Migrate existing error branches only: Load errors distinguish configuration,
module document syntax and metadata/zone declarations; current/projected graph
errors use C02; module/change item syntax uses C03 except duplicate identities
C05; queue identity errors use C13; invalid change module targets use C07;
invalid set declarations use C12. Preserve C04–C09 causes already migrated,
especially absent preceding targets C08 and tag errors C06. Keep every existing
error condition and message; adding private cause information at its existing
source is allowed when needed for accurate propagation, without changing the
public interfaces. Do not add a general error framework, speculative checks,
unsupported features or irrelevant probes. Controller owns specs and archive.
