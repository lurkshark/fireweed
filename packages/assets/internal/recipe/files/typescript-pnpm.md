# TypeScript/pnpm recipe

Prerequisites: Node.js 24 or newer and pnpm 10.15.1. Create a root package.json with `private: true`, `type: module`, and `packageManager: pnpm@10.15.1`; create pnpm-workspace.yaml with `packages: ['packages/*']`. Each module has its own package.json and exports its retained `./src/index.ts` interface. Run `pnpm install --offline` when the pnpm version is cached; otherwise install dependencies through your normal package registry before working offline. Node's built-in TypeScript stripping and test runner APIs run the eval files, so no compiler package is required for these templates.

Derive the dependency graph from `FIREWEED_BIN modules --json`: use every module `dir`, `depends_on`, and resolved `zones` for boundary checks. Do not maintain a second dependency graph in package metadata. Keep implementation in `src/internal/`, and expose only retained entry points through package exports. The hook checks source imports against Fireweed's graph.

Run the setup hook after creating the root workspace files. Keep test names linked to spec IDs, for example `ALC001Pass` for ALC-001.
