package evals

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"
	"time"

	g "github.com/lurkshark/fireweed/packages/gitstore"
)

func gitReadPath(t *testing.T, tools gitReadTools, root, name string) string {
	t.Helper()
	p := gitReadOK(t, tools, root, "rev-parse", "--git-path", name)
	if !filepath.IsAbs(p) {
		p = filepath.Join(root, p)
	}
	return filepath.Clean(p)
}

func Test_GIT011_GIT003ConfiguredAvailableEdits(t *testing.T) {
	tools := gitReadPrepare(t)
	modes := []string{"autocrlf", "included-config", "info-attributes", "attributes-file", "excludes", "split-index", "timestamp-only", "literal-filenames"}
	if runtime.GOOS != "windows" {
		modes = append(modes, "filemode-false", "filemode-true")
	}
	for _, format := range []string{"sha1", "sha256"} {
		for _, linked := range []bool{false, true} {
			for _, mode := range modes {
				t.Run(fmt.Sprintf("%s/linked_%t/%s", format, linked, mode), func(t *testing.T) {
					f := gitReadRepository(t, tools, format, linked, "")
					base := f.Head
					external := t.TempDir()
					want := []g.Edit{}
					tracked := filepath.Join(f.Root, "tracked.txt")
					switch mode {
					case "autocrlf", "included-config":
						if mode == "included-config" {
							p := filepath.Join(external, "included config ' λ")
							write(t, external, filepath.Base(p), "[core]\n\tautocrlf = true\n")
							gitReadOK(t, tools, f.Root, "config", "include.path", p)
						} else {
							gitReadOK(t, tools, f.Root, "config", "core.autocrlf", "true")
						}
						write(t, f.Root, "tracked.txt", "new content\r\n")
					case "info-attributes":
						write(t, f.Root, ".gitattributes", "tracked.txt text\n")
						gitReadOK(t, tools, f.Root, "add", ".gitattributes")
						gitReadOK(t, tools, f.Root, "-c", "user.name=Local Git Eval", "-c", "user.email=local@eval.invalid", "commit", "-qm", "tracked attribute baseline")
						base = gitReadOK(t, tools, f.Root, "rev-parse", "HEAD")
						p := gitReadPath(t, tools, f.Root, "info/attributes")
						write(t, filepath.Dir(p), filepath.Base(p), "tracked.txt -text\n")
						write(t, f.Root, "tracked.txt", "new content\r\n")
						want = []g.Edit{{Status: "M", Path: "tracked.txt"}}
					case "attributes-file":
						p := filepath.Join(external, "attributes ' λ")
						write(t, external, filepath.Base(p), "tracked.txt text\n")
						gitReadOK(t, tools, f.Root, "config", "core.attributesFile", p)
						write(t, f.Root, "tracked.txt", "new content\r\n")
					case "excludes":
						p := filepath.Join(external, "excludes ' λ")
						write(t, external, filepath.Base(p), "global-hidden\n")
						gitReadOK(t, tools, f.Root, "config", "core.excludesFile", p)
						info := gitReadPath(t, tools, f.Root, "info/exclude")
						write(t, filepath.Dir(info), filepath.Base(info), "info-hidden\n")
						for _, name := range []string{"global-hidden", "info-hidden", "visible untracked"} {
							write(t, f.Root, name, "new untracked bytes\n")
						}
						want = []g.Edit{{Status: "A", Path: "visible untracked"}}
					case "split-index":
						gitReadOK(t, tools, f.Root, "update-index", "--split-index")
						raw, e := os.ReadFile(filepath.Join(f.GitDir, "index"))
						if e != nil {
							t.Fatal(e)
						}
						if !strings.Contains(string(raw), "link") {
							t.Fatal("native split-index control did not create a linked index")
						}
						write(t, f.Root, "tracked.txt", "dirty split-index bytes\n")
						want = []g.Edit{{Status: "M", Path: "tracked.txt"}}
					case "timestamp-only":
						stamp := time.Unix(1000000000, 0)
						if e := os.Chtimes(tracked, stamp, stamp); e != nil {
							t.Fatal(e)
						}
						info, e := os.Stat(tracked)
						if e != nil || info.ModTime().Unix() != stamp.Unix() {
							t.Fatal("timestamp control did not change workfile stat", e)
						}
					case "literal-filenames":
						names := []string{"-leading ; literal ' λ.txt", "untracked\twith\nnewline.txt"}
						for _, name := range names {
							write(t, f.Root, name, "literal new bytes\n")
						}
						write(t, f.Root, "nested/file.txt", "staged new content\n")
						gitReadOK(t, tools, f.Root, "add", "--", "nested/file.txt")
						write(t, f.Root, "tracked.txt", "unstaged new content\n")
						want = []g.Edit{{Status: "A", Path: names[0]}, {Status: "M", Path: "nested/file.txt"}, {Status: "M", Path: "tracked.txt"}, {Status: "A", Path: names[1]}}
					case "filemode-false", "filemode-true":
						enabled := mode == "filemode-true"
						gitReadOK(t, tools, f.Root, "config", "core.filemode", fmt.Sprint(enabled))
						if e := os.Chmod(tracked, 0755); e != nil {
							t.Fatal(e)
						}
						if enabled {
							want = []g.Edit{{Status: "M", Path: "tracked.txt"}}
						}
					}
					native := gitReadNativeEdits(t, tools, f.Root, base)
					if !reflect.DeepEqual(native, want) {
						t.Fatalf("native configuration premise got %+v, want %+v", native, want)
					}
					gitReadNoTransport(t, tools)
					gitReadReset(t, tools)
					temporary := gitReadTemporary(t)
					before := gitReadStateImage(t, f, temporary)
					externalBefore := gitReadTreeImage(t, external, false)
					got, e := g.Edits(f.Root, base)
					gitReadNoTransport(t, tools)
					if e != nil || !reflect.DeepEqual(got, native) {
						t.Errorf("configured available edits got %+v/%v, native %+v", got, e, native)
					}
					gitReadSameState(t, f, temporary, before)
					if !reflect.DeepEqual(externalBefore, gitReadTreeImage(t, external, false)) {
						t.Error("API changed external user configuration/attributes/excludes")
					}
				})
			}
		}
	}
}
