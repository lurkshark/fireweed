---
module: gitstore
---

## Motivation

The design requires local inspection without networking. Ordinary public Clean probes on installed Git2.30.1 prove a successful clean result invokes a configured agent-shaped fsmonitor callback. With required HEAD commit or root-tree objects unavailable, Clean attempts fetch before returning an error. Both are violations even though the final boolean/error is otherwise useful. Native marker controls are live; original index/config/HEAD/files remain preserved in these discovery fixtures.

The same discovery proves missing irrelevant parents/old blobs and a missing current blob body in a simple dirty comparison can remain usable without fetch. Do not require a complete object graph. Preserve built-in normalization, attributes, excludes and index/worktree settings; inspection is not permission to run configured external conversion or monitor commands.

The preceding git-read-locality implementation also corrected a public/native compatibility regression for case-distinct packed tags and existing reflog/pseudoref expressions. Retain permanent passing controls for those GIT-011 cases in this owning-module EVALS step, alongside genuinely failing new Clean regressions. Those controls do not replace the new fail-first obligation.

## Interface

No public API/type changes. Keep Clean, Edits and every retained signature unchanged. Add public-API evals with independent native built-in clean/dirty premises, live configured monitor/content-conversion markers, real allowed/denied transport observations, required versus irrelevant missing objects, ordinary/linked SHA-1/SHA-256 repositories, ordinary unborn and invalid/nested roots, state preservation and owned-temp cleanup. Retain built-in CRLF/config include/attributes/excludes/split-index/timestamp/filemode behavior. Add passing native/public case-distinct packed-tag, HEAD/branch/symbolic reflog and selected-worktree pseudoref Edits compatibility controls for GIT-011. Existing evals and interfaces remain intact; no bootstrap or authenticated agent is used as the subject or a fixture dependency.

## Added

### new:clean-local-only
Clean MUST inspect only locally available repository state without attempting fetch, push, branch creation or network commands. Configured monitor and external content-conversion callbacks MUST NOT run, and no bootstrap or authenticated agent MAY be required or invoked. It MUST retain GIT-003 clean/dirty meaning for staged, unstaged, deleted and untracked files while ignoring ignored files, including native built-in normalization, attributes, excludes and index/worktree configuration on available ordinary and linked SHA-1/SHA-256 repositories. Valid ordinary unborn repositories and nested checkout paths MUST remain usable. Missing objects required to determine cleanliness and invalid roots MUST return false with a useful non-nil error; missing unrelated ancestors, trees or blob bodies not needed by the comparison MUST NOT prevent otherwise available clean or dirty results. Successful and rejected calls MUST be noninteractive, preserve original HEAD, refs, index bytes, configuration, working/ignored files and parent cwd/environment, and clean their owned temporary storage.

## Modified

## Removed

## Decisions

Use only public gitstore APIs in permanent evals. First establish each available fixture's expected built-in clean/dirty result independently with real native Git and a manual boolean expectation. Configure marker-only callbacks after that premise, prove the callbacks are live natively, clear their observations, then test Clean. Required-object fixtures must calibrate actual transport observations, including denied protocol fetch attempts. Keep useful available comparisons with missing bodies; neither object reachability nor complete object inventory is a prerequisite.

The original caller configuration and environment cannot be rewritten to suppress callbacks or fetch. Implementation choices must preserve required built-in behavior and original state. The existing owned Git view is a possible reusable implementation, not a mandated architecture or a pre-EVALS change. Frozen eval/interface immutability and a fresh restricted formal audit remain required. Additional passing GIT-011 metadata-expression controls exercise corrected behavior; do not use already_satisfied or maintenance to bypass genuine Clean failures.

The user's full rebuild goal authorizes local authored execution and commits under the pinned bootstrap controller. Record the normal unmerged manual-start warning without claiming a human merge. Keep NEW product and governing bootstrap separate; no pushes or publication. Workflow approval/exact-HEAD locality, checked Verify/Done, runner/recovery, real regeneration/self-hosting and full acceptance remain subsequent required work.

## Rejected alternatives

A successful boolean after an unexpected configured callback, or an error after a fetch attempt, still violates locality. GIT_NO_LAZY_FETCH alone is ineffective on installed Git2.30.1. Requiring every reachable object would reject proven valid comparisons. Rewriting original Git configuration/index, ignoring ignored-file rules or built-in normalization, weakening frozen evals, bypassing audits, or delegating product behavior to the bootstrap is not a fix.
