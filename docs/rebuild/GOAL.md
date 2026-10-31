# Early core-cutover goal prompt

Restart prompt for the implementing agent. Editing it does not run it or switch
controllers. [PROGRESS.md](PROGRESS.md) records the current pin and state;
[PLAN.md](PLAN.md) orders the milestones; [ACCEPTANCE.md](ACCEPTANCE.md) defines
both early-cutover and later full-product gates.

```text
/goal Complete the early Go core-loop cutover in /Users/joel/Workspace/fireweed: finish 0134–0138 and the preceding diagnostic queue, prove NEW can govern real spec changes with restricted audits, and promote a separately pinned NEW controller only after ACCEPTANCE EC1–EC5 pass. Complete EC6 as the controller handoff. Drills/regeneration and the deferred product surface remain later work.

Resume:
- Read AGENTS.md, SYSTEM.md and docs/rebuild/{PROGRESS,PLAN,ACCEPTANCE,DECISIONS}.md. The design remains product authority; SYSTEM's guardrails govern execution.
- Inspect live Git status, queue, controller processes and shared claims. Finish/resume existing work with the controller that owns it. Never reset another worktree's work or release its claim.
- Author commits on codex/stub-handoff-guidance must be integrated onto the current main only after its active queue run has returned and the checkout is clean and unclaimed. The author branch contains an older lifecycle snapshot; do not execute its stale pending phases as if they were main's current state.
- Until EC6 is recorded, use /Users/joel/Workspace/fireweed/.fireweed/bin/fireweed-bootstrap for workflow commands. After promotion use the NEW controller pin recorded in PROGRESS, never the mutable bin/fireweed subject. Use GOCACHE=/tmp/fireweed-rebuild-cache.

Work:
- Complete 0134–0137 as cutover prerequisites, followed by 0138's audit input-workspace isolation. Keep normal fail-first, independent audit, phase-zone checks, approval/P07 and retry/round bounds enabled.
- Follow PLAN's per-change loop and SYSTEM's scope guardrails. Only concrete core failures or unmet core clauses justify additional queued fixes; return to the gates after affected checks pass.
- Pass the clean-checkout Go checks, core controller/NEW checks, one reference lifecycle smoke and isolated real-worker self-hosting in ACCEPTANCE EC1–EC5. Real workers/auditors use codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c 'model_reasoning_effort="high"' -s workspace-write {prompt}. On access failure, stop and report it rather than substituting a model or effort.
- Never commit during a worker step. The active controller owns phases, Spec/Decision application and archives. Protected setup, planning and promotion edits belong in separate author commits with no active run/claim.
- Stop on any blocked change. Push only on explicit owner instruction. Never alter refs, bounds or requirements to evade enforcement.

Promotion and completion:
- EC6 pins a verified NEW build separately from bin/fireweed and updates AGENTS, SYSTEM, PROGRESS, setup/recipe instructions and worker PATH consistently. Keep bootstrap available for explicit recovery; never silently mix controllers or resume another controller's live claim.
- Keep the design scope intact. No bootstrap source/tests, rebuild-v1 or deleted implementation as an oracle; NEW has no runtime bootstrap dependency.
- No real drill/regeneration work is authorized by this milestone. Existing deterministic fixture tests may still run as part of the normal Go checks.
- Record concise gate results and the controller pin in PROGRESS; per-change evidence remains in archived Provenance. No new report inventories or hash manifests.
- This goal is complete only when EC1–EC6 are proven. Full product completion remains open until every later full-completion gate passes.
```
