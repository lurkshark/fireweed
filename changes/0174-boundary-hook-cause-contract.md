---
already_satisfied:
    - CHK-043
bases:
    CHK-043: 99f4260f
module: checks
---

## Motivation

The owner approved clarifying the boundary hook failure returned to callers
such as lifecycle Verify. D006 assigns X02 to a required hook that cannot
complete, and FLOW-062 requires preserving the assigned cause. The second
independent checks drill put X02 into both the diagnostic and the returned
error, and workflow rendered X02 twice. Specify the split already exercised by
the validated product so callers can preserve the cause and add the code once.

## Interface

The existing Boundaries signature continues to return structured diagnostics
and an error.

## Added

## Modified

### CHK-043
Boundaries MUST report hook launch, cancellation or timeout failures as errors with a diagnostic whose Code is X02. The returned error MUST preserve the hook failure cause without a leading X02 prefix.

## Removed

## Decisions

## Notes

Extend the existing public Test_CHK_043_LaunchCancellationAndTimeoutStayErrors
case to assert the diagnostic and error separation for launch, cancellation and
timeout. Require an error-level X02 diagnostic and preserve the launch cause
naming the missing executable, the cancellation cause and the timeout cause,
without requiring incidental prose. Keep the opt-out and execution bound
assertions and all other existing coverage. Verify that current implementation
already satisfies the clauses. The existing FLOW-062 eval retains the exact
cause and single-prefix assertion. Any demonstrated runtime gap stops for
review under the approved scope.

Use the pinned controller, context before each step, independent restricted
audit, normal already_satisfied verification and full implementation checks.
The worker runs pinned fireweed verify --quick before returning. The owning
controller runs full phase verification and advances the lifecycle.
The auditor edits only the existing Audit body. Preserve every byte outside it.

## Audit
