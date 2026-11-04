---
audit_pending:
    - AST-050
    - AST-056
audit_rounds: 1
bases:
    AST-004: a9742599
module: assets
phase: audit
run: 20261104T165115-6eb1fbc842333208fb055f7a
---

## Motivation

ACCEPTANCE C7 selects the design's TypeScript/pnpm and Python/uv recipes.
Each must include usable project guidance, module templates, canaries and
default hooks, with an actual fresh-init doctor proof. Init's generic built-in
selection is now applied; add both bundles coherently in their owning module.

## Interface

Keep the public Recipe model and BuiltinRecipe signature unchanged. Supply
both named bundles through the existing data interface and embedded files.

## Added

### AST-049
BuiltinRecipe for typescript-pnpm and python-uv MUST provide nonempty instructions, interface/evals/disposable defaults, eval/boundaries/setup commands, protected glue, language lockfile worker_writable entries and project-relative instruction, hook, module-template and canary files.

### AST-050
The additional recipe instructions MUST explain stack prerequisites and workspace setup, using pnpm workspaces with package exports for TypeScript and uv workspaces with an import linter for Python. They MUST explain deriving the module graph from FIREWEED_BIN modules --json rather than maintaining another dependency graph.

### AST-051
The additional recipes' module templates and canary snippets MUST support name, prefix, package, module_dir and other_dir substitutions used by doctor.

### AST-052
Each additional recipe MUST provide a version-1 canary map containing baseline, undeclared, internal_eval, failing_eval and missing_id entries.

### AST-053
Rendered baseline modules for typescript-pnpm and python-uv MUST pass their declared eval and boundaries hooks.

### AST-054
Each additional recipe's eval hook MUST run the requested module's language evals and write JUnit testcase names, classnames, pass/fail/skip and elapsed seconds while preserving test failure exit status.

### AST-055
The additional eval hooks MUST retain started-but-unfinished tests and collection or compile failures as visible failures rather than passing or empty evidence.

### AST-056
The additional boundaries hooks MUST query FIREWEED_BIN modules --json and derive their checks from its module directories, depends_on and resolved zone map.

### AST-057
The additional boundaries hooks MUST report module-prefixed violations for imports of undeclared modules and for eval imports reaching beyond retained interfaces, including their own implementation.

### AST-058
Each supplied additional-recipe canary MUST expose its named violation through the corresponding hook. The hooks MUST pass cleanly again after the canary is removed.

## Modified

### AST-004
BuiltinRecipe MUST return the embedded recipes for go, typescript-pnpm and python-uv and reject unknown names.

## Removed

## Decisions

## Notes

Fail-first: `GOCACHE=/tmp/fireweed-rebuild-cache go test ./packages/assets/evals -run 'Test_AST_(004_Additional|049_|050_|051_|052_|053_|054_|055_|056_|057_|058_)' -count=1` failed for every assigned ID because the current `BuiltinRecipe` rejects `typescript-pnpm` and `python-uv`. Missing behavior is implementation-writable in `packages/assets/internal/recipe/recipe.go` and new embedded files under `packages/assets/internal/recipe/files/`; no retained interface edit is needed.

Evals step: `PATH=/Users/joel/Workspace/fireweed/.fireweed/bin:$PATH GOCACHE=/tmp/fireweed-rebuild-cache /Users/joel/Workspace/fireweed/.fireweed/bin/fireweed verify --quick` passed (`verify: evals`).

Audit-gap correction: AST-050 now asserts guidance against a second dependency graph; AST-056 now varies a module directory and a resolved interface zone. `GOCACHE=/tmp/fireweed-rebuild-cache go test ./packages/assets/evals -run 'Test_AST_(050_|056_)' -count=1` still fails for both IDs on the missing built-ins, and pinned `fireweed verify --quick` passed (`verify: evals`).

Preserve every existing Go recipe assertion, hook behavior, payload and other
assets behavior. Keep specs atomic and new metadata consistent with the files
and instructions actually returned. Use real language packages and protected
glue, without a second copy of the module graph. Existing Go/current project
setup is not a target for speculative cleanup.

C7 prerequisites were checked without installing tools: pnpm 10.15.1 is already
cached and runs when package.json selects that packageManager version; uv is
available. Use that existing pnpm runtime for the TypeScript proof, avoiding
the unpinned Corepack default's registry lookup. Do not install or replace pnpm
or uv, change models/effort, weaken canaries, opt out of boundaries or raise
bounds/timeouts. Ordinary documented project/workspace setup follows each recipe.

Before completing implementation, build the NEW subject and prove EACH named
recipe in its own fresh scratch Git project created by init --recipe <name>.
Follow the returned recipe's prerequisite/workspace instructions, then run
doctor --no-agent with all hook canaries and recoveries enabled. Setup copies
the invoking controller into scratch, and the current scratch .fireweed/bin
comes first on worker PATH. Use the prescribed Go cache and exact configured
agent argv. Record the actual commands and outcomes concisely in this change's
Notes; temporary paths, extra hashes or supplemental evidence files are not
needed. Neither proof can be replaced by a local-recipe alias or a drill.

Use public Recipe/bundle evals and representative temporary fixtures for both
stacks, including real hook results and their cleanup. Keep genuine fail-first
for every assigned ID, separate restricted audit, frozen evals/interfaces and
normal full verification. Keep own Verify under five minutes and full impl
Verify under twenty-five. Do not add process instrumentation or redundant
expensive fixture matrices. Assets evals must respect their module boundary;
the fresh-init CLI/doctor proofs are separate native integration checks.

Do not edit other modules, root config, protected setup, module SPEC/Decisions,
bounds or timeouts. Workers run pinned fireweed context, verify with fireweed
verify --quick, poll until terminal and return. They never commit, invoke
lifecycle mutators in the held checkout, edit run/phase/Provenance or alter
frozen evals/interfaces. Only the outside owning controller advances, commits,
applies and archives. C8 remains unstarted until all of C7 is complete.

The separate restricted auditor edits only the existing Audit body. Findings
use - ID: "exact clause" — missing assertion. Passing audit writes No findings
and Checked IDs. Re-audits cover only flagged IDs, retaining prior clearances
and all counters/bounds.

## Provenance
- evals 2026-11-04T17:00:02.536427Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=117 failing=31 passing=86
- audit 2026-11-04T17:01:16.619258Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
- evals 2026-11-04T17:04:41.518012Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=117 failing=31 passing=86
## Audit

- AST-050: "They MUST explain deriving the module graph from FIREWEED_BIN modules --json rather than maintaining another dependency graph." — No assertion checks the instruction against maintaining another dependency graph.
- AST-056: "The additional boundaries hooks MUST query FIREWEED_BIN modules --json and derive their checks from its module directories, depends_on and resolved zone map." — The tests vary depends_on, but no assertion checks that changing module directories or resolved zones changes the boundary result.
