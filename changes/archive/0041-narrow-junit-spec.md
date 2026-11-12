---
already_satisfied:
    - CHK-001
    - CHK-027
    - CHK-028
    - CHK-029
    - CHK-030
audit_rounds: 0
bases:
    CHK-001: 6794875b
module: checks
phase: applied
run: 20261028T053206-0fd8210f3a61dc9c5c2d10c5
---



## Motivation

Continue the identified CHK-001 split under SYSTEM's spec-shape rule, preserving
the design's JUnit coverage contract. Separate parsing supported document
shapes from XML/ID rejection and ID/result mapping. CHK-002 already requires
non-skipped coverage, so skipped coverage remains there rather than being
bundled into ParseJUnit's contract. All existing assertions are retained.

## Interface



## Added

### CHK-027
ParseJUnit MUST reject malformed XML.

### CHK-028
ParseJUnit MUST reject testcases with missing or unknown spec IDs.

### CHK-029
ParseJUnit MUST map testcase name and classname to current spec IDs using contracts.MatchesID.

### CHK-030
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

Implementation already satisfies the projected split; no code changes were
needed. Quick verification passed all 28 checks cases.

## Provenance

- evals 2026-10-28T05:32:51.870903Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=146 failing=0 passing=146
- audit 2026-10-28T05:33:28.948786Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0
- impl 2026-10-28T05:34:21.276553Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=07554a3249d8fca0ba70df7504135f297491465caa63458734a4b8351734d0e5 tests=146 failing=0 passing=146

## Audit

No findings. Checked IDs: CHK-001, CHK-002, CHK-027, CHK-028, CHK-029, CHK-030.
