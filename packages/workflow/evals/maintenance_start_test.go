package evals_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow"
)

func Test_FLOW_077_MaintenanceStartsInImplAndArchivesDirectly(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	write(t, root, "changes/0001-manual.md", `---
module: alpha
bases: {}
---

## Motivation
Simplify the existing alpha implementation while keeping its behavior.

## Interface

## Added

## Modified

## Removed

## Decisions

`)
	write(t, root, "packages/alpha/internal/ready", "original implementation\n")
	write(t, root, "packages/alpha/evals/alpha_test.go", "package alpha_test\n\n// Retained coverage fixture.\n")
	marker := hookMarker(t, root)
	for _, d := range project.Check(root) {
		if !d.Warning {
			t.Fatalf("invalid maintenance fixture: %#v", d)
		}
	}
	retained := map[string]string{}
	for _, name := range []string{"packages/alpha/SPEC.md", "packages/alpha/api.go", "packages/alpha/evals/alpha_test.go"} {
		b, err := os.ReadFile(filepath.Join(root, name))
		if err != nil {
			t.Fatal(err)
		}
		retained[name] = string(b)
	}
	before := head(t, root)
	s := start(t, root)
	if s.Phase != "impl" || s.Run == "" || s.ClaimOID == "" || s.PhaseBase != before || head(t, root) != before {
		t.Fatalf("maintenance Start state = %#v; HEAD before %s after %s", s, before, head(t, root))
	}
	if _, err := os.Stat(marker); !os.IsNotExist(err) {
		t.Errorf("Start ran an eval hook: %v", err)
	}
	current, err := workflow.Current(root, "manual")
	if err != nil || current.Phase != "impl" || current.Run != s.Run || current.ClaimOID != s.ClaimOID {
		t.Fatalf("persisted maintenance state = %#v, %v", current, err)
	}
	pending, err := os.ReadFile(filepath.Join(root, "changes/0001-manual.md"))
	if err != nil || !strings.Contains(string(pending), "phase: impl") {
		t.Fatalf("pending maintenance phase = %q, %v", pending, err)
	}
	write(t, root, "packages/alpha/internal/ready", "simplified implementation\n")
	checked, err := verify(root, false)
	if err != nil {
		t.Fatalf("full maintenance Verify: %v, reports: %#v", err, checked.Reports)
	}
	if len(checked.Reports["alpha"].Cases) != 2 || len(checked.Reports["beta"].Cases) != 1 {
		t.Errorf("full Verify did not run current and direct dependent coverage: %#v", checked.Reports)
	}
	if head(t, root) != before {
		t.Fatal("Verify changed HEAD")
	}
	applied, err := done(root)
	if err != nil || applied.State.Phase != "applied" {
		t.Fatalf("maintenance Done = %#v, %v", applied, err)
	}
	after := head(t, root)
	if after == before || git(t, root, "rev-parse", "HEAD^") != before {
		t.Errorf("maintenance archive was not one commit: before %s after %s", before, after)
	}
	if _, err := os.Stat(filepath.Join(root, "changes/0001-manual.md")); !os.IsNotExist(err) {
		t.Errorf("pending maintenance change remains: %v", err)
	}
	archive, err := os.ReadFile(filepath.Join(root, "changes/archive/0001-manual.md"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(archive), "phase: applied") || !strings.Contains(string(archive), "- impl ") || strings.Contains(string(archive), "- evals ") || strings.Contains(string(archive), "- audit ") {
		t.Errorf("maintenance archive has unexpected phase history: %s", archive)
	}
	for name, want := range retained {
		got, err := os.ReadFile(filepath.Join(root, name))
		if err != nil || string(got) != want {
			t.Errorf("retained %s changed: %q, %v", name, got, err)
		}
	}
	if _, err := workflow.Current(root, "manual"); err == nil {
		t.Error("applied maintenance claim remained held")
	}
}
