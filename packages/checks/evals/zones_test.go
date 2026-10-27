package evals_test

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/lurkshark/fireweed/packages/checks"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project"
)

func zoneFixture(t *testing.T) (project.Snapshot, string) {
	t.Helper()
	root := t.TempDir()
	for path := range map[string]string{
		"packages/alpha/SPEC.md": "alpha", "packages/beta/SPEC.md": "beta",
		"packages/alpha/evals/old_test.go": "old", "packages/alpha/api.go": "api",
		"packages/alpha/internal/existing.go": "old", "packages/alpha/gen/old.sql": "old",
		"packages/alpha/keep.txt": "old", "changes/0011-current.md": "current",
	} {
		full := filepath.Join(root, path)
		if err := os.MkdirAll(filepath.Dir(full), 0755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(full, []byte("fixture"), 0644); err != nil {
			t.Fatal(err)
		}
	}
	run := func(args ...string) string {
		cmd := exec.Command("git", args...)
		cmd.Dir = root
		cmd.Env = append(os.Environ(), "GIT_AUTHOR_NAME=Fixture", "GIT_AUTHOR_EMAIL=fixture@example.test", "GIT_COMMITTER_NAME=Fixture", "GIT_COMMITTER_EMAIL=fixture@example.test")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("git %v: %v: %s", args, err, out)
		}
		return string(out)
	}
	run("init", "-q")
	run("add", ".")
	run("commit", "-qm", "base")
	base := run("rev-parse", "HEAD")
	zones := contracts.Zones{Evals: []string{"evals/**"}, Interface: []string{"api.go"}, Disposable: []string{"internal/**"}, Preserve: []string{"keep.txt"}, AppendOnly: []string{"gen/**"}}
	snap := project.Snapshot{Root: root, Config: map[string]any{"protected": []any{"locked/**"}, "worker_writable": []any{"go.sum"}}, Modules: []project.Module{
		{Name: "alpha", Dir: "packages/alpha", Zones: zones},
		{Name: "beta", Dir: "packages/beta", Zones: zones},
	}}
	return snap, base[:len(base)-1]
}

func zoneRequest(snap project.Snapshot, base, phase string, changes ...gitstore.Change) checks.ZoneRequest {
	return checks.ZoneRequest{Snapshot: snap, Module: "alpha", Phase: phase, ChangePath: "changes/0011-current.md", Base: base, InterfaceDeclared: true, Changes: changes}
}

func zoneAllowed(t *testing.T, request checks.ZoneRequest) {
	t.Helper()
	if diags := checks.Zones(request); len(diags) != 0 {
		t.Errorf("wanted allowed change, got %#v", diags)
	}
}

func zoneRejected(t *testing.T, request checks.ZoneRequest) {
	t.Helper()
	if diags := checks.Zones(request); len(diags) == 0 {
		t.Errorf("forbidden change accepted: %#v", request.Changes)
	}
}

func Test_CHK_004_AuditOnlyCurrentChange(t *testing.T) {
	snap, base := zoneFixture(t)
	zoneAllowed(t, zoneRequest(snap, base, "audit", gitstore.Change{Path: "changes/0011-current.md", Status: "M"}))
	for _, path := range []string{"changes/0012-other.md", "packages/alpha/evals/old_test.go", "go.sum"} {
		zoneRejected(t, zoneRequest(snap, base, "audit", gitstore.Change{Path: path, Status: "M"}))
	}
}

func Test_CHK_004_CHK_006_CHK_007_CHK_008_CHK_009_RejectProtectedSpecForeignModuleAndChangeEveryPhase(t *testing.T) {
	snap, base := zoneFixture(t)
	for _, phase := range []string{"evals", "impl", "audit"} {
		for _, path := range []string{"locked/config", ".fireweed/hooks.go", "packages/alpha/SPEC.md", "packages/beta/internal/code.go", "changes/0012-other.md"} {
			zoneRejected(t, zoneRequest(snap, base, phase, gitstore.Change{Path: path, Status: "M"}))
		}
	}
}

func Test_CHK_004_CHK_011_CHK_012_CHK_013_CHK_014_CHK_015_EvalsAllowedZonesAndNewStubs(t *testing.T) {
	snap, base := zoneFixture(t)
	for _, change := range []gitstore.Change{
		{Path: "packages/alpha/evals/new_test.go", Status: "A"},
		{Path: "packages/alpha/api.go", Status: "M"},
		{Path: "packages/alpha/internal/new_stub.go", Status: "A"},
		{Path: "packages/alpha/manifest.json", Status: "M"},
		{Path: "changes/0011-current.md", Status: "M"},
	} {
		zoneAllowed(t, zoneRequest(snap, base, "evals", change))
	}
	zoneRejected(t, zoneRequest(snap, base, "evals", gitstore.Change{Path: "packages/alpha/internal/existing.go", Status: "M"}))
	zoneRejected(t, zoneRequest(snap, base, "evals", gitstore.Change{Path: "packages/alpha/keep.txt", Status: "M"}))
	request := zoneRequest(snap, base, "evals", gitstore.Change{Path: "packages/alpha/api.go", Status: "M"})
	request.InterfaceDeclared = false
	zoneRejected(t, request)
}

func Test_CHK_004_CHK_016_CHK_017_CHK_018_CHK_019_CHK_020_ImplFreezesEvalsAndInterfaces(t *testing.T) {
	snap, base := zoneFixture(t)
	for _, path := range []string{"packages/alpha/evals/old_test.go", "packages/alpha/api.go"} {
		zoneRejected(t, zoneRequest(snap, base, "impl", gitstore.Change{Path: path, Status: "M"}))
	}
	for _, path := range []string{"packages/alpha/internal/existing.go", "packages/alpha/keep.txt", "packages/alpha/manifest.json", "changes/0011-current.md"} {
		zoneAllowed(t, zoneRequest(snap, base, "impl", gitstore.Change{Path: path, Status: "M"}))
	}
}

func Test_CHK_004_ImplAppendOnlyNewVersusExisting(t *testing.T) {
	snap, base := zoneFixture(t)
	zoneAllowed(t, zoneRequest(snap, base, "impl", gitstore.Change{Path: "packages/alpha/gen/new.sql", Status: "A"}))
	zoneRejected(t, zoneRequest(snap, base, "impl", gitstore.Change{Path: "packages/alpha/gen/old.sql", Status: "M"}))
	zoneRejected(t, zoneRequest(snap, base, "impl", gitstore.Change{Path: "packages/alpha/gen/old.sql", Status: "D"}))
	zoneRejected(t, zoneRequest(snap, base, "evals", gitstore.Change{Path: "packages/alpha/gen/new.sql", Status: "A"}))
}

func Test_CHK_004_RenamesCheckBothPaths(t *testing.T) {
	snap, base := zoneFixture(t)
	zoneRejected(t, zoneRequest(snap, base, "impl", gitstore.Change{OldPath: "packages/beta/internal/old.go", Path: "packages/alpha/internal/new.go", Status: "R100"}))
	zoneRejected(t, zoneRequest(snap, base, "impl", gitstore.Change{OldPath: "packages/alpha/internal/existing.go", Path: "packages/beta/internal/new.go", Status: "R100"}))
	zoneRejected(t, zoneRequest(snap, base, "impl", gitstore.Change{OldPath: "packages/alpha/gen/old.sql", Path: "packages/alpha/internal/new.go", Status: "R100"}))
	zoneRejected(t, zoneRequest(snap, base, "impl", gitstore.Change{OldPath: "packages/alpha/internal/existing.go", Path: "packages/alpha/gen/new.sql", Status: "R100"}))
}

func Test_CHK_004_CHK_010_WorkerWritableOnlyEvalsAndImpl(t *testing.T) {
	snap, base := zoneFixture(t)
	for _, phase := range []string{"evals", "impl"} {
		zoneAllowed(t, zoneRequest(snap, base, phase, gitstore.Change{Path: "go.sum", Status: "M"}))
		zoneRejected(t, zoneRequest(snap, base, phase, gitstore.Change{Path: "outside.txt", Status: "M"}))
	}
	zoneRejected(t, zoneRequest(snap, base, "audit", gitstore.Change{Path: "go.sum", Status: "M"}))
}
