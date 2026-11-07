# Failed workflow retry checkpoint

[ACCEPTANCE](ACCEPTANCE.md#planned-milestone--workflow-contracts-and-regeneration-proof-wf1wf4)
is the ordered plan. [PROGRESS](PROGRESS.md) records status, the original deadline
and controller pin; [SYSTEM](../../SYSTEM.md) defines execution rules, and
[DECISIONS](DECISIONS.md) retains behavioral clarifications.

The separately approved safer retry ran exactly one native workflow invocation on
candidate `79cf33a7690082f28c06bf41a72719d282bc00fe`, with its two configured workers
and the protected [external wrapper/inspector](workflow-drill/README.md). Both raw
worker streams were durably preserved and forwarded unchanged. Full observed
commands/scripts and incremental semantic patches were reviewed. No local build,
product test, module setup, outside implementation/history read or retained-input
edit was observed. The actual live isolation gates passed before both releases;
only native verification feedback was added to the second prompt.

The invocation **failed** and exited with status 1. Workflow, runner and app each
reported PackageFailure/C11 and missing-coverage C10 diagnostics. Native own and
dependent case sections are empty. Source review also found `project.Project`
without a project import in the fresh `attempt.go`; this is a reconstruction
compile defect, not a retained contract/eval finding. No additional build or
correction was run. The [current native report and review](../../.fireweed/drills/workflow.md)
and failed reconstruction ref
`refs/fireweed/drills/workflow/20261107T172136.182088000` are retained.
Normal failed-report publication was observed; native reporting after rejection
or emergency interruption was not exercised and remains unestablished.

Stop here for owner review. This approval is consumed; it permits no further
native invocation or correction. WF4 and full product completion remain open.
The worker's Abandon/PRBody notes authorize no retained-input fix. Any proposed
change to the pinned controller, product, spec or evals requires a concrete scope
and separate owner agreement before authoring. A further drill needs its own
explicit approval; do not silently restart this invocation or add an attempt.

WF1/WF2 validation and approved follow-ups 0178/0179 remain valid evidence for the
unchanged candidate and range
`be35552b706aa3000ff8f11306d129ff1d701b85..79cf33a7690082f28c06bf41a72719d282bc00fe`.
The entire [prior workflow report](../../.fireweed/drills/workflow-before-retry-a79f34c.md)
is preserved byte-for-byte, including stopped-drill review limits and the safer
preparation/fake rehearsals. Preserve it, the valid checks proof, every prior
proof/attempt ref, archive counters, configured bounds, unrelated worktree and
controller pin. The original deadline is still **2026-11-07 22:51:05 UTC**.

The native input/verifier and owned caller clone are removed, all recorded owned
processes exited, the queue is empty and no live run/shared claim remains. No
reconstructed source was applied, controller promoted or push performed. Retain
this clean protected-evidence handoff and await the owner's next decision.
