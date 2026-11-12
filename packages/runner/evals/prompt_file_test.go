package evals_test

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/runner"
)

func promptFileTempDir(t *testing.T) string {
	t.Helper()
	root := filepath.Join(t.TempDir(), "temporary files with spaces")
	if err := os.Mkdir(root, 0755); err != nil {
		t.Fatal(err)
	}
	t.Setenv("TMPDIR", root)
	return root
}

func promptFileObservation(t *testing.T, root, name string) string {
	t.Helper()
	data, err := os.ReadFile(filepath.Join(root, name))
	if err != nil {
		t.Fatalf("read %s: %v", name, err)
	}
	return string(data)
}

func promptFileRemoved(t *testing.T, path, workspaceRoot string) {
	t.Helper()
	if !filepath.IsAbs(path) || path == workspaceRoot || strings.HasPrefix(path, workspaceRoot+string(os.PathSeparator)) {
		t.Errorf("prompt file path %q is not absolute outside input workspace %q", path, workspaceRoot)
	}
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		t.Errorf("prompt file %q remains after invocation: %v", path, err)
	}
}

func Test_RUN_003_PromptFileExpansionInShellAndArgvPreservesValuesOnce(t *testing.T) {
	snap := agentFixture(t)
	common, err := gitstore.Inspect(snap.Root)
	if err != nil {
		t.Fatal(err)
	}
	observations := t.TempDir()
	script := agentScript(t, snap.Root, "prompt-file-args.sh", `printf '%s' "$2" > "$1/prompt"
printf '%s' "$3" > "$1/path"
printf '%s' "$4" > "$1/common"
cat "$3" > "$1/content" || exit 12
printf 'agent-ok'`)
	prompt := "a'b ; $(printf injected) {git_common_dir} {prompt_file}\nsecond line Ω"
	for _, tc := range []struct {
		name, role string
		command    any
	}{
		{"worker shell", "worker", shellQuote(script) + " " + shellQuote(observations) + " {prompt} {prompt_file} {git_common_dir}"},
		{"worker argv", "worker", []any{script, observations, "{prompt}", "{prompt_file}", "{git_common_dir}"}},
		{"audit shell", "audit", shellQuote(script) + " " + shellQuote(observations) + " {prompt} {prompt_file} {git_common_dir}"},
		{"audit argv", "audit", []any{script, observations, "{prompt}", "{prompt_file}", "{git_common_dir}"}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			promptFileTempDir(t)
			snap.Config["agent"] = map[string]any{"command": []any{"/bin/sh", "-c", "printf unselected"}}
			delete(snap.Config, "audit")
			if tc.role == "worker" {
				snap.Config["agent"] = map[string]any{"command": tc.command}
			} else {
				snap.Config["audit"] = map[string]any{"command": tc.command}
			}
			result, err := runner.InvokeAgent(context.Background(), request(snap, tc.role, prompt))
			if err != nil || result.Stdout != "agent-ok" || result.ExitCode != 0 {
				t.Errorf("InvokeAgent = %#v, %v", result, err)
			}
			if got := promptFileObservation(t, observations, "prompt"); got != prompt {
				t.Errorf("expanded prompt = %q, want %q", got, prompt)
			}
			if got := promptFileObservation(t, observations, "content"); got != prompt {
				t.Errorf("prompt file content = %q, want %q", got, prompt)
			}
			if got := promptFileObservation(t, observations, "common"); got != common.CommonDir {
				t.Errorf("expanded git common dir = %q, want %q", got, common.CommonDir)
			}
			promptFileRemoved(t, promptFileObservation(t, observations, "path"), snap.Root)
		})
	}
}

func Test_RUN_003_RUN_067_WrapperUsesOnePromptFileForBothRolesAndFallback(t *testing.T) {
	snap := agentFixture(t)
	common, err := gitstore.Inspect(snap.Root)
	if err != nil {
		t.Fatal(err)
	}
	observations := t.TempDir()
	worker := agentScript(t, snap.Root, "prompt-file-worker.sh", `printf '%s' "$2" > "$1/command-path"
cat "$2" > "$1/command-content" || exit 12
printf worker`)
	auditor := agentScript(t, snap.Root, "prompt-file-auditor.sh", `printf '%s' "$2" > "$1/command-path"
cat "$2" > "$1/command-content" || exit 12
printf audit`)
	wrapper := agentScript(t, snap.Root, "prompt-file-wrapper.sh", `printf '%s' "$2" > "$1/prompt"
printf '%s' "$3" > "$1/path"
printf '%s' "$4" > "$1/common"
printf '%s' "$5" > "$1/second-path"
cat "$3" > "$1/content" || exit 12
shift 5
exec "$@"`)
	snap.Config["agent"] = map[string]any{
		"command": []any{worker, observations, "{prompt_file}"},
		"wrap":    []any{wrapper, observations, "{prompt}", "{prompt_file}", "{git_common_dir}", "{prompt_file}"},
	}
	snap.Config["audit"] = map[string]any{"command": []any{auditor, observations, "{prompt_file}"}}
	prompt := "wrapper 'literal' {prompt_file} {git_common_dir}\nΩ"
	for _, tc := range []struct{ name, role, want string }{
		{"worker", "worker", "worker"},
		{"audit", "audit", "audit"},
		{"audit fallback", "audit", "worker"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			promptFileTempDir(t)
			if tc.name == "audit fallback" {
				delete(snap.Config, "audit")
			}
			result, err := runner.InvokeAgent(context.Background(), request(snap, tc.role, prompt))
			if err != nil || result.Stdout != tc.want || result.ExitCode != 0 {
				t.Errorf("wrapped %s = %#v, %v", tc.role, result, err)
			}
			for _, name := range []string{"prompt", "content", "command-content"} {
				if got := promptFileObservation(t, observations, name); got != prompt {
					t.Errorf("%s = %q, want %q", name, got, prompt)
				}
			}
			if got := promptFileObservation(t, observations, "common"); got != common.CommonDir {
				t.Errorf("wrapper common dir = %q, want %q", got, common.CommonDir)
			}
			path := promptFileObservation(t, observations, "path")
			for _, name := range []string{"second-path", "command-path"} {
				if duplicate := promptFileObservation(t, observations, name); duplicate != path {
					t.Errorf("%s received separate prompt file %q, want %q", name, duplicate, path)
				}
			}
			promptFileRemoved(t, path, snap.Root)
		})
	}
}

func Test_RUN_067_WrapperOnlyPlaceholderCreatesPromptFile(t *testing.T) {
	snap := agentFixture(t)
	observations := t.TempDir()
	promptFileTempDir(t)
	wrapper := agentScript(t, snap.Root, "wrapper-only.sh", `printf '%s' "$1" > "$2/path"
cat "$1" > "$2/content" || exit 12
printf '%s' "$3" > "$2/second-path"
shift 3
exec "$@"`)
	snap.Config["agent"] = map[string]any{
		"command": []any{"/bin/sh", "-c", "printf wrapper-only"},
		"wrap":    []any{wrapper, "{prompt_file}", observations, "{prompt_file}"},
	}
	prompt := "wrapper-only 'literal' {prompt_file}\nΩ"
	result, err := runner.InvokeAgent(context.Background(), request(snap, "worker", prompt))
	if err != nil || result.Stdout != "wrapper-only" || result.ExitCode != 0 {
		t.Errorf("wrapper-only invocation = %#v, %v", result, err)
	}
	path := promptFileObservation(t, observations, "path")
	if got := promptFileObservation(t, observations, "content"); got != prompt {
		t.Errorf("wrapper-only prompt file content = %q, want %q", got, prompt)
	}
	if second := promptFileObservation(t, observations, "second-path"); second != path {
		t.Errorf("wrapper-only placeholders used different files: %q, %q", path, second)
	}
	promptFileRemoved(t, path, snap.Root)
}

func Test_RUN_067_PromptFileLifetimeOnExitTimeoutCancellationAndLaunchFailure(t *testing.T) {
	for _, tc := range []struct{ name string }{
		{"success"}, {"nonzero"}, {"timeout"}, {"cancellation"}, {"launch failure"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			snap := agentFixture(t)
			observations := t.TempDir()
			tempRoot := promptFileTempDir(t)
			prompt := "exact prompt\nwith trailing newline Ω\n"
			if tc.name == "launch failure" {
				snap.Config["agent"] = map[string]any{"command": []any{filepath.Join(snap.Root, "missing-agent"), "{prompt_file}"}}
			} else {
				script := agentScript(t, snap.Root, "prompt-file-lifetime.sh", `printf '%s' "$1" > "$2/path"
test -r "$1" || exit 12
prompt_bytes=
IFS= read -r -d '' prompt_bytes < "$1" || test -n "$prompt_bytes" || exit 12
printf '%s' "$prompt_bytes" > "$2/content" || exit 12
printf ready > "$2/ready"
case "$3" in
  nonzero) printf partial; printf warning >&2; exit 9;;
  timeout|cancellation) printf partial; printf warning >&2; sleep 5;;
esac
printf success`)
				snap.Config["agent"] = map[string]any{"command": []any{"/bin/bash", script, "{prompt_file}", observations, tc.name}}
			}
			req := request(snap, "worker", prompt)
			req.Timeout = 2 * time.Second
			ctx := context.Background()
			if tc.name == "timeout" {
				req.Timeout = 150 * time.Millisecond
			}
			if tc.name == "cancellation" {
				var cancel context.CancelFunc
				ctx, cancel = context.WithCancel(ctx)
				defer cancel()
				go func() {
					ready := filepath.Join(observations, "ready")
					for i := 0; i < 200; i++ {
						if _, err := os.Stat(ready); err == nil {
							cancel()
							return
						}
						time.Sleep(10 * time.Millisecond)
					}
					cancel()
				}()
			}
			result, err := runner.InvokeAgent(ctx, req)
			if tc.name == "success" {
				if err != nil || result.Stdout != "success" || result.ExitCode != 0 {
					t.Errorf("successful invocation = %#v, %v", result, err)
				}
			} else if err == nil {
				t.Errorf("%s returned no error: %#v", tc.name, result)
			}
			if tc.name == "cancellation" && !errors.Is(err, context.Canceled) {
				t.Errorf("cancellation error = %v", err)
			}
			if (tc.name == "nonzero" && result.ExitCode != 9) || (tc.name == "timeout" && !errors.Is(err, context.DeadlineExceeded)) {
				t.Errorf("%s result = %#v, %v", tc.name, result, err)
			}
			if tc.name != "launch failure" {
				if got := promptFileObservation(t, observations, "content"); got != prompt {
					t.Errorf("file during %s = %q, want %q", tc.name, got, prompt)
				}
				promptFileRemoved(t, promptFileObservation(t, observations, "path"), snap.Root)
			}
			entries, readErr := os.ReadDir(tempRoot)
			if readErr != nil {
				t.Errorf("read temporary directory after %s: %v", tc.name, readErr)
			}
			for _, entry := range entries {
				// macOS command-line tooling can populate this cache in TMPDIR.
				if !strings.Contains(entry.Name(), "xcrun_db") {
					t.Errorf("temporary file %q remains after %s", entry.Name(), tc.name)
				}
			}
		})
	}
}

func Test_RUN_067_UnselectedTemplateDoesNotRequirePromptFileIO(t *testing.T) {
	snap := agentFixture(t)
	snap.Config["agent"] = map[string]any{"command": []any{"/bin/sh", "-c", "printf worker"}}
	snap.Config["audit"] = map[string]any{"command": []any{"/bin/sh", "-c", "printf '{prompt_file}'"}}
	t.Setenv("TMPDIR", filepath.Join(t.TempDir(), "missing-temp-directory"))
	result, err := runner.InvokeAgent(context.Background(), request(snap, "worker", "unused prompt"))
	if err != nil || result.Stdout != "worker" || result.ExitCode != 0 {
		t.Errorf("worker selected without prompt file = %#v, %v", result, err)
	}
	snap.Config["agent"] = map[string]any{"command": []any{"/bin/sh", "-c", "printf '{prompt_file}'"}}
	snap.Config["audit"] = map[string]any{"command": []any{"/bin/sh", "-c", "printf audit"}}
	result, err = runner.InvokeAgent(context.Background(), request(snap, "audit", "unused prompt"))
	if err != nil || result.Stdout != "audit" || result.ExitCode != 0 {
		t.Errorf("audit selected without prompt file = %#v, %v", result, err)
	}
	snap.Config["agent"] = map[string]any{"command": []any{"/bin/sh", "-c", "printf worker"}}
	delete(snap.Config, "audit")
	result, err = runner.InvokeAgent(context.Background(), request(snap, "audit", "unused prompt"))
	if err != nil || result.Stdout != "worker" || result.ExitCode != 0 {
		t.Errorf("audit fallback without prompt file = %#v, %v", result, err)
	}
}
