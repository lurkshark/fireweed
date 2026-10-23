---
module: workflow
phase: evals
run: db1ac372ae0406218e1ccab6
---

## Motivation

Starting a change and selecting unattended work must agree on whether its authored intent matches the base branch. A filename alone is not approval, while runtime state, late IDs and worker feedback do not revoke already approved intent. Establish the pure content comparison before composing it with local base-file reads and checked lifecycle mutations.

## Interface

Add ApprovalMatches(module contracts.Module, baseText, currentText, path string) (bool, error), with a thin api.go forwarding function and disposable implementation. Keep existing interfaces and public models unchanged. Inputs are supplied change document bytes and the target module identity/prefix; this operation does not discover a repository, select a queue entry or establish that the supplied base text came from a branch.

## Added

### FLOW-028
ApprovalMatches MUST parse currentText with the public change grammar at path using the supplied module prefix, require its target to equal module.Name, and return false with corrective C05 and useful public parser detail for invalid current input or target mismatch. Current input MUST be checked even when baseText is absent or malformed. Once current input is valid, empty, malformed or differently targeted baseText MUST return false with nil error. A valid content mismatch MUST also return false with nil error; a matching valid pair MUST return true with nil error. It MUST NOT certify static project validity, phase advancement, clean state, claim ownership, base provenance or queue eligibility.

### FLOW-029
After public Markdown normalization and handle mapping below, approval MUST compare the parsed authored Module, Set, Bases, Fixes, AlreadySatisfied and Regenerate metadata values, preamble, and ordered retained human/edit sections. Runtime Run, Phase and Blocked metadata and the exact sections Decisions, Questions, Audit, Provenance and Notes MUST be excluded. All other authored sections, including Motivation, Interface, Added, Modified, Removed, Removed decisions, Rejected alternatives and unknown names, MUST remain significant, including their names, order, tags, complete bodies and interior whitespace. Fenced headings and differently cased names MUST NOT become excluded sections. YAML key order, quoting/comments, CRLF and insignificant outer blank lines normalized by the public parser MUST NOT change approval. Omitted and explicitly empty top-level Set/Bases/Fixes/AlreadySatisfied collections MUST be equivalent; explicit empty values inside Set MUST remain significant. Ordered list values MUST NOT be silently sorted or treated as sets.

### FLOW-030
When both valid documents have the same number of Added items, every local Added handle in baseText MUST map positionally to the corresponding current Added ID or handle, and every occurrence of that declared handle in canonical metadata, preamble and retained section names/bodies MUST use that same binding before comparison. This MUST cover nested set hook strings, already_satisfied, Interface and fenced examples. Overlapping handle names MUST NOT corrupt one another, replacements MUST be simultaneous rather than cascade through another binding, and a base Added real ID MUST remain fixed. Differing Added counts or differing authored intent after this mapping MUST return false with nil error. This comparison MUST NOT allocate IDs, reserve numbers or infer approval merely from equal Added IDs or spec hashes.

### FLOW-031
ApprovalMatches MUST operate solely on supplied values outside a repository, remain deterministic/local/noninteractive, and MUST NOT inspect files, Git, claims, hooks, agents, bootstrap, network, cwd or environment, execute processes, allocate IDs or mutate files/refs/temporary storage or parent process state on success or rejection. It MUST preserve every supplied module field, nested map/slice/document and empty append backing storage after repeated true, false and error calls. Irrelevant module specs/zones/hooks/Decisions MUST NOT change its bool/error or trigger validation or execution. Ambient cwd/environment/files MUST NOT change results or error details.

## Modified

## Removed

## Decisions

This compares supplied content only. The future lifecycle/runner caller must independently obtain the exact local base-branch file through safe Git reads and use an absent/unavailable document as unapproved; this function cannot prove where text came from. Start's warning and unattended approval enforcement remain later composition. Existing LoadApproval restores claim authority and is a different operation.

Public parsing defines normalized Markdown/YAML values. Compare meaningful authored metadata rather than YAML presentation. Runtime state and the five worker-authored feedback sections are omitted; remaining section structure and prose stay approval-sensitive, rather than using normalized spec hashes as an authority check. Added declaration order determines late-handle correspondence, matching allocation's ordered binding model; current handles may remain handles before start. Neither input is changed or written back. A malformed base is unapproved rather than an operational error, while malformed current intent must be corrected before checking approval.

## Rejected alternatives

Approving every file that exists on base accepts changed unauthorized intent. Byte comparison revokes approval for runtime state, late IDs and harmless YAML formatting. Ignoring all human sections permits unauthorized Motivation/Interface/unknown author instructions. Comparing only IDs or normalized spec hashes loses tags, metadata and author text. Reading Git here would make a pure comparison inseparable from base provenance and I/O policy.

## Provenance

- evals | 2026-10-23T04:48:41Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 1392 | failing: 73 | passing: 1319
- audit | 2026-10-23T04:54:30Z | agent: manual | fireweed: 0.4.0 | playbook: a9c982d24b476621316049b527e6c7bbc319f06b12268af150fa7d8c6df92230 | tests: 0 | failing: 0 | passing: 0

## Audit

- FLOW-028: "Current input MUST be checked even when baseText is absent or malformed." — missing assertion that a validly parsed current document targeting another module still returns corrective C05 when the base is absent or malformed. Those base conditions are combined only with current parser errors; the wrong-target case uses a valid base, so checking target equality after an early base rejection would pass.
- FLOW-029: "the exact sections Decisions, Questions, Audit, Provenance and Notes MUST be excluded" — missing assertion that same-case prefix lookalikes such as `Notes appendix` or `Audit details` remain significant. Existing excluded-name negatives exercise case folding, but a prefix-based exclusion would still pass them.
- FLOW-030: "every local Added handle in baseText MUST map positionally to the corresponding current Added ID or handle" — missing a matching valid pair whose base Added list mixes a fixed real ID and a local handle. Current positive bases have only handles or only real IDs, so an implementation that maps handles only when every base Added item is a handle would pass. Construct the mixed fixture directly under the public change grammar, retaining the real ID while mapping the local handle.
