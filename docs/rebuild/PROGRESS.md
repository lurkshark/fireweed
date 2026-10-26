# Rebuild progress

Keep this file short: update sections in place and add one line per applied change.

## Status — 2026-10-26

- Stage-2 preparation reviewed (D004); the rebuild has NOT started. No changes
  authored/applied and no claims found during review. All R rows remain unverified.
- The executable pin, live config/hook adaptation and controller doctor checks are
  startup work. `fireweed.yaml` still pins `>=0.4, <0.5`; do not run the new
  controller against that setup before reconciling compatibility.
- A NEW full setup commit SHA must be recorded after successful startup checks;
  use its explicit nonempty `baseline..HEAD` for history/CI acceptance.
- Historical restart baseline: `d991c26472ad990a5d050a07ce7dc873d1900433` (D001).
  Old build/check/doctor results belong to the retired setup, not the new pin.

## Controller candidate and agents

- Tested source executable: `/Users/joel/Workspace/fireweed-bootstrap/bin/fireweed`.
- Candidate: version `0.1.0`, embedded source revision
  `7f925190878e38fc38f30d16bb1cf47aca098cf7`, `vcs.modified=false`; metadata checked
  against HANDOFF.md during preparation. It has NOT been installed for stage 2.
- Destination: `/Users/joel/Workspace/fireweed/.fireweed/bin/fireweed-bootstrap`.
  Copy the tested executable; do not rebuild or consult bootstrap source.
- Codex-only workers/auditors/regeneration: `gpt-6-sol`, high reasoning effort,
  exact argv in GOAL.md. A read-only Sol/low-effort probe exited 0 with the expected
  reply. High effort and both full doctor probes remain unrun startup checks.
- The new controller covers the revised design's bounds, eval_fix and usability;
  retain SYSTEM.md's human checkpoints and scope limits as well.

## Next / resume

Follow GOAL.md preflight: inspect status/claims, verify and install the controller,
adapt protected setup using a disposable recipe fixture, configure Sol/high,
pass controller check/doctor, and commit the new setup baseline. Then build the
NEW binary's narrow manual lifecycle slice before expanding the surface.

If a change is already active on resume, read its context and preserve its work;
protected edits belong between active changes in separate author commits.

## Applied

(none yet)
