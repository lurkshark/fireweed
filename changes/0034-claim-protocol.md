---
fixes:
    - FLOW-008
    - FLOW-009
    - FLOW-010
module: workflow
phase: audit
run: 288cf71ec47c45aee437c6c6
---

## Motivation

Executable context must trust exact stored claim fields and object tokens. The current claim adapter parses combined Git stdout/stderr, so a successful local command that warns on stderr can corrupt object format/ID/ref/JSON parsing and reject otherwise valid create, update, release, load or inventory calls. Git replacement views can also substitute JSON bytes for the claim blob identified by its stored Object. Add failing public regressions for the existing exact-store/read/CAS contracts before binding context to them.

## Interface

No public interface or data change. Add only permanent public-API regressions; keep all behavior disposable under internal/claims. Reuse the applied gitstore Capture primitive for protocol stdout when implementing the fix.

## Added

## Modified

## Removed

## Decisions

This fixes the existing success and exact stored-field clauses without redefining the five-field wire format, token semantics, errors or ownership rules. Successful Git stderr is diagnostic output, never object-format/ref/object-ID/JSON data. Failure guidance may include captured stderr. Literal metadata operations must continue to suppress unrelated native hooks and preserve the existing local/no-network/temporary-storage contract.

Claim Object identifies the actual content-addressed blob named by the shared ref. Git replacement views must not replace those bytes or their type, even when replacement JSON is otherwise valid and retains the original identity fields. Preserve replacement refs unchanged; disabling their interpretation for owned metadata operations does not delete them or change the user's Git configuration. All operations still return exact full tokens and original State bytes, including CRLF/non-ASCII/whitespace.

## Rejected alternatives

Parsing or stripping warning-looking lines from merged stdout loses byte-exact protocol data and can discard legitimate State content. Returning replacement-view JSON with an original blob token breaks content identity. Requiring users to remove harmless warnings or their replacement refs makes stored claim interpretation depend on unrelated local settings.

## Provenance

- evals | 2026-10-23T00:31:10Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 1098 | failing: 20 | passing: 1078
