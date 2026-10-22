# Goal prompt

Prepared 2026-10-21. This file describes future work; preparing it has not started the rebuild.

Paste the following into `/goal` in the `fireweed` repository:

```text
/goal Recreate fireweed from scratch in /Users/joel/Workspace/fireweed as a complete, working, self-hostable Go CLI implementing /Users/joel/Workspace/fireweed/fireweed-design-v1.md. Use a compiled, pinned binary built from /Users/joel/Workspace/fireweed-bootstrap to govern the entire rebuild. Read docs/rebuild/PLAN.md and docs/rebuild/ACCEPTANCE.md first and follow their workflow and completion gates.

Treat the design as the product authority and the bootstrap source, tests, README, and executable as behavioral references. Write new implementation and evals; do not copy the bootstrap codebase, import it as a dependency, invoke it to implement the new tool's commands, or preserve its internal architecture by default. Resolve discrepancies with design-derived tests and a recorded decision; do not silently reproduce reference bugs or omit designed behavior.

Begin by inspecting current state, pinning/building the bootstrap, initializing the Go recipe, completing SYSTEM.md, establishing module boundaries and protected hooks, and committing the setup. Then author small single-module changes and execute each through the bootstrap's start, context, verify, and done commands, including real fail-first evals and audits. Keep the bootstrap executable separate from the new executable under test. Only fireweed applies module specs and archives changes. Do not weaken hooks, zones, coverage, audit, or history checks to make progress.

This goal authorizes local setup, authoring, planning updates, and Git commits necessary for this rebuild. Make protected setup/infrastructure/documentation edits in separate author commits between active changes. Keep worker edits within the current phase and module. Prefer the manual lifecycle in this chat. Manual start may warn about changes absent from the configured base; record that local execution uses this goal's authorization and do not claim a human merge occurred. If using unattended run/next for this rebuild, --allow-unmerged is explicitly permitted for the locally authored planned changes only. Preserve the product's normal approval enforcement and test it. Do not modify fireweed-bootstrap, push, publish, tag a release, or change remote services.

Continue through the full design, not just an MVP. Maintain the requirement-to-spec/eval evidence matrix and a restartable progress record at safe author boundaries. Completion requires all acceptance gates: permanent ID-linked black-box evals, archived bootstrap-managed changes with valid history, passing Go tests/race/vet, bootstrap CI over the explicit rebuild commit range, fresh-project lifecycle/doctor/recovery tests against the NEW binary, successful regeneration evidence, and a documented self-hosting rehearsal. Report commands, commits, artifacts, and any remaining uncertainty honestly. Do not mark the goal complete while required behavior or verification remains unfinished. If progress requires unavailable access or a product decision that cannot be inferred, preserve recoverable state, document the exact blocker and needed input, and continue any independent work that remains possible.
```

The prompt names an outcome, evidence, boundaries, and continuation behavior, following [official OpenAI guidance for Goals](https://developers.openai.com/cookbook/examples/codex/using_goals_in_codex). No token budget is specified.
