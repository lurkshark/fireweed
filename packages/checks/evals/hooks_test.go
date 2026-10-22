package evals

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"

	k "github.com/lurkshark/fireweed/packages/checks"
	c "github.com/lurkshark/fireweed/packages/contracts"
	g "github.com/lurkshark/fireweed/packages/gitstore"
	p "github.com/lurkshark/fireweed/packages/project"
)

const hookHelper = "Test_CHK005_CHK006_CHK007_CHK008HookHelper"

func shellLiteral(s string) string { return "'" + strings.ReplaceAll(s, "'", `'"'"'`) + "'" }
func hookCommand(t *testing.T, mode string, args ...string) string {
	t.Helper()
	exe, err := os.Executable()
	if err != nil {
		t.Fatal(err)
	}
	command := "FIREWEED_HOOK_EVAL=1 GORACE=atexit_sleep_ms=0 exec " + shellLiteral(exe) + " -test.run=^" + hookHelper + "$ -- " + shellLiteral(mode)
	for _, arg := range args {
		command += " " + shellLiteral(arg)
	}
	return command
}

type hookDescription struct {
	Args                                           []string
	Dir, Module, ModuleDir, Report, Bin, Inherited string
}
type reportRecord struct {
	Report, Argument string
	Existed          bool
	Started          time.Time
}

func Test_CHK005_CHK006_CHK007_CHK008HookHelper(t *testing.T) {
	if os.Getenv("FIREWEED_HOOK_EVAL") != "1" {
		return
	}
	var args []string
	for i, arg := range os.Args {
		if arg == "--" {
			args = os.Args[i+1:]
			break
		}
	}
	if len(args) == 0 {
		os.Exit(90)
	}
	put := func(path string, data []byte) {
		if err := os.WriteFile(path, data, 0600); err != nil {
			os.Exit(91)
		}
	}
	switch args[0] {
	case "describe", "describehold":
		dir, _ := os.Getwd()
		data, _ := json.Marshal(hookDescription{args[1:], dir, os.Getenv("FIREWEED_MODULE"), os.Getenv("FIREWEED_MODULE_DIR"), os.Getenv("FIREWEED_REPORT"), os.Getenv("FIREWEED_BIN"), os.Getenv("FIREWEED_INHERITED_HOOK")})
		fmt.Fprintln(os.Stdout, "OUT")
		fmt.Fprintln(os.Stdout, string(data))
		if args[0] == "describehold" {
			put(os.Getenv("FIREWEED_REPORT"), []byte("ready"))
			time.Sleep(250 * time.Millisecond)
		}
		fmt.Fprintln(os.Stderr, "ERR")
	case "marker":
		put(args[1], []byte("unexpected"))
	case "treehold":
		child := exec.Command(os.Args[0], "-test.run=^"+hookHelper+"$", "--", "hold", args[1]+".child", args[2]+".child")
		child.Env = os.Environ()
		child.Stdout = os.Stdout
		child.Stderr = os.Stderr
		if err := child.Start(); err != nil {
			os.Exit(96)
		}
		fmt.Fprintln(os.Stdout, "hook started")
		put(args[1], []byte("ready"))
		time.Sleep(3 * time.Second)
		put(args[2], []byte("late"))
		_ = child.Wait()
	case "hold":
		fmt.Fprintln(os.Stdout, "hook started")
		put(args[1], []byte("ready"))
		time.Sleep(3 * time.Second)
		put(args[2], []byte("late"))
	case "report", "missing", "directory", "partial":
		path := os.Getenv("FIREWEED_REPORT")
		_, statErr := os.Stat(path)
		record := reportRecord{path, args[4], statErr == nil, time.Now()}
		data, _ := json.Marshal(record)
		put(args[3], data)
		fmt.Fprintln(os.Stdout, "eval output")
		fmt.Fprintln(os.Stderr, "eval err")
		if args[0] == "directory" {
			if err := os.Mkdir(path, 0700); err != nil {
				os.Exit(92)
			}
		} else if args[0] != "missing" {
			data, err := os.ReadFile(args[1])
			if err != nil {
				os.Exit(93)
			}
			put(path, data)
		}
		if args[0] == "partial" {
			time.Sleep(3 * time.Second)
			put(args[3]+".late", []byte("late"))
		}
		code, err := strconv.Atoi(args[2])
		if err != nil {
			os.Exit(94)
		}
		os.Exit(code)
	default:
		os.Exit(95)
	}
	os.Exit(0)
}
func hookGit(t *testing.T, root string, args ...string) string {
	t.Helper()
	argv := append([]string{"-C", root, "-c", "user.name=Hook Eval", "-c", "user.email=hook@example.invalid"}, args...)
	command := exec.Command("git", argv...)
	command.Env = append(os.Environ(), "GIT_CONFIG_NOSYSTEM=1", "GIT_CONFIG_GLOBAL="+os.DevNull)
	out, err := command.CombinedOutput()
	if err != nil {
		t.Fatalf("git %v: %v %s", args, err, out)
	}
	return string(out)
}
func hookFixture(t *testing.T) (p.Snapshot, c.Module) {
	t.Helper()
	root := filepath.Join(t.TempDir(), "root '雪 {report} $HOME")
	if err := os.MkdirAll(filepath.Join(root, "nested"), 0755); err != nil {
		t.Fatal(err)
	}
	hookGit(t, root, "init", "-q", "--initial-branch=main")
	repo, err := g.Open(root)
	if err != nil {
		t.Fatal(err)
	}
	m := evidenceModule("ALP-001", "ALP-002")
	m.Dir = "units/a '雪 {report} $HOME ; $(touch INJECTED)"
	m.DependsOn = []string{"beta"}
	m.Interface = []string{"api.go"}
	m.Document = c.Document{Metadata: map[string]any{"nested": []any{"keep"}}, Sections: []c.Section{{Name: "Decisions", Body: "Keep."}}}
	trap := hookCommand(t, "marker", filepath.Join(repo.Root, "MUST-NOT-RUN"))
	m.Hooks = c.Hooks{Eval: trap, Setup: trap, Boundaries: trap}
	return p.Snapshot{Root: repo.Root, Config: c.Config{Timeout: "5s", Hooks: c.Hooks{Eval: trap, Setup: trap, Boundaries: trap}}, Modules: []c.Module{m, {Name: "beta", IDPrefix: "BET", Dir: "units/beta"}}}, m
}
func parseDescription(t *testing.T, r g.CommandResult) hookDescription {
	t.Helper()
	if r.Err != nil || r.Exit != 0 || !strings.HasPrefix(r.Output, "OUT\n") || !strings.HasSuffix(r.Output, "\nERR\n") {
		t.Fatalf("hook output %+v", r)
	}
	var data hookDescription
	if err := json.Unmarshal([]byte(strings.TrimSuffix(strings.TrimPrefix(r.Output, "OUT\n"), "\nERR\n")), &data); err != nil {
		t.Fatal(err)
	}
	return data
}
func Test_CHK005QuotedPlaceholdersAndIdentity(t *testing.T) {
	s, m := hookFixture(t)
	root := s.Root
	s.Root = filepath.Join(root, "nested")
	for _, key := range []string{"FIREWEED_MODULE", "FIREWEED_MODULE_DIR", "FIREWEED_REPORT", "FIREWEED_BIN"} {
		t.Setenv(key, "STALE-BOOTSTRAP-IDENTITY")
	}
	before := jsonSnapshot(t, []any{s, m})
	envBefore := os.Environ()
	cwd, _ := os.Getwd()
	report := filepath.Join(root, "report ' 雪 {module} {git_common_dir} $(touch REPORT-INJECTION) .xml")
	command := hookCommand(t, "describe") + " {module} {module_dir} {report} {git_common_dir} '{unknown}'"
	r := k.RunHook(context.Background(), s, &m, command, report)
	data := parseDescription(t, r)
	repo, err := g.Open(root)
	if err != nil {
		t.Fatal(err)
	}
	exe, _ := os.Executable()
	if !reflect.DeepEqual(data.Args, []string{m.Name, m.Dir, report, repo.CommonDir, "{unknown}"}) || data.Dir != root || data.Module != m.Name || data.ModuleDir != m.Dir || data.Report != report || data.Bin != exe {
		t.Fatalf("expanded hook %+v", data)
	}
	nilData := parseDescription(t, k.RunHook(context.Background(), s, nil, command, report))
	if !reflect.DeepEqual(nilData.Args, []string{"", "", report, repo.CommonDir, "{unknown}"}) || nilData.Module != "" || nilData.ModuleDir != "" || nilData.Bin != exe || nilData.Report != report {
		t.Fatalf("nil module %+v", nilData)
	}
	afterCWD, _ := os.Getwd()
	if before != jsonSnapshot(t, []any{s, m}) || cwd != afterCWD || !reflect.DeepEqual(envBefore, os.Environ()) {
		t.Fatal("hook mutated inputs/parent")
	}
	for _, name := range []string{"INJECTED", "REPORT-INJECTION", "MUST-NOT-RUN"} {
		if _, err := os.Stat(filepath.Join(root, name)); !os.IsNotExist(err) {
			t.Fatalf("unexpected command %s", name)
		}
	}
}
func Test_CHK005LinkedWorktreeCommonDirectory(t *testing.T) {
	s, m := hookFixture(t)
	root := s.Root
	if err := os.WriteFile(filepath.Join(root, "seed"), []byte("seed"), 0600); err != nil {
		t.Fatal(err)
	}
	hookGit(t, root, "add", "seed")
	hookGit(t, root, "commit", "-qm", "seed")
	linked := filepath.Join(t.TempDir(), "linked ' {module_dir}")
	hookGit(t, root, "worktree", "add", "--detach", linked)
	if err := os.Mkdir(filepath.Join(linked, "nested"), 0755); err != nil {
		t.Fatal(err)
	}
	repo, err := g.Open(linked)
	if err != nil {
		t.Fatal(err)
	}
	s.Root = filepath.Join(linked, "nested")
	data := parseDescription(t, k.RunHook(context.Background(), s, &m, hookCommand(t, "describe")+" {git_common_dir}", ""))
	base, err := g.Open(root)
	if err != nil {
		t.Fatal(err)
	}
	if data.Dir != repo.Root || !reflect.DeepEqual(data.Args, []string{base.CommonDir}) {
		t.Fatalf("linked hook %+v", data)
	}
}
func Test_CHK006InvalidInputsAndEOF(t *testing.T) {
	s, m := hookFixture(t)
	marker := filepath.Join(s.Root, "started")
	command := hookCommand(t, "marker", marker)
	envBefore := os.Environ()
	cwd, _ := os.Getwd()
	before := jsonSnapshot(t, []any{s, m})
	for _, tt := range []struct {
		name, timeout, root, command string
		nilContext                   bool
	}{
		{name: "nil context", timeout: "5s", root: s.Root, command: command, nilContext: true},
		{name: "blank command", timeout: "5s", root: s.Root, command: " \n\t"},
		{name: "bad duration", timeout: "tomorrow", root: s.Root, command: command},
		{name: "zero duration", timeout: "0s", root: s.Root, command: command},
		{name: "negative duration", timeout: "-1s", root: s.Root, command: command},
		{name: "nongit", timeout: "5s", root: t.TempDir(), command: command},
		{name: "absent root", timeout: "5s", root: filepath.Join(t.TempDir(), "absent"), command: command},
	} {
		copy := s
		copy.Config.Timeout = tt.timeout
		copy.Root = tt.root
		var ctx context.Context = context.Background()
		if tt.nilContext {
			ctx = nil
		}
		r := k.RunHook(ctx, copy, &m, tt.command, "literal report")
		if r.Exit != 2 || r.Err == nil || strings.TrimSpace(r.Err.Error()) == "" {
			t.Fatalf("%s: %+v", tt.name, r)
		}
		if _, err := os.Stat(marker); !os.IsNotExist(err) {
			t.Fatalf("%s launched hook", tt.name)
		}
	}
	cancelled, cancel := context.WithCancel(context.Background())
	cancel()
	r := k.RunHook(cancelled, s, &m, command, "")
	if r.Exit != -1 || !errors.Is(r.Err, context.Canceled) {
		t.Fatalf("pre-cancel hook %+v", r)
	}
	if _, err := os.Stat(marker); !os.IsNotExist(err) {
		t.Fatal("pre-cancel started hook")
	}
	// read must receive EOF, even though the parent test runner has an open stdin.
	r = k.RunHook(context.Background(), s, &m, "if read answer; then exit 41; else printf EOF; fi", "")
	if r.Exit != 0 || r.Err != nil || r.Output != "EOF" {
		t.Fatalf("interactive stdin %+v", r)
	}
	afterCWD, _ := os.Getwd()
	if before != jsonSnapshot(t, []any{s, m}) || cwd != afterCWD || !reflect.DeepEqual(envBefore, os.Environ()) {
		t.Fatal("invalid hook changed inputs/parent")
	}
}
func awaitHookReady(t *testing.T, path string, done <-chan g.CommandResult) {
	t.Helper()
	until := time.Now().Add(1500 * time.Millisecond)
	for time.Now().Before(until) {
		if _, err := os.Stat(path); err == nil {
			return
		}
		select {
		case r := <-done:
			t.Fatalf("hook ended before ready %+v", r)
		default:
		}
		time.Sleep(5 * time.Millisecond)
	}
	t.Fatal("hook did not start")
}
func Test_CHK006ContextsBoundHook(t *testing.T) {
	for _, mode := range []string{"cancel", "parent deadline", "configured deadline"} {
		t.Run(mode, func(t *testing.T) {
			s, m := hookFixture(t)
			ready, late := filepath.Join(s.Root, "ready"), filepath.Join(s.Root, "late")
			ctx, cancel := context.WithCancel(context.Background())
			if mode == "parent deadline" {
				cancel()
				ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
			}
			defer cancel()
			if mode == "configured deadline" {
				s.Config.Timeout = "2s"
			}
			done := make(chan g.CommandResult, 1)
			start := time.Now()
			go func() { done <- k.RunHook(ctx, s, &m, hookCommand(t, "treehold", ready, late), "") }()
			awaitHookReady(t, ready, done)
			awaitHookReady(t, ready+".child", done)
			readyAt := time.Now()
			if mode == "cancel" {
				cancel()
			}
			var r g.CommandResult
			select {
			case r = <-done:
			case <-time.After(3 * time.Second):
				t.Fatal("hook context failed to bound execution")
			}
			want := context.DeadlineExceeded
			if mode == "cancel" {
				want = context.Canceled
			}
			if r.Exit != -1 || !errors.Is(r.Err, want) || !strings.Contains(r.Output, "hook started") || time.Since(start) > 2800*time.Millisecond {
				t.Fatalf("bounded hook %+v", r)
			}
			if wait := time.Until(readyAt.Add(3300 * time.Millisecond)); wait > 0 {
				time.Sleep(wait)
			}
			for _, path := range []string{late, late + ".child"} {
				if _, err := os.Stat(path); !os.IsNotExist(err) {
					t.Fatalf("hook tree survived cancellation: %s", path)
				}
			}
		})
	}
}
func evaluationCommand(t *testing.T, mode, source string, exit int, record string) string {
	return hookCommand(t, mode, source, strconv.Itoa(exit), record) + " {report}"
}
func readReportRecord(t *testing.T, path string) reportRecord {
	t.Helper()
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	var record reportRecord
	if err := json.Unmarshal(raw, &record); err != nil {
		t.Fatal(err)
	}
	if !filepath.IsAbs(record.Report) || record.Report != record.Argument || record.Existed || record.Started.IsZero() {
		t.Fatalf("report identity %+v", record)
	}
	if _, err := os.Stat(filepath.Dir(record.Report)); !os.IsNotExist(err) {
		t.Fatalf("report directory not cleaned %s", record.Report)
	}
	return record
}
func invalidCoverage(t *testing.T, cov k.Coverage, module c.Module, code string) {
	t.Helper()
	if cov.Module != module.Name || cov.Tests == nil || len(cov.Tests) != 0 || cov.BySpec == nil || len(cov.BySpec) != 0 || cov.Diagnostics == nil || len(cov.Diagnostics) != 1 {
		t.Fatalf("invalid evidence %+v", cov)
	}
	d := cov.Diagnostics[0]
	if d.Code != code || d.Path != module.Dir || d.Severity != "error" || strings.TrimSpace(d.Message) == "" {
		t.Fatalf("invalid diagnostic %+v", d)
	}
	lower := strings.ToLower(d.Message)
	corrective := false
	for _, verb := range []string{"repair", "configure", "write", "provide", "correct", "ensure", "check", "use ", "add "} {
		corrective = corrective || strings.Contains(lower, verb)
	}
	if !corrective {
		t.Fatalf("diagnostic lacks corrective guidance %+v", d)
	}
}
func Test_CHK007InvalidReportLifecycle(t *testing.T) {
	for _, mode := range []string{"missing", "directory", "report"} {
		t.Run(mode, func(t *testing.T) {
			s, m := hookFixture(t)
			source, record := filepath.Join(s.Root, "input.xml"), filepath.Join(s.Root, "record.json")
			if err := os.WriteFile(source, []byte(`<testsuite><testcase name="ALP001"/><broken>`), 0600); err != nil {
				t.Fatal(err)
			}
			m.Hooks.Eval = evaluationCommand(t, mode, source, 0, record)
			before := jsonSnapshot(t, []any{s, m})
			cov := k.Evaluate(context.Background(), s, m, []string{"ALP-001"}, nil)
			invalidCoverage(t, cov, m, "C11")
			if cov.Output != "eval output\neval err\n" || cov.HookExit != 0 {
				t.Fatalf("lost hook evidence %+v", cov)
			}
			readReportRecord(t, record)
			if before != jsonSnapshot(t, []any{s, m}) {
				t.Fatal("invalid report mutated inputs")
			}
		})
	}
	s, m := hookFixture(t)
	ownedTemp := t.TempDir()
	for _, key := range []string{"TMPDIR", "TEMP", "TMP"} {
		t.Setenv(key, ownedTemp)
	}
	for _, hook := range []string{"", " \n\t"} {
		copy := m
		copy.Hooks.Eval = hook
		invalidCoverage(t, k.Evaluate(context.Background(), s, copy, nil, nil), copy, "C07")
		entries, err := os.ReadDir(ownedTemp)
		if err != nil || len(entries) != 0 {
			t.Fatalf("missing/blank hook leaked reports %v %v", entries, err)
		}
	}
	if _, err := os.Stat(filepath.Join(s.Root, "MUST-NOT-RUN")); !os.IsNotExist(err) {
		t.Fatal("missing eval ran another hook")
	}
}
func Test_CHK007CancelledValidReportIsInvalid(t *testing.T) {
	s, m := hookFixture(t)
	s.Config.Timeout = "2s"
	source, record := filepath.Join(s.Root, "input.xml"), filepath.Join(s.Root, "record.json")
	xml := `<testsuite><testcase name="ALP001"><failure/></testcase><testcase name="ALP002"/></testsuite>`
	if err := os.WriteFile(source, []byte(xml), 0600); err != nil {
		t.Fatal(err)
	}
	m.Hooks.Eval = evaluationCommand(t, "partial", source, 1, record)
	cov := k.Evaluate(context.Background(), s, m, []string{"ALP-001"}, nil)
	invalidCoverage(t, cov, m, "C11")
	if cov.HookExit != -1 || cov.Output != "eval output\neval err\n" {
		t.Fatalf("cancelled report %+v", cov)
	}
	written := readReportRecord(t, record)
	if wait := time.Until(written.Started.Add(3300 * time.Millisecond)); wait > 0 {
		time.Sleep(wait)
	}
	if _, err := os.Stat(record + ".late"); !os.IsNotExist(err) {
		t.Fatal("cancelled report hook survived")
	}
}
func Test_CHK008RealEvidenceAndOwnership(t *testing.T) {
	s, m := hookFixture(t)
	trapMarker := installHookTraps(t, s.Root)
	source := filepath.Join(s.Root, "input.xml")
	record := filepath.Join(s.Root, "record.json")
	xml := `<testsuites tests="999"><testsuite><testcase name="ALP001"><error message="failure">trace</error></testcase><testcase name="ALP002"/><testcase name="BET001"><skipped/></testcase></testsuite></testsuites>`
	if err := os.WriteFile(source, []byte(xml), 0600); err != nil {
		t.Fatal(err)
	}
	m.Hooks.Eval = evaluationCommand(t, "report", source, 37, record)
	first := []string{"ALP-001"}
	already := []string{}
	before := jsonSnapshot(t, []any{s, m, first, already})
	parentEnv := os.Environ()
	parentDir, _ := os.Getwd()
	refs := hookGit(t, s.Root, "for-each-ref")
	parsed, err := k.ParseJUnit([]byte(xml))
	if err != nil {
		t.Fatal(err)
	}
	expected := k.Analyze(m, []string{"ALP", "BET"}, parsed, 37, first, already)
	expected.Output = "eval output\neval err\n"
	if len(expected.Tests) != 3 || len(expected.Diagnostics) != 1 || expected.Diagnostics[0].Code != "C12" || !strings.Contains(expected.Diagnostics[0].Message, "BET-001") {
		t.Fatalf("evidence fixture %+v", expected)
	}
	a := k.Evaluate(context.Background(), s, m, first, already)
	if !reflect.DeepEqual(a, expected) {
		t.Fatalf("actual %+v expected %+v", a, expected)
	}
	pathA := readReportRecord(t, record)
	b := k.Evaluate(context.Background(), s, m, first, already)
	pathB := readReportRecord(t, record)
	if !reflect.DeepEqual(b, expected) || pathA.Report == pathB.Report {
		t.Fatalf("reused report/evidence %+v %+v", a, b)
	}
	saved := jsonSnapshot(t, b)
	a.Tests[0].Errors[0].Message = "changed"
	a.Tests[0].IDs[0] = "changed"
	a.BySpec["ALP-001"][0] = 99
	a.Diagnostics[0].Message = "changed"
	a.Tests[2].Skipped.Message = "changed"
	if saved != jsonSnapshot(t, b) || before != jsonSnapshot(t, []any{s, m, first, already}) {
		t.Fatal("evaluation evidence aliases input/other call")
	}
	afterDir, _ := os.Getwd()
	if parentDir != afterDir || !reflect.DeepEqual(parentEnv, os.Environ()) || refs != hookGit(t, s.Root, "for-each-ref") {
		t.Fatal("evaluation mutated parent/refs")
	}
	if _, err := os.Stat(filepath.Join(s.Root, "MUST-NOT-RUN")); !os.IsNotExist(err) {
		t.Fatal("setup/boundary/default hook ran")
	}
	if _, err := os.Stat(trapMarker); !os.IsNotExist(err) {
		t.Fatal("evaluation invoked an agent/bootstrap/network/Git mutation command")
	}
}

func Test_CHK006OverlapAndActualExit(t *testing.T) {
	s, m := hookFixture(t)
	r := k.RunHook(context.Background(), s, &m, "printf stdout; printf stderr >&2; exit 37", "")
	if r.Exit != 37 || r.Err == nil || r.Output != "stdoutstderr" || errors.Is(r.Err, context.Canceled) || errors.Is(r.Err, context.DeadlineExceeded) {
		t.Fatalf("ordinary hook exit %+v", r)
	}
	other := m
	other.Name = "beta"
	other.Dir = "units/beta"
	reportA, reportB := filepath.Join(s.Root, "a.ready"), filepath.Join(s.Root, "b.ready")
	command := hookCommand(t, "describehold") + " {module} {report}"
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	a, b := make(chan g.CommandResult, 1), make(chan g.CommandResult, 1)
	before := jsonSnapshot(t, []any{s, m, other})
	cwd, _ := os.Getwd()
	env := os.Environ()
	go func() { a <- k.RunHook(ctx, s, &m, command, reportA) }()
	go func() { b <- k.RunHook(context.Background(), s, &other, command, reportB) }()
	awaitHookReady(t, reportA, a)
	awaitHookReady(t, reportB, b)
	cancel()
	var first, second g.CommandResult
	select {
	case first = <-a:
	case <-time.After(2 * time.Second):
		t.Fatal("overlap cancellation stuck")
	}
	select {
	case second = <-b:
	case <-time.After(2 * time.Second):
		t.Fatal("independent hook cancelled or stuck")
	}
	if first.Exit != -1 || !errors.Is(first.Err, context.Canceled) || !strings.Contains(first.Output, `"Module":"alpha"`) || strings.Contains(first.Output, reportB) {
		t.Fatalf("mixed cancelled call %+v", first)
	}
	desc := parseDescription(t, second)
	if desc.Module != "beta" || desc.ModuleDir != other.Dir || desc.Report != reportB || !reflect.DeepEqual(desc.Args, []string{"beta", reportB}) {
		t.Fatalf("mixed independent call %+v", desc)
	}
	after, _ := os.Getwd()
	if cwd != after || !reflect.DeepEqual(env, os.Environ()) || before != jsonSnapshot(t, []any{s, m, other}) {
		t.Fatal("overlapping hook changed caller state")
	}
}
func Test_CHK007FatalLaunchConfigSignalAndCleanup(t *testing.T) {
	for _, mode := range []string{"config", "discovery", "signal", "launch"} {
		t.Run(mode, func(t *testing.T) {
			s, m := hookFixture(t)
			ownedTemp := t.TempDir()
			for _, key := range []string{"TMPDIR", "TEMP", "TMP"} {
				t.Setenv(key, ownedTemp)
			}
			source, record := filepath.Join(s.Root, "input.xml"), filepath.Join(s.Root, "record.json")
			if err := os.WriteFile(source, []byte(`<testsuite><testcase name="ALP001"><failure/></testcase><testcase name="ALP002"/></testsuite>`), 0600); err != nil {
				t.Fatal(err)
			}
			m.Hooks.Eval = evaluationCommand(t, "report", source, 1, record)
			switch mode {
			case "config":
				s.Config.Timeout = "invalid"
			case "discovery":
				s.Root = t.TempDir()
			case "signal":
				m.Hooks.Eval = strings.Replace(m.Hooks.Eval, "exec ", "", 1) + "; kill -TERM $$"
			case "launch":
				// Git discovery remains available, but a shell launch cannot succeed.
				realGit, err := exec.LookPath("git")
				if err != nil {
					t.Fatal(err)
				}
				bin := t.TempDir()
				if err := os.Symlink(realGit, filepath.Join(bin, "git")); err != nil {
					t.Skip("local symbolic links unavailable")
				}
				t.Setenv("PATH", bin)
			}
			cov := k.Evaluate(context.Background(), s, m, []string{"ALP-001"}, nil)
			invalidCoverage(t, cov, m, "C11")
			if mode == "signal" {
				if cov.HookExit != -1 || cov.Output != "eval output\neval err\n" {
					t.Fatalf("signal evidence %+v", cov)
				}
				readReportRecord(t, record)
			} else {
				if cov.HookExit != 2 {
					t.Fatalf("fatal hook exit %+v", cov)
				}
				if _, err := os.Stat(record); !os.IsNotExist(err) {
					t.Fatal("fatal setup launched eval")
				}
			}
			entries, err := os.ReadDir(ownedTemp)
			if err != nil {
				t.Fatal(err)
			}
			if len(entries) != 0 {
				t.Fatalf("failed hook leaked temp reports %v", entries)
			}
		})
	}
}
func Test_CHK007InvalidResultsAreOwned(t *testing.T) {
	s, m := hookFixture(t)
	source, record := filepath.Join(s.Root, "invalid.xml"), filepath.Join(s.Root, "record.json")
	if err := os.WriteFile(source, []byte(`<testsuite><`), 0600); err != nil {
		t.Fatal(err)
	}
	m.Hooks.Eval = evaluationCommand(t, "report", source, 37, record)
	a := k.Evaluate(context.Background(), s, m, nil, nil)
	invalidCoverage(t, a, m, "C11")
	firstPath := readReportRecord(t, record)
	b := k.Evaluate(context.Background(), s, m, nil, nil)
	invalidCoverage(t, b, m, "C11")
	secondPath := readReportRecord(t, record)
	if a.HookExit != 37 || b.HookExit != 37 || a.Output != "eval output\neval err\n" || firstPath.Report == secondPath.Report {
		t.Fatal("invalid results lost actual exits/output or reused reports")
	}
	saved := jsonSnapshot(t, b)
	a.Diagnostics[0].Message = "changed"
	a.BySpec["changed"] = []int{1}
	a.Tests = append(a.Tests, k.TestCase{Name: "first"})
	if saved != jsonSnapshot(t, b) {
		t.Fatal("invalid coverage aliases another result")
	}
	b.Tests = append(b.Tests, k.TestCase{Name: "second"})
	if a.Tests[0].Name != "first" {
		t.Fatal("empty invalid testcase storage shared")
	}
}

func installHookTraps(t *testing.T, root string) string {
	t.Helper()
	realGit, err := exec.LookPath("git")
	if err != nil {
		t.Fatal(err)
	}
	bin, marker := filepath.Join(root, "trap-bin"), filepath.Join(root, "UNEXPECTED-COMMAND")
	if err := os.Mkdir(bin, 0755); err != nil {
		t.Fatal(err)
	}
	script := "#!/bin/sh\nprintf trap >> " + shellLiteral(marker) + "\nexit 97\n"
	for _, name := range []string{"codex", "claude", "fireweed", "fireweed-bootstrap", "bootstrap-controller", "curl", "wget", "ssh", "nc"} {
		if err := os.WriteFile(filepath.Join(bin, name), []byte(script), 0755); err != nil {
			t.Fatal(err)
		}
	}
	gitScript := "#!/bin/sh\nfor arg in \"$@\"; do case \"$arg\" in fetch|push|pull|clone|ls-remote|commit|update-ref|reset|checkout|switch|add|stash|clean|rm|mv|branch|tag|init) printf trap >> " + shellLiteral(marker) + "; exit 97;; esac; done\nexec " + shellLiteral(realGit) + " \"$@\"\n"
	if err := os.WriteFile(filepath.Join(bin, "git"), []byte(gitScript), 0755); err != nil {
		t.Fatal(err)
	}
	t.Setenv("PATH", bin+string(os.PathListSeparator)+os.Getenv("PATH"))
	t.Setenv("FIREWEED_BIN", filepath.Join(bin, "fireweed-bootstrap"))
	control := exec.Command("sh", "-c", "codex")
	if err := control.Run(); err == nil {
		t.Fatal("command trap positive control did not fail")
	}
	if _, err := os.Stat(marker); err != nil {
		t.Fatal("command trap positive control did not execute")
	}
	if err := os.Remove(marker); err != nil {
		t.Fatal(err)
	}
	return marker
}

func Test_CHK005RepeatedRelativeAndInheritedValues(t *testing.T) {
	s, m := hookFixture(t)
	t.Setenv("FIREWEED_INHERITED_HOOK", "parent sentinel ' 雪 {module} $HOME")
	report := "relative ' report {module_dir} $(touch RELATIVE-INJECTION).xml"
	repo, err := g.Open(s.Root)
	if err != nil {
		t.Fatal(err)
	}
	command := hookCommand(t, "describe") + " {module} {module_dir} {report} {git_common_dir} {module} {module_dir} {report} {git_common_dir} '{unrecognized}'"
	data := parseDescription(t, k.RunHook(context.Background(), s, &m, command, report))
	want := []string{m.Name, m.Dir, report, repo.CommonDir, m.Name, m.Dir, report, repo.CommonDir, "{unrecognized}"}
	if !reflect.DeepEqual(data.Args, want) || data.Report != report || data.Inherited != os.Getenv("FIREWEED_INHERITED_HOOK") {
		t.Fatalf("literal repeated/inherited values %+v", data)
	}
	if _, err := os.Stat(filepath.Join(s.Root, "RELATIVE-INJECTION")); !os.IsNotExist(err) {
		t.Fatal("relative report was shell-expanded")
	}
}
func Test_CHK007ParentContextCancellation(t *testing.T) {
	s, m := hookFixture(t)
	source, record := filepath.Join(s.Root, "input.xml"), filepath.Join(s.Root, "record.json")
	if err := os.WriteFile(source, []byte(`<testsuite><testcase name="ALP001"><failure/></testcase><testcase name="ALP002"/></testsuite>`), 0600); err != nil {
		t.Fatal(err)
	}
	m.Hooks.Eval = evaluationCommand(t, "partial", source, 1, record)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	done := make(chan k.Coverage, 1)
	go func() { done <- k.Evaluate(ctx, s, m, []string{"ALP-001"}, nil) }()
	awaitHookReady(t, record, nil)
	cancel()
	var cov k.Coverage
	select {
	case cov = <-done:
	case <-time.After(time.Second):
		t.Fatal("Evaluate ignored parent cancellation")
	}
	invalidCoverage(t, cov, m, "C11")
	if cov.HookExit != -1 || cov.Output != "eval output\neval err\n" {
		t.Fatalf("parent cancellation outcome %+v", cov)
	}
	written := readReportRecord(t, record)
	if wait := time.Until(written.Started.Add(3300 * time.Millisecond)); wait > 0 {
		time.Sleep(wait)
	}
	if _, err := os.Stat(record + ".late"); !os.IsNotExist(err) {
		t.Fatal("Evaluate parent cancellation left hook running")
	}
}
func Test_CHK008SelectedModulePoliciesAndZeroExit(t *testing.T) {
	for _, mode := range []string{"zero exit and exemptions", "multiple policy errors", "removed ID", "quarantine and missing", "zero exit reporter contradiction"} {
		t.Run(mode, func(t *testing.T) {
			s, old := hookFixture(t)
			marker := installHookTraps(t, s.Root)
			selected := old
			selected.Specs = []c.Spec{{ID: "ALP-001", Body: "It MUST persist."}, {ID: "ALP-003", Body: "It MUST be projected."}}
			xml := `<testsuite><testcase name="ALP001"/><testcase name="ALP003"/></testsuite>`
			exit := 0
			first, already := []string{"ALP-003", "ALP-001"}, []string{"ALP-003", "ALP-001"}
			codes := []string{}
			switch mode {
			case "multiple policy errors":
				selected.Specs = []c.Spec{{ID: "ALP-004", Body: "It MUST pass."}, {ID: "ALP-003", Body: "It MUST be covered."}, {ID: "ALP-001", Tag: "[review]", Body: "It MUST be reviewed."}}
				xml = `<testsuite><testcase name="ALP001"/><testcase name="ALP004"><failure/></testcase></testsuite>`
				exit = 37
				first = []string{"ALP-003"}
				already = []string{"ALP-001"}
				codes = []string{"C13", "C14", "P1-02"}
			case "removed ID":
				xml = `<testsuite><testcase name="ALP001"/><testcase name="ALP003"/><testcase name="ALP002"><skipped/></testcase></testsuite>`
				codes = []string{"C12"}
			case "quarantine and missing":
				selected.Specs[0].Tag = "[quarantine: known temporary issue]"
				selected.Specs = append(selected.Specs, c.Spec{ID: "ALP-004", Body: "It MUST have coverage."})
				xml = `<testsuite><testcase name="ALP001"><failure/></testcase><testcase name="ALP003"/></testsuite>`
				exit = 37
				first = []string{"ALP-003"}
				already = []string{"ALP-003"}
				codes = []string{"C13"}
			case "zero exit reporter contradiction":
				xml = `<testsuite><testcase name="ALP001"><error/></testcase><testcase name="ALP003"/></testsuite>`
				first = []string{"ALP-001"}
				already = nil
				codes = []string{"C11"}
			}
			source, record := filepath.Join(s.Root, "input.xml"), filepath.Join(s.Root, "record.json")
			if err := os.WriteFile(source, []byte(xml), 0600); err != nil {
				t.Fatal(err)
			}
			selected.Hooks.Eval = evaluationCommand(t, "report", source, exit, record)
			if reflect.DeepEqual(selected.Specs, s.Modules[0].Specs) {
				t.Fatal("fixture failed to distinguish selected and stale snapshot module")
			}
			before := jsonSnapshot(t, []any{s, selected, first, already})
			cases, err := k.ParseJUnit([]byte(xml))
			if err != nil {
				t.Fatal(err)
			}
			expected := k.Analyze(selected, []string{"ALP", "BET"}, cases, exit, first, already)
			expected.Output = "eval output\neval err\n"
			actualCodes := []string{}
			for _, d := range expected.Diagnostics {
				actualCodes = append(actualCodes, d.Code)
			}
			if !reflect.DeepEqual(actualCodes, codes) {
				t.Fatalf("policy fixture %+v expected codes %v", expected, codes)
			}
			got := k.Evaluate(context.Background(), s, selected, first, already)
			if !reflect.DeepEqual(got, expected) {
				t.Fatalf("policy composition got%+v expected%+v", got, expected)
			}
			readReportRecord(t, record)
			if before != jsonSnapshot(t, []any{s, selected, first, already}) {
				t.Fatal("policy composition mutated inputs")
			}
			if _, err := os.Stat(marker); !os.IsNotExist(err) {
				t.Fatal("policy composition invoked unrelated command")
			}
		})
	}
}
