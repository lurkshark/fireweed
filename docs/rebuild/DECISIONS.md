# Rebuild decisions

Clarifications of the design and intentional differences from the bootstrap, one
short entry each. Decisions from the first run are on the `rebuild-v1` branch;
[LESSONS.md](LESSONS.md) summarizes the ones that still matter.

## D001–D004 — History (2026-10-26)

The first run is preserved on `rebuild-v1`; `main` restarted from the setup
baseline. The owner then accepted three design revisions, all reflected in
fireweed-design-v1.md and tracked in ACCEPTANCE: usability (R26–R28), any merge
strategy (R29) and bounded effort (R22–R25). The rebuild uses the tested
bootstrap as its pre-cutover controller (0.1.1), never as an implementation
oracle, and proves behavior in vertical slices through NEW. PROGRESS.md records
the current controller and the retained bootstrap recovery pin after EC6.

## D005 — Initial product diagnostic allocation

The design requires stable codes and explanations but does not assign a meaning
to P2-01. AST-003 assigns it to missing genuine fail-first evidence.
The bootstrap's operational `explain P2-01` instead describes a forbidden-zone
edit. The completed controller/product smoke records this intentional difference;
the authored product meaning remains authoritative. Bootstrap code assignments are not
product authority. Other code mappings remain to be authored with their behavior.

## D006 — Diagnostic allocation (human review 3)

Codes identify a stable cause. Their prefix does not determine severity; the
diagnostic carries its warning status. Assignments below apply to the NEW
product. This allocation did not itself change the controller pin.

| Range | Area and fixed assignments |
| --- | --- |
| C01–C99 | Deterministic configuration and static checks. C01 configuration discovery/decoding; C02 module graph; C03 document/item syntax; C04 normative text; C05 ID ownership/uniqueness; C06 spec tags; C07 change structure/Motivation; C08 bases; C09 handles; C10 coverage; C11 eval-hook/JUnit contract; C12 metadata/zone declarations; C13 queue identity. C14–C16 reserved. C17 spec shape, C18 copied conventions, C19 Decision length retain design meanings. |
| P01–P99, P2-01 | Lifecycle and approval. P01 preconditions; P02 claims; P03 approved state; P04 approval warning; P05 stale base; P06 blocked/Questions; P07 unapproved-work limit; P08 repeated-module warning; P09 audit submission. Existing P2-01 remains missing genuine fail-first evidence. |
| Z01–Z99 | Phase-zone permissions. Existing Z01 remains a forbidden path edit. |
| X01–X99 | Boundaries. X01 unsuppressed violation; X02 required hook missing or unable to complete; X03 reasoned opt-out; X04 boundary enforcement unnecessary. B05 is retired. |
| U01–U99 | CLI usage and lookup. U01 invalid arguments/options; U02 unknown lookup name. Replace the legacy CLI_* labels. |
| E01–E99 | Required eval results. E01 an enforced testcase failed where passing results are required. |
| H01–H04 | History, with the design's exact meanings. |
| B01–B04 | Budget/report warnings only, with the design's exact meanings. No B05 allocation. |

Unlisted numbers are reserved. Every emitted code must have an accurate cause
and corrective action in explain. Wrong label assertions are repaired through
eval_fix without losing the behavior assertions or ID coverage. Normal changes
then establish exact assigned-code assertions and migrate implementation. The
spec split removes the old blanket C01 promise before that migration; no existing
Specs or Decisions are edited by hand.

## D007 — Doctor diagnostics

Doctor uses the previously reserved C14–C16 slots:
C14 for doctor hook/canary setup, detection and recovery failures; C15 for agent
probe failures or usage limits, and the shared-auditor-command warning; C16 for
stale or missing managed guidance and playbook override base metadata. Severity
remains a property of the diagnostic. D006's existing assignments remain unchanged.

## D008 — Regenerate warning

The regenerate lifecycle feature assigns P10 to the warning that its module
lacks a passing drill dated after the latest applied implementation Provenance.
This allocation grants no drill authorization; PROGRESS records current
status and GOAL links the planned milestone.

## D009 — Early core-loop cutover milestone (owner, 2026-10-30)

The owner selected an early Go/current-format core-loop cutover before drills,
regeneration validation and the remaining product surface. It required 0134–0138
and ACCEPTANCE EC1–EC5 under bootstrap, then EC6 as a separate controller
pin/handoff; PROGRESS records completion. Full design obligations and
full-completion gates remain open; this changed sequencing, not product scope
or enforcement.

## D010 — Held-change continuation

NEW `next` resumes this worktree's held change directly. The pinned bootstrap
instead rejects that live claim with P01; handing a manually advanced phase to
its unattended loop requires releasing the owning checkout's claim first.
This controller difference adds no product requirement.

## D011 — Reference lifecycle output and audit metadata

The reference smoke returns string data for bootstrap next --json; NEW returns
the change/run/phase/empty object required by APP-079. Both use the schema-1
envelope and complete the same three lifecycle commits and held ref shape.
NEW's independent audit metadata has an input-only initial commit, while
bootstrap uses empty metadata; neither exposes held history or saved attempts.
