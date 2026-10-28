# Contracts drill review

The single authorized controller drill ran at base
`7afe9a0570a8885aadc33bed5a6b5b38c93fa453`. Native result: pass.
[Native report](../../../.fireweed/drills/contracts.md) records Codex CLI
`gpt-6-sol`, high effort, with 34 own and 339 dependent cases passing (373 total).
No contradicted Decision was reported.

Rebuilt code is retained at
`refs/fireweed/drills/contracts/20261028T103005.180922000`
(commit `b101ddebc6d55400308ad56a74df68f5b3ddd465`). The controller removed its
scratch worktree; only the original worktree remains. No pending change or claim
remains. The primary implementation was not replaced, and the applied count
remains four since the last checkpoint.

## Reported gap

CON-011/CON-012, target contracts: item parsing rejects malformed and duplicate
headings, but its specs/Decisions do not define the `malformed item heading: `
and `duplicate item ID: ` prefixes used by dependent diagnostic classification.
The native report repeats this gap in both final sections; it explicitly reports
no contradicted Decisions. No gap fix has been authored.

## Independence limitation

The configured regeneration instructions prohibited deleted implementation and
Git-history consultation. The saved invocation confirms the required model and
effort, but the agent read dependent internal implementations outside the goal's
retained-input-only boundary. Its recorded commands include:

- Before reconstruction, `rg ... packages/project/api.go packages/*/internal --glob '*.go'`.
- After dependent failures, `rg ... packages/project/internal packages/app/internal/entry`.

These are visible in ignored auxiliary log
`.fireweed/runs/drill-contracts/2026-10-28T10:22:49.049835Z/1.log`
(at lines 967 and 3145). No deleted contracts source, old Git source/history or
bootstrap source read was observed in the recorded commands reviewed. Nevertheless,
this run does not establish the strict independence acceptance gate. Preserve
its native pass and useful gap evidence without presenting it as final acceptance.

## Follow-up boundary

The owner authorized exactly this drill and no automatic gap changes or other
module drills. No rerun or gap fix is authorized by this report. Defer those
follow-ups pending human agreement. Continue the separately authorized
Milestone 2 work; full regeneration and final acceptance remain incomplete.
