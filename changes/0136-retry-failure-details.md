---
bases: {}
module: runner
---



## Motivation

During 0129, workers passed quick own-module checks while repeated full checks
failed in a dependent module. Manual focused runs exposed the actual assertions.
NEW already forwards verification errors and hook streams, but the assertions
can exist only in JUnit. Consume Case.Failure from 0135 so the next attempt sees
the failed module, testcase and assertion without reconstructing them from logs.
Close this slice with NEW next in a temporary fixture whose fake worker saves
its retry prompt, demonstrating the JUnit-only assertion and bounded excerpt.

## Interface

No public API changes. Next's worker retry prompt gains a bounded failed-testcase
detail section using checks.Case.Failure from 0135-junit-failure-details.

## Added

### new:retry-case-assertions
On a verification retry, Next MUST append a detail section identifying the module, classname, testcase name and Case.Failure for failed cases from every reported module, including dependents, subject to new:retry-detail-limits.
It MUST select cases by lexical module name and then report order, excluding passing and skipped cases.

### new:retry-detail-limits
The failed-case detail section MUST fit within 16,384 Unicode code points, including labels and truncation markers.
Each Case.Failure excerpt MUST contain at most its first 4,096 code points.
Shortened diagnostic text or omitted cases MUST be marked.

## Modified



## Removed



## Decisions



## Rejected alternatives

Requiring repeated full-suite runs just to discover the assertion: deliver the
failure already produced by verification to the retry worker.

## Notes

Depends on the checks report field introduced by 0135. Existing workflow.Result
already carries checks reports, so no workflow API or implementation change is
needed. Preserve RUN-022's existing verification feedback and hook streams;
limits apply only to the new failed-case section. Keep attempt counts, retained
work, usage/time stops and audit restrictions unchanged.

Use the existing bounded fake-agent and local JUnit-hook fixtures. Include an
assertion sentinel in JUnit that does not appear in hook stdout/stderr, including
a dependent-module failure. Check readable bounded excerpts and truncation;
headings need not have an exact prescribed wording. No real agent or network.
