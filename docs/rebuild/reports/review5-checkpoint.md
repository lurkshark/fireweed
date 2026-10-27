# Review 5 repair checkpoint

At the owner's explicit request, the reviewed main through the prior checkpoint
was pushed to origin/main and P07 cleared. Human review 5 restarted the applied
count at zero. Five further changes have now applied: 0025–0029. Stop here for
review; these new changes have not been pushed.

FLOW-003 is narrowed to the full-verification gate. CHK-004 is narrowed to audit's
current-file restriction; CHK-006–024 retain its other permissions and prohibitions.
The four original examples from review 3 are now split. The broader spec repair
remains incomplete. Workflow and checks implementation, interfaces and models
are unchanged throughout this batch; reported case counts remain 32 and 28.
All worker/audit provenance records the real configured Sol/high command.

0026 had one audit-findings round: the existing unmatched outside-path rejection
was moved into the existing evals/impl loop to cover both named phases. No new
test cases or fixture instances were added. 0025 reused existing full-failure
fixtures to assert that Done rejects even when quick Verify succeeds.

## Verification

Checked source: `db03f6210e9f80cd7a5546cca66411eee1695baf`.
History range: `be35552b706aa3000ff8f11306d129ff1d701b85..db03f6210e9f80cd7a5546cca66411eee1695baf`.

| Check | Result |
| --- | --- |
| go test -race ./...; go vet ./... | Passed |
| Pinned controller CI on explicit range | Passed; doctor canaries/recovery passed; no history diagnostics |
| CI coverage | 261 passes: app 86, assets 20, checks 28, contracts 34, gitstore 8, project 53, workflow 32; runner remains empty |
| NEW build | Passed, CGO disabled, clean checked-source metadata |
| Pending changes / claims / scratch worktrees | None |

CI reports only the known embedded recipe SPEC.md warning. Status reports P07
at five unapproved archives and P08 for the explicitly authorized consecutive
checks splits. No bounds were changed. Archived phase commits and permanent
evals are the reviewable evidence; ignored runtime logs are auxiliary.

## Remaining work

- Continue atomic splits for FLOW-004/005, CHK-001/002/003/005 and the other bundled contracts identified in the prior report. Prioritize PRJ-007's validation rules so its blanket C01 promise can be removed before diagnostic migration.
- D006 is authored, but wrong-code eval_fix changes, normal implementation migration and complete accurate explain coverage remain pending.
- Before Milestone 2 touches workflow implementation, split lifecycle/stub.go by responsibility as maintenance only. No other cleanup is authorized.
- The single real contracts drill remains gated on completing both spec and diagnostic repairs. Commit its report and gaps, author no gap changes and drill no other module.
- Milestone 2 then centers projected blocking and fake-agent next/run through NEW. No NEW unattended real-agent execution before that fake-agent run succeeds.

The applied count remains five until human review and shared-base synchronization.
