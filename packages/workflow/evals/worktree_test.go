package evals

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"testing"

	k "github.com/lurkshark/fireweed/packages/checks"
	c "github.com/lurkshark/fireweed/packages/contracts"
	g "github.com/lurkshark/fireweed/packages/gitstore"
	p "github.com/lurkshark/fireweed/packages/project"
	w "github.com/lurkshark/fireweed/packages/workflow"
)

type worktreeFixtureData struct {
	Root, Path string
	Snapshot   p.Snapshot
	Module     c.Module
	Claim      w.Claim
}

func worktreePut(t *testing.T, root, path, text string) {
	t.Helper()
	full := filepath.Join(root, filepath.FromSlash(path))
	if err := os.MkdirAll(filepath.Dir(full), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(full, []byte(text), 0644); err != nil {
		t.Fatal(err)
	}
}
func worktreeFixture(t *testing.T, phase, format string, slugs ...string) worktreeFixtureData {
	t.Helper()
	root := claimRepo(t, format)
	slug := "current"
	if len(slugs) > 0 {
		slug = slugs[0]
	}
	path := "changes/0001-" + slug + ".md"
	config := `fireweed: '>=0.4, <0.5'
version: 1
base: main
modules: [units/*]
module_defaults:
  interface: [api.go]
  evals: [evals/**]
  disposable: [internal/**]
hooks:
  eval: hook-trap
  boundaries: hook-trap
protected: [private/**]
worker_writable: [go.sum, shared/**]
agent:
  command: [agent-trap]
`
	spec := `---
name: alpha
id_prefix: ALP
preserve: [keep/**]
append_only: [migrations/**]
---

## Specs

### ALP-001
The old behavior MUST remain retained.

## Decisions
`
	queue := `---
module: alpha
---

## Motivation

Check the approved live worktree.

## Interface

Describe the retained contract edit.

## Added

### new:case
The new behavior MUST obey the approved workflow.
`
	for path, text := range map[string]string{"fireweed.yaml": config, "units/a/SPEC.md": spec, "units/b/SPEC.md": strings.ReplaceAll(strings.ReplaceAll(spec, "alpha", "beta"), "ALP", "BET"), path: queue, ".gitignore": "ignored\n", "ignored": "retained ignored bytes", "units/a/internal/existing.go": "distinct existing implementation bytes must survive\n", "units/a/evals/frozen.txt": "distinct frozen evaluation bytes must survive\n", "units/a/api.go": "retained public interface bytes\n", "units/a/keep/owned.txt": "retained preserved bytes\n", "units/a/migrations/001.sql": "append-only bytes\n", "go.sum": "shared lock bytes\n"} {
		worktreePut(t, root, path, text)
	}
	claimGit(t, root, "add", ".")
	claimGit(t, root, "-c", "user.name=Worktree Fixture", "-c", "user.email=worktree@example.invalid", "commit", "-m", "authored queue")
	state := strings.Replace(queue, "module: alpha\n", "module: alpha\nrun: approved-run\nphase: "+phase+"\n", 1)
	state = strings.ReplaceAll(state, "new:case", "ALP-002")
	worktreePut(t, root, path, state)
	claim := w.Claim{Worktree: root, Change: slug, Run: "approved-run", Base: claimBase(t, root), State: state}
	claim.Object = claimRaw(t, root, slug, claimPayload(t, claim))
	snapshot, err := p.Load(root)
	if err != nil {
		t.Fatal(err)
	}
	if ds := p.Check(snapshot); len(ds) != 0 {
		t.Fatalf("validated snapshot premise %+v", ds)
	}
	var module c.Module
	for _, m := range snapshot.Modules {
		if m.Name == "alpha" {
			module = m
		}
	}
	parsed, err := c.ParseChange(state, path, module.IDPrefix)
	if err != nil {
		t.Fatal(err)
	}
	module, err = p.ApplyChange(module, parsed)
	if err != nil {
		t.Fatal(err)
	}
	return worktreeFixtureData{root, path, snapshot, module, claim}
}
func worktreeEdits(t *testing.T, f worktreeFixtureData) []g.Edit {
	t.Helper()
	edits, err := g.Edits(f.Root, f.Claim.Base)
	if err != nil {
		t.Fatal(err)
	}
	return edits
}
func worktreeShape(t *testing.T, got w.WorktreeCheck, edits []g.Edit) {
	t.Helper()
	if got.Edits == nil || got.Diagnostics == nil {
		t.Fatalf("null worktree result %+v", got)
	}
	if !reflect.DeepEqual(got.Edits, edits) {
		t.Fatalf("edits %+v; want actual public Git records %+v", got.Edits, edits)
	}
	if !sort.SliceIsSorted(got.Diagnostics, func(i, j int) bool { return policyLess(got.Diagnostics[i], got.Diagnostics[j]) }) {
		t.Fatal("unsorted diagnostics")
	}
	seen := map[k.Diagnostic]bool{}
	for _, d := range got.Diagnostics {
		if d.Severity != "error" || !policyCorrective(d.Message) || seen[d] {
			t.Fatalf("non-corrective/duplicate diagnostic %+v", d)
		}
		seen[d] = true
	}
}
func worktreePublicDiagnostics(t *testing.T, f worktreeFixtureData, current string, edits []g.Edit) []k.Diagnostic {
	t.Helper()
	parsed, err := c.ParseChange(f.Claim.State, f.Path, f.Module.IDPrefix)
	if err != nil {
		t.Fatal(err)
	}
	expected := append(w.CheckState(f.Module, f.Claim.State, current, f.Path), w.CheckEdits(f.Snapshot, f.Module, parsed, edits)...)
	sort.Slice(expected, func(i, j int) bool { return policyLess(expected[i], expected[j]) })
	return expected
}
func worktreeCheckPublic(t *testing.T, got w.WorktreeCheck, edits []g.Edit, expected []k.Diagnostic, extraPaths ...string) {
	t.Helper()
	worktreeShape(t, got, edits)
	allowed := map[string]bool{}
	for _, path := range extraPaths {
		allowed[path] = true
	}
	remaining := append([]k.Diagnostic{}, got.Diagnostics...)
	for _, d := range expected {
		index := -1
		for i, actual := range remaining {
			if actual == d {
				index = i
				break
			}
		}
		if index < 0 {
			t.Fatalf("lost public composition diagnostic %+v in %+v", d, got.Diagnostics)
		}
		remaining = append(remaining[:index], remaining[index+1:]...)
	}
	extraSeen := map[string]bool{}
	for _, d := range remaining {
		if d.Code != "P2-01" || !allowed[d.Path] {
			t.Fatalf("unexpected extra diagnostic %+v", d)
		}
		extraSeen[d.Path] = true
	}
	for path := range allowed {
		if !extraSeen[path] {
			t.Fatalf("missing path safety/read rejection at %s", path)
		}
	}
}
func worktreeReseed(t *testing.T, f *worktreeFixtureData) {
	t.Helper()
	f.Claim.Object = claimRaw(t, f.Root, f.Claim.Change, claimPayload(t, f.Claim))
}
func worktreePrecondition(t *testing.T, f worktreeFixtureData) {
	t.Helper()
	got := w.CheckWorktree(f.Snapshot, f.Module, f.Path, f.Claim)
	worktreeShape(t, got, []g.Edit{})
	policyAssert(t, got.Diagnostics, policyWant(f.Path, "P02"))
}
func Test_FLOW012ValidClaimedRoots(t *testing.T) {
	for _, format := range []string{"sha1", "sha256"} {
		for _, location := range []string{"root", "nested", "peer"} {
			t.Run(format+"_"+location, func(t *testing.T) {
				f := worktreeFixture(t, "impl", format)
				if location == "nested" {
					nested := filepath.Join(f.Root, "nested")
					if err := os.Mkdir(nested, 0755); err != nil {
						t.Fatal(err)
					}
					f.Snapshot.Root = nested
				}
				if location == "peer" {
					peer := claimPeer(t, f.Root)
					worktreePut(t, peer, f.Path, f.Claim.State)
					f.Root = peer
					f.Snapshot.Root = peer
					f.Claim.Worktree = peer
					worktreeReseed(t, &f)
				}
				edits := worktreeEdits(t, f)
				got := w.CheckWorktree(f.Snapshot, f.Module, f.Path, f.Claim)
				worktreeCheckPublic(t, got, edits, nil)
			})
		}
	}
}
func Test_FLOW012ClaimHeadAndApprovedPreconditions(t *testing.T) {
	for _, name := range []string{"missing_claim", "stale_token", "tampered_field", "foreign_owner", "moved_head", "no_repository", "invalid_state", "missing_run", "wrong_phase", "wrong_module", "wrong_run", "wrong_slug", "wrong_prefix", "archive_path", "absolute_path", "parent_path", "nested_path"} {
		t.Run(name, func(t *testing.T) {
			f := worktreeFixture(t, "impl", "sha1")
			switch name {
			case "missing_claim":
				claimGit(t, f.Root, "update-ref", "-d", "refs/fireweed/claims/current", f.Claim.Object)
			case "stale_token":
				f.Claim.Object = strings.Repeat("f", len(f.Claim.Object))
			case "tampered_field":
				f.Claim.State += "worker supplied forged baseline"
			case "foreign_owner":
				f.Snapshot.Root = claimPeer(t, f.Root)
			case "moved_head":
				claimGit(t, f.Root, "-c", "user.name=Worktree Fixture", "-c", "user.email=worktree@example.invalid", "commit", "--allow-empty", "-m", "forbidden worker commit")
			case "no_repository":
				f.Snapshot.Root = t.TempDir()
			case "invalid_state":
				f.Claim.State = "not a change"
				worktreeReseed(t, &f)
			case "missing_run":
				f.Claim.State = strings.Replace(f.Claim.State, "run: approved-run\n", "", 1)
				worktreeReseed(t, &f)
			case "wrong_phase":
				f.Claim.State = strings.Replace(f.Claim.State, "phase: impl", "phase: applied", 1)
				worktreeReseed(t, &f)
			case "wrong_module":
				f.Module.Name = "beta"
			case "wrong_run":
				f.Claim.State = strings.Replace(f.Claim.State, "run: approved-run", "run: worker-run", 1)
				worktreeReseed(t, &f)
			case "wrong_slug":
				f.Path = "changes/0001-other.md"
			case "wrong_prefix":
				f.Module.IDPrefix = "BET"
			case "archive_path":
				f.Path = "changes/archive/0001-current.md"
			case "absolute_path":
				f.Path = filepath.Join(f.Root, f.Path)
			case "parent_path":
				f.Path = "changes/../changes/0001-current.md"
			case "nested_path":
				f.Path = "changes/nested/0001-current.md"
			}
			worktreePrecondition(t, f)
		})
	}
}
func Test_FLOW013CurrentIntentComposition(t *testing.T) {
	for _, phase := range []string{"evals", "audit", "impl"} {
		for _, name := range []string{"unchanged", "permitted", "blocked", "metadata", "body", "metadata_and_body", "invalid_grammar", "phase_escalation"} {
			t.Run(phase+"_"+name, func(t *testing.T) {
				f := worktreeFixture(t, phase, "sha1")
				current := f.Claim.State
				switch name {
				case "permitted":
					current += "\n## Questions\n\nAsk only an allowed question.\n"
					if phase == "audit" {
						current += "\n## Audit\n\nNo findings. Checked ALP-002.\n"
					} else {
						current += "\n## Decisions\n\nRecord the allowed choice.\n\n## Notes\n\nWorker notes.\n"
					}
				case "blocked":
					current = strings.Replace(current, "module: alpha\n", "module: alpha\nblocked: Needs clarification\n", 1)
				case "metadata":
					current = strings.Replace(current, "run: approved-run", "run: worker-run", 1)
				case "body":
					current = strings.Replace(current, "Check the approved live worktree.", "Worker changed protected motivation.", 1)
				case "metadata_and_body":
					current = strings.Replace(strings.Replace(current, "run: approved-run", "run: worker-run", 1), "Check the approved live worktree.", "Worker changed protected motivation.", 1)
				case "invalid_grammar":
					current = "invalid current change"
				case "phase_escalation":
					current = strings.Replace(current, "phase: "+phase, "phase: impl", 1) + "\n## Decisions\n\nAn auditor cannot grant this permission.\n"
				}
				worktreePut(t, f.Root, f.Path, current)
				edits := worktreeEdits(t, f)
				expected := worktreePublicDiagnostics(t, f, current, edits)
				got := w.CheckWorktree(f.Snapshot, f.Module, f.Path, f.Claim)
				worktreeCheckPublic(t, got, edits, expected)
			})
		}
	}
}
func Test_FLOW013MissingAndLinkedCurrentFiles(t *testing.T) {
	for _, name := range []string{"missing", "directory", "unreadable", "outside_link", "dangling_link", "inside_link"} {
		t.Run(name, func(t *testing.T) {
			f := worktreeFixture(t, "impl", "sha1")
			full := filepath.Join(f.Root, f.Path)
			switch name {
			case "missing":
				if err := os.Remove(full); err != nil {
					t.Fatal(err)
				}
			case "directory":
				if err := os.Remove(full); err != nil {
					t.Fatal(err)
				}
				if err := os.Mkdir(full, 0755); err != nil {
					t.Fatal(err)
				}
			case "unreadable":
				if os.Geteuid() == 0 {
					t.Skip("read-permission fixture requires unprivileged user")
				}
				if err := os.Chmod(full, 0000); err != nil {
					t.Fatal(err)
				}
				defer os.Chmod(full, 0644)
			case "outside_link", "dangling_link":
				outside := filepath.Join(t.TempDir(), "private")
				if name == "outside_link" {
					if err := os.WriteFile(outside, []byte("private MUST NOT be read"), 0644); err != nil {
						t.Fatal(err)
					}
				}
				if err := os.Remove(full); err != nil {
					t.Fatal(err)
				}
				if err := os.Symlink(outside, full); err != nil {
					t.Skipf("symlink fixture unavailable: %v", err)
				}
			case "inside_link":
				worktreePut(t, f.Root, "shared/approved.md", f.Claim.State)
				if err := os.Remove(full); err != nil {
					t.Fatal(err)
				}
				if err := os.Symlink(filepath.Join(f.Root, "shared/approved.md"), full); err != nil {
					t.Skipf("symlink fixture unavailable: %v", err)
				}
			}
			edits := worktreeEdits(t, f)
			parsed, err := c.ParseChange(f.Claim.State, f.Path, f.Module.IDPrefix)
			if err != nil {
				t.Fatal(err)
			}
			expected := w.CheckEdits(f.Snapshot, f.Module, parsed, edits)
			got := w.CheckWorktree(f.Snapshot, f.Module, f.Path, f.Claim)
			if name == "inside_link" {
				worktreeCheckPublic(t, got, edits, expected)
			} else {
				worktreeCheckPublic(t, got, edits, expected, f.Path)
			}
		})
	}
}
func Test_FLOW014ActualEditsAndPhasePolicy(t *testing.T) {
	for _, row := range []struct{ phase, path, status string }{
		{"evals", "units/a/evals/new.txt", "A"}, {"evals", "units/a/api.go", "M"}, {"evals", "units/a/internal/stub.go", "A"}, {"evals", "units/a/internal/existing.go", "M"},
		{"impl", "units/a/internal/existing.go", "M"}, {"impl", "units/a/keep/owned.txt", "M"}, {"impl", "units/a/evals/frozen.txt", "M"}, {"impl", "units/a/api.go", "M"}, {"impl", "units/a/migrations/001.sql", "M"},
		{"impl", "private/tooling", "A"}, {"impl", "fireweed.yaml", "M"}, {"impl", "units/a/SPEC.md", "M"}, {"impl", "units/b/internal/file.go", "A"}, {"impl", "changes/0002-other.md", "A"}, {"impl", "go.sum", "M"}, {"impl", "shared/lock", "A"}, {"audit", "units/a/internal/existing.go", "M"},
	} {
		t.Run(row.phase+"_"+row.path, func(t *testing.T) {
			f := worktreeFixture(t, row.phase, "sha1")
			worktreePut(t, f.Root, row.path, "worker bytes")
			edits := worktreeEdits(t, f)
			found := false
			for _, e := range edits {
				found = found || (e.Status == row.status && e.Path == row.path)
			}
			if !found {
				t.Fatalf("fixture did not produce %s %s: %+v", row.status, row.path, edits)
			}
			expected := worktreePublicDiagnostics(t, f, f.Claim.State, edits)
			worktreeCheckPublic(t, w.CheckWorktree(f.Snapshot, f.Module, f.Path, f.Claim), edits, expected)
		})
	}
}
func Test_FLOW014RenameSourcesAndStagedUnstagedUntracked(t *testing.T) {
	for _, staged := range []bool{false, true} {
		t.Run(fmt.Sprint(staged), func(t *testing.T) {
			f := worktreeFixture(t, "impl", "sha1")
			from, to := "units/a/evals/frozen.txt", "units/a/internal/moved.txt"
			if err := os.Rename(filepath.Join(f.Root, from), filepath.Join(f.Root, to)); err != nil {
				t.Fatal(err)
			}
			if staged {
				claimGit(t, f.Root, "add", "units/a")
			}
			worktreePut(t, f.Root, "units/a/internal/untracked.go", "untracked bytes")
			worktreePut(t, f.Root, "units/a/keep/owned.txt", "unstaged preserved bytes")
			edits := worktreeEdits(t, f)
			want := g.Edit{Status: "D", Path: from}
			if staged {
				want = g.Edit{Status: "R", OldPath: from, Path: to}
			}
			found := false
			for _, e := range edits {
				found = found || e == want
			}
			if !found {
				t.Fatalf("rename fixture %+v lacks %+v", edits, want)
			}
			expected := worktreePublicDiagnostics(t, f, f.Claim.State, edits)
			worktreeCheckPublic(t, w.CheckWorktree(f.Snapshot, f.Module, f.Path, f.Claim), edits, expected)
		})
	}
}
func Test_FLOW014UnsafeEditEndpoints(t *testing.T) {
	for _, name := range []string{"untracked_link", "renamed_link", "directory_escape"} {
		t.Run(name, func(t *testing.T) {
			f := worktreeFixture(t, "impl", "sha1")
			outside := t.TempDir()
			worktreePut(t, outside, "private", "private contents")
			unsafe := []string{}
			switch name {
			case "untracked_link":
				path := "units/a/internal/escape"
				if err := os.Symlink(filepath.Join(outside, "private"), filepath.Join(f.Root, path)); err != nil {
					t.Skipf("symlink fixture unavailable: %v", err)
				}
				unsafe = append(unsafe, path)
			case "renamed_link":
				path := "units/a/internal/escape"
				if err := os.Symlink(filepath.Join(outside, "private"), filepath.Join(f.Root, path)); err != nil {
					t.Skipf("symlink fixture unavailable: %v", err)
				}
				claimGit(t, f.Root, "add", path)
				claimGit(t, f.Root, "-c", "user.name=Worktree Fixture", "-c", "user.email=worktree@example.invalid", "commit", "-m", "fixture unsafe retained link")
				f.Claim.Base = claimBase(t, f.Root)
				worktreeReseed(t, &f)
				to := "units/a/internal/moved-link"
				if err := os.Rename(filepath.Join(f.Root, path), filepath.Join(f.Root, to)); err != nil {
					t.Fatal(err)
				}
				claimGit(t, f.Root, "add", "units/a/internal")
				unsafe = append(unsafe, to)
			case "directory_escape":
				if err := os.RemoveAll(filepath.Join(f.Root, "units/a/internal")); err != nil {
					t.Fatal(err)
				}
				if err := os.Symlink(outside, filepath.Join(f.Root, "units/a/internal")); err != nil {
					t.Skipf("symlink fixture unavailable: %v", err)
				}
				unsafe = append(unsafe, "units/a/internal", "units/a/internal/existing.go")
			}
			edits := worktreeEdits(t, f)
			expected := worktreePublicDiagnostics(t, f, f.Claim.State, edits)
			worktreeCheckPublic(t, w.CheckWorktree(f.Snapshot, f.Module, f.Path, f.Claim), edits, expected, unsafe...)
		})
	}
}
func Test_FLOW012_FLOW013_FLOW014_FLOW015WorktreeFixtureControls(t *testing.T) {
	for _, format := range []string{"sha1", "sha256"} {
		t.Run(format, func(t *testing.T) {
			f := worktreeFixture(t, "evals", format)
			parsed, err := c.ParseChange(f.Claim.State, f.Path, f.Module.IDPrefix)
			if err != nil || parsed.Run != f.Claim.Run || parsed.Slug != f.Claim.Change || parsed.Phase != "evals" || len(f.Module.Specs) != 2 {
				t.Fatalf("approved fixture %+v %v", parsed, err)
			}
			edits := worktreeEdits(t, f)
			if !reflect.DeepEqual(edits, []g.Edit{{Status: "M", Path: f.Path}}) {
				t.Fatalf("started-state-only fixture %+v", edits)
			}
			stored, err := w.LoadClaim(f.Root, f.Claim.Change)
			if err != nil || stored != f.Claim {
				t.Fatalf("known existing claim contract fixture %+v %v", stored, err)
			}
			if ds := worktreePublicDiagnostics(t, f, f.Claim.State, edits); len(ds) != 0 {
				t.Fatalf("unchanged composition premise %+v", ds)
			}
		})
	}
}

func Test_FLOW015ResultAndInputOwnership(t *testing.T) {
	f := worktreeFixture(t, "impl", "sha1")
	current := strings.Replace(strings.Replace(f.Claim.State, "run: approved-run", "run: worker-run", 1), "Check the approved live worktree.", "Protected motivation changed.", 1)
	worktreePut(t, f.Root, f.Path, current)
	worktreePut(t, f.Root, "units/a/migrations/001.sql", "worker changed append-only bytes")
	edits := worktreeEdits(t, f)
	expected := worktreePublicDiagnostics(t, f, current, edits)
	before := policyJSON(t, []any{f.Snapshot, f.Module})
	claim := f.Claim
	first := w.CheckWorktree(f.Snapshot, f.Module, f.Path, f.Claim)
	second := w.CheckWorktree(f.Snapshot, f.Module, f.Path, f.Claim)
	worktreeCheckPublic(t, first, edits, expected)
	worktreeCheckPublic(t, second, edits, expected)
	saved := policyJSON(t, second)
	first.Edits[0].Path = "mutated"
	first.Diagnostics[0].Message = "mutated"
	first.Edits = append(first.Edits, g.Edit{Path: "new"})
	first.Diagnostics = append(first.Diagnostics, k.Diagnostic{Message: "new"})
	if before != policyJSON(t, []any{f.Snapshot, f.Module}) || f.Claim != claim || saved != policyJSON(t, second) {
		t.Fatal("input or prior result aliases")
	}
	// Commit the fixture's started text as setup and refresh its independently stored baseline.
	worktreePut(t, f.Root, f.Path, f.Claim.State)
	claimGit(t, f.Root, "checkout", "--", "units/a/migrations/001.sql")
	claimGit(t, f.Root, "add", f.Path)
	claimGit(t, f.Root, "-c", "user.name=Worktree Fixture", "-c", "user.email=worktree@example.invalid", "commit", "-m", "fixture resumed started state")
	f.Claim.Base = claimBase(t, f.Root)
	worktreeReseed(t, &f)
	a, b := w.CheckWorktree(f.Snapshot, f.Module, f.Path, f.Claim), w.CheckWorktree(f.Snapshot, f.Module, f.Path, f.Claim)
	worktreeShape(t, a, []g.Edit{})
	worktreeShape(t, b, []g.Edit{})
	policyAssert(t, a.Diagnostics, nil)
	policyAssert(t, b.Diagnostics, nil)
	a.Edits = append(a.Edits, g.Edit{Path: "a"})
	b.Edits = append(b.Edits, g.Edit{Path: "b"})
	a.Diagnostics = append(a.Diagnostics, k.Diagnostic{Message: "a"})
	b.Diagnostics = append(b.Diagnostics, k.Diagnostic{Message: "b"})
	if a.Edits[0].Path != "a" || b.Edits[0].Path != "b" || a.Diagnostics[0].Message != "a" || b.Diagnostics[0].Message != "b" {
		t.Fatal("writable empty results alias")
	}
}
func worktreeTraps(t *testing.T, denyCollection bool) (string, string) {
	t.Helper()
	traps := t.TempDir()
	marker := filepath.Join(traps, "forbidden")
	realGit, err := exec.LookPath("git")
	if err != nil {
		t.Fatal(err)
	}
	deny := "fetch|push|clone|ls-remote|commit|branch|checkout|switch|reset|restore|add|worktree|update-ref|hash-object"
	if denyCollection {
		deny += "|diff|ls-files"
	}
	script := "#!/bin/sh\nfor a in \"$@\"; do case \"$a\" in " + deny + ") printf attempted >> " + claimQuote(marker) + ";exit 97;;esac;done\nexec " + claimQuote(realGit) + " \"$@\"\n"
	if err := os.WriteFile(filepath.Join(traps, "git"), []byte(script), 0755); err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"hook-trap", "agent-trap", "codex", "claude", "fireweed", "fireweed-bootstrap", "curl", "wget", "ssh", "nc"} {
		if err := os.WriteFile(filepath.Join(traps, name), []byte("#!/bin/sh\nprintf attempted >> "+claimQuote(marker)+"\nexit 97\n"), 0755); err != nil {
			t.Fatal(err)
		}
	}
	t.Setenv("PATH", traps+string(os.PathListSeparator)+os.Getenv("PATH"))
	t.Setenv("FIREWEED_BIN", filepath.Join(traps, "fireweed-bootstrap"))
	if err := exec.Command("git", "fetch", "forbidden.invalid").Run(); err == nil {
		t.Fatal("Git/network trap control unexpectedly passed")
	}
	if _, err := os.Stat(marker); err != nil {
		t.Fatal("positive command trap control did not execute")
	}
	if err := os.Remove(marker); err != nil {
		t.Fatal(err)
	}
	return traps, marker
}
func Test_FLOW015InspectionSafetyAndPreconditionShortCircuit(t *testing.T) {
	if os.PathSeparator == '\\' {
		t.Skip("POSIX command-trap fixture; ordinary worktree checks remain portable")
	}
	for _, mode := range []string{"success", "state_error", "zone_error", "precondition_error", "read_error"} {
		t.Run(mode, func(t *testing.T) {
			f := worktreeFixture(t, "impl", "sha1")
			if mode == "state_error" {
				worktreePut(t, f.Root, f.Path, strings.Replace(f.Claim.State, "run: approved-run", "run: worker", 1))
			}
			if mode == "zone_error" {
				worktreePut(t, f.Root, "units/a/evals/frozen.txt", "forbidden worker bytes")
			}
			if mode == "precondition_error" {
				f.Claim.Object = strings.Repeat("f", len(f.Claim.Object))
			}
			if mode == "read_error" {
				if err := os.Remove(filepath.Join(f.Root, f.Path)); err != nil {
					t.Fatal(err)
				}
			}
			_, marker := worktreeTraps(t, mode == "precondition_error")
			temporary := stateCanonicalDirectory(t, t.TempDir())
			t.Setenv("TMPDIR", temporary)
			t.Setenv("TMP", temporary)
			t.Setenv("TEMP", temporary)
			claimGit(t, f.Root, "--version")
			beforeTemporary := claimTree(t, temporary, func(string) bool { return false })
			common := claimCommon(t, f.Root)
			files := claimTree(t, f.Root, func(path string) bool { return path == ".git" })
			metadata := claimTree(t, common, func(string) bool { return false })
			env := os.Environ()
			cwd, err := os.Getwd()
			if err != nil {
				t.Fatal(err)
			}
			inputs := policyJSON(t, []any{f.Snapshot, f.Module})
			claim := f.Claim
			edits := []g.Edit{{Status: "M", Path: f.Path}}
			if mode == "precondition_error" {
				edits = []g.Edit{}
			}
			if mode == "read_error" {
				edits[0].Status = "D"
			}
			if mode == "zone_error" {
				edits = append(edits, g.Edit{Status: "M", Path: "units/a/evals/frozen.txt"})
			}
			got := w.CheckWorktree(f.Snapshot, f.Module, f.Path, f.Claim)
			worktreeShape(t, got, edits)
			if mode == "success" && len(got.Diagnostics) != 0 {
				t.Fatalf("unexpected diagnostics %+v", got.Diagnostics)
			}
			if mode != "success" && len(got.Diagnostics) == 0 {
				t.Fatal("unsafe state accepted")
			}
			after, err := os.Getwd()
			if err != nil || after != cwd || !reflect.DeepEqual(env, os.Environ()) || inputs != policyJSON(t, []any{f.Snapshot, f.Module}) || f.Claim != claim || !reflect.DeepEqual(files, claimTree(t, f.Root, func(path string) bool { return path == ".git" })) || !reflect.DeepEqual(metadata, claimTree(t, common, func(string) bool { return false })) {
				t.Fatal("inspection mutated parent inputs, index/HEAD/refs, user/ignored files or Git storage")
			}
			if _, err := os.Stat(marker); !os.IsNotExist(err) {
				t.Fatal("inspection executed a forbidden command or collected edits after a bad claim")
			}
			if !reflect.DeepEqual(beforeTemporary, claimTree(t, temporary, func(string) bool { return false })) {
				t.Fatal("inspection leaked or changed dedicated temporary storage")
			}
		})
	}
}

func Test_FLOW014CollectionFailure(t *testing.T) {
	if os.PathSeparator == '\\' {
		t.Skip("POSIX Git failure injection")
	}
	f := worktreeFixture(t, "impl", "sha1")
	_, marker := worktreeTraps(t, true)
	got := w.CheckWorktree(f.Snapshot, f.Module, f.Path, f.Claim)
	worktreeShape(t, got, []g.Edit{})
	policyAssert(t, got.Diagnostics, policyWant(f.Path, "P02"))
	if _, err := os.Stat(marker); err != nil {
		t.Fatal("collection-failure fixture was never reached")
	}
}
func Test_FLOW013_FLOW014IndependentViolationsCompose(t *testing.T) {
	f := worktreeFixture(t, "impl", "sha1")
	current := strings.Replace(strings.Replace(f.Claim.State, "run: approved-run", "run: worker", 1), "Check the approved live worktree.", "Worker changed protected motivation.", 1)
	worktreePut(t, f.Root, f.Path, current)
	worktreePut(t, f.Root, "units/a/migrations/001.sql", "forbidden mutation")
	private := filepath.Join(t.TempDir(), "private")
	if err := os.WriteFile(private, []byte("private retained bytes"), 0644); err != nil {
		t.Fatal(err)
	}
	path := "units/a/internal/unsafe"
	if err := os.Symlink(private, filepath.Join(f.Root, path)); err != nil {
		t.Skipf("symlink fixture unavailable: %v", err)
	}
	edits := worktreeEdits(t, f)
	expected := worktreePublicDiagnostics(t, f, current, edits)
	got := w.CheckWorktree(f.Snapshot, f.Module, f.Path, f.Claim)
	worktreeCheckPublic(t, got, edits, expected, path)
}
