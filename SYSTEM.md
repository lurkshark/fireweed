# Fireweed

Fireweed is a local, language-agnostic Go CLI for checked spec-driven development and regeneration. The unchanged fireweed-design-v1.md is product authority. docs/rebuild/GOAL.md, PLAN.md and ACCEPTANCE.md define rebuild process and completion gates. The bootstrap is a behavioral reference, never a product dependency or command implementation. Record discrepancies in docs/rebuild/DECISIONS.md and design-derived evals.

## Modules

contracts owns public data, Markdown grammar, IDs, hashes and zones. gitstore owns safe local Git and process primitives. assets owns embedded recipes, playbooks and versions. project owns discovery, projection, validation and authoring. checks owns JUnit coverage, hooks, boundaries, history and doctor. workflow owns claims, lifecycle, provenance and context. runner owns agents, queues, retries, recovery and drills. app composes the CLI. SPEC.md frontmatter is the only graph source.

## Conventions

Module path: github.com/lurkshark/fireweed. Thin api.go facades and public contracts are retained interfaces; all behavior belongs under internal/. Evals import public packages only and use temporary repositories, explicit fixture Git identities, deterministic local hooks and bounded fake agents. No runtime network calls, bootstrap delegation, or private implementation tests. Keep dependencies minimal (YAML parser only initially). Never import a public parent from its own internal implementation; shared data belongs in contracts.

## Controller and product

The pinned controller is /Users/joel/Workspace/fireweed/.fireweed/bin/fireweed-bootstrap. The independently built subject is /Users/joel/Workspace/fireweed/bin/fireweed. Invoke explicit paths. FIREWEED_BIN during hooks always names the governing controller. Doctor and drill worktrees use protected Go hooks and graph metadata, including randomly named modules.

## Author and worker policy

This local goal authorizes authored changes and local commits. Manual execution on main uses that authority; it does not assert a human merge. Keep normal product approval enforcement intact. Authors edit protected setup/infrastructure/evidence only between active changes and commit separately. Only fireweed applies Specs/Decisions, phase commits and archives. Workers obey context and phase/module zones, never commit or weaken enforcement, and never edit protected planning/design/tooling. Audit stays enabled; separate restricted auditor context checks exact spec clauses against eval assertions. Regeneration agents receive retained contracts, evals and allowed dependency interfaces without old implementation/history/bootstrap as an oracle. No pushes, publication or release tags.
