---
module: project
phase: applied
run: 20261114T004020-1e354415175fed6ba88bab65
---

## Motivation

The design's Specs grammar requires one Specs section per module (R03).
Main 3b3b64d ignores a second Specs section, including invalid items inside
it, and check reports success. Validate module structure without changing
the general-purpose first-section contract in CON-002/CON-009.

## Interface

No public signature or type changes.

## Added

### PRJ-077
Check MUST reject a module SPEC.md with zero or more than one level-two Specs section using a non-warning C03 diagnostic naming that file. Heading recognition MUST follow CON-002's fenced-code and ATX-heading rules. A single empty Specs section MUST remain valid.

## Modified



## Removed

## Decisions

## Notes

Add project.Check fixtures for absent, duplicate, single empty and single
populated sections, plus a fenced example containing ## Specs that does not
count. Put invalid text or a duplicate ID in a second real section to ensure
it cannot disappear silently. Keep contracts.Section and ReplaceSection
first-match behavior unchanged; their duplicate-section eval remains valid.
This validates module documents, not an invented required section list for
change files.

Drafted at the owner's request after the R01–R29 traceability review.
No lifecycle execution, agent invocation, drill or release is authorized by
this drafting request. Workers must obtain context before their step and use
fireweed verify; only the controller may advance or commit lifecycle work.

The targeted PRJ-077 Go run recorded fail-first evidence. All five invalid
module fixtures returned no diagnostics where a non-warning C03 was expected.
All six valid fixtures passed. The existing CON-002 and CON-009 section evals
passed, and `fireweed verify single-specs-section` passed in evals.
The implementation can add the module structure check in the disposable
`packages/project/internal/queue/validate.go` file.
## Audit

No findings; checked PRJ-077

## Provenance

- evals 2026-11-14T00:41:56.212423Z agent=codex fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=180 failing=6 passing=174
- audit 2026-11-14T00:42:09.072481Z agent=codex fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
- impl 2026-11-14T01:09:19.488971Z agent=codex fireweed=0.1.0 playbook=5f856d6bcab308b9ff1a8a7c03b9d2547744112aba81dd30dea8f61ec8062b03 tests=1055 failing=0 passing=1055

