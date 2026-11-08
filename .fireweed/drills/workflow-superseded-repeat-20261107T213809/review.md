# Superseded diagnostic repeat

> Retention update: this review describes the evidence at `d63c16e`.
> Superseded attachments are recoverable from [Git history](../README.md#historical-evidence);
> current module reports and selected outcome receipts remain in this checkout.

The owner changed the task to applying the two proposed fixes, then validating
and updating the controller, then running a corrected-subject drill. The already
started unchanged-input diagnostic invocation was rejected through the prepared
inspector while its first worker was held. No release file or worker process
record was created. The first wrapper recorded exit 125 with launched false; the
later wrapper saw the latched rejection and exited 125 without readiness/release.

The native controller exited 1 without emergency escalation and published the
retained fail report with empty eval sections. It created no reconstruction ref.
This establishes native failure publication for this pre-release rejection only,
not output drain or reporting after interrupting a running reconstruction worker.
Input and verifier were removed natively; the owned caller was removed after
evidence retention, with no claims/worktrees or owned native group remaining.
The prior full workflow report/ref remain unchanged. This invocation is not an
independent workflow proof and no product/spec/eval correction was made in it.
