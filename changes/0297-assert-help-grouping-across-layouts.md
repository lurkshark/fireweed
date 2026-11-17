---
eval_fix:
    - AST-008
module: assets
phase: audit
run: 20261117T003204-f999b24a830cb04837f0fe36
---

## Motivation

AST-008 requires the help overview to group commands by purpose. The assets
eval's label-to-command proximity regex led the retained reconstruction to
repeat command names in headings above a readable list of entries. Users
benefit from grouping and recognizable command names in either a concise
row or a headed list.

## Interface

## Added

## Modified

## Removed

## Decisions

## Notes

Keep AST-008 coverage for Fireweed identification, complete command listing,
purpose grouping, navigation, a concrete example and keeping command-specific
options and configuration details on command pages. Check that actual
command entries appear under their relevant purposes. Support both grouped
rows and commands in lists under headings, with flexible label wording.
A heading may simply name the purpose, and the test should associate the
entries that follow it with that group. Avoid counting examples or incidental
command mentions as entries.
## Provenance

- evals 2026-11-17T00:37:45.312842Z agent=codex fireweed=0.2.1 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=151 failing=0 passing=151

