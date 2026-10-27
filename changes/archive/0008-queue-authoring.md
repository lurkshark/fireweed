---
audit_rounds: 0
bases: {}
module: project
phase: applied
run: 20261027T011358-d724b4b7a8213be798cd2f15
---



## Motivation

The NEW CLI has created the manual slice's Added change and contracts can parse
its handle. The design's Change files and stale-change rules now require queue
projection, stamp and static author checks before start can assign IDs. Add only
these operations, not history, budgets, init or lifecycle mutation.

## Interface

Extend retained model/ with data-only Change (Number, Slug, Path, Module,
Document, Added, Modified, Removed) and Diagnostic (Code, Message, Path, Warning)
types; use contracts Document/Item types and expose public aliases. Thin wrappers:
Queue(root string) ([]Change,error), Project(root,through string) (Snapshot,error),
Stamp(root,change string) error, Check(root string) []Diagnostic. A change argument
accepts its unambiguous slug, queued basename or repository-relative queued path.
Keep all parsing/projection/validation behavior under internal/.

## Added

### PRJ-004
Queue MUST read pending NNNN-slug.md documents in numeric-then-slug order, allowing tied numbers and excluding archives. It MUST return each change's module, frontmatter and parsed Added/Modified/Removed items, rejecting malformed documents, unknown modules and ambiguous slugs. A change argument MUST resolve its unique slug, queued basename or repository-relative queued path, rejecting missing or ambiguous targets.

### PRJ-005
Project MUST apply pending changes through the named change inclusively, or the whole queue when through is empty, preserving Added/Modified full text and removing Removed items. It MUST apply set values for depends_on, interface, evals, disposable, preserve, append_only, boundaries_opt_out and hooks, rejecting identity/unknown fields and an invalid projected graph. Added handles MUST remain local to their declaring change; different changes may reuse handle text.

### PRJ-006
Stamp MUST write the canonical SpecHash of each Modified/Removed item from the projection preceding the target change into its bases. It MUST reject IDs absent from that projection and preserve other frontmatter and human sections. Stamping an Added-only change MUST succeed without assigning IDs.

### PRJ-007
Check MUST reject missing Motivation, non-normative spec bodies, invalid review/quarantine tags, duplicate or foreign module IDs, duplicate edit identities within a change and references to undeclared handles. It MUST reject missing/stale bases and Modified/Removed IDs absent from the preceding projection. Check results MUST have no error diagnostics for valid module specs/changes, and nonempty C01 code/message diagnostics for failures.

## Modified



## Removed



## Decisions



## Provenance

- evals 2026-10-27T01:20:43.362625Z agent=manual fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=100 failing=24 passing=76
- audit 2026-10-27T01:23:04.979543Z agent=manual fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0
- impl 2026-10-27T01:27:18.160841Z agent=manual fireweed=0.1.0 playbook=07554a3249d8fca0ba70df7504135f297491465caa63458734a4b8351734d0e5 tests=104 failing=0 passing=104

## Audit

No findings. Checked PRJ-004, PRJ-005, PRJ-006, and PRJ-007.
