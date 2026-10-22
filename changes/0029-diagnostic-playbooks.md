---
bases:
    AST-003: c7f349e8
module: assets
---

## Motivation

Keep the embedded diagnostic explanation aligned with the project resolver and executable: D05 now rejects unsafe/unreadable selected overrides as well as describing doctor staleness warnings. An explanation mentioning only stale metadata gives the wrong remedy for a contained-read failure. This follows the design requirement that embedded protocol instructions match verification.

## Interface

No public signature or inventory changes. Extend the embedded D05 explanation returned by assets.Explain, still without runtime file/network reads.

## Modified

### AST-003
The diagnostic catalog MUST contain U01, C01 through C16, P01 through P06, P1-01 through P1-03, P2-01, D01 through D07 and H01 through H04, sorted and without duplicates. Each explanation MUST identify the problem and a corrective action. H01 MUST cover protected worker files, H02 immutable impl evals/interfaces, H03 append-only changes by anyone including merge resolutions, and H04 immutable archives with deletion for revert allowed plus duplicate recovery by aborting the merge, listing commits, reverting the incoming duplicate and merging again. D05 MUST explain both unsafe or unreadable playbook overrides and stale based_on metadata. Its corrective guidance MUST tell the user to restore a readable contained override and, for staleness, compare the embedded hash, review instructions and update based_on. Unknown codes MUST return errors. Returned code inventories MUST be independent slices.

## Decisions

D05 retains its existing code and covers both operational override failures (PRJ-023 and APP-018) and advisory stale metadata. Restore readable contained instructions for read/containment failures; review the embedded hash and update based_on for staleness. This change updates explanation data only and does not add doctor or change warning/error policy. Existing catalog membership and other explanations remain covered by permanent evals.

## Rejected alternatives

An explanation describing only stale hashes cannot guide the unsafe/unreadable error path. Reassigning stable diagnostic codes would needlessly change the already checked resolver/CLI contracts. Silently editing embedded data without a failing regression would bypass the governed process.
