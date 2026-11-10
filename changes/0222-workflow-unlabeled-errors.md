---
bases:
    FLOW-133: 7de84872
module: workflow
---

## Motivation

SYSTEM.md's diagnostic-code convention says error text does not repeat its
code once the app reads codes only through `contracts.ErrorCode`, which
0214 (app-reads-error-codes) does. Workflow's error text still begins with
the code, and FLOW-133 fixes the whole text including the label. Remove the
labels and restate FLOW-133.

## Interface

No signature changes. Error text no longer includes the code label.

## Added

## Modified

### FLOW-133
For a valid queue, Current(root, "") MUST return an error whose contracts ErrorCode is P02 and whose text is `worktree holds no change` when the calling worktree owns none of its change claims.

## Removed

## Decisions

## Notes

Remove each leading `<code>: ` from workflow error text and keep the code in
`contracts.CodedError`. Update workflow's own evals that still match a
leading P code (0211 moved the modified items' evals to
`contracts.ErrorCode`). Run after 0220 and 0221, so that wrapped project and
checks causes are already unlabeled.
