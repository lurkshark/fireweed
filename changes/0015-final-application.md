---
module: project
---

## Motivation

Projection and late allocation exist, but implementation completion still needs a final module document with real specs, metadata and the worker's durable Decisions. Keep this pure transform separate from the workflow's commit/archive transaction. Remove obsolete grammar wording from the retained project Decisions so reconstruction follows the current contracts.

## Interface

Add Finalize(module contracts.Module, change contracts.Change) (contracts.Module, error). It applies the existing ApplyChange transform, rejects remaining local IDs and returns a complete canonical module, including appended/removed Decisions. It does not mutate files, archive the change or advance lifecycle state. Keep ApplyChange's preview behavior unchanged.

## Added

### new:final-spec
Finalize MUST apply Added/Modified/Removed specs and all allowed set values through the existing public projection contract, retaining item order, tags, module identity, preamble and unrelated sections. Its result MUST be a valid current module with real IDs only, with resolved fields, Specs and metadata agreeing under public Markdown normalization. It MUST reject invalid applications and remaining local handles in current or resulting specs, returning a zero Module on error. The change MUST have a nonempty valid slug consistent with its filename; differing numeric filename/typed Number values MUST remain permitted. Finalize MUST NOT assign IDs, change run/phase, write the change, or incorporate the change's Motivation, Interface, Notes, Audit, Questions or Provenance into the module.

### new:decisions
Finalize MUST start from the module's typed Decisions, overriding a stale retained Decisions body. For each nonempty line in the change's Removed decisions section, it MUST trim surrounding whitespace and one optional leading blockquote marker and remove every exact occurrence of that quoted text from existing Decisions; unmatched quotes MUST be harmless, and unrelated text MUST remain. It MUST then append the change's nonempty Decisions body, preceded by a bracketed change slug on its own line and separated from existing nonempty Decisions by one blank line. It MUST preserve multiline/fenced decision content and existing slug tags under public Markdown normalization. No new Decisions section MUST be invented when none exists and neither removal nor addition requests one; an explicit removal request MUST retain a Decisions section even if its resulting body is empty. The typed Decisions field and document body MUST agree.

### new:final-purity
Finalize MUST be deterministic and pure on successful and failed inputs, without filesystem/network I/O or execution of Git, hooks or agents. It MUST NOT mutate or return mutable aliases to the supplied module, change, retained documents, metadata maps or slices. Separate calls and input values MUST remain unchanged immediately when a returned result is mutated, including valid maintenance changes with no spec edits. Errors MUST return no partial module. ApplyChange MUST remain a preview retaining existing Decisions without performing final decision additions/removals.

## Modified

## Removed

## Removed decisions

> Metadata/body validation still uses the public change grammar, including its current real-ID-only fixes/already_satisfied fields.

## Decisions

Removed-decision quoting follows the pinned reference's line-wise exact-text convention, including optional Markdown blockquote markers and harmless unmatched quotes. Public rendering/parsing defines normalization of remaining whitespace and human content. Each change appends one slug-tagged body, preserving multiline choices and fenced examples as text. Finalize is not an idempotent transaction: the workflow validates, writes and archives once, and handles rollback separately. Existing ApplyChange remains the placeholder-capable preview and therefore does not perform final Decisions edits. Current CON-016 permits locally declared Added handles in already_satisfied before allocation; fixes and bases remain real-ID-only. Allocation must rewrite those handles before this final transform can produce a current SPEC.md.

## Rejected alternatives

Writing final SPEC.md while previews still contain handles would leak temporary names into the permanent contract. Appending every change's author sections would mix transient implementation context into the module's durable Decisions. Performing commits here would couple a deterministic transform to workflow recovery.
