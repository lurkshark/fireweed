---
module: workflow
phase: audit
run: d6540bfc1bc15d3673735a4d
---

## Motivation

The NEW modules command must publish the approved inclusive graph while a local change is held, including when a real boundary hook queries FIREWEED_BIN. The existing public approval loader intentionally supports an unclaimed current-source fallback. Calling it after a separate optional claim lookup can therefore fall back to worker metadata if ownership disappears. Provide a strict optional source selection that callers can safely compose without running Verify recursively. D042 records the equivalent-fixture graph mismatch; app integration belongs in a later single-module change.

## Interface

Add LoadHeldApproval(start string) (project.Snapshot, Claim, bool, error) to the retained workflow facade. It uses the existing public Snapshot and Claim types. Preserve all existing signatures and LoadApproval behavior. Add only a new private stub file/package needed for EVALS to compile.

## Added

### FLOW-048
LoadHeldApproval MUST discover the canonical checkout and common Git directory from start, including nested ordinary and linked SHA-1/SHA-256 worktrees, and select the unique claim owned by that checkout. No owned claim MUST return a zero Snapshot, zero Claim, false and nil error without loading current project configuration, SPEC or queue documents. Valid foreign claims MUST NOT count as owned. An unavailable or malformed claim payload that cannot be safely classified, or ambiguous local ownership, MUST return zero values, false and a useful corrective error.

### FLOW-049
For a unique valid owned claim LoadHeldApproval MUST return true, the exact stored Claim including its object token, and the same approved Base snapshot with the exact started State inserted that claimed LoadApproval provides. It MUST use only the selected committed approval and state as project authority, retaining the canonical active Root and approved project configuration, modules, changes and archive. Current worker configuration, SPEC and queue edits MUST NOT replace approved metadata. The returned snapshot MUST remain before inclusive projection, so public Project before the selected change followed by ApplyChange for that change applies preceding and selected effects exactly once and excludes later effects.

### FLOW-050
LoadHeldApproval MUST bind approved source loading to the exact selected claim token. After loading and before returning success it MUST revalidate the unique locally owned selection and that exact token; disappearance, replacement or newly ambiguous local ownership observed at that validation MUST return zero values, false and a useful corrective error. It MUST NOT silently return unclaimed absence or fall back to current project source after selection. Unavailable required claim payload, Base or project-document objects, malformed started State or other selected approved content, and unsafe selected source paths MUST return zero values, false and useful corrective errors. Unavailable unrelated non-claim document bodies or unneeded ancestors MUST NOT prevent a locally complete selected result.

### FLOW-051
Successful and rejected LoadHeldApproval calls MUST remain local, bounded, noninteractive and read-only: no lazy fetch, push, branch/worktree creation, hooks, agents, bootstrap or network commands; no file, ignored-content, index, HEAD, ref, claim, attempts or drill-report mutations; no parent cwd or environment changes. Required local Git failures MUST retain useful command details rather than treating stderr as successful source data. Repeated successful results MUST have independent writable nested slices/maps and deterministic approved values.

## Modified

## Removed

## Decisions

This optional source API does not verify a worker step or validate the effective queue/graph. Callers retain public Project/ApplyChange/Check composition and must not recursively invoke Verify merely to read the graph. Existing LoadApproval retains its documented unclaimed behavior. Optional implicit selection cannot safely ignore an unavailable claim payload: its worktree identity is unknown, so it could be a second local claim. This differs from an explicitly named LoadApproval selection, which can ignore unrelated claim bodies. Absence can be established using local Git ownership metadata before current project documents are loaded.

The rebuild goal authorizes local authored execution and controller-owned phase/archive commits. Use the pinned bootstrap as controller and separately built NEW binaries as subjects; no push, publication or tags, no claimed human merge.

Final ownership validation is the successful read operation's validation point; the API cannot prevent a later independent owner update. Immutable-object caching is allowed but cannot skip that validation. Evals must not require a particular Git command/body-read strategy. Race/fault observers need independent live controls and fresh-source premises; a stub that never selects anything does not establish an ownership-disappearance race. Positive locally complete sources with missing unrelated bodies/ancestors cover the fail-closed/locality contract without treating a fixture-premise failure as fail-first. The saved core/race drafts are planning inputs only and must be compiled, calibrated and corrected during EVALS.

Object-availability measurements and their reference controls enter separate fresh public processes after object removal; fixture setup must not make a legitimate immutable cache look like a missing-data violation. Missing nonselected archive content retains the existing public loader's categorized nonfatal C05 source diagnostic. No universal error code or literal injected process exit number is added to FLOW-048–051: assertions require useful corrective errors and relevant command details. A failed optional Git probe may recover through an independently complete native reader.

The read-only and ownership observers are eval instrumentation, never product protocols. Their controls distinguish native source-file/ref reads, Git stdout, interactive batch requests, absolute native Git, compressed archive output and stripped ref names. On macOS the observer uses a test-owned copy of the actual selected developer Git engine with an ad-hoc signature so system-image loader restrictions do not hide its actual object/ref I/O; the selected Git exec path and argv are retained. External exact-token CAS controls run outside the measured process. No bootstrap, private implementation entry point or particular Git body-read command supplies source authority. Inert stub failures are ordinary source-selection failures, not claimed exercised race evidence. Empty writable maps are tested by insertion and empty slices by interleaved appends, in addition to populated nested-container mutations.

The probe wire explicitly preserves Claim.Object separately because the retained Claim JSON deliberately omits that field. Nonzero success tokens and Object-only partial-claim leaks have independent live wire controls. Native ownership controls include relative/dot/parent paths and exec/spawn/spawnp with minimal or empty environments. If an exact successful source uses an unobserved route, incomplete source observation is an eval-observer coverage failure to investigate, not evidence that the API violated ownership validation. Only complete observed source followed by missing validation establishes that assertion's premise. Restricted pre-freeze reviews corrected the draft's identified overconstraints before phase freeze; they do not replace the separate formal Audit.

Fault-recovery operation assertions run on every return path, including successful native fallback. Required Git failure details may identify any actually observed failed command, rather than presuming the first optional probe was required; independent native controls exercise distinct probe/read failures and forbidden recovery operations. Canonical ownership identity also applies through nested symlink aliases of ordinary and linked checkouts. Valid Markdown syntax alone does not establish a valid started State: missing/invalid/applied phase, missing run, and missing/mismatched module remain corrective atomic rejections, calibrated with the existing public claimed loader and state checker.

Native source-delivery evidence requires complete unique byte coverage of each tracked loose approval object, using shared range bitmaps for read/pread/readable mmap, including overlapping and out-of-order chunks. Header-only reads and unreadable mappings cannot arm the final-ownership barrier. Textual stdout markers do not establish completeness. Independent partial/readable/overlap/chunk controls validate this eval premise; it does not prescribe a product read size, body command or private entry point. Native observation transport follows exec/spawn/spawnp children independently of their application environment so Git body observation remains live.

Read-only Git commands can still execute repository-configured helpers. A real local fsmonitor hook is therefore installed and invoked twice by native Git status with optional locks disabled before each measured fixture. Its marker is cleared after fixture setup and observed on every measured return, including rejected ownership and successful native fallback paths. Disabling an unnecessary hook for a particular read is permitted; invoking it is not. No hook or Git-command reading strategy is required by this observer.

Git replacement views are not stored approval authority. Separate native replacement controls substitute the selected claim blob and Base commit while the public claimed loader and held results retain the original token, State and snapshot; preexisting replacement refs remain unchanged. No-owned current-document observation includes the caller and actual Git children, for both working files and committed HEAD document blobs. Native hash-object/show controls establish those read paths for config, SPEC and queue bodies, and public Claims independently establishes absence of a local owner before cold measurements.

## Rejected alternatives

Separate Claims plus fallback-capable LoadApproval admits a disappearance race. Loading current project metadata before checking ownership lets worker edits obstruct approved graph queries. Running Verify to find ownership recursively executes boundary hooks. Returning an already inclusive projected snapshot would risk applying the selected change twice in existing public composition.

## Audit

- FLOW-049: "the exact stored Claim including its object token, and the same approved Base snapshot with the exact started State inserted that claimed LoadApproval provides" — No held fixture installs replacement refs for the selected claim blob or Base commit. A replacement-aware reader can substitute claim State or Base metadata and pass current held cases. Calibrate ordinary Git replacement views against replacement-immune public LoadApproval and assert identical held source/token plus unchanged replacement refs.
- FLOW-048: "without loading current project configuration, SPEC or queue documents" — Caller-only file guards do not observe current-document reads inside actual native Git children. An unnecessary git hash-object --no-filters fireweed.yaml can consume current configuration before absence without changing results, files, refs or fsmonitor markers. Independently calibrate native Git current-document read observation and assert no current configuration/SPEC/queue bytes are consumed on no-owned paths.

## Provenance

- evals | 2026-10-26T06:46:09Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 2346 | failing: 86 | passing: 2260
- audit | 2026-10-26T06:54:34Z | agent: manual | fireweed: 0.4.0 | playbook: a9c982d24b476621316049b527e6c7bbc319f06b12268af150fa7d8c6df92230 | tests: 0 | failing: 0 | passing: 0
- evals | 2026-10-26T08:08:17Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 2356 | failing: 90 | passing: 2266
- audit | 2026-10-26T08:19:24Z | agent: manual | fireweed: 0.4.0 | playbook: a9c982d24b476621316049b527e6c7bbc319f06b12268af150fa7d8c6df92230 | tests: 0 | failing: 0 | passing: 0
- evals | 2026-10-26T13:50:45Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 2362 | failing: 90 | passing: 2272
- audit | 2026-10-26T13:59:14Z | agent: manual | fireweed: 0.4.0 | playbook: a9c982d24b476621316049b527e6c7bbc319f06b12268af150fa7d8c6df92230 | tests: 0 | failing: 0 | passing: 0
- evals | 2026-10-26T14:35:04Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 2362 | failing: 90 | passing: 2272
- audit | 2026-10-26T14:44:27Z | agent: manual | fireweed: 0.4.0 | playbook: a9c982d24b476621316049b527e6c7bbc319f06b12268af150fa7d8c6df92230 | tests: 0 | failing: 0 | passing: 0
- evals | 2026-10-26T15:24:59Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 2381 | failing: 103 | passing: 2278
