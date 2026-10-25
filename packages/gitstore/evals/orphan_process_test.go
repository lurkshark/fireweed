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
	"strings"
	"testing"
	"time"
)

const orphanHelperName = "Test_GIT006_GIT010OrphanProcessHelper"
const orphanLiteral = "literal $HOME ' ; λ"

type orphanReady struct {
	PID, Parent, Group, Session int
	Root, Label                 string
	owned                       *nestedOwned
}

func orphanArgs(stage, root, label, mode string, code int) []string {
	exe, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return []string{exe, "-test.run=^" + orphanHelperName + "$", "--", stage, root, label, mode, strconv.Itoa(code), orphanLiteral}
}

func orphanEnv() []string {
	return []string{"FIREWEED_ORPHAN_EVAL=1", "GORACE=atexit_sleep_ms=0", "FW_ORPHAN_MARKER=marker-λ=x", "FW_ORPHAN_OPTION=first", "FW_ORPHAN_OPTION="}
}

func orphanWriteReady(root, label string) {
	session, err := orphanSession()
	if err != nil {
		panic(err)
	}
	value := orphanReady{PID: os.Getpid(), Parent: os.Getppid(), Group: nestedSelfGroup(), Session: session, Root: root, Label: label}
	raw, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	f, err := os.CreateTemp(root, ".orphan-ready-")
	if err != nil {
		panic(err)
	}
	if _, err = f.Write(raw); err != nil {
		panic(err)
	}
	if err = f.Close(); err != nil {
		panic(err)
	}
	if err = os.Rename(f.Name(), nestedFile(root, label, "ready")); err != nil {
		panic(err)
	}
}

func orphanWaitFile(path string, limit time.Duration) bool {
	end := time.Now().Add(limit)
	for time.Now().Before(end) {
		if _, err := os.Stat(path); err == nil {
			return true
		}
		time.Sleep(5 * time.Millisecond)
	}
	return false
}

func Test_GIT006_GIT010OrphanProcessHelper(t *testing.T) {
	if os.Getenv("FIREWEED_ORPHAN_EVAL") != "1" {
		return
	}
	var args []string
	for i, arg := range os.Args {
		if arg == "--" {
			args = os.Args[i+1:]
			break
		}
	}
	if len(args) != 6 || args[5] != orphanLiteral || os.Getenv("FW_ORPHAN_MARKER") != "marker-λ=x" {
		os.Exit(110)
	}
	if value, ok := os.LookupEnv("FW_ORPHAN_OPTION"); !ok || value != "" {
		os.Exit(111)
	}
	stage, root, label, mode := args[0], args[1], args[2], args[3]
	code, err := strconv.Atoi(args[4])
	if err != nil {
		os.Exit(112)
	}
	cwd, err := os.Getwd()
	if err != nil {
		os.Exit(113)
	}
	actual, err := filepath.EvalSymlinks(cwd)
	if err != nil {
		os.Exit(114)
	}
	want, err := filepath.EvalSymlinks(root)
	if err != nil || actual != want {
		os.Exit(115)
	}
	if stage == "bridge" {
		fmt.Fprint(os.Stdout, nestedOut(label+"-parent"))
		fmt.Fprint(os.Stderr, nestedErr(label+"-parent"))
		orphanWriteReady(root, label+"-parent")
		a := orphanArgs("leaf", root, label+"-leaf", mode, 0)
		cmd := exec.Command(a[0], a[1:]...)
		cmd.Dir, cmd.Env = root, append(os.Environ(), orphanEnv()...)
		cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
		orphanSeparate(cmd, mode)
		if err := cmd.Start(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(116)
		}
		if !orphanWaitFile(nestedFile(root, label, "exit-parent"), 15*time.Second) {
			os.Exit(117)
		}
		os.Exit(code) // Deliberately leave the ready leaf alive with both inherited streams.
	}
	if stage != "leaf" {
		os.Exit(118)
	}
	orphanWriteReady(root, label)
	if !orphanWaitFile(nestedFile(root, label, "arm"), 15*time.Second) {
		os.Exit(119)
	}
	time.Sleep(200 * time.Millisecond)
	// Observe a surviving orphan before a closed stream could kill it through SIGPIPE.
	if err := os.WriteFile(nestedFile(root, label, "late"), []byte(label), 0600); err != nil {
		os.Exit(120)
	}
	fmt.Fprint(os.Stdout, nestedOut(label))
	fmt.Fprint(os.Stderr, nestedErr(label))
	os.Exit(code)
}

func orphanNative(t *testing.T, root, label, mode string, code int) (<-chan nestedOutcome, context.CancelFunc) {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	a := orphanArgs("bridge", root, label, mode, code)
	cmd := exec.CommandContext(ctx, a[0], a[1:]...)
	cmd.Dir, cmd.Env = root, append(os.Environ(), orphanEnv()...)
	var out, stderr bytes.Buffer
	cmd.Stdout, cmd.Stderr = &out, &stderr
	done := make(chan nestedOutcome, 1)
	if err := cmd.Start(); err != nil {
		cancel()
		t.Fatal("native orphan bridge launch", err)
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

func orphanReadiness(t *testing.T, root string, labels []string, done <-chan nestedOutcome) map[string]orphanReady {
	t.Helper()
	end := time.Now().Add(3 * time.Second)
	found := map[string]orphanReady{}
	seen := map[int]bool{}
	for time.Now().Before(end) {
		for _, label := range labels {
			if _, ok := found[label]; ok {
				continue
			}
			var ready orphanReady
			raw, err := os.ReadFile(nestedFile(root, label, "ready"))
			if err != nil || json.Unmarshal(raw, &ready) != nil || ready.PID <= 0 || ready.Root != root || ready.Label != label {
				continue
			}
			if seen[ready.PID] {
				t.Fatal("readiness reused helper PID", found, ready)
			}
			owned, err := orphanPin(ready.PID, root, label)
			if err != nil {
				t.Fatal("pin live orphan helper", label, err)
			}
			// Register each exact instance immediately, including incomplete-readiness failures.
			t.Cleanup(func() {
				defer owned.close()
				alive, err := owned.alive()
				if err == nil && alive {
					err = owned.kill()
				}
				if err != nil {
					t.Errorf("owned orphan helper cleanup %s: %v", label, err)
				}
			})
			ready.owned = owned
			found[label] = ready
			seen[ready.PID] = true
		}
		if len(found) == len(labels) {
			for label, ready := range found {
				alive, err := ready.owned.alive()
				if err != nil || !alive {
					t.Fatal("ready helper was not alive before parent exit", label, alive, err)
				}
			}
			return found
		}
		select {
		case result := <-done:
			t.Fatalf("outer returned before helper readiness: %+v", result)
		default:
		}
		time.Sleep(5 * time.Millisecond)
	}
	t.Fatal("bounded orphan helper readiness timed out")
	return nil
}

func orphanProve(t *testing.T, root, label, mode string, ready map[string]orphanReady, done <-chan nestedOutcome) {
	t.Helper()
	parent, leaf := ready[label+"-parent"], ready[label+"-leaf"]
	if leaf.Parent != parent.PID {
		t.Fatal("native original parent premise absent", parent, leaf)
	}
	if err := orphanIndependent(parent, leaf, mode); err != nil {
		t.Fatal("native independent group/session premise", err)
	}
	if err := os.WriteFile(nestedFile(root, label, "exit-parent"), []byte("exit now"), 0600); err != nil {
		t.Fatal(err)
	}
	end := time.Now().Add(time.Second)
	for {
		alive, err := parent.owned.alive()
		if err != nil {
			t.Fatal("native parent state", err)
		}
		if !alive {
			active, err := leaf.owned.alive()
			if err != nil || !active {
				t.Fatal("orphan not live after native parent termination", active, err)
			}
			if err := orphanReparented(leaf.PID, parent.PID); err != nil {
				t.Fatal("native orphan reparenting", err)
			}
			select {
			case result := <-done:
				t.Fatalf("outer returned before meaningful orphan cancellation: %+v", result)
			default:
			}
			return
		}
		if time.Now().After(end) {
			t.Fatal("parent did not terminate within native observation bound")
		}
		time.Sleep(5 * time.Millisecond)
	}
}

func orphanCleanup(t *testing.T, root string, stops ...context.CancelFunc) {
	t.Helper()
	t.Cleanup(func() {
		for _, stop := range stops {
			stop()
		}
		for _, label := range []string{"subject", "outside", "control"} {
			_ = os.WriteFile(nestedFile(root, label, "exit-parent"), []byte("cleanup"), 0600)
		}
		nestedArm(root, "subject-leaf", "outside-leaf", "overlap", "control-leaf")
	})
}

func orphanStreams(t *testing.T, kind string, result nestedOutcome, labels ...string) {
	t.Helper()
	if kind == "Run" {
		for _, label := range labels {
			if strings.Count(result.Out, nestedOut(label)) != 1 || strings.Count(result.Out, nestedErr(label)) != 1 {
				t.Error("Run lost/duplicated orphan stream bytes", label, result)
			}
		}
		return
	}
	out, stderr := "", ""
	for _, label := range labels {
		out += nestedOut(label)
		stderr += nestedErr(label)
	}
	if result.Out != out || result.Stderr != stderr {
		t.Error("orphan streams lost or mixed", result, out, stderr)
	}
}

func Test_GIT006_GIT010OrphanControls(t *testing.T) {
	nestedSupported(t)
	for _, mode := range orphanModes() {
		t.Run(mode, func(t *testing.T) {
			for _, kind := range []string{"Native", "Run", "Capture"} {
				t.Run(kind, func(t *testing.T) {
					for _, code := range []int{0, 37} {
						t.Run(strconv.Itoa(code), func(t *testing.T) {
							root := t.TempDir()
							var done <-chan nestedOutcome
							var stop context.CancelFunc
							if kind == "Native" {
								done, stop = orphanNative(t, root, "control", mode, code)
							} else {
								ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
								stop = cancel
								ch := make(chan nestedOutcome, 1)
								done = ch
								go func() {
									ch <- nestedCall(kind, ctx, root, orphanArgs("bridge", root, "control", mode, code), orphanEnv())
								}()
							}
							orphanCleanup(t, root, stop)
							ready := orphanReadiness(t, root, []string{"control-parent", "control-leaf"}, done)
							orphanProve(t, root, "control", mode, ready, done)
							nestedExpectLate(t, root, "control-leaf", false)
							nestedArm(root, "control-leaf")
							result := nestedWaitOutcome(t, done, 3*time.Second)
							if result.Exit != code || (result.Err == nil) != (code == 0) || errors.Is(result.Err, context.Canceled) || errors.Is(result.Err, context.DeadlineExceeded) {
								t.Error("ordinary orphan outcome", result, code)
							}
							orphanStreams(t, kind, result, "control-parent", "control-leaf")
							nestedExpectLate(t, root, "control-leaf", true)
						})
					}
				})
			}
		})
	}
}

func orphanCancellation(t *testing.T, kind string) {
	nestedSupported(t)
	for _, termination := range []string{"cancel", "deadline"} {
		t.Run(termination, func(t *testing.T) {
			for _, mode := range orphanModes() {
				t.Run(mode, func(t *testing.T) {
					root := t.TempDir()
					cwd, _ := os.Getwd()
					parentEnv := os.Environ()
					outside, stopOutside := orphanNative(t, root, "outside", mode, 0)
					overlapCtx, stopOverlap := context.WithTimeout(context.Background(), 12*time.Second)
					overlap := make(chan nestedOutcome, 1)
					go func() {
						overlap <- nestedCall("Capture", overlapCtx, root, orphanArgs("leaf", root, "overlap", mode, 0), orphanEnv())
					}()
					var ctx context.Context
					var stop context.CancelFunc
					if termination == "deadline" {
						ctx, stop = context.WithTimeout(context.Background(), 5*time.Second)
					} else {
						ctx, stop = context.WithCancel(context.Background())
					}
					orphanCleanup(t, root, stop, stopOutside, stopOverlap)
					argv, env := orphanArgs("bridge", root, "subject", mode, 0), orphanEnv()
					originalArgv, originalEnv := append([]string(nil), argv...), append([]string(nil), env...)
					done := make(chan nestedOutcome, 1)
					go func() { done <- nestedCall(kind, ctx, root, argv, env) }()
					ready := orphanReadiness(t, root, []string{"subject-parent", "subject-leaf", "outside-parent", "outside-leaf", "overlap"}, done)
					orphanProve(t, root, "outside", mode, ready, outside)
					if termination == "deadline" {
						deadline, ok := ctx.Deadline()
						if !ok || ctx.Err() != nil {
							t.Fatal("real pending deadline premise absent", ctx.Err())
						}
						pause := time.Until(deadline.Add(-700 * time.Millisecond))
						if pause <= 0 {
							t.Fatal("insufficient readiness margin before real deadline")
						}
						time.Sleep(pause)
					}
					orphanProve(t, root, "subject", mode, ready, done)
					if ctx.Err() != nil {
						t.Fatal("context ended before native orphan premise", ctx.Err())
					}
					if termination == "cancel" {
						stop()
					} else {
						select {
						case <-ctx.Done():
						case <-time.After(2 * time.Second):
							t.Fatal("real deadline did not expire")
						}
					}
					result := nestedWaitOutcome(t, done, 2*time.Second)
					if result.Exit != -1 || !errors.Is(result.Err, ctx.Err()) {
						t.Error("orphan context result", result, ctx.Err())
					}
					orphanStreams(t, kind, result, "subject-parent")
					for label, ch := range map[string]<-chan nestedOutcome{"outside": outside, "overlap": overlap} {
						select {
						case other := <-ch:
							t.Errorf("unrelated %s ended at subject cancellation: %+v", label, other)
						default:
						}
						key := label
						if label == "outside" {
							key = "outside-leaf"
						}
						alive, err := ready[key].owned.alive()
						if err != nil || !alive {
							t.Error("unrelated live process lost", label, alive, err)
						}
					}
					// Only now can a survivor demonstrate a post-return write.
					nestedArm(root, "subject-leaf", "outside-leaf", "overlap")
					other := nestedWaitOutcome(t, outside, 3*time.Second)
					if other.Exit != 0 || other.Err != nil {
						t.Error("unrelated native orphan failed", other)
					}
					orphanStreams(t, "Native", other, "outside-parent", "outside-leaf")
					nestedExpectLate(t, root, "outside-leaf", true)
					other = nestedWaitOutcome(t, overlap, 3*time.Second)
					if other.Exit != 0 || other.Err != nil {
						t.Error("overlapping public call failed", other)
					}
					orphanStreams(t, "Capture", other, "overlap")
					nestedExpectLate(t, root, "overlap", true)
					time.Sleep(300 * time.Millisecond)
					nestedExpectLate(t, root, "subject-leaf", false)
					end := time.Now().Add(time.Second)
					for {
						alive, err := ready["subject-leaf"].owned.alive()
						if err != nil {
							t.Fatal("owned orphan state", err)
						}
						if !alive {
							break
						}
						if time.Now().After(end) {
							t.Error("owned orphan remains live after cancellation")
							break
						}
						time.Sleep(10 * time.Millisecond)
					}
					after, _ := os.Getwd()
					if cwd != after || !reflect.DeepEqual(parentEnv, os.Environ()) || !reflect.DeepEqual(argv, originalArgv) || !reflect.DeepEqual(env, originalEnv) {
						t.Error("orphan cancellation changed caller state or inputs")
					}
					laterCtx, cancelLater := context.WithTimeout(context.Background(), 3*time.Second)
					later := nestedCall(kind, laterCtx, root, helperCommand(t, "exit"), processEnvironment())
					cancelLater()
					if later.Exit != 37 || later.Err == nil || errors.Is(later.Err, context.Canceled) || errors.Is(later.Err, context.DeadlineExceeded) {
						t.Error("orphan cancellation poisoned ordinary exit", later)
					}
				})
			}
		})
	}
}

func Test_GIT006OrphanRunCancellation(t *testing.T)     { orphanCancellation(t, "Run") }
func Test_GIT010OrphanCaptureCancellation(t *testing.T) { orphanCancellation(t, "Capture") }
