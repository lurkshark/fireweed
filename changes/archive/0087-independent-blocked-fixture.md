---
audit_rounds: 0
bases: {}
eval_fix:
    - RUN-011
module: runner
phase: applied
run: 20261029T070406-ea5196b72786ec9e4118faa0
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

## Provenance

- evals 2026-10-29T07:10:35.574684Z agent=manual fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=213 failing=0 passing=213
- audit 2026-10-29T07:13:42.81725Z agent=manual fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0
- impl 2026-10-29T07:13:46.482161Z agent=manual fireweed=0.1.0 playbook=07554a3249d8fca0ba70df7504135f297491465caa63458734a4b8351734d0e5 tests=213 failing=0 passing=213

## Audit

No findings
Checked IDs: RUN-011
