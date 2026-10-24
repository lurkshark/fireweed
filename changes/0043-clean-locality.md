---
module: gitstore
phase: audit
run: 0e11477d477d697832d05857
---

## Motivation

The design requires local inspection without networking. Ordinary public Clean probes on installed Git2.30.1 prove a successful clean result invokes a configured agent-shaped fsmonitor callback. With required HEAD commit or root-tree objects unavailable, Clean attempts fetch before returning an error. Both are violations even though the final boolean/error is otherwise useful. Native marker controls are live; original index/config/HEAD/files remain preserved in these discovery fixtures.

The same discovery proves missing irrelevant parents/old blobs and a missing current blob body in a simple dirty comparison can remain usable without fetch. Do not require a complete object graph. Preserve built-in normalization, attributes, excludes and index/worktree settings; inspection is not permission to run configured external conversion or monitor commands.

The preceding git-read-locality implementation also corrected a public/native compatibility regression for case-distinct packed tags and existing reflog/pseudoref expressions. Retain permanent passing controls for those GIT-011 cases in this owning-module EVALS step, alongside genuinely failing new Clean regressions. Those controls do not replace the new fail-first obligation.

## Interface

No public API/type changes. Keep Clean, Edits and every retained signature unchanged. Add public-API evals with independent native built-in clean/dirty premises, live configured monitor/content-conversion markers, real allowed/denied transport observations, required versus irrelevant missing objects, ordinary/linked SHA-1/SHA-256 repositories, ordinary unborn and invalid/nested roots, state preservation and owned-temp cleanup. Retain built-in CRLF/config include/attributes/excludes/split-index/timestamp/filemode behavior. Add passing native/public case-distinct packed-tag, HEAD/branch/symbolic reflog and selected-worktree pseudoref Edits compatibility controls for GIT-011. Existing evals and interfaces remain intact; no bootstrap or authenticated agent is used as the subject or a fixture dependency.

## Added

### GIT-012
Clean MUST inspect only locally available repository state without attempting fetch, push, branch creation or network commands. Configured monitor and external content-conversion callbacks MUST NOT run, and no bootstrap or authenticated agent MAY be required or invoked. It MUST retain GIT-003 clean/dirty meaning for staged, unstaged, deleted and untracked files while ignoring ignored files, including native built-in normalization, attributes, excludes and index/worktree configuration on available ordinary and linked SHA-1/SHA-256 repositories. Valid ordinary unborn repositories and nested checkout paths MUST remain usable. Missing objects required to determine cleanliness and invalid roots MUST return false with a useful non-nil error; missing unrelated ancestors, trees or blob bodies not needed by the comparison MUST NOT prevent otherwise available clean or dirty results. Successful and rejected calls MUST be noninteractive, preserve original HEAD, refs, index bytes, configuration, working/ignored files and parent cwd/environment, and clean their owned temporary storage.

## Modified

## Removed

## Decisions

Use only public gitstore APIs in permanent evals. First establish each available fixture's expected built-in clean/dirty result independently with real native Git and a manual boolean expectation. Configure marker-only callbacks after that premise, prove the callbacks are live natively, clear their observations, then test Clean. Required-object fixtures must calibrate actual transport observations, including denied protocol fetch attempts. Keep useful available comparisons with missing bodies; neither object reachability nor complete object inventory is a prerequisite.

The original caller configuration and environment cannot be rewritten to suppress callbacks or fetch. Implementation choices must preserve required built-in behavior and original state. The existing owned Git view is a possible reusable implementation, not a mandated architecture or a pre-EVALS change. Frozen eval/interface immutability and a fresh restricted formal audit remain required. Additional passing GIT-011 metadata-expression controls exercise corrected behavior; do not use already_satisfied or maintenance to bypass genuine Clean failures.

The user's full rebuild goal authorizes local authored execution and commits under the pinned bootstrap controller. Record the normal unmerged manual-start warning without claiming a human merge. Keep NEW product and governing bootstrap separate; no pushes or publication. Workflow approval/exact-HEAD locality, checked Verify/Done, runner/recovery, real regeneration/self-hosting and full acceptance remain subsequent required work.

Manual start assigns GIT-012/run0e11477d477d697832d05857 in EVALS and emits the normal P04 unmerged warning; execution uses the user's local goal authority without claiming a human merge. Only new eval files are being added. Existing238 evals/API and all disposable implementation hashes remain unchanged. The initial converter premise was excluded because a size-only dirty result skipped content conversion; use equal-size changed bytes and an explicit timestamp so native callback observation is live. Corrected preflight has24 genuine callback failures and24 passing GIT-011 metadata controls, with no fixture-premise failures. Complete fail-first and separate formal Audit are still required.

Pre-freeze restricted review rejects forcing a stored HEAD-tree body when an available index can establish the correct comparison, and rejects forcing an error when an already-dirty worktree establishes false. Eval outcomes allow those valid local alternatives while always rejecting callbacks/fetches; clean-index unavailable HEAD commit and invalid roots retain strict false/useful-error checks. Add clean/dirty/unavailable-HEAD variants for both converter types, staged unborn coverage and invalid-root sentinel preservation. Built-in CRLF fixtures record and verify the normalized HEAD/index blob, then force a fresh timestamp comparison; initial false native premises are excluded. No implementation or retained eval/API is edited.

Final source-only pre-freeze review reports No findings for projected GIT-012, retained GIT-003 meaning and GIT-011 metadata compatibility, after calibrated attribute precedence on the same normalized workfile/index. New-test preflight182 records has75 genuine failure records and107 passing,0 skips or fixture-premise failures; separate24-leaf metadata matrix passes. Existing238-record evals/API and all implementation files remain byte-identical to applied a62ec6a. Complete445-record controller fail-first still required before EVALS done and fresh formal Audit.

The first complete445-record controller check runs187.919s and confirms75 failures/370 passing, but correctly rejects C14 for GIT-003: the new configuration test name referenced both GIT-012 and GIT-003, so new index-byte preservation failures were attributed to the older ID whose boolean semantics pass. Rename only that new eval to GIT-012, which owns preservation; retain every original eval/API/private byte and all assertions. This rejected check is excluded from the accepted fail-first gate; repeat full verification. No fixes/already_satisfied or coverage-policy bypass is added.

## Rejected alternatives

A successful boolean after an unexpected configured callback, or an error after a fetch attempt, still violates locality. GIT_NO_LAZY_FETCH alone is ineffective on installed Git2.30.1. Requiring every reachable object would reject proven valid comparisons. Rewriting original Git configuration/index, ignoring ignored-file rules or built-in normalization, weakening frozen evals, bypassing audits, or delegating product behavior to the bootstrap is not a fix.

## Provenance

- evals | 2026-10-24T00:41:20Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 445 | failing: 75 | passing: 370
