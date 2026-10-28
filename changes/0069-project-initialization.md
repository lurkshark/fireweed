---
bases: {}
module: app
---



## Motivation

Milestone 2 needs a project created by NEW before doctor and bounded execution.
The design's Distribution and init first-run flow writes setup files, copies a
recipe, creates an optional composition root and prints instructions for the
agent to complete the glue. Wire this through the CLI as one cohesive init
change, using the embedded Go recipe and supporting none/local recipe inputs.
Do not run init over the rebuild checkout.

## Interface

The public app Run signature stays unchanged. New behavior is in app internal
implementation; use the existing assets.Recipe interface and YAML dependency.
No other module interface or implementation change is required.

## Added

### new:init-configuration
init MUST create fireweed.yaml in the current directory with format version 1, a binary range containing assets.Version, base main, remote origin and modules [packages/*].

### new:init-recipe-config
init MUST write the selected recipe's module_defaults, hooks, protected and worker_writable values into the generated configuration.

### new:init-system
init MUST create an editable SYSTEM.md project scaffold when absent and retain an existing SYSTEM.md unchanged.

### new:init-agent-instructions
init MUST create or update the fireweed managed AGENTS.md block, pointing to author-change and context, while preserving text outside the markers.

### new:init-go-payload
init --recipe go MUST copy assets.BuiltinRecipe(go)'s Files into their project-relative destinations without changing their contents.

### new:init-instructions
Successful init MUST print the selected recipe's instructions on stdout and return exit 0 with empty stderr.

### new:init-none
init --recipe none MUST write minimal setup without language hooks, templates or canaries and print the eval, boundaries and setup hook contract.

### new:init-local-metadata
init --recipe ./directory MUST load recipe.yaml module_defaults, hooks, protected and worker_writable metadata and recipe.md instructions from that directory.

### new:init-local-payload
init with a local recipe MUST copy the directory's files/ contents into matching project-relative destinations and copy its instructions to .fireweed/recipes/local.md.

### new:init-app-module
By default init MUST create a normal packages/app module named app with prefix APP and no dependencies, rendering available module templates with name app, prefix APP and package app.

### new:init-no-app
init --no-app MUST skip creation of the app module.

### new:init-fallback-module
When the selected recipe supplies no module templates, init MUST create a minimal app SPEC.md with empty Specs and Decisions sections.

### new:init-agent-config
init MUST configure codex or claude command presets, or preserve a custom agent command template containing {prompt}, in agent.command.

### new:init-auditor-config
init MUST enable audit by default and write a separately selected --audit-agent command when supplied, otherwise leaving audit.command unset for agent fallback.

### new:init-usage
init MUST reject missing --recipe or --agent values, malformed options or extra positional arguments with corrective U01 and exit 2 before creating files; missing-input messages MUST list valid recipe and agent choices.

### new:init-recipe-rejections
init MUST return U02 and exit 2 for an unknown recipe name, C03 for unreadable or malformed local recipe documents, and C12 for invalid local recipe metadata types, without creating setup files.

### new:init-existing-project
init MUST refuse an existing fireweed.yaml with P01 and exit 2, preserving existing project files.

### new:init-conflicting-files
init MUST refuse conflicting generated recipe or module files with P01 and exit 2 rather than overwrite them.

### new:init-ignore
init MUST retain existing .gitignore entries and ensure .fireweed/runs/ is ignored so lifecycle reports do not become forbidden worktree edits.

### new:init-no-execution
init MUST copy setup and print instructions without executing configured agents or hooks.

### new:init-json
init --json MUST return schema 1 with instruction text as string data and empty diagnostics on success, or null data and code/message diagnostics on failure, with empty stderr.

## Modified



## Removed



## Decisions

Init accepts codex and claude presets as argv commands; other selections are shell command templates containing {prompt}, stored as strings for one-pass quoted substitution when execution is implemented. Local recipes use recipe.yaml metadata, recipe.md instructions and a files/ project-relative payload; none uses empty language defaults and a generic module scaffold.

## Notes

Read the design first-run/recipe clauses and retained assets API. Built-in Go
files/templates are already implemented; consume them through assets rather
than duplicate their payload in app. Codex preset is [codex, exec, {prompt}];
Claude preset follows the design's [claude, -p, {prompt}, --permission-mode,
acceptEdits]. The same selection rules apply to --audit-agent. Keep actual
rebuild workers on the existing configured Sol/high command; init tests only
inspect generated config and use deterministic local commands, never auth.

Require explicit recipe and agent selections. Accept the init flags only for
init; existing commands retain their behavior and errors. Go is the currently
available built-in; other language bundles remain later recipe work. For none
use empty language defaults/hooks and the generic optional app module. A local
manifest uses the same key names as configuration for the four recipe metadata
fields; use existing zone list shapes. Copy instructions into the project and
print them. Render the selected .fireweed/templates/module tree for app (strip
.tmpl from filenames, retain SPEC.md); use the three documented substitutions.
For Go create the declared interface/disposable/eval layout. Preserve existing
SYSTEM and non-managed AGENTS text. Validate selections and payload destinations
before overwriting files; no transactional-I/O or platform edge-case framework.

The old APP-014 fixture currently chooses init as an unimplemented example.
With init implemented, retarget that example and its JSON call to still-pending
status/gc, preserving the generic U01/unimplemented/JSON assertions. This is the
planned feature transition, not permission to weaken any other existing test.
APP-015's init-without-required-inputs failure still needs its existing envelope.

Use a few successful Go/none/local setup fixtures to assert grouped independent
outputs, --no-app, preserved instructions and custom auditor configuration.
Reuse fixture/build/assertion helpers. Add representative usage/rejection cases,
not a every-field/command matrix. A local recipe with sentinel hook/agent commands
can show init does not execute them. Exercise generated Go metadata through NEW
modules/check in a fixture with an existing minimal go.mod. Each new ID needs
genuine runtime fail-first evidence; compilation errors/skips do not qualify.
Keep own Verify under five minutes and full Verify under fifteen. Never invoke
init with real configuration against this checkout. Do not implement doctor,
new-module, agents-md, runner behavior or new language bundles here. Do not
address drill gaps, rerun a drill or consult bootstrap/deleted implementation.
Controller owns spec application, phase commits and archive.
