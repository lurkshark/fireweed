# 0122 commit-denial checkpoint and scoped recovery proposal

The approved 0117–0121 gate passed and origin/main is7c5e90e. Local authored
0122-status-command is incomplete: no implementation, independent audit or
actual NEW status proof has run. Count remains0 since review; no new push.

0122 run20261030T163917-0fc34507bd1994f10aea0b08 assigned APP-118–126.
Author HEAD0b3b1927cd6534b9bcc4bc6c020d25f87f35c984 remains unchanged.
The first native evals Done rejected unchanged APP-001 because a newly added
convention extension asserted status success before implementation. A fresh
fixed Sol/high worker removed only that new success assertion; original
assertions and new hook/agent/fetch/push checks remain. APP-014's obsolete
status-pending example was retargeted to pending pr-body, preserving assertions.
All three dedicated status cases genuinely fail at unsupported status; APP-001
passes. Other pre-freeze refinements retain nonnull slices, optional diagnostic
paths, warning stderr and complete resolved-zone coverage.

The replacement native eval verification passed253 cases:3 failing/250 passing,
then the controller changed phase to audit and appended its eval provenance.
Its subsequent git add -A failed creating .git/index.lock: Operation not permitted.
No phase commit exists. Both own checks were below5m. The coordinator should
have requested write-capable execution for Done. This is a sandbox commit denial,
not an automatic approval-review rejection or an assertion audit finding.

The claimed public document compared with the current one differs only by
phase evals→audit and the appended eval provenance. Public inputs are preserved
under /tmp/fireweed-0122-commit-recovery: both documents, their exact textual
diff and the two eval files. No implementation was saved or discarded.
Owned claim was released by native controller release (exit0). No phase/ref/
counter rewrite, commit retry or audit dispatch followed the failed persistence.
Current pending file remains audit, round0, same run/IDs; two eval files remain
uncommitted. No valid fail-first commit or independent audit is claimed.

SYSTEM.md reserves phase transitions/commits to the controller; the current
approval explicitly grants no phase exception. Proposed new one-time exception:

1. Restore ONLY changes/0122-status-command.md to the exact saved claim-document-
   evals.md. This reverses only the uncommitted phase transition and provenance;
   it preserves the exact run, IDs, round0 and all eval assertions. No eval/code,
   claim/ref/history/bound/count edits are included.
2. Invoke native start to resume the same run/IDs. If native resume refuses,
   stop and report it; do not manufacture a claim or reset other state.
3. Run context and native full Done with write-capable execution so the
   controller can commit normally. Inspect actual phase/run/rounds/commit.
4. Complete the isolated fixed Sol/high assertion audit and fresh implementation
   worker, native full verification, clean build and actual NEW status proof.
5. Record results and stop with the caller local at count1. No push, drill,
   additional change, weakened assertions or new bounds is proposed.

This approval would authorize exactly one document restoration in step1, not a
standing phase-write exception. It does not assert the native resume will pass.

Evidence: /tmp/fireweed-0122-{evals-done,evals-done2,convention-worker,
release-after-denial}.log; /tmp/fireweed-0122-evals-convention-focused.json;
/tmp/fireweed-0122-commit-recovery/native-document-diff.txt.

## Approved recovery result — stopped at native resume

Owner answered "Approve scoped recovery". The one-time document restoration
was applied exactly to the saved claimed evals document. Both eval files were
checked against their saved public inputs and remain unchanged. Same run
20261030T163917-0fc34507bd1994f10aea0b08, APP-118–126, audit_rounds0.
The uncommitted audit transition/provenance is removed; phase is now evals.
This exception is consumed and authorizes no further manual phase write.

Native start, invoked with write-capable execution, returned exit1:
`P01: working tree must be clean`. No claim was created, no phase commit exists,
and all two eval edits plus the same-run evals document remain uncommitted.
The approved proposal explicitly says to stop if native resume refuses; work
stopped there. No further Done, audit, implementation, push, drill, claim/ref
write or discarded input followed. Resume log:
`/tmp/fireweed-0122-approved-resume.log`.

The next owner decision must address the dirty uncommitted public inputs before
native start. The existing approval does not authorize discarding/re-authoring
them, manually manufacturing a claim or committing a worker phase. An isolated
CLI-only diagnosis could establish a supported recovery path without touching
main; it has not been run or authorized at this checkpoint. Full goal incomplete.

## Approved native cancellation — replacement pending

Owner approved /tmp/fireweed-0122-native-reauthor-checkpoint.md. The second
one-time exception restored only the uncommitted change to its committed author
version; only the two named eval edits were restored/removed after byte-for-byte
preservation. This exception is consumed. Native abandon succeeded as
000984d76e36e187606117a41a88f6d471d40bbb, with a clean, unclaimed tree.
No implementation existed or was saved/discarded/reused. The old run never had
a valid phase commit or audit, and its missing-persistence history remains here.

The approved replacement uses the identical public contract, fresh native run
and normal allocation, preserved public assertions with a fresh eval worker,
independent isolated audit, fresh impl worker/full checks and actual NEW proof.
Count remains0 until that caller applies. Stop with the caller local at count1;
no push, drill, altered bounds/counts or further phase exception is authorized.

## Resolution

The approved replacement completed and archived at74e2818 after normal fail-first,
one persisted APP-121 finding/scoped correction/re-audit PASS and253 final passes.
Actual clean-built NEW status fixture and main-checkout read both pass. Full
phase/proof evidence is in [status CLI](status-cli-0122.md). Both exceptions are
consumed; no further recovery is pending. Caller stays local at count1 as approved.
