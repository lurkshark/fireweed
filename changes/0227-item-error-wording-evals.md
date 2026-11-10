---
eval_fix:
    - CON-017
    - CON-018
    - CON-019
module: contracts
---

## Motivation

The design makes codes, not wording, the stable diagnostic contract.
CON-017 to CON-019 require the C05 or C03 code and, for duplicates, the
duplicated ID as subject. Their evals also require exact messages
(`duplicate item ID: <id>`, `malformed item heading: "<heading>"`) and
message prefixes for frontmatter errors, which no spec states. The
contracts drill lists this as a gap.

## Interface

Audited eval_fix only. Edit only the contracts evals zone and this change.

## Added

## Modified

## Removed

## Decisions

## Notes

In `Test_CON_017_ParseItemsDuplicateCode`, `Test_CON_018_ParseItemsMalformedCode`
and `Test_CON_019_ParseDocumentFrontmatterCode`, keep the `contracts.ErrorCode`
code and subject checks. Replace each exact message or prefix check with a
check that the text names the duplicated ID or the malformed heading where
the fixture has one; otherwise drop it.
