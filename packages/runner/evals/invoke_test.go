package evals_test

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/runner"
)

func agentFixture(t *testing.T) project.Snapshot {
	t.Helper()
	root, err := filepath.EvalSymlinks(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	for _, args := range [][]string{
		{"init", "-q", root},
		{"-C", root, "config", "user.name", "Fixture"},
		{"-C", root, "config", "user.email", "fixture@example.invalid"},
		{"-C", root, "commit", "-q", "--allow-empty", "-m", "fixture"},
	} {
		if out, err := exec.Command("git", args...).CombinedOutput(); err != nil {
			t.Fatalf("git %v: %v: %s", args, err, out)
		}
	}
	return project.Snapshot{Root: root, Config: map[string]any{}}
}

func agentScript(t *testing.T, root, name, body string) string {
	t.Helper()
	path := filepath.Join(root, name)
	if err := os.WriteFile(path, []byte("#!/bin/sh\n"+body+"\n"), 0755); err != nil {
		t.Fatal(err)
	}
	return path
}

func shellQuote(s string) string { return "'" + strings.ReplaceAll(s, "'", "'\\''") + "'" }

func request(s project.Snapshot, role, prompt string) runner.AgentRequest {
	return runner.AgentRequest{Snapshot: s, Role: role, Prompt: prompt, Bin: "/fixture/fireweed"}
}

func Test_RUN_001_RUN_002_RUN_005_RUN_006_CommandSelectionDirectoryEnvAndOutput(t *testing.T) {
	snap := agentFixture(t)
	worker := agentScript(t, snap.Root, "worker.sh", `printf 'worker|%s|%s' "$PWD" "$FIREWEED_BIN"; printf 'worker-stderr' >&2`)
	audit := agentScript(t, snap.Root, "audit.sh", `printf 'audit|%s|%s' "$PWD" "$FIREWEED_BIN"; printf 'audit-stderr' >&2`)
	snap.Config["agent"] = map[string]any{"command": []any{worker}}
	snap.Config["audit"] = map[string]any{"command": []any{audit}}
	for _, tc := range []struct{ name, role, want string }{
		{"worker", "worker", "worker"},
		{"audit", "audit", "audit"},
		{"audit_fallback", "audit", "worker"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if tc.name == "audit_fallback" {
				delete(snap.Config, "audit")
			}
			result, err := runner.InvokeAgent(context.Background(), request(snap, tc.role, "prompt"))
			if err != nil || result.ExitCode != 0 || result.Stdout != tc.want+"|"+snap.Root+"|/fixture/fireweed" || result.Stderr != tc.want+"-stderr" || result.UsageLimit {
				t.Errorf("InvokeAgent(%s) = %#v, %v", tc.role, result, err)
			}
		})
	}
}

func Test_RUN_003_ShellAndArgvExpandOnceAndPreserveValues(t *testing.T) {
	snap := agentFixture(t)
	common, err := gitstore.Inspect(snap.Root)
	if err != nil {
		t.Fatal(err)
	}
	printArgs := agentScript(t, snap.Root, "args.sh", `printf '%s\n' "$1" "$2"`)
	prompt := `a'b ; $(printf hacked) {git_common_dir}`
	for _, tc := range []struct {
		name    string
		command any
	}{
		{"shell", shellQuote(printArgs) + " {prompt} {git_common_dir}"},
		{"argv", []any{printArgs, "{prompt}", "{git_common_dir}"}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			snap.Config["agent"] = map[string]any{"command": tc.command}
			result, err := runner.InvokeAgent(context.Background(), request(snap, "worker", prompt))
			want := prompt + "\n" + common.CommonDir + "\n"
			if err != nil || result.Stdout != want || result.ExitCode != 0 {
				t.Errorf("%s substitution = %#v, %v; want stdout %q", tc.name, result, err, want)
			}
		})
	}
}

func Test_RUN_003_RUN_004_WrapperPrefixForBothRoles(t *testing.T) {
	snap := agentFixture(t)
	common, err := gitstore.Inspect(snap.Root)
	if err != nil {
		t.Fatal(err)
	}
	log := filepath.Join(snap.Root, "wrap.log")
	wrap := agentScript(t, snap.Root, "wrap.sh", `printf '%s\n%s\n' "$1" "$2" > "$3"; shift 3; exec "$@"`)
	worker := agentScript(t, snap.Root, "worker.sh", `printf worker`)
	audit := agentScript(t, snap.Root, "audit.sh", `printf audit`)
	snap.Config["agent"] = map[string]any{"command": []any{worker}, "wrap": []any{wrap, "{prompt}", "{git_common_dir}", log}}
	snap.Config["audit"] = map[string]any{"command": []any{audit}}
	prompt := `wrap 'quoted' {git_common_dir}`
	for _, role := range []string{"worker", "audit"} {
		t.Run(role, func(t *testing.T) {
			result, err := runner.InvokeAgent(context.Background(), request(snap, role, prompt))
			data, readErr := os.ReadFile(log)
			if readErr != nil {
				t.Fatalf("wrapper did not run: %v; InvokeAgent = %#v, %v", readErr, result, err)
			}
			if err != nil || result.Stdout != role || string(data) != prompt+"\n"+common.CommonDir+"\n" {
				t.Errorf("wrapped %s = %#v, %v; log=%q", role, result, err, data)
			}
		})
	}
}

func Test_RUN_006_RUN_009_NonzeroExitRetainsStreamsAndCode(t *testing.T) {
	snap := agentFixture(t)
	snap.Config["agent"] = map[string]any{"command": []any{"/bin/sh", "-c", `printf output; printf warning >&2; exit 7`}}
	result, err := runner.InvokeAgent(context.Background(), request(snap, "worker", ""))
	if err == nil || result.ExitCode != 7 || result.Stdout != "output" || result.Stderr != "warning" || result.UsageLimit {
		t.Errorf("nonzero exit = %#v, %v", result, err)
	}
}

func Test_RUN_007_RUN_009_TimeoutSelectionAndCapturedOutput(t *testing.T) {
	snap := agentFixture(t)
	snap.Config["agent"] = map[string]any{"command": []any{"/bin/sh", "-c", `printf begun; printf warning >&2; sleep 2`}}
	for _, tc := range []struct {
		name    string
		config  string
		request time.Duration
	}{
		{"configured", "120ms", 0},
		{"request", "2s", 120 * time.Millisecond},
	} {
		t.Run(tc.name, func(t *testing.T) {
			snap.Config["timeout"] = tc.config
			req := request(snap, "worker", "")
			req.Timeout = tc.request
			start := time.Now()
			result, err := runner.InvokeAgent(context.Background(), req)
			if err == nil || result.Stdout != "begun" || result.Stderr != "warning" || time.Since(start) > 1500*time.Millisecond {
				t.Errorf("%s timeout = %#v, %v after %s", tc.name, result, err, time.Since(start))
			}
		})
	}
	snap.Config["timeout"] = "50ms"
	snap.Config["agent"] = map[string]any{"command": []any{"/bin/sh", "-c", `sleep 0.3; printf complete`}}
	req := request(snap, "worker", "")
	req.Timeout = 2 * time.Second
	result, err := runner.InvokeAgent(context.Background(), req)
	if err != nil || result.Stdout != "complete" {
		t.Errorf("positive request timeout did not override shorter configured timeout: %#v, %v", result, err)
	}
}

func Test_RUN_007_RUN_009_TimeoutAndCancellationKillOrdinaryDescendant(t *testing.T) {
	snap := agentFixture(t)
	script := agentScript(t, snap.Root, "descendant.sh", `(sleep 1; printf child > "$2") & printf ready > "$1"; printf begun; printf warning >&2; wait`)
	for _, mode := range []string{"timeout", "cancel"} {
		t.Run(mode, func(t *testing.T) {
			ready := filepath.Join(snap.Root, mode+"-ready")
			marker := filepath.Join(snap.Root, mode+"-child-finished")
			snap.Config["agent"] = map[string]any{"command": []any{script, ready, marker}}
			snap.Config["timeout"] = "500ms"
			if mode == "cancel" {
				snap.Config["timeout"] = "5s"
			}
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			type outcome struct {
				result runner.AgentResult
				err    error
			}
			done := make(chan outcome, 1)
			go func() {
				result, err := runner.InvokeAgent(ctx, request(snap, "worker", ""))
				done <- outcome{result, err}
			}()
			deadline := time.After(2 * time.Second)
			for {
				if _, err := os.Stat(ready); err == nil {
					break
				}
				select {
				case got := <-done:
					t.Fatalf("agent exited before descendant was ready: %#v, %v", got.result, got.err)
				case <-deadline:
					t.Fatal("agent never started descendant")
				case <-time.After(10 * time.Millisecond):
				}
			}
			if mode == "cancel" {
				cancel()
			}
			select {
			case got := <-done:
				if got.err == nil || got.result.Stdout != "begun" || got.result.Stderr != "warning" {
					t.Errorf("terminated agent = %#v, %v", got.result, got.err)
				}
			case <-time.After(2 * time.Second):
				t.Fatal("terminated agent did not stop promptly")
			}
			time.Sleep(1200 * time.Millisecond)
			if _, err := os.Stat(marker); !os.IsNotExist(err) {
				t.Errorf("descendant completed after %s: stat error %v", mode, err)
			}
		})
	}
}

func Test_RUN_008_UsageLimitOnlyForNonzeroMatchInLastTwentyLines(t *testing.T) {
	snap := agentFixture(t)
	snap.Config["limit_patterns"] = []any{"RATE LIMIT"}
	for _, tc := range []struct {
		name, command string
		want          bool
	}{
		{"late_stderr_match", `printf normal; printf 'rAtE LiMiT\n' >&2; exit 9`, true},
		{"successful_match", `printf 'rate limit\n'; exit 0`, false},
		{"outside_last_twenty", `printf 'rate limit\n'; i=0; while [ "$i" -lt 21 ]; do printf 'noise\n'; i=$((i+1)); done; exit 9`, false},
		{"stdout_before_stderr_window", `printf 'rate limit\n'; i=0; while [ "$i" -lt 21 ]; do printf 'noise\n' >&2; i=$((i+1)); done; exit 9`, false},
		{"literal_substring", `printf 'rating limit\n'; exit 9`, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			snap.Config["agent"] = map[string]any{"command": []any{"/bin/sh", "-c", tc.command}}
			result, err := runner.InvokeAgent(context.Background(), request(snap, "worker", ""))
			if result.UsageLimit != tc.want || (err != nil) != (result.ExitCode != 0) {
				t.Errorf("%s = %#v, %v; want UsageLimit=%v", tc.name, result, err, tc.want)
			}
		})
	}
}

func Test_RUN_009_LaunchFailureReturnsError(t *testing.T) {
	snap := agentFixture(t)
	missing := filepath.Join(snap.Root, "missing-agent")
	snap.Config["agent"] = map[string]any{"command": []any{missing}}
	result, err := runner.InvokeAgent(context.Background(), request(snap, "worker", ""))
	if err == nil || !strings.Contains(err.Error(), "missing-agent") || result.Stdout != "" || result.Stderr != "" {
		t.Errorf("launch failure = %#v, %v", result, err)
	}
}
