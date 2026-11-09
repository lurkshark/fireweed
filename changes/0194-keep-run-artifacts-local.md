---
bases:
    FLOW-064: 82d7d776
    FLOW-065: eeb380bc
    FLOW-067: c0e89300
module: workflow
phase: impl
run: 20261109T070854-c99988411c0cf00ba64bcde1
---

## Motivation

The design keeps per-run logs under the local `.fireweed/runs/` directory. In
the supervised workflow drill, staging this directory prevented Done and
Block from recording their normal commits. Make the distinction between
step work and generated run artifacts explicit so staging and restoration
work with or without a matching Git ignore rule.

## Interface

Keep existing report and attempt locations. Done and Block keep generated
run artifacts local while recording the same phase and attempt history for
ordinary step work.

## Added

### FLOW-135
A phase commit created by Done MUST exclude generated run artifacts under
`.fireweed/runs/`, including when Git does not ignore those files.

## Modified

### FLOW-064
Block MUST save the held step's tracked and untracked work as a commit at the
next unused positive refs/fireweed/attempts/<slug>/<n>, preserving prior
attempt refs. It MUST exclude generated run artifacts under `.fireweed/runs/`
even when Git does not ignore them. BlockResult MUST identify the ref and its
commit OID.

### FLOW-065
After saving the attempt, Block MUST restore the step's tracked files,
untracked files and index to the phase base, removing the step's edits from
the checkout. It MUST leave generated run artifacts under `.fireweed/runs/`
on disk with their contents unchanged.

### FLOW-067
On successful blocking, Block MUST release the current claim and return the
blocked State with no held ClaimOID. The resulting working tree MUST be clean
apart from untracked generated run artifacts under `.fireweed/runs/`.

## Removed

## Decisions

## Notes

Extend the existing Done and Block fixtures. Cover an ignored run directory
and one with no matching ignore rule, with artifacts generated after Start.
Check that the report or log stays readable on disk and ordinary staged and
unstaged step edits appear in the
appropriate trees, and no generated run artifact appears in Done's phase
commit, Block's attempt, or the block commit. Keep the existing assertions
for phase transition, the state-only block commit, and claim release.

Prefer selecting the ordinary step paths before staging them and reusing
the existing scoped Git operations. Check the selected commands on the
supported verifier's stock Git. Avoid broad force-add or cleanup commands
that would include or remove the run directory.

The evals verify passed with the expected fail-first cases. With no ignore
rule, Done rejected its generated report as a protected path. Block saved the
run log in the attempt and removed it from disk. The ignored cases passed.

Done and its verification and application paths are implementation-writable
in `packages/workflow/internal/lifecycle/stub.go`. Block's attempt and
restoration paths are in `packages/workflow/internal/lifecycle/block_stub.go`.
On Apple Git 2.30.1, scoped `git add -A -- <paths>` and
`git restore --source=HEAD --staged --worktree -- <paths>`, both with
`--literal-pathspecs`, staged and restored the ordinary edits while leaving
the run log on disk in both fixture setups.
## Provenance
- evals 2026-11-09T07:19:18.004095Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=181 failing=4 passing=177
- audit 2026-11-09T07:19:41.654347Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-astra -c model_reasoning_effort="medium" -s workspace-write {prompt} fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
## Audit

No findings; checked FLOW-064, FLOW-065, FLOW-067, FLOW-135
