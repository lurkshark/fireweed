---
already_satisfied:
    - CHK-052
bases:
    CHK-052: ce8f7780
module: checks
phase: audit
run: 20261116T222202-3681a0b47fed5d05c652e86e
---

## Motivation

When doctor stops before the probes, the setup hook's stderr can explain
which prerequisite a user needs to fix. The latest checks drill passed
its own doctor evals after 0285 but failed the APP-065/066/067 setup-failure
case because it reported only the process exit status. CHK-052 requires
C14 and the hook name and leaves the failure detail unspecified.

## Interface

## Added

## Modified

### CHK-052
Doctor MUST run the configured setup hook in scratch before probes. A setup
failure MUST produce C14 naming the setup hook and including any captured
stderr. It MUST prevent later probes.

## Removed

## Decisions

## Notes

Use a setup script that exits nonzero and writes a distinctive message to
stderr. Put the marker inside the script body so it cannot be echoed by displaying
the configured invocation. Assert the C14 diagnostic names setup, contains the
marker and no later hook or agent runs. Keep the scratch behavior and coverage
from the existing CHK-052 eval. Checks already reads the captured stream
from gitstore.Result when building its setup failure report.
## Provenance

- evals 2026-11-16T22:31:06.421602Z agent=codex fireweed=0.2.1 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=135 failing=0 passing=135

