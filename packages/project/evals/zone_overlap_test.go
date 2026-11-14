package evals_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/project"
)

func writeZoneSpec(t *testing.T, root, module, header, zones string) {
	t.Helper()
	put(t, root, "packages/"+module+"/SPEC.md", strings.Replace(header, "\n---\n", "\n"+zones+"---\n", 1))
}

func requireZoneOverlap(t *testing.T, root, module, evals, disposable string) {
	t.Helper()
	diagnostics := project.Check(root)
	found := false
	for _, d := range diagnostics {
		if !d.Warning && d.Code == "C02" {
			t.Errorf("valid dependency graph received C02: %#v", d)
		}
		if !d.Warning && d.Code == "C12" && strings.Contains(d.Message, module) &&
			strings.Contains(d.Message, evals) && strings.Contains(d.Message, disposable) {
			found = true
		}
	}
	if !found {
		t.Errorf("Check diagnostics = %#v; want error C12 naming module %q and globs %q and %q",
			diagnostics, module, evals, disposable)
	}
}

func requireNoZoneErrors(t *testing.T, root string) {
	t.Helper()
	for _, d := range project.Check(root) {
		if !d.Warning {
			t.Errorf("valid zones received an error diagnostic: %#v", d)
		}
	}
}

func Test_PRJ_075_CheckPossibleEvalDisposablePaths(t *testing.T) {
	for _, tc := range []struct {
		name, evals, disposable string
		overlap                 bool
	}{
		{"identical", "evals/**", "evals/**", true},
		{"recursive_and_suffix", "evals/**", "**/*_test.go", true},
		{"zero_recursive_components", "checks/**/smoke-*.go", "checks/smoke-test.*", true},
		{"disjoint_siblings", "evals/**", "evals-extra/**", false},
		{"single_component_wildcard", "checks/*/smoke.go", "checks/smoke.go", false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := fixture(t)
			// api.go and data/** already belong to the interface and preserve zones.
			zones := fmt.Sprintf("evals: [%q]\ndisposable: [%q, api.go, data/**]\n", tc.evals, tc.disposable)
			writeZoneSpec(t, root, "alpha", alphaSpecHeader, zones)
			if tc.overlap {
				requireZoneOverlap(t, root, "alpha", tc.evals, tc.disposable)
			} else {
				requireNoZoneErrors(t, root)
			}
		})
	}
}

func Test_PRJ_076_CheckResolvedZonesThroughTheQueue(t *testing.T) {
	const evals = "evals/**"
	const disposable = "**/*_test.go"
	overlappingDefaults := strings.Replace(fixtureConfig, "disposable: [internal/**]", `disposable: ["**/*_test.go"]`, 1)

	t.Run("current_overlap_cleared_by_first_change", func(t *testing.T) {
		root := fixture(t)
		put(t, root, "fireweed.yaml", overlappingDefaults)
		put(t, root, "packages/alpha/SPEC.md", alphaSpecHeader)
		writeZoneSpec(t, root, "beta", betaSpecHeader, "disposable: []\n")
		queued(t, root, "0001-clear.md", "set:\n  disposable: []\n", "Keep evals out of disposable code.\n", "", "", "")

		requireZoneOverlap(t, root, "alpha", evals, disposable)
	})

	t.Run("explicit_empty_lists_override_defaults", func(t *testing.T) {
		root := fixture(t)
		put(t, root, "fireweed.yaml", overlappingDefaults)
		writeZoneSpec(t, root, "alpha", alphaSpecHeader, "disposable: []\n")
		writeZoneSpec(t, root, "beta", betaSpecHeader, "evals: []\n")
		requireNoZoneErrors(t, root)

		queued(t, root, "0001-reassign.md", "set:\n  evals: []\n  disposable: [\"**/*_test.go\"]\n", "Move files to the disposable zone.\n", "", "", "")
		requireNoZoneErrors(t, root)
	})

	t.Run("set_edits_create_overlap_at_tip_and_intermediate_projection", func(t *testing.T) {
		root := fixture(t)
		writeZoneSpec(t, root, "alpha", alphaSpecHeader, "evals: []\n")
		queued(t, root, "0001-add-evals.md", "set:\n  evals: [checks/**]\n", "Declare the eval directory.\n", "", "", "")
		queued(t, root, "0002-overlap.md", "set:\n  disposable: [\"**/*_test.go\"]\n", "Declare disposable test files.\n", "", "", "")
		t.Run("tip", func(t *testing.T) {
			requireZoneOverlap(t, root, "alpha", "checks/**", disposable)
		})

		queued(t, root, "0003-clear.md", "set:\n  evals: []\n", "Clear the eval zone.\n", "", "", "")
		t.Run("intermediate", func(t *testing.T) {
			requireZoneOverlap(t, root, "alpha", "checks/**", disposable)
		})
	})
}
