package evals_test

import (
	"crypto/sha256"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/assets"
	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow"
)

// The shared fixture supplies a Modified item by default. Regeneration is a
// maintenance change, so author its complete approved change before committing.
func regenerateFixture(t *testing.T, configExtra string) string {
	t.Helper()
	root := fixture(t, configExtra, "", "", "", "")
	write(t, root, "changes/0001-manual.md", `---
module: alpha
bases: {}
regenerate: true
---

## Motivation
Rebuild alpha's implementation from retained contracts and evals.

## Interface

## Added

## Modified

## Removed

## Decisions

`)
	return root
}

func archivedImpl(t *testing.T, root, name, module, date string) {
	t.Helper()
	playbook, err := assets.Playbook("impl")
	if err != nil {
		t.Fatal(err)
	}
	hash := sha256.Sum256([]byte(playbook))
	write(t, root, "changes/archive/"+name+".md", fmt.Sprintf(`---
module: %s
bases: {}
run: fixture-%s
phase: applied
---

## Motivation
Earlier implementation change.

## Interface

## Added

## Modified

## Removed

## Decisions

## Provenance
- impl %s agent=manual fireweed=%s playbook=%x tests=1 failing=0 passing=1
`, module, name, date, assets.Version(), hash))
}

func Test_FLOW_080_FLOW_081_FLOW_083_RegenerateDeletesProjectedDisposableAndRebuildsThroughArchive(t *testing.T) {
	root := regenerateFixture(t, "protected: [packages/alpha/internal/global.go]\n")
	change := readFile(t, root, "changes/0001-manual.md")
	write(t, root, "changes/0001-manual.md", strings.Replace(change, "regenerate: true\n", "regenerate: true\nset:\n  preserve: [internal/keep.go]\n", 1))
	spec := readFile(t, root, "packages/alpha/SPEC.md")
	spec = strings.Replace(spec, "depends_on: []\n", "depends_on: []\ninterface: [api.go, internal/contract.go]\nappend_only: [internal/log.go]\n", 1)
	write(t, root, "packages/alpha/SPEC.md", spec)
	files := map[string]string{
		"packages/alpha/api.go":               "package alpha\n",
		"packages/alpha/evals/alpha_test.go":  "package alpha_test\n// retained eval\n",
		"packages/alpha/internal/contract.go": "package alpha\n// retained interface\n",
		"packages/alpha/internal/log.go":      "retained append-only\n",
		"packages/alpha/internal/keep.go":     "retained preserved\n",
		"packages/alpha/internal/global.go":   "retained global protected\n",
		"packages/alpha/manifest.txt":         "retained module manifest\n",
		"packages/beta/internal/other.go":     "another module\n",
	}
	for name, body := range files {
		write(t, root, name, body)
	}
	write(t, root, "packages/alpha/internal/ready", "old implementation\n")
	write(t, root, "packages/alpha/internal/old.go", "old disposable\n")
	for _, d := range project.Check(root) {
		if !d.Warning {
			t.Fatalf("invalid regenerate fixture: %#v", d)
		}
	}
	commit(t, root, "approved regenerate fixture")
	before := head(t, root)
	s := start(t, root)
	if s.Phase != "impl" || s.Run == "" || s.ClaimOID == "" || s.PhaseBase != before || head(t, root) != before {
		t.Errorf("regenerate Start phase/claim/base = %q/%q/%q; HEAD before %s after %s", s.Phase, s.ClaimOID, s.PhaseBase, before, head(t, root))
	}
	for _, name := range []string{"packages/alpha/internal/ready", "packages/alpha/internal/old.go"} {
		if _, err := os.Stat(filepath.Join(root, name)); !os.IsNotExist(err) {
			t.Errorf("disposable %s remained: %v", name, err)
		}
	}
	for name, want := range files {
		if got := readFile(t, root, name); got != want {
			t.Errorf("retained %s changed: %q", name, got)
		}
	}
	current, err := workflow.Current(root, "manual")
	if err != nil || current.Phase != "impl" || current.Run != s.Run {
		t.Fatalf("held regenerate state = %#v, %v", current, err)
	}
	regenerate, err := assets.Playbook("regenerate")
	if err != nil {
		t.Fatal(err)
	}
	impl, err := assets.Playbook("impl")
	if err != nil {
		t.Fatal(err)
	}
	context, err := workflow.Context(root, "manual", false)
	if err != nil || !strings.Contains(context, regenerate) || strings.Contains(context, impl) {
		t.Errorf("regenerate Context did not select playbook: %v, %q", err, context)
	}
	if s.Phase != "impl" {
		return // Archive reconstruction requires the newly specified starting phase.
	}
	write(t, root, "packages/alpha/internal/ready", "reconstructed implementation\n")
	checked, err := verify(root, false)
	if err != nil || len(checked.Reports["alpha"].Cases) != 2 || len(checked.Reports["beta"].Cases) != 1 {
		t.Fatalf("checked reconstruction = %#v, %v", checked.Reports, err)
	}
	applied, err := done(root)
	if err != nil || applied.State.Phase != "applied" {
		t.Fatalf("regenerate Done = %#v, %v", applied, err)
	}
	if head(t, root) == before {
		t.Error("archive did not commit checked reconstruction")
	}
	archive := readFile(t, root, "changes/archive/0001-manual.md")
	if !strings.Contains(archive, "phase: applied") || !strings.Contains(archive, "- impl ") {
		t.Errorf("regenerate archive lacks implementation provenance: %q", archive)
	}
	for name, want := range files {
		if got := readFile(t, root, name); got != want {
			t.Errorf("archived reconstruction changed retained %s: %q", name, got)
		}
	}
}

func Test_FLOW_082_RegenerateWarningUsesLatestTargetImplProvenanceAndReportDate(t *testing.T) {
	for _, tc := range []struct {
		name, date, result string
		warn               bool
	}{
		{name: "missing", warn: true},
		{name: "older passing", date: "2026-06-14T12:00:00Z", result: "pass", warn: true},
		{name: "newer failed", date: "2026-06-16T12:00:00Z", result: "fail", warn: true},
		{name: "newer passing", date: "2026-06-16T12:00:00Z", result: "pass", warn: false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := regenerateFixture(t, "")
			archivedImpl(t, root, "0000-alpha-earlier", "alpha", "2026-06-01T12:00:00Z")
			archivedImpl(t, root, "0001-alpha-latest", "alpha", "2026-06-15T12:00:00Z")
			archivedImpl(t, root, "0002-beta-later", "beta", "2026-06-20T12:00:00Z")
			if tc.date != "" {
				// An ordinary YAML timestamp is a valid authored report date.
				write(t, root, ".fireweed/drills/alpha.md", fmt.Sprintf("---\ndate: %s\nbase_commit: %s\nagent: fixture\nmodel: fixture\nresult: %s\n---\n\n## Own evals\n\n## Dependent evals\n\n## Contradicted Decisions\n\n## Gaps\n", tc.date, head(t, root), tc.result))
			}
			for _, d := range project.Check(root) {
				if !d.Warning {
					t.Fatalf("invalid date/report fixture: %#v", d)
				}
			}
			commit(t, root, "approved archives and authored report")
			s := start(t, root)
			if got := containsCode(s.Diagnostics, "P10"); got != tc.warn {
				t.Errorf("P10 warning = %t; want %t; diagnostics = %#v", got, tc.warn, s.Diagnostics)
			}
			if tc.warn {
				for _, d := range s.Diagnostics {
					if d.Code == "P10" && !d.Warning {
						t.Errorf("P10 diagnostic is not a warning: %#v", d)
					}
				}
			}
		})
	}
}

func Test_FLOW_083_RecordedRegenerateImplContextUsesRegeneratePlaybook(t *testing.T) {
	root := regenerateFixture(t, "")
	change := readFile(t, root, "changes/0001-manual.md")
	write(t, root, "changes/0001-manual.md", strings.Replace(change, "regenerate: true\n", "regenerate: true\nphase: impl\nrun: fixture-recorded\n", 1))
	write(t, root, "packages/alpha/internal/ready", "recorded phase implementation\n")
	for _, d := range project.Check(root) {
		if !d.Warning {
			t.Fatalf("invalid recorded-phase fixture: %#v", d)
		}
	}
	commit(t, root, "approved recorded regenerate phase")
	s := start(t, root)
	if s.Phase != "impl" || s.Run != "fixture-recorded" {
		t.Fatalf("recorded regenerate state = %#v", s)
	}
	if got := readFile(t, root, "packages/alpha/internal/ready"); got != "recorded phase implementation\n" {
		t.Errorf("recorded phase repeated deletion: %q", got)
	}
	context, err := workflow.Context(root, "manual", false)
	if err != nil {
		t.Fatal(err)
	}
	playbook, err := assets.Playbook("regenerate")
	if err != nil {
		t.Fatal(err)
	}
	ordinary, err := assets.Playbook("impl")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(context, playbook) || strings.Contains(context, ordinary) {
		t.Errorf("recorded regenerate Context selected wrong playbook: %q", context)
	}
}
