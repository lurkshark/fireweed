# Rebuild history

Historical evidence moved out of the current handoff on 2026-10-28. These
entries describe results and next-work statements at their recorded revisions;
they are not current instructions, permissions or readiness claims. Resume from
[PROGRESS.md](PROGRESS.md). Preserve evidence and checkpoint history; do not
repeat passed work because an older entry calls it pending.

## Applied changes through 0102

[0079–0083 checkpoint](reports/checkpoint-0079-0083.md) PASS: separate subagent
review and fresh build/vet/race/static/history at clean 6ce22e9; 38 new history
commits have zero diagnostics. Actual NEW Next and bounded Run fixtures pass.
The standing policy authorizes push and continuation; final gates remain open.

[Local 0084–0088 checkpoint](reports/checkpoint-0084-0088.md) PASS: separate
Sol/high review and build/vet/race/static/history at clean 940ef0e; 37 history
commits have zero diagnostics. The proposed push was rejected by automatic
approval review under the latest explicit owner stop; nothing was pushed and
further work awaits an owner decision.

[0094–0098 checkpoint](reports/checkpoint-0094-0098.md) PASS: separate Sol/high quality/goal review plus clean9949dfe build/vet/race/static/history;33 new history commits have zero diagnostics. Actual NEW usage/time stops, resume and maintenance proofs pass; fixtures removed. Standing owner authorization covers push and continuation; final gates remain open.

- 0102 regenerate-warning-explanation — AST-032; dc6da04/59499c3/31c28d1. One failure/576 passes →577 full passes; isolated audit PASS,60 own passes. Actual NEW Go regenerate start/owned-next/P10 proof passes with deletion/context/warning, one impl agent/commit, retained contracts and clean release. Separate ordinary-fixes next passes5.33s with real fail-first/pass and three commits; both fixtures removed. [Report](reports/regenerate-fixes-cli-0102.md).
- 0101 regenerate-start — FLOW-080–083; 73fa079/24ec135/006fda8. Six failure events/362 passes →368 full passes; isolated four-ID audit PASS. Untested late nested guard corrected after scoped public ancestor-target failure; seven focused passes and fixed repro, then final native pass. Initial native impl check interrupted before archive; no phase exception. Library dependency only, P10 explanation and actual Go CLI proof next. [Report](reports/regenerate-api-0101.md).
- 0100 eval-fix-playbooks — AST-029/030/031; b2841bf/6361924/f418f88. Three failures/566 passes →569 full passes within15m; isolated three-ID audit PASS,59 own passes. Actual NEW evals/audit/impl playbook JSON commands pass with conditional rules, schema1/string data, no diagnostics/stderr. Reuses0099 lifecycle proof. [Report](reports/eval-fix-playbooks-0100.md).
- 0099 eval-fix-lifecycle — FLOW-078/079, modified FLOW-036/023 (023 already_satisfied); 35ce303/3e30063d/23ad9de. Six failure events/355 passes →361 full passes; isolated four-ID audit PASS,89 own passes. Actual NEW Go eval-fix proof passes4.74s: wrong baseline test fails, correction passes under evals/audit/no-op impl, three commits, contracts/implementation unchanged, clean/unclaimed; fixture removed. [Report](reports/eval-fix-cli-0099.md).
- 0098 maintenance-start — FLOW-077; 6548562/c60e32d/d3272a3. One failure/352 passes →353 full native passes; isolated audit PASS. Corrected full own81 passes after private no-op SPEC write guard and eval_fix presence correction. Actual NEW Go maintenance next passes3.26s with only one impl agent/one archive commit, retained contracts, clean/unclaimed; fixture removed. Fifth-change checkpoint now required. [Report](reports/maintenance-cli-0098.md).
- 0097 run-for-command — APP-089/090, modified APP-084; 79d6ffd/32ede66/0184b69/81cc839/ab7b317. Both native eval rounds10 failures/186 passes →196 impl passes; scoped APP-089 audit correction/re-audit PASS. Actual NEW time-stop/resume proof PASS: exit3 in2.09s, original claim/run/phase/ID/work retained without commit; next archives same run in5.84s, four calls/three commits, later queue unchanged, clean/unclaimed. Fixture removed. [Report](reports/run-time-cli-0097.md).
- 0096 run-time-limit — RUN-036–040 and qualified RUN-016/030; f83f4ac/a847f9f/2f6870f. Complete native262 cases: five failures/257 passes →262 passes; isolated seven-ID audit PASS,76 own passes in194.04s. Initial pre-freeze240.503s timeout was incomplete and superseded after fixture simplification/direct identity assertions. Library dependency only; app --for and actual Go proof next. [Report](reports/run-time-runner-0096.md).
- 0095 usage-stop-command — APP-076/077/079/083/084/086; 6978e2a/51924d9/6dfe616/e546a01/dbda392. Seven genuine failures/179 passes →186 full passes after one APP-076 identity assertion correction and scoped re-audit PASS. Implementation wording correction passed focused config then full native checks. Actual NEW Go shared-budget stop/resume proof PASS: exit3 after Alpha applies/Beta holds, then same-run Beta archive, eight calls/six phase commits, Gamma unchanged, clean/unclaimed; fixture removed. [Report](reports/usage-stop-cli-0095.md).
- 0094 usage-limit-control — RUN-029–035, qualified RUN-015; c9a5729/03f9476/5ff9b3d/b3cbdc2/a720cf6. 15 failures/231 passes →16 failures/231 passes after one scoped RUN-033 approved-policy assertion →247 full passes; isolated re-audit PASS. Fresh implementation corrected one attempt-counting error and passed69 own cases. Runner dependency only; app exit3 and actual NEW proof immediately next. [Report](reports/usage-limit-runner-0094.md).
- 0093 applied-result-count — eval_fix RUN-016; b8a9260/bd4d862/5b0e642. Both full phases passed 231 cases; 53 own passes and isolated RUN-016 audit PASS. Exact-count helper preserves the ordinary three-step wrapper and every existing assertion/caller. Earlier usage change natively abandoned at b08c4d0 under explicit owner approval after a frozen three-versus-four-step conflict; fresh re-authoring remains pending. [Report](reports/applied-result-count-0093.md).
- 0092 agent-commit-recovery — RUN-027/028; 0ea029d/8bf710c/96211bd. Two genuine failures/229 passes → 231 full phase passes; isolated audit PASS. Actual NEW Go next proof PASS in 6.12s: four agent commits undone, work retained, three checked phase commits and P05 warnings, clean/unclaimed; fixture removed. [Report](reports/agent-commits-0092.md).
- 0091 saved-attempt-command — APP-087/088 and already-satisfied APP-009 qualification; 79a5849/7d51927/523f908. Ten genuine failures/168 passes → 178 full phase passes; isolated audit PASS. Actual NEW saved-attempt resume/correction/archive proof PASS in 2.86s, retaining run/phase/IDs/HEAD/ref and later independent work; fixture removed. Harness cwd repair occurred before any NEW call, preserving its existing unblock commit. [Report](reports/saved-resume-cli-0091.md).
- 0090 saved-attempt-resume — FLOW-073–076; 8cbcac8/d6cb397/9affb4c/2d2c483/da16b68. Eight genuine failures/287 passes → ten failures/287 passes after scoped round-1 correction → 297 full phase passes. FLOW-074/075 findings corrected and isolated re-audit passed; same run/IDs, no exception. StartFromAttempt dependency ready; app flag/actual NEW proof immediately next. [Report](reports/saved-resume-api-0090.md).
- 0089 independent-continuation — RUN-010/011/016/025/026; 423a9f7/5edbc1d/bf27309. Five genuine failures/212 passes → 217 full phase passes; isolated audit PASS, interfaces frozen. Actual NEW Go run proof PASS in 14.46s: Alpha saved/blocked, Beta waiting, Gamma applied; exit 1/E01/ref, seven calls/six commits, clean/unclaimed. Fixture retained for immediate saved-attempt resume proof. [Report](reports/independent-run-0089.md).
- 0088 projected-blocking-eligibility — FLOW-011/068–072; 49fecf4/d5e477d/199f373. Fresh implementation from retained contracts only; 13 genuine failures/270 passes → 283 full phase passes, isolated audit no findings. Canonical projected selection and Start dependency ready; runner/CLI continuation still pending after fifth checkpoint. [Report](reports/projected-eligibility-0088.md).
- 0087 independent-blocked-fixture — eval_fix RUN-011; f898bd7/2c44359/931c17f. Both full phases passed 213 cases; 47 own pass and isolated audit no findings. Fixture inputs only; all assertions/IDs preserved. Earlier 0087-projected-blocking-eligibility was natively abandoned at 3d65370 after dependent-fixture incompatibility, with only contracts retained for fresh reuse. Owner requested stop after this archive; [correction](reports/independent-blocked-fixture-0087.md), [progress review](reports/progress-review-after-0087.md).
- 0086 exhausted-verification-blocking — RUN-023/024, narrowed already-satisfied RUN-015; e0b5c8f/3af876d/c797264. Two genuine failures/211 passes → 213 full phase passes; isolated audit no findings. Existing assertions retained with changed terminal-state expectations; ten-line caller handoff. Actual NEW Go next exhaustion proof passes in 5.61s (expected exit 1), saving final work/ref and restoring state-only clean release; [report](reports/exhausted-next-0086.md).
- 0085 saved-attempt-blocking — FLOW-063–067; initial evals 5b9076e, valid findings 188b41f, corrected evals a7134cf, passing audit 7875cce, impl 526f90e. Two genuine failures/266 passes → 268 full phase passes; 54 own pass. Native round 1/same run/IDs retained through two explicitly approved author revert recoveries; ignored mechanical passes remain invalid. Scoped FLOW-066 correction/re-audit passed. Public lifecycle dependency ready; runner/CLI exhaustion proof follows. [Report](reports/saved-attempt-blocking-0085.md).
- 0084 verification-retries — RUN-015/021/022; ffa53d0/59d1a85/667e419. Five runtime failures/209 passes → 214 full phase passes; valid isolated audit no findings. Both capacity-stopped workers completed their separately approved single Sol/high retries. Initial mislocated audit is invalid/retained; fresh isolated audit accepted. Actual NEW next retry Go proof passes in 6.53s with retained work/feedback and claim cleanup; [report](reports/verification-retries-0084.md).
- 0083 audit-passing-example — AST-028; initial evals/findings 2033c79/ee2d097, corrected phases 3aa5475/da2335d/6ce22e9. One runtime failure/460 passes → 461 full phase passes. Native findings round 1 retained through author clarification; isolated AST-028-only re-audit passed. Adds the accepted two-line example only; [report](reports/audit-example-0083.md). Five-application checkpoint review follows.
- 0082 run-command — APP-080–086; 0e17dae/8220c66/58351b2. Seven runtime failures/159 passes → 166 full phase passes; isolated Sol/high audit no findings. The diagnostic fixture was corrected/validated before freezing. Actual NEW Go queue --max/continuation/empty smoke passes with nine phase commits and claim cleanup; [report](reports/bounded-run-cli.md).
- 0081 bounded-queue-execution — RUN-016–020; 4c8c330/b788753/6bdfb43. Seven runtime failures/197 passes → 204 full phase passes; isolated Sol/high audit no findings. Small independent Go queue, Max/continuation, Empty/Limited, request forwarding and held-error progress are covered; app wiring and actual NEW run proof follow.
- 0080 next-command — APP-074–079; 5e6d0bb/fff9dba/a11d713. Owner-approved native reauthoring corrected one fixture exit code; public Next preflight confirms C10/ALP-002. Six runtime failures/153 passes → 159 full phase passes. The authorized single Sol/high audit retry passed after a capacity-only stop. Fresh NEW next Go smoke passes in 5.73s with checked phases/archive/provenance/claim cleanup; [report](reports/next-0080-block.md).
- 0079 single-change-execution — RUN-010–015; 7ed35d4/54f7f5c/1ddea1c. Reauthored under explicit owner approval after native abandon e5059ae; corrected only the unsupported final commit-subject assertion. 12 genuine runtime failures/179 passes → 191 full phase passes; isolated audit no findings. Runner library works; CLI wiring follows. Prior block/attempt retained in the [recovery record](reports/next-0079-block.md).
- 0078 stable-descendant-timeout-eval — eval_fix RUN-007/009; 88119c5/7b8a843/0f6f7b2. One existing case gets bounded startup/observation windows and output-before-readiness; 178 eval/impl phase passes, focused race pass and isolated audit no findings. No product/interface/spec change.
- 0077 baseline-stop-assertion — eval_fix CHK-055; 064d9ab/0fd7393/435b176. Added stop-before-canary assertions in existing cases; 298 eval/impl phase passes, isolated audit no findings, no product/interface/spec change. Coordinator provenance is manual; real fixed-profile CLI logs retained. B02: 16 CHK-055 cases is informational; this change adds no cases.
- 0076 doctor-boundary-cleanup-fixes — CHK-051/055/056/057; e2b5b3b/37fc2e9/995eebb; three runtime failures → 298 phase passes. Native audit input deviation and supplemental CHK-055 assertion gap remain open in the correction report; no correction push.
- 0075 bounded-agent-invocation — RUN-001–009; evals 53d9279/2e40588, audits 6fdcb9f/6f898ea, impl 4b2fe8e. One scoped correction; 25 runtime failures → 25 own/178 phase passes. Retained API/model types and execution primitive are implemented; next/run queue wiring remains pending.
- 0074 doctor-setup-commands — APP-062–073; evals 5620604/eb05e41, audits a02d10d/3769100, impl 19d2b31. One scoped correction; six runtime failures → 153 app passes. Fresh NEW Go/doctor/guidance smoke passes; see the linked report.
- 0073 doctor-diagnostic-explanations — AST-026/027; b364af7/53cae57/8d27f34. Six runtime failures → 55 own/389 phase passes; no findings, restricted auditor input reads confirmed. Explains emitted C14–C16 and existing init P01 causes, retaining prior causes/assertions.
- 0072 doctor-helper-portability — eval_fix CHK-050/051; dc2b1e7/426e259/762052f. Five-line helper-only correction retains all assertions and 71 own cases; 260 phase passes with ordinary Apple Git, no findings, no implementation edit.
- 0071 doctor-probes — CHK-049–069; final evals ebf127f (34 runtime failures/226 passes, 260 total), restricted audit 8cafea6, impl e50061d (260 passes). Full recovery and earlier phases are in the linked block report; native application/cleanup passed under Git 2.53.0. CLI wiring remains pending.
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

- 0047 narrow-application-spec — FLOW-004/049–050; 80b6992/6598b5a/2a087d7; 32 own cases and 118 total phase cases pass; no findings, archive/SPEC commit assertion strengthened, implementation unchanged.

- 0048 narrow-context-spec — FLOW-005/051–053; evals c14ce8a/a146cbc, audits 63af469/a7969b6, impl 95a404a; one findings round, 32 own cases and 118 total phase cases pass; per-phase provenance assertions strengthened, implementation unchanged.

- 0049 split-contract-primitives — CON-001/002/005/006/007–015; 208fd0a/e2809a5/71ac66e; 13 affected items, 34 own cases and 261 total phase cases pass; no findings, assertions/test data and implementation unchanged.

[0045–0049 checkpoint](reports/checkpoint-0045-0049.md): five applied changes.
Subagent checkpoint_0045_0049 returned PASS with no issues. The explicit
baseline history check passed over 286 commits. The owner conditionally
authorizes the push and continuation; the applied count restarts after review.

- 0050 split-git-primitive-apis — GIT-003/004/006–010; d0179bd/bf6acb6/3519faa; seven affected items, eight own cases and 207 total phase cases pass; no findings, assertions and implementation unchanged.

- 0051 split-asset-contracts — AST-001/002/004–023; c9210cf/0d856fc/40cceb7; 22 affected items, 20 own cases and 219 total phase cases pass; no findings, assertions/fixtures and implementation unchanged.

- 0052 split-authoring-cli-contracts — APP-003–008/013–026; evals ff54a43/4ad8388, audits 56a829c/bcc1400, impl ed94c03; 20 affected items, one findings round, 86 cases pass; existing zone-array assertion strengthened, implementation/diagnostics unchanged.

- 0053 split-lifecycle-cli-contracts — APP-009–012/027–036; evals 863b578/28baed4, audits 62d07a5/ca18597, impl 11cb39a; 14 affected items, one findings round, 86 cases pass; existing omitted-Done/blocked-Verify assertions completed, implementation unchanged.

- 0054 repair-validation-code-assertions — eval_fix PRJ-007; aa756c5/7aa96d6/706ce7e; 53 own cases and 199 total phase cases pass; no findings, stale blanket C01 comparisons replaced with nonempty-code checks while retaining reason assertions; implementation unchanged.

[0050–0054 checkpoint](reports/checkpoint-0050-0054.md): five applied changes.
Subagent checkpoint_0050_0054 returned PASS with no issues. The explicit baseline
history check passed over 316 commits, authorizing push and continuation. The
applied count restarts after this passing review.

- 0055 validation-cause-codes — PRJ-008–012; 0b91a02/8f875e4/405db73; nine reported failures → 54 own/200 total phase cases pass, no findings; existing C04/C05/C06/C07 paths migrated. One existing table gained a separate reporting parent; no scenarios/fixtures added.

- 0056 validation-code-explanations — AST-003; d12c3a2/f8e16b6/c7b4bea; five reported failures → 24 own/224 total phase cases pass, no findings; four emitted-code explanations added, existing explanation and unknown-code checks retained.

- 0057 edit-reference-codes — PRJ-013–016; 7f099ed/80a1760/12cd5bc; seven reported failures → 54 own/200 total phase cases pass, no findings; existing duplicate edit/handle/base/absent-target paths now C05/C09/C08, fixtures and earlier assertions retained.

- 0058 boundary-cause-codes — CHK-039/040/043–045; 69ae618/700c317/0bfd581; three reported failures → 28 own/146 total phase cases pass, no findings; existing violations/hook failures/opt-out/unnecessary warnings now X01–X04, B05 retired here.

- 0059 reference-boundary-explanations — AST-003; 9d109e6/5ec2a2b/c7eabf9; eight reported failures → 30 own/230 total phase cases pass, no findings; C08/C09/X01–X04 explanations added and C05 repeated-edit repair clarified.

[0055–0059 checkpoint](reports/checkpoint-0055-0059.md): five applied changes.
Subagent checkpoint_0055_0059 returned PASS with no issues. Baseline history
passed over 341 commits; build/race/vet and NEW explanation JSON passed for
source c7eabf9. The passing review authorizes push and continuation; the applied
count restarts. No queue or claim remains. Diagnostic migration is incomplete.

- 0060 usage-lookup-codes — APP-003/013/014; ab0f0d6/9b95b13/df1c10d; 18 reported failures → 94 app cases pass, no findings; existing usage/unsupported rejections now U01 and information lookups U02; eight arity table cases added, all prior behavior assertions retained.

- 0061 usage-code-explanations — AST-003; 3e801bf/dd38afc/875aff7; three reported failures → 32 own/240 total phase cases pass, no findings; U01/U02 explanations added with existing causes and unknown-code checks retained.

- 0062 static-diagnostic-causes — PRJ-040–045 plus fixes PRJ-013; 7e2b3f6/d16c471/1e478b8; 17 reported failures → 74 own/228 total phase cases pass, no findings; existing Check graph/syntax/metadata/queue/target paths classified, C01 retained and duplicate-section C05 corrected; 20 focused public cases added with existing fixture helpers.

- 0063 eval-diagnostic-causes — CHK-002/025–028/036–038/046–048; 4db7bf5/4b1e611/f6b9921; 17 reported failures → 37 own/163 total phase cases pass, no findings; existing eval/coverage/result errors now carry C10/C11/E01/P2-01, with static and lookup causes preserved; signatures unchanged.

- 0064 lifecycle-diagnostic-causes — FLOW-009/015/054–062; evals 6479da3/22f0543, audits 23444aa/8f50b18, impl 4066479; one findings round, 12 reported failures → 45 own/139 total phase cases pass; P04-warning/configured-P07/Done-cause assertions completed and scoped re-audit passed; existing lifecycle causes prefixed, behavior/interfaces retained.

[0060–0064 checkpoint](reports/checkpoint-0060-0064.md): five applied changes.
Subagent checkpoint_0060_0064 returned PASS with no issues. Baseline history
passed over 369 commits; current-source build/race/vet and NEW U01/U02 JSON
checks passed. The owner-authorized push and continuation follow this passing
review; the applied count restarts. No queue or claim remains; D006 is incomplete.

- 0065 remaining-diagnostic-causes — APP-018/019/037–040; 2157abf/7697050/7eb9898; 30 reported failures → 121 app cases pass, no findings; legacy CLI labels removed. Post-application NEW smoke exposed show returning U02 for duplicate graph names and duplicate module spec IDs; correct those static causes next while retaining legitimate projected-handle ambiguity.

- 0066 preserve-show-static-causes — APP-038 plus fixes APP-037; afdc02d/edbac55/39ecdb8; three reported failures → 124 app cases pass, no findings; duplicate graph names/spec IDs retain C02/C05 and valid projected local-handle ambiguity retains U02; rebuilt NEW confirms both observed regressions fixed.

- 0067 emitted-cause-explanations — AST-003/024–026; 1360324/5d3e724/72a2860; 22 reported failures → 51 own/331 total phase cases pass, no findings; sixteen explanations added and C07/U02 clarified. All 30 emitted codes resolve through NEW; race/vet pass. Text review found C02 identity causes and P02 missing/unreadable-claim recovery underdescribed; correct those before the drill.

- 0068 complete-graph-claim-explanations — fixes AST-024/026; d6fb6b1/54892c9/79e57f4; five reported failures → 51 own/331 total phase cases pass, no findings; C02 covers module identities and dependencies, P02 covers missing/invalid/foreign claims and appropriate recovery, P01 includes existing authoring preconditions. NEW C02/P02 JSON smoke passes.

- 0069 project-initialization — APP-041–061; evals b3cbc19/734e476, audits 6fe48c5/07fcb6c, impl cdc64ae; one findings round, 17 reported runtime failures → 144 app cases pass; version-range membership and recipe default-protected assertions completed, scoped re-audit passed. Go/none/local init, optional app, managed instructions, preservation, usage/errors and JSON are wired through NEW; its Go init/modules/check fixture passes. Doctor and execution remain pending.

- 0070 recipe-error-explanations — fixes AST-003/024; 39f82e1/95b998c/4b36a84; four reported runtime failures → 51 own/351 total phase cases pass, no findings; seven assertions added to existing cases and only C12/U02 explanation strings extended. Existing causes/guidance, assertions, public interfaces and init behavior are retained. NEW C12/U02 JSON smoke and owner-required checkpoint re-review pass.

[0065–0069 checkpoint](reports/checkpoint-0065-0069.md): five applications since
the passing 0060–0064 review. Subagent checkpoint_0065_0069 initially returned
BLOCK for missing init recipe causes/repairs in explain C12/U02. The owner
approved 0070 and conditional push/continuation; its scoped re-review now returns
PASS. Build/race/vet and NEW Go init/modules/check passed at cdc64ae; the text-only
correction's full phase checks and NEW C12/U02 JSON pass at 4b36a84. History
validation passed over 397 baseline commits plus the seven-commit correction
extension, both with no diagnostics. The count resets after this passing review.
The owner's intervening direct push
approval synchronized committed history through 3897b1a; it did not reset the
five-application review count.

The planned D006 existing-path slice is complete through 0068: no legacy CLI
labels remain, and all 30 emitted codes have NEW JSON explanations. Full phase
checks passed; race/vet at 72a2860 remains valid for 0068's text-only correction.
No new feature, public interface or speculative error was added. Full diagnostic
surface for unimplemented commands is not claimed. The single contracts drill is now recorded below; no gap changes, repeat drill
or other module drills are authorized.

[Contracts drill review](reports/contracts-drill.md): native pass at 7afe9a0,
34 own/339 dependent cases; retained ref and scratch cleanup verified. One
CON-011/012 error-prefix gap is recorded. Agent reads of dependent internal
implementations exceed the retained-input boundary, so strict independence
acceptance remains unverified. No gap fix/rerun is authorized or performed.


## Historical acceptance evidence

The following is the prior ACCEPTANCE evidence log, preserved without changing
its recorded results. Current status belongs in PROGRESS; final obligations
remain in ACCEPTANCE. Historical pending-work statements are superseded by
later changes and reports.

### Partial requirement evidence at recorded revisions

The entries below record results and outstanding work at their cited revisions. Later entries may resolve earlier gaps; use PROGRESS.md and linked current reports for status. Reuse valid evidence without treating an incomplete row as a requirement to redo its passing portions.

- R17/R20: AST-001–AST-003 cover embedded version/help, initial diagnostic lookups
  and five playbooks through public assets evals `Test_AST_001_*`,
  `Test_AST_002_*`, `Test_AST_003_Explanations`. Change
  `changes/archive/0001-standalone-information.md`; phase commits
  `a4ceb17` / `826b9e9` / `2248b8e`; 8 assertion failures before and 8 passes after.
  CLI wiring, full diagnostics/versioning/playbooks/AGENTS remain unverified.

- R17: APP-001–APP-003 cover standalone NEW-binary invocation outside a project
  and with invalid config, schema-1 information JSON, usage errors and explicit
  unsupported-command failures; 30 assertion failures before, 30 passes after.
  Archive `0002-standalone-cli.md`; evals `1298acf`, controller audit `7de7d30`.
  Independent audit found APP-003 JSON gaps discarded by the controller because
  quotes folded line breaks. Correction through eval_fix remains pending.

- R17: `json-error-assertions` eval_fix repairs the APP-003 JSON gaps; phase SHAs
  `eb4704d` / `e184bb6` / `8796d37`. Separate restricted Sol/high audit passed
  APP-003; all 26 app cases pass, with unchanged implementation. This is rebuild
  process evidence, not evidence that the NEW product implements eval_fix.

- R02/R03/R04 partial: CON-001–CON-005 provide document/section/item parsing,
  canonical hashes, ID matching and zone precedence/globs. Archive
  `0004-lifecycle-primitives.md`; evals `11d0180` (18 failures, 9 passes; all changed IDs fail), audit
  `6e341a4` (independent no findings), impl `8e37dd7` (27 passes, existing app/assets
  and direct dependents/boundaries pass). Full grammar, project validation and
  lifecycle remain unverified.

- R02/R04 partial: PRJ-001–PRJ-003 load the minimal fixture, validate module graph
  identities/dependencies and author queued skeletons without assigned IDs. Archive
  `0005-fixture-authoring.md`; initial evals `1bcba65`, audit findings `5c647a2`,
  corrected evals `bfee6b8` (24 failures, 1 pass), passing independent re-audit
  `06f17c6`, impl `91f036c` (25 own and 26 app cases pass). One audit round;
  projection, stamping, other checks and CLI wiring remain unverified.

### First human checkpoint — five applied changes

Current-source Go build/test/race/vet and controller check/history/CI passed at
`91f036cbb33981a5fb72798da3edec9bdc90b077`; exact range starts at
`be35552b706aa3000ff8f11306d129ff1d701b85`. CI included doctor and all 86 reported
cases (including parent cases), with no diagnostics. See PROGRESS for logs and
limits. This is not fresh-checkout evidence or completion of any full acceptance
row. The manual NEW lifecycle, remaining surface, real drills/self-hosting and
delivery gates remain incomplete. Human review is required before change six.

- R04/R17 partial: APP-004–APP-005 wire modules and new-change through the NEW
  executable, including ancestor config discovery, graph/zone JSON and author
  error semantics. Archive `0006-project-commands.md`; phases `da5dba8` /
  `98e5852` / `5f09706`; 12 failures/39 passes before, all 51 app cases pass after.
  NEW smoke created changes/0001-answer.md in the manual slice fixture; stamping
  and lifecycle remain unverified.

- R04 partial: CON-006 extends ParseItems to the design-required new: handles
  without a parallel parser. Archive `0007-change-handles.md`; phases `a22fcd4` /
  `a194629` / `7580897`; one genuine failing case, then all 34 contracts cases pass
  with dependents/boundaries. Allocation, cross-change scope and rewriting remain
  for project/workflow evidence.

- R04/R05 partial: PRJ-004–PRJ-007 cover ordered queue parsing, projected edits/set
  metadata, per-change handle ownership, preceding-projection stamp bases and
  initial static diagnostics. Archive `0008-queue-authoring.md`; phases `6e3ca5f` /
  `776b4d0` / `2ef37df`; 24 failures/25 passes before, all 53 project cases pass after
  plus 51 app cases and boundaries. Full validation/views/lifecycle remain open.

- R04/R05/R17 partial: APP-006–APP-008 expose stamp/check and current/projected
  show through the NEW binary. Archive `0009-authoring-commands.md`; phases
  `5720dce` / `b2aaa14` / `de65319`; 27 failures/51 passes before, all 78 app cases
  pass after. Manual fixture stamp/check/projected and selected new:answer show
  exited 0 with no assigned IDs. Product lifecycle remains unbuilt/unverified.

### Shared-base approval gate after review 1

At clean product source `de65319d646fd0e39e9810415a181dd28911762c`, build/test/race/
vet, controller check and CI on explicit setup baseline `be35552...` passed. CI
includes doctor canaries/recovery, boundaries, history and all 173 reported cases,
with no diagnostics. This is current-source intermediate evidence, not final
acceptance or a fresh-checkout/drill/self-hosting claim. Controller unapproved
count is 5 against origin/main `6e341a4`; P07 requires human shared-base approval
before another change. Four changes were applied since the explicit review at
`3311f67`; no bound/ref was altered.

- R06/R13/R14 partial: GIT-001–GIT-005 supply canonical primary/linked repository
  metadata, local base selection, working/index/committed diffs, selected commits,
  blobs/CAS and bounded process groups. Archive `0010-manual-lifecycle-git.md`;
  evals `3e82adc` / `93ba9a6` / `5d00cac`, audits `98a799f` / `f2ddd0f` / `5911fab`,
  impl `92f999a`; 7 fail-first cases, then 8 permanent passes plus 2 temporary
  public smoke probes and dependent/boundary passes. Escaped process groups remain
  unsupported. Claims/lifecycle and full Git surface remain unverified.

- R07/R09/R10/R11 partial: CHK-001–CHK-005 cover public JUnit/results, hooks, zones and boundaries; archive `0011-lifecycle-checks.md`, phases `69933f7` / `eff6823` / `3c5b40b`. 25 failures before, 28 permanent cases and 4 public smoke probes pass after. Full dependents/boundaries pass. Product lifecycle and doctor integration remain unverified.

- R05–R09/R14/R17 partial: FLOW-001–FLOW-005 provide manual claims/start/verify/done/audit/apply/archive/provenance/context through public APIs. Archive `0012-manual-lifecycle.md`, evals `c5c02de` / `ca11056` / `aec7154`, findings `0b5aeef` / `25e4775`, final audit `8627014`, impl `563d2ee`: 28 initial failures, 32 final permanent cases and 3 public smoke probes pass; full dependents/boundaries pass. Sol reset/retry recorded in ignored evidence. NEW CLI lifecycle, remaining kinds/recovery/history and full acceptance remain unverified.

- Milestone 1 / R05–R11/R17 partial: APP-004, APP-009–APP-012 connect the manual
  lifecycle through NEW source `89b7028`; original phases `678f7bb` / `7db8c47` /
  `a9d195c` / `804cb84` / `89b7028`. Eight genuine initial failures, 86 final app
  cases and 2 public coverage probes pass. Real Go fixture TOY-001 at
  `/tmp/fireweed-stage2-manual-slice/project` completed native evals `f04e9f7`,
  separate audit `bfb84d0`, apply `1323a31`, with Z01 and P2-01 rejections,
  verify/no commit, frozen eval/facade, passing coverage, inspected provenance,
  clean tree and released claim. All workers/auditor were real Sol/high CLI calls.
  This verifies the manual milestone, not init/doctor/run or full acceptance.

- R03/R04/R05/R17 refinement: PRJ-007's Check contract is now split into
  PRJ-007–016 by archives 0030/0031. Existing validation assertions are retained
  and linked to their atomic IDs; cause-specific diagnostic migration remains
  pending. See [review 6 checkpoint](reports/review6-checkpoint.md) and PROGRESS
  for phase commits, audit corrections and verification evidence. This does
  not establish the remaining static-validation or CLI requirements.

- R02/R04 refinement: archives 0032–0034 split PRJ-001–003 into focused
  discovery, graph and authoring contracts, with PRJ-017–026 retaining their
  separated rules. Existing assertions and cases are unchanged. See
  [review 7 checkpoint](reports/review7-checkpoint.md) and PROGRESS for evidence
  and the P07 shared-base gate; PRJ-004–006 and full validation remain open.

- R04/R05 refinement: archives 0035–0039 complete the PRJ-004–006 split, with
  PRJ-027–039 retaining queue, projection and Stamp rules. Existing public
  assertions remain; 0038 strengthens complete-body checks after its audit.
  See [the passing checkpoint review](reports/checkpoint-0035-0039.md) and
  PROGRESS for phase evidence. Remaining module splits and D006 migration
  still precede Milestone 2; full acceptance remains unverified.

- R07/R10/R11 refinement: 0040–0044 split CHK-001–003 and retain CHK-005's
  suppression rules under CHK-039–043 before its final narrowing. Assertions,
  fixtures and implementation remain unchanged. See the
  [passing checkpoint review](reports/checkpoint-0040-0044.md) and PROGRESS
  for phase evidence; remaining splits and diagnostic migration are pending.

- R02/R03/R05/R06/R09/R11/R14/R27 refinement: 0045–0049 finish CHK-005,
  FLOW-004/005 and the identified contracts primitive bundles. Requirements
  and assertions are retained; existing lifecycle fixtures gain direct
  metadata/phase-base/shared-commit/provenance checks. See the
  [passing checkpoint review](reports/checkpoint-0045-0049.md) and PROGRESS
  for actual phase counts and approved execution-policy changes. Remaining
  gitstore/assets/app splits and D006 migration still precede the contracts drill.

- R02/R04/R05/R06/R17 refinement: 0050–0053 split gitstore/assets/app contracts
  with existing obligations and assertions retained. 0054 repairs PRJ-007's
  stale blanket C01 eval comparisons; exact cause migration remains pending.
  See the [passing checkpoint review](reports/checkpoint-0050-0054.md) and
  PROGRESS for phases and counts. This establishes no remaining acceptance gate.

- R04/R07/R11/R17 diagnostic refinement: 0055–0059 assign existing validation
  causes C04–C09 and boundary causes X01–X04, with exact cause assertions and
  accurate explanations. All phase audits passed; current-source build/race/vet
  and NEW JSON explanation smoke checks passed. See the
  [checkpoint review](reports/checkpoint-0055-0059.md) and PROGRESS for scope and
  phase counts. Remaining diagnostic migration and full acceptance are unverified.

- R04/R07/R09/R11/R17 diagnostic refinement: 0060–0064 migrate scoped CLI,
  static, eval and lifecycle causes while preserving existing behavior. All phase
  audits passed, including one scoped correction/re-audit in 0064. Current-source
  build/race/vet and NEW U01/U02 JSON smoke checks passed; see the
  [checkpoint review](reports/checkpoint-0060-0064.md) and PROGRESS for evidence
  and owner-approved policy updates. Remaining CLI classification/explanations,
  contracts drill, Milestone 2 and full acceptance remain unverified.

- R17 existing-path diagnostic slice: 0065–0068 finish current CLI classification
  and emitted-code explanations, including the observed show C02/C05 correction.
  All phase audits passed; NEW resolves all 30 emitted codes, targeted regressions
  pass, and race/vet passed at 72a2860 (0068 only corrects explanation text).
  PROGRESS and the archives hold phases/counts. This does not establish missing
  commands, the full future diagnostic surface, regeneration or final acceptance.

- R16 diagnostic drill evidence: the single contracts controller drill at 7afe9a0
  reports 34 own/339 dependent passes, retains its ref and removes scratch.
  [The review](reports/contracts-drill.md) records a CON-011/012 prefix gap and
  dependent-implementation reads outside the retained-input boundary. This run
  does not satisfy strict independence acceptance; no gap fix or repeat drill
  is authorized. Other module regeneration and final acceptance remain open.

- R01/R17/R19 partial: 0069 (APP-041–061) implements Go/none/local init and passes
  144 app cases after one scoped audit correction. NEW Go init/modules/check,
  current-source build/race/vet pass at cdc64ae. The
  [0065–0069 checkpoint](reports/checkpoint-0065-0069.md) initially blocked for
  missing recipe guidance in explain C12/U02. Owner-approved 0070 fixes
  AST-003/024 with four runtime failures → 51 own/351 total phase passes;
  scoped subagent re-review and NEW explanation JSON pass at 4b36a84. Prior
  causes/assertions are retained. History evidence covers 397 baseline commits
  plus the seven-commit correction extension, with no diagnostics. The passing
  re-review authorizes push/continuation. Doctor,
  remaining recipes, bounded execution and full acceptance remain unverified.

- R08/R11/R28 remain unverified for Doctor: 0071 is blocked in evals after two
  findings rounds, with the default-timeout assertion unresolved and a first-audit
  input-boundary deviation. [The block report](reports/doctor-0071-block.md)
  records exact phase evidence, retained attempt and proposed human-approved
  recovery. No Doctor implementation, CLI pass or restricted full audit is claimed.

  The subsequent approved recovery committed d78d618 with a real default-timeout
  assertion (255 phase cases, 29 runtime failures, 226 passes). Its fresh full
  audit stayed within permitted inputs and no longer flags CHK-065, but reports
  eight other assertion gaps. Controller block a81b723 retains audit_rounds 2,
  releases the claim and preserves attempt refs. Doctor remains unimplemented;
  the recovery result is a stop for further human review, not acceptance.

  A later approved eight-ID correction could not begin: native saved-attempt
  recovery retained audit phase, so tests stayed frozen. No test edit or worker
  launch occurred; native block 4c8439b released the claim. A one-time phase-only
  repair requires explicit approval before that correction can proceed. No
  counters, requirements or bounds changed, and no additional acceptance is claimed.

  The explicitly approved one-time phase repair and eight-ID correction completed:
  ebf127f records 260 phase cases (34 failures, 226 passes), and a restricted
  re-review plus valid prior passes led to audit commit 8cafea6. The implementation
  worker passed 71 own quick-check cases using Git 2.53.0, including the real
  default-timeout case. Final full verification did not complete; coordinator
  interruption produced native block ed8a9c7 and saved attempt 4. The frozen
  helper's demonstrated Apple Git incompatibility needs a normal eval_fix after
  approved recovery. No Doctor implementation application or acceptance gate is
  claimed; see the updated block report for the exact state and proposed recovery.

  Approved recovery now passed full native implementation verification and
  applied/archived 0071 at e50061d: 260 tests, all passing, with claim released.
  This establishes the CHK-049–069 backend eval evidence under Git 2.53.0;
  ordinary-Apple-Git helper portability, C14–C16 explanations and app doctor
  wiring remain pending. This is partial R08/R11/R28 evidence, not completion.

  0072 eval_fix restores ordinary Apple Git helper compatibility while preserving
  all CHK-050/051 placement, physical-cleanup and registration assertions. Evals
  dc2b1e7 and impl 762052f each report 260 passes; audit 426e259 has no findings
  and no implementation reads. Product implementation is unchanged.

- R17 partial: 0073 implements AST-026/027 cause/action guidance for emitted
  doctor C14–C16 and existing init P01 failures. Six fail-first cases precede
  389 full phase passes at 8d27f34, with no independent audit findings. Doctor
  CLI wiring, bounded execution and full acceptance remain open.

- R08/R11/R28 partial: 0074 exposes doctor and agents-md --write through NEW,
  with 153 app passes after one scoped audit correction. The fresh Go/fake-agent
  [CLI smoke](reports/doctor-setup-cli.md) at 19d2b31 passes doctor with and without
  agents, exact surrounding-text preservation, explanation lookup and scratch
  cleanup. This completes the currently scoped Go doctor setup slice; remaining
  recipes, bounded next/run, real regeneration and final acceptance remain open.

- R08/R09/R28 partial: 0075 establishes RUN-001–009 bounded worker/auditor
  invocation through the retained runner interface. Corrected evals record 25
  runtime failures; implementation 4b2fe8e passes 25 own/178 phase cases after one
  scoped re-audit. Queue selection, next/run orchestration, retries/recovery and
  the required NEW fake-agent queue fixture remain pending. This foundation
  does not establish unattended execution or final acceptance.

- The [0071–0075 checkpoint](reports/checkpoint-0071-0075.md) is BLOCK: required
  boundary hooks can be skipped, and cancelled doctor work leaves Git registration.
  A NEW missing-hook repro returned false success. Full race also exposes a
  timing-dependent RUN-007/009 eval; a focused pass does not close that failed
  full gate. Build, vet, static check and the 52-commit history extension pass.
  Existing healthy setup smoke remains valid, but doctor conformance, bounded
  execution and final acceptance stay open. P07/human synchronization is required
  before the proposed normal checks fix and runner eval_fix; no push occurred.


- The approved 0076–0078 correction gate now [passes](reports/checkpoint-corrections-0076.md):
  required-boundary/cancellation cleanup defects and baseline-stop coverage are
  corrected, valid isolated audits pass, the same scoped checkpoint reviewer
  returns PASS, and fresh full race/build/vet/static/history and NEW doctor smoke
  pass at a34a32a. Earlier phase/audit deviations remain historical. This is
  intermediate setup/execution-primitive evidence; next/run fixtures and every
  remaining completion gate stay open under the updated vertical-slice goal.

## Historical reference inspection record — superseded by D004

These entries describe the retired controller and original preparation. They supply historical provenance only, not startup instructions or current readiness. PROGRESS.md records the installed stage-2 controller, completed setup and baseline.

- Bootstrap clean source commit: `583b3c53740cf2341e0b0a79afee5b7f646ec1d4`.
- Local design snapshot: `fireweed-design-v1.md` at the repository root, copied byte-for-byte from `/Users/joel/Workspace/fireweed-bootstrap/fireweed-design-v1.md` at that commit on 2026-10-21.
- Design SHA-256: `1c17d0ab9496b1b15ad5ca18eed4f4ca623d01cff596bc557bf4c97400d00789`. Revised 2026-10-26: `a1fe8a2a2335c04443d4e7d3e6aa3a138e221cd5bd8302c6c66bfd947761d919`; usability revision 2026-10-26: `5d7cfcc9e3987fb594d16cce94f5ffa057c583f17f985114fe1de946aef5cf9d`; merge-strategy revision 2026-10-26: `52623b89d8184680f535519e82ceda8a49490b3d4dec000347cbd2cde429b21a`.
- Target initial commit: `d543462579d3fe182460c07476846476d38c6dc5` on main; superseded by the stage-2 baseline in PROGRESS.md.
- The original temporary 0.4.0 build used `go1.27.2 darwin/arm64`. A disposable Go-recipe init/doctor probe passed four canaries without real agent/auditor execution. Neither workspace was initialized by it. The temporary binary is not a durable input or a reason to rebuild/repin on resume.
