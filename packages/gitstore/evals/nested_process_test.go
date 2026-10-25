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
	"testing"
	"time"

	g "github.com/lurkshark/fireweed/packages/gitstore"
)

const nestedHelperName = "Test_GIT006_GIT010NestedProcessHelper"

type nestedReady struct {
	PID, Group int
	Label      string
	Root       string
	owned      *nestedOwned
}

type nestedOwned struct {
	alive func() (bool, error)
	kill  func() error
	close func()
}

type nestedOutcome struct {
	Out, Stderr string
	Exit        int
	Err         error
}

func nestedSupported(t *testing.T) {
	t.Helper()
	if runtime.GOOS != "darwin" && runtime.GOOS != "linux" && runtime.GOOS != "windows" {
		t.Skip("owned-tree regression targets Linux/macOS/Windows")
	}
}

func nestedArgs(stage, root, label, inner string) []string {
	exe, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return []string{exe, "-test.run=^" + nestedHelperName + "$", "--", stage, root, label, inner}
}

func nestedEnv() []string {
	return []string{"FIREWEED_NESTED_EVAL=1", "GORACE=atexit_sleep_ms=0"}
}

func nestedFile(root, label, suffix string) string {
	return filepath.Join(root, label+"."+suffix)
}

func nestedOut(label string) string { return label + "-out\x00\xff\n" }
func nestedErr(label string) string { return label + "-err\x00\x80\n" }

func nestedCall(kind string, ctx context.Context, root string, argv, env []string) nestedOutcome {
	if kind == "Run" {
		r := g.Run(ctx, root, argv, env)
		return nestedOutcome{Out: r.Output, Exit: r.Exit, Err: r.Err}
	}
	r := g.Capture(ctx, root, argv, env)
	return nestedOutcome{Out: r.Stdout, Stderr: r.Stderr, Exit: r.Exit, Err: r.Err}
}

func nestedSignalReady(root, label string) {
	value := nestedReady{PID: os.Getpid(), Group: nestedSelfGroup(), Root: root, Label: label}
	raw, err := json.Marshal(value)
	if err != nil {
		os.Exit(90)
	}
	f, err := os.CreateTemp(root, ".ready-")
	if err != nil {
		os.Exit(91)
	}
	if _, err = f.Write(raw); err != nil {
		os.Exit(92)
	}
	if err = f.Close(); err != nil {
		os.Exit(93)
	}
	if err = os.Rename(f.Name(), nestedFile(root, label, "ready")); err != nil {
		os.Exit(94)
	}
}

func Test_GIT006_GIT010NestedProcessHelper(t *testing.T) {
	if os.Getenv("FIREWEED_NESTED_EVAL") != "1" {
		return
	}
	var args []string
	for i, arg := range os.Args {
		if arg == "--" {
			args = os.Args[i+1:]
			break
		}
	}
	if len(args) != 4 {
		os.Exit(95)
	}
	stage, root, label, inner := args[0], args[1], args[2], args[3]
	fmt.Fprint(os.Stdout, nestedOut(label))
	fmt.Fprint(os.Stderr, nestedErr(label))
	nestedSignalReady(root, label)
	switch stage {
	case "bridge":
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		result := nestedCall(inner, ctx, root, nestedArgs("inner", root, "inner", inner), nestedEnv())
		cancel()
		if result.Err != nil || result.Exit != 0 {
			os.Exit(96)
		}
		fmt.Fprint(os.Stdout, result.Out)
		fmt.Fprint(os.Stderr, result.Stderr)
	case "inner":
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()
		a := nestedArgs("leaf", root, "leaf", inner)
		cmd := exec.CommandContext(ctx, a[0], a[1:]...)
		cmd.Dir = root
		cmd.Env = append(os.Environ(), nestedEnv()...)
		cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
		nestedSeparateGroup(cmd)
		if err := cmd.Run(); err != nil {
			os.Exit(97)
		}
	case "leaf":
		// The outer caller arms this only after its adapter has returned.
		// Write the marker before any output: a closed inherited pipe must
		// not turn a surviving child into a false negative through SIGPIPE.
		limit := time.Now().Add(15 * time.Second)
		for {
			if _, err := os.Stat(nestedFile(root, label, "arm")); err == nil {
				time.Sleep(200 * time.Millisecond)
				if err := os.WriteFile(nestedFile(root, label, "late"), []byte(label), 0600); err != nil {
					os.Exit(98)
				}
				break
			}
			if time.Now().After(limit) {
				os.Exit(99)
			}
			time.Sleep(5 * time.Millisecond)
		}
	default:
		os.Exit(100)
	}
	os.Exit(0)
}

func nestedArm(root string, labels ...string) {
	for _, label := range labels {
		_ = os.WriteFile(nestedFile(root, label, "arm"), []byte(label), 0600)
	}
}

func nestedCleanup(t *testing.T, root string, cancel ...context.CancelFunc) {
	t.Helper()
	t.Cleanup(func() {
		for _, stop := range cancel {
			stop()
		}
		nestedArm(root, "bridge", "inner", "leaf", "outside", "overlap", "control")
	})
}

func nestedWaitReady(t *testing.T, root string, labels []string, done <-chan nestedOutcome) map[string]nestedReady {
	t.Helper()
	end := time.Now().Add(3 * time.Second)
	for time.Now().Before(end) {
		found := map[string]nestedReady{}
		for _, label := range labels {
			raw, err := os.ReadFile(nestedFile(root, label, "ready"))
			var ready nestedReady
			if err == nil && json.Unmarshal(raw, &ready) == nil && ready.PID > 0 && ready.Root == root && ready.Label == label {
				found[label] = ready
			}
		}
		if len(found) == len(labels) {
			seen := map[int]bool{}
			for _, r := range found {
				if seen[r.PID] {
					t.Fatal("readiness did not identify distinct helper processes", found)
				}
				seen[r.PID] = true
			}
			if a, ok := found["inner"]; ok && runtime.GOOS != "windows" {
				if b, ok := found["leaf"]; ok && (a.Group <= 0 || b.Group != b.PID || a.Group == b.Group) {
					t.Fatal("independent native descendant group premise absent", found)
				}
			}
			for label, ready := range found {
				owned, err := nestedPin(ready.PID, root)
				if err != nil {
					t.Fatal("pin ready process instance", label, err)
				}
				t.Cleanup(func() {
					defer owned.close()
					alive, err := owned.alive()
					if err == nil && alive {
						err = owned.kill()
					}
					if err != nil {
						t.Errorf("owned helper cleanup %s: %v", label, err)
					}
				})
				alive, err := owned.alive()
				if err != nil || !alive {
					t.Fatal("ready helper not live before cancellation", label, alive, err)
				}
				ready.owned = owned
				found[label] = ready
			}
			return found
		}
		select {
		case result := <-done:
			t.Fatalf("helper ended before readiness: %+v", result)
		default:
		}
		time.Sleep(5 * time.Millisecond)
	}
	t.Fatal("helper readiness timed out")
	return nil
}

func nestedWaitOutcome(t *testing.T, done <-chan nestedOutcome, limit time.Duration) nestedOutcome {
	t.Helper()
	select {
	case result := <-done:
		return result
	case <-time.After(limit):
		t.Fatal("bounded process outcome did not arrive")
		return nestedOutcome{}
	}
}

func nestedNative(t *testing.T, root, stage, label, inner string) (<-chan nestedOutcome, context.CancelFunc) {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 12*time.Second)
	a := nestedArgs(stage, root, label, inner)
	cmd := exec.CommandContext(ctx, a[0], a[1:]...)
	cmd.Dir = root
	cmd.Env = append(os.Environ(), nestedEnv()...)
	var out, stderr bytes.Buffer
	cmd.Stdout, cmd.Stderr = &out, &stderr
	done := make(chan nestedOutcome, 1)
	if err := cmd.Start(); err != nil {
		cancel()
		t.Fatal("native helper launch", err)
	}
	go func() {
		err := cmd.Wait()
		code := 0
		if err != nil {
			code = -1
			var exit *exec.ExitError
			if errors.As(err, &exit) {
				code = exit.ExitCode()
			}
		}
		done <- nestedOutcome{Out: out.String(), Stderr: stderr.String(), Exit: code, Err: err}
	}()
	return done, cancel
}

func nestedExpectLate(t *testing.T, root, label string, want bool) {
	t.Helper()
	raw, err := os.ReadFile(nestedFile(root, label, "late"))
	if want {
		if err != nil || string(raw) != label {
			t.Errorf("live delayed-write control %s missing: %q %v", label, raw, err)
		}
	} else if !os.IsNotExist(err) {
		t.Errorf("descendant %s survived cancellation and wrote after return: %q %v", label, raw, err)
	}
}

func Test_GIT006_GIT010NestedProcessControls(t *testing.T) {
	nestedSupported(t)
	for _, inner := range []string{"Run", "Capture"} {
		t.Run(inner, func(t *testing.T) {
			root := t.TempDir()
			done, cancel := nestedNative(t, root, "bridge", "bridge", inner)
			nestedCleanup(t, root, cancel)
			nestedWaitReady(t, root, []string{"bridge", "inner", "leaf"}, done)
			select {
			case result := <-done:
				t.Fatal("unarmed delayed-write helper ended", result)
			default:
			}
			nestedExpectLate(t, root, "leaf", false)
			nestedArm(root, "leaf")
			result := nestedWaitOutcome(t, done, 3*time.Second)
			if result.Exit != 0 || result.Err != nil {
				t.Fatal("native nested control failed", result)
			}
			for _, label := range []string{"bridge", "inner", "leaf"} {
				if !strings.Contains(result.Out+result.Stderr, nestedOut(label)) || !strings.Contains(result.Out+result.Stderr, nestedErr(label)) {
					t.Error("nested control lost output", label, result)
				}
			}
			nestedExpectLate(t, root, "leaf", true)
		})
	}
}

func nestedCancellation(t *testing.T, kind string) {
	nestedSupported(t)
	for _, termination := range []string{"cancel", "deadline"} {
		t.Run(termination, func(t *testing.T) {
			for _, inner := range []string{"Run", "Capture"} {
				t.Run(inner, func(t *testing.T) {
					root := t.TempDir()
					cwd, _ := os.Getwd()
					parent := os.Environ()
					outside, stopOutside := nestedNative(t, root, "leaf", "outside", inner)
					overlapCtx, stopOverlap := context.WithTimeout(context.Background(), 12*time.Second)
					overlap := make(chan nestedOutcome, 1)
					go func() {
						overlap <- nestedCall("Capture", overlapCtx, root, nestedArgs("leaf", root, "overlap", inner), nestedEnv())
					}()
					var ctx context.Context
					var cancel context.CancelFunc
					if termination == "deadline" {
						ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
					} else {
						ctx, cancel = context.WithCancel(context.Background())
					}
					nestedCleanup(t, root, cancel, stopOverlap, stopOutside)
					argv, env := nestedArgs("bridge", root, "bridge", inner), nestedEnv()
					originalArgv, originalEnv := append([]string(nil), argv...), append([]string(nil), env...)
					done := make(chan nestedOutcome, 1)
					go func() { done <- nestedCall(kind, ctx, root, argv, env) }()
					ready := nestedWaitReady(t, root, []string{"bridge", "inner", "leaf", "outside", "overlap"}, done)
					if ctx.Err() != nil {
						t.Fatal("context expired before meaningful readiness", ctx.Err())
					}
					select {
					case result := <-done:
						t.Fatal("outer call returned before cancellation", result)
					default:
					}
					if termination == "cancel" {
						cancel()
					} else {
						select {
						case <-ctx.Done():
						case <-time.After(6 * time.Second):
							t.Fatal("deadline control did not expire")
						}
					}
					returnStart := time.Now()
					result := nestedWaitOutcome(t, done, 2*time.Second)
					if time.Since(returnStart) > 2*time.Second || result.Exit != -1 || !errors.Is(result.Err, ctx.Err()) {
						t.Error("owned tree context outcome", result, ctx.Err())
					}
					if kind == "Run" {
						if strings.Count(result.Out, nestedOut("bridge")) != 1 || strings.Count(result.Out, nestedErr("bridge")) != 1 {
							t.Error("Run lost pre-cancellation bridge streams", result)
						}
					} else if result.Out != nestedOut("bridge") || result.Stderr != nestedErr("bridge") {
						t.Error("Capture lost or mixed pre-cancellation bridge streams", result)
					}
					for label, channel := range map[string]<-chan nestedOutcome{"outside": outside, "overlap": overlap} {
						select {
						case other := <-channel:
							t.Errorf("unrelated %s ended at cancellation: %+v", label, other)
						default:
						}
					}
					nestedArm(root, "leaf", "outside", "overlap")
					for label, channel := range map[string]<-chan nestedOutcome{"outside": outside, "overlap": overlap} {
						other := nestedWaitOutcome(t, channel, 3*time.Second)
						if other.Exit != 0 || other.Err != nil || other.Out != nestedOut(label) || other.Stderr != nestedErr(label) {
							t.Errorf("unrelated/overlap %s did not survive independently: %+v", label, other)
						}
						nestedExpectLate(t, root, label, true)
					}
					time.Sleep(300 * time.Millisecond)
					nestedExpectLate(t, root, "leaf", false)
					// A stopped survivor cannot pass merely by not writing.
					until := time.Now().Add(time.Second)
					for _, label := range []string{"bridge", "inner", "leaf"} {
						for {
							alive, err := ready[label].owned.alive()
							if err != nil {
								t.Fatal("native process-state observation", err)
							}
							if !alive {
								break
							}
							if time.Now().After(until) {
								t.Errorf("owned descendant %s remains live after cancellation", label)
								break
							}
							time.Sleep(10 * time.Millisecond)
						}
					}
					after, _ := os.Getwd()
					if cwd != after || !reflect.DeepEqual(parent, os.Environ()) || !reflect.DeepEqual(argv, originalArgv) || !reflect.DeepEqual(env, originalEnv) {
						t.Error("cancellation changed caller state or inputs")
					}
					afterCtx, stopAfter := context.WithTimeout(context.Background(), 3*time.Second)
					later := nestedCall(kind, afterCtx, root, helperCommand(t, "exit"), processEnvironment())
					stopAfter()
					if later.Exit != 37 || later.Err == nil || errors.Is(later.Err, context.Canceled) || errors.Is(later.Err, context.DeadlineExceeded) {
						t.Error("cancelled call poisoned later ordinary exit", later)
					}
				})
			}
		})
	}
}

func Test_GIT006NestedRunCancellation(t *testing.T) {
	nestedCancellation(t, "Run")
}

func Test_GIT010NestedCaptureCancellation(t *testing.T) {
	nestedCancellation(t, "Capture")
}
