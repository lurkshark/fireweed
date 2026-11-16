---
module: checks
eval_fix: [CHK-048]
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
