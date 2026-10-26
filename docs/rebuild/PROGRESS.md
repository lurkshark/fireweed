# Rebuild progress

Current state for restarting work. Keep this file short: update the sections below in place and add one line per applied change. Per-change history through 2026-10-26 is in [PROGRESS-LEDGER.md](PROGRESS-LEDGER.md); read it only when tracing a past change.

## Status (2026-10-26)

- 51 changes archived. The 2026-10-26 review found the rebuild had stalled in hardening beyond the design: 11 of 30 product commands, no `done`, runner empty. See D048.
- In-flight `0049-held-approval-source` was abandoned and its ten phase commits plus authored commit reverted. FLOW-048 to FLOW-051 were never applied.
- The design was revised for bounded effort (new SHA in PLAN.md), SYSTEM.md gained scope guardrails, ACCEPTANCE.md gained R22 to R25 and narrower R13, R21 and gate 6, and `fireweed.yaml` now uses `verify_dependents: direct` (`ci` still runs everything).
- Queue: descoping changes 0050 to 0056 (app, workflow, checks, project, assets, gitstore, contracts), stamped; controller `check` passes.
- No real regeneration or NEW self-hosting has run. All R rows remain unverified.

## Controller pin

- Source: /Users/joel/Workspace/fireweed-bootstrap at 583b3c53740cf2341e0b0a79afee5b7f646ec1d4.
- Build: `GOCACHE=/tmp/fireweed-rebuild-cache go build -o /Users/joel/Workspace/fireweed/.fireweed/bin/fireweed-bootstrap .` from the bootstrap checkout. Version 0.4.0, schema 1.
- Binary SHA-256 at the original pin: f9c48ae6b14887ebc549110478c081bedc450847507c6967735ef8486c98a3ed. Rebuild and check if absent.
- The controller implements the original v1 design. The revised design's bounds (audit rounds, unapproved-change limit, budgets) are enforced by hand per SYSTEM.md.
- Setup baseline for history and CI ranges: 14cec73e1126ee37d33f1138d19936b6a05a59a6. Always use the explicit range `14cec73..HEAD`.
- Worker and auditor: `codex exec --ephemeral --model gpt-6-sol --sandbox workspace-write {prompt}`.

## Restart procedure

Inspect `git status`, the controller checksum and version, controller `status` and `list`, and `refs/fireweed/claims` before any edit. If a change is claimed, read its context and resume its phase without resetting work. Edit protected files only between active changes, in separate author commits. Follow the SYSTEM.md scope guardrails, including the two-round audit limit and the human checkpoint after five applied changes.

## Next

1. Apply 0050 to 0056 in order.
2. Remove the four-worker sharding from `.fireweed/hooks.go` once complete Verify is under 15 minutes; run doctor after.
3. Vertical slice per PLAN.md "Priority after the 2026-10-26 review".

## Open issues

- D042: `modules` must publish the held change's projected graph for boundary hooks. Plan: compose existing LoadApproval with Project/ApplyChange in app; claim disappearance mid-read is accepted.
- D046: processes that leave their process group can outlive cancellation. Accepted limitation after 0055.
- APP-023 timeouts seen once during 0048 Verify, cause unknown.

## Applied since review

(none yet)
