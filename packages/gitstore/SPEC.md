---
depends_on:
    - contracts
id_prefix: GIT
name: gitstore
---

# gitstore

## Specs

### GIT-001
Repository discovery MUST work from a checkout or its nested directory and return the checkout root, worktree Git directory and common Git directory as canonical absolute paths. Linked worktrees MUST have distinct roots and Git directories but share their common Git directory. Non-repositories MUST return an error. Inspection MUST NOT create branches, change refs, fetch or push.

### GIT-002
Base resolution MUST use the configured remote (default origin) and configured branch; when the branch is omitted it MUST use that remote's locally known symbolic default, falling back to main. It MUST prefer an existing refs/remotes/<remote>/<base> over refs/heads/<base>, otherwise return the local ref name even when absent. It MUST NOT contact a remote. Branch names containing slashes MUST be retained.

### GIT-003
Clean-tree inspection MUST detect tracked staged or unstaged edits, deletions and untracked files, while ignoring ignored files. Worktree diffs against an explicit local commit MUST include staged and unstaged additions/modifications/deletions, untracked files and renames with old and new paths. Spaces, tabs and newlines in filenames MUST be preserved. Results MUST be deterministic and invalid Git roots or revision arguments MUST return errors rather than empty success.

### GIT-004
Safe path resolution MUST accept only nonempty repository-relative forward-slash paths without absolute roots, drive prefixes, dot/traversal components, repeated/trailing slashes, backslashes or NUL. It MUST return an absolute path inside the canonical root, allow nonexistent descendants and internal symlinks, and reject existing or ancestor symlinks escaping the root, including sibling paths sharing its string prefix. It MUST NOT create or modify files.

### GIT-005
Run MUST execute exactly the supplied executable and argument vector in dir, without shell expansion, and collect complete stdout and stderr into Output while returning the actual exit code and a non-nil Err for a nonzero exit. Successful execution MUST return Exit 0 and nil Err. It MUST inherit the parent environment, apply supplied KEY=value overrides with the last entry for each key winning (including empty values and values containing equals signs), preserve all caller inputs, and leave the caller's working directory and environment unchanged. Empty argv, an empty executable, nil context, malformed environment overrides, nonexistent working directories and launch failures MUST return Exit 2 with a useful non-nil Err without launching a fallback executable. Long output MUST be drained without deadlock or truncation.

### GIT-006
A context already cancelled or expired before Run MUST prevent process startup, return Exit -1 and an Err matching the context cancellation/deadline via errors.Is. Cancellation or deadline during execution MUST terminate the entire spawned process tree on supported platforms, including children and grandchildren that retain output streams, return promptly with Exit -1 and the matching context Err, and retain output emitted before termination. Descendants MUST NOT survive to perform delayed writes after Run returns. Cancellation MUST leave unrelated processes alive, including when the calling executable is instrumented for race detection. Ordinary successful and failing process termination MUST remain distinguishable from context termination.

### GIT-007
Independent Run calls MUST NOT share captured output or cancellation state, even when overlapping. Run MUST NOT mutate caller argv or override slices, including backing storage beyond their lengths, or the parent environment/current directory. Empty or invalid calls and pre-cancelled calls MUST leave no process-created marker. Invalid inputs and completed successful, failed or cancelled calls MUST NOT poison a later valid call. Process tests MUST use local bounded helpers and no agent/bootstrap/network dependency.

## Decisions

[local-repository]
Git operations use argument vectors, not shell interpolation. Edit statuses are A/M/D/R (or another first status character reported by Git); rename Path is the new path and OldPath the old path. Untracked files are A. The diff is relative to the requested commit, not merely the index. Identical entries are deduplicated and results sorted by Path then OldPath then Status. SafePath rejects dangling symlinks and roots that do not exist. Root canonicalization follows symlinks. Base resolution reads only local refs; names are validated with git check-ref-format to prevent revision/options injection. Invalid base/remote names return errors. No temporary branches or network commands are used by inspection.

[command-process]
Exit -1 identifies context termination inside this process API, while CLI layers retain their separately specified 0/1/2/3 meanings. The original context error is exposed for callers to classify timeouts and cancellation without parsing process messages. Launch/configuration errors use Exit 2. Explicit shell invocation remains possible but substitutions and quoting belong to the hook/agent caller, not this adapter. This primitive does not create Git branches or invoke Git by itself. Windows environment keys follow the operating system's case-insensitive semantics. Runtime platform evidence must be reported honestly; cross-compilation cannot establish cancellation behavior on an unexecuted platform.

[race-helper-exit]
Use the Go race runtime's exit-sleep setting only for bounded helper subprocesses. Do not increase the tested process termination bound or remove the live unrelated-process predicate. This is a test harness repair for already implemented behavior, governed through evals/audit/impl/archive so permanent eval changes are reviewable and never smuggled into an implementation commit. Implementation phase requires no product code edits. Keep the original race failure in .fireweed/evidence/go-race-process-checkpoint.log and record the corrected full-run result separately.
