---
bases:
    AST-004: a9742599
module: assets
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

### new:additional-recipe-layout
BuiltinRecipe for typescript-pnpm and python-uv MUST provide nonempty instructions, interface/evals/disposable defaults, eval/boundaries/setup commands, protected glue, language lockfile worker_writable entries and project-relative instruction, hook, module-template and canary files.

### new:additional-recipe-instructions
The additional recipe instructions MUST explain stack prerequisites and workspace setup, using pnpm workspaces with package exports for TypeScript and uv workspaces with an import linter for Python. They MUST explain deriving the module graph from FIREWEED_BIN modules --json rather than maintaining another dependency graph.

### new:additional-recipe-bindings
The additional recipes' module templates and canary snippets MUST support name, prefix, package, module_dir and other_dir substitutions used by doctor.

### new:additional-recipe-canary-maps
Each additional recipe MUST provide a version-1 canary map containing baseline, undeclared, internal_eval, failing_eval and missing_id entries.

### new:additional-recipe-baselines
Rendered baseline modules for typescript-pnpm and python-uv MUST pass their declared eval and boundaries hooks.

### new:additional-recipe-evals
Each additional recipe's eval hook MUST run the requested module's language evals and write JUnit testcase names, classnames, pass/fail/skip and elapsed seconds while preserving test failure exit status.

### new:additional-recipe-eval-errors
The additional eval hooks MUST retain started-but-unfinished tests and collection or compile failures as visible failures rather than passing or empty evidence.

### new:additional-recipe-graph
The additional boundaries hooks MUST query FIREWEED_BIN modules --json and derive their checks from its module directories, depends_on and resolved zone map.

### new:additional-recipe-violations
The additional boundaries hooks MUST report module-prefixed violations for imports of undeclared modules and for eval imports reaching beyond retained interfaces, including their own implementation.

### new:additional-recipe-recovery
Each supplied additional-recipe canary MUST expose its named violation through the corresponding hook. The hooks MUST pass cleanly again after the canary is removed.

## Modified

### AST-004
BuiltinRecipe MUST return the embedded recipes for go, typescript-pnpm and python-uv and reject unknown names.

## Removed

## Decisions

## Notes

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

