---
audit_rounds: 0
bases: {}
module: checks
phase: audit
run: 20261031T065710-5addb41e87d5cd80494011f0
---



## Motivation

The 0129 recoveries required manual focused test runs to recover the assertion
behind a dependent failure. The product's JUnit parser currently retains only
failure presence, discarding the message and body that the Go hook records.
The owner requested preserving those details for retry workers. This dependency
makes them available through checks reports; 0136 returns immediately to the
caller and closes the slice with NEW next and an inspectable fake-worker retry
prompt containing the failed assertion.

## Interface

Add an exported Failure string field to checks.Case, whose retained definition
is model.Case. It carries human-readable JUnit failure and error message/body
text. Existing function signatures and other Case/Report fields stay unchanged.
Implementation belongs under internal/; the evals step adds the retained field.

## Added

### CHK-079
ParseJUnit MUST retain the message attributes and text bodies of JUnit failure and error elements in each affected Case.Failure, including message-only and body-only diagnostics.
A testcase without such text MUST have an empty Failure.

### CHK-080
Evaluate MUST retain parsed Case.Failure strings in its returned Report.Cases, including when failing testcases make the eval hook exit nonzero.

## Modified



## Removed



## Decisions



## Rejected alternatives

Relying on hook stdout: the Go hook writes testcase details into JUnit, while
its ordinary stdout can contain only package summaries.

## Notes

This is the checks dependency for 0136-retry-failure-details. Preserve existing
result classification, ID mapping, coverage, fail-first and exit-status checks.
Retain diagnostic text in reports; excerpt limits belong to the retry prompt
consumer in 0136. Reuse public ParseJUnit/Evaluate fixtures with literal JUnit
and deterministic local hooks. Keep the change within checks.

## Provenance

- evals 2026-10-31T07:26:18.166228Z agent=manual fireweed=0.1.1 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=610 failing=3 passing=607
