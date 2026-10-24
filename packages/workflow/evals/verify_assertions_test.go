package evals

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	k "github.com/lurkshark/fireweed/packages/checks"
	c "github.com/lurkshark/fireweed/packages/contracts"
	p "github.com/lurkshark/fireweed/packages/project"
	w "github.com/lurkshark/fireweed/packages/workflow"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"testing"
	"time"
)

func verifyInclusive(t *testing.T, f verifyFixtureData) p.Snapshot {
	t.Helper()
	s, e := p.Project(f.Snapshot, "current")
	if e != nil {
		t.Fatal(e)
	}
	for i, m := range s.Modules {
		if m.Name == "alpha" {
			s.Modules[i] = f.Module
		}
	}
	return s
}
func verifyGuardEvidence(t *testing.T, f verifyFixtureData, v w.Verification) {
	t.Helper()
	want := w.CheckWorktree(verifyInclusive(t, f), f.Module, f.Path, f.Claim)
	if !reflect.DeepEqual(v.Edits, want.Edits) {
		t.Errorf("actual guard edits got %+v want %+v", v.Edits, want.Edits)
	}
	for _, d := range want.Diagnostics {
		found := false
		for _, got := range v.Diagnostics {
			found = found || got == d
		}
		if !found {
			t.Errorf("dropped exact independent guard diagnostic %+v in %+v", d, v.Diagnostics)
		}
	}
}
func verifyModuleCoverage(t *testing.T, f verifyFixtureData, names []string) []k.Coverage {
	t.Helper()
	s := verifyInclusive(t, f)
	out := []k.Coverage{}
	for _, name := range names {
		found := false
		for _, m := range s.Modules {
			if m.Name == name {
				out = append(out, verifyNormalizeCoverage(t, f, k.Evaluate(context.Background(), s, m, nil, nil)))
				found = true
			}
		}
		if !found {
			t.Fatal("missing expected public module", name)
		}
	}
	verifyClear(t, f)
	return out
}
func verifyDiagnosticOrder(t *testing.T, ds []k.Diagnostic) {
	t.Helper()
	sorted := append([]k.Diagnostic{}, ds...)
	sort.Slice(sorted, func(i, j int) bool {
		a, b := sorted[i], sorted[j]
		if a.Path != b.Path {
			return a.Path < b.Path
		}
		if a.Code != b.Code {
			return a.Code < b.Code
		}
		if a.Severity != b.Severity {
			return a.Severity < b.Severity
		}
		return a.Message < b.Message
	})
	if !reflect.DeepEqual(sorted, ds) {
		t.Errorf("diagnostics not in required complete order: %+v", ds)
	}
	for i := 1; i < len(ds); i++ {
		if ds[i] == ds[i-1] {
			t.Error("duplicate exact diagnostic", ds[i])
		}
	}
}
func verifyRefreshFixture(t *testing.T, f *verifyFixtureData) {
	t.Helper()
	startCommit(t, f.Root, "Retain calibrated approved verification metadata")
	cl, e := w.UpdateClaim(f.Root, f.Claim, "HEAD", f.Raw)
	if e != nil {
		t.Fatal(e)
	}
	f.Claim = cl
	f.Snapshot, _, e = w.LoadApproval(f.Root, "current")
	if e != nil {
		t.Fatal(e)
	}
	s, e := p.Project(f.Snapshot, "current")
	if e != nil {
		t.Fatal(e)
	}
	for _, m := range s.Modules {
		if m.Name == "alpha" {
			f.Module, e = p.ApplyChange(m, f.Change)
			if e != nil {
				t.Fatal(e)
			}
		}
	}
}
func verifyPutModuleMetadata(t *testing.T, f verifyFixtureData, name string, key string, value any) {
	t.Helper()
	path := "units/" + name + "/SPEC.md"
	doc, e := c.ParseMarkdown(mustReadVerify(t, filepath.Join(f.Root, path)))
	if e != nil {
		t.Fatal(e)
	}
	doc.Metadata[key] = value
	raw, e := c.RenderMarkdown(doc)
	if e != nil {
		t.Fatal(e)
	}
	worktreePut(t, f.Root, path, raw)
}

func Test_FLOW042VerifyIndependentGuardsCompose(t *testing.T) {
	tool := verifyTool(t)
	for _, mode := range []string{"multiple-errors", "invalid-current", "unreadable-current", "current-directory"} {
		t.Run(mode, func(t *testing.T) {
			f := verifyFixture(t, tool, "impl", "all", "sha1", false)
			switch mode {
			case "multiple-errors":
				worktreePut(t, f.Root, f.Path, strings.Replace(strings.Replace(f.Raw, "run: verify-run", "run: worker-run", 1), "Verify approved phases.", "Worker changed motivation.", 1))
				worktreePut(t, f.Root, "units/alpha/api/old", "frozen interface mutation")
				worktreePut(t, f.Root, "units/beta/impl/old", "foreign mutation")
				worktreePut(t, f.Root, "private/forbidden", "protected mutation")
			case "invalid-current":
				worktreePut(t, f.Root, f.Path, "---\nmodule: [invalid\n")
			case "unreadable-current":
				full := filepath.Join(f.Root, f.Path)
				if e := os.Chmod(full, 0); e != nil {
					t.Fatal(e)
				}
				defer os.Chmod(full, 0644)
				if _, e := os.ReadFile(full); e == nil {
					t.Skip("host does not enforce the configured unreadable file mode")
				}
			case "current-directory":
				if e := os.Remove(filepath.Join(f.Root, f.Path)); e != nil {
					t.Fatal(e)
				}
				if e := os.Mkdir(filepath.Join(f.Root, f.Path), 0755); e != nil {
					t.Fatal(e)
				}
			}
			want := w.CheckWorktree(verifyInclusive(t, f), f.Module, f.Path, f.Claim)
			if len(want.Diagnostics) == 0 {
				t.Fatal("independent invalid guard premise absent")
			}
			if mode == "multiple-errors" && len(want.Diagnostics) < 4 {
				t.Fatal("multiple independent guard premise absent", want)
			}
			v := verifyInvoke(t, f, "current")
			verifyGuardEvidence(t, f, v)
			verifyDiagnosticOrder(t, v.Diagnostics)
			if len(v.Coverage) != 0 || len(verifyCalls(t, f.Plan)) != 0 {
				t.Error("failed guards ran hooks")
			}
		})
	}
}
func Test_FLOW045VerifyBoundaryUsesCompleteGraph(t *testing.T) {
	tool := verifyTool(t)
	f := verifyFixture(t, tool, "impl", "none", "sha1", false)
	other := stateCanonicalDirectory(t, t.TempDir())
	worktreePut(t, other, "boundary.output", "extra: illegal edge in unrelated module\n")
	verifyPutModuleMetadata(t, f, "extra", "hooks", map[string]any{"boundaries": verifyCommand(tool, "boundary", other)})
	verifyRefreshFixture(t, &f)
	s := verifyInclusive(t, f)
	ds := k.Boundaries(context.Background(), s)
	found := false
	for _, d := range ds {
		found = found || d.Code == "C15" && d.Path == "units/extra"
	}
	if !found {
		t.Fatal("complete-graph distinct boundary control absent", ds)
	}
	if len(verifyCalls(t, other)) != 1 {
		t.Fatal("distinct boundary command not live")
	}
	verifyClear(t, f)
	if e := os.Remove(filepath.Join(other, "calls")); e != nil {
		t.Fatal(e)
	}
	v := verifyInvoke(t, f, "current")
	verifyExpectedCoverage(t, f, []string{"alpha"}, v)
	verifyCode(t, v, "C15", "units/extra")
	calls := verifyCalls(t, other)
	if len(calls) != 1 || calls[0].Kind != "boundary" {
		t.Error("complete graph boundary omitted or duplicated", calls)
	}
}

func Test_FLOW043VerifyApprovedSelectedHookAndUnrelatedFailure(t *testing.T) {
	tool := verifyTool(t)
	for _, mode := range []string{"selected-override", "unrelated-failure"} {
		t.Run(mode, func(t *testing.T) {
			f := verifyFixture(t, tool, "evals", "all", "sha1", false)
			if mode == "selected-override" {
				other := stateCanonicalDirectory(t, t.TempDir())
				worktreePut(t, other, "alpha.xml", mustReadVerify(t, filepath.Join(f.Plan, "alpha.xml")))
				worktreePut(t, other, "alpha.exit", "1")
				doc := f.Change.Document
				set := doc.Metadata["set"].(map[string]any)
				set["hooks"] = map[string]any{"eval": verifyCommand(tool, "eval", other)}
				raw, e := c.RenderMarkdown(doc)
				if e != nil {
					t.Fatal(e)
				}
				verifyReseed(t, &f, raw)
				v := verifyInvoke(t, f, "current")
				verifyNoErrors(t, v.Diagnostics)
				verifyExpectedCoverage(t, f, []string{"alpha"}, v)
				if len(verifyCalls(t, f.Plan)) != 0 || len(verifyCalls(t, other)) != 1 {
					t.Error("selected effective hook override not used")
				}
				if len(v.Coverage) == 1 && v.Coverage[0].HookExit != 1 {
					t.Error("override's failed-before report omitted")
				}
			} else {
				id := f.Change.Added[0].ID
				verifyReport(t, f, "alpha", []string{"ALP-001", "ALP-002", "ALP-006", id}, map[string]bool{id: true, "ALP-006": true})
				want := verifyManualCoverage(t, f, []string{id}, nil)
				if len(want.Diagnostics) == 0 {
					t.Fatal("unrelated failure control absent")
				}
				v := verifyInvoke(t, f, "current")
				verifyCode(t, v, "C14", "units/alpha")
				if len(v.Coverage) != 1 || !reflect.DeepEqual(verifyNormalizeCoverage(t, f, v.Coverage[0]), want) {
					t.Error("unrelated failure evidence not retained", v)
				}
			}
		})
	}
}
func Test_FLOW044VerifyAuditOnlyParsedSection(t *testing.T) {
	tool := verifyTool(t)
	for _, mode := range []string{"outside-no-findings", "approved-notes-decoy", "outside-finding", "retained-finding", "fenced-heading"} {
		t.Run(mode, func(t *testing.T) {
			f := verifyFixture(t, tool, "audit", "all", "sha1", false)
			id := f.Change.Added[0].ID
			decoy := "No findings. Checked " + id + ".\n- " + id + ": \"MUST be checked\" — decoy"
			raw := f.Raw
			want := []w.Finding{}
			switch mode {
			case "approved-notes-decoy":
				verifyReseed(t, &f, f.Raw+"\n## Notes\n\n"+decoy+"\n")
				verifyRefreshFixture(t, &f)
				raw = f.Raw
			case "outside-no-findings":
				raw += "\n## Questions\n\n" + decoy + "\n"
			case "outside-finding":
				raw += "\n## Questions\n\n" + decoy + "\n\n## Audit\n\nNo findings. Checked " + id + ".\n"
			case "retained-finding":
				line := "- ALP-001: \"MUST remain\" — retained behavior gap"
				raw += "\n## Audit\n\n" + line + "\n"
				want = []w.Finding{{ID: "ALP-001", Clause: "MUST remain", Text: line}}
			case "fenced-heading":
				raw += "\n## Questions\n\n```markdown\n## Audit\n" + decoy + "\n```\n"
			}
			worktreePut(t, f.Root, f.Path, raw)
			if ds := w.CheckState(f.Module, f.Claim.State, raw, f.Path); len(ds) != 0 {
				t.Fatal("allowed audit/Questions fixture invalid", ds)
			}
			v := verifyInvoke(t, f, "current")
			if !reflect.DeepEqual(v.Findings, want) {
				t.Error("audit scanned other sections or dropped retained finding", v.Findings, want)
			}
			if mode == "outside-no-findings" || mode == "fenced-heading" || mode == "approved-notes-decoy" {
				verifyCode(t, v, "P1-03", f.Path)
			} else {
				verifyNoErrors(t, v.Diagnostics)
			}
			if len(verifyCalls(t, f.Plan)) != 0 {
				t.Error("audit ran hooks")
			}
		})
	}
}

// Snapshot all parent/repository/temp state after native startup is calibrated.
func verifyReadOnly(t *testing.T, f verifyFixtureData, call func()) {
	t.Helper()
	temp := stateCanonicalDirectory(t, t.TempDir())
	t.Setenv("TMPDIR", temp)
	t.Setenv("TMP", temp)
	t.Setenv("TEMP", temp)
	claimGit(t, f.Root, "--version")
	files := startTree(t, f.Root)
	meta := claimTree(t, claimCommon(t, f.Root), func(string) bool { return false })
	scratch := claimTree(t, temp, func(string) bool { return false })
	cwd, _ := os.Getwd()
	env := os.Environ()
	call()
	after, _ := os.Getwd()
	if after != cwd || !reflect.DeepEqual(env, os.Environ()) || !reflect.DeepEqual(files, startTree(t, f.Root)) || !reflect.DeepEqual(meta, claimTree(t, claimCommon(t, f.Root), func(string) bool { return false })) || !reflect.DeepEqual(scratch, claimTree(t, temp, func(string) bool { return false })) {
		t.Error("success/rejection/cancellation changed parent/files/Git metadata or leaked temporary storage")
	}
}
func Test_FLOW041_FLOW047VerifyFatalSelectionPreservesAllState(t *testing.T) {
	tool := verifyTool(t)
	for _, mode := range []string{"missing", "malformed", "foreign", "invalid-selector"} {
		t.Run(mode, func(t *testing.T) {
			f := verifyFixture(t, tool, "impl", "none", "sha1", false)
			slug, code := "current", "P02"
			switch mode {
			case "missing":
				if e := w.ReleaseClaim(f.Root, "current", f.Claim.Object); e != nil {
					t.Fatal(e)
				}
			case "malformed":
				claimRaw(t, f.Root, "current", "invalid JSON")
			case "foreign":
				cl := f.Claim
				cl.Worktree = claimPeer(t, f.Root)
				claimRaw(t, f.Root, "current", claimPayload(t, cl))
			case "invalid-selector":
				slug = "../bad"
				code = "C05"
			}
			verifyReadOnly(t, f, func() {
				v, e := w.Verify(context.Background(), f.Root, slug)
				if !reflect.DeepEqual(v, w.Verification{}) || e == nil || !strings.Contains(e.Error(), code) || !policyCorrective(e.Error()) {
					t.Error("fatal selection result", v, e)
				}
			})
			if len(verifyCalls(t, f.Plan)) != 0 {
				t.Error("fatal selection launched hooks")
			}
		})
	}
}
func Test_FLOW046VerifyDependentAndBoundaryMutations(t *testing.T) {
	tool := verifyTool(t)
	for _, stage := range []string{"beta", "boundary"} {
		t.Run(stage, func(t *testing.T) {
			f := verifyFixture(t, tool, "impl", "all", "sha1", false)
			path := "private/late-hook"
			raw, _ := json.Marshal(map[string]string{path: "late hook mutation survives"})
			worktreePut(t, f.Plan, stage+".mutation", string(raw))
			v := verifyInvoke(t, f, "current")
			verifyGuardEvidence(t, f, v)
			verifyExpectedCoverage(t, f, []string{"alpha", "beta", "delta", "gamma"}, v)
			verifyCode(t, v, "P2-01", path)
			if got := mustReadVerify(t, filepath.Join(f.Root, path)); got != "late hook mutation survives" {
				t.Error("late hook change repaired")
			}
			found := false
			for _, e := range v.Edits {
				found = found || e.Path == path
			}
			if !found {
				t.Error("late hook edit omitted")
			}
		})
	}
}
func Test_FLOW046_FLOW047VerifyLaterCancellationRetainsEvidence(t *testing.T) {
	tool := verifyTool(t)
	f := verifyFixture(t, tool, "impl", "all", "sha1", false)
	own := verifyManualCoverage(t, f, nil, nil)
	verifyNoErrors(t, own.Diagnostics)
	worktreePut(t, f.Plan, "beta.hold", "6s")
	verifyReadOnly(t, f, func() {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		type answer struct {
			v w.Verification
			e error
		}
		ch := make(chan answer, 1)
		go func() { v, e := w.Verify(ctx, f.Root, "current"); ch <- answer{v, e} }()
		until := time.Now().Add(5 * time.Second)
		for {
			if _, e := os.Stat(filepath.Join(f.Plan, "beta.ready")); e == nil {
				break
			}
			select {
			case r := <-ch:
				t.Fatalf("ended before later hook: %+v %v", r.v, r.e)
			default:
			}
			if time.Now().After(until) {
				t.Fatal("later hook ready control absent")
			}
			time.Sleep(10 * time.Millisecond)
		}
		cancel()
		select {
		case r := <-ch:
			if !errors.Is(r.e, context.Canceled) {
				t.Error("lost context cancellation", r.e)
			}
			verifyShape(t, f, r.v)
			verifyExpectedCoverage(t, f, []string{"alpha", "beta"}, r.v)
			if len(r.v.Coverage) == 2 {
				if !reflect.DeepEqual(verifyNormalizeCoverage(t, f, r.v.Coverage[0]), own) {
					t.Error("completed own coverage/output/status lost", r.v.Coverage[0], own)
				}
				if !strings.Contains(r.v.Coverage[1].Output, "hook-ready:beta") || r.v.Coverage[1].HookExit != -1 {
					t.Error("active hook partial output/status lost", r.v.Coverage[1])
				}
			}
		case <-time.After(2 * time.Second):
			t.Fatal("later cancellation did not terminate")
		}
		calls := verifyCalls(t, f.Plan)
		if len(calls) != 2 || calls[0].Module != "alpha" || calls[1].Module != "beta" {
			t.Error("ran later hooks after cancellation", calls)
		}
		time.Sleep(6100 * time.Millisecond)
		if _, e := os.Stat(filepath.Join(f.Plan, "beta.late")); !os.IsNotExist(e) {
			t.Error("cancelled later hook survived")
		}
	})
}

func Test_FLOW047VerifyMissingObjectLocality(t *testing.T) {
	tool := verifyTool(t)
	claimLocalTransport(t)
	for _, format := range []string{"sha1", "sha256"} {
		for _, linked := range []bool{false, true} {
			for _, denied := range []bool{false, true} {
				for _, kind := range []string{"selected-claim", "moved-head", "unrelated-ancestor"} {
					t.Run(fmt.Sprintf("%s/linked_%v/denied_%v/%s", format, linked, denied, kind), func(t *testing.T) {
						f := verifyFixture(t, tool, "audit", "none", format, linked)
						worktreePut(t, f.Root, f.Path, f.Raw+"\n## Audit\n\nNo findings. Checked "+f.Change.Added[0].ID+".\n")
						oid := f.Claim.Object
						if kind == "moved-head" {
							claimGit(t, f.Root, "-c", "user.name=Verify Locality", "-c", "user.email=verify@local.invalid", "commit", "--allow-empty", "-qm", "moved unavailable head")
							oid = claimBase(t, f.Root)
						} else if kind == "unrelated-ancestor" {
							oid = claimGit(t, f.Root, "rev-parse", f.Claim.Base+"~1")
						}
						marker, trace := verifyLocalMissing(t, f.Root, oid, denied)
						verifyReadOnly(t, f, func() {
							v, e := w.Verify(context.Background(), f.Root, "current")
							claimLocalNoFetch(t, marker, trace)
							if kind == "selected-claim" {
								if !reflect.DeepEqual(v, w.Verification{}) || e == nil || !strings.Contains(e.Error(), "P02") || !policyCorrective(e.Error()) {
									t.Error("required selected claim did not reject atomically", v, e)
								}
							} else {
								if e != nil {
									t.Error("local source unexpectedly unavailable", e)
								}
								verifyShape(t, f, v)
								if kind == "moved-head" {
									verifyCode(t, v, "P02", f.Path)
								} else {
									verifyNoErrors(t, v.Diagnostics)
								}
								if len(v.Coverage) != 0 || len(verifyCalls(t, f.Plan)) != 0 {
									t.Error("audit/rejected precondition invoked hooks")
								}
							}
						})
					})
				}
			}
		}
	}
}

func Test_FLOW047VerifyLiveConfiguredCallbacks(t *testing.T) {
	tool := verifyTool(t)
	for _, callback := range []string{"monitor", "conversion"} {
		for _, rejected := range []bool{false, true} {
			t.Run(fmt.Sprintf("%s/rejected_%v", callback, rejected), func(t *testing.T) {
				f := verifyFixture(t, tool, "impl", "none", "sha1", false)
				if callback == "conversion" {
					worktreePut(t, f.Root, ".gitattributes", "units/alpha/impl/old filter=verify\n")
					verifyRefreshFixture(t, &f)
				}
				marker := verifyForbiddenCommands(t, tool)
				if callback == "monitor" {
					claimGit(t, f.Root, "config", "core.fsmonitor", claimQuote(filepath.Join(filepath.Dir(marker), "codex")))
					_, _ = claimCommand(t, f.Root, "", "update-index", "--refresh")
				} else {
					command := claimQuote(filepath.Join(filepath.Dir(marker), "codex"))
					claimGit(t, f.Root, "config", "filter.verify.clean", command)
					_, _ = claimCommand(t, f.Root, "probe content\n", "hash-object", "--stdin", "--path=units/alpha/impl/old")
				}
				if _, e := os.Stat(marker); e != nil {
					t.Fatal("configured native callback control not live", e)
				}
				if e := os.Remove(marker); e != nil {
					t.Fatal(e)
				}
				worktreePut(t, f.Root, "units/alpha/impl/old", "RETAINED implementation\n")
				old := time.Unix(1, 0)
				if e := os.Chtimes(filepath.Join(f.Root, "units/alpha/impl/old"), old, old); e != nil {
					t.Fatal(e)
				}
				if rejected {
					worktreePut(t, f.Root, "private/forbidden", "rejected working edit")
				}
				verifyReadOnly(t, f, func() {
					v, e := w.Verify(context.Background(), f.Root, "current")
					if e != nil {
						t.Error(e)
					}
					verifyShape(t, f, v)
					if rejected {
						verifyCode(t, v, "P2-01", "private/forbidden")
						if len(verifyCalls(t, f.Plan)) != 0 {
							t.Error("rejected guard launched hooks")
						}
					} else {
						verifyNoErrors(t, v.Diagnostics)
						verifyExpectedCoverage(t, f, []string{"alpha"}, v)
					}
					if _, e := os.Stat(marker); !os.IsNotExist(e) {
						t.Error("Verify invoked configured monitor/conversion/agent/controller callback", e)
					}
				})
			})
		}
	}
}

func Test_FLOW044VerifyAuditHasNoAgentCalls(t *testing.T) {
	tool := verifyTool(t)
	for _, passing := range []bool{false, true} {
		t.Run(fmt.Sprint(passing), func(t *testing.T) {
			f := verifyFixture(t, tool, "audit", "all", "sha1", false)
			text := "- " + f.Change.Added[0].ID + ": \"MUST be checked\" — missing assertion"
			if passing {
				text = "No findings. Checked " + f.Change.Added[0].ID + "."
			}
			worktreePut(t, f.Root, f.Path, f.Raw+"\n## Audit\n\n"+text+"\n")
			marker := verifyForbiddenCommands(t, tool)
			v := verifyInvoke(t, f, "current")
			verifyNoErrors(t, v.Diagnostics)
			if _, e := os.Stat(marker); !os.IsNotExist(e) {
				t.Error("audit invoked an agent or controller", e)
			}
			if len(verifyCalls(t, f.Plan)) != 0 {
				t.Error("audit executed project hooks")
			}
		})
	}
}

func Test_FLOW046VerifyRejectsReplacedClaimToken(t *testing.T) {
	tool := verifyTool(t)
	f := verifyFixture(t, tool, "impl", "none", "sha1", false)
	replacement, e := w.UpdateClaim(f.Root, f.Claim, f.Claim.Base, f.Raw+"\n## Decisions\n\nReplacement claim state.\n")
	if e != nil || replacement.Object == f.Claim.Object {
		t.Fatal("replacement token control", replacement, e)
	}
	claimGit(t, f.Root, "update-ref", "refs/fireweed/claims/current", f.Claim.Object, replacement.Object)
	original, e := w.LoadClaim(f.Root, "current")
	if e != nil || original != f.Claim {
		t.Fatal("restored original-token premise", original, e)
	}
	args, _ := json.Marshal([]string{"update-ref", "refs/fireweed/claims/current", replacement.Object, f.Claim.Object})
	worktreePut(t, f.Plan, "alpha.gitargs", string(args))
	v := verifyInvoke(t, f, "current")
	verifyGuardEvidence(t, f, v)
	verifyCode(t, v, "P02", f.Path)
	after, e := w.LoadClaim(f.Root, "current")
	if e != nil || after != replacement {
		t.Error("verification erased/replaced the hook's winning claim", after, e)
	}
}
func Test_FLOW045VerifyDependentPolicyIsIndependent(t *testing.T) {
	tool := verifyTool(t)
	f := verifyFixture(t, tool, "impl", "all", "sha1", false)
	id := f.Change.Added[0].ID
	doc := f.Change.Document
	doc.Metadata["already_satisfied"] = []any{id}
	raw, e := c.RenderMarkdown(doc)
	if e != nil {
		t.Fatal(e)
	}
	verifyReseed(t, &f, raw)
	want := verifyModuleCoverage(t, f, []string{"alpha", "beta", "delta", "gamma"})
	v := verifyInvoke(t, f, "current")
	verifyNoErrors(t, v.Diagnostics)
	verifyExpectedCoverage(t, f, []string{"alpha", "beta", "delta", "gamma"}, v)
	for i, cov := range v.Coverage {
		if i < len(want) && !reflect.DeepEqual(verifyNormalizeCoverage(t, f, cov), want[i]) {
			t.Error("selected change policy leaked into dependent coverage", cov, want[i])
		}
	}
}
func Test_FLOW043VerifyApprovedTimeoutAndFatalHook(t *testing.T) {
	tool := verifyTool(t)
	for _, mode := range []string{"timeout", "launch"} {
		t.Run(mode, func(t *testing.T) {
			f := verifyFixture(t, tool, "evals", "all", "sha1", false)
			cfg, e := c.ParseMarkdown("---\n" + mustReadVerify(t, filepath.Join(f.Root, "fireweed.yaml")) + "---\n")
			if e != nil {
				t.Fatal(e)
			}
			cfg.Metadata["timeout"] = "700ms"
			cfgText, e := c.RenderMarkdown(cfg)
			if e != nil {
				t.Fatal(e)
			}
			cfgText = strings.TrimPrefix(cfgText, "---\n")
			cfgText = strings.TrimSuffix(strings.TrimSpace(cfgText), "---")
			worktreePut(t, f.Root, "fireweed.yaml", strings.TrimSpace(cfgText)+"\n")
			verifyRefreshFixture(t, &f)
			if mode == "timeout" {
				worktreePut(t, f.Plan, "alpha.hold", "3s")
			} else {
				doc := f.Change.Document
				set := doc.Metadata["set"].(map[string]any)
				set["hooks"] = map[string]any{"eval": claimQuote(filepath.Join(f.Plan, "missing-hook-executable"))}
				raw, e := c.RenderMarkdown(doc)
				if e != nil {
					t.Fatal(e)
				}
				verifyReseed(t, &f, raw)
			}
			if f.Snapshot.Config.Timeout != "700ms" {
				t.Fatal("approved timeout metadata premise", f.Snapshot.Config.Timeout)
			}
			want := verifyManualCoverage(t, f, []string{f.Change.Added[0].ID}, nil)
			fatal := false
			for _, d := range want.Diagnostics {
				fatal = fatal || d.Code == "C11"
			}
			if !fatal || len(want.Tests) != 0 || len(want.BySpec) != 0 {
				t.Fatal("fatal hook control not live", want)
			}
			v := verifyInvoke(t, f, "current")
			verifyExpectedCoverage(t, f, []string{"alpha"}, v)
			if len(v.Coverage) == 1 && !reflect.DeepEqual(verifyNormalizeCoverage(t, f, v.Coverage[0]), want) {
				t.Error("fatal hook evidence lost", v.Coverage[0], want)
			}
			verifyCode(t, v, "C11", "units/alpha")
			if len(v.Coverage) == 1 {
				if len(v.Coverage[0].Tests) != 0 || len(v.Coverage[0].BySpec) != 0 || v.Coverage[0].Diagnostics == nil {
					t.Error("fatal hook accepted partial test evidence", v.Coverage[0])
				}
				if mode == "timeout" && (!strings.Contains(v.Coverage[0].Output, "hook-ready:alpha") || v.Coverage[0].HookExit != -1) {
					t.Error("approved timeout did not retain actual failed invocation", v.Coverage[0])
				}
			}
			if mode == "timeout" {
				time.Sleep(3100 * time.Millisecond)
				if _, e := os.Stat(filepath.Join(f.Plan, "alpha.late")); !os.IsNotExist(e) {
					t.Error("timed-out hook survived")
				}
			}
		})
	}
}
func Test_FLOW044VerifyMixedFindingsAndDecoys(t *testing.T) {
	tool := verifyTool(t)
	f := verifyFixture(t, tool, "audit", "all", "sha1", false)
	id := f.Change.Added[0].ID
	good := "- " + id + ": \"MUST be checked\" — valid gap"
	invalid := []string{"- BET-001: \"MUST remain\" — foreign module", "- ALP-999: \"MUST be checked\" — unknown", "- ALP-003: \"MUST remain visible\" — review", "- ALP-004: \"MUST remain visible\" — quarantine", "- " + id + ": \"wrong clause\" — inexact"}
	audit := strings.Join(append(append([]string{invalid[0]}, good), invalid[1:]...), "\n")
	raw := f.Raw + "\n## Questions\n\n- ALP-006: \"MUST remain\" — outside audit\n\n## Audit\n\n" + audit + "\n"
	worktreePut(t, f.Root, f.Path, raw)
	v := verifyInvoke(t, f, "current")
	want := []w.Finding{{ID: id, Clause: "MUST be checked", Text: good}}
	if !reflect.DeepEqual(v.Findings, want) {
		t.Error("mixed findings were not filtered in order", v.Findings, want)
	}
	verifyNoErrors(t, v.Diagnostics)
}
func Test_FLOW042VerifyUnsafeEditEndpointRetainsEvidence(t *testing.T) {
	tool := verifyTool(t)
	f := verifyFixture(t, tool, "impl", "none", "sha1", false)
	outside := filepath.Join(t.TempDir(), "private")
	if e := os.WriteFile(outside, []byte("outside retained content"), 0600); e != nil {
		t.Fatal(e)
	}
	path := "units/alpha/impl/escape"
	if e := os.Symlink(outside, filepath.Join(f.Root, path)); e != nil {
		t.Fatal(e)
	}
	want := w.CheckWorktree(verifyInclusive(t, f), f.Module, f.Path, f.Claim)
	if len(want.Diagnostics) == 0 {
		t.Fatal("unsafe endpoint public guard control absent")
	}
	v := verifyInvoke(t, f, "current")
	verifyGuardEvidence(t, f, v)
	verifyCode(t, v, "P2-01", path)
	if len(v.Coverage) != 0 || len(verifyCalls(t, f.Plan)) != 0 {
		t.Error("unsafe endpoint launched hooks")
	}
	if got := mustReadVerify(t, outside); got != "outside retained content" {
		t.Error("outside contents changed")
	}
}
