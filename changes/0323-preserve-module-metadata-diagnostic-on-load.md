---
already_satisfied:
    - PRJ-089
module: project
phase: audit
run: 20261118T221926-515921cad4dae6f5a225516a
---

## Motivation

The November 18 project drill classified invalid depends_on as C12 in its diagnostic
list but returned an uncoded Load error. The modules CLI fell back to C01. APP-018
needs the metadata classification at the loader boundary, and live and recorded
loaders share that classification on main.
[Drill reports](https://github.com/lurkshark/fireweed/pull/25)

## Interface

Specify the existing error code and subject exposed by project.Load and
project.LoadAt.

## Added

### PRJ-089
Load and LoadAt MUST return a C12 error through contracts.ErrorCode for a matched
module's invalid metadata or invalid zone field type. The error's subject MUST
identify the repository-relative SPEC.md path.

## Modified

## Removed

## Decisions

## Notes

Assert the code and subject on the public return value. Use invalid depends_on and
reuse an invalid zone field fixture. Use the committed-spec fixture to confirm
LoadAt reports the recorded invalid metadata even when the working copy has been
repaired. Keep the existing Check and recorded syntax or duplicate item assertions.
The current loaders return C12 with the matching path.
## Provenance

- evals 2026-11-18T22:20:55.783536Z agent=codex fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=325 failing=0 passing=325

