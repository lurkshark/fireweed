---
already_satisfied:
    - CHK-001
    - new:junit-malformed
    - new:junit-invalid-ids
    - new:junit-id-mapping
    - new:junit-results
bases:
    CHK-001: 6794875b
module: checks
---



## Motivation

Continue the identified CHK-001 split under SYSTEM's spec-shape rule, preserving
the design's JUnit coverage contract. Separate parsing supported document
shapes from XML/ID rejection and ID/result mapping. CHK-002 already requires
non-skipped coverage, so skipped coverage remains there rather than being
bundled into ParseJUnit's contract. All existing assertions are retained.

## Interface



## Added

### new:junit-malformed
ParseJUnit MUST reject malformed XML.

### new:junit-invalid-ids
ParseJUnit MUST reject testcases with missing or unknown spec IDs.

### new:junit-id-mapping
ParseJUnit MUST map testcase name and classname to current spec IDs using contracts.MatchesID.

### new:junit-results
ParseJUnit MUST classify failure and error elements as fail, and skipped elements as skip.

## Modified

### CHK-001
ParseJUnit MUST accept testsuite or testsuites documents, including nested suites, returning testcase name, classname and time.

## Removed



## Decisions



## Notes

Rename existing junit_test.go evals only, keeping all assertions and 28 checks
cases. NestedJUnitIDsResultsAndTime retains CHK-001 and gains mapping/results
IDs; SingleSuiteSkippedDoesNotCover retains CHK-001 and gains result mapping
and CHK-002 for its existing skipped-coverage assertion. RejectMalformedXML
uses the malformed ID; RejectMissingAndUnknownIDs uses invalid-ID/mapping IDs.
Do not change checkSpecs, XML fixtures, helper behavior, code or interfaces.
No new Test functions, t.Run cases or fixture instances. Implementation may
edit only this change file. Preserve all other coverage labels.
