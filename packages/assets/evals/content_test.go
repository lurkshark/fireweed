package evals

import (
	"crypto/sha256"
	"fmt"
	a "github.com/lurkshark/fireweed/packages/assets"
	"os"
	"reflect"
	"sort"
	"strings"
	"testing"
)

func isolatedCWD(t *testing.T) {
	t.Helper()
	old, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(t.TempDir()); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := os.Chdir(old); err != nil {
			t.Error(err)
		}
	})
}
func Test_AST001BuildMetadata(t *testing.T) {
	isolatedCWD(t)
	got := a.Info()
	if got.Version != "0.4.0" || got.CurrentFormat != 1 || got.PreviousFormat != 0 {
		t.Errorf("build info %+v", got)
	}
}
func Test_AST002EmbeddedPlaybooks(t *testing.T) {
	isolatedCWD(t)
	want := []string{"audit", "author-change", "evals", "impl", "regenerate"}
	names := a.Playbooks()
	if !reflect.DeepEqual(names, want) {
		t.Fatalf("playbooks %v, want %v", names, want)
	}
	names[0] = "mutated"
	if !reflect.DeepEqual(a.Playbooks(), want) {
		t.Error("inventory mutation leaked")
	}
	clauses := map[string][]string{
		"author-change": {"one module", "new:<slug>", "MUST", "Modified", "headings only", "fireweed stamp", "base branch", "approval"},
		"evals":         {"fireweed context", "black-box", "fail", "testcase", "compile", "new disposable", "interface", "fireweed verify", "fireweed done"},
		"audit":         {"independent", "exact", "spec ID", "No findings", "checked IDs", "untestable", "Questions", "fireweed done"},
		"impl":          {"MUST NOT edit", "evals", "interface", "SPEC.md", "protected", "append-only", "dependents", "boundaries", "fireweed verify", "fireweed done"},
		"regenerate":    {"retained", "interfaces", "evals", "MUST NOT consult", "old implementation", "history", "bootstrap", "dependent", "Decisions", "gap", "spec ID", "module"},
	}
	for _, name := range want {
		asset, err := a.Playbook(name)
		if err != nil {
			t.Fatal(err)
		}
		if asset.Name != name || strings.TrimSpace(asset.Text) == "" {
			t.Errorf("asset %+v", asset)
		}
		expected := fmt.Sprintf("%x", sha256.Sum256([]byte(asset.Text)))
		if asset.Hash != expected || len(asset.Hash) != 64 {
			t.Errorf("hash %q, want %q", asset.Hash, expected)
		}
		for _, clause := range clauses[name] {
			if !strings.Contains(asset.Text, clause) {
				t.Errorf("playbook %s missing instruction %q", name, clause)
			}
		}
		again, err := a.Playbook(name)
		if err != nil || again != asset {
			t.Errorf("unstable playbook %s: %+v %v", name, again, err)
		}
	}
	for _, name := range []string{"", "missing", "../author-change", "AUDIT"} {
		if _, err := a.Playbook(name); err == nil {
			t.Errorf("unknown playbook accepted %q", name)
		}
	}
}
func Test_AST003DiagnosticCatalog(t *testing.T) {
	expected := []string{"U01", "P1-01", "P1-02", "P1-03", "P2-01"}
	for _, group := range []struct {
		prefix string
		count  int
	}{{"C", 16}, {"P", 6}, {"D", 7}, {"H", 4}} {
		for i := 1; i <= group.count; i++ {
			expected = append(expected, fmt.Sprintf("%s%02d", group.prefix, i))
		}
	}
	sort.Strings(expected)
	got := a.Codes()
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("codes %v, want %v", got, expected)
	}
	got[0] = "mutated"
	if !reflect.DeepEqual(a.Codes(), expected) {
		t.Error("catalog mutation leaked")
	}
	for _, code := range expected {
		explanation, err := a.Explain(code)
		if err != nil || len(strings.Fields(explanation)) < 8 {
			t.Errorf("unhelpful %s: %q %v", code, explanation, err)
		}
	}
	meanings := map[string][]string{
		"H01": {"worker", "protected"}, "H02": {"implementation", "evals", "interfaces"},
		"H03": {"append-only", "anyone", "merge", "new file"},
		"H04": {"archive", "immutable", "deletion", "revert", "abort", "fireweed commits", "incoming", "merge again"},
	}
	for code, terms := range meanings {
		text, err := a.Explain(code)
		if err != nil {
			t.Fatal(err)
		}
		for _, term := range terms {
			if !strings.Contains(strings.ToLower(text), term) {
				t.Errorf("explain %s missing %q: %s", code, term, text)
			}
		}
	}
	for _, code := range []string{"", "unknown", "C00", "H05", "c01"} {
		if _, err := a.Explain(code); err == nil {
			t.Errorf("unknown code accepted %q", code)
		}
	}
}
