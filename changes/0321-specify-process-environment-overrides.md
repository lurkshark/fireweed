---
already_satisfied:
    - "new:process-environment-overrides"
    - "new:process-executable-path"
module: gitstore
---

## Motivation

The November 18 gitstore reconstruction treated Request.Env as a replacement
environment, removing PATH from invoked hooks. Applying it as overrides restored all
52 dependent cases. A public Run probe also distinguished child PATH overrides from
lookup of a bare executable name. Both behaviors affect how callers launch tools.
[Drill reports](https://github.com/lurkshark/fireweed/pull/25)

## Interface

Specify the existing gitstore.Run semantics for Request.Env and bare executable
lookup.

## Added

### new:process-environment-overrides
Run MUST inherit the caller's environment and apply Request.Env as per-variable
overrides to the launched process. Omitted keys MUST retain their inherited values.

### new:process-executable-path
When Args[0] is a bare executable name, Run MUST use the caller's PATH for
executable lookup. A PATH override in Request.Env MUST apply to the launched
process's environment.

## Modified

## Removed

## Decisions

## Notes

Use t.Setenv and fixture-local values. For an explicit override, assert both that
the overridden value wins and an unrelated inherited sentinel survives. Exercise a
nil or empty override map in the same fixture. For lookup, use executable fixture
files with the same bare name in two directories and distinct output markers. Set
the parent PATH and pass a different child PATH, then assert which executable ran
and what PATH it observed. The current launcher already exhibits both behaviors.
