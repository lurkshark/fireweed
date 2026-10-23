package evals

import (
	"bytes"
	"compress/zlib"
	"context"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"testing"
	"time"

	c "github.com/lurkshark/fireweed/packages/contracts"
	p "github.com/lurkshark/fireweed/packages/project"
	w "github.com/lurkshark/fireweed/packages/workflow"
)

type startFixtureData struct {
	Root, Path, Raw string
	Mode            fs.FileMode
}

func startCommit(t *testing.T, root, message string) {
	t.Helper()
	claimGit(t, root, "add", ".")
	claimGit(t, root, "-c", "user.name=Start Fixture", "-c", "user.email=start@example.invalid", "commit", "-m", message)
}
func startFixture(t *testing.T, format, kind, phase string) startFixtureData {
	t.Helper()
	root := claimRepo(t, format)
	config := `version: 1
fireweed: '>=0.4, <0.5'
base: approval-base
modules: [units/*]
module_defaults:
  interface: [api.txt]
  evals: [evals/**]
  disposable: [impl/**]
  preserve: [keep/**]
  append_only: [ledger/**]
hooks:
  eval: start-hook-trap
  boundaries: start-hook-trap
agent: {command: [start-agent-trap]}
protected: [private/**, units/a/impl/protected.txt]
worker_writable: [shared/**]
`
	module := func(name, prefix string) string {
		return "---\nname: " + name + "\nid_prefix: " + prefix + "\n---\n\n## Specs\n\n### " + prefix + "-001 [review]\nThe existing behavior MUST remain.\n\n## Decisions\n\nRetain the initial choice.\n"
	}
	meta := ""
	if phase != "" {
		meta = "run: retained-start-run\nphase: " + phase + "\n"
	}
	switch kind {
	case "fix":
		meta += "fixes: [ALP-001]\n"
	case "regenerate":
		meta += "regenerate: true\n"
	}
	raw := "---\nmodule: alpha\n" + meta + "---\n\nRetained preamble.\n\n## Motivation\n\nStart the exact authored change.\n\n## Empty author section\n\n## Interface\n\nRetain the public shape.\n"
	if kind == "spec" {
		id := "new:feature"
		if phase != "" {
			id = "ALP-051"
		}
		raw += "\n## Added\n\n### " + id + "\nThe new behavior MUST retain " + id + ".\n"
	}
	raw += "\n## Notes\n\n```markdown\n## Modified\n## Removed\n```\n"
	files := map[string]string{
		"fireweed.yaml": config, "SYSTEM.md": "The start fixture retains its instructions.\n", ".gitignore": "ignored*\n",
		"units/a/SPEC.md": module("alpha", "ALP"), "units/b/SPEC.md": module("beta", "BET"),
		"changes/0001-current.md":          raw,
		"changes/archive/0000-reserved.md": "---\nmodule: alpha\nrun: archived-run\nphase: applied\n---\n\n## Motivation\n\nReserve a removed number.\n\n## Removed\n\n### ALP-050\n",
		"units/a/api.txt":                  "Public retained bytes.\r\n", "units/a/evals/case.txt": "Retained eval bytes.\n",
		"units/a/impl/code.txt": "Old disposable implementation.\n", "units/a/impl/protected.txt": "Protected despite disposable glob.\n",
		"units/a/keep/value.txt": "Preserved value.\n", "units/a/ledger/001.txt": "Append-only value.\n", "units/a/manifest.txt": "Other owned data.\n",
		"units/b/impl/peer.txt": "Other module disposable.\n", "private/secret.txt": "Private bytes.\n", "shared/keep.txt": "Shared bytes.\n",
	}
	for path, text := range files {
		worktreePut(t, root, path, text)
	}
	if err := os.Chmod(filepath.Join(root, "changes/0001-current.md"), 0600); err != nil {
		t.Fatal(err)
	}
	if err := os.Chmod(filepath.Join(root, "units/a/impl/code.txt"), 0751); err != nil {
		t.Fatal(err)
	}
	startCommit(t, root, "Authored fixture without phase-subject conventions")
	claimGit(t, root, "branch", "approval-base")
	worktreePut(t, root, "ignored-cache/value.txt", "Ignored data remains.\n")
	startFixtureValid(t, root)
	return startFixtureData{root, "changes/0001-current.md", raw, 0600}
}
func startFixtureValid(t *testing.T, root string) p.Snapshot {
	t.Helper()
	s, err := p.Load(root)
	if err != nil {
		t.Fatal("public fixture load", err)
	}
	for _, ds := range [][]p.Diagnostic{p.Check(s), p.CheckQueue(s)} {
		for _, d := range ds {
			if d.Severity == "error" {
				t.Fatalf("public fixture check: %+v", ds)
			}
		}
	}
	return s
}

type startNode struct {
	Mode        fs.FileMode
	Bytes, Link string
}

func startTree(t *testing.T, root string) map[string]startNode {
	t.Helper()
	out := map[string]startNode{}
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		if rel == ".git" {
			if d.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}
		if rel == "." {
			return nil
		}
		info, err := d.Info()
		if err != nil {
			return err
		}
		n := startNode{Mode: info.Mode()}
		if info.Mode()&os.ModeSymlink != 0 {
			n.Link, err = os.Readlink(path)
		} else if info.Mode().IsRegular() {
			var b []byte
			b, err = os.ReadFile(path)
			n.Bytes = string(b)
		}
		if err != nil {
			return err
		}
		out[filepath.ToSlash(rel)] = n
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	return out
}

type startGitState struct {
	Head, Refs string
	Index      []byte
	Cwd        string
	Env        []string
}

func startGitStdout(t *testing.T, root string, args ...string) string {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, "git", args...)
	cmd.Dir = root
	cmd.Env = append(os.Environ(), "GIT_CONFIG_NOSYSTEM=1", "GIT_CONFIG_GLOBAL="+os.DevNull, "GIT_OPTIONAL_LOCKS=0")
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	out, err := cmd.Output()
	if err != nil {
		t.Fatalf("fixture Git stdout %q: %v %s", args, err, stderr.String())
	}
	return strings.TrimSpace(string(out))
}
func startGitSnapshot(t *testing.T, root string) startGitState {
	t.Helper()
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	path := startGitStdout(t, root, "rev-parse", "--git-path", "index")
	if !filepath.IsAbs(path) {
		path = filepath.Join(root, path)
	}
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return startGitState{startGitStdout(t, root, "rev-parse", "HEAD"), startGitStdout(t, root, "show-ref"), b, cwd, os.Environ()}
}
func startUnrelated(t *testing.T, root string, before startGitState, owned string) {
	t.Helper()
	after := startGitSnapshot(t, root)
	lines := []string{}
	for _, line := range strings.Split(after.Refs, "\n") {
		if !strings.HasSuffix(line, " refs/fireweed/claims/"+owned) {
			lines = append(lines, line)
		}
	}
	after.Refs = strings.Join(lines, "\n")
	if !reflect.DeepEqual(before, after) {
		t.Errorf("changed unrelated HEAD/refs/index/cwd/environment\nbefore %+v\nafter %+v", before, after)
	}
}
func startFailure(t *testing.T, root, slug, code string) error {
	t.Helper()
	checkStorage := startStorage(t, root, slug)
	defer checkStorage()
	tree := startTree(t, root)
	git := startGitSnapshot(t, root)
	got, err := w.Start(root, slug)
	if !reflect.DeepEqual(got, w.StartResult{}) {
		t.Errorf("partial result on rejection: %+v", got)
	}
	if err == nil || !strings.Contains(err.Error(), code) || !policyCorrective(err.Error()) {
		t.Errorf("want useful corrective %s, got %v", code, err)
	}
	if !reflect.DeepEqual(tree, startTree(t, root)) {
		t.Error("rejection changed working files or modes")
	}
	if !reflect.DeepEqual(git, startGitSnapshot(t, root)) {
		t.Error("rejection changed Git or parent state")
	}
	return err
}
func startFixtureSlug(t *testing.T, f startFixtureData) string {
	t.Helper()
	ch, err := c.ParseChange(f.Raw, f.Path, "ALP")
	if err != nil {
		t.Fatal(err)
	}
	return ch.Slug
}
func startSuccess(t *testing.T, f startFixtureData, start, phase string) w.StartResult {
	t.Helper()
	checkStorage := startStorage(t, f.Root, startFixtureSlug(t, f))
	defer checkStorage()
	before := startGitSnapshot(t, f.Root)
	got, err := w.Start(start, startFixtureSlug(t, f))
	if err != nil {
		t.Fatalf("valid start: %v", err)
	}
	startAssertSuccess(t, f, phase, before, got)
	return got
}
func startAssertSuccess(t *testing.T, f startFixtureData, phase string, before startGitState, got w.StartResult) {
	t.Helper()
	slug := startFixtureSlug(t, f)
	if got.Diagnostics == nil {
		t.Error("nil diagnostics")
	}
	cl, err := w.LoadClaim(f.Root, slug)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(cl, got.Claim) || cl.Object == "" || cl.Worktree != f.Root || cl.Change != slug || cl.Base != before.Head || cl.Run == "" {
		t.Errorf("stored/returned claim disagreement: %+v %+v", cl, got.Claim)
	}
	raw, err := os.ReadFile(filepath.Join(f.Root, f.Path))
	if err != nil {
		t.Fatal(err)
	}
	if cl.State != string(raw) {
		t.Error("claim state is not exact saved bytes")
	}
	ch, err := c.ParseChange(string(raw), f.Path, "ALP")
	if err != nil {
		t.Fatal(err)
	}
	if ch.Phase != phase || ch.Run != cl.Run || !reflect.DeepEqual(ch, got.Context.Change) || got.Context.ChangeText != string(raw) {
		t.Errorf("started context mismatch: %+v %+v", ch, got.Context.Change)
	}
	s, owned, err := w.LoadApproval(f.Root, slug)
	if err != nil {
		t.Fatal(err)
	}
	expected, err := w.BuildContext(s, slug)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got.Context, expected) || owned != cl {
		t.Error("context is not complete owned BuildContext")
	}
	original, err := p.LoadAt(f.Root, before.Head)
	if err != nil {
		t.Fatal("original committed source premise", err)
	}
	allocated, _, err := p.Allocate(original, slug)
	if err != nil {
		t.Fatal("public original allocation premise", err)
	}
	allocated.Document.Metadata["run"] = cl.Run
	allocated.Document.Metadata["phase"] = phase
	canonical, err := c.RenderMarkdown(allocated.Document)
	if err != nil {
		t.Fatal(err)
	}
	wanted, err := c.ParseChange(canonical, f.Path, "ALP")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(ch, wanted) {
		t.Error("start changed original allocated author intent, structure, identity or metadata", ch, wanted)
	}

	info, err := os.Stat(filepath.Join(f.Root, f.Path))
	if err != nil || info.Mode().Perm() != f.Mode {
		t.Error("lost original change mode", err)
	}
	startUnrelated(t, f.Root, before, slug)
}
func Test_FLOW032_FLOW033_FLOW034_FLOW035_FLOW036_FLOW037_FLOW038_FLOW039_FLOW040StartFixtureControls(t *testing.T) {
	for _, kind := range []string{"spec", "fix", "maintenance", "regenerate"} {
		for _, phase := range []string{"", "evals", "audit", "impl"} {
			f := startFixture(t, "sha1", kind, phase)
			s := startFixtureValid(t, f.Root)
			if len(s.Changes) != 1 || s.Changes[0].Slug != "current" || s.Changes[0].Kind != kind {
				t.Fatalf("fixture classification %+v", s.Changes)
			}
			text, err := c.RenderMarkdown(s.Changes[0].Document)
			if err != nil {
				t.Fatal(err)
			}
			ch, err := c.ParseChange(text, f.Path, "ALP")
			if err != nil || !reflect.DeepEqual(ch, s.Changes[0]) {
				t.Fatal("fixture round trip", err)
			}
			if clean := claimGit(t, f.Root, "status", "--porcelain"); clean != "" {
				t.Fatal("fixture not clean", clean)
			}
		}
	}
}
func Test_FLOW032_FLOW034_FLOW037_FLOW040StartFreshStateAndContext(t *testing.T) {
	runs := map[string]bool{}
	for _, format := range []string{"sha1", "sha256"} {
		for _, kind := range []string{"spec", "fix", "maintenance"} {
			t.Run(format+"-"+kind, func(t *testing.T) {
				f := startFixture(t, format, kind, "")
				before := startTree(t, f.Root)
				nested := filepath.Join(f.Root, "units/a/evals")
				phase := "evals"
				if kind == "maintenance" {
					phase = "impl"
				}
				got := startSuccess(t, f, nested, phase)
				if runs[got.Claim.Run] {
					t.Error("reused fresh run")
				}
				runs[got.Claim.Run] = true
				after := startTree(t, f.Root)
				delete(before, f.Path)
				delete(after, f.Path)
				if !reflect.DeepEqual(before, after) {
					t.Error("start changed unrelated working files")
				}
				if kind == "spec" {
					if len(got.Context.Change.Added) != 1 || got.Context.Change.Added[0].ID != "ALP-051" || !strings.Contains(got.Context.Change.Added[0].Body, "ALP-051") || strings.Contains(got.Context.ChangeText, "new:feature") {
						t.Error("wrong reservation/binding", got.Context.Change.Added)
					}
					names := []string{}
					for _, s := range got.Context.Change.Document.Sections {
						names = append(names, s.Name)
					}
					if !reflect.DeepEqual(names, []string{"Motivation", "Empty author section", "Interface", "Added", "Notes"}) {
						t.Error("changed authored heading presence", names)
					}
				}
				for _, d := range got.Diagnostics {
					if d.Severity == "error" || d.Code == "P04" || d.Code == "D01" {
						t.Errorf("unexpected success diagnostic %+v", d)
					}
				}
				b, err := json.Marshal(got)
				if err != nil {
					t.Fatal(err)
				}
				var wire map[string]json.RawMessage
				if json.Unmarshal(b, &wire) != nil || len(wire) != 3 || wire["context"] == nil || wire["claim"] == nil || string(wire["diagnostics"]) == "null" {
					t.Error("result wire shape", string(b))
				}
			})
		}
	}
}
func Test_FLOW035StartCommittedResume(t *testing.T) {
	for _, phase := range []string{"evals", "audit", "impl"} {
		for _, unblocked := range []bool{false, true} {
			t.Run(fmt.Sprintf("%s-unblocked-%v", phase, unblocked), func(t *testing.T) {
				f := startFixture(t, "sha1", "spec", phase)
				if unblocked {
					blocked := strings.Replace(f.Raw, "module: alpha\n", "module: alpha\nblocked: awaiting repair\n", 1)
					worktreePut(t, f.Root, f.Path, blocked)
					startCommit(t, f.Root, "Record blocked state")
					worktreePut(t, f.Root, f.Path, f.Raw)
					startCommit(t, f.Root, "Remove blocked field")
				}
				got := startSuccess(t, f, f.Root, phase)
				if got.Claim.Run != "retained-start-run" || len(got.Context.Change.Added) != 1 || got.Context.Change.Added[0].ID != "ALP-051" {
					t.Error("resume reallocated state", got.Context.Change)
				}
			})
		}
	}
}

// Exact Start codes are asserted only where FLOW-032–040 prescribe them.
// Other failures still require corrective detail, zero results and preservation.
func Test_FLOW032_FLOW033StartPreconditionRejections(t *testing.T) {
	binary := startProxyBuild(t)
	for _, kind := range []string{"tracked", "staged", "untracked", "blocked", "behind-head", "existing-own", "existing-peer", "existing-other", "orphan-own", "malformed-claim", "invalid-config", "invalid-queue"} {
		t.Run(kind, func(t *testing.T) {
			f := startFixture(t, "sha1", "spec", "")
			code := ""
			switch kind {
			case "tracked":
				code = "P01"
				worktreePut(t, f.Root, "retained", "dirty")
			case "staged":
				code = "P01"
				worktreePut(t, f.Root, "retained", "staged")
				claimGit(t, f.Root, "add", "retained")
			case "untracked":
				code = "P01"
				worktreePut(t, f.Root, "new-untracked", "dirty")
			case "blocked":
				worktreePut(t, f.Root, f.Path, strings.Replace(f.Raw, "module: alpha\n", "module: alpha\nblocked: pending repair\n", 1))
				startCommit(t, f.Root, "Record block")
				code = "P03"
			case "behind-head":
				worktreePut(t, f.Root, "changes/0000-before.md", strings.Replace(f.Raw, "new:feature", "new:prior", -1))
				startCommit(t, f.Root, "Earlier queue head")
				code = "P03"
			case "existing-own":
				claimSeed(t, f.Root, "current", f.Raw)
				code = ""
			case "existing-other":
				raw := strings.ReplaceAll(f.Raw, "new:feature", "new:other")
				worktreePut(t, f.Root, "changes/0002-other.md", raw)
				startCommit(t, f.Root, "Another queued change")
				claimSeed(t, f.Root, "other", raw)
				code = ""
			case "orphan-own":
				claimSeed(t, f.Root, "orphan", f.Raw)
				code = ""
			case "existing-peer":
				peer := claimPeer(t, f.Root)
				claimSeed(t, peer, "current", f.Raw)
				code = "P02"
			case "malformed-claim":
				claimRaw(t, f.Root, "orphan", "invalid JSON")
				code = "P02"
			case "invalid-config":
				worktreePut(t, f.Root, "fireweed.yaml", "version: nonsense\n")
				startCommit(t, f.Root, "Invalid current configuration")
				code = ""
			case "invalid-queue":
				worktreePut(t, f.Root, f.Path, strings.Replace(f.Raw, "new:feature", "ALP-099", -1))
				startCommit(t, f.Root, "Invalid author IDs")
				code = ""
			}
			checkNoClaim := startPreconditionProxy(t, binary, f)
			defer checkNoClaim()
			startFailure(t, f.Root, "current", code)
		})
	}
	for _, slug := range []string{"", "Current", "a/b", "-a", "a b", "a\n"} {
		got, err := w.Start("/missing/start-repository", slug)
		if !reflect.DeepEqual(got, w.StartResult{}) || err == nil || !strings.Contains(err.Error(), "C05") || !policyCorrective(err.Error()) {
			t.Errorf("invalid selector %q must fail before repository access: %+v %v", slug, got, err)
		}
	}
	f := startFixture(t, "sha1", "spec", "")
	checkNoClaim := startPreconditionProxy(t, binary, f)
	defer checkNoClaim()
	startFailure(t, f.Root, "missing", "")
}
func Test_FLOW038StartContextFailureRollsBack(t *testing.T) {
	for _, kind := range []string{"spec", "regenerate"} {
		t.Run(kind, func(t *testing.T) {
			f := startFixture(t, "sha1", kind, "")
			phase := "evals"
			if kind == "regenerate" {
				phase = "impl"
			}
			startBrokenBook(t, f, phase)
			startFailure(t, f.Root, "current", "")
		})
	}
}
func Test_FLOW039StartRegenerateTrackedZoneDeletion(t *testing.T) {
	for _, phase := range []string{"", "impl"} {
		t.Run("phase-"+phase, func(t *testing.T) {
			f := startFixture(t, "sha1", "regenerate", phase)
			before := startTree(t, f.Root)
			got := startSuccess(t, f, f.Root, "impl")
			after := startTree(t, f.Root)
			if _, ok := after["units/a/impl/code.txt"]; ok {
				t.Error("disposable code survived")
			}
			delete(before, f.Path)
			delete(after, f.Path)
			delete(before, "units/a/impl/code.txt")
			if !reflect.DeepEqual(before, after) {
				t.Error("regeneration changed retained files or directories")
			}
			found := false
			for _, d := range got.Diagnostics {
				if d.Code == "D01" && d.Severity == "warning" && d.Path == f.Path && policyCorrective(d.Message) {
					found = true
				}
			}
			if !found {
				t.Error("missing useful no-drill warning", got.Diagnostics)
			}
		})
	}
}

func Test_FLOW036StartBaseWarnings(t *testing.T) {
	for _, kind := range []string{"approved-layout", "missing-base-file", "unapproved", "invalid-base", "wrong-target-base", "stale-approved", "unavailable-base"} {
		t.Run(kind, func(t *testing.T) {
			f := startFixture(t, "sha1", "spec", "")
			want := []string{}
			switch kind {
			case "missing-base-file":
				claimGit(t, f.Root, "branch", "-f", "approval-base", "HEAD^")
				want = []string{"unapproved"}
			case "unapproved":
				worktreePut(t, f.Root, f.Path, strings.Replace(f.Raw, "Start the exact authored change.", "Different authored motivation.", 1))
				startCommit(t, f.Root, "Unapproved authored change")
				want = []string{"unapproved"}
			case "invalid-base", "wrong-target-base":
				base := "not valid YAML or a change"
				if kind == "wrong-target-base" {
					base = strings.ReplaceAll(strings.ReplaceAll(f.Raw, "module: alpha", "module: beta"), "ALP-", "BET-")
				}
				worktreePut(t, f.Root, f.Path, base)
				startCommit(t, f.Root, "Base content differs")
				claimGit(t, f.Root, "branch", "-f", "approval-base", "HEAD")
				worktreePut(t, f.Root, f.Path, f.Raw)
				startCommit(t, f.Root, "Valid current content")
				want = []string{"unapproved"}
			case "stale-approved":
				tip := claimGit(t, f.Root, "-c", "user.name=Start Fixture", "-c", "user.email=start@example.invalid", "commit-tree", claimGit(t, f.Root, "rev-parse", "HEAD^{tree}"), "-p", "HEAD", "-m", "Locally known later base tip")
				claimGit(t, f.Root, "update-ref", "refs/heads/approval-base", tip)
				want = []string{"stale"}
			case "unavailable-base":
				claimGit(t, f.Root, "branch", "-D", "approval-base")
				want = []string{"stale", "unapproved"}
			}
			got := startSuccess(t, f, f.Root, "evals")
			actual := []string{}
			for _, d := range got.Diagnostics {
				if d.Code != "P04" {
					continue
				}
				if d.Severity != "warning" || d.Path != f.Path || !policyCorrective(d.Message) {
					t.Errorf("warning lacks useful path/remedy: %+v", d)
				}
				text := strings.ToLower(d.Message)
				if strings.Contains(text, "not approved") || strings.Contains(text, "unapproved") {
					actual = append(actual, "unapproved")
				} else if strings.Contains(text, "base") && (strings.Contains(text, "tip") || strings.Contains(text, "ancestor") || strings.Contains(text, "contain")) {
					actual = append(actual, "stale")
				} else {
					t.Errorf("warning has no meaningful cause: %+v", d)
				}
			}
			sort.Strings(actual)
			sort.Strings(want)
			if !reflect.DeepEqual(actual, want) {
				t.Errorf("warnings %v want %v: %+v", actual, want, got.Diagnostics)
			}
		})
	}
}
func Test_FLOW032_FLOW035_FLOW038StartInvalidStatesAndPaths(t *testing.T) {
	binary := startProxyBuild(t)
	for _, kind := range []string{"duplicate-slug", "run-only", "phase-only", "applied", "started-handles", "bad-other-module", "change-symlink"} {
		t.Run(kind, func(t *testing.T) {
			f := startFixture(t, "sha1", "spec", "")
			raw := f.Raw
			switch kind {
			case "duplicate-slug":
				worktreePut(t, f.Root, "changes/0002-current.md", raw)
			case "run-only":
				raw = strings.Replace(raw, "module: alpha\n", "module: alpha\nrun: incomplete\n", 1)
			case "phase-only":
				raw = strings.Replace(raw, "module: alpha\n", "module: alpha\nphase: evals\n", 1)
			case "applied":
				raw = strings.Replace(raw, "module: alpha\n", "module: alpha\nrun: ended\nphase: applied\n", 1)
				raw = strings.ReplaceAll(raw, "new:feature", "ALP-051")
			case "started-handles":
				raw = strings.Replace(raw, "module: alpha\n", "module: alpha\nrun: started\nphase: audit\n", 1)
			case "bad-other-module":
				worktreePut(t, f.Root, "units/b/SPEC.md", "---\nname: beta\nid_prefix: BET\ndepends_on: [missing]\n---\n\n## Specs\n\n### BET-001\nThe peer MUST remain valid.\n")
			case "change-symlink":
				if err := os.Remove(filepath.Join(f.Root, f.Path)); err != nil {
					t.Fatal(err)
				}
				worktreePut(t, f.Root, "private/source.md", raw)
				if err := os.Symlink("../private/source.md", filepath.Join(f.Root, f.Path)); err != nil {
					t.Fatal(err)
				}
			}
			if kind != "change-symlink" {
				worktreePut(t, f.Root, f.Path, raw)
			}
			startCommit(t, f.Root, "Invalid selected fixture state")
			checkNoClaim := startPreconditionProxy(t, binary, f)
			defer checkNoClaim()
			code := ""
			err := startFailure(t, f.Root, "current", code)
			if kind == "bad-other-module" && err != nil && (!strings.Contains(err.Error(), "missing") || !strings.Contains(err.Error(), "beta") && !strings.Contains(err.Error(), "units/b")) {
				t.Error("lost useful module/dependency detail", err)
			}
		})
	}
}
func Test_FLOW039StartRegenerateOverridesAndLinks(t *testing.T) {
	for _, kind := range []string{"stricter-overlap", "new-interface", "new-preserve", "tracked-leaf-link", "ignored-disposable"} {
		t.Run(kind, func(t *testing.T) {
			f := startFixture(t, "sha1", "regenerate", "")
			deleted := []string{"units/a/impl/code.txt"}
			outside := ""
			switch kind {
			case "stricter-overlap":
				raw := strings.Replace(f.Raw, "module: alpha\n", "module: alpha\nset:\n  disposable: [impl/**, api.txt, keep/**, ledger/**]\n", 1)
				worktreePut(t, f.Root, f.Path, raw)
			case "new-interface":
				raw := strings.Replace(f.Raw, "module: alpha\n", "module: alpha\nset:\n  interface: [api.txt, impl/**]\n", 1)
				worktreePut(t, f.Root, f.Path, raw)
				deleted = nil
			case "new-preserve":
				raw := strings.Replace(f.Raw, "module: alpha\n", "module: alpha\nset:\n  preserve: [keep/**, impl/code.txt]\n", 1)
				worktreePut(t, f.Root, f.Path, raw)
				deleted = nil
			case "tracked-leaf-link":
				outside = filepath.Join(t.TempDir(), "external retained target")
				if err := os.WriteFile(outside, []byte("Outside bytes remain.\n"), 0600); err != nil {
					t.Fatal(err)
				}
				if err := os.Symlink(outside, filepath.Join(f.Root, "units/a/impl/link.txt")); err != nil {
					t.Fatal(err)
				}
				deleted = append(deleted, "units/a/impl/link.txt")
			case "ignored-disposable":
				worktreePut(t, f.Root, "units/a/impl/ignored-new.txt", "Ignored untracked implementation remains.\n")
			}
			if kind != "ignored-disposable" {
				startCommit(t, f.Root, "Authored projected deletion fixture")
			}
			startFixtureValid(t, f.Root)
			before := startTree(t, f.Root)
			startSuccess(t, f, f.Root, "impl")
			after := startTree(t, f.Root)
			for _, path := range deleted {
				if _, ok := after[path]; ok {
					t.Error("tracked disposable survived", path)
				}
				delete(before, path)
			}
			delete(before, f.Path)
			delete(after, f.Path)
			if !reflect.DeepEqual(before, after) {
				t.Error("projected deletion lost or changed retained paths")
			}
			if outside != "" {
				b, err := os.ReadFile(outside)
				if err != nil || string(b) != "Outside bytes remain.\n" {
					t.Error("followed leaf symlink", err)
				}
			}
		})
	}
}
func startReport(t *testing.T, f startFixtureData, result, base string) {
	t.Helper()
	worktreePut(t, f.Root, ".fireweed/drills/alpha.md", "---\ndate: '2026-01-01'\nbase_commit: '"+base+"'\nagent: fixture\nmodel: deterministic\nresult: "+result+"\n---\n\n## Own evals\n\nRetained report.\n")
	startCommit(t, f.Root, "Commit reviewed report")
}
func Test_FLOW039StartDrillFreshnessWarnings(t *testing.T) {
	for _, kind := range []string{"current", "other-module-apply", "own-module-apply", "failing", "malformed", "unknown-base", "nonancestor"} {
		t.Run(kind, func(t *testing.T) {
			f := startFixture(t, "sha1", "regenerate", "")
			base := claimBase(t, f.Root)
			result := "pass"
			warn := true
			switch kind {
			case "current", "other-module-apply":
				warn = false
			case "failing":
				result = "fail"
			case "unknown-base":
				base = strings.Repeat("a", 40)
			case "nonancestor":
				base = claimGit(t, f.Root, "-c", "user.name=Start Fixture", "-c", "user.email=start@example.invalid", "commit-tree", claimGit(t, f.Root, "rev-parse", "HEAD^{tree}"), "-m", "Unrelated local report base")
			}
			startReport(t, f, result, base)
			if kind == "malformed" {
				worktreePut(t, f.Root, ".fireweed/drills/alpha.md", "---\nresult: [invalid\n---\n")
				startCommit(t, f.Root, "Malformed report")
			}
			if kind == "other-module-apply" || kind == "own-module-apply" {
				name, prefix := "beta", "BET"
				if kind == "own-module-apply" {
					name, prefix = "alpha", "ALP"
				}
				worktreePut(t, f.Root, "changes/archive/0010-later.md", "---\nmodule: "+name+"\nrun: later-apply\nphase: applied\n---\n\n## Motivation\n\nRecord an actual later apply.\n\n## Added\n\n### "+prefix+"-060 [review]\nThe later behavior MUST work.\n")
				startCommit(t, f.Root, "Misleading subject: not an implementation commit")
			}
			startFixtureValid(t, f.Root)
			before := startTree(t, f.Root)
			got := startSuccess(t, f, f.Root, "impl")
			found := false
			for _, d := range got.Diagnostics {
				if d.Code == "D01" {
					found = true
					if d.Severity != "warning" || d.Path != f.Path || !policyCorrective(d.Message) {
						t.Error("invalid drill diagnostic", d)
					}
				}
			}
			if found != warn {
				t.Errorf("D01 %v want %v: %+v", found, warn, got.Diagnostics)
			}
			after := startTree(t, f.Root)
			for path, node := range before {
				if strings.HasPrefix(path, ".fireweed/drills/") && after[path] != node {
					t.Error("rewrote drill report", path)
				}
			}
		})
	}
}
func Test_FLOW033_FLOW040StartConcurrentClaims(t *testing.T) {
	for _, format := range []string{"sha1", "sha256"} {
		for _, linked := range []bool{false, true} {
			t.Run(fmt.Sprintf("%s-linked-%v", format, linked), func(t *testing.T) {
				f := startFixture(t, format, "spec", "")
				roots := []string{f.Root, f.Root}
				if linked {
					roots[1] = claimPeer(t, f.Root)
					if err := os.Chmod(filepath.Join(roots[1], f.Path), 0600); err != nil {
						t.Fatal(err)
					}
					startFixtureValid(t, roots[1])
				}
				trees := []map[string]startNode{startTree(t, roots[0]), startTree(t, roots[1])}
				git := []startGitState{startGitSnapshot(t, roots[0]), startGitSnapshot(t, roots[1])}
				type response struct {
					Index  int
					Result w.StartResult
					Err    error
				}
				barrier := make(chan struct{})
				done := make(chan response, 2)
				for i, root := range roots {
					go func(i int, root string) {
						<-barrier
						got, err := w.Start(root, "current")
						done <- response{i, got, err}
					}(i, root)
				}
				close(barrier)
				results := []response{}
				for len(results) < 2 {
					select {
					case r := <-done:
						results = append(results, r)
					case <-time.After(35 * time.Second):
						t.Fatal("local competing starts did not finish within their bounded deadline")
					}
				}
				winners := []response{}
				for _, r := range results {
					if r.Err == nil {
						winners = append(winners, r)
					} else if !reflect.DeepEqual(r.Result, w.StartResult{}) || !policyCorrective(r.Err.Error()) {
						t.Errorf("loser result/error %+v %v", r.Result, r.Err)
					}
				}
				if len(winners) != 1 {
					t.Fatalf("winner count %d want 1: %+v", len(winners), results)
				}
				winner := winners[0]
				finalBytes, readErr := os.ReadFile(filepath.Join(roots[winner.Index], f.Path))
				if readErr != nil || string(finalBytes) != winner.Result.Claim.State {
					t.Error("loser clobbered winner selected bytes", readErr)
				}
				cl, err := w.LoadClaim(f.Root, "current")
				if err != nil || cl != winner.Result.Claim || cl.Worktree != roots[winner.Index] || cl.Object == "" {
					t.Error("winner not the exact shared reservation", cl, err)
				}
				for i, root := range roots {
					startUnrelated(t, root, git[i], "current")
					after := startTree(t, root)
					before := trees[i]
					if !linked || i == winner.Index {
						delete(before, f.Path)
						delete(after, f.Path)
					}
					if !reflect.DeepEqual(before, after) {
						t.Errorf("contender %d changed unrelated or losing files", i)
					}
				}
			})
		}
	}
}
func Test_FLOW032_FLOW033_FLOW037StartWarningsAndInclusiveEligibility(t *testing.T) {
	for _, kind := range []string{"unrelated-blocker", "forward-blocker", "reverse-interface-blocker", "static-warning"} {
		t.Run(kind, func(t *testing.T) {
			f := startFixture(t, "sha1", "spec", "")
			if kind == "static-warning" {
				worktreePut(t, f.Root, "outside/SPEC.md", "Not a configured module.\n")
			} else {
				before := "---\nmodule: beta\nrun: held-run\nphase: evals\nblocked: waiting for repair\n---\n\n## Motivation\n\nHold the beta work.\n\n## Added\n\n### BET-051\nThe held behavior MUST remain blocked.\n"
				worktreePut(t, f.Root, "changes/0000-held.md", before)
				path := "units/a/SPEC.md"
				dependency := "beta"
				if kind == "reverse-interface-blocker" {
					path = "units/b/SPEC.md"
					dependency = "alpha"
				}
				if kind != "unrelated-blocker" {
					b, err := os.ReadFile(filepath.Join(f.Root, path))
					if err != nil {
						t.Fatal(err)
					}
					prefix := "ALP"
					if path == "units/b/SPEC.md" {
						prefix = "BET"
					}
					text := strings.Replace(string(b), "id_prefix: "+prefix+"\n", "id_prefix: "+prefix+"\ndepends_on: ["+dependency+"]\n", 1)
					worktreePut(t, f.Root, path, text)
				}
			}
			startCommit(t, f.Root, "Compose checked prefix and warning fixture")
			s := startFixtureValid(t, f.Root)
			eligible, err := w.Eligible(s, "current")
			if err != nil {
				t.Fatal("public eligibility premise", err)
			}
			if kind == "forward-blocker" || kind == "reverse-interface-blocker" {
				if eligible {
					t.Fatal("blocked prefix premise")
				}
				startFailure(t, f.Root, "current", "P03")
				return
			}
			if !eligible {
				t.Fatal("independent selected head premise")
			}
			got := startSuccess(t, f, f.Root, "evals")
			if kind == "static-warning" {
				found := false
				for _, d := range got.Diagnostics {
					if d.Code == "C02" && d.Severity == "warning" && d.Path == "outside/SPEC.md" {
						found = true
					}
				}
				if !found {
					t.Error("static warning disappeared", got.Diagnostics)
				}
			}
		})
	}
}
func Test_FLOW034StartReservationsAndAllBindings(t *testing.T) {
	f := startFixture(t, "sha1", "spec", "")
	later := "---\nmodule: alpha\nrun: previously-assigned\nphase: evals\n---\n\n## Motivation\n\nRetain the earlier assigned number.\n\n## Added\n\n### ALP-060\nThe assigned behavior MUST retain its number.\n"
	worktreePut(t, f.Root, "changes/0002-assigned.md", later)
	raw := strings.Replace(f.Raw, "module: alpha\n", "module: alpha\nalready_satisfied: [new:feature]\nset:\n  hooks:\n    setup: 'record new:feature'\n", 1)
	raw = strings.Replace(raw, "Retained preamble.", "Retained new:feature preamble.", 1)
	raw = strings.Replace(raw, "Retain the public shape.", "Retain the public new:feature shape.", 1)
	raw += "\n## Extra new:feature\n\nA new:feature example MUST remain bound.\n\n```text\nnew:feature\n```\n"
	worktreePut(t, f.Root, f.Path, raw)
	startCommit(t, f.Root, "Declared bindings and existing assigned reservation")
	startFixtureValid(t, f.Root)
	got := startSuccess(t, f, f.Root, "evals")
	if len(got.Context.Change.Added) != 1 || got.Context.Change.Added[0].ID != "ALP-061" || strings.Contains(got.Context.ChangeText, "new:feature") || !strings.Contains(got.Context.ChangeText, "Extra ALP-061") || !reflect.DeepEqual(got.Context.Change.AlreadySatisfied, []string{"ALP-061"}) {
		t.Errorf("late reservations/bindings lost: %+v", got.Context.Change)
	}
	hooks, ok := got.Context.Change.Set["hooks"].(map[string]any)
	if !ok || hooks["setup"] != "record ALP-061" {
		t.Error("nested metadata was not bound", got.Context.Change.Set)
	}
	b, err := os.ReadFile(filepath.Join(f.Root, "changes/0002-assigned.md"))
	if err != nil || string(b) != later {
		t.Error("changed unrelated assigned change", err)
	}
}
func Test_FLOW032StartValidSelectorBoundaries(t *testing.T) {
	for _, slug := range []string{"0", "a-", "1alpha", "a--b"} {
		t.Run(slug, func(t *testing.T) {
			f := startFixture(t, "sha1", "spec", "")
			old := f.Path
			f.Path = "changes/0001-" + slug + ".md"
			if err := os.Rename(filepath.Join(f.Root, old), filepath.Join(f.Root, f.Path)); err != nil {
				t.Fatal(err)
			}
			startCommit(t, f.Root, "Exact boundary selector")
			startFixtureValid(t, f.Root)
			got := startSuccess(t, f, f.Root, "evals")
			if got.Claim.Change != slug || got.Context.Change.Slug != slug {
				t.Error("selector not exact", got.Claim.Change, got.Context.Change.Slug)
			}
		})
	}
}
func Test_FLOW034StartFreshRunAfterRelease(t *testing.T) {
	f := startFixture(t, "sha1", "spec", "")
	first := startSuccess(t, f, f.Root, "evals")
	if err := w.ReleaseClaim(f.Root, "current", first.Claim.Object); err != nil {
		t.Fatal(err)
	}
	worktreePut(t, f.Root, f.Path, f.Raw)
	if clean := claimGit(t, f.Root, "status", "--porcelain"); clean != "" {
		t.Fatal("fresh replay premise", clean)
	}
	second := startSuccess(t, f, f.Root, "evals")
	if first.Claim.Run == second.Claim.Run || first.Claim.Base != second.Claim.Base || !reflect.DeepEqual(first.Context.Change.Added, second.Context.Change.Added) {
		t.Error("fresh replay reused run or renumbered unchanged IDs", first.Claim, second.Claim)
	}
}
func Test_FLOW032_FLOW040StartNonRepository(t *testing.T) {
	root := t.TempDir()
	worktreePut(t, root, "sentinel", "Unchanged nonrepository data.")
	before := startTree(t, root)
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	env := os.Environ()
	got, err := w.Start(root, "current")
	if !reflect.DeepEqual(got, w.StartResult{}) || err == nil || !policyCorrective(err.Error()) {
		t.Error("nonrepository rejection", got, err)
	}
	afterCwd, e := os.Getwd()
	if e != nil || cwd != afterCwd || !reflect.DeepEqual(env, os.Environ()) || !reflect.DeepEqual(before, startTree(t, root)) {
		t.Error("nonrepository failure changed files or parent state")
	}
}
func Test_FLOW037_FLOW038StartUnsafeRequiredContext(t *testing.T) {
	for _, kind := range []string{"interface", "playbook"} {
		t.Run(kind, func(t *testing.T) {
			f := startFixture(t, "sha1", "spec", "")
			outside := filepath.Join(t.TempDir(), "retained outside bytes")
			if err := os.WriteFile(outside, []byte("External bytes MUST remain intact.\n"), 0640); err != nil {
				t.Fatal(err)
			}
			path := "units/a/api.txt"
			code := ""
			if kind == "playbook" {
				path = ".fireweed/playbooks/evals.md"
				code = ""
				if err := os.MkdirAll(filepath.Dir(filepath.Join(f.Root, path)), 0755); err != nil {
					t.Fatal(err)
				}
			} else {
				if err := os.Remove(filepath.Join(f.Root, path)); err != nil {
					t.Fatal(err)
				}
			}
			if err := os.Symlink(outside, filepath.Join(f.Root, path)); err != nil {
				t.Fatal(err)
			}
			startCommit(t, f.Root, "Unsafe required context path")
			err := startFailure(t, f.Root, "current", code)
			if err != nil && !strings.Contains(err.Error(), path) {
				t.Error("unsafe context lacks required-path detail", err)
			}
			b, err := os.ReadFile(outside)
			info, e := os.Stat(outside)
			if err != nil || e != nil || string(b) != "External bytes MUST remain intact.\n" || info.Mode().Perm() != 0640 {
				t.Error("changed external context target", err, e)
			}
		})
	}
}
func Test_FLOW033_FLOW040StartPreservesForeignClaimsAndBookkeeping(t *testing.T) {
	f := startFixture(t, "sha1", "spec", "")
	peer := claimPeer(t, f.Root)
	foreign, err := w.CreateClaim(peer, "unrelated", "foreign-run", "HEAD", "Unrelated opaque state.\n")
	if err != nil {
		t.Fatal(err)
	}
	for _, ref := range []string{"refs/fireweed/attempts/current/1", "refs/fireweed/drills/alpha/20260101"} {
		claimGit(t, f.Root, "update-ref", ref, "HEAD")
	}
	worktreePut(t, f.Root, ".fireweed/drills/alpha.md", "Retained reviewed report without metadata.\n")
	startCommit(t, f.Root, "Retained report for a different operation")
	common := claimCommon(t, f.Root)
	metadata := func() map[string]startNode {
		t.Helper()
		all := startTree(t, common)
		for key, node := range all {
			if node.Mode.IsDir() || key == "refs/fireweed/claims/current" || key == "logs/refs/fireweed/claims/current" {
				delete(all, key)
			}
		}
		return all
	}
	before := metadata()
	peerFiles := startTree(t, peer)
	startSuccess(t, f, f.Root, "evals")
	after := metadata()
	for key, node := range before {
		if actual, ok := after[key]; !ok || actual != node {
			t.Error("changed existing common-Git metadata/object", key)
		}
	}
	for key := range after {
		if _, ok := before[key]; !ok && !strings.HasPrefix(key, "objects/") {
			t.Error("created unrelated common-Git metadata", key)
		}
	}
	loaded, err := w.LoadClaim(f.Root, "unrelated")
	if err != nil || loaded != foreign {
		t.Error("mutated foreign claim", loaded, err)
	}
	if !reflect.DeepEqual(peerFiles, startTree(t, peer)) {
		t.Error("changed another linked worktree's files")
	}
}
func Test_FLOW033_FLOW037StartInclusiveContextOnce(t *testing.T) {
	f := startFixture(t, "sha1", "spec", "")
	old := f.Path
	f.Path = "changes/0002-current.md"
	if err := os.Rename(filepath.Join(f.Root, old), filepath.Join(f.Root, f.Path)); err != nil {
		t.Fatal(err)
	}
	raw := strings.Replace(f.Raw, "module: alpha\n", "module: alpha\nset:\n  depends_on: [beta]\n  interface: [pub/**]\n", 1)
	worktreePut(t, f.Root, f.Path, raw)
	worktreePut(t, f.Root, "units/e/SPEC.md", "---\nname: echo\nid_prefix: ECH\ndepends_on: [beta]\n---\n\n## Specs\n\n### ECH-001 [review]\nThe blocker MUST retain its dependency.\n")
	worktreePut(t, f.Root, "changes/0000-blocked.md", "---\nmodule: echo\nrun: blocked-run\nphase: evals\nblocked: pending repair\n---\n\n## Motivation\n\nBlock echo while its dependency may have held work.\n")
	before := "---\nmodule: beta\nrun: earlier-run\nphase: impl\nset:\n  interface: [futurepub/**]\n---\n\n## Motivation\n\nProject the held dependency contract.\n\n## Interface\n\nChange beta exports while echo holds it.\n\n## Added\n\n### BET-052\nThe earlier behavior MUST be projected exactly once.\n"
	worktreePut(t, f.Root, "changes/0001-before.md", before)
	worktreePut(t, f.Root, "units/a/pub/current.txt", "Projected alpha interface.\n")
	worktreePut(t, f.Root, "units/b/futurepub/dependency.txt", "Projected direct beta interface.\n")
	startCommit(t, f.Root, "Projected earlier held interface and current metadata")
	s := startFixtureValid(t, f.Root)
	eligible, err := w.Eligible(s, "current")
	if err != nil || !eligible {
		t.Fatal("inclusive runnable fixture premise", eligible, err)
	}
	got := startSuccess(t, f, f.Root, "evals")
	if !reflect.DeepEqual(got.Context.Module.Interface, []string{"pub/**"}) || !reflect.DeepEqual(got.Context.Module.DependsOn, []string{"beta"}) || len(got.Context.Module.Specs) != 2 || len(got.Context.Dependencies) != 1 || got.Context.Dependencies[0].Name != "beta" || len(got.Context.Dependencies[0].Specs) != 2 || !reflect.DeepEqual(got.Context.Dependencies[0].Interface, []string{"futurepub/**"}) {
		t.Error("lost or double-applied selected/prefix graph", got.Context.Module, got.Context.Dependencies)
	}
	files := map[string]string{}
	for _, file := range got.Context.Files {
		files[file.Path] = file.Content
	}
	if files["units/a/pub/current.txt"] != "Projected alpha interface.\n" || files["units/b/futurepub/dependency.txt"] != "Projected direct beta interface.\n" {
		t.Error("missing projected interface files", files)
	}
	for _, path := range []string{"units/a/api.txt", "units/b/api.txt", "units/e/SPEC.md"} {
		if _, ok := files[path]; ok {
			t.Error("included stale/transitive context file", path)
		}
	}
	b, err := os.ReadFile(filepath.Join(f.Root, "changes/0001-before.md"))
	if err != nil || string(b) != before {
		t.Error("renumbered or rewrote prior held change", err)
	}
}
func Test_FLOW032_FLOW035StartUsesCurrentCleanState(t *testing.T) {
	f := startFixture(t, "sha1", "spec", "")
	state := strings.Replace(f.Raw, "module: alpha\n", "module: alpha\nrun: existing-run\nphase: evals\n", 1)
	state = strings.ReplaceAll(state, "new:feature", "ALP-051")
	if _, err := w.CreateClaim(f.Root, "current", "existing-run", "HEAD", state); err != nil {
		t.Fatal(err)
	}
	worktreePut(t, f.Root, f.Path, "Broken current feedback bytes.\n")
	worktreePut(t, f.Root, "fireweed.yaml", "broken current config [\n")
	if _, _, err := w.LoadApproval(f.Root, "current"); err != nil {
		t.Fatal("owned repair snapshot must be independently readable", err)
	}
	startFailure(t, f.Root, "current", "P01")
}
func Test_FLOW032_FLOW039StartRejectsUnsafeAncestorBeforeDeletion(t *testing.T) {
	f := startFixture(t, "sha1", "regenerate", "")
	outside := t.TempDir()
	worktreePut(t, outside, "code.txt", "Outside implementation target.")
	worktreePut(t, outside, "protected.txt", "Outside protected target.")
	if err := os.RemoveAll(filepath.Join(f.Root, "units/a/impl")); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(outside, filepath.Join(f.Root, "units/a/impl")); err != nil {
		t.Fatal(err)
	}
	before := startTree(t, outside)
	startFailure(t, f.Root, "current", "P01")
	if !reflect.DeepEqual(before, startTree(t, outside)) {
		t.Error("followed unsafe ancestor into outside deletion")
	}
}
func Test_FLOW033StartUnreadableClaim(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("POSIX permission control; malformed and foreign claim assertions run separately")
	}
	f := startFixture(t, "sha1", "spec", "")
	cl := claimSeed(t, f.Root, "orphan", f.Raw)
	object := filepath.Join(claimCommon(t, f.Root), "objects", cl.Object[:2], cl.Object[2:])
	originalInfo, err := os.Lstat(object)
	if err != nil {
		t.Fatal(err)
	}
	// Capture bytes before making the object unreadable; restore fixture mode only
	// after checking Start retained the fault, then compare the complete snapshot.
	checkStorage := startStorage(t, f.Root, "current")
	defer checkStorage()
	tree := startTree(t, f.Root)
	git := startGitSnapshot(t, f.Root)
	if err := os.Chmod(object, 0000); err != nil {
		t.Fatal(err)
	}
	defer func() {
		info, err := os.Lstat(object)
		if err != nil || !info.Mode().IsRegular() {
			t.Error("Start replaced or removed unreadable claim object", err)
			return
		}
		if info.Mode().Perm() != 0 {
			t.Error("Start changed unreadable claim object mode")
		}
		if err := os.Chmod(object, originalInfo.Mode().Perm()); err != nil {
			t.Error("fixture mode restoration", err)
		}
	}()
	if _, err := os.ReadFile(object); err == nil {
		t.Skip("host privileges bypass the permission control")
	}
	if _, err := w.Claims(f.Root); err == nil || !strings.Contains(err.Error(), "P02") {
		t.Fatal("public unreadable claim control", err)
	}
	got, err := w.Start(f.Root, "current")
	if !reflect.DeepEqual(got, w.StartResult{}) || err == nil || !strings.Contains(err.Error(), "P02") || !policyCorrective(err.Error()) {
		t.Error("unreadable claim must fail closed with useful P02", got, err)
	}
	if !reflect.DeepEqual(tree, startTree(t, f.Root)) || !reflect.DeepEqual(git, startGitSnapshot(t, f.Root)) {
		t.Error("unreadable claim rejection changed working/Git state")
	}

}
func Test_FLOW034_FLOW038StartRetainsChangeModes(t *testing.T) {
	for _, mode := range []fs.FileMode{0644, 0751} {
		t.Run(fmt.Sprintf("%o", mode), func(t *testing.T) {
			f := startFixture(t, "sha1", "spec", "")
			f.Mode = mode
			if err := os.Chmod(filepath.Join(f.Root, f.Path), mode); err != nil {
				t.Fatal(err)
			}
			if mode&0111 != 0 {
				startCommit(t, f.Root, "Commit original executable mode")
			}
			startFixtureValid(t, f.Root)
			startSuccess(t, f, f.Root, "evals")
		})
	}
}

// Snapshot external temporary storage and common Git metadata around the subject.
// New immutable loose objects and the owned claim/ref log are legitimate artifacts.
func startStorage(t *testing.T, root, slug string) func() {
	t.Helper()
	temporary := t.TempDir()
	for _, name := range []string{"TMPDIR", "TMP", "TEMP"} {
		t.Setenv(name, temporary)
	}
	claimGit(t, root, "--version") // warm toolchain-owned Git shim caches before subject snapshot
	tempBefore := startTree(t, temporary)
	common := startGitStdout(t, root, "rev-parse", "--git-common-dir")
	if !filepath.IsAbs(common) {
		common = filepath.Join(root, common)
	}
	common = stateCanonicalDirectory(t, common)
	before := startTree(t, common)
	objectWidth := len(startGitStdout(t, root, "rev-parse", "HEAD"))
	return func() {
		t.Helper()
		if got := startTree(t, temporary); !reflect.DeepEqual(got, tempBefore) {
			t.Error("leaked operation temporary files outside working tree", got)
		}
		for _, issue := range startStorageIssues(before, startTree(t, common), slug, objectWidth) {
			t.Error(issue)
		}
	}
}
func startStorageIssues(before, after map[string]startNode, slug string, objectWidth int) []string {
	issues := []string{}
	allowed := map[string]bool{}
	// Native Git retains the fireweed ref namespace after the final claim release.
	if node, ok := after["refs/fireweed"]; ok && node.Mode.IsDir() {
		allowed["refs/fireweed"] = true
	}
	owned := func(path string) bool {
		return path == "refs/fireweed/claims/"+slug || path == "logs/refs/fireweed/claims/"+slug
	}
	for path, node := range after {
		if _, existed := before[path]; !existed && node.Mode.IsRegular() && (owned(path) || startLooseObject(path, node, objectWidth)) {
			allowed[path] = true
			for parent := filepath.ToSlash(filepath.Dir(path)); parent != "."; parent = filepath.ToSlash(filepath.Dir(parent)) {
				allowed[parent] = true
			}
		}
	}
	for path, node := range before {
		if got, ok := after[path]; !ok || got != node {
			issues = append(issues, "changed existing common-Git bookkeeping/object/type/mode: "+path)
		}
	}
	for path := range after {
		if _, existed := before[path]; !existed && !allowed[path] {
			issues = append(issues, "leaked temporary/unrelated common-Git node: "+path)
		}
	}
	sort.Strings(issues)
	return issues
}
func startLooseObject(path string, node startNode, objectWidth int) bool {
	if (objectWidth != 40 && objectWidth != 64) || len(strings.ReplaceAll(strings.TrimPrefix(path, "objects/"), "/", "")) != objectWidth || !node.Mode.IsRegular() || !regexp.MustCompile(`^objects/[0-9a-f]{2}/(?:[0-9a-f]{38}|[0-9a-f]{62})$`).MatchString(path) {
		return false
	}
	compressed := bytes.NewBufferString(node.Bytes)
	r, err := zlib.NewReader(compressed)
	if err != nil {
		return false
	}
	raw, err := io.ReadAll(r)
	closed := r.Close()
	if err != nil || closed != nil || compressed.Len() != 0 {
		return false
	}
	header, body, found := bytes.Cut(raw, []byte{0})
	if !found {
		return false
	}
	if !bytes.HasPrefix(header, []byte("blob ")) {
		return false
	}
	sizeText := string(header[len("blob "):])
	if sizeText == "" || (len(sizeText) > 1 && sizeText[0] == '0') {
		return false
	}
	for _, digit := range sizeText {
		if digit < '0' || digit > '9' {
			return false
		}
	}
	size, err := strconv.ParseUint(sizeText, 10, 64)
	if err != nil || size != uint64(len(body)) {
		return false
	}

	want := strings.ReplaceAll(strings.TrimPrefix(path, "objects/"), "/", "")
	if len(want) == 40 {
		return fmt.Sprintf("%x", sha1.Sum(raw)) == want
	}
	return fmt.Sprintf("%x", sha256.Sum256(raw)) == want
}

func Test_FLOW037StartReadableStaleOverride(t *testing.T) {
	for _, kind := range []string{"spec", "regenerate"} {
		t.Run(kind, func(t *testing.T) {
			f := startFixture(t, "sha1", kind, "")
			phase := "evals"
			if kind == "regenerate" {
				phase = "impl"
			}
			worktreePut(t, f.Root, ".fireweed/playbooks/"+phase+".md", "---\nbased_on: 00000000\n---\n\nReadable stale instructions.\n")
			startCommit(t, f.Root, "Readable stale playbook")
			book, err := p.ResolvePlaybook(f.Root, phase)
			if err != nil || !book.Overridden || !book.Stale {
				t.Fatal("stale readable premise", book, err)
			}
			got := startSuccess(t, f, f.Root, phase)
			if got.Context.Playbook != book.Text {
				t.Error("lost readable stale override context")
			}
		})
	}
}

func Test_FLOW034StartOriginalDocumentFixtureControls(t *testing.T) {
	for _, kind := range []string{"spec", "fix", "maintenance", "regenerate"} {
		for _, phase := range []string{"", "evals", "audit", "impl"} {
			t.Run(kind+"-"+phase, func(t *testing.T) {
				f := startFixture(t, "sha1", kind, phase)
				snapshot, err := p.LoadAt(f.Root, claimBase(t, f.Root))
				if err != nil {
					t.Fatal(err)
				}
				allocated, _, err := p.Allocate(snapshot, "current")
				if err != nil {
					t.Fatal(err)
				}
				allocated.Document.Metadata["run"] = "fixture-run"
				expectedPhase := phase
				if phase == "" {
					expectedPhase = "evals"
					if kind == "maintenance" || kind == "regenerate" {
						expectedPhase = "impl"
					}
				}
				allocated.Document.Metadata["phase"] = expectedPhase
				raw, err := c.RenderMarkdown(allocated.Document)
				if err != nil {
					t.Fatal(err)
				}
				parsed, err := c.ParseChange(raw, f.Path, "ALP")
				if err != nil || parsed.Run != "fixture-run" || parsed.Phase != expectedPhase {
					t.Fatal("original-document expected-state premise", parsed, err)
				}
				if parsed.Document.Preamble != allocated.Document.Preamble || len(parsed.Document.Sections) != len(allocated.Document.Sections) {
					t.Fatal("fixture loses retained author document")
				}
			})
		}
	}
}

func startBrokenBook(t *testing.T, f startFixtureData, phase string) {
	t.Helper()
	worktreePut(t, f.Root, ".fireweed/playbooks/"+phase+".md/entry.txt", "A directory cannot be read as the selected playbook.\n")
	startCommit(t, f.Root, "Selected playbook is an unreadable directory")
	if _, err := p.ResolvePlaybook(f.Root, phase); err == nil || !strings.Contains(err.Error(), "D05") {
		t.Fatal("public context/playbook failure premise", err)
	}
}
func Test_FLOW037StartContextFailureFixtureControls(t *testing.T) {
	for _, kind := range []string{"spec", "regenerate"} {
		t.Run(kind, func(t *testing.T) {
			f := startFixture(t, "sha1", kind, "")
			phase := "evals"
			if kind == "regenerate" {
				phase = "impl"
			}
			startBrokenBook(t, f, phase)
		})
	}
}

func Test_FLOW038_FLOW040StartStorageFixtureControls(t *testing.T) {
	for _, format := range []string{"sha1", "sha256"} {
		for _, slug := range []string{"current", "0", "a-", "1alpha", "a--b"} {
			t.Run(format+"-"+slug, func(t *testing.T) {
				f := startFixture(t, format, "spec", "")
				objectWidth := len(startGitStdout(t, f.Root, "rev-parse", "HEAD"))
				common := claimCommon(t, f.Root)
				before := startTree(t, common)
				claim, err := w.CreateClaim(f.Root, slug, "storage-control", "HEAD", "Opaque control state.\n")
				if err != nil {
					t.Fatal(err)
				}
				after := startTree(t, common)
				if issues := startStorageIssues(before, after, slug, objectWidth); len(issues) != 0 {
					t.Fatal("legitimate selected claim/object artifacts rejected", issues)
				}
				if err := w.ReleaseClaim(f.Root, slug, claim.Object); err != nil {
					t.Fatal(err)
				}
				if issues := startStorageIssues(before, startTree(t, common), slug, objectWidth); len(issues) != 0 {
					t.Fatal("legitimate Git release artifacts rejected", issues)
				}
				// Use independent captured-node mutations to prove each rejection channel.
				for _, mode := range []string{"empty-directory", "fake-object", "wrong-object-hash", "existing-directory-mode"} {
					t.Run(mode, func(t *testing.T) {
						changed := map[string]startNode{}
						for path, node := range after {
							changed[path] = node
						}
						switch mode {
						case "empty-directory":
							changed["fireweed-temporary-empty"] = startNode{Mode: os.ModeDir | 0700}
						case "fake-object":
							suffix := strings.Repeat("b", 38)
							if format == "sha256" {
								suffix = strings.Repeat("b", 62)
							}
							changed["objects/aa/"+suffix] = startNode{Mode: 0444, Bytes: "not a compressed Git object"}
						case "wrong-object-hash":
							for path, node := range after {
								if _, old := before[path]; !old && startLooseObject(path, node, objectWidth) {
									suffix := strings.Repeat("c", 38)
									if format == "sha256" {
										suffix = strings.Repeat("c", 62)
									}
									changed["objects/aa/"+suffix] = node
									break
								}
							}
						case "existing-directory-mode":
							for path, node := range before {
								if node.Mode.IsDir() {
									node.Mode ^= 0100
									changed[path] = node
									break
								}
							}
						}
						if issues := startStorageIssues(before, changed, slug, objectWidth); len(issues) == 0 {
							t.Fatal("storage assertion does not detect control", mode)
						}
					})
				}
			})
		}
	}
}
func Test_FLOW035_FLOW040StartCommittedResumeDeterminism(t *testing.T) {
	for _, phase := range []string{"evals", "audit", "impl"} {
		t.Run(phase, func(t *testing.T) {
			f := startFixture(t, "sha1", "spec", phase)
			first := startSuccess(t, f, f.Root, phase)
			if err := w.ReleaseClaim(f.Root, "current", first.Claim.Object); err != nil {
				t.Fatal(err)
			}
			if err := os.WriteFile(filepath.Join(f.Root, f.Path), []byte(f.Raw), f.Mode); err != nil {
				t.Fatal(err)
			}
			startFixtureValid(t, f.Root)
			if status := claimGit(t, f.Root, "status", "--porcelain"); status != "" {
				t.Fatal("identical committed resume premise", status)
			}
			second := startSuccess(t, f, f.Root, phase)
			if !reflect.DeepEqual(first, second) {
				t.Error("identical committed resume varied complete result including bytes/token/context/diagnostics", first, second)
			}
		})
	}
}

func Test_FLOW038_FLOW040StartLooseObjectFixtureControls(t *testing.T) {
	for _, format := range []string{"sha1", "sha256"} {
		for _, mode := range []string{"valid-blob", "wrong-type", "wrong-size", "bad-compression", "wrong-hash", "symlink", "tab-header", "signed-size", "zero-padded-size", "trailing-size-space", "opposite-format"} {
			t.Run(format+"-"+mode, func(t *testing.T) {
				payload := []byte("blob 4\x00data")
				switch mode {
				case "wrong-type":
					payload = []byte("tree 4\x00data")
				case "wrong-size":
					payload = []byte("blob 40\x00data")
				case "tab-header":
					payload = []byte("blob\t4\x00data")
				case "signed-size":
					payload = []byte("blob +4\x00data")
				case "zero-padded-size":
					payload = []byte("blob 04\x00data")
				case "trailing-size-space":
					payload = []byte("blob 4 \x00data")
				}
				var zipped bytes.Buffer
				writer := zlib.NewWriter(&zipped)
				if _, err := writer.Write(payload); err != nil {
					t.Fatal(err)
				}
				if err := writer.Close(); err != nil {
					t.Fatal(err)
				}
				oid := fmt.Sprintf("%x", sha1.Sum(payload))
				if format == "sha256" {
					oid = fmt.Sprintf("%x", sha256.Sum256(payload))
				}
				objectWidth := len(oid)
				if mode == "opposite-format" {
					if format == "sha1" {
						oid = fmt.Sprintf("%x", sha256.Sum256(payload))
					} else {
						oid = fmt.Sprintf("%x", sha1.Sum(payload))
					}
				}
				node := startNode{Mode: 0444, Bytes: zipped.String()}
				switch mode {
				case "bad-compression":
					node.Bytes = "not zlib"
				case "wrong-hash":
					oid = strings.Repeat("a", len(oid))
				case "symlink":
					node.Mode = os.ModeSymlink | 0777
				}
				path := "objects/" + oid[:2] + "/" + oid[2:]
				if got := startLooseObject(path, node, objectWidth); got != (mode == "valid-blob") {
					t.Error("object validator positive/negative control", mode, got)
				}
			})
		}
	}
}

func Test_FLOW032_FLOW037StartPublicErrorFixtureControls(t *testing.T) {
	t.Run("dependency", func(t *testing.T) {
		f := startFixture(t, "sha1", "spec", "")
		worktreePut(t, f.Root, "units/b/SPEC.md", "---\nname: beta\nid_prefix: BET\ndepends_on: [missing]\n---\n\n## Specs\n\n### BET-001\nThe peer MUST remain valid.\n")
		startCommit(t, f.Root, "Unknown dependency public-error fixture")
		snapshot, err := p.Load(f.Root)
		if err != nil {
			t.Fatal(err)
		}
		found := false
		for _, d := range p.Check(snapshot) {
			if d.Code == "C06" && strings.Contains(d.Message, "missing") {
				found = true
			}
		}
		if !found {
			t.Fatal("unknown dependency does not establish public C06 premise")
		}
	})
	t.Run("unsafe-interface", func(t *testing.T) {
		f := startFixture(t, "sha1", "spec", "evals")
		outside := filepath.Join(t.TempDir(), "outside-api")
		if err := os.WriteFile(outside, []byte("External interface bytes.\n"), 0640); err != nil {
			t.Fatal(err)
		}
		if err := os.Remove(filepath.Join(f.Root, "units/a/api.txt")); err != nil {
			t.Fatal(err)
		}
		if err := os.Symlink(outside, filepath.Join(f.Root, "units/a/api.txt")); err != nil {
			t.Fatal(err)
		}
		startCommit(t, f.Root, "Unsafe context public-error fixture")
		snapshot, err := p.Load(f.Root)
		if err != nil {
			t.Fatal(err)
		}
		if _, err := w.BuildContext(snapshot, "current"); err == nil || !strings.Contains(err.Error(), "C01") || !strings.Contains(err.Error(), "units/a/api.txt") {
			t.Fatal("unsafe interface does not establish public C01 premise", err)
		}
	})
}

func Test_FLOW038_FLOW040StartWarnedStorageControl(t *testing.T) {
	f := startFixture(t, "sha1", "spec", "")
	binary := startProxyBuild(t)
	startProxyInstall(t, binary, startProxyConfig{Mode: "warn"})
	checkStorage := startStorage(t, f.Root, "current")
	defer checkStorage()
	claim, err := w.CreateClaim(f.Root, "current", "warned-storage-control", "HEAD", "Opaque warned storage bytes.\n")
	if err != nil {
		t.Fatal(err)
	}
	if err := w.ReleaseClaim(f.Root, "current", claim.Object); err != nil {
		t.Fatal(err)
	}
}
