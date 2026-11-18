---
already_satisfied:
    - "new:load-module-metadata-code"
module: project
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

### new:load-module-metadata-code
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
