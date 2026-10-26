# Rebuild progress

Keep this file short: update sections in place and add one line per applied change.

## Status — 2026-10-26

Stage-2 startup and standalone NEW CLI checkpoint passed. Four changes applied
since the last human review; stop after the next applied change. The manual
lifecycle fixture is not built and every acceptance row remains unverified.

## Controller and baseline

- Pinned controller: `.fireweed/bin/fireweed-bootstrap`, copied from the tested
  sibling executable; version `0.1.0`, embedded source
  `7f925190878e38fc38f30d16bb1cf47aca098cf7`, vcs.modified=false.
  Candidate/destination SHA-256:
  `f1516b5686435c23b54b71c5ac554aba6859b5454f072e0c5eb031dfd6f0fd4f`.
- NEW stage-2 setup baseline: `be35552b706aa3000ff8f11306d129ff1d701b85`.
  Use this explicit nonempty baseline..HEAD for history/CI acceptance. Historical
  restart `d991c26472ad990a5d050a07ce7dc873d1900433` is not the new baseline.
- Product: `bin/fireweed`, built from this checkout; outside-project version
  --json passed. help/version/explain/playbook work; other commands explicitly
  fail as unimplemented. No controller delegation.

## Startup evidence

- Initial tree at `0fee737d7bc2c21c6bf5e28faf299301e2542cb1` was clean, with no
  changes/archives/claims. Read all required design/supporting docs and HANDOFF.
  Design SHA-256 matches handoff:
  `52623b89d8184680f535519e82ceda8a49490b3d4dec000347cbd2cde429b21a`.
- Disposable Go init probe: `/tmp/fireweed-stage2-setup.pyaO9N`. Adapted protected
  version range to >=0.1,<0.2, canaries.yaml maps for evals/ layout, {{name}}
  templates and api.go.tmpl. Retained Go JUnit/boundary hook including unfinished
  tests and package failures; setup copies invoking FIREWEED_BIN into scratch.
  Managed AGENTS already matched. Protected paths and bounds not weakened.
- Worker/auditor argv: Codex ephemeral, ignore-user-config, gpt-6-sol, high effort,
  workspace-write. Audit enabled, max_rounds=2; max_unapproved=5. Startup check and
  full doctor passed, including four canaries/removal recovery and BOTH real
  Sol/high probes, no diagnostics. Scratch worktree removed. Evidence is ignored
  under `.fireweed/evidence/startup-doctor.*`.

## Applied

- `standalone-information` / AST-001–AST-003: evals `a4ceb17` (8 genuine failures), independent audit `826b9e9` (no findings), impl `2248b8e` (8/8 pass, direct dependents/boundaries pass); archive `0001-standalone-information.md`.
- `standalone-cli` / APP-001–APP-003: evals `1298acf` (30 genuine failures), controller audit `7de7d30` (independent findings discarded due folded quote line breaks), impl `db77de7` (30/30 pass); archive `0002-standalone-cli.md`. Independent audit did NOT pass; corrected in next eval_fix.
- `json-error-assertions` / APP-003 eval_fix: evals `eb4704d`, independent audit `e184bb6` (no findings), impl `8796d37` (no code edits; all 26 app cases pass); archive `0003-json-error-assertions.md`.
- `lifecycle-primitives` / CON-001–CON-005: evals `11d0180` (27 genuine failures), independent audit `6e341a4` (no findings), impl `8e37dd7` (27/27 own, 26 app and 8 assets cases pass; all direct dependents/boundaries pass); archive `0004-lifecycle-primitives.md`.

## Process findings and limits

- Real workers/auditors used required Sol/high. Manually coordinated controller
  transitions record agent=manual in Provenance; actual worker transcripts are
  `.fireweed/evidence/<slug>-*`. Auditors ran separately in `/tmp/fireweed-audit-*`
  with only copied change/projected spec/evals, never implementation or SYSTEM.
- Standalone-cli audit quotes folded spec line breaks. Controller discarded them
  and advanced to impl; an attempted eval edit was rejected and restored exactly
  before implementation. A separately audited eval_fix repaired the known gaps.
  Future quotes must be literal contiguous substrings; inspect phase/audit_rounds
  after every transition. No forbidden eval/interface edits committed in impl.
- Evals-only B03 warnings disappeared after implementation. No claim remains.
- Contracts worker hit account usage limit after writing its implementation. On
  resumption, full controller verification passed the actual code; controller
  applied it without another worker. Usage API now permits ordinary use (1% of
  5-hour window used). No model/provider substitution or fake-agent evidence.
- On resumption origin/main was observed at `6e341a4`; this process did not push
  or move the base. Review count remains four; do not silently reset it.

## Next / resume

Author one focused project change for loading the minimal fixture and authoring,
queue projection and stamping. Then stop at five applied changes for human review.
After review, wire those operations through app promptly and continue the NEW
manual fail-first/audit/impl/archive fixture, adding only needed Git/checks/workflow
behavior. Do not polish working primitives before that end-to-end slice.

If a change is active on resume, inspect its context, tree and claims first.
Protected author edits belong only between changes, in separate commits.
