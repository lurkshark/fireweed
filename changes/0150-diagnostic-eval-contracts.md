---
bases: {}
eval_fix:
    - APP-107
    - APP-108
    - APP-112
    - APP-113
module: app
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

