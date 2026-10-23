package evals

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"runtime"
	"sort"
	"strings"
	"testing"
	"time"

	g "github.com/lurkshark/fireweed/packages/gitstore"
)

type gitReadTools struct{ Git, Transport, Trap, Trace, RefTrace string }
type gitReadRun struct {
	Out, Err []byte
	Exit     int
}

func gitReadCommand(t *testing.T, tools gitReadTools, root string, input []byte, args ...string) gitReadRun {
	t.Helper()
	return gitReadWithEnv(t, tools, root, input, nil, args...)
}
func gitReadWithEnv(t *testing.T, tools gitReadTools, root string, input []byte, overrides []string, args ...string) gitReadRun {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, tools.Git, args...)
	cmd.Dir = root
	cmd.Env = append(append(os.Environ(), "GIT_OPTIONAL_LOCKS=0"), overrides...)
	cmd.Stdin = bytes.NewReader(input)
	var out, diagnostics bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &diagnostics
	err := cmd.Run()
	code := 0
	if err != nil {
		if e, ok := err.(*exec.ExitError); ok {
			code = e.ExitCode()
		} else {
			t.Fatalf("fixture Git launch %v: %v", args, err)
		}
	}
	if ctx.Err() != nil {
		t.Fatalf("fixture Git deadline %v: %v", args, ctx.Err())
	}
	return gitReadRun{out.Bytes(), diagnostics.Bytes(), code}
}
func gitReadOK(t *testing.T, tools gitReadTools, root string, args ...string) string {
	t.Helper()
	r := gitReadCommand(t, tools, root, nil, args...)
	if r.Exit != 0 {
		t.Fatalf("fixture Git %v exit%d: %s %s", args, r.Exit, r.Out, r.Err)
	}
	return strings.TrimSpace(string(r.Out))
}
func gitReadPrepare(t *testing.T) gitReadTools {
	t.Helper()
	real, e := exec.LookPath("git")
	if e != nil {
		t.Fatal(e)
	}
	tools := t.TempDir()
	marker := filepath.Join(t.TempDir(), "transport")
	trap := filepath.Join(t.TempDir(), "unexpected-command")
	trace := filepath.Join(t.TempDir(), "trace.jsonl")
	refs := filepath.Join(t.TempDir(), "refs.log")
	source := filepath.Join(t.TempDir(), "helper.go")
	program := `package main
import("os";"path/filepath";"strings")
func main(){name:=filepath.Base(os.Args[0]);key:="FIREWEED_GITREAD_TRAP";if strings.HasPrefix(name,"git-remote-"){key="FIREWEED_GITREAD_TRANSPORT"};p:=os.Getenv(key);if p==""{os.Exit(93)};f,e:=os.OpenFile(p,os.O_CREATE|os.O_WRONLY|os.O_APPEND,0600);if e!=nil{os.Exit(94)};f.WriteString(name+"\n");f.Close();os.Exit(1)}`
	if e := os.WriteFile(source, []byte(program), 0600); e != nil {
		t.Fatal(e)
	}
	suffix := ""
	if runtime.GOOS == "windows" {
		suffix = ".exe"
	}
	helper := filepath.Join(tools, "git-remote-fwgitread"+suffix)
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("fixture source location")
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	cmd := exec.CommandContext(ctx, "go", "build", "-o", helper, source)
	cmd.Dir = filepath.Clean(filepath.Join(filepath.Dir(file), "../../.."))
	cmd.Env = append(os.Environ(), "GOPROXY=off", "GOTOOLCHAIN=local")
	if output, e := cmd.CombinedOutput(); e != nil {
		t.Fatalf("compile local marker helper: %v %s", e, output)
	}
	executable, e := os.ReadFile(helper)
	if e != nil {
		t.Fatal(e)
	}
	for _, name := range []string{"fireweed", "fireweed-bootstrap", "codex", "claude", "curl", "wget", "ssh"} {
		if e := os.WriteFile(filepath.Join(tools, name+suffix), executable, 0755); e != nil {
			t.Fatal(e)
		}
	}
	t.Setenv("PATH", tools+string(os.PathListSeparator)+os.Getenv("PATH"))
	t.Setenv("FIREWEED_GITREAD_TRANSPORT", marker)
	t.Setenv("FIREWEED_GITREAD_TRAP", trap)
	t.Setenv("FIREWEED_BIN", filepath.Join(tools, "fireweed-bootstrap"+suffix))
	t.Setenv("GIT_TRACE2_EVENT", trace)
	t.Setenv("GIT_TRACE_REFS", refs)
	t.Setenv("GIT_NO_LAZY_FETCH", "1")
	t.Setenv("GIT_CONFIG_NOSYSTEM", "1")
	t.Setenv("GIT_CONFIG_GLOBAL", os.DevNull)
	t.Setenv("GIT_OPTIONAL_LOCKS", "1")
	return gitReadTools{real, marker, trap, trace, refs}
}
func gitReadReset(t *testing.T, tools gitReadTools) {
	t.Helper()
	for _, p := range []string{tools.Transport, tools.Trap, tools.Trace, tools.RefTrace} {
		if e := os.Remove(p); e != nil && !os.IsNotExist(e) {
			t.Fatal(e)
		}
	}
}
func gitReadOperation(argv []string) (string, []string) {
	if len(argv) == 0 {
		return "", nil
	}
	start := 0
	name := strings.TrimSuffix(filepath.Base(argv[0]), ".exe")
	if name == "git" {
		start = 1
	}
	for i := start; i < len(argv); i++ {
		a := argv[i]
		switch a {
		case "-c", "-C", "--git-dir", "--work-tree", "--namespace", "--config-env":
			i++
			continue
		}
		if strings.HasPrefix(a, "-") {
			continue
		}
		return a, argv[i+1:]
	}
	return "", nil
}
func gitReadNetworkEvents(t *testing.T, tools gitReadTools) []string {
	t.Helper()
	raw, e := os.ReadFile(tools.Trace)
	if os.IsNotExist(e) {
		return nil
	}
	if e != nil {
		t.Fatal(e)
	}
	found := []string{}
	for _, line := range bytes.Split(raw, []byte{'\n'}) {
		if len(line) == 0 {
			continue
		}
		var event struct {
			Event, Name string
			Argv        []string
		}
		if e := json.Unmarshal(line, &event); e != nil {
			t.Fatalf("fixture Trace2 JSON %q: %v", line, e)
		}
		op, tail := "", []string(nil)
		if event.Event == "start" {
			op, tail = gitReadOperation(event.Argv)
		} else if event.Event == "cmd_name" {
			op = event.Name
		} else if event.Event == "child_start" {
			op, tail = gitReadOperation(event.Argv)
		}
		switch op {
		case "fetch", "push", "pull", "clone", "ls-remote", "fetch-pack", "send-pack":
			found = append(found, string(line))
		case "branch":
			if gitReadBranchCreation(tail) {
				found = append(found, string(line))
			}
		case "checkout", "switch":
			for _, a := range tail {
				if a == "--" {
					break
				}
				if a == "-b" || a == "-B" || a == "-c" || a == "-C" || a == "--create" || strings.HasPrefix(a, "--create=") || a == "--force-create" || strings.HasPrefix(a, "--force-create=") {
					found = append(found, string(line))
				}
			}
		case "remote":
			if remoteOp, _ := gitReadOperation(tail); remoteOp == "update" {
				found = append(found, string(line))
			}
		}
	}
	refBytes, err := os.ReadFile(tools.RefTrace)
	if err != nil && !os.IsNotExist(err) {
		t.Fatal(err)
	}
	for _, line := range strings.Split(string(refBytes), "\n") {
		// Ref transactions expose stdin-based update-ref writes as well as branch.
		// Read-only raw-ref and iterator lines never have an old -> new pair.
		fields := strings.Fields(line)
		for i, field := range fields {
			if field == "->" && i >= 2 && i+1 < len(fields) && strings.HasPrefix(fields[i-2], "refs/heads/") && strings.Trim(fields[i-1], "0") == "" && strings.Trim(fields[i+1], "0") != "" {
				found = append(found, line)
			}
		}
	}
	return found
}
func gitReadBranchCreation(args []string) bool {
	listing := false
	operands := 0
	for i := 0; i < len(args); i++ {
		a := args[i]
		if a == "--format" || a == "--sort" {
			i++
			continue
		}
		if a == "--list" || a == "-a" || a == "--all" || a == "-r" || a == "--remotes" || a == "--show-current" || a == "--contains" || a == "--no-contains" || a == "--merged" || a == "--no-merged" || a == "--points-at" {
			listing = true
		}
		if !strings.HasPrefix(a, "-") {
			operands++
		}
	}
	return !listing && operands > 0
}
func gitReadNoTransport(t *testing.T, tools gitReadTools) {
	t.Helper()
	for _, p := range []string{tools.Transport, tools.Trap} {
		if raw, e := os.ReadFile(p); e == nil {
			t.Errorf("API attempted transport or unexpected command: %s", raw)
		} else if !os.IsNotExist(e) {
			t.Fatal(e)
		}
	}
	if events := gitReadNetworkEvents(t, tools); len(events) > 0 {
		t.Errorf("API attempted prohibited Git operation: %s", strings.Join(events, "\n"))
	}
}
func gitReadLiveTransport(t *testing.T, tools gitReadTools, root, missing string) {
	t.Helper()
	gitReadReset(t, tools)
	r := gitReadWithEnv(t, tools, root, nil, []string{"GIT_NO_LAZY_FETCH=0"}, "cat-file", "-t", missing)
	if r.Exit == 0 {
		t.Fatal("fixture missing-object control unexpectedly succeeded")
	}
	if _, e := os.Stat(tools.Transport); e != nil {
		t.Fatalf("fixture allowed transport marker is not live: %v, %s", e, r.Err)
	}
	if len(gitReadNetworkEvents(t, tools)) == 0 {
		t.Fatal("fixture allowed fetch Trace2 observer is not live")
	}
	gitReadReset(t, tools)
	r = gitReadWithEnv(t, tools, root, nil, []string{"GIT_NO_LAZY_FETCH=0"}, "-c", "protocol.fwgitread.allow=never", "cat-file", "-t", missing)
	if r.Exit == 0 {
		t.Fatal("fixture denied transport unexpectedly succeeded")
	}
	if _, e := os.Stat(tools.Transport); !os.IsNotExist(e) {
		t.Fatalf("fixture denied helper ran or cannot be inspected: %v", e)
	}
	if len(gitReadNetworkEvents(t, tools)) == 0 {
		t.Fatal("fixture denied fetch Trace2 observer is not live")
	}
	gitReadReset(t, tools)
}

type gitReadFixture struct{ Root, Common, GitDir, Old, Head, OldTree, OldNested, OldBlob, Tag, TagObject string }

func gitReadRepository(t *testing.T, tools gitReadTools, format string, linked bool, rename string) gitReadFixture {
	t.Helper()
	root := filepath.Join(t.TempDir(), "repo ' λ spaces")
	if e := os.MkdirAll(root, 0755); e != nil {
		t.Fatal(e)
	}
	root, e := filepath.EvalSymlinks(root)
	if e != nil {
		t.Fatal(e)
	}
	gitReadOK(t, tools, root, "init", "-q", "--initial-branch=main", "--object-format="+format)
	gitReadOK(t, tools, root, "config", "user.name", "Local Git Eval")
	gitReadOK(t, tools, root, "config", "user.email", "local@eval.invalid")
	write(t, root, "tracked.txt", "old content\n")
	write(t, root, "nested/file.txt", "old nested content\n")
	write(t, root, ".gitignore", "ignored-user\n")
	if rename != "" {
		var text strings.Builder
		for i := 0; i < 120; i++ {
			fmt.Fprintf(&text, "retained distinct line %03d for rename detection\n", i)
		}
		write(t, root, "rename old.txt", text.String())
	}
	gitReadOK(t, tools, root, "add", ".")
	gitReadOK(t, tools, root, "commit", "-qm", "old baseline")
	old := gitReadOK(t, tools, root, "rev-parse", "HEAD")
	tree := gitReadOK(t, tools, root, "rev-parse", "HEAD^{tree}")
	nested := gitReadOK(t, tools, root, "rev-parse", "HEAD:nested")
	blob := gitReadOK(t, tools, root, "rev-parse", "HEAD:tracked.txt")
	write(t, root, "tracked.txt", "new content\n")
	write(t, root, "nested/file.txt", "new nested content\n")
	if rename != "" {
		if e := os.Rename(filepath.Join(root, "rename old.txt"), filepath.Join(root, "rename new.txt")); e != nil {
			t.Fatal(e)
		}
		if rename == "fuzzy" {
			p := filepath.Join(root, "rename new.txt")
			raw, e := os.ReadFile(p)
			if e != nil {
				t.Fatal(e)
			}
			raw = append(raw, []byte("one additional changed line\n")...)
			if e = os.WriteFile(p, raw, 0644); e != nil {
				t.Fatal(e)
			}
		}
	}
	gitReadOK(t, tools, root, "add", ".")
	gitReadOK(t, tools, root, "commit", "-qm", "current baseline")
	head := gitReadOK(t, tools, root, "rev-parse", "HEAD")
	gitReadOK(t, tools, root, "tag", "-a", "selected", "-m", "older selected baseline", old)
	gitReadOK(t, tools, root, "tag", "-a", "selected-chain", "-m", "selected tag chain", "selected")
	tag := gitReadOK(t, tools, root, "rev-parse", "refs/tags/selected")
	opened, e := g.Open(root)
	if e != nil {
		t.Fatal(e)
	}
	common := opened.CommonDir
	if linked {
		peer := filepath.Join(t.TempDir(), "linked ' λ checkout")
		gitReadOK(t, tools, root, "worktree", "add", "--detach", peer, head)
		root, e = filepath.EvalSymlinks(peer)
		if e != nil {
			t.Fatal(e)
		}
	}
	write(t, root, "ignored-user", "preserved ignored user bytes\n")
	write(t, common, "fireweed-user-retained.txt", "preserved common Git user bytes\n")
	gitReadOK(t, tools, root, "config", "--unset", "user.name")
	gitReadOK(t, tools, root, "config", "--unset", "user.email")
	gitReadReset(t, tools)
	opened, e = g.Open(root)
	if e != nil {
		t.Fatal(e)
	}
	gitReadReset(t, tools)
	return gitReadFixture{root, common, opened.GitDir, old, head, tree, nested, blob, "selected", tag}
}
func gitReadPromisor(t *testing.T, tools gitReadTools, f gitReadFixture, missing string) {
	t.Helper()
	for _, kv := range [][2]string{{"extensions.partialClone", "origin"}, {"remote.origin.promisor", "true"}, {"remote.origin.partialclonefilter", "blob:none"}, {"remote.origin.url", "fwgitread::marker-only"}, {"protocol.fwgitread.allow", "always"}} {
		gitReadOK(t, tools, f.Root, "config", kv[0], kv[1])
	}
	p := filepath.Join(f.Common, "objects", missing[:2], missing[2:])
	raw, e := os.ReadFile(p)
	if e != nil {
		t.Fatal(e)
	}
	info, e := os.Stat(p)
	if e != nil {
		t.Fatal(e)
	}
	if e = os.Remove(p); e != nil {
		t.Fatal(e)
	}
	t.Cleanup(func() {
		if e := os.MkdirAll(filepath.Dir(p), 0755); e != nil {
			t.Error(e)
			return
		}
		if e := os.WriteFile(p, raw, info.Mode().Perm()); e != nil {
			t.Error(e)
		}
	})
	gitReadLiveTransport(t, tools, f.Root, missing)
}

type gitReadFile struct {
	Mode       fs.FileMode
	Data, Link string
}
type gitReadTree map[string]gitReadFile

func gitReadTreeImage(t *testing.T, root string, excludeGit bool) gitReadTree {
	t.Helper()
	result := gitReadTree{}
	e := filepath.WalkDir(root, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, e := filepath.Rel(root, path)
		if e != nil {
			return e
		}
		rel = filepath.ToSlash(rel)
		if excludeGit && rel == ".git" && entry.IsDir() {
			return filepath.SkipDir
		}
		info, e := entry.Info()
		if e != nil {
			return e
		}
		v := gitReadFile{Mode: info.Mode()}
		if info.Mode()&os.ModeSymlink != 0 {
			v.Link, e = os.Readlink(path)
		} else if info.Mode().IsRegular() {
			var b []byte
			b, e = os.ReadFile(path)
			v.Data = string(b)
		}
		if e != nil {
			return e
		}
		result[rel] = v
		return nil
	})
	if e != nil {
		t.Fatal(e)
	}
	return result
}

type gitReadState struct {
	Root, Common, Temporary gitReadTree
	GitPaths                map[string]bool
	Cwd                     string
	Env                     []string
}

func gitReadStateImage(t *testing.T, f gitReadFixture, temporary string) gitReadState {
	t.Helper()
	cwd, e := os.Getwd()
	if e != nil {
		t.Fatal(e)
	}
	return gitReadState{Root: gitReadTreeImage(t, f.Root, true), Common: gitReadGitImage(t, f), Temporary: gitReadTreeImage(t, temporary, false), GitPaths: gitReadGitPaths(t, f), Cwd: cwd, Env: os.Environ()}
}

// Only promised Git state is immutable: unrelated object-store layout and caches
// are not part of the contract. Include the selected linked index and shared index.
func gitReadGitImage(t *testing.T, f gitReadFixture) gitReadTree {
	result := gitReadTree{}
	for prefix, root := range map[string]string{"common": f.Common, "selected": f.GitDir} {
		for rel, entry := range gitReadTreeImage(t, root, false) {
			if rel == "config" || rel == "config.worktree" || rel == "HEAD" || rel == "index" || rel == "packed-refs" || rel == "refs" || strings.HasPrefix(rel, "refs/") || strings.HasPrefix(rel, "sharedindex.") {
				result[prefix+"/"+rel] = entry
			}
		}
	}
	return result
}

// Observe repository-local scratch without making unrelated object/cache bytes
// immutable. New paths outside native object/cache storage must be cleaned.
func gitReadGitPaths(t *testing.T, f gitReadFixture) map[string]bool {
	paths := map[string]bool{}
	for prefix, root := range map[string]string{"common": f.Common, "selected": f.GitDir} {
		for rel := range gitReadTreeImage(t, root, false) {
			cache := false
			for _, dir := range []string{"objects", "rr-cache", "fsmonitor--daemon"} {
				if rel == dir || strings.HasPrefix(rel, dir+"/") {
					cache = true
				}
			}
			if gitReadSharedIndex(rel) {
				cache = true
			}
			if strings.HasPrefix(filepath.Base(rel), "tmp_") || strings.HasSuffix(rel, ".lock") {
				cache = false
			}
			if !cache {
				paths[prefix+"/"+rel] = true
			}
		}
	}
	return paths
}
func gitReadSharedIndex(path string) bool {
	name := filepath.Base(path)
	if !strings.HasPrefix(name, "sharedindex.") {
		return false
	}
	oid := strings.TrimPrefix(name, "sharedindex.")
	if len(oid) != 40 && len(oid) != 64 {
		return false
	}
	return strings.Trim(oid, "0123456789abcdef") == ""
}
func gitReadNewScratch(before, after gitReadState) []string {
	paths := []string{}
	for path := range after.GitPaths {
		if !before.GitPaths[path] {
			paths = append(paths, path)
		}
	}
	sort.Strings(paths)
	return paths
}
func gitReadTemporary(t *testing.T) string {
	t.Helper()
	p, e := filepath.EvalSymlinks(t.TempDir())
	if e != nil {
		t.Fatal(e)
	}
	write(t, p, "user/retained.txt", "temporary parent user bytes\n")
	for _, key := range []string{"TMPDIR", "TMP", "TEMP"} {
		t.Setenv(key, p)
	}
	return p
}
func gitReadSameState(t *testing.T, f gitReadFixture, temp string, before gitReadState) {
	t.Helper()
	after := gitReadStateImage(t, f, temp)
	if !reflect.DeepEqual(before.Root, after.Root) {
		t.Error("API changed working/ignored files or modes")
	}
	// Existing shared-index bytes belong to the original index. New shared-index
	// cache files need not imply a mutation of that original index.
	for path := range after.Common {
		if gitReadSharedIndex(path) {
			if _, existed := before.Common[path]; !existed {
				delete(after.Common, path)
			}
		}
	}
	if !reflect.DeepEqual(before.Common, after.Common) {
		t.Error("API changed HEAD/refs/original index bytes or repository configuration")
	}
	if scratch := gitReadNewScratch(before, after); len(scratch) != 0 {
		t.Errorf("API left repository-local owned scratch: %v", scratch)
	}
	if !reflect.DeepEqual(before.Temporary, after.Temporary) {
		t.Error("API changed temporary parent or left owned storage")
	}
	if before.Cwd != after.Cwd || !reflect.DeepEqual(before.Env, after.Env) {
		t.Error("API changed parent cwd/environment")
	}
}
func gitReadNativeEdits(t *testing.T, tools gitReadTools, root, base string) []g.Edit {
	t.Helper()
	r := gitReadCommand(t, tools, root, nil, "diff", "--name-status", "-z", "--find-renames", "--no-ext-diff", "--no-textconv", base, "--")
	if r.Exit != 0 {
		t.Fatalf("available native diff control failed: %s %s", r.Out, r.Err)
	}
	tokens := bytes.Split(r.Out, []byte{0})
	result := []g.Edit{}
	for i := 0; i < len(tokens) && len(tokens[i]) > 0; {
		status := string(tokens[i])
		i++
		if i >= len(tokens) || len(tokens[i]) == 0 {
			t.Fatal("native diff control missing path")
		}
		first := string(tokens[i])
		i++
		entry := g.Edit{Status: status[:1], Path: first}
		if entry.Status == "R" || entry.Status == "C" {
			if i >= len(tokens) || len(tokens[i]) == 0 {
				t.Fatal("native rename control missing destination")
			}
			entry.OldPath = first
			entry.Path = string(tokens[i])
			i++
		}
		result = append(result, entry)
	}
	r = gitReadCommand(t, tools, root, nil, "ls-files", "--others", "--exclude-standard", "-z")
	if r.Exit != 0 {
		t.Fatalf("native untracked control failed: %s", r.Err)
	}
	for _, p := range bytes.Split(r.Out, []byte{0}) {
		if len(p) > 0 {
			result = append(result, g.Edit{Status: "A", Path: string(p)})
		}
	}
	sort.Slice(result, func(i, j int) bool {
		if result[i].Path == result[j].Path {
			if result[i].OldPath == result[j].OldPath {
				return result[i].Status < result[j].Status
			}
			return result[i].OldPath < result[j].OldPath
		}
		return result[i].Path < result[j].Path
	})
	return result
}
