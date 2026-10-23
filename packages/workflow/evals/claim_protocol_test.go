package evals

import (
	"bytes"
	"context"
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
	"time"

	w "github.com/lurkshark/fireweed/packages/workflow"
)

const protocolState = "approved\r\nλ\x00\nlocal claim warning\n  "

func protocolCommand(t *testing.T, executable, root, input string, args ...string) (string, string, error) {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, executable, args...)
	cmd.Dir = root
	cmd.Stdin = strings.NewReader(input)
	cmd.Env = append(os.Environ(), "GIT_CONFIG_NOSYSTEM=1", "GIT_CONFIG_GLOBAL="+os.DevNull, "GIT_OPTIONAL_LOCKS=0")
	var out, stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	return out.String(), stderr.String(), err
}
func protocolGit(t *testing.T, executable, root, input string, args ...string) string {
	t.Helper()
	out, stderr, err := protocolCommand(t, executable, root, input, args...)
	if err != nil {
		t.Fatalf("protocol fixture Git %q: %v stderr %q", args, err, stderr)
	}
	return out
}
func protocolRealGit(t *testing.T) string {
	t.Helper()
	real, err := exec.LookPath("git")
	if err != nil {
		t.Fatal(err)
	}
	return real
}
func protocolWarnings(t *testing.T, root string) string {
	t.Helper()
	if os.PathSeparator == '\\' {
		t.Skip("POSIX warning wrapper")
	}
	real := protocolRealGit(t)
	format := protocolGit(t, real, root, "", "rev-parse", "--show-object-format")
	dir := t.TempDir()
	script := "#!/bin/sh\nprintf 'local claim warning\\n' >&2\nexec " + claimQuote(real) + " \"$@\"\n"
	if err := os.WriteFile(filepath.Join(dir, "git"), []byte(script), 0755); err != nil {
		t.Fatal(err)
	}
	t.Setenv("PATH", dir+string(os.PathListSeparator)+os.Getenv("PATH"))
	out, stderr, err := protocolCommand(t, filepath.Join(dir, "git"), root, "", "rev-parse", "--show-object-format")
	if err != nil || out != format || stderr != "local claim warning\n" {
		t.Fatalf("warning stdout/stderr control %v %q/%q expected %q", err, out, stderr, format)
	}
	return real
}
func protocolStored(t *testing.T, real, root string, c w.Claim) {
	t.Helper()
	token := strings.TrimSpace(protocolGit(t, real, root, "", "rev-parse", "refs/fireweed/claims/"+c.Change))
	if token != c.Object || c.Object == "" {
		t.Fatalf("stored token %q != %+v", token, c)
	}
	raw := protocolGit(t, real, root, "", "--no-replace-objects", "cat-file", "blob", c.Object)
	var value w.Claim
	var fields map[string]json.RawMessage
	if json.Unmarshal([]byte(raw), &value) != nil || json.Unmarshal([]byte(raw), &fields) != nil {
		t.Fatal("invalid stored JSON", raw)
	}
	want := c
	want.Object = ""
	if value != want || len(fields) != 5 {
		t.Fatalf("stored fields %+v want %+v", value, want)
	}
}
func protocolParent(t *testing.T) func() {
	t.Helper()
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	env := os.Environ()
	return func() {
		t.Helper()
		after, err := os.Getwd()
		if err != nil || after != cwd || !reflect.DeepEqual(env, os.Environ()) {
			t.Fatal("claim operation changed parent cwd/environment")
		}
	}
}
func Test_FLOW008ClaimCreateWithGitWarnings(t *testing.T) {
	root := claimRepo(t, "sha1")
	base := claimBase(t, root)
	real := protocolWarnings(t, root)
	check := protocolParent(t)
	got, err := w.CreateClaim(root, "warning-create", "run-warnings", "HEAD", protocolState)
	if err != nil {
		t.Fatalf("successful Git stderr corrupted create: %v", err)
	}
	want := w.Claim{Worktree: root, Change: "warning-create", Run: "run-warnings", Base: base, State: protocolState, Object: got.Object}
	if got != want {
		t.Fatalf("created %+v want %+v", got, want)
	}
	protocolStored(t, real, root, got)
	check()
}
func Test_FLOW009ClaimUpdateAndReleaseWithGitWarnings(t *testing.T) {
	for _, mode := range []string{"update", "release"} {
		t.Run(mode, func(t *testing.T) {
			root := claimRepo(t, "sha1")
			old := claimSeed(t, root, "warning-operation", protocolState)
			real := protocolWarnings(t, root)
			check := protocolParent(t)
			if mode == "update" {
				nextState := protocolState + "updated\r\n"
				got, err := w.UpdateClaim(root, old, "HEAD", nextState)
				if err != nil {
					t.Fatalf("successful Git stderr corrupted update: %v", err)
				}
				want := old
				want.State = nextState
				want.Object = got.Object
				if got != want || got.Object == old.Object {
					t.Fatalf("updated %+v want %+v with new token", got, want)
				}
				protocolStored(t, real, root, got)
				if err := w.ReleaseClaim(root, old.Change, old.Object); err == nil {
					t.Fatal("stale token accepted after warned update")
				}
				protocolStored(t, real, root, got)
			} else {
				if err := w.ReleaseClaim(root, old.Change, old.Object); err != nil {
					t.Fatalf("successful Git stderr corrupted release: %v", err)
				}
				if refs := protocolGit(t, real, root, "", "for-each-ref", "--format=%(refname)", "refs/fireweed/claims/"+old.Change); refs != "" {
					t.Fatal("released ref remains", refs)
				}
			}
			check()
		})
	}
}
func Test_FLOW010ClaimReadsWithGitWarnings(t *testing.T) {
	for _, mode := range []string{"load", "inventory"} {
		t.Run(mode, func(t *testing.T) {
			root := claimRepo(t, "sha1")
			old := claimSeed(t, root, "warning-read", protocolState)
			protocolWarnings(t, root)
			files := claimTree(t, root, func(string) bool { return false })
			check := protocolParent(t)
			if mode == "load" {
				got, err := w.LoadClaim(root, old.Change)
				if err != nil || got != old {
					t.Fatalf("warned read %+v %v want %+v", got, err, old)
				}
			} else {
				got, err := w.Claims(root)
				if err != nil || !reflect.DeepEqual(got, []w.Claim{old}) {
					t.Fatalf("warned inventory %+v %v", got, err)
				}
			}
			check()
			if !reflect.DeepEqual(files, claimTree(t, root, func(string) bool { return false })) {
				t.Fatal("warned read changed files/index/refs")
			}
		})
	}
}
func protocolEnableReplacements(t *testing.T) {
	t.Helper()
	old, present := os.LookupEnv("GIT_NO_REPLACE_OBJECTS")
	if err := os.Unsetenv("GIT_NO_REPLACE_OBJECTS"); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if present {
			_ = os.Setenv("GIT_NO_REPLACE_OBJECTS", old)
		} else {
			_ = os.Unsetenv("GIT_NO_REPLACE_OBJECTS")
		}
	})
}
func Test_FLOW010ClaimObjectIdentityIgnoresReplacementViews(t *testing.T) {
	for _, format := range []string{"sha1", "sha256"} {
		for _, kind := range []string{"state", "type"} {
			t.Run(format+"/"+kind, func(t *testing.T) {
				protocolEnableReplacements(t)
				root := claimRepo(t, format)
				old := claimSeed(t, root, "replace-view", protocolState)
				real := protocolRealGit(t)
				replacement := old.Base
				if kind == "state" {
					altered := old
					altered.State = "substituted unrelated state"
					altered.Object = ""
					replacement = strings.TrimSpace(protocolGit(t, real, root, claimPayload(t, altered), "hash-object", "-w", "--stdin"))
				}
				protocolGit(t, real, root, "", "replace", "-f", old.Object, replacement)
				if kind == "state" {
					raw := protocolGit(t, real, root, "", "cat-file", "blob", old.Object)
					var got w.Claim
					if json.Unmarshal([]byte(raw), &got) != nil || got.State != "substituted unrelated state" {
						t.Fatal("replacement JSON control did not activate", raw)
					}
				} else if typ := protocolGit(t, real, root, "", "cat-file", "-t", old.Object); typ != "commit\n" {
					t.Fatal("replacement type control did not activate", typ)
				}
				raw := protocolGit(t, real, root, "", "--no-replace-objects", "cat-file", "blob", old.Object)
				var original w.Claim
				if json.Unmarshal([]byte(raw), &original) != nil || original.State != protocolState {
					t.Fatal("original blob control lost state", raw)
				}
				files := claimTree(t, root, func(string) bool { return false })
				check := protocolParent(t)
				// Both calls must inspect the stored object, and leave replacement refs untouched.
				t.Run("load", func(t *testing.T) {
					got, err := w.LoadClaim(root, old.Change)
					if err != nil || got != old {
						t.Fatalf("replacement changed exact claim: %+v %v want %+v", got, err, old)
					}
				})
				t.Run("inventory", func(t *testing.T) {
					got, err := w.Claims(root)
					if err != nil || !reflect.DeepEqual(got, []w.Claim{old}) {
						t.Fatalf("replacement changed inventory: %+v %v", got, err)
					}
				})
				check()
				if !reflect.DeepEqual(files, claimTree(t, root, func(string) bool { return false })) {
					t.Fatal("replacement reads changed repository")
				}
			})
		}
	}
}
