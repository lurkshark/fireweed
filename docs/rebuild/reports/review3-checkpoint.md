# Review 3 repair checkpoint

Five changes have applied since the review-3 reset: 0015–0019. Work stops here
for the requested human checkpoint. FLOW-001 is now one claim-creation rule;
FLOW-006–FLOW-019 name its other observable guarantees. Workflow implementation,
public interface and retained model are unchanged from review commit 1fe462f.
The eval_fix in 0016 corrected 0015's extra test functions: workflow remains at
32 reported cases using the existing fixtures and assertions.

## Verification

Checked source: `bde237f4f093056cd578ab47c1f677bdcecb9cd2`.
History range: `be35552b706aa3000ff8f11306d129ff1d701b85..bde237f4f093056cd578ab47c1f677bdcecb9cd2`.

| Check | Result |
| --- | --- |
| go test -race ./... | Passed |
| go vet ./... | Passed |
| Pinned controller CI on explicit range | Passed |
| CI history diagnostics | None |
| CI doctor canaries and recovery | Passed; agent probes not run by CI |
| CI coverage | 261 reported cases passed: app 86, assets 20, checks 28, contracts 34, gitstore 8, project 53, workflow 32; runner has no cases |
| Build | Passed with CGO_ENABLED=0; clean source metadata |
| Pending changes / active claims / scratch worktrees | None |

The controller reports the known template-file warning for
packages/assets/internal/recipe/files/SPEC.md. Status reports P07 at five
unapproved archives. Its P08 warning reflects the explicitly authorized
consecutive workflow spec-split changes. No bound was changed.

## Remaining review work

- FLOW-002, FLOW-003 and CHK-004 remain bundled; inspect the other existing specs as directed.
- D006 assigns diagnostic ranges, but wrong-code evals, implementation migration and complete explain coverage remain to be changed through their normal workflows.
- The optional workflow file split remains a maintenance change before Milestone 2 touches its implementation.
- The single real contracts drill has not run: it is gated on completing the spec and diagnostic repairs. Its report will be committed; no gap changes will be authored automatically.
- Milestone 2 and NEW unattended real-agent runs remain gated on these repairs and the required fake-agent NEW run.

Phase commits, checked IDs and actual configured Sol/high commands are recorded
in the archived changes. This committed summary and permanent evals are the
reviewable evidence; ignored runtime logs are auxiliary only.
