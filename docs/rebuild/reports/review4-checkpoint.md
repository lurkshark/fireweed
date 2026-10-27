# Review 4 repair checkpoint

Five changes have applied since review 4: 0020–0024. FLOW-002 is narrowed to
phase-zone rejection; FLOW-020–031 retain its other guarantees. FLOW-032–040
separate nine FLOW-003 behaviors; its final narrowing remains pending. Workflow
implementation, interface and model are unchanged throughout this batch. Existing
evals were retargeted, with the same 32 reported workflow cases. All phase
provenance in this batch records the configured Sol/high command.

0021 stopped at a Sol usage limit after its passing audit. Following human
confirmation and a successful configured access probe, the controller restored
saved attempt 1 and completed impl. The review count was not reset.

## Verification

Checked source: `c64765f2bb21c4449cdf6d5b1e1778a71dd1eb4f`.
History range: `be35552b706aa3000ff8f11306d129ff1d701b85..c64765f2bb21c4449cdf6d5b1e1778a71dd1eb4f`.

| Check | Result |
| --- | --- |
| go test -race ./...; go vet ./... | Passed |
| Pinned controller CI on explicit range | Passed; doctor canaries/recovery passed; no history diagnostics |
| CI coverage | 261 passes: app 86, assets 20, checks 28, contracts 34, gitstore 8, project 53, workflow 32; runner remains empty |
| NEW build/version | Passed, CGO disabled, version 0.1.0, clean checked-source metadata |
| Pending changes / claims / scratch worktrees | None |

CI reports only the known embedded recipe SPEC.md warning. Status reports P07
at five unapproved archives and P08 for the explicitly authorized consecutive
workflow splits. No bound or base ref was changed. Permanent evals and archived
phase commits are the reviewable evidence; ignored runtime logs are auxiliary.

## Remaining work

- Finish FLOW-003: retain passing-audit, Questions and claim-release behavior under separate specs before narrowing the old item. FLOW-004/005 also bundle independent behaviors.
- Split CHK-004 and the other checks items. Inspection also found separate APIs packed into CON-001/002/005 and GIT-003/004; project, assets and app still contain bundled authoring, configuration, reporting and validation behaviors. Review those splits by observable behavior, preserving assertions rather than multiplying cases.
- D006 is authored; wrong-code eval_fix work, normal implementation migration and accurate explain coverage have not begun. Remove the blanket C01 promise in PRJ-007 through the split before migrating it.
- The workflow file split remains a behavior-preserving maintenance change before Milestone 2 touches workflow implementation. No other cleanup is authorized.
- The one real contracts drill has not run. Complete the spec and diagnostic repairs first, commit its report and gaps, author no gap changes and drill no other module.
- Milestone 2 remains gated; its core is projected blocking and a fake-agent next/run through NEW. No NEW unattended real-agent execution before that fake-agent run succeeds.

Stop here for human review and synchronization of approved work with the shared
base. The applied count remains five until that checkpoint is cleared.
