---
module: workflow
phase: audit
run: 4f474a62c65de2db4e4cc261
---

## Motivation

Separate context and lifecycle invocations need the metadata approved by the owning claim, even after a worker has changed or deleted current configuration, specs or queued text. Compose the immutable project loader and exact claim store before binding the executable context command. Current source bytes remain worker feedback, not permission authority.

## Interface

Add public LoadApproval(start, slug string) (project.Snapshot, Claim, error). It returns source metadata with the selected approved Change substituted when claimed, or current discovered metadata and a zero Claim when unclaimed. Keep existing snapshot/claim wire shapes and BuildContext unchanged. Implement behavior only under workflow/internal; a thin facade forwards the call.

## Added

### FLOW-020
LoadApproval MUST accept only an exact valid claim slug and discover the canonical repository from start, including nested and linked checkouts. Invalid selectors MUST fail C05 before repository or content access. It MUST inspect only the exact selected shared claim ref, independently of other claims, and distinguish an absent ref from a malformed payload, symbolic alias or inspection failure. With no selected ref it MUST load the current project through public Load and return its metadata with a zero Claim. With a valid selected claim held by this canonical checkout it MUST load metadata from exactly Claim.Base through public LoadAt and return that exact Claim, including Object. Current config, module SPEC, queued/archive document bytes or moving/replacement refs MUST NOT replace this committed metadata. Both SHA-1 and SHA-256 repositories supported by installed Git MUST work.

### FLOW-021
In either source LoadApproval MUST select exactly one queued change by exact Slug and exactly one matching target module, and require its canonical immediate changes/ path to match the public filename grammar and slug. A selected archived change MUST NOT substitute for a missing queued change. For a claimed source it MUST parse Claim.State using that baseline path and target prefix, require a valid unchanged started document through public CheckState, and require its Module and Run to match the baseline target and Claim.Run, with phase evals/audit/impl. It MUST replace only that queued Change with the complete parsed approved record, retaining its assigned IDs, metadata, sections, path and order; all other baseline config, modules, changes, archive, diagnostics and BaseRef MUST remain as loaded. Initial late-ID/run/phase assignment and later committed phase states MUST both work. A missing or invalid current change MUST NOT prevent metadata restoration. It MUST NOT project preceding or selected changes itself; public Project/ApplyChange/BuildContext remain responsible for projection exactly once.

### FLOW-022
Missing/ambiguous queued selectors, missing/ambiguous target modules and invalid selected paths MUST fail with corrective C05; claim lookup/decoding, ownership, baseline loading or approved State/identity failures MUST fail with corrective P02 and preserve useful underlying public error details. Other discovery failures MUST retain corrective C01 details. A foreign-worktree, empty reservation, invalid started state, run/target mismatch, unavailable baseline, malformed selected ref or failed selected-ref inspection MUST NOT fall back to current metadata. HEAD differing from Claim.Base MUST NOT prevent read-only restoration: CheckWorktree separately rejects that mismatch before advancement. Every failed call MUST return a zero Snapshot and zero Claim, never partial metadata.

### FLOW-023
LoadApproval MUST remain local and noninteractive and preserve HEAD, branches, claim/other refs, index bytes, tracked/untracked/ignored/user/outside files, common Git and temporary storage, and parent cwd/environment on success and rejection. It MUST NOT commit, stage, reset, fetch, push, create branches, execute hooks/agents/bootstrap/network commands, or write files/refs. For a claimed source it MUST NOT open current config/SPEC/queue/archive contents, SYSTEM.md, playbooks, retained contracts/evals, disposable code or unrelated private contents; only necessary local Git metadata and committed source objects selected by LoadAt may be read. For an unclaimed source only the normal public Load discovery reads are permitted. Unsafe paths MUST reject before outside content opens, including ignored-error attempts. Returned snapshots and their nested maps/slices MUST be independent across repeated calls and after mutation, including empty append backing storage. Returning metadata alone MUST NOT execute coverage, boundaries or other project hooks, and MUST NOT validate current worker edits or change ownership.

## Modified

## Removed

## Decisions

An absent selected claim is a normal author/resume inspection source, represented by a zero Claim; malformed or foreign ownership is not absence. Other malformed claims cannot block an explicitly selected change. The source loader chooses metadata authority before opening current source documents, so worker-edited configuration cannot redirect inspection. Current text is read later by BuildContext under its existing containment and feedback contract. Static project/queue checks, clean-tree/eligibility, HEAD equality and advancement remain separate caller responsibilities; restoring a snapshot does not certify them. Loader diagnostics remain data rather than being discarded or silently repaired.

Claim.State is the controller's approved step document. It may legitimately differ from the baseline queued text because start assigns IDs/run/phase without committing, and phase transitions refresh the claim after commits. Identity is checked against the baseline path/target and exact claim run; all state grammar and recognized phase requirements remain enforced. A moved HEAD remains inspectable for repair while CheckWorktree still prevents advancement. No new public result model or persistent export/cache is needed.

## Rejected alternatives

Reloading current config/SPEC before examining ownership lets worker edits select their own permission metadata. Treating every claim read error as absence hides corrupted ownership. Looking up all claims makes an unrelated malformed ref block explicit inspection. Projecting during restoration would apply queued effects twice when BuildContext performs its existing projection. Requiring clean files or matching HEAD for this read would hide the context needed to repair invalid work.

## Provenance

- evals | 2026-10-23T01:19:08Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 1220 | failing: 86 | passing: 1134
- audit | 2026-10-23T01:24:56Z | agent: manual | fireweed: 0.4.0 | playbook: a9c982d24b476621316049b527e6c7bbc319f06b12268af150fa7d8c6df92230 | tests: 0 | failing: 0 | passing: 0
- evals | 2026-10-23T01:39:14Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 1232 | failing: 98 | passing: 1134
- audit | 2026-10-23T01:42:56Z | agent: manual | fireweed: 0.4.0 | playbook: a9c982d24b476621316049b527e6c7bbc319f06b12268af150fa7d8c6df92230 | tests: 0 | failing: 0 | passing: 0
- evals | 2026-10-23T01:52:29Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 1232 | failing: 98 | passing: 1134

## Audit

- FLOW-021: "In either source LoadApproval MUST select exactly one queued change by exact Slug" — the cur selector has no selected cur claim in either claimed loop variant, so both negative partial-selector cases use the unclaimed source. Claimed current with a prefix-neighbor rejects prefix-as-ambiguity but permits a prefix-first matcher. Add a valid exact cur claim whose baseline contains current/current-neighbor but no cur change, establish that claim through public LoadClaim, and require C05 from the claimed source.
