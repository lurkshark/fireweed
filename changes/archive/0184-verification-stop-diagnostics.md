---
bases: {}
eval_fix:
    - RUN-037
module: runner
phase: applied
run: 20261108T022045-f396635dcaaf7eeeec9c6ac0
---

## Motivation

The owner approved better diagnostics for the existing verification-deadline
retention testcase. Its native failure stopped at a fatal read of a hook-created
marker before reporting agent work or the full held identity. Focused and full
runner replay passed, so no runtime retention defect is established. The design
requires retaining work on time stop; distinguish the available fixture evidence
without changing that requirement, the deadline or any assertion.

## Interface

Diagnostic-only normal independently audited eval_fix. During evals edit only
runner public evals and this change; implementation edits only this change.
No runtime source, public interface/model, spec, Decision, hook configuration,
worker command or other-module correction. Preserve every existing assertion.

## Added

## Modified

## Removed

## Decisions

## Notes

Narrow the edit to Test_RUN_037_VerificationDeadlineRetainsHeldWork and any small
helper confined to its diagnostics. Report the existing fixture agent invocation
log (including read errors), existence/read result for agent-added work, hook
marker/verification-entry evidence, clean time-stop result and held phase/run/
claim/assigned-ID/HEAD before an early fatal marker read can hide them. Capture
available evidence through ordinary fixture data and public workflow/runner APIs;
read failures are diagnostic data, not a reason to omit the rest of the evidence.

Ensure the original applied-result, marker-content, agent-work, no-commit,
phase/run/claim/assigned-ID/HEAD and identity-retention assertions remain and can
be evaluated even if the hook marker is missing. Restructure that named case's
fatal read into nonfatal assertions if needed; do not change shared-helper
behavior for unrelated cases. Missing marker, missing agent work or altered
identity must still fail the testcase. Do not turn a pre-verification stop into
an accepted pass or skip. Keep all RUN-037 coverage and every other test intact.

Preserve the exact two-second For budget, existing four-second hook sleep,
request timeout, fake worker/hook behavior and all configured limits. Add no
extra sleep, retry, timing matrix, load simulation, private implementation test,
filesystem/process instrumentation, command interception or runtime clock hook.
This change improves failure evidence only; it does not claim to solve the
unreproduced native failure or establish a new timing root cause. All runner evals
must pass with RUN-037 still covered; implementation needs no runtime correction.

Use current public retained contracts/evals. Do not read bootstrap/rebuild-v1 or
saved reconstruction/private investigation contents, import reconstructed code,
weaken a verifier, expand regeneration inputs or edit protected planning/tooling.

Workers use the current pin on PATH and GOCACHE from SYSTEM, run pinned context
before each held step and pinned verify --quick, poll to terminal and return.
They never commit, invoke lifecycle mutators, edit runtime/Provenance fields,
change counters or promote a controller. The outside controller owns independent
restricted audit, full phase verification, commits, archive and claim release.
Audit edits Audit only, checks RUN-037 and keeps the configured round bound.

Keep the current pin unchanged through both approved changes. Preserve candidates,
prior reports/archives/proof/attempt refs, counters/bounds, the existing extended
deadline, recoveries and unrelated worktrees. After both changes complete and
all run/claim processes exit, the outside owner will validate clean source,
update the pin and prepare a stopped retry handoff. No native drill or push is
authorized in this sequence; a retry requires separate owner approval.

## Audit

No findings; checked RUN-037
## Provenance
- evals 2026-11-08T02:34:04.260227Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=159 failing=0 passing=159
- audit 2026-11-08T02:34:30.501964Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
- impl 2026-11-08T02:40:56.606276Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=5f856d6bcab308b9ff1a8a7c03b9d2547744112aba81dd30dea8f61ec8062b03 tests=482 failing=0 passing=482
