# Author decisions

## D001 — Initial graph and interfaces

Adopt PLAN.md eight-module DAG. api.go is the public facade, evals/ permanent black-box tests, internal/ disposable implementation. Share public values through contracts to avoid parent/internal import cycles. The initial app facade forwards to a compiling disposable placeholder; no product behavior is implemented in setup. Use the YAML v3 parser as the sole initial dependency; its checksum file is dependency metadata, not copied implementation.

## D002 — Local authority and controller

Work in the existing clean main checkout. The goal authorizes local changes and commits, including manual start without claiming a human merge. Product approval tests will use an independent base ref. Do not change product enforcement. Keep the bootstrap source clean and all generated caches/build artifacts outside its checkout.

## D003 — Working real-agent CLI

Use explicit gpt-6-sol only in this repository worker/audit argv after the configured default returned unsupported-model errors and a cached supported-model probe succeeded. Keep CLI sessions ephemeral. Do not alter global user configuration. Fresh independent auditing contexts are used throughout; identical model choice is disclosed rather than claimed to be a different-model review.

## D004 — Retained data types without import cycles

The upcoming contracts configuration change will add a retained `model/*.go` interface subpackage with public aliases through api.go. Disposable contracts implementation can import the retained types without importing its public parent; dependent modules still import only the public contracts facade. Update interface zones through that authored change, not a protected setup shortcut.

## D005 — Vet-only facade repair between claims

Moving Spec to the retained model package made an existing unkeyed facade conversion fail vet. With no active claim, an author commit keys that exact literal; it changes neither the public contract nor executable behavior and does not edit evals, specs or implementation. This is retained entry/facade glue repair, not an unverified behavioral change. The prior evidence claim that vet passed is corrected; rerun vet before recording success.

## D006 — Assets runtime-isolation evidence

Native sandbox policies must permit the OS loader to start. The asset eval therefore combines source/network denial with a test-only Darwin file/socket interposer armed after startup; positive controls prove ignored errors cannot hide attempts. Its cc/dylib dependency belongs only to a temporary eval fixture, never the product. Other-platform isolation is explicitly unverified; portable content checks still run. Keep this platform/test prerequisite visible in final delivery evidence and README testing instructions.

## D007 — Embedded CLI access before init

Design playbooks/explanations are embedded binary protocol data. The first app slice makes help/version/explain and embedded playbook lookup usable without a project. The reference cli.go resolves explain before Load, but playbook after Load and therefore fails before init. Allowing embedded playbooks before init supports the design's setup/authoring flow. Later project integration must honor project-owned overrides and binary pins when a project is present; this first slice does not claim that integration. The reference error is an intentional comparison difference, not silently reproduced.

## D008 — correct discovery eval fixtures through a full revert

The first discovery-validation eval/audit run (9191c67 through d9fe6ed) passed assertion review, but implementation exposed three invalid fixture premises. CON-013 rejects unknown module metadata; the retention fixture's `custom` key was invalid. The retained document-grammar Decision requires filenames with at least four digits, so `2-alpha.md` never reaches numeric ordering. CON-005's Decision treats adjacent stars within a segment as ordinary wildcards, so `bad/**suffix` is valid. These are eval mistakes, not product behavior to reproduce or reasons to weaken contracts.

Preserved the implementation and eval drafts locally, released the claim, and fully reverted all eight phase commits in reverse order (b3de999 through 0f4ffb1), restoring the queued authored change and unimplemented project facade. No frozen eval/interface was edited during implementation. Correct fixtures use known metadata, valid variable-width numbers (`00010` versus `0002`) and genuinely invalid repeated separators. Restart the complete fail-first/audit/impl cycle. The earlier commits/audits remain visible as superseded evidence; they do not prove this change complete.

## D009 — local handles retain per-change scope in projection

Design Change files says handles are unique within a change and prohibits referencing another change's handle. A separate temporary reference fixture (/tmp/fireweed-projection-reference-qegy0yqr) with two valid queue files each declaring `new:local` made bootstrap `show alpha --projected --json` exit 2/U01 `spec new:local already exists`. The new pure Project API permits those two independent declarations and retains both placeholder items until real IDs are allocated on start. It still rejects a reference satisfied only by a different change. PRJ-007/PRJ-009 and Test_PRJ007_PRJ009LocalHandleScopes lock down this design-based difference. Preview documents may repeat placeholder headings; they are never written as current SPEC.md before allocation. No bootstrap source was modified.

## D010 — archive human prose is historical context

After 0010 was archived, the NEW binary's live `check` returned C05 on that immutable archive because a worker Decision explains the literal `new:local` example after allocation. Queued documents must validate every live handle reference, including human/meta sections, so cross-change references remain forbidden. Archived human sections instead record completed history; their illustrative handle text is not an unresolved binding. Load will validate archive metadata and edit sections through the public change grammar using a validation-only view of human bodies, then retain the complete original document. Archive edit bodies and metadata still validate live references. No archive edit or deletion is needed.

This distinction was clarified before queue-validation evals were committed: saved uncommitted work under /tmp/fireweed-queue-author-recovery, released the claim, restored only our own started state/stubs/evals to the authored baseline, and reauthored/stamped the change between claims. Restart its full eval/audit/impl cycle. This fixes actual live repository behavior rather than weakening archive immutability.


## D011 — Local Added handles in already_satisfied

Design Change files defines already_satisfied for new or modified specs while new IDs are assigned only on start. Initial CON-016 incorrectly required real IDs for this field, making the authored new-spec case impossible. Change0013 corrects this to accept locally declared Added handles until allocation; fixes and bases remain real-ID-only. Separate fixture confirms pinned bootstrap already accepts this authored form, and corrected NEW product check agrees. No fail-first policy is relaxed: explicit justified already_satisfied remains required and workflow enforcement is still pending.

## D012 — History uses implementation chronology

Design Specs requires archived spec history newest first; Change files says numbers order the queue and need not be unique. An otherwise equivalent scratch fixture has two ALP-001 modifications: newer-low at 2026-10-22T12:00:00Z and older-high at 11:00:00Z. Pinned bootstrap show --history puts older-high first. Evidence .fireweed/evidence/show-history-reference-order.json retains exact outputs and fixture path. The new executable will use the final parseable impl provenance RFC3339 timestamp descending, with dated rows before undated legacy rows and numeric/slug/path descending ties/fallback. This is a deliberate design-derived difference, with a total order when dates are absent, rather than reproducing the reference's chronology bug. Permanent APP history evals must assert it. Raw provenance remains unchanged, and no command consults the bootstrap at runtime.


## D013 — trustworthy JUnit structure and complete missing evidence diagnostics

Ordinary reference inspection of bootstrap internal/fireweed/coverage.go at pinned583b3c5 found parseJUnit accepts a token stream after its first valid root, and Coverage continues after C13 before checking fail-first. The new CHK-001 contract requires one well-formed XML document, rejecting multiple roots and text outside it atomically. CHK-003 reports both C13 and P1-02 when an enforced failFirst spec has neither executed coverage nor an already exemption, rather than hiding the unmet failing-before requirement behind C13. The design requires real executed coverage and failing-before evidence separately; neither malformed report fragments nor another missing condition supplies them. Permanent public tests exercise both choices. Other C11–C14/P1-02 meanings and actual failed/errored/skipped evidence remain consistent with the reference. This is a deliberate diagnostic detail, not a weaker exemption.
