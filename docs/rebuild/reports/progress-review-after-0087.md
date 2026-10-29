# Current progress review after the fixture correction

The rebuild now has a demonstrated checked CLI lifecycle, bounded successful
queue execution, verification retries, and saved-work blocking on exhaustion.
It is still a partial implementation of the full design. Block-aware queue
scheduling and much of the remaining command surface are unfinished, and final
regeneration, self-hosting and delivery acceptance are unproved.

Work stopped at the requested boundary: 0087-independent-blocked-fixture is
archived. This is a progress assessment, not the final focused design/inventory
review or an independent code audit.

## Checkpoint state and evidence

Verification head: 931c17feb97659a96f3eb4b4ceff0eb1a43218cc. The correction's
full eval and impl phases each passed all 213 cases; its isolated RUN-011 audit
passed. The working tree was clean, pending queue empty and claim refs absent.
Controller status reports four unapproved applied changes since the reviewed
push at 8aff1e4: 0084, 0085, 0086 and the new fixture eval_fix. Generating this
review does not reset the four-change count. No new push occurred.

Controller static checking passed. Explicit history
6ce22e9e9d618e7c5c4451f48b3f1b522ab4cf08..931c17feb97659a96f3eb4b4ceff0eb1a43218cc
passed all 31 commits with zero diagnostics, extending the recorded valid
baseline portions. The existing C01 embedded-recipe SPEC discovery warning and
B02 for 16 CHK-055 cases remain informational and documented.

The last built NEW executable is from clean c7972645f5079fd93c1267fbf5aee4eab7f21f5e,
Go 1.27.2 darwin/arm64, CGO disabled, vcs.modified=false. A source diff confirms
product code/interfaces/specs outside evals are unchanged since that build;
later work is the fixture correction, documentation and abandoned eligibility
attempt. The binary was not rebuilt merely to update its revision label.

## What is demonstrated

| Capability | Evidence and practical limit |
| --- | --- |
| Manual checked lifecycle | [NEW lifecycle](manual-lifecycle.md): real fail-first, separate audit, implementation, archive, provenance and claim cleanup. |
| Setup and guidance | [Doctor/setup](doctor-setup-cli.md) and approved corrections: Go/none/local init, doctor canaries/cleanup, guidance and agent invocation. Other built-in language recipes remain outside the completed slice. |
| Successful next and bounded run | [Next](next-0080-block.md) and [Run](bounded-run-cli.md): actual Go fixtures, --max, later continuation and empty queue. |
| Verification retry | [0084 proof](verification-retries-0084.md): retained edits and actual latest verification feedback; successful CLI retry in 6.53s. |
| Exhaustion with recoverable work | [0085 API](saved-attempt-blocking-0085.md) and [0086 CLI](exhausted-next-0086.md): final failed edits on a real attempt ref, phase-base restoration, state-only block, original cause, clean claim release; actual CLI proof in 5.61s. |
| Fixture compatibility | [0087 correction](independent-blocked-fixture-0087.md): all assertions preserved while removing an invalid same-module eligibility assumption. |

These are bounded fake-agent orchestration proofs. They establish runnable
behavior; they do not establish real-agent regeneration or the final package.

## Main findings

1. **Projected blocking and independent continuation are the immediate gap.**
   The eligibility attempt passed its own 70 cases and assertion audit, but
   full verification exposed a dependent runner fixture conflict. The owner
   approved abandonment and the preparatory eval_fix. Consequently the new
   eligibility API/Start rules are not in the current product. Manual Start
   and Next still use the earlier unblocked-head policy; Run still stops on a
   failure. Do not count blocked dependency waiting or independent progress
   as completed capabilities.
2. **Recovery remains partial.** Saved-attempt loading/resume, agent-commit
   undo, nonzero-agent retry, usage waits/budgets, --for and clean exit 3,
   remaining change kinds and persistent-bound coverage still need focused
   scenarios. The existence of an attempt ref does not prove resume works.
3. **The command and acceptance backlog is substantial.** Source dispatch still
   explicitly rejects new-module, list/status, check-commits/ci/pr-body,
   release/commits/abandon, drill/gc/migrate. show history/stale and check --fix
   are also unfinished. The authoritative inventory remains
   [ACCEPTANCE](../ACCEPTANCE.md); no completion percentage is justified by
   archive/spec counts.
4. **Validation recency has limits.** Latest phase checks and the history/static
   checks above pass. The last whole-root build/vet/race checkpoint was at
   6ce22e9, before the later recovery code; no fresh-checkout/final Go gate or
   current full controller doctor/CI gate is claimed. Reuse unaffected evidence,
   but refresh the required final gates against the eventual source.
5. **Independence and delivery are still open.** The sole contracts drill has
   recorded limitations, including dependent-internal reads, so strict
   regeneration independence is unverified. Other real module drills, isolated
   NEW self-hosting, the final comparison and packaging/delivery remain open.
   No new drill or gap-fix permission follows from this review.

## Process assessment and resumption

Recent work has returned to observable vertical slices: retry and exhaustion
both ended in actual NEW CLI results. Full dependent verification was useful:
it caught the fixture contract assumption that own-module tests could not.
The two approved 0085 phase recoveries and earlier invalid mechanical audit
passes remain documented; passing corrections do not retroactively validate
those submissions. No verification or ownership bound was raised or bypassed.

On resumption, re-author projected eligibility from retained requirements,
interfaces and evals, with a fresh implementation rather than the discarded
code. Then wire Next/Run to it and prove Alpha blocked, dependent Beta waiting,
and independent Gamma applied through actual NEW run. Keep status views and
broader recovery out of that immediate slice. With the current count of four,
the next applied change reaches the standing five-change review checkpoint
unless the owner explicitly changes the checkpoint authority.

The user's stop overrides automatic continuation. No next change, push, drill
or eligibility re-authoring is started. The full rebuild goal remains incomplete.
