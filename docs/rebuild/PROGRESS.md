# Rebuild handoff

**Stopped at CK3.** The owner-authorized drill on the validated FLOW-062
candidate exhausted its two configured attempts. Both workers passed live
input and verifier isolation checks. Verification reported PackageFailure
and missing coverage across checks, app, runner and workflow. A compile-only
diagnostic on the final proof found an undefined variable and an unused
import in reconstructed checks. The [reviewed report](../../.fireweed/drills/checks.md)
and `refs/fireweed/drills/checks/20261106T155422.780454000` retain the result.
CK4 remains open. A further invocation or correction needs separate owner
authorization.

The approved FLOW-062 fix
`0175-preserve-x02-boundary-cause` is applied and archived. Fail-first,
independent audit and full implementation verification passed. The clean
candidate passed the affected CK2 gates, including both boundary-failure
fixtures through the candidate CLI and explicit-range CI with hook doctor
enabled. The [reviewed checks report](../../.fireweed/drills/checks.md)
records the exact validation.

## Prior CK3 workflow failure

The previously authorized proof failed. The rehearsed
gate checked live isolation and released the first worker. The second
configured launch followed review of the first trace and verifier feedback
and a fresh live isolation check. Final checks, app and runner evals passed.
Workflow FLOW-062 failed because Verify duplicated the hook-launch detail.
The [reviewed checks report](../../.fireweed/drills/checks.md) retains the
native result and review. The rebuilt ref is
`refs/fireweed/drills/checks/20261106T072400.155233000`.

The owner-approved recovery is
applied and archived in `0172-doctor-probe-assertions`,
`0173-boundary-override-diagnostic` and `0174-boundary-hook-cause-contract`.
The held 0172 audit recovered on its original run. The separate audits and
required implementation verification passed. The fresh recovery candidate
passed all CK2 gates, including explicit-range CI with hook doctor enabled.
The [reviewed checks report](../../.fireweed/drills/checks.md) records the
recovery and exact candidate validation.

CK4 remains open. The retained proof passed the doctor canary and global
boundary-baseline cases and retained coverage across all four modules in
that earlier invocation.
FLOW-062 was corrected on main and the candidate was validated. A passing
independent proof is still required. Another native invocation needs
separate owner authorization. Main and the validated candidate are clean,
with no active run or shared claim. The controller pin, original deadline, audit/retry
bounds, retained proofs and unrelated worktree remain in place.

## Prior CK3 evidence

The additional authorized proof exhausted its attempts and failed.
The CHK-055 clarification and scoped workflow eval_fix remain applied and archived
as `0170-doctor-baseline-scope` and `0171-boundary-cause-eval`. Their audits and
required implementation verification passed. The clean recovery candidate passed
normal/race/vet suites and actual candidate check, doctor --no-agent and explicit-
baseline check-commits/CI with hook doctor enabled. Exact identity, range and
results remain in the reviewed native checks report. No runtime implementation,
interface, hook or configuration changes were made on main, and no reconstructed
source was imported.

After recovery validation, the owner separately authorized one additional native
checks drill and CK4 review if it passed. That single invocation exhausted its
configured attempts and failed enforced checks/app/workflow cases; runner passed.
CK4's passing-proof/completion gate remains open. The reviewed native
[checks report](../../.fireweed/drills/checks.md) preserves current attempts,
cases, diagnostics, independence review and copied prior validation evidence.
Its failed rebuilt ref is `refs/fireweed/drills/checks/20261105T051932.421753000`.
The earlier failed ref `refs/fireweed/drills/checks/20261105T000910.693992000`
remains unchanged; its complete report remains in committed history.

**Owner checkpoint after that proof:** review missing_id detection, global boundary-baseline
handling and doubled X02 cause propagation, together with the CHK-057 recovery
check selection and CHK-060 diagnostic wording questions. No automatic fix or
further drill is authorized. Owner agreement precedes scoped audited recovery;
another real drill still needs separate explicit authorization. Preserve gates,
obligations, counters and the exhausted attempts. Native input/verifier cleanup
and local report/ref retention are complete; the task-owned invoking clone/gate
were removed. No active run/shared claim remains. Main and validated candidate
checkouts are clean. The pin, bounds, prior proof/attempt refs, unrelated worktree
and original deadline remain unchanged.

[ACCEPTANCE](ACCEPTANCE.md#planned-milestone--checks-contracts-and-proof-ck1ck4)
holds the ordered plan; [GOAL](GOAL.md) is its restart prompt.
Approved plan budget: **three days**, starting **2026-11-04 22:51:05 UTC**
(14:51:05 PST), with deadline **2026-11-07 22:51:05 UTC** (14:51:05 PST).
Retain this original start/deadline on resume. The completed stabilization budget
is closed. Full product completion remains open; every full-completion gate is
retained. Prior proof refs and the unrelated worktree are preserved.

**Handoff authorization (2026-11-04):** the previous pin returned U01 because
native drill was not implemented. The owner explicitly authorized building from
clean committed current source, completing all CK2 gates, then replacing the pin
with no active run or shared claim and committing the protected pin instructions
separately. Validation passed and live status showed an empty queue, no shared
claim and no active change. Approval, audit/retry bounds and the original budget
were preserved; no product or eval correction was needed.

**Controller pin:** `/Users/joel/Workspace/fireweed/.fireweed/bin/fireweed`,
version 0.1.0, source `193634686aa6a1c4eb32e69df1c85e1d5ce48bd7`, SHA-256
`23e87b00217ab1a6990352393311bd150b7d60d3a4c138b75e285f518fba117f`.
Built with Go 1.27.2 and `-trimpath` from clean source on darwin/arm64.
The 2026-11-06 handoff passed normal and race Go tests, vet, candidate check,
hook doctor, and the full explicit history check. Candidate CI with hook
doctor passed for the range since the validated FLOW-062 candidate.
The controller stays separate from mutable `bin/fireweed`.

The immediate predecessor is preserved at `.fireweed/bin/fireweed-99a1d6d`.
Earlier recovery executables remain at `.fireweed/bin/fireweed-pre-ck-handoff`
and `.fireweed/bin/fireweed-bootstrap`. Use these only for explicit
owner-directed recovery, with no automatic fallback or cross-controller
claim resume. Worker PATH starts with the current tree's `.fireweed/bin`.
Use the prescribed rebuild cache and exact agent argv in GOAL.

**Recovery authorization (2026-11-04):** the owner approved clarifying CHK-055's
baseline coverage as the throwaway canary modules used to prove hooks, and one
scoped workflow eval_fix for FLOW-062's incidental boundary-error wording. Execute
both through the normal audited lifecycle, preserving existing assertions,
coverage, approval enforcement and counters. Main's passing doctor behavior and
CHK-058/CHK-029/CON-004 remain unchanged; the failed reconstruction is not imported.
Complete affected candidate validation from clean committed recovery source,
including the original explicit history baseline and CI with hook doctor enabled.
Then stop before another drill. Any further real drill needs separate explicit
owner authorization. Do not restart the exhausted drill, reset attempts, promote
another controller or broaden scope automatically. The original deadline remains
unchanged.

Native input/verifier worktrees and registrations were removed, with no live run
or shared claim. The pin, configured bounds, prior refs and unrelated worktree
remain unchanged. The original deadline still applies; full product completion
and all remaining full-completion gates remain open.

**Push authorization (2026-11-04):** the owner's separate one-time planning push
is completed; it authorizes no future push, publication, tags or bookkeeping-ref
push.

**Push authorization (2026-11-06).** The owner explicitly approved a one-time
push of all 43 commits from `ac4bf36` through `13a2685` to `origin/main`, with
`ac4bf36` as the exclusive base.

Existing testcase-budget warnings remain outside scoped cleanup. Preserve normal
approval/audit/retry bounds; read live approval state rather than treating a
prior handoff as current approval. Evidence remains in archived Notes/Provenance
and native drill reports; no supplemental inventories or manifests.
