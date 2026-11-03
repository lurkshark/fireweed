---
audit_pending:
    - CHK-063
audit_rounds: 1
bases: {}
eval_fix:
    - CHK-063
    - CHK-064
    - CHK-065
    - CHK-066
    - CHK-067
module: checks
phase: audit
run: 20261103T052009-5096bd1be0486adcc6a17841
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

- CHK-063: "Doctor MUST run the worker and audit commands once each in scratch through the configured agent.wrap with a fixed-reply prompt, requiring exit 0 and the requested reply for success." — The eval asserts the fixed-reply prompt on the audit line, but not on the worker line. Its general FIREWEED_DOCTOR_OK log check can be satisfied by the audit prompt alone.
## Provenance
- evals 2026-11-03T05:30:18.999922Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=101 failing=0 passing=101
- audit 2026-11-03T05:32:35.5279Z agent=manual fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
- evals 2026-11-03T05:41:38.739674Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=101 failing=0 passing=101
