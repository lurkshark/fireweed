---
already_satisfied:
    - APP-067
bases:
    APP-067: 27d57f83
module: app
---

## Motivation

The app's APP-067 eval requires short probe names in `doctor --json`, and
the app takes them from the first word of checks' DoctorCheck names. APP-067
does not mention names, so an app rebuild has to infer that rule from the
eval. The app drill lists this as a gap. 0229 states the checks name format;
state the JSON rule here.

## Interface

None.

## Added

## Modified

### APP-067
doctor --json MUST emit schema 1 with an object containing a checks array on success or null data on failure, retaining code/message/path/warning diagnostics and leaving stderr empty. Each canary check MUST use its probe name as name and keep the hook scope in detail.

## Removed

## Decisions

## Notes

The current app already does this, so APP-067 is listed in
already_satisfied. Run after 0229.
