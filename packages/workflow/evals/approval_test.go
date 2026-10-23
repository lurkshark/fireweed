package evals

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	c "github.com/lurkshark/fireweed/packages/contracts"
	p "github.com/lurkshark/fireweed/packages/project"
	w "github.com/lurkshark/fireweed/packages/workflow"
)

type approvalFixtureData struct {
	contextFixtureData
	Claim w.Claim
	Want  p.Snapshot
}

func approvalCommit(t *testing.T, root string) string {
	t.Helper()
	claimGit(t, root, "add", "fireweed.yaml", "units", "changes")
	claimGit(t, root, "-c", "user.name=Approval Fixture", "-c", "user.email=approval@example.invalid", "commit", "-m", "approved source")
	return claimBase(t, root)
}
func approvalFixture(t *testing.T, phase, format string, late bool) approvalFixtureData {
	t.Helper()
	sourcePhase := phase
	if late {
		sourcePhase = ""
	}
	f := contextFixture(t, sourcePhase, format)
	// Preserve nontrivial archive data and a loader warning as part of source equality.
	archived := "---\nmodule: alpha\nphase: applied\nrun: archived-run\n---\n\n## Motivation\n\nPreserve historical source.\n\n## Added\n\n### ALP-090\nArchived behavior MUST remain historical.\n"
	worktreePut(t, f.Root, "changes/archive/0009-earlier.md", archived)
	worktreePut(t, f.Root, "outside-glob/SPEC.md", "not a selected module\n")
	if late {
		f.Raw = strings.ReplaceAll(f.Raw, "ALP-002", "new:late")
	}
	worktreePut(t, f.Root, f.Path, f.Raw)
	claimGit(t, f.Root, "add", "outside-glob/SPEC.md")
	base := approvalCommit(t, f.Root)
	source, err := p.LoadAt(f.Root, base)
	if err != nil {
		t.Fatal(err)
	}
	if len(source.Archive) != 1 || len(source.Diagnostics) == 0 {
		t.Fatalf("archive/diagnostic source premise missing: archives=%+v diagnostics=%+v", source.Archive, source.Diagnostics)
	}
	if late {
		allocated, _, err := p.Allocate(source, "current")
		if err != nil {
			t.Fatal(err)
		}
		if len(allocated.Added) != 1 || strings.HasPrefix(allocated.Added[0].ID, "new:") {
			t.Fatal("late assignment premise")
		}
		doc := allocated.Document
		doc.Metadata["run"] = "context-run"
		doc.Metadata["phase"] = phase
		f.Raw, err = c.RenderMarkdown(doc)
		if err != nil {
			t.Fatal(err)
		}
	}
	selected, err := c.ParseChange(f.Raw, f.Path, "ALP")
	if err != nil {
		t.Fatal(err)
	}
	if ds := w.CheckState(source.Modules[0], f.Raw, f.Raw, f.Path); len(ds) != 0 {
		t.Fatalf("valid approved-state premise %+v", ds)
	}
	cl := w.Claim{Worktree: f.Root, Change: "current", Run: "context-run", Base: base, State: f.Raw}
	cl.Object = claimRaw(t, f.Root, cl.Change, claimPayload(t, cl))
	// An unrelated corrupt claim must not affect an exact selected lookup.
	claimRaw(t, f.Root, "unrelated-malformed", "not JSON")
	source.Changes[1] = selected
	f.Snapshot = source
	return approvalFixtureData{f, cl, source}
}
func approvalError(t *testing.T, start, slug, code string) error {
	t.Helper()
	got, cl, err := w.LoadApproval(start, slug)
	if !reflect.DeepEqual(got, p.Snapshot{}) || cl != (w.Claim{}) {
		t.Fatalf("partial rejected source %+v %+v", got, cl)
	}
	if err == nil || !strings.Contains(err.Error(), code) || !policyCorrective(err.Error()) {
		t.Fatalf("want corrective %s: %v", code, err)
	}
	return err
}
func approvalAssert(t *testing.T, start string, f approvalFixtureData) p.Snapshot {
	t.Helper()
	got, cl, err := w.LoadApproval(start, "current")
	if err != nil || cl != f.Claim || !reflect.DeepEqual(got, f.Want) {
		t.Fatalf("approved source mismatch: err %v claim %+v want %+v\nsource %s\nwant %s", err, cl, f.Claim, policyJSON(t, got), policyJSON(t, f.Want))
	}
	return got
}
func Test_FLOW020_FLOW021ApprovalSources(t *testing.T) {
	for _, format := range []string{"sha1", "sha256"} {
		for _, mode := range []string{"current", "evals", "audit", "impl", "late", "nested", "linked", "moved-head", "replacement"} {
			t.Run(format+"/"+mode, func(t *testing.T) {
				phase := "impl"
				if mode == "evals" || mode == "audit" {
					phase = mode
				}
				f := approvalFixture(t, phase, format, mode == "late")
				start := f.Root
				if mode == "current" {
					claimGit(t, f.Root, "update-ref", "-d", "refs/fireweed/claims/current", f.Claim.Object)
					// Use a new valid current config value to distinguish Load from LoadAt.
					raw, err := os.ReadFile(filepath.Join(f.Root, "fireweed.yaml"))
					if err != nil {
						t.Fatal(err)
					}
					worktreePut(t, f.Root, "fireweed.yaml", string(raw)+"\nverify_dependents: none\n")
					want, err := p.Load(f.Root)
					if err != nil {
						t.Fatal(err)
					}
					if reflect.DeepEqual(want, f.Want) {
						t.Fatal("current source control did not differ")
					}
					got, cl, err := w.LoadApproval(start, "current")
					if err != nil || cl != (w.Claim{}) || !reflect.DeepEqual(got, want) {
						t.Fatalf("current source %+v %+v %v", got, cl, err)
					}
					return
				}
				if mode == "linked" {
					start = claimPeer(t, f.Root)
					f.Claim.Worktree = start
					f.Claim.Object = claimRaw(t, start, "current", claimPayload(t, f.Claim))
					f.Want.Root = start
				}
				if mode == "nested" {
					start = filepath.Join(f.Root, "units/a/e")
				}
				if mode == "moved-head" || mode == "replacement" {
					worktreePut(t, f.Root, "fireweed.yaml", "invalid current metadata\n")
					next := approvalCommit(t, f.Root)
					if next == f.Claim.Base {
						t.Fatal("HEAD movement premise")
					}
					if mode == "replacement" {
						protocolEnableReplacements(t)
						real := protocolRealGit(t)
						protocolGit(t, real, f.Root, "", "replace", "-f", f.Claim.Base, next)
						if raw := protocolGit(t, real, f.Root, "", "show", f.Claim.Base+":fireweed.yaml"); raw != "invalid current metadata\n" {
							t.Fatal("committed replacement view control", raw)
						}
					}
				}
				// These are feedback bytes, never metadata authority.
				root := f.Want.Root
				worktreePut(t, root, "fireweed.yaml", "[: broken current config")
				worktreePut(t, root, "units/a/SPEC.md", "broken worker module")
				worktreePut(t, root, f.Path, "invalid worker change")
				worktreePut(t, root, "changes/0001-before.md", "invalid worker predecessor")
				worktreePut(t, root, "changes/archive/0009-earlier.md", "invalid worker archive")
				got := approvalAssert(t, start, f)
				// The result is still unprojected; composing the public projection applies once.
				before, err := p.Project(got, "current")
				if err != nil {
					t.Fatal(err)
				}
				projected, err := p.ApplyChange(before.Modules[0], got.Changes[1])
				if err != nil {
					t.Fatal(err)
				}
				if len(got.Modules[0].Specs) != 1 || len(before.Modules[1].Specs) != 2 || len(projected.Specs) != 2 {
					t.Fatal("restoration projected queued effects prematurely")
				}
			})
		}
	}
}
func Test_FLOW020_FLOW021DeletedCurrentMetadata(t *testing.T) {
	for _, path := range []string{"fireweed.yaml", "units/a/SPEC.md", "changes/0002-current.md", "changes", "units"} {
		t.Run(path, func(t *testing.T) {
			f := approvalFixture(t, "evals", "sha1", true)
			if err := os.RemoveAll(filepath.Join(f.Root, path)); err != nil {
				t.Fatal(err)
			}
			approvalAssert(t, f.Root, f)
		})
	}
}
func Test_FLOW020_FLOW022InvalidApprovalSources(t *testing.T) {
	for _, mode := range []string{"foreign", "reservation", "invalid-state", "missing-run", "missing-phase", "applied", "wrong-run", "wrong-target", "missing-base", "malformed", "symbolic", "dangling-symbolic", "missing-object", "nonblob", "absent-selected", "archive-only", "missing-target", "duplicate-target", "duplicate-change", "nested-path", "inspection-failure", "no-config"} {
		t.Run(mode, func(t *testing.T) {
			f := approvalFixture(t, "impl", "sha1", false)
			code := "P02"
			switch mode {
			case "foreign":
				f.Claim.Worktree = claimPeer(t, f.Root)
			case "reservation":
				f.Claim.State = ""
			case "invalid-state":
				f.Claim.State = "---\n[broken"
			case "missing-run":
				f.Claim.State = strings.ReplaceAll(f.Raw, "run: context-run\n", "")
			case "missing-phase":
				f.Claim.State = strings.ReplaceAll(f.Raw, "phase: impl\n", "")
			case "applied":
				f.Claim.State = strings.ReplaceAll(f.Raw, "phase: impl", "phase: applied")
			case "wrong-run":
				f.Claim.Run = "different-run"
			case "wrong-target":
				f.Claim.State = strings.ReplaceAll(f.Raw, "module: alpha", "module: beta")
			case "missing-base":
				f.Claim.Base = strings.Repeat("f", len(f.Claim.Base))
			case "malformed":
				claimRaw(t, f.Root, "current", "not JSON")
			case "symbolic", "dangling-symbolic":
				target := "refs/heads/main"
				if mode == "dangling-symbolic" {
					target = "refs/heads/missing-owner-target"
				}
				claimGit(t, f.Root, "symbolic-ref", "refs/fireweed/claims/current", target)
			case "missing-object":
				worktreePut(t, claimCommon(t, f.Root), "refs/fireweed/claims/current", strings.Repeat("f", len(f.Claim.Object))+"\n")
			case "nonblob":
				claimGit(t, f.Root, "update-ref", "refs/fireweed/claims/current", f.Claim.Base)

			case "inspection-failure":
				if os.PathSeparator == '\\' {
					t.Skip("POSIX selected-ref fault injection")
				}
				real := protocolRealGit(t)
				dir := t.TempDir()
				marker := filepath.Join(dir, "selected")
				script := "#!/bin/sh\nfor a in \"$@\"; do if [ \"$a\" = refs/fireweed/claims/current ]; then printf inspected >> " + claimQuote(marker) + "; printf 'selected claim read failed' >&2; exit 92; fi; done\nexec " + claimQuote(real) + " \"$@\"\n"
				worktreePut(t, dir, "git", script)
				if err := os.Chmod(filepath.Join(dir, "git"), 0755); err != nil {
					t.Fatal(err)
				}
				t.Setenv("PATH", dir+string(os.PathListSeparator)+os.Getenv("PATH"))
				out, stderr, err := protocolCommand(t, filepath.Join(dir, "git"), f.Root, "", "rev-parse", "--verify", "refs/fireweed/claims/current")
				if err == nil || out != "" || stderr != "selected claim read failed" {
					t.Fatal("selected inspection failure control")
				}
			default:
				claimGit(t, f.Root, "update-ref", "-d", "refs/fireweed/claims/current", f.Claim.Object)
				code = "C05"
				switch mode {
				case "absent-selected":
					if err := os.Remove(filepath.Join(f.Root, f.Path)); err != nil {
						t.Fatal(err)
					}
				case "archive-only":
					if err := os.Rename(filepath.Join(f.Root, f.Path), filepath.Join(f.Root, "changes/archive/0002-current.md")); err != nil {
						t.Fatal(err)
					}
				case "missing-target":
					if err := os.Remove(filepath.Join(f.Root, "units/a/SPEC.md")); err != nil {
						t.Fatal(err)
					}
				case "duplicate-target":
					worktreePut(t, f.Root, "units/x/SPEC.md", f.Spec)
				case "duplicate-change":
					worktreePut(t, f.Root, "changes/0042-current.md", f.Raw)
				case "nested-path":
					if err := os.Remove(filepath.Join(f.Root, f.Path)); err != nil {
						t.Fatal(err)
					}
					worktreePut(t, f.Root, "changes/nested/0002-current.md", f.Raw)
				case "no-config":
					if err := os.Remove(filepath.Join(f.Root, "fireweed.yaml")); err != nil {
						t.Fatal(err)
					}
					code = "C01"
				}
			}
			if mode == "foreign" || mode == "reservation" || mode == "invalid-state" || mode == "missing-run" || mode == "missing-phase" || mode == "applied" || mode == "wrong-run" || mode == "wrong-target" || mode == "missing-base" {
				f.Claim.Object = claimRaw(t, f.Root, "current", claimPayload(t, f.Claim))
			}
			var detail string
			var stateDetails []string
			switch mode {
			case "malformed", "symbolic", "dangling-symbolic", "missing-object", "nonblob", "missing-base":
				_, underlying := w.LoadClaim(f.Root, "current")
				if underlying == nil {
					t.Fatal("public claim error premise")
				}
				detail = underlying.Error()
			case "invalid-state", "missing-run", "missing-phase", "applied", "wrong-target":
				_, underlying := c.ParseChange(f.Claim.State, f.Path, "ALP")
				if underlying != nil {
					detail = underlying.Error()
				} else {
					ds := w.CheckState(f.Want.Modules[0], f.Claim.State, f.Claim.State, f.Path)
					if len(ds) == 0 {
						t.Fatal("public approved-state rejection premise missing")
					}
					for _, d := range ds {
						stateDetails = append(stateDetails, d.Message)
					}
				}
			case "inspection-failure":
				detail = "selected claim read failed"
			case "no-config":
				_, underlying := p.Load(f.Root)
				if underlying == nil {
					t.Fatal("public discovery error premise")
				}
				detail = underlying.Error()
			}
			err := approvalError(t, f.Root, "current", code)
			if detail != "" && !strings.Contains(err.Error(), detail) {
				t.Fatalf("lost underlying public detail %q: %v", detail, err)
			}
			for _, detail := range stateDetails {
				if detail == "" || !strings.Contains(err.Error(), detail) {
					t.Fatalf("lost public approved-state problem %q: %v", detail, err)
				}
			}

		})
	}
	for _, slug := range []string{"", "Upper", " current", "../current", "current/other", "-option", "current\x00"} {
		t.Run("selector/"+fmt.Sprintf("%q", slug), func(t *testing.T) { approvalError(t, filepath.Join(t.TempDir(), "must-not-open"), slug, "C05") })
	}
}
func Test_FLOW022ClaimedBaselineSelectionErrors(t *testing.T) {
	for _, mode := range []string{"missing-change", "archive-only", "missing-module", "duplicate-module", "duplicate-change"} {
		t.Run(mode, func(t *testing.T) {
			f := approvalFixture(t, "impl", "sha1", false)
			switch mode {
			case "missing-change":
				if err := os.Remove(filepath.Join(f.Root, f.Path)); err != nil {
					t.Fatal(err)
				}
			case "archive-only":
				if err := os.Rename(filepath.Join(f.Root, f.Path), filepath.Join(f.Root, "changes/archive/0002-current.md")); err != nil {
					t.Fatal(err)
				}
			case "missing-module":
				if err := os.Remove(filepath.Join(f.Root, "units/a/SPEC.md")); err != nil {
					t.Fatal(err)
				}
			case "duplicate-module":
				worktreePut(t, f.Root, "units/x/SPEC.md", f.Spec)
			case "duplicate-change":
				worktreePut(t, f.Root, "changes/0042-current.md", f.Raw)
			}
			f.Claim.Base = approvalCommit(t, f.Root)
			f.Claim.Object = claimRaw(t, f.Root, "current", claimPayload(t, f.Claim))
			// A current repaired document cannot substitute for a missing committed source.
			worktreePut(t, f.Root, f.Path, f.Raw)
			worktreePut(t, f.Root, "units/a/SPEC.md", f.Spec)
			approvalError(t, f.Root, "current", "C05")
		})
	}
}
func Test_FLOW023ApprovalResultOwnership(t *testing.T) {
	for _, claimed := range []bool{false, true} {
		t.Run(fmt.Sprint(claimed), func(t *testing.T) {
			f := approvalFixture(t, "impl", "sha1", true)
			if !claimed {
				claimGit(t, f.Root, "update-ref", "-d", "refs/fireweed/claims/current", f.Claim.Object)
			}
			a, ac, ae := w.LoadApproval(f.Root, "current")
			b, bc, be := w.LoadApproval(f.Root, "current")
			if ae != nil || be != nil {
				t.Fatalf("ownership setup %v %v", ae, be)
			}
			before := policyJSON(t, b)
			a.Modules[0].Name = "mutated"
			a.Modules[1].Document.Metadata["name"] = "mutated"
			a.Config.Protected[0] = "mutated"
			a.Changes[1].Set["depends_on"].([]any)[0] = "mutated"
			a.Changes[1].Document.Sections[0].Body = "mutated"
			a.Archive[0].Added[0].Body = "mutated"
			a.Diagnostics[0].Message = "mutated"
			ac.State = "mutated"
			if policyJSON(t, b) != before || bc.State == ac.State {
				t.Fatal("peer source or claim aliases result")
			}
			again, _, err := w.LoadApproval(f.Root, "current")
			if err != nil || policyJSON(t, again) != before {
				t.Fatal("mutation changed later source", err)
			}
		})
	}
}
func Test_FLOW023ApprovalReadOnlyState(t *testing.T) {
	if os.PathSeparator == '\\' {
		t.Skip("POSIX configured/native hook controls; portable approval cases run separately")
	}
	for _, linked := range []bool{false, true} {
		for _, mode := range []string{"claimed", "unclaimed", "foreign", "invalid", "absent"} {
			t.Run(fmt.Sprint(linked)+"/"+mode, func(t *testing.T) {
				f := approvalFixture(t, "audit", "sha1", false)
				originalRoot := f.Root
				if linked {
					f.Root = claimPeer(t, f.Root)
					if mode != "foreign" {
						f.Claim.Worktree = f.Root
						f.Claim.Object = claimRaw(t, f.Root, "current", claimPayload(t, f.Claim))
					}
				}

				if mode == "unclaimed" {
					claimGit(t, f.Root, "update-ref", "-d", "refs/fireweed/claims/current", f.Claim.Object)
				}
				if mode == "foreign" && !linked {
					f.Claim.Worktree = claimPeer(t, f.Root)
					f.Claim.Object = claimRaw(t, f.Root, "current", claimPayload(t, f.Claim))
				}
				if mode == "invalid" {
					claimRaw(t, f.Root, "current", "invalid")
				}
				native := contextNativeHook(t, f.Root)
				traps, marker := worktreeTraps(t, false)
				for _, name := range []string{"context-hook-trap", "context-agent-trap"} {
					worktreePut(t, traps, name, "#!/bin/sh\nprintf attempted >> "+claimQuote(marker)+"\nexit 97\n")
					if err := os.Chmod(filepath.Join(traps, name), 0755); err != nil {
						t.Fatal(err)
					}
					_, _, err := protocolCommand(t, filepath.Join(traps, name), f.Root, "")
					exit, ok := err.(*exec.ExitError)
					if !ok || exit.ExitCode() != 97 {
						t.Fatal("configured command control did not exit as expected", name, err)
					}
					if _, err := os.Stat(marker); err != nil {
						t.Fatal("configured command marker control", name, err)
					}
					if err := os.Remove(marker); err != nil {
						t.Fatal(err)
					}

				}
				tmp := stateCanonicalDirectory(t, t.TempDir())
				t.Setenv("TMPDIR", tmp)
				t.Setenv("TMP", tmp)
				t.Setenv("TEMP", tmp)
				common := claimCommon(t, f.Root)
				files := claimTree(t, f.Root, func(path string) bool { return path == ".git" })
				originalFiles := claimTree(t, originalRoot, func(path string) bool { return path == ".git" })
				metadata := claimTree(t, common, func(string) bool { return false })
				temps := claimTree(t, tmp, func(string) bool { return false })
				check := protocolParent(t)
				defer check()
				slug := "current"
				if mode == "absent" {
					slug = "absent"
				}
				got, cl, err := w.LoadApproval(f.Root, slug)
				if mode == "foreign" || mode == "invalid" || mode == "absent" {
					if err == nil || !reflect.DeepEqual(got, p.Snapshot{}) || cl != (w.Claim{}) {
						t.Fatal("bad rejection shape", err)
					}
				} else if err != nil {
					t.Fatal(err)
				}
				if !reflect.DeepEqual(files, claimTree(t, f.Root, func(path string) bool { return path == ".git" })) || !reflect.DeepEqual(metadata, claimTree(t, common, func(string) bool { return false })) || !reflect.DeepEqual(temps, claimTree(t, tmp, func(string) bool { return false })) {
					t.Fatal("approval read changed files/index/refs/common Git/temp")
				}
				if !reflect.DeepEqual(originalFiles, claimTree(t, originalRoot, func(path string) bool { return path == ".git" })) {
					t.Fatal("approval changed original linked peer worktree files")
				}
				for _, path := range []string{marker, native} {
					if _, err := os.Stat(path); !os.IsNotExist(err) {
						t.Fatal("approval ran a forbidden command", path)
					}
				}
			})
		}
	}
}

func Test_FLOW020_FLOW021ApprovalFixtureControls(t *testing.T) {
	for _, late := range []bool{false, true} {
		f := approvalFixture(t, "impl", "sha1", late)
		got, err := w.LoadClaim(f.Root, "current")
		if err != nil || got != f.Claim {
			t.Fatal("claim fixture control", err)
		}
		base, err := p.LoadAt(f.Root, f.Claim.Base)
		if err != nil {
			t.Fatal(err)
		}
		if late && base.Changes[1].Added[0].ID == f.Want.Changes[1].Added[0].ID {
			t.Fatal("late source failed to distinguish assigned state")
		}
		if _, err = json.Marshal(f.Want); err != nil {
			t.Fatal(err)
		}
	}
}

func Test_FLOW023ApprovalEmptyResultOwnership(t *testing.T) {
	f := approvalFixture(t, "impl", "sha1", false)
	config, err := os.ReadFile(filepath.Join(f.Root, "fireweed.yaml"))
	if err != nil {
		t.Fatal(err)
	}
	worktreePut(t, f.Root, "fireweed.yaml", strings.ReplaceAll(strings.ReplaceAll(string(config), "worker_writable: [shared/**]", "worker_writable: []"), "protected: [private/**, units/a/e/protected.txt]", "protected: []"))
	worktreePut(t, f.Root, "units/a/SPEC.md", strings.Replace(f.Spec, "depends_on: [beta, gamma]", "depends_on: []\ninterface: []\npreserve: []", 1))
	f.Claim.Base = approvalCommit(t, f.Root)
	f.Claim.Object = claimRaw(t, f.Root, "current", claimPayload(t, f.Claim))
	source, err := p.LoadAt(f.Root, f.Claim.Base)
	if err != nil {
		t.Fatal(err)
	}
	if len(source.Config.Protected) != 0 || len(source.Config.WorkerWritable) != 0 || len(source.Modules[0].DependsOn) != 0 || len(source.Modules[0].Interface) != 0 || len(source.Modules[0].Preserve) != 0 {
		t.Fatal("independent explicit empty source control")
	}
	a, _, ea := w.LoadApproval(f.Root, "current")
	b, _, eb := w.LoadApproval(f.Root, "current")
	if ea != nil || eb != nil {
		t.Fatal("empty result setup", ea, eb)
	}
	if len(a.Config.Protected) != 0 || len(a.Config.WorkerWritable) != 0 || len(a.Modules[0].DependsOn) != 0 || len(a.Modules[0].Interface) != 0 || len(a.Modules[0].Preserve) != 0 {
		t.Fatal("explicit empty result premise")
	}
	a.Config.Protected = append(a.Config.Protected, "first")
	b.Config.Protected = append(b.Config.Protected, "second")
	a.Config.WorkerWritable = append(a.Config.WorkerWritable, "first")
	b.Config.WorkerWritable = append(b.Config.WorkerWritable, "second")
	a.Modules[0].DependsOn = append(a.Modules[0].DependsOn, "first")
	b.Modules[0].DependsOn = append(b.Modules[0].DependsOn, "second")
	a.Modules[0].Interface = append(a.Modules[0].Interface, "first")
	b.Modules[0].Interface = append(b.Modules[0].Interface, "second")
	a.Modules[0].Preserve = append(a.Modules[0].Preserve, "first")
	b.Modules[0].Preserve = append(b.Modules[0].Preserve, "second")
	if a.Config.Protected[0] != "first" || a.Config.WorkerWritable[0] != "first" || a.Modules[0].DependsOn[0] != "first" || a.Modules[0].Interface[0] != "first" || a.Modules[0].Preserve[0] != "first" {
		t.Fatal("empty result peers share append storage")
	}
	b.Modules[0].Interface[0] = "peer mutation"
	if a.Modules[0].Interface[0] != "first" {
		t.Fatal("empty result peer mutation aliased")
	}
	again, _, err := w.LoadApproval(f.Root, "current")
	if err != nil || len(again.Modules[0].Interface) != 0 || len(again.Config.Protected) != 0 {
		t.Fatal("appended result changed later source", err)
	}
}

func Test_FLOW021_FLOW022ApprovalExactStateAndBaselineFailures(t *testing.T) {
	for _, kind := range []string{"blocked-fix", "regenerate", "bad-baseline-config", "valid-selector-no-repository"} {
		t.Run(kind, func(t *testing.T) {
			f := approvalFixture(t, "impl", "sha1", true)
			if kind == "valid-selector-no-repository" {
				approvalError(t, filepath.Join(t.TempDir(), "missing"), "current", "C01")
				return
			}
			if kind == "bad-baseline-config" {
				original, err := os.ReadFile(filepath.Join(f.Root, "fireweed.yaml"))
				if err != nil {
					t.Fatal(err)
				}
				worktreePut(t, f.Root, "fireweed.yaml", "[: invalid approved config")
				f.Claim.Base = approvalCommit(t, f.Root)
				f.Claim.Object = claimRaw(t, f.Root, "current", claimPayload(t, f.Claim))
				_, underlying := p.LoadAt(f.Root, f.Claim.Base)
				if underlying == nil {
					t.Fatal("invalid baseline config must fail its public loader")
				}
				worktreePut(t, f.Root, "fireweed.yaml", string(original))
				err = approvalError(t, f.Root, "current", "P02")
				if !strings.Contains(err.Error(), underlying.Error()) {
					t.Fatalf("lost immutable source detail %q: %v", underlying, err)
				}
				return
			}
			selected := f.Want.Changes[1]
			doc := selected.Document
			if kind == "blocked-fix" {
				doc.Metadata["blocked"] = "repair source metadata"
				doc.Metadata["fixes"] = []string{"ALP-001"}
				doc.Metadata["already_satisfied"] = []string{selected.Added[0].ID}
			} else {
				doc.Metadata["regenerate"] = true
				for i := range doc.Sections {
					if doc.Sections[i].Name == "Added" {
						doc.Sections[i].Body = ""
					}
				}
			}
			doc.Sections = append(doc.Sections, c.Section{Name: "Decisions", Body: "Exact approved choice λ\r\nsecond line"}, c.Section{Name: "Questions", Body: "Answered and retained."}, c.Section{Name: "Notes", Body: "Step notes."})
			raw, err := c.RenderMarkdown(doc)
			if err != nil {
				t.Fatal(err)
			}
			parsed, err := c.ParseChange(raw, f.Path, "ALP")
			if err != nil {
				t.Fatal(err)
			}
			if ds := w.CheckState(f.Want.Modules[0], raw, raw, f.Path); len(ds) != 0 {
				t.Fatal("valid detailed state premise", ds)
			}
			f.Claim.State = raw
			f.Claim.Object = claimRaw(t, f.Root, "current", claimPayload(t, f.Claim))
			f.Want.Changes[1] = parsed
			approvalAssert(t, f.Root, f)
		})
	}
}

func Test_FLOW020ApprovalWithGitWarnings(t *testing.T) {
	for _, format := range []string{"sha1", "sha256"} {
		t.Run(format, func(t *testing.T) {
			f := approvalFixture(t, "evals", format, true)
			protocolWarnings(t, f.Root)
			approvalAssert(t, f.Root, f)
		})
	}
}

func Test_FLOW020_FLOW021ExactApprovalSelectors(t *testing.T) {
	for _, claimed := range []bool{false, true} {
		for _, slug := range []string{"current", "cur", "current-neighbor", "current-bad"} {
			t.Run(fmt.Sprint(claimed)+"/"+slug, func(t *testing.T) {
				f := approvalFixture(t, "impl", "sha1", false)
				neighbor := "---\nmodule: alpha\n---\n\n## Motivation\n\nA separate prefix-neighbor maintenance change.\n"
				worktreePut(t, f.Root, "changes/0042-current-neighbor.md", neighbor)
				f.Claim.Base = approvalCommit(t, f.Root)
				f.Claim.Object = claimRaw(t, f.Root, "current", claimPayload(t, f.Claim))
				claimRaw(t, f.Root, "current-bad", "invalid prefix-neighbor claim")
				baseline, err := p.LoadAt(f.Root, f.Claim.Base)
				if err != nil {
					t.Fatal(err)
				}
				if len(baseline.Changes) != 4 || baseline.Changes[3].Slug != "current-neighbor" {
					t.Fatal("prefix-neighbor queue premise")
				}
				parsed, err := c.ParseChange(f.Claim.State, f.Path, "ALP")
				if err != nil {
					t.Fatal(err)
				}
				baseline.Changes[1] = parsed
				f.Want = baseline
				if !claimed {
					claimGit(t, f.Root, "update-ref", "-d", "refs/fireweed/claims/current", f.Claim.Object)
				}
				if slug == "cur" {
					if claimed {
						exact := f.Claim
						exact.Change = "cur"
						exact.Object = claimRaw(t, f.Root, "cur", claimPayload(t, exact))
						stored, err := w.LoadClaim(f.Root, "cur")
						if err != nil || stored != exact {
							t.Fatalf("exact partial-selector claim control: %+v %v", stored, err)
						}
						for _, change := range baseline.Changes {
							if change.Slug == "cur" {
								t.Fatal("partial-selector baseline unexpectedly contains exact slug")
							}
						}
					}
					approvalError(t, f.Root, slug, "C05")
					return
				}
				if slug == "current-bad" {
					approvalError(t, f.Root, slug, "P02")
					return
				}
				if slug == "current" && claimed {
					approvalAssert(t, f.Root, f)
					return
				}
				want, err := p.Load(f.Root)
				if err != nil {
					t.Fatal(err)
				}
				got, cl, err := w.LoadApproval(f.Root, slug)
				if err != nil || cl != (w.Claim{}) || !reflect.DeepEqual(got, want) {
					t.Fatalf("exact unclaimed selector %s: %+v %v", slug, cl, err)
				}
			})
		}
	}
}
