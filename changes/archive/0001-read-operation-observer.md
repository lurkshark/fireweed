---
already_satisfied:
    - APP-021
bases:
    APP-021: "16461911"
module: app
phase: applied
run: dc1b2eb21e4aa3eb71be25d0
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

Pre-freeze independent review found that the first actual-verb observer missed remote add/notes add mutations formerly caught by argv-wide matching. Its running preflight was deliberately terminated and is not accepted evidence. The corrected observer explicitly denies nested mutations (with live remote/notes/bisect/submodule controls), while preserving harmless remote names and selected metadata. A real nested fixture now calibrates both root and subtree reads. A second pre-freeze review found sparse-checkout add and submodule foreach execution were still delegated; its incomplete preflight was also stopped, not accepted. The final observer delegates an explicit set of read operations, classifies config/symbolic-ref/remote/submodule read forms, and rejects remaining operations by default, with independent sparse-checkout/foreach/config/symbolic-ref mutation controls. A third review found flag-looking --file/-f operands were mistaken for config read actions; its incomplete preflight was stopped. Config option operands are now consumed explicitly and both filename forms have denial controls, alongside successful ordinary and explicit-file config reads. No product implementation or frozen eval was changed.

## Rejected alternatives

Deleting dependent checks, allowing actual commit/fetch operations, or exempting private contents would weaken the contract. Editing app evals during project implementation would violate module and phase boundaries. Adapting product metadata reads to word matching would encode an observer bug as a product requirement.

## Audit

No findings. Checked APP-021.

The modification legitimately clarifies the existing read-only source contract; already_satisfied does not exempt new product functionality. Reviewed actual operation parsing after global options, consumed config operands, Windows executable naming, and observer installation after fixture setup. Allowed controls require successful real Git execution, byte-identical observer output and no marker for selected commit headers, OID-only inventory, selected no-walk resolution, nonempty root/subtree IDs and harmless command-looking arguments. Denied controls cover mutation/network verbs, nested remote/notes/bisect/submodule/sparse-checkout operations, config and symbolic-ref writes, and flag-looking config filenames.

The specialized selected-source and exact-claim observers, private/disposable/transitive/outside-content denial controls, native mutation/network/noninteractive controls, hook/fsmonitor markers, and complete file/index/HEAD/ref/common-Git/temp/cwd/environment preservation assertions remain in force. This audit inspected only the permitted public source and evals; it did not execute tests or inspect implementation, history, bootstrap or evidence.

## Provenance

- evals | 2026-10-23T17:43:07Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 476 | failing: 0 | passing: 476
- audit | 2026-10-23T17:44:34Z | agent: manual | fireweed: 0.4.0 | playbook: a9c982d24b476621316049b527e6c7bbc319f06b12268af150fa7d8c6df92230 | tests: 0 | failing: 0 | passing: 0
- impl | 2026-10-23T17:52:40Z | agent: manual | fireweed: 0.4.0 | playbook: 653912b194316c62181f62d0a69dc213f7668481ea8b1b9926ef889dd48a9e08 | tests: 476 | failing: 0 | passing: 476
