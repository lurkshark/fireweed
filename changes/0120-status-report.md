---
bases: {}
module: workflow
---



## Motivation

The design Reading surface and blocking, approval, cost and drill-drift clauses
require status to explain the queue and expose project state. NEW rejects
status. Recorded costs are available from0119; this missing read model is the
next dependency before the immediate app caller and actual NEW proof.

Owner approved /tmp/fireweed-0120-frozen-phase-checkpoint.md. Native abandon
cee4dbc canceled the former0120 after its frozen maintenance-phase assertion
expected evals instead of the recorded impl. Re-author the same contract with
that one-token public correction and all prior FLOW-113 assertion fixes.
No discarded implementation or old worker-log reuse. After this application
(count4), add scoped B01–B03 Explain at count5, review/checks/push only on PASS,
then return immediately to app status/NEW proof and stop with that caller local.

## Interface

Add Status(root string) (StatusReport, error). StatusReport contains
Modules []project.Module, Queue []StatusEntry, Quarantines []Quarantine,
BoundaryOptOuts []BoundaryOptOut, DriftSinceDrill map[string]int,
Unapproved int, PhaseCommits map[string]int, Costs []ModuleCost and
Diagnostics []project.Diagnostic. StatusEntry contains Slug, Path, Module,
Phase, Blocked, WaitingOn and WaitingRule strings plus AuditRounds int.
Quarantine contains Module, ID and Reason strings; BoundaryOptOut contains
Module and Reason strings. Thin public aliases/wrapper; shared data in model/.
WaitingOn identifies the earlier change by its numbered slug. WaitingRule is
same_module, depends_on_blocked or interface_dependency. No other public edits.

## Added

### new:status-queue
Status MUST return every project.Queue entry in queue order with slug, path, module, phase, blocked reason and audit rounds.
Absent phase, reason and rounds MUST appear as queued, empty string and zero respectively.

### new:status-wait-same-module
For an unblocked queued change targeting an earlier blocked change's module, Status MUST name that earlier change in WaitingOn with WaitingRule same_module.

### new:status-wait-dependent
For an unblocked queued change whose module depends directly or transitively on an earlier blocked change's module, Status MUST name that earlier change with WaitingRule depends_on_blocked.

### new:status-wait-interface
For an unblocked queued change with a nonempty Interface section, Status MUST name an earlier blocked change with WaitingRule interface_dependency when that blocked module depends directly or transitively on the candidate's module.
An empty Interface MUST not trigger this rule.

### new:status-wait-projection
Status MUST determine waiting reasons using the graph projected through each candidate, including preceding and candidate set edits.
Blocked entries and unblocked entries unaffected by the blocking rules MUST have empty WaitingOn and WaitingRule.

### new:status-modules
Status MUST expose the discovered modules with their declared identities, dependency graph and resolved zones.

### new:status-quarantines
Status MUST list every currently quarantined spec with its module, ID and quarantine reason.

### new:status-boundary-opt-outs
Status MUST list every module with boundaries_opt_out and its authored reason.

### new:status-drill-drift
Status MUST report each module's count of applied archived changes whose implementation Provenance date is later than its latest drill report date.
With no drill report, it MUST count all that module's applied archived changes, including zero when there are none.

### new:status-unapproved
Status MUST count applied changes archived on HEAD whose change file is absent from the locally known configured remote/base ref, falling back to the local base when the remote ref is absent.
The base and remote MUST use the existing configured/default base-ref resolution.

### new:status-repeated-module
Status MUST warn with P08 when the three or more most recently applied changes consecutively target one module, identifying the module and consecutive count.

### new:status-phase-counts
Status MUST report the number of phase-changing commits reachable from HEAD for each applied archived change, matching Commits for that change.

### new:status-recorded-costs
Status MUST expose the latest retained ModuleCost for every module with recorded costs, together with the saved budget diagnostics associated with those selected measurements.
Modules without a cost record MUST not receive a fabricated measurement.

## Modified



## Removed



## Decisions

When several earlier changes hold one entry, report the first in queue order,
with same-module, dependent and interface-dependency precedence. Applied order
uses implementation Provenance dates, newest first with archive path as the
tie-breaker; a drill report date counts regardless of result. Budget diagnostics
from selected snapshots are returned once each, rather than recomputed from
later settings or source edits. PhaseCommits keys are archived change filenames
without .md.

## Rejected alternatives

No hooks to collect status, second graph/state directory, extra diagnostics,
regeneration/drill execution or whole-library completion.

## Notes

Owner: "Approve recovery and revised slice". Native cancellation freed the
unapplied IDs; let Start allocate IDs and a new run normally. Historical audit
round1/FLOW-113 findings remain in the old checkpoint report. Fresh workers and
full independent assertion audit apply here; this is not a manual counter or
phase reset. No phase exception, direct refs, history rewrite or drill.

Restore only public inputs from /tmp/fireweed-0120-public-inputs: exact corrected
status evals, described interface/model. Preserve every other assertion and the
beta-presence/latest-warning corrections. The maintenance fixture records impl;
Status must report its actual stored phase. New implementation comes from a
fresh fixed Sol/high worker, with no canceled/deleted implementation or old logs.

One module, thin api/model interfaces, current public dependencies; no protected
or other-module implementation changes. Workers never commit/advance phases or
write claims. Normal max2 audit rounds, re-audit only prior finding IDs. Own
Verify under5m/full under15m, representative fixture hooks/local Git only,
no process instrumentation or increased timeouts. Read conventions remain
review plus the existing single app PATH-shim eval, extended by the caller.
