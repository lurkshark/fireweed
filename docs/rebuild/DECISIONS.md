# Author decisions

## D001 — Initial graph and interfaces

Adopt PLAN.md eight-module DAG. api.go is the public facade, evals/ permanent black-box tests, internal/ disposable implementation. Share public values through contracts to avoid parent/internal import cycles. The initial app facade forwards to a compiling disposable placeholder; no product behavior is implemented in setup. Use the YAML v3 parser as the sole initial dependency; its checksum file is dependency metadata, not copied implementation.

## D002 — Local authority and controller

Work in the existing clean main checkout. The goal authorizes local changes and commits, including manual start without claiming a human merge. Product approval tests will use an independent base ref. Do not change product enforcement. Keep the bootstrap source clean and all generated caches/build artifacts outside its checkout.

## D003 — Working real-agent CLI

Use explicit gpt-6-sol only in this repository worker/audit argv after the configured default returned unsupported-model errors and a cached supported-model probe succeeded. Keep CLI sessions ephemeral. Do not alter global user configuration. Fresh independent auditing contexts are used throughout; identical model choice is disclosed rather than claimed to be a different-model review.

## D004 — Retained data types without import cycles

The upcoming contracts configuration change will add a retained `model/*.go` interface subpackage with public aliases through api.go. Disposable contracts implementation can import the retained types without importing its public parent; dependent modules still import only the public contracts facade. Update interface zones through that authored change, not a protected setup shortcut.
