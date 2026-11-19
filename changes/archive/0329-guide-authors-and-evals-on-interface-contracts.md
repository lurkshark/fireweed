---
module: assets
phase: applied
run: 20261119T002320-1d7087d5d40dd94906df1936
---

## Motivation

The November 18 drills found that callers depended on which inputs a supplier reads,
which settings survive, and which details are meaningful. Existing author-change
guidance addresses distinguishable failure outcomes. Users need similarly focused
guidance for choosing an authoritative input and for testing behavior without
freezing incidental formatting. The design places observable contracts at the
supplying module interface.
[Drill reports](https://github.com/lurkshark/fireweed/pull/25)

## Interface

Add focused guidance to the built-in author-change and evals playbooks through
Playbook.

## Added

### AST-074
The author-change playbook MUST direct authors to specify which inputs govern an
operation when more than one source can supply them. It MUST direct authors to state
fallback, merge or preservation behavior when callers rely on it.

### AST-075
The evals playbook MUST guide workers to test source selection through the supplying
public interface by making a competing input conflicting or invalid when the
contract permits ignoring it.

### AST-076
The evals playbook MUST guide workers to assert the contracted result of structured
outputs and assert exact representation only when required by the spec. For a
regression caused by incidental format or order, it MUST advise using a second
valid representation to exercise the distinction.

## Modified

## Removed

## Decisions

## Notes

Add short examples using conflicting or malformed inactive configuration and
equivalent YAML field order. Test the meaning of the guidance through Playbook,
allowing equivalent wording. Keep the existing failure-contract guidance, audit
scope, module ownership and independent dependent-eval signal.

Fireweed verification passed for the evals step. AST-074, AST-075 and AST-076
each failed on missing guidance in the current `Playbook` output.
## Provenance

- evals 2026-11-19T00:29:27.972428Z agent=codex fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=160 failing=3 passing=157
- audit 2026-11-19T00:30:01.66117Z agent=manual fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
- impl 2026-11-19T01:00:12.275509Z agent=codex fireweed=0.2.1 playbook=5f856d6bcab308b9ff1a8a7c03b9d2547744112aba81dd30dea8f61ec8062b03 tests=1481 failing=0 passing=1481

## Audit

No findings; checked AST-074, AST-075, AST-076

