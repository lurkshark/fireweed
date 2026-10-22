---
bases:
    CON-016: 922d6d1b
module: contracts
phase: evals
run: a0503a38234a37bc8153e97a
---

## Motivation

The design allows an author to mark a newly added spec already satisfied before start assigns real IDs. The initial grammar required real IDs in this field, leaving that designed use case impossible. Preserve local handle scope while allowing those declarations.

## Interface

No public interface change. Existing ParseChange retains local Added handles in AlreadySatisfied until the later start/allocation transform rewrites every occurrence.

## Added

## Modified

### CON-016
Change set metadata MUST allow only depends_on, interface, evals, disposable, preserve, append_only, boundaries_opt_out and per-module hooks. It MUST reject name/id_prefix identity edits, unknown fields and incorrectly typed values, and retain explicit empty lists or hook strings. Bases MUST contain eight lowercase hexadecimal hashes keyed by real IDs; fixes MUST contain real IDs rather than handles. already_satisfied MUST accept real IDs and locally declared Added handles, retain their order and exact strings, and reject undeclared/malformed handles and incorrectly typed values. A declaration in another change MUST NOT satisfy a handle reference; parsing MUST retain independence between results.

## Removed

## Decisions

This corrects the earlier real-ID-only already_satisfied restriction using the design's frontmatter definition: new or modified specs may legitimately pass before implementation, and new specs use local handles before start. fixes continues to refer to preexisting real IDs. AlreadySatisfied retains local handles during queue preview; the later all-occurrence allocation transform will replace them with assigned real IDs. Scope remains validated by the existing full-document local-reference rule, including metadata. No automatic already_satisfied exemptions are introduced.

## Rejected alternatives

Requiring authors to predict real IDs would defeat late allocation. Requiring start before authoring already_satisfied would make the authored approved change incomplete and separate it from the normal authoring workflow.

## Provenance

- evals | 2026-10-22T14:45:32Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 27 | failing: 2 | passing: 25
- audit | 2026-10-22T14:47:05Z | agent: manual | fireweed: 0.4.0 | playbook: a9c982d24b476621316049b527e6c7bbc319f06b12268af150fa7d8c6df92230 | tests: 0 | failing: 0 | passing: 0

## Audit

- CON-016: "parsing MUST retain independence between results." — assert immediately after mutating the first result that the already parsed second result still contains its original slice, nested metadata and sections, before parsing again. Shared reusable storage could reset both second and third results and evade the current comparison. Independent auditor /root/audit_document_grammar used permitted public context only, without implementation/history/bootstrap or test execution.
