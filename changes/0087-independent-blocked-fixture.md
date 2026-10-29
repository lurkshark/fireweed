---
bases: {}
eval_fix:
    - RUN-011
module: runner
---



## Motivation

Full verification of projected blocking eligibility rejected the runner's
first_unblocked fixture: it expects a change to start behind a blocked change
in the same module. The design's Blocking rule requires that work to wait.
RUN-011 tests first-unblocked selection and propagation of lifecycle Start
failures; it does not promise that Start accepts a same-module blocked queue.
Move the preceding blocked fixture to an independent module, preserving every
existing assertion and coverage ID. A disposable-copy probe passed this exact
fixture correction under both current committed and new Start selection.

## Interface

No interface or product behavior change. Only the existing runner fixture
setup changes; blocked skipping, selected change, held phase/claim and worker
failure assertions remain identical.

## Added



## Modified



## Removed



## Decisions



## Notes

Apply the reviewed runner-fixture-candidate.patch via the normal eval_fix
workflow, preserving all tests/assertions/IDs and passing all own evals. No
product, interface, SPEC, other-module or enforcement edit. Independent scoped
assertion audit remains required. Governing controller owns commits. This
preparatory fix precedes any future fresh authoring of projected eligibility.
The owner now requires stopping immediately after this change archives and
generating a current-progress review; do not re-author eligibility or push. Do not use
any discarded eligibility implementation as an oracle.
