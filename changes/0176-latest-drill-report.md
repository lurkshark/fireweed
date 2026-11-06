---
already_satisfied:
    - RUN-060
bases:
    RUN-060: 147a023a
module: runner
---

## Motivation

The design defines one drill report per module. Make the replacement behavior
explicit so the active report represents the latest reported invocation and
reviewed prior versions stay accessible through their Git commits. This keeps
old attempts and review notes from accumulating in the active report.

## Interface

The report path and format stay the same.

## Added

## Modified

### RUN-060

After reconstruction has been invoked, Drill MUST write the current invocation's report to .fireweed/drills/<module>.md, replacing any previous report there, with date, base_commit, agent, model and result frontmatter and Own evals, Dependent evals, Contradicted Decisions and Gaps sections. Eval entries MUST identify module, testcase, spec IDs and result. Agent observations and verification failures MUST remain visible, with each gap identifying a nearest spec ID or "no spec" and target module.

## Removed

## Decisions

## Notes

The existing writer already replaces the report. Extend the public runner eval
with a fake-agent fixture that starts with a committed report carrying a unique
prior marker. Verify the next report at the same path contains current evidence
and omits the marker, while the prior committed version is still readable in
Git. Retain the existing report-content and proof-ref assertions.

Reviewed report versions are retained by the author in Git. Keep prior proof
refs, the active milestone's evidence and the current invocation's worker
observations.

## Audit
