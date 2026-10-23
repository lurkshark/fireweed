package evals

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	w "github.com/lurkshark/fireweed/packages/workflow"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func Test_FLOW011ClaimLocalityOracleControls(t *testing.T) {
	tools := claimLocalTransport(t)
	for _, kind := range []string{"claim-body", "stored-base"} {
		t.Run("allowed-cas/"+kind, func(t *testing.T) {
			f := claimLocalSetup(t, tools, "sha1", kind, "update", false)
			claimLocalCall(t, f, "update", func() (w.Claim, []w.Claim, error) {
				// Independent exact Git CAS demonstrates a valid local proof using the
				// complete caller-supplied Claim, without reading unavailable old bodies.
				replacement := f.Previous
				replacement.Base = f.Head
				replacement.State = "new opaque state"
				oid, e := claimCommand(t, f.Root, claimPayload(t, replacement), "hash-object", "-w", "--stdin")
				if e != nil {
					t.Fatal(e)
				}
				claimGit(t, f.Root, "update-ref", "--no-deref", "refs/fireweed/claims/current", oid, f.Previous.Object)
				replacement.Object = oid
				return replacement, nil, nil
			})
		})
	}
	t.Run("empty-nonnil-inventory", func(t *testing.T) {
		f := claimLocalSetup(t, tools, "sha1", "claim-body", "list", false)
		claimLocalCall(t, f, "list", func() (w.Claim, []w.Claim, error) {
			return w.Claim{}, []w.Claim{}, fmt.Errorf("P02: required selected object is unavailable. Use fireweed release after checking the owner.")
		})
	})
	t.Run("unreachable-attempt-blob", func(t *testing.T) {
		f := claimLocalSetup(t, tools, "sha1", "claim-body", "update", false)
		claimLocalCall(t, f, "update", func() (w.Claim, []w.Claim, error) {
			replacement := f.Previous
			replacement.Base = f.Head
			replacement.State = "new opaque state"
			if _, e := claimCommand(t, f.Root, claimPayload(t, replacement), "hash-object", "-w", "--stdin"); e != nil {
				t.Fatal(e)
			}
			return w.Claim{}, nil, fmt.Errorf("P02: ownership could not be proven. Use fireweed release after checking the owner.")
		})
	})
	t.Run("fabricated-success-token", func(t *testing.T) {
		f := claimLocalSetup(t, tools, "sha1", "claim-body", "update", false)
		objects := claimTree(t, filepath.Join(f.Common, "objects"), func(string) bool { return false })
		fabricated := f.Previous
		fabricated.State = "new opaque state"
		fabricated.Base = f.Head
		if claimGit(t, f.Root, "for-each-ref", "--format=%(objectname)", "refs/fireweed/claims/current") != fabricated.Object {
			t.Fatal("unchanged dangling ref premise")
		}
		if e := claimLocalStoredToken(f, fabricated, objects); e == nil {
			t.Fatal("fabricated success with missing unchanged token accepted")
		}
		fabricated.Object = f.Other.Object
		if e := claimLocalStoredToken(f, fabricated, objects); e == nil {
			t.Fatal("fabricated success with unrelated preexisting token accepted")
		}
		if e := claimLocalStoredToken(f, f.Other, objects); e != nil {
			t.Fatal("valid preexisting stored token rejected", e)
		}
	})

	for _, format := range []string{"sha1", "sha256"} {
		t.Run("storage/"+format, func(t *testing.T) {
			root := claimRepo(t, format)
			head := claimBase(t, root)
			f := claimLocalFixture{Root: root, Head: head, Old: head, Base: head, Slug: "fresh"}
			before := map[string]string{".": "drwxr-xr-x"}
			add := func(dst map[string]string, body string) {
				raw := []byte(fmt.Sprintf("blob %d\x00%s", len(body), body))
				id := ""
				if len(head) == 40 {
					h := sha1.Sum(raw)
					id = fmt.Sprintf("%x", h)
				} else {
					h := sha256.Sum256(raw)
					id = fmt.Sprintf("%x", h)
				}
				var buf bytes.Buffer
				z := zlib.NewWriter(&buf)
				if _, e := z.Write(raw); e != nil {
					t.Fatal(e)
				}
				if e := z.Close(); e != nil {
					t.Fatal(e)
				}
				dst[id[:2]] = "drwxr-xr-x"
				dst[id[:2]+"/"+id[2:]] = "-r--r--r--\n" + buf.String()
			}
			candidate := w.Claim{Worktree: root, Change: "fresh", Run: "new-run", Base: head, State: "new opaque state"}
			good := map[string]string{".": "drwxr-xr-x"}
			add(good, claimPayload(t, candidate))
			if errs := claimLocalObjectStorage(f, "create", before, good); len(errs) != 0 {
				t.Fatal("valid selected blob rejected", errs)
			}
			for _, kind := range []string{"empty-dir", "garbage", "pack", "unrelated-blob", "wrong-hash", "excess-encodings", "existing-mutation", "read-write"} {
				t.Run(kind, func(t *testing.T) {
					bad := map[string]string{}
					for k, v := range good {
						bad[k] = v
					}
					action := "create"
					switch kind {
					case "empty-dir":
						for i := 0; i < 256; i++ {
							p := fmt.Sprintf("%02x", i)
							if _, ok := bad[p]; !ok {
								bad[p] = "drwxr-xr-x"
								break
							}
						}
					case "garbage":
						bad["temporary"] = "-rw-------\ngarbage"
					case "pack":
						bad["pack"] = "drwxr-xr-x"
						bad["pack/arbitrary.pack"] = "-rw-------\ngarbage"
					case "unrelated-blob":
						add(bad, "unrelated garbage bytes")
					case "wrong-hash":
						for k, v := range bad {
							if strings.Contains(k, "/") {
								delete(bad, k)
								q := k[:len(k)-1] + "0"
								if q == k {
									q = k[:len(k)-1] + "1"
								}
								bad[q] = v
								break
							}
						}
					case "excess-encodings":
						raw, _ := json.Marshal(candidate)
						add(bad, string(raw)+"\n")
						add(bad, string(raw)+" \n")
					case "existing-mutation":
						bad["."] = "drwx------"
					case "read-write":
						action = "load"
					}
					if errs := claimLocalObjectStorage(f, action, before, bad); len(errs) == 0 {
						t.Fatal("invalid storage accepted", kind)
					}
				})
			}
			// The decoder reads snapshot bytes directly; it must never open a path or
			// depend on an object query that could fetch from the fixture's remote.
			if _, e := os.Stat(filepath.Join(root, "temporary")); !os.IsNotExist(e) {
				t.Fatal(e)
			}
		})
	}
}
