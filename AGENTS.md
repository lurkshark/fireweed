<!-- fireweed:start -->
This project uses fireweed. Specs and change files are the source of truth.
Run `fireweed playbook author-change` before authoring changes and
`fireweed context <change>` before each step.
<!-- fireweed:end -->

## Rebuild controller

Use the separately pinned NEW controller `/Users/joel/Workspace/fireweed/.fireweed/bin/fireweed`; docs/rebuild/PROGRESS.md records its verified source and hash. Put `/Users/joel/Workspace/fireweed/.fireweed/bin` first on worker PATH and use `GOCACHE=/tmp/fireweed-rebuild-cache`. `bin/fireweed` remains the mutable product under test; do not use it as the workflow controller. Scratch setup copies the invoking controller and workers use that scratch tree's `.fireweed/bin` first on PATH. Keep `.fireweed/bin/fireweed-bootstrap` for explicit owner-directed recovery only, with no automatic fallback. Never switch controllers while a run or claim is active or resume another controller's claim.

Use `fireweed verify` to check your work. Never commit during a worker step; fireweed owns commits. Read SYSTEM.md (including its scope guardrails) and docs/rebuild/PROGRESS.md before work; docs/rebuild/DECISIONS.md retains design clarifications. Do not consult `rebuild-v1` or bootstrap source/tests as an implementation oracle. Protected author edits belong in separate commits between active changes. Never read the bootstrap or deleted code during a regeneration drill.
