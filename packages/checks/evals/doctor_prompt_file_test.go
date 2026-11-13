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
)

const doctorProbePrompt = "Reply with exactly FIREWEED_DOCTOR_OK"

func doctorRecordedPath(t *testing.T, log, label string) string {
	t.Helper()
	return doctorRecordedPaths(t, log, label, 1)[0]
}

func doctorRecordedPaths(t *testing.T, log, label string, count int) []string {
	t.Helper()
	var paths []string
	for _, line := range strings.Split(log, "\n") {
		if strings.HasPrefix(line, label+":") {
			paths = append(paths, strings.TrimPrefix(line, label+":"))
		}
	}
	if len(paths) != count {
		t.Fatalf("%s paths = %q; want %d absolute paths", label, paths, count)
	}
	for _, path := range paths {
		if !filepath.IsAbs(path) {
			t.Fatalf("%s path = %q; want an absolute path", label, path)
		}
	}
	return paths
}

func doctorPromptFileRemoved(t *testing.T, path string) {
	t.Helper()
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		t.Errorf("prompt file %q remains after probe: %v", path, err)
	}
}

func Test_CHK_082_CHK_083_DoctorExpandsSelectedCommandsAndWrapperOnce(t *testing.T) {
	f := localDoctorFixture(t)
	f.editConfig(func(config map[string]any) {
		config["agent"].(map[string]any)["command"] = []any{"bash", ".fireweed/agent.sh", "pre{prompt}post", "file={prompt_file}", "common={git_common_dir}", "literal {unknown}"}
		// The shell command leaves placeholders unquoted. The worker and wrapper use YAML sequences.
		config["audit"].(map[string]any)["command"] = "bash .fireweed/audit.sh {prompt} {prompt_file} {git_common_dir}"
		config["agent"].(map[string]any)["wrap"] = []any{"bash", ".fireweed/prompt-wrap.sh", "{prompt_file}", "{prompt}", "{git_common_dir}"}
	})
	f.put(".fireweed/prompt-wrap.sh", fmt.Sprintf(`#!/bin/bash
printf 'wrap-file:%%s\nwrap-prompt:%%s\nwrap-common:%%s\n' "$1" "$2" "$3" >> %q
test -f "$1" && printf '%%s' "$2" | cmp -s - "$1" || exit 11
shift 3
exec "$@"
`, f.log))
	f.put(".fireweed/agent.sh", fmt.Sprintf(`#!/bin/bash
test "$#" -eq 4 || exit 12
printf 'worker-prompt:%%s\nworker-file:%%s\nworker-common:%%s\nworker-literal:%%s\n' "$1" "${2#file=}" "${3#common=}" "$4" >> %q
test "$1" = "pre%spost" || exit 12
test "$2" != "file=" || exit 13
test -f "${2#file=}" && printf '%%s' %q | cmp -s - "${2#file=}" || exit 14
printf 'FIREWEED_DOCTOR_OK\n'
`, f.log, doctorProbePrompt, doctorProbePrompt))
	f.put(".fireweed/audit.sh", fmt.Sprintf(`#!/bin/bash
test "$#" -eq 3 || exit 15
printf 'audit-prompt:%%s\naudit-file:%%s\naudit-common:%%s\n' "$1" "$2" "$3" >> %q
test -f "$2" && printf '%%s' "$1" | cmp -s - "$2" || exit 15
printf 'FIREWEED_DOCTOR_OK\n'
`, f.log))
	f.commit()
	// Relocate the Git directory so its absolute path contains a quote, a space,
	// and another placeholder. The expanded value must remain literal.
	commonDir := filepath.Join(t.TempDir(), "space 'quote' {prompt}", "repo.git")
	if err := os.MkdirAll(filepath.Dir(commonDir), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.Rename(filepath.Join(f.root, ".git"), commonDir); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(f.root, ".git"), []byte("gitdir: "+commonDir+"\n"), 0644); err != nil {
		t.Fatal(err)
	}
	t.Setenv("TMPDIR", filepath.Dir(commonDir))
	status, registrations := f.git("status", "--porcelain", "--untracked-files=all"), worktrees(f)
	report := checks.Doctor(context.Background(), f.request(false))
	if doctorErrors(report) || !hasDoctorCheck(report, "worker", true) || !hasDoctorCheck(report, "audit", true) {
		t.Errorf("prompt-file probes failed: %s", doctorText(report))
	}
	log := doctorLog(f)
	common, err := filepath.EvalSymlinks(commonDir)
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{
		"worker-prompt:pre" + doctorProbePrompt + "post", "audit-prompt:" + doctorProbePrompt,
		"worker-common:" + common, "audit-common:" + common,
		"worker-literal:literal {unknown}",
		"wrap-prompt:" + doctorProbePrompt, "wrap-common:" + common,
	} {
		if !strings.Contains(log, want+"\n") {
			t.Errorf("missing exact argument %q", want)
		}
	}
	workerFile := doctorRecordedPath(t, log, "worker-file")
	auditFile := doctorRecordedPath(t, log, "audit-file")
	if workerFile == auditFile {
		t.Errorf("worker and audit reused prompt file %q", workerFile)
	}
	for _, path := range []string{workerFile, auditFile} {
		if !strings.Contains(log, "wrap-file:"+path+"\n") {
			t.Errorf("wrapper did not receive selected probe file %q", path)
		}
		doctorPromptFileRemoved(t, path)
	}
	doctorProbeLog(t, f, report, "worker", "FIREWEED_DOCTOR_OK")
	doctorProbeLog(t, f, report, "audit", "FIREWEED_DOCTOR_OK")
	doctorCallerClean(t, f, status, registrations)
}

func Test_CHK_083_DoctorRemovesPromptFileOnProbeFailures(t *testing.T) {
	for _, mode := range []string{"launch", "nonzero", "timeout", "cancel"} {
		t.Run(mode, func(t *testing.T) {
			f := localDoctorFixture(t)
			ready := filepath.Join(t.TempDir(), "ready")
			f.editConfig(func(config map[string]any) {
				config["agent"].(map[string]any)["command"] = []any{"bash", ".fireweed/agent.sh", "{prompt_file}"}
				if mode == "launch" {
					config["agent"].(map[string]any)["command"] = []any{filepath.Join(f.root, "missing-agent"), "{prompt_file}"}
					config["agent"].(map[string]any)["wrap"] = []any{"bash", ".fireweed/prompt-wrap.sh", "{prompt_file}"}
				}
			})
			f.put(".fireweed/prompt-wrap.sh", fmt.Sprintf("#!/bin/bash\nprintf 'worker-file:%%s\\n' \"$1\" >> %q\ntest -f \"$1\" || exit 11\nshift\nexec \"$@\"\n", f.log))
			script := fmt.Sprintf("#!/bin/bash\nprintf 'worker-file:%%s\\nworker-content:%%s\\n' \"$1\" \"$(cat \"$1\")\" >> %q\ntest -f \"$1\" && printf '%%s' %q | cmp -s - \"$1\" || exit 18\nprintf ready > %q\n", f.log, doctorProbePrompt, ready)
			switch mode {
			case "nonzero":
				script += "printf 'failed worker stderr\\n' >&2\nexit 7\n"
			case "timeout", "cancel":
				script += "printf 'waiting worker stdout\\n'\nsleep 5\n"
			}
			f.put(".fireweed/agent.sh", script)
			f.commit()
			status, registrations := f.git("status", "--porcelain", "--untracked-files=all"), worktrees(f)
			req := f.request(false)
			if mode == "timeout" {
				req.Timeout = 2 * time.Second
			}
			ctx := context.Background()
			var cancel context.CancelFunc
			if mode == "cancel" {
				ctx, cancel = context.WithCancel(ctx)
				defer cancel()
			}
			var report checks.DoctorReport
			if mode == "cancel" {
				reports := make(chan checks.DoctorReport, 1)
				go func() { reports <- checks.Doctor(ctx, req) }()
				deadline := time.After(8 * time.Second)
				for {
					if _, err := os.Stat(ready); err == nil {
						break
					} else if !os.IsNotExist(err) {
						t.Fatal(err)
					}
					select {
					case report = <-reports:
						t.Fatalf("worker exited before cancellation: %s", doctorText(report))
					case <-deadline:
						t.Fatal("worker never signalled readiness")
					case <-time.After(10 * time.Millisecond):
					}
				}
				livePath := doctorRecordedPath(t, doctorLog(f), "worker-file")
				if data, err := os.ReadFile(livePath); err != nil || string(data) != doctorProbePrompt {
					t.Errorf("prompt file unavailable before cancellation: %q, %v", data, err)
				}
				cancel()
				select {
				case report = <-reports:
				case <-time.After(4 * time.Second):
					t.Fatal("cancelled Doctor did not return")
				}
			} else {
				report = checks.Doctor(ctx, req)
			}
			if !hasDoctorCheck(report, "worker", false) || !hasDoctorDiagnostic(report, "C15", false, "worker") {
				t.Errorf("%s probe outcome changed: %s", mode, doctorText(report))
			}
			log := doctorLog(f)
			count := 1
			if mode == "launch" {
				// The shared wrapper records both the failed worker launch and
				// the subsequent audit probe; both files must be cleaned up.
				count = 2
			}
			paths := doctorRecordedPaths(t, log, "worker-file", count)
			if mode != "launch" && !strings.Contains(log, "worker-content:"+doctorProbePrompt+"\n") {
				t.Errorf("%s worker could not read exact prompt: %q", mode, log)
			}
			for _, path := range paths {
				doctorPromptFileRemoved(t, path)
			}
			doctorProbeLog(t, f, report, "worker")
			doctorCallerClean(t, f, status, registrations)
		})
	}
}

func Test_CHK_083_DoctorDoesNotCreatePromptFileWithoutSelectedPlaceholder(t *testing.T) {
	f := localDoctorFixture(t)
	expected := filepath.Join(t.TempDir(), "expected-prompt")
	if err := os.WriteFile(expected, []byte(doctorProbePrompt), 0644); err != nil {
		t.Fatal(err)
	}
	script := fmt.Sprintf(`#!/bin/bash
while IFS= read -r candidate; do
  if test -f "$candidate" && cmp -s "$candidate" %q; then
    printf 'unexpected prompt file: %%s\n' "$candidate" >&2
    exit 19
  fi
done < <(find "$TMPDIR" "$PWD" -type f -print)
printf 'FIREWEED_DOCTOR_OK\n'
`, expected)
	f.put(".fireweed/agent.sh", script)
	f.put(".fireweed/audit.sh", script)
	f.commit()
	temp := t.TempDir()
	t.Setenv("TMPDIR", temp)
	status, registrations := f.git("status", "--porcelain", "--untracked-files=all"), worktrees(f)
	report := checks.Doctor(context.Background(), f.request(false))
	if doctorErrors(report) || !hasDoctorCheck(report, "worker", true) || !hasDoctorCheck(report, "audit", true) {
		t.Errorf("unneeded prompt file appeared: %s", doctorText(report))
	}
	doctorProbeLog(t, f, report, "worker", "FIREWEED_DOCTOR_OK")
	doctorProbeLog(t, f, report, "audit", "FIREWEED_DOCTOR_OK")
	doctorCallerClean(t, f, status, registrations)
}

func Test_CHK_082_CHK_083_DoctorWrapperOnlyFileWithAuditFallback(t *testing.T) {
	f := localDoctorFixture(t)
	f.editConfig(func(config map[string]any) {
		delete(config["audit"].(map[string]any), "command")
		config["agent"].(map[string]any)["command"] = []any{"bash", ".fireweed/agent.sh", "{prompt}"}
		config["agent"].(map[string]any)["wrap"] = []any{"bash", ".fireweed/prompt-wrap.sh", "{prompt_file}"}
	})
	f.put(".fireweed/prompt-wrap.sh", fmt.Sprintf(`#!/bin/bash
printf 'wrap-file:%%s\n' "$1" >> %q
test -f "$1" && printf '%%s' %q | cmp -s - "$1" || exit 11
shift
exec "$@"
`, f.log, doctorProbePrompt))
	f.put(".fireweed/agent.sh", fmt.Sprintf(`#!/bin/bash
printf 'fallback-prompt:%%s\n' "$1" >> %q
test "$1" = %q || exit 12
printf 'FIREWEED_DOCTOR_OK\n'
`, f.log, doctorProbePrompt))
	f.commit()
	status, registrations := f.git("status", "--porcelain", "--untracked-files=all"), worktrees(f)
	report := checks.Doctor(context.Background(), f.request(false))
	if doctorErrors(report) || !hasDoctorDiagnostic(report, "C15", true, "shared") || !hasDoctorCheck(report, "worker", true) || !hasDoctorCheck(report, "audit", true) {
		t.Errorf("wrapper-only fallback changed: %s", doctorText(report))
	}
	log := doctorLog(f)
	if strings.Count(log, "fallback-prompt:"+doctorProbePrompt+"\n") != 2 {
		t.Errorf("fallback did not receive fixed prompt twice")
	}
	var paths []string
	for _, line := range strings.Split(log, "\n") {
		if strings.HasPrefix(line, "wrap-file:") {
			paths = append(paths, strings.TrimPrefix(line, "wrap-file:"))
		}
	}
	if len(paths) != 2 || paths[0] == paths[1] {
		t.Errorf("wrapper prompt files = %q; want distinct worker/audit paths", paths)
	}
	for _, path := range paths {
		if !filepath.IsAbs(path) {
			t.Errorf("wrapper prompt path is relative: %q", path)
		}
		doctorPromptFileRemoved(t, path)
	}
	doctorProbeLog(t, f, report, "worker", "FIREWEED_DOCTOR_OK")
	doctorProbeLog(t, f, report, "audit", "FIREWEED_DOCTOR_OK")
	doctorCallerClean(t, f, status, registrations)
}
