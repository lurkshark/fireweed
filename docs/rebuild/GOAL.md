# Stabilization and cleanup restart prompt

**Current selected milestone: A → B → C1–C8.** This update prepares documentation
only; execution has not started. The owner supplied the execution scope below.
Checks contracts/proof (CK1–CK4) are deferred; all full-completion gates remain.
**Budget: three days from execution start**, selected by the owner. Record the
start/deadline when execution begins. Stop for owner review at that deadline or
after C8, whichever comes first.

The unchanged [design](../../fireweed-design-v1.md) is product authority;
[SYSTEM.md](../../SYSTEM.md) governs process and stops.
[ACCEPTANCE.md](ACCEPTANCE.md#current-milestone--stabilization-and-cleanup) holds
all ordered tasks, proof requirements and the final report contract.
[PROGRESS.md](PROGRESS.md) holds the handoff and single controller pin.
[DECISIONS.md](DECISIONS.md) records design clarifications;
[LESSONS.md](LESSONS.md) is bounded historical input.

```text
/goal Execute the current stabilization and cleanup milestone in ACCEPTANCE, in order: A identical-controller restoration/validation; B protected process cleanup committed separately; C1–C8 normal product changes through one real gitstore regenerate change. Preserve every full-completion requirement and defer CK1–CK4. Stop at the first blocked change, unresolved verification failure under SYSTEM's correction limit, uncovered owner decision, milestone budget or C8 completion. Do not push.

Resume:
- Read AGENTS.md, SYSTEM.md, the relevant unchanged design clauses and docs/rebuild/{PROGRESS,ACCEPTANCE,DECISIONS,LESSONS}.md. Record the three-day budget's execution start/deadline in PROGRESS, or retain the existing deadline when resuming. Inspect live Git status, queue, runs and shared claims; preserve unrelated worktrees. Never switch controllers during a run/claim or resume another controller's claim.
- Execute Part A using the recorded source and exact pin. The earlier missing-file snapshot differs from planning-time observation: the matching pin is present. Do not infer restoration/validation success from its presence. Build from the specified clean source with the prescribed cache and matching -trimpath settings where discoverable; compare the full hash. A mismatch stops work with both hashes reported to the owner; no automatic repin or bootstrap fallback. Then run pinned status --json, check and doctor --no-agent, and report results.
- Main workflow commands use /Users/joel/Workspace/fireweed/.fireweed/bin/fireweed, never mutable bin/fireweed. Put the current tree's .fireweed/bin first on worker PATH. Use GOCACHE=/tmp/fireweed-rebuild-cache. Setup copies the invoking controller into scratch. Read live approval state from status --json and bounds from fireweed.yaml without recording their values in planning documents. For later history checks, retain baseline be35552b706aa3000ff8f11306d129ff1d701b85 and pass an explicit nonempty range.
- Workers, separate restricted auditors and regeneration workers use exactly: codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c 'model_reasoning_effort="high"' -s workspace-write {prompt}. Stop on access failure; do not substitute model/effort. Accept the supplied macOS sh diagnosis without re-investigating it.
- The owner authorized committing/pushing the initial B1–B3 planning update before A on 2026-11-02. After A, review those committed edits and finish any remaining B cleanup in a separate protected author commit before C, with no active change/run/claim. This authorizes no product-work push. Product specs and Decisions are not edited manually.
- For each C change, read the pinned controller's author-change playbook and projected specs, use new: handles, stamp, check and show --projected, then commit authoring separately with no active run/claim. Run context before each step; workers verify with fireweed verify and return. The owning controller alone advances phases, commits and archives. Workers never invoke lifecycle mutators or edit run/phase/Provenance. Preserve fail-first (or legitimate already_satisfied), audited eval_fix, frozen phases, separate restricted audits and normal full verification. Batch coherent changes by module, execute one at a time, finish each C stage before the next, and never reset retries/rounds.
- For C7, stop if pnpm or uv is absent; prove each recipe with doctor --no-agent in a fresh init scratch project. C8 is one real regenerate: true gitstore change through the normal lifecycle, not a drill. Never read bootstrap, deleted implementation or old history as regeneration inputs.
- After C8, stop. No module drills, checks milestone, controller promotion, delivery, publication or push. Report applied change slugs and commits, blocked/skipped work with reasons, the live status --json unapproved count and local go test ./... result. Keep evidence in committed archives/Provenance and native drill reports only; no temporary paths, extra binary hashes, manifests or case counts in planning documents.
```
