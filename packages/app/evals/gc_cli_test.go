package evals_test

import (
	"encoding/json"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/gitstore"
)

func gcCLIRef(t *testing.T, root, name, oid string) {
	t.Helper()
	if err := gitstore.CASRef(root, name, "", oid); err != nil {
		t.Fatalf("create %s: %v", name, err)
	}
}

func gcCLIClaim(t *testing.T, root, name, worktree string) {
	t.Helper()
	data, err := json.Marshal(map[string]string{"worktree": worktree, "change": name})
	if err != nil {
		t.Fatal(err)
	}
	oid, err := gitstore.WriteBlob(root, data)
	if err != nil {
		t.Fatal(err)
	}
	gcCLIRef(t, root, "refs/fireweed/claims/"+name, oid)
}

func gcCLIRefs(t *testing.T, root string) []string {
	t.Helper()
	out := cliGit(t, root, "for-each-ref", "--format=%(refname)", "refs/fireweed")
	if out == "" {
		return []string{}
	}
	return strings.Split(out, "\n")
}

func gcCLIProject(t *testing.T) (root, outer, nested string) {
	t.Helper()
	outer = appProjectFixture(t)
	cliGit(t, outer, "init", "-q", "-b", "main")
	cliCommit(t, outer, "outer project")
	root = filepath.Join(outer, "inner-project")
	appPut(t, root, "fireweed.yaml", appFixtureConfig)
	appPut(t, root, "packages/alpha/SPEC.md", authoringSpec)
	cliGit(t, root, "init", "-q", "-b", "main")
	cliCommit(t, root, "gc project")
	nested = appNestedDir(t, root)
	return root, outer, nested
}

func gcCLIJSONPruned(t *testing.T, output string) []string {
	t.Helper()
	envelope := decodeEnvelope(t, output)
	var data map[string]json.RawMessage
	if err := json.Unmarshal(envelope["data"], &data); err != nil || len(data) != 1 {
		t.Fatalf("gc data=%s; want object with pruned field: %v", envelope["data"], err)
	}
	var pruned []string
	if err := json.Unmarshal(data["pruned"], &pruned); err != nil || pruned == nil {
		t.Fatalf("gc pruned=%s; want array: %v", data["pruned"], err)
	}
	var diagnostics []json.RawMessage
	if err := json.Unmarshal(envelope["diagnostics"], &diagnostics); err != nil || diagnostics == nil || len(diagnostics) != 0 {
		t.Errorf("gc diagnostics=%s; want empty array: %v", envelope["diagnostics"], err)
	}
	return pruned
}

func Test_APP_150_APP_151_GCNearestProjectSortedAndEmpty(t *testing.T) {
	bin := buildAppBinary(t)
	for _, jsonMode := range []bool{false, true} {
		name := "plain"
		if jsonMode {
			name = "JSON"
		}
		t.Run(name, func(t *testing.T) {
			root, outer, nested := gcCLIProject(t)
			head := cliGit(t, root, "rev-parse", "HEAD")
			outerHead := cliGit(t, outer, "rev-parse", "HEAD")
			oldDrill := "refs/fireweed/drills/alpha/20200101T000000.000000000"
			futureDrill := "refs/fireweed/drills/alpha/20990101T000000.000000000"
			attempt := "refs/fireweed/attempts/finished/2"
			staleClaim := "refs/fireweed/claims/z-stale"
			gcCLIRef(t, outer, oldDrill, outerHead)
			appPut(t, root, "changes/archive/0002-finished.md", "---\nmodule: alpha\nphase: applied\n---\n")
			linked := filepath.Join(t.TempDir(), "linked")
			cliGit(t, root, "worktree", "add", "-b", "linked", linked, "HEAD")
			gcCLIClaim(t, root, "main-live", root)
			gcCLIClaim(t, root, "linked-live", linked)
			gcCLIClaim(t, root, "z-stale", filepath.Join(t.TempDir(), "gone"))
			for _, ref := range []string{attempt, oldDrill, futureDrill, "refs/fireweed/other/keep"} {
				gcCLIRef(t, root, ref, head)
			}
			beforeHead := cliGit(t, root, "rev-parse", "HEAD")
			beforeIndex := cliGit(t, root, "write-tree")
			beforeStatus := cliGit(t, root, "status", "--porcelain=v1", "--untracked-files=all")
			want := []string{attempt, staleClaim, oldDrill}
			sort.Strings(want)
			args := []string{"gc"}
			if jsonMode {
				args = append(args, "--json")
			}
			status, out, stderr := runAppBinary(t, bin, nested, args...)
			if status != 0 || stderr != "" {
				t.Fatalf("gc: exit=%d stdout=%q stderr=%q", status, out, stderr)
			}
			if jsonMode {
				if got := gcCLIJSONPruned(t, out); !reflect.DeepEqual(got, want) {
					t.Errorf("JSON pruned=%v; want sorted full names %v", got, want)
				}
			} else {
				position := -1
				for _, ref := range want {
					at := strings.Index(out, ref)
					if at <= position || strings.Count(out, ref) != 1 {
						t.Errorf("plain output %q does not list %q once in order", out, ref)
					}
					position = at
				}
			}
			if got := gcCLIRefs(t, root); !reflect.DeepEqual(got, []string{"refs/fireweed/claims/linked-live", "refs/fireweed/claims/main-live", futureDrill, "refs/fireweed/other/keep"}) {
				t.Errorf("gc retained refs=%v", got)
			}
			if got := gcCLIRefs(t, outer); !reflect.DeepEqual(got, []string{oldDrill}) {
				t.Errorf("gc changed enclosing project refs=%v", got)
			}
			if cliGit(t, root, "rev-parse", "HEAD") != beforeHead || cliGit(t, root, "write-tree") != beforeIndex || cliGit(t, root, "status", "--porcelain=v1", "--untracked-files=all") != beforeStatus {
				t.Error("gc changed HEAD, index or working files")
			}
			status, out, stderr = runAppBinary(t, bin, nested, args...)
			if status != 0 || stderr != "" {
				t.Fatalf("second gc: exit=%d stdout=%q stderr=%q", status, out, stderr)
			}
			if jsonMode {
				if got := gcCLIJSONPruned(t, out); len(got) != 0 {
					t.Errorf("second gc pruned=%v; want []", got)
				}
			} else if !strings.Contains(strings.ToLower(out), "no refs") || !strings.Contains(strings.ToLower(out), "pruned") {
				t.Errorf("empty plain gc must explicitly report no refs pruned: %q", out)
			}
		})
	}
}

func gcCLIFailure(t *testing.T, bin, dir string, args []string, wantStatus int, code string) {
	t.Helper()
	jsonMode := len(args) > 0 && args[len(args)-1] == "--json"
	status, out, stderr := runAppBinary(t, bin, dir, args...)
	message := commandFailure(t, status, wantStatus, out, stderr, jsonMode)
	if jsonMode {
		message = assertFailedJSON(t, out, code)
	} else if out != "" || !strings.HasPrefix(strings.TrimSpace(stderr), code+":") {
		t.Errorf("plain failure stdout=%q stderr=%q; want %s on stderr", out, stderr, code)
	}
	if !corrective(message) {
		t.Errorf("gc diagnostic lacks correction: %q", message)
	}
}

func Test_APP_150_APP_152_GCRejectsUsageBeforePruning(t *testing.T) {
	bin := buildAppBinary(t)
	root, _, nested := gcCLIProject(t)
	oldDrill := "refs/fireweed/drills/alpha/20200101T000000.000000000"
	gcCLIRef(t, root, oldDrill, cliGit(t, root, "rev-parse", "HEAD"))
	before := gcCLIRefs(t, root)
	for _, args := range [][]string{{"gc", "extra"}, {"gc", "--bogus"}, {"gc", "--max", "1"}} {
		for _, jsonMode := range []bool{false, true} {
			call := append([]string(nil), args...)
			if jsonMode {
				call = append(call, "--json")
			}
			gcCLIFailure(t, bin, nested, call, 2, "U01")
			if got := gcCLIRefs(t, root); !reflect.DeepEqual(got, before) {
				t.Errorf("invalid gc usage pruned refs: before=%v after=%v", before, got)
			}
		}
	}
}

func Test_APP_152_GCConfigurationAndGitFailures(t *testing.T) {
	bin := buildAppBinary(t)
	for _, tc := range []struct {
		name, config string
	}{
		{"malformed-config", "modules: [broken\n"},
		{"invalid-retention", appFixtureConfig + "gc:\n  drill_max_age: 0s\n"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root, _, nested := gcCLIProject(t)
			oldDrill := "refs/fireweed/drills/alpha/20200101T000000.000000000"
			gcCLIRef(t, root, oldDrill, cliGit(t, root, "rev-parse", "HEAD"))
			before := gcCLIRefs(t, root)
			appPut(t, root, "fireweed.yaml", tc.config)
			for _, jsonMode := range []bool{false, true} {
				args := []string{"gc"}
				if jsonMode {
					args = append(args, "--json")
				}
				gcCLIFailure(t, bin, nested, args, 2, "C01")
				if got := gcCLIRefs(t, root); !reflect.DeepEqual(got, before) {
					t.Errorf("failed gc changed refs: before=%v after=%v", before, got)
				}
			}
		})
	}
	noGit := appProjectFixture(t)
	for _, jsonMode := range []bool{false, true} {
		args := []string{"gc"}
		if jsonMode {
			args = append(args, "--json")
		}
		gcCLIFailure(t, bin, appNestedDir(t, noGit), args, 1, "P01")
	}
	for _, jsonMode := range []bool{false, true} {
		args := []string{"gc"}
		if jsonMode {
			args = append(args, "--json")
		}
		gcCLIFailure(t, bin, t.TempDir(), args, 2, "C01")
	}
}
