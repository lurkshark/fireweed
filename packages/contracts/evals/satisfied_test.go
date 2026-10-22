package evals

import (
	"reflect"
	"strings"
	"testing"

	c "github.com/lurkshark/fireweed/packages/contracts"
)

func Test_CON016AlreadySatisfiedLocalHandles(t *testing.T) {
	text := changeDoc("already_satisfied: [new:second, AUTH-003, new:first]\nset: {hooks: {setup: 'explain new:first'}}\n", "## Added\n\n### new:first\nIt MUST retain the first behavior.\n\n### new:second\nIt MUST retain the second behavior.\n\n## Interface\n\nExpose new:second and new:first.\n")
	want := []string{"new:second", "AUTH-003", "new:first"}
	first, err := c.ParseChange(text, "changes/0001-satisfied.md", "AUTH")
	if err != nil {
		t.Fatalf("local newly added declarations rejected: %v", err)
	}
	if !reflect.DeepEqual(first.AlreadySatisfied, want) || first.Kind != "spec" || len(first.Added) != 2 || first.Added[0].ID != "new:first" || first.Added[1].ID != "new:second" {
		t.Fatalf("lost local declaration values %+v", first)
	}
	second, err := c.ParseChange(text, "changes/0001-satisfied.md", "AUTH")
	if err != nil {
		t.Fatal(err)
	}
	first.AlreadySatisfied[0] = "poison"
	first.Added[0].Body = "poison"
	first.Document.Metadata["already_satisfied"].([]any)[0] = "poison"
	first.Document.Metadata["set"].(map[string]any)["hooks"].(map[string]any)["setup"] = "poison"
	first.Document.Sections[0].Body = "poison"
	third, err := c.ParseChange(text, "changes/0001-satisfied.md", "AUTH")
	if err != nil || !reflect.DeepEqual(second, third) {
		t.Fatalf("parses share mutable results: %v", err)
	}
	if !strings.Contains(text, "already_satisfied: [new:second, AUTH-003, new:first]") {
		t.Fatal("source text changed")
	}
}
func Test_CON016SatisfiedHandleScopeAndTypes(t *testing.T) {
	declared := "## Added\n\n### new:local\nIt MUST remain local.\n"
	valid := changeDoc("already_satisfied: [new:local]\n", declared)
	if _, err := c.ParseChange(valid, "changes/0001-declaration.md", "AUTH"); err != nil {
		t.Fatalf("valid local premise: %v", err)
	}
	// A prior successful parse must never supply a namespace to another change.
	foreign := changeDoc("already_satisfied: [new:local]\n", "")
	if _, err := c.ParseChange(foreign, "changes/0002-foreign.md", "AUTH"); err == nil {
		t.Fatal("another change's declaration leaked")
	}
	for _, meta := range []string{
		"already_satisfied: [new:unknown]\n", "already_satisfied: [new:UPPER]\n", "already_satisfied: ['new:']\n", "already_satisfied: [new:local/path]\n", "already_satisfied: [invalid]\n",
		"already_satisfied: new:local\n", "already_satisfied: true\n", "already_satisfied: [23]\n", "already_satisfied: {new:local: true}\n",
		"fixes: [new:local]\n", "bases: {'new:local': abcdef12}\n",
	} {
		if _, err := c.ParseChange(changeDoc(meta, declared), "changes/0003-invalid.md", "AUTH"); err == nil {
			t.Fatalf("invalid metadata accepted %q", meta)
		}
	}
	// Existing real IDs remain legal with and without local declarations.
	for _, body := range []string{"", declared} {
		got, err := c.ParseChange(changeDoc("already_satisfied: [AUTH-003, AUTH-001]\nfixes: [AUTH-002]\nbases: {AUTH-001: abcdef12}\n", body), "changes/0004-real.md", "AUTH")
		if err != nil || !reflect.DeepEqual(got.AlreadySatisfied, []string{"AUTH-003", "AUTH-001"}) || !reflect.DeepEqual(got.Fixes, []string{"AUTH-002"}) || got.Bases["AUTH-001"] != "abcdef12" {
			t.Fatalf("real-ID compatibility %+v %v", got, err)
		}
	}
}
