---
module: workflow
phase: audit
run: d61fa60bbbecd202f1616c09
---

## Motivation

Checked start and unattended queue driving need the same deterministic answer about which change may proceed around blocked work. Establish the read-only queue decision before adding claim/file mutations. Eligibility is separate from clean-tree, static approval, claim ownership and step verification.

## Interface

Add public Eligible(snapshot project.Snapshot, slug string) (bool, error), forwarding through api.go to disposable workflow/internal code. Keep existing public models and functions unchanged. This pure operation consumes supplied public source metadata and never reloads a repository or claims it.

## Added

### FLOW-024
Eligible MUST select exactly one queued change by an exact valid claim slug, rejecting invalid, missing or ambiguous selectors with corrective C05 and false. An archived-only selector MUST NOT substitute for a queued change. Queue precedence MUST be numeric Number ascending, then Slug and Path ascending, independently of supplied slice order and without sorting that slice in place. The selected change and preceding entries participate; later entries MUST NOT affect the result, except that a duplicate selected Slug remains an ambiguity. A selected blocked change MUST be ineligible. Ineligibility in an otherwise valid prefix MUST return false with nil error rather than inventing a failure diagnostic.

### FLOW-025
An unblocked selected change MUST be eligible only when no earlier runnable unblocked change exists and no earlier blocked change holds it. A blocked change holds a later change when they target the same module, when the later target depends directly or transitively on the blocked target, or when the later change has a nonempty Interface section and the blocked target depends directly or transitively on the later target. Earlier unblocked changes held by earlier blockers MUST be skipped rather than stop independent work; every earlier blocker MUST remain relevant, not only the nearest one. Later blockers MUST NOT hold earlier work. Multiple independently runnable changes MUST still have only the earliest eligible queue head. Empty or whitespace-only Interface bodies MUST NOT activate reverse dependency blocking, and other/fenced heading text MUST NOT substitute for the exact parsed Interface section.

### FLOW-026
For each candidate whose dependency relation is evaluated, Eligible MUST obtain its inclusive projected graph through public Project before that candidate and ApplyChange for the candidate exactly once. The selected inclusive graph MUST be validated even when the selected change is blocked or behind an earlier runnable head. This MUST include preceding set/depends_on changes, including blocked intent, and the candidate's own set values, while excluding later intent. It MUST use that graph for direct/transitive forward and Interface-dependent reverse relationships. Unknown or ambiguous target modules, missing dependency modules, self-dependencies, dependency cycles and failed prefix projection/application MUST fail atomically with corrective C05 and false, preserving useful public projection error details. It MUST NOT silently use an unprojected graph, ignore a failed projection, expand a cycle forever or treat a missing module as independent. This is essential graph/selection validation, not certification of all static project/queue requirements; unrelated hook/coverage/zone validation remains separate.

### FLOW-027
Eligible MUST operate only on supplied values outside any repository, remain local and noninteractive, and MUST NOT inspect files, Git, claims, hooks, agents, bootstrap, network, cwd or environment, or mutate files/refs/temporary storage or parent process state on success or rejection. Inputs and every nested map/slice/document MUST remain unchanged, including queue/module order and empty append backing storage, after successful and rejected repeated calls. The same supplied values MUST produce the same bool/error regardless of Root, current working directory or ambient files/environment. It MUST NOT allocate IDs, rewrite changes, apply specs to disk or run clean-tree, approval, coverage, boundary or worker-edit checks.

## Modified

## Removed

## Decisions

This is the single-head queue decision for later manual start and runner composition. Blocked fields, dependency names and exact parsed Interface content come from the supplied source snapshot; no current worker source is read. Every nonempty Blocked string represents a blocked change. The prefix graph includes queued metadata effects consistently with the existing public projection contract. Evaluating an earlier unblocked candidate uses its own inclusive prefix, so a later change cannot retroactively make that earlier candidate eligible or held. Caller-owned static validation, clean state and exact claim acquisition remain prerequisites for advancement.

Only essential selection/graph failures are errors; normal blocked/held/behind-head decisions are false. A duplicate selected slug is never an order tie that can choose an arbitrary change. Full public snapshots may retain unrelated discovery diagnostics without this pure policy certifying or executing them. No new persistent result/cache or public model is needed. The design's blocking rules motivate this prerequisite; complete start/verify/done, attempts, regeneration deletion and runner behavior remain subsequent changes.

## Rejected alternatives

Treating the first blocked entry as a global stop prevents independent work. Looking only at direct dependencies misses transitive blocking; blocking all reverse dependents prevents a non-interface dependency implementation from proceeding. Ignoring projected dependency changes lets queued intent evade its own ordering rule. Reading claims here would couple a pure ordering decision to worktree ownership and make it unusable for deterministic planning.

## Provenance

- evals | 2026-10-23T03:57:05Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 1287 | failing: 42 | passing: 1245
