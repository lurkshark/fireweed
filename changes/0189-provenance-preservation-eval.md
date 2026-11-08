---
bases: {}
eval_fix:
    - RUN-047
    - RUN-048
module: runner
phase: audit
run: 20261108T213829-a3f8bb244d3dde95398de65e
---

## Motivation

The owner selected the proposed proper-stage provenance preservation eval after
reviewing retained drill failures. RUN-047/RUN-048 constrain Audit copying itself;
their existing final-archive raw prefix assertion also constrains later successful
Done formatting. Preserve byte-exact copying assertions before successful audit
Done, then check the unchanged ordered nonblank provenance prefix after Done as
clarified in workflow. Keep RUN-047/RUN-048 unchanged.

## Interface

Audited eval_fix only. Edit the existing public runner audit eval and this change
during evals. Impl edits only this change. No spec/runtime/interface or other
behavioral change. Every runner eval and covered ID must continue to pass.

## Added

## Modified

## Removed

## Decisions

## Notes

Limit edits to Test_RUN_047_RUN_048_FirstAuditBeforeProvenanceIsImported and a
small local helper if necessary in packages/runner/evals/audit_workspace_test.go.
Retain its first Audit insertion before Provenance with all other original bytes
unchanged. Make the first fake auditor return exit 9 after that permitted edit.
The next fresh audit invocation captures its copied input before returning 0.
Before any successful audit Done, require that retry input's raw Provenance
section exactly equals the original raw section and contains the imported No
findings/Checked IDs content. Require four ordered agent steps (evals, failed
audit, successful audit, impl), the nonzero/zero audit results and applied state.
Use bounded fake agents and normal retries; do not mutate the held checkout from
the fake auditor. Keep the first-input/candidate insertion-only check.

In the final archive replace only the raw strings.HasPrefix provenance comparison
with an ordered-prefix comparison of nonblank lines. Blank means
strings.TrimSpace(line) == ""; preserve each remaining line byte for byte, with no
trim or normalization of its content. Keep exact original author-section checks,
Audit import, applied-state and workspace-cleanup assertions. Do not relax the
raw pre-Done preservation assertion or RUN-047/RUN-048 themselves.

Keep packages/runner/evals/retry_failure_details_test.go unchanged, preserving
module/classname/testcase/Case.Failure details from all reported modules, lexical
module then report ordering, passing/skipped exclusion, 16,384 Unicode code point
total, 4,096 excerpt bounds and truncation/omission markers. Retain every other
audit rejection, copy-on-failure/stop, isolation and cleanup eval. All own evals
must pass; no broader formatting change or newly mandated serializer spelling.

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
## Provenance
- evals 2026-11-08T21:49:33.629417Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=159 failing=0 passing=159
