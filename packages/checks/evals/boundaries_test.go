package evals_test

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
	"time"
	"unicode/utf8"

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

func requireBoundaryDiagnostic(t *testing.T, diags []project.Diagnostic, code string, warning bool, message string) {
	t.Helper()
	for _, d := range diags {
		if message == "" {
			if d.Code == code && d.Warning == warning {
				return
			}
			continue
		}
		if strings.Contains(d.Message, message) {
			if d.Code != code || d.Warning != warning {
				t.Errorf("boundary diagnostic for %q = %#v; want code %s, warning %t", message, d, code, warning)
			}
			return
		}
	}
	t.Errorf("boundary diagnostic for %q absent: %#v", message, diags)
}

func Test_CHK_005_GlobalAndModuleBoundaryHooks(t *testing.T) {
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

func Test_CHK_039_CHK_040_OptOutSuppressesOnlyItsModuleAndWarns(t *testing.T) {
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
	requireBoundaryDiagnostic(t, diags, "X03", true, "legacy dependency")
	snap.Config["hooks"] = map[string]any{"boundaries": boundaryScript(t, snap.Root, `printf 'alpha: ignored\nbeta: enforced\n'`)}
	diags, err = checks.Boundaries(context.Background(), checks.BoundaryRequest{Snapshot: snap, Timeout: time.Second})
	if !boundaryHasError(diags, err) {
		t.Error("opt-out hid another module's violation")
	}
	if err == nil {
		t.Error("unsuppressed beta violation returned no error")
	}
	requireBoundaryDiagnostic(t, diags, "X01", false, "beta: enforced")
}

func Test_CHK_044_CHK_045_MissingHookRequiredVersusOptional(t *testing.T) {
	snap := boundarySnapshot(t, nil)
	diags, err := checks.Boundaries(context.Background(), checks.BoundaryRequest{Snapshot: snap, Timeout: time.Second})
	if !boundaryHasError(diags, err) {
		t.Error("missing hook for disposable module accepted")
	}
	if err == nil {
		t.Error("missing required hook returned no error")
	}
	requireBoundaryDiagnostic(t, diags, "X02", false, "alpha: missing boundaries hook")
	for _, d := range diags {
		if d.Code == "X04" {
			t.Errorf("missing required hook returned X04: %#v", d)
		}
	}
	snap.Modules[0].Zones.Disposable = nil
	diags, err = checks.Boundaries(context.Background(), checks.BoundaryRequest{Snapshot: snap, Timeout: time.Second})
	if boundaryHasError(diags, err) || !boundaryHasWarning(diags) {
		t.Errorf("missing optional hook = %#v, %v", diags, err)
	}
	requireBoundaryDiagnostic(t, diags, "X04", true, "no boundary enforcement required")
	snap.Modules[0].Zones.Disposable = []string{"internal/**"}
	snap.Modules[0].Document.Frontmatter["boundaries_opt_out"] = "legacy"
	diags, err = checks.Boundaries(context.Background(), checks.BoundaryRequest{Snapshot: snap, Timeout: time.Second})
	if boundaryHasError(diags, err) || !boundaryHasWarning(diags) {
		t.Errorf("missing hook with reasoned opt-out = %#v, %v", diags, err)
	}
	requireBoundaryDiagnostic(t, diags, "X03", true, "legacy")
	requireBoundaryDiagnostic(t, diags, "X04", true, "no boundary enforcement required")
}

func Test_CHK_041_CHK_042_NormalNonzeroExitCanUseBaselineAndOptOut(t *testing.T) {
	snap := boundarySnapshot(t, nil)
	if err := os.MkdirAll(filepath.Join(snap.Root, ".fireweed"), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(snap.Root, ".fireweed", "boundaries-baseline"), []byte("alpha: known\n"), 0644); err != nil {
		t.Fatal(err)
	}
	snap.Modules = append(snap.Modules, project.Module{
		Name: "beta", Dir: "packages/beta",
		Document: contracts.Document{Frontmatter: map[string]any{"boundaries_opt_out": "legacy"}},
	})
	snap.Config["hooks"] = map[string]any{"boundaries": boundaryScript(t, snap.Root, `printf 'alpha: known\nbeta: opted out\n'; exit 3`)}
	diags, err := checks.Boundaries(context.Background(), checks.BoundaryRequest{Snapshot: snap, Timeout: time.Second})
	if boundaryHasError(diags, err) {
		t.Errorf("normally exited hook with fully suppressed lines rejected: %#v, %v", diags, err)
	}
}

func requireBoundaryExitFailure(t *testing.T, diags []project.Diagnostic, err error, status int) string {
	t.Helper()
	exitStatus := regexp.MustCompile(fmt.Sprintf(`(?i)\b(?:exit(?:ed)?|status|code)\b[^0-9\n]*\b%d\b`, status))
	if err == nil || !exitStatus.MatchString(err.Error()) || strings.HasPrefix(err.Error(), "X02:") {
		t.Errorf("boundary hook error = %v; want exit status %d without X02 prefix", err, status)
	}
	for _, d := range diags {
		if d.Code == "X02" && !d.Warning {
			return d.Message
		}
	}
	t.Fatalf("missing X02 error diagnostic: %#v", diags)
	return ""
}

func Test_CHK_043_CHK_086_RequiredGlobalHookRejectsUnattributedNonzeroExit(t *testing.T) {
	const reason = "startup error: missing tool"
	snap := boundarySnapshot(t, nil)
	snap.Config["hooks"] = map[string]any{"boundaries": boundaryScript(t, snap.Root, fmt.Sprintf(`printf '%%s\n' '%s' >&2; exit 3`, reason))}
	diags, err := checks.Boundaries(context.Background(), checks.BoundaryRequest{Snapshot: snap, Timeout: time.Second})
	message := requireBoundaryExitFailure(t, diags, err, 3)
	if !regexp.MustCompile(`(?i)\bglobal\b`).MatchString(message) || !strings.Contains(message, reason) {
		t.Errorf("X02 message = %q; want the global hook scope and stderr %q", message, reason)
	}
}

func Test_CHK_086_OverrideFailureIncludesBoundedUTF8StderrTail(t *testing.T) {
	const (
		head = "override-stderr-head"
		fill = "界"
		tail = "override-stderr-tail"
	)
	// A 4,096-byte suffix of this payload starts inside a three-byte rune.
	stderr := head + strings.Repeat(fill, 2000) + tail
	snap := boundarySnapshot(t, nil)
	snap.Modules[0].Document.Frontmatter["hooks"] = map[string]any{
		"boundaries": boundaryScript(t, snap.Root, fmt.Sprintf(`printf '%%s' '%s' >&2; exit 7`, stderr)),
	}
	diags, err := checks.Boundaries(context.Background(), checks.BoundaryRequest{Snapshot: snap, Timeout: time.Second})
	message := requireBoundaryExitFailure(t, diags, err, 7)
	if !regexp.MustCompile(`\balpha\b`).MatchString(message) {
		t.Error("X02 message does not identify the alpha override")
	}
	if !utf8.ValidString(message) {
		t.Error("X02 message is not valid UTF-8")
	}
	if strings.Contains(message, head) || !strings.Contains(message, fill+tail) {
		t.Error("X02 message must omit the stderr head and retain its UTF-8 tail")
	}
	excerptBytes := strings.Count(message, fill)*len(fill) + strings.Count(message, tail)*len(tail)
	excerptBytes += strings.Count(message, "\uFFFD") * len("\uFFFD")
	if excerptBytes > 4096 {
		t.Errorf("X02 contains at least %d bytes of stderr; want at most 4096", excerptBytes)
	}
	omission := regexp.MustCompile(`(?i)\b(?:truncat(?:ed|ion)|omitted|elided)\b|…|\.{3}`)
	if !omission.MatchString(message) {
		t.Error("X02 message lacks a visible truncation marker")
	}
}

func Test_CHK_043_SilentNonzeroOverrideIsIndependentOfGlobalViolation(t *testing.T) {
	snap := boundarySnapshot(t, `printf 'alpha: global violation\n'; exit 3`)
	snap.Modules[0].Document.Frontmatter["hooks"] = map[string]any{"boundaries": boundaryScript(t, snap.Root, `exit 7`)}
	diags, err := checks.Boundaries(context.Background(), checks.BoundaryRequest{Snapshot: snap, Timeout: time.Second})
	requireBoundaryDiagnostic(t, diags, "X01", false, "alpha: global violation")
	requireBoundaryExitFailure(t, diags, err, 7)
}

func Test_CHK_043_LaunchCancellationAndTimeoutStayErrors(t *testing.T) {
	snap := boundarySnapshot(t, nil)
	snap.Modules[0].Document.Frontmatter["boundaries_opt_out"] = "legacy"
	snap.Config["hooks"] = map[string]any{"boundaries": []any{filepath.Join(snap.Root, "missing-executable")}}
	diags, err := checks.Boundaries(context.Background(), checks.BoundaryRequest{Snapshot: snap, Timeout: time.Second})
	if !boundaryHasError(diags, err) {
		t.Error("launch failure hidden by opt-out")
	}
	if err == nil {
		t.Error("launch failure returned no error")
	} else if !strings.Contains(err.Error(), "missing-executable") || strings.HasPrefix(err.Error(), "X02:") {
		t.Errorf("launch error = %q; want missing executable cause without X02 prefix", err)
	}
	requireBoundaryDiagnostic(t, diags, "X02", false, "")
	snap.Config["hooks"] = map[string]any{"boundaries": boundaryScript(t, snap.Root, `sleep 2`)}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	diags, err = checks.Boundaries(ctx, checks.BoundaryRequest{Snapshot: snap, Timeout: time.Second})
	if !boundaryHasError(diags, err) {
		t.Error("cancelled hook hidden by opt-out")
	}
	if err == nil {
		t.Error("cancelled hook returned no error")
	} else if !errors.Is(err, context.Canceled) || strings.HasPrefix(err.Error(), "X02:") {
		t.Errorf("cancellation error = %q; want canceled cause without X02 prefix", err)
	}
	requireBoundaryDiagnostic(t, diags, "X02", false, "")
	start := time.Now()
	diags, err = checks.Boundaries(context.Background(), checks.BoundaryRequest{Snapshot: snap, Timeout: 80 * time.Millisecond})
	if !boundaryHasError(diags, err) {
		t.Error("timed-out hook hidden by opt-out")
	}
	if err == nil {
		t.Error("timed-out hook returned no error")
	} else if !errors.Is(err, context.DeadlineExceeded) || strings.HasPrefix(err.Error(), "X02:") {
		t.Errorf("timeout error = %q; want deadline cause without X02 prefix", err)
	}
	requireBoundaryDiagnostic(t, diags, "X02", false, "")
	if time.Since(start) > 3*time.Second {
		t.Error("boundary timeout did not bound execution")
	}
}
