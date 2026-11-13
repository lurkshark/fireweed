---
module: assets
---

## Motivation

The design's New spec IDs are assigned late and Stale changes are caught
sections explicitly require authoring guidance to combine cross-change handle
dependencies and run stamp (R04/R20). The embedded author-change playbook on
main 3b3b64d contains neither instruction, although the validators exist.

## Interface

No public signature or type changes.

## Added

### new:author-handles-and-stamp
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
