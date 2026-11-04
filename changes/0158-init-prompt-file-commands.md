---
bases:
    APP-053: 724ef7fa
    APP-054: 35e56075
module: app
phase: audit
run: 20261104T082316-c3cd0c4234bd8471a3269d76
---

## Motivation

ACCEPTANCE C3 adds prompt-file transport to configured agent and audit commands.
Runner and doctor now accept {prompt_file}, but init rejects a custom command
unless it contains {prompt}. Let users select the supported file transport
at initialization while retaining existing presets and prompt templates.

## Interface

No public fields or signatures change. Init accepts either supported prompt
placeholder for custom worker and audit templates and retains their text.

## Added

## Modified

### APP-053
init MUST configure codex or claude argv command presets, or preserve a custom agent command template containing {prompt} or {prompt_file} as a shell string, in agent.command.

### APP-054
init MUST enable audit by default and write a separately selected --audit-agent command when supplied, accepting the same presets and custom template forms as --agent, otherwise leaving audit.command unset for agent fallback.

## Removed

## Removed decisions

[0069-project-initialization] Init accepts codex and claude presets as argv commands; other selections are shell command templates containing {prompt}, stored as strings for one-pass quoted substitution when execution is implemented. Local recipes use recipe.yaml metadata, recipe.md instructions and a files/ project-relative payload; none uses empty language defaults and a generic module scaffold.

## Decisions

Local recipes use recipe.yaml metadata, recipe.md instructions and a files/ project-relative payload; none uses empty language defaults and a generic module scaffold.

## Notes

Preserve every existing assertion, presets, literal {prompt} command storage,
audit enablement/fallback, recipe handling, diagnostics, generated layout and
non-execution of agents/hooks. The superseded Decision's command forms are
now explicit in APP-053/054; its recipe conventions remain in Decisions.
Do not manually edit module SPEC/Decisions or change version-range output,
recipe inventory, controller configuration, protected infrastructure or timeouts.
C4 follows only after C3 completes.

Use public CLI evals. Exercise file-only worker and auditor selections
independently so both Modified IDs have genuine fail-first assertions; retain
all earlier coverage. Reuse fixtures and keep own Verify below five minutes.

Workers run pinned fireweed context, verify with fireweed verify --quick,
poll until terminal and return. They never commit, call lifecycle mutators,
edit run/phase/Provenance or change frozen evals/interfaces. Only the outside
owning controller advances, commits, applies and archives with normal full
verification.

The separate restricted auditor edits only the existing Audit body, retaining
all other bytes. Findings use - ID: "exact clause" — missing assertion. Passing
audit writes No findings and Checked IDs. Re-audits cover only flagged IDs,
retaining previous clearances and every counter/bound.

## Audit

## Provenance
- evals 2026-11-04T08:38:36.574541Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=306 failing=2 passing=304
