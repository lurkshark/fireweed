package evals_test

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/lurkshark/fireweed/packages/checks"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
)

func boundarySnapshot(t *testing.T, hook any) project.Snapshot {
	t.Helper()
	root := t.TempDir()
	snap := project.Snapshot{Root: root, Config: map[string]any{}, Modules: []project.Module{{
		Name: "alpha", Dir: "packages/alpha",
		Zones:    contracts.Zones{Disposable: []string{"internal/**"}},
		Document: contracts.Document{Frontmatter: map[string]any{}},
	}}}
	if hook != nil {
		snap.Config["hooks"] = map[string]any{"boundaries": hook}
	}
	return snap
}

func boundaryScript(t *testing.T, root, body string) []any {
	t.Helper()
	path := filepath.Join(root, "boundary-hook.sh")
	if err := os.WriteFile(path, []byte("#!/bin/sh\n"+body+"\n"), 0700); err != nil {
		t.Fatal(err)
	}
	return []any{path}
}

func boundaryHasError(diags []project.Diagnostic, err error) bool {
	if err != nil {
		return true
	}
	for _, d := range diags {
		if !d.Warning {
			return true
		}
	}
	return false
}

func boundaryHasWarning(diags []project.Diagnostic) bool {
	for _, d := range diags {
		if d.Warning {
			return true
		}
	}
	return false
}

func Test_CHK_005_CHK_039_GlobalAndModuleBoundaryHooks(t *testing.T) {
	snap := boundarySnapshot(t, nil)
	globalMarker := filepath.Join(snap.Root, "global-ran")
	moduleMarker := filepath.Join(snap.Root, "module-ran")
	snap.Config["hooks"] = map[string]any{"boundaries": boundaryScript(t, snap.Root, fmt.Sprintf("printf run > %q; printf 'alpha: global violation\\n'", globalMarker))}
	moduleHook := filepath.Join(snap.Root, "module-hook.sh")
	if err := os.WriteFile(moduleHook, []byte(fmt.Sprintf("#!/bin/sh\nprintf run > %q; printf 'beta: module violation\\n'\n", moduleMarker)), 0700); err != nil {
		t.Fatal(err)
	}
	snap.Modules = append(snap.Modules, project.Module{Name: "beta", Dir: "packages/beta", Zones: contracts.Zones{Disposable: []string{"internal/**"}}, Document: contracts.Document{Frontmatter: map[string]any{"hooks": map[string]any{"boundaries": []any{moduleHook}}}}})
	diags, err := checks.Boundaries(context.Background(), checks.BoundaryRequest{Snapshot: snap, Timeout: 2 * time.Second})
	if !boundaryHasError(diags, err) {
		t.Errorf("unsuppressed violations accepted: %#v, %v", diags, err)
	}
	for _, module := range []string{"alpha", "beta"} {
		if !strings.Contains(fmt.Sprint(diags, err), module) {
			t.Errorf("violation for %s not attributed: %#v, %v", module, diags, err)
		}
	}
	for _, path := range []string{globalMarker, moduleMarker} {
		if _, statErr := os.Stat(path); statErr != nil {
			t.Errorf("applicable hook did not run: %s: %v", path, statErr)
		}
	}
}

func Test_CHK_041_ExactBaselineOnly(t *testing.T) {
	snap := boundarySnapshot(t, nil)
	if err := os.MkdirAll(filepath.Join(snap.Root, ".fireweed"), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(snap.Root, ".fireweed", "boundaries-baseline"), []byte("alpha: old violation\n"), 0644); err != nil {
		t.Fatal(err)
	}
	snap.Config["hooks"] = map[string]any{"boundaries": boundaryScript(t, snap.Root, `printf 'alpha: old violation\n'`)}
	diags, err := checks.Boundaries(context.Background(), checks.BoundaryRequest{Snapshot: snap, Timeout: time.Second})
	if boundaryHasError(diags, err) {
		t.Errorf("exact baseline line rejected: %#v, %v", diags, err)
	}
	snap.Config["hooks"] = map[string]any{"boundaries": boundaryScript(t, snap.Root, `printf 'alpha: old violation changed\n'`)}
	diags, err = checks.Boundaries(context.Background(), checks.BoundaryRequest{Snapshot: snap, Timeout: time.Second})
	if !boundaryHasError(diags, err) {
		t.Error("non-exact baseline line suppressed")
	}
}

func Test_CHK_040_CHK_042_OptOutSuppressesOnlyItsModuleAndWarns(t *testing.T) {
	snap := boundarySnapshot(t, nil)
	snap.Modules[0].Document.Frontmatter["boundaries_opt_out"] = "legacy dependency"
	snap.Modules = append(snap.Modules, project.Module{Name: "beta", Dir: "packages/beta", Zones: contracts.Zones{Disposable: []string{"internal/**"}}, Document: contracts.Document{Frontmatter: map[string]any{}}})
	snap.Config["hooks"] = map[string]any{"boundaries": boundaryScript(t, snap.Root, `printf 'alpha: ignored\n'`)}
	diags, err := checks.Boundaries(context.Background(), checks.BoundaryRequest{Snapshot: snap, Timeout: time.Second})
	if boundaryHasError(diags, err) || !boundaryHasWarning(diags) {
		t.Errorf("reasoned opt-out = %#v, %v", diags, err)
	}
	if !strings.Contains(fmt.Sprint(diags), "legacy dependency") {
		t.Errorf("opt-out reason absent from warning: %#v", diags)
	}
	snap.Config["hooks"] = map[string]any{"boundaries": boundaryScript(t, snap.Root, `printf 'alpha: ignored\nbeta: enforced\n'`)}
	diags, err = checks.Boundaries(context.Background(), checks.BoundaryRequest{Snapshot: snap, Timeout: time.Second})
	if !boundaryHasError(diags, err) {
		t.Error("opt-out hid another module's violation")
	}
}

func Test_CHK_044_CHK_045_MissingHookRequiredVersusOptional(t *testing.T) {
	snap := boundarySnapshot(t, nil)
	diags, err := checks.Boundaries(context.Background(), checks.BoundaryRequest{Snapshot: snap, Timeout: time.Second})
	if !boundaryHasError(diags, err) {
		t.Error("missing hook for disposable module accepted")
	}
	snap.Modules[0].Zones.Disposable = nil
	diags, err = checks.Boundaries(context.Background(), checks.BoundaryRequest{Snapshot: snap, Timeout: time.Second})
	if boundaryHasError(diags, err) || !boundaryHasWarning(diags) {
		t.Errorf("missing optional hook = %#v, %v", diags, err)
	}
	snap.Modules[0].Zones.Disposable = []string{"internal/**"}
	snap.Modules[0].Document.Frontmatter["boundaries_opt_out"] = "legacy"
	diags, err = checks.Boundaries(context.Background(), checks.BoundaryRequest{Snapshot: snap, Timeout: time.Second})
	if boundaryHasError(diags, err) || !boundaryHasWarning(diags) {
		t.Errorf("missing hook with reasoned opt-out = %#v, %v", diags, err)
	}
}

func Test_CHK_041_CHK_042_NormalNonzeroExitCanUseBaselineAndOptOut(t *testing.T) {
	snap := boundarySnapshot(t, nil)
	if err := os.MkdirAll(filepath.Join(snap.Root, ".fireweed"), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(snap.Root, ".fireweed", "boundaries-baseline"), []byte("alpha: known\n"), 0644); err != nil {
		t.Fatal(err)
	}
	snap.Modules[0].Document.Frontmatter["boundaries_opt_out"] = "legacy"
	snap.Config["hooks"] = map[string]any{"boundaries": boundaryScript(t, snap.Root, `printf 'alpha: known\nalpha: opted out\n'; exit 3`)}
	diags, err := checks.Boundaries(context.Background(), checks.BoundaryRequest{Snapshot: snap, Timeout: time.Second})
	if boundaryHasError(diags, err) {
		t.Errorf("normally exited hook with fully suppressed lines rejected: %#v, %v", diags, err)
	}
}

func Test_CHK_043_LaunchCancellationAndTimeoutStayErrors(t *testing.T) {
	snap := boundarySnapshot(t, nil)
	snap.Modules[0].Document.Frontmatter["boundaries_opt_out"] = "legacy"
	snap.Config["hooks"] = map[string]any{"boundaries": []any{filepath.Join(snap.Root, "missing-executable")}}
	diags, err := checks.Boundaries(context.Background(), checks.BoundaryRequest{Snapshot: snap, Timeout: time.Second})
	if !boundaryHasError(diags, err) {
		t.Error("launch failure hidden by opt-out")
	}
	snap.Config["hooks"] = map[string]any{"boundaries": boundaryScript(t, snap.Root, `sleep 2`)}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	diags, err = checks.Boundaries(ctx, checks.BoundaryRequest{Snapshot: snap, Timeout: time.Second})
	if !boundaryHasError(diags, err) {
		t.Error("cancelled hook hidden by opt-out")
	}
	start := time.Now()
	diags, err = checks.Boundaries(context.Background(), checks.BoundaryRequest{Snapshot: snap, Timeout: 80 * time.Millisecond})
	if !boundaryHasError(diags, err) {
		t.Error("timed-out hook hidden by opt-out")
	}
	if time.Since(start) > 3*time.Second {
		t.Error("boundary timeout did not bound execution")
	}
}
