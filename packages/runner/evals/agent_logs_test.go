package evals_test

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/lurkshark/fireweed/packages/runner"
)

func assertLogFile(t *testing.T, path, root, run, phase string, attempt int, streams ...string) []byte {
	t.Helper()
	want := filepath.Join(root, ".fireweed", "runs", run, fmt.Sprintf("%s-%d.log", phase, attempt))
	if path != want || !filepath.IsAbs(path) {
		t.Errorf("log path = %q, want absolute %q", path, want)
	}
	if path == "" {
		return nil
	}
	data, err := os.ReadFile(path)
	if err != nil {
		t.Errorf("read agent log %q: %v", path, err)
		return nil
	}
	for _, stream := range streams {
		if !strings.Contains(string(data), stream) {
			t.Errorf("log %q lacks complete stream %q: %q", path, stream, data)
		}
	}
	return data
}

func Test_RUN_063_InvokeAgentPersistsStreamsOnFailureTimeoutCancellationAndAppend(t *testing.T) {
	snap := agentFixture(t)
	logRoot := t.TempDir()
	request := request(snap, "worker", "")
	request.LogRoot = logRoot
	request.Run = "fixture-run"
	request.Phase = "evals"
	request.Attempt = 2
	var prior []byte
	for _, tc := range []struct {
		name, command, stdout, stderr string
		cancel                        bool
		wantError                     bool
	}{
		{"failure", `printf 'failure stdout Ω'; printf 'failure stderr β' >&2; exit 7`, "failure stdout Ω", "failure stderr β", false, true},
		{"timeout", `printf 'timeout stdout Ω'; printf 'timeout stderr β' >&2; sleep 5`, "timeout stdout Ω", "timeout stderr β", false, true},
		{"cancellation", `printf 'cancel stdout Ω'; printf 'cancel stderr β' >&2; sleep 5`, "cancel stdout Ω", "cancel stderr β", true, true},
		{"append", `printf 'second stdout Ω'; printf 'second stderr β' >&2`, "second stdout Ω", "second stderr β", false, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			command := tc.command
			request.Timeout = 150 * time.Millisecond
			ctx := context.Background()
			if tc.cancel {
				ready := filepath.Join(logRoot, "cancel.ready")
				command = `printf 'cancel stdout Ω'; printf 'cancel stderr β' >&2; printf ready > ` + shellQuote(ready) + `; sleep 5`
				request.Timeout = 3 * time.Second
				var cancel context.CancelFunc
				ctx, cancel = context.WithCancel(ctx)
				go func() {
					for i := 0; i < 200; i++ {
						if _, err := os.Stat(ready); err == nil {
							cancel()
							return
						}
						time.Sleep(10 * time.Millisecond)
					}
					cancel()
				}()
				defer cancel()
			}
			snap.Config["agent"] = map[string]any{"command": []any{"/bin/sh", "-c", command}}
			request.Snapshot = snap
			result, err := runner.InvokeAgent(ctx, request)
			if (err != nil) != tc.wantError || result.Stdout != tc.stdout || result.Stderr != tc.stderr {
				t.Errorf("invocation = %#v, %v", result, err)
			}
			if tc.cancel && !errors.Is(err, context.Canceled) {
				t.Errorf("cancellation returned %v", err)
			}
			path := result.LogPath
			data := assertLogFile(t, path, logRoot, "fixture-run", "evals", 2, tc.stdout, tc.stderr)
			if len(prior) > 0 && !strings.HasPrefix(string(data), string(prior)) {
				t.Error("repeated invocation did not append after previous log records")
			}
			prior = data
			if tc.name == "append" {
				for _, earlier := range []string{"failure stdout Ω", "failure stderr β", "timeout stdout Ω", "timeout stderr β", "cancel stdout Ω", "cancel stderr β"} {
					if !strings.Contains(string(data), earlier) {
						t.Errorf("append discarded earlier output %q", earlier)
					}
				}
			}
		})
	}
}

func Test_RUN_064_NextLogsHeldRunPhaseAndOneBasedAttemptsOutsideAuditInput(t *testing.T) {
	f := newNextFixture(t, true, true, "retry-both")
	result, err := f.run(false)
	if err != nil || result.State.Phase != "applied" || len(result.Steps) != 5 {
		t.Fatalf("Next fixture = %#v, %v", result, err)
	}
	for i, want := range []struct {
		phase   string
		attempt int
	}{
		{"evals", 1}, {"evals", 2}, {"audit", 1}, {"impl", 1}, {"impl", 2},
	} {
		step := result.Steps[i]
		if step.Phase != want.phase {
			t.Errorf("step %d phase = %q", i, step.Phase)
		}
		path := step.Agent.LogPath
		assertLogFile(t, path, f.root, result.State.Run, want.phase, want.attempt, step.Agent.Stdout)
		if want.phase == "audit" {
			cwd := nextRead(t, f.logDir, "audit.cwd")
			if cwd == f.root || strings.HasPrefix(path, cwd+string(os.PathSeparator)) {
				t.Errorf("audit log %q placed in restricted input %q", path, cwd)
			}
		}
	}
}

func Test_RUN_065_DrillLogRunIsStableAcrossUsageAndReconstructionRetries(t *testing.T) {
	for _, mode := range []string{"usage_once", "agent_retry"} {
		t.Run(mode, func(t *testing.T) {
			f := newDrillFixture(t)
			f.mode(t, mode)
			result, err := f.run(context.Background(), 0)
			if err != nil || result.Result != "pass" || len(result.Steps) != 2 {
				t.Fatalf("Drill fixture = %#v, %v", result, err)
			}
			var run string
			for i, step := range result.Steps {
				path := step.LogPath
				rel, relErr := filepath.Rel(filepath.Join(f.root, ".fireweed", "runs"), path)
				if relErr != nil || strings.HasPrefix(rel, "..") || !filepath.IsAbs(path) {
					t.Errorf("drill log %q is outside invoking project: %v", path, relErr)
					continue
				}
				parts := strings.Split(filepath.ToSlash(rel), "/")
				if len(parts) != 2 || parts[0] == "" {
					t.Errorf("drill log relative path = %q", rel)
					continue
				}
				if i == 0 {
					run = parts[0]
				} else if parts[0] != run {
					t.Errorf("reconstruction retry changed log run from %q to %q", run, parts[0])
				}
				attempt := i + 1
				if mode == "usage_once" {
					attempt = 1
				}
				assertLogFile(t, path, f.root, run, "drill", attempt, step.Stdout, step.Stderr)
			}
			if mode == "usage_once" && result.Steps[0].LogPath != result.Steps[1].LogPath {
				t.Error("usage retry spent a reconstruction log attempt")
			}
			f.mode(t, "success")
			second, secondErr := f.run(context.Background(), 0)
			if secondErr != nil || second.Result != "pass" || len(second.Steps) == 0 {
				t.Errorf("second Drill invocation = %#v, %v", second, secondErr)
			} else {
				secondPath := second.Steps[0].LogPath
				secondRel, relErr := filepath.Rel(filepath.Join(f.root, ".fireweed", "runs"), secondPath)
				if relErr != nil || strings.HasPrefix(secondRel, "..") || strings.Split(filepath.ToSlash(secondRel), "/")[0] == run {
					t.Errorf("separate Drill invocation reused log run %q: %v", secondRel, relErr)
				}
			}
			drillCleaned(t, f)
		})
	}
}

func Test_RUN_066_InvokeAgentLogContextDefaultsAndExplicitFields(t *testing.T) {
	snap := agentFixture(t)
	snap.Config["agent"] = map[string]any{"command": []any{"/bin/sh", "-c", `printf 'default output'; printf 'default error' >&2`}}
	var previous string
	for i := 0; i < 2; i++ {
		result, err := runner.InvokeAgent(context.Background(), request(snap, "worker", ""))
		if err != nil {
			t.Fatalf("default invocation: %v", err)
		}
		path := result.LogPath
		rel, relErr := filepath.Rel(filepath.Join(snap.Root, ".fireweed", "runs"), path)
		if relErr != nil || strings.HasPrefix(rel, "..") {
			t.Errorf("default log %q outside Snapshot.Root: %v", path, relErr)
			continue
		}
		parts := strings.Split(filepath.ToSlash(rel), "/")
		if len(parts) != 2 || parts[0] == "" || parts[1] != "worker-1.log" {
			t.Errorf("default log relative path = %q", rel)
			continue
		}
		assertLogFile(t, path, snap.Root, parts[0], "worker", 1, "default output", "default error")
		if i == 1 && path == previous {
			t.Error("default run identifier was reused")
		}
		previous = path
	}
	chosenRoot := t.TempDir()
	explicit := request(snap, "worker", "")
	explicit.LogRoot = chosenRoot
	explicit.Run = "chosen-run"
	explicit.Phase = "chosen-phase"
	explicit.Attempt = 3
	result, err := runner.InvokeAgent(context.Background(), explicit)
	if err != nil {
		t.Fatalf("explicit invocation: %v", err)
	}
	assertLogFile(t, result.LogPath, chosenRoot, "chosen-run", "chosen-phase", 3, "default output", "default error")
}
