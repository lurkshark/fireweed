# Rebuild progress

## Current status — 2026-10-27

Human review 8 approved 0032–0034; the authorized push synchronized main and
origin/main at 16cc688 and cleared P07. **Applied since review: 2.** Resume
the remaining previously identified bundled specs; PRJ-004–006 are complete.

The owner's goal update authorizes a subagent review at every five applied
changes: push to origin and continue when the review finds the work on track, and
stop for human review with details when it reports an issue. Record each review. Separate
phase audits, blocked-change stops and P07 remain enforced; this policy does not
alter bounds. The owner explicitly authorizes each push only after its
passing five-change subagent review. Milestone 1 passed;
Milestone 2 waits for the spec splits, diagnostic migration and contracts drill.

Controller: `.fireweed/bin/fireweed-bootstrap` 0.1.0, source `7f925190878e38fc38f30d16bb1cf47aca098cf7`, pin SHA-256 `f1516b5686435c23b54b71c5ac554aba6859b5454f072e0c5eb031dfd6f0fd4f`; unchanged.
Stage-2 history baseline: `be35552b706aa3000ff8f11306d129ff1d701b85`; always check an explicit nonempty baseline..HEAD range. Use `GOCACHE=/tmp/fireweed-rebuild-cache`.
NEW binary: `bin/fireweed`, clean source `ac3779b71bbdab0df016a99e74db38c4b190a659` (CGO disabled); [committed manual-lifecycle report](reports/manual-lifecycle.md). Unbuilt commands remain explicit errors; the empty runner is expected.

Workers and separate auditors use `codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c 'model_reasoning_effort="high"' -s workspace-write {prompt}`. Interactive work through 0014's initial evals was followed by coordinator `done`, hence legacy `agent=manual` lines. Configured controller execution began at 0014's audit; its provenance records the actual command/model. Continue configured execution. Ignored runtime logs are auxiliary, not reviewable evidence; archives, permanent evals and committed reports carry the evidence.

## Applied changes
- 0001 standalone-information — AST-001–003; phases a4ceb17/826b9e9/2248b8e; 8 fail-first → 8 pass.
- 0002 standalone-cli — APP-001–003; phases 1298acf/7de7d30/db77de7; 30 fail-first → 30 pass; auditor findings were discarded by folded quotes, corrected in 0003.
- 0003 json-error-assertions — APP-003 eval_fix; eb4704d/e184bb6/8796d37; 26 app cases pass, independent audit passed.
- 0004 lifecycle-primitives — CON-001–005; 11d0180/6e341a4/8e37dd7; 18 failures → 27 own cases pass, dependents pass.
- 0005 fixture-authoring — PRJ-001–003; 1bcba65/5c647a2/bfee6b8/06f17c6/91f036c; one findings round, 24 failures → 25 own cases pass.
- 0006 project-commands — APP-004–005; da5dba8/98e5852/5f09706; 12 failures → 51 app cases pass.
- 0007 change-handles — CON-006; a22fcd4/a194629/7580897; 1 failure → 34 own cases pass, dependents pass.
- 0008 queue-authoring — PRJ-004–007; 6e3ca5f/776b4d0/2ef37df; 24 failures → 53 own cases pass, dependents pass.
- 0009 authoring-commands — APP-006–008; 5720dce/b2aaa14/de65319; 27 failures → 78 app cases pass.
- 0010 manual-lifecycle-git — GIT-001–005; evals 3e82adc/93ba9a6/5d00cac, audits 98a799f/f2ddd0f/5911fab, impl 92f999a; two rounds, 7 failures → 8 own cases and 2 public probes pass.
- 0011 lifecycle-checks — CHK-001–005; 69933f7/eff6823/3c5b40b; 25 failures → 28 own cases and 4 public probes pass, dependents pass.
- 0012 manual-lifecycle — FLOW-001–005; evals c5c02de/ca11056/aec7154, audits 0b5aeef/25e4775/8627014, impl 563d2ee; two rounds, 28 failures → 32 own cases and 3 public probes pass.
- 0013 lifecycle-cli — APP-004/009–012; 678f7bb/7db8c47/a9d195c/804cb84/89b7028; one round, 8 failures → 86 app cases, held-coverage probes and real Go lifecycle pass.
- 0014 go-setup-bundle — AST-004–007; 491bffd/7789d62/d24f799/715a1be/03c2375; one round, 6 then 8 failures → 20 own cases and 199 dependent cases pass; configured-agent provenance from audit onward.
- 0015 start-state-specs — FLOW-006–009; ec03162/5223bcb/891d152; already_satisfied, 36 own cases and dependents pass; eval retargeting correction follows in 0016.
- 0016 retarget-start-evals — eval_fix FLOW-001/006–009; 129178c/0b7ea8b/5337b77; restored 32 own cases with every assertion/ID retained, dependents pass.
- 0017 start-precondition-specs — FLOW-010–013; 438072c/b35410d/6b2af96; already_satisfied, 32 own cases and dependents pass; implementation unchanged.
- 0018 start-claim-specs — FLOW-014–017; 9cfb547/31df564/9669555; already_satisfied, 32 own cases and dependents pass; implementation unchanged.
- 0019 narrow-start-claim-spec — FLOW-001/018–019; 0bc1106/dc436e7/bde237f; already_satisfied, 32 own cases and dependents pass; FLOW-001 narrowed, implementation unchanged.
- 0020 verify-evidence-specs — FLOW-020–023; evals 2d7912c/15139aa, audits b43448f/7bbf4eb, impl 42dae8e; one findings round, 32 own cases and dependents pass; implementation unchanged.
- 0021 verify-pass-specs — FLOW-024–027; 3067213/0521c85/8d786c6; 32 own cases and dependents pass, implementation unchanged; resumed saved attempt 1 after human-confirmed Sol access recovery.
- 0022 narrow-verify-zone-spec — FLOW-002/028–031; 534fef4/ea82977/17441fd; 32 own cases and dependents pass; FLOW-002 narrowed, implementation unchanged.
- 0023 evals-transition-specs — FLOW-032–035; 5ebcac7/2481621/5fbaf95; 32 own cases and dependents pass; existing tests retargeted, implementation unchanged.
- 0024 audit-state-specs — FLOW-036–040; 1b1c8f4/8aa115e/c64765f; 32 own cases and dependents pass; existing tests retargeted, implementation unchanged.
- 0025 narrow-done-verification-spec — FLOW-003/041–043; fe63ecf/44f5aa4/9042469; 32 own cases and dependents pass; FLOW-003 narrowed, implementation unchanged.
- 0026 zone-prohibition-specs — CHK-006–010; evals 36fffae/69c2898, audits a643dcb/7bfbe66, impl 7b5634c; one findings round, 28 own cases and dependents pass; implementation unchanged.
- 0027 evals-zone-permission-specs — CHK-011–015; 82bf116/09b7274/d384192; 28 own cases and dependents pass; existing eval retargeted, implementation unchanged.
- 0028 impl-zone-permission-specs — CHK-016–020; e4bb1be/1a73eb6/5c54ca1; 28 own cases and dependents pass; existing eval retargeted, implementation unchanged.
- 0029 narrow-audit-zone-spec — CHK-004/021–024; cdfb41c/e38d19b/db03f62; 28 own cases and dependents pass; CHK-004 narrowed, implementation unchanged.

[Review 5 checkpoint](reports/review5-checkpoint.md): race/vet/controller CI passed, including 261 cases, doctor recovery and history; no pending change or claim. Approved in human review 6; the next slice starts with PRJ-007.

- 0030 check-document-specs — PRJ-008–012; evals f8d1bc2/05825f0, audits 2909f64/4be3541, impl fcc385f; one findings round, 53 own cases and 199 total phase cases pass; implementation unchanged.

- 0031 narrow-check-diagnostics-spec — PRJ-007/013–016; evals 6960718/4441d55, audits 04eb482/1e91828, impl ac3779b; one findings round, 53 own cases and 199 total phase cases pass; PRJ-007 narrowed and blanket C01 removed, implementation unchanged.

[Review 6 checkpoint](reports/review6-checkpoint.md): race/vet/controller CI passed,
including 261 cases, doctor and history; no pending change, claim or scratch
worktree. Approved in human review 7, including consecutive project splits.

- 0032 narrow-load-spec — PRJ-001/017–018; 58d1946/b7d1061/f36a557; 53 own cases and 199 total phase cases pass; no findings, implementation unchanged.

- 0033 narrow-graph-spec — PRJ-002/019–022; d170452/8f084e2/5d1bb0f; 53 own cases and 199 total phase cases pass; no findings, implementation unchanged.

- 0034 narrow-new-change-spec — PRJ-003/023–026; 305d5a3/ea7cd72/2519e8a; 53 own cases and 199 total phase cases pass; no findings, implementation unchanged.

[Review 7 checkpoint](reports/review7-checkpoint.md): three applications; P07
was awaiting shared-base synchronization; human review 8 now approves the
checkpoint and authorizes pushing reviewed main. No pending
change, claim or scratch worktree. Prior race/vet/CI/doctor evidence is reused
because this slice changed only spec wording and eval labels/table metadata.

- 0035 queue-document-specs — PRJ-027–031; c22892f/ac152fa/6187c97; 53 own cases and 199 total phase cases pass; no findings, implementation unchanged.

- 0036 narrow-change-resolution-spec — PRJ-004; 49b9f0e/9bdc62d/0fc27f1; 53 own cases and 199 total phase cases pass; no findings, implementation unchanged.

- 0037 projection-rule-specs — PRJ-032–036; 3002d69/1f13327/136ca0c; 53 own cases and 199 total phase cases pass; no findings, implementation unchanged.

- 0038 narrow-projection-edits-spec — PRJ-005; evals dc8847c/a6ecd35, audits ec740ab/e705509, impl 0850274; one findings round, 53 own cases and 199 total phase cases pass; full Modified-body assertion strengthened, implementation unchanged.

- 0039 narrow-stamp-spec — PRJ-006/037–039; f33d236/7166bc5/430e56c; 53 own cases and 199 total phase cases pass; no findings, implementation unchanged.

[0035–0039 checkpoint](reports/checkpoint-0035-0039.md): five applied project
splits. Subagent checkpoint_0035_0039 returned PASS with no issues. The baseline
history check passed over 232 commits; the owner conditionally authorizes the
push and continuation. The applied count restarts after this passing review.

- 0040 narrow-coverage-spec — CHK-002/025–026; 9d54aea/ef8495d/40bfafb; 28 own cases and 146 total phase cases pass; no findings, implementation unchanged.

- 0041 narrow-junit-spec — CHK-001/027–030; fb61726/73f0bc7/dd588b3; 28 own cases and 146 total phase cases pass; no findings, implementation unchanged.

- 0042 eval-hook-data-specs — CHK-031–035; 9ae18d0/53dd0be/5c585d6; 28 own cases and 146 total phase cases pass; no findings, implementation unchanged.

- 0043 narrow-eval-hook-spec — CHK-003/036–038; be6ed03/5536992/07181b1; 28 own cases and 146 total phase cases pass; no findings, implementation unchanged.

- 0044 boundary-suppression-specs — CHK-039–043; 8492280/7cbffac/06699f9; 28 own cases and 146 total phase cases pass; no findings, implementation unchanged.

[0040–0044 checkpoint](reports/checkpoint-0040-0044.md): five applied checks
splits. Subagent checkpoint_0040_0044 returned PASS with no issues. The explicit
baseline history check passed over 258 commits, authorizing the conditional
push and continuation. The applied count restarts after this passing review.

- 0045 narrow-boundary-hook-spec — CHK-005/044–045; 8d23fd6/cb9e161/ab79c6c; 28 own cases and 146 total phase cases pass; no findings, implementation unchanged.

- 0046 application-state-specs — FLOW-044–048; 5a5df43/7ab36b4/e52698f; 32 own cases and 118 total phase cases pass; no findings, existing metadata/Decision/phase-base assertions strengthened, implementation unchanged.

## Next

1. Split permanent bundled behaviors through authored changes. Use already_satisfied and retarget existing evals; never hand-edit Specs/Decisions. FLOW-001/FLOW-002/FLOW-003/CHK-004/PRJ-001–007 are split. Human review 7 authorizes consecutive project splits for the remaining project bundles; PRJ-001–007 are now complete; other previously identified bundled specs remain. Human review 8 approves the P07 checkpoint and authorizes a push; verify the shared-base gate has cleared before resuming. Add atomic specs before narrowing an old item when one change cannot hold the full split.
2. Apply D006: remove wrong label assertions through eval_fix, migrate implementation in normal changes, and keep explain accurate. Preserve all behavior assertions and coverage.
3. Before Milestone 2 touches workflow implementation, split lifecycle/stub.go by responsibility as maintenance only. No other cleanup/hardening.
4. After items 1 and 2, run exactly one real controller drill on contracts. Commit its report and list gaps here; author no gap changes and drill no other module.
5. Milestone 2 then centers projected blocking and fake-agent next/run through NEW. No NEW unattended real-agent execution before that fake-agent run completes. Full acceptance remains unverified.
