<!-- fireweed:start -->
This project uses fireweed. Specs and change files are the source of truth.
Run `fireweed playbook author-change` before authoring changes and
`fireweed context <change>` before each step. Use `fireweed verify`
to check your work. Never commit during a worker step; fireweed owns commits.
<!-- fireweed:end -->

## Rebuild controller

Use the controller explicitly pinned in docs/rebuild/PROGRESS.md. Until the early core-cutover gates in docs/rebuild/ACCEPTANCE.md pass and promotion is recorded, this remains `/Users/joel/Workspace/fireweed/.fireweed/bin/fireweed-bootstrap`. After promotion, use the separately pinned NEW controller at the recorded path. `bin/fireweed` remains the mutable product under test; do not use it as the workflow controller. Never switch controllers while a run or claim is active. Read SYSTEM.md (including its scope guardrails) and docs/rebuild/PROGRESS.md before work; docs/rebuild/LESSONS.md summarizes historical findings; do not consult `rebuild-v1` or bootstrap source/tests as an implementation oracle. Protected author edits belong in separate commits between active changes. Use `GOCACHE=/tmp/fireweed-rebuild-cache` in this sandbox. Never read the bootstrap or deleted code during a regeneration drill.
