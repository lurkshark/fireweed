---
already_satisfied:
    - APP-007
    - APP-018
    - APP-021
    - APP-022
    - APP-024
    - APP-025
    - APP-028
bases:
    APP-007: d6fccd18
    APP-018: 1de04788
    APP-021: e68934bb
    APP-022: 1d70becb
    APP-024: b9f3e8be
    APP-025: 4e3a0ed7
    APP-028: 0342b5e3
module: app
---

## Motivation

The 2026-10-26 project review found effort going into hardening beyond the design while core commands stayed unbuilt. The revised design (Spec shape; SYSTEM.md is a plain document) and the SYSTEM.md scope guardrails say that locality and safety properties are conventions checked by review, that specs state observable behavior without aliasing or purity requirements, and that the supported environment is macOS and Linux with ordinary full SHA-1 clones. See D048.

app goes first so that its frozen locality observers cannot block the dependency changes that follow (0051 to 0056).

## Interface

No public API change.

## Added

## Modified

### APP-007
modules and check MUST accept --json before or after the command, accept --help/-h without loading a project, and reject excess operands or unknown/valued flags with U01 exit 2 before loading a project. They MUST NOT run hooks, agents, or git fetch or push. Project validation MUST expose invalid graph/zone/hook/document/reservation diagnostics from the public project validator rather than silently succeeding.

### APP-018
Project-bound config errors, incompatible pins, unreadable config or a config path that is a directory MUST return C01 severity error at fireweed.yaml with exit 2, null JSON data, corrective guidance and no fallback. Git discovery failure after binding MUST also return that C01 error. Public override read/containment failures MUST return D05 severity error at .fireweed/playbooks/<name>.md with exit 2 and null data, preserving corrective resolver detail. Plain errors MUST have empty stdout and coded/path-attributed stderr; JSON MUST have only the schema-1 canonical playbook envelope and non-null diagnostics on stdout with empty stderr. Unknown names, missing/excess operands, unknown/valued flags and help MUST be handled before project/config/override access, preserving the existing -- and --json rules. Output writer failures MUST return exit 2 for embedded, override and operational-error output.

### APP-021
context MUST accept --json before or after the command/operand and valid --help/-h without loading a project. Missing/excess operands, unknown or valued boolean flags, show-only modes and coverage-only --with-change MUST be corrective U01 exit 2 before repository/content access. -- MUST stop option parsing, so later flag-looking text is an operand and cannot enable JSON. The command MUST NOT run hooks, agents, or git fetch or push, commit, or create claims.

### APP-022
start <change> MUST invoke public workflow Start from the current directory with the exact operand, including nested and linked checkouts, and expose the complete returned Context as its data without dropping or adding fields. It MUST NOT independently load current project metadata first, allocate IDs, certify or advance a step, run verification/setup/boundary hooks or agents, or commit. Successful fresh spec/fix/maintenance/regenerate and committed evals/audit/impl resume calls MUST retain public Start's saved state, claim ownership, phase/IDs, projection, context and deletion behavior. Its warning diagnostics MUST appear exactly in the public returned order and retain code/severity/message/path; warnings alone MUST return exit0 with a non-null diagnostics array.

### APP-024
start MUST accept --json before or after command/operand and valid --help/-h without opening a repository. Missing/excess operands, unknown or valued boolean flags, show-only modes and coverage-only --with-change MUST be corrective U01 exit2 before repository/content access or mutation. -- MUST end option parsing so later flag-looking text is an operand and cannot enable JSON. Invalid exact slug operands MUST retain the public C05-before-I/O behavior rather than become usage aliases or select another change.

### APP-025
verify with zero or one exact slug MUST compose public workflow.Verify from the current directory, including nested and linked checkouts, using implicit held ownership when the operand is absent. It MUST NOT independently load current project metadata, select a different claim, project again or advance the workflow. The returned data MUST retain the full public Change/Phase, ordered Coverage and every nested testcase/ID/mapping/hook exit/output/diagnostic, actual Edits and ordered Findings with ID/Clause/Text/Untestable. A nil public error with no error-severity verification diagnostics MUST return exit0, including legitimate evals fail-first, warnings alone, passing audit and valid audit findings. Any structured checking error MUST return exit1 with all available data/diagnostics intact. Approved hooks and dependent policy MUST remain authoritative despite invalid or worker-edited current metadata; evals, audit and impl MUST retain public Verify's hook selection and guard behavior.

### APP-028
verify MUST NOT run hooks other than the selected step's approved eval and boundaries hooks, and MUST NOT run agents, commit, stage, or update or release claims. Forbidden worker edits MUST prevent hooks, and post-hook forbidden edits or ownership/HEAD changes MUST be reported. FIREWEED_BIN observed by an approved hook MUST identify the actual executable rather than an inherited controller path.

## Removed

## Decisions

Narrowed specs are listed in already_satisfied because the current implementation already meets the narrower text; nothing here adds behavior. Clauses removed from a spec become unspecified, not forbidden: existing code that still does more is conforming, and the implementation step may simplify it.

One app eval replaces the process-level isolation tests: it runs help, version, modules, check, show, context and playbook against a fixture project with a PATH shim that records git, hook and agent invocations, and asserts no hook or agent ran and git was never called with fetch or push. It references APP-007 and APP-021.

## Rejected alternatives

Keeping the clauses and only deleting their evals: coverage requires every enforced spec to stay covered, and the next audit would reinstate the tests. Rewriting every over-long spec to the new shape limits at once: most long specs carry no hardening and rewriting them adds risk without benefit; they are left for later changes.

## Notes

Evals step: delete context_isolation_test.go, playbooks_isolation_test.go, start_isolation_test.go, verify_cli_isolation_test.go, verify_cli_observer_test.go, verify_cli_operations_test.go, context_git_proxy_test.go, start_git_proxy_test.go, read_operation_git_test.go and verify_cli_effects_test.go, and the cgo parts of verify_cli_metadata_test.go. Move any helper still used elsewhere into a plain test file. Remove subtests in the remaining files that assert only removed clauses (escaping-path reads, preserved index bytes, temp storage, SHA-256 fixtures, environment and cwd preservation). Add the PATH-shim test described in Decisions. Every Modified ID must keep at least one referencing test. Implementation step: no product change is required.
