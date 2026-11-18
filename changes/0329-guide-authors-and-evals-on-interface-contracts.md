---
module: assets
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

### new:author-playbook-input-authority
The author-change playbook MUST direct authors to specify which inputs govern an
operation when more than one source can supply them. It MUST direct authors to state
fallback, merge or preservation behavior when callers rely on it.

### new:eval-playbook-input-authority
The evals playbook MUST guide workers to test source selection through the supplying
public interface by making a competing input conflicting or invalid when the
contract permits ignoring it.

### new:eval-playbook-semantic-contracts
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
