---
module: app
regenerate: true
---

## Motivation

The owner approved rebuilding app's disposable implementation from its
complete current specifications to remove accumulated implementation cruft
while preserving behavior and retained contracts. This is the real regeneration
maintenance workflow described in the design's Regeneration section.

## Interface

## Added

## Modified

## Removed

## Decisions

## Notes

Reconstruct only app from its current Specs, Decisions, retained interfaces,
evals and permitted controller context after Fireweed removes disposable files.
Read `fireweed context regenerate-app` before each worker step and use
`fireweed verify` with the configured dependent checks. Keep `.fireweed/bin`
first on PATH and `GOCACHE=/tmp/fireweed-rebuild-cache`; the pinned controller
owns removal, transitions, commits and archiving. Wait for each verification
command to finish before editing sources again so its reports describe one
source state.

Do not consult deleted implementations, prior target implementations, Git
history, saved attempts or rebuild-v1 as implementation oracles. Do not read
docs/rebuild/PROGRESS.md or ACCEPTANCE.md. Preserve all Specs, retained
interfaces, eval assertions and unrelated modules, including the APP-156–APP-158
help contracts and APP-159/APP-160 prerequisites. Do not add behavior changes,
eval fixes, compatibility scaffolding or new Decisions. If a genuine spec or
interface ambiguity prevents reconstruction, record the concrete issue in
Questions and block rather than changing the contract.

