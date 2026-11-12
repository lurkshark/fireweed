package evals_test

import (
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/assets"
)

var helpForms = []struct {
	args      []string
	command   string
	badConfig bool
}{
	{[]string{"help"}, "", false},
	{[]string{"--help"}, "", true},
	{[]string{"-h"}, "", false},
	{[]string{"help", "coverage"}, "coverage", false},
	{[]string{"agents-md", "help"}, "agents-md", false},
	{[]string{"doctor", "help"}, "doctor", true},
	{[]string{"init", "--help"}, "init", false},
	{[]string{"init", "-h"}, "init", true},
	{[]string{"new-module", "--help"}, "new-module", true},
	{[]string{"new-module", "-h"}, "new-module", false},
}

func helpFixture(t *testing.T) (outside, invalid string) {
	t.Helper()
	outside, invalid = t.TempDir(), t.TempDir()
	appPut(t, invalid, "fireweed.yaml", "modules: [broken\n")
	return outside, invalid
}

func helpText(t *testing.T, command string) string {
	t.Helper()
	if command == "" {
		return assets.Help()
	}
	page, err := assets.CommandHelp(command)
	if err != nil || page == "" {
		t.Fatalf("CommandHelp(%q) returned empty help or an error: %v", command, err)
	}
	return page
}

func Test_APP_156_HelpForms(t *testing.T) {
	outside, invalid := helpFixture(t)
	for _, form := range helpForms {
		t.Run(strings.Join(form.args, " "), func(t *testing.T) {
			dir := outside
			if form.badConfig {
				dir = invalid
			}
			status, stdout, stderr := runSavedAttemptApp(t, dir, form.args...)
			if want := helpText(t, form.command); status != 0 || stderr != "" || stdout != want {
				t.Errorf("exit=%d stdout=%q stderr=%q; want exit 0, help text %q, and empty stderr", status, stdout, stderr, want)
			}
		})
	}
	if entries, err := os.ReadDir(outside); err != nil || len(entries) != 0 {
		t.Errorf("help changed the empty directory: entries=%v error=%v", entries, err)
	}
}

func Test_APP_156_EachListedCommandHasCLIHelp(t *testing.T) {
	// Overview rows put two or more spaces between the group and its command list.
	rows := regexp.MustCompile(`(?m)^[A-Za-z]+ {2,}([a-z][a-z-]*(?:, [a-z][a-z-]*)*)$`).FindAllStringSubmatch(assets.Help(), -1)
	if len(rows) == 0 {
		t.Fatal("no command rows in the help overview")
	}
	outside := t.TempDir()
	for _, row := range rows {
		for _, command := range strings.Split(row[1], ", ") {
			status, stdout, stderr := runSavedAttemptApp(t, outside, "help", command)
			if want := helpText(t, command); status != 0 || stderr != "" || stdout != want {
				t.Errorf("help %s: exit=%d stdout=%q stderr=%q; want its assets command page and empty stderr", command, status, stdout, stderr)
			}
		}
	}
}

func Test_APP_156_ExactHelpWordForm(t *testing.T) {
	root := appProjectFixture(t)
	appPut(t, root, "packages/help/SPEC.md", "---\nname: help\nid_prefix: HLP\n---\n\n# Help module\n\n## Specs\n\n### HLP-001\nThe module MUST expose this item.\n")
	status, stdout, stderr := runSavedAttemptApp(t, root, "show", "help")
	if status != 0 || stdout != helpText(t, "show") || stderr != "" {
		t.Errorf("show help: exit=%d stdout=%q stderr=%q; want the show command page", status, stdout, stderr)
	}
	status, stdout, stderr = runSavedAttemptApp(t, root, "show", "help", "HLP-001")
	if status != 0 || stderr != "" || !strings.Contains(stdout, "The module MUST expose this item.") {
		t.Errorf("show help HLP-001: exit=%d stdout=%q stderr=%q; want the selected spec", status, stdout, stderr)
	}
}

func Test_APP_157_HelpJSONMatchesPlain(t *testing.T) {
	outside, invalid := helpFixture(t)
	for _, form := range helpForms {
		t.Run(strings.Join(form.args, " "), func(t *testing.T) {
			dir := outside
			if form.badConfig {
				dir = invalid
			}
			plainStatus, plain, plainErr := runSavedAttemptApp(t, dir, form.args...)
			args := append(append([]string(nil), form.args...), "--json")
			status, stdout, stderr := runSavedAttemptApp(t, dir, args...)
			if status != 0 || stderr != "" {
				t.Fatalf("JSON exit=%d stdout=%q stderr=%q; want exit 0 and empty stderr", status, stdout, stderr)
			}
			if got := commandStringData(t, stdout); plainStatus != 0 || plainErr != "" || got != plain {
				t.Errorf("JSON data=%q; want successful plain output, got exit=%d stdout=%q stderr=%q", got, plainStatus, plain, plainErr)
			}
		})
	}
}

func Test_APP_158_HelpFailures(t *testing.T) {
	outside, invalid := helpFixture(t)
	const unknown = "not-a-fireweed-command"
	validHelp := regexp.MustCompile(`(?i)\bfireweed (?:help\b|--help\b|-h\b)|\bhelp <command>|(?:<command>|doctor|agents-md|show) (?:help\b|--help\b|-h\b)|\b(?:use|try|run|see) [\x60'"]?help\b`)
	for _, tc := range []struct {
		dir, code string
		args      []string
		json      bool
	}{
		{outside, "U02", []string{"help", unknown}, false},
		{invalid, "U02", []string{unknown, "help"}, true},
		{invalid, "U02", []string{unknown, "--help"}, false},
		{outside, "U02", []string{unknown, "-h"}, true},
		{outside, "U01", []string{"help", "doctor", "extra"}, false},
		{invalid, "U01", []string{"-h", "doctor"}, true},
		{outside, "U01", []string{"show", "alpha", "-h"}, false},
		{invalid, "U01", []string{"doctor", "--help", "--no-agent"}, true},
		{outside, "U01", []string{"agents-md", "--write", "--help"}, false},
		{outside, "U01", []string{"help", "--bogus"}, true},
		{invalid, "U01", []string{"--help", "-h"}, false},
	} {
		t.Run(strings.Join(tc.args, " "), func(t *testing.T) {
			args := tc.args
			if tc.json {
				args = append(append([]string(nil), args...), "--json")
			}
			status, stdout, stderr := runSavedAttemptApp(t, tc.dir, args...)
			if status != 2 {
				t.Errorf("exit=%d stdout=%q stderr=%q; want exit 2", status, stdout, stderr)
			}
			message := stderr
			if tc.json {
				if stderr != "" {
					t.Errorf("JSON stderr=%q; want empty", stderr)
				}
				message = assertFailedJSON(t, stdout, tc.code)
			} else if !strings.HasPrefix(stderr, tc.code+":") {
				t.Errorf("stderr=%q; want %s diagnostic", stderr, tc.code)
			}
			if tc.code == "U02" {
				if !strings.Contains(message, unknown) || !strings.Contains(message, "fireweed help") {
					t.Errorf("diagnostic %q must name %q and direct the user to fireweed help", message, unknown)
				}
			} else if !corrective(message) || !validHelp.MatchString(message) {
				t.Errorf("diagnostic %q must show a valid help invocation", message)
			}
		})
	}
}
