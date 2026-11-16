---
module: checks
fixes: [CHK-065, CHK-066]
---

## Motivation

Doctor should diagnose the same usage-limit signal users encounter during a
run. On current main a failed command that emits RATE LIMIT on stdout,
followed by 19 stderr lines, is classified as an ordinary failure by Doctor.
InvokeAgent classifies the same output as a usage limit. Doctor inserts a
blank line at the stream boundary and drops the oldest actual line from its
20-line window.

## Interface

## Added

## Modified

## Removed

## Decisions

## Notes

Add focused Doctor cases with one matched stdout line followed by 19 and
20 stderr lines. Use a nonzero exit for both. Assert C15 usage-limit
classification when the match is twentieth from the end, then ordinary
failure and exclusion when the match is twenty-first from the end. Keep the
captured streams and full probe logs intact. Count real emitted lines,
including deliberate blank lines. Use the existing stdout-then-stderr order
for independently captured streams.
