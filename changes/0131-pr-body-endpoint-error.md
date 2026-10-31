---
module: workflow
---

## Motivation

App pr-body decides between exit 2 with configuration advice and exit 1 with
range advice by matching `load endpoint project at ` in the PRBody error. The
design's distinct exit codes (2 for configuration errors) then depend on
unspecified message text. Make the endpoint configuration failure identifiable.

## Interface

Add `IsEndpointConfigError(err error) bool` to the workflow public API.

Keep the public facade delegating to disposable implementation. During evals,
add the implementation stub in a new internal file so the impl step can replace
it without editing the frozen interface (the wiring repair from 0128).

## Added

### new:endpoint-config-error
PRBody MUST return an error for which IsEndpointConfigError reports true when the project at the range endpoint fails to load.

### new:endpoint-other-errors
IsEndpointConfigError MUST report false when PRBody fails because the range cannot be resolved.

## Modified

## Removed

## Decisions

## Rejected alternatives
