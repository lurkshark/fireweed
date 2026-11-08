---
already_satisfied:
    - FLOW-007
    - FLOW-023
bases:
    FLOW-007: fcf4408d
    FLOW-023: 4849e489
module: workflow
phase: impl
run: 20261108T020656-6bd845e9af54ab3537d4deea
---

## Motivation

The owner approved making two existing design rules explicit in workflow's retained
contract: removed IDs stay reserved, and review-tagged specs need no fail-first
eval. The design's IDs, Tags and Evals-step clauses already require these rules.
The current candidate passes the original lifecycle cases; bounded public probes
also confirm correct allocation and review exclusion. This is a legitimate
already_satisfied clarification. Focused assertions should identify these causes
under FLOW-007/FLOW-023 before audit or impl integration tests can abort early.

## Interface

No runtime source, interface/model, Decision, hook, configuration, other-module
or worker-input expansion. During evals edit only workflow public evals and this
change. Impl needs no implementation edit; the controller alone applies this
contract clarification and archives the change. Preserve every existing assertion.

## Added

## Modified

### FLOW-007
Start MUST replace each local Added handle throughout the change with the next
available spec ID for its module. IDs of current and previously removed specs
MUST remain reserved, including specs Removed by the starting change.

### FLOW-023
During evals Verify for spec and fix changes MUST require genuine fail-first
evidence for every enforced change ID that is not already_satisfied.
Review-tagged specs MUST be excluded from fail-first requirements.

## Removed

## Decisions

## Notes

Use the existing temporary-repository fixtures and public APIs. Keep original
late-ID replacement, audit convergence, impl application, fail-first rejection,
coverage, quarantine and already_satisfied assertions intact. Do not weaken any
checker, require fake failing cases for review items or rename unrelated tests.
Keep FLOW-062 and its boundary cause-text test unchanged.

Give FLOW-007 focused public assertions for reserved current/previously removed
IDs, including removal in the starting change. Assert the returned and recorded
new ID and replacement throughout the change without running later lifecycle
steps that could hide allocation failures. Reuse representative existing coverage
where it already asserts a clause; avoid duplicate fixture systems or a broad
new matrix. Preserve abandonment's existing rule for assigned-but-unapplied IDs.

Give FLOW-023 a focused case that actually reaches evals Verify with review-tagged
change content and genuine fail-first only for the enforced non-review change ID.
Require success with no fake failing review case, while preserving the existing
negative assertions for missing fail-first on enforced IDs. Keep current coverage
and passing-result requirements distinct. Current implementation must remain
unchanged and all workflow evals must pass. Stop and report if already_satisfied
cannot be supported; do not manufacture a failure or author a runtime correction.

Use only current retained public contracts/evals and the unchanged design. Do not
read bootstrap/rebuild-v1 or saved reconstruction/private investigation contents,
or import reconstructed source. No regeneration worker input, hint, import list,
command interceptor or instrumentation is added by this normal change.

Workers run pinned context before each held step, pinned verify --quick, poll to
terminal and return. They never commit, invoke lifecycle mutators, edit runtime
fields/Provenance, change counters or protected planning/tooling. The outside
controller owns full phase verification, commits, application/archive and claim
release. Keep independent restricted audit enabled and its configured round bound;
a passing audit checks FLOW-007 and FLOW-023 and edits Audit only.

Keep the current controller pin unchanged through both approved changes. Preserve
prior candidates, reports/archives/proof/attempt refs, existing counters/bounds,
2026-11-10 22:51:05 UTC deadline, recovery copies and unrelated worktrees. Only
after both changes complete without any live run/claim, the outside owner will
validate clean source, update the pin and prepare a stopped retry handoff.
No worker launches a native drill, promotes a controller or pushes. The next
native workflow retry still requires separate owner approval.

## Audit

No findings; checked FLOW-007, FLOW-023
## Provenance
- evals 2026-11-08T02:13:49.686743Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=170 failing=0 passing=170
- audit 2026-11-08T02:14:27.385843Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
