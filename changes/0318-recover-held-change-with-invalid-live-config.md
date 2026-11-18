---
module: app
---

## Motivation

A malformed live fireweed.yaml blocked the CLI from showing held context or
releasing a claim, while the workflow APIs could do both. Users need those recovery
operations after a configuration edit goes wrong. APP-011 and APP-091 describe
operations on the held change.

## Interface

Keep the existing context and release arguments, output, and ownership rules.

## Added

### new:held-context-invalid-live-config
After normal project discovery, context for a selected change owned by the caller's
worktree MUST return the held context even when live fireweed.yaml is malformed or
its binary pin excludes the running binary.

### new:release-invalid-live-config
After normal project discovery, release for a selected claim owned by the caller's
worktree MUST release that claim even when live fireweed.yaml is malformed or its
binary pin excludes the running binary.

## Modified

## Removed

## Decisions

## Notes

Use the existing nested CLI fixture and start an owned change with compatible
configuration. Edit the live file after Start. Exercise malformed YAML and an
excluding live pin against context and release. Check that context returns the held
change and projected spec. Verify that release drops only the owned claim and
preserves HEAD, the index and edits. Give release an unrelated change selector
while the live configuration is malformed and confirm the caller's claim survives.
Reuse the existing selector and ownership assertions. Keep fresh start and other
live-config consumers on their existing admission paths.
