---
bases:
    APP-164: b5a418dc
module: app
---

## Motivation

The current CLI reads live configuration before dispatching verify or done. A
malformed protected configuration edit produces C01 before the workflow can identify
the zone violation. With an owned change, the compatibility gate should check the
same configuration that supplies that step's hooks and zones. This is the same
source-selection issue seen in the runner drill.

## Interface

Extend APP-164's held compatibility admission to verify and done. These commands
still run workflow's ordinary author, phase, zone and live project checks.

## Added

## Modified

### APP-164
While a change is owned, modules, coverage, verify, done, next and run MUST check
compatibility against the held approved configuration used by their operations.
Compatibility admission MUST NOT add live-config parsing to their approved-snapshot
paths or replace live configuration required by other commands with a held snapshot.
An approved-configuration rejection MUST return exit 2 with C01 and corrective
guidance before command effects or hook/agent execution.

## Removed

## Decisions

## Notes

Extend the alternate-binary approved-pin fixture to verify and done. Use a valid
live pin and an incompatible held pin to show the held pin governs admission. With a
compatible held pin and a malformed protected live config, assert the workflow's
zone rejection identifying fireweed.yaml and zero hook calls or phase commits. Keep
the live static and applied-project checks exercised by the existing fixtures.
