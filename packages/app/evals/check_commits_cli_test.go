package evals_test

import (
	"encoding/json"
	"reflect"
	"regexp"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/checks"
)

const checkCommitsChange = "---\nmodule: alpha\n%s---\n\n## Motivation\nRecord a local change.\n\n## Interface\n\n## Added\n\n## Modified\n\n## Removed\n\n## Decisions\n\n## Notes\n"

func checkCommitsFixture(t *testing.T) (root, initial, configuredBase string) {
	t.Helper()
	root = t.TempDir()
	config := strings.Replace(appFixtureConfig, "base: main", "base: review-base", 1) + "protected: [locked/**]\nmodule_defaults:\n  interface: [api.go]\n  evals: [evals/**]\n  append_only: [ledger/**]\n"
	// Keep one module_defaults mapping in the recorded configuration.
	config = strings.Replace(config, "module_defaults:\n  interface: [api.go]\n  evals: [evals/**]\n  disposable: [internal/**]\n  preserve: [data/**]\n", "", 1)
	appPut(t, root, "fireweed.yaml", config)
	appPut(t, root, "packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\n---\n# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST retain its initial behavior.\n")
	appPut(t, root, "packages/alpha/api.go", "package alpha\n")
	appPut(t, root, "packages/alpha/ledger/entry.txt", "initial\n")
	appPut(t, root, "locked/config.txt", "initial\n")
	appPut(t, root, "changes/archive/0000-old.md", strings.Replace(checkCommitsChange, "%s", "phase: applied\n", 1))
	cliGit(t, root, "init", "-q", "-b", "main")
	cliGit(t, root, "config", "user.name", "Fixture")
	cliGit(t, root, "config", "user.email", "fixture@example.invalid")
	cliCommit(t, root, "initial project")
	initial = cliGit(t, root, "rev-parse", "HEAD")
	cliGit(t, root, "checkout", "-qb", "review-base")
	appPut(t, root, "notes/common.txt", "common\n")
	cliCommit(t, root, "review base content")
	configuredBase = cliGit(t, root, "rev-parse", "HEAD")
	cliGit(t, root, "checkout", "-qb", "feature")
	return root, initial, configuredBase
}

type checkCommitsSquash struct {
	Commit string `json:"commit"`
	Change string `json:"change"`
}

func checkCommitsJSON(t *testing.T, output string) (commits []string, squashed []checkCommitsSquash) {
	t.Helper()
	envelope := decodeEnvelope(t, output)
	var diagnostics []json.RawMessage
	if err := json.Unmarshal(envelope["diagnostics"], &diagnostics); err != nil || diagnostics == nil || len(diagnostics) != 0 {
		t.Fatalf("success diagnostics=%s; want empty array: %v", envelope["diagnostics"], err)
	}
	var data struct {
		Commits  []string             `json:"commits"`
		Squashed []checkCommitsSquash `json:"squashed"`
	}
	if err := json.Unmarshal(envelope["data"], &data); err != nil || data.Commits == nil || data.Squashed == nil {
		t.Fatalf("data=%s; want object with commits and squashed arrays: %v", envelope["data"], err)
	}
	return data.Commits, data.Squashed
}

func Test_APP_105_APP_108_CheckCommitsRangesAndRecordedAuthorWork(t *testing.T) {
	root, initial, configuredBase := checkCommitsFixture(t)
	queued := strings.Replace(checkCommitsChange, "%s", "", 1)
	appPut(t, root, "changes/0001-queued.md", queued)
	cliCommit(t, root, "ordinary queued document without phase")
	queuedOID := cliGit(t, root, "rev-parse", "HEAD")
	cliGit(t, root, "mv", "changes/0001-queued.md", "changes/archive/0001-queued.md")
	appPut(t, root, "changes/archive/0001-queued.md", strings.Replace(checkCommitsChange, "%s", "phase: applied\n", 1))
	appPut(t, root, "packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\n---\n# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST retain its initial behavior.\n\n### ALP-002\nAlpha MUST expose the recorded addition.\n")
	cliCommit(t, root, "apply change and legitimate SPEC addition")
	specOID := cliGit(t, root, "rev-parse", "HEAD")
	for _, tc := range []struct {
		rangeValue string
		want       []string
	}{
		{initial + "..HEAD", []string{configuredBase, queuedOID, specOID}},
		{"", []string{queuedOID, specOID}},
	} {
		report, err := checks.History(root, tc.rangeValue)
		if err != nil || !reflect.DeepEqual(report.Commits, tc.want) || len(report.Diagnostics) != 0 || len(report.Squashed) != 1 || report.Squashed[0].Commit != specOID || report.Squashed[0].Change != "changes/archive/0001-queued.md" {
			t.Fatalf("clean history fixture range %q: report=%+v err=%v", tc.rangeValue, report, err)
		}
	}

	nested := appNestedDir(t, root)
	status, plain, stderr := runSavedAttemptApp(t, nested, "check-commits", initial+"..HEAD")
	if status != 0 || stderr != "" || !regexp.MustCompile(`\b3\b`).MatchString(plain) {
		t.Errorf("explicit range: exit=%d stdout=%q stderr=%q; want three clean commits", status, plain, stderr)
	}
	status, plain, stderr = runSavedAttemptApp(t, nested, "check-commits")
	if status != 0 || stderr != "" || !regexp.MustCompile(`\b2\b`).MatchString(plain) {
		t.Errorf("configured-base range: exit=%d stdout=%q stderr=%q; want two clean commits", status, plain, stderr)
	}
	status, output, stderr := runSavedAttemptApp(t, nested, "check-commits", "--json")
	if status != 0 || stderr != "" {
		t.Fatalf("configured-base JSON: exit=%d stdout=%q stderr=%q", status, output, stderr)
	}
	commits, squashed := checkCommitsJSON(t, output)
	if !reflect.DeepEqual(commits, []string{queuedOID, specOID}) || !reflect.DeepEqual(squashed, []checkCommitsSquash{{Commit: specOID, Change: "changes/archive/0001-queued.md"}}) {
		t.Errorf("configured-base data: commits=%v squashed=%v; want queued and applied SPEC commits with archive", commits, squashed)
	}
	status, output, stderr = runSavedAttemptApp(t, nested, "check-commits", initial+"..HEAD", "--json")
	if status != 0 || stderr != "" {
		t.Fatalf("explicit JSON: exit=%d stdout=%q stderr=%q", status, output, stderr)
	}
	commits, _ = checkCommitsJSON(t, output)
	if !reflect.DeepEqual(commits, []string{configuredBase, queuedOID, specOID}) {
		t.Errorf("explicit commits=%v; want configured-base, queued, SPEC order", commits)
	}
}

func Test_APP_105_APP_108_CheckCommitsSquashAndEmptyJSON(t *testing.T) {
	root, _, base := checkCommitsFixture(t)
	nested := appNestedDir(t, root)
	paths := []string{"changes/archive/0001-one.md", "changes/archive/0002-two.md"}
	for _, path := range paths {
		appPut(t, root, path, strings.Replace(checkCommitsChange, "%s", "phase: applied\n", 1))
	}
	cliCommit(t, root, "squash two applied changes")
	oid := cliGit(t, root, "rev-parse", "HEAD")
	report, err := checks.History(root, base+"..HEAD")
	if err != nil || !reflect.DeepEqual(report.Commits, []string{oid}) || len(report.Squashed) != 2 || len(report.Diagnostics) != 0 {
		t.Fatalf("clean squash fixture: report=%+v err=%v", report, err)
	}
	if !regexp.MustCompile(`^[0-9a-f]{40}$`).MatchString(oid) {
		t.Fatalf("fixture commit is not a full object ID: %q", oid)
	}
	status, output, stderr := runSavedAttemptApp(t, nested, "check-commits", base+".."+base, "--json")
	if status != 0 || stderr != "" {
		t.Fatalf("empty JSON: exit=%d stdout=%q stderr=%q", status, output, stderr)
	}
	commits, squashed := checkCommitsJSON(t, output)
	if len(commits) != 0 || len(squashed) != 0 {
		t.Errorf("empty range: commits=%v squashed=%v", commits, squashed)
	}
	status, plain, stderr := runSavedAttemptApp(t, nested, "check-commits", base+"..HEAD")
	if status != 0 || stderr != "" || !regexp.MustCompile(`\b1\b`).MatchString(plain) || strings.Contains(strings.ToLower(plain), "warning") {
		t.Errorf("squash plain: exit=%d stdout=%q stderr=%q; want one clean commit", status, plain, stderr)
	}
	for _, path := range paths {
		if !strings.Contains(plain, oid) || !strings.Contains(plain, path) {
			t.Errorf("plain squash %q must name full %s and %s", plain, oid, path)
		}
	}
	status, output, stderr = runSavedAttemptApp(t, nested, "check-commits", base+"..HEAD", "--json")
	if status != 0 || stderr != "" {
		t.Fatalf("squash JSON: exit=%d stdout=%q stderr=%q", status, output, stderr)
	}
	commits, squashed = checkCommitsJSON(t, output)
	if !reflect.DeepEqual(commits, []string{oid}) || len(squashed) != 2 {
		t.Fatalf("squash data: commits=%v squashed=%v; want one commit and two entries", commits, squashed)
	}
	for i, entry := range squashed {
		if entry.Commit != oid || entry.Change != paths[i] {
			t.Errorf("squashed[%d]=%+v; want %s %s", i, entry, oid, paths[i])
		}
	}
}

func Test_APP_106_APP_108_CheckCommitsReportsEveryHistoryRule(t *testing.T) {
	root, _, base := checkCommitsFixture(t)
	pending := "changes/0001-work.md"
	appPut(t, root, pending, strings.Replace(checkCommitsChange, "%s", "phase: impl\n", 1))
	cliCommit(t, root, "ready for application")
	archive := "changes/archive/0001-work.md"
	cliGit(t, root, "mv", pending, archive)
	appPut(t, root, archive, strings.Replace(checkCommitsChange, "%s", "phase: applied\n", 1))
	appPut(t, root, "locked/config.txt", "worker edit\n")
	appPut(t, root, "packages/alpha/api.go", "package alpha // changed at application\n")
	appPut(t, root, "packages/alpha/ledger/entry.txt", "modified\n")
	appPut(t, root, "changes/archive/0000-old.md", strings.Replace(checkCommitsChange, "%s", "phase: applied\n", 1)+"\narchive edit\n")
	cliCommit(t, root, "application with four rule failures")
	oid := cliGit(t, root, "rev-parse", "HEAD")
	want := map[string]string{"H01": "locked/config.txt", "H02": "packages/alpha/api.go", "H03": "packages/alpha/ledger/entry.txt", "H04": "changes/archive/0000-old.md"}
	report, err := checks.History(root, base+"..HEAD")
	if err != nil || len(report.Diagnostics) != len(want) {
		t.Fatalf("four-rule history fixture: report=%+v err=%v", report, err)
	}
	for _, d := range report.Diagnostics {
		if want[d.Code] != d.Path || !strings.Contains(d.Message, oid) {
			t.Fatalf("unexpected history fixture diagnostic: %+v", d)
		}
	}
	nested := appNestedDir(t, root)
	status, plain, stderr := runSavedAttemptApp(t, nested, "check-commits", base+"..HEAD")
	if status != 1 || stderr == "" && plain == "" {
		t.Errorf("plain rules: exit=%d stdout=%q stderr=%q", status, plain, stderr)
	}
	for code, path := range want {
		found := false
		for _, line := range strings.Split(plain+stderr, "\n") {
			if strings.Contains(line, code) && strings.Contains(line, path) && strings.Contains(line, oid) && corrective(line) {
				found = true
			}
		}
		if !found {
			t.Errorf("plain diagnostics lack actionable %s, %s, %s: %q", code, path, oid, plain+stderr)
		}
	}
	status, output, stderr := runSavedAttemptApp(t, nested, "check-commits", base+"..HEAD", "--json")
	if status != 1 || stderr != "" {
		t.Errorf("JSON rules: exit=%d stdout=%q stderr=%q", status, output, stderr)
	}
	envelope := decodeEnvelope(t, output)
	if string(envelope["data"]) != "null" {
		t.Errorf("rule-failure data=%s; want null", envelope["data"])
	}
	var diagnostics []struct{ Code, Message, Path string }
	if err := json.Unmarshal(envelope["diagnostics"], &diagnostics); err != nil {
		t.Fatal(err)
	}
	if len(diagnostics) != len(want) {
		t.Errorf("diagnostics=%+v; want four returned rules", diagnostics)
	}
	for code, path := range want {
		found := false
		for _, d := range diagnostics {
			if d.Code == code && d.Path == path && strings.Contains(d.Message, oid) && corrective(d.Message) {
				found = true
			}
		}
		if !found {
			t.Errorf("JSON diagnostics lack actionable %s, %s, %s: %+v", code, path, oid, diagnostics)
		}
	}
}

func Test_APP_107_APP_108_CheckCommitsFailures(t *testing.T) {
	root, _, base := checkCommitsFixture(t)
	nested := appNestedDir(t, root)
	for _, tc := range []struct {
		name   string
		args   []string
		status int
		code   string
	}{
		{"extra", []string{"check-commits", base + "..HEAD", "extra", "--json"}, 2, "U01"},
		{"option", []string{"check-commits", "--bogus"}, 2, "U01"},
		{"bad-range", []string{"check-commits", "missing-ref..HEAD", "--json"}, 1, "P01"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			status, output, stderr := runSavedAttemptApp(t, nested, tc.args...)
			message := commandFailure(t, status, tc.status, output, stderr, strings.Contains(strings.Join(tc.args, " "), "--json"))
			if strings.Contains(strings.Join(tc.args, " "), "--json") {
				message = assertFailedJSON(t, output, tc.code)
			} else if !strings.Contains(output+stderr, tc.code+":") {
				t.Errorf("diagnostic=%q; want %s", output+stderr, tc.code)
			}
			if !corrective(message) {
				t.Errorf("diagnostic %q needs corrective action", message)
			}
		})
	}
	appPut(t, root, "fireweed.yaml", "modules: [broken\n")
	status, output, stderr := runSavedAttemptApp(t, nested, "check-commits", "--json")
	message := commandFailure(t, status, 2, output, stderr, true)
	message = assertFailedJSON(t, output, "C01")
	if !corrective(message) {
		t.Errorf("configuration diagnostic %q needs corrective action", message)
	}

	codedRoot, _, codedBase := checkCommitsFixture(t)
	appPut(t, codedRoot, "packages/alpha/SPEC.md", "---\nname: [broken\n---\n")
	cliCommit(t, codedRoot, "record malformed SPEC")
	cliGit(t, codedRoot, "checkout", "HEAD^", "--", "packages/alpha/SPEC.md")
	cliCommit(t, codedRoot, "restore current SPEC")
	if _, err := checks.History(codedRoot, codedBase+"..HEAD"); err == nil || !strings.Contains(err.Error(), "parse packages/alpha/SPEC.md") {
		t.Fatalf("coded operational fixture: want malformed recorded SPEC, got %v", err)
	}
	status, output, stderr = runSavedAttemptApp(t, appNestedDir(t, codedRoot), "check-commits", codedBase+"..HEAD", "--json")
	message = commandFailure(t, status, 1, output, stderr, true)
	message = assertFailedJSON(t, output, "C03")
	if !corrective(message) {
		t.Errorf("coded operational diagnostic %q needs corrective action", message)
	}
}
