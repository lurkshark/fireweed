package evals

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	g "github.com/lurkshark/fireweed/packages/gitstore"
)

func Test_GIT011NativeRevisionMetadata(t *testing.T) {
	tools := gitReadPrepare(t)
	selectors := []string{"topic", "TOPIC", "HEAD@{1}", "main@{1}", "upstream/HEAD@{1}", "ORIG_HEAD"}
	for _, format := range []string{"sha1", "sha256"} {
		for _, linked := range []bool{false, true} {
			for _, selector := range selectors {
				t.Run(fmt.Sprintf("%s/linked_%t/%s", format, linked, selector), func(t *testing.T) {
					f := gitReadRepository(t, tools, format, linked, "")
					// Store both identities in the case-sensitive native representation even
					// when the host cannot represent them as distinct loose filenames.
					write(t, f.Common, "packed-refs", fmt.Sprintf("# pack-refs with: sorted\n%s refs/tags/TOPIC\n%s refs/tags/topic\n", f.Head, f.Old))
					if got := gitReadOK(t, tools, f.Root, "rev-parse", "--verify", "topic"); got != f.Old {
						t.Fatal("native lowercase packed tag premise", got)
					}
					if got := gitReadOK(t, tools, f.Root, "rev-parse", "--verify", "TOPIC"); got != f.Head || got == f.Old {
						t.Fatal("native uppercase packed tag premise", got)
					}
					gitReadOK(t, tools, f.Root, "-c", "user.name=Ref Eval", "-c", "user.email=ref@eval.invalid", "reset", "--soft", f.Old)
					gitReadOK(t, tools, f.Root, "-c", "user.name=Ref Eval", "-c", "user.email=ref@eval.invalid", "reset", "--soft", f.Head)
					mainRoot := filepath.Dir(f.Common)
					if linked {
						// Give the main checkout a different HEAD reflog. Reading common logs/HEAD
						// instead of this selected worktree's log must change the expected result.
						gitReadOK(t, tools, mainRoot, "-c", "user.name=Ref Eval", "-c", "user.email=ref@eval.invalid", "commit", "--allow-empty", "-qm", "different main reflog")
						commonPrevious := gitReadOK(t, tools, mainRoot, "rev-parse", "HEAD@{1}")
						selectedPrevious := gitReadOK(t, tools, f.Root, "rev-parse", "HEAD@{1}")
						if commonPrevious != f.Head || selectedPrevious != f.Old || commonPrevious == selectedPrevious {
							t.Fatal("selected/common reflog distinction is not live")
						}
						write(t, f.Common, "ORIG_HEAD", f.Head+"\n")
					}
					gitReadOK(t, tools, f.Root, "symbolic-ref", "refs/remotes/upstream/HEAD", "refs/heads/main")
					aliasLog := gitReadPath(t, tools, f.Root, "logs/refs/remotes/upstream/HEAD")
					if e := os.Remove(aliasLog); e != nil && !os.IsNotExist(e) {
						t.Fatal(e)
					}
					gitReadOK(t, tools, f.Root, "update-ref", "ORIG_HEAD", f.Old)
					if linked {
						common, e := os.ReadFile(filepath.Join(f.Common, "ORIG_HEAD"))
						if e != nil || strings.TrimSpace(string(common)) != f.Head {
							t.Fatal("main pseudoref distinction premise failed", e)
						}
						selected, e := os.ReadFile(filepath.Join(f.GitDir, "ORIG_HEAD"))
						if e != nil || strings.TrimSpace(string(selected)) != f.Old {
							t.Fatal("selected pseudoref distinction premise failed", e)
						}
					}
					oid := f.Old
					if selector == "TOPIC" || (linked && (selector == "main@{1}" || selector == "upstream/HEAD@{1}")) {
						oid = f.Head
					}
					if resolved := gitReadOK(t, tools, f.Root, "rev-parse", "--verify", selector+"^{commit}"); resolved != oid {
						t.Fatalf("native revision premise got %s, want %s", resolved, oid)
					}
					want := []g.Edit{}
					if oid == f.Old {
						want = []g.Edit{{Status: "M", Path: "nested/file.txt"}, {Status: "M", Path: "tracked.txt"}}
					}
					native := gitReadNativeEdits(t, tools, f.Root, selector)
					if !reflect.DeepEqual(native, want) {
						t.Fatalf("native metadata diff got %+v, manual want %+v", native, want)
					}
					gitReadNoTransport(t, tools)
					gitReadReset(t, tools)
					temporary := gitReadTemporary(t)
					before := gitReadStateImage(t, f, temporary)
					metadata := cleanMetadata(t, f)
					mainBefore := gitReadTreeImage(t, mainRoot, true)
					got, e := g.Edits(f.Root, selector)
					if e != nil || !reflect.DeepEqual(got, want) {
						t.Errorf("metadata expression %q got %+v/%v, native %+v", selector, got, e, native)
					}
					gitReadNoTransport(t, tools)
					gitReadSameState(t, f, temporary, before)
					if !reflect.DeepEqual(metadata, cleanMetadata(t, f)) {
						t.Error("Edits changed original reflog/pseudoref/attributes metadata")
					}
					if !reflect.DeepEqual(mainBefore, gitReadTreeImage(t, mainRoot, true)) {
						t.Error("Edits changed main checkout files")
					}
				})
			}
		}
	}
}
