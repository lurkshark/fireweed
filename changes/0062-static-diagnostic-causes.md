---
already_satisfied:
    - PRJ-040
audit_rounds: 0
bases: {}
fixes:
    - PRJ-013
module: project
phase: impl
run: 20261028T083708-d4e70feeb59862f196802fca
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

### PRJ-040
Check MUST report configuration discovery or decoding failures with diagnostic C01.

### PRJ-041
Check MUST report invalid current or projected module graphs with diagnostic C02.

### PRJ-042
Check MUST report unreadable or malformed module/change documents and malformed item syntax with diagnostic C03, retaining C05 for duplicate identities.

### PRJ-043
Check MUST report invalid module metadata or zone field types, and invalid set keys or value types, with diagnostic C12.

### PRJ-044
Check MUST report pending queue identity errors with diagnostic C13.

### PRJ-045
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

Eval quick verification accepted the fail-first set: existing cases and C01
configuration cases pass; 14 changed-cause scenarios fail on their assigned
codes. The PRJ-013 table retains its cross-section case, while a separate case
covers duplicate identities within Added. Its parent test label no longer
attributes every table row to PRJ-013, keeping the per-spec case count bounded.

## Provenance

- evals 2026-10-28T08:43:18.618304Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=228 failing=17 passing=211
- audit 2026-10-28T08:44:17.52309Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0

## Audit

No findings. Checked PRJ-013, PRJ-040, PRJ-041, PRJ-042, PRJ-043, PRJ-044, and PRJ-045.
