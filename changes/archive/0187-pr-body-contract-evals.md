---
bases: {}
eval_fix:
    - APP-127
    - APP-128
    - APP-129
    - APP-130
    - APP-131
module: app
phase: applied
run: 20261108T081712-dc9c6122b165e3c07cd35635
---

## Motivation

The owner requested correction of the PRBody eval mismatches exposed by review
of rejected drill d63c16e. The CLI fixture requires every workflow B04 diagnostic
to have Path although FLOW-124 does not mandate that field, and APP-129 preserves
path only when present. Its empty-range assertion also rejects compliant "No
selected changes." wording. Align those assertions with the existing contracts
without losing artifact identification, CLI fidelity or failure behavior.

## Interface

Audited eval-only correction against the existing public contracts. Edit only
the target module evals and this change in evals; implementation changes only
this change. Preserve behavioral coverage and require all own evals to pass.

## Added

## Modified

## Removed

## Decisions

## Notes

Limit edits to packages/app/evals/pr_body_cli_test.go and a small local helper
only if useful. Remove the fixture's unconditional demand that workflow.PRBody
populate B04 Diagnostic.Path. Keep artifact identification and byte-size behavior
where already asserted, B04 warning presence and all B01-B03 warnings. Preserve
the exact plain body and diagnostic code/message checks, JSON schema/string body,
empty stderr, diagnostic count and equality to all diagnostics actually returned
by workflow.PRBody, including any present Path. Do not normalize away, drop or
ignore Path during CLI/public-result comparison. A nonempty artifact Path may
still be checked for correctness without making it mandatory.

Accept explicit "No selected changes." as well as previously accepted empty
selection forms. Blank/non-explicit empty reporting must still fail. Retain
valid-empty success, nested-root/historical-endpoint authority, warning-only
success and all invalid-range/options/config/discovery exit, cause, null-data
and stderr assertions for APP-127 through APP-131. No runtime/signature/contract
change, cross-module helper or verifier weakening.

Workers use the unchanged pinned controller with the current tree's .fireweed/bin
first on PATH and GOCACHE=/tmp/fireweed-rebuild-cache. Read pinned context before
each held step, run pinned verify --quick, wait for its terminal result and return.
Do not commit, invoke lifecycle mutators, edit run/phase/Provenance or protected
planning/tooling. The outside controller owns full phase verification, independent
restricted audit, commits, archive and claim release. The auditor edits Audit
only, checks the listed eval_fix IDs and preserves configured rounds. A passing
submission says No findings and lists the checked IDs.

This is one of exactly three owner-authorized eval_fix changes. No runtime,
interface/model, spec, Decision, configuration, hook, worker-policy or tooling
change is authorized. Keep every unrelated assertion and every covered ID.
Implementation edits only this change file; frozen evals stay unchanged.
Use current public contracts/evals, without bootstrap/rebuild-v1, saved
reconstruction source or other implementation oracles. Preserve prior evidence,
proof/attempt refs, bounds/counters, controller/recoveries and unrelated worktrees.
Verify and stop after all three changes; no drill, reconstruction import,
controller promotion, later work or push is authorized. The existing deadline
2026-11-10 22:51:05 UTC remains unchanged.
## Provenance
- evals 2026-11-08T08:28:07.969254Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=323 failing=0 passing=323
- audit 2026-11-08T08:28:25.413553Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-astra -c model_reasoning_effort="medium" -s workspace-write {prompt} fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
- impl 2026-11-08T08:39:45.751265Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=5f856d6bcab308b9ff1a8a7c03b9d2547744112aba81dd30dea8f61ec8062b03 tests=323 failing=0 passing=323
## Audit

No findings; checked APP-127, APP-128, APP-129, APP-130, APP-131
