# Workflow proof approval handoff

[ACCEPTANCE](ACCEPTANCE.md#planned-milestone--workflow-contracts-and-regeneration-proof-wf1wf4)
is the ordered plan. [PROGRESS](PROGRESS.md) records status, deadline and the
controller pin. [SYSTEM](../../SYSTEM.md) defines execution rules and the exact
agent command. [DECISIONS](DECISIONS.md) retains behavioral clarifications.

WF1 review and WF2 validation are complete for candidate
`3d7d210fb7300a3cf48406247ce08db7a14e11fe`, with no correction needed.
The appended workflow preparation in the [checks report](../../.fireweed/drills/checks.md)
records the exact validation and history range. WF3 awaits separate explicit
owner approval for one native workflow invocation on that exact candidate.
No native workflow drill ran during preparation.

After that approval, use this prompt to resume:

```text
/goal Execute WF3–WF4 in docs/rebuild/ACCEPTANCE.md on validated candidate 3d7d210fb7300a3cf48406247ce08db7a14e11fe after my separate explicit approval for one native workflow invocation. WF1–WF2 are complete; reuse their committed preparation evidence. Follow AGENTS.md, SYSTEM.md, PROGRESS.md and DECISIONS.md. Preserve the pinned controller, recorded deadline, configured bounds/counters, prior proofs and unrelated worktrees. Use the isolated clone, inspector and live isolation gates required by WF3, reviewing actual retained inputs, full commands/scripts/patches and native feedback before each permitted release. Monitor the same invocation to its terminal result. A native failure or owner decision stops the milestone; further drills and gap corrections require separate agreement. Commit the reviewed native evidence and clean handoff, then stop at WF4. Do not import rebuilt code, promote a controller or push.
```
