---
already_satisfied:
    - APP-008
bases:
    APP-008: 01a6e117
module: app
phase: applied
run: 20261101T195004-729a090e3389f1b64d28cef5
---

## Motivation

The design's Fixing a stale change clause requires show --stale <change> to
explain what moved beneath stamped edits and check --fix to fill missing bases
without discarding conflicts. Project's applied 0146 supplies StaleSpecs and
StampMissing; expose them through the actual NEW CLI. Keep main's original
controller for workflow control, and preserve existing command behavior.

## Interface

No public Go signatures, metadata, hooks or configuration changes. Extend the
existing app.Run CLI with show --stale <change> and check --fix using project's
retained APIs. Existing plain/JSON envelope conventions continue to apply.

## Added

### APP-143
show --stale <change> MUST discover the nearest project root and query project.StaleSpecs for the pending change selector.
For each recovered entry in returned order, it MUST identify the spec and display a line diff from Before to Current, distinguishing removed and added lines.

### APP-144
For an unrecovered stale entry, show --stale MUST identify the spec, display its RecordedHash and Current text, and indicate that stamped text was unavailable.
It MUST NOT present an unverified historical text as a recovered diff.

### APP-145
The CLI MUST accept show --stale with exactly one change selector and reject its combination with --history or --projected.
Use of --stale outside show or invalid stale-view arguments MUST return exit 2 with corrective U01 guidance describing valid show --stale usage.

### APP-146
show --stale MUST return exit 2 for change lookup failures, retaining U02 for an unknown selector or C13 for an ambiguous slug and naming the rejected selector.
Project discovery or configuration errors MUST return exit 2 retaining their cause code or C01 when uncoded; other query failures MUST return exit 1 retaining their cause code or P01.

### APP-147
check --fix MUST fill missing Modified and Removed bases across the pending queue using project.StampMissing before ordinary project.Check diagnostics.
Every existing bases entry, including valid and stale hashes, MUST remain unchanged; stale conflicts MUST still produce C08 and exit 1.

### APP-148
check --fix MUST use check's existing plain/JSON success and diagnostic shapes, returning exit 0 only when no error remains.
Stamping rejections MUST return exit 1 with their cause code and corrective guidance; project discovery or configuration errors MUST retain exit 2.

### APP-149
The CLI MUST accept check --fix without positional arguments and reject --fix with other commands.
Invalid fix-command usage MUST return exit 2 with corrective U01 guidance describing valid check --fix usage before mutation.

## Modified

### APP-008
Without --history or --stale, show <module> [ID] MUST return the current module document or selected item, using the whole-queue project projection with --projected.

## Removed

## Decisions

## Notes

APP-008 only narrows its command-mode scope and legitimately uses
already_satisfied: preserve every existing case, assertion and ID. Preserve
existing APP-025 and APP-015 string-data/error envelopes for stale views and
APP-024 check envelopes for fix requests. Stale-view success exits 0; JSON uses
identical plain text as schema-1 string data, empty diagnostics and empty
stderr. Failures use null JSON data and corrective code/message diagnostics;
plain failures have empty stdout and diagnostics on stderr.

Use actual explicitly built NEW CLI fixtures for representative historical
recovery, unavailable and mismatched-history fallback, aliases, missing and
ambiguous selectors, invalid arguments/options, and applicable configuration
and stamping rejections. Recovery must distinguish an unrelated later change
file edit from the last bases-changing commit, include tagged normalized text,
and compare to the preceding queue projection rather than the target's proposed
edit. Do not query or substitute the pinned old controller for the new options.

Evals use a freshly built NEW CLI and fail first against the current implementation:
the new options return U01 before their stale-view and fix behaviors can run.
Implementation can change the disposable packages/app/internal/entry/entry.go
parser and add command handlers under packages/app/internal/entry/; the public
interface and evals remain frozen after this phase.
Controller `verify --quick` passed in evals phase; only the pre-existing
APP-103/104 B02 case-count warnings were reported.

For fix, exercise missing/valid/stale bases together, Modified and Removed,
multiple queued changes, preserved extra bases/frontmatter/human text, a
missing-only success, and a repeated no-op. Confirm ordinary check stays
read-only and fix changes only queued files with missing bases; show preserves
repository bytes, index and refs. Reuse existing public fixtures and build a
NEW binary once per new test group where practical. No redundant Cartesian
matrices, process instrumentation, new dependencies or broad helper refactoring.

Run context before each step. Use fireweed verify --quick for worker feedback;
the owning controller still runs full verification and boundaries. Workers
never commit, call done or advance phases. Keep restricted independent audit,
frozen interfaces/evals and all configured bounds; SYSTEM's owner-updated
full-verification process guardrail is 25 minutes, with own Verify under 5.

## Provenance
- evals 2026-11-01T20:08:19.104116Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=304 failing=13 passing=291
- audit 2026-11-01T20:09:23.21348Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
- impl 2026-11-01T20:26:10.329834Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=5f856d6bcab308b9ff1a8a7c03b9d2547744112aba81dd30dea8f61ec8062b03 tests=304 failing=0 passing=304
## Audit

No findings; checked APP-008, APP-143, APP-144, APP-145, APP-146, APP-147, APP-148, APP-149
