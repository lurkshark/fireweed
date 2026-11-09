# Workflow report discovery and X01 detail

The owner approved the two workflow changes and Git compatibility guidance
identified in the failure investigation. The change process applied both
archives. The candidate for owner review is
`7d17a76186fd94d603a018713bdcc1005c89e020`.
[Failure investigation](workflow-native-retry-20261108T233724/investigation.md).

## Audited changes

| Change | Result | Evals / audit / application |
| --- | --- | --- |
| [0190](../../changes/archive/0190-retained-report-discovery.md) | FLOW-052 makes each evaluated module's saved JSON report discoverable. Ordinary reports keep the flat filename. Module names matching `costs` without regard to case use `<phase>/reports/<module>.json`, preserving the phase cost snapshot. | `ccf0874` / `7ec0853` / `ab4f19d` |
| [0191](../../changes/archive/0191-boundary-violation-detail.md) | FLOW-062 explicitly requires the selected module-attributed violation in a returned X01 error. The strengthened public eval passed against the existing behavior. | `0a402d1` / `90e8882` / `7d17a76` |

The protected guidance commit `5c2b61c` added the existing supported-host
policy and actual Apple Git 2.30.1 version to the workflow worker instructions.
It leaves the choice of Git operations to the worker. It supplies no
reconstruction source or failure-specific instructions. The independent
controller pin remains the `75efa9a` build.
[Worker instructions](../../docs/rebuild/workflow-drill/worker-system.md),
[pin record](controller-pin-75efa9a.md).

The queue was authored at `efc8a89`. The installed controller ran
`fireweed run --max 2 --json` with normal approval, worker and audit checks.
It made all six phase commits and returned exit 0 with two applied changes and
an empty queue. Separate restricted gpt-6-astra/medium audits returned no
findings for FLOW-052 and FLOW-062. Their scratch inputs were removed.

0190's fail-first phase had 175 passing rows and three failing FLOW-052 rows,
including the parent row. Both fresh fixtures completed Done and reached the
missing report assertion for `costs` or `CoStS`, with the phase snapshot
still present. After the report path fix, all rows passed. 0191 passed all
178 own rows with the existing runtime, including the X01 and X02 cases.
Full impl verification for each change reported 660 passing rows across
workflow, app and runner, with all enforced IDs referenced. Worker quick
verification took 146.448–187.958 seconds. Both full implementation phases
finished within the 25-minute bound. [Phase verification](workflow-report-boundary-corrections/phase-evidence.json),
[controller result](workflow-report-boundary-corrections/run.json).

## Candidate validation

A no-hardlinks clean checkout at the exact candidate revision built a separate
Go 1.27.2 darwin/arm64 subject with `-trimpath` and `vcs.modified=false`.
The setup hook copied that invoker byte for byte in the validation checkout.
The compiled CLI fixture confirmed that evals and impl Done retain both
`reports/costs.json` and the cost snapshot, that the ordinary dependent report
stays flat, and that a failed full Verify returns the concrete X01 message
without advancing HEAD. The fixture then applied successfully and released its
claim. [Candidate build](workflow-report-boundary-corrections/candidate.json),
[CLI fixture](workflow-report-boundary-corrections/fixture.json).

The first default `go test ./...` exited 1 in
`Test_RUN_036_RUN_037_OneBudgetAcrossChangesRetainsApplied`. The ten-second
fixture budget returned zero applied changes after invoking alpha's evals and
audit. The other suites passed. A full serial `go test -p 1 ./...` passed
with the same test bounds. This comparison does not establish why the
parallel run exhausted its budget. Earlier timing exceptions remain open.
The affected workflow, app and runner race suites were run serially.
The unaffected race proof is reused from the earlier validated candidate
because its package sources, dependencies and hooks are identical here.
[Earlier race proof](workflow-contract-diagnostics-handoff/validation/results.json).

The explicit new history delta is
`75efa9a03c4eb7a17673a53bdd8058074e5bb7e1..7d17a76186fd94d603a018713bdcc1005c89e020`.
It covers 13 commits. The composed history range from the standing
`be35552b706aa3000ff8f11306d129ff1d701b85` baseline covers 1,274 commits,
using the existing verified prefix.
Candidate `check-commits` passed on the delta. Candidate static check, all ten
agent-free hook doctor checks, `go vet ./...` and candidate CI with hook
doctor enabled passed. CI checked all modules and reported no diagnostics.
The six lifecycle commits separately passed the installed controller's
explicit-range history check. [Validation commands and results](workflow-report-boundary-corrections/validation.json).

## Preservation and stop

The main checkout and both scratch checkouts were clean at the validation
handoff. The two restricted audit inputs, validation clone, compiled subject
and CLI fixture were removed. The queue and shared claims are empty. The
live unapproved count went from zero to two with the two archives. Existing
B02 warnings for CHK-055, APP-103 and APP-104 and their bounds were preserved.
The existing native workflow report and proof refs, controller pin,
recoveries, configuration and unrelated linked worktree were preserved.
Idle Context U02 and Verify P02 record that the checkout holds no change.
[Preservation and cleanup](workflow-report-boundary-corrections/preservation.json).

The approved change and validation scope is consumed. WF4 and full completion
remain open. A new supervised native invocation on the exact candidate
requires separate owner approval. The existing deadline remains
**2026-11-10 22:51:05 UTC**.
