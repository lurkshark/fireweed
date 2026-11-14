---
module: assets
phase: applied
run: 20261114T055334-d6def17ce8950c830534cd93
---

## Motivation

The design's New spec IDs are assigned late and Stale changes are caught
sections explicitly require authoring guidance to combine cross-change handle
dependencies and run stamp (R04/R20). The embedded author-change playbook on
main 3b3b64d contains neither instruction, although the validators exist.

## Interface

No public signature or type changes.

## Added

### AST-064
The author-change playbook MUST instruct authors to use unique change-local new:<slug> handles for Added specs and combine changes when one would need another's handle. It MUST direct authors to run fireweed stamp <change> for Modified or Removed specs and resolve stale conflicts rather than overwrite stale bases.

## Modified



## Removed

## Decisions

## Notes

Keep the existing Motivation, narrowing/already_satisfied and no-worker-commit
guidance. Add a short example rather than a second protocol document. Use
public assets.Playbook assertions for the instructions, accepting equivalent
prose. Do not alter copied project playbooks or managed AGENTS blocks; stale
override detection already follows the embedded hash through CHK-069.

Drafted at the owner's request after the R01–R29 traceability review.
No lifecycle execution, agent invocation, drill or release is authorized by
this drafting request. Workers must obtain context before their step and use
fireweed verify; only the controller may advance or commit lifecycle work.

The AST-064 eval fails against the current `assets.Playbook("author-change")`
output because it lacks the handle, stamping and stale-conflict instructions.
The implementation step can add the guidance in the disposable
`packages/assets/internal/content/content.go` playbook entry.
## Audit

No findings; checked AST-064

## Provenance

- evals 2026-11-14T05:58:12.334767Z agent=codex fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=141 failing=1 passing=140
- audit 2026-11-14T05:58:31.292857Z agent=codex fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
- impl 2026-11-14T06:26:11.088057Z agent=codex fireweed=0.1.0 playbook=5f856d6bcab308b9ff1a8a7c03b9d2547744112aba81dd30dea8f61ec8062b03 tests=1325 failing=0 passing=1325

