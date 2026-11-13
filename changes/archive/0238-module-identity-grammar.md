---
already_satisfied:
    - APP-159
bases: {}
module: app
phase: applied
run: 20261113T012218-a5369281eb32f3b76b2fe98d
---

## Motivation

The owner approved making the new-module identity grammar explicit before app regeneration. APP-103 requires invalid-input rejection but leaves the accepted names and prefixes implicit. Preserve the existing observable grammar so a reconstruction can reproduce it without inferring it from implementation or incidental eval inputs.

## Interface

No public API changes.

## Added

### APP-159
new-module MUST accept names matching `[a-z][a-z0-9]*(?:-[a-z0-9]+)*` and prefixes matching `[A-Z][A-Z0-9]*`, with each match covering the entire value. Values outside those ASCII grammars MUST produce corrective U01 and exit 2 before writing module files.


## Modified



## Removed

## Decisions

## Notes

The owner authorized a one-time archive correction during the PR #2 merge: APP-156 became APP-159 to avoid the IDs already used by main. Requirement text, eval behavior, lifecycle phases, run IDs and Provenance are unchanged.

This records the current moduleNamePattern/modulePrefixPattern behavior. Retain APP-103 error behavior and extend existing new-module fixtures with representative accepted digit/hyphen names and letter-led prefixes, and rejected case/separator/leading-digit forms only where existing assertions leave the grammar uncovered. No product behavior change is intended; already_satisfied is legitimate.
Owner authorized these five regeneration prerequisites and their necessary module-local changes. Stop after this queue; no regeneration or drill is authorized. Follow SYSTEM.md and the controller context. Workers must not read docs/rebuild/PROGRESS.md or ACCEPTANCE.md. Preserve unrelated assertions and retained interfaces after evals. Use existing public-interface fixtures and focused assertions; do not redesign helpers or add a test matrix. Evals and implementation workers run fireweed context for this change before each step and fireweed verify for verification. Restricted auditors follow only the supplied audit playbook and inputs. Only the controller advances phases and commits.
## Provenance
- evals 2026-11-13T01:41:34.421315Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=332 failing=0 passing=332
- audit 2026-11-13T01:42:02.100525Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-astra -c model_reasoning_effort="medium" -s workspace-write {prompt} fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
- impl 2026-11-13T01:49:47.062022Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=5f856d6bcab308b9ff1a8a7c03b9d2547744112aba81dd30dea8f61ec8062b03 tests=332 failing=0 passing=332
## Audit

No findings; checked APP-159
