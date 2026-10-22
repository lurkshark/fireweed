---
module: project
phase: audit
run: 72ae08640c3a077e700e83a5
set:
    interface:
        - api.go
        - model/*.go
---

## Motivation

Resolve the exact project-owned instructions and their embedded provenance before constructing worker context or adding doctor and CLI override checks. Overrides are author-owned text; a stale based_on hash warns without silently replacing the author's instructions.

## Interface

Add retained project/model Playbook with string fields Name, Text, Hash, EmbeddedHash and BasedOn, and booleans Overridden and Stale. JSON keys are name, text, hash, embedded_hash, based_on, overridden and stale. Add a public project Playbook alias and ResolvePlaybook(start, name string) (Playbook, error), forwarding to new disposable implementation. Existing APIs remain intact. The caller chooses project binding; this API discovers the Git checkout but does not load project configuration. Later CLI and context changes compose it, and doctor will render stale information as D05 warnings.

## Added

### PRJ-021
ResolvePlaybook MUST recognize exactly the five names in public assets.Playbooks and resolve the canonical Git checkout from start, including nested directories and linked worktrees. Without the selected .fireweed/playbooks/<name>.md override it MUST return the exact public embedded Text and Hash, EmbeddedHash equal to that embedded Hash, Name equal to the requested name, BasedOn empty, Overridden false and Stale false. With a readable regular selected override it MUST return every exact byte of that file as Text, SHA-256 of those exact bytes as lowercase full Hash, the unchanged embedded hash as EmbeddedHash, and Overridden true, including empty files and CRLF text. Only the selected override may affect the result; unrelated overrides and project configuration MUST NOT be read. Unknown or path-like names MUST error with valid-name guidance before repository discovery or override reads; a known name outside any Git repository MUST error with corrective repository guidance. Every error MUST return a zero Playbook.

### PRJ-022
For an override, ResolvePlaybook MUST extract BasedOn only from a string-valued based_on key in valid public Markdown frontmatter. Stale MUST be false exactly when that string equals the selected current embedded hash, and true for missing, wrong, differently cased or non-string hashes and malformed frontmatter. Staleness MUST NOT replace or reject readable override text: Text and its exact-byte Hash remain the override's values, even for invalid frontmatter, absent based_on or an empty file. The full file including frontmatter MUST remain in Text and Hash. Embedded-only resolution MUST remain non-stale independently of other overrides.

### PRJ-023
ResolvePlaybook MUST check selected override containment using the public Git SafePath contract before opening its contents, including when a read error would otherwise be ignored. Missing selected overrides MAY fall back to embedded text only when the missing path remains contained. Unreadable files, directories and escaping or dangling symlinks, including missing descendants of escaping directory symlinks, MUST return a corrective error mentioning D05 and the selected repository-relative override path, never silently fall back or read outside contents. Contained regular-file symlink targets MAY be used. Resolution MUST be deterministic, noninteractive and read-only on success, stale overrides and all errors: preserve tracked/untracked/ignored/outside files, index bytes, HEAD/branches/claim/other refs, parent cwd/environment and temporary storage; never run hooks/agents/bootstrap/network commands or mutate project configuration. Calls and returned values MUST remain independent without mutable aliases or input mutation.

## Decisions

This API binds explicitly to an existing Git checkout without reading fireweed.yaml. The CLI retains its existing embedded-only behavior outside a project and will opt into this resolver when project binding is appropriate. Hashes cover the complete exact override file, including frontmatter, because that is the text supplied to workers and recorded in provenance. based_on inspection uses the public Markdown parser; malformed metadata is advisory staleness, while failed or unsafe reads are errors. Empty override files remain author-owned instructions and produce a stale flag. SafePath defines containment; inside links remain permitted without weakening worker zone rules.

The reference Project.Playbook reads complete override bytes and hashes them, and Doctor compares based_on with its embedded hash (internal/fireweed/playbooks.go and doctor.go). This fresh API exposes those two observations together without running doctor or hooks. Static input validation and command rendering remain callers' responsibilities.

## Rejected alternatives

Automatically replacing stale overrides discards author instructions. Silently treating permission or containment errors as absence hides broken or unsafe project state. Reading all overrides or loading configuration couples a single instruction lookup to unrelated project problems. Hashing only the Markdown body would record a different artifact from the exact worker instructions.

## Provenance

- evals | 2026-10-22T20:43:05Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 180 | failing: 102 | passing: 78
- audit | 2026-10-22T20:48:20Z | agent: manual | fireweed: 0.4.0 | playbook: a9c982d24b476621316049b527e6c7bbc319f06b12268af150fa7d8c6df92230 | tests: 0 | failing: 0 | passing: 0
- evals | 2026-10-22T20:53:11Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 223 | failing: 137 | passing: 86
- audit | 2026-10-22T20:55:38Z | agent: manual | fireweed: 0.4.0 | playbook: a9c982d24b476621316049b527e6c7bbc319f06b12268af150fa7d8c6df92230 | tests: 0 | failing: 0 | passing: 0
- evals | 2026-10-22T20:56:54Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 229 | failing: 142 | passing: 87

## Audit

- PRJ-023: "ResolvePlaybook MUST check selected override containment using the public Git SafePath contract before opening its contents, including when a read error would otherwise be ignored." — the missing-descendant native probe traps the selected alias, but still supplies outside/private as its exact forbidden target. The fully resolved missing endpoint outside/evals.md is not trapped. An incorrect resolver can resolve the escaping parent, attempt to read that outside endpoint, receive ENOENT and then return the expected D05 error while passing every current assertion. Pass the missing outside/evals.md endpoint as the forbidden target for this case and add a direct ignored-error read positive control for that same endpoint, while retaining the selected-alias control.
