---
eval_fix:
    - CON-017
    - CON-018
    - CON-019
module: contracts
phase: applied
run: 20261111T213207-39b0efee633a4ff9b1ad4295
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

The contracts evals passed `fireweed verify item-error-wording-evals`.
## Provenance
- evals 2026-11-11T21:33:13.792968Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=68 failing=0 passing=68
- audit 2026-11-11T21:33:28.758395Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-astra -c model_reasoning_effort="medium" -s workspace-write {prompt} fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
- impl 2026-11-11T22:00:06.918157Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=5f856d6bcab308b9ff1a8a7c03b9d2547744112aba81dd30dea8f61ec8062b03 tests=1148 failing=0 passing=1148
## Audit

No findings; checked CON-017, CON-018, CON-019
