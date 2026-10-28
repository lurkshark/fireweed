# Rebuild goal prompt

Revised 2026-10-28. This is a restart prompt; editing it does not execute it.
[PROGRESS.md](PROGRESS.md) owns current state, authorizations and the next slice.
[PLAN.md](PLAN.md) owns sequencing; [ACCEPTANCE.md](ACCEPTANCE.md) owns final gates.

```text
/goal Build the independent fireweed CLI in /Users/joel/Workspace/fireweed to the complete fireweed-design-v1.md design. Deliver it through small vertical slices: each slice ends in observable behavior through the NEW CLI. Finish the approved checkpoint corrections first, then prove one queued change through NEW next and a bounded queue through NEW run.

Resume:
- Read AGENTS.md, SYSTEM.md and docs/rebuild/{PROGRESS,PLAN,ACCEPTANCE,DECISIONS}.md. Consult LESSONS.md for relevant historical pitfalls. The design is product authority; SYSTEM's guardrails govern execution. Planning/history adds no product requirements. Read the design initially, then the clauses needed by the current slice.
- Inspect Git status, pending changes and claims. Resume the live phase or the next slice in PROGRESS; honor checkpoints and preserve unrelated work. Never reset work, release another worktree's claim or restart passed milestones to obtain a clean state.
- Use /Users/joel/Workspace/fireweed/.fireweed/bin/fireweed-bootstrap for all rebuild workflow commands. bin/fireweed is NEW, the independent product under test. Use GOCACHE=/tmp/fireweed-rebuild-cache and explicit binary paths. FIREWEED_BIN names the invoking executable, including NEW in product fixtures.
- Reuse the recorded controller pin, setup and nonempty history baseline. If setup is absent, follow PLAN's one-time setup. Do not repin or rebuild the controller. Real workers, separate restricted auditors and regeneration agents use: codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c 'model_reasoning_effort="high"' -s workspace-write {prompt}. On access failure, record it and ask the user; do not switch models or retry indefinitely.

Work by observable slice:
- PLAN orders manual lifecycle (already proved), one-change next, bounded run, then recovery/limits/change kinds and the remaining surface. Complete only the dependencies needed for the current scenario and wire app promptly. Do not finish a whole module before its caller can exercise it.
- Before authoring each change, name the user-visible slice, the missing design/spec behavior it enables, and how NEW will demonstrate it. If a dependency takes a separate single-module change, return to its caller as soon as usable. Keep this statement in the existing Motivation/PROGRESS; add no new proof artifact or checklist.
- Finish fake-agent next/run before NEW unattended real-agent execution. Reuse the manual lifecycle and init/doctor evidence, repeating only portions invalidated by relevant changes or unresolved failures. Support code or passing library evals alone do not complete a slice.
- Preserve full design scope and all ACCEPTANCE gates. Missing features remain pending; a narrow slice does not waive their requirements. Do not restart completed spec splits, diagnostic migration, or setup. Deferred cleanup and drills stay deferred as recorded in PROGRESS.

Execution:
- This goal authorizes local authoring/execution and separate protected author commits between active changes. Workers never commit; the controller owns claims, phases, Spec/Decision application and archives. Preserve normal approval and P07 enforcement. --allow-unmerged bypasses neither bounds nor human checkpoints. Push only under explicit owner authorization recorded in PROGRESS; no publication, release tags or history rewrites.
- Run controller playbook author-change before authoring and context <change> before each worker step. Target one module; batch coherent related work without a numeric spec-count cap. Assign no IDs; stamp/check/review projection and commit author work separately. Use verify --quick to iterate; required full verify/done provide phase evidence. Frozen eval corrections use eval_fix.
- Follow all SYSTEM guardrails: atomic specs, brief behavioral Decisions, public-interface evals, supported environments, process groups, bounded eval cost, correction limit, restricted independent audits and at most two findings rounds. Keep audit enabled. Preserve assertions and design obligations. No speculative hardening, repository-wide rewrites or new acceptance ceremonies.
- Stop after five applied changes since review, on any blocked change, or before a third consecutive change to the same module/behavior unless explicitly authorized. Follow PROGRESS's standing checkpoint review policy. Record issues and leave the goal incomplete; never alter refs, counts or bounds to evade a checkpoint.

Independence and completion:
- Never use bootstrap source/tests, rebuild-v1, deleted implementation, old binaries or saved implementation attempts as an oracle. NEW never imports/invokes/embeds/delegates to bootstrap. Allowed operational references and the single lifecycle comparison are bounded by PLAN/SYSTEM.
- Regeneration receives retained specs/Decisions/interfaces/evals and permitted dependency interfaces only. Run only authorized drills, retain reports/refs and clean scratch worktrees. Gaps authorize no automatic fixes or reruns. Keep this checkout under the pinned controller; NEW self-hosting happens in isolation.
- PROGRESS stays a short current handoff: active slice, exit evidence, next work, blockers, authorizations and links. HISTORY holds one entry per applied change; reports hold detailed results. Record actual results once, never predicted successes.
- Complete all ACCEPTANCE gates, including fresh-source Go checks, controller history/CI on the recorded explicit range, NEW command fixtures, one reference smoke, real module regeneration, isolated self-hosting and local delivery. Account for warnings and limitations; milestones/checkpoints are not completion. Reuse valid evidence and stop verification when all gates pass unless later changes invalidate it.
```
