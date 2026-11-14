---
already_satisfied:
    - FLOW-143
    - FLOW-144
module: workflow
phase: impl
run: 20261114T080752-882730f123ec22d21b064c0d
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

### FLOW-143
Context MUST use the complete text of .fireweed/playbooks/<name>.md when that override exists, selecting the held phase name or regenerate under FLOW-083. When the override is absent it MUST use the corresponding assets.Playbook text.

### FLOW-144
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
## Audit

No findings; checked FLOW-143, FLOW-144

## Provenance

- evals 2026-11-14T08:13:26.797701Z agent=codex fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=211 failing=0 passing=211
- audit 2026-11-14T08:13:54.580213Z agent=codex fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0

