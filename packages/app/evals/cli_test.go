package evals

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"
	"time"

	app "github.com/lurkshark/fireweed/packages/app"
	a "github.com/lurkshark/fireweed/packages/assets"
)

var subject string

func TestMain(m *testing.M) {
	directory, err := os.MkdirTemp("", "fireweed-product-evals-")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
	subject = filepath.Join(directory, "fireweed-product")
	if runtime.GOOS == "windows" {
		subject += ".exe"
	}
	root, err := filepath.Abs("../../..")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.RemoveAll(directory)
		os.Exit(2)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	command := exec.CommandContext(ctx, "go", "build", "-o", subject, ".")
	command.Dir = root
	command.Env = append(os.Environ(), "GOPROXY=off", "GOTOOLCHAIN=local")
	output, err := command.CombinedOutput()
	cancel()
	if err != nil {
		fmt.Fprintf(os.Stderr, "build NEW product subject: %v %s\n", err, output)
		os.RemoveAll(directory)
		os.Exit(2)
	}
	code := m.Run()
	os.RemoveAll(directory)
	os.Exit(code)
}

type result struct {
	code     int
	out, err string
}

func invoke(t *testing.T, dir string, args ...string) result {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	command := exec.CommandContext(ctx, subject, args...)
	command.Dir = dir
	command.Stdin = strings.NewReader("")
	var out, errOut bytes.Buffer
	command.Stdout = &out
	command.Stderr = &errOut
	err := command.Run()
	if ctx.Err() != nil {
		t.Fatalf("command hung: %v", args)
	}
	code := 0
	if err != nil {
		var exit *exec.ExitError
		if !errors.As(err, &exit) {
			t.Fatalf("execute NEW binary: %v", err)
		}
		code = exit.ExitCode()
	}
	return result{code, out.String(), errOut.String()}
}

type diagnostic struct{ Code, Severity, Message string }
type envelope struct {
	SchemaVersion int             `json:"schema_version"`
	Command       string          `json:"command"`
	Data          json.RawMessage `json:"data"`
	Diagnostics   []diagnostic    `json:"diagnostics"`
}

func decode(t *testing.T, got result, command string) envelope {
	t.Helper()
	var value envelope
	decoder := json.NewDecoder(strings.NewReader(got.out))
	if err := decoder.Decode(&value); err != nil {
		t.Fatalf("not JSON: %q: %v", got.out, err)
	}
	var extra any
	if err := decoder.Decode(&extra); err != io.EOF {
		t.Errorf("more than one JSON value: %v", err)
	}
	if value.SchemaVersion != 1 || value.Command != command || value.Data == nil || value.Diagnostics == nil {
		t.Errorf("bad envelope %+v", value)
	}
	if got.err != "" {
		t.Errorf("JSON command wrote stderr: %q", got.err)
	}
	return value
}
func successJSON(t *testing.T, got result, command string) envelope {
	t.Helper()
	value := decode(t, got, command)
	if got.code != 0 || len(value.Diagnostics) != 0 {
		t.Errorf("success code/diagnostics %+v %+v", got, value)
	}
	return value
}
func assertUsage(t *testing.T, got result) {
	t.Helper()
	if got.code != 2 || got.out != "" || !strings.Contains(got.err, "U01") || !strings.Contains(got.err, "help") {
		t.Errorf("usage result %+v", got)
	}
}
func Test_APP001HelpAndVersion(t *testing.T) {
	dir := t.TempDir()
	inventory := []string{"help", "version", "init", "doctor", "new-module", "new-change", "stamp", "list", "status", "modules", "show", "context", "explain", "playbook", "check", "coverage", "check-commits", "ci", "pr-body", "start", "verify", "done", "next", "run", "release", "commits", "drill", "gc", "migrate", "agents-md"}
	for _, args := range [][]string{nil, {"help"}, {"-h"}, {"--help"}} {
		got := invoke(t, dir, args...)
		if got.code != 0 || got.err != "" {
			t.Fatalf("help %v: %+v", args, got)
		}
		for _, name := range inventory {
			if !strings.Contains(got.out, name) {
				t.Errorf("help missing %s", name)
			}
		}
		for _, term := range []string{"--json", "0", "success", "1", "failed", "2", "usage", "3", "limit"} {
			if !strings.Contains(strings.ToLower(got.out), term) {
				t.Errorf("help missing exit/options %q", term)
			}
		}
	}
	for _, args := range [][]string{{"version"}, {"-v"}, {"--version"}} {
		got := invoke(t, dir, args...)
		if got.code != 0 || got.err != "" || got.out != a.Info().Version+"\n" {
			t.Errorf("version %v: %+v", args, got)
		}
	}
	files, err := os.ReadDir(dir)
	if err != nil || len(files) != 0 {
		t.Errorf("inspection created files %v %v", files, err)
	}
}
func Test_APP002EmbeddedReads(t *testing.T) {
	dir := t.TempDir()
	for _, code := range a.Codes() {
		got := invoke(t, dir, "explain", code)
		var data map[string]string
		if err := json.Unmarshal([]byte(got.out), &data); err != nil {
			t.Fatalf("plain explain: %v %+v", err, got)
		}
		text, err := a.Explain(code)
		if err != nil {
			t.Fatal(err)
		}
		if got.code != 0 || got.err != "" || data["code"] != code || data["explanation"] != text {
			t.Errorf("explain %s: %+v %+v", code, got, data)
		}
	}
	for _, name := range a.Playbooks() {
		asset, err := a.Playbook(name)
		if err != nil {
			t.Fatal(err)
		}
		got := invoke(t, dir, "playbook", name)
		if got.code != 0 || got.err != "" || got.out != asset.Text {
			t.Errorf("playbook %s: %+v", name, got)
		}
		value := successJSON(t, invoke(t, dir, "playbook", name, "--json"), "playbook")
		var data struct {
			Name, Text, Hash string
			EmbeddedHash     string `json:"embedded_hash"`
			Overridden       bool
		}
		if err := json.Unmarshal(value.Data, &data); err != nil {
			t.Fatal(err)
		}
		sum := fmt.Sprintf("%x", sha256.Sum256([]byte(data.Text)))
		if data.Name != name || data.Text != asset.Text || data.Hash != sum || data.EmbeddedHash != asset.Hash || data.Overridden {
			t.Errorf("playbook data %+v", data)
		}
	}
	for _, tc := range []struct {
		command string
		choices []string
	}{{"explain", a.Codes()}, {"playbook", a.Playbooks()}} {
		got := invoke(t, dir, tc.command, "unknown")
		assertUsage(t, got)
		for _, choice := range tc.choices {
			if !strings.Contains(got.err, choice) {
				t.Errorf("unknown %s omits valid %s", tc.command, choice)
			}
		}
		assertUsage(t, invoke(t, dir, tc.command))
		assertUsage(t, invoke(t, dir, tc.command, tc.choices[0], "extra"))
	}
}
func Test_APP003JSONEnvelope(t *testing.T) {
	dir := t.TempDir()
	for _, args := range [][]string{{"help"}, {"version"}, {"-h"}, {"--help"}, {"-v"}, {"--version"}, {"explain", "H04"}, {"playbook", "audit"}} {
		canonical := args[0]
		if canonical == "-h" || canonical == "--help" {
			canonical = "help"
		}
		if canonical == "-v" || canonical == "--version" {
			canonical = "version"
		}
		before := append([]string{"--json"}, args...)
		after := append(append([]string{}, args...), "--json")
		first := successJSON(t, invoke(t, dir, before...), canonical)
		second := successJSON(t, invoke(t, dir, after...), canonical)
		if !reflect.DeepEqual(first, second) {
			t.Errorf("flag position changed envelope %v: %+v %+v", args, first, second)
		}
		if canonical == "version" {
			var version string
			if err := json.Unmarshal(first.Data, &version); err != nil || version != a.Info().Version {
				t.Errorf("version JSON %q %v", version, err)
			}
		}
	}
	for _, args := range [][]string{{"bogus"}, {"--bogus"}, {"version", "extra"}, {"explain"}, {"playbook", "unknown"}, {"version", "--json=true"}} {
		args = append(args, "--json")
		got := invoke(t, dir, args...)
		command := "help"
		if !strings.HasPrefix(args[0], "--") {
			command = args[0]
		}
		value := decode(t, got, command)
		if got.code != 2 || string(value.Data) != "null" || len(value.Diagnostics) != 1 || value.Diagnostics[0].Code != "U01" || value.Diagnostics[0].Severity != "error" || !strings.Contains(value.Diagnostics[0].Message, "help") {
			t.Errorf("JSON usage %+v %+v", got, value)
		}
	}
}

type failedWriter struct{}

func (failedWriter) Write([]byte) (int, error) { return 0, errors.New("controlled write failure") }
func Test_APP003WriteFailures(t *testing.T) {
	for _, args := range [][]string{{"help"}, {"version", "--json"}, {"explain", "H01"}, {"playbook", "audit"}, {"unknown", "--json"}} {
		if code := app.Run(args, failedWriter{}, io.Discard); code != 2 {
			t.Errorf("write failure %v exit %d", args, code)
		}
	}
	if code := app.Run([]string{"unknown"}, io.Discard, failedWriter{}); code != 2 {
		t.Errorf("stderr failure exit %d", code)
	}
}
func Test_APP004ArgumentRules(t *testing.T) {
	dir := t.TempDir()
	for _, args := range [][]string{{"bogus"}, {"bogus", "--help"}, {"--bogus"}, {"help", "extra"}, {"version", "extra"}, {"version", "--json=true"}, {"version", "--help=1"}, {"explain", "--unknown", "--help"}, {"--", "--json"}, {"explain", "--", "--json"}} {
		assertUsage(t, invoke(t, dir, args...))
	}
	for _, args := range [][]string{{"explain", "--help"}, {"playbook", "-h"}} {
		got := invoke(t, dir, args...)
		if got.code != 0 || got.err != "" || !strings.Contains(got.out, "--json") {
			t.Errorf("global help %+v", got)
		}
	}
	successJSON(t, invoke(t, dir, "playbook", "--help", "--json"), "help")
	successJSON(t, invoke(t, dir, "--json", "version", "--json"), "version")
	// End-of-options keeps a known code usable as an operand.
	got := invoke(t, dir, "explain", "--", "H01")
	if got.code != 0 {
		t.Errorf("end-of-options known operand %+v", got)
	}
}
func gitFixtureCommand(t *testing.T, root string, args ...string) string {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, "git", args...)
	cmd.Dir = root
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("fixture git %v: %v %s", args, err, out)
	}
	return string(out)
}
func Test_APP004NoAgentOrMutation(t *testing.T) {
	dir := t.TempDir()
	gitFixtureCommand(t, dir, "init", "--initial-branch=main")
	if err := os.WriteFile(filepath.Join(dir, "README"), []byte("fixture\n"), 0644); err != nil {
		t.Fatal(err)
	}
	gitFixtureCommand(t, dir, "add", ".")
	gitFixtureCommand(t, dir, "-c", "user.name=Fixture", "-c", "user.email=fixture@example.invalid", "commit", "-m", "fixture")
	refs := gitFixtureCommand(t, dir, "show-ref")
	files, err := os.ReadDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	if runtime.GOOS != "windows" {
		traps := t.TempDir()
		marker := filepath.Join(traps, "invoked")
		for _, name := range []string{"codex", "claude", "fireweed", "fireweed-bootstrap", "git"} {
			script := "#!/bin/sh\n: > '" + marker + "'\nexit 99\n"
			if err := os.WriteFile(filepath.Join(traps, name), []byte(script), 0755); err != nil {
				t.Fatal(err)
			}
		}
		previous := os.Getenv("PATH")
		t.Setenv("PATH", traps)
		t.Setenv("FIREWEED_BIN", filepath.Join(traps, "fireweed-bootstrap"))
		for _, args := range [][]string{{"help"}, {"version"}, {"explain", "H01"}, {"playbook", "evals"}} {
			if got := invoke(t, dir, args...); got.code != 0 {
				t.Errorf("standalone command %v %+v", args, got)
			}
		}
		if _, err := os.Stat(marker); !os.IsNotExist(err) {
			t.Errorf("inspection invoked external agent/controller/Git: %v", err)
		}
		t.Setenv("PATH", previous)
	}
	if gitFixtureCommand(t, dir, "show-ref") != refs || gitFixtureCommand(t, dir, "status", "--porcelain") != "" {
		t.Error("inspection mutated refs or checkout")
	}
	after, err := os.ReadDir(dir)
	if err != nil || len(after) != len(files) {
		t.Errorf("inspection created directory entries %v %v", after, err)
	}
}
