---
already_satisfied:
    - FLOW-051
    - new:complete-module-reports
bases:
    FLOW-051: eb653782
module: workflow
---

## Motivation

The owner selected the proposed provenance and dependent-report contracts after
reviewing retained drill failures. The design requires successful Done to append
provenance and impl verification to run selected dependents; retry feedback needs
the reports of completed selected modules even when mapped testcases fail.
Make those observable obligations explicit, allowing blank-line normalization
only during Done. Existing behavior is expected to satisfy them; public evals
must establish that before this already_satisfied clarification can advance.

## Interface

No public interface or runtime change. During evals edit only workflow public
evals and this change. Impl edits only this change; the controller applies the
contract and archives it. Preserve all existing assertions and covered IDs.

## Added

### new:complete-module-reports
During full impl Verify, mapped testcase failures in valid reports MUST NOT
stop evaluation of the remaining selected modules.
Verify MUST return every completed module's report in Result.Reports.
Cancellation, timeouts, hook execution errors, and invalid or missing reports
MAY stop further evaluation.

## Modified

### FLOW-051
Each successful Done MUST append one provenance line containing phase, date,
manual or configured agent, version, playbook hash and result counts.
It MUST preserve existing nonblank provenance lines unchanged and in order.
It MAY normalize blank lines within the Provenance section.

## Removed

## Decisions

## Notes

Keep FLOW-025/FLOW-026 and RUN-043/RUN-044 unchanged. This change clarifies
workflow's producer obligation; it does not broaden runner retry details or
require continuing after infrastructure/report failures.

Extend Test_FLOW_005_FLOW_051_FLOW_052_FLOW_053_ContextAndProvenanceAndRetainedReports
in packages/workflow/evals/lifecycle_test.go with pre-existing valid provenance
entries separated by blank lines before Start. After each successful Done assert
that previous nonblank lines are an unchanged ordered prefix and exactly one new
entry was appended. A blank line satisfies strings.TrimSpace(line) == ""; do not
trim or normalize text in nonblank lines. Keep phase/date/manual/configured-agent,
version, full playbook hash, report counts, retained output and context assertions.

Add a focused public lifecycle eval for new:complete-module-reports using existing
alpha -> beta -> gamma temporary fixtures and deterministic local hooks. Exercise
these four selections: own fail/direct requires alpha and failing beta reports;
own fail/all requires alpha plus failing beta and gamma reports; own fail/none
requires alpha only; own pass/all requires alpha plus failing beta and gamma.
For each required module assert mapped testcase identity and failure-text sentinel
where failing, with alpha passing in the last case. Assert unselected modules
were not invoked, Verify fails, and HEAD and held impl phase remain unchanged.
Use valid JUnit reports, including ordinary hook nonzero exits with mapped failed
cases. Do not manufacture hook-launch errors. Keep all existing quick-mode,
cancellation, timeout, hook/report-error and pre-hook rejection evals unchanged.
Do not replace or weaken any runner retry-failure detail assertions.

All workflow evals must pass against unchanged runtime before advance. If either
already_satisfied claim cannot be established, stop and report; do not fake a
failure, edit runtime, narrow the contract or drop an assertion to get through.

Workers use the pinned .fireweed/bin/fireweed with that directory first on PATH
and GOCACHE=/tmp/fireweed-rebuild-cache. Read context before each held step, run
pinned verify --quick, wait for its terminal result and return. Never commit,
invoke lifecycle mutators, edit phase/run/Provenance or protected planning/tooling.
The outside controller owns full phase verification, commits, application/archive
and claim release. Keep the separate restricted audit and configured bounds.
A passing auditor edits Audit only, says No findings and lists every changed ID.

This is one of exactly two owner-authorized changes resolving provenance-prefix
and dependent-report gaps. No runtime/interface/model, hook/configuration, policy,
tooling, unrelated eval or Decision changes. Use current retained public contracts
and evals; do not read bootstrap/rebuild-v1, deleted code, saved reconstruction or
private investigation contents as implementation inputs. Preserve prior reports,
archives/proof/attempt refs, controller/recoveries, bounds/counters, deadline
2026-11-10 22:51:05 UTC and unrelated worktrees. Stop after these changes; no drill,
repin, source import, later milestone or push is authorized.
