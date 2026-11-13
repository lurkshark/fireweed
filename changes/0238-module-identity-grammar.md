---
already_satisfied:
    - new:module-grammar
bases: {}
module: app
---

## Motivation

The owner approved making the new-module identity grammar explicit before app regeneration. APP-103 requires invalid-input rejection but leaves the accepted names and prefixes implicit. Preserve the existing observable grammar so a reconstruction can reproduce it without inferring it from implementation or incidental eval inputs.

## Interface

No public API changes.

## Added

### new:module-grammar
new-module MUST accept names matching `[a-z][a-z0-9]*(?:-[a-z0-9]+)*` and prefixes matching `[A-Z][A-Z0-9]*`, with each match covering the entire value. Values outside those ASCII grammars MUST produce corrective U01 and exit 2 before writing module files.


## Modified



## Removed

## Decisions

## Notes

This records the current moduleNamePattern/modulePrefixPattern behavior. Retain APP-103 error behavior and extend existing new-module fixtures with representative accepted digit/hyphen names and letter-led prefixes, and rejected case/separator/leading-digit forms only where existing assertions leave the grammar uncovered. No product behavior change is intended; already_satisfied is legitimate.
Owner authorized these five regeneration prerequisites and their necessary module-local changes. Stop after this queue; no regeneration or drill is authorized. Follow SYSTEM.md and the controller context. Workers must not read docs/rebuild/PROGRESS.md or ACCEPTANCE.md. Preserve unrelated assertions and retained interfaces after evals. Use existing public-interface fixtures and focused assertions; do not redesign helpers or add a test matrix. Evals and implementation workers run fireweed context for this change before each step and fireweed verify for verification. Restricted auditors follow only the supplied audit playbook and inputs. Only the controller advances phases and commits.
