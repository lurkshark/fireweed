---
audit_pending:
    - FLOW-113
audit_rounds: 1
bases: {}
module: workflow
phase: impl
run: 20261030T144326-24f855074c2173810865ffc2
---



## Motivation

The design's Reading surface and blocking, approval, cost and drill-drift
clauses require status to explain the queue and expose project state. NEW
currently rejects status. The owner approved workflow cost recording (0119,
complete), this status read model, then immediate app wiring and an actual NEW
proof. Build only that caller's missing data, reusing the existing graph,
eligibility rules, archive/provenance, phase history and recorded cost snapshots.

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

### FLOW-101
Status MUST return every project.Queue entry in queue order with slug, path, module, phase, blocked reason and audit rounds.
Absent phase, reason and rounds MUST appear as queued, empty string and zero respectively.

### FLOW-102
For an unblocked queued change targeting an earlier blocked change's module, Status MUST name that earlier change in WaitingOn with WaitingRule same_module.

### FLOW-103
For an unblocked queued change whose module depends directly or transitively on an earlier blocked change's module, Status MUST name that earlier change with WaitingRule depends_on_blocked.

### FLOW-104
For an unblocked queued change with a nonempty Interface section, Status MUST name an earlier blocked change with WaitingRule interface_dependency when that blocked module depends directly or transitively on the candidate's module.
An empty Interface MUST not trigger this rule.

### FLOW-105
Status MUST determine waiting reasons using the graph projected through each candidate, including preceding and candidate set edits.
Blocked entries and unblocked entries unaffected by the blocking rules MUST have empty WaitingOn and WaitingRule.

### FLOW-106
Status MUST expose the discovered modules with their declared identities, dependency graph and resolved zones.

### FLOW-107
Status MUST list every currently quarantined spec with its module, ID and quarantine reason.

### FLOW-108
Status MUST list every module with boundaries_opt_out and its authored reason.

### FLOW-109
Status MUST report each module's count of applied archived changes whose implementation Provenance date is later than its latest drill report date.
With no drill report, it MUST count all that module's applied archived changes, including zero when there are none.

### FLOW-110
Status MUST count applied changes archived on HEAD whose change file is absent from the locally known configured remote/base ref, falling back to the local base when the remote ref is absent.
The base and remote MUST use the existing configured/default base-ref resolution.

### FLOW-111
Status MUST warn with P08 when the three or more most recently applied changes consecutively target one module, identifying the module and consecutive count.

### FLOW-112
Status MUST report the number of phase-changing commits reachable from HEAD for each applied archived change, matching Commits for that change.

### FLOW-113
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

This is the second of three approved applications. Count3 before this change,
4 after; return immediately to app status/NEW proof, then the mandatory fifth
Sol/high quality/goal review and appropriate checks. Push only on PASS and stop;
stop earlier on any native block. B01–B03 Explain remains a known integration
gap outside these three applications; do not hide it or claim review PASS.
Fresh fixed Sol/high eval/impl workers and isolated restricted assertion audit;
normal max2 audit rounds. Workers never commit/advance phases/write claims.
Only described interface/evals/new stubs during evals, only internal during impl.
Use representative existing fixture/public helpers and keep own Verify under5m,
full under15m; no process instrumentation or increased timeouts. Read conventions
remain review plus the single app PATH-shim eval, extended only by the caller.
No bootstrap source/tests, rebuild-v1, deleted/saved/abandoned implementation,
old worker logs, protected edits, other-module implementation or drill gap fixes.

## Provenance

- evals 2026-10-30T15:02:18.346522Z agent=manual fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=457 failing=7 passing=450
- audit 2026-10-30T15:05:03.208732Z agent=manual fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0
- evals 2026-10-30T15:08:43.685351Z agent=manual fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=457 failing=7 passing=450
- audit 2026-10-30T15:09:35.428263Z agent=manual fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0

## Audit

No findings
Checked IDs: FLOW-113
