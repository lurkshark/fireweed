# Review 6 spec-split checkpoint

The owner approved 0025–0029 and authorized the next spec-split slice beginning
with PRJ-007. The shared base was already synchronized at c2da5e2. Two changes
have applied since that review: 0030 and 0031. Stop before a third consecutive
project split under SYSTEM's human-checkpoint rule; the existing consecutive
split exception covers workflow/checks only. No third project change is authored.

PRJ-008–016 now retain PRJ-007's validation rules, and PRJ-007 states the
valid/error diagnostic contract without the blanket C01 allocation. Product
implementation and interfaces are unchanged. Existing assertions are preserved;
the existing cases were relabeled and extended for the clauses identified by the
separate audits. There are still 53 project cases. Each change used one findings
round; each re-audit cleared only the previously reported IDs. See the archived
changes and PROGRESS for IDs, phase commits and counts.

## Verification

Checked source: `ac3779b71bbdab0df016a99e74db38c4b190a659`.
History range:
`be35552b706aa3000ff8f11306d129ff1d701b85..ac3779b71bbdab0df016a99e74db38c4b190a659`.

| Check | Result |
| --- | --- |
| go test -race ./...; go vet ./... | Passed |
| Pinned controller CI on the explicit range | Passed; doctor enabled; no history diagnostics |
| CI coverage | 261 passes: app 86, assets 20, checks 28, contracts 34, gitstore 8, project 53, workflow 32; runner empty |
| NEW build and version JSON | Passed; CGO disabled; clean checked-source revision recorded in Go build metadata |
| Pending changes / claims / scratch worktrees | None |

The only controller warning is the known embedded recipe SPEC.md outside the
configured module globs. Status reports two unapproved archives and no P08.
No controller, bound, ref or requirement was changed to pass verification. The
build exited 0 with a nonfatal sandbox denial writing Go's module-version cache;
`go version -m` confirms the complete source revision and vcs.modified=false.
No agent access probes or drills were repeated; passed setup and earlier
lifecycle evidence remain reusable. This is a checkpoint, not full acceptance.

## Remaining work

- Review this slice before any third consecutive project split. Remaining
  project bundles include PRJ-001–006; other previously identified module
  bundles also remain. Keep changes atomic and reuse existing assertions.
- D006's wrong-code eval_fix and normal implementation migrations remain
  pending. C01 assertions are deliberately retained until audited eval_fix;
  this slice removed the spec obstacle, not the implementation labels.
- The workflow responsibility split, single contracts drill and Milestone 2
  work remain sequenced as recorded in PROGRESS. No drill ran in this slice.
- No pushes, publication or tags occurred. Two applications since review remain
  recorded; review and shared-base synchronization govern the next slice.
