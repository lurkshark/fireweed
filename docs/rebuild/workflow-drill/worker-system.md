# Workflow reconstruction

Reconstruct workflow from the supplied retained workflow SPEC (including Decisions), interface/model and public evals. Only the declared contracts, gitstore, assets, project and checks public api.go/model interfaces are provided as dependency inputs.

Write only packages/workflow/internal/** and DRILL_NOTES.md. All supplied inputs are immutable. No go.mod, replacement dependency, hook, manifest, tests or root setup may be added. Do not run local builds, tests, or module setup, including temporary module setup. Missing compilation infrastructure is intentional. Format only the fresh code you write, then return it to the separate outside verifier; that verifier runs full builds, own evals and all transitive dependent evals.

Support ordinary full SHA-1 clones on macOS and Linux with the host's installed Git. The selected macOS verifier uses Apple Git 2.30.1 (Apple Git-130). Choose Git operations compatible with that verifier.

Read only this supplied input tree and the fresh disposable code you write here. Its Git history is input-only. Do not read original/deleted implementation, old history, refs/attempts/proofs/reports/reconstructions, bootstrap or rebuild-v1, outside repositories, verifier paths, binaries, or source caches. Do not search outside this tree, use controller queries as an oracle, run other agents, or contact the network. Do not infer any permission from a visible outside path or environment variable.

Do not commit, change Git refs/history/config, invoke lifecycle mutators, modify retained inputs, or weaken any verification requirement. Return to the outside controller after reconstruction; it owns all transitions and commits. Use the configured workspace-write command and model/effort only.

The current tree's .fireweed/bin is first on PATH and GOCACHE=/tmp/fireweed-rebuild-cache. Do not invoke any existing Fireweed binary to obtain behavior or implementation. No controller/model fallback is permitted. Record contradicted Decisions and targeted gaps in DRILL_NOTES.md, naming the nearest spec ID (or no spec) and owning module. A gap authorizes no correction of retained inputs.
