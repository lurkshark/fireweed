---
bases: {}
module: app
---

## Motivation

The design's CLI table and Regeneration section require `fireweed drill
contracts` to invoke the native rebuild/report/ref workflow. 0141 supplies the
runner API, but the app still calls drill unimplemented. Return to this caller
and expose the required command, documented exit behavior and schema-1 output,
so the DR1 slice can close with an actual clean-candidate CLI fixture before the
one isolated real contracts proof. No other deferred commands are included.

## Interface

## Added

### new:drill-routing
drill <module> MUST discover the nearest project root and invoke runner.Drill with that Root, the requested Module and the actual invoking executable as Bin. It MUST leave Timeout unset so configured backend bounds apply.

### new:drill-plain-success
A successful drill MUST return exit 0 with empty stderr and plain output identifying the module, passing result, project-relative native report path and retained rebuilt-code ref.

### new:drill-json-success
Successful drill --json MUST return exit 0 with schema 1 object data containing module, result, report, ref and base_commit string fields associated with the runner result. It MUST emit empty diagnostics and empty stderr.

### new:drill-preflight
drill MUST reject missing or extra module arguments with exit 2 and corrective U01, and an unknown module with exit 2 and corrective U02 naming it. These rejections MUST invoke neither setup/eval hooks nor agents.

### new:drill-config
drill project-discovery or configuration-loading failures MUST return exit 2 with corrective diagnostics retaining their cause code or C01 when uncoded. JSON MUST use schema 1 with null data and empty stderr; plain failures MUST use empty stdout and stderr diagnostics.

### new:drill-failure
Ordinary runner.Drill failures MUST return exit 1 with corrective diagnostics retaining the cause code or P01 when uncoded. Diagnostics MUST identify the native report and rebuilt ref when those artifacts exist. JSON MUST use schema 1 with null data and empty stderr; plain failures use empty stdout and stderr diagnostics.

### new:drill-usage-stop
A clean runner.IsUsageStop error with drill Result stopped MUST return exit 3 with the stop reason and any native report/ref identifiers. JSON MUST use schema 1 object data containing module, result, report, ref, base_commit and reason, with empty diagnostics and stderr; plain output MUST name the stopped module and reason.

## Modified

## Removed

## Decisions

## Notes

Use public app.Run/runner behavior with temporary repositories, fixture Git
identity and bounded fake agents/hooks only. Reuse existing CLI output helpers
and diagnostic conventions. Preserve the runner's report, retained ref and
scratch cleanup; cover a passing drill, failed own/dependent evals with retained
evidence, clean usage stop, setup failure without imaginary report artifacts,
argument/lookup/configuration rejection and invocation Bin through local fixture
observations. Avoid a Cartesian product of formats and failure modes.

The command is now implemented, so retarget APP-014's unimplemented-command
assertions from drill to a still-pending command such as gc, preserving those
assertions and APP-014/APP-015 coverage. Do not weaken generic failure envelopes
or other CLI behavior. New permanent tests remain independent of real Codex,
bootstrap and network access. An actual candidate executable fixture closes
DR2 later; no real module drill belongs in this change's evals.

Normal fail-first and separate restricted audit apply; only the unchanged main
NEW pin commits, applies and archives. Context is required before each worker
step. Passing Audit submissions use No findings and a complete Checked IDs list;
findings use an exact clause quote followed by a separator and explanation.
No public app facade changes, protected tooling edits, extra flags, other
deferred commands or automatic gap authoring are included.

