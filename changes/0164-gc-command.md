---
bases: {}
module: app
---

## Motivation

ACCEPTANCE C5 selects the design's local bookkeeping cleanup. The workflow
backend now supports gc; expose it through the CLI with normal discovery,
plain and JSON results, and actionable diagnostics. Retention documentation
follows in assets before C6 begins.

## Interface

No public signatures change. Use workflow.GC with the discovered project Root
and zero Now so the backend uses current UTC time.

## Added

### new:gc-result
gc MUST discover the nearest project root, invoke workflow.GC and return exit 0 on success. Plain output MUST list the sorted full pruned ref names or explicitly report that no refs were pruned.

### new:gc-json
Successful gc --json MUST emit schema 1 object data containing a pruned array of sorted full deleted ref names, including an empty array when nothing was pruned, with empty diagnostics and stderr.

### new:gc-failures
gc MUST return exit 2 for project discovery or configuration failures, retaining the cause code or C01 when uncoded. Other cleanup failures MUST return exit 1 retaining the cause code or P01 when uncoded. Failures MUST give corrective diagnostics, with null data and empty stderr in JSON mode.

## Modified

## Removed

## Decisions

## Notes

Accept no positional arguments or command-specific options other than --json;
retain existing U01 usage rejection before mutation. Exercise all cleanup only
in temporary eval repositories, NEVER this actual project. Preserve main proof
refs and unrelated worktrees. Reuse the backend instead of duplicating pruning
policy. Keep configuration validation before pruning and live claims intact.

Preserve every existing assertion. APP-014's unfinished-command representative
is gc today; move it to the still-unimplemented migrate command without dropping
its help-listing, nonzero exit, U01, explicit-message and JSON assertions. This
is the intentional gc product change, not an eval_fix. Preserve every other
case and all logging, prompt transport and audit controls. Do not edit other
modules, protected setup, module SPEC/Decisions, configuration bounds or timeouts.

Use public-interface CLI evals and deterministic local fixtures. Cover nested
root discovery, sorted pruned refs and a second empty result, both output modes,
usage before mutation, malformed config/invalid retention and a normal Git
precondition failure. Keep fixtures representative and own Verify under five
minutes; full implementation Verify remains under twenty-five minutes.

Workers run pinned fireweed context, verify with fireweed verify --quick,
poll until terminal and return. They never commit, call lifecycle mutators in
the held checkout, edit run/phase/Provenance or change frozen evals/interfaces.
Only the outside owning controller advances, commits, applies and archives
with normal full verification. Fixture API calls remain allowed.

The separate restricted auditor edits only the existing Audit body. Findings
use - ID: "exact clause" — missing assertion. Passing audit writes No findings
and Checked IDs. Re-audits cover only flagged IDs, retaining prior clearances
and all counters/bounds.

