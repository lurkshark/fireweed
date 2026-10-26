# Rebuild progress

Keep this file short: update sections in place and add one line per applied change.

## Status — 2026-10-26

- Stage-2 startup passed. Clean main at `0fee737d7bc2c21c6bf5e28faf299301e2542cb1`;
  no pending/archive changes or local claims. All R rows remain unverified.
- Pinned the tested controller and adapted protected setup; `check --json` passed
  with no diagnostics. Doctor passed all four canaries/removal recovery and both Sol/high probes
  with no diagnostics; scratch worktree removed.
- Stage-2 setup baseline: `be35552b706aa3000ff8f11306d129ff1d701b85`;
  use its explicit nonempty `baseline..HEAD` for history/CI acceptance.
- Historical restart baseline: `d991c26472ad990a5d050a07ce7dc873d1900433` (D001).
  Old build/check/doctor results belong to the retired setup, not the new pin.

## Controller candidate and agents

- Tested source executable: `/Users/joel/Workspace/fireweed-bootstrap/bin/fireweed`.
- Candidate: version `0.1.0`, embedded source revision
  `7f925190878e38fc38f30d16bb1cf47aca098cf7`, `vcs.modified=false`; metadata checked
  against HANDOFF.md at startup and installed for stage 2. SHA-256 of source
  candidate and destination: `f1516b5686435c23b54b71c5ac554aba6859b5454f072e0c5eb031dfd6f0fd4f`.
- Destination: `/Users/joel/Workspace/fireweed/.fireweed/bin/fireweed-bootstrap`.
  Copy the tested executable; do not rebuild or consult bootstrap source.
- Codex-only workers/auditors/regeneration: `gpt-6-sol`, high reasoning effort,
  exact argv in GOAL.md. A read-only Sol/low-effort probe exited 0 with the expected
  reply. High effort and both full doctor probes remain unrun startup checks.
- The new controller covers the revised design's bounds, eval_fix and usability;
  retain SYSTEM.md's human checkpoints and scope limits as well.

## Next / resume

Build the NEW binary's standalone command checkpoint, then the narrow manual
lifecycle slice before expanding the surface.

If a change is already active on resume, read its context and preserve its work;
protected edits belong between active changes in separate author commits.

## Applied

(none yet)

## Startup compatibility

- Disposable recipe probe: `/tmp/fireweed-stage2-setup.pyaO9N`, initialized by the
  pinned controller with Go, worker Codex and auditor Codex. Never initialized the
  real checkout. Managed AGENTS block already matches; design checksum matches
  HANDOFF (`52623b89d8184680f535519e82ceda8a49490b3d4dec000347cbd2cde429b21a`).
- Updated version range to `>=0.1, <0.2`, explicit audit enabled/2-round limit and
  5-change unapproved bound; worker/auditor commands now use the exact Sol/high
  ephemeral, ignore-user-config argv. No scope or protected paths broadened.
- Replaced obsolete canary folders with the recipe's version-1 YAML maps, adapted
  to this project's evals/ layout; template now uses {{name}} and api.go.tmpl.
  Retained existing stdlib Go eval/boundary enforcement, including package failures
  and unfinished tests; no generated Python hook adopted. Setup copies the invoking
  FIREWEED_BIN to scratch trees and provides a controller PATH alias.
- `doctor --json` exited 0; evidence `.fireweed/evidence/startup-doctor.json`
  and `.stderr` (empty). The command included agent checks, not --no-agent.
  Initial scaffold `go build -o bin/fireweed .` exited 0 (no usable commands yet).
- No implementation or worker change has started.
