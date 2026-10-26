# Goal prompt

Revised 2026-10-26 after the project review (D048). The original open-ended rebuild goal from 2026-10-21 is in git history at commit 1b27763. This prompt covers one bounded session; update its scope section before each new session.

Paste the following into `/goal` in the `fireweed` repository:

```text
/goal Apply the seven approved descoping changes in /Users/joel/Workspace/fireweed, then prepare the next vertical-slice change for human approval. Stop when that is done or when a stop condition below is hit.

Before starting, read these, in this order: AGENTS.md; SYSTEM.md, especially "Scope guardrails"; docs/rebuild/PROGRESS.md; and the "Priority after the 2026-10-26 review" section of docs/rebuild/PLAN.md. Read the change files themselves before working on each one. Don't read docs/rebuild/PROGRESS-LEDGER.md or old DECISIONS entries unless a change file points you to one. For workflow commands, use the controller at /Users/joel/Workspace/fireweed/.fireweed/bin/fireweed-bootstrap with GOCACHE=/tmp/fireweed-rebuild-cache.

Scope for this session:
1. Run changes/0050-descope-app.md through 0056-descope-contracts.md, in numeric order, through the controller's full lifecycle: start, context, evals, verify, done, audit, done, impl, verify, done. The project owner wrote and committed these changes; that commit is their approval, so the controller's warning that they aren't on the base branch is expected. Each change's Notes section lists which tests to delete and what to simplify. Do exactly that. Don't add new tests beyond what the Notes ask for.
2. After 0055 is applied, check how long a complete implementation Verify takes. If it's under 15 minutes, make a separate author commit that removes the four-worker eval sharding from .fireweed/hooks.go. Then run controller doctor and confirm all four canaries pass.
3. Write the next vertical-slice change for the workflow `done` step (apply, archive, provenance), following `fireweed playbook author-change` and the SYSTEM.md spec-shape rules: at most 5 specs, each with at most 3 MUST/SHOULD clauses and about 80 words, and a Motivation that cites the design section it implements. Stamp it, run check, and commit it as an author commit. Do not start it; a human approves it first.

Rules for this session:
- An audit's findings must be about spec clauses that no test asserts. If a finding is about how the tests are built, a platform, or a convention, the auditor discards it.
- A change gets at most two rounds of audit findings. If a third audit still has findings, write them into the change's Questions, add `blocked`, run done (which commits the block), and stop.
- In evals and implementation, do the minimum the spec requires. Don't use cgo, syscall interposers, git proxies, native code or embedded binaries, and don't add Windows-specific, SHA-256-specific or partial-clone handling.
- Never commit during a worker step; the controller makes those commits. Author edits to protected files go in their own commits, made only between changes.
- For evidence, add one line per applied change under "Applied since review" in PROGRESS.md, with the change, its assigned IDs, its phase commit SHAs, testcase pass/fail counts, and anything that didn't run. Don't write SHA-256 manifests or testcase inventory proofs.

Stop and report, without trying to work around the problem, if any of these happen:
- A change needs a spec or Interface edit beyond what its file says, or a frozen eval turns out to be wrong.
- Fixing something would mean weakening a hook, a zone, coverage, the audit, or the history checks.
- The same verification failure happens twice in a row.
- A single change takes more than about 2 hours, or needs more than 3 implementation iterations.
- A complete Verify takes longer than 30 minutes.
- Anything needs a product decision that the design and SYSTEM.md don't settle.

Final report, 20 lines or fewer: which changes were applied, with their final commit SHAs; the controller check and check-commits result for 14cec73..HEAD; Verify duration before and after; anything blocked, skipped or not run; and the name of the vertical-slice change waiting for approval. No push, no tags, no publishing.
```

Why it is shaped this way:

- The scope names already-approved work, so the human checkpoint after five applied changes doesn't stop it partway.
- New scope is authored but not started, so each new piece of work gets human approval first.
- Stop conditions end the session instead of inviting deeper investigation.
- It avoids open-ended phrases ("the full design", "do not weaken anything", "honestly") that the previous goal prompt's agent escalated on.
