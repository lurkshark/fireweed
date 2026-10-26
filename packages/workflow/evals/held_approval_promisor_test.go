package evals

import (
	"context"
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

// This local helper records transport attempts and immediately fails. It never
// contacts a network endpoint, another repository or a controller.
func heldPromisorTransport(t *testing.T) string {
	t.Helper()
	directory := stateCanonicalDirectory(t, t.TempDir())
	worktreePut(t, directory, "transport.go", `package main
import "os"
func main(){if os.WriteFile(os.Getenv("HELD_PROMISOR_MARKER"),[]byte("lazy transport attempted"),0600)!=nil{os.Exit(90)};os.Exit(91)}
`)
	binary := filepath.Join(directory, "git-remote-heldfetch")
	if os.PathSeparator == '\\' {
		binary += ".exe"
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	command := exec.CommandContext(ctx, "go", "build", "-o", binary, filepath.Join(directory, "transport.go"))
	command.Env = append(os.Environ(), "GOPROXY=off", "GOTOOLCHAIN=local")
	if out, err := command.CombinedOutput(); err != nil {
		t.Fatal("local transport compile", err, string(out))
	}
	t.Setenv("PATH", directory+string(os.PathListSeparator)+os.Getenv("PATH"))
	t.Setenv("GIT_NO_LAZY_FETCH", "0")
	return directory
}

func heldFetchObserved(t *testing.T, trace string) bool {
	t.Helper()
	raw, err := os.ReadFile(trace)
	if err != nil {
		t.Fatal(err)
	}
	for _, line := range strings.Split(strings.TrimSpace(string(raw)), "\n") {
		if line == "" {
			continue
		}
		var event struct {
			Event, Name string
			Argv        []string
		}
		if err := json.Unmarshal([]byte(line), &event); err != nil {
			t.Fatal("Git Trace2 observer", err)
		}
		if event.Event == "cmd_name" && (event.Name == "fetch" || event.Name == "fetch-pack") {
			return true
		}
		if event.Event == "child_start" && len(event.Argv) > 1 && strings.HasPrefix(filepath.Base(event.Argv[0]), "git") {
			for index := 1; index < len(event.Argv); index++ {
				argument := event.Argv[index]
				if argument == "-c" || argument == "-C" || argument == "--git-dir" || argument == "--work-tree" || argument == "--exec-path" {
					index++
					continue
				}
				if strings.HasPrefix(argument, "-") {
					continue
				}
				if argument == "fetch" || argument == "fetch-pack" {
					return true
				}
				break
			}
		}
	}
	return false
}

func Test_FLOW048_FLOW049_FLOW050_FLOW051HeldPromisorLocality(t *testing.T) {
	probe := heldBuildColdProbe(t)
	heldPromisorTransport(t)
	real := protocolRealGit(t)
	for _, format := range []string{"sha1", "sha256"} {
		for _, mode := range []string{"claim", "base", "config", "module", "unrelated-body", "unneeded-parent"} {
			t.Run(format+"/"+mode, func(t *testing.T) {
				f := heldSourceFixture(t, "impl", format, true)
				selector := f.Claim.Base
				switch mode {
				case "claim":
					selector = f.Claim.Object
				case "config":
					selector += ":fireweed.yaml"
				case "module":
					selector += ":units/a/SPEC.md"
				case "unrelated-body":
					selector += ":units/a/impl/private.txt"
				case "unneeded-parent":
					selector += "^"
				}
				object := claimGit(t, f.Root, "rev-parse", selector)
				for key, value := range map[string]string{"extensions.partialClone": "origin", "remote.origin.promisor": "true", "remote.origin.partialclonefilter": "blob:none", "remote.origin.url": "heldfetch::marker-only", "protocol.heldfetch.allow": "always"} {
					claimGit(t, f.Root, "config", key, value)
				}
				artifacts := t.TempDir()
				trace, marker := filepath.Join(artifacts, "trace"), filepath.Join(artifacts, "transport")
				t.Setenv("GIT_TRACE2_EVENT", trace)
				t.Setenv("HELD_PROMISOR_MARKER", marker)
				reset := func() {
					if err := os.WriteFile(trace, nil, 0600); err != nil {
						t.Fatal(err)
					}
				}
				path := filepath.Join(claimCommon(t, f.Root), "objects", object[:2], object[2:])
				payload, err := os.ReadFile(path)
				if err != nil {
					t.Fatal("loose selected object premise", err)
				}
				if err := os.Remove(path); err != nil {
					t.Fatal(err)
				}
				t.Cleanup(func() { _ = os.WriteFile(path, payload, 0444) })
				for _, denied := range []bool{false, true} {
					reset()
					args := []string{}
					if denied {
						args = append(args, "-c", "protocol.heldfetch.allow=never")
					}
					args = append(args, "cat-file", "-p", object)
					_, detail, code := heldRunGitControl(t, real, f.Root, "", args...)
					if code == 0 || !heldFetchObserved(t, trace) {
						t.Fatal("native missing-object fetch control", code, detail)
					}
					bytes, err := os.ReadFile(marker)
					if denied {
						if !os.IsNotExist(err) {
							t.Fatal("denied transport executed helper", err)
						}
					} else {
						if err != nil || string(bytes) != "lazy transport attempted" {
							t.Fatal("live local transport control", string(bytes), err)
						}
						if err := os.Remove(marker); err != nil {
							t.Fatal(err)
						}
					}
				}
				heldPoisonCurrent(t, f.Root)
				reset()
				heldColdAssert(t, heldColdCall(t, probe, "held", f.Root), f, mode == "unrelated-body" || mode == "unneeded-parent")
				if heldFetchObserved(t, trace) {
					t.Error("held source attempted proven forbidden lazy fetch")
				}
				if bytes, err := os.ReadFile(marker); !os.IsNotExist(err) {
					t.Error("held source invoked proven local transport", string(bytes), err)
				}
			})
		}
	}
}
