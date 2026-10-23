---
module: app
phase: audit
run: b11651ba600996cbe270568a
---

## Motivation

Expose the already governed public manual Start transaction through the independently built executable. This remains a CLI composition slice; saved attempts and later verify/done/runner commands require their own changes. The full command inventory and acceptance gates remain unfinished.

The frozen app context Git observer currently rejects necessary metadata-only object inventory, selected no-walk commit resolution, and pinned subtree IDs. An independent public-only review confirms these are allowed by APP-021 and PRJ-024/027. The pending project baseline-local-only fix passes its frozen focused suite/race/vet, but the observer rejects it before context restoration. Correct that observer and independently calibrate allowed metadata versus forbidden private contents in this module's eval phase. Do not weaken any normative source/claim/content restriction.

This author-selected queue number puts the app slice before the unclaimed committed impl-phase project regression. Its frozen eval/audit commits and saved fresh implementation remain intact; resume that project change after the app observer correction is applied. No dependency or enforcement setting is changed.

## Interface

No public Go interface change. The existing Run facade remains thin; dispatch behavior stays disposable under internal/. Permanent executable evals must build the NEW binary separately from the controller, use only public APIs and local fixtures, and never invoke the bootstrap as subject. Correct/calibrate the existing context Git observer without editing public contracts or source-selection expectations.

## Added

### APP-022
start <change> MUST invoke public workflow Start from the current directory with the exact operand, including nested and linked checkouts, and expose the complete returned Context as its data without dropping or adding fields. It MUST NOT independently load current project metadata first, allocate IDs, certify or advance a step, run verification/setup/boundary hooks or agents, or commit. Successful fresh spec/fix/maintenance/regenerate and committed evals/audit/impl resume calls MUST retain public Start's saved state, claim ownership, phase/IDs, projection, context and deletion behavior. Its warning diagnostics MUST appear exactly in the public returned order and retain code/severity/message/path; warnings alone MUST return exit0 with a non-null diagnostics array. The CLI MUST add no file/ref/index/HEAD/temp/cwd/environment changes beyond the public Start transaction and its output.

### APP-023
start MUST use the schema-1 canonical start/data/non-null diagnostics envelope in JSON with exactly one value on stdout and empty stderr. Plain success MUST write pretty JSON of that same complete Context on stdout and coded/path-attributed warnings on stderr. An error whose public message starts with P01, P02, P03 or C05 MUST retain its full useful message as one error diagnostic with that code and exit1; leading C01 or D05 MUST retain its code with exit2. Other public errors MUST retain their full useful message as C01 error/exit2. Failures MUST have null data and no partial context, with empty plain stdout and coded stderr or the JSON envelope alone. Output writer failures MUST return exit2, including usage/error/warning output. Failure to emit an already successful Start MUST retain its valid owned claim and started files for context-based recovery; the CLI MUST NOT discard or advance that transaction.

### APP-024
start MUST accept --json before or after command/operand and valid --help/-h without opening a repository. Missing/excess operands, unknown or valued boolean flags, show-only modes and coverage-only --with-change MUST be corrective U01 exit2 before repository/content access or mutation. -- MUST end option parsing so later flag-looking text is an operand and cannot enable JSON. Invalid exact slug operands MUST retain the public C05-before-I/O behavior rather than become usage aliases or select another change. Success, rejected calls and output failures MUST remain local/noninteractive and MUST NOT independently execute bootstrap/network commands or mutate other claims or user content.

## Modified

## Removed

## Decisions

The JSON data matches the reference CLI's WorkerContext surface; the public StartResult's Claim token is not added to that wire shape. Successful transaction output errors retain recoverable ownership. Context observer calibration must distinguish object metadata from blob contents, exact selected-claim inspection from inventories of claims, and pinned selected trees from mutable/unrelated sources. Keep live negative controls and add real allowed controls for OID-only inventory, selected no-walk resolution and selected subtree metadata.

## Rejected alternatives

Duplicating lifecycle logic in the app would diverge from the public guarded transaction. Adapting production reads to an invalid closed command allowlist would encode fixture assumptions as product restrictions. Editing app evals during a project worker step, disabling dependent checks, or claiming a passing project archive before dependent checks is not permitted.

## Provenance

- evals | 2026-10-23T16:01:48Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 475 | failing: 60 | passing: 415
