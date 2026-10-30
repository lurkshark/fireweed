package evals_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/lurkshark/fireweed/packages/project"
)

func appHistoryArchive(t *testing.T, root, path, module, motivation, added, modified, decisions, removedDecisions, provenance string) {
	t.Helper()
	appPut(t, root, path, "---\nmodule: "+module+"\nphase: applied\n---\n\n## Motivation\n"+motivation+
		"## Added\n"+added+"## Modified\n"+modified+"## Removed\n\n## Decisions\n"+decisions+
		"## Removed decisions\n"+removedDecisions+"## Provenance\n"+provenance)
}

func appHistoryFixture(t *testing.T) string {
	t.Helper()
	root := authoringFixture(t)
	appHistoryArchive(t, root, "changes/archive/0001-first.md", "alpha", "Initial need.\n\n",
		"### ALP-002\nSecond behavior.\n### ALP-001\nFirst behavior.\n", "",
		"Original decision.\n\n", "Retired decision.\n\n",
		"- evals 2024-01-01T00:00:00Z agent=fixture tests=2\n- impl 2024-02-03T04:05:06Z agent=fixture commit=abc123\n")
	appHistoryArchive(t, root, "changes/archive/0002-revised.md", "alpha", "Revised need.\n\n", "",
		"### ALP-001 [review]\nRevised behavior.\n", "New decision.\n\n", "Superseded choice.\n\n",
		"- impl 2025-06-07T08:09:10Z agent=fixture commit=def456\n- done 2025-06-08T08:09:10Z agent=fixture\n")
	appHistoryArchive(t, root, "changes/archive/0003-undated.md", "alpha", "Undated need.\n\n", "",
		"### ALP-001\nUndated behavior.\n", "", "", "- evals 2026-01-01T00:00:00Z agent=fixture\n")
	return root
}

func appHistoryEntries(t *testing.T, root, id string, paths []string, ids [][]string) []project.SpecHistoryEntry {
	t.Helper()
	entries, err := project.SpecHistory(root, "alpha", id)
	if err != nil {
		t.Fatalf("public SpecHistory fixture: %v", err)
	}
	if len(entries) != len(paths) {
		t.Fatalf("public SpecHistory entries=%#v; want paths %v", entries, paths)
	}
	for i, entry := range entries {
		if entry.Path != paths[i] || entry.Module != "alpha" || !reflect.DeepEqual(entry.IDs, ids[i]) {
			t.Fatalf("public SpecHistory entry %d=%#v; want path %q and IDs %v", i, entry, paths[i], ids[i])
		}
	}
	return entries
}

func assertAppHistoryJSON(t *testing.T, output string, entries []project.SpecHistoryEntry) {
	t.Helper()
	envelope := decodeEnvelope(t, output)
	var diagnostics []json.RawMessage
	if err := json.Unmarshal(envelope["diagnostics"], &diagnostics); err != nil || diagnostics == nil || len(diagnostics) != 0 {
		t.Errorf("diagnostics=%s; want empty array", envelope["diagnostics"])
	}
	var rows []map[string]json.RawMessage
	if err := json.Unmarshal(envelope["data"], &rows); err != nil || rows == nil || len(rows) != len(entries) {
		t.Fatalf("data=%s; want %d history rows: %v", envelope["data"], len(entries), err)
	}
	for i, row := range rows {
		want := entries[i]
		for key := range row {
			if key != strings.ToLower(key) {
				t.Errorf("row %d has non-lowercase key %q", i, key)
			}
		}
		for key, value := range map[string]string{
			"path": want.Path, "module": want.Module, "motivation": want.Motivation,
			"decisions": want.Decisions, "removed_decisions": want.RemovedDecisions,
			"provenance": want.Provenance,
		} {
			var got string
			if err := json.Unmarshal(row[key], &got); err != nil || got != value {
				t.Errorf("row %d %s=%s; want %q: %v", i, key, row[key], value, err)
			}
		}
		var gotIDs []string
		if err := json.Unmarshal(row["ids"], &gotIDs); err != nil || !reflect.DeepEqual(gotIDs, want.IDs) {
			t.Errorf("row %d ids=%s; want %v: %v", i, row["ids"], want.IDs, err)
		}
		var date string
		if err := json.Unmarshal(row["implementation_date"], &date); err != nil {
			t.Errorf("row %d implementation_date=%s; want string: %v", i, row["implementation_date"], err)
		}
		if want.ImplementationDate.IsZero() {
			if date != "" {
				t.Errorf("row %d undated implementation_date=%q; want empty", i, date)
			}
		} else {
			parsed, err := time.Parse(time.RFC3339, date)
			if err != nil || !parsed.Equal(want.ImplementationDate) {
				t.Errorf("row %d implementation_date=%q; want RFC3339 %s: %v", i, date, want.ImplementationDate.Format(time.RFC3339), err)
			}
		}
	}
}

func Test_APP_132_APP_133_ShowHistoryOrderedMetadataAndFilter(t *testing.T) {
	bin := buildAppBinary(t)
	root := appHistoryFixture(t)
	nested := appNestedDir(t, root)
	for _, tc := range []struct {
		name  string
		id    string
		paths []string
		ids   [][]string
	}{
		{"selected-ID", "ALP-001", []string{"changes/archive/0002-revised.md", "changes/archive/0001-first.md", "changes/archive/0003-undated.md"}, [][]string{{"ALP-001"}, {"ALP-001"}, {"ALP-001"}}},
		{"module-wide", "", []string{"changes/archive/0002-revised.md", "changes/archive/0001-first.md", "changes/archive/0003-undated.md"}, [][]string{{"ALP-001"}, {"ALP-001", "ALP-002"}, {"ALP-001"}}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			entries := appHistoryEntries(t, root, tc.id, tc.paths, tc.ids)
			args := []string{"show", "alpha"}
			if tc.id != "" {
				args = append(args, tc.id)
			}
			args = append(args, "--history")
			status, plain, stderr := runAppBinary(t, bin, nested, args...)
			if status != 0 || stderr != "" {
				t.Fatalf("plain status=%d stdout=%q stderr=%q; want success", status, plain, stderr)
			}
			starts := make([]int, len(entries))
			position := 0
			for i, entry := range entries {
				at := strings.Index(plain[position:], entry.Path)
				if at < 0 {
					t.Errorf("plain history omits ordered path %q: %q", entry.Path, plain)
					continue
				}
				starts[i] = position + at
				position = starts[i] + len(entry.Path)
				if count := strings.Count(plain, entry.Path); count != 1 {
					t.Errorf("plain history repeats path %q %d times", entry.Path, count)
				}
			}
			for i, entry := range entries {
				end := len(plain)
				if i+1 < len(entries) {
					end = starts[i+1]
				}
				if starts[i] >= end {
					continue
				}
				block := plain[starts[i]:end]
				for _, value := range append([]string{entry.Module, entry.Motivation, entry.Decisions, entry.RemovedDecisions, entry.Provenance}, entry.IDs...) {
					if value = strings.TrimSpace(value); value != "" && !strings.Contains(block, value) {
						t.Errorf("plain history %q omits metadata %q: %q", entry.Path, value, block)
					}
				}
				for _, label := range []string{"motivation", "decisions", "removed decisions", "provenance", "implementation date"} {
					if !strings.Contains(strings.ToLower(block), label) {
						t.Errorf("plain history %q omits %q label: %q", entry.Path, label, block)
					}
				}
				if entry.ImplementationDate.IsZero() {
					if !strings.Contains(strings.ToLower(block), "undated") {
						t.Errorf("plain history %q must identify undated entry: %q", entry.Path, block)
					}
				} else if !strings.Contains(block, entry.ImplementationDate.Format(time.RFC3339)) {
					t.Errorf("plain history %q omits implementation date %s: %q", entry.Path, entry.ImplementationDate, block)
				}
			}
			status, output, stderr := runAppBinary(t, bin, nested, append(args, "--json")...)
			if status != 0 || stderr != "" {
				t.Fatalf("JSON status=%d stdout=%q stderr=%q; want success", status, output, stderr)
			}
			assertAppHistoryJSON(t, output, entries)
		})
	}
}

func Test_APP_132_APP_133_ShowHistoryNearestRootAndEmpty(t *testing.T) {
	bin := buildAppBinary(t)
	outer := appHistoryFixture(t)
	inner := filepath.Join(outer, "other-project")
	appPut(t, inner, "fireweed.yaml", appFixtureConfig)
	appPut(t, inner, "packages/alpha/SPEC.md", authoringSpec)
	appHistoryArchive(t, inner, "changes/archive/0007-inner.md", "alpha", "Inner project only.\n\n",
		"### ALP-001\nInner behavior.\n", "", "", "", "- impl 2026-01-02T03:04:05Z agent=fixture\n")
	nested := appNestedDir(t, inner)
	status, plain, stderr := runAppBinary(t, bin, nested, "show", "alpha", "ALP-001", "--history")
	if status != 0 || stderr != "" || !strings.Contains(plain, "changes/archive/0007-inner.md") || strings.Contains(plain, "changes/archive/0002-revised.md") {
		t.Errorf("nearest root plain: status=%d stdout=%q stderr=%q", status, plain, stderr)
	}
	status, output, stderr := runAppBinary(t, bin, nested, "show", "alpha", "ALP-001", "--history", "--json")
	if status != 0 || stderr != "" {
		t.Fatalf("nearest root JSON: status=%d stdout=%q stderr=%q", status, output, stderr)
	}
	entries := appHistoryEntries(t, inner, "ALP-001", []string{"changes/archive/0007-inner.md"}, [][]string{{"ALP-001"}})
	assertAppHistoryJSON(t, output, entries)

	for _, jsonMode := range []bool{false, true} {
		args := []string{"show", "alpha", "--history"}
		if jsonMode {
			args = append(args, "--json")
		}
		status, stdout, stderr := runAppBinary(t, bin, appNestedDir(t, appProjectFixture(t)), args...)
		if status != 0 || stderr != "" {
			t.Errorf("empty history: status=%d stdout=%q stderr=%q", status, stdout, stderr)
			continue
		}
		if jsonMode {
			assertAppHistoryJSON(t, stdout, nil)
		} else if !strings.Contains(strings.ToLower(stdout), "no history") {
			t.Errorf("plain empty history needs explicit no-history message: %q", stdout)
		}
	}
}

func assertAppHistoryFailure(t *testing.T, status, wantStatus int, stdout, stderr string, jsonMode bool, code, named string) {
	t.Helper()
	if status != wantStatus {
		t.Errorf("status=%d; want %d; stdout=%q stderr=%q", status, wantStatus, stdout, stderr)
	}
	message := stderr
	if jsonMode {
		if stderr != "" {
			t.Errorf("JSON stderr=%q; want empty", stderr)
		}
		message = assertFailedJSON(t, stdout, code)
	} else if stdout != "" || !strings.HasPrefix(strings.TrimSpace(stderr), code+":") {
		t.Errorf("plain failure stdout=%q stderr=%q; want empty stdout and %s diagnostic", stdout, stderr, code)
	}
	if named != "" && !strings.Contains(message, named) {
		t.Errorf("diagnostic %q must name %q", message, named)
	}
	if !corrective(message) {
		t.Errorf("diagnostic %q needs corrective guidance", message)
	}
}

func Test_APP_134_ShowHistoryOptionMisuse(t *testing.T) {
	bin := buildAppBinary(t)
	root := authoringFixture(t)
	for _, tc := range []struct {
		name  string
		args  []string
		named string
	}{
		{"other-command", []string{"modules", "--history"}, "--history"},
		{"projected-combination", []string{"show", "alpha", "--history", "--projected"}, "--projected"},
	} {
		for _, jsonMode := range []bool{false, true} {
			args := append([]string(nil), tc.args...)
			if jsonMode {
				args = append(args, "--json")
			}
			t.Run(tc.name+map[bool]string{false: "/plain", true: "/JSON"}[jsonMode], func(t *testing.T) {
				status, stdout, stderr := runAppBinary(t, bin, appNestedDir(t, root), args...)
				assertAppHistoryFailure(t, status, 2, stdout, stderr, jsonMode, "U01", tc.named)
			})
		}
	}
}

func Test_APP_135_ShowHistoryLookupAndArchiveFailures(t *testing.T) {
	bin := buildAppBinary(t)
	for _, tc := range []struct {
		name   string
		args   []string
		code   string
		exit   int
		named  string
		modify func(*testing.T, string)
	}{
		{"unknown-module", []string{"show", "absent-module", "--history"}, "U02", 2, "absent-module", nil},
		{"unknown-ID", []string{"show", "alpha", "ALP-999", "--history"}, "U02", 2, "ALP-999", nil},
		{"missing-config", []string{"show", "alpha", "--history"}, "C01", 2, "", func(t *testing.T, root string) {
			if err := os.Remove(filepath.Join(root, "fireweed.yaml")); err != nil {
				t.Fatal(err)
			}
		}},
		{"malformed-config", []string{"show", "alpha", "--history"}, "C01", 2, "", func(t *testing.T, root string) {
			appPut(t, root, "fireweed.yaml", "modules: [broken\n")
		}},
		{"malformed-archive", []string{"show", "alpha", "--history"}, "C03", 1, "changes/archive/0001-invalid.md", func(t *testing.T, root string) {
			appPut(t, root, "changes/archive/0001-invalid.md", "---\nmodule: alpha\n---\n## Added\n### ALP-001 [broken\nBehavior.\n")
		}},
	} {
		for _, jsonMode := range []bool{false, true} {
			args := append([]string(nil), tc.args...)
			if jsonMode {
				args = append(args, "--json")
			}
			t.Run(tc.name+map[bool]string{false: "/plain", true: "/JSON"}[jsonMode], func(t *testing.T) {
				root := authoringFixture(t)
				if tc.modify != nil {
					tc.modify(t, root)
				}
				status, stdout, stderr := runAppBinary(t, bin, appNestedDir(t, root), args...)
				assertAppHistoryFailure(t, status, tc.exit, stdout, stderr, jsonMode, tc.code, tc.named)
			})
		}
	}
}
