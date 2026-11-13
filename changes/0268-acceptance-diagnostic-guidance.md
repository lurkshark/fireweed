---
module: assets
---

## Motivation

R17/R20 require explanations to match enforced checks, and the design's Git
behavior explicitly promises recovery instructions for duplicate-machine runs
(R14). The current C07 explanation omits the new change-shape/target causes,
P05 omits the stale branch-tip warning, and no
explanation gives the duplicate-run merge recovery procedure.

## Interface

No public signature or type changes.

## Added

### new:change-shape-explanation
Explain("C07") MUST describe invalid fixes/eval_fix/already_satisfied lists or targets and non-heading content in Removed as invalid change structure. It MUST advise correcting the listed identities or moving removal commentary into Motivation or Notes.

### new:stale-tip-explanation
Explain("P05") MUST distinguish a warning that HEAD lacks the locally resolved base tip from a stale spec-hash rejection, and advise reviewing and updating the branch before continuing a potentially stale queue.

### new:duplicate-run-recovery
Explain("P02") MUST describe duplicate runs on different machines as a conflict in the archived change's run identity. It MUST advise aborting the merge, using fireweed commits <change> to identify the incoming duplicate's commits, reverting those commits on the incoming branch, then merging again. It MUST warn that choosing one archive side can retain both implementations.

## Modified



## Removed

## Decisions

## Notes

Keep all existing cause/correction assertions and codes; this change extends
guidance without changing severity or runtime checks. C01 already describes
invalid required configuration and C03/C12 already cover malformed document
and zone definitions, so avoid duplicating their generic explanations.
Use public assets.Explain tests accepting equivalent wording. Git branch,
merge, revert and push operations remain user actions, never product actions.
The duplicate-run explanation uses existing claim/conflict code P02 rather
than allocating a new reserved diagnostic number.
This is discoverable recovery guidance, not a claim that Git merge conflicts
emit P02. No nested-module prohibition or C02 ownership guidance is added.
Coordinate C07 with 0260/0261 and P05 with 0262/0272: the explanations and
behavior must both be present before closing their acceptance gaps. These
assets edits need no new project or workflow API and may be applied first.

Drafted at the owner's request after the R01–R29 traceability review.
No lifecycle execution, agent invocation, drill or release is authorized by
this drafting request. Workers must obtain context before their step and use
fireweed verify; only the controller may advance or commit lifecycle work.
