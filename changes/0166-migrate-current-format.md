---
audit_pending:
    - APP-154
audit_rounds: 1
bases:
    APP-014: 1ce99378
module: app
phase: audit
run: 20261104T121536-5791bf8c68d7e282cc14f914
---

## Motivation

ACCEPTANCE C6 selects the design's migrate command and explicitly permits a
verified no-op for format 1. Users need a checked confirmation that their
current project already uses the supported format, including JSON and causes
when verification fails. Implement no older-format rewrites in this slice.

## Interface

No public signatures change. Compose project loading and static Check through
the app CLI; validation runs no hooks or agents and performs no writes.

## Added

### APP-153
migrate MUST discover the nearest project root and validate a format-1 project with project.Check before reporting already current with exit 0. The format-1 operation MUST leave working files, index, HEAD and refs unchanged.

### APP-154
migrate --json MUST emit schema 1 with string data equal to the plain confirmation on success or null data on failure. It MUST retain diagnostic codes and messages, with empty stderr.

### APP-155
migrate MUST return exit 2 for usage, project discovery or loading failures and unsupported format versions, using U01 for usage, C01 for unsupported formats and the existing loading cause code otherwise. Static validation errors MUST return exit 1 with their existing cause codes. All failures MUST include corrective guidance.

## Modified

### APP-014
gc and migrate listed by help MUST dispatch supported maintenance operations instead of reporting that they are unimplemented or under construction.

## Removed

## Decisions

## Notes

Accept no positional arguments or command-specific flags other than --json.
Require explicit configuration format version 1; unsupported versions are
diagnosed, never rewritten or guessed. Reuse project.Check and preserve its
diagnostics and warnings. This is only the owner-selected current-format no-op;
the full compatibility/migration obligations remain open for later work.

The last unfinished-command representative is migrate. APP-014 intentionally
changes that temporary requirement to supported maintenance dispatch as C6
implements the last placeholder. Update its expected result through this normal
Modified spec. Retain help-listing, plain/JSON status/output and diagnostics
assertions under supported and invalid requests; keep U01 rejection assertions
for invalid maintenance usage. Do not weaken any other existing assertion.

Use public CLI evals in temporary repositories for nested-root discovery,
valid/empty current-format projects, exact confirmation/JSON, repeated no-op
with dirty/staged/untracked files and all refs preserved, malformed config,
unsupported format, a static validation error and argument rejection before
mutation. Keep fixtures representative, own Verify below five minutes and full
impl Verify below twenty-five. No process instrumentation or time-bound changes.

Do not edit other modules, root config, protected setup, module SPEC/Decisions,
bounds or timeouts. C7 tooling/recipes and C8 regeneration follow only after C6.
Workers run pinned fireweed context, verify with fireweed verify --quick,
poll until terminal and return. They never commit, invoke lifecycle mutators in
the held checkout, edit run/phase/Provenance or alter frozen evals/interfaces.
Only the outside owning controller advances, commits, applies and archives
with normal full verification.

The separate restricted auditor edits only the existing Audit body. Findings
use - ID: "exact clause" — missing assertion. Passing audit writes No findings
and Checked IDs. Re-audits cover only flagged IDs, retaining prior clearances
and all counters/bounds.

Owner decision on 2026-11-04: narrow APP-154 to currently observable diagnostic
codes and messages. project.Check currently emits only errors with Warning
false, so its requested positive-warning fixture has no public input. Preserve
every existing eval assertion, including diagnostic warning-status comparisons;
full-completion warning obligations remain open. The original finding and stop
remain in committed history. The prior automatic correction returned without
eval edits, and the controller interrupted before another audit. Resume this
same run and re-audit only APP-154 against its narrowed requirement, retaining
the other IDs' prior clearances and all counters/bounds. No new warning-producing
project behavior is authorized by this decision.

## Provenance
- evals 2026-11-04T12:33:20.065941Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=323 failing=11 passing=312
- audit 2026-11-04T12:33:58.303303Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
- evals 2026-11-04T15:11:53.705668Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=323 failing=11 passing=312
## Audit

- APP-154: "It MUST retain diagnostic codes, messages and warning status, with empty stderr." — missing assertion that a warning returned by project.Check is retained with warning status true.

## Questions
