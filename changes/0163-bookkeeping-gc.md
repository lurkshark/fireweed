---
bases: {}
module: workflow
---

## Motivation

ACCEPTANCE C5 selects the design's gc behavior: prune stale local claims,
saved attempts for archived changes and expired drill refs. Workflow owns
these bookkeeping namespaces. Add the backend first; normal app and assets
changes will expose the command and document its retention setting.

## Interface

Add GC(ctx context.Context, request GCRequest) (GCResult, error), delegating
through api.go to disposable lifecycle code. GCRequest has Root string and
Now time.Time; zero Now uses current UTC time. GCResult has Pruned []string.
Declare the data in retained model/ and expose aliases from api.go. No behavior
belongs in the retained facade/model. Evals may add a new disposable stub file
for compilation; it remains implementation-writable after interfaces freeze.

## Added

### new:gc-stale-claims
GC MUST prune refs/fireweed/claims entries whose recorded worktree is missing or is no longer registered with the same Git common directory, while retaining claims for live registered worktrees.

### new:gc-archived-attempts
GC MUST prune refs/fireweed/attempts/<slug>/* only for changes present in changes/archive, retaining attempt refs for unapplied changes.

### new:gc-drill-expiry
GC MUST prune refs/fireweed/drills/<module>/<date> when the UTC timestamp encoded in the native drill ref name is older than gc.drill_max_age relative to GCRequest.Now, or current UTC time when Now is zero, retaining refs at or newer than the cutoff.

### new:gc-retention-config
GC MUST read gc.drill_max_age as a positive Go duration, defaulting to 720h (30 days) when unset. An invalid configured value MUST return C01 before any pruning.

### new:gc-result-scope
GC MUST return the sorted full names of only its successfully deleted refs. It MUST leave all other refs, HEAD, index and working files unchanged.

## Modified

## Removed

## Decisions

## Notes

Implement only the selected bookkeeping cleanup. Preserve every existing
assertion and all lifecycle/approval/audit/retry controls. Reuse gitstore's
public local Git/process/ref primitives and compare-and-swap deletion of
observed objects. Native drill dates use UTC 20060102T150405.000000000 in the
ref suffix. Do not guess ages from unrelated commits or run drills to test gc.
Keep ordinary live linked worktrees/claims intact. Do not prune unrelated,
unapplied or unrecognized refs speculatively. No networking, branch changes,
hooks, agents, history rewrites or main-root cleanup belongs in GC.

Run gc behavior only in temporary eval repositories. NEVER invoke it against
the actual project: existing required proof refs and unrelated worktrees must
remain intact. Do not edit other modules, root config, module SPEC/Decisions,
protected setup or bounds. App CLI/config documentation follow after this
backend applies; no public gitstore additions are needed.

Use public-interface evals with deterministic Now inputs and representative
claims/attempts/drills in temporary repositories. Preserve all existing cases,
keep own Verify below five minutes and full impl Verify below twenty-five.
Do not add process instrumentation, exhaustive fault matrices or duplicate
expensive lifecycle fixtures. Fail-first must have real failed testcases for
every Added handle, not compile/collection errors or a complete stub.

Workers run pinned fireweed context, verify with fireweed verify --quick,
poll until terminal and return. They never commit, call lifecycle mutators in
the held checkout, edit run/phase/Provenance or change frozen evals/interfaces.
Fixture API lifecycle calls remain permitted public-interface tests. Only the
outside owning controller advances, commits, applies and archives with normal
full verification.

The separate restricted auditor edits only the existing Audit body, retaining
all other bytes. Findings use - ID: "exact clause" — missing assertion. Passing
audit writes No findings and Checked IDs. Re-audits cover only flagged IDs,
retaining prior clearances and all counters/bounds.

## Audit

