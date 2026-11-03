---
bases: {}
eval_fix:
    - CHK-063
    - CHK-064
    - CHK-065
    - CHK-066
    - CHK-067
module: checks
---

## Motivation

ACCEPTANCE C1 requires every prompt-bearing fixture script to use bash or
receive its prompt through stdin, accepting the supplied macOS sh diagnosis.
The checks doctor fixtures still pass {prompt} to sh worker/auditor scripts and
a wrapper. Correct those fixture commands while preserving every existing
assertion, duration, coverage obligation and product behavior. This is the
remaining checks-owned C1 fixture work, not the deferred CK milestone.

## Interface

None. This eval_fix changes no specs or interfaces and needs no new handles.

## Added

## Modified

## Removed

## Decisions

## Notes

Change only the checks eval fixtures' prompt-bearing fake-agent and wrapper
interpreters/transport, including the shell-string audit-fallback command.
Retain shell-string versus argv coverage, quoting, all testcase names, assertions,
configured timeouts and hook canary/recovery behavior. Do not change hooks that
receive no prompt merely for consistency. Do not investigate the OS diagnosis.

Evals/impl workers run the current pinned fireweed context before their step,
use fireweed verify --quick for feedback and return. Poll a live verification
session until terminal; an observation timeout alone is no reason to restart it.
Only the outside owning controller advances phases, commits or archives and
performs required normal full verification. Never edit run/phase/Provenance.

The separate restricted auditor edits only the body of the existing Audit
section, retaining every byte outside it, including whitespace. Its review must
confirm the corrected evals preserve assertions of all listed spec clauses.

## Audit

