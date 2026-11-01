# Foundation drills goal prompt

Restart prompt for future execution. This planning update does not start a goal,
change, worker or drill. SYSTEM.md governs execution; PLAN.md orders the work;
ACCEPTANCE.md defines FD1–FD4; PROGRESS.md records the pin, state and evidence.
Completed EC1–EC6 and DR1–DR4 evidence is reused where still valid.

```text
/goal Prove independent gitstore and assets drills in isolated NEW candidate clones, gitstore first and assets second, one at a time. Complete ACCEPTANCE FD1–FD4 under the unchanged main NEW controller pin and configured bounds. Stop at any blocked change, independence/access failure or gap needing the owner's decision. Full product completion remains open.

Resume:
- Read AGENTS.md, SYSTEM.md and docs/rebuild/{PROGRESS,PLAN,ACCEPTANCE,DECISIONS}.md. The unchanged design is product authority. Inspect live Git status, queue, controller processes and shared claims; preserve unrelated worktrees and finish a claim only with its owning controller.
- Verify main's controller source/hash against PROGRESS. Use /Users/joel/Workspace/fireweed/.fireweed/bin/fireweed for main workflow commands; never mutable bin/fireweed or an automatic bootstrap fallback. Keep the pin and all configured bounds unchanged. Put the current tree's .fireweed/bin first on worker PATH and use GOCACHE=/tmp/fireweed-rebuild-cache.
- Select an explicitly identified clean NEW candidate for the isolated proofs. Align clone/scratch instructions, setup and PATH; hooks copy the actual invoker. Reuse valid checks and the completed contracts proof; repeat checks only when relevant changes invalidate evidence.

Sequential proofs:
- Inspect independence before each real invocation. Supply the target's retained spec/Decisions/interfaces/evals and approved retained files, plus only declared dependency interfaces. Both targets depend only on contracts. Keep dependent implementations in the separate verifier checkout, outside the rebuilding worker's inputs.
- Exclude target disposable implementation and embedded bundle data, source history, saved attempts, product binaries, bootstrap and other modules' internals as reconstruction inputs. Do not query product executables for target content; assets recipes/playbooks/help/explanations must not be recovered from the invoking candidate.
- Workers, auditors and drill workers use exactly: codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c 'model_reasoning_effort="high"' -s workspace-write {prompt}. Stop on access failure; never substitute model or effort.
- Run gitstore within existing bounds. Require own and every transitive dependent module's evals to pass, retained inputs unchanged, reviewed native report and retained rebuilt-code ref, verified independence and complete scratch cleanup. Resolve any required owner checkpoint and commit reviewed evidence before starting assets. Run assets with the same requirements; never overlap the two drills.
- Stop on any blocked change, independence/access failure or gap/contradicted Decision needing owner agreement. Do not automatically author gap fixes, reset retries/rounds or expand to another module. Owner-approved corrections follow PLAN's normal fail-first, restricted independent audit and immutable phase rules, preserving cases, assertions and ID coverage.

Completion:
- Record concise FD1–FD4 results, candidate identities, native reports/refs and inspectable clone locations in PROGRESS. Commit reviewed native reports as protected author evidence with no active run/shared claim. Per-change evidence stays in archived Provenance; create no supplemental report inventories or hash manifests.
- Finish clean, with no live runs/shared claims, scratch removed and main's original pin verified unchanged. Preserve completed evidence and unrelated worktrees. No main implementation import, real regenerate change, controller promotion, other module drill, deferred product work or push is included.
- Complete this goal only after both module proofs and FD1–FD4 pass. A failed drill or unresolved owner checkpoint is not completion. All later full-completion gates remain required.
```
