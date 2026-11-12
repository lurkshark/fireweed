package content

import "fmt"

const commonHelpOptions = `
Common options
  --help, -h  Show this command's help page.
  --json      Return output in the schema-1 JSON envelope.
`

func CommandHelp(name string) (string, error) {
	page, ok := commandPages[name]
	if !ok {
		return "", fmt.Errorf("unknown command %q", name)
	}
	return page + commonHelpOptions, nil
}

var commandPages = map[string]string{
	"help": `Show the Fireweed command overview or the help page for one command.
Works outside a project.

Usage
  fireweed help [--json]
  fireweed --help [--json]
  fireweed -h [--json]
  fireweed help <command> [--json]
  fireweed <command> help [--json]
  fireweed <command> --help [--json]
  fireweed <command> -h [--json]

The first three forms show the overview. The remaining forms take a command
name and show its page. Square brackets mark optional inputs.

Examples
  fireweed help
  fireweed help agents-md
  fireweed doctor --help

JSON output is one object with schema set to 1, data for the command result,
and a diagnostics array. Help text is a string in data. Usage and configuration
failures return null data and diagnostics with codes and messages.

Exit statuses
  0  Success or empty queue.
  1  Failed checks or blocked execution.
  2  Usage or configuration failure.
  3  Clean early stop on a time or agent usage limit.

See fireweed help explain for diagnostic codes and fireweed help playbook for
workflow instructions.
`,
	"version": `Print the installed Fireweed binary version. Works outside a project.

Usage
  fireweed version [--json]

Example
  fireweed version --json

See fireweed help init for project setup.
`,
	"init": `Initialize Fireweed in the current directory, which must lack fireweed.yaml.
Creates configuration and recipe files, scaffolds SYSTEM.md if needed, updates
AGENTS.md and .gitignore, and prints setup instructions. It runs no hooks or agents.

Usage
  fireweed init --recipe <recipe> --agent <agent> [--audit-agent <agent>] [--no-app] [--json]

Options
  --recipe <recipe>       Required recipe selection. Built-ins are go,
                          typescript-pnpm, and python-uv. Use none for a minimal
                          setup with hooks, templates, and canaries left for
                          you to configure. A local recipe directory path,
                          such as ./recipe or ../recipe, must contain
                          recipe.yaml and recipe.md. An optional files/
                          directory supplies project files to copy.
  --agent <agent>         Required worker command. Use codex, claude, or a quoted
                          shell command template containing {prompt} for inline
                          prompt text or {prompt_file} for the prompt file path.
  --audit-agent <agent>   Optional separate auditor command. It accepts the same
                          forms as --agent. Auditing is enabled and falls back
                          to the worker command when this option is omitted.
  --no-app                Skip scaffolding the default packages/app module.

Example
  fireweed init --recipe go --agent codex --audit-agent claude

See fireweed help doctor to check the setup and fireweed help new-module to
scaffold additional modules.
`,
	"doctor": `Check a Git project's configuration, managed AGENTS.md guidance, and playbook
overrides. Run setup, hook canaries and recovery checks, and worker and auditor
probes in a temporary worktree. The caller's checkout stays unchanged.

Usage
  fireweed doctor [--no-agent] [--json]

Option
  --no-agent  Skip worker and auditor probes. Hook canary and recovery checks still run.

Example
  fireweed doctor --no-agent

See fireweed help ci for deterministic checks and fireweed help agents-md to
refresh managed guidance.
`,
	"new-module": `Scaffold a module in an initialized project. Writes SPEC.md and other files
from the project's module templates into a new directory under a configured
module glob such as packages/*. It runs no hooks or agents.

Usage
  fireweed new-module <name> --prefix <prefix> [--json]

Input
  <name>             Required unused lowercase module name. Start with a letter.
                     Use letters and digits, with hyphens between segments.
  --prefix <prefix>  Required unused spec ID prefix. Start with an uppercase
                     letter and use uppercase letters and digits.

Example
  fireweed new-module ledger --prefix LED

See fireweed help modules to inspect existing names and prefixes.
`,
	"new-change": `Create the next numbered change file in the project's pending queue.
Prints the new path and runs no hooks or agents.

Usage
  fireweed new-change <module> <slug> [--json]

Arguments
  <module>  Required name of an existing module.
  <slug>    Required unused slug containing only lowercase letters, digits,
            and hyphens.

Example
  fireweed new-change ledger add-balance

See fireweed help playbook for authoring instructions and fireweed help stamp
to record the base hashes after drafting spec edits.
`,
	"stamp": `Record base hashes for Modified and Removed specs against the projection
before the selected pending change. Updates the change file in the project and
runs no hooks or agents. Added handles receive their IDs at start.

Usage
  fireweed stamp <change> [--json]

Argument
  <change>  Required pending change slug, filename, or repository-relative path.

Example
  fireweed stamp add-balance

See fireweed help check for queue validation and fireweed help show for stale
base details.
`,
	"list": `List the project's pending change queue in order, with each change's module,
phase, blocked reason, and audit rounds. Read-only.

Usage
  fireweed list [--json]

Example
  fireweed list --json

See fireweed help status for queue eligibility and project-wide details.
`,
	"status": `Summarize project state, including why changes are waiting, blocked changes,
audit rounds, quarantines, boundary opt-outs, drift since the last drill,
unapproved changes, and budget metrics and warnings. Read-only.

Usage
  fireweed status [--json]

Example
  fireweed status

See fireweed help list for the pending queue and fireweed help drill for
reconstruction results.
`,
	"show": `Inspect a project's current specs, projected specs, archived spec history,
or stale change bases. Read-only.

Usage
  fireweed show <module> [<ID>] [--projected | --history] [--json]
  fireweed show --stale <change> [--json]

Arguments and options
  <module>          Required module name for the spec and history views.
  <ID>              Optional spec item ID. Omit it for the whole module.
  --projected       Show specs as they would read after applying the pending queue.
  --history         Show archived changes and provenance for the module or item.
  --stale <change>  Compare recorded bases with the preceding projection for one
                    pending change, selected by slug, filename, or repository-relative
                    path. Shows the stamped text when it can be recovered.

The --projected, --history, and --stale modes are mutually exclusive.

Examples
  fireweed show ledger LED-001 --history
  fireweed show --stale add-balance

See fireweed help modules to find module names and fireweed help stamp for
recording base hashes.
`,
	"modules": `Show the project's resolved module dependency graph and zone map, including
directories and ID prefixes. When this worktree holds a change, the command
uses its approved projected graph. Read-only.

Usage
  fireweed modules [--json]

Example
  fireweed modules --json

See fireweed help show for module specs.
`,
	"context": `Show worker context for the current phase of a change claimed by this
worktree in a Git project. Includes the change, projected spec, interface,
SYSTEM.md, and playbook, with paths and sizes for other relevant files. Read-only.

Usage
  fireweed context <change> [--full] [--json]

Argument and option
  <change>  Required held change slug, filename, or repository-relative path.
  --full    Include the contents of the other relevant files.

Example
  fireweed context add-balance --full

See fireweed help start to claim a change and fireweed help playbook for
step instructions.
`,
	"explain": `Look up a diagnostic code to learn what caused it and how to address it.
Works outside a project.

Usage
  fireweed explain <code> [--json]

Argument
  <code>  Required code from a Fireweed diagnostic, such as C10.

Example
  fireweed explain C10

See fireweed help playbook for workflow guidance.
`,
	"playbook": `Read built-in workflow instructions for authoring, lifecycle steps, or
reconstruction. Works outside a project.

Usage
  fireweed playbook <name> [--json]

Argument
  <name>  Required playbook name. Choose one of the following.

  author-change  Draft a change with a clear motivation and observable specs.
  evals          Write tests and collect fail-first evidence when required.
  audit          Review evals for assertions that cover the changed specs.
  impl           Implement the change in disposable code and verify it.
  regenerate     Rebuild disposable code from the retained specs, interface, and evals.

Example
  fireweed playbook author-change

See fireweed help context for the playbook and files for an active step.
`,
	"check": `Run static validation of project configuration, specs, changes, and the
pending queue. It runs no hooks or agents. By default it only reports diagnostics.

Usage
  fireweed check [--fix] [--json]

Option
  --fix  Fill missing bases for Modified and Removed items in pending change
         files before validation. Existing hashes are preserved.

Example
  fireweed check --fix

See fireweed help stamp for stamping one change and fireweed help ci for
running evals and other deterministic checks.
`,
	"coverage": `Run a project's eval hook for one module, map JUnit testcases to spec IDs,
and check required coverage and passing results. Uses a temporary JUnit report.
It invokes no agents and leaves the change phase as it is.

Usage
  fireweed coverage <module> [--with-change <change>] [--json]

Argument and option
  <module>                Required name of the module to evaluate.
  --with-change <change>  Check against the projection through a pending change
                          selected by slug, filename, or repository-relative path.
                          By default, check current specs.

Example
  fireweed coverage ledger --with-change add-balance

See fireweed help verify for checks on an active change.
`,
	"check-commits": `Check a Git project's history for violations of protected paths, phase
zones, append-only files, and archived change files (H01–H04). Read-only.

Usage
  fireweed check-commits <range> [--json]
  fireweed check-commits [--json]

Argument
  <range>  Optional Git revision range. By default, check from the merge base
           with the configured base ref through HEAD, using the locally stored
           remote-tracking ref when available and the local branch otherwise.

Example
  fireweed check-commits main..HEAD

See fireweed help ci to combine history checks with other checks, and
fireweed help explain for diagnostic guidance.
`,
	"ci": `Run the Git project's deterministic checks. Validates specs and changes,
checks history, runs every module's evals and configured boundary checks, and
runs doctor canaries in scratch by default. Agent probes are disabled.

Usage
  fireweed ci [<range>] [--no-doctor] [--json]

Argument and option
  <range>      Optional Git revision range for history checks. Defaults to the
               range from the merge base with the configured base ref to HEAD.
  --no-doctor  Skip doctor canaries and their setup. Other checks still run.

Example
  fireweed ci main..HEAD

See fireweed help check-commits for history checks and fireweed help doctor
for setup diagnostics.
`,
	"pr-body": `Print a pull request description for a range in the Git project, including
change summaries, audit results, spec decisions, and review or budget warnings.
Read-only.

Usage
  fireweed pr-body <range> [--json]

Argument
  <range>  Required Git revision range in the form <from>..<to>.

Example
  fireweed pr-body main..HEAD

See fireweed help check-commits to validate the same history.
`,
	"start": `Claim the first eligible queued change in a clean Git worktree and print its
step context. Records the phase, run, and assigned spec IDs in the change file.
Starting a fresh regeneration change also removes the module's disposable files.
Start runs no hooks or agents and makes no commit.

Usage
  fireweed start <change> [--from-attempt <N>] [--json]

Argument and option
  <change>            Required slug, filename, or repository-relative path of the
                      first eligible change in the queue.
  --from-attempt <N>  Restore saved work for the change's current phase from
                      attempt N, a positive integer.

Example
  fireweed start add-balance

See fireweed help context for step inputs and fireweed help verify for checks.
`,
	"verify": `Run checks for the current step of a held change in the Git project. Runs
applicable eval and boundary hooks, and can write run reports. It invokes no
agents and makes no phase commit.

Usage
  fireweed verify [<change>] [--quick] [--json]

Argument and option
  <change>  Optional slug, filename, or repository-relative path. Defaults to
            the change held by this worktree.
  --quick   Skip dependent-module evals and boundary checks in implementation.

Example
  fireweed verify --quick

See fireweed help done to advance the phase after full verification.
`,
	"done": `Complete the held change's current step in the Git project. Runs full
verification, including applicable hooks, then commits the phase transition.
Completing implementation applies spec and Decision edits and archives the
change. It runs no agents.

Usage
  fireweed done [<change>] [--json]

Argument
  <change>  Optional slug, filename, or repository-relative path. Defaults to
            the change held by this worktree.

Example
  fireweed done

See fireweed help verify to check the phase before committing.
`,
	"next": `Run the configured agents and verification hooks through the steps of one
change in a Git project. Resume the change held by this worktree or start the
first eligible queued change. Records phase commits and archives the change on
success. An empty or entirely waiting queue is a no-op.

Usage
  fireweed next [--allow-unmerged] [--json]

Option
  --allow-unmerged  Start a change without approval on either base branch.
                    By default it must be on the local or shared base.

Example
  fireweed next

See fireweed help run to process multiple changes and fireweed help status
to see why changes are waiting.
`,
	"run": `Process the Git project's queue until no eligible change remains, a bound
is reached, or execution stops. A blocked change leaves independent changes
eligible to run. The final result reports the block as a failure.
Runs agents and verification hooks, commits phase transitions, and archives
applied changes.

Usage
  fireweed run [--allow-unmerged] [--max <N>] [--for <duration>] [--json]

Options
  --allow-unmerged  Start changes without approval on either base branch.
                    By default a change must be on the local or shared base.
  --max <N>         Apply at most N changes. N must be a positive integer.
  --for <duration>  Set a time limit for the whole run. Use a positive Go
                    duration such as 1h30m. A clean time stop returns exit 3.

Example
  fireweed run --max 3 --for 1h

See fireweed help next for single-change execution.
`,
	"release": `Drop a change claim held by this worktree. Requires a Git project. Working
files, recorded change state, and HEAD stay as they are. Runs no hooks or agents.

Usage
  fireweed release <change> [--json]

Argument
  <change>  Required slug, filename, or repository-relative path of the change
            claimed by this worktree.

Example
  fireweed release add-balance

See fireweed help start to claim work and fireweed help abandon to undo an
unapplied change.
`,
	"commits": `List full commit IDs for the change's phase transitions on the current
branch, newest first. Requires a Git project and is read-only.

Usage
  fireweed commits <change> [--json]

Argument
  <change>  Required repository-relative path, filename, or unambiguous slug
            for a pending or archived change.

Example
  fireweed commits add-balance

See fireweed help check-commits to check history rules.
`,
	"abandon": `Undo the phase commits for an unapplied change and remove its pending file
in one new commit. Requires a clean Git worktree. The change must be unclaimed
or claimed by this worktree. Runs no hooks or agents.

Usage
  fireweed abandon <change> [--json]

Argument
  <change>  Required slug, filename, or repository-relative path for the pending change.

Example
  fireweed abandon add-balance

See fireweed help commits to inspect the affected commits and fireweed help
release to drop just the claim.
`,
	"drill": `Measure whether a module can be rebuilt from retained material. Requires a
clean Git worktree with no live change claims. Runs setup and reconstruction
in scratch spaces, then runs the module's evals and the evals of all its
dependents. Writes .fireweed/drills/<module>.md and retains accepted rebuilt
code in a drill ref.

Usage
  fireweed drill <module> [--json]

Argument
  <module>  Required name of an existing module to rebuild using the configured
            worker agent. The invoking implementation and HEAD stay in place.

Example
  fireweed drill ledger

See fireweed help playbook for reconstruction guidance and fireweed help gc
for drill ref retention.
`,
	"gc": `In a Git project, gc prunes stale claims, attempt refs for archived changes, and drill refs older than gc.drill_max_age.
Set gc.drill_max_age to a positive Go duration in fireweed.yaml.
The gc.drill_max_age default is 720h (30 days).
Prints pruned Git refs and leaves working files alone. Runs no hooks or agents.

Usage
  fireweed gc [--json]

Example
  fireweed gc --json

See fireweed help release to drop a claim owned by this worktree.
`,
	"migrate": `Check the project's file format. This version validates format 1 and reports
that it is current. Other format versions return an error. Runs no hooks or
agents and changes no files.

Usage
  fireweed migrate [--json]

Example
  fireweed migrate

See fireweed help check for static validation.
`,
	"agents-md": `Refresh the Fireweed managed guidance in the project's AGENTS.md.
Creates the file if needed and preserves text outside the managed markers.
Runs no hooks or agents.

Usage
  fireweed agents-md --write [--json]

Option
  --write  Required to write or update the managed block.

Example
  fireweed agents-md --write

See fireweed help doctor to check managed guidance and other project setup.
`,
}
