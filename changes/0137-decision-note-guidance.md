---
bases: {}
module: assets
---



## Motivation

The design's Decisions section says retained choices describe previously
unspecified behavior in one or two sentences, while tests and process records
belong in Notes. LESSONS.md records that workers once accumulated about 13,000
words of Decisions, burdening later workers and regeneration. Project-local
instructions now separate them, but the product's worker playbooks omit that
rule. The owner requested carrying it into the product. NEW playbook evals,
impl and regenerate should each expose the guidance and short examples, which
existing worker context already delivers.

## Interface

No public API changes. Expand the text returned by Playbook for evals, impl and
regenerate. This is instructional guidance using existing change sections.

## Added

### new:worker-decision-note-guidance
The evals, impl and regenerate playbooks MUST instruct workers to record only previously unspecified behavioral choices in Decisions, using one or two sentences per choice, and to put verification, debugging and recovery records in Notes.

### new:worker-decision-note-examples
Each of those playbooks MUST include short illustrative examples classifying a lasting behavioral choice as a Decision and a verification or recovery event as a Note.

## Modified



## Removed



## Decisions



## Rejected alternatives

Project-only policy: it does not help other users receiving embedded playbooks.
New blocking checks: the requested improvement is guidance for classification,
not automatic semantic judgment of a worker's prose.

## Notes

Preserve 0134's eval stub handoff guidance and existing eval_fix, phase-zone and
regeneration instructions. Keep audit scoped to its existing restricted role;
this change covers the three worker playbooks. No new sections, configuration,
size enforcement or changes to how Decisions/Notes are stored or applied.
