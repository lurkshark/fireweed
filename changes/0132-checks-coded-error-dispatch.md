---
module: checks
---

## Motivation

Maintenance after 0128-coded-parse-errors: checks classifies ParseItems failures
by the `duplicate item ID:` message prefix (internal/engine/hooks.go) and doctor
classifies its own C11 failures by matching message text
(internal/engine/doctor.go). Read codes through contracts.ErrorCode and typed
errors instead, so a regenerated contracts or checks module cannot silently
change the diagnostics. Behavior and output stay the same.

## Interface

No interface changes.

## Added

## Modified

## Removed

## Decisions

## Rejected alternatives
