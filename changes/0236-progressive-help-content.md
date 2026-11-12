---
audit_pending:
    - AST-008
    - AST-061
audit_rounds: 1
bases:
    AST-008: 581b745d
    AST-048: 9a5d6a47
module: assets
phase: evals
run: 20261112T192706-b138d77f5fa82c3569d46a5e
---

## Motivation

The owner could not discover command usage through `fireweed help agents-md`,
`fireweed agents-md help` or `fireweed doctor help`, and prefers progressive
disclosure. The current overview lists commands but provides no route to their
usage, while exposing gc configuration details at the top level. Supply a short
overview and useful command pages; the following app change exposes them.

## Interface

Add `CommandHelp(name string) (string, error)` to the assets facade, delegating
to disposable implementation. Retain `Help() string` as the overview. Command
pages include help itself and every command in the overview.

## Added

### AST-061
CommandHelp MUST return a command page containing its purpose, usage syntax,
descriptions of its accepted arguments and options, and at least one invocation
example for each command listed by Help. Each page MUST document --help, -h
and --json and point to related command help. Unknown command names MUST return
an error naming the rejected command.

### AST-062
CommandHelp("help") MUST describe the overview, help <command>, <command> help,
<command> --help and <command> -h forms, and the schema-1 JSON envelope.
It MUST explain exit statuses 0 (success or empty queue), 1 (failed checks or
blocked execution), 2 (usage or configuration failure) and 3 (clean early stop).

### AST-063
CommandHelp("playbook") MUST list author-change, evals, audit, impl and
regenerate with their purposes and show how to read one. CommandHelp("explain")
MUST describe diagnostic-code lookup with a concrete code example and point to
playbook help for workflow guidance.

## Modified

### AST-008
Help MUST identify fireweed and list every command in the design's CLI table,
including help, version, explain and playbook, grouped by purpose. It MUST
advertise help <command> and <command> --help, with a concrete command example.
It MUST keep command-specific option and configuration descriptions on the
command pages rather than the overview.

### AST-048
CommandHelp("gc") MUST document that gc prunes stale claims, attempt refs for
archived changes and drill refs older than gc.drill_max_age. It MUST identify
gc.drill_max_age as a positive Go duration in fireweed.yaml with default 720h
(30 days).

## Removed

## Decisions

## Notes

Use the design CLI table and current app command contracts to author accurate
pages. Cover options absent from the abbreviated design table too:
init --audit-agent, doctor --no-agent, new-module --prefix, show --stale,
--history and --projected, context --full, check --fix, coverage --with-change,
ci --no-doctor, start --from-attempt, verify --quick, next/run --allow-unmerged,
run --max and --for, and agents-md --write. Explain required versus optional
inputs, mutually exclusive show modes, and the effect of options. Init help
lists built-in recipes, local recipe paths and supported agent forms.

Keep the overview compact. The help page can explain common conventions;
individual pages should give enough information to use their own command,
including whether it needs a project and whether it runs hooks or agents or
changes files. Examples are documentation, not commands to execute in evals.
Do not make help a full copy of the design document.

Update the AST-048 eval to query the gc page while retaining its assertions.
Exercise the public assets API with lightweight text tests. Assert useful
content and syntax, not exact prose or whitespace. New interface code delegates
to an implementation-writable disposable stub during evals. No recipe,
playbook, diagnostic-allocation or unrelated module changes are needed.

The focused assets evals failed against the existing overview and temporary
stub for AST-008, AST-048 and AST-061 through AST-063. The controller's
`fireweed verify progressive-help-content` passed in the evals phase.

The command pages belong in the disposable
`packages/assets/internal/content/command_help.go` stub file. The overview
lives in `packages/assets/internal/content/content.go`, also disposable.

## Provenance
- evals 2026-11-12T19:34:26.323341Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=135 failing=8 passing=127
- audit 2026-11-12T19:37:54.541Z agent=manual fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
## Audit

- AST-008: "command-specific option and configuration descriptions" — `Test_AST_008_Help` excludes gc configuration details, but no assertion excludes command-specific option descriptions from the overview.
- AST-061: "descriptions of its accepted arguments and options" — No assertion checks descriptions of the accepted arguments for `start`, `release`, `commits` or `drill`. Their page assertions check purpose, argument-shaped syntax and examples, but `Test_AST_061_CommandInputDetails` does not check these commands' argument descriptions.
