# Checkpoint: changes 0065–0069

Five applications since the passing 0060–0064 checkpoint.
Previous reviewed base: `ba82fcfbfeb769b04b09bc0a8b421d583f75ef0c`.
Checked source: `cdc64ae65c088b61029d560eeaab136e66de14d5`.
Review range: `ba82fcf..cdc64ae`.
History range:
`be35552b706aa3000ff8f11306d129ff1d701b85..cdc64ae65c088b61029d560eeaab136e66de14d5`.

0065 migrates remaining existing app error labels, and 0066 corrects the observed
show regression so duplicate graph names and fixed spec IDs retain C02/C05.
0067 explains the remaining emitted codes, and 0068 corrects C02/P02 causes and
recovery guidance identified in text review. These finish the planned D006
existing-path slice, without claiming diagnostics for missing commands.

0069 begins Milestone 2 with Go, none and local recipe initialization through
NEW: configuration, optional app, copied payload/templates, managed instructions,
preservation, errors and JSON. APP-014's unimplemented-command examples now use
status/gc while retaining its assertions. One independent audit round found
missing assertions for full version-range membership and default-protected
metadata; the correction passed re-audit of APP-041/042 only. Required full
phase checks pass, with 17 initial runtime failures and 144 final app cases.
The archives and PROGRESS retain all phase hashes and actual counts.

## Drill and scope limitations

The single authorized contracts drill reports 34 own/339 dependent passes,
retains its ref and cleans scratch. Its error-prefix gap and dependent-internal
reads are recorded in [the drill review](contracts-drill.md). The native pass
does not establish strict independence acceptance. No gap fix, repeat drill or
other module drill was performed or is authorized. Milestone 2 continues under
the existing owner authorization; doctor, bounded fake-agent execution and full
acceptance remain incomplete.

## Verification and checkpoint review

NEW built with CGO disabled from the clean source above; build metadata confirms
vcs.modified=false. A denied module stat-cache write was nonfatal. Vet passed.
In `/tmp/fireweed-stage2-setup-yqr_um46`, NEW Go init, modules and check returned
exit 0, schema 1, empty stderr and no diagnostics; generated config, guidance,
recipe hook and app SPEC were inspected. No configured agent or hook was run.

Current-source `go test -race ./...` and `go vet ./...` passed. Explicit baseline
history validation passed over 397 commits with no diagnostics against the
fixed source above.
No queue or claim remains. The embedded
recipe SPEC warning is unchanged. Setup doctor/canary evidence is reused because
controller enforcement/setup did not change. This is intermediate current-checkout
evidence, not fresh-checkout or final acceptance evidence.

The owner's direct approval pushed committed history through 3897b1a during the
live 0069 eval step. That push left worker files untouched and did not reset the
checkpoint count. Further implementation waits for this checkpoint's review;
any reported issue requires human review.

**BLOCK — human review required.** Read-only subagent
`checkpoint_0065_0069` found one concrete D006 explanation issue. Init returns
C12 for invalid local recipe metadata (`packages/app/internal/entry/init.go:238`),
but explain C12 identifies only module metadata/zones and directs repair in the
module specification (`packages/assets/internal/content/content.go:53`). Init
returns U02 for unknown recipe names (`init.go:215`), but explain U02 lists only
code/playbook/module/item/change lookup causes and repairs (`content.go:61`).
D006 requires every emitted code to have an accurate cause and corrective
action. The coordinator inspected these exact paths and confirms the mismatch.

The reviewer otherwise found the batch on track: diagnostic assertions retained,
init contracts/evals meaningful, phase bounds and checkpoint count respected,
and drill/final-gate limitations reported honestly. It edited nothing and ran
no tests or drills. No correction change was authored and no further push made.
The application count stays at five until the failed checkpoint is resolved
through human review.

Proposed scoped resolution, awaiting approval: a normal assets change with
`fixes: [AST-003, AST-024]` adds assertions that C12 identifies recipe metadata
and repairs in recipe.yaml, and U02 identifies an unknown recipe and advises
selecting a supported name or readable local directory. Extend those two
explanations while retaining all existing causes, corrective guidance, unknown-code
checks and other assertions. No init behavior or public interface change is needed.

## Next

Obtain human review of the scoped C12/U02 explanation correction before authoring
it or continuing implementation/pushes. Preserve existing causes/assertions and
use the normal change/audit workflow if approved. Doctor and bounded fake-agent
execution follow resolution; contracts drill limitations and final gates remain
explicitly unverified.
