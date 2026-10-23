---
already_satisfied:
    - APP-021
bases:
    APP-021: "16461911"
module: app
---

## Motivation

Full dependent verification of the preserved project baseline-local-only fix passes all581 project records but rejects a legitimate cat-file commit <selectedBase> read: playbookCLITraps matches every argument word as a command. Independent restricted review confirms the observer overconstrains APP-021. Clarify already implemented read-only behavior and repair the observer through its owning module. APP-021 is already_satisfied because no product functionality changes; this is a legitimate clarification, not an exemption for unfinished functionality.

## Interface

No public API or product implementation edits. Correct the generic app Git observer to parse the actual operation after global options, retain every forbidden operation and non-Git execution trap, and independently calibrate real-success byte-equal selected commit/metadata and harmless command-looking path/config arguments. Retain live mutation/network controls and all file/ref/index/HEAD/private-content/hook/fsmonitor/context assertions. Existing specialized context content selection controls remain in force.

## Added

## Modified

### APP-021
context MUST accept --json before or after the command/operand and valid --help/-h without loading a project. Missing/excess operands, unknown or valued boolean flags, show-only modes and coverage-only --with-change MUST be corrective U01 exit 2 before repository/content access. -- MUST stop option parsing, so later flag-looking text is an operand and cannot enable JSON. All calls MUST remain local and noninteractive and preserve files, ignored/outside contents, index, HEAD, branches, claim/other refs, common Git, temporary storage and parent cwd/environment on success and rejection. The command MUST NOT independently run eval/setup/boundaries hooks, agents, bootstrap or network commands, commit, stage, reset, write files/refs or create claims. It MUST inspect only the exact selected claim rather than require every claim to decode, and MUST NOT open disposable code, unrelated private content, transitive dependency contracts or inputs excluded by the public source/context selection contracts. Escaping paths MUST reject before outside content opens, including ignored-error attempts. Necessary local Git object metadata reads, including the selected commit header, OID-name-only object inventory, selected no-walk commit resolution and pinned root/subtree IDs, MUST remain permitted under those source-selection restrictions. A command-looking argument to a read operation MUST NOT be treated as a mutation or network operation; restrictions apply to the actual Git operation and its effects.

## Removed

## Decisions

Do not alter the preserved project's frozen evals or production algorithms to satisfy an invalid argv-wide observer. Product implementation is already conformant; the impl phase needs no source edits. Full app evals and boundaries must still pass, and project application requires fresh own and all dependent verification after owning-module corrections. Author ordering at0001 precedes the unclaimed0041 entry without changing eligibility rules.

## Rejected alternatives

Deleting dependent checks, allowing actual commit/fetch operations, or exempting private contents would weaken the contract. Editing app evals during project implementation would violate module and phase boundaries. Adapting product metadata reads to word matching would encode an observer bug as a product requirement.
