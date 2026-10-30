package evals_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/project"
)

func listRun(t *testing.T, dir string, args ...string) (int, string, string) {
	t.Helper()
	previous, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(dir); err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := os.Chdir(previous); err != nil {
			t.Fatalf("restore working directory: %v", err)
		}
	}()
	return run(args...)
}

func listChange(t *testing.T, root, filename, module, front string) {
	t.Helper()
	appPut(t, root, "changes/"+filename, "---\nmodule: "+module+"\n"+front+"---\n\n## Motivation\nUsers need this queued change.\n\n## Interface\n\n## Added\n\n## Modified\n\n## Removed\n\n## Decisions\n")
}

func listJSON(t *testing.T, output string) []map[string]json.RawMessage {
	t.Helper()
	envelope := decodeEnvelope(t, output)
	var diagnostics []json.RawMessage
	if err := json.Unmarshal(envelope["diagnostics"], &diagnostics); err != nil || diagnostics == nil || len(diagnostics) != 0 {
		t.Errorf("diagnostics=%s; want empty array: %v", envelope["diagnostics"], err)
	}
	var entries []map[string]json.RawMessage
	if err := json.Unmarshal(envelope["data"], &entries); err != nil || entries == nil {
		t.Fatalf("data=%s; want array: %v", envelope["data"], err)
	}
	return entries
}

func Test_APP_114_APP_115_APP_117_ListQueueOrderAndFields(t *testing.T) {
	root := appProjectFixture(t)
	listChange(t, root, "0002-zed.md", "alpha", "")
	listChange(t, root, "0001-first.md", "beta", "phase: evals\naudit_rounds: 1\n")
	listChange(t, root, "0002-ace.md", "alpha", "phase: audit\nblocked: review required\naudit_rounds: 2\n")
	want := []struct {
		slug, path, module, phase, blocked string
		rounds                             int
	}{
		{"first", "changes/0001-first.md", "beta", "evals", "", 1},
		{"ace", "changes/0002-ace.md", "alpha", "audit", "review required", 2},
		{"zed", "changes/0002-zed.md", "alpha", "queued", "", 0},
	}
	queued, err := project.Queue(root)
	if err != nil || len(queued) != len(want) {
		t.Fatalf("public Queue fixture: %v, entries=%v", err, queued)
	}
	for i, entry := range queued {
		if entry.Slug != want[i].slug || entry.Path != want[i].path || entry.Module != want[i].module {
			t.Fatalf("public Queue order[%d]=%+v", i, entry)
		}
	}
	status, plain, stderr := listRun(t, appNestedDir(t, root), "list")
	if status != 0 || stderr != "" {
		t.Fatalf("plain list exit=%d stdout=%q stderr=%q", status, plain, stderr)
	}
	previousPath := -1
	for _, entry := range want {
		for _, field := range []string{entry.slug, entry.path, entry.module, entry.phase} {
			if !strings.Contains(plain, field) {
				t.Errorf("plain list=%q missing %q", plain, field)
			}
		}
		if strings.Count(plain, entry.slug) < 2 {
			t.Errorf("plain list=%q must identify slug %q separately from its path", plain, entry.slug)
		}
		position := strings.Index(plain, entry.path)
		if position <= previousPath {
			t.Errorf("plain path %q at %d after previous at %d; output=%q", entry.path, position, previousPath, plain)
		}
		previousPath = position
		if entry.blocked != "" && !strings.Contains(plain, entry.blocked) {
			t.Errorf("plain list=%q missing blocked reason %q", plain, entry.blocked)
		}
		if entry.blocked == "" {
			found := false
			for _, row := range strings.Split(plain, "\n") {
				if strings.Contains(row, entry.path) && strings.Contains(row, `blocked=""`) {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("plain list=%q missing blocked=\"\" on row for %q", plain, entry.path)
			}
		}
		if !regexp.MustCompile(`\b` + strconv.Itoa(entry.rounds) + `\b`).MatchString(plain) {
			t.Errorf("plain list=%q missing audit rounds %d", plain, entry.rounds)
		}
	}
	status, output, stderr := listRun(t, root, "list", "--json")
	if status != 0 || stderr != "" {
		t.Fatalf("JSON list exit=%d stdout=%q stderr=%q", status, output, stderr)
	}
	entries := listJSON(t, output)
	if len(entries) != len(want) {
		t.Fatalf("JSON entries=%s; want %d", output, len(want))
	}
	for i, entry := range want {
		fields := entries[i]
		for key, expected := range map[string]any{
			"slug": entry.slug, "path": entry.path, "module": entry.module,
			"phase": entry.phase, "blocked": entry.blocked, "audit_rounds": float64(entry.rounds),
		} {
			var got any
			if err := json.Unmarshal(fields[key], &got); err != nil || !reflect.DeepEqual(got, expected) {
				t.Errorf("JSON entry[%d].%s=%s; want %v: %v", i, key, fields[key], expected, err)
			}
		}
	}
}

func Test_APP_114_APP_117_ListNearestRootAndEmptyQueue(t *testing.T) {
	outer := appProjectFixture(t)
	listChange(t, outer, "0001-outer.md", "alpha", "")
	inner := filepath.Join(outer, "workspace", "inner")
	appPut(t, inner, "fireweed.yaml", appFixtureConfig)
	appPut(t, inner, "packages/gamma/SPEC.md", "---\nname: gamma\nid_prefix: GAM\ndepends_on: []\n---\n# Gamma\n\n## Specs\n")
	deep := filepath.Join(inner, "packages", "gamma", "deep")
	if err := os.MkdirAll(deep, 0755); err != nil {
		t.Fatal(err)
	}
	status, plain, stderr := listRun(t, deep, "list")
	if status != 0 || plain != "" || stderr != "" {
		t.Errorf("nearest empty plain exit=%d stdout=%q stderr=%q", status, plain, stderr)
	}
	status, output, stderr := listRun(t, deep, "list", "--json")
	if status != 0 || stderr != "" {
		t.Fatalf("nearest empty JSON exit=%d stdout=%q stderr=%q", status, output, stderr)
	}
	if entries := listJSON(t, output); len(entries) != 0 {
		t.Errorf("nearest JSON data=%s; want empty array", output)
	}
	listChange(t, inner, "0001-inner.md", "gamma", "")
	status, output, stderr = listRun(t, deep, "list", "--json")
	if status != 0 || stderr != "" {
		t.Fatalf("nearest populated JSON exit=%d stdout=%q stderr=%q", status, output, stderr)
	}
	entries := listJSON(t, output)
	if len(entries) != 1 || string(entries[0]["slug"]) != `"inner"` || string(entries[0]["module"]) != `"gamma"` {
		t.Errorf("nearest JSON entries=%s; want inner gamma only", output)
	}
}

func listFailure(t *testing.T, root string, args []string, wantStatus int, code, cause string) {
	t.Helper()
	status, output, stderr := listRun(t, root, args...)
	if status != wantStatus || !strings.Contains(output+stderr, code+":") || !strings.Contains(output+stderr, cause) || !corrective(output+stderr) {
		t.Errorf("plain %v exit=%d stdout=%q stderr=%q; want %s cause %q and action", args, status, output, stderr, code, cause)
	}
	jsonArgs := append(append([]string(nil), args...), "--json")
	status, output, stderr = listRun(t, root, jsonArgs...)
	if status != wantStatus || stderr != "" {
		t.Errorf("JSON %v exit=%d stdout=%q stderr=%q; want %d and empty stderr", jsonArgs, status, output, stderr, wantStatus)
	}
	message := assertFailedJSON(t, output, code)
	if !strings.Contains(message, cause) || !corrective(message) {
		t.Errorf("JSON %v diagnostic=%q; want cause %q and action", jsonArgs, message, cause)
	}
}

func Test_APP_116_APP_117_ListUsageConfigurationAndQueueErrors(t *testing.T) {
	root := appProjectFixture(t)
	listFailure(t, root, []string{"list", "extra"}, 2, "U01", "extra")
	appPut(t, root, "fireweed.yaml", "modules: [broken\n")
	listFailure(t, root, []string{"list"}, 2, "C01", "fireweed.yaml")
	appPut(t, root, "fireweed.yaml", appFixtureConfig)
	listChange(t, root, "0001-repeat.md", "alpha", "")
	listChange(t, root, "0002-repeat.md", "alpha", "")
	_, cause := project.Queue(root)
	if cause == nil || !strings.Contains(cause.Error(), "ambiguous change slug") {
		t.Fatalf("public Queue fixture cause: %v", cause)
	}
	listFailure(t, root, []string{"list"}, 1, "C13", cause.Error())
	regularFile := appProjectFixture(t)
	appPut(t, regularFile, "changes", "not a directory")
	regularFile, err := filepath.EvalSymlinks(regularFile)
	if err != nil {
		t.Fatal(err)
	}
	_, cause = project.Queue(regularFile)
	if cause == nil {
		t.Fatal("public Queue accepted a regular file as changes directory")
	}
	hasC03 := false
	for _, diagnostic := range project.Check(regularFile) {
		if diagnostic.Code == "C03" && !diagnostic.Warning && strings.Contains(diagnostic.Path+diagnostic.Message, "changes") {
			hasC03 = true
			break
		}
	}
	if !hasC03 {
		t.Fatal("public Check did not report a C03 error for the regular-file changes path")
	}
	listFailure(t, regularFile, []string{"list"}, 1, "C03", cause.Error())
}
