---
bases:
    APP-002: 101697fd
    APP-003: e86b3428
module: app
phase: audit
run: 004e57240af95b8e0d0fb8ab
---

## Motivation

Expose the checked project playbook resolver through the real executable while preserving standalone embedded authoring instructions. Project instruction reads need exact bytes and reliable errors without requiring valid module/spec/queue documents or starting enforcement hooks.

## Interface

No public Go API changes. Extend existing app.Run command composition using public project.ResolvePlaybook, contracts.ParseConfig and embedded assets. Preserve the public Run signature and schema-1 envelope. Implement all behavior under app internal/.

## Modified

### APP-002
explain <code> MUST return the embedded diagnostic's code and corrective explanation. playbook <name> MUST return exact embedded text with its byte-exact SHA-256 and name in JSON when no project override is selected, and MUST support the project binding contract below. Known codes and all five playbooks MUST succeed with exit 0 outside a project, using only the new binary. Unknown names/codes MUST be usage errors listing valid choices. Missing and extra positional inputs MUST be rejected. Embedded reads MUST NOT modify files or refs.

### APP-003
Every command in this foundational slice MUST accept --json before or after the command/operands and emit exactly one JSON value to stdout with schema_version 1, canonical command name, data and a non-null diagnostics array. Successful diagnostics MUST be empty. JSON usage errors MUST return exit 2, data null and a diagnostic with code U01, severity error and a nonempty corrective message, without plaintext on stdout or stderr. Plaintext usage errors MUST have empty stdout and a U01 diagnostic on stderr. Project-bound operational errors MUST retain their command-specific diagnostic contract. Output write failures MUST return exit 2.

## Added

### APP-017
playbook MUST discover a project from a nested working directory, binding only to the nearest enclosing checkout's root fireweed.yaml and ignoring nested config/override decoys and parent-project instructions across a nested checkout boundary. A checkout includes ordinary .git directories and linked-worktree .git files. When that root has no config, or no enclosing checkout exists, it MUST return embedded instructions without executing Git or reading any config/override contents. When root config exists it MUST safely read and validate it with public ParseConfig and the embedded binary version, then use public ResolvePlaybook; it MUST NOT require valid unrelated modules/specs/queued changes or read their contents. Project override absence MUST return embedded fallback. Readable selected overrides, including stale/malformed frontmatter, empty files and CRLF bytes, MUST succeed with exit 0 and empty diagnostics. Plain stdout MUST be exactly the selected Text with no added newline; JSON data MUST contain exactly name, text, hash, embedded_hash and overridden, retaining public resolver values. No based_on or stale field is added to this CLI surface; stale warnings belong to doctor.

### APP-018
Project-bound config errors, incompatible pins, unreadable config, directories or unsafe/dangling config links MUST return C01 severity error at fireweed.yaml with exit 2, null JSON data, corrective guidance and no fallback. Git discovery failure after binding MUST also return that C01 error. Public override read/containment failures MUST return D05 severity error at .fireweed/playbooks/<name>.md with exit 2 and null data, preserving corrective resolver detail. Plain errors MUST have empty stdout and coded/path-attributed stderr; JSON MUST have only the schema-1 canonical playbook envelope and non-null diagnostics on stdout with empty stderr. Unknown names, missing/excess operands, unknown/valued flags and help MUST be handled before project/config/override access, preserving the existing -- and --json rules. Output writer failures MUST return exit 2 for embedded, override and operational-error output. Successful and rejected calls MUST remain noninteractive, preserve files/ignored/outside content/index/HEAD/branches/claim/other refs/common Git and parent cwd/environment/temp storage, and MUST NOT execute hooks, agents, bootstrap or network commands. Config, override and unrelated contents MUST NOT be read through escaping paths, including when an eventual read error would be ignored.

## Decisions

Project binding is a local filesystem search for the nearest enclosing .git marker followed by checking only that checkout root's config. It does not execute Git merely to serve standalone embedded instructions. This preserves the previously audited no-Git embedded read behavior. A linked checkout's root config and override win over the primary checkout. ParseConfig validates the project pin/settings without discovering unrelated modules or queued documents, so a worker can read its instructions while repairing those documents. Unknown names and help remain globally available before binding. File containment is checked before content reads; config failures are C01 while a selected override failure is D05. This extends the command, not the resolver's no-configuration-read API contract.

The five existing JSON data keys remain unchanged. Full override bytes and their hash flow through exactly, including frontmatter. Stale metadata does not become an execution failure or warning in this read command, consistent with the design's doctor-only warning. The existing standalone behavior outside a project remains the deliberate D007 difference from the reference; D019 containment also applies here. No full lifecycle/doctor behavior is implemented by this change.

## Rejected alternatives

Loading all modules and queued documents for one playbook read introduces unrelated failures and unnecessary reads. Running Git before any standalone lookup breaks embedded-only operation. Falling back after unsafe or unreadable project state hides a broken configuration. Normalizing override bytes or adding a plain newline changes the hashed worker instructions.

## Provenance

- evals | 2026-10-22T21:20:16Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 134 | failing: 66 | passing: 68
