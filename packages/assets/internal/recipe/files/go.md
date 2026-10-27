# Go recipe

Create a Go module with `go mod init <module-path>` if the project has no `go.mod` yet. For a repository with several Go modules, create a workspace with `go work init` and add each module with `go work use <directory>`. Adapt the hook command and `go mod download` setup command to the chosen module or workspace layout; setup runs after a Go module exists.

The dependency graph is derived from `FIREWEED_BIN modules --json`: build each module's `depends_on` edges and `zones` from the returned graph, including its `dir`. Set the `interface` zone to the Go packages other modules' evals may import, and put tests in the `evals` zone. The boundaries hook combines that graph with the workspace `go.mod` module path. Keep `api.go` as the thin retained facade and put implementation under `internal/`.
