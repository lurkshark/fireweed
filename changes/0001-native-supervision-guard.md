---
already_satisfied:
    - FLOW-040
bases:
    FLOW-040: 22cfe698
module: workflow
---

## Motivation

Complete dependent verification of the unclaimed vanished-ancestor Gitstore change exposes an execution observer that rejects every executable whose basename differs from git. Independent restricted source-only review confirms FLOW-037/FLOW-040 forbid hooks, agents, bootstrap and network effects, while explicitly permitting bounded local subprocesses; they do not require a Git basename for a native supervision helper. The current workflow implementation already meets this contract through direct bounded local Git execution. Correct the observer through its owning module without adapting production to an invalid basename test.

already_satisfied is explicit for this legitimate clarification of existing FLOW-040 behavior. This is not an exemption for unfinished cancellation functionality: Gitstore0048 retains its genuine fail-first535/Audit and must resume with complete owning/dependent verification after this repair. Its private16-file work is hash-preserved outside the checkout, the claim is released and its committed baseline is restored. The rejected complete report remains evidence.

## Interface

No public API or product implementation change. Admit only native supervision from the source executable in the existing native Start guard, alongside requested Git commands. Calibrate a real native-only self subprocess positively and reject self reentry into Go initialization before main. Retain the existing unrelated-process, hook, agent, invalid-selector-before-access, network, stdin/stdout, private-content, mutation, rollback, deletion and ownership controls. Do not encode the Gitstore helper's private argument marker or request format in the oracle.

## Added

## Modified

### FLOW-040
Start MUST remain local, deterministic in all fields except fresh run identities, and noninteractive. It MUST NOT fetch, push, create branches/worktrees, commit, move HEAD, stage/reset the index, mutate other refs/claims, invoke Git hooks or fsmonitor, execute hooks/agents/bootstrap, use the network, or change parent cwd/environment. Ordinary and linked SHA-1/SHA-256 repositories MUST retain all unrelated tracked/ignored files, file modes, index bytes, refs, common-Git bookkeeping, attempts and drill reports. Failure paths MUST preserve the same unrelated state. The operation MUST use bounded local subprocesses and retain useful failed-command details without parsing stderr as a successful Git stdout protocol. A bounded local native supervision helper MAY carry out the requested Git subprocess lifecycle, provided it preserves these same operation, state and interaction restrictions and does not enter the caller's Go initializers or main. Merely using such supervision MUST NOT be classified as invoking a hook, agent or bootstrap.

## Removed

## Decisions

This is an eval-observer repair for already conformant workflow behavior, governed through EVALS, an independent source-only assertion Audit, implementation verification and archive. The implementation phase requires no workflow source edits. All permanent cases remain enabled and positive/negative controls must prove the refined observer is live. The source-executable permission is conditional on native-only supervision: a separately calibrated Go-initializer trap must reject ordinary Go self reexecution rather than silently allow unguarded main or initializer effects. A filename or executable-content copy must not qualify as the exact source executable. All other non-Git targets stay denied.

Author ordering at0001 places this prerequisite repair before the released0048 queue entry, using the existing numeric/slug ordering without changing eligibility rules. The goal authorizes local execution and controller-owned phase/archive commits, not a claimed human merge. Keep bootstrap and NEW subjects distinct, GOCACHE=/tmp/fireweed-rebuild-cache, and no pushes/publication/tags.

## Rejected alternatives

Editing workflow evals during Gitstore implementation violates phase/module boundaries. Deleting execution enforcement, allowing arbitrary executables or Go initializer/main self reentry, or disabling dependent verification would weaken the actual contract. Renaming a product helper to git or bypassing native interposition would encode or hide the oracle bug. Marking the unfinished Gitstore fix already satisfied would be false; its original fail-first evidence and complete verification remain mandatory.
