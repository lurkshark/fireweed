# Acceptance and traceability

Prepared 2026-10-21. All rows start **unverified**. R01–R21 are planning labels, not fireweed spec IDs. Expand them into atomic observable specs through changes; this inventory does not replace reading the full design.

The authoritative design snapshot is [fireweed-design-v1.md](../../fireweed-design-v1.md) in this repository; its origin and checksum are recorded in [PLAN.md](PLAN.md).

At safe author boundaries map each row to actual assigned spec IDs, permanent eval names/assertions, archived changes/phase commits and results. Referencing an ID is not proof that every clause is asserted. Record exceptions and design/reference discrepancies explicitly.

## Requirement inventory

| Ref | Design section / proposed owner | Required behavior and representative evidence |
| --- | --- | --- |
| R01 | Repository layout; Distribution and init / app, assets, project | Init creates config, SYSTEM, managed AGENTS, queue/archive, protected recipe copy and optional app. Exercise no-app, invalid input, existing files and safe non-overwrite behavior. |
| R02 | Modules and zone map / contracts, project | Omitted defaults differ from explicit empty lists; unique ownership; strict zone precedence; interface/append-only survive regeneration; eval/disposable overlap rejected even for future paths; unique prefixes and acyclic existing dependencies; warn about out-of-glob specs. |
| R03 | Specs / contracts, project | Single Specs section, valid headings/normative text, review/quarantine semantics and quarantine reason; precise ID matching in name/classname across case/separators/token boundaries; removed IDs never reused; review/quarantine visibility. |
| R04 | Change files / contracts, project | Number/slug ordering including ties; unambiguous slugs; unique local handles and cross-change rejection; late IDs and all-occurrence rewriting; normalized SHA-256 bases including tags; stale/missing bases rejected; allowed set fields/immutable identity; spec/fix/maintenance/regenerate classification. |
| R05 | Specs and Change files / project, workflow | Projection includes preceding queue changes and current zones; Added/Modified full text, Removed headings only; apply Decisions with slug, remove through Removed decisions; archive-backed history/provenance; archived changes immutable. |
| R06 | Lifecycle / workflow, gitstore | Clean tree including untracked files; queue eligibility; atomic owner/worktree claim; late IDs/run/phase, start does not commit; stale-base/unapproved warnings; resume without new IDs; verify never commits; done advances/applies/archives with recovery on failure. |
| R07 | Evals step / checks, workflow | Only allowed eval/interface/new-stub/manifest/shared edits; no existing impl mutation. Every Added/Modified/fixes ID needs a failing testcase unless legitimately already_satisfied. No compile/collection/skipped/unrelated-failure loophole. Remove references to removed IDs. Test acceptance and rejection. |
| R08 | Audit step / workflow, runner | On by default/configurable; real ID and exact clause validation; findings return to evals/count attempts; untestable findings block with Questions; passing audit lists checked IDs; audit-only edits; command fallback and doctor warning. |
| R09 | Implementation step / workflow, checks | Immutable eval/interface/SPEC/protected/other-module files; only current change editable; append-only additions only. Own/transitive dependents and boundaries pass; all/direct/none setting and CI always-all. Apply metadata/specs/Decisions and archive intact. Negative tests for forbidden zones. |
| R10 | Language support / checks | JUnit testsuite/testsuites; every testcase mapped to current IDs; unknown/missing IDs and coverage rejected; skips do not cover; malformed/missing reports and hook errors fail; only actual failed/errored cases prove fail-first; reviews/quarantines respected. |
| R11 | Hooks/recipes/doctor / checks, assets | Shell placeholder quoting versus argv expansion; hook environment/module overrides; module-attributed boundary output; exact baseline handling; disposable modules require boundaries or visible reasoned opt-out; graph derived from module metadata. All four canaries detected and recovery verified, including each hook override. |
| R12 | Retries/blocking / runner | Attempts and timeouts; retries retain work/verification feedback; exhausted attempts saved as refs; block commit includes only state/Questions; checkout restored/claim released; projected dependency blocking with independent continuation; from-attempt resume; agent commits soft-reset before checking. |
| R13 | Limits / runner, gitstore | Usage patterns inspect final 20 lines; usage waits differ from failed attempts; wait budget/--for/--max; exit 3 preserves claim/edits; resume works; supported-platform process cancellation kills children. Tests use bounded durations, not real 30-minute waits. |
| R14 | Git behavior / gitstore, workflow | Locally known remote/base defaults/fallback; no product fetch/push/branch creation; linked-worktree claims; retained attempts/drill refs; release/gc; common-Git-dir scratch paths/setup/cleanup; duplicate-run conflict and commits/recovery guidance. |
| R15 | History / checks | H01 protected worker edits; H02 impl eval/interface edits; H03 append-only edits by everyone including combined merge resolutions; H04 archive mutation with deletion/revert allowed. Before/after zones, merge/rebase/squash limits and full revert to queued state. Parse state, not commit subjects. |
| R16 | Regeneration / runner, workflow | Maintenance regenerate deletes only disposable code and passes normal checks. Real-agent drill uses spec/Decisions/interfaces/evals in scratch, runs own/transitive dependents, reports contradicted Decisions/targeted gaps, retains ref, cleans up and reports drift. No automatic change authoring from gaps. |
| R17 | CLI / app, contracts, assets | Full inventory below, noninteractive argument/usage behavior, stable diagnostic codes/explain, versioned JSON envelope/diagnostics, documented exits 0/1/2/3 including empty queue; no hidden bootstrap delegation/network behavior. |
| R18 | CI/PR body / checks, app | CI runs static/history/all coverage/boundaries/doctor with documented skip; explicit ranges work. PR body uses range endpoint, covers changes, interface/stub/manifest/shared/preserved edits, review specs, audits, Decisions, quarantines, opt-outs and blocked changes. |
| R19 | Recipes / assets, app | Go/TypeScript-pnpm/Python-uv, none, local recipes/defaults/templates/canaries; accurate ready-made versus user-built glue instructions; substitutions work for real/random modules; copied recipe becomes project-owned and never auto-upgrades. |
| R20 | Versioning/playbooks/AGENTS / contracts, assets, app | Binary version-range enforcement; distinct format version with current/previous support/migration; embedded playbook hash, override based_on/stale detection; managed AGENTS refresh preserves surrounding content; explain/playbooks match checks. |
| R21 | Packaging/principles / app and author infrastructure | Single dependency-light Go binary; no runtime networking/update/telemetry; build/install docs; Linux/macOS/Windows packaging/checksums, explicit-version installer, Go install/Homebrew/release configuration. Validate without publishing or requiring a tap token. |

## Required command inventory

Exercise the new executable, beyond public function tests. Cover applicable documented options, JSON output, invalid input and exit semantics. State-changing JSON commands are also part of the bootstrap reference surface; preserve them where consistent with the design.

```text
help, version
init --recipe <go|typescript-pnpm|python-uv|none|local-path> --agent <agent> [--audit-agent <agent>] [--no-app]
doctor
new-module <name> --prefix <P>
new-change <module> <slug>
stamp <change>
list, status, modules
show <module> [ID] [--projected | --history]
context <change>
explain <code>, playbook <name>
check
coverage <module> [--with-change <change>]
check-commits <range>
ci [<range>] [--no-doctor]
pr-body <range>
start <change> [--from-attempt <n>]
verify [<change>], done [<change>]
next [--allow-unmerged] [--for <duration>]
run [--max N] [--for <duration>] [--allow-unmerged]
release <change>, commits <change>
drill <module>
gc, migrate, agents-md --write
```

## Completion gates

1. **Traceability:** all design requirements mapped to assigned specs, assertion-level black-box evidence or honest review requirements, and archived changes. All R rows verified. Do not evade acceptance through review/quarantine/already_satisfied/maintenance.
2. **Process:** bootstrap-managed fail-first/audit/impl/archive commits for behavioral work; immutable evals/interfaces during impl; separate protected setup/author commits; no pending/blocked rebuild changes or active claims.
3. **Go:** fresh checkout builds the new binary and passes `go test ./...`, `go test -race ./...`, `go vet ./...`. Permanent tests run with local fixtures/fake agents and without bootstrap/agent authentication. Cross-compilation is not platform runtime testing; report limits accurately.
4. **Bootstrap checks:** controller `check`, `doctor`, `check-commits <setup-sha>..HEAD`, and `ci <setup-sha>..HEAD` with doctor enabled. Record the exact nonempty SHA range and per-module coverage. Resolve errors and account for warnings. An empty default range on main proves no history.
5. **Product checks:** NEW binary in fresh repositories performs init/doctor, author/stamp/check, documented base approval, start/failing-evals/audit/impl/done, coverage/history/PR/CI. Exercise fix, maintenance, blocked recovery and regenerate. Normal unattended approval rejection is tested without --allow-unmerged. Pin subject paths so bootstrap cannot be used accidentally.
6. **Reference comparison:** separate equivalent fixture runs, normalized stable comparisons, and every intentional difference recorded with design basis. Resolve unexplained mismatches; reference bugs do not define acceptance.
7. **Real regeneration:** successful bootstrap-driven drills for every rebuilt module declaring disposable code, using a working agent reconstructing from specs/Decisions/interfaces/evals without consulting old implementation/history/bootstrap source as an oracle. Run dependency-first; repeat only when relevant changes stale evidence. Own and transitive dependent evals pass without preserved-contract edits. Commit reviewed reports/retain refs. Fix agreed gaps through changes; never automatically approve all agent suggestions. Missing agent access blocks this gate; a fake agent cannot satisfy it.
8. **Self-hosting:** in an isolated clone/fixture, the NEW binary governs a small real spec change and a real drill, with checked fail-first/audit/impl commits, archive/provenance, CI and scratch cleanup. Keep the original repository governed by the pinned bootstrap throughout acceptance.
9. **Delivery:** README build/test/use instructions, locally validated distribution/installer checksum behavior, clean committed target tree, restartable evidence record and independent product executable. No product runtime read/import/command requires bootstrap. Report exact checks/commits/reports and limitations. Do not push, publish or tag.

## Evidence record shape

Create execution evidence when work runs; never prefill predicted successes.

```text
Requirement: R07 (specific clause)
Module / actual spec ID:
Change / archive path:
Permanent eval names and asserted edge cases:
Fail-first command, testcase IDs, and evidence location:
Evals / audit / implementation commit SHAs:
Final verification command, exit status, and evidence location:
Reference discrepancy / decision, if any:
Status: unverified | in progress | verified | blocked
```

Bulky logs may remain ignored; commit concise findings/counts and durable commit/test references. Independent audits, platform execution, real regeneration and self-hosting claims must match what actually ran.


## Execution evidence checkpoint: Verify CLI0045

Status: R17 in progress; R06–R11/R13/R16 and all remaining rows remain incompletely verified. This entry proves the Verify command slice, not the full command inventory, lifecycle completion, regeneration or self-hosting.

- Assigned specs: APP-025–APP-028; archive changes/archive/0045-verify-cli.md. Final EVALS6c13e6cdbff9e2c0f11fa672f5fe6d3636b9803e; independent passing Audit925b3e7d1355468de3df3908a202531833405b63; implementationdacb9977a60b9a1bb0b3567a1db6424004b1dfea.
- Assertions: Test_APP025_APP026_APP028VerifyExecutablePhases and Test_APP025_APP026VerifyEvidenceAndPolicies retain full public phase/wire/coverage/findings/edits/diagnostic/exit/hook behavior across formats, linked/hash variants and policies. Test_APP026_APP028VerifyFatalDiscovery and Test_APP026_APP027VerifyFatalAndArguments assert useful fatal classification, options/help and exact selectors. Test_APP027_APP028VerifyBeforeAccess plus Test_APP025_APP028VerifyCurrentMetadataReads calibrate before-I/O and approved-source selection. Test_APP028VerifyNoninteractiveStdio, Test_APP026_APP028VerifyWritersAndHookChanges and Test_APP025_APP026_APP028VerifyRetainsHookMutations assert live streams/stdin, writer failures and retained hook/file/ref/ownership effects. Test_APP028VerifyReadonlyControls and presentation/public controls calibrate repeated observers, exact envelope keys, permitted pretty/diagnostic presentation and genuine public behavior.
- Fail-first: complete659 records,476 legacy passing,183 new,104 failed records/75 terminal failure leaves and58 passing controls;0 collection errors/skips/diagnostics. Verify1025.561s and Done1030.845s; strict report/source/inventory proofs under verify-cli-presentation-evals-{verify,done}. Earlier rejected/incomplete reports are excluded.
- Implementation: all183 new cases pass730.714s; child-instrumented scoped race33 passes101.798s; whole vet/whitespace pass. Complete controller659-case Verify1274.111s and Done with boundaries pass with0 failures/errors/skips/diagnostics. Frozen170 and implementation178 source files match their manifests. The controller applies/archives/commits/releases; no manual worker commit or frozen eval/interface edit.
- Independent NEW executable: clean dacb997 build SHA78433276f3f59c24558c49de2d5eda3277aec7957c96ea0c0bb52c0ad3a7afb0; NEW check valid:true and unowned verify --json P02/exit1/null data. Permanent evals invoke separately built NEW binaries/public app.Run probes, without bootstrap delegation. No fresh whole-app race/whole-repository checkpoint or separate post-archive NEW coverage app result is claimed.
- Remaining explicit defects: D042 approved held graph publication and D043 nested owned-process cancellation. The latter is a confirmed GIT-006/GIT-010/R13 violation; all affected rows remain in progress. Complete command inventory, every completion gate, real dependency-first drills and isolated NEW self-hosting remain mandatory.


## Execution evidence checkpoint: nested-process cancellation0046

R13/R14 remain in progress; the complete goal is not achieved. Existing GIT-006/GIT-010 live-parent regressions are archived in0046 at1eafc5ca82125a57b1ab1bf0dcc4a75b464b8c23 (EVALS50cf52d, independent Audit02dd24b). Complete fail-first463 has445 legacy passing,18 new,14 failed records/8 genuine delayedwrite leaves and4 passing controls,0 errors/skips/diagnostics. New18 focus and whole Gitstore463 child-instrumented race pass. Complete controller Verify/Done each pass3957 own/dependent records and boundaries, with exact counts/classes/names except3 source-confirmed volatile fixture object suffixes and frozen source proofs. NEW clean applied build SHAe9b8a2ef4b57b3ab776f4150245c27b9b3ee8ac941ff17ddb2b7231d6da51f9e passes actual check. Crossbuilds are compile-only.

Ordinary orphaned-group-cancellation-probe confirms a separate existing-contract violation after a parent has exited: native leaf group is independent and ppid1 before cancellation, outer is still pending, both adapters return context.Canceled/-1, then the orphan writes after outer return. D044 requires a permanent fail-first/audit/implementation follow-up; passing0046 does not complete GIT-006/GIT-010/R13. Full command inventory, final whole Go/race/vet/controller CI, all acceptance rows, real dependency-first regeneration and isolated NEW self-hosting remain unfinished.


## Execution evidence checkpoint: orphan cancellation0047

R13/R14 and the full goal remain in progress. Archive0047 applies at e34f9d04d6d46226d7455b8148e8faf5ace5c54e after genuine499-case EVALS (463 legacy passing,36 new,14 failed records/8 actual delayed-write leaves,22 positive controls/ancestors), freezecc22d099 and restricted independent Audit09efd344. Native macOS36 focus, complete499 race and native Linux87 process cases pass with exact shared case names. Final16 crossbuilds, whole vet and clean NEW build/check pass; Windows and other architectures are compile-only. Complete controller implementation Verify/Done each pass3993 records/boundaries,0 failures/errors/skips/inner+outer diagnostics, exact inventories except3 calibrated fixture object suffixes, frozen179/final222 source proofs. Done raw logs honestly retain normal Go caching of unchanged suites. Actual archive/claim release/clean state inspected; no manual worker commit.

D046 is a confirmed remaining GIT-006/GIT-010/R13 violation on macOS after two ancestors have vanished: both public adapters return matching cancellation/-1 in990/995ms but their live native ppid1/session-separated descendant writes after a marker armed only after outer return. All probe helpers terminate before cleanup. This discovery is not yet a permanent fail-first regression or a fix, and0047 is not universal tree coverage. D042 held graph, full commands/lifecycle/runner, final fresh Go/race/vet/controller CI, all requirement rows, real dependency-first regeneration and isolated NEW self-hosting remain unfinished. No gate is waived.
