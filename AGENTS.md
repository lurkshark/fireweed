<!-- fireweed:start -->
This project uses fireweed. Specs and change files are the source of truth.
Run `fireweed playbook author-change` before authoring changes and
`fireweed context <change>` before each step.
<!-- fireweed:end -->

## Rebuild controller

Drive the workflow with the controller `.fireweed/bin/fireweed` (see SYSTEM.md), not `bin/fireweed`, which is the product under test. Put `.fireweed/bin` first on worker PATH and use `GOCACHE=/tmp/fireweed-rebuild-cache`. Never switch controllers while a run or claim is active.

Use `fireweed verify` to check your work. Never commit during a worker step; fireweed owns commits. Read SYSTEM.md (including its scope guardrails) before work; docs/rebuild/DECISIONS.md retains design clarifications. Do not consult `rebuild-v1` or deleted code as an implementation oracle.

Only sessions that coordinate the rebuild (owner sessions and the agent driving the controller) read docs/rebuild/PROGRESS.md and ACCEPTANCE.md. Workers in a lifecycle step and drill reconstruction workers must not read them: they record earlier runs' results, including how earlier reconstructions failed, which would bias the work being checked. Use `fireweed context` and the drill prompt instead.
