---
already_satisfied:
    - PRJ-025
    - PRJ-029
    - PRJ-050
    - PRJ-060
bases:
    PRJ-025: cc04908f
    PRJ-029: cdf6f8bd
    PRJ-050: 9291801b
    PRJ-060: e2b83fcd
module: project
phase: impl
run: 20261116T213742-336c0c4cba00f9d0358d04ae
---

## Motivation

The CLI needs a stable code and an actionable location when users select a
missing module or a malformed project document. D006 assigns U02 to unknown
names, C03 to document syntax, C05 to duplicate identities and C08 to bases.
In the project drill,
new-change lost U02, recorded SPEC and live queued-change parse errors lost
their paths, and show --stale reported no stale specs for a nonmapping bases
value. Current project code already supplies all four outcomes.

## Interface

## Added

## Modified

### PRJ-025
NewChange MUST reject an unknown target module with contracts.ErrorCode U02
and an error beginning `unknown module ` and naming the rejected module.

### PRJ-029
Queue and Project MUST reject a pending change document when its frontmatter
or items fail to parse. The error MUST preserve the parser's code and use
the document's repository-relative path as the contracts.ErrorCode subject
and in the message.

### PRJ-050
LoadAt MUST return an error when the revision cannot be resolved, its recorded
fireweed.yaml is missing or malformed, or a matched recorded module directory
lacks a readable SPEC.md. When a recorded SPEC.md fails to parse, the error
MUST preserve the parser's code and use the document's repository-relative
path as the contracts.ErrorCode subject and in the message.

### PRJ-060
StaleSpecs MUST return the target change's Modified and Removed items whose
recorded bases differ from the corresponding preceding projected spec hashes,
in edit order, with ID, RecordedHash and Current text. Items without a bases
entry MUST be omitted. A nonmapping bases value MUST produce C08 naming the
queued path.

## Removed

## Decisions

## Notes

Use public project evals to check U02 for an unknown NewChange target with
files preserved, and C08 with the selected path for StaleSpecs with
nonmapping bases. For Queue and Project and for LoadAt, cover malformed
frontmatter (C03) and duplicate item identities (C05). Assert the exact
repository-relative path in the contracts.ErrorCode subject and its presence
in the diagnostic message. Use a valid working copy for the recorded SPEC
cases. CON-017 through CON-019 define parser codes. Preserve the existing
PRJ-050 and PRJ-060 coverage. CLI evidence includes APP-019, APP-038, APP-107,
APP-108, APP-146 and APP-160.
## Audit

No findings; checked PRJ-025, PRJ-029, PRJ-050, PRJ-060

## Provenance

- evals 2026-11-16T21:41:03.168083Z agent=codex fireweed=0.2.1 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=317 failing=0 passing=317
- audit 2026-11-16T21:41:21.491964Z agent=codex fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0

