---
audit_rounds: 0
bases: {}
module: app
phase: applied
run: 20261030T013002-9857da7756618210344049b4
---



## Motivation

Design CLI Authoring requires new-module <name> --prefix <P> to create a module
from the project's templates. NEW still rejects it. Reuse the copied recipe,
public project/contracts APIs and existing app creation/output helpers; no new
public subsystem is needed. Close this slice with actual NEW init --no-app,
new-module, modules/check and a second duplicate request in a disposable Go
fixture, inspecting rendered project-owned templates and non-overwrite behavior.
Checkpoint0106–0111 passed and pushed; this is the next application after count0.

## Interface

No public API/model changes. Implement app-owned authoring composition under
app/internal, using project.Load/ValidateGraph and contracts as needed. Reuse
existing initialization template/destination helpers where appropriate.

## Added

### APP-100
new-module <name> --prefix <P> MUST discover the nearest project root and create the named module from .fireweed/templates/module in a new directory matched by the configured module globs.
It MUST use that project-owned copy rather than replace it with embedded templates.

### APP-101
Module creation MUST render template contents using name, prefix, package, module_dir and go_module substitutions, removing a terminal .tmpl suffix from output filenames.
The resulting SPEC.md MUST declare the requested module name and prefix and be loadable as part of the project graph.

### APP-102
new-module MUST reject an existing module name, existing ID prefix or occupied destination before writing module files, preserving the existing module and destination contents.

### APP-103
new-module MUST return exit2 with corrective U01 for invalid arguments and options, the existing cause code for project errors, C03 for missing or malformed templates, and P01 for conflicting module identity or destination.
It MUST reject requests before writing module files when these errors are present.

### APP-104
Successful new-module MUST return exit0 and print the repository-relative module directory.
With --json it MUST emit schema1 string data equal to that directory and empty diagnostics; failures MUST emit null data and code/message diagnostics with empty stderr.

## Modified



## Removed



## Decisions

Creation uses the first configured module glob ending in a single wildcard
component under a literal parent, placing the requested name in that component;
a configuration with no such destination is reported as C01 before writing.
The package substitution replaces name hyphens with underscores; go_module is
read from go.mod when present and is otherwise empty.

## Notes

Use existing app public fixtures/temporary identities; no API changes or new
backend abstraction. Cover Go/project-owned modified templates and a custom
module parent in representative fixtures, including nearest versus enclosing
root, substitutions/output suffix/plain/JSON, discovery and duplicate/missing/
malformed/usage/config refusal. Avoid a layout/language/option Cartesian matrix,
instrumentation, new recipe features or unsupported environment cases. Preserve
all existing assertions. Workers never commit, alter phase/claims or edit
protected files/other modules/SPEC/interface. Evals freeze, separate isolated
Sol/high audit, then implementation/full native verification and actual NEW
proof. Own under5m/full under15m with unchanged bounds. No bootstrap source/tests,
rebuild-v1, deleted/abandoned/saved implementation or old worker-log oracles.
This is the second consecutive app application after0111; before a third app
change require the normal checkpoint unless separately authorized.

## Provenance

- evals 2026-10-30T01:38:38.959698Z agent=manual fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=236 failing=12 passing=224
- audit 2026-10-30T01:41:15.440201Z agent=manual fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0
- impl 2026-10-30T01:48:25.830755Z agent=manual fireweed=0.1.0 playbook=07554a3249d8fca0ba70df7504135f297491465caa63458734a4b8351734d0e5 tests=236 failing=0 passing=236

## Audit

No findings
Checked IDs: APP-100, APP-101, APP-102, APP-103, APP-104
