package evals_test

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"gopkg.in/yaml.v3"
)

const doctorCLIManagedBlock = "<!-- fireweed:start -->\nThis project uses fireweed. Specs and change files are the source of truth.\nRun `fireweed playbook author-change` before authoring changes and\n`fireweed context <change>` before each step.\n<!-- fireweed:end -->"

func doctorCLIGit(t *testing.T, root string, args ...string) {
	t.Helper()
	cmd := exec.Command("git", args...)
	cmd.Dir = root
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("git %v: %v: %s", args, err, out)
	}
}

func doctorCLIConfig(t *testing.T, root string, edit func(map[string]any)) {
	t.Helper()
	var config map[string]any
	if err := yaml.Unmarshal([]byte(readInitFile(t, root, "fireweed.yaml")), &config); err != nil {
		t.Fatal(err)
	}
	edit(config)
	data, err := yaml.Marshal(config)
	if err != nil {
		t.Fatal(err)
	}
	appPut(t, root, "fireweed.yaml", string(data))
}

func doctorCLIFile(t *testing.T, path string) string {
	t.Helper()
	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return ""
	}
	if err != nil {
		t.Fatal(err)
	}
	return string(data)
}

func doctorCLIFixture(t *testing.T, bin string) (string, string, string) {
	t.Helper()
	root := t.TempDir()
	log := filepath.Join(t.TempDir(), "probes.log")
	appPut(t, root, "go.mod", "module example.test/doctorcli\n\ngo 1.23\n")
	status, out, errOut := runAppBinary(t, bin, root, "init", "--recipe", "go", "--agent", "sh .fireweed/agent.sh {prompt}", "--audit-agent", "sh .fireweed/audit.sh {prompt}", "--no-app")
	if status != 0 || errOut != "" {
		t.Fatalf("initialize Go fixture: status=%d stdout=%q stderr=%q", status, out, errOut)
	}
	appPut(t, root, "packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\n---\n# Alpha\n\n## Specs\n\n### ALP-001\nThe fixture MUST pass.\n")
	appPut(t, root, "packages/alpha/api.go", "package alpha\n")
	appPut(t, root, ".fireweed/setup.sh", fmt.Sprintf("#!/bin/sh\nprintf 'setup:%%s:%%s\\n' \"$PWD\" \"$FIREWEED_BIN\" >> %q\n\"$FIREWEED_BIN\" modules --json >/dev/null\n", log))
	appPut(t, root, ".fireweed/agent.sh", fmt.Sprintf("#!/bin/sh\nprintf 'worker:%%s\\n' \"$PWD\" >> %q\nprintf 'FIREWEED_DOCTOR_OK\\n'\n", log))
	appPut(t, root, ".fireweed/audit.sh", fmt.Sprintf("#!/bin/sh\nprintf 'audit:%%s\\n' \"$PWD\" >> %q\nprintf 'FIREWEED_DOCTOR_OK\\n'\n", log))
	doctorCLIConfig(t, root, func(config map[string]any) {
		config["modules"] = []string{"packages/alpha"}
		config["hooks"].(map[string]any)["setup"] = "sh .fireweed/setup.sh"
	})
	doctorCLIGit(t, root, "init", "-b", "main")
	doctorCLIGit(t, root, "config", "user.name", "Doctor CLI Fixture")
	doctorCLIGit(t, root, "config", "user.email", "doctor-cli@example.test")
	doctorCLIGit(t, root, "add", ".")
	doctorCLIGit(t, root, "commit", "-m", "fixture")
	nested := appNestedDir(t, root)
	return root, nested, log
}

func doctorCLIChecks(t *testing.T, output string) ([]map[string]json.RawMessage, []map[string]json.RawMessage) {
	t.Helper()
	envelope := decodeEnvelope(t, output)
	var data struct {
		Checks []map[string]json.RawMessage `json:"checks"`
	}
	if err := json.Unmarshal(envelope["data"], &data); err != nil || data.Checks == nil {
		t.Fatalf("doctor data=%s: %v", envelope["data"], err)
	}
	var diagnostics []map[string]json.RawMessage
	if err := json.Unmarshal(envelope["diagnostics"], &diagnostics); err != nil || diagnostics == nil {
		t.Fatalf("doctor diagnostics=%s: %v", envelope["diagnostics"], err)
	}
	return data.Checks, diagnostics
}

func doctorCLIHasCheck(checks []map[string]json.RawMessage, name string, passed bool) bool {
	for _, check := range checks {
		var gotName string
		var gotPassed bool
		_ = json.Unmarshal(check["name"], &gotName)
		_ = json.Unmarshal(check["passed"], &gotPassed)
		if gotName == name && gotPassed == passed {
			for _, field := range []string{"module", "detail"} {
				if _, ok := check[field]; !ok {
					return false
				}
			}
			return true
		}
	}
	return false
}

func doctorCLIPlainOutcome(output, name, outcome string) bool {
	for _, line := range strings.Split(output, "\n") {
		if strings.Contains(line, name) && strings.Contains(strings.ToLower(line), outcome) {
			return true
		}
	}
	return false
}

func Test_APP_062_APP_063_APP_064_APP_065_APP_066_APP_067_DoctorCLI(t *testing.T) {
	bin := buildAppBinary(t)
	root, nested, log := doctorCLIFixture(t, bin)
	status, plain, errOut := runAppBinary(t, bin, nested, "doctor", "--no-agent")
	if status != 0 {
		t.Fatalf("APP-062/066 doctor status=%d stdout=%q stderr=%q", status, plain, errOut)
	}
	plain += errOut
	for _, name := range []string{"baseline", "undeclared", "internal_eval", "failing_eval", "missing_id", "recovery"} {
		if !doctorCLIPlainOutcome(plain, name, "pass") {
			t.Errorf("APP-065 plain output lacks outcome %s: %q", name, plain)
		}
	}
	probeLog := doctorCLIFile(t, log)
	if !strings.Contains(probeLog, "setup:") || !strings.Contains(probeLog, bin) || strings.Contains(probeLog, "worker:") || strings.Contains(probeLog, "audit:") {
		t.Errorf("APP-062/063/064 scratch hook did not query NEW or agent ran: %q", probeLog)
	}
	status, jsonOut, errOut := runAppBinary(t, bin, nested, "doctor", "--no-agent", "--json")
	if status != 0 || errOut != "" {
		t.Fatalf("APP-067 JSON status=%d stdout=%q stderr=%q", status, jsonOut, errOut)
	}
	checks, diagnostics := doctorCLIChecks(t, jsonOut)
	for _, check := range checks {
		var name string
		var passed bool
		if err := json.Unmarshal(check["name"], &name); err != nil || name == "" {
			t.Errorf("APP-067 unnamed check: %v", check)
			continue
		}
		if err := json.Unmarshal(check["passed"], &passed); err != nil {
			t.Errorf("APP-067 check %s lacks outcome: %v", name, check)
			continue
		}
		outcome := "fail"
		if passed {
			outcome = "pass"
		}
		if !doctorCLIPlainOutcome(plain, name, outcome) {
			t.Errorf("APP-065 plain output lacks %s outcome for reported %s: %q", outcome, name, plain)
		}
	}
	for _, name := range []string{"baseline", "undeclared", "internal_eval", "failing_eval", "missing_id", "recovery"} {
		if !doctorCLIHasCheck(checks, name, true) {
			t.Errorf("APP-067 missing named passing check %s: %s", name, jsonOut)
		}
	}
	if len(diagnostics) != 0 {
		t.Errorf("APP-067 healthy diagnostics=%v", diagnostics)
	}
	if strings.Contains(doctorCLIFile(t, log), "worker:") || strings.Contains(doctorCLIFile(t, log), "audit:") {
		t.Error("APP-064 --no-agent ran worker or audit")
	}

	// A stale managed block is a warning: the command still succeeds and retains C16.
	appPut(t, root, "AGENTS.md", "Before.\n<!-- fireweed:start -->\nstale\n<!-- fireweed:end -->\nAfter.\n")
	status, plain, errOut = runAppBinary(t, bin, nested, "doctor", "--no-agent")
	if status != 0 || !strings.Contains(plain+errOut, "C16") || !strings.Contains(strings.ToLower(plain+errOut), "warn") || !strings.Contains(plain+errOut, "AGENTS.md") {
		t.Errorf("APP-065/066 warning-only status=%d stdout=%q stderr=%q", status, plain, errOut)
	}
	status, jsonOut, errOut = runAppBinary(t, bin, nested, "doctor", "--no-agent", "--json")
	if status != 0 || errOut != "" {
		t.Errorf("APP-067 warning JSON status=%d stderr=%q", status, errOut)
	}
	_, diagnostics = doctorCLIChecks(t, jsonOut)
	foundWarning := false
	for _, d := range diagnostics {
		if string(d["code"]) == `"C16"` && string(d["warning"]) == "true" && strings.Contains(string(d["message"]), "AGENTS.md") {
			foundWarning = true
		}
	}
	if !foundWarning {
		t.Errorf("APP-067 C16 warning lost: %s", jsonOut)
	}

	// The ordinary doctor path runs both bounded local agents.
	status, plain, errOut = runAppBinary(t, bin, nested, "doctor")
	if status != 0 || !strings.Contains(plain+errOut, "C16") {
		t.Errorf("APP-066 agent doctor status=%d stdout=%q stderr=%q", status, plain, errOut)
	}
	if got := doctorCLIFile(t, log); !strings.Contains(got, "worker:") || !strings.Contains(got, "audit:") {
		t.Errorf("APP-064 configured agents not exercised: %q", got)
	}
}

func Test_APP_065_APP_066_APP_067_DoctorFailures(t *testing.T) {
	bin := buildAppBinary(t)
	root, nested, _ := doctorCLIFixture(t, bin)
	doctorCLIConfig(t, root, func(config map[string]any) {
		config["hooks"].(map[string]any)["setup"] = "sh -c 'echo named-setup-failure >&2; exit 17'"
	})
	status, plain, errOut := runAppBinary(t, bin, nested, "doctor", "--no-agent")
	if status != 1 || !strings.Contains(plain+errOut, "C14") || !strings.Contains(plain+errOut, "named-setup-failure") || !doctorCLIPlainOutcome(plain+errOut, "setup", "fail") {
		t.Errorf("APP-065/066 failed probe status=%d output=%q", status, plain+errOut)
	}
	status, output, errOut := runAppBinary(t, bin, nested, "doctor", "--no-agent", "--json")
	if status != 1 || errOut != "" {
		t.Errorf("APP-067 failed JSON status=%d stderr=%q", status, errOut)
	}
	envelope := decodeEnvelope(t, output)
	if string(envelope["data"]) != "null" {
		t.Errorf("APP-067 failed data=%s", envelope["data"])
	}
	var diags []map[string]json.RawMessage
	if err := json.Unmarshal(envelope["diagnostics"], &diags); err != nil || len(diags) == 0 {
		t.Fatalf("APP-067 diagnostics=%s: %v", envelope["diagnostics"], err)
	}
	if string(diags[0]["code"]) != `"C14"` || !strings.Contains(string(diags[0]["message"]), "named-setup-failure") {
		t.Errorf("APP-067 failed diagnostic=%v", diags[0])
	}

	appPut(t, root, "packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\n---\n# Alpha\n\n## Specs\n\n### ALP-001\nOne MUST pass.\n\n### ALP-001\nAnother MUST pass.\n")
	status, plain, errOut = runAppBinary(t, bin, nested, "doctor", "--no-agent")
	if status != 1 || !strings.Contains(plain+errOut, "C05") {
		t.Errorf("APP-066/037 static failure status=%d output=%q", status, plain+errOut)
	}
	status, output, errOut = runAppBinary(t, bin, nested, "doctor", "--no-agent", "--json")
	if status != 1 || errOut != "" {
		t.Errorf("APP-067 static JSON status=%d stderr=%q", status, errOut)
	}
	envelope = decodeEnvelope(t, output)
	if string(envelope["data"]) != "null" {
		t.Errorf("APP-067 static data=%s", envelope["data"])
	}
	if err := json.Unmarshal(envelope["diagnostics"], &diags); err != nil || len(diags) == 0 {
		t.Fatalf("APP-067 static diagnostics=%s: %v", envelope["diagnostics"], err)
	}
	if string(diags[0]["code"]) != `"C05"` || !strings.Contains(string(diags[0]["path"]), "SPEC.md") || string(diags[0]["warning"]) != "false" || len(diags[0]["message"]) == 0 {
		t.Errorf("APP-067 static diagnostic=%v", diags[0])
	}

	for _, tc := range []struct {
		name string
		args []string
	}{
		{"extra", []string{"doctor", "unexpected"}},
		{"unknown-flag", []string{"doctor", "--bad"}},
		{"wrong-command-flag", []string{"modules", "--no-agent"}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			status, stdout, stderr := runAppBinary(t, bin, nested, tc.args...)
			if status != 2 || !strings.Contains(stdout+stderr, "U01") || !corrective(stdout+stderr) {
				t.Errorf("APP-066 usage status=%d output=%q", status, stdout+stderr)
			}
			status, stdout, stderr = runAppBinary(t, bin, nested, append(tc.args, "--json")...)
			if status != 2 || stderr != "" {
				t.Errorf("APP-067 usage JSON status=%d stderr=%q", status, stderr)
			}
			assertFailedJSON(t, stdout, "U01")
		})
	}
	appPut(t, root, "fireweed.yaml", "[invalid yaml\n")
	status, plain, errOut = runAppBinary(t, bin, nested, "doctor")
	if status != 2 || !strings.Contains(plain+errOut, "C01") {
		t.Errorf("APP-066 config status=%d output=%q", status, plain+errOut)
	}
	status, output, errOut = runAppBinary(t, bin, nested, "doctor", "--json")
	if status != 2 || errOut != "" {
		t.Errorf("APP-067 config JSON status=%d stderr=%q", status, errOut)
	}
	assertFailedJSON(t, output, "C01")
	if err := os.Remove(filepath.Join(root, "fireweed.yaml")); err != nil {
		t.Fatal(err)
	}
	status, plain, errOut = runAppBinary(t, bin, nested, "doctor")
	if status != 2 || !strings.Contains(plain+errOut, "C01") {
		t.Errorf("APP-066 missing config status=%d output=%q", status, plain+errOut)
	}
}

func Test_APP_068_APP_069_APP_070_APP_071_APP_072_APP_073_AgentsMDWrite(t *testing.T) {
	bin := buildAppBinary(t)
	root := appProjectFixture(t)
	nested := appNestedDir(t, root)
	system := "# Existing system\n"
	appPut(t, root, "SYSTEM.md", system)
	before := "Before.\n\n<!-- fireweed:start -->\nold guidance\n<!-- fireweed:end -->\n\nAfter.\n"
	appPut(t, root, "AGENTS.md", before)
	status, stdout, stderr := runAppBinary(t, bin, nested, "agents-md", "--write")
	if status != 0 || strings.TrimSpace(stdout) == "" {
		t.Errorf("APP-068/072 status=%d stdout=%q stderr=%q", status, stdout, stderr)
	}
	agents := readInitFile(t, root, "AGENTS.md")
	if agents != "Before.\n\n"+doctorCLIManagedBlock+"\n\nAfter.\n" {
		t.Errorf("APP-069 managed guidance=%q", agents)
	}
	if got := readInitFile(t, root, "SYSTEM.md"); got != system {
		t.Errorf("APP-068 changed SYSTEM.md: %q", got)
	}
	status, output, stderr := runAppBinary(t, bin, nested, "agents-md", "--write", "--json")
	if status != 0 || stderr != "" {
		t.Errorf("APP-073 status=%d stderr=%q", status, stderr)
	}
	envelope := decodeEnvelope(t, output)
	var confirmation string
	if err := json.Unmarshal(envelope["data"], &confirmation); err != nil || confirmation != strings.TrimSpace(stdout) {
		t.Errorf("APP-073 confirmation=%s: %v", envelope["data"], err)
	}
	var diags []json.RawMessage
	if err := json.Unmarshal(envelope["diagnostics"], &diags); err != nil || diags == nil || len(diags) != 0 {
		t.Errorf("APP-073 diagnostics=%s: %v", envelope["diagnostics"], err)
	}

	for _, tc := range []struct {
		name, initial string
		absent        bool
	}{
		{"no-block", "Personal instructions.\n", false},
		{"missing-file", "", true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if tc.absent {
				if err := os.Remove(filepath.Join(root, "AGENTS.md")); err != nil {
					t.Fatal(err)
				}
			} else {
				appPut(t, root, "AGENTS.md", tc.initial)
			}
			status, stdout, stderr := runAppBinary(t, bin, nested, "agents-md", "--write")
			if status != 0 || strings.TrimSpace(stdout) == "" {
				t.Errorf("APP-070/072 status=%d stdout=%q stderr=%q", status, stdout, stderr)
			}
			got := readInitFile(t, root, "AGENTS.md")
			if !strings.HasPrefix(got, tc.initial) || strings.Count(got, "<!-- fireweed:start -->") != 1 || !strings.Contains(got, "playbook author-change") {
				t.Errorf("APP-070 content=%q", got)
			}
		})
	}

	for _, malformed := range []string{"Personal.\n<!-- fireweed:start -->\nold\n", "Personal.\n<!-- fireweed:end -->\n"} {
		appPut(t, root, "AGENTS.md", malformed)
		status, stdout, stderr := runAppBinary(t, bin, nested, "agents-md", "--write")
		if status != 2 || !strings.Contains(stdout+stderr, "P01") || !corrective(stdout+stderr) || readInitFile(t, root, "AGENTS.md") != malformed {
			t.Errorf("APP-071/072 malformed status=%d output=%q", status, stdout+stderr)
		}
		status, output, stderr := runAppBinary(t, bin, nested, "agents-md", "--write", "--json")
		if status != 2 || stderr != "" {
			t.Errorf("APP-073 malformed JSON status=%d stderr=%q", status, stderr)
		}
		assertFailedJSON(t, output, "P01")
	}

	for _, args := range [][]string{{"agents-md"}, {"agents-md", "--bad"}, {"agents-md", "--write", "extra"}, {"doctor", "--write"}} {
		status, stdout, stderr := runAppBinary(t, bin, nested, args...)
		if status != 2 || !strings.Contains(stdout+stderr, "U01") || !corrective(stdout+stderr) {
			t.Errorf("APP-072 usage %v: status=%d output=%q", args, status, stdout+stderr)
		}
		status, output, errOut := runAppBinary(t, bin, nested, append(args, "--json")...)
		if status != 2 || errOut != "" {
			t.Errorf("APP-073 usage JSON %v: status=%d stderr=%q", args, status, errOut)
		}
		assertFailedJSON(t, output, "U01")
	}
	if got := readInitFile(t, root, "AGENTS.md"); got != "Personal.\n<!-- fireweed:end -->\n" {
		t.Errorf("APP-072 usage changed guidance: %q", got)
	}
	if err := os.Remove(filepath.Join(root, "AGENTS.md")); err != nil {
		t.Fatal(err)
	}
	if err := os.Mkdir(filepath.Join(root, "AGENTS.md"), 0755); err != nil {
		t.Fatal(err)
	}
	status, stdout, stderr = runAppBinary(t, bin, nested, "agents-md", "--write")
	if status != 2 || !strings.Contains(stdout+stderr, "P01") || !corrective(stdout+stderr) {
		t.Errorf("APP-072 operational failure status=%d output=%q", status, stdout+stderr)
	}
	status, output, stderr = runAppBinary(t, bin, nested, "agents-md", "--write", "--json")
	if status != 2 || stderr != "" {
		t.Errorf("APP-073 operational JSON status=%d stderr=%q", status, stderr)
	}
	assertFailedJSON(t, output, "P01")
	if err := os.Remove(filepath.Join(root, "AGENTS.md")); err != nil {
		t.Fatal(err)
	}

	appPut(t, root, "fireweed.yaml", "[invalid yaml\n")
	status, stdout, stderr = runAppBinary(t, bin, nested, "agents-md", "--write")
	if status != 2 || !strings.Contains(stdout+stderr, "C01") {
		t.Errorf("APP-072 config status=%d output=%q", status, stdout+stderr)
	}
	status, output, stderr = runAppBinary(t, bin, nested, "agents-md", "--write", "--json")
	if status != 2 || stderr != "" {
		t.Errorf("APP-073 config JSON status=%d stderr=%q", status, stderr)
	}
	assertFailedJSON(t, output, "C01")
	if err := os.Remove(filepath.Join(root, "fireweed.yaml")); err != nil {
		t.Fatal(err)
	}
	status, stdout, stderr = runAppBinary(t, bin, nested, "agents-md", "--write")
	if status != 2 || !strings.Contains(stdout+stderr, "C01") {
		t.Errorf("APP-072 missing config status=%d output=%q", status, stdout+stderr)
	}
}
