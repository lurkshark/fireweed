---
module: runner
phase: applied
regenerate: true
run: 20261113T072516-2b52ce34b6d40c44d79ca9dc
---

## Motivation

The owner approved rebuilding runner's disposable implementation from its
complete current specifications to remove accumulated implementation cruft
while preserving behavior and retained contracts. This is the real regeneration
maintenance workflow described in the design's Regeneration section.

## Interface

## Added

## Modified

## Removed

## Decisions

## Notes

Reconstruct only runner from its current Specs, Decisions, retained interfaces,
evals and permitted controller context after Fireweed removes disposable files.
Read `fireweed context regenerate-runner` before each worker step and use
`fireweed verify` with the configured dependent checks. Keep `.fireweed/bin`
first on PATH and `GOCACHE=/tmp/fireweed-rebuild-cache`; the pinned controller
owns removal, transitions, commits and archiving. Wait for each verification
command to finish before editing sources again so its reports describe one
source state.

Do not consult deleted implementations, prior target implementations, Git
history, saved attempts or rebuild-v1 as implementation oracles. Do not read
docs/rebuild/PROGRESS.md or ACCEPTANCE.md. Preserve all Specs, retained
interfaces, eval assertions and unrelated modules. Do not add behavior changes,
eval fixes, compatibility scaffolding or new Decisions. If a genuine spec or
interface ambiguity prevents reconstruction, record the concrete issue in
Questions and block rather than changing the contract.

## Provenance
- impl 2026-11-13T08:00:08.370228Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=5f856d6bcab308b9ff1a8a7c03b9d2547744112aba81dd30dea8f61ec8062b03 tests=539 failing=0 passing=539
