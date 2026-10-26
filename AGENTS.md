<!-- fireweed:start -->
This project uses fireweed. Specs and change files are the source of truth.
Run `fireweed playbook author-change` before authoring changes and
`fireweed context <change>` before each step. Use `fireweed verify`
to check your work. Never commit during a worker step; fireweed owns commits.
<!-- fireweed:end -->

## Rebuild controller

Use `/Users/joel/Workspace/fireweed/.fireweed/bin/fireweed-bootstrap` for all workflow commands in this checkout. The new `bin/fireweed` is the product under test, never the controller. Read SYSTEM.md (including its scope guardrails) and docs/rebuild/PROGRESS.md before work; docs/rebuild/LESSONS.md summarizes historical findings; do not consult `rebuild-v1` or bootstrap source/tests as an implementation oracle. Protected author edits belong in separate commits between active changes. Use `GOCACHE=/tmp/fireweed-rebuild-cache` in this sandbox. Never read the bootstrap or deleted code during a regeneration drill.
