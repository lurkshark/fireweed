---
already_satisfied:
    - PRJ-015
    - PRJ-025
    - PRJ-055
    - PRJ-056
    - new:create-change-cause
    - new:document-diagnostic-path
bases:
    PRJ-015: 12d385ec
    PRJ-025: e5686358
    PRJ-055: ab809b96
    PRJ-056: adc7f8db
module: project
---

## Motivation

The design's Regeneration section requires owner-reviewed drill gaps to become
explicit retained contracts. The owner authorized recovery of the failed project
drill's diagnostic contracts and one fresh drill after affected checks pass.
All 126 own cases passed while dependent CLI/workflow cases exposed missing
details about error prefixes, actionable causes, diagnostic paths and blank
bases. Current main already supplies these behaviors; preserve them through
legitimate already_satisfied clarification and public evals rather than replacing
main with the drill implementation. These details support D006's existing causes
and preserve the callers' current observable behavior.

## Interface

No public signature, type, configuration or module metadata changes.

## Added

### new:create-change-cause
When NewChange cannot create its queued file, it MUST return an error beginning
with `create <queued path>:` and retaining the filesystem cause, where the queued
path is repository-relative.

### new:document-diagnostic-path
When Check rejects syntax or duplicate identities in a module SPEC.md, the
failure diagnostic MUST identify that document's repository-relative path in
Diagnostic.Path.

## Modified

### PRJ-015
Check MUST reject Modified or Removed entries whose bases hash is missing or
differs from the preceding projected spec's canonical hash with diagnostic C08.
An absent, null or empty-string entry MUST report `missing base for <ID>`;
a nonempty differing hash MUST report `stale base for <ID>`.

### PRJ-025
NewChange MUST reject an unknown target module with an error beginning
`unknown module ` and naming the rejected module.

### PRJ-055
An unknown module or an ID absent from both current module Specs and matching
archive edits MUST return an error beginning `U02: ` and naming the rejected name.
A current ID with no matching archive edits MUST return an empty slice.
An ID in matching archive edits MUST remain queryable when it no longer exists
in current Specs.

### PRJ-056
Missing or malformed fireweed.yaml MUST produce an error beginning `C01: `.
Malformed archived change documents or malformed Added/Modified item syntax
MUST produce an error beginning `C03: ` and identifying the archive path.

## Removed

## Decisions

## Notes

Preserve every existing case, assertion and covered ID. Extend existing public
cases where practical: missing-base cases need null/empty-string variants;
history errors need the colon-and-space prefix assertions; module parse errors
need Diagnostic.Path; NewChange needs an owned blocked-output-path fixture.
All listed items are already satisfied by main. No runtime implementation change
or interface change is intended. Use fireweed verify --quick for worker feedback;
the controller owns full checks, restricted audit, phase commits and archive.

