package evals_test

import (
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"
)

func Test_AST_073_GlobalRecipeBoundaries(t *testing.T) {
	t.Run("go", func(t *testing.T) {
		w := newGoRecipeWorkspace(t, goRecipe(t))
		checkGlobalBoundaries(t, w.root, w.boundaries, func() {
			w.write("packages/beta/foreign.go", `package beta
import _ "example.com/recipefixture/packages/alpha"
`)
		})
	})
	for _, name := range additionalRecipeNames {
		t.Run(name, func(t *testing.T) {
			w := newExtraWorkspace(t, name)
			checkGlobalBoundaries(t, w.root, func() (string, error) {
				_, output, err := w.hook("boundaries", "")
				return output, err
			}, func() {
				w.apply(w.canaries.Undeclared, "beta")
			})
		})
	}
}

func checkGlobalBoundaries(t *testing.T, root string, run func() (string, error), addUndeclaredImport func()) {
	t.Helper()
	output, err := run()
	if err != nil || hasBoundaryLine(output, "alpha") || hasBoundaryLine(output, "beta") {
		t.Fatalf("global hook rejected the clean graph: %v\n%s", err, output)
	}
	query, err := os.ReadFile(filepath.Join(root, "modules-query.log"))
	if err != nil || !slices.Contains(strings.Split(string(query), "\n"), "modules --json") {
		t.Fatalf("global hook did not run FIREWEED_BIN modules --json: %v, log %q", err, query)
	}

	addUndeclaredImport()
	output, err = run()
	if !hasBoundaryLine(output, "beta") || hasBoundaryLine(output, "alpha") {
		t.Errorf("global hook must attribute the undeclared import to beta: %v\n%s", err, output)
	}
}

func hasBoundaryLine(output, module string) bool {
	for _, line := range strings.Split(output, "\n") {
		if strings.HasPrefix(line, module+":") {
			return true
		}
	}
	return false
}

func expandRecipeHook(command, module, report string) string {
	// Global hooks receive empty module and directory values, as in checks.Boundaries.
	var moduleDir string
	if module != "" {
		moduleDir = "packages/" + module
	}
	quote := func(value string) string {
		return "'" + strings.ReplaceAll(value, "'", "'\"'\"'") + "'"
	}
	return strings.NewReplacer("{module}", quote(module), "{module_dir}", quote(moduleDir), "{report}", quote(report)).Replace(command)
}
