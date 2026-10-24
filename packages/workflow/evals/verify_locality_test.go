package evals

import (
	"context"
	"fmt"
	w "github.com/lurkshark/fireweed/packages/workflow"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
	"time"
)

// Configure a live local marker-only promisor, remove one known loose object,
// and prove both allowed-transport and denied-protocol fetch controls.
func verifyLocalMissing(t *testing.T, root, oid string, denied bool) (string, string) {
	t.Helper()
	common := claimCommon(t, root)
	for k, v := range map[string]string{"extensions.partialClone": "origin", "remote.origin.promisor": "true", "remote.origin.partialclonefilter": "blob:none", "remote.origin.url": "claimlocal::marker-only", "protocol.claimlocal.allow": "always"} {
		claimGit(t, root, "config", k, v)
	}
	object := filepath.Join(common, "objects", oid[:2], oid[2:])
	raw, e := os.ReadFile(object)
	if e != nil {
		t.Fatal("known loose object premise", e)
	}
	if e = os.Remove(object); e != nil {
		t.Fatal(e)
	}
	t.Cleanup(func() { os.MkdirAll(filepath.Dir(object), 0755); os.WriteFile(object, raw, 0444) })
	dir := stateCanonicalDirectory(t, t.TempDir())
	marker, trace := filepath.Join(dir, "marker"), filepath.Join(dir, "trace.json")
	t.Setenv("FW_CLAIM_LOCAL_MARKER", marker)
	t.Setenv("GIT_TRACE2_EVENT", trace)
	real, e := exec.LookPath("git")
	if e != nil {
		t.Fatal(e)
	}
	for _, deny := range []bool{false, true} {
		claimLocalReset(t, marker, trace)
		args := []string{}
		if deny {
			args = append(args, "-c", "protocol.claimlocal.allow=never")
		}
		args = append(args, "cat-file", "-t", oid)
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		cmd := exec.CommandContext(ctx, real, args...)
		cmd.Dir = root
		cmd.Env = append(os.Environ(), "GIT_NO_LAZY_FETCH=0")
		out, e := cmd.CombinedOutput()
		cancel()
		if e == nil || !claimLocalFetch(t, trace) {
			t.Fatal("live missing-object fetch control absent", deny, e, string(out))
		}
		data, e := os.ReadFile(marker)
		if deny {
			if !os.IsNotExist(e) {
				t.Fatal("denied protocol still invoked transport", e)
			}
		} else if e != nil || !strings.Contains(string(data), "transport attempted") {
			t.Fatal("transport control absent", e, string(out))
		}
	}
	if denied {
		claimGit(t, root, "config", "protocol.claimlocal.allow", "never")
	}
	claimLocalReset(t, marker, trace)
	return marker, trace
}
func verifyLocalUnchanged(t *testing.T, root string, call func()) {
	t.Helper()
	temp := stateCanonicalDirectory(t, t.TempDir())
	t.Setenv("TMPDIR", temp)
	t.Setenv("TMP", temp)
	t.Setenv("TEMP", temp)
	claimGit(t, root, "--version")
	common := claimCommon(t, root)
	files := startTree(t, root)
	meta := claimTree(t, common, func(string) bool { return false })
	scratch := claimTree(t, temp, func(string) bool { return false })
	cwd, _ := os.Getwd()
	env := os.Environ()
	call()
	after, _ := os.Getwd()
	if after != cwd || !reflect.DeepEqual(env, os.Environ()) || !reflect.DeepEqual(files, startTree(t, root)) || !reflect.DeepEqual(meta, claimTree(t, common, func(string) bool { return false })) || !reflect.DeepEqual(scratch, claimTree(t, temp, func(string) bool { return false })) {
		t.Error("locality call changed original files/index/refs/Git metadata/parent state or leaked owned storage")
	}
}
func Test_FLOW023ApprovalSelectedObjectNeverFetches(t *testing.T) {
	claimLocalTransport(t)
	for _, format := range []string{"sha1", "sha256"} {
		for _, linked := range []bool{false, true} {
			for _, denied := range []bool{false, true} {
				for _, kind := range []string{"selected-claim", "unrelated-claim"} {
					t.Run(fmt.Sprintf("%s/linked_%v/denied_%v/%s", format, linked, denied, kind), func(t *testing.T) {
						f := approvalFixture(t, "impl", format, true)
						if linked {
							peer := claimPeer(t, f.Root)
							if e := w.ReleaseClaim(f.Root, "current", f.Claim.Object); e != nil {
								t.Fatal(e)
							}
							f.Root = peer
							cl, e := w.CreateClaim(peer, "current", f.Claim.Run, f.Claim.Base, f.Claim.State)
							if e != nil {
								t.Fatal(e)
							}
							f.Claim = cl
							f.Want.Root = peer
						}
						if _, cl, e := w.LoadApproval(f.Root, "current"); e != nil || cl != f.Claim {
							t.Fatal("available exact approval control", cl, e)
						}
						oid := f.Claim.Object
						if kind == "unrelated-claim" {
							oid = claimRaw(t, f.Root, "unrelated-missing", "independent unavailable body")
						}
						marker, trace := verifyLocalMissing(t, f.Root, oid, denied)
						verifyLocalUnchanged(t, f.Root, func() {
							s, cl, e := w.LoadApproval(f.Root, "current")
							claimLocalNoFetch(t, marker, trace)
							if kind == "selected-claim" {
								if !reflect.DeepEqual(s, reflect.Zero(reflect.TypeOf(s)).Interface()) || cl != (w.Claim{}) || e == nil || !strings.Contains(e.Error(), "P02") || !policyCorrective(e.Error()) {
									t.Errorf("selected missing claim did not fail atomically/usefully %+v %+v %v", s, cl, e)
								}
							} else {
								if e != nil || cl != f.Claim || !reflect.DeepEqual(s, f.Want) {
									t.Errorf("unrelated missing body obstructed complete exact approval %+v %+v %v", s, cl, e)
								}
							}
						})
					})
				}
			}
		}
	}
}
func Test_FLOW015WorktreeMovedUnavailableHeadNeverFetches(t *testing.T) {
	claimLocalTransport(t)
	for _, format := range []string{"sha1", "sha256"} {
		for _, linked := range []bool{false, true} {
			for _, denied := range []bool{false, true} {
				for _, kind := range []string{"moved-head", "unrelated-ancestor"} {
					t.Run(fmt.Sprintf("%s/linked_%v/denied_%v/%s", format, linked, denied, kind), func(t *testing.T) {
						f := worktreeFixture(t, "impl", format)
						if linked {
							peer := claimPeer(t, f.Root)
							if e := w.ReleaseClaim(f.Root, "current", f.Claim.Object); e != nil {
								t.Fatal(e)
							}
							cl, e := w.CreateClaim(peer, "current", f.Claim.Run, f.Claim.Base, f.Claim.State)
							if e != nil {
								t.Fatal(e)
							}
							f.Root = peer
							f.Snapshot.Root = peer
							f.Claim = cl
							worktreePut(t, peer, f.Path, cl.State)
						}
						before := w.CheckWorktree(f.Snapshot, f.Module, f.Path, f.Claim)
						if len(before.Diagnostics) != 0 {
							t.Fatal("valid exact Object-token guard control", before)
						}
						oid := ""
						if kind == "moved-head" {
							claimGit(t, f.Root, "-c", "user.name=Verify Locality", "-c", "user.email=local@example.invalid", "commit", "--allow-empty", "-qm", "moved worker head")
							oid = claimBase(t, f.Root)
						} else {
							oid = claimGit(t, f.Root, "rev-parse", f.Claim.Base+"~1")
						}
						marker, trace := verifyLocalMissing(t, f.Root, oid, denied)
						verifyLocalUnchanged(t, f.Root, func() {
							v := w.CheckWorktree(f.Snapshot, f.Module, f.Path, f.Claim)
							claimLocalNoFetch(t, marker, trace)
							if kind == "moved-head" {
								if len(v.Edits) != 0 || v.Edits == nil || len(v.Diagnostics) != 1 || v.Diagnostics[0].Code != "P02" || v.Diagnostics[0].Path != f.Path || !policyCorrective(v.Diagnostics[0].Message) {
									t.Errorf("unavailable moved HEAD not rejected before current comparison %+v", v)
								}
							} else {
								if !reflect.DeepEqual(v, before) {
									t.Errorf("unrelated ancestor obstructed available guard got %+v want %+v", v, before)
								}
							}
						})
					})
				}
			}
		}
	}
}
