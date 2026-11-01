package evals_test

import (
	"bytes"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
)

func checkFixFixture(t *testing.T) string {
	t.Helper()
	root := authoringFixture(t)
	appPut(t, root, "packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\n---\n# Alpha\n\n## Specs\n\n### ALP-001\nOne MUST work.\n\n### ALP-002\nTwo MUST work.\n\n### ALP-003\nThree MUST work.\n\n### ALP-004\nFour MUST work.\n")
	return root
}

func checkFixHash(t *testing.T, root, through, id string) string {
	t.Helper()
	var snapshot project.Snapshot
	var err error
	if through == "" {
		snapshot, err = project.Load(root)
	} else {
		snapshot, err = project.Project(root, through)
	}
	if err != nil {
		t.Fatal(err)
	}
	for _, module := range snapshot.Modules {
		if module.Name != "alpha" {
			continue
		}
		section, _ := contracts.Section(module.Document.Body, "Specs")
		items, err := contracts.ParseItems(section)
		if err != nil {
			t.Fatal(err)
		}
		for _, item := range items {
			if item.ID == id {
				return contracts.SpecHash(item)
			}
		}
	}
	t.Fatalf("fixture lacks %s in projection %q", id, through)
	return ""
}

func checkFixBases(t *testing.T, root, path string) (map[string]any, contracts.Document) {
	t.Helper()
	doc, err := contracts.ParseDocument(string(staleCLIText(t, root, path)))
	if err != nil {
		t.Fatal(err)
	}
	bases, ok := doc.Frontmatter["bases"].(map[string]any)
	if !ok {
		t.Fatalf("%s bases=%#v", path, doc.Frontmatter["bases"])
	}
	return bases, doc
}

func checkFixCLI(t *testing.T, bin, root string, jsonMode bool, wantStatus int, wantCode string) string {
	t.Helper()
	args := []string{"check", "--fix"}
	if jsonMode {
		args = append(args, "--json")
	}
	status, out, stderr := runAppBinary(t, bin, appNestedDir(t, root), args...)
	if status != wantStatus {
		t.Errorf("check --fix: exit=%d stdout=%q stderr=%q; want %d", status, out, stderr, wantStatus)
	}
	if wantStatus == 0 {
		if stderr != "" {
			t.Errorf("success stderr=%q", stderr)
		}
		if jsonMode {
			return commandStringData(t, out)
		}
		return out
	}
	if jsonMode {
		if stderr != "" {
			t.Errorf("JSON failure stderr=%q", stderr)
		}
		message := assertFailedJSON(t, out, wantCode)
		if !corrective(message) {
			t.Errorf("diagnostic lacks correction: %q", message)
		}
		return message
	}
	if out != "" {
		t.Errorf("plain failure stdout=%q; want empty", out)
	}
	if wantCode != "" && !strings.Contains(stderr, wantCode+":") || !corrective(stderr) {
		t.Errorf("plain failure stderr=%q; want corrective %s", stderr, wantCode)
	}
	return stderr
}

func Test_APP_147_APP_148_CheckFixAcrossQueuePreservesRecordedBases(t *testing.T) {
	bin := buildAppBinary(t)
	root := checkFixFixture(t)
	valid := checkFixHash(t, root, "", "ALP-002")
	staleCLIChange(t, root, "changes/0001-first.md", "custom: keep-first\nbases:\n  ALP-001: deadbeef\n  ALP-002: "+valid+"\n  ALP-999: extra-base\n", "Keep first human motivation.\n", "### ALP-001\nOne MUST change.\n\n### ALP-003\nThree MUST change.\n", "### ALP-002\n\n### ALP-004\n")
	staleCLIChange(t, root, "changes/0002-second.md", "custom: keep-second\n", "Keep second human motivation.\n", "### ALP-001\nOne MUST change again.\n", "### ALP-003\n")
	first, second := "changes/0001-first.md", "changes/0002-second.md"
	secondOne := checkFixHash(t, root, "first", "ALP-001")
	secondThree := checkFixHash(t, root, "first", "ALP-003")
	beforeFirst, beforeSecond := staleCLIText(t, root, first), staleCLIText(t, root, second)
	beforeSpec, beforeConfig := staleCLIText(t, root, "packages/alpha/SPEC.md"), staleCLIText(t, root, "fireweed.yaml")
	status, _, _ := runAppBinary(t, bin, appNestedDir(t, root), "check")
	if status != 1 || !bytes.Equal(beforeFirst, staleCLIText(t, root, first)) || !bytes.Equal(beforeSecond, staleCLIText(t, root, second)) {
		t.Error("ordinary check did not report missing bases read-only")
	}
	plain := checkFixCLI(t, bin, root, false, 1, "C08")
	if !strings.Contains(plain, "stale base for ALP-001") || strings.Contains(plain, "missing base for") {
		t.Errorf("post-fix diagnostics=%q; want only preserved stale conflict", plain)
	}
	bases, doc := checkFixBases(t, root, first)
	originalFirst, err := contracts.ParseDocument(string(beforeFirst))
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]any{
		"ALP-001": "deadbeef", "ALP-002": valid, "ALP-003": checkFixHash(t, root, "", "ALP-003"),
		"ALP-004": checkFixHash(t, root, "", "ALP-004"), "ALP-999": "extra-base",
	}
	if !reflect.DeepEqual(bases, want) || doc.Frontmatter["custom"] != "keep-first" || doc.Body != originalFirst.Body {
		t.Errorf("first change lost recorded bases or human content: bases=%v doc=%+v", bases, doc)
	}
	secondBases, secondDoc := checkFixBases(t, root, second)
	originalSecond, err := contracts.ParseDocument(string(beforeSecond))
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(secondBases, map[string]any{"ALP-001": secondOne, "ALP-003": secondThree}) || secondDoc.Frontmatter["custom"] != "keep-second" || secondDoc.Body != originalSecond.Body {
		t.Errorf("second change not stamped against preceding projection: bases=%v doc=%+v", secondBases, secondDoc)
	}
	if bytes.Equal(beforeFirst, staleCLIText(t, root, first)) || bytes.Equal(beforeSecond, staleCLIText(t, root, second)) {
		t.Error("fix did not change both missing-base queued files")
	}
	if !bytes.Equal(beforeSpec, staleCLIText(t, root, "packages/alpha/SPEC.md")) || !bytes.Equal(beforeConfig, staleCLIText(t, root, "fireweed.yaml")) {
		t.Error("fix changed nonqueue project files")
	}
	afterFirst, afterSecond := staleCLIText(t, root, first), staleCLIText(t, root, second)
	message := checkFixCLI(t, bin, root, true, 1, "C08")
	if !strings.Contains(message, "stale base for ALP-001") || !bytes.Equal(afterFirst, staleCLIText(t, root, first)) || !bytes.Equal(afterSecond, staleCLIText(t, root, second)) {
		t.Error("repeated fix changed files or lost stale C08")
	}
}

func Test_APP_147_APP_148_CheckFixMissingOnlySuccessAndNoOp(t *testing.T) {
	bin := buildAppBinary(t)
	root := checkFixFixture(t)
	staleCLIChange(t, root, "changes/0001-first.md", "custom: retain\n", "Keep human text.\n", "### ALP-001\nOne MUST change.\n", "### ALP-002\n")
	path := "changes/0001-first.md"
	before, beforeSpec := staleCLIText(t, root, path), staleCLIText(t, root, "packages/alpha/SPEC.md")
	plain := checkFixCLI(t, bin, root, false, 0, "")
	if !strings.Contains(plain, "Check passed") {
		t.Errorf("plain success=%q", plain)
	}
	bases, doc := checkFixBases(t, root, path)
	if !reflect.DeepEqual(bases, map[string]any{"ALP-001": checkFixHash(t, root, "", "ALP-001"), "ALP-002": checkFixHash(t, root, "", "ALP-002")}) || doc.Frontmatter["custom"] != "retain" || !strings.Contains(doc.Body, "Keep human text.") {
		t.Errorf("missing-only fix: bases=%v doc=%+v", bases, doc)
	}
	after := staleCLIText(t, root, path)
	if bytes.Equal(before, after) || !bytes.Equal(beforeSpec, staleCLIText(t, root, "packages/alpha/SPEC.md")) {
		t.Error("fix changed wrong files")
	}
	if got := checkFixCLI(t, bin, root, true, 0, ""); got != plain || !bytes.Equal(after, staleCLIText(t, root, path)) {
		t.Errorf("repeated fix changed output or bytes: %q", got)
	}
}

func Test_APP_148_APP_149_CheckFixRejectsBeforeMutation(t *testing.T) {
	bin := buildAppBinary(t)
	root := checkFixFixture(t)
	staleCLIChange(t, root, "changes/0001-first.md", "", "Needed.\n", "### ALP-001\nOne MUST change.\n", "")
	path := "changes/0001-first.md"
	before := staleCLIText(t, root, path)
	for _, args := range [][]string{{"check", "--fix", "extra"}, {"show", "alpha", "--fix"}, {"modules", "--fix"}, {"check", "--fix", "--bogus"}} {
		for _, jsonMode := range []bool{false, true} {
			call := append([]string(nil), args...)
			if jsonMode {
				call = append(call, "--json")
			}
			status, out, stderr := runAppBinary(t, bin, appNestedDir(t, root), call...)
			assertAppHistoryFailure(t, status, 2, out, stderr, jsonMode, "U01", "check --fix")
			if !bytes.Equal(before, staleCLIText(t, root, path)) {
				t.Error("invalid fix usage changed queued file")
			}
		}
	}
	staleCLIChange(t, root, path, "bases: [invalid]\n", "Needed.\n", "### ALP-001\nOne MUST change.\n", "")
	before = staleCLIText(t, root, path)
	for _, jsonMode := range []bool{false, true} {
		checkFixCLI(t, bin, root, jsonMode, 1, "C08")
		if !bytes.Equal(before, staleCLIText(t, root, path)) {
			t.Error("stamping rejection changed queued file")
		}
	}
	appPut(t, root, "fireweed.yaml", "modules: [broken\n")
	checkFixCLI(t, bin, root, true, 2, "C01")
	if !bytes.Equal(before, staleCLIText(t, root, path)) {
		t.Error("configuration rejection changed queued file")
	}
	if err := os.Remove(filepath.Join(root, "fireweed.yaml")); err != nil {
		t.Fatal(err)
	}
	checkFixCLI(t, bin, root, false, 2, "C01")
}
