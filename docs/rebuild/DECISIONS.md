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


## D014 — hook template text versus substituted values

RunHook expands every recognized placeholder in the original authored command and never rescans inserted values. The first0020 fixture used a literal {report} segment in source/marker paths embedded into that original command, so correct expansion changed those literal paths and invalidated the helper protocol. This is a fixture error; changing production expansion to ignore quoted template tokens would violate the authored contract and reference quoting model. Fully revert the incomplete phase cycle, preserve the attempt, and restart with ordinary inline paths plus adversarial paths supplied as substitution values. A direct literal-argv helper positive control establishes fixture validity independently. No applied spec or archived change is rewritten.

## D015 — boundary suppressions require completed execution and readable policy

Ordinary source inspection of pinned bootstrap internal/fireweed/coverage.go, Boundaries, shows it filters output lines and checks a nonzero exit only when no module was recognized; it never checks CommandResult.Err. A checker cancelled after a recognized baselined or opted-out line can therefore return no diagnostics. The optional baseline read also discards every os.ReadFile error. The design requires a working boundary hook and ignores exact existing violation lines, not incomplete execution or unreadable policy. New CHK boundary specs will always report fatal launch/configuration/signal/context errors regardless of recognized output, allow ordinary completed nonzero exits explained entirely by intentionally suppressed module lines, and fail closed for baseline errors other than absence, including escaping symlinks. These are explicit source-derived decisions, not yet a runtime comparison or implemented evidence. Permanent fail-first evals must assert them before implementation.


D015 evidence update at52dd898: archived0022 implements CHK-009–012 with108 permanent checks cases. A scratch reference maintenance verify returned0/diagnostics[] after its boundary checker printed an exact baselined module violation and killed itself with SIGTERM. A separately compiled NEW public Boundaries probe on that same fixture returned C15, retained exit-1 and full checker output. Evidence .fireweed/evidence/boundary-hooks-reference-cancelled-baseline.json; the scratch claim was released. This demonstrates the intentional difference at the public boundary API; product verify CLI composition remains future work. All full Go/race/vet/bootstrap CI gates at this endpoint pass.


## D016 — immutable explicit metadata key presence

Ordinary inspection of bootstrap internal/fireweed/lifecycle.go verifyState and model.go ChangeMeta shows it compares JSON of typed metadata with omitempty fields. A separate scratch maintenance change starts without fixes, then a worker adds fixes: []; bootstrap verify returns0/diagnostics[]. Evidence .fireweed/evidence/approved-state-reference-empty-key.json records approved/current text and released scratch claim. The authored workflow state guard compares parsed metadata key presence as well as nested values, excluding only blocked, so an explicit empty immutable key remains a worker edit. This conservative permission detail follows approved-intent immutability and is explicit in0023's specs/Decisions. It is an intentional comparison difference, not a relaxation of enforcement; NEW passing evidence remains pending after the fixture recovery below.

## D017 — canonical temporary current-directory fixtures

The initial0023 eval/audit cycle passed assertion review, but implementation verification exposed a macOS fixture alias error: t.TempDir returned /var/folders/... and os.Getwd returned the same directory as /private/var/folders/.... The safety test compared the raw path string and therefore failed even for a CheckState call that preserved the actual working directory. A separately compiled public-API probe proves before==after, diagnostics[], canonical==after and literal_matches:false; evidence approved-state-cwd-alias-probe.json.

Preserved API/evals/stub/implementation/change under /tmp/fireweed-state-recovery-c5ee487, released only our claim, restored only our uncommitted implementation to its committed stub, and fully reverted all four incomplete phase commits in977bbcf,f17b272,36cca42,94fec7a. No workflow spec was applied and no frozen eval was edited during implementation. Canonicalize the fixture path before changing directories, establish that premise with a standalone control, and restart the complete fail-first/audit/impl cycle. Keep the original failed evidence; it is not a product failure or a reason to mutate current directory in the state guard.


Recovery release correction: the first0023 release was run without the required Git metadata permissions and returned U01, which was missed in a multi-command shell. Its stale own claim persisted through the author recovery commits; the next start correctly rejected with P02. No restarted phase or implementation was accepted. Reissued release with the authorized normal-permissions parent, inspected diagnostics[] and an empty claims ref list, preserved the unclaimed corrected eval draft and restored the tree to its author baseline before starting again. This supersedes the earlier statement that the initial release had succeeded. Keep approved-state-recovery-release.json, approved-state-recovery-release-permitted.json and the rejected restart reports; check every state-changing command result before dependent work.


D016 evidence update at15bfe036: corrected0023 is archived and all220 workflow cases pass. A separately compiled public CheckState probe reads the exact recorded reference approved/current documents and returns P2-01 for the added fixes: [] key, while reference verify remains0/diagnostics[]. The updated approved-state-reference-empty-key.json records both actual observations. CLI lifecycle composition remains pending.


## D018 — claim updates preserve newer owners and unrelated refs

Ordinary inspection of pinned bootstrap internal/fireweed/git.go writeClaim/release shows create uses an empty expected-old argument, but later updates and deletion are unconditional. New FLOW-008–011 storage retains exact five-field wire compatibility while requiring full object tokens for refresh/release, owner matching for refresh, atomic expected-old updates, and --no-deref plus explicit symbolic-ref rejection. This prevents a stale operation or selected symbolic alias from modifying newer ownership or an unrelated target. Product metadata operations suppress native reference-transaction hooks so a claim never starts unrelated project commands; fireweed's own eval/boundary/audit enforcement remains enabled. These are design-consistent local ownership safeguards, not controller changes or a claim that reference manual lifecycle was self-hosted.

Permanent claims evals assert these races/aliases and native Git CAS arguments with real Git protocol controls. The actual two-way scratch wire experiment at a38c826 proves NEW LoadClaim reads bootstrap's record exactly and bootstrap verify accepts the NEW-created identical blob; NEW matching-token release cleans up. Evidence local-claims-reference-wire.json. No unsafe stale overwrite was needed for this compatibility experiment; the unconditional reference helper observation is explicitly source-derived.


## D019 — project-owned override containment

The design places project-owned instructions in .fireweed/playbooks and requires checked file paths; it does not spell out override symlink reads. ResolvePlaybook uses the established public SafePath containment contract before opening a selected override, permits contained regular-file links, and returns D05 for escaping/dangling/unreadable paths. This is an explicit conservative clarification; readable stale metadata remains advisory and does not discard author instructions.

A separate scratch fixture at9017aa5 demonstrates the difference: bootstrap playbook evals --json follows a selected symlink to harmless outside fixture instructions and returns them; a separately compiled NEW public ResolvePlaybook probe on the same fixture returns zero data and D05. Contained exact text and full override hashes agree. Native permanent evals independently prove rejection before attempted content reads, including canonical missing outside endpoints whose reads would return ENOENT. Report .fireweed/evidence/project-playbooks-reference-containment.json. This is reference CLI versus NEW API evidence; NEW override CLI/context/doctor integration is still pending. Independently authored embedded instructions have different embedded hashes by construction; no hash normalization or complete JSON equivalence is claimed.

Audit caught two successive missing-endpoint probe gaps before implementation: the guard first omitted the fully resolved ENOENT path, then used a raw /var alias instead of its canonical /private/var parent. Both were corrected through normal eval/audit retries, with direct ignored-error controls and the selected-alias control retained. No frozen eval was edited in implementation and no incomplete implementation cycle was accepted.


## D020 — complete projected worker context

The design describes context as everything a worker needs for its step and makes queued projection the contract for future work. FLOW-016–019 compose that contract from an approved captured snapshot: the target and its direct dependencies include preceding queue effects, while exact current source text remains repair feedback rather than authority for permission metadata. Required selected inputs fail atomically instead of silently delivering an incomplete prompt. Optional missing SYSTEM.md remains empty; a missing current change may recover only from its valid matching owned claim. Native Git fsmonitor hooks are suppressed during file enumeration, so a read command cannot launch unrelated project code. Fireweed's own enforcement remains enabled.

Pinned reference internal/fireweed/context.go projects the selected target but obtains dependency modules with p.Module, and skips selected file reads whose os.ReadFile fails. Ordinary separate fixtures and a compiled NEW public BuildContext probe using the implementation archived at8224693 demonstrate both differences. With an earlier beta change adding BET-002 and moving its interface to public/**, reference context still reports only BET-001 and units/b/api.txt; NEW reports BET-001/BET-002 and units/b/public/new.txt. Replacing the required alpha interface file with a directory leaves reference context successful with that file omitted; NEW returns corrective C01 and a zero WorkerContext. The API probe chooses exit2 for library errors; this is not a NEW context CLI exit claim.

The normal equivalent fixtures agree on exact change/spec/system text, selected file paths/content/order, identical project override text/hash, module identity/specs, direct dependency names and canonical common-Git directory after root-only normalization. Report .fireweed/evidence/worker-context-reference.json and scripts /tmp/fireweed-context-reference-{compare,projection}.py; fixture /tmp/fireweed-context-reference-h0daknfd. This is a narrow reference CLI versus NEW API comparison before regeneration, not full JSON/schema equivalence, self-hosting or a real drill. Public permanent tests assert complete projection, safe reads, error details, no partial results and phase/file/metadata/claim isolation. Executable context/lifecycle composition still owns restoring approved snapshots across processes and remains future work.


## D021 — separate process streams before immutable baselines

Before any baseline-snapshot phase commit, its accepted fail-first run produced362 project records with128 genuine expected failures and all229 existing cases/five controls passing (baseline-snapshot-fail-first.json). Inspection exposed a prerequisite: gitstore.Run deliberately combines stdout/stderr, which is unsuitable for byte-exact Git tree/blob protocols. Preserve that API and add separate-stream Capture in gitstore before project LoadAt, rather than introducing a shell/temp-file workaround or a second lifecycle adapter in project.

Verified all six draft files against SHA-256 manifests under /tmp/fireweed-baseline-capture-dependency-sbk5yxrt, released only the owned claim with diagnostics[], and confirmed no remaining claim refs. No eval/audit/impl phase commit existed. Restored only our tracked draft edits and removed only manifest-verified untracked files. Reorder the still-unstarted authored baseline change from0031 to0033 after new0032-process-capture. Source metadata remains local new: handles until a fresh start; saved started metadata must not be restored. Reuse saved API/evals/stub only after the prerequisite is archived, then rerun fail-first and independent audit. No frozen eval or bootstrap source was changed.
