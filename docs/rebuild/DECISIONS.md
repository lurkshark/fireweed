# Author decisions

## D001 — Initial graph and interfaces

Adopt PLAN.md eight-module DAG. api.go is the public facade, evals/ permanent black-box tests, internal/ disposable implementation. Share public values through contracts to avoid parent/internal import cycles. The initial app facade forwards to a compiling disposable placeholder; no product behavior is implemented in setup. Use the YAML v3 parser as the sole initial dependency; its checksum file is dependency metadata, not copied implementation.

## D002 — Local authority and controller

Work in the existing clean main checkout. The goal authorizes local changes and commits, including manual start without claiming a human merge. Product approval tests will use an independent base ref. Do not change product enforcement. Keep the bootstrap source clean and all generated caches/build artifacts outside its checkout.
