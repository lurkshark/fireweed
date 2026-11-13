---
module: project
---

## Motivation

The design's Specs grammar requires one Specs section per module (R03).
Main 3b3b64d ignores a second Specs section, including invalid items inside
it, and check reports success. Validate module structure without changing
the general-purpose first-section contract in CON-002/CON-009.

## Interface

No public signature or type changes.

## Added

### new:single-module-specs-section
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
