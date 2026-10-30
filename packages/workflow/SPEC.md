---
depends_on:
    - contracts
    - gitstore
    - assets
    - project
    - checks
id_prefix: FLOW
interface:
    - api.go
    - model/**
name: workflow
---

# workflow

## Specs

### FLOW-001
Start MUST atomically create the local claim for the selected queued change.

### FLOW-002
Verify MUST reject edits forbidden by the approved projected phase zones before running hooks, checking differences since the held phase base.

### FLOW-003
Done MUST require full verification to succeed before committing a phase transition.

### FLOW-004
Successful impl Done MUST apply Added, Modified and Removed spec edits to the module's SPEC.md.

### FLOW-005
Context MUST include the change, projected spec, interface contents, SYSTEM and matching playbook. It MUST list other relevant files by path and size, or include their contents when full is requested.

### FLOW-006
Start MUST include the approved projected project snapshot in its returned State.

### FLOW-007
Start MUST replace each local Added handle throughout the change with the next available spec ID for its module.

### FLOW-008
Start MUST assign a unique run identifier when starting an unstarted change.

### FLOW-009
Start MUST report P04 as a warning when neither the local base nor the shared base contains the change file.

### FLOW-010
Start MUST reject a worktree containing uncommitted changes.

### FLOW-011
Start MUST reject a requested change unless it is the first change returned by
Eligible. With no requested target, Start MUST start that first eligible
change. When Eligible returns no change, Start MUST reject with P01.

### FLOW-012
Start MUST return the canonical absolute worktree root in State.Snapshot.Root.

### FLOW-013
Start MUST set phase to evals when starting an unstarted spec change.

### FLOW-014
Start MUST leave HEAD unchanged.

### FLOW-015
Start MUST reject with P07 when the shared-base unapproved archive count reaches authoring.max_unapproved, whose default is 5.

### FLOW-016
Start MUST reject a competing worktree that attempts to claim the same change.

### FLOW-017
Current MUST resolve only a claim owned by the calling worktree.

### FLOW-018
Start MUST use the current HEAD as the initial phase base of its claim.

### FLOW-019
While a claim is held, Current MUST return the project snapshot approved at Start.

### FLOW-020
Verify MUST reject an unapproved alteration of the active change's author fields before running a hook.

### FLOW-021
Verify MUST reject an unapproved alteration of the active change's recorded runtime state before running a hook.

### FLOW-022
Whenever Verify runs the target module's evals it MUST require coverage of every enforced current spec.

### FLOW-023
During evals Verify for spec and fix changes MUST require genuine fail-first
evidence for every enforced change ID that is not already_satisfied.

### FLOW-024
During impl Verify MUST reject a failing current-module testcase that references a non-quarantined spec.

### FLOW-025
Full impl Verify MUST use verify_dependents with a default of direct to select dependent modules.

### FLOW-026
Full impl Verify MUST reject a failed enforced testcase from a selected dependent module.

### FLOW-027
Full impl Verify MUST require boundary verification to succeed.

### FLOW-028
Verify MUST select hooks from the held approved projected project snapshot.

### FLOW-029
Quick Verify MUST skip dependent-module evals.

### FLOW-030
Quick Verify MUST skip boundary verification.

### FLOW-031
Verify MUST leave HEAD unchanged.

### FLOW-032
Done MUST commit a successfully verified evals phase and advance it to audit when audit is enabled.

### FLOW-033
Done MUST commit a successfully verified evals phase and advance it to impl when audit.enabled is false.

### FLOW-034
Audit Done MUST discard a finding whose spec ID is absent or whose quoted clause is not an exact substring of that spec.

### FLOW-035
On re-audit, Audit Done MUST discard findings for spec IDs that had no retained finding in the preceding audit round.

### FLOW-036
Audit Done MUST require every enforced ID in the change's Added, Modified,
fixes or eval_fix entries to appear in a No findings submission's checked-ID list.

### FLOW-037
Retained testable audit findings within the round limit MUST return the change to evals.

### FLOW-038
An audit return to evals MUST increment and persist audit_rounds.

### FLOW-039
Audit Done MUST block the change when retained findings would exceed audit.max_rounds, whose default is 2.

### FLOW-040
Audit Done MUST block the change when a retained finding marks its quoted clause untestable.

### FLOW-041
Audit Done MUST require No findings with a complete checked-ID list before advancing to impl.

### FLOW-042
When audit blocks a change, Done MUST include the blocking findings in its Questions section.

### FLOW-043
Audit blocking MUST release the calling worktree's held change claim.

### FLOW-044
Successful impl Done MUST apply the change's set values to module metadata.

### FLOW-045
Successful impl Done MUST apply Decisions attributed to the change slug and remove entries quoted in Removed decisions.

### FLOW-046
Each successful Done MUST advance the phase base to its new commit.

### FLOW-047
Successful impl Done MUST release only the exact held claim upon application.

### FLOW-048
Failed Done verification MUST leave the phase unadvanced and create no commit.

### FLOW-049
Impl Done MUST require the applied project to pass static validation before committing.

### FLOW-050
Successful impl Done MUST move the active change to the archive with phase applied and commit it together with the application.

### FLOW-051
Each successful Done MUST append provenance containing phase, date, manual or configured agent, version, playbook hash and result counts.

### FLOW-052
Successful Done MUST retain verification reports and process output under the run directory.

### FLOW-053
State MUST expose the held projected snapshot so composing commands use the approved graph during hooks.

### FLOW-054
Start MUST prefix dirty-worktree and queue-eligibility rejections with P01.

### FLOW-055
Verify and Done MUST prefix unsupported-phase rejections with P01.

### FLOW-056
Start and Current MUST prefix existing claim ownership or contention failures with P02.

### FLOW-057
Verify and Done MUST prefix unapproved author, runtime or audit-state alteration errors with P03 before running hooks.

### FLOW-058
Start, Verify and Done MUST prefix stale-base rejections with P05, retaining the underlying cause.

### FLOW-059
Audit Done MUST prefix malformed or incomplete No findings submission errors with P09.

### FLOW-060
Done MUST prefix its existing blocked-change error with P06 while preserving the committed blocking state and claim release.

### FLOW-061
Current and Context MUST prefix unknown requested change lookup errors with U02.

### FLOW-062
Lifecycle operations MUST preserve assigned static, eval, zone and boundary causes when returning their existing failures, except for lifecycle stale-base classification P05.

### FLOW-063
Block MUST require the calling worktree's current claim with unchanged approved
metadata and phase. It MUST reject a changed HEAD rather than blocking a stale
phase base.

### FLOW-064
Block MUST save the held step's tracked and untracked work as a commit at the
next unused positive refs/fireweed/attempts/<slug>/<n>, preserving prior attempt
refs. BlockResult MUST identify that ref and its commit OID.

### FLOW-065
After saving the attempt, Block MUST restore the step's tracked files, untracked
files and index to the phase base, removing the step's edits from the checkout.

### FLOW-066
Block MUST commit only the held change's assigned IDs, run, current phase and
requested blocked reason over the phase base. The block commit MUST preserve
the approved author and already-committed audit content and exclude step work.

### FLOW-067
On successful blocking, Block MUST release the current claim and return the
blocked State with no held ClaimOID. The resulting working tree MUST be clean.

### FLOW-068
Eligible MUST return unblocked queued changes not held back by earlier blocked
changes, in queue order. Empty or entirely waiting queues MUST return an empty
result without an error.

### FLOW-069
Eligible MUST omit a later change targeting the same module as an earlier
blocked change.

### FLOW-070
Eligible MUST omit a later change whose module depends directly or transitively
on an earlier blocked change's module.

### FLOW-071
When a later change has a nonempty Interface section, Eligible MUST omit it if
an earlier blocked change's module depends directly or transitively on its
module. With an empty Interface, that dependency change MUST remain eligible
unless another blocking rule applies.

### FLOW-072
For each candidate, Eligible MUST apply the blocking rules to the graph
projected through that candidate, including preceding queued set edits and its
own set edits.

### FLOW-073
StartFromAttempt MUST resume a committed phase with the same clean-worktree,
eligibility and claim-ownership preconditions as Start. It MUST preserve the
existing run, phase and assigned IDs and leave HEAD unchanged.

### FLOW-074
StartFromAttempt MUST restore the saved step's modified, deleted and newly
added files from refs/fireweed/attempts/<canonical-slug>/<attempt> into the
working tree without reverting later independent changes.

### FLOW-075
While restoring saved worker-editable body sections for the resumed phase,
StartFromAttempt MUST preserve the current change's approved author fields,
runtime state and controller-written provenance.

### FLOW-076
StartFromAttempt MUST reject a non-positive attempt number or missing saved
attempt ref with P01 before claiming or writing the change.

### FLOW-077
For an unstarted change with no Added, Modified or Removed specs, fixes or
eval_fix, and regenerate not true, Start MUST set phase to impl.

### FLOW-078
For an eval_fix change, Verify MUST permit evals-phase edits only in the target
module's evals zone and the current change file. During impl it MUST permit
edits only to that change file.

### FLOW-079
During evals for an eval_fix change, Verify MUST require all target-module
evals to pass rather than fail first, with every listed eval_fix ID covered by
at least one referencing testcase.

### FLOW-080
Start MUST set phase to impl for an unstarted change with regenerate true.

### FLOW-081
Before returning from successful Start of an unstarted regenerate change,
Fireweed MUST delete only the target module's files classified as disposable
by the approved projected zone map, retaining all other files.

### FLOW-082
For a regenerate change, Start MUST emit warning P10 when the module has no
passing drill dated after its latest applied change's implementation Provenance.
A newer passing drill MUST suppress that warning.

### FLOW-083
For a held regenerate change in impl, Context MUST supply the regenerate
playbook in place of the ordinary impl playbook.

### FLOW-084
Release MUST drop only the named claim owned by the calling worktree while
preserving its working files, HEAD and recorded change state.

### FLOW-085
Release MUST reject a missing claim or a claim owned by another worktree with
P02, leaving refs unchanged.

### FLOW-086
Release MUST accept the same unambiguous change path, filename and slug
selectors as Current.

### FLOW-087
Commits MUST return the full object IDs, newest first, of commits reachable
from HEAD that change the selected change document's phase relative to their
first parent, following its pending and archive paths. Selection MUST use file
contents rather than commit messages.

### FLOW-088
Commits MUST resolve pending or archived changes by repository-relative path,
filename or unambiguous slug. It MUST reject missing or ambiguous selectors
with U02.

### FLOW-089
Abandon MUST undo the unapplied change's phase-commit effects identified by Commits and delete its pending change file in one commit, returning that commit's full object ID.

### FLOW-090
Abandon MUST reject a dirty working tree with P01 before changing files, HEAD or refs.

### FLOW-091
Abandon MUST allow an unclaimed change or a claim owned by the calling worktree.
It MUST reject a foreign claim with P02 before changing files, HEAD or refs.
On success it MUST release the selected owned claim.

### FLOW-092
Abandon MUST reject an applied archived change with P01.
After git revert returns that change to the pending queue, Abandon MUST allow dropping it under the same preconditions.

### FLOW-093
Abandon MUST retain saved attempt refs unchanged.

### FLOW-094
The Abandon commit MUST not assign a new phase value.

### FLOW-095
After successful Abandon, its assigned but unapplied spec IDs MUST be available for subsequent Start allocation.

### FLOW-096
Abandon MUST accept the same unambiguous repository-relative path, filename and slug selectors as Commits, rejecting missing or ambiguous selectors with U02.

### FLOW-097
Each Verify that evaluates modules, including through Done, MUST retain a cost snapshot under its existing run/phase directory and return those module costs in Result.Costs.
The snapshot MUST identify the run, phase and UTC recording time, replacing the previous cost snapshot for that run/phase.

### FLOW-098
Each evaluated module's cost MUST report its name, summed JUnit testcase seconds and number of matching testcases per spec ID, counting each testcase once per matching ID.
It MUST report eval and disposable line counts for regular files selected by the held approved projected zone map, counting an unterminated final line and excluding higher-priority zones.

### FLOW-099
Verify and Done MUST return warnings B01 for total evaluated-module seconds above budgets.verify_time (default15m), B02 for each spec's testcase count above budgets.cases_per_spec (default15), and B03 for each module's eval/disposable line ratio above budgets.eval_ratio (default4).
Warnings MUST identify measurement and limit without affecting verification success or failure.
Positive eval lines with zero disposable lines MUST exceed the ratio limit.

### FLOW-100
Costs MUST return the retained run/phase cost snapshots in recording-time order with their module values and budget diagnostics, or an empty slice when none exist.
The retained values MUST reflect the evaluated state even if source files or budget settings later change.

### FLOW-101
Status MUST return every project.Queue entry in queue order with slug, path, module, phase, blocked reason and audit rounds.
Absent phase, reason and rounds MUST appear as queued, empty string and zero respectively.

### FLOW-102
For an unblocked queued change targeting an earlier blocked change's module, Status MUST name that earlier change in WaitingOn with WaitingRule same_module.

### FLOW-103
For an unblocked queued change whose module depends directly or transitively on an earlier blocked change's module, Status MUST name that earlier change with WaitingRule depends_on_blocked.

### FLOW-104
For an unblocked queued change with a nonempty Interface section, Status MUST name an earlier blocked change with WaitingRule interface_dependency when that blocked module depends directly or transitively on the candidate's module.
An empty Interface MUST not trigger this rule.

### FLOW-105
Status MUST determine waiting reasons using the graph projected through each candidate, including preceding and candidate set edits.
Blocked entries and unblocked entries unaffected by the blocking rules MUST have empty WaitingOn and WaitingRule.

### FLOW-106
Status MUST expose the discovered modules with their declared identities, dependency graph and resolved zones.

### FLOW-107
Status MUST list every currently quarantined spec with its module, ID and quarantine reason.

### FLOW-108
Status MUST list every module with boundaries_opt_out and its authored reason.

### FLOW-109
Status MUST report each module's count of applied archived changes whose implementation Provenance date is later than its latest drill report date.
With no drill report, it MUST count all that module's applied archived changes, including zero when there are none.

### FLOW-110
Status MUST count applied changes archived on HEAD whose change file is absent from the locally known configured remote/base ref, falling back to the local base when the remote ref is absent.
The base and remote MUST use the existing configured/default base-ref resolution.

### FLOW-111
Status MUST warn with P08 when the three or more most recently applied changes consecutively target one module, identifying the module and consecutive count.

### FLOW-112
Status MUST report the number of phase-changing commits reachable from HEAD for each applied archived change, matching Commits for that change.

### FLOW-113
Status MUST expose the latest retained ModuleCost for every module with recorded costs, together with the saved budget diagnostics associated with those selected measurements.
Modules without a cost record MUST not receive a fabricated measurement.

### FLOW-114
PRBody MUST render a Markdown description for an explicit Git commit range, identifying its resolved endpoints.
Project configuration, module/spec metadata and change documents MUST come from the range endpoint rather than later HEAD or working-tree edits.

### FLOW-115
PRBody MUST identify each pending or archived change present at the endpoint whose document identity was touched in the range, including its path, module and change kind.
An empty selection MUST be reported explicitly.

### FLOW-116
For every selected change, PRBody MUST include its recorded eval result counts from Provenance and identify its eval-file edits.

### FLOW-117
PRBody MUST identify interface edits, new disposable stubs added in evals, module manifest/other-file edits and worker_writable edits associated with each selected change.

### FLOW-118
PRBody MUST identify preserved-file edits associated with selected changes.

### FLOW-119
PRBody MUST list every review-tagged Added or Modified spec of each selected change with its ID and normative text.
It MUST not list an unrelated pre-existing review spec as changed.

### FLOW-120
PRBody MUST include each selected change's audit_rounds and recorded audit findings, including findings in earlier phase commits that were corrected before the endpoint.

### FLOW-121
PRBody MUST include the selected changes' added and removed Decisions.

### FLOW-122
PRBody MUST list all quarantined endpoint specs with module, ID and reason, every endpoint boundary opt-out with module and reason, and every blocked pending endpoint change with its reason.

### FLOW-123
PRBody MUST show the latest retained measurements and associated B01–B03 diagnostics for modules represented at the endpoint, without fabricating missing measurements.
Saved local cost snapshots MUST remain the measurement source rather than later file counts or recalculated budgets.

### FLOW-124
PRBody MUST list every binary file and every file larger than 256 KB newly added to a disposable zone between the range endpoints, identifying path and byte size.
Each such file MUST produce a B04 warning in Diagnostics and the Markdown body; ordinary smaller text additions MUST not produce B04.

### FLOW-125
An empty or unresolvable range MUST return an error with P01 and the rejected range or resolution cause.
A valid empty range MUST return a report rather than an error.

## Decisions

[0012-manual-lifecycle] Provenance lines begin `- <phase>` and use `playbook=<full SHA-256>` plus `tests=N failing=N passing=N` for the reports verified by that Done call.

[0119-verification-costs] Cost snapshots use .fireweed/runs/<run>/<phase>/costs.json beside existing module
reports; repeated verification in that phase replaces its snapshot, as it does
the module reports. Each regular file's line count is newline count plus one
when nonempty content lacks a trailing newline.

[0120-status-report] When several earlier changes hold one entry, report the first in queue order,
with same-module, dependent and interface-dependency precedence. Applied order
uses implementation Provenance dates, newest first with archive path as the
tie-breaker; a drill report date counts regardless of result. Budget diagnostics
from selected snapshots are returned once each, rather than recomputed from
later settings or source edits. PhaseCommits keys are archived change filenames
without .md.

[0123-pr-body-report] Selected changes are identities touched by range commits and present at its end;
removed change documents are absent from that endpoint report. Review context for
a selected change includes its phase commits reachable at the endpoint, so a
range covering only its final phase retains earlier eval/stub/audit information.
Binary follows Git's text/binary classification; 256 KB is262144 bytes, with
strictly larger text files warned. File categories use that phase's declared
metadata; endpoint exceptions use the endpoint's declarations.
Cost measurements use a Cost/Metric Markdown section with one row per measured
module: module, seconds, cases per spec (`ID: count`), eval lines, disposable lines.
