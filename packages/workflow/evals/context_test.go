package evals

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"testing"

	a "github.com/lurkshark/fireweed/packages/assets"
	c "github.com/lurkshark/fireweed/packages/contracts"
	g "github.com/lurkshark/fireweed/packages/gitstore"
	p "github.com/lurkshark/fireweed/packages/project"
	w "github.com/lurkshark/fireweed/packages/workflow"
)

type contextFixtureData struct {
	Root, Path, Raw, Spec, System string
	Snapshot                      p.Snapshot
	Expected                      map[string]string
}

func contextFixture(t *testing.T, phase, format string) contextFixtureData {
	t.Helper()
	root := claimRepo(t, format)
	config := `version: 1
fireweed: '>=0.4, <0.5'
base: main
modules: [units/*]
module_defaults:
  evals: [e/**]
  interface: [api/**]
  disposable: [impl/**]
  preserve: [keep/**]
  append_only: [append/**]
hooks:
  eval: context-hook-trap
  boundaries: context-hook-trap
protected: [private/**, units/a/e/protected.txt]
worker_writable: [shared/**]
agent:
  command: [context-agent-trap]
`
	spec := func(name, prefix, dep string) string {
		return "---\nname: " + name + "\nid_prefix: " + prefix + "\n" + dep + "---\n\n# exact " + name + "\n\n## Specs\n\n### " + prefix + "-001\nThe retained behavior MUST stay.\n\n## Decisions\n\noriginal choice\n"
	}
	specs := map[string]string{"units/a/SPEC.md": spec("alpha", "ALP", "depends_on: [beta, gamma]\n"), "units/b/SPEC.md": spec("beta", "BET", "depends_on: [delta]\n"), "units/c/SPEC.md": spec("gamma", "GAM", ""), "units/d/SPEC.md": spec("delta", "DEL", ""), "units/x/SPEC.md": spec("extra", "EXT", "")}
	raw := "---\nmodule: alpha\nset:\n  interface: [pub/**]\n  preserve: [saved/**]\n  depends_on: [gamma, beta]\n  hooks:\n    eval: context-hook-trap\n    setup: context-hook-trap\n"
	if phase != "" {
		raw += "run: context-run\nphase: " + phase + "\n"
	}
	raw += "---\n\n## Motivation\n\nBuild exact worker context.\n\n## Interface\n\nDeclare fresh public contracts.\n\n## Added\n\n### ALP-002\nThe current behavior MUST be projected.\n"
	before := "---\nmodule: beta\nset:\n  interface: [future/**]\n---\n\n## Motivation\n\nPrecede target.\n\n## Added\n\n### BET-002\nThe preceding behavior MUST be included.\n"
	later := "---\nmodule: alpha\nset:\n  interface: [later/**]\n---\n\n## Motivation\n\nFollow target.\n\n## Added\n\n### ALP-003\nThe later behavior MUST stay excluded.\n"
	path := "changes/0002-current.md"
	system := "project exact instructions λ\r\nno final newline"
	files := map[string]string{"fireweed.yaml": config, "SYSTEM.md": system, path: raw, "changes/0001-before.md": before, "changes/0003-later.md": later, ".gitignore": "ignored*\n", "units/a/e/current.txt": "eval bytes\r\n\x00", "units/a/pub/contract ' λ.txt": "projected interface\r\n", "units/a/saved/keep.txt": "projected preserved", "units/a/append/001.txt": "append-only", "units/c/api/gamma.txt": "first direct interface", "units/b/future/beta.txt": "preceding direct interface", "units/a/api/old.txt": "old interface excluded", "units/a/keep/old.txt": "old preserved excluded", "units/a/later/future.txt": "later interface excluded", "units/a/impl/private.txt": "implementation excluded", "units/a/manifest.txt": "other excluded", "units/b/api/old.txt": "old dependency interface excluded", "units/b/e/dependency.txt": "dependency eval excluded", "units/b/keep/private.txt": "dependency preserved excluded", "units/b/impl/private.txt": "dependency impl excluded", "units/d/api/transitive.txt": "transitive excluded", "units/x/api/unrelated.txt": "unrelated excluded", "units/a/e/protected.txt": "protected eval excluded", "private/secret.txt": "private excluded", "shared/contract.txt": "worker writable excluded"}
	for name, text := range specs {
		files[name] = text
	}
	for name, text := range files {
		worktreePut(t, root, name, text)
	}
	worktreePut(t, root, "units/a/e/ignored-tracked.txt", "tracked ignored still included")
	claimGit(t, root, "add", ".")
	claimGit(t, root, "add", "-f", "units/a/e/ignored-tracked.txt")
	claimGit(t, root, "-c", "user.name=Context Fixture", "-c", "user.email=context@example.invalid", "commit", "-m", "context fixture")
	worktreePut(t, root, "units/a/e/ignored-new.txt", "ignored untracked excluded")
	worktreePut(t, root, "units/a/e/new stub.txt", "untracked retained included")
	snapshot, err := p.Load(root)
	if err != nil {
		t.Fatal(err)
	}
	if ds := p.Check(snapshot); len(ds) != 0 {
		t.Fatalf("snapshot premise %+v", ds)
	}
	expected := map[string]string{}
	for _, name := range []string{"units/a/e/current.txt", "units/a/e/ignored-tracked.txt", "units/a/e/new stub.txt", "units/a/pub/contract ' λ.txt", "units/a/saved/keep.txt", "units/a/append/001.txt", "units/c/api/gamma.txt", "units/b/future/beta.txt"} {
		if b, e := os.ReadFile(filepath.Join(root, name)); e != nil {
			t.Fatal(e)
		} else {
			expected[name] = string(b)
		}
	}
	return contextFixtureData{root, path, raw, specs["units/a/SPEC.md"], system, snapshot, expected}
}
func contextChange(t *testing.T, f contextFixtureData) c.Change {
	t.Helper()
	for _, ch := range f.Snapshot.Changes {
		if ch.Slug == "current" {
			return ch
		}
	}
	t.Fatal("no fixture change")
	return c.Change{}
}
func contextExpectedFiles(f contextFixtureData, phase string) []w.FileContent {
	names := []string{}
	for name := range f.Expected {
		if phase != "audit" || strings.HasPrefix(name, "units/a/e/") || strings.HasPrefix(name, "units/a/pub/") {
			names = append(names, name)
		}
	}
	sort.Strings(names)
	out := []w.FileContent{}
	for _, name := range names {
		out = append(out, w.FileContent{Path: name, Content: f.Expected[name]})
	}
	return out
}
func contextError(t *testing.T, f contextFixtureData, slug, code, path string) {
	t.Helper()
	got, err := w.BuildContext(f.Snapshot, slug)
	if !reflect.DeepEqual(got, w.WorkerContext{}) {
		t.Fatalf("partial context on rejection: %+v", got)
	}
	if err == nil || !strings.Contains(err.Error(), code) || !policyCorrective(err.Error()) || (path != "" && !strings.Contains(err.Error(), path)) {
		t.Fatalf("want corrective %s at %q; got %v", code, path, err)
	}
}
func contextAssert(t *testing.T, f contextFixtureData, phase string) w.WorkerContext {
	t.Helper()
	got, err := w.BuildContext(f.Snapshot, "current")
	if err != nil {
		t.Fatalf("context: %v", err)
	}
	selected := contextChange(t, f)
	if !reflect.DeepEqual(got.Change, selected) || got.ChangeText != f.Raw || got.SpecText != f.Spec || got.System != f.System {
		t.Fatalf("typed/exact source data mismatch: %+v", got)
	}
	before, err := p.Project(f.Snapshot, "current")
	if err != nil {
		t.Fatal(err)
	}
	var current c.Module
	for _, m := range before.Modules {
		if m.Name == "alpha" {
			current = m
		}
	}
	projected, err := p.ApplyChange(current, selected)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got.Module, projected) || len(got.Module.Specs) != 2 || got.Module.Specs[1].ID != "ALP-002" || !reflect.DeepEqual(got.Module.Interface, []string{"pub/**"}) {
		t.Fatalf("wrong current projection %+v", got.Module)
	}
	deps := []c.Module{}
	for _, name := range []string{"gamma", "beta"} {
		for _, m := range before.Modules {
			if m.Name == name {
				deps = append(deps, m)
			}
		}
	}
	if !reflect.DeepEqual(got.Dependencies, deps) || len(got.Dependencies[1].Specs) != 2 || got.Dependencies[1].Specs[1].ID != "BET-002" {
		t.Fatalf("direct projected dependencies %+v", got.Dependencies)
	}
	if !reflect.DeepEqual(got.Files, contextExpectedFiles(f, phase)) {
		t.Fatalf("selected files %+v; want %+v", got.Files, contextExpectedFiles(f, phase))
	}
	name := phase
	if name == "" {
		name = "author-change"
	}
	if selected.Regenerate {
		name = "regenerate"
	}
	play, err := p.ResolvePlaybook(f.Root, name)
	if err != nil {
		t.Fatal(err)
	}
	if got.Playbook != play.Text || got.PlaybookHash != play.Hash {
		t.Fatalf("instruction bytes/hash differ")
	}
	repo, err := g.Open(f.Root)
	if err != nil {
		t.Fatal(err)
	}
	if got.GitCommonDir != repo.CommonDir {
		t.Fatalf("common dir %q want %q", got.GitCommonDir, repo.CommonDir)
	}
	data, err := json.Marshal(got)
	if err != nil {
		t.Fatal(err)
	}
	var values map[string]json.RawMessage
	if err = json.Unmarshal(data, &values); err != nil {
		t.Fatal(err)
	}
	want := []string{"change", "change_text", "module", "spec_text", "system", "dependencies", "files", "playbook", "playbook_hash", "git_common_dir"}
	if len(values) != len(want) {
		t.Fatalf("context keys %+v", values)
	}
	for _, key := range want {
		if _, ok := values[key]; !ok {
			t.Fatalf("missing JSON %s", key)
		}
	}
	return got
}
func Test_FLOW016_FLOW017_FLOW018ProjectedPhaseContext(t *testing.T) {
	for _, format := range []string{"sha1", "sha256"} {
		for _, phase := range []string{"", "evals", "audit", "impl"} {
			t.Run(format+"/"+phase, func(t *testing.T) { f := contextFixture(t, phase, format); contextAssert(t, f, phase) })
		}
	}
}
func Test_FLOW016_FLOW017CurrentBytesAndSnapshotAuthority(t *testing.T) {
	f := contextFixture(t, "impl", "sha1")
	f.Raw = strings.ReplaceAll(f.Raw, "\n", "\r\n") + "\r\n## Notes\r\n\r\nworker feedback λ"
	worktreePut(t, f.Root, f.Path, f.Raw)
	f.Spec = strings.ReplaceAll(f.Spec, "\n", "\r\n")
	worktreePut(t, f.Root, "units/a/SPEC.md", f.Spec)
	worktreePut(t, f.Root, "fireweed.yaml", "malformed current config: [")
	contextAssert(t, f, "impl")
}
func Test_FLOW016InvalidContextSelection(t *testing.T) {
	for _, kind := range []string{"missing", "ambiguous", "missing-module", "phase", "archive", "nested", "slug-path", "bad-projection", "missing-dependency"} {
		t.Run(kind, func(t *testing.T) {
			f := contextFixture(t, "evals", "sha1")
			slug := "current"
			idx := 1
			switch kind {
			case "missing":
				slug = "absent"
			case "ambiguous":
				f.Snapshot.Changes = append(f.Snapshot.Changes, f.Snapshot.Changes[idx])
			case "missing-module":
				f.Snapshot.Changes[idx].Module = "absent"
			case "phase":
				f.Snapshot.Changes[idx].Phase = "applied"
			case "archive":
				f.Snapshot.Changes[idx].Path = "changes/archive/0002-current.md"
			case "nested":
				f.Snapshot.Changes[idx].Path = "changes/sub/0002-current.md"
			case "slug-path":
				f.Snapshot.Changes[idx].Path = "changes/0002-other.md"
			case "bad-projection":
				f.Snapshot.Changes[0].Added = append(f.Snapshot.Changes[0].Added, c.Spec{ID: "BET-001", Body: "Duplicate MUST fail."})
			case "missing-dependency":
				f.Snapshot.Changes[idx].Set["depends_on"] = []string{"absent"}
			}
			f.Snapshot.Root = filepath.Join(t.TempDir(), "no repository")
			contextError(t, f, slug, "C05", "")
		})
	}
}
func Test_FLOW017ContextInstructionVariants(t *testing.T) {
	for _, kind := range []string{"embedded", "current", "stale", "malformed", "empty", "crlf", "regenerate"} {
		t.Run(kind, func(t *testing.T) {
			f := contextFixture(t, "impl", "sha1")
			name := "impl"
			if kind == "regenerate" {
				f.Snapshot.Changes[1].Regenerate = true
				name = "regenerate"
			}
			if kind != "embedded" {
				base, err := a.Playbook(name)
				if err != nil {
					t.Fatal(err)
				}
				text := "---\nbased_on: " + base.Hash + "\n---\nexact selected λ"
				switch kind {
				case "stale":
					text = strings.Replace(text, base.Hash, strings.Repeat("0", 64), 1)
				case "malformed":
					text = "---\ninvalid: [\n---\ninstructions"
				case "empty":
					text = ""
				case "crlf":
					text = strings.ReplaceAll(text, "\n", "\r\n")
				}
				worktreePut(t, f.Root, ".fireweed/playbooks/"+name+".md", text)
			}
			contextAssert(t, f, "impl")
		})
	}
}
func Test_FLOW017NestedAndLinkedContext(t *testing.T) {
	for _, linked := range []bool{false, true} {
		t.Run(fmt.Sprint(linked), func(t *testing.T) {
			f := contextFixture(t, "impl", "sha1")
			if linked {
				peer := claimPeer(t, f.Root)
				f.Root = peer
				f.Snapshot.Root = peer
				worktreePut(t, peer, "units/a/e/new stub.txt", f.Expected["units/a/e/new stub.txt"])
				worktreePut(t, peer, ".fireweed/playbooks/impl.md", "linked instructions")
				worktreePut(t, filepath.Dir(peer), "SYSTEM.md", "outside decoy")
			}
			nested := filepath.Join(f.Root, "nested ' λ")
			if err := os.MkdirAll(nested, 0755); err != nil {
				t.Fatal(err)
			}
			worktreePut(t, nested, "SYSTEM.md", "nested decoy")
			f.Snapshot.Root = nested
			contextAssert(t, f, "impl")
		})
	}
}
func Test_FLOW017MissingSourcesAndClaimRecovery(t *testing.T) {
	for _, kind := range []string{"system-absent", "spec-missing", "change-missing", "claim-match", "claim-foreign", "claim-run", "claim-phase", "claim-target", "claim-malformed", "queued-claim"} {
		t.Run(kind, func(t *testing.T) {
			phase := "impl"
			if kind == "queued-claim" {
				phase = ""
			}
			f := contextFixture(t, phase, "sha1")
			if kind == "system-absent" {
				if err := os.Remove(filepath.Join(f.Root, "SYSTEM.md")); err != nil {
					t.Fatal(err)
				}
				f.System = ""
				contextAssert(t, f, phase)
				return
			}
			target := f.Path
			if kind == "spec-missing" {
				target = "units/a/SPEC.md"
			}
			if err := os.Remove(filepath.Join(f.Root, target)); err != nil {
				t.Fatal(err)
			}
			if strings.HasPrefix(kind, "claim-") || kind == "queued-claim" {
				state := f.Raw
				claim := w.Claim{Worktree: f.Root, Change: "current", Run: "context-run", Base: claimBase(t, f.Root), State: state}
				switch kind {
				case "claim-foreign":
					claim.Worktree = stateCanonicalDirectory(t, t.TempDir())
				case "claim-run":
					claim.Run = "different"
				case "claim-phase":
					claim.State = strings.Replace(state, "phase: impl", "phase: evals", 1)
				case "claim-target":
					claim.State = strings.Replace(state, "module: alpha", "module: beta", 1)
				case "claim-malformed":
					claim.State = "malformed"
				}
				claimRaw(t, f.Root, "current", claimPayload(t, claim))
			}
			if kind == "claim-match" {
				contextAssert(t, f, phase)
			} else {
				contextError(t, f, "current", "C01", target)
			}
		})
	}
}
func Test_FLOW018_FLOW019SelectedReadErrors(t *testing.T) {
	for _, target := range []string{"SYSTEM.md", "units/a/SPEC.md", "changes/0002-current.md", "units/a/pub/contract ' λ.txt", "units/c/api/gamma.txt", ".fireweed/playbooks/impl.md"} {
		for _, kind := range []string{"directory", "outside", "dangling", "missing-outside", "unreadable"} {
			t.Run(target+"/"+kind, func(t *testing.T) {
				f := contextFixture(t, "impl", "sha1")
				full := filepath.Join(f.Root, target)
				if target == ".fireweed/playbooks/impl.md" {
					worktreePut(t, f.Root, target, "selected")
				}
				if err := os.Remove(full); err != nil {
					t.Fatal(err)
				}
				outside := stateCanonicalDirectory(t, t.TempDir())
				private := filepath.Join(outside, "private")
				if err := os.WriteFile(private, []byte("outside sentinel"), 0600); err != nil {
					t.Fatal(err)
				}
				switch kind {
				case "directory":
					if err := os.Mkdir(full, 0755); err != nil {
						t.Fatal(err)
					}
				case "outside":
					if err := os.Symlink(private, full); err != nil {
						t.Fatal(err)
					}
				case "dangling":
					if err := os.Symlink(filepath.Join(f.Root, "missing"), full); err != nil {
						t.Fatal(err)
					}
				case "missing-outside":
					if err := os.Symlink(filepath.Join(outside, "missing", "file"), full); err != nil {
						t.Fatal(err)
					}
				case "unreadable":
					if err := os.WriteFile(full, []byte("denied"), 0000); err != nil {
						t.Fatal(err)
					}
					t.Cleanup(func() { _ = os.Chmod(full, 0600) })
					if _, err := os.ReadFile(full); err == nil {
						t.Skip("cannot establish unreadable file premise under current identity")
					}
				}
				code := "C01"
				if target == ".fireweed/playbooks/impl.md" {
					code = "D05"
				}
				contextError(t, f, "current", code, target)
				if b, err := os.ReadFile(private); err != nil || string(b) != "outside sentinel" {
					t.Fatal("outside fixture changed")
				}
			})
		}
	}
}
func Test_FLOW018MissingTrackedAndContainedLinks(t *testing.T) {
	f := contextFixture(t, "impl", "sha1")
	target := "units/a/pub/contract ' λ.txt"
	full := filepath.Join(f.Root, target)
	if err := os.Remove(full); err != nil {
		t.Fatal(err)
	}
	contextError(t, f, "current", "C01", target)
	worktreePut(t, f.Root, "contained-content", f.Expected[target])
	if err := os.Symlink(filepath.Join(f.Root, "contained-content"), full); err != nil {
		t.Fatal(err)
	}
	contextAssert(t, f, "impl")
}
func Test_FLOW019ContextOwnershipAndEmptyCollections(t *testing.T) {
	f := contextFixture(t, "impl", "sha1")
	before := f.Snapshot
	original, err := json.Marshal(before)
	if err != nil {
		t.Fatal(err)
	}
	got := contextAssert(t, f, "impl")
	got.Change.Document.Metadata["module"] = "mutated"
	got.Change.Set["depends_on"].([]any)[0] = "mutated"
	got.Change.Added[0].Body = "mutated"
	got.Change.Set["hooks"].(map[string]any)["eval"] = "mutated"
	got.Change.Bases = map[string]string{"new": "mutated"}
	got.Module.Interface[0] = "mutated"
	got.Module.Document.Metadata["name"] = "mutated"
	got.Module.Specs[0].Body = "mutated"
	got.Dependencies[0].Name = "mutated"
	got.Dependencies[1].Document.Metadata["name"] = "mutated"
	got.Dependencies[1].Interface[0] = "mutated"
	got.Files[0].Content = "mutated"
	after, _ := json.Marshal(f.Snapshot)
	if string(after) != string(original) {
		t.Fatal("result mutation changed input immediately")
	}
	contextAssert(t, f, "impl")
	for i := range f.Snapshot.Modules {
		if f.Snapshot.Modules[i].Name == "alpha" {
			f.Snapshot.Modules[i].Zones = c.Zones{}
			f.Snapshot.Modules[i].DependsOn = []string{}
		}
	}
	f.Snapshot.Changes[1].Set = map[string]any{"depends_on": []string{}, "evals": []string{}, "interface": []string{}, "preserve": []string{}, "append_only": []string{}, "disposable": []string{}}
	empty, err := w.BuildContext(f.Snapshot, "current")
	if err != nil {
		t.Fatal(err)
	}
	if empty.Files == nil || empty.Dependencies == nil || len(empty.Files) != 0 || len(empty.Dependencies) != 0 {
		t.Fatalf("nonnull empties %+v", empty)
	}
	empty.Files = append(empty.Files, w.FileContent{Path: "mutated"})
	empty.Dependencies = append(empty.Dependencies, c.Module{Name: "mutated"})
	again, err := w.BuildContext(f.Snapshot, "current")
	if err != nil || len(again.Files) != 0 || len(again.Dependencies) != 0 {
		t.Fatalf("empty collection alias %+v %v", again, err)
	}
}

func Test_FLOW016_FLOW017_FLOW018_FLOW019ContextFixturePremises(t *testing.T) {
	f := contextFixture(t, "impl", "sha1")
	selected := contextChange(t, f)
	if _, ok := selected.Set["depends_on"].([]any); !ok {
		t.Fatalf("parsed nested mutation premise %T", selected.Set["depends_on"])
	}
	before, err := p.Project(f.Snapshot, "current")
	if err != nil {
		t.Fatal(err)
	}
	var alpha, beta c.Module
	for _, m := range before.Modules {
		switch m.Name {
		case "alpha":
			alpha = m
		case "beta":
			beta = m
		}
	}
	projected, err := p.ApplyChange(alpha, selected)
	if err != nil {
		t.Fatal(err)
	}
	if len(projected.Specs) != 2 || projected.Specs[1].ID != "ALP-002" || len(beta.Specs) != 2 || beta.Specs[1].ID != "BET-002" || !reflect.DeepEqual(projected.DependsOn, []string{"gamma", "beta"}) || !reflect.DeepEqual(beta.Interface, []string{"future/**"}) {
		t.Fatal("projected fixture premise")
	}
	for _, path := range []string{f.Path, "units/a/SPEC.md", "SYSTEM.md"} {
		if raw, err := os.ReadFile(filepath.Join(f.Root, path)); err != nil || len(raw) == 0 {
			t.Fatalf("exact source premise %q %v", path, err)
		}
	}
	for _, phase := range []string{"", "evals", "audit", "impl", "regenerate"} {
		name := phase
		if name == "" {
			name = "author-change"
		}
		asset, err := p.ResolvePlaybook(f.Root, name)
		if err != nil || asset.Text == "" || len(asset.Hash) != 64 {
			t.Fatalf("instruction premise %q %v", name, err)
		}
	}
	input, _ := json.Marshal(f.Snapshot)
	var snapshot p.Snapshot
	if err = json.Unmarshal(input, &snapshot); err != nil {
		t.Fatal(err)
	}
	if _, err = p.Project(snapshot, "current"); err != nil {
		t.Fatalf("native serialized projection premise %v", err)
	}
}
func Test_FLOW017ExactEmptySources(t *testing.T) {
	for _, path := range []string{"SYSTEM.md", "units/a/SPEC.md", "changes/0002-current.md"} {
		t.Run(path, func(t *testing.T) {
			f := contextFixture(t, "impl", "sha1")
			worktreePut(t, f.Root, path, "")
			switch path {
			case "SYSTEM.md":
				f.System = ""
			case "units/a/SPEC.md":
				f.Spec = ""
			default:
				f.Raw = ""
			}
			contextAssert(t, f, "impl")
		})
	}
}
func Test_FLOW019ContextFilesystemAndEnumerationErrors(t *testing.T) {
	for _, mode := range []string{"repository", "enumeration"} {
		t.Run(mode, func(t *testing.T) {
			f := contextFixture(t, "impl", "sha1")
			if mode == "repository" {
				f.Snapshot.Root = t.TempDir()
			} else {
				if os.PathSeparator == '\\' {
					t.Skip("POSIX Git failure injection")
				}
				real, err := exec.LookPath("git")
				if err != nil {
					t.Fatal(err)
				}
				traps := t.TempDir()
				script := "#!/bin/sh\nfor a in \"$@\"; do if [ \"$a\" = ls-files ]; then exit 73; fi; done\nexec " + claimQuote(real) + " \"$@\"\n"
				if err = os.WriteFile(filepath.Join(traps, "git"), []byte(script), 0755); err != nil {
					t.Fatal(err)
				}
				t.Setenv("PATH", traps+string(os.PathListSeparator)+os.Getenv("PATH"))
			}
			contextError(t, f, "current", "C01", "")
		})
	}
}
func Test_FLOW019ContextPreservesCompleteState(t *testing.T) {
	if os.PathSeparator == '\\' {
		t.Skip("POSIX command traps")
	}
	for _, linked := range []bool{false, true} {
		for _, mode := range []string{"success", "audit", "override", "read-error", "selection-error", "claim-recovery"} {
			t.Run(fmt.Sprint(linked)+"/"+mode, func(t *testing.T) {
				phase := "impl"
				if mode == "audit" {
					phase = "audit"
				}
				f := contextFixture(t, phase, "sha1")
				if linked {
					f.Root = claimPeer(t, f.Root)
					f.Snapshot.Root = f.Root
					worktreePut(t, f.Root, "units/a/e/new stub.txt", f.Expected["units/a/e/new stub.txt"])
				}
				claim := w.Claim{Worktree: f.Root, Change: "current", Run: "context-run", Base: claimBase(t, f.Root), State: f.Raw}
				claimRaw(t, f.Root, "current", claimPayload(t, claim))
				claimGit(t, f.Root, "update-ref", "refs/fireweed/drills/retained", claim.Base)
				if mode == "override" {
					worktreePut(t, f.Root, ".fireweed/playbooks/impl.md", "---\nbased_on: stale\n---\nselected")
				}
				if mode == "read-error" {
					if err := os.Remove(filepath.Join(f.Root, "units/a/e/current.txt")); err != nil {
						t.Fatal(err)
					}
				}
				if mode == "claim-recovery" {
					if err := os.Remove(filepath.Join(f.Root, f.Path)); err != nil {
						t.Fatal(err)
					}
				}
				traps, marker := worktreeTraps(t, false)
				for _, name := range []string{"context-hook-trap", "context-agent-trap"} {
					if err := os.WriteFile(filepath.Join(traps, name), []byte("#!/bin/sh\nprintf attempted >> "+claimQuote(marker)+"\nexit 97\n"), 0755); err != nil {
						t.Fatal(err)
					}
				}
				temporary := stateCanonicalDirectory(t, t.TempDir())
				t.Setenv("TMPDIR", temporary)
				t.Setenv("TMP", temporary)
				t.Setenv("TEMP", temporary)
				common := claimCommon(t, f.Root)
				files := claimTree(t, f.Root, func(path string) bool { return path == ".git" })
				metadata := claimTree(t, common, func(string) bool { return false })
				temps := claimTree(t, temporary, func(string) bool { return false })
				env := os.Environ()
				cwd, err := os.Getwd()
				if err != nil {
					t.Fatal(err)
				}
				input := policyJSON(t, f.Snapshot)
				if mode == "selection-error" {
					contextError(t, f, "absent", "C05", "")
				} else if mode == "read-error" {
					contextError(t, f, "current", "C01", "units/a/e/current.txt")
				} else {
					contextAssert(t, f, phase)
				}
				if !reflect.DeepEqual(files, claimTree(t, f.Root, func(path string) bool { return path == ".git" })) || !reflect.DeepEqual(metadata, claimTree(t, common, func(string) bool { return false })) || !reflect.DeepEqual(temps, claimTree(t, temporary, func(string) bool { return false })) {
					t.Fatal("context changed files/common Git/temp storage")
				}
				afterCWD, err := os.Getwd()
				if err != nil || afterCWD != cwd || !reflect.DeepEqual(env, os.Environ()) || input != policyJSON(t, f.Snapshot) {
					t.Fatal("context changed parent/input state")
				}
				if _, err = os.Stat(marker); !os.IsNotExist(err) {
					t.Fatal("forbidden command executed")
				}
			})
		}
	}
}

func Test_FLOW018DeduplicatedNULInventory(t *testing.T) {
	if os.PathSeparator == '\\' {
		t.Skip("POSIX duplicate Git enumeration adapter")
	}
	f := contextFixture(t, "impl", "sha1")
	path := "units/a/e/newline\nfile.txt"
	worktreePut(t, f.Root, path, "newline filename bytes")
	f.Expected[path] = "newline filename bytes"
	real, err := exec.LookPath("git")
	if err != nil {
		t.Fatal(err)
	}
	traps := t.TempDir()
	script := "#!/bin/sh\nfor a in \"$@\"; do if [ \"$a\" = ls-files ]; then " + claimQuote(real) + " \"$@\"\n" + claimQuote(real) + " \"$@\"\nexit $?\nfi;done\nexec " + claimQuote(real) + " \"$@\"\n"
	if err = os.WriteFile(filepath.Join(traps, "git"), []byte(script), 0755); err != nil {
		t.Fatal(err)
	}
	t.Setenv("PATH", traps+string(os.PathListSeparator)+os.Getenv("PATH"))
	contextAssert(t, f, "impl")
}
