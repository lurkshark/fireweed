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
and exposes it as .fireweed/bin/fireweed. Launch workers with that directory on
PATH. Product fixtures must set FIREWEED_BIN to the NEW product executable.

Go and Git are required. The controller and setup remain protected. Run doctor
after a coherent setup commit; no package downloads or bootstrap source reads
are part of setup.
