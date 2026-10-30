package evals_test

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/lurkshark/fireweed/packages/project"
)

func historyArchive(t *testing.T, root, path, module, motivation, added, modified, removed, decisions, removedDecisions, provenance string) {
	t.Helper()
	put(t, root, path, fmt.Sprintf("---\nmodule: %s\nphase: applied\n---\n\n## Motivation\n%s## Added\n%s## Modified\n%s## Removed\n%s## Decisions\n%s## Removed decisions\n%s## Provenance\n%s", module, motivation, added, modified, removed, decisions, removedDecisions, provenance))
}

func historyProject(t *testing.T) string {
	t.Helper()
	root := fixture(t)
	revisionGit(t, root, "init", "-q")
	historyArchive(t, root, "changes/archive/0009-added.md", "alpha",
		"Why the behavior exists.\nSecond motivation line.\n\n",
		"### ALP-004\nAnother behavior.\n### ALP-002\nRetained behavior.\n", "", "",
		"Choose the stable form.\nKeep this second decision.\n\n",
		"Prior choice was dropped.\n\n",
		"- evals 2026-01-01T00:00:00Z agent=fixture tests=2\n- impl 2023-01-01T00:00:00Z agent=fixture\n- impl not-a-date agent=fixture\n- impl 2026-06-01T12:30:00Z agent=fixture\n")
	revisionCommit(t, root, "archive newest implementation date first")
	historyArchive(t, root, "changes/archive/0001-retag.md", "alpha", "Retag the behavior.\n", "",
		"### ALP-002 [review]\nRetained behavior.\n", "", "", "", "- impl 2024-01-01T00:00:00Z agent=fixture\n")
	historyArchive(t, root, "changes/archive/0003-equal.md", "alpha", "First equal date.\n", "",
		"### ALP-002\nEqual date revision.\n", "", "", "", "- impl 2025-01-01T00:00:00Z agent=fixture\n")
	historyArchive(t, root, "changes/archive/0004-equal.md", "alpha", "Second equal date.\n", "",
		"### ALP-002\nAnother equal date revision.\n", "", "", "", "- impl 2025-01-01T00:00:00Z agent=fixture\n")
	historyArchive(t, root, "changes/archive/0005-bad-date.md", "alpha", "Invalid date.\n", "",
		"### ALP-002\nUndated revision.\n", "", "", "", "- impl not-a-date agent=fixture\n")
	historyArchive(t, root, "changes/archive/0006-no-impl.md", "alpha", "No implementation line.\n", "",
		"### ALP-002\nUndated revision.\n", "", "", "", "- evals 2030-01-01T00:00:00Z agent=fixture\n")
	historyArchive(t, root, "changes/archive/0010-other-id.md", "alpha", "Other specs.\n",
		"### ALP-005\nFifth behavior.\n### ALP-003\nThird behavior.\n",
		"### ALP-003 [review]\nThird behavior retagged.\n", "", "", "", "")
	historyArchive(t, root, "changes/archive/0011-removed-only.md", "alpha", "Only a removal.\n", "", "",
		"### ALP-002\n", "", "", "- impl 2031-01-01T00:00:00Z agent=fixture\n")
	historyArchive(t, root, "changes/archive/0012-beta.md", "beta", "Beta history.\n",
		"### BET-001\nBeta behavior.\n", "", "", "", "", "- impl 2032-01-01T00:00:00Z agent=fixture\n")
	queued(t, root, "0013-pending.md", "", "Not archived.\n", "### ALP-002\nPending behavior.\n", "", "")
	revisionCommit(t, root, "archive older dates and pending change later")
	return root
}

func historyPaths(entries []project.SpecHistoryEntry) []string {
	paths := make([]string, len(entries))
	for i, entry := range entries {
		paths[i] = entry.Path
	}
	return paths
}

func Test_PRJ_051_SpecHistorySelectsArchivedAddsAndTagOnlyModifications(t *testing.T) {
	root := historyProject(t)
	entries, err := project.SpecHistory(root, "alpha", "ALP-002")
	if err != nil {
		t.Fatalf("SpecHistory: %v", err)
	}
	want := map[string]bool{
		"changes/archive/0009-added.md":    true,
		"changes/archive/0001-retag.md":    true,
		"changes/archive/0003-equal.md":    true,
		"changes/archive/0004-equal.md":    true,
		"changes/archive/0005-bad-date.md": true,
		"changes/archive/0006-no-impl.md":  true,
	}
	if len(entries) != len(want) {
		t.Fatalf("selected archives = %#v; want paths %#v", historyPaths(entries), want)
	}
	for _, entry := range entries {
		if !want[entry.Path] || entry.Module != "alpha" || !reflect.DeepEqual(entry.IDs, []string{"ALP-002"}) {
			t.Errorf("unexpected path/module/filtered IDs = %q, %q, %#v", entry.Path, entry.Module, entry.IDs)
		}
		delete(want, entry.Path)
	}
	if len(want) != 0 {
		t.Errorf("missing archives: %#v", want)
	}
}

func Test_PRJ_053_SpecHistoryOrdersByLatestValidImplementationDate(t *testing.T) {
	root := historyProject(t)
	entries, err := project.SpecHistory(root, "alpha", "ALP-002")
	if err != nil {
		t.Fatalf("SpecHistory: %v", err)
	}
	wantPaths := []string{
		"changes/archive/0009-added.md", "changes/archive/0003-equal.md",
		"changes/archive/0004-equal.md", "changes/archive/0001-retag.md",
		"changes/archive/0005-bad-date.md", "changes/archive/0006-no-impl.md",
	}
	if got := historyPaths(entries); !reflect.DeepEqual(got, wantPaths) {
		t.Fatalf("archive order/selection = %#v; want %#v", got, wantPaths)
	}
	wantDates := []time.Time{
		time.Date(2026, 6, 1, 12, 30, 0, 0, time.UTC),
		time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		{}, {},
	}
	for i, entry := range entries {
		if !entry.ImplementationDate.Equal(wantDates[i]) {
			t.Errorf("%s implementation date = %s; want %s", entry.Path, entry.ImplementationDate, wantDates[i])
		}
	}
}

func Test_PRJ_052_SpecHistoryPreservesCompleteEntryMetadata(t *testing.T) {
	root := historyProject(t)
	entries, err := project.SpecHistory(root, "alpha", "ALP-002")
	if err != nil {
		t.Fatalf("SpecHistory: %v", err)
	}
	if len(entries) == 0 {
		t.Fatal("missing archived ALP-002 history")
	}
	want := project.SpecHistoryEntry{
		Path: "changes/archive/0009-added.md", Module: "alpha", IDs: []string{"ALP-002"},
		Motivation:         "Why the behavior exists.\nSecond motivation line.\n\n",
		Decisions:          "Choose the stable form.\nKeep this second decision.\n\n",
		RemovedDecisions:   "Prior choice was dropped.\n\n",
		Provenance:         "- evals 2026-01-01T00:00:00Z agent=fixture tests=2\n- impl 2023-01-01T00:00:00Z agent=fixture\n- impl not-a-date agent=fixture\n- impl 2026-06-01T12:30:00Z agent=fixture\n",
		ImplementationDate: time.Date(2026, 6, 1, 12, 30, 0, 0, time.UTC),
	}
	if !reflect.DeepEqual(entries[0], want) {
		t.Errorf("entry = %#v; want %#v", entries[0], want)
	}
}

func Test_PRJ_054_SpecHistoryModuleWideUniqueSortedIDsAndEmptyModule(t *testing.T) {
	root := historyProject(t)
	entries, err := project.SpecHistory(root, "alpha", "")
	if err != nil {
		t.Fatalf("SpecHistory module-wide: %v", err)
	}
	wantPaths := []string{
		"changes/archive/0009-added.md", "changes/archive/0003-equal.md",
		"changes/archive/0004-equal.md", "changes/archive/0001-retag.md",
		"changes/archive/0005-bad-date.md", "changes/archive/0006-no-impl.md",
		"changes/archive/0010-other-id.md",
	}
	if got := historyPaths(entries); !reflect.DeepEqual(got, wantPaths) {
		t.Fatalf("module-wide archive order/selection = %#v; want %#v", got, wantPaths)
	}
	for _, entry := range entries {
		wantIDs := []string{"ALP-002"}
		switch entry.Path {
		case "changes/archive/0009-added.md":
			wantIDs = []string{"ALP-002", "ALP-004"}
		case "changes/archive/0010-other-id.md":
			wantIDs = []string{"ALP-003", "ALP-005"}
		}
		if entry.Module != "alpha" || !reflect.DeepEqual(entry.IDs, wantIDs) {
			t.Errorf("%s module/unique IDs = %q, %#v; want alpha, %#v", entry.Path, entry.Module, entry.IDs, wantIDs)
		}
	}
	emptyRoot := fixture(t)
	empty, err := project.SpecHistory(emptyRoot, "beta", "")
	if err != nil || len(empty) != 0 {
		t.Errorf("module without archives = %#v, %v; want empty slice", empty, err)
	}
}

func Test_PRJ_055_SpecHistoryLookupNamesAndRemovedCurrentSpec(t *testing.T) {
	root := historyProject(t)
	archived, err := project.SpecHistory(root, "alpha", "ALP-002")
	if err != nil || len(archived) == 0 {
		t.Errorf("archived ID absent from current Specs = %#v, %v", archived, err)
	}
	currentOnly, err := project.SpecHistory(root, "alpha", "ALP-001")
	if err != nil || len(currentOnly) != 0 {
		t.Errorf("current ID with no archive edits = %#v, %v; want empty slice", currentOnly, err)
	}
	for _, tc := range []struct{ module, id, rejected string }{
		{"absent-module", "", "absent-module"},
		{"alpha", "ALP-999", "ALP-999"},
		{"alpha", "BET-001", "BET-001"},
	} {
		_, err := project.SpecHistory(root, tc.module, tc.id)
		if err == nil || !strings.HasPrefix(err.Error(), "U02") || !strings.Contains(err.Error(), tc.rejected) {
			t.Errorf("lookup module %q ID %q = %v; want U02 naming %q", tc.module, tc.id, err, tc.rejected)
		}
	}
}

func Test_PRJ_056_SpecHistoryConfigurationAndArchiveSyntaxErrors(t *testing.T) {
	for _, tc := range []struct{ name, config string }{
		{"missing", ""},
		{"malformed", "modules: [broken\n"},
	} {
		t.Run("config_"+tc.name, func(t *testing.T) {
			root := fixture(t)
			if tc.name == "missing" {
				if err := os.Remove(filepath.Join(root, "fireweed.yaml")); err != nil {
					t.Fatal(err)
				}
			} else {
				put(t, root, "fireweed.yaml", tc.config)
			}
			_, err := project.SpecHistory(root, "alpha", "")
			if err == nil || !strings.HasPrefix(err.Error(), "C01") {
				t.Errorf("invalid config = %v; want C01", err)
			}
		})
	}
	for _, tc := range []struct{ name, document string }{
		{"document", "---\nmodule: [broken\n---\n## Added\n### ALP-002\nBehavior.\n"},
		{"added", "---\nmodule: alpha\n---\n## Added\n### ALP-002 [broken\nBehavior.\n"},
		{"modified", "---\nmodule: alpha\n---\n## Modified\n### ALP-002 [broken\nBehavior.\n"},
	} {
		t.Run("archive_"+tc.name, func(t *testing.T) {
			root := fixture(t)
			path := "changes/archive/0001-invalid.md"
			put(t, root, path, tc.document)
			_, err := project.SpecHistory(root, "alpha", "ALP-002")
			if err == nil || !strings.HasPrefix(err.Error(), "C03") || !strings.Contains(err.Error(), path) {
				t.Errorf("malformed archive = %v; want C03 naming %s", err, path)
			}
		})
	}
}
