package evals

import (
	"bytes"
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
	"sync"
	"testing"
	"time"

	g "github.com/lurkshark/fireweed/packages/gitstore"
)

const captureHelper = "Test_GIT008_GIT009_GIT010CaptureHelper"

func captureArgs(t *testing.T, mode string, args ...string) []string {
	t.Helper()
	exe, err := os.Executable()
	if err != nil {
		t.Fatal(err)
	}
	return append([]string{exe, "-test.run=^" + captureHelper + "$", "--", mode}, args...)
}
func captureEnv(extra ...string) []string {
	return append([]string{"FIREWEED_CAPTURE_HELPER=1", "GORACE=atexit_sleep_ms=0"}, extra...)
}
func captureBytes(stream byte) string {
	return strings.Repeat(string([]byte{stream, 0, 255, 128, '\r', '\n'}), 90000)
}
func Test_GIT008_GIT009_GIT010CaptureHelper(t *testing.T) {
	if os.Getenv("FIREWEED_CAPTURE_HELPER") != "1" {
		return
	}
	var a []string
	for i, v := range os.Args {
		if v == "--" {
			a = os.Args[i+1:]
			break
		}
	}
	if len(a) == 0 {
		os.Exit(91)
	}
	switch a[0] {
	case "streams":
		var wg sync.WaitGroup
		wg.Add(2)
		go func() { defer wg.Done(); fmt.Fprint(os.Stdout, captureBytes('O')) }()
		go func() { defer wg.Done(); fmt.Fprint(os.Stderr, captureBytes('E')) }()
		wg.Wait()
	case "echo":
		cwd, err := os.Getwd()
		if err != nil {
			os.Exit(92)
		}
		v, present := os.LookupEnv("FIREWEED_CAPTURE_EMPTY")
		b, err := json.Marshal(map[string]any{"cwd": cwd, "argv": a[1:], "inherited": os.Getenv("FIREWEED_CAPTURE_PARENT"), "override": os.Getenv("FIREWEED_CAPTURE_OVERRIDE"), "empty": v, "empty_present": present, "equals": os.Getenv("FIREWEED_CAPTURE_EQUALS")})
		if err != nil {
			os.Exit(93)
		}
		fmt.Fprint(os.Stdout, string(b))
		fmt.Fprint(os.Stderr, "warning\x00\xff\n")
	case "silent":
	case "out":
		fmt.Fprint(os.Stdout, "stdout only\x00\xff")
	case "err":
		fmt.Fprint(os.Stderr, "stderr only\x00\xff")
	case "exit":
		fmt.Fprint(os.Stdout, "before\x00\xff")
		fmt.Fprint(os.Stderr, "failure\x80\n")
		os.Exit(37)
	default:
		os.Exit(94)
	}
	os.Exit(0)
}
func boundedCapture(t *testing.T, dir string, argv, env []string) g.CaptureResult {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	defer cancel()
	cwd, _ := os.Getwd()
	parent := os.Environ()
	a, e := append([]string(nil), argv[:cap(argv)]...), append([]string(nil), env[:cap(env)]...)
	got := g.Capture(ctx, dir, argv, env)
	after, _ := os.Getwd()
	if cwd != after || !reflect.DeepEqual(parent, os.Environ()) || !reflect.DeepEqual(a, append([]string(nil), argv[:cap(argv)]...)) || !reflect.DeepEqual(e, append([]string(nil), env[:cap(env)]...)) {
		t.Fatal("capture changed caller state or inputs")
	}
	return got
}
func Test_GIT008_GIT009CaptureHelperControls(t *testing.T) {
	// This oracle does not call the subject. It proves binary/concurrent helper output.
	for _, mode := range []string{"streams", "silent", "out", "err", "exit"} {
		a := captureArgs(t, mode)
		ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
		cmd := exec.CommandContext(ctx, a[0], a[1:]...)
		cmd.Env = append(os.Environ(), captureEnv()...)
		var out, stderr bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &stderr
		err := cmd.Run()
		cancel()
		wantOut, wantErr, code := captureExpected(mode)
		if out.String() != wantOut || stderr.String() != wantErr || cmd.ProcessState == nil || cmd.ProcessState.ExitCode() != code || (err == nil) != (code == 0) {
			t.Fatalf("helper %s output lengths %d/%d exit %v: %v", mode, out.Len(), stderr.Len(), cmd.ProcessState, err)
		}
	}
}
func captureExpected(mode string) (string, string, int) {
	switch mode {
	case "streams":
		return captureBytes('O'), captureBytes('E'), 0
	case "out":
		return "stdout only\x00\xff", "", 0
	case "err":
		return "", "stderr only\x00\xff", 0
	case "exit":
		return "before\x00\xff", "failure\x80\n", 37
	default:
		return "", "", 0
	}
}
func Test_GIT008_GIT009CaptureExactStreamsAndExits(t *testing.T) {
	for _, mode := range []string{"streams", "silent", "out", "err", "exit"} {
		t.Run(mode, func(t *testing.T) {
			got := boundedCapture(t, t.TempDir(), captureArgs(t, mode), captureEnv())
			out, stderr, code := captureExpected(mode)
			if got.Stdout != out || got.Stderr != stderr || got.Exit != code || (got.Err == nil) != (code == 0) {
				t.Fatalf("capture %s: stdout %d/%d stderr %d/%d exit %d/%d err %v", mode, len(got.Stdout), len(out), len(got.Stderr), len(stderr), got.Exit, code, got.Err)
			}
			if errors.Is(got.Err, context.Canceled) || errors.Is(got.Err, context.DeadlineExceeded) {
				t.Fatal("ordinary exit mislabeled context termination")
			}
		})
	}
}
func Test_GIT008CaptureLiteralEnvironmentAndInputs(t *testing.T) {
	t.Setenv("FIREWEED_CAPTURE_PARENT", "inherited")
	t.Setenv("FIREWEED_CAPTURE_OVERRIDE", "parent")
	t.Setenv("FIREWEED_CAPTURE_EMPTY", "parent")
	dir := filepath.Join(t.TempDir(), "space 'quote' λ")
	if err := os.Mkdir(dir, 0700); err != nil {
		t.Fatal(err)
	}
	dir, err := filepath.EvalSymlinks(dir)
	if err != nil {
		t.Fatal(err)
	}
	marker := filepath.Join(dir, "expanded")
	args := []string{"", "two words", "'quote'", `"double"`, "$FIREWEED_CAPTURE_PARENT", "$(touch " + marker + ")", "`touch " + marker + "`", "; touch " + marker, "*.go", "tab\tline\nλ"}
	backing := append(captureArgs(t, "echo", args...), "spare argv")
	envBacking := append(captureEnv("FIREWEED_CAPTURE_OVERRIDE=first", "FIREWEED_CAPTURE_OVERRIDE=last", "FIREWEED_CAPTURE_EMPTY=", "FIREWEED_CAPTURE_EQUALS=a=b=c"), "spare env")
	beforeA, beforeE := append([]string(nil), backing...), append([]string(nil), envBacking...)
	cwd, _ := os.Getwd()
	env := os.Environ()
	got := boundedCapture(t, dir, backing[:len(backing)-1], envBacking[:len(envBacking)-1])
	if got.Exit != 0 || got.Err != nil || got.Stderr != "warning\x00\xff\n" {
		t.Fatalf("literal capture %+v", got)
	}
	var data struct {
		Cwd                                string
		Argv                               []string
		Inherited, Override, Empty, Equals string
		EmptyPresent                       bool `json:"empty_present"`
	}
	if err := json.Unmarshal([]byte(got.Stdout), &data); err != nil {
		t.Fatal(err)
	}
	if data.Cwd != dir || !reflect.DeepEqual(data.Argv, args) || data.Inherited != "inherited" || data.Override != "last" || data.Empty != "" || !data.EmptyPresent || data.Equals != "a=b=c" {
		t.Fatalf("child data %+v", data)
	}
	after, _ := os.Getwd()
	if cwd != after || !reflect.DeepEqual(env, os.Environ()) || !reflect.DeepEqual(backing, beforeA) || !reflect.DeepEqual(envBacking, beforeE) {
		t.Fatal("caller state/backing storage changed")
	}
	if _, err := os.Stat(marker); !os.IsNotExist(err) {
		t.Fatal("shell interpolation executed")
	}
	if runtime.GOOS == "windows" {
		got = boundedCapture(t, dir, captureArgs(t, "echo"), captureEnv("fireweed_capture_override=lower", "FIREWEED_CAPTURE_OVERRIDE=upper"))
		if got.Exit != 0 || got.Err != nil || !strings.Contains(got.Stdout, `"override":"upper"`) {
			t.Fatalf("Windows keys %+v", got)
		}
	}
	// Run still combines exact sequential streams on an ordinary failing process.
	r := boundedProcess(t, dir, captureArgs(t, "exit"), captureEnv())
	if r.Exit != 37 || r.Err == nil || r.Output != "before\x00\xfffailure\x80\n" {
		t.Fatalf("Run changed %+v", r)
	}
}
func Test_GIT009CaptureInvalidPrecancelAndRecovery(t *testing.T) {
	root := t.TempDir()
	marker := filepath.Join(root, "started")
	good := helperCommand(t, "marker", marker)
	cancelled, cancel := context.WithCancel(context.Background())
	cancel()
	expired, finish := context.WithDeadline(context.Background(), time.Now().Add(-time.Second))
	defer finish()
	for _, tt := range []struct {
		name, dir string
		ctx       context.Context
		argv, env []string
		code      int
		cause     error
	}{
		{"nil context", root, nil, good, processEnvironment(), 2, nil},
		{"nil argv", root, context.Background(), nil, processEnvironment(), 2, nil},
		{"empty argv", root, context.Background(), []string{}, processEnvironment(), 2, nil},
		{"empty executable", root, context.Background(), []string{"", marker}, processEnvironment(), 2, nil},
		{"bad directory", filepath.Join(root, "missing"), context.Background(), good, processEnvironment(), 2, nil},
		{"launch failure", root, context.Background(), []string{filepath.Join(root, "missing-exe"), marker}, processEnvironment(), 2, nil},
		{"override no equals", root, context.Background(), good, append(processEnvironment(), "INVALID"), 2, nil},
		{"override no key", root, context.Background(), good, append(processEnvironment(), "=bad"), 2, nil},
		{"override NUL", root, context.Background(), good, append(processEnvironment(), "KEY=bad\x00value"), 2, nil},
		{"cancelled", root, cancelled, good, processEnvironment(), -1, context.Canceled},
		{"expired", root, expired, good, processEnvironment(), -1, context.DeadlineExceeded},
	} {
		t.Run(tt.name, func(t *testing.T) {
			a, e := append([]string(nil), tt.argv...), append([]string(nil), tt.env...)
			cwd, _ := os.Getwd()
			env := os.Environ()
			got := g.Capture(tt.ctx, tt.dir, tt.argv, tt.env)
			if got.Exit != tt.code || got.Err == nil || strings.TrimSpace(got.Err.Error()) == "" || got.Stdout != "" || got.Stderr != "" {
				t.Fatalf("invalid outcome %+v", got)
			}
			if tt.cause != nil && !errors.Is(got.Err, tt.cause) {
				t.Fatalf("lost cause %v", got.Err)
			}
			if tt.cause == nil && (errors.Is(got.Err, context.Canceled) || errors.Is(got.Err, context.DeadlineExceeded)) {
				t.Fatal("invalid input confused with cancellation")
			}
			after, _ := os.Getwd()
			if cwd != after || !reflect.DeepEqual(env, os.Environ()) || !reflect.DeepEqual(a, append([]string(nil), tt.argv...)) || !reflect.DeepEqual(e, append([]string(nil), tt.env...)) {
				t.Fatal("invalid call mutated caller")
			}
			if _, err := os.Stat(marker); !os.IsNotExist(err) {
				t.Fatal("invalid call launched helper")
			}
			r := boundedCapture(t, root, captureArgs(t, "out"), captureEnv())
			if r.Exit != 0 || r.Err != nil || r.Stdout != "stdout only\x00\xff" || r.Stderr != "" {
				t.Fatalf("later valid call poisoned %+v", r)
			}
		})
	}
}

func waitCaptureReady(t *testing.T, paths []string, done <-chan g.CaptureResult) {
	t.Helper()
	until := time.Now().Add(1500 * time.Millisecond)
	for time.Now().Before(until) {
		ready := true
		for _, p := range paths {
			if _, err := os.Stat(p); err != nil {
				ready = false
			}
		}
		if ready {
			return
		}
		select {
		case result := <-done:
			t.Fatalf("capture ended before helper readiness: %+v", result)
		default:
		}
		time.Sleep(5 * time.Millisecond)
	}
	t.Fatal("capture helpers did not become ready")
}
func Test_GIT010CaptureTreeAndOverlappingOwnership(t *testing.T) {
	if runtime.GOOS != "darwin" && runtime.GOOS != "linux" && runtime.GOOS != "windows" {
		t.Skip("supported process-tree platforms")
	}
	for _, deadline := range []bool{false, true} {
		t.Run(fmt.Sprintf("deadline_%v", deadline), func(t *testing.T) {
			root := t.TempDir()
			cwd, _ := os.Getwd()
			parentEnv := os.Environ()
			argvBacking := append(helperCommand(t, "tree", root), "spare argv")
			envBacking := append(processEnvironment(), "spare env")
			savedA, savedE := append([]string(nil), argvBacking...), append([]string(nil), envBacking...)
			var ctx context.Context
			var cancel context.CancelFunc
			if deadline {
				ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
			} else {
				ctx, cancel = context.WithCancel(context.Background())
			}
			defer cancel()
			first := make(chan g.CaptureResult, 1)
			go func() {
				first <- g.Capture(ctx, root, argvBacking[:len(argvBacking)-1], envBacking[:len(envBacking)-1])
			}()
			waitCaptureReady(t, []string{filepath.Join(root, "tree.ready"), filepath.Join(root, "branch.ready"), filepath.Join(root, "leaf.ready")}, first)
			readyAt := time.Now()
			// Two independently owned live subject calls must both survive this cancellation.
			captureReady, runReady := filepath.Join(root, "capture.ready"), filepath.Join(root, "run.ready")
			peer := make(chan g.CaptureResult, 1)
			run := make(chan g.CommandResult, 1)
			peerCtx, peerCancel := context.WithTimeout(context.Background(), 6*time.Second)
			defer peerCancel()
			go func() {
				peer <- g.Capture(peerCtx, root, helperCommand(t, "hold", captureReady, "peer capture", "3s"), processEnvironment())
			}()
			go func() {
				run <- g.Run(peerCtx, root, helperCommand(t, "hold", runReady, "peer run", "3s"), processEnvironment())
			}()
			waitCaptureReady(t, []string{captureReady, runReady}, peer)
			if !deadline {
				cancel()
			}
			var got g.CaptureResult
			select {
			case got = <-first:
			case <-time.After(2800 * time.Millisecond):
				t.Fatal("capture failed to bound tree/held streams")
			}
			if time.Since(readyAt) > 2600*time.Millisecond || got.Exit != -1 || ctx.Err() == nil || !errors.Is(got.Err, ctx.Err()) {
				t.Fatalf("termination %+v cause %v", got, ctx.Err())
			}
			if (got.Stdout != "tree ready\nbranch ready\n" && got.Stdout != "branch ready\ntree ready\n") || got.Stderr != "leaf ready\n" {
				t.Fatalf("pretermination streams %+v", got)
			}
			select {
			case p := <-peer:
				t.Fatalf("unrelated Capture ended early %+v", p)
			default:
			}
			select {
			case r := <-run:
				t.Fatalf("unrelated Run ended early %+v", r)
			default:
			}
			p, r := <-peer, <-run
			if p.Exit != 0 || p.Err != nil || p.Stdout != "peer capture\n" || p.Stderr != "held end\n" || r.Exit != 0 || r.Err != nil || r.Output != "peer run\nheld end\n" {
				t.Fatalf("peers mixed or killed %+v %+v", p, r)
			}
			if wait := time.Until(readyAt.Add(3300 * time.Millisecond)); wait > 0 {
				time.Sleep(wait)
			}
			for _, name := range []string{"tree.late", "branch.late", "leaf.late"} {
				if _, err := os.Stat(filepath.Join(root, name)); !os.IsNotExist(err) {
					t.Fatalf("descendant survived: %s", name)
				}
			}
			after, _ := os.Getwd()
			if cwd != after || !reflect.DeepEqual(parentEnv, os.Environ()) || !reflect.DeepEqual(savedA, argvBacking) || !reflect.DeepEqual(savedE, envBacking) {
				t.Fatal("cancelled/concurrent calls mutated caller")
			}
			later := boundedCapture(t, root, captureArgs(t, "exit"), captureEnv())
			if later.Exit != 37 || later.Err == nil || later.Stdout != "before\x00\xff" || later.Stderr != "failure\x80\n" {
				t.Fatalf("post-cancellation call poisoned %+v", later)
			}
		})
	}
}
