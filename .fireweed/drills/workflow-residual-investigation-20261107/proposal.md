# Residual workflow drill failures: evidence and scoped proposals

> Retention update: this review describes the evidence at `d63c16e`.
> Superseded attachments are recoverable from [Git history](../README.md#historical-evidence);
> current module reports and selected outcome receipts remain in this checkout.

The owner requested investigation and fix proposals, followed by one further
drill while keeping candidate `79cf33a7690082f28c06bf41a72719d282bc00fe` and the
fixed pinned controller unchanged. This document proposes corrections; it applies
none. The requested new invocation is an unchanged-input diagnostic repeat, not
validation of these unapplied proposals. Any passing result leaves this retained-
eval correction checkpoint open. No investigation finding or saved reconstruction
will be supplied as an extra worker input or hint.

## Reproduction and classification

The retained failed ref
`refs/fireweed/drills/workflow/20261107T202952.649365000` points to
`50bd33c5526fb495d541e3290868d198ed010b78`. Existing public evals were run, unchanged,
in a disposable clone at that exact ref with bounded fake fixture workers only.
Both failures reproduced. The same two existing tests passed after checking out
the frozen candidate in that owned clone. No tracked source, spec or eval was
edited; no real agent or native drill ran during investigation. The clone and
probe fixture were cleaned up with claims released. Commands, bounds, outputs and
comparison are retained beside this proposal.

### FLOW-083 is an early Start deletion failure

The failing test aborts at regenerate_start_test.go:216 while reading the recorded
`packages/alpha/internal/ready` file. Its Context assertion is never reached.
The saved reconstruction's start function executes deleteDisposable for every
ordinary Start with regenerate true, including a change already at phase impl
with run fixture-recorded; only StartFromAttempt avoids that branch. It checks
whether phase is absent when assigning a phase, but does not preserve that
unstarted condition for deletion.

A separate bounded public-API probe confirms the distinction: recorded code is
deleted, while Context still contains the exact regenerate playbook and excludes
the ordinary impl playbook. Its fixture claim was released. This is a saved-
reconstruction initialization defect, not demonstrated wrong playbook selection
or an established bug in candidate 79cf33a. FLOW-081 already names an unstarted
regenerate change as its trigger; FLOW-083 describes Context only. The current
single testcase makes the deletion failure look like a Context failure.

**Proposed workflow scope:** a normal audited clarification of FLOW-081, retaining
its existing deletion/zone obligations and explicitly stating that claiming a
recorded regenerate phase retains existing disposable files. Suggested atomic
wording: “Start MUST delete only the target module's approved disposable files
when starting an unstarted regenerate change. When claiming an already recorded
regenerate phase, Start MUST retain its existing disposable files.” The current
candidate's passing baseline supports legitimate already_satisfied classification
for the clarification; no fail-first bypass or implementation import is proposed.

Keep the existing unstarted-deletion assertion. Isolate recorded-phase retention
under FLOW-081 and keep a distinct FLOW-083 assertion for exact regenerate
playbook selection and exclusion of ordinary impl. Preserve every existing
assertion and ID obligation through the audited change; do not manually edit
frozen evals. The intended fresh implementation behavior is first-entry deletion,
not deletion every time a recorded phase is claimed. Do not repair the saved
reconstruction or transplant its code into the product.

### RUN-027/RUN-028 fail an uncontracted commit-subject assertion

The targeted test reports only commit_recovery_test.go:76. The three subjects are
“fireweed next evals”, “fireweed next audit”, and “fireweed next impl”; its final
subject assertion requires the substring “apply”. Its earlier and later reset,
index/working/untracked retention, warning, run/phase, archive, applied-spec/code
and claim-release assertions all finish without another test error.

RUN-027 requires soft reset to the held phase base with working/index work kept;
RUN-028 requires retained P05 warnings. Neither specifies these subject strings.
FLOW-087 identifies phase commits using change-file contents rather than messages,
and the unchanged design's history section says messages are never parsed.
The test therefore rejects this reconstruction for an incidental workflow commit
subject, not a demonstrated reset/warning failure. This is a retained runner-eval
coupling issue; making fresh workflow emit “apply” would merely satisfy it.

**Proposed runner scope:** one normal independently audited eval_fix for RUN-027
and RUN-028, with no runtime/spec/interface change. Replace subject-word checks
with public observable history checks: exactly the three checked phase commits,
phase contents advancing through audit, impl and applied, and no fixture-agent
commit reachable from final HEAD. Use workflow.Commits and/or change-document
contents and exact fixture commit OIDs; the fake workers can record their own
OIDs in the existing outside fixture log. Keep all reset-before-verification,
retained index/working/untracked content, warning propagation, archived content,
applied behavior and claim-release assertions, including invocation-failure
coverage. Do not replace “apply” with another required spelling such as “impl”.

## Approval and next invocation

These are scoped proposals, not authored product/spec/eval changes. Candidate,
controller and recovery bytes, configured bounds/counters, deadline, old reports,
proof/attempt refs and unrelated worktrees remain untouched. Applying a correction
would require its normal audited lifecycle and a separately identified corrected
subject; it cannot be hidden in the setup child while claiming unchanged 79cf33a.
No controller update is needed or authorized for this investigation/repeat.

The owner's explicit request authorizes one new diagnostic workflow invocation
using exactly unchanged 79cf33a, the fixed pin, prepared inspector and original
input boundaries. It retains the known subject-text assertion and tests none of
the proposals above. Keep both observations separate from that invocation's fresh
worker prompts and input tree. The current deadline remains
2026-11-10 22:51:05 UTC, with no further extension. Preserve every isolation gate,
exact worker command/model/effort and configured attempt/timeout bounds. Retain the
new terminal evidence and stop; no additional invocation, correction, reconstructed-
code application, controller promotion or push is authorized.
