package evals_test

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/assets"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
)

func binaryVersionParts(t *testing.T) (major string, minor, patch int) {
	t.Helper()
	version := assets.Version()
	parts := strings.Split(version, ".")
	if len(parts) != 3 {
		t.Fatalf("fixture expects three version components in %q", version)
	}
	minor, minorErr := strconv.Atoi(parts[1])
	patch, patchErr := strconv.Atoi(parts[2])
	if minorErr != nil || patchErr != nil {
		t.Fatalf("fixture expects decimal minor and patch components in %q", version)
	}
	return parts[0], minor, patch
}

func Test_PRJ_082_CheckCompatibilityAcceptsMatchingPins(t *testing.T) {
	version := assets.Version()
	major, minor, patch := binaryVersionParts(t)
	minorVersion := fmt.Sprintf("%s.%d", major, minor)
	nextMinor := fmt.Sprintf("%s.%d", major, minor+1)
	nextPatch := fmt.Sprintf("%s.%d.%d", major, minor, patch+1)
	for _, tc := range []struct{ name, pin string }{
		{"init_range", ">=" + minorVersion + ", <" + nextMinor},
		{"full_equality", "=" + version},
		{"strict_lower_bound", ">0.0"},
		{"inclusive_bounds", ">=" + version + ", <=" + version},
		{"strict_upper_patch", "<" + nextPatch},
		{"numeric_order_with_leading_zero", fmt.Sprintf("<%s.0%d", major, minor+1)},
		{"large_decimal_component", fmt.Sprintf("<%s.999999999999999999999999999999", major)},
		{"surrounding_whitespace", " \t>= " + minorVersion + " ,\t < " + nextMinor + " \n"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if err := project.CheckCompatibility(map[string]any{"fireweed": tc.pin}); err != nil {
				t.Errorf("CheckCompatibility(%q) rejected %s: %v", tc.pin, version, err)
			}
		})
	}
	t.Run("omitted", func(t *testing.T) {
		if err := project.CheckCompatibility(map[string]any{"version": 1}); err != nil {
			t.Errorf("CheckCompatibility rejected configuration with no pin: %v", err)
		}
	})
	t.Run("omitted_patch_is_zero", func(t *testing.T) {
		err := project.CheckCompatibility(map[string]any{"fireweed": "=" + minorVersion})
		if (err == nil) != (patch == 0) {
			t.Errorf("CheckCompatibility(=%s) for %s returned %v; want equality only when patch is zero", minorVersion, version, err)
		}
	})
}

func Test_PRJ_083_CheckCompatibilityReportsInvalidAndIncompatiblePins(t *testing.T) {
	version := assets.Version()
	major, minor, patch := binaryVersionParts(t)
	invalid := regexp.MustCompile(`(?i)invalid|syntax|malformed|must be a .*string|expected|parse|empty|unrecognized|unsupported`)
	incompatible := regexp.MustCompile(`(?i)incompatib|exclud|outside|out of range|does not (match|satisfy|meet|support)|requires|not (compatible|in|within|satisfied|supported)`)
	for _, tc := range []struct {
		name string
		pin  any
		want *regexp.Regexp
	}{
		{"empty", "", invalid},
		{"nonstring", 42, invalid},
		{"double_equals", "==" + version, invalid},
		{"missing_operator", version, invalid},
		{"missing_minor", "=" + major, invalid},
		{"extra_component", "=" + version + ".0", invalid},
		{"negative_component", ">=" + major + ".-1", invalid},
		{"prerelease", "=" + version + "-rc.1", invalid},
		{"empty_comparison", "=" + version + ",", invalid},
		{"invalid_after_exclusion", ">" + version + ", bad", invalid},
		{"strict_upper_endpoint", "<" + version, incompatible},
		{"strict_lower_endpoint", ">" + version, incompatible},
		{"second_comparison_excludes", fmt.Sprintf(">=%s, =%s.%d.%d", version, major, minor, patch+1), incompatible},
		{"documented_range", ">=0.4, <0.5", incompatible},
	} {
		t.Run(tc.name, func(t *testing.T) {
			err := project.CheckCompatibility(map[string]any{"fireweed": tc.pin})
			if code, _, ok := contracts.ErrorCode(err); !ok || code != "C01" {
				t.Fatalf("CheckCompatibility(%#v) = %v; want a C01 CodedError", tc.pin, err)
			}
			message := err.Error()
			pinText := fmt.Sprint(tc.pin)
			identifiesPin := strings.Contains(message, pinText)
			if pinText == "" {
				identifiesPin = strings.Contains(strings.ToLower(message), "fireweed") || strings.Contains(message, `""`)
			}
			if !identifiesPin || !strings.Contains(message, version) {
				t.Errorf("error %q must identify pin %#v and running version %q", message, tc.pin, version)
			}
			if !tc.want.MatchString(message) {
				t.Errorf("error %q does not explain the reason for rejecting pin %#v", message, tc.pin)
			}
		})
	}
}

func Test_PRJ_084_LoadAndCheckUseLiveBinaryPin(t *testing.T) {
	version := assets.Version()
	for _, tc := range []struct {
		name, pin string
		wantError bool
	}{
		{"omitted", "", false},
		{"compatible", fmt.Sprintf("fireweed: %q\n", "="+version), false},
		{"incompatible", fmt.Sprintf("fireweed: %q\n", ">"+version), true},
		{"malformed", fmt.Sprintf("fireweed: %q\n", "~"+version), true},
		{"null", "fireweed: null\n", true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := fixture(t)
			put(t, root, "packages/alpha/SPEC.md", alphaSpecHeader+"### ALP-001\nAlpha MUST work.\n")
			put(t, root, "fireweed.yaml", tc.pin+"modules: [packages/*]\n")
			_, err := project.Load(root)
			if tc.wantError {
				if code, _, ok := contracts.ErrorCode(err); !ok || code != "C01" {
					t.Errorf("Load with %q = %v; want C01", tc.pin, err)
				}
			} else if err != nil {
				t.Errorf("Load with %q rejected valid project: %v", tc.pin, err)
			}
			diagnostics := project.Check(root)
			found := false
			for _, d := range diagnostics {
				if d.Warning {
					continue
				}
				if !tc.wantError {
					t.Errorf("Check with %q reported unexpected error: %#v", tc.pin, d)
				}
				if d.Code == "C01" && d.Path == "fireweed.yaml" && strings.Contains(d.Message, version) {
					found = true
				}
			}
			if tc.wantError && !found {
				t.Errorf("Check with %q = %#v; want non-warning C01 at fireweed.yaml naming version %s", tc.pin, diagnostics, version)
			}
		})
	}
}

func Test_PRJ_047_LoadAtReadsRecordedBinaryPin(t *testing.T) {
	root := fixture(t)
	revisionGit(t, root, "init", "-q")
	pin := ">" + assets.Version()
	put(t, root, "fireweed.yaml", fmt.Sprintf("fireweed: %q\nmodules: [packages/*]\n", pin))
	recorded := revisionCommit(t, root, "record project with a newer binary pin")
	put(t, root, "fireweed.yaml", "modules: [broken\n")
	snapshot, err := project.LoadAt(root, recorded)
	if err != nil {
		t.Fatalf("LoadAt rejected the recorded project: %v", err)
	}
	if got := snapshot.Config["fireweed"]; got != pin {
		t.Errorf("LoadAt fireweed pin = %#v; want %q", got, pin)
	}
}
