# Rebuild goal prompt

Restart prompt for the implementing agent. Editing it does not run it.
[PROGRESS.md](PROGRESS.md) holds current state and the next slice,
[PLAN.md](PLAN.md) the working method, [ACCEPTANCE.md](ACCEPTANCE.md) the final gates.

```text
/goal Build the independent fireweed CLI in /Users/joel/Workspace/fireweed to the complete fireweed-design-v1.md design, in small vertical slices that each end in observable behavior through the NEW CLI.

Resume:
- Read AGENTS.md, SYSTEM.md and docs/rebuild/{PROGRESS,PLAN,ACCEPTANCE,DECISIONS}.md; consult LESSONS.md for historical pitfalls. The design is product authority; SYSTEM's guardrails govern execution.
- Inspect git status, pending changes and claims. Resume a live phase, or the next slice the owner agreed in PROGRESS. Never reset work or release another worktree's claim.
- Use /Users/joel/Workspace/fireweed/.fireweed/bin/fireweed-bootstrap for all workflow commands. bin/fireweed is NEW, the product under test. Use explicit binary paths and GOCACHE=/tmp/fireweed-rebuild-cache. Workers, auditors and regeneration use the Codex command in PROGRESS; on access failure, record it and ask the owner.

Work:
- Before authoring, name in the change Motivation the user-visible slice, the design behavior it enables and how NEW will show it. Build only the dependencies the slice needs and return to the caller as soon as they work. Library evals alone do not close a slice.
- Follow PLAN's per-change loop and every SYSTEM guardrail. Keep audit enabled. Workers never commit; the controller owns phases, Spec/Decision application and archives. Author edits to protected files go in separate commits between active changes.
- Stop on any blocked change. Push only on explicit owner instruction. Never alter refs or bounds to evade enforcement.

Independence and completion:
- Never use bootstrap source/tests, rebuild-v1, deleted code, old binaries or saved attempts as an oracle. NEW never imports, invokes or embeds bootstrap.
- Regeneration receives retained specs, Decisions, interfaces, evals and permitted dependency interfaces only. Run only drills the owner authorized; gaps authorize no automatic fixes.
- Keep PROGRESS a short handoff. Per-change evidence belongs in the archived change's Provenance, not in new report files.
- The goal is complete only when every ACCEPTANCE gate passes.
```
