---
agent: opencode
base_commit: 029e78c6804931a9274a47c82a2586bec0282b3e
date: "2026-11-19T11:30:30.673302Z"
model: opencode/inkling-2-small-free#high
result: fail
---

## Own evals

## Dependent evals

## Contradicted Decisions
None found. The four retained decisions are faithfully implementable and consistent with the FLOW specs and the evals:

- [0012-manual-lifecycle] Provenance lines carry the phase, a UTC RFC3339 timestamp, `agent`, `fireweed` and a `playbook=<full SHA-256>` of the exact text selected under the FLOW-143 playbook rule, plus `tests=N failing=N passing=N` for the reports verified by that Done call. Existing nonblank provenance lines are preserved verbatim and in order; blank separators inside the section are normalized.
- [0119-verification-costs] Cost snapshots land in `.fireweed/runs/<run>/<phase>/costs.json` beside the module reports and are replaced, whole, on each repeated verification of that phase. Line counts are newline-count plus one when nonempty content lacks a trailing newline, regular files only, excluding higher-priority zones via `ResolveZone`.
- [0120-status-report] Waiting annotations are the first earlier blocked change in queue order, with same-module, depends-on-blocked and interface-dependency precedence; applied order is implementation Provenance date, newest first, archive path as tie-breaker; PhaseCommits keys are archived filenames without `.md`; budget diagnostics from the selected snapshots are returned once each.
- [0123-pr-body-report] Selected changes are identities touched in the range and present at its end; removed change documents are absent from that endpoint report; phase commits reachable at the endpoint supply the review context; binary follows Git's text/binary classification (NUL in the payload); `256 KB is 262144` bytes with strictly larger text files warned; cost measurements render one row per measured module in a Cost/Metric section.

## Gaps
Each entry names the nearest spec ID (or "no spec") and the module that should own the fixing change. All are packages/workflow unless noted.

- FLOW-021 vs FLOW-137: the specs do not state that the `blocked:` frontmatter key is worker-writable and is exempt from the "recorded runtime state" (P03) comparison. I treat it that way, as FLOW-137's worker-declared block requires it; FLOW-021 never mentions the exception. Nearest: FLOW-021. Fixing change: packages/workflow.
- FLOW-058 vs FLOW-141/142: the stale-base *error* P05 (FLOW-058) and the base-tip *warning* P05 (FLOW-141/142) reuse one code for two severities with no spec text distinguishing them; only the presence of an error/warning flag disambiguates. Nearest: FLOW-141. Fixing change: packages/workflow.
- FLOW-062: "Lifecycle operations MUST preserve assigned static, eval, zone and boundary causes … except for lifecycle stale-base classification P05" does not pin down whether the absent-from-projection C08 (a Modified/Removed spec that no longer exists) is stale-base (P05) or a plain static failure (C08). I classify runtime staleness by recorded-hash mismatch; absent projections are left as-is (C08). Nearest: FLOW-062. Fixing change: packages/workflow.
- FLOW-134 vs FLOW-024/FLOW-026: "Cancellation, timeouts, hook execution errors, and invalid reports MAY stop further evaluation" is a permissive clause, so the order in which own-module vs dependent-module failures surface in `Verify` is unspecified; I stop evaluation early on hook execution errors (spec permits) but continue across mapped testcase failures (FLOW-134 requires). Nearest: FLOW-134. Fixing change: packages/workflow.
- FLOW-009 and FLOW-015 vs FLOW-110: Start's P04/P07 computation and Status's FLOW-110 unapproved count both need a shared-base identity set, but the specs do not say whether Start should also count/error when the shared base ref is missing (I fail open with no warning, mirroring FLOW-015's silent-success ambiguity) or whether it should error as FLOW-110 does. Nearest: FLOW-009. Fixing change: packages/workflow.
- FLOW-064/065: when a *tracked* file under `.fireweed/runs/` is modified during the step, FLOW-064 requires the attempt to exclude run artifacts but says nothing about restore; I restore such tracked run artifacts from the phase base rather than keep them, which is stricter than spelled out. Nearest: FLOW-064. Fixing change: packages/workflow.
- FLOW-073/075: the saved-attempt restore merges only worker-editable body sections (`Decisions`, `Questions`, `Audit`, `Notes`) from the attempt document over the current approved document; the spec does not enumerate which body sections are worker-editable nor the precedence rule for whole-section replacement vs. line-level merge. Nearest: FLOW-075. Fixing change: packages/workflow.
- FLOW-035: the spec does not say how Done recovers the *preceding* audit round's retained finding IDs (required for the re-audit discard rule). I read the change document at HEAD (the pre-submission commit) as the preceding round's retained set, which is correct for the tested flow but assumes the submission is never committed before Done runs — true in every exercised path but unstated. Nearest: FLOW-035. Fixing change: packages/workflow.
- FLOW-034: which spec text is authoritative for the "quoted clause is an exact substring of that spec" check after the change modifies that same spec is unstated; I compare against the approved *projected* spec from the held snapshot. Nearest: FLOW-034. Fixing change: packages/workflow.
- FLOW-011/FLOW-016 vs FLOW-061: the valid-changes list embedded in U02 errors is unspecified — whether it should list pending only, pending+archive, or be ordered is up to the implementation; I list pending+archived slugs in resolution order. Nearest: FLOW-061. Fixing change: packages/workflow.
- FLOW-031/FLOW-146: Done's post-verification zone re-scan raises Z01 for hooks that edit protected/paths during verification, but FLOW-002 only zones edits *before* hooks; the post-hook zone check and its ordering relative to FLOW-137's worker-block detection before advancing is unspecified. Nearest: FLOW-136. Fixing change: packages/workflow.
- FLOW-045 vs FLOW-050: the exact rendering of module metadata keys under `set:` when both the spec `set` and the module frontmatter carry YAML-typed values (int vs. float preservation under round-trip) is unspecified; the implementation hand-tags scalars to preserve YAML-typed configuration verbatim through the claim. Nearest: FLOW-045. Fixing change: packages/workflow.
- no spec: how the assigned spec ID width is chosen (3-digit, zero-padded) and what happens past 999 (or non-digit-prefix IDs) is unstated. Nearest: FLOW-007. Fixing change: packages/workflow.
- no spec: the canonical shape of the claim blob (`refs/fireweed/claims/<slug>`), its JSON/YAML encoding choice and the ordering between claim CAS and subsequent worktree mutation are implementation-freedom areas; the evals require atomic claim creation and exact-OID release but say nothing else. Nearest: FLOW-001/FLOW-047. Fixing change: packages/workflow.
- no spec: whether a `blocked` change that is later re-queued for eligibility is considered a "known" or "unknown" explicit selector for FLOW-011/FLOW-061 disambiguation is unspecified; I treat it as known-but-not-eligible (P01), not U02. Nearest: FLOW-011. Fixing change: packages/workflow.
- no spec: the subject field for X01/X02 boundary diagnostics when the violation message has no module prefix is unspecified; I derive the subject from the first colon in the module-attributed message when available, else leave it empty. Nearest: FLOW-027/FLOW-062. Fixing change: packages/workflow.

## Agent observations

### Invocation 1
Log: .fireweed/runs/20261119T113030-6f3deacfc761e2825b900688/drill-1.log

Now let me read the eval tests — these define the expected behavior. Starting with the smaller ones.
Let me examine the dependency internals I'll build against.
Now let me study the project module internals to understand queue/projection semantics.
I now have a complete picture. Let me verify the toolchain and study a few remaining pieces (assets playbooks, checks.Evaluate tail, source helpers) before designing.
Now I have the full picture. Let me set up a scratch branch and start building the internal package.
CheckCompatibility found a critical constraint: claim snapshots must round-trip Go value types (yaml-typed), not encoding/json. Let me fix the claim serialization to use yaml.
Now the core resolution and shared held-change machinery:
Now the audit submission engine:
Now the shared eligibility and base-ref helpers, then Start itself.
Let me rewrite `start_flow.go` cleanly — I rushed it:
Let me consolidate — rewriting `start_flow.go` and trimming `start.go` of duplicated helpers:
All stubs. Now let me build up from the core. First, how far Start gets:
Type fidelity between YAML roundtrips breaks integral float64 values. Let me probe yaml.v3 behavior:
YAML can't roundtrip int vs float64. I'll add a scalar tagging codec to the claim:
Now the document editing helpers, then the verification engine:
Now the core verification engine:
Now the Done transitions and blocking:
Good progress — 38 pass. Let me implement the remaining pieces. First Release, Commits, GC:
Now the Block flow:
The committed Audit section needs normalization on No-findings acceptance. Let me fix:
Two bugs: audit submissions with findings must not require "No findings", and the git add exclusion is wrong. Let me verify the exclusion:
`git add` with exclusions is fragile. Let me build the tree from explicit paths instead:
Found it — `filepath.Join(root, absoluteSpecPath)` doubled the path. Fixing:
Context is next:
Let me implement StartFromAttempt (resume):
Now Abandon:
The registered worktree path contains a newline, which git C-quotes in porcelain output. Handling quoted paths:
Git prints the newline path raw, breaking line-splitting. Parsing by blank-line-separated records instead:
Now the Status implementation:
All Status tests pass. Now PRBody — the last major piece:
Let me rewrite prbody.go completely and carefully:
Three fixes needed: coverage check in impl Verify, X02 code for boundary hook errors, and byte-preserving Decisions sections:
Two more surgical fixes — byte-preserving frontmatter and config-level protected patterns:
All 109 tests pass. Let me check robustness — race detector and repeat runs:
Now let me run the target eval hook as instructed:
All 218 test cases pass. Let me check repo conventions and then write DRILL_NOTES.md:
Repo conventions honored (no commits). Now let me write DRILL_NOTES.md — reflecting honestly on unresolved design points:
Let me do a final hygiene pass — dead code removal and a fresh look at the main flow:


## Failures
worker agent opencode: context deadline exceeded
