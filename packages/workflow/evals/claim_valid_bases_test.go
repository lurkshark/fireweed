package evals

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	w "github.com/lurkshark/fireweed/packages/workflow"
)

// Missing-object regressions cannot be satisfied by rejecting every tag or
// ancestor. The selected targets deliberately differ from the current HEAD.
func Test_FLOW008_FLOW009_FLOW011AvailableClaimBases(t *testing.T) {
	tools := claimLocalTransport(t)
	for _, format := range []string{"sha1", "sha256"} {
		for _, linked := range []bool{false, true} {
			t.Run(fmt.Sprintf("%s_linked_%t", format, linked), func(t *testing.T) {
				f := claimLocalSetup(t, tools, format, "unrelated-blob", "create", linked)
				claimGit(t, f.Root, "-c", "user.name=Local Claim", "-c", "user.email=claim@local.invalid", "tag", "-a", "selected", "-m", "available older commit", f.Old)
				claimGit(t, f.Root, "-c", "user.name=Local Claim", "-c", "user.email=claim@local.invalid", "tag", "-a", "selected-chain", "-m", "available tag chain", "selected")
				claimGit(t, f.Root, "config", "--unset", "extensions.partialClone")
				claimGit(t, f.Root, "config", "--unset", "remote.origin.promisor")
				temp := stateCanonicalDirectory(t, t.TempDir())
				t.Setenv("TMPDIR", temp)
				t.Setenv("TMP", temp)
				t.Setenv("TEMP", temp)
				claimGit(t, f.Root, "--version")
				metadata := func() map[string]string {
					m := claimTree(t, f.Common, func(p string) bool { return p == "objects" || strings.HasPrefix(p, "objects/") })
					delete(m, "refs/fireweed/claims/valid-forms")
					return m
				}
				meta := metadata()
				files := claimTree(t, f.Root, func(p string) bool { return p == ".git" })
				temporary := claimTree(t, temp, func(string) bool { return false })
				objects := claimTree(t, filepath.Join(f.Common, "objects"), func(string) bool { return false })
				storage := func(allowed ...w.Claim) {
					t.Helper()
					after := claimTree(t, filepath.Join(f.Common, "objects"), func(string) bool { return false })
					if errors := claimLocalAllowedObjectStorage(f, allowed, objects, after); len(errors) != 0 {
						t.Fatal("selected-operation storage", errors)
					}
					objects = after
				}
				env := os.Environ()
				cwd, e := os.Getwd()
				if e != nil {
					t.Fatal(e)
				}
				claimLocalNoFetch(t, f.Marker, f.Trace)
				claimLocalReset(t, f.Marker, f.Trace)
				created, e := w.CreateClaim(f.Root, "valid-forms", "valid-run", "selected", "tagged state")
				claimLocalNoFetch(t, f.Marker, f.Trace)
				if e != nil {
					t.Fatal("available annotated tag rejected", e)
				}
				claimAssert(t, f.Root, created, "valid-forms", "valid-run", f.Old, "tagged state")
				storage(created)
				chained, e := w.UpdateClaim(f.Root, created, "selected-chain", "chained state")
				claimLocalNoFetch(t, f.Marker, f.Trace)
				if e != nil {
					t.Fatal("available nested tag rejected", e)
				}
				claimAssert(t, f.Root, chained, "valid-forms", "valid-run", f.Old, "chained state")
				storage(created, chained)
				ancestor, e := w.UpdateClaim(f.Root, chained, "HEAD~1", "ancestor state")
				claimLocalNoFetch(t, f.Marker, f.Trace)
				if e != nil {
					t.Fatal("available ancestor rejected", e)
				}
				claimAssert(t, f.Root, ancestor, "valid-forms", "valid-run", f.Old, "ancestor state")
				storage(chained, ancestor)
				refs := claimGit(t, f.Root, "for-each-ref", "--format=%(refname) %(objectname)")
				candidates := func(slug string) []w.Claim {
					result := []w.Claim{}
					for _, base := range []string{f.Old, f.Head} {
						result = append(result, w.Claim{Worktree: f.Root, Change: slug, Run: "valid-run", Base: base, State: "not a single commit"})
					}
					return result
				}
				bad, e := w.CreateClaim(f.Root, "invalid-range", "valid-run", "HEAD~1..HEAD", "not a single commit")
				claimError(t, bad, e)
				storage(candidates("invalid-range")...)
				bad, e = w.UpdateClaim(f.Root, ancestor, "HEAD~1..HEAD", "not a single commit")
				claimError(t, bad, e)
				storage(append(candidates("valid-forms"), ancestor)...)
				claimLocalNoFetch(t, f.Marker, f.Trace)
				if refs != claimGit(t, f.Root, "for-each-ref", "--format=%(refname) %(objectname)") {
					t.Fatal("range rejection changed refs")
				}
				if claimBase(t, f.Root) != f.Head || !reflect.DeepEqual(meta, metadata()) || !reflect.DeepEqual(files, claimTree(t, f.Root, func(p string) bool { return p == ".git" })) || !reflect.DeepEqual(temporary, claimTree(t, temp, func(string) bool { return false })) {
					t.Fatal("available-base operations changed unrelated state")
				}
				after, e := os.Getwd()
				if e != nil || after != cwd || !reflect.DeepEqual(env, os.Environ()) {
					t.Fatal("available-base operations changed parent state", e)
				}
				if e := claimLocalStoredToken(f, ancestor, objects); e != nil {
					t.Fatal(e)
				}
			})
		}
	}
}
