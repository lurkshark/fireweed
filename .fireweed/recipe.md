# Project-owned Go recipe

Each module exposes a thin api.go facade, keeps black-box Go tests in evals/,
and places disposable behavior under internal/. The protected Go hook runs
`go test -json ./<module>/...`, writes JUnit, and rejects unfinished tests and
package failures. Its boundary check reads the invoking FIREWEED_BIN's modules
JSON; module SPEC.md files are the only graph source. It enforces declared
cross-module dependencies and eval imports through public interfaces.

Canaries use the controller's version-1 canaries.yaml contract, with Go module
and random module-name substitutions. Templates use {{name}}, {{package}} and
{{prefix}}. The setup hook copies the invoking controller into scratch trees
and exposes it as a regular, independent .fireweed/bin/fireweed copy. Launch
workers with the current tree's .fireweed/bin first on PATH. Main uses the
verified NEW pin recorded in docs/rebuild/PROGRESS.md; mutable bin/fireweed is
the subject, and fireweed-bootstrap is retained only for explicit recovery.
FIREWEED_BIN always names the invoking controller, including in scratch hooks.
Product fixtures explicitly select the NEW product executable. Never fall back
automatically or switch controllers while a run or claim is active.

Go and Git are required. The controller and setup remain protected. Run doctor
after a coherent setup commit; no package downloads or bootstrap source reads
are part of setup.
