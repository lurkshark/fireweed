---
audit_rounds: 1
bases: {}
module: runner
phase: impl
run: 20261031T220623-e826454d4fd1bcfca2375a01
---

## Motivation

The design's Regeneration section and R16 require `fireweed drill contracts`
to rebuild disposable code, verify own and all transitive dependent evals,
retain the rebuilt ref and write a reviewable native report. Runner currently
has no drill API. Implement that prerequisite for the app command, including
separation of reconstruction inputs from the verifier's full scratch checkout,
so the bounded independent contracts proof can use NEW without changing main's
controller pin. Return to app once this API is usable.

## Interface

Add `Drill(context.Context, DrillRequest) (DrillResult, error)` to api.go,
delegating to internal implementation, and matching aliases/types in model/.
DrillRequest has Root, Module, Bin string and Timeout time.Duration.
DrillResult has Module, Base, Date, Agent, Model, Result, ReportPath, Ref string;
Own checks.Report; Dependents map[string]checks.Report; Steps []AgentResult.
Result is pass, fail or stopped; ReportPath is project-relative and Ref is the
retained rebuilt-code ref. Existing APIs and behavior remain unchanged.

## Added

### RUN-050
Drill MUST reject an unknown module, dirty invoking checkout or any live shared change claim before setup or agent invocation, leaving the invoking checkout and refs unchanged.

### RUN-051
Drill MUST create a detached scratch worktree at the invoking HEAD under the common Git directory's fireweed/worktrees directory and run the configured setup hook there with FIREWEED_BIN set to request Bin. Setup failure MUST prevent agent invocation and return its captured failure.

### RUN-052
Drill MUST delete only target-module files classified disposable by the approved zone precedence before reconstruction. It MUST retain every other target-module file and every other module's files in the verification worktree.

### RUN-053
Drill MUST invoke reconstruction in a separate input workspace containing the target's retained files and only interface files from its transitive dependencies, plus SYSTEM.md and fresh independent Git metadata. It MUST exclude original implementation, other-module internals, source history, saved attempts and binaries from those inputs.

### RUN-054
The reconstruction prompt MUST include the target spec and Decisions, identify its retained interfaces and evals, and require rebuilding only disposable paths without reading outside supplied inputs. It MUST request separate Contradicted Decisions and Gaps sections in DRILL_NOTES.md, each gap naming the nearest spec ID or "no spec" and a target module.

### RUN-055
Drill MUST accept reconstruction edits only to target disposable paths and DRILL_NOTES.md, rejecting retained-file changes, deletions or other additions before importing code. It MUST import accepted rebuilt disposable files into the verification worktree without changing retained files or the invoking implementation.

### RUN-056
Drill MUST use the configured worker command, wrapper and timeout through InvokeAgent, retaining ordered AgentResults in Steps. Non-usage failures and verification retries MUST share the configured attempts bound, retain rebuilt work and include the preceding failure in retry prompts.

### RUN-057
Drill MUST treat usage-limit responses as the same attempt, applying configured limit_wait and limit_budget without spending a failed attempt. An exhausted usage budget MUST return an IsUsageStop error with Result stopped; cancellation or timeout MUST return an error and stop reconstruction.

### RUN-058
After accepted reconstruction, Drill MUST run the target module's eval hook and require passing enforced cases and coverage through checks. It MUST retain the complete returned own report in Own, including failed testcase details when verification fails.

### RUN-059
Drill MUST evaluate every direct and transitive dependent in lexical module-name order regardless of verify_dependents, retaining each report in Dependents. Every dependent MUST satisfy passing enforced cases and coverage for a passing drill, even when own or another dependent's evals fail.

### RUN-060
After reconstruction has been invoked, Drill MUST write .fireweed/drills/<module>.md with date, base_commit, agent, model and result frontmatter and Own evals, Dependent evals, Contradicted Decisions and Gaps sections. Eval entries MUST identify module, testcase, spec IDs and result; agent observations and verification failures MUST remain visible, with each gap identifying a nearest spec ID or "no spec" and target module.

### RUN-061
After accepted reconstruction, Drill MUST retain the rebuilt verification tree at refs/fireweed/drills/<module>/<date>, returning that ref without moving the invoking HEAD. A passing result MUST require accepted reconstruction and passing own and all dependent evals; failure MUST return a nonnil error with Result fail.

### RUN-062
Drill MUST remove its verification worktree and input workspace on success, agent or verification failure, invalid edits, usage stop and cancellation. Cleanup failure MUST be returned as an error rather than reported as a passing drill.

## Modified

## Removed

## Decisions

## Notes

The owner explicitly approved a one-time frozen-eval recovery exception on
2026-10-31 after the normal eval_fix path proved unavailable for this unapplied
spec change. Author commit e19a9d6 corrects only the usage_once, agent_retry and
verify_retry fake-agent branches. Their successful new writes require the prior
PARTIAL_WORK/wrong value, preserving the retry-work retention assertion; the
mixed_exhaust scenario remains failing. Testcase names, IDs, assertions, coverage,
interfaces and configured bounds are unchanged.

A separate input-only Codex review raised a RUN-056 retention finding on an
unconditional draft write. The scoped guarded-write correction passed the fresh
restricted re-audit: No findings; Checked IDs RUN-056 and RUN-057, with RUN-057
carrying forward its prior pass. Both invocations used the exact configured
model/effort, changed only their copied Audit section, had fresh one-commit Git
metadata and were cleaned up. The corrected RUN-056/RUN-057 tests passed against
saved attempt 1. This owner exception preserves the original run, assigned IDs,
phase, audit_rounds and retained attempt refs; it does not replace controller
Provenance or the original passing Audit section. Full controller checks remain
required when resuming that attempt.

The remaining implementation review is confined to existing native report
obligations: the report date must be readable by existing date-based drift
tracking, and model metadata must identify the configured command's actual
model rather than be blank when only -m supplies it. Keep ref-safe date naming
and the frozen interface compatible. Do not change frozen evals to hide these
integration details, add unrelated behavior or author automatic drill gap fixes.

Use public deterministic fixture hooks and bounded fake agents; no real Codex
access or bootstrap in Go tests. Reuse existing public checks, contracts zone
resolution, agent invocation and Git/process primitives. A clean input Git
repository may contain only its retained-input initial commit, never alternates
or shared history. The verifier keeps the full project manifests and dependent
implementations; a worker need not be able to run full-project tests in its
restricted workspace. SYSTEM is retained guidance, not permission to access its
mentioned paths. The prompt must explicitly prohibit old code/history/binaries
and external trees, including attempts to repair dependent internals. Never copy
the setup-installed binary, root hooks/configuration, AGENTS or source repository
selectors into reconstruction inputs. Do not expose the verification worktree
or original checkout in reconstruction prompts. DRILL_NOTES.md is an observation
channel, not an automatically authored change. Failed case details must inform
retries and the final report; no automatic gap fixes or change files.

Normal fail-first, independent restricted audit, frozen interfaces/evals,
phase zones and retry/audit bounds remain in force. Workers run context before
each step and never commit; only the unchanged pinned main controller commits.
Keep Decisions behavioral and short; process/evidence belongs in Provenance.
Do not consult old drill implementations, bootstrap or deleted code. No other
product surface, hardening campaign or speculative prerequisite is authorized.

## Provenance
- evals 2026-10-31T22:27:25.744814Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=124 failing=25 passing=99
- audit 2026-10-31T22:29:55.525964Z agent=manual fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
- evals 2026-10-31T22:50:05.029594Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=128 failing=29 passing=99
- audit 2026-10-31T22:51:33.041786Z agent=manual fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
## Audit

No findings; checked RUN-050, RUN-051, RUN-052, RUN-053, RUN-054, RUN-055, RUN-056, RUN-057, RUN-058, RUN-059, RUN-060, RUN-061, RUN-062
