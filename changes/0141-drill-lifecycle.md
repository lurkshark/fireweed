---
audit_pending:
    - RUN-051
    - RUN-053
    - RUN-056
    - RUN-057
    - RUN-058
    - RUN-059
    - RUN-060
    - RUN-062
audit_rounds: 1
bases: {}
module: runner
phase: audit
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
## Audit

- RUN-051: "Drill MUST create a detached scratch worktree at the invoking HEAD under the common Git directory's fireweed/worktrees directory and run the configured setup hook there with FIREWEED_BIN set to request Bin." — The setup assertion checks its path and detached state, but never checks the scratch HEAD against the invoking HEAD.
- RUN-053: "Drill MUST invoke reconstruction in a separate input workspace containing the target's retained files and only interface files from its transitive dependencies, plus SYSTEM.md and fresh independent Git metadata." — The input fixture has only one direct dependency of the target; no assertion checks an interface from a dependency beyond it.
- RUN-056: "Drill MUST use the configured worker command, wrapper and timeout through InvokeAgent, retaining ordered AgentResults in Steps." — The drill tests check the worker and wrapper, and a request timeout override, but never assert that the configured timeout bounds a drill invocation when the request timeout is unset.
- RUN-057: "Drill MUST treat usage-limit responses as the same attempt, applying configured limit_wait and limit_budget without spending a failed attempt." — The usage tests count calls and budget exhaustion, but do not assert that the configured wait occurs.
- RUN-058: "After accepted reconstruction, Drill MUST run the target module's eval hook and require passing enforced cases and coverage through checks." — The own-eval tests assert case failures but never require coverage to pass.
- RUN-059: "Every dependent MUST satisfy passing enforced cases and coverage for a passing drill, even when own or another dependent's evals fail." — The dependent-eval tests assert failed cases but never require dependent coverage to pass.
- RUN-060: "Eval entries MUST identify module, testcase, spec IDs and result" — The report tests search for tokens anywhere in the document; no assertion binds module, testcase, IDs and result to each eval entry or spec ID and target module to each gap. They also do not assert that the successful worker observation appears in the report.
- RUN-062: "Cleanup failure MUST be returned as an error rather than reported as a passing drill." — No test induces a cleanup failure or asserts its returned error and result.
