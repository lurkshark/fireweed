---
bases: {}
eval_fix:
    - RUN-027
    - RUN-028
module: runner
---

## Motivation

The owner approved the scoped residual-failure proposals. A bounded replay of the
retained failed reconstruction reported only the runner recovery test's commit-
subject assertion: its final subject said impl rather than containing apply.
Every reset, retained index/working/untracked content, P05 warning, lifecycle,
archive and claim assertion completed without another error. RUN-027/RUN-028
specify recovery and warnings, not those subject strings; FLOW-087 and the design
identify phase commits from change-file contents rather than message parsing.
Correct this retained-eval coupling without weakening the recovery assertions.

## Interface

No runtime source, public interface/model, spec, Decision, hook or configuration
edit. During evals only runner public evals and this change may change. Impl may
change only this change document. This is a normal independently audited eval_fix.

## Added

## Modified

## Removed

## Decisions

## Notes

In Test_RUN_027_RUN_028_RecoversAgentCommitsAcrossLifecycle replace the uncontracted
subject-word checks with public history evidence. Preserve the exact three
checked phase commits, and assert their change-document phase transitions through
audit, impl and applied using public workflow.Commits and/or Git reads of the
pending/archive document contents. Do not substitute another required spelling
such as impl for apply, or require a particular controller commit subject.

Preserve the old exclusion of fixture-agent commits using commit identities rather
than message wording. The bounded fake workers/auditor may record their own full
commit OIDs in the existing outside fixture log when making the successful
fixture commits. Require the expected fixture commits to be observed so absence
checks cannot pass vacuously, and ensure none is reachable from final held HEAD.
Use ordinary Git metadata and public inputs, not process/filesystem instrumentation
or private implementation tests. Keep fake invocation behavior and failure modes
unchanged outside this focused commit-success evidence recording.

Retain every existing assertion for HEAD reset to the held evals base before
verification, staged/unstaged/untracked work, all returned ordered steps, phase/
run/claim state, P05 warnings for evals and impl and propagation across advancement,
archive/provenance and applied spec/code contents, and final claim release.
Preserve the separate invocation-failure recovery coverage and every other test
and ID assertion. Reuse the established fixture; avoid duplicate helper systems,
new timing/platform matrices or unrelated cleanup. All runner evals must pass,
with both RUN-027/RUN-028 still covered. No runtime adjustment is needed or allowed.

Workers use the current pinned controller on PATH and GOCACHE from SYSTEM, run
context before each held step and pinned verify --quick, poll to terminal and
return. They never commit, invoke lifecycle mutators, edit runtime fields or
Provenance, change counters or protected planning/tooling. The outside controller
owns full phase verification, commits, apply/archive and claim release. Independent
restricted audit stays enabled. The auditor edits Audit only, quotes exact
normalized clauses for findings, checks RUN-027 and RUN-028, and preserves the
configured round bound; a passing submission is No findings with both IDs checked.
Do not read bootstrap/rebuild-v1 or saved reconstruction implementation, import
reconstructed source, or weaken any retained verification requirement.

Keep the current pin through both authorized changes and preserve the original
candidate object, reports/proof/attempt refs, existing archive counters, all
configured bounds, extended 2026-11-10 22:51:05 UTC deadline, recoveries and unrelated
worktrees. After both changes are archived with no active run/claim, the owner has
authorized validation of the clean corrected subject, verified controller update,
then one native workflow drill with the new controller. No worker launches a drill,
promotes a controller or pushes. Stop after that separately owned final drill.
## Audit
