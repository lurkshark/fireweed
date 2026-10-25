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
	"strconv"
	"testing"
	"time"
)

const vanishedHelperName = "Test_GIT006_GIT010VanishedAncestorHelper"

func vanishedArgs(root, label, mode string, code int) []string {
	a := orphanArgs("ancestor", root, label, mode, code)
	a[1] = "-test.run=^" + vanishedHelperName + "$"
	return a
}
func vanishedEnv() []string { return append(orphanEnv(), "FIREWEED_VANISHED_EVAL=1") }

func Test_GIT006_GIT010VanishedAncestorHelper(t *testing.T) {
	if os.Getenv("FIREWEED_VANISHED_EVAL") != "1" {
		return
	}
	var args []string
	for i, arg := range os.Args {
		if arg == "--" {
			args = os.Args[i+1:]
			break
		}
	}
	if len(args) != 6 || args[0] != "ancestor" || args[5] != orphanLiteral || os.Getenv("FW_ORPHAN_MARKER") != "marker-λ=x" {
		os.Exit(130)
	}
	if value, ok := os.LookupEnv("FW_ORPHAN_OPTION"); !ok || value != "" {
		os.Exit(131)
	}
	root, label, mode := args[1], args[2], args[3]
	code, err := strconv.Atoi(args[4])
	if err != nil {
		os.Exit(132)
	}
	cwd, err := os.Getwd()
	if err != nil {
		os.Exit(133)
	}
	actual, err := filepath.EvalSymlinks(cwd)
	if err != nil {
		os.Exit(134)
	}
	want, err := filepath.EvalSymlinks(root)
	if err != nil || actual != want {
		os.Exit(135)
	}
	fmt.Fprint(os.Stdout, nestedOut(label+"-ancestor"))
	fmt.Fprint(os.Stderr, nestedErr(label+"-ancestor"))
	orphanWriteReady(root, label+"-ancestor")
	a := orphanArgs("bridge", root, label, mode, code)
	command := exec.Command(a[0], a[1:]...)
	command.Dir, command.Env = root, append(os.Environ(), orphanEnv()...)
	command.Stdout, command.Stderr = os.Stdout, os.Stderr
	orphanSeparate(command, "group")
	if err := command.Start(); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(136)
	}
	err = command.Wait() // File streams avoid a Go copier waiting for the leaf.
	if err != nil {
		var exited *exec.ExitError
		if !errors.As(err, &exited) || exited.ExitCode() != code {
			fmt.Fprint(os.Stderr, err)
			os.Exit(137)
		}
	} else if code != 0 {
		os.Exit(138)
	}
	if err := os.WriteFile(nestedFile(root, label, "ancestor-reaped-middle"), []byte(label), 0600); err != nil {
		os.Exit(139)
	}
	os.Exit(code)
}

type vanishedCall struct {
	done    <-chan nestedOutcome
	settled <-chan struct{}
}

func vanishedLaunch(t *testing.T, kind string, ctx context.Context, cancel context.CancelFunc, root string, argv, env []string) vanishedCall {
	t.Helper()
	done, settled := make(chan nestedOutcome, 1), make(chan struct{})
	if kind == "Native" {
		command := exec.CommandContext(ctx, argv[0], argv[1:]...)
		command.Dir, command.Env = root, append(os.Environ(), env...)
		var out, stderr bytes.Buffer
		command.Stdout, command.Stderr = &out, &stderr
		if err := command.Start(); err != nil {
			cancel()
			t.Fatal("native ancestor launch", err)
		}
		go func() {
			defer close(settled)
			err := command.Wait()
			code := 0
			if err != nil {
				code = -1
				var exited *exec.ExitError
				if errors.As(err, &exited) {
					code = exited.ExitCode()
				}
			}
			done <- nestedOutcome{Out: out.String(), Stderr: stderr.String(), Exit: code, Err: err}
		}()
	} else {
		go func() { defer close(settled); done <- nestedCall(kind, ctx, root, argv, env) }()
	}
	// This completion channel is independent of consuming the result. On every
	// failure path, release/arm helpers before waiting for inherited streams and
	// before TempDir cleanup. Later per-PID callbacks retain exact owned cleanup.
	t.Cleanup(func() {
		for _, label := range []string{"subject", "outside", "control"} {
			_ = os.WriteFile(nestedFile(root, label, "exit-parent"), []byte("cleanup"), 0600)
		}
		nestedArm(root, "subject-leaf", "outside-leaf", "control-leaf", "overlap")
		cancel()
		select {
		case <-settled:
		case <-time.After(3 * time.Second):
			t.Error("bounded ancestor call did not finish before temporary cleanup", kind)
		}
	})
	return vanishedCall{done, settled}
}

func vanishedAncestorReady(t *testing.T, root, label string, done <-chan nestedOutcome) orphanReady {
	t.Helper()
	end := time.Now().Add(3 * time.Second)
	for time.Now().Before(end) {
		var ready orphanReady
		raw, err := os.ReadFile(nestedFile(root, label+"-ancestor", "ready"))
		if err == nil && json.Unmarshal(raw, &ready) == nil && ready.PID > 0 && ready.Root == root && ready.Label == label+"-ancestor" {
			owned, err := vanishedPinAncestor(ready.PID, root, label)
			if err != nil {
				t.Fatal("pin ancestor immediately on readiness", err)
			}
			ready.owned = owned
			t.Cleanup(func() {
				defer owned.close()
				live, err := owned.alive()
				if err == nil && live {
					err = owned.kill()
				}
				if err != nil {
					t.Error("exact ancestor cleanup", err)
				}
			})
			return ready
		}
		select {
		case result := <-done:
			t.Fatal("call returned before ancestor readiness", result)
		default:
		}
		time.Sleep(5 * time.Millisecond)
	}
	t.Fatal("bounded ancestor readiness missing")
	return orphanReady{}
}

func vanishedProve(t *testing.T, root, label, mode string, ancestor orphanReady, ready map[string]orphanReady, done <-chan nestedOutcome) {
	t.Helper()
	middle, leaf := ready[label+"-parent"], ready[label+"-leaf"]
	if ancestor.PID == middle.PID || ancestor.PID == leaf.PID || middle.Parent != ancestor.PID || leaf.Parent != middle.PID {
		t.Fatal("original native ancestor chain missing", ancestor, middle, leaf)
	}
	live, err := ancestor.owned.alive()
	if err != nil || !live {
		t.Fatal("ancestor not live before coordinated exit", live, err)
	}
	orphanProve(t, root, label, mode, ready, done)
	end := time.Now().Add(time.Second)
	for {
		rootErr, middleErr := vanishedReaped(ancestor.PID, ancestor.owned), vanishedReaped(middle.PID, middle.owned)
		marker, markerErr := os.ReadFile(nestedFile(root, label, "ancestor-reaped-middle"))
		if rootErr == nil && middleErr == nil && markerErr == nil && string(marker) == label {
			live, err := leaf.owned.alive()
			if err != nil || !live {
				t.Fatal("leaf not live after both ancestors reaped", live, err)
			}
			if err := orphanReparented(leaf.PID, middle.PID); err != nil {
				t.Fatal("native leaf reparenting missing", err)
			}
			select {
			case result := <-done:
				t.Fatal("outer returned before multiple-ancestor cancellation premise", result)
			default:
			}
			return
		}
		if time.Now().After(end) {
			t.Fatal("both ancestors were not natively reaped", rootErr, middleErr, markerErr)
		}
		time.Sleep(5 * time.Millisecond)
	}
}

func Test_GIT006_GIT010VanishedAncestorControls(t *testing.T) {
	nestedSupported(t)
	for _, mode := range orphanModes() {
		t.Run(mode, func(t *testing.T) {
			for _, kind := range []string{"Native", "Run", "Capture"} {
				t.Run(kind, func(t *testing.T) {
					for _, code := range []int{0, 37} {
						t.Run(strconv.Itoa(code), func(t *testing.T) {
							root := t.TempDir()
							ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
							call := vanishedLaunch(t, kind, ctx, cancel, root, vanishedArgs(root, "control", mode, code), vanishedEnv())
							ancestor := vanishedAncestorReady(t, root, "control", call.done)
							ready := orphanReadiness(t, root, []string{"control-parent", "control-leaf"}, call.done)
							vanishedProve(t, root, "control", mode, ancestor, ready, call.done)
							nestedExpectLate(t, root, "control-leaf", false)
							nestedArm(root, "control-leaf")
							result := nestedWaitOutcome(t, call.done, 3*time.Second)
							if result.Exit != code || (result.Err == nil) != (code == 0) || errors.Is(result.Err, context.Canceled) || errors.Is(result.Err, context.DeadlineExceeded) {
								t.Error("ordinary multi-ancestor result", result, code)
							}
							orphanStreams(t, kind, result, "control-ancestor", "control-parent", "control-leaf")
							nestedExpectLate(t, root, "control-leaf", true)
						})
					}
				})
			}
		})
	}
}

func vanishedCancellation(t *testing.T, kind string) {
	nestedSupported(t)
	for _, termination := range []string{"cancel", "deadline"} {
		t.Run(termination, func(t *testing.T) {
			for _, mode := range orphanModes() {
				t.Run(mode, func(t *testing.T) {
					root := t.TempDir()
					cwd, _ := os.Getwd()
					parentEnv := os.Environ()
					outsideCtx, stopOutside := context.WithTimeout(context.Background(), 12*time.Second)
					outside := vanishedLaunch(t, "Native", outsideCtx, stopOutside, root, vanishedArgs(root, "outside", mode, 0), vanishedEnv())
					outsideAncestor := vanishedAncestorReady(t, root, "outside", outside.done)
					overlapCtx, stopOverlap := context.WithTimeout(context.Background(), 12*time.Second)
					overlap := vanishedLaunch(t, "Capture", overlapCtx, stopOverlap, root, orphanArgs("leaf", root, "overlap", mode, 0), vanishedEnv())
					var ctx context.Context
					var cancel context.CancelFunc
					if termination == "deadline" {
						ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
					} else {
						ctx, cancel = context.WithCancel(context.Background())
					}
					argv := vanishedArgs(root, "subject", mode, 0)
					originalArgv := append([]string(nil), argv...)
					env := vanishedEnv()
					envBefore := append([]string(nil), env...)
					subject := vanishedLaunch(t, kind, ctx, cancel, root, argv, env)
					ancestor := vanishedAncestorReady(t, root, "subject", subject.done)
					ready := orphanReadiness(t, root, []string{"subject-parent", "subject-leaf", "outside-parent", "outside-leaf", "overlap"}, subject.done)
					if ancestor.PID == outsideAncestor.PID {
						t.Fatal("independent ancestor identity reused")
					}
					vanishedProve(t, root, "outside", mode, outsideAncestor, ready, outside.done)
					if termination == "deadline" {
						deadline, ok := ctx.Deadline()
						if !ok || ctx.Err() != nil {
							t.Fatal("real pending deadline missing", ctx.Err())
						}
						pause := time.Until(deadline.Add(-700 * time.Millisecond))
						if pause <= 0 {
							t.Fatal("insufficient readiness margin before actual deadline")
						}
						time.Sleep(pause)
					}
					vanishedProve(t, root, "subject", mode, ancestor, ready, subject.done)
					if ctx.Err() != nil {
						t.Fatal("context ended before both ancestors reaped", ctx.Err())
					}
					if termination == "cancel" {
						cancel()
					} else {
						select {
						case <-ctx.Done():
						case <-time.After(2 * time.Second):
							t.Fatal("real deadline did not expire")
						}
					}
					result := nestedWaitOutcome(t, subject.done, 2*time.Second)
					if result.Exit != -1 || !errors.Is(result.Err, ctx.Err()) {
						t.Error("multi-ancestor context result", result, ctx.Err())
					}
					orphanStreams(t, kind, result, "subject-ancestor", "subject-parent")
					for label, call := range map[string]vanishedCall{"outside": outside, "overlap": overlap} {
						select {
						case other := <-call.done:
							t.Error("unrelated work ended at subject cancellation", label, other)
						default:
						}
						key := label
						if label == "outside" {
							key = "outside-leaf"
						}
						live, err := ready[key].owned.alive()
						if err != nil || !live {
							t.Error("unrelated native process lost", label, live, err)
						}
					}
					// The only stimulus permitting the leaf's delayed write is after return.
					nestedArm(root, "subject-leaf", "outside-leaf", "overlap")
					other := nestedWaitOutcome(t, outside.done, 3*time.Second)
					if other.Exit != 0 || other.Err != nil {
						t.Error("unrelated native multi-orphan failed", other)
					}
					orphanStreams(t, "Native", other, "outside-ancestor", "outside-parent", "outside-leaf")
					nestedExpectLate(t, root, "outside-leaf", true)
					other = nestedWaitOutcome(t, overlap.done, 3*time.Second)
					if other.Exit != 0 || other.Err != nil {
						t.Error("overlapping public work failed", other)
					}
					orphanStreams(t, "Capture", other, "overlap")
					nestedExpectLate(t, root, "overlap", true)
					time.Sleep(300 * time.Millisecond)
					if _, err := os.Stat(nestedFile(root, "subject-leaf", "late")); err == nil {
						t.Error("descendant survived two reaped ancestors and wrote after cancellation return")
					} else if !os.IsNotExist(err) {
						t.Error("delayed-write observation failed", err)
					}
					end := time.Now().Add(time.Second)
					for {
						live, err := ready["subject-leaf"].owned.alive()
						if err != nil {
							t.Fatal("owned leaf state", err)
						}
						if !live {
							break
						}
						if time.Now().After(end) {
							t.Error("owned leaf still live after cancellation")
							break
						}
						time.Sleep(5 * time.Millisecond)
					}
					after, _ := os.Getwd()
					if cwd != after || !reflect.DeepEqual(parentEnv, os.Environ()) || !reflect.DeepEqual(argv, originalArgv) || !reflect.DeepEqual(env, envBefore) {
						t.Error("multi-ancestor cancellation changed caller state/inputs")
					}
					laterCtx, stopLater := context.WithTimeout(context.Background(), 3*time.Second)
					later := nestedCall(kind, laterCtx, root, helperCommand(t, "exit"), processEnvironment())
					stopLater()
					if later.Exit != 37 || later.Err == nil || errors.Is(later.Err, context.Canceled) || errors.Is(later.Err, context.DeadlineExceeded) {
						t.Error("multi-ancestor cancellation poisoned later ordinary exit", later)
					}
				})
			}
		})
	}
}
func Test_GIT006VanishedAncestorRunCancellation(t *testing.T)     { vanishedCancellation(t, "Run") }
func Test_GIT010VanishedAncestorCaptureCancellation(t *testing.T) { vanishedCancellation(t, "Capture") }
