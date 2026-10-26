---
already_satisfied:
    - CON-009
    - CON-013
    - CON-016
bases:
    CON-009: 3ef0e798
    CON-013: a58e5592
    CON-016: dc455fb6
module: contracts
---

## Motivation

The 2026-10-26 project review found effort going into hardening beyond the design while core commands stayed unbuilt. The revised design (Spec shape; SYSTEM.md is a plain document) and the SYSTEM.md scope guardrails say that locality and safety properties are conventions checked by review, that specs state observable behavior without aliasing or purity requirements, and that the supported environment is macOS and Linux with ordinary full SHA-1 clones. See D048.

## Interface

No public API change.

## Added

## Modified

### CON-009
Configuration parsing MUST retain every design field, including shell hooks and agent argv/wrap values, per-project zone defaults and audit settings, without executing commands or expanding placeholders. It MUST reject malformed YAML, duplicate keys, multiple YAML documents, non-mapping roots and fields of the wrong type.

### CON-013
Module parsing MUST require valid name and id_prefix, exactly one Specs section, valid normative spec items, and typed known metadata fields. Omitted zone lists and hooks MUST inherit supplied project defaults; explicit empty lists and explicitly empty hook strings MUST override them. Depends_on, boundaries_opt_out and Decisions MUST be retained.

### CON-016
Change set metadata MUST allow only depends_on, interface, evals, disposable, preserve, append_only, boundaries_opt_out and per-module hooks. It MUST reject name/id_prefix identity edits, unknown fields and incorrectly typed values, and retain explicit empty lists or hook strings. Bases MUST contain eight lowercase hexadecimal hashes keyed by real IDs; fixes MUST contain real IDs rather than handles. already_satisfied MUST accept real IDs and locally declared Added handles, retain their order and exact strings, and reject undeclared/malformed handles and incorrectly typed values. A declaration in another change MUST NOT satisfy a handle reference.

## Removed

## Decisions

Narrowed specs are listed in already_satisfied because the current implementation already meets the narrower text; nothing here adds behavior. Clauses removed from a spec become unspecified, not forbidden: existing code that still does more is conforming, and the implementation step may simplify it.

## Rejected alternatives

Keeping the clauses and only deleting their evals: coverage requires every enforced spec to stay covered, and the next audit would reinstate the tests. Rewriting every over-long spec to the new shape limits at once: most long specs carry no hardening and rewriting them adds risk without benefit; they are left for later changes.

## Notes

Evals step: remove subtests that only assert independence of parse results or slices shared with inputs. Every Modified ID must keep at least one referencing test. Implementation step: no product change is required.
