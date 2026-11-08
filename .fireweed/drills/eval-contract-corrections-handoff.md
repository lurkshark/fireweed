# Three audited eval corrections — verified and stopped

The owner requested the three eval/contract mismatch fixes from review of rejected
workflow attempt `d63c16e`, through normal audited changes preserving behavioral
coverage, with verification and a stop before another drill. The three-change
budget is complete. No drill was invoked in this sequence. WF4 remains open.

## Source and controller

Authoring began at `c4b5fd0a428d69ab5d59ca563e642d61d4c0d6f8`; owner scope and all
three change files were committed separately at `39dfb9b`. The applied phase head
is `6b027b05dcfbded9692ff0f76d40fb9a5123eee2`. Local execution used normal base
approval on main; no `--allow-unmerged`, enforcement override or counter reset.

The unchanged controller is `/Users/joel/Workspace/fireweed/.fireweed/bin/fireweed`,
source `9a206c558e8e5625c0ae55bb57c42ae8130b4f15`, SHA-256
`cc22f47c7aa607c96a13273d29d5bf4ade4f912453840f621331d73dc385f112`.
Its directory was first on PATH and GOCACHE was `/tmp/fireweed-rebuild-cache`.
No fallback or controller promotion occurred.

Each change ran through pinned `fireweed next --json`, with the configured
`gpt-6.1-sol` high-effort workers and separate restricted `gpt-6-astra`
medium-effort auditor. Both commands retained `--ephemeral --ignore-user-config`
and `-s workspace-write`. Workers ran pinned context and `verify --quick`, then
returned without commits or transitions. The owning controller ran required phase
verification, made all phase commits, archived the changes and released claims.
All three Next invocations returned exit 0, applied, with no diagnostics. Each
audit returned No findings for all listed eval_fix IDs. No findings round trip
or blocked/replacement change was used.

## Changes and retained behavior

| Change | Correction | Evals / audit / apply commits | Evals cases | Full impl cases |
| --- | --- | --- | ---: | ---: |
| [0185](../../changes/archive/0185-explicit-empty-selection-eval.md) | Accept explicit empty-selection wording including “No selected changes.” and attribute that assertion to FLOW-115. | `7678b0b` / `8ea8244` / `aee58cc` | 170 | 652 |
| [0186](../../changes/archive/0186-context-content-eval.md) | Assert context contents and matching playbook instructions instead of unspecified heading spelling. | `2c40c03` / `ba3a28d` / `61813e8` | 159 | 482 |
| [0187](../../changes/archive/0187-pr-body-contract-evals.md) | Permit absent workflow B04 Path, preserve returned diagnostic fields, and accept the compliant empty-selection wording. | `fbff84c` / `f610ddc` / `6b027b0` | 323 | 323 |

Every recorded eval and implementation case passed. Full implementation checks
covered workflow plus app/runner for 0185, runner plus app for 0186, and app for
0187, with normal boundary enforcement. Archive Provenance records the exact
commands, playbooks and phase result counts.

0185 preserves binary/large-artifact path and byte-size checks, small/exact-size
exclusions, B04 warning/body/count checks, no B04 for an empty range, valid-empty
success and invalid-range P01/cause checks. Its empty-selection predicate still
rejects blank output and unrelated “None” in another section.

0186 checks distinguishable change, retained projected-spec, interface, SYSTEM,
implementation/eval contents and phase playbook instructions. It preserves roles,
cwd/bin/request-timeout behavior, ordered results, verification/applied-state and
output checks. Audit change/spec/eval/playbook assertions and secret exclusions
remain; matching phase playbook instructions replace the audit title assumption.

0187 keeps B01–B04 warnings and artifact identification, historical endpoint and
nested-root behavior, plain body/code/message checks, JSON schema/string body,
empty stderr, diagnostic count and equality to every public returned field. A
present Path remains checked and compared; absence is permitted. Empty-success
and all invalid-option/range/config/discovery status/cause/null-data assertions
remain. No test was skipped or removed.

Only these three package files changed:

- `packages/workflow/evals/pr_body_test.go`
- `packages/runner/evals/next_test.go`
- `packages/app/evals/pr_body_cli_test.go`

Specs, Decisions, runtime source, public interfaces/models, configuration, hooks,
worker policy and inspector tooling are unchanged. No reconstruction was imported.

## Coverage and closure verification

The latest own-module implementation reports were compared with the current
SPEC IDs. Case rows include ordinary parent/subtest rows; they are not a count
of independent behaviors. Case counts remain unchanged from entry.

| Module | Passing case rows | Covered current IDs | Missing / unknown IDs | Nonpassing rows / duplicate case names |
| --- | ---: | ---: | --- | --- |
| workflow | 170 | 133 | none / none | none / none |
| runner | 159 | 67 | none / none | none / none |
| app | 323 | 155 | none / none | none / none |
| Total | 652 | 355 | none / none | none / none |

Pinned `fireweed check --json` passed with no diagnostics.
Pinned `fireweed check-commits 39dfb9b..6b027b05dcfbded9692ff0f76d40fb9a5123eee2
--json` checked exactly the nine phase commits listed above, with no diagnostics
and no squashed entries. `git diff --check c4b5fd0..6b027b0` passed. The authoring
commit precedes this explicit phase delta; prior valid history/proof prefixes
were reused. Final idle context returned U02 and idle Verify returned P02 with
“worktree holds no change”; these establish idle state, not product verification.

**Own-verification timing target remains unmet in observed runs.** Worker Go
suite output reported runner 306.478 seconds in evals and 318.719 seconds in impl,
and app 294.370 seconds in evals and 315.695 seconds in impl. The runner runs and
app impl exceed the five-minute own-Verify target even before controller overhead.
All cases passed. No timeout/budget was raised, no assertion was removed to meet
timing, and no extra benchmark or performance correction was attempted. This
limitation is retained for owner review; the scoped fixes do not claim to resolve
it. Full implementation verification completed within its existing 25-minute bound.

Final live status has an empty queue, no shared claim, and three unapproved
applied changes, against the unchanged configured cap of ten. Each new change
has three phase commits; every pre-existing phase count is unchanged. Existing
B02 warnings are unchanged: CHK-055 has 16 cases, APP-103 has 17 and APP-104 has
19 against the configured 15-case warning threshold. These pre-existing warning
sources are outside the three selected corrections; no count or limit was reset.
Derived drill drift is workflow 1, runner 26 and app 40 after the new applications;
other module drift values are unchanged. No passing-drill claim follows.

## Preservation and stop

The failed [workflow report](workflow.md) is byte-unchanged from entry. Its ref
`refs/fireweed/drills/workflow/20261108T055420.712886000` still identifies
`8859d4433b16bf721907009e58b269937b8af4fe`. Original 79cf33a, corrected ccd3d6b and
9a206c5 commit objects remain available. Prior archives/reports/proof/attempt refs
and recovery executables were not pruned or replaced. The unrelated worktree
remains at `b59e9cc15a7d6c82f6436ef72fe5c2c3ccded211`; no additional linked worktree
remains from this sequence. All three owning controller sessions terminated
successfully before protected handoff authoring, with no queued change or claim.

The three requested corrections are complete. The normal phase evidence does not
close WF4 or change the rejected drill outcome. Stop before another drill. No
further correction, policy/tooling change, replay, reconstructed-source import,
controller promotion or push is authorized. Existing deadline remains
**2026-11-10 22:51:05 UTC**; no extension. Follow current GOAL/PROGRESS for the
stopped state and obtain a new owner-selected scope before further work.
