package evals_test

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/assets"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow"
)

const fixtureBinaryPin = ">=0.1, <0.2"

func incompatiblePin() string { return ">" + assets.Version() }

func incompatibleConfig(config string) string {
	return strings.Replace(config, fixtureBinaryPin, incompatiblePin(), 1)
}

var compatibilityGuidance = regexp.MustCompile(`(?i)\b(?:correct|edit|fix|update|change|adjust|check|use|install|select|choose|run|try|switch|upgrade|downgrade)\b[^.\n]*\b(?:fireweed\.yaml|config(?:uration)?|project inputs|pin|range|binary|version|format|compatible)\b`)

func assertCompatibilityRejection(t *testing.T, bin, dir, cause string, args ...string) {
	t.Helper()
	status, out, stderr := runLifecycleBinary(t, bin, dir, args...)
	if status != 2 {
		t.Errorf("exit=%d, want 2; stdout=%q stderr=%q", status, out, stderr)
	}
	message := stderr
	if args[len(args)-1] == "--json" {
		if stderr != "" {
			t.Errorf("JSON stderr=%q; want empty", stderr)
		}
		message = assertFailedJSON(t, out, "C01")
	} else if out != "" || !strings.HasPrefix(strings.TrimSpace(stderr), "C01:") {
		t.Errorf("plain failure stdout=%q stderr=%q; want empty stdout and C01 diagnostic", out, stderr)
	}
	if !strings.Contains(message, cause) {
		t.Errorf("diagnostic %q must identify %q", message, cause)
	}
	if !compatibilityGuidance.MatchString(message) {
		t.Errorf("diagnostic lacks advice to correct the configuration or use a compatible binary: %q", message)
	}
}

func assertCompatibilityState(t *testing.T, root string, before migrateState, calls string) {
	t.Helper()
	after := migrateCapture(t, root)
	if !reflect.DeepEqual(after, before) {
		t.Error("compatibility rejection changed working files, index, HEAD or refs")
	}
	if log, err := os.ReadFile(calls); !os.IsNotExist(err) {
		t.Errorf("compatibility rejection ran fixture commands: %q (%v)", log, err)
	}
}

func compatibilityProject(t *testing.T, bin string) (root, calls string) {
	t.Helper()
	root = lifecycleCLIProject(t, bin, false)
	calls = filepath.Join(t.TempDir(), "calls")
	config := string(staleCLIText(t, root, "fireweed.yaml"))
	config += fmt.Sprintf(`agent:
  command: [sh, .fireweed/agent.sh, %q, "{prompt_file}"]
limit_patterns: [RATE LIMIT]
limit_wait: 5ms
limit_budget: 0s
`, calls)
	appPut(t, root, "fireweed.yaml", config)
	appPut(t, root, ".fireweed/agent.sh", "#!/bin/sh\nprintf 'agent\\n' >> \"$1\"\nprintf 'RATE LIMIT\\n' >&2\nexit 9\n")
	appPut(t, root, ".fireweed/eval.sh", fmt.Sprintf(`#!/bin/sh
printf 'eval\n' >> %q
printf '<testsuite><testcase name="ALP-001"/><testcase name="ALP-002"/></testsuite>' > "$2"
`, calls))
	cliCommit(t, root, "configure bounded compatibility fixture commands")
	return root, calls
}

func Test_APP_163_LiveCompatibilityBeforeCommandEffects(t *testing.T) {
	bin := buildAppBinary(t)
	root, calls := compatibilityProject(t, bin)
	dir := appNestedDir(t, root)
	config := string(staleCLIText(t, root, "fireweed.yaml"))
	appPut(t, root, "fireweed.yaml", incompatibleConfig(config))
	cliCommit(t, root, "record incompatible binary pin")
	for _, args := range [][]string{
		{"show", "alpha", "--history"},
		{"check-commits", "HEAD..HEAD", "--json"},
		{"start", "manual", "--json"},
		{"new-change", "alpha", "new-work"},
	} {
		t.Run(strings.Join(args, " "), func(t *testing.T) {
			before := migrateCapture(t, root)
			t.Cleanup(func() { assertCompatibilityState(t, root, before, calls) })
			assertCompatibilityRejection(t, bin, dir, incompatiblePin(), args...)
		})
	}
	t.Run("unsupported-format", func(t *testing.T) {
		appPut(t, root, "fireweed.yaml", strings.Replace(config, "version: 1", "version: 999", 1))
		before := migrateCapture(t, root)
		t.Cleanup(func() { assertCompatibilityState(t, root, before, calls) })
		assertCompatibilityRejection(t, bin, dir, "999", "coverage", "alpha", "--json")
	})

	t.Run("live-pin-for-history", func(t *testing.T) {
		// The inspected commit still has the excluding pin, but the live pin admits it.
		appPut(t, root, "fireweed.yaml", config)
		status, out, stderr := runAppBinary(t, bin, dir, "check-commits", "HEAD^..HEAD", "--json")
		if status != 0 || stderr != "" {
			t.Fatalf("check-commits: exit=%d stdout=%q stderr=%q", status, out, stderr)
		}
		commits, _ := checkCommitsJSON(t, out)
		if want := []string{cliGit(t, root, "rev-parse", "HEAD")}; !reflect.DeepEqual(commits, want) {
			t.Errorf("checked commits = %v, want %v", commits, want)
		}
	})
}

func Test_APP_165_StandaloneAndSetupPreconditions(t *testing.T) {
	bin := buildAppBinary(t)
	root := migrateProject(t)
	badConfig := incompatibleConfig(appFixtureConfig)
	appPut(t, root, "fireweed.yaml", badConfig)

	for _, tc := range informationCases(t) {
		status, out, stderr := runAppBinary(t, bin, root, tc.args...)
		if status != 0 {
			t.Errorf("%v exit=%d, want 0", tc.args, status)
		}
		assertInformationOutput(t, out, stderr, tc.want)
	}
	status, out, stderr := runAppBinary(t, bin, root, "coverage", "--help", "--json")
	if status != 0 || stderr != "" || commandStringData(t, out) != helpText(t, "coverage") {
		t.Errorf("command help with an incompatible pin: exit=%d stdout=%q stderr=%q", status, out, stderr)
	}

	before := migrateCapture(t, root)
	status, out, stderr = runAppBinary(t, bin, root, "init", "--recipe", "none", "--agent", "codex")
	assertAppHistoryFailure(t, status, 2, out, stderr, false, "P01", "fireweed.yaml")
	migrateUnchanged(t, root, before)

	appPut(t, root, "fireweed.yaml", strings.Replace(appFixtureConfig, "version: 1\n", "", 1))
	before = migrateCapture(t, root)
	status, out, stderr = runAppBinary(t, bin, root, "migrate", "--json")
	assertAppHistoryFailure(t, status, 2, out, stderr, true, "C01", "")
	migrateUnchanged(t, root, before)

	appPut(t, root, "fireweed.yaml", appFixtureConfig)
	before = migrateCapture(t, root)
	status, out, stderr = runAppBinary(t, bin, root, "migrate", "--json")
	if status != 0 || stderr != "" || commandStringData(t, out) != migrateConfirmation {
		t.Errorf("migrate format 1: exit=%d stdout=%q stderr=%q", status, out, stderr)
	}
	migrateUnchanged(t, root, before)
}

// Build a recorded claim with the pin of a different binary. Start creates
// the fixture claim with a supported pin because live loading enforces it.
func recordIncompatibleClaim(t *testing.T, root string) {
	t.Helper()
	const ref = "refs/fireweed/claims/manual"
	oid := cliGit(t, root, "rev-parse", ref)
	blob, err := gitstore.ReadBlob(root, oid)
	if err != nil {
		t.Fatal(err)
	}
	recorded := string(blob)
	escape := func(pin string) string {
		encoded, err := json.Marshal(pin)
		if err != nil {
			t.Fatal(err)
		}
		return strings.Trim(string(encoded), `"`)
	}
	original := escape(fixtureBinaryPin)
	if strings.Count(recorded, original) != 1 {
		t.Fatal("fixture claim must contain exactly one binary pin")
	}
	recorded = strings.Replace(recorded, original, escape(incompatiblePin()), 1)
	updated, err := gitstore.WriteBlob(root, []byte(recorded))
	if err != nil {
		t.Fatal(err)
	}
	if err := gitstore.CASRef(root, ref, oid, updated); err != nil {
		t.Fatal(err)
	}
	approved, held, err := workflow.Approved(root)
	if err != nil || !held || approved.Snapshot.Config["fireweed"] != incompatiblePin() || project.CheckCompatibility(approved.Snapshot.Config) == nil {
		t.Fatalf("fixture must expose the incompatible approved configuration: held=%v error=%v", held, err)
	}
}

func Test_APP_164_RejectsIncompatibleApprovedConfiguration(t *testing.T) {
	bin := buildAppBinary(t)
	for _, args := range [][]string{
		{"modules"},
		{"coverage", "alpha", "--with-change", "manual", "--json"},
		{"next", "--json"},
		{"run"},
	} {
		t.Run(args[0], func(t *testing.T) {
			root, calls := compatibilityProject(t, bin)
			cliHold(t, root)
			recordIncompatibleClaim(t, root)
			dir := appNestedDir(t, root)
			before := migrateCapture(t, root)
			t.Cleanup(func() { assertCompatibilityState(t, root, before, calls) })
			assertCompatibilityRejection(t, bin, dir, incompatiblePin(), args...)
		})
	}
}

func Test_APP_164_HeldRunnerAdmissionUsesApprovedConfiguration(t *testing.T) {
	bin := buildAppBinary(t)
	root, calls := compatibilityProject(t, bin)
	cliHold(t, root)
	dir := appNestedDir(t, root)
	appPut(t, root, "fireweed.yaml", incompatibleConfig(string(staleCLIText(t, root, "fireweed.yaml"))))
	for _, args := range [][]string{{"check", "--json"}, {"show", "alpha"}, {"new-change", "alpha", "new-work"}} {
		before := migrateCapture(t, root)
		t.Run(strings.Join(args, " "), func(t *testing.T) {
			t.Cleanup(func() { assertCompatibilityState(t, root, before, calls) })
			assertCompatibilityRejection(t, bin, dir, incompatiblePin(), args...)
		})
	}
	for i, tc := range []struct{ command, liveConfig string }{
		{"next", "modules: [broken\n"},
		{"run", incompatibleConfig(appFixtureConfig)},
	} {
		t.Run(tc.command, func(t *testing.T) {
			appPut(t, root, "fireweed.yaml", tc.liveConfig)
			status, out, stderr := runLifecycleBinary(t, bin, dir, tc.command, "--json")
			if status != 3 || stderr != "" {
				t.Fatalf("held %s: exit=%d stdout=%q stderr=%q; want a usage stop", tc.command, status, out, stderr)
			}
			cliStopEnvelope(t, out, "0s", "5ms", "0s")
			if got := string(mustReadNextCLI(t, calls)); got != strings.Repeat("agent\n", i+1) {
				t.Errorf("fixture calls = %q; want only the approved agent for each invocation", got)
			}
		})
	}
}
