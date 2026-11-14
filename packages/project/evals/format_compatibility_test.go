package evals_test

import (
	"regexp"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/assets"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
)

var supportedProjectFormat = regexp.MustCompile(`(?i)\b(?:supported|supports?|accepted|accepts?|expected|expects?|only|current|use|want)\b[^0-9\n]*\b1\b|\b1\b\s+is\s+(?:the\s+)?(?:supported|accepted|expected|only|current)\b`)

func requireFormatRejection(t *testing.T, message, rejected string) {
	t.Helper()
	supported := supportedProjectFormat.FindStringIndex(message)
	if supported == nil {
		t.Errorf("message %q must identify format 1 as supported", message)
		return
	}
	value := regexp.MustCompile(rejected)
	if !value.MatchString(message[:supported[0]]) && !value.MatchString(message[supported[1]:]) {
		t.Errorf("message %q must identify the rejected format value matching %q", message, rejected)
	}
}

func Test_PRJ_085_CheckCompatibilityProjectFormat(t *testing.T) {
	matchingPin := "=" + assets.Version()
	excludingPin := ">" + assets.Version()
	for _, tc := range []struct {
		name     string
		config   map[string]any
		rejected string
	}{
		{"explicit_initial", map[string]any{"version": 1}, ""},
		{"omitted", map[string]any{}, ""},
		{"compatible_pin_with_initial", map[string]any{"version": 1, "fireweed": matchingPin}, ""},
		{"future", map[string]any{"version": 999}, `\b999\b`},
		{"compatible_pin_with_future", map[string]any{"version": 999, "fireweed": matchingPin}, `\b999\b`},
		{"zero", map[string]any{"version": 0}, `\b0\b`},
		{"string", map[string]any{"version": "1"}, `\b1\b`},
		{"null", map[string]any{"version": nil}, `(?i)\b(?:null|nil)\b`},
		{"float", map[string]any{"version": 1.0}, `\b1\b`},
		{"boolean", map[string]any{"version": true}, `\btrue\b`},
		{"sequence", map[string]any{"version": []any{"preview"}}, `\bpreview\b`},
		{"mapping", map[string]any{"version": map[string]any{"format": "experimental"}}, `\bexperimental\b`},
	} {
		t.Run(tc.name, func(t *testing.T) {
			err := project.CheckCompatibility(tc.config)
			if tc.rejected == "" {
				if err != nil {
					t.Errorf("CheckCompatibility(%#v) = %v; want success", tc.config, err)
				}
				return
			}
			if code, _, ok := contracts.ErrorCode(err); !ok || code != "C01" {
				t.Fatalf("CheckCompatibility(%#v) = %v; want a C01 CodedError", tc.config, err)
			}
			requireFormatRejection(t, err.Error(), tc.rejected)
		})
	}

	t.Run("excluding_pin_with_initial", func(t *testing.T) {
		err := project.CheckCompatibility(map[string]any{"version": 1, "fireweed": excludingPin})
		if code, _, ok := contracts.ErrorCode(err); !ok || code != "C01" || !strings.Contains(err.Error(), excludingPin) {
			t.Errorf("CheckCompatibility for format 1 and pin %q = %v; want C01 identifying the excluding pin", excludingPin, err)
		}
	})
}

func Test_PRJ_085_LoadAndCheckRejectUnsupportedLiveFormat(t *testing.T) {
	root := t.TempDir()
	put(t, root, "fireweed.yaml", "version: 999\nmodules: []\n")
	_, err := project.Load(root)
	if code, _, ok := contracts.ErrorCode(err); !ok || code != "C01" {
		t.Errorf("Load = %v; want C01 for format 999", err)
	} else {
		requireFormatRejection(t, err.Error(), `\b999\b`)
	}
	for _, d := range project.Check(root) {
		if !d.Warning && d.Code == "C01" && d.Path == "fireweed.yaml" {
			requireFormatRejection(t, d.Message, `\b999\b`)
			return
		}
	}
	t.Error("Check must report a non-warning C01 at fireweed.yaml for format 999")
}

func Test_PRJ_047_LoadAtReadsUnsupportedRecordedFormat(t *testing.T) {
	root := t.TempDir()
	revisionGit(t, root, "init", "-q")
	put(t, root, "fireweed.yaml", "version: 999\nmodules: []\n")
	recorded := revisionCommit(t, root, "record a future project format")
	put(t, root, "fireweed.yaml", "version: 1\nmodules: []\n")
	snapshot, err := project.LoadAt(root, recorded)
	if err != nil {
		t.Fatalf("LoadAt rejected the recorded project: %v", err)
	}
	if got := snapshot.Config["version"]; got != 999 {
		t.Errorf("LoadAt format version = %#v; want 999", got)
	}
}
