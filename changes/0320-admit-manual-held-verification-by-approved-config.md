---
module: app
---

## Motivation

The current CLI reads live configuration before dispatching verify or done. A
malformed protected configuration edit produces C01 before the workflow can identify
the zone violation. With an owned change, the compatibility gate should check the
same configuration that supplies that step's hooks and zones. This is the same
source-selection issue seen in the runner drill.

## Interface

Use held compatibility admission when verify or done selects an owned change. These
commands still run workflow's ordinary author, phase, zone and live project checks.

## Added

### new:manual-held-verification-config
When verify or done selects a change owned by the caller's worktree, compatibility
admission MUST use that change's held approved configuration without parsing live
fireweed.yaml for admission. An incompatible approved configuration MUST return
exit 2 with C01 and corrective guidance before running hooks or changing phase.

## Modified

## Removed

## Decisions

## Notes

Extend the alternate-binary approved-pin fixture to verify and done. Use a valid
live pin and an incompatible held pin to show the selected claim governs admission.
Check that an unknown selector still reports the selector error in that fixture.
With a compatible held pin and malformed protected live configuration, assert Z01
for fireweed.yaml and zero hook calls or phase commits. Keep the existing live
static and applied-project checks.
