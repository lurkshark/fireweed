---
bases: {}
eval_fix:
    - APP-107
    - APP-108
    - APP-112
    - APP-113
module: app
phase: applied
run: 20261102T023131-6d7ad8ebd7e0e1fc2fbde48e
---

## Motivation

The owner authorized normal audited corrections for the failed independent
project drill's diagnostic contracts. Its app evals reject valid corrective
instructions containing "Correct", "remove" or "rename" because the shared
corrective helper omits those verbs. A check-commits fixture also requires the
incidental text "parse packages/alpha/SPEC.md" before checking the CLI, although
APP-107/108 require the assigned cause code and actionable output, not that
producer wording. Repair those frozen eval assumptions through eval_fix while
preserving their behavioral assertions and every existing case and covered ID.

## Interface

No API, runtime implementation, spec, configuration or module metadata changes.

## Added

## Modified

## Removed

## Decisions

## Notes

Keep the requirement for a real corrective action: extend corrective() only to
recognize the meaningful verbs correct/remove/rename in addition to its existing
terms. Keep the malformed recorded-SPEC fixture's failure assertion, checking its
assigned C03 through public contracts.ErrorCode or a `C03:` error prefix, together
with its module path, instead of requiring the word "parse". This fixture checks
the cause, without imposing one error representation. Preserve all subsequent CLI exit, code, JSON,
stderr, cause and guidance assertions. Do not change application implementation
or other modules. Use fireweed verify --quick for worker feedback; the controller
owns full checks, restricted audit, phase commits and archive.

Worker phase control: do not invoke done, start, release, next, run, abandon,
git commit or git reset, and do not edit run/phase fields or Provenance. Use
verify --quick for feedback and return. The outside owning controller alone
performs every phase transition and commit.

## Provenance
- evals 2026-11-02T02:48:15.408863Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=304 failing=0 passing=304
- audit 2026-11-02T02:48:57.700527Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
- impl 2026-11-02T02:58:56.793981Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=5f856d6bcab308b9ff1a8a7c03b9d2547744112aba81dd30dea8f61ec8062b03 tests=304 failing=0 passing=304
## Audit

No findings; checked APP-107, APP-108, APP-112, APP-113
