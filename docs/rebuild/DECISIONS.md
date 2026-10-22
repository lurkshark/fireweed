# Author decisions

## D001 — Initial graph and interfaces

Adopt PLAN.md eight-module DAG. api.go is the public facade, evals/ permanent black-box tests, internal/ disposable implementation. Share public values through contracts to avoid parent/internal import cycles. The initial app facade forwards to a compiling disposable placeholder; no product behavior is implemented in setup. Use the YAML v3 parser as the sole initial dependency; its checksum file is dependency metadata, not copied implementation.

## D002 — Local authority and controller

Work in the existing clean main checkout. The goal authorizes local changes and commits, including manual start without claiming a human merge. Product approval tests will use an independent base ref. Do not change product enforcement. Keep the bootstrap source clean and all generated caches/build artifacts outside its checkout.

## D003 — Working real-agent CLI

Use explicit gpt-6-sol only in this repository worker/audit argv after the configured default returned unsupported-model errors and a cached supported-model probe succeeded. Keep CLI sessions ephemeral. Do not alter global user configuration. Fresh independent auditing contexts are used throughout; identical model choice is disclosed rather than claimed to be a different-model review.

## D004 — Retained data types without import cycles

The upcoming contracts configuration change will add a retained `model/*.go` interface subpackage with public aliases through api.go. Disposable contracts implementation can import the retained types without importing its public parent; dependent modules still import only the public contracts facade. Update interface zones through that authored change, not a protected setup shortcut.

## D005 — Vet-only facade repair between claims

Moving Spec to the retained model package made an existing unkeyed facade conversion fail vet. With no active claim, an author commit keys that exact literal; it changes neither the public contract nor executable behavior and does not edit evals, specs or implementation. This is retained entry/facade glue repair, not an unverified behavioral change. The prior evidence claim that vet passed is corrected; rerun vet before recording success.

## D006 — Assets runtime-isolation evidence

Native sandbox policies must permit the OS loader to start. The asset eval therefore combines source/network denial with a test-only Darwin file/socket interposer armed after startup; positive controls prove ignored errors cannot hide attempts. Its cc/dylib dependency belongs only to a temporary eval fixture, never the product. Other-platform isolation is explicitly unverified; portable content checks still run. Keep this platform/test prerequisite visible in final delivery evidence and README testing instructions.

## D007 — Embedded CLI access before init

Design playbooks/explanations are embedded binary protocol data. The first app slice makes help/version/explain and embedded playbook lookup usable without a project. The reference cli.go resolves explain before Load, but playbook after Load and therefore fails before init. Allowing embedded playbooks before init supports the design's setup/authoring flow. Later project integration must honor project-owned overrides and binary pins when a project is present; this first slice does not claim that integration. The reference error is an intentional comparison difference, not silently reproduced.

## D008 — correct discovery eval fixtures through a full revert

The first discovery-validation eval/audit run (9191c67 through d9fe6ed) passed assertion review, but implementation exposed three invalid fixture premises. CON-013 rejects unknown module metadata; the retention fixture's `custom` key was invalid. The retained document-grammar Decision requires filenames with at least four digits, so `2-alpha.md` never reaches numeric ordering. CON-005's Decision treats adjacent stars within a segment as ordinary wildcards, so `bad/**suffix` is valid. These are eval mistakes, not product behavior to reproduce or reasons to weaken contracts.

Preserved the implementation and eval drafts locally, released the claim, and fully reverted all eight phase commits in reverse order (b3de999 through 0f4ffb1), restoring the queued authored change and unimplemented project facade. No frozen eval/interface was edited during implementation. Correct fixtures use known metadata, valid variable-width numbers (`00010` versus `0002`) and genuinely invalid repeated separators. Restart the complete fail-first/audit/impl cycle. The earlier commits/audits remain visible as superseded evidence; they do not prove this change complete.
