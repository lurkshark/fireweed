# Workflow proof preparation restart prompt

The next planning target is workflow contracts and regeneration proof
(WF1–WF4). The completed checks milestone and passing proof remain valid.
[ACCEPTANCE](ACCEPTANCE.md#planned-milestone--workflow-contracts-and-regeneration-proof-wf1wf4)
defines the ordered plan and all full-completion gates. [PROGRESS](PROGRESS.md)
records current state, the controller pin and execution window.
[SYSTEM](../../SYSTEM.md) governs scope and stops;
[DECISIONS](DECISIONS.md) retains design clarifications. The unchanged
[design](../../fireweed-design-v1.md) remains product authority.

The current documentation update plans the work. WF1 contract review and WF2
clean candidate validation are pending. Their execution handoff stops before
the native workflow drill for explicit owner approval on the exact candidate.
Preserve the controller pin, all counters and configured bounds. Preparation
uses the existing deadline of 2026-11-07 22:51:05 UTC without extension.

Preparation restart prompt

```text
/goal Execute WF1–WF2 in ACCEPTANCE for the workflow regeneration proof. Review workflow contracts, Decisions, interfaces and public evals, then identify and validate a clean candidate. Reuse valid completed checks evidence. Stop at any correction needing owner agreement or other SYSTEM stop. After successful validation, commit the preparation evidence and handoff and stop before invoking a native workflow drill for my approval. Preserve the exact controller pin, configured bounds, counters, prior reports/proof refs and unrelated worktrees. Retain the existing deadline of 2026-11-07 22:51:05 UTC. Do not push.

Resume
- Read AGENTS.md, SYSTEM.md, relevant unchanged design clauses and docs/rebuild/{PROGRESS,ACCEPTANCE,DECISIONS}.md. Follow WF1 then WF2. Preserve completed CK1–CK4 evidence and the recorded deadline. Stop if the deadline has elapsed.
- Inspect live Git status, queue, runs and common claims. Read approval state from status --json and limits from fireweed.yaml. Preserve attempts, audit rounds, approval enforcement and all other bounds. Confirm the controller identity against PROGRESS before use.
- Use /Users/joel/Workspace/fireweed/.fireweed/bin/fireweed as the sole workflow controller. Keep it separate from the subject executable. Put the current tree's .fireweed/bin first on worker PATH and use GOCACHE=/tmp/fireweed-rebuild-cache. Scratch setup copies the invoking controller; isolated candidate CLI proofs copy their explicitly selected subject. A pin mismatch or active claim from another controller stops execution.
- Workers and separate restricted auditors use exactly codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c 'model_reasoning_effort="high"' -s workspace-write {prompt}. Stop on access failure. Preserve the accepted macOS sh fixture diagnosis.
- WF1 reviews claims and approved state, projection/zones, lifecycle and audit transitions, error causes, full Done verification, apply/archive/provenance, blocked recovery/release, context, eval_fix and regenerate behavior. Include FLOW-055 and FLOW-062. Record the focused conclusion and stop for owner agreement before authoring product/eval corrections. Approved corrections follow the normal audited lifecycle with context before every step, frozen evals and assertions, fail-first or legitimate already_satisfied, required full verification and controller-owned commits/archive. Workers never advance phases or edit run/phase/Provenance.
- WF2 uses a fresh clean checkout and an explicit, separately built candidate. Run applicable go test ./..., go test -race ./... and go vet ./..., candidate check, doctor --no-agent, check-commits and CI with hook doctor enabled. Preserve baseline be35552b706aa3000ff8f11306d129ff1d701b85 and resolve a nonempty exact range. Exercise affected candidate CLI paths. Reuse valid prefixes and unaffected evidence without skipping required phase checks. Append exact workflow preparation evidence to the existing reviewed checks report, preserving its native result. Keep temporary paths/hash manifests out of planning documents and carry the preparation evidence into the later workflow native report.
- Verify clean main/candidate checkouts, scratch cleanup, no live run/shared claim, unchanged pin/bounds and preserved previous proofs/unrelated worktrees. Commit the review and candidate validation handoff, report existing warnings and live unapproved count, and stop for explicit drill approval on the identified candidate.
- WF3/WF4 remain conditional on that approval. The later drill uses an isolated workflow input with declared dependency interfaces, a separate full verifier and the rehearsed live isolation gate. The inspector starts alongside the controller. Release only after inspecting live input/verifier isolation. Before a second configured attempt, review the full first trace/native feedback and recheck live isolation. Monitor the same invocation to its terminal result. Review all worker commands/patches and enforce workflow plus runner/app evals and coverage. Retain reviewed native report/ref and clean scratch state. Further invocations or gap corrections need separate authorization.
```
