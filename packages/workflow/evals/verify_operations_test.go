package evals

import (
	"context"
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"testing"
	"time"
)

// Trace actual native Git entry, not substring matches in arbitrary operands.
// Temporary metadata setup and read-only branch/worktree queries remain legal.
func verifyForbiddenOperation(argv []string) bool {
	args := verifyGitArgs(argv)
	if len(args) == 0 {
		return false
	}
	verb, args := args[0], args[1:]
	values := map[string]bool{}
	switch verb {
	case "commit":
		for _, v := range []string{"-m", "--message", "-F", "--file", "-C", "--reuse-message", "-c", "--reedit-message", "--author", "--date", "--fixup", "--squash", "--cleanup", "-t", "--template", "--pathspec-from-file", "--trailer"} {
			values[v] = true
		}
	case "commit-tree":
		for _, v := range []string{"-m", "-F", "-p"} {
			values[v] = true
		}
	case "push", "send-pack":
		for _, v := range []string{"-o", "--push-option", "--receive-pack", "--exec", "--repo"} {
			values[v] = true
		}
	case "fetch":
		for _, v := range []string{"--upload-pack", "--depth", "--deepen", "--shallow-since", "--shallow-exclude", "--negotiation-tip", "--filter", "--server-option", "-o", "-j", "--jobs"} {
			values[v] = true
		}
	case "fetch-pack":
		for _, v := range []string{"--upload-pack", "--depth"} {
			values[v] = true
		}
	case "branch":
		for _, v := range []string{"--format", "--sort", "--contains", "--no-contains", "--merged", "--no-merged", "--points-at", "-u", "--set-upstream-to"} {
			values[v] = true
		}
	}
	tokens := []string{}
	options := true
	for i := 0; i < len(args); i++ {
		a := args[i]
		if options && a == "--" {
			options = false
			tokens = append(tokens, "--")
			continue
		}
		if options && (a == "-h" || a == "--help") {
			return false
		}
		tokens = append(tokens, a)
		if options && strings.HasPrefix(a, "-") && !strings.HasPrefix(a, "--") && len(a) > 2 {
			for j := 1; j < len(a); j++ {
				if values["-"+a[j:j+1]] {
					if j == len(a)-1 && i+1 < len(args) {
						i++
					}
					break
				}
			}
			continue
		}
		if options && values[a] && i+1 < len(args) {
			// contains/merged have optional revisions; a following flag is another option.
			optional := a == "--contains" || a == "--no-contains" || a == "--merged" || a == "--no-merged"
			if !optional || !strings.HasPrefix(args[i+1], "-") {
				i++
			}
		}
	}
	switch verb {
	case "commit", "commit-tree", "fetch", "fetch-pack", "push", "send-pack":
		return true
	case "worktree":
		return len(tokens) > 0 && tokens[0] == "add"
	case "branch":
		readOnly := false
		operand := false
		branchOptions := true
		for _, a := range tokens {
			if a == "--" && branchOptions {
				branchOptions = false
				continue
			}
			if !branchOptions {
				operand = true
				continue
			}
			switch a {
			case "-d", "-D", "-m", "-M", "-c", "-C", "--delete", "--move", "--copy", "--edit-description", "--unset-upstream", "-u", "--set-upstream-to":
				return true
			case "--list", "-l", "--show-current", "--contains", "--no-contains", "--merged", "--no-merged", "--points-at", "--format", "--sort":
				readOnly = true
			}
			if strings.HasPrefix(a, "-") && !strings.HasPrefix(a, "--") && len(a) > 2 && strings.ContainsAny(a[1:], "dDmMcCu") {
				return true
			}
			if strings.HasPrefix(a, "--set-upstream-to=") {
				return true
			}
			for _, prefix := range []string{"--contains=", "--no-contains=", "--merged=", "--no-merged=", "--points-at=", "--format=", "--sort="} {
				if strings.HasPrefix(a, prefix) {
					readOnly = true
				}
			}
			if !strings.HasPrefix(a, "-") {
				operand = true
			}
		}
		return operand && !readOnly
	}
	return false
}
func verifyGitArgs(argv []string) []string {
	args := append([]string{}, argv...)
	if len(args) > 0 && (filepath.Base(args[0]) == "git" || filepath.Base(args[0]) == "git.exe") {
		args = args[1:]
	}
	for len(args) > 0 && strings.HasPrefix(args[0], "-") {
		a := args[0]
		args = args[1:]
		if a == "--" {
			break
		}
		switch a {
		case "-C", "-c", "--git-dir", "--work-tree", "--namespace", "--super-prefix", "--config-env":
			if len(args) > 0 {
				args = args[1:]
			}
		}
	}
	return args
}
func verifyGitVerb(argv []string) string {
	args := verifyGitArgs(argv)
	if len(args) == 0 {
		return ""
	}
	return args[0]
}
func verifyGitLocation(argv []string, root, inherited string) string {
	args := append([]string{}, argv...)
	if len(args) > 0 && (filepath.Base(args[0]) == "git" || filepath.Base(args[0]) == "git.exe") {
		args = args[1:]
	}
	cwd, dir := root, inherited
	for len(args) > 0 && strings.HasPrefix(args[0], "-") {
		a := args[0]
		args = args[1:]
		if a == "--" {
			break
		}
		switch a {
		case "-C":
			if len(args) > 0 {
				if filepath.IsAbs(args[0]) {
					cwd = args[0]
				} else {
					cwd = filepath.Join(cwd, args[0])
				}
				args = args[1:]
			}
		case "--git-dir":
			if len(args) > 0 {
				dir = args[0]
				args = args[1:]
			}
		case "-c", "--work-tree", "--namespace", "--super-prefix", "--config-env":
			if len(args) > 0 {
				args = args[1:]
			}
		default:
			if strings.HasPrefix(a, "--git-dir=") {
				dir = strings.TrimPrefix(a, "--git-dir=")
			}
		}
	}
	if dir == "" {
		return filepath.Join(cwd, ".git")
	}
	if filepath.IsAbs(dir) {
		return filepath.Clean(dir)
	}
	return filepath.Join(cwd, dir)
}

// The recorder execs the real Git in place: native Trace2 sees the same PID,
// arguments, environment and streams. It supplies cwd and canonical metadata at
// invocation because native Trace2 omits cwd for bare repositories.
const verifyGitRecorderSource = `package main
import("encoding/json";"os";"path/filepath";"strings";"syscall")
func canonical(p string)string{if c,e:=filepath.EvalSymlinks(p);e==nil{return c};return filepath.Clean(p)}
func main(){cwd,e:=os.Getwd();if e!=nil{os.Exit(121)};cwd=canonical(cwd);dir:=os.Getenv("GIT_DIR");args:=os.Args[1:];for len(args)>0&&strings.HasPrefix(args[0],"-"){a:=args[0];args=args[1:];if a=="--"{break};switch a{case "-C":if len(args)>0{if filepath.IsAbs(args[0]){cwd=args[0]}else{cwd=filepath.Join(cwd,args[0])};args=args[1:]};case "--git-dir":if len(args)>0{dir=args[0];args=args[1:]};case "-c","--work-tree","--namespace","--super-prefix","--config-env":if len(args)>0{args=args[1:]};default:if strings.HasPrefix(a,"--git-dir="){dir=strings.TrimPrefix(a,"--git-dir=")}}};if dir==""{for at:=cwd;;at=filepath.Dir(at){candidate:=filepath.Join(at,".git");if _,e:=os.Lstat(candidate);e==nil{dir=candidate;if raw,e:=os.ReadFile(dir);e==nil&&strings.HasPrefix(string(raw),"gitdir: "){dir=strings.TrimSpace(strings.TrimPrefix(string(raw),"gitdir: "));if !filepath.IsAbs(dir){dir=filepath.Join(at,dir)}};break};if _,e:=os.Stat(filepath.Join(at,"HEAD"));e==nil{if _,e:=os.Stat(filepath.Join(at,"objects"));e==nil{dir=at;break}};if filepath.Dir(at)==at{dir=filepath.Join(cwd,".git");break}}}else if !filepath.IsAbs(dir){dir=filepath.Join(cwd,dir)};dir=canonical(dir);if common:=os.Getenv("GIT_COMMON_DIR");common!=""{if !filepath.IsAbs(common){common=filepath.Join(cwd,common)};dir=canonical(common)}else if raw,e:=os.ReadFile(filepath.Join(dir,"commondir"));e==nil{common:=strings.TrimSpace(string(raw));if !filepath.IsAbs(common){common=filepath.Join(dir,common)};dir=canonical(common)};rec,_:=json.Marshal(map[string]any{"pid":os.Getpid(),"cwd":canonical(cwd),"dir":canonical(dir)});f,e:=os.OpenFile(os.Getenv("FW_VERIFY_GIT_LOCATIONS"),os.O_CREATE|os.O_APPEND|os.O_WRONLY,0600);if e!=nil{os.Exit(122)};if _,e=f.Write(append(rec,'\n'));e!=nil{os.Exit(123)};f.Close();real:=os.Getenv("FW_VERIFY_REAL_GIT");if syscall.Exec(real,append([]string{real},os.Args[1:]...),os.Environ())!=nil{os.Exit(124)}}
`

func verifyGitRecorder(t *testing.T, trace string) {
	t.Helper()
	if runtime.GOOS == "windows" {
		t.Skip("native exec recorder requires POSIX")
	}
	real, e := exec.LookPath("git")
	if e != nil {
		t.Fatal(e)
	}
	dir := t.TempDir()
	src := filepath.Join(dir, "recorder.go")
	bin := filepath.Join(dir, "git")
	if e = os.WriteFile(src, []byte(verifyGitRecorderSource), 0600); e != nil {
		t.Fatal(e)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	cmd := exec.CommandContext(ctx, "go", "build", "-o", bin, src)
	cmd.Env = append(os.Environ(), "GOPROXY=off", "GOTOOLCHAIN=local")
	if out, e := cmd.CombinedOutput(); e != nil {
		t.Fatalf("Git recorder build: %v %s", e, out)
	}
	t.Setenv("FW_VERIFY_REAL_GIT", real)
	t.Setenv("FW_VERIFY_GIT_LOCATIONS", trace+".locations")
	t.Setenv("PATH", dir+string(os.PathListSeparator)+os.Getenv("PATH"))
}
func verifyCanonicalGitDir(root string) string {
	dir := filepath.Join(root, ".git")
	if raw, e := os.ReadFile(dir); e == nil && strings.HasPrefix(string(raw), "gitdir: ") {
		dir = strings.TrimSpace(strings.TrimPrefix(string(raw), "gitdir: "))
		if !filepath.IsAbs(dir) {
			dir = filepath.Join(root, dir)
		}
	}
	if p, e := filepath.EvalSymlinks(dir); e == nil {
		return p
	}
	return filepath.Clean(dir)
}
func verifyForbiddenGitTrace(t *testing.T, path string, roots ...string) []string {
	t.Helper()
	raw, e := os.ReadFile(path)
	if os.IsNotExist(e) {
		return []string{}
	}
	if e != nil {
		t.Fatal(e)
	}
	type event struct {
		Event, Sid, Param, Value string
		Argv                     []string
	}
	locations := map[int]string{}
	if data, e := os.ReadFile(path + ".locations"); e == nil {
		for _, line := range strings.Split(strings.TrimSpace(string(data)), "\n") {
			if line == "" {
				continue
			}
			var loc struct {
				Pid int
				Dir string
			}
			if e := json.Unmarshal([]byte(line), &loc); e != nil {
				t.Fatal("malformed Git location record", e)
			}
			locations[loc.Pid] = loc.Dir
		}
	} else if !os.IsNotExist(e) {
		t.Fatal(e)
	}
	events := []event{}
	dirs := map[string]string{}
	for _, line := range strings.Split(strings.TrimSpace(string(raw)), "\n") {
		if line == "" {
			continue
		}
		var x event
		if e := json.Unmarshal([]byte(line), &x); e != nil {
			t.Fatal("malformed actual Git trace", e)
		}
		events = append(events, x)
		if x.Event == "def_param" && x.Param == "GIT_DIR" {
			dirs[x.Sid] = x.Value
		}
	}
	root, owned := "", ""
	if len(roots) > 0 {
		root = roots[0]
	}
	if len(roots) > 1 {
		owned = roots[1]
	}
	bad := []string{}
	for _, x := range events {
		if (x.Event != "start" && x.Event != "alias") || !verifyForbiddenOperation(x.Argv) {
			continue
		}
		// Metadata views may materialize copies of existing branch refs in owned
		// temporary storage. This never permits changes to the original Git dirs.
		if verifyGitVerb(x.Argv) == "branch" {
			dir := verifyGitLocation(x.Argv, root, dirs[x.Sid])
			sid := x.Sid
			if i := strings.LastIndex(sid, "-P"); i >= 0 {
				if pid, e := strconv.ParseUint(sid[i+2:], 16, 32); e == nil && locations[int(pid)] != "" {
					dir = locations[int(pid)]
				}
			}
			// Resolve aliases before deciding whether storage belongs to the observer.
			if p, e := filepath.EvalSymlinks(dir); e == nil {
				dir = p
			}
			if p, e := filepath.EvalSymlinks(owned); e == nil {
				owned = p
			}
			if owned != "" {
				rel, e := filepath.Rel(owned, dir)
				if e == nil && rel != ".." && !strings.HasPrefix(rel, ".."+string(filepath.Separator)) && !filepath.IsAbs(rel) {
					continue
				}
			}
			// A cleaned temporary view can also live under common Git scratch
			// storage rather than TMPDIR. Existing caller repositories remain
			// observable; a view whose HEAD/storage was removed is not one.
			original := verifyCanonicalGitDir(root)
			canonical := dir
			if p, e := filepath.EvalSymlinks(dir); e == nil {
				canonical = p
			}
			if p, e := filepath.EvalSymlinks(original); e == nil {
				original = p
			}
			if filepath.Clean(canonical) != filepath.Clean(original) {
				if _, e := os.Stat(filepath.Join(dir, "HEAD")); os.IsNotExist(e) {
					continue
				}
			}
		}
		bad = append(bad, strings.Join(x.Argv, " "))
	}
	return bad
}
func verifyNoForbiddenGit(t *testing.T, path string, roots ...string) {
	t.Helper()
	if bad := verifyForbiddenGitTrace(t, path, roots...); len(bad) > 0 {
		t.Error("Verify executed forbidden native Git operations", bad)
	}
}
func Test_FLOW047VerifyNativeOperationObserverControls(t *testing.T) {
	root := claimRepo(t, "sha1")
	base := claimBase(t, root)
	trace := filepath.Join(t.TempDir(), "trace")
	t.Setenv("GIT_TRACE2_EVENT", trace)
	t.Setenv("GIT_TRACE2_ENV_VARS", "GIT_DIR")
	verifyGitRecorder(t, trace)
	clear := func() {
		t.Helper()
		if e := os.WriteFile(trace+".locations", nil, 0600); e != nil {
			t.Fatal(e)
		}
		if e := os.WriteFile(trace, nil, 0600); e != nil {
			t.Fatal(e)
		}
	}
	expect := func(forbidden bool) {
		t.Helper()
		bad := verifyForbiddenGitTrace(t, trace, root, filepath.Join(filepath.Dir(trace), "owned"))
		if (len(bad) > 0) != forbidden {
			t.Fatalf("native operation observer not live: forbidden=%v actual=%v", forbidden, bad)
		}
	}
	for _, args := range [][]string{{"--version"}, {"-c", "core.fsmonitor=", "rev-parse", "--shared-index-path"}, {"branch", "--list"}, {"branch", "--show-current"}, {"branch", "--format", "-m", "--list"}, {"branch", "--list", "--", "-m"}, {"worktree", "list", "--porcelain"}, {"config", "--get", "core.repositoryformatversion"}} {
		clear()
		claimGit(t, root, args...)
		expect(false)
	}
	nested := filepath.Join(root, "nested", "deeper")
	if e := os.MkdirAll(nested, 0755); e != nil {
		t.Fatal(e)
	}
	clear()
	claimGit(t, nested, "branch", "nested-caller-control", base)
	expect(true)
	claimGit(t, root, "branch", "-D", "nested-caller-control")
	// Existing selected commit can populate an owned metadata view without mutating
	// the caller's refs or index. Such setup is not a new caller branch.
	owned := filepath.Join(filepath.Dir(trace), "owned", "metadata")
	clear()
	claimGit(t, root, "init", "--bare", owned)
	claimGit(t, root, "--git-dir", owned, "config", "core.bare", "true")
	worktreePut(t, owned, "objects/info/alternates", filepath.Join(claimCommon(t, root), "objects")+"\n")
	claimGit(t, root, "--git-dir", owned, "branch", "copied", base)
	expect(false)
	clear()
	claimGit(t, owned, "branch", "copied-via-cwd", base)
	expect(false)
	alias := filepath.Join(filepath.Dir(owned), "original-alias")
	if e := os.Symlink(claimCommon(t, root), alias); e != nil {
		t.Fatal(e)
	}
	clear()
	claimGit(t, root, "--git-dir", alias, "branch", "symlink-caller-control", base)
	claimGit(t, root, "branch", "-D", "symlink-caller-control")
	expect(true)
	clear()
	claimGit(t, root, "-c", "user.name=Observer Fixture", "-c", "user.email=observer@example.invalid", "commit", "--allow-empty", "-qm", "-h")
	expect(true)
	claimGit(t, root, "update-ref", "HEAD", base)
	clear()
	claimGit(t, root, "branch", "--create-reflog", "reflog-caller-control", base)
	expect(true)
	claimGit(t, root, "branch", "-D", "reflog-caller-control")
	tree := claimGit(t, root, "rev-parse", base+"^{tree}")
	clear()
	claimGit(t, root, "--git-dir", owned, "-c", "user.name=Observer Fixture", "-c", "user.email=observer@example.invalid", "commit-tree", tree, "-m", "branch")
	expect(true)
	t.Run("owned-env-metadata", func(t *testing.T) {
		clear()
		t.Setenv("GIT_DIR", owned)
		claimGit(t, root, "branch", "copied-via-env", base)
		expect(false)
	})
	clear()
	claimGit(t, root, "branch", "transient-observer", base)
	claimGit(t, root, "branch", "-D", "transient-observer")
	expect(true)
	clear()
	claimGit(t, root, "-c", "user.name=Observer Fixture", "-c", "user.email=observer@example.invalid", "commit", "--allow-empty", "-qm", "local operation control")
	expect(true)
	claimGit(t, root, "update-ref", "HEAD", base)
	clear()
	claimGit(t, root, "push", ".", "HEAD:refs/heads/transient-push-observer")
	expect(true)
	claimGit(t, root, "update-ref", "-d", "refs/heads/transient-push-observer")
	clear()
	claimGit(t, root, "send-pack", ".", "HEAD:refs/heads/transient-send-pack-observer")
	expect(true)
	claimGit(t, root, "update-ref", "-d", "refs/heads/transient-send-pack-observer")
	path := filepath.Join(t.TempDir(), "detached")
	clear()
	claimGit(t, root, "worktree", "add", "--detach", path, base)
	claimGit(t, root, "worktree", "remove", path)
	expect(true)
	// Even a denied local transport command is forbidden; no server is contacted.
	clear()
	_, _ = claimCommand(t, root, "", "-c", "protocol.file.allow=never", "fetch", ".")
	expect(true)
	for _, args := range [][]string{{"git", "-C", root, "-c", "note.value=push branch worktree add commit", "rev-parse", "HEAD"}, {"git", "branch", "--list", "push*"}, {"git", "worktree", "list"}, {"git", "--git-dir=" + owned, "update-ref", "refs/heads/copied", base}} {
		if verifyForbiddenOperation(args) {
			t.Error("read/owned metadata operands falsely classified", args)
		}
	}
}
