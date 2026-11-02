---
already_satisfied:
    - PRJ-015
    - PRJ-025
    - PRJ-055
    - PRJ-056
    - PRJ-067
    - PRJ-068
bases:
    PRJ-015: 12d385ec
    PRJ-025: e5686358
    PRJ-055: ab809b96
    PRJ-056: adc7f8db
module: project
phase: applied
run: 20261102T021403-224e6f0fcfb6b2f14f07be9d
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

### PRJ-067
When NewChange cannot create its queued file, it MUST return an error beginning
with `create <queued path>:` and retaining the filesystem cause, where the queued
path is repository-relative.

### PRJ-068
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

## Provenance
- evals 2026-11-02T02:21:52.817484Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=127 failing=0 passing=127
- audit 2026-11-02T02:22:29.445391Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
- impl 2026-11-02T02:30:05.307767Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=5f856d6bcab308b9ff1a8a7c03b9d2547744112aba81dd30dea8f61ec8062b03 tests=812 failing=0 passing=812
## Audit

No findings; checked PRJ-015, PRJ-025, PRJ-055, PRJ-056, PRJ-067, PRJ-068
