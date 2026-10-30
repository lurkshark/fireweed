---
bases: {}
module: app
---



## Motivation

The design Reading surface requires status to show modules, waiting/blocked
queue state, quarantines, opt-outs, drift, approval/audit counts and measured
cost/budget information. Workflow0120 now provides that complete read model,
but NEW still reports status unimplemented. Return immediately to this caller
after the approved fifth-change gate: review/build/vet/race/static/history and
actual Explain proof PASS, batch pushed to origin7c5e90e. Close the status slice
with an actual NEW disposable-fixture proof, then stop with this caller local.

## Interface

No public API/model changes. Compose workflow.Status through internal CLI data
types, existing moduleResults conversion and schema1/diagnostic helpers. Do not
recompute eligibility, costs, approval, drift or history in app.

## Added

### new:status-command
status MUST discover the nearest project root and expose workflow.Status for it.
Successful status MUST return exit0, including empty queues and warning-only reports.

### new:status-plain-queue
Plain status MUST identify every queued entry in returned order by slug, path and module, with its phase, blocked reason, audit rounds, waiting change and blocking rule.
It MUST report an empty queue explicitly when there are no entries.

### new:status-plain-overview
Plain status MUST identify discovered module names, prefixes, directories, dependencies and resolved zones.
It MUST report quarantines with module/ID/reason, boundary opt-outs with module/reason, each module's drift, the unapproved count and each archived change's phase-commit count from the returned report.

### new:status-plain-costs
Plain status MUST show every returned module cost with its JUnit seconds, testcase counts per spec, eval lines and disposable lines.
Modules without a recorded cost MUST not receive a fabricated measurement.

### new:status-json-shape
Successful status --json MUST emit schema1 object data containing modules, queue, quarantines, boundary_opt_outs, drift_since_drill, unapproved, phase_commits and metrics.
Empty array/map containers MUST appear as arrays/maps rather than null.
It MUST leave stderr empty.

### new:status-json-rows
JSON module rows MUST use the existing name, id_prefix, dir, depends_on and zones shape.
Queue rows MUST use slug, path, module, phase, blocked, audit_rounds, waiting_on and waiting_rule.
Quarantine and opt-out rows MUST use module/id/reason and module/reason fields respectively.

### new:status-json-costs
JSON metrics MUST use module, seconds, cases_per_spec, eval_lines and disposable_lines fields.
The cases_per_spec, drift_since_drill and phase_commits maps MUST retain their returned identifier/count associations.

### new:status-diagnostics
status MUST retain the returned diagnostic codes and messages, including budget and P08 warnings.
JSON diagnostics MUST retain path when present and warning status; plain output MUST show each diagnostic code and message.

### new:status-failures
Usage and project configuration errors MUST return exit2 with corrective guidance, using U01 for usage and the existing configuration cause code.
Status failures MUST return exit1 with corrective guidance, retaining assigned cause codes and using P01 for an unavailable base ref.
JSON failures MUST use null data with code/message diagnostics and empty stderr.

## Modified



## Removed



## Decisions



## Rejected alternatives

No second status model, library completion, new flags, diagnostic allocation,
independent convention observer, hook/agent execution or broad error hardening.

## Notes

Owner approved revised status slice. The mandatory fifth quality/goal review and
appropriate checks passed; origin/main7c5e90e was verified. Review count0 before
this application,1 afterward. Independent audit/full native verification and
actual NEW proof, then stop with this caller local; no further push or drill.

Fresh fixed Sol/high workers, isolated restricted assertion auditor, normal
max2 rounds/re-audit only previous finding IDs. Evals use public app/workflow/
project interfaces and temporary local fixtures. Preserve all existing assertions.
No API/model/protected/other-module changes; no commits, phase/claim/ref writes
by workers. Own Verify under5m/full under15m; keep representative cases compact.
No bootstrap source/tests, rebuild-v1, canceled/deleted/saved implementation or
old worker logs. The maintenance fixture phase is impl; derive fixture state
from its public data, never force an incorrect evals phase.

Extend only the existing APP-001 PATH-recorder convention test. Preserve its
zero-external-invocation assertion for existing commands, then exercise status
plain/JSON in that same test with a blind Git recorder/pass-through and the
existing hook/agent sentinels. Read-only Git is permitted. No Git proxy that
classifies argv, calibration cases, syscall observer or second convention test.
