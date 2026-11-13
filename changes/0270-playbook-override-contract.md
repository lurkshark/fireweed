---
module: workflow
already_satisfied: [new:override-context, new:override-provenance]
---

## Motivation

The design's Playbooks and Provenance sections permit project-owned playbook
overrides and record the text actually used (R20). FLOW-005/051 currently
name a matching playbook and its hash but never specify override selection.
Current workflow code already selects the project override for context and
provenance; preserve that behavior explicitly for reconstruction.

## Interface

No public signature or type changes.

## Added

### new:override-context
Context MUST use the complete text of .fireweed/playbooks/<name>.md when that override exists, selecting the held phase name or regenerate under FLOW-083. When the override is absent it MUST use the corresponding assets.Playbook text.

### new:override-provenance
The playbook hash appended by Done MUST be the full hexadecimal SHA-256 of the exact override or embedded text selected for that completed phase under the Context playbook-selection rule.

## Modified

## Removed

## Decisions

## Notes

Both additions are legitimately already_satisfied by the current public
behavior; no implementation change is requested. Extend the existing context
and provenance fixtures with project-owned override text distinguishable from
the embedded text and a missing override fallback.
Use public Context/Done and ordinary fixture filesystem operations. Keep
existing FLOW-005/051 assertions, preserved provenance lines, regeneration
playbook selection, protected zones and CHK-069 stale embedded_base warning.
Do not turn hash metadata into a new blocking compatibility requirement.
Do not add unreadable-override hardening or repeat provenance-line structure
already owned by FLOW-051. The new provenance assertion covers which text is
hashed, including the existing full SHA-256 representation.

Draft only. No lifecycle execution, agent invocation or drill is authorized.
Workers obtain context before their step and use fireweed verify; the
controller alone advances and commits lifecycle work.
