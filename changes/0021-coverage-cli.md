---
module: app
phase: audit
run: 2c1fa4000b6b828fb9cb84f4
---

## Motivation

Expose real eval execution and JUnit coverage through the independent executable. Authors and workers need normal passing coverage for current or selected queued intent; lifecycle-specific fail-first remains a separate check. Use the verified public hook and evidence APIs instead of introducing another reporter or process path.

## Interface

No public Go API change. Extend Run with coverage <module> [--with-change <slug>], split/equals values and shared JSON/help/output semantics. Compose project.Load, Project/ApplyChange and checks.Evaluate through disposable app code. Coverage data uses the complete retained checks.Coverage JSON shape.

## Added

### APP-013
coverage MUST discover from a nested project directory, resolve the exact case-sensitive named module and run its resolved eval hook through public Evaluate with empty failFirst/already lists. It MUST return the complete Coverage result, including real testcase evidence/IDs/BySpec, module name, hook exit, captured Output and diagnostics, preserving deterministic public diagnostic order. No error diagnostics MUST mean exit 0; any coverage error MUST mean exit 1 while retaining Coverage data. Plain output MUST be pretty JSON of the same data on stdout with coded/path-attributed diagnostics on stderr; JSON MUST use the schema-1 canonical coverage/data/non-null diagnostics envelope with no plaintext on either stream. The envelope diagnostics MUST equal the data diagnostics. Unknown modules MUST be U01 exit 2 with exact valid choices and no hook startup. Fatal Git/configuration loading failures MUST be C01 exit 2 with null JSON data. Current coverage MUST use current specs/hooks and MUST NOT project or validate pending intent as a prerequisite, so unrelated queue errors cannot replace selected-module coverage.

### APP-014
coverage --with-change MUST resolve an exact queued slug targeting the selected module and use public Project before that change followed by ApplyChange for the selected change, exposing the inclusive numeric/slug ordered projected module's specs and resolved hook metadata to Evaluate. It MUST exclude later queue entries and archives, and MUST NOT allocate IDs or edit current state. Unknown or archived-only slugs and a module/target mismatch MUST produce corrective U01 exit 2 without running a hook. Projection/application failure MUST produce C05 exit 1 attributed to the selected change path, null data and corrective guidance without hook startup. This command MUST still require ordinary passing coverage: Added/Modified/fixes and already_satisfied metadata MUST NOT infer failFirst or grant coverage exemptions. Stale bases alone MUST NOT prevent an otherwise valid preview, consistent with public Project; static queue validation remains check's responsibility.

### APP-015
The executable MUST retain public Evaluate's C07/C11/C12/C13/C14 errors, target module paths, corrective messages, actual hook exits and output for missing hooks/reports, unreadable/malformed reports, timeouts, reporter exit inconsistencies, missing/unknown/removed IDs, skips/uncovered specs and ordinary failures. Review and reasoned quarantine suspensions MUST remain visible in testcase mappings without inventing enforcement errors. Malformed or fatally incomplete hook evidence MUST keep non-null empty Tests/BySpec and MUST NOT become successful coverage, including a timeout after a valid failing report was written. Errors MUST return exit 1 rather than a configuration/usage exit; arbitrary hook stdout/stderr MUST remain only inside Coverage.Output, without corrupting either output format.

### APP-016
coverage MUST accept --json before or after command/operands, --with-change <slug> and --with-change=<slug> anywhere before --, and valid --help/-h without loading a project or running hooks. Repeated with-change values MUST use the last value. Missing/excess operands, empty/missing with-change values, unknown/valued boolean flags, show-only mode flags and with-change on other implemented commands MUST produce corrective U01 exit 2 before loading a project. -- MUST stop option parsing, so subsequent flag-looking text is an operand and cannot enable JSON or with-change. Output write failures MUST return exit 2. Current and projected coverage MUST remain noninteractive and MUST NOT independently run setup/boundaries hooks, agents, the bootstrap or network commands, create Git commits/refs/claims, rewrite specs/changes or modify ignored/user files, except explicitly authored eval-hook actions. Owned report directories MUST be cleaned after success and failure, and FIREWEED_BIN observed by the hook MUST identify the new subject executable rather than an inherited controller path.

## Modified

## Removed

## Decisions

coverage is selected-module evidence, not ci or check. It emits Evaluate diagnostics only and does not merge unrelated project static/queue validation. Fatal discovery/configuration errors retain the established C01 contract; a readable selected module remains evaluable despite unrelated pending/static problems. --with-change uses the existing full-queue projection before the selected slug, then applies that change once, so preceding cross-module intent participates consistently with the public projection contract. A failed preview is C05 at the selected change path with null data. It never reserves IDs or writes lifecycle fields.

Normal coverage always passes nil failFirst/already to Evaluate, matching the reference command. A change's already_satisfied metadata cannot excuse an ordinary failing or missing testcase. A structurally valid stale-base preview is still available; check reports C10 separately. Values accept both split and equals syntax, with last repeated value winning. Boolean flags keep their existing no-value contract. Existing command validation stays before help and project loading. Module and slug errors list exact choices or the actual target and a usable corrective command. Hook output is a field of structured evidence, not an extra stdout/stderr stream.

## Rejected alternatives

Automatically treating --with-change as evals-phase fail-first would make a normal coverage command accept failures. Reusing the controller to execute hooks would falsify product independence. Sharing a report file or forwarding raw hook output would corrupt evidence and JSON.

## Provenance

- evals | 2026-10-22T17:45:43Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 43 | failing: 21 | passing: 22
- audit | 2026-10-22T17:49:29Z | agent: manual | fireweed: 0.4.0 | playbook: a9c982d24b476621316049b527e6c7bbc319f06b12268af150fa7d8c6df92230 | tests: 0 | failing: 0 | passing: 0
- evals | 2026-10-22T17:52:37Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 45 | failing: 22 | passing: 23
- audit | 2026-10-22T17:53:52Z | agent: manual | fireweed: 0.4.0 | playbook: a9c982d24b476621316049b527e6c7bbc319f06b12268af150fa7d8c6df92230 | tests: 0 | failing: 0 | passing: 0
- evals | 2026-10-22T17:54:46Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 45 | failing: 22 | passing: 23

## Audit

- APP-013: "Fatal Git/configuration loading failures MUST be C01 exit 2 with null JSON data." — Test_APP013FatalCoverageConfiguration calls coveragePretty on plain stdout and then requires that stdout to be empty. coveragePretty rejects empty input, making these assertions mutually exclusive; restrict the pretty-data assertion to executions that return Coverage data.
