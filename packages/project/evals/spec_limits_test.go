package evals_test

import (
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
)

func Test_PRJ_071_LoadSpecLimits(t *testing.T) {
	for _, tc := range []struct {
		name, config                           string
		maxClauses, maxWords, maxDecisionWords int
	}{
		{"omitted", "", 3, 80, 300},
		{"empty_mapping", "spec_limits: {}\n", 3, 80, 300},
		{"partial", "spec_limits:\n  max_words: 45\n", 3, 45, 300},
		{"configured", "spec_limits:\n  max_clauses: 1\n  max_words: 20\n  max_decision_words: 30\n", 1, 20, 30},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := fixture(t)
			put(t, root, "fireweed.yaml", fixtureConfig+tc.config)
			snapshot, err := project.Load(root)
			if err != nil {
				t.Fatalf("Load: %v", err)
			}
			limits, ok := snapshot.Config["spec_limits"].(map[string]any)
			if !ok {
				t.Fatalf("loaded spec_limits = %#v; want a mapping", snapshot.Config["spec_limits"])
			}
			for key, want := range map[string]int{
				"max_clauses": tc.maxClauses, "max_words": tc.maxWords, "max_decision_words": tc.maxDecisionWords,
			} {
				if got, ok := limits[key].(int); !ok || got != want {
					t.Errorf("loaded spec_limits.%s = %#v; want integer %d", key, limits[key], want)
				}
			}
		})
	}
}

func Test_PRJ_071_LoadRejectsInvalidSpecLimitsWithC01(t *testing.T) {
	for _, tc := range []struct{ name, config string }{
		{"scalar_section", "spec_limits: 3\n"},
		{"null_section", "spec_limits: null\n"},
		{"sequence_section", "spec_limits: [3, 80, 300]\n"},
		{"zero_clauses", "spec_limits: {max_clauses: 0}\n"},
		{"negative_words", "spec_limits: {max_words: -1}\n"},
		{"quoted_decision_words", `spec_limits: {max_decision_words: "300"}` + "\n"},
		{"floating_point_clauses", "spec_limits: {max_clauses: 3.0}\n"},
		{"boolean_words", "spec_limits: {max_words: true}\n"},
		{"null_decision_words", "spec_limits: {max_decision_words: null}\n"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := fixture(t)
			put(t, root, "fireweed.yaml", fixtureConfig+tc.config)
			_, err := project.Load(root)
			if code, _, ok := contracts.ErrorCode(err); !ok || code != "C01" {
				t.Errorf("Load error = %v; want C01 for %q", err, strings.TrimSpace(tc.config))
			}
		})
	}
}

const alphaSpecHeader = "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\n---\n# Alpha\n\n## Specs\n\n"
const betaSpecHeader = "---\nname: beta\nid_prefix: BET\ndepends_on: [alpha]\n---\n# Beta\n\n## Specs\n\n"

func specLimitWarnings(t *testing.T, root, code string) []project.Diagnostic {
	t.Helper()
	var warnings []project.Diagnostic
	for _, diagnostic := range project.Check(root) {
		if !diagnostic.Warning {
			t.Errorf("valid fixture received an error diagnostic: %#v", diagnostic)
		}
		if diagnostic.Code == code {
			warnings = append(warnings, diagnostic)
		}
	}
	return warnings
}

func requireNamedWarnings(t *testing.T, warnings []project.Diagnostic, names ...string) {
	t.Helper()
	if len(warnings) != len(names) {
		t.Errorf("diagnostics = %#v; want one warning each for %q", warnings, names)
	}
	for _, name := range names {
		count := 0
		for _, warning := range warnings {
			if warning.Warning && strings.Contains(warning.Message, name) {
				count++
			}
		}
		if count != 1 {
			t.Errorf("warning count naming %q = %d; want 1 in %#v", name, count, warnings)
		}
	}
}

func Test_PRJ_072_CheckWarnsForOversizedCurrentAndPendingItems(t *testing.T) {
	root := fixture(t)
	put(t, root, "fireweed.yaml", fixtureConfig+"spec_limits: {max_clauses: 2, max_words: 9}\n")
	put(t, root, "packages/alpha/SPEC.md", alphaSpecHeader+`### ALP-001
Alpha MUST read, MUST write, SHOULD retry.

### ALP-002
Alpha MUST return one entry for every source row
requested.

### ALP-003 [review]
Alpha MUST NOT discard data and SHOULD record
failures.

### ALP-004
Alpha MUST parse MUSTARD and SHOULDERS and preserve must.

### ALP-005
Alpha MUST save revisions.
`)
	t.Run("current", func(t *testing.T) {
		requireNamedWarnings(t, specLimitWarnings(t, root, "C17"), "ALP-001", "ALP-002")
	})

	queued(t, root, "0001-oversized.md", "", "Expand the behavior.\n",
		"### new:wide\nThe addition MUST report\tall available dates for each request.\n",
		"### ALP-004\nAlpha SHOULD parse, SHOULD validate, MUST reject.\n", "")
	queued(t, root, "0002-revise.md", "", "Narrow the behavior.\n", "",
		"### ALP-004\nAlpha MUST reject invalid input.\n",
		"### ALP-005\nThe removed item MUST read, MUST write, and SHOULD retry indefinitely.\n")
	for _, change := range []string{"oversized", "revise"} {
		if err := project.Stamp(root, change); err != nil {
			t.Fatalf("Stamp(%s): %v", change, err)
		}
	}
	requireNamedWarnings(t, specLimitWarnings(t, root, "C17"), "ALP-001", "ALP-002", "ALP-004", "new:wide")
}

func Test_PRJ_073_CheckWarnsForNormalizedSentencesAcrossProjectedItems(t *testing.T) {
	root := fixture(t)
	put(t, root, "packages/alpha/SPEC.md", alphaSpecHeader+`### ALP-001
Alpha MUST index rows.

### ALP-002
Alpha MUST sort rows.
`)
	put(t, root, "packages/beta/SPEC.md", betaSpecHeader+`### BET-001
Keep the cache-private. Beta MUST return a status.
`)
	queued(t, root, "0001-revise.md", "", "Revise the sorter.\n", "",
		"### ALP-002\nKEEP THE CACHE PRIVATE! Alpha MUST sort names.\n", "")
	if err := project.Stamp(root, "revise"); err != nil {
		t.Fatal(err)
	}
	queued(t, root, "0002-add.md", "", "Add a lookup.\n",
		"### new:lookup\nThe lookup MUST accept a name. Keep the\ncache, private?\n", "", "")
	requireRepeatedSentenceWarning(t, root, "ALP-002", "BET-001", "new:lookup")
}

func requireRepeatedSentenceWarning(t *testing.T, root string, names ...string) {
	t.Helper()
	warnings := specLimitWarnings(t, root, "C18")
	if len(warnings) != 1 || !warnings[0].Warning {
		t.Fatalf("C18 diagnostics = %#v; want one warning for the repeated sentence", warnings)
	}
	for _, name := range names {
		if !strings.Contains(warnings[0].Message, name) {
			t.Errorf("C18 warning %q does not name %s", warnings[0].Message, name)
		}
	}
}

func Test_PRJ_073_CheckCountsDistinctItemsInTheFinalProjection(t *testing.T) {
	root := fixture(t)
	put(t, root, "packages/alpha/SPEC.md", alphaSpecHeader+`### ALP-001
The cache MUST stay private. The cache MUST stay private. The cache MUST stay private.

### ALP-002
The cache MUST stay private.

### ALP-003
The cache MUST stay private.
`)
	put(t, root, "packages/beta/SPEC.md", betaSpecHeader+"### BET-001\nThe cache MUST stay private.\n")
	t.Run("empty_queue", func(t *testing.T) {
		requireRepeatedSentenceWarning(t, root, "ALP-001", "ALP-002", "ALP-003", "BET-001")
	})

	queued(t, root, "0001-add.md", "", "Add a lookup.\n",
		"### new:lookup\nThe lookup MUST accept a name.\n", "", "")
	queued(t, root, "0002-cleanup.md", "", "Drop repeated text.\n", "",
		"### ALP-002\nAlpha MUST sort names.\n", "### ALP-003\n")
	if err := project.Stamp(root, "cleanup"); err != nil {
		t.Fatal(err)
	}
	if warnings := specLimitWarnings(t, root, "C18"); len(warnings) != 0 {
		t.Errorf("C18 diagnostics = %#v; only two final items contain the sentence", warnings)
	}
}

func Test_PRJ_074_CheckWarnsOnlyForModulesOverTheDecisionsLimit(t *testing.T) {
	root := fixture(t)
	put(t, root, "fireweed.yaml", fixtureConfig+"spec_limits: {max_decision_words: 4}\n")
	put(t, root, "packages/alpha/SPEC.md", alphaSpecHeader+`### ALP-001
Alpha MUST write records.

## Decisions

Keep records in insertion
order.
`)
	put(t, root, "packages/beta/SPEC.md", betaSpecHeader+`### BET-001
Beta MUST compare all names in the catalog.

## Decisions

Use case insensitive comparisons.

## Notes

These notes appear after the decisions section.
`)
	requireNamedWarnings(t, specLimitWarnings(t, root, "C19"), "alpha")
}
