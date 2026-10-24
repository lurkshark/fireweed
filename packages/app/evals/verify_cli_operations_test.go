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
func cliVerifyForbiddenOperation(argv []string) bool {
	args := cliVerifyGitArgs(argv)
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
	if verb == "add" || verb == "rm" || verb == "mv" {
		for _, a := range tokens {
			if a == "--" {
				break
			}
			if a == "--dry-run" || a == "-n" {
				return false
			}
		}
	}
	if verb == "update-index" && len(tokens) == 1 && tokens[0] == "--test-untracked-cache" {
		return false
	}
	switch verb {
	case "commit", "commit-tree", "fetch", "fetch-pack", "push", "send-pack", "add", "reset", "update-index", "read-tree", "rm", "mv", "checkout", "switch", "restore", "update-ref":
		return true
	case "checkout-index":
		for _, a := range tokens {
			if a == "-u" || a == "--index" {
				return true
			}
		}
		return false
	case "symbolic-ref":
		operands := 0
		for _, a := range tokens {
			if a == "--delete" || a == "-d" {
				return true
			}
			if !strings.HasPrefix(a, "-") {
				operands++
			}
		}
		return operands > 1
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
func cliVerifyGitArgs(argv []string) []string {
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
func cliVerifyGitVerb(argv []string) string {
	args := cliVerifyGitArgs(argv)
	if len(args) == 0 {
		return ""
	}
	return args[0]
}
func cliVerifyGitLocation(argv []string, root, inherited string) string {
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
const cliVerifyGitRecorderSource = `package main
import("encoding/json";"os";"path/filepath";"strings";"syscall")
func canonical(p string)string{if c,e:=filepath.EvalSymlinks(p);e==nil{return c};return filepath.Clean(p)}
func main(){cwd,e:=os.Getwd();if e!=nil{os.Exit(121)};cwd=canonical(cwd);dir:=os.Getenv("GIT_DIR");args:=os.Args[1:];for len(args)>0&&strings.HasPrefix(args[0],"-"){a:=args[0];args=args[1:];if a=="--"{break};switch a{case "-C":if len(args)>0{if filepath.IsAbs(args[0]){cwd=args[0]}else{cwd=filepath.Join(cwd,args[0])};args=args[1:]};case "--git-dir":if len(args)>0{dir=args[0];args=args[1:]};case "-c","--work-tree","--namespace","--super-prefix","--config-env":if len(args)>0{args=args[1:]};default:if strings.HasPrefix(a,"--git-dir="){dir=strings.TrimPrefix(a,"--git-dir=")}}};if dir==""{for at:=cwd;;at=filepath.Dir(at){candidate:=filepath.Join(at,".git");if _,e:=os.Lstat(candidate);e==nil{dir=candidate;if raw,e:=os.ReadFile(dir);e==nil&&strings.HasPrefix(string(raw),"gitdir: "){dir=strings.TrimSpace(strings.TrimPrefix(string(raw),"gitdir: "));if !filepath.IsAbs(dir){dir=filepath.Join(at,dir)}};break};if _,e:=os.Stat(filepath.Join(at,"HEAD"));e==nil{if _,e:=os.Stat(filepath.Join(at,"objects"));e==nil{dir=at;break}};if filepath.Dir(at)==at{dir=filepath.Join(cwd,".git");break}}}else if !filepath.IsAbs(dir){dir=filepath.Join(cwd,dir)};index:=os.Getenv("GIT_INDEX_FILE");if index==""{index=filepath.Join(dir,"index")}else if !filepath.IsAbs(index){index=filepath.Join(cwd,index)};if len(args)>0&&args[0]=="read-tree"{for i:=1;i<len(args);i++{if args[i]=="--"{break};if args[i]=="--index-output"&&i+1<len(args){index=args[i+1];i++}else if strings.HasPrefix(args[i],"--index-output="){index=strings.TrimPrefix(args[i],"--index-output=")}}};if !filepath.IsAbs(index){index=filepath.Join(cwd,index)};index=canonical(index);dir=canonical(dir);if common:=os.Getenv("GIT_COMMON_DIR");common!=""{if !filepath.IsAbs(common){common=filepath.Join(cwd,common)};dir=canonical(common)}else if raw,e:=os.ReadFile(filepath.Join(dir,"commondir"));e==nil{common:=strings.TrimSpace(string(raw));if !filepath.IsAbs(common){common=filepath.Join(dir,common)};dir=canonical(common)};rec,_:=json.Marshal(map[string]any{"pid":os.Getpid(),"cwd":canonical(cwd),"dir":canonical(dir),"index":index});f,e:=os.OpenFile(os.Getenv("FW_VERIFY_GIT_LOCATIONS"),os.O_CREATE|os.O_APPEND|os.O_WRONLY,0600);if e!=nil{os.Exit(122)};if _,e=f.Write(append(rec,'\n'));e!=nil{os.Exit(123)};f.Close();real:=os.Getenv("FW_VERIFY_REAL_GIT");if syscall.Exec(real,append([]string{real},os.Args[1:]...),os.Environ())!=nil{os.Exit(124)}}
`

func cliVerifyGitRecorder(t *testing.T, trace string) {
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
	if e = os.WriteFile(src, []byte(cliVerifyGitRecorderSource), 0600); e != nil {
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
func cliVerifyCanonicalGitDir(root string) string {
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
func cliVerifyForbiddenGitTrace(t *testing.T, path string, roots ...string) []string {
	t.Helper()
	raw, e := os.ReadFile(path)
	if os.IsNotExist(e) {
		return []string{}
	}
	if e != nil {
		t.Fatal(e)
	}
	type event struct {
		Event, Sid, Param, Value, Name string
		Argv                           []string
	}
	type location struct{ Dir, Index string }
	locations := map[int]location{}
	if data, e := os.ReadFile(path + ".locations"); e == nil {
		for _, line := range strings.Split(strings.TrimSpace(string(data)), "\n") {
			if line == "" {
				continue
			}
			var loc struct {
				Pid   int
				Dir   string
				Index string
			}
			if e := json.Unmarshal([]byte(line), &loc); e != nil {
				t.Fatal("malformed Git location record", e)
			}
			locations[loc.Pid] = location{Dir: loc.Dir, Index: loc.Index}
		}
	} else if !os.IsNotExist(e) {
		t.Fatal(e)
	}
	events := []event{}
	dirs := map[string]string{}
	modes := map[string]string{}
	for _, line := range strings.Split(strings.TrimSpace(string(raw)), "\n") {
		if line == "" {
			continue
		}
		var x event
		if e := json.Unmarshal([]byte(line), &x); e != nil {
			t.Fatal("malformed actual Git trace", e)
		}
		events = append(events, x)
		if x.Event == "cmd_mode" {
			modes[x.Sid] = x.Name
		}
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
		if (x.Event != "start" && x.Event != "alias") || !cliVerifyForbiddenOperation(x.Argv) {
			continue
		}
		// Metadata views may materialize copies of existing branch refs in owned
		// temporary storage. This never permits changes to the original Git dirs.
		verb := cliVerifyGitVerb(x.Argv)
		pathMode := modes[x.Sid] == "path"
		if modes[x.Sid] == "" && verb == "reset" {
			pathMode = cliVerifyResetPaths(x.Argv)
		}
		softMode := modes[x.Sid] == "soft" || modes[x.Sid] == "" && cliVerifyResetSoft(x.Argv)
		indexOnly := cliVerifyIndexOnlyOperation(verb) || (verb == "reset" || verb == "checkout") && pathMode
		metadata := verb == "branch" || verb == "reset" || verb == "checkout" || verb == "switch" || verb == "update-ref" || verb == "symbolic-ref"
		indexes := indexOnly || verb == "checkout" || verb == "switch" || verb == "reset" && !softMode
		if metadata || indexes {
			dir := cliVerifyGitLocation(x.Argv, root, dirs[x.Sid])
			index := filepath.Join(dir, "index")
			if i := strings.LastIndex(x.Sid, "-P"); i >= 0 {
				if pid, e := strconv.ParseUint(x.Sid[i+2:], 16, 32); e == nil && locations[int(pid)].Dir != "" {
					dir = locations[int(pid)].Dir
					index = locations[int(pid)].Index
				}
			}
			original := cliVerifyCanonicalGitDir(root)
			safe := true
			if metadata && !indexOnly {
				safe = cliVerifyOwnedGitStorage(dir, original, owned, false)
			}
			if indexes {
				safe = safe && cliVerifyOwnedGitStorage(index, filepath.Join(original, "index"), owned, true)
			}
			if safe {
				continue
			}
		}
		bad = append(bad, strings.Join(x.Argv, " "))
	}
	return bad
}
func cliVerifyResetPaths(argv []string) bool {
	args := cliVerifyGitArgs(argv)
	if len(args) == 0 {
		return false
	}
	args = args[1:]
	operands := 0
	for i, a := range args {
		if a == "--" {
			return i+1 < len(args)
		}
		if a == "-p" || a == "--patch" {
			return true
		}
		if strings.HasPrefix(a, "--pathspec-from-file") {
			return false
		}
		if !strings.HasPrefix(a, "-") {
			operands++
		}
	}
	return operands > 1
}
func cliVerifyResetSoft(argv []string) bool {
	args := cliVerifyGitArgs(argv)
	if len(args) == 0 {
		return false
	}
	soft := false
	for _, a := range args[1:] {
		if a == "--" {
			break
		}
		switch a {
		case "--soft", "--mixed", "--hard", "--merge", "--keep":
			soft = a == "--soft"
		}
	}
	return soft
}
func cliVerifyOwnedGitStorage(storage, original, owned string, index bool) bool {
	if p, e := filepath.EvalSymlinks(storage); e == nil {
		storage = p
	}
	if p, e := filepath.EvalSymlinks(original); e == nil {
		original = p
	}
	if p, e := filepath.EvalSymlinks(owned); e == nil {
		owned = p
	}
	if owned != "" {
		rel, e := filepath.Rel(owned, storage)
		if e == nil && rel != ".." && !strings.HasPrefix(rel, ".."+string(filepath.Separator)) && !filepath.IsAbs(rel) {
			return true
		}
	}
	// Removed owned views can also live under common-Git scratch storage.
	if filepath.Clean(storage) != filepath.Clean(original) {
		target := storage
		if !index {
			target = filepath.Join(storage, "HEAD")
		}
		if _, e := os.Stat(target); os.IsNotExist(e) {
			return true
		}
	}
	return false
}
func cliVerifyIndexOnlyOperation(verb string) bool {
	switch verb {
	case "add", "update-index", "read-tree", "checkout-index", "rm", "mv", "restore":
		return true
	}
	return false
}
func cliVerifyNoForbiddenGit(t *testing.T, path string, roots ...string) {
	t.Helper()
	if bad := cliVerifyForbiddenGitTrace(t, path, roots...); len(bad) > 0 {
		t.Error("Verify executed forbidden native Git operations", bad)
	}
}

func Test_APP028VerifyNativeOperationControls(t *testing.T) {
	f := cliVerifyFixture(t, "impl", "sha1", "none", false, false)
	root, base := f.Root, f.Head
	trace := filepath.Join(t.TempDir(), "trace")
	owned := filepath.Join(t.TempDir(), "owned")
	t.Setenv("GIT_TRACE2_EVENT", trace)
	t.Setenv("GIT_TRACE2_ENV_VARS", "GIT_DIR,GIT_INDEX_FILE")
	cliVerifyGitRecorder(t, trace)
	clear := func() {
		t.Helper()
		for _, p := range []string{trace, trace + ".locations"} {
			if e := os.WriteFile(p, nil, 0600); e != nil {
				t.Fatal(e)
			}
		}
	}
	expect := func(forbidden bool) {
		t.Helper()
		bad := cliVerifyForbiddenGitTrace(t, trace, root, owned)
		if (len(bad) > 0) != forbidden {
			t.Fatalf("native observer forbidden=%v actual=%v", forbidden, bad)
		}
	}
	git := func(args ...string) string { t.Helper(); return strings.TrimSpace(gitFixtureCommand(t, root, args...)) }
	for _, args := range [][]string{{"--version"}, {"-c", "note.value=commit add reset worktree", "rev-parse", "HEAD"}, {"branch", "--list"}, {"branch", "--format", "-m", "--list"}, {"worktree", "list", "--porcelain"}, {"add", "--dry-run", "--", "units/a/internal/private.txt"}} {
		clear()
		git(args...)
		expect(false)
	}
	clear()
	git("branch", "transient-branch", base)
	git("branch", "-D", "transient-branch")
	expect(true)
	clear()
	git("commit", "--allow-empty", "-qm", "transient commit")
	git("update-ref", "HEAD", base)
	expect(true)
	index := filepath.Join(cliVerifyCanonicalGitDir(root), "index")
	raw, e := os.ReadFile(index)
	if e != nil {
		t.Fatal(e)
	}
	projectPut(t, root, "units/a/internal/private.txt", "transient staged content")
	for _, args := range [][]string{{"add", "--", "units/a/internal/private.txt"}, {"update-index", "--add", "--", "units/a/internal/private.txt"}, {"reset", base, "--", "units/a/internal/private.txt"}, {"read-tree", base}} {
		clear()
		git(args...)
		if args[0] == "add" {
			after, e := os.ReadFile(index)
			if e != nil || string(after) == string(raw) {
				t.Fatal("staging premise absent", e)
			}
		}
		if e := os.WriteFile(index, raw, 0644); e != nil {
			t.Fatal(e)
		}
		expect(true)
	}
	detached := filepath.Join(t.TempDir(), "detached")
	clear()
	git("worktree", "add", "--detach", detached, base)
	git("worktree", "remove", detached)
	expect(true)
	metadata := filepath.Join(owned, "metadata")
	clear()
	git("init", "--bare", metadata)
	common := git("rev-parse", "--git-common-dir")
	if !filepath.IsAbs(common) {
		common = filepath.Join(root, common)
	}
	projectPut(t, metadata, "objects/info/alternates", filepath.Join(common, "objects")+"\n")
	git("--git-dir", metadata, "branch", "copied", base)
	expect(false)
	t.Run("owned-index", func(t *testing.T) {
		t.Setenv("GIT_INDEX_FILE", filepath.Join(owned, "index"))
		clear()
		git("read-tree", base)
		git("add", "--", "units/a/internal/private.txt")
		expect(false)
		after, e := os.ReadFile(index)
		if e != nil || string(after) != string(raw) {
			t.Fatal("owned index changed caller index", e)
		}
	})
	t.Run("owned-metadata-caller-index", func(t *testing.T) {
		t.Setenv("GIT_INDEX_FILE", index)
		clear()
		git("--git-dir", metadata, "--work-tree", root, "reset", base, "--", "units/a/internal/private.txt")
		expect(true)
	})
}
