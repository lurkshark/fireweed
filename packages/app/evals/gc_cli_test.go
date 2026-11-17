package evals_test

import (
	"encoding/json"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/workflow"
)

func gcCLIRef(t *testing.T, root, name, oid string) {
	t.Helper()
	if err := gitstore.CASRef(root, name, "", oid); err != nil {
		t.Fatalf("create %s: %v", name, err)
	}
}

func gcCLIStaleClaim(t *testing.T, root, name string) {
	t.Helper()
	missing := filepath.Join(t.TempDir(), "gone")
	data, err := json.Marshal(map[string]string{"worktree": missing, "change": name})
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
			outerHead := cliGit(t, outer, "rev-parse", "HEAD")
			oldDrill := "refs/fireweed/drills/alpha/20200101T000000.000000000"
			futureDrill := "refs/fireweed/drills/alpha/20990101T000000.000000000"
			attempt := "refs/fireweed/attempts/finished/2"
			staleClaim := "refs/fireweed/claims/z-stale"
			gcCLIRef(t, outer, oldDrill, outerHead)
			appPut(t, root, "changes/archive/0002-finished.md", "---\nmodule: alpha\nphase: applied\n---\n")
			cliCommit(t, root, "record archived change")
			linked := filepath.Join(t.TempDir(), "linked")
			cliGit(t, root, "worktree", "add", "-b", "linked", linked, "HEAD")
			authoringChange(t, root, "0003-main-live.md", "Main needs an additional value.\n", "### new:main\nAlpha MUST expose the main value.\n", "")
			authoringChange(t, linked, "0004-linked-live.md", "Linked needs an additional value.\n", "### new:linked\nAlpha MUST expose the linked value.\n", "")
			cliCommit(t, root, "queue main worktree change")
			cliCommit(t, linked, "queue linked worktree change")
			live := []struct {
				root, slug string
				oid        string
			}{{root: root, slug: "main-live"}, {root: linked, slug: "linked-live"}}
			for i, claim := range live {
				started, err := workflow.Start(claim.root, claim.slug)
				if err != nil || started.ClaimOID == "" {
					t.Fatalf("start %s: claim=%q err=%v", claim.slug, started.ClaimOID, err)
				}
				live[i].oid = started.ClaimOID
			}
			gcCLIStaleClaim(t, root, "z-stale")
			head := cliGit(t, root, "rev-parse", "HEAD")
			for _, ref := range []string{attempt, oldDrill, futureDrill, "refs/fireweed/other/keep"} {
				gcCLIRef(t, root, ref, head)
			}
			before := migrateCapture(t, root)
			retained := []string{"refs/fireweed/claims/linked-live", "refs/fireweed/claims/main-live", futureDrill, "refs/fireweed/other/keep"}
			args := []string{"gc"}
			if jsonMode {
				args = append(args, "--json")
			}
			for i, want := range [][]string{{attempt, staleClaim, oldDrill}, {}} {
				status, out, stderr := runAppBinary(t, bin, nested, args...)
				if status != 0 || stderr != "" {
					t.Fatalf("gc call %d: exit=%d stdout=%q stderr=%q", i+1, status, out, stderr)
				}
				switch {
				case jsonMode:
					if got := gcCLIJSONPruned(t, out); !reflect.DeepEqual(got, want) {
						t.Errorf("JSON pruned=%v; want sorted full names %v", got, want)
					}
				case len(want) == 0:
					if lower := strings.ToLower(out); !strings.Contains(lower, "no refs") || !strings.Contains(lower, "pruned") {
						t.Errorf("empty plain gc must explicitly report no refs pruned: %q", out)
					}
				default:
					var got []string
					for _, field := range strings.Fields(out) {
						if strings.HasPrefix(field, "refs/") {
							got = append(got, field)
						}
					}
					if !reflect.DeepEqual(got, want) {
						t.Errorf("plain pruned=%v; want sorted full names %v", got, want)
					}
				}
				if got := gcCLIRefs(t, root); !reflect.DeepEqual(got, retained) {
					t.Errorf("gc retained refs=%v; want %v", got, retained)
				}
				if got := gcCLIRefs(t, outer); !reflect.DeepEqual(got, []string{oldDrill}) {
					t.Errorf("gc changed enclosing project refs=%v", got)
				}
				after := migrateCapture(t, root)
				if after.head != before.head || !reflect.DeepEqual(after.index, before.index) || !reflect.DeepEqual(after.files, before.files) {
					t.Error("gc changed HEAD, index or working files")
				}
				for _, claim := range live {
					current, err := workflow.Current(claim.root, claim.slug)
					if err != nil || current.ClaimOID != claim.oid {
						t.Errorf("gc changed %s claim: got %q, want %q, err=%v", claim.slug, current.ClaimOID, claim.oid, err)
					}
				}
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
