---
bases:
    AST-008: 581b745d
    AST-048: 9a5d6a47
module: assets
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

### new:command-help-content
CommandHelp MUST return a command page containing its purpose, usage syntax,
descriptions of its accepted arguments and options, and at least one invocation
example for each command listed by Help. Each page MUST document --help, -h
and --json and point to related command help. Unknown command names MUST return
an error naming the rejected command.

### new:help-navigation
CommandHelp("help") MUST describe the overview, help <command>, <command> help,
<command> --help and <command> -h forms, and the schema-1 JSON envelope.
It MUST explain exit statuses 0 (success or empty queue), 1 (failed checks or
blocked execution), 2 (usage or configuration failure) and 3 (clean early stop).

### new:playbook-discovery
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

