package evals_test

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/app"
	"github.com/lurkshark/fireweed/packages/assets"
)

type informationCase struct {
	name string
	args []string
	want string
}

func informationCases(t *testing.T) []informationCase {
	t.Helper()
	explanation, err := assets.Explain("C01")
	if err != nil {
		t.Fatal(err)
	}
	playbook, err := assets.Playbook("evals")
	if err != nil {
		t.Fatal(err)
	}
	return []informationCase{
		{"help", []string{"help"}, assets.Help()},
		{"version", []string{"version"}, assets.Version()},
		{"explain", []string{"explain", "C01"}, explanation},
		{"playbook", []string{"playbook", "evals"}, playbook},
	}
}

func Test_APP_001_StandaloneBinary(t *testing.T) {
	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("locate eval source")
	}
	root := filepath.Clean(filepath.Join(filepath.Dir(thisFile), "..", "..", ".."))
	bin := filepath.Join(t.TempDir(), "fireweed-new")
	build := exec.Command("go", "build", "-o", bin, ".")
	build.Dir = root
	build.Env = append(os.Environ(), "GOCACHE=/tmp/fireweed-rebuild-cache")
	if output, err := build.CombinedOutput(); err != nil {
		t.Fatalf("build NEW binary: %v\n%s", err, output)
	}

	outside := t.TempDir()
	malformed := t.TempDir()
	if err := os.WriteFile(filepath.Join(malformed, "fireweed.yaml"), []byte("[malformed yaml\n"), 0600); err != nil {
		t.Fatal(err)
	}
	for _, location := range []struct {
		name string
		dir  string
	}{{"outside-project", outside}, {"invalid-config", malformed}} {
		for _, tc := range informationCases(t) {
			t.Run(location.name+"/"+tc.name, func(t *testing.T) {
				cmd := exec.Command(bin, tc.args...)
				cmd.Dir = location.dir
				var stdout, stderr bytes.Buffer
				cmd.Stdout, cmd.Stderr = &stdout, &stderr
				if err := cmd.Run(); err != nil {
					t.Errorf("%v: %v", tc.args, err)
				}
				assertInformationOutput(t, stdout.String(), stderr.String(), tc.want)
			})
		}
	}
}

func Test_APP_002_InformationJSON(t *testing.T) {
	for _, tc := range informationCases(t) {
		t.Run(tc.name, func(t *testing.T) {
			plainStatus, plain, plainErr := run(tc.args...)
			if plainStatus != 0 || plainErr != "" {
				t.Errorf("plain status=%d stderr=%q; want 0 and empty", plainStatus, plainErr)
			}
			assertInformationOutput(t, plain, plainErr, tc.want)
			args := append(append([]string{}, tc.args...), "--json")
			status, stdout, stderr := run(args...)
			if status != 0 || stderr != "" {
				t.Errorf("status=%d stderr=%q; want 0 and empty", status, stderr)
			}
			envelope := decodeEnvelope(t, stdout)
			var data string
			if err := json.Unmarshal(envelope["data"], &data); err != nil {
				t.Fatalf("data must be a string: %v", err)
			}
			if data != plain {
				t.Errorf("data=%q; want plain result %q", data, plain)
			}
			var diagnostics []json.RawMessage
			if err := json.Unmarshal(envelope["diagnostics"], &diagnostics); err != nil || diagnostics == nil || len(diagnostics) != 0 {
				t.Errorf("diagnostics=%s; want empty array", envelope["diagnostics"])
			}
		})
	}
}

func Test_APP_003_InvalidRequests(t *testing.T) {
	for _, tc := range []struct {
		name string
		args []string
		bad  string
	}{
		{"unknown-command", []string{"not-a-command"}, "not-a-command"},
		{"unknown-option", []string{"help", "--bogus"}, "--bogus"},
		{"missing-argument", []string{"explain"}, "explain"},
		{"extra-argument", []string{"version", "surplus"}, "version"},
		{"unknown-explanation", []string{"explain", "NO-SUCH-CODE"}, "NO-SUCH-CODE"},
		{"unknown-playbook", []string{"playbook", "no-such-playbook"}, "no-such-playbook"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			for _, format := range []struct {
				name string
				json bool
			}{{"plain", false}, {"JSON", true}} {
				args := append([]string(nil), tc.args...)
				if format.json {
					args = append(args, "--json")
				}
				status, stdout, stderr := run(args...)
				if status != 2 {
					t.Errorf("%s status=%d; want 2", format.name, status)
				}
				diagnostic := stdout + stderr
				if format.json {
					diagnostic = assertFailedJSON(t, stdout)
				}
				if !strings.Contains(diagnostic, tc.bad) || !corrective(diagnostic) {
					t.Errorf("%s diagnostic %q must identify %q and suggest a correction", format.name, diagnostic, tc.bad)
				}
			}
		})
	}
}

func Test_APP_003_UnimplementedCommands(t *testing.T) {
	for _, command := range []string{"init", "status"} {
		t.Run(command, func(t *testing.T) {
			if !strings.Contains(assets.Help(), command) {
				t.Fatalf("%q is not listed by help", command)
			}
			status, stdout, stderr := run(command)
			if status == 0 || !regexp.MustCompile(`(?i)unimplemented|not implemented|under construction`).MatchString(stdout+stderr) {
				t.Errorf("status=%d diagnostic=%q; want nonzero and explicit unimplemented notice", status, stdout+stderr)
			}
		})
	}
}

func Test_APP_003_FailedJSON(t *testing.T) {
	status, stdout, _ := run("init", "--json")
	if status == 0 {
		t.Error("failed request returned exit 0")
	}
	assertFailedJSON(t, stdout)
}

func assertFailedJSON(t *testing.T, stdout string) string {
	t.Helper()
	envelope := decodeEnvelope(t, stdout)
	if string(envelope["data"]) != "null" {
		t.Errorf("data=%s; want null", envelope["data"])
	}
	var diagnostics []struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}
	if err := json.Unmarshal(envelope["diagnostics"], &diagnostics); err != nil || len(diagnostics) == 0 {
		t.Fatalf("diagnostics=%s; want nonempty array: %v", envelope["diagnostics"], err)
	}
	if strings.TrimSpace(diagnostics[0].Code) == "" || strings.TrimSpace(diagnostics[0].Message) == "" {
		t.Errorf("diagnostic needs code and message: %+v", diagnostics[0])
	}
	return diagnostics[0].Message
}

// The app-level convention probe uses one PATH recorder for external commands.
func Test_APP_001_ReadCommandsDoNotInvokeGitHooksOrAgents(t *testing.T) {
	binDir := t.TempDir()
	logPath := filepath.Join(t.TempDir(), "invocations")
	recorder := []byte("#!/bin/sh\nprintf '%s\\n' \"$0 $*\" >> \"$FIREWEED_RECORDER\"\nexit 77\n")
	for _, name := range []string{"git", "hook", "codex"} {
		if err := os.WriteFile(filepath.Join(binDir, name), recorder, 0700); err != nil {
			t.Fatal(err)
		}
	}
	t.Setenv("FIREWEED_RECORDER", logPath)
	t.Setenv("PATH", binDir+string(os.PathListSeparator)+os.Getenv("PATH"))
	for _, tc := range informationCases(t) {
		status, stdout, stderr := run(tc.args...)
		if status != 0 {
			t.Errorf("%s status=%d stderr=%q", tc.name, status, stderr)
		}
		assertInformationOutput(t, stdout, stderr, tc.want)
	}
	if calls, err := os.ReadFile(logPath); err == nil {
		t.Errorf("read commands invoked external tools: %s", calls)
	} else if !os.IsNotExist(err) {
		t.Fatal(err)
	}
}

func run(args ...string) (int, string, string) {
	var stdout, stderr bytes.Buffer
	status := app.Run(args, &stdout, &stderr)
	return status, stdout.String(), stderr.String()
}

func assertInformationOutput(t *testing.T, stdout, stderr, want string) {
	t.Helper()
	if stderr != "" {
		t.Errorf("stderr=%q; want empty", stderr)
	}
	if strings.TrimSuffix(stdout, "\n") != strings.TrimSuffix(want, "\n") {
		t.Errorf("stdout=%q; want assets text %q", stdout, want)
	}
}

func decodeEnvelope(t *testing.T, output string) map[string]json.RawMessage {
	t.Helper()
	dec := json.NewDecoder(strings.NewReader(output))
	var envelope map[string]json.RawMessage
	if err := dec.Decode(&envelope); err != nil || envelope == nil {
		t.Fatalf("stdout=%q; want JSON object: %v", output, err)
	}
	var extra any
	if err := dec.Decode(&extra); err != io.EOF {
		t.Fatalf("stdout=%q; want exactly one JSON object: %v", output, err)
	}
	var schema int
	if err := json.Unmarshal(envelope["schema"], &schema); err != nil || schema != 1 {
		t.Errorf("schema=%s; want 1", envelope["schema"])
	}
	return envelope
}

func corrective(message string) bool {
	return regexp.MustCompile(`(?i)\b(help|usage|valid|available|expected|required|requires|try|use|choose|provide|specify)\b`).MatchString(message)
}
