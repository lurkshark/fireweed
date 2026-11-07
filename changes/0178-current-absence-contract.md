---
module: workflow
bases: {}
already_satisfied:
  - new:current-absence
---

## Motivation

The owner approved the two focused follow-ups to the failed WF3 workflow drill.
The design makes next/run thin lifecycle loops that resume a held change or
start eligible queued work. The independent reconstruction returned an error
for no held claim that runner treated as fatal, preventing fresh queue work.
Make the existing public absence signal a retained workflow obligation so it
can be reconstructed without reading runner internals. Preserve D006's P02
claim cause allocation and the existing interface and behavior.

## Interface

No signature or behavior change. Document and assert the existing Current
empty-selector error used by composing callers. Explicit selectors keep their
existing unknown/foreign-claim behavior.

## Added

### new:current-absence
For a valid queue, Current(root, "") MUST return an error whose text is `P02: worktree holds no change` when the calling worktree owns none of its change claims.

## Modified

## Removed

## Decisions

## Notes

This new item is legitimately already_satisfied: the current public operation
already emits exactly this P02 signal, which the current runner uses to select
fresh queue work. Add a focused public workflow eval using the existing local
fixture. Cover an unclaimed queued change, an empty valid queue, and a linked
worktree whose queue has a claim owned by the other worktree. Require the exact
error text in each absence case. Keep existing owned/explicit selector and
claim tests intact; do not broaden the absence signal to other errors.
Use the ID assigned by Start in the testcase name. No new API/classifier,
dependency, runner edit, error-code reassignment or private test is needed.

Implementation already satisfies the added clause. Evals add the public
assertions; implementation changes no runtime source. Preserve every retained
assertion and all current Specs/Decisions. Do not import failed reconstruction
source or repair its other mistakes. Do not read bootstrap/rebuild-v1.

Workers use pinned fireweed context before each held step, run pinned
fireweed verify --quick and poll to terminal, then return. They never commit,
invoke lifecycle mutators, edit lifecycle metadata/Provenance, or edit protected
planning/setup. Only the outside controller advances, commits, applies and
archives. Use exactly the configured model/effort, normal approval enforcement
and unchanged bounds/counters. Stop on a blocked change or access failure.

The independent restricted auditor reviews only retained change/spec/eval inputs
and edits the Audit body only. Passing audit says No findings and lists the
assigned ID; findings name that ID and quote an exact clause without an asserting
test. Re-audits retain clearances, audit rounds and attempt bounds.

This owner-selected follow-up is limited to this contract change and the
separate FLOW-002 eval_fix. After both apply, validate a clean identified
candidate and commit the evidence/handoff. Stop at the exact-candidate WF2
approval gate before another native drill. Preserve prior reports and refs,
the controller pin, recorded deadline and unrelated worktrees. No push.
