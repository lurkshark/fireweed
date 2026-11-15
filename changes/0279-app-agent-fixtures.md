---
eval_fix: [APP-063, APP-074, APP-080, APP-136]
module: app
---

## Motivation

The CLI integration fixtures encode agent commands inline. Runner and checks
will switch to named commands in separate changes, and their implementation
verification runs the app evals. Give those fixtures equivalent named entries
before either dependency switches its command lookup.

## Interface

## Added

## Modified

## Removed

## Decisions

## Notes

In the doctor, next, run, and drill CLI fixtures, define named entries for the
same fake commands and retain the inline entries until both dependencies have
switched. For doctor, prepare these entries in the fixture config callback.
Keep the existing CLI assertions, fake commands, and fixture bounds. The
app CLI change removes the temporary inline fixture entries.
