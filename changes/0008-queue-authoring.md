---
bases: {}
module: project
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

### new:queue-documents
Queue MUST read pending NNNN-slug.md documents in numeric-then-slug order, allowing tied numbers and excluding archives. It MUST return each change's module, frontmatter and parsed Added/Modified/Removed items, rejecting malformed documents, unknown modules and ambiguous slugs. A change argument MUST resolve its unique slug, queued basename or repository-relative queued path, rejecting missing or ambiguous targets.

### new:projection
Project MUST apply pending changes through the named change inclusively, or the whole queue when through is empty, preserving Added/Modified full text and removing Removed items. It MUST apply set values for depends_on, interface, evals, disposable, preserve, append_only, boundaries_opt_out and hooks, rejecting identity/unknown fields and an invalid projected graph. Added handles MUST remain local to their declaring change; different changes may reuse handle text.

### new:stamp-bases
Stamp MUST write the canonical SpecHash of each Modified/Removed item from the projection preceding the target change into its bases. It MUST reject IDs absent from that projection and preserve other frontmatter and human sections. Stamping an Added-only change MUST succeed without assigning IDs.

### new:author-checks
Check MUST reject missing Motivation, non-normative spec bodies, invalid review/quarantine tags, duplicate or foreign module IDs, duplicate edit identities within a change and references to undeclared handles. It MUST reject missing/stale bases and Modified/Removed IDs absent from the preceding projection. Check results MUST have no error diagnostics for valid module specs/changes, and nonempty C01 code/message diagnostics for failures.

## Modified



## Removed



## Decisions
