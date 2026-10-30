# Rebuild plan

How the agent works. [PROGRESS.md](PROGRESS.md) holds live state, [SYSTEM.md](../../SYSTEM.md)
the guardrails, [ACCEPTANCE.md](ACCEPTANCE.md) the full inventory and final gates.

## Route

The manual lifecycle, setup, `next`, bounded `run`, recovery, all change kinds
and most of the reading surface are done (see PROGRESS). What remains:

1. Remaining commands and options: `drill`, `gc`, `migrate`, `show --stale`,
   `check --fix`.
2. Recipes, version-range enforcement, format migration and packaging.
3. Final gates: real drills of every disposable module, isolated self-hosting,
   reference smoke, fresh-checkout checks and delivery.

## Choosing and closing a change

- Name the CLI scenario and missing behavior in the change Motivation. Target
  one module; batch related work in it with no cap on spec count, but keep each
  spec atomic and the change reviewable.
- A dependency can land as its own change, but return to its caller as soon as
  the interface works.
- Every slice ends with an actual NEW command and an inspectable result. Prefer
  a permanent app eval that drives NEW over a one-off manual run.
- Fix frozen evals with `eval_fix` and behavior with normal changes. Keep
  `explain` accurate for codes actually emitted; don't pre-allocate diagnostics.

## Controller and product

| Role | Path | Use |
| --- | --- | --- |
| Pinned bootstrap controller | `.fireweed/bin/fireweed-bootstrap` | Claims, checks, commits, application and archives in this checkout |
| NEW | `bin/fireweed`, built with `go build -o bin/fireweed .` | Product fixtures and isolated self-hosting |

Worker commands named `fireweed` resolve to the controller; product fixtures
select NEW. `FIREWEED_BIN` names the binary that invoked a hook. Never
substitute the controller for product acceptance. The controller governs this
checkout until acceptance; NEW self-hosts only in an isolated clone.

## Per-change loop

1. As author: read SYSTEM, the projected specs, the relevant design clauses and
   `playbook author-change`. Run `new-change`, write Added specs with `new:`
   handles (or Modified/Removed/`fixes`), `stamp`, `check`, and commit the
   authored change separately.
2. Controller `start <slug>`, then `context <slug>` before each worker step.
3. Evals step: permanent black-box evals, the described interface and allowed
   stubs. Every Added/Modified/fixes ID needs a genuinely failing case unless
   `already_satisfied`. Iterate with `verify --quick`; `done` runs full checks.
4. Audit step: the controller's audit step. Findings are exact
   spec clauses no test asserts. Re-audits cover only IDs with prior findings;
   at most two findings rounds, then the change blocks for the owner.
5. Implementation step: only allowed files. `done` applies, archives and
   commits.
6. Update PROGRESS only if state, next work or standing rules changed.

## Verification and regeneration

Permanent evals use public APIs, temporary repositories, explicit Git
identities, local hooks and bounded fake agents, with no bootstrap, agent login
or network. Fake-agent runs prove orchestration, not regeneration.

Drills need owner authorization. Reconstruction gets retained contracts, evals
and permitted dependency interfaces only. Keep reports and refs, clean scratch
worktrees, and get owner agreement before any gap-driven change.

The single controller/NEW smoke compares exit codes, JSON envelope shape and
resulting commits/refs. Record intentional differences in DECISIONS.
