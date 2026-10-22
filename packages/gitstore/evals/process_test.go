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
	"runtime"
	"strings"
	"testing"
	"time"

	g "github.com/lurkshark/fireweed/packages/gitstore"
)

const processHelperName = "Test_GIT005_GIT006_GIT007ProcessHelper"

func helperCommand(t *testing.T, mode string, args ...string) []string {
	t.Helper()
	exe, err := os.Executable()
	if err != nil {
		t.Fatal(err)
	}
	return append([]string{exe, "-test.run=^" + processHelperName + "$", "--", mode}, args...)
}
func Test_GIT005_GIT006_GIT007ProcessHelper(t *testing.T) {
	if os.Getenv("FIREWEED_PROCESS_EVAL") != "1" {
		return
	}
	args := []string{}
	for i, arg := range os.Args {
		if arg == "--" {
			args = os.Args[i+1:]
			break
		}
	}
	if len(args) == 0 {
		os.Exit(90)
	}
	marker := func(path string) {
		if err := os.WriteFile(path, []byte("helper"), 0600); err != nil {
			os.Exit(91)
		}
	}
	switch args[0] {
	case "echo":
		dir, err := os.Getwd()
		if err != nil {
			os.Exit(92)
		}
		b, _ := json.Marshal(map[string]any{"args": args[1:], "dir": dir, "inherited": os.Getenv("FIREWEED_INHERITED"), "override": os.Getenv("FIREWEED_OVERRIDE"), "empty": os.Getenv("FIREWEED_EMPTY"), "equals": os.Getenv("FIREWEED_EQUALS")})
		fmt.Fprintln(os.Stdout, "OUT")
		fmt.Fprintln(os.Stdout, string(b))
		fmt.Fprintln(os.Stderr, "ERR")
	case "exit":
		fmt.Fprintln(os.Stdout, "before exit")
		fmt.Fprintln(os.Stderr, "error exit")
		os.Exit(37)
	case "long":
		fmt.Fprint(os.Stdout, strings.Repeat("O", 524288))
		fmt.Fprint(os.Stderr, strings.Repeat("E", 524288))
		fmt.Fprint(os.Stdout, "END")
	case "marker":
		marker(args[1])
	case "hold":
		fmt.Fprintln(os.Stdout, args[2])
		marker(args[1])
		time.Sleep(250 * time.Millisecond)
		fmt.Fprintln(os.Stderr, "held end")
	case "tree", "branch":
		mode := "branch"
		if args[0] == "branch" {
			mode = "leaf"
		}
		child := exec.Command(os.Args[0], "-test.run=^"+processHelperName+"$", "--", mode, args[1])
		child.Env = os.Environ()
		child.Stdout = os.Stdout
		child.Stderr = os.Stderr
		if err := child.Start(); err != nil {
			os.Exit(93)
		}
		fmt.Fprintln(os.Stdout, args[0]+" ready")
		marker(filepath.Join(args[1], args[0]+".ready"))
		time.Sleep(4 * time.Second)
		_ = child.Wait()
	case "leaf":
		fmt.Fprintln(os.Stderr, "leaf ready")
		marker(filepath.Join(args[1], "leaf.ready"))
		time.Sleep(3 * time.Second)
		marker(filepath.Join(args[1], "leaf.late"))
		fmt.Fprintln(os.Stdout, "leaf late")
	default:
		os.Exit(94)
	}
	os.Exit(0)
}
func processEnvironment() []string { return []string{"FIREWEED_PROCESS_EVAL=1"} }
func boundedProcess(t *testing.T, dir string, argv, env []string) g.CommandResult {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	defer cancel()
	return g.Run(ctx, dir, argv, env)
}
func Test_GIT005LiteralArgumentsAndEnvironment(t *testing.T) {
	t.Setenv("FIREWEED_INHERITED", "parent-value")
	t.Setenv("FIREWEED_OVERRIDE", "parent-override")
	t.Setenv("FIREWEED_EMPTY", "parent-nonempty")
	dir := filepath.Join(t.TempDir(), "dir with 'quotes' 雪")
	if err := os.Mkdir(dir, 0755); err != nil {
		t.Fatal(err)
	}
	dir, err := filepath.EvalSymlinks(dir)
	if err != nil {
		t.Fatal(err)
	}
	beforeDir, _ := os.Getwd()
	beforeEnv := os.Environ()
	marker := filepath.Join(dir, "SHELL-EXPANDED")
	args := []string{"two words", "'literal quote'", `"double"`, "$FIREWEED_INHERITED", "$(touch " + marker + ")", "`touch " + marker + "`", "; echo nope", "*.go", "line1\nline2", "", "雪"}
	argv := helperCommand(t, "echo", args...)
	env := append(processEnvironment(), "FIREWEED_OVERRIDE=first", "FIREWEED_EMPTY=", "FIREWEED_EQUALS=a=b=c", "FIREWEED_OVERRIDE=last")
	savedArgv, savedEnv := append([]string(nil), argv...), append([]string(nil), env...)
	got := boundedProcess(t, dir, argv, env)
	if got.Exit != 0 || got.Err != nil || !strings.HasPrefix(got.Output, "OUT\n") || !strings.HasSuffix(got.Output, "\nERR\n") {
		t.Fatalf("literal process %+v", got)
	}
	var data struct {
		Args                                    []string
		Dir, Inherited, Override, Empty, Equals string
	}
	if err := json.Unmarshal([]byte(strings.TrimSuffix(strings.TrimPrefix(got.Output, "OUT\n"), "\nERR\n")), &data); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(data.Args, args) || data.Dir != dir || data.Inherited != "parent-value" || data.Override != "last" || data.Empty != "" || data.Equals != "a=b=c" {
		t.Fatalf("child data %+v", data)
	}
	afterDir, _ := os.Getwd()
	if beforeDir != afterDir || !reflect.DeepEqual(beforeEnv, os.Environ()) || !reflect.DeepEqual(argv, savedArgv) || !reflect.DeepEqual(env, savedEnv) {
		t.Fatal("changed parent or inputs")
	}
	if _, err := os.Stat(marker); !os.IsNotExist(err) {
		t.Fatal("shell expansion ran")
	}
	if runtime.GOOS == "windows" {
		got = boundedProcess(t, dir, helperCommand(t, "echo"), append(processEnvironment(), "fireweed_override=lower", "FIREWEED_OVERRIDE=upper"))
		if got.Err != nil || !strings.Contains(got.Output, `"override":"upper"`) {
			t.Fatalf("Windows environment key semantics %+v", got)
		}
	}
}
func Test_GIT005ExitCodesAndCompleteOutput(t *testing.T) {
	root := t.TempDir()
	got := boundedProcess(t, root, helperCommand(t, "exit"), processEnvironment())
	if got.Exit != 37 || got.Err == nil || got.Output != "before exit\nerror exit\n" {
		t.Fatalf("failure %+v", got)
	}
	if errors.Is(got.Err, context.Canceled) || errors.Is(got.Err, context.DeadlineExceeded) {
		t.Fatal("ordinary failure mislabeled context error")
	}
	got = boundedProcess(t, root, helperCommand(t, "long"), processEnvironment())
	if got.Err != nil || got.Exit != 0 || got.Output != strings.Repeat("O", 524288)+strings.Repeat("E", 524288)+"END" {
		t.Fatalf("long output code%d err%v len%d", got.Exit, got.Err, len(got.Output))
	}
}
func Test_GIT005_GIT007InvalidCallsAndRecovery(t *testing.T) {
	root := t.TempDir()
	marker := filepath.Join(root, "started")
	good := helperCommand(t, "marker", marker)
	for _, tt := range []struct {
		name, dir  string
		argv, env  []string
		nilContext bool
	}{
		{name: "nil argv", dir: root, env: processEnvironment()},
		{name: "empty argv", dir: root, argv: []string{}, env: processEnvironment()},
		{name: "empty executable", dir: root, argv: []string{"", "marker", marker}, env: processEnvironment()},
		{name: "nil context", dir: root, argv: good, env: processEnvironment(), nilContext: true},
		{name: "bad directory", dir: filepath.Join(root, "absent"), argv: good, env: processEnvironment()},
		{name: "missing executable", dir: root, argv: []string{filepath.Join(root, "absent-executable"), marker}, env: processEnvironment()},
		{name: "override without equals", dir: root, argv: good, env: append(processEnvironment(), "INVALID_OVERRIDE")},
		{name: "empty override key", dir: root, argv: good, env: append(processEnvironment(), "=value")},
		{name: "NUL override", dir: root, argv: good, env: append(processEnvironment(), "BAD=one\x00two")},
	} {
		ctx := context.Background()
		if tt.nilContext {
			ctx = nil
		}
		got := g.Run(ctx, tt.dir, tt.argv, tt.env)
		if got.Exit != 2 || got.Err == nil || strings.TrimSpace(got.Err.Error()) == "" {
			t.Fatalf("%s: %+v", tt.name, got)
		}
		if _, err := os.Stat(marker); !os.IsNotExist(err) {
			t.Fatalf("%s started fallback", tt.name)
		}
		recovered := boundedProcess(t, root, helperCommand(t, "echo", "recovered"), processEnvironment())
		if recovered.Exit != 0 || recovered.Err != nil || !strings.Contains(recovered.Output, "recovered") {
			t.Fatalf("%s poisoned later call %+v", tt.name, recovered)
		}
	}
}
func Test_GIT006PreCancelledDoesNotStart(t *testing.T) {
	root := t.TempDir()
	marker := filepath.Join(root, "started")
	cancelled, cancel := context.WithCancel(context.Background())
	cancel()
	expired, finish := context.WithDeadline(context.Background(), time.Now().Add(-time.Second))
	defer finish()
	for _, ctx := range []context.Context{cancelled, expired} {
		got := g.Run(ctx, root, helperCommand(t, "marker", marker), processEnvironment())
		if got.Exit != -1 || !errors.Is(got.Err, ctx.Err()) || got.Output != "" {
			t.Fatalf("pre-cancelled %+v", got)
		}
		if _, err := os.Stat(marker); !os.IsNotExist(err) {
			t.Fatal("pre-cancelled process started")
		}
		after := boundedProcess(t, root, helperCommand(t, "echo", "after cancellation"), processEnvironment())
		if after.Exit != 0 || after.Err != nil {
			t.Fatal(after)
		}
	}
}
func waitProcessMarkers(t *testing.T, paths []string, done <-chan g.CommandResult) {
	t.Helper()
	deadline := time.Now().Add(1500 * time.Millisecond)
	for time.Now().Before(deadline) {
		all := true
		for _, path := range paths {
			if _, err := os.Stat(path); err != nil {
				all = false
			}
		}
		if all {
			return
		}
		select {
		case result := <-done:
			t.Fatalf("process ended before ready: %+v", result)
		default:
		}
		time.Sleep(5 * time.Millisecond)
	}
	t.Fatal("local process helper did not start its complete tree")
}
func Test_GIT006CancellationKillsTree(t *testing.T) {
	if runtime.GOOS != "darwin" && runtime.GOOS != "linux" && runtime.GOOS != "windows" {
		t.Skip("process-tree contract targets Linux/macOS/Windows")
	}
	for _, deadline := range []bool{false, true} {
		t.Run(fmt.Sprintf("deadline_%v", deadline), func(t *testing.T) {
			root := t.TempDir()
			start := time.Now()
			var ctx context.Context
			var cancel context.CancelFunc
			if deadline {
				ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
			} else {
				ctx, cancel = context.WithCancel(context.Background())
			}
			defer cancel()
			done := make(chan g.CommandResult, 1)
			go func() { done <- g.Run(ctx, root, helperCommand(t, "tree", root), processEnvironment()) }()
			waitProcessMarkers(t, []string{filepath.Join(root, "tree.ready"), filepath.Join(root, "branch.ready"), filepath.Join(root, "leaf.ready")}, done)
			readyTime := time.Now()
			// This process is outside Run's tree and must survive the cancellation.
			outside := filepath.Join(root, "unrelated.ready")
			unrelated := exec.Command(helperCommand(t, "hold", outside, "unrelated")[0], helperCommand(t, "hold", outside, "unrelated")[1:]...)
			unrelated.Env = append(os.Environ(), processEnvironment()...)
			unrelatedDone := make(chan error, 1)
			if err := unrelated.Start(); err != nil {
				t.Fatal(err)
			}
			go func() { unrelatedDone <- unrelated.Wait() }()
			if !deadline {
				cancel()
			}
			var got g.CommandResult
			select {
			case got = <-done:
			case <-time.After(3 * time.Second):
				t.Fatal("cancellation failed to bound held output pipes")
			}
			if got.Exit != -1 || !errors.Is(got.Err, ctx.Err()) || ctx.Err() == nil {
				t.Fatalf("context outcome %+v expected %v", got, ctx.Err())
			}
			for _, text := range []string{"tree ready", "branch ready", "leaf ready"} {
				if !strings.Contains(got.Output, text) {
					t.Fatalf("lost pre-cancellation output %+v", got)
				}
			}
			if time.Since(start) > 2800*time.Millisecond {
				t.Fatal("context termination did not return promptly")
			}
			select {
			case err := <-unrelatedDone:
				if err != nil {
					t.Fatalf("unrelated process killed: %v", err)
				}
			case <-time.After(time.Second):
				t.Fatal("unrelated process stuck")
			}
			if _, err := os.Stat(outside); err != nil {
				t.Fatal("unrelated marker missing")
			}
			if wait := time.Until(readyTime.Add(3300 * time.Millisecond)); wait > 0 {
				time.Sleep(wait)
			}
			if _, err := os.Stat(filepath.Join(root, "leaf.late")); !os.IsNotExist(err) {
				t.Fatal("grandchild survived cancellation and performed delayed write")
			}
			after := boundedProcess(t, root, helperCommand(t, "echo", "after tree"), processEnvironment())
			if after.Exit != 0 || after.Err != nil {
				t.Fatal(after)
			}
		})
	}
}
func Test_GIT007ConcurrentOwnership(t *testing.T) {
	root := t.TempDir()
	cwd, _ := os.Getwd()
	envBefore := os.Environ()
	firstReady, secondReady := filepath.Join(root, "first.ready"), filepath.Join(root, "second.ready")
	argvA := helperCommand(t, "hold", firstReady, "first output")
	argvBacking := append(append([]string{}, argvA...), "spare argv sentinel")
	argvA = argvBacking[:len(argvA)]
	envBacking := append(processEnvironment(), "FIREWEED_OVERRIDE=first", "spare override sentinel")
	envA := envBacking[:len(envBacking)-1]
	argvBefore, envInputBefore := append([]string(nil), argvBacking...), append([]string(nil), envBacking...)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	first, second := make(chan g.CommandResult, 1), make(chan g.CommandResult, 1)
	go func() { first <- g.Run(ctx, root, argvA, envA) }()
	go func() {
		second <- boundedProcess(t, root, helperCommand(t, "hold", secondReady, "second output"), append(processEnvironment(), "FIREWEED_OVERRIDE=second"))
	}()
	waitProcessMarkers(t, []string{firstReady, secondReady}, first)
	cancel()
	a, b := <-first, <-second
	if a.Exit != -1 || !errors.Is(a.Err, context.Canceled) || b.Exit != 0 || b.Err != nil || !strings.Contains(a.Output, "first output") || strings.Contains(a.Output, "second output") || !strings.Contains(b.Output, "second output") || strings.Contains(b.Output, "first output") {
		t.Fatalf("overlapping calls mixed %+v %+v", a, b)
	}
	afterCWD, _ := os.Getwd()
	if cwd != afterCWD || !reflect.DeepEqual(envBefore, os.Environ()) || !reflect.DeepEqual(argvBefore, argvBacking) || !reflect.DeepEqual(envInputBefore, envBacking) {
		t.Fatal("overlapping calls changed caller state or spare storage")
	}
}
