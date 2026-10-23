package evals

import (
	"bytes"
	"compress/zlib"
	"context"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	w "github.com/lurkshark/fireweed/packages/workflow"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"testing"
	"time"
)

// The transport endpoint only leaves a local marker and fails. No server,
// credentials, bootstrap or production implementation participates.
func claimLocalTransport(t *testing.T) string {
	t.Helper()
	dir := stateCanonicalDirectory(t, t.TempDir())
	source := filepath.Join(dir, "transport.go")
	worktreePut(t, dir, "transport.go", `package main
import "os"
func main(){f,e:=os.OpenFile(os.Getenv("FW_CLAIM_LOCAL_MARKER"),os.O_CREATE|os.O_WRONLY|os.O_APPEND,0600);if e!=nil{os.Exit(92)};f.WriteString("transport attempted\n");f.Close();os.Exit(91)}`)
	name := "git-remote-claimlocal"
	if runtime.GOOS == "windows" {
		name += ".exe"
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	c := exec.CommandContext(ctx, "go", "build", "-o", filepath.Join(dir, name), source)
	c.Env = append(os.Environ(), "GOPROXY=off", "GOTOOLCHAIN=local")
	if out, e := c.CombinedOutput(); e != nil {
		t.Fatalf("local transport build %v %s", e, out)
	}
	t.Setenv("PATH", dir+string(os.PathListSeparator)+os.Getenv("PATH"))
	t.Setenv("GIT_CONFIG_NOSYSTEM", "1")
	t.Setenv("GIT_CONFIG_GLOBAL", os.DevNull)
	t.Setenv("GIT_NO_LAZY_FETCH", "0")
	return dir
}

func claimLocalFetch(t *testing.T, trace string) bool {
	t.Helper()
	raw, e := os.ReadFile(trace)
	if e != nil {
		t.Fatal(e)
	}
	for _, line := range strings.Split(strings.TrimSpace(string(raw)), "\n") {
		if line == "" {
			continue
		}
		var event struct {
			Event, Name string
			Argv        []string
		}
		if e = json.Unmarshal([]byte(line), &event); e != nil {
			t.Fatal("malformed independent Git trace", e)
		}
		if event.Event == "cmd_name" && (event.Name == "fetch" || event.Name == "fetch-pack") {
			return true
		}
		if event.Event == "child_start" && len(event.Argv) > 1 && (filepath.Base(event.Argv[0]) == "git" || filepath.Base(event.Argv[0]) == "git.exe") {
			for i := 1; i < len(event.Argv); i++ {
				a := event.Argv[i]
				if a == "-c" || a == "-C" || a == "--git-dir" || a == "--work-tree" {
					i++
					continue
				}
				if strings.HasPrefix(a, "-") {
					continue
				}
				if a == "fetch" || a == "fetch-pack" {
					return true
				}
				break
			}
		}
	}
	return false
}

func claimLocalReset(t *testing.T, marker, trace string) {
	t.Helper()
	if e := os.Remove(marker); e != nil && !os.IsNotExist(e) {
		t.Fatal(e)
	}
	if e := os.WriteFile(trace, nil, 0600); e != nil {
		t.Fatal(e)
	}
}

func claimLocalNoFetch(t *testing.T, marker, trace string) {
	t.Helper()
	if _, e := os.Stat(marker); !os.IsNotExist(e) {
		t.Error("claim operation attempted local transport", e)
	}
	if claimLocalFetch(t, trace) {
		t.Error("claim operation attempted Git fetch even if transport was denied")
	}
}

type claimLocalFixture struct {
	Root, Common, Base, Slug, Object, Missing, Marker, Trace, Head, Old string
	Bytes                                                               []byte
	Previous, Other                                                     w.Claim
	Reject                                                              bool
	OptionalUpdate                                                      bool
}

func claimLocalSetup(t *testing.T, tools, format, kind, action string, linked bool) claimLocalFixture {
	t.Helper()
	root := claimRepo(t, format)
	old := claimBase(t, root)
	worktreePut(t, root, "retained", "second tracked bytes\n")
	claimGit(t, root, "add", ".")
	claimGit(t, root, "-c", "user.name=Local Claim", "-c", "user.email=claim@local.invalid", "commit", "-qm", "second")
	head := claimBase(t, root)
	if linked {
		root = claimPeer(t, root)
	}
	common := claimCommon(t, root)
	base := head
	if kind == "stored-base" {
		base = old
	}
	previous, e := w.CreateClaim(root, "current", "local-run", base, "exact opaque saved state\n")
	if e != nil {
		t.Fatal("available selected claim premise", e)
	}
	other, e := w.CreateClaim(root, "alpha", "other-run", head, "independent state")
	if e != nil {
		t.Fatal(e)
	}
	got, e := w.LoadClaim(root, "current")
	if e != nil || got != previous {
		t.Fatal("available read premise", got, e)
	}
	all, e := w.Claims(root)
	if e != nil || len(all) != 2 {
		t.Fatal("available inventory premise", all, e)
	}
	slug := "current"
	if action == "create" || action == "create-other" {
		slug = "fresh"
	}
	requested := head
	missing := ""
	reject := false
	switch kind {
	case "claim-body":
		missing = previous.Object
		reject = action == "load" || action == "list" || action == "update"
	case "stored-base":
		missing = old
		reject = action == "load" || action == "list" || action == "update"
	case "requested-commit":
		missing = old
		requested = old
		reject = true
	case "requested-ancestor":
		missing = old
		requested = "HEAD~1"
		reject = true
	case "requested-tag", "requested-tag-target":
		claimGit(t, root, "-c", "user.name=Local Claim", "-c", "user.email=claim@local.invalid", "tag", "-a", "selected", "-m", "local tag", old)
		requested = "selected"
		missing = old
		if kind == "requested-tag" {
			missing = claimGit(t, root, "rev-parse", "selected")
		}
		reject = true
	case "unrelated-claim":
		missing = other.Object
		reject = action == "list"
	case "unrelated-tree":
		missing = claimGit(t, root, "rev-parse", head+"^{tree}")
	case "unrelated-blob":
		missing = claimGit(t, root, "rev-parse", head+":retained")
	case "unrelated-ancestor":
		missing = old
	default:
		t.Fatal("unknown kind", kind)
	}
	// Dirty/staged/untracked/ignored bytes must survive all API operations.
	worktreePut(t, root, "retained", "staged user bytes\n")
	claimGit(t, root, "add", "retained")
	worktreePut(t, root, "retained", "unstaged user bytes\n")
	worktreePut(t, root, ".gitignore", "ignored\n")
	worktreePut(t, root, "ignored", "private ignored bytes")
	worktreePut(t, root, "untracked", "untracked bytes")
	for k, v := range map[string]string{"extensions.partialClone": "origin", "remote.origin.promisor": "true", "remote.origin.partialclonefilter": "blob:none", "remote.origin.url": "claimlocal::marker-only", "protocol.claimlocal.allow": "always"} {
		claimGit(t, root, "config", k, v)
	}
	object := filepath.Join(common, "objects", missing[:2], missing[2:])
	raw, e := os.ReadFile(object)
	if e != nil {
		t.Fatal("loose object premise", e)
	}
	if e = os.Remove(object); e != nil {
		t.Fatal(e)
	}
	t.Cleanup(func() { os.MkdirAll(filepath.Dir(object), 0755); os.WriteFile(object, raw, 0444) })
	dir := stateCanonicalDirectory(t, t.TempDir())
	marker := filepath.Join(dir, "marker")
	trace := filepath.Join(dir, "trace.json")
	t.Setenv("FW_CLAIM_LOCAL_MARKER", marker)
	t.Setenv("GIT_TRACE2_EVENT", trace)
	real, e := exec.LookPath("git")
	if e != nil {
		t.Fatal(e)
	}
	for _, denied := range []bool{false, true} {
		claimLocalReset(t, marker, trace)
		args := []string{}
		if denied {
			args = append(args, "-c", "protocol.claimlocal.allow=never")
		}
		args = append(args, "cat-file", "-t", missing)
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		cmd := exec.CommandContext(ctx, real, args...)
		cmd.Dir = root
		cmd.Env = append(os.Environ(), "GIT_NO_LAZY_FETCH=0")
		out, e := cmd.CombinedOutput()
		cancel()
		if e == nil {
			t.Fatal("missing-object positive control unexpectedly readable", string(out))
		}
		if !claimLocalFetch(t, trace) {
			t.Fatal("fetch positive control not observed", denied, string(out))
		}
		data, e := os.ReadFile(marker)
		if denied {
			if !os.IsNotExist(e) {
				t.Fatal("protocol denied transport still ran", e)
			}
		} else if e != nil || !strings.Contains(string(data), "transport attempted") {
			t.Fatal("transport marker positive control absent", e, string(out))
		}
	}
	claimLocalReset(t, marker, trace)
	_ = tools
	return claimLocalFixture{Root: root, Common: common, Base: requested, Slug: slug, Object: object, Missing: missing, Marker: marker, Trace: trace, Bytes: raw, Previous: previous, Other: other, Reject: reject, Head: head, Old: old, OptionalUpdate: action == "update" && (kind == "claim-body" || kind == "stored-base")}
}

func claimLocalCall(t *testing.T, f claimLocalFixture, action string, injected ...func() (w.Claim, []w.Claim, error)) {
	t.Helper()
	temp := stateCanonicalDirectory(t, t.TempDir())
	t.Setenv("TMPDIR", temp)
	t.Setenv("TMP", temp)
	t.Setenv("TEMP", temp)
	claimGit(t, f.Root, "--version") // Establish Git's own native startup cache.
	selected := func(path string) bool { return path == "objects" || strings.HasPrefix(path, "objects/") }
	beforeMeta := claimTree(t, f.Common, selected)
	beforeFiles := claimTree(t, f.Root, func(path string) bool { return path == ".git" })
	beforeTemp := claimTree(t, temp, func(string) bool { return false })
	beforeObjects := claimTree(t, filepath.Join(f.Common, "objects"), func(string) bool { return false })
	beforeRefs := claimGit(t, f.Root, "for-each-ref", "--format=%(refname) %(objectname)")
	previous := f.Previous
	env := os.Environ()
	cwd, e := os.Getwd()
	if e != nil {
		t.Fatal(e)
	}
	claimLocalNoFetch(t, f.Marker, f.Trace)
	claimLocalReset(t, f.Marker, f.Trace)
	var got w.Claim
	var inventory []w.Claim
	var err error
	if len(injected) > 0 {
		got, inventory, err = injected[0]()
	} else {
		switch action {
		case "create", "create-other":
			got, err = w.CreateClaim(f.Root, f.Slug, "new-run", f.Base, "new opaque state")
		case "load":
			got, err = w.LoadClaim(f.Root, "current")
		case "list":
			inventory, err = w.Claims(f.Root)
		case "update":
			got, err = w.UpdateClaim(f.Root, f.Previous, f.Base, "new opaque state")
		case "release":
			err = w.ReleaseClaim(f.Root, "current", f.Previous.Object)
		default:
			t.Fatal("unknown action", action)
		}
	}
	// Observe before fixture follow-up reads, independently of API error/success.
	claimLocalNoFetch(t, f.Marker, f.Trace)
	rejected := f.Reject
	if f.OptionalUpdate && err == nil {
		rejected = false
	}
	if rejected {
		claimError(t, got, err)
		if len(inventory) != 0 {
			t.Error("partial inventory after unavailable claim", inventory)
		}
	} else {
		if err != nil {
			t.Fatalf("independent/available operation rejected: %v", err)
		}
		switch action {
		case "create", "create-other":
			if got.Worktree != f.Root || got.Change != f.Slug || got.Run != "new-run" || got.Base != f.Head || got.State != "new opaque state" || got.Object == "" {
				t.Error("created claim mismatch", got)
			}
		case "load":
			if got != f.Previous {
				t.Error("loaded claim changed", got, f.Previous)
			}
		case "update":
			want := f.Previous
			want.Base = f.Head
			want.State = "new opaque state"
			want.Object = got.Object
			if got != want || got.Object == "" {
				t.Error("updated claim mismatch", got, want)
			}
		case "list":
			if !reflect.DeepEqual(inventory, []w.Claim{f.Other, f.Previous}) {
				t.Error("successful inventory mismatch", inventory)
			}
		}
	}
	if f.Previous != previous {
		t.Error("input Claim changed")
	}
	afterCwd, e := os.Getwd()
	if e != nil || afterCwd != cwd || !reflect.DeepEqual(env, os.Environ()) {
		t.Error("parent process changed", e)
	}
	if !reflect.DeepEqual(beforeFiles, claimTree(t, f.Root, func(path string) bool { return path == ".git" })) || !reflect.DeepEqual(beforeTemp, claimTree(t, temp, func(string) bool { return false })) {
		t.Error("working files/index link/temp changed")
	}
	afterObjects := claimTree(t, filepath.Join(f.Common, "objects"), func(string) bool { return false })
	if !rejected && action != "release" {
		claims := []w.Claim{got}
		if action == "list" {
			claims = inventory
		}
		for _, cl := range claims {
			if e := claimLocalStoredToken(f, cl, afterObjects); e != nil {
				t.Error("success token/payload mismatch", e)
			}
		}
	}

	for _, detail := range claimLocalObjectStorage(f, action, beforeObjects, afterObjects) {
		t.Error(detail)
	}
	if rejected || action == "load" || action == "list" {
		if !reflect.DeepEqual(beforeMeta, claimTree(t, f.Common, selected)) {
			t.Error("rejected/read call changed Git state")
		}
		if beforeRefs != claimGit(t, f.Root, "for-each-ref", "--format=%(refname) %(objectname)") {
			t.Error("rejected/read call changed refs")
		}
	} else {
		afterMeta := claimTree(t, f.Common, selected)
		path := "refs/fireweed/claims/" + f.Slug
		// Ancestor directories already exist because two independent claims are seeded.
		delete(beforeMeta, path)
		delete(afterMeta, path)
		if !reflect.DeepEqual(beforeMeta, afterMeta) {
			t.Error("successful call changed unrelated Git state")
		}
		refs := claimGit(t, f.Root, "for-each-ref", "--format=%(refname) %(objectname)")
		filter := func(raw string) string {
			var out []string
			for _, line := range strings.Split(raw, "\n") {
				if !strings.HasPrefix(line, "refs/fireweed/claims/"+f.Slug+" ") {
					out = append(out, line)
				}
			}
			return strings.Join(out, "\n")
		}
		if filter(beforeRefs) != filter(refs) {
			t.Error("successful call changed other refs")
		}
		selectedRef := claimGit(t, f.Root, "for-each-ref", "--format=%(objectname)", "refs/fireweed/claims/"+f.Slug)
		if action == "release" {
			if selectedRef != "" {
				t.Error("release retained selected ref")
			}
		} else if selectedRef != got.Object {
			t.Error("selected ref does not match returned token")
		}
	}
}

// Inspect new loose storage directly, so validating a malformed subject token
// cannot itself fetch or be misattributed to the public operation.
func claimLocalObjectStorage(f claimLocalFixture, action string, before, after map[string]string) []string {
	allowed := []w.Claim{}
	if action == "create" || action == "create-other" || action == "update" {
		run := "new-run"
		if action == "update" {
			run = f.Previous.Run
			allowed = append(allowed, f.Previous)
		}
		base := f.Head
		if f.Base != f.Head {
			base = f.Old
		}
		allowed = append(allowed, w.Claim{Worktree: f.Root, Change: f.Slug, Run: run, Base: base, State: "new opaque state"})
	}
	return claimLocalAllowedObjectStorage(f, allowed, before, after)
}

func claimLocalAllowedObjectStorage(f claimLocalFixture, allowed []w.Claim, before, after map[string]string) []string {
	errors := []string{}
	fail := func(v ...any) { errors = append(errors, fmt.Sprint(v...)) }
	for path, value := range before {
		if after[path] != value {
			fail("existing object storage changed", path)
		}
	}
	added := 0
	for path, value := range after {
		if _, ok := before[path]; ok {
			continue
		}
		parts := strings.Split(path, "/")
		if len(parts) == 1 && len(path) == 2 && strings.Trim(path, "0123456789abcdef") == "" && strings.HasPrefix(value, "d") {
			populated := false
			for child := range after {
				if strings.HasPrefix(child, path+"/") {
					populated = true
					break
				}
			}
			if !populated {
				fail("empty leaked object directory", path)
			}
			continue
		}
		if len(parts) != 2 || len(parts[0]) != 2 || len(parts[0]+parts[1]) != len(f.Head) || strings.Trim(parts[0]+parts[1], "0123456789abcdef") != "" || !strings.HasPrefix(value, "-") {
			fail("unexpected new object storage", path)
			continue
		}
		added++
		values, e := claimLocalBlob(f, parts[0]+parts[1], after)
		if e != nil {
			fail(e)
			continue
		}
		matched := false
		for _, candidate := range allowed {
			want := map[string]any{"worktree": candidate.Worktree, "change": candidate.Change, "run": candidate.Run, "base": candidate.Base, "state": candidate.State}
			if reflect.DeepEqual(values, want) {
				matched = true
				break
			}
		}
		if !matched {
			fail("new object does not belong to selected operation", path)
		}
	}
	if added > len(allowed) {
		fail("excess new selected-operation objects", added, len(allowed))
	}

	return errors
}

func claimLocalBlob(f claimLocalFixture, oid string, objects map[string]string) (map[string]any, error) {
	if len(oid) != len(f.Head) || len(oid) < 2 || strings.Trim(oid, "0123456789abcdef") != "" {
		return nil, fmt.Errorf("invalid selected object token %q", oid)
	}
	value, exists := objects[oid[:2]+"/"+oid[2:]]
	if !exists || !strings.HasPrefix(value, "-") {
		return nil, fmt.Errorf("selected token %s has no local regular object", oid)
	}
	_, compressed, ok := strings.Cut(value, "\n")
	if !ok {
		return nil, fmt.Errorf("selected object bytes missing %s", oid)
	}
	z, e := zlib.NewReader(bytes.NewReader([]byte(compressed)))
	if e != nil {
		return nil, fmt.Errorf("invalid object compression %s: %w", oid, e)
	}
	raw, e := io.ReadAll(io.LimitReader(z, 1<<20))
	closeErr := z.Close()
	if e != nil || closeErr != nil {
		return nil, fmt.Errorf("invalid object stream %s: %v %v", oid, e, closeErr)
	}
	header, body, ok := strings.Cut(string(raw), "\x00")
	fields := strings.Fields(header)
	if !ok || len(fields) != 2 || fields[0] != "blob" {
		return nil, fmt.Errorf("selected object is not a blob %s", oid)
	}
	n, e := strconv.Atoi(fields[1])
	if e != nil || n != len(body) {
		return nil, fmt.Errorf("invalid blob length %s", oid)
	}
	actual := ""
	if len(f.Head) == 40 {
		sum := sha1.Sum(raw)
		actual = fmt.Sprintf("%x", sum)
	} else {
		sum := sha256.Sum256(raw)
		actual = fmt.Sprintf("%x", sum)
	}
	if actual != oid {
		return nil, fmt.Errorf("selected object hash mismatch %s", oid)
	}
	var values map[string]any
	if e = json.Unmarshal([]byte(body), &values); e != nil || len(values) != 5 {
		return nil, fmt.Errorf("selected object is not exact claim JSON %s: %v", oid, e)
	}
	return values, nil
}
func claimLocalStoredToken(f claimLocalFixture, cl w.Claim, objects map[string]string) error {
	values, e := claimLocalBlob(f, cl.Object, objects)
	if e != nil {
		return e
	}
	want := map[string]any{"worktree": cl.Worktree, "change": cl.Change, "run": cl.Run, "base": cl.Base, "state": cl.State}
	if !reflect.DeepEqual(values, want) {
		return fmt.Errorf("stored token %s does not encode returned Claim", cl.Object)
	}
	return nil
}

func Test_FLOW011ClaimsNeverLazyFetchRequiredObjects(t *testing.T) {
	tools := claimLocalTransport(t)
	actions := map[string][]string{
		"claim-body":       {"load", "list", "update", "release", "create-other"},
		"stored-base":      {"load", "list", "update", "release", "create-other"},
		"requested-commit": {"create", "update"}, "requested-ancestor": {"create", "update"}, "requested-tag": {"create", "update"}, "requested-tag-target": {"create", "update"},
		"unrelated-claim": {"load", "list", "create", "update", "release"}, "unrelated-tree": {"create", "load", "list", "update", "release"}, "unrelated-blob": {"create", "load", "list", "update", "release"}, "unrelated-ancestor": {"create", "load", "list", "update", "release"},
	}
	for _, format := range []string{"sha1", "sha256"} {
		for _, linked := range []bool{false, true} {
			for _, kind := range []string{"claim-body", "stored-base", "requested-commit", "requested-ancestor", "requested-tag", "requested-tag-target", "unrelated-claim", "unrelated-tree", "unrelated-blob", "unrelated-ancestor"} {
				for _, action := range actions[kind] {
					t.Run(fmt.Sprintf("%s/linked_%t/%s/%s", format, linked, kind, action), func(t *testing.T) {
						f := claimLocalSetup(t, tools, format, kind, action, linked)
						claimLocalCall(t, f, action)
					})
				}
			}
		}
	}
}
