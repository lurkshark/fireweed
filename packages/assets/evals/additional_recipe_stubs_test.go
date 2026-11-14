package evals_test

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/contracts"
)

func Test_AST_065_AdditionalRecipeStubHandoff(t *testing.T) {
	const assertion = "greet should return a personalized greeting"
	// The Python hook runs again immediately after the stub file is replaced.
	t.Setenv("PYTHONDONTWRITEBYTECODE", "1")
	for _, recipe := range []struct {
		name, language, evalPath, eval, replacement string
	}{
		{"typescript-pnpm", "ts|typescript", "evals/greet.test.ts", `
import { strictEqual } from 'node:assert';
import { test } from 'node:test';
import { greet } from 'alpha';

test('ALC001Greeting', () => {
  const actual = greet('Ada');
  strictEqual(actual, 'Hello, Ada!', '` + assertion + `');
});
`, "export function greet(name: string): string { return `Hello, ${name}!`; }\n"},
		{"python-uv", "py|python", "evals/test_greet.py", `
import unittest
from alpha import greet

class TestGreeting(unittest.TestCase):
    def test_ALC001Greeting(self):
        actual = greet('Ada')
        self.assertEqual(actual, 'Hello, Ada!', '` + assertion + `')
`, "def greet(name: str) -> str:\n    return f'Hello, {name}!'\n"},
	} {
		r := additionalRecipe(t, recipe.name)
		for _, guide := range []struct{ name, text string }{
			{"instructions", r.Instructions},
			{"bundled guide", r.Files[".fireweed/recipes/"+recipe.name+".md"]},
		} {
			t.Run(recipe.name+"/"+guide.name, func(t *testing.T) {
				w := newExtraWorkspace(t, recipe.name)
				stubPath := installRecipeStubExample(t, w, guide.text, recipe.language)
				w.write(filepath.Join("packages/alpha", recipe.evalPath), recipe.eval)
				if _, out, err := w.hook("boundaries", "alpha"); err != nil {
					t.Fatalf("example violates recipe boundaries: %v %s", err, out)
				}
				report, out, err := w.hook("eval", "alpha")
				greeting := greetingCase(report)
				if err == nil || greeting == nil || greeting.Failure == nil {
					t.Fatalf("temporary stub must reach and fail the greeting assertion: %v %s %+v", err, out, report.Cases)
				}
				if !strings.Contains(*greeting.Failure, assertion) {
					t.Fatalf("expected greeting assertion failure, got %q", *greeting.Failure)
				}

				w.write(stubPath, recipe.replacement)
				report, out, err = w.hook("eval", "alpha")
				greeting = greetingCase(report)
				if err != nil || report.hasFailure() || greeting == nil || greeting.state() != "pass" {
					t.Errorf("replacing only the disposable stub must pass the greeting assertion: %v %s %+v", err, out, report.Cases)
				}
			})
		}
	}
}

func greetingCase(report junitSuite) *junitCase {
	for _, c := range report.Cases {
		if strings.HasSuffix(c.Name, "ALC001Greeting") {
			return &c
		}
	}
	return nil
}

func installRecipeStubExample(t *testing.T, w *extraWorkspace, guide, language string) string {
	t.Helper()
	// Read the file labels and fenced snippets from the guide itself.
	fence := regexp.MustCompile(`(?ms)^\x60([\w./-]+)\x60([^\n]*)\n\s*^\x60{3}(?:` + language + `)[ \t]*\n(.*?)^\x60{3}[ \t]*$`)
	var retained, stub []string
	for _, snippet := range fence.FindAllStringSubmatch(guide, -1) {
		switch contracts.ResolveZone(snippet[1], w.recipe.ModuleDefaults) {
		case "interface":
			retained = snippet
		case "disposable":
			stub = snippet
		}
	}
	if retained == nil || stub == nil {
		t.Fatal("guide must show a retained interface and a new disposable stub in file-labeled code blocks")
	}

	prose := strings.Join(strings.Fields(fence.ReplaceAllString(guide, "$1$2")), " ")
	requirePattern(t, prose, `(?i)\btemporar\w*\b[^.!?]*\bstub\b|\bstub\b[^.!?]*\btemporar\w*\b`, "temporary stub guidance")
	requirePattern(t, prose, `(?i)\b(?:implement(?:ation|ing)|impl step)\b[^!?]*\breplac\w*\b[^!?]*\bstub\b(?:[^!?]*\bwithout\s+(?:edit|chang)\w*\b[^!?]*\b(?:interface|facade)\b|[^!?]*\b(?:interface|facade)\b[^!?]*\b(?:unchanged|untouched|as is)\b)`, "implementation replaces the stub while preserving the retained interface")

	interfacePath := filepath.Join("packages/alpha", retained[1])
	stubPath := filepath.Join("packages/alpha", stub[1])
	if _, err := os.Stat(filepath.Join(w.root, stubPath)); !errors.Is(err, fs.ErrNotExist) {
		t.Fatalf("disposable example must add a new file %s: stat returned %v", stub[1], err)
	}
	baseline, err := os.ReadFile(filepath.Join(w.root, interfacePath))
	if err != nil {
		t.Fatal(err)
	}
	w.write(interfacePath, string(baseline)+"\n"+retained[3])
	w.write(stubPath, stub[3])
	return stubPath
}
