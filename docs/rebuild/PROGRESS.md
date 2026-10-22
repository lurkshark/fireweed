# Restartable rebuild progress

Status: stage 1, seven behavioral changes archived; project discovery/static graph next. Goal remains active.

## Controller pin

- Source: /Users/joel/Workspace/fireweed-bootstrap at 583b3c53740cf2341e0b0a79afee5b7f646ec1d4 (clean before/after build).
- Build: GOCACHE=/tmp/fireweed-rebuild-cache go build -o /Users/joel/Workspace/fireweed/.fireweed/bin/fireweed-bootstrap . (from bootstrap checkout).
- Toolchain: go1.27.2 darwin/arm64. Bootstrap version JSON: 0.4.0, schema 1.
- Binary SHA-256: f9c48ae6b14887ebc549110478c081bedc450847507c6967735ef8486c98a3ed. Ignored binary must be rebuilt and checksum checked if absent.
- Design SHA-256: 1c17d0ab9496b1b15ad5ca18eed4f4ca623d01cff596bc557bf4c97400d00789.
- Initial target: 1b27763, clean main. No remote changes, pushes or publication.
- Bootstrap help/version/author-change read; Go recipe initialized with codex worker and codex auditor. Codex CLI present and login status reports ChatGPT authentication; a real invocation was verified later (see Agent access).

## Setup and next work

Eight scaffold modules with proposed dependency DAG; protected design/planning/build/enforcement files; updated Go canary imports; compiling app facade/placeholder. Controller check and go test ./... pass (no behavioral tests yet). Setup baseline: 14cec73e1126ee37d33f1138d19936b6a05a59a6. Controller doctor --json passed boundaries/interface/failing-eval/missing-id, including recovery, with zero diagnostics. Git metadata writes and scratch creation require the sandbox escalation used for authorized local operations. Next: contracts spec-primitives through fail-first, independent audit and impl.

All R01–R21 remain unverified. No product acceptance, real regeneration or self-hosting is claimed.

## Restart procedure

Inspect git status, pinned controller checksum/version, controller status/list and refs/fireweed/claims before any edit. Read current change/context if claimed; resume its phase without resetting work. Only update protected evidence between active changes. Final history/CI uses explicit setup-SHA..HEAD, never main default range.

## spec-primitives (CON-001–CON-004)

- Archive: changes/archive/0001-spec-primitives.md. Four permanent public-API evals in packages/contracts/evals/specs_test.go cover precise ID matching, normalized heading/tag/body hashes, fenced Markdown grammar, tags and local handles.
- Evals commit 372cc4f; independent passing audit 4a7d559; implementation ac75deecc8c6ec9a3e871035639c6ed9aadfd880.
- Controller fail-first JSON has four compiled/executed failed cases, one per ID; no coverage/zone diagnostics. `done` committed evals before implementation. Local ignored evidence: .fireweed/evidence/spec-primitives/{fail-first,evals-done,audit-done,impl-verify,impl-done}.json.
- Auditor /root/audit_spec_primitives used fresh context restricted to change, spec, interface and evals. No findings, all four IDs listed. Controller impl verify ran own + all seven transitive dependents and boundaries: no diagnostics; four own cases passed. Empty scaffold dependents do not prove product functionality.
- Normal P04 unapproved-on-base warning acknowledged under goal authority; no human merge claimed.

## Agent access

The default CLI model gpt-6.1-sol was rejected by ChatGPT auth (HTTP 400 unsupported model). Local models_cache.json listed gpt-6-sol; an explicit ephemeral read-only probe returned FIREWEED_AGENT_READY with exit 0. Pin `codex exec --ephemeral --model gpt-6-sol --sandbox workspace-write {prompt}` for worker and audit commands. Do not change global config. Logs are local ignored evidence. This is capability evidence only, not a regeneration pass. Official model documentation: https://learn.chatgpt.com/docs/models .

## zone-primitives (CON-005–CON-007)

Archive changes/archive/0002-zone-primitives.md. Permanent Test_CON005GlobMatching, Test_CON006FutureOverlap, Test_CON007ZonePrecedence and Test_CON007AuditRegressions cover whole relative-path globs, nonexistent-path intersection, all precedence levels, module-relative matching, protected documents and invalid inputs. Initial fail-first commit 5402f70990472dd91ff84db90679c359fd6d5a0c. Independent audit found four concrete CON-007 assertion gaps: recorded in 9bc68250a4322bc4046f0b944ffccf80b99b289d and returned to evals. Strengthened fail-first evals b5624fd1362dff7fc2e1ee430a8cc87e74a986bc; passing independent re-audit 7d7f600e92a8ff6ff6ab61e45adfe27b834b4e0f; applied implementation d7dfbd90bf4cebfdeb168268bd710f67dd6b43bb. All eight contracts cases plus dependent scaffold hooks and boundaries pass. Explicit setup..HEAD history check has no diagnostics. Local JSON evidence is .fireweed/evidence/zone-primitives/. No active claims or pending changes at this boundary.

Next: retained typed data/configuration grammar, then Git/assets/project discovery and early CLI inspection slices. Product binary still only the setup placeholder; do not confuse its successful build with implemented command behavior.

## config-contract (CON-008–CON-011)

Archive changes/archive/0003-config-contract.md. Seven added permanent testcases in packages/contracts/evals/config_test.go verify defaults/explicit empty lists, exact typed field retention and opaque commands, independent parsed slices, strict YAML roots/types, current/previous formats, numeric version conjunctions and duration/retry settings. Initial fail-first b61d96d7e569dd497ca3f67722a14fefbd3a21f3. Audit 75045c563aaf490ef66e75debd98c034301f0b15 captured seven assertion gaps; retry evals 85375462e3f24fd0d0f820abb696d725ea90ed4f resolved them. Independent re-audit pass 3d67cc6be9f5ba0c81013b210e98cfe0b651806a. Applied implementation be9fe8b312380bfbc4a02334f375cdb2fa29ae57. Retained model/*.go was introduced using the change's projected interface set; only the controller applied that zone to SPEC.md.

Checkpoint after be9fe8b: go test -race ./..., go vet ./..., go build -o bin/fireweed . all exit 0. Controller ci 14cec73e1126ee37d33f1138d19936b6a05a59a6..be9fe8b312380bfbc4a02334f375cdb2fa29ae57 passed with no diagnostics, including doctor (not skipped), own coverage and all scaffold modules. Fifteen actual contracts testcases. Saved JSON .fireweed/evidence/ci-contracts-checkpoint.json. Bootstrap checkout remains clean; no claims or pending changes. This is a partial checkpoint, not the final acceptance range. The Go build emits a harmless denied module-stat-cache write warning in the outer sandbox; build still succeeds.

## document-grammar (CON-012–CON-016)

Archive changes/archive/0004-document-grammar.md. Ten permanent testcases in packages/contracts/evals/documents_test.go assert fenced Markdown and independent metadata round trips, module default inheritance/overrides and slice isolation, complete typed change state/classification, normative edit sections/Removed headings, local handles, retained author prose, set allowlist and hash/ID grammar. Initial fail-first dd6098a; independent audit ea3f868 records twelve concrete missing-assertion findings; retry fail-first a38906e; independent passing re-audit b6a39ae. Applied implementation 45f5bb5b878b6e46f06f7c034c9f96a59f41c6dd. Controller verify: 25 contracts cases, all own/dependent hooks and boundaries pass; explicit baseline history check passes. One implementation iteration fixed a real YAML block-scalar trailing-newline failure; evals/interfaces stayed unchanged. Local evidence .fireweed/evidence/document-grammar/.

Restart audit: prior goal turn classified as progress (authoritative archived code, evals and CI evidence). At this turn's start checksum matched, queue/claims empty, target clean. Goal still requires the full design; contracts primitives do not satisfy CLI/lifecycle/regeneration/self-hosting gates. Next implement gitstore local discovery/base refs/clean-tree/diffs/safe paths, then assets/project and an early CLI inspection slice.

## local-repository (GIT-001–GIT-004)

Archive changes/archive/0005-local-repository.md. Eight permanent Git fixture testcases verify canonical discovery and linked-worktree directories, local base/default remote resolution, no-network Git traces and ref immutability, clean-tree cases, explicit older-base NUL-delimited diffs/renames/unusual filenames, and symlink-safe paths without file mutations. Fail-first 83f6008d70787f25877aa8fac8822c582332cb37; eight independent audit findings e04fc7295c642442ab871ef88adc163ca8c89d22; stronger fail-first 849cb6cd9f71174957971105a8e6562486b48e80; passing independent re-audit f731985fdcebceb480685531797797aec0e628a7; applied aea039d2aa9dfff238435df66403859edb6e6f80. Own/dependent coverage and boundaries pass. An implementation iteration removed unsupported Git --path-format=absolute and resolved relative --git-common-dir output against the checkout root; tests stayed fixed. No bootstrap source edits.

After aea039d, go test -race ./... passed; go vet ./... reported an unkeyed external model.Spec literal in the retained facade. The preceding evidence commit incorrectly recorded vet as passing; this correction supersedes that claim. The author repair keys that unchanged conversion literal without changing API shape or behavior, between claims. Explicit baseline history check passes. Next: assets version/playbooks, then a real help/version CLI slice while project discovery is built. Remaining CLI/lifecycle/recovery/drill/packaging gates are still unfinished.

Vet repair verification: keyed facade literal; go vet ./..., go test ./... and controller check now exit 0. No behavioral specs/evals or disposable implementation were edited by the author repair.

## embedded-protocol (AST-001–AST-003)

Archive changes/archive/0006-embedded-protocol.md. Six permanent testcases assert compiled version/current/previous format, five complete step playbooks with exact SHA-256 text hashes and independent inventories, all 38 exact diagnostic problem/remedy messages, and a public compiled lookup probe without source/network access. Playbooks and diagnostics are fresh design-derived text; diagnostic codes/meaning were compared with the pinned reference catalog. No bootstrap content is read at runtime.

Phase history: initial fail-first ac1cf4a; semantic/isolation findings 5161809; stronger fail-first 7e464aa; remaining format/command/I/O findings 96fdc7e; interception fail-first 6b0e96a; error-path finding bbbbb33; final failing evals a8bb493; independent passing audit 3868be2. Applied implementation a8eabefafe937d11c5de10b1510192319ee701ee. All findings are real separate-context audit outputs and retained in phase commits. Eval/interface files stayed immutable during implementation. Local JSON evidence .fireweed/evidence/embedded-protocol/.

Isolation details: a separately compiled Go public-API probe loads a test-only Darwin I/O interposer, armed after OS loader startup. Positive controls terminate on private/repository file reads and real loopback network attempts even with ignored errors. A native policy independently denies workspace/sentinel reads and networking. Paths must be canonical to avoid /var aliases. Early overly broad native policies aborted in the OS loader, not the product; final controls pass and the implemented asset lookups pass under the armed guard on both valid and invalid inputs. Only the macOS isolation eval requires local cc and an unsandboxed parent to install the tighter child profile. Other-platform native-isolation execution is not claimed; the case skips there. The product remains Go-only.

## Checkpoint 2026-10-22

At a8eabefafe937d11c5de10b1510192319ee701ee, go test ./..., go test -race ./..., go vet ./... and go build -o bin/fireweed . all exited 0 (including the native guarded probe). Controller ci 14cec73e1126ee37d33f1138d19936b6a05a59a6..a8eabefafe937d11c5de10b1510192319ee701ee passed with doctor enabled and zero diagnostics; invocation used ..HEAD at that exact endpoint. Coverage: contracts 25 cases, gitstore 8, assets 6; other five scaffold modules 0. Evidence .fireweed/evidence/ci-assets-checkpoint.json. Queue and claims empty, bootstrap checkout clean.

Next action: author an app command slice for actual executable help/version and embedded explain/playbook reads with versioned JSON/usage exits. Use explicit NEW subject binaries in permanent app evals, separate from controller. Then build project discovery/graph/projection/static validation and extend CLI inspection/authoring; continue all later lifecycle/runner/history/recipes/distribution gates. bin/fireweed still contains the setup app placeholder at this boundary. No full design completion, real regeneration, self-hosting or final acceptance is claimed.

## cli-inspection (APP-001–APP-004)

Archive changes/archive/0007-cli-inspection.md. Eight permanent app evals build and execute a separately named NEW product binary from temporary directories, checking help/inventory/exit meanings, version aliases, all embedded explanations/playbooks, JSON envelope/data/error paths, option boundaries, writer failures and external-agent/controller traps. No controller is used as subject. Initial fail-first 3ac20ca; two independent audit findings cb63082; stronger fail-first c3fc980; passing independent re-audit c01421b; applied c3e9428dc7c61344d7cee6f4dbb2fdbb2f560463. Controller impl verify passed app's eight cases and boundaries. bin/fireweed was rebuilt and its version --json returns schema 1, version 0.4.0, diagnostics [].

Playbook JSON retains embedded_hash and adds overridden:false for later project-owned overrides. Embedded playbooks now work outside a project (intentional reference difference D007); pin/override integration is still pending. Remaining commands printed in help are not yet implemented. Next author project discovery/config/module/queue loading and graph/zone/hook validation, then extend actual executable inspection. Goal remains active and full acceptance incomplete.

Restart at this boundary: previous goal turn was progress (archived code/evals and full checkpoint evidence); current turn added real product CLI behavior. Queue/claims clear after archive; bootstrap remains the pinned controller.

## Discovery eval recovery

First discovery-validation cycle assigned PRJ-001–006 and passed a separate assertion audit, then implementation exposed three fixture errors (unknown metadata, too-short filenames, permitted adjacent-star pattern). See D008. Released the claim and fully reverted that cycle's eight phase commits to restore the original queue entry. Implementation/evals preserved under /tmp/fireweed-discovery-recovery-5380157 and earlier commits; only our uncommitted implementation was restored to its stub before reverting. Restart from evals with valid contract-derived fixtures; do not treat the reverted cycle as completed behavior or edit frozen evals in impl.
