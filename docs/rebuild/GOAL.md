# Goal prompt

Revised 2026-10-26 for the restarted rebuild (D001). Earlier versions are in git history and on the `rebuild-v1` branch. This prompt covers one bounded session; update its scope section before each new session.

Paste the following into `/goal` in the `fireweed` repository:

```text
/goal Author the first slice of the restarted fireweed rebuild in /Users/joel/Workspace/fireweed, then stop for human approval. Stop when that is done or when a stop condition below is hit.

Before starting, read these, in this order: AGENTS.md; SYSTEM.md, especially "Scope guardrails"; docs/rebuild/PROGRESS.md; docs/rebuild/LESSONS.md; the "Restart" section of docs/rebuild/PLAN.md; and the design sections each change implements in fireweed-design-v1.md. For workflow commands, use the controller at /Users/joel/Workspace/fireweed/.fireweed/bin/fireweed-bootstrap with GOCACHE=/tmp/fireweed-rebuild-cache. The rebuild-v1 branch is a reference for the edge cases LESSONS.md names; never copy its specs, Decisions or evals.

Scope for this session:
1. Author up to five changes covering PLAN.md "Restart" items 1 and 2 (contracts, then gitstore), following `fireweed playbook author-change`. Each change targets one module and has at most five specs. Each spec states one observable behavior in at most three MUST/SHOULD clauses and about 80 words. Each Motivation cites the design section it implements. Stamp each change, run controller check, and commit each as an author commit.
2. Do not start them. Stop and report so a human can review and approve the specs before any evals are written.

Rules for this session:
- Specs describe design behavior only. Locality, read-only and no-network properties are SYSTEM.md conventions, not specs. Platforms are macOS and Linux with ordinary full SHA-1 clones.
- Interfaces stay small: public functions in api.go and shared types in a retained model/ subpackage (LESSONS.md).
- No cgo, native code, embedded binaries or process-level instrumentation, in specs or plans.
- Don't author changes beyond items 1 and 2.

Stop and report if anything needs a product decision that the design, SYSTEM.md and LESSONS.md don't settle.

Final report, 20 lines or fewer: the changes authored, with spec counts and commit SHAs; the controller check result; and any open questions.
```

Later sessions run approved changes through the full lifecycle under the SYSTEM.md rules (two audit rounds at most, stop after five applied changes) and author the next slice items.
