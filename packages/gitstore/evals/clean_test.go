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

func cleanNative(t *testing.T, tools gitReadTools, root string) bool {
	t.Helper()
	r := gitReadCommand(t, tools, root, nil, "status", "--porcelain=v1", "-z", "--untracked-files=all")
	if r.Exit != 0 {
		t.Fatalf("native available cleanliness premise failed: %s %s", r.Out, r.Err)
	}
	return len(r.Out) == 0
}
func cleanMetadata(t *testing.T, f gitReadFixture) gitReadTree {
	t.Helper()
	result := gitReadTree{}
	for label, root := range map[string]string{"common": f.Common, "selected": f.GitDir} {
		for path, v := range gitReadTreeImage(t, root, false) {
			if path == "info" || strings.HasPrefix(path, "info/") || path == "logs" || strings.HasPrefix(path, "logs/") || strings.HasSuffix(path, "_HEAD") || path == "shallow" || path == "commondir" || path == "gitdir" {
				result[label+"/"+path] = v
			}
		}
	}
	return result
}
func cleanAPI(t *testing.T, tools gitReadTools, f gitReadFixture, root string, want, wantError bool, allowUnavailable ...bool) {
	t.Helper()
	gitReadReset(t, tools)
	temporary := gitReadTemporary(t)
	before := gitReadStateImage(t, f, temporary)
	metadata := cleanMetadata(t, f)
	got, err := g.Clean(root)
	gitReadNoTransport(t, tools)
	if wantError {
		if err == nil || strings.TrimSpace(err.Error()) == "" || got {
			t.Errorf("unavailable/invalid cleanliness got %t/%v; want false and useful error", got, err)
		}
	} else if err != nil {
		if len(allowUnavailable) == 0 || !allowUnavailable[0] || got || strings.TrimSpace(err.Error()) == "" {
			t.Errorf("available cleanliness got %t/%v, want %t", got, err, want)
		}
	} else if got != want {
		t.Errorf("available cleanliness got %t/%v, want %t", got, err, want)
	}
	gitReadSameState(t, f, temporary, before)
	if !reflect.DeepEqual(metadata, cleanMetadata(t, f)) {
		t.Error("Clean changed local attributes/excludes/logs/pseudoref metadata")
	}
}

func Test_GIT012CleanBuiltInConfiguration(t *testing.T) {
	tools := gitReadPrepare(t)
	modes := []string{"clean", "staged", "unstaged", "deleted", "untracked", "ignored", "literal-untracked", "autocrlf", "included-config", "info-attributes", "attributes-file", "excludes", "split-index", "timestamp-only"}
	if runtime.GOOS != "windows" {
		modes = append(modes, "filemode-false", "filemode-true")
	}
	for _, format := range []string{"sha1", "sha256"} {
		for _, linked := range []bool{false, true} {
			for _, mode := range modes {
				t.Run(fmt.Sprintf("%s/linked_%t/%s", format, linked, mode), func(t *testing.T) {
					f := gitReadRepository(t, tools, format, linked, "")
					external := t.TempDir()
					want := true
					switch mode {
					case "staged":
						write(t, f.Root, "nested/file.txt", "staged clean inspection\n")
						gitReadOK(t, tools, f.Root, "add", "--", "nested/file.txt")
						want = false
					case "unstaged":
						write(t, f.Root, "tracked.txt", "unstaged clean inspection\n")
						want = false
					case "deleted":
						if e := os.Remove(filepath.Join(f.Root, "tracked.txt")); e != nil {
							t.Fatal(e)
						}
						want = false
					case "untracked":
						write(t, f.Root, "visible new file", "new\n")
						want = false
					case "ignored":
						write(t, f.Root, "ignored-user", "changed ignored bytes\n")
					case "literal-untracked":
						write(t, f.Root, "-literal ; ' λ\twith\nnewline", "literal new bytes\n")
						want = false
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
						gitReadOK(t, tools, f.Root, "-c", "user.name=Clean Eval", "-c", "user.email=clean@eval.invalid", "commit", "-qm", "attribute baseline")
						write(t, f.Root, "tracked.txt", "new content\r\n")
						gitReadOK(t, tools, f.Root, "add", "--", "tracked.txt")
						if blob := gitReadOK(t, tools, f.Root, "rev-parse", ":tracked.txt"); blob != gitReadOK(t, tools, f.Root, "rev-parse", "HEAD:tracked.txt") {
							t.Fatal("native tracked attribute normalized blob premise failed")
						}
						stamp := time.Unix(1000000000, 0)
						if e := os.Chtimes(filepath.Join(f.Root, "tracked.txt"), stamp, stamp); e != nil {
							t.Fatal(e)
						}
						if !cleanNative(t, tools, f.Root) {
							t.Fatal("tracked text attribute must first make this same CRLF workfile clean")
						}
						p := gitReadPath(t, tools, f.Root, "info/attributes")
						write(t, filepath.Dir(p), filepath.Base(p), "tracked.txt -text\n")
						want = false
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
						write(t, f.Root, "global-hidden", "ignored\n")
						write(t, f.Root, "info-hidden", "ignored\n")
					case "split-index":
						gitReadOK(t, tools, f.Root, "update-index", "--split-index")
						raw, e := os.ReadFile(filepath.Join(f.GitDir, "index"))
						if e != nil || !strings.Contains(string(raw), "link") {
							t.Fatal("split index premise missing", e)
						}
						write(t, f.Root, "tracked.txt", "dirty split index\n")
						want = false
					case "timestamp-only":
						stamp := time.Unix(1000000000, 0)
						p := filepath.Join(f.Root, "tracked.txt")
						if e := os.Chtimes(p, stamp, stamp); e != nil {
							t.Fatal(e)
						}
						info, e := os.Stat(p)
						if e != nil || info.ModTime().Unix() != stamp.Unix() {
							t.Fatal("timestamp premise missing", e)
						}
					case "filemode-false", "filemode-true":
						enabled := mode == "filemode-true"
						gitReadOK(t, tools, f.Root, "config", "core.filemode", fmt.Sprint(enabled))
						if e := os.Chmod(filepath.Join(f.Root, "tracked.txt"), 0755); e != nil {
							t.Fatal(e)
						}
						want = !enabled
					}
					if mode == "autocrlf" || mode == "included-config" || mode == "attributes-file" {
						// Record a normalized LF blob with the actual CRLF workfile
						// stat, then force a fresh content comparison without a size shortcut.
						expected := gitReadOK(t, tools, f.Root, "rev-parse", "HEAD:tracked.txt")
						gitReadOK(t, tools, f.Root, "add", "--", "tracked.txt")
						if blob := gitReadOK(t, tools, f.Root, "rev-parse", ":tracked.txt"); blob != expected {
							t.Fatal("native normalized index blob premise failed")
						}
						stamp := time.Unix(1000000000, 0)
						if e := os.Chtimes(filepath.Join(f.Root, "tracked.txt"), stamp, stamp); e != nil {
							t.Fatal(e)
						}
					}
					if got := cleanNative(t, tools, f.Root); got != want {
						t.Fatalf("native built-in premise got %t, manual want %t", got, want)
					}
					gitReadNoTransport(t, tools)
					originalExternal := gitReadTreeImage(t, external, false)
					cleanAPI(t, tools, f, f.Root, want, false)
					if !reflect.DeepEqual(originalExternal, gitReadTreeImage(t, external, false)) {
						t.Error("Clean changed external built-in configuration files")
					}
				})
			}
		}
	}
}

func Test_GIT012CleanObjectLocality(t *testing.T) {
	tools := gitReadPrepare(t)
	modes := []string{"required-commit", "missing-head-tree", "irrelevant-parent", "irrelevant-old-tree", "irrelevant-old-nested-tree", "irrelevant-old-blob", "current-blob-dirty"}
	for _, format := range []string{"sha1", "sha256"} {
		for _, linked := range []bool{false, true} {
			for _, protocol := range []string{"always", "never"} {
				for _, mode := range modes {
					t.Run(fmt.Sprintf("%s/linked_%t/%s/%s", format, linked, protocol, mode), func(t *testing.T) {
						f := gitReadRepository(t, tools, format, linked, "")
						want := mode != "current-blob-dirty"
						missing := f.Old
						switch mode {
						case "required-commit":
							missing = f.Head
						case "missing-head-tree":
							missing = gitReadOK(t, tools, f.Root, "rev-parse", "HEAD^{tree}")
						case "irrelevant-old-tree":
							missing = f.OldTree
						case "irrelevant-old-nested-tree":
							missing = f.OldNested
						case "irrelevant-old-blob":
							missing = f.OldBlob
						case "current-blob-dirty":
							missing = gitReadOK(t, tools, f.Root, "rev-parse", "HEAD:tracked.txt")
							write(t, f.Root, "tracked.txt", "changed current comparison bytes\n")
						}
						if cleanNative(t, tools, f.Root) != want {
							t.Fatal("available pre-removal manual cleanliness premise failed")
						}
						gitReadPromisor(t, tools, f, missing)
						gitReadOK(t, tools, f.Root, "config", "protocol.fwgitread.allow", protocol)
						required := mode == "required-commit"
						missingTree := mode == "missing-head-tree"
						native := gitReadCommand(t, tools, f.Root, nil, "status", "--porcelain=v1", "-z", "--untracked-files=all")
						if required || missingTree {
							if native.Exit == 0 || len(native.Err) == 0 {
								t.Fatal("native required-object rejection premise failed")
							}
						} else if native.Exit != 0 || (len(native.Out) == 0) != want {
							t.Fatalf("native missing irrelevant object premise failed: exit%d %s %s", native.Exit, native.Out, native.Err)
						}
						cleanAPI(t, tools, f, f.Root, want, required, missingTree)
					})
				}
			}
		}
	}
}

func Test_GIT012CleanConfiguredCallbacks(t *testing.T) {
	tools := gitReadPrepare(t)
	modes := []string{"monitor-clean", "monitor-dirty", "monitor-unavailable-commit", "monitor-unavailable-tree", "clean-filter-clean", "clean-filter-dirty", "clean-filter-unavailable-commit", "process-filter-clean", "process-filter-dirty", "process-filter-unavailable-commit"}
	for _, format := range []string{"sha1", "sha256"} {
		for _, linked := range []bool{false, true} {
			for _, mode := range modes {
				t.Run(fmt.Sprintf("%s/linked_%t/%s", format, linked, mode), func(t *testing.T) {
					f := gitReadRepository(t, tools, format, linked, "")
					filter := strings.Contains(mode, "filter-")
					want := strings.HasSuffix(mode, "-clean") || (filter && strings.HasSuffix(mode, "unavailable-commit"))
					if !want {
						write(t, f.Root, "tracked.txt", "NEW content\n")
					}
					stamp := time.Unix(1000000000, 0)
					if e := os.Chtimes(filepath.Join(f.Root, "tracked.txt"), stamp, stamp); e != nil {
						t.Fatal(e)
					}
					if cleanNative(t, tools, f.Root) != want {
						t.Fatal("native built-in callback comparison premise failed")
					}
					unavailable := strings.Contains(mode, "unavailable-")
					if strings.HasSuffix(mode, "unavailable-commit") {
						gitReadPromisor(t, tools, f, f.Head)
					}
					if strings.HasSuffix(mode, "unavailable-tree") {
						gitReadPromisor(t, tools, f, gitReadOK(t, tools, f.Root, "rev-parse", "HEAD^{tree}"))
					}
					suffix := ""
					if runtime.GOOS == "windows" {
						suffix = ".exe"
					}
					hook := filepath.Join(filepath.Dir(os.Getenv("FIREWEED_BIN")), "codex"+suffix)
					program, e := os.ReadFile(hook)
					if e != nil {
						t.Fatal(e)
					}
					if e = os.WriteFile(filepath.Join(filepath.Dir(hook), "false"+suffix), program, 0755); e != nil {
						t.Fatal(e)
					}
					if !filter {
						gitReadOK(t, tools, f.Root, "config", "core.fsmonitor", hook)
					} else {
						p := gitReadPath(t, tools, f.Root, "info/attributes")
						write(t, filepath.Dir(p), filepath.Base(p), "tracked.txt filter=marker\n")
						key := "filter.marker.clean"
						if strings.HasPrefix(mode, "process-") {
							key = "filter.marker.process"
						}
						gitReadOK(t, tools, f.Root, "config", key, hook)
					}
					gitReadReset(t, tools)
					if !filter {
						gitReadCommand(t, tools, f.Root, nil, "rev-parse", "--shared-index-path")
					} else {
						// This control reaches conversion independently of HEAD availability, and
						// the same timestamp-forced file comparison is present for the API call.
						gitReadCommand(t, tools, f.Root, []byte("new content\n"), "hash-object", "--stdin", "--path=tracked.txt")
					}
					raw, e := os.ReadFile(tools.Trap)
					if e != nil || !strings.Contains(string(raw), "codex") {
						t.Fatalf("native configured callback is not live: %s %v", raw, e)
					}
					// A dirty index/worktree can establish false without reading HEAD's tree.
					// Keep that valid alternative; every callback/fetch remains forbidden.
					cleanAPI(t, tools, f, f.Root, want, filter && unavailable, !filter && unavailable)
				})
			}
		}
	}
}

func Test_GIT012CleanUnbornNestedAndInvalid(t *testing.T) {
	tools := gitReadPrepare(t)
	for _, format := range []string{"sha1", "sha256"} {
		t.Run(format+"/unborn", func(t *testing.T) {
			root := t.TempDir()
			gitReadOK(t, tools, root, "init", "-q", "--initial-branch=main", "--object-format="+format)
			opened, e := g.Open(root)
			if e != nil {
				t.Fatal(e)
			}
			f := gitReadFixture{Root: opened.Root, Common: opened.CommonDir, GitDir: opened.GitDir}
			if !cleanNative(t, tools, f.Root) {
				t.Fatal("native unborn clean premise failed")
			}
			cleanAPI(t, tools, f, f.Root, true, false)
			write(t, f.Root, "untracked λ", "unborn data\n")
			if cleanNative(t, tools, f.Root) {
				t.Fatal("native unborn dirty premise failed")
			}
			cleanAPI(t, tools, f, f.Root, false, false)
			gitReadOK(t, tools, f.Root, "add", "--", "untracked λ")
			if cleanNative(t, tools, f.Root) {
				t.Fatal("native staged unborn dirty premise failed")
			}
			cleanAPI(t, tools, f, f.Root, false, false)
		})
		for _, linked := range []bool{false, true} {
			for _, mode := range []string{"nested-clean", "nested-dirty", "invalid-root", "missing-root"} {
				t.Run(fmt.Sprintf("%s/linked_%t/%s", format, linked, mode), func(t *testing.T) {
					f := gitReadRepository(t, tools, format, linked, "")
					root := filepath.Join(f.Root, "nested")
					want := mode != "nested-dirty"
					invalid := strings.Contains(mode, "root")
					if mode == "nested-dirty" {
						write(t, f.Root, "tracked.txt", "dirty from nested cwd\n")
					}
					invalidParent := ""
					var invalidBefore gitReadTree
					if invalid {
						root = t.TempDir()
						invalidParent = root
						write(t, invalidParent, "retained invalid-root user file", "preserve invalid-root bytes\n")
						invalidBefore = gitReadTreeImage(t, invalidParent, false)
						if mode == "missing-root" {
							root = filepath.Join(root, "missing")
						}
					} else if cleanNative(t, tools, root) != want {
						t.Fatal("native nested-root premise failed")
					}
					cleanAPI(t, tools, f, root, want, invalid)
					if invalid && !reflect.DeepEqual(invalidBefore, gitReadTreeImage(t, invalidParent, false)) {
						t.Error("Clean changed invalid-root user files or created the missing root")
					}
				})
			}
		}
	}
}
