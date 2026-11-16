---
module: app
---

## Motivation

APP-103 requires a corrective diagnostic for invalid module templates before
creating files. During the project drill, project.Load reported success with
an empty module list and new-module panicked in validateModuleTemplate.
Guarding the expected single result lets the command fail with a diagnostic
if template discovery violates that expectation again.

## Interface

## Added

## Modified

## Removed

## Decisions

## Notes

Check the template discovery result has exactly one module before indexing
it. On unexpected cardinality, return a template-validation error through
the existing corrective C03 response. Keep APP-100/101/103/104 and APP-159
behavior and exercise the existing valid and malformed template evals.
Current project.Load returns one module or an error for the temporary
dot-glob fixture. The cardinality failure was reproduced against the
retained project reconstruction. This maintenance draft adds a local guard
to the existing behavior.
