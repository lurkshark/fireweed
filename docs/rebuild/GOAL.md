# Checks milestone restart prompt

**Planning only.** Execution requires a separate owner instruction. No change,
worker, verification run or drill is started by maintaining these documents.

The unchanged [design](../../fireweed-design-v1.md) is product authority;
[SYSTEM.md](../../SYSTEM.md) governs process and owner checkpoints.
[ACCEPTANCE.md](ACCEPTANCE.md) holds the ordered milestone gates and full-product
obligations. [PROGRESS.md](PROGRESS.md) holds status, the controller pin and proof
locations. [DECISIONS.md](DECISIONS.md) records design clarifications;
[LESSONS.md](LESSONS.md) is bounded historical input.

```text
/goal Execute ACCEPTANCE CK1–CK4: one focused checks contract review and only necessary scoped corrections, candidate verification, then one independent checks drill and reviewed evidence handoff. Follow those gates in order under SYSTEM's process rules. Keep main's pin and configured bounds unchanged, preserve existing proofs, and stop at owner checkpoints. Full product completion remains open.

Resume:
- Read AGENTS.md, SYSTEM.md, the relevant unchanged design clauses and docs/rebuild/{PROGRESS,ACCEPTANCE,DECISIONS,LESSONS}.md. Inspect live Git status, queue, controller processes and shared claims; preserve unrelated worktrees. Finish a claim only with its owning controller. This prompt needs separate execution authorization.
- Verify main's controller source/hash against PROGRESS. Use /Users/joel/Workspace/fireweed/.fireweed/bin/fireweed for main workflow commands, never mutable bin/fireweed or an automatic bootstrap fallback. Put the current tree's .fireweed/bin first on worker PATH and use GOCACHE=/tmp/fireweed-rebuild-cache. Setup copies the actual invoking controller into scratch. Preserve the standing history baseline and valid prior evidence.
- Workers, separate restricted auditors and drill workers use exactly: codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c 'model_reasoning_effort="high"' -s workspace-write {prompt}. Stop on access failure; never substitute model or effort. Read live approval state from status --json and bounds from fireweed.yaml without recording their values in planning documents.
- For an authorized change, read projected specs and the pinned controller's author-change playbook first; use new: handles, stamp, check and show --projected, then commit authoring separately with no active run/claim. Run context before each step and verify work with fireweed verify. The owning controller alone advances phases, commits and archives; workers never edit run/phase/Provenance or invoke lifecycle mutators. SYSTEM and CK1 govern fail-first, restricted audits, frozen phases and eval_fix; historical recovery exceptions confer no new authorization.
- Complete CK1 → CK2 → CK3 → CK4. No automatic drill-gap fixes, counter resets, fresh drill after exhausted attempts or expansion outside checks. Stop on any blocked change, independence/access failure, ambiguous behavior or gap/contradicted Decision needing owner agreement. Later drills, main regenerated-code import, real regenerate changes, controller promotion, deferred product work and push require separate owner selection.
```
