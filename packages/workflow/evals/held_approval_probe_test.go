package evals

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
	"time"

	p "github.com/lurkshark/fireweed/packages/project"
	w "github.com/lurkshark/fireweed/packages/workflow"
)

// Object-availability measurements always enter a fresh public process. Setup
// and reference controls cannot populate a measured immutable-object cache.
func heldBuildColdProbe(t *testing.T) string {
	t.Helper()
	directory := stateCanonicalDirectory(t, t.TempDir())
	worktreePut(t, directory, "probe.go", `package main
import("encoding/json";"os";p "github.com/lurkshark/fireweed/packages/project";w "github.com/lurkshark/fireweed/packages/workflow")
type output struct{Snapshot p.Snapshot;Claim w.Claim;ClaimObject string;Present bool;Error string;StateSame bool}
func main(){var result output;var e error;if os.Args[1]=="approval"{result.Snapshot,result.Claim,e=w.LoadApproval(os.Args[2],"current");result.Present=e==nil&&result.Claim!=(w.Claim{})}else if os.Args[1]=="object-only-wire"{result.Claim.Object="held-object-only-wire-control";result.Error="restore source for wire control"}else{result.Snapshot,result.Claim,result.Present,e=w.LoadHeldApproval(os.Args[2])};if e!=nil{result.Error=e.Error()};result.ClaimObject=result.Claim.Object;if json.NewEncoder(os.Stdout).Encode(result)!=nil{os.Exit(80)}}
`)
	source, err := filepath.Abs("../../..")
	if err != nil {
		t.Fatal(err)
	}
	worktreePut(t, directory, "go.mod", fmt.Sprintf("module heldcoldprobe\n\ngo 1.23\nrequire github.com/lurkshark/fireweed v0.0.0\nreplace github.com/lurkshark/fireweed => %q\n", source))
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	binary := filepath.Join(directory, "probe")
	if os.PathSeparator == '\\' {
		binary += ".exe"
	}
	command := exec.CommandContext(ctx, "go", "build", "-mod=mod", "-o", binary, filepath.Join(directory, "probe.go"))
	command.Dir = directory
	command.Env = append(os.Environ(), "GOPROXY=off", "GOTOOLCHAIN=local")
	if out, err := command.CombinedOutput(); err != nil {
		t.Fatal("cold public source probe compile", err, string(out))
	}
	return binary
}

func heldColdCall(t *testing.T, executable, mode, root string) heldNativeOutput {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	command := exec.CommandContext(ctx, executable, mode, root)
	command.Dir = root
	var diagnostic strings.Builder
	command.Stderr = &diagnostic
	raw, err := command.Output()
	if err != nil || ctx.Err() != nil {
		t.Fatal("cold bounded public source call", err, ctx.Err(), string(raw), diagnostic.String())
	}
	var result heldNativeOutput
	if err := json.Unmarshal(raw, &result); err != nil {
		t.Fatal("cold public source output", err, string(raw), diagnostic.String())
	}
	result.Claim.Object = result.ClaimObject
	return result
}

func heldColdAssert(t *testing.T, result heldNativeOutput, f approvalFixtureData, available bool) {
	t.Helper()
	if available {
		if result.Error != "" || !result.Present || result.Claim != f.Claim || !reflect.DeepEqual(result.Snapshot, f.Want) {
			t.Error("cold locally complete approved source", result.Error, result.Present, result.Claim)
		}
	} else if result.Error == "" || !policyCorrective(result.Error) || result.Present || result.Claim != (w.Claim{}) || !reflect.DeepEqual(result.Snapshot, p.Snapshot{}) {
		t.Error("cold unavailable source leaked authority or lacked useful corrective error", result.Error, result.Present, result.Claim)
	}
}

func Test_FLOW049_FLOW051HeldColdProbeWireControls(t *testing.T) {
	probe := heldBuildColdProbe(t)
	f := heldSourceFixture(t, "impl", "sha1", true)
	result := heldColdCall(t, probe, "approval", f.Root)
	if result.Error != "" || !result.Present || result.Claim != f.Claim || !reflect.DeepEqual(result.Snapshot, f.Want) {
		t.Fatal("cold nonzero exact-token wire control", result.Error, result.Claim)
	}
	result = heldColdCall(t, probe, "object-only-wire", f.Root)
	if result.Claim != (w.Claim{Object: "held-object-only-wire-control"}) || result.ClaimObject == "" {
		t.Fatal("cold wire concealed an Object-only partial claim", result)
	}
}
