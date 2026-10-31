# Rebuild plan

[PROGRESS.md](PROGRESS.md) records the current controller and handoff.
[ACCEPTANCE.md](ACCEPTANCE.md) separates early-cutover evidence from full product
completion. The owner selected early Go core-loop cutover on 2026-10-30.

## Milestone 1 — Finish the cutover prerequisites

1. Let the existing primary-checkout run finish 0132 and 0133, or resolve a new
   checkpoint before proceeding. Respect its live claim and owning controller.
2. Once that run has returned and main is clean and unclaimed, integrate the
   author commits from codex/stub-handoff-guidance. Apply their authored files
   onto current main; do not restore the branch's older lifecycle snapshot.
3. Execute 0134–0138 with bootstrap and all ordinary phase checks:
   - 0134: implementation-ready stub guidance and Go example.
   - 0135 then 0136: retain JUnit failure text and supply bounded retry details.
   - 0137: worker Decisions/Notes guidance and examples.
   - 0138: isolated audit input workspaces, validated Audit-only import and
     cleanup, including partial work on failure/limit exits.
4. Return to the core-readiness gates. Add further changes only for a concrete
   failing core scenario or an unmet core design/spec clause discovered there.
   The current focused inspection identified audit workspace isolation as the
   additional implementation blocker; it did not establish the gates as passed.

All of 0134–0138 are prerequisites, including the guidance improvements. A
blocked change still stops the run for an owner checkpoint.

## Milestone 2 — Prove and promote the core loop

Pass EC1–EC5 in ACCEPTANCE: scoped core review, fresh-checkout Go and CLI checks,
one controller/NEW comparison smoke, and NEW governing a real change in an
isolated clone with real workers and a restricted auditor. Reuse permanent
fixtures for negative paths and bounded recovery/limit behavior.

Then perform EC6 in a separate author commit: pin NEW separately from the
mutable subject, align controller instructions/setup/PATH, record source/hash
and remaining limitations, and smoke-check status/check/doctor with
the promoted controller before subsequent core-loop changes. Keep bootstrap available for explicit recovery.
No active run or shared claim may cross the handoff.

## Milestone 3 — Deferred product completion

After early cutover, resume the full design and full-completion gates. Deferred
work includes real drills/regeneration validation; gc/migrate, show --stale and
check --fix; remaining recipes; version-range/file-format migration; C17–C19
warning completeness; packaging and final distribution. These remain required
for full completion and do not block the scoped core milestone. Real drills
still need their own owner authorization.

## Controller and subject

| Stage | Workflow controller | Product under test |
| --- | --- | --- |
| Current / EC1–EC5 | .fireweed/bin/fireweed-bootstrap, existing 0.1.1 pin | bin/fireweed |
| After EC6 | separately pinned NEW at the path recorded in PROGRESS | bin/fireweed |

Use explicit paths and GOCACHE=/tmp/fireweed-rebuild-cache. Worker commands
named fireweed and FIREWEED_BIN must resolve to the selected controller;
product fixtures explicitly use the subject binary. A worker build of the
subject must not replace the pinned controller. Never repin or fall back during
a live claim. After a cutover failure, stop and recover explicitly with the
owning controller before considering a controller change.

## Choosing and closing a change

- Motivation names the CLI scenario, user-visible need and design/SYSTEM basis.
- Target one module; batch coherent work and keep every spec atomic. Return
  from a dependency to its caller as soon as the interface is usable.
- Close each slice with an actual NEW invocation and inspectable result.
- Use eval_fix for wrong frozen evals, normal changes for behavior. Keep explain
  accurate for emitted codes; do not allocate speculative diagnostics.

## Per-change loop

1. Author: read SYSTEM, projected specs, relevant design clauses and the active
   controller's author-change playbook. Use new: handles, stamp, check and review
   show --projected, then commit only the authored change separately.
2. Controller start, then context before each worker step.
3. Evals: black-box tests, described interfaces and allowed new disposable stubs.
   Every Added/Modified/fixes ID needs a genuine failure unless already_satisfied.
   Iterate with verify --quick; done runs full checks.
4. Independent audit: only unasserted spec clauses, restricted inputs, re-audits
   limited to prior finding IDs, and at most two configured findings rounds.
5. Implementation: allowed files only. The controller's done verifies, applies,
   archives and commits. Workers never commit.
6. Update PROGRESS only when state, next work or standing rules change.

## Evidence and limits

Permanent evals use public APIs, temporary repositories, explicit fixture Git
identities, local hooks and bounded fake agents. Real-agent self-hosting is a
separate acceptance run. Neither requires a real drill for early cutover.
Record concise gate outcomes in PROGRESS and intentional comparison differences
in DECISIONS. Preserve existing evidence; rerun only what changed or remains
unproven. Do not create separate per-case reports or hash inventories.
