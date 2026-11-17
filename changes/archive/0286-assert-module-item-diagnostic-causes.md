---
eval_fix:
    - CHK-048
module: checks
phase: applied
run: 20261116T200508-28b7a066e6581509a04f5869
---

## Motivation

The contracts reconstruction preserved the diagnostic codes and named the
offending heading or ID. The checks eval rejected punctuation and quoting that
CON-017, CON-018 and CHK-048 leave open. D006 allocates stable cause codes so
callers can recognize the failure independently of wording.

## Interface

## Added

## Modified

## Removed

## Decisions

## Notes

Keep both malformed-heading and duplicate-ID cases and CHK-048 coverage.
Assert contracts.ErrorCode C03 and C05 respectively, the explanatory cause,
and the offending heading or ID in the message. Check the cause and subject
text independently of punctuation and quotation marks. Leave other diagnostic
assertions intact. The existing eval fails on the retained contracts build
and passes on current main.

Full implementation-phase `fireweed verify` passed, including all 135
checks cases and both CHK-048 cases.

## Provenance

- evals 2026-11-16T20:15:50.647434Z agent=codex fireweed=0.2.1 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=135 failing=0 passing=135
- audit 2026-11-16T20:16:12.9103Z agent=codex fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
- impl 2026-11-16T20:40:30.620305Z agent=codex fireweed=0.2.1 playbook=5f856d6bcab308b9ff1a8a7c03b9d2547744112aba81dd30dea8f61ec8062b03 tests=966 failing=0 passing=966

## Audit

No findings; checked CHK-048

