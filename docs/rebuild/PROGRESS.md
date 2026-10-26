# Rebuild progress

Current state for restarting work. Keep this file short: update sections in place and add one line per applied change under "Applied". The first run is on the `rebuild-v1` branch; see [LESSONS.md](LESSONS.md).

## Status (2026-10-26)

- Restarted from the setup baseline (D001). No changes authored or applied. All R rows unverified.
- History baseline for check-commits and ci ranges: the restart docs commit, recorded below. Always use an explicit range from it.

## Controller pin

- Source: /Users/joel/Workspace/fireweed-bootstrap at 583b3c53740cf2341e0b0a79afee5b7f646ec1d4.
- Build: `GOCACHE=/tmp/fireweed-rebuild-cache go build -o /Users/joel/Workspace/fireweed/.fireweed/bin/fireweed-bootstrap .` from the bootstrap checkout. Version 0.4.0.
- The controller implements the original v1 design. The revised design's bounds are enforced by hand per SYSTEM.md.
- Worker and auditor: `codex exec --ephemeral --model gpt-6-sol --sandbox workspace-write {prompt}`.

## Restart procedure

Inspect `git status`, the controller version, controller `status` and `list`, and `refs/fireweed/claims` before any edit. If a change is claimed, read its context and resume its phase without resetting work. Edit protected files only between active changes, in separate author commits. Follow the SYSTEM.md scope guardrails.

## Next

Author and apply the first slice changes in PLAN.md "Restart" order, starting with contracts.

## Applied

(none yet)
