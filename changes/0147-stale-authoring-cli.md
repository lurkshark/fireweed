---
already_satisfied:
    - APP-008
bases:
    APP-008: 01a6e117
module: app
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

### new:stale-diff
show --stale <change> MUST discover the nearest project root and query project.StaleSpecs for the pending change selector.
For each recovered entry in returned order, it MUST identify the spec and display a line diff from Before to Current, distinguishing removed and added lines.

### new:stale-fallback
For an unrecovered stale entry, show --stale MUST identify the spec, display its RecordedHash and Current text, and indicate that stamped text was unavailable.
It MUST NOT present an unverified historical text as a recovered diff.

### new:stale-usage
The CLI MUST accept show --stale with exactly one change selector and reject its combination with --history or --projected.
Use of --stale outside show or invalid stale-view arguments MUST return exit 2 with corrective U01 guidance describing valid show --stale usage.

### new:stale-errors
show --stale MUST return exit 2 for change lookup failures, retaining U02 for an unknown selector or C13 for an ambiguous slug and naming the rejected selector.
Project discovery or configuration errors MUST return exit 2 retaining their cause code or C01 when uncoded; other query failures MUST return exit 1 retaining their cause code or P01.

### new:fix-bases
check --fix MUST fill missing Modified and Removed bases across the pending queue using project.StampMissing before ordinary project.Check diagnostics.
Every existing bases entry, including valid and stale hashes, MUST remain unchanged; stale conflicts MUST still produce C08 and exit 1.

### new:fix-results
check --fix MUST use check's existing plain/JSON success and diagnostic shapes, returning exit 0 only when no error remains.
Stamping rejections MUST return exit 1 with their cause code and corrective guidance; project discovery or configuration errors MUST retain exit 2.

### new:fix-usage
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

