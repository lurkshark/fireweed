# Go recipe

Create a Go module with `go mod init <module-path>` if the project has no `go.mod` yet. For a repository with several Go modules, create a workspace with `go work init` and add each module with `go work use <directory>`. Adapt the hook command and `go mod download` setup command to the chosen module or workspace layout; setup runs after a Go module exists.

The dependency graph is derived from `FIREWEED_BIN modules --json`: build each module's `depends_on` edges and `zones` from the returned graph, including its `dir`. Set the `interface` zone to the Go packages other modules' evals may import, and put tests in the `evals` zone. The boundaries hook combines that graph with the workspace `go.mod` module path. Keep `api.go` as the thin retained facade and put implementation under `internal/`.

For missing behavior, keep the retained `api.go` function as a delegation point and put the temporary stub in a new disposable `internal/` file. For example, with module path `example.com/myproject`:

`api.go` (retained interface):

```go
package myproject

import "example.com/myproject/internal/errorcode"

func ErrorCode(err error) string { return errorcode.From(err) }
```

`internal/errorcode/errorcode.go` (new disposable stub):

```go
package errorcode

func From(err error) string { return "" } // Temporary until implementation.
```

During implementation, replace the disposable stub with the required behavior while leaving the retained `api.go` facade unchanged.
