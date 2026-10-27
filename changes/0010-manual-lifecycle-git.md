---
audit_pending:
    - GIT-003
audit_rounds: 2
bases: {}
module: gitstore
phase: impl
run: 20261027T051014-bb3072879b91b81e7176489d
set:
    interface:
        - api.go
        - model/**
---



## Motivation

The design's Git behavior, checked lifecycle and hook contracts require local
repository snapshots, phase diffs/commits, atomic claim refs and bounded process
execution. These are the minimum Git/process operations needed to connect the
existing manual fixture to start/verify/done; history, scratch worktrees and
retry/drill orchestration are not part of this change.

## Interface

Expose retained data-only Repository (Root, CommonDir, Head, Branch strings),
Change (Path, OldPath, Status strings), Request (Args []string, Dir string,
Env map[string]string, Input []byte, Timeout time.Duration) and Result (Stdout,
Stderr strings, ExitCode int), with thin public aliases/wrappers. API:
Inspect(dir) (Repository,error), BaseRef(root,remote,base) (string,error),
Changes(root,base) ([]Change,error), ReadAt(root,revision,path) ([]byte,error),
Commit(root,subject string,paths []string) (string,error), WriteBlob(root string,
data []byte) (string,error), ReadBlob(root,oid) ([]byte,error), Ref(root,name)
(string,bool,error), CASRef(root,name,old,new string) error and Run(ctx
context.Context,request Request) (Result,error). String arguments are strings;
BaseRef returns full refs/heads/... or refs/remotes/... names. Empty CAS old means
create-only; empty new means delete. Behavior stays under internal/.

## Added

### GIT-001
Inspect MUST return canonical absolute worktree/common-Git directories plus HEAD and branch, from a directory inside a primary or linked worktree. The branch MUST be empty for detached HEAD, and non-repositories MUST return an error.

### GIT-002
BaseRef MUST prefer a locally known remote/base ref, falling back to the local base branch and rejecting a missing base. An omitted remote MUST mean origin. An omitted base MUST use the local remote HEAD's branch when known, otherwise main.

### GIT-003
Changes MUST report repository-relative staged/unstaged tracked edits and untracked paths against the given base, retaining both names for staged renames. ReadAt MUST return a revision's file bytes or an error for absent paths. Commit MUST commit only explicitly selected paths with the given subject, return the new HEAD, and leave unrelated working/index changes uncommitted.

### GIT-004
WriteBlob and ReadBlob MUST round-trip byte payloads through local Git objects. Ref MUST return the named local ref's object ID and existence. CASRef MUST create, update or delete a ref only when its exact current object matches the expected old value, returning an error without overwriting a newer value on mismatch.

### GIT-005
Run MUST launch argv in the requested directory with the supplied environment/stdin and return captured stdout, stderr and exit code. A nonzero exit MUST return an error while preserving output. Timeout or context cancellation MUST terminate the launched process group and return an error; processes escaping that group are outside the supported contract.

## Modified



## Removed



## Decisions



## Provenance

- evals 2026-10-27T05:17:38.756278Z agent=manual fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=138 failing=7 passing=131
- audit 2026-10-27T05:20:39.059751Z agent=manual fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0
- evals 2026-10-27T05:23:03.480159Z agent=manual fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=138 failing=7 passing=131
- audit 2026-10-27T05:30:15.420149Z agent=manual fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0
- evals 2026-10-27T05:32:57.623355Z agent=manual fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=138 failing=7 passing=131
- audit 2026-10-27T05:37:33.111254Z agent=manual fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0

## Audit

No findings. Checked GIT-003.
