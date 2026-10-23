package evals

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	g "github.com/lurkshark/fireweed/packages/gitstore"
)

func Test_GIT002PreferredRefNeverFetches(t *testing.T) {
	tools := gitReadPrepare(t)
	for _, format := range []string{"sha1", "sha256"} {
		for _, linked := range []bool{false, true} {
			for _, mode := range []string{"explicit-tracking", "remote-default", "available-tracking", "absent-tracking", "default-fallback", "invalid-operand"} {
				t.Run(fmt.Sprintf("%s/linked_%t/%s", format, linked, mode), func(t *testing.T) {
					f := gitReadRepository(t, tools, format, linked, "")
					base, remote, want := "main", "", "refs/remotes/origin/main"
					missing := f.Old
					switch mode {
					case "explicit-tracking":
						gitReadOK(t, tools, f.Root, "update-ref", want, f.Old)
					case "remote-default":
						base = ""
						remote = "upstream"
						want = "refs/remotes/upstream/team/trunk"
						gitReadOK(t, tools, f.Root, "update-ref", want, f.Old)
						gitReadOK(t, tools, f.Root, "symbolic-ref", "refs/remotes/upstream/HEAD", want)
					case "available-tracking":
						missing = f.OldBlob
						gitReadOK(t, tools, f.Root, "update-ref", want, f.Head)
					case "absent-tracking":
						base = "team/local"
						remote = "upstream"
						want = "refs/heads/team/local"
					case "default-fallback":
						base = ""
						want = "refs/heads/main"
					case "invalid-operand":
						base = "--help"
						want = ""
					}
					gitReadPromisor(t, tools, f, missing)
					temporary := gitReadTemporary(t)
					before := gitReadStateImage(t, f, temporary)
					got, err := g.BaseRef(f.Root, base, remote)
					gitReadNoTransport(t, tools)
					if mode == "invalid-operand" {
						if err == nil {
							t.Errorf("invalid base accepted: %q %v", got, err)
						}
					} else if err != nil || got != want {
						t.Errorf("preferred local ref got %q/%v, want %q", got, err, want)
					}
					gitReadSameState(t, f, temporary, before)
				})
			}
		}
	}
}

func Test_GIT011EditObjectLocality(t *testing.T) {
	tools := gitReadPrepare(t)
	modes := []string{"required-commit", "required-tag", "required-tag-target", "required-ancestor", "required-root-tree", "required-nested-tree", "required-rename-body", "available-modification", "available-tag", "available-nested-tag", "available-ancestor", "available-exact-rename", "irrelevant-parent", "irrelevant-tree"}
	for _, format := range []string{"sha1", "sha256"} {
		for _, linked := range []bool{false, true} {
			for _, mode := range modes {
				t.Run(fmt.Sprintf("%s/linked_%t/%s", format, linked, mode), func(t *testing.T) {
					rename := ""
					if mode == "required-rename-body" {
						rename = "fuzzy"
					}
					if mode == "available-exact-rename" {
						rename = "exact"
					}
					f := gitReadRepository(t, tools, format, linked, rename)
					base, missing, reject := f.Old, f.Old, true
					switch mode {
					case "required-tag":
						base = f.Tag
						missing = f.TagObject
					case "required-tag-target":
						base = "selected-chain"
					case "required-ancestor":
						base = "HEAD~1"
					case "required-root-tree":
						missing = f.OldTree
					case "required-nested-tree":
						missing = f.OldNested
					case "required-rename-body":
						missing = gitReadOK(t, tools, f.Root, "rev-parse", f.Old+":rename old.txt")
					case "available-modification":
						missing = f.OldBlob
						reject = false
					case "available-tag":
						base = f.Tag
						missing = f.OldBlob
						reject = false
					case "available-nested-tag":
						base = "selected-chain"
						missing = f.OldBlob
						reject = false
					case "available-ancestor":
						base = "HEAD~1"
						missing = f.OldBlob
						reject = false
					case "available-exact-rename":
						missing = gitReadOK(t, tools, f.Root, "rev-parse", f.Old+":rename old.txt")
						reject = false
					case "irrelevant-parent":
						base = f.Head
						reject = false
					case "irrelevant-tree":
						base = f.Head
						missing = f.OldNested
						reject = false
					}
					want := gitReadNativeEdits(t, tools, f.Root, base)
					if rename != "" {
						found := false
						for _, e := range want {
							if e.Status == "R" && e.Path == "rename new.txt" && e.OldPath == "rename old.txt" {
								found = true
							}
						}
						if !found {
							t.Fatalf("fixture available native %s rename did not detect expected old/new paths: %+v", rename, want)
						}
					}
					if strings.HasPrefix(mode, "available-") && len(want) == 0 {
						t.Fatal("fixture available older-baseline comparison is vacuous")
					}
					if strings.HasPrefix(mode, "irrelevant-") && len(want) != 0 {
						t.Fatalf("fixture clean current comparison is not empty: %+v", want)
					}
					gitReadPromisor(t, tools, f, missing)
					if !reject {
						native := gitReadNativeEdits(t, tools, f.Root, base)
						gitReadNoTransport(t, tools)
						if !reflect.DeepEqual(native, want) {
							t.Fatalf("fixture unavailable unrelated object changes native result: %+v vs %+v", native, want)
						}
						gitReadReset(t, tools)
					}
					temporary := gitReadTemporary(t)
					before := gitReadStateImage(t, f, temporary)
					got, err := g.Edits(f.Root, base)
					gitReadNoTransport(t, tools)
					if reject {
						if err == nil || strings.TrimSpace(err.Error()) == "" || len(got) != 0 {
							t.Errorf("unavailable required object returned partial/successful edits: %+v %v", got, err)
						}
					} else if err != nil || !reflect.DeepEqual(got, want) {
						t.Errorf("available native comparison got %+v/%v, want %+v", got, err, want)
					}
					gitReadSameState(t, f, temporary, before)
				})
			}
		}
	}
}

func Test_GIT011SingleCommitSelectors(t *testing.T) {
	tools := gitReadPrepare(t)
	for _, format := range []string{"sha1", "sha256"} {
		for _, linked := range []bool{false, true} {
			t.Run(fmt.Sprintf("%s/linked_%t", format, linked), func(t *testing.T) {
				f := gitReadRepository(t, tools, format, linked, "")
				temporary := gitReadTemporary(t)
				for _, base := range []string{"HEAD~1..HEAD", "HEAD^!", "HEAD^@", "--help", "", "missing-revision", "HEAD; fireweed; #", "HEAD$(fireweed)", "HEAD`fireweed`", "HEAD & fireweed & rem "} {
					t.Run(fmt.Sprintf("reject_%q", base), func(t *testing.T) {
						gitReadReset(t, tools)
						before := gitReadStateImage(t, f, temporary)
						got, err := g.Edits(f.Root, base)
						gitReadNoTransport(t, tools)
						if err == nil || len(got) != 0 {
							t.Errorf("non-single baseline accepted %q: %+v %v", base, got, err)
						}
						gitReadSameState(t, f, temporary, before)
					})
				}
			})
		}
	}
}

func Test_GIT011LocalityObserverControls(t *testing.T) {
	tools := gitReadPrepare(t)
	t.Run("external_command_trap_is_live", func(t *testing.T) {
		gitReadReset(t, tools)
		cmd := exec.Command(os.Getenv("FIREWEED_BIN"))
		err := cmd.Run()
		if err == nil {
			t.Fatal("marker command unexpectedly succeeded")
		}
		raw, e := os.ReadFile(tools.Trap)
		if e != nil || !strings.Contains(string(raw), "fireweed-bootstrap") {
			t.Fatalf("external command marker is not live: %s %v", raw, e)
		}
		gitReadReset(t, tools)
	})
	t.Run("operation_arguments_are_not_commands", func(t *testing.T) {
		f := gitReadRepository(t, tools, "sha1", false, "")
		gitReadReset(t, tools)
		gitReadOK(t, tools, f.Root, "rev-parse", "--sq-quote", "fetch", "push", "branch")
		gitReadOK(t, tools, f.Root, "branch", "--format", "%(refname)")
		gitReadOK(t, tools, f.Root, "branch", "--sort", "refname")
		gitReadNoTransport(t, tools)
		write(t, f.Root, "-b", "literal checkout path\n")
		gitReadOK(t, tools, f.Root, "add", "--", "-b")
		gitReadOK(t, tools, f.Root, "-c", "user.name=Local Git Eval", "-c", "user.email=local@eval.invalid", "commit", "-qm", "literal option path")
		gitReadOK(t, tools, f.Root, "update-ref", "refs/heads/-c", "HEAD")
		gitReadReset(t, tools)
		gitReadOK(t, tools, f.Root, "checkout", "--", "-b")
		gitReadOK(t, tools, f.Root, "switch", "--", "-c")
		gitReadNoTransport(t, tools)
		if op, _ := gitReadOperation([]string{"git", "-c", "user.name=fetch", "-C", "fetch", "rev-parse", "--sq-quote", "fetch"}); op != "rev-parse" {
			t.Fatalf("observer misclassified operand as command: %s", op)
		}
		if op, _ := gitReadOperation([]string{"git", "-c", "fetch.negotiationAlgorithm=noop", "fetch", "origin"}); op != "fetch" {
			t.Fatalf("observer missed actual fetch: %s", op)
		}
	})
	t.Run("allowed_and_denied_actual_fetch", func(t *testing.T) {
		f := gitReadRepository(t, tools, "sha1", false, "")
		gitReadPromisor(t, tools, f, f.OldBlob)
		gitReadNoTransport(t, tools)
	})
	t.Run("transient_branch_and_pack_commands", func(t *testing.T) {
		f := gitReadRepository(t, tools, "sha1", false, "")
		for _, kind := range []string{"branch", "stdin", "fetch-pack", "send-pack", "remote-verbose", "remote-short-verbose"} {
			t.Run(kind, func(t *testing.T) {
				gitReadReset(t, tools)
				switch kind {
				case "branch":
					gitReadOK(t, tools, f.Root, "branch", "transient-control")
					gitReadOK(t, tools, f.Root, "branch", "-D", "transient-control")
				case "stdin":
					r := gitReadCommand(t, tools, f.Root, []byte("create refs/heads/transient-control "+f.Head+"\n"), "update-ref", "--stdin")
					if r.Exit != 0 {
						t.Fatal(string(r.Err))
					}
					r = gitReadCommand(t, tools, f.Root, []byte("delete refs/heads/transient-control\n"), "update-ref", "--stdin")
					if r.Exit != 0 {
						t.Fatal(string(r.Err))
					}
				case "remote-verbose", "remote-short-verbose":
					flag := "--verbose"
					if kind == "remote-short-verbose" {
						flag = "-v"
					}
					r := gitReadCommand(t, tools, f.Root, nil, "remote", flag, "update", "absent-control-remote")
					if r.Exit == 0 {
						t.Fatal("nonexistent remote unexpectedly succeeded")
					}
				default:
					r := gitReadCommand(t, tools, f.Root, nil, kind)
					if r.Exit == 0 {
						t.Fatal("missing-operand command control unexpectedly succeeded")
					}
				}
				if len(gitReadNetworkEvents(t, tools)) == 0 {
					t.Fatal("actual forbidden operation was not observed")
				}
				gitReadReset(t, tools)
				gitReadOK(t, tools, f.Root, "branch", "--list", "transient-control")
				gitReadNoTransport(t, tools)
			})
		}
	})
	t.Run("state_observer_detects_mutations", func(t *testing.T) {
		f := gitReadRepository(t, tools, "sha1", true, "")
		gitReadOK(t, tools, f.Root, "config", "extensions.worktreeConfig", "true")
		gitReadOK(t, tools, f.Root, "config", "--worktree", "core.autocrlf", "false")
		temporary := gitReadTemporary(t)
		for _, p := range []string{filepath.Join(f.Root, "tracked.txt"), filepath.Join(f.Common, "HEAD"), filepath.Join(f.GitDir, "index"), filepath.Join(f.Common, "config"), filepath.Join(f.GitDir, "config.worktree"), filepath.Join(f.GitDir, "HEAD"), filepath.Join(f.Common, "refs", "heads", "main"), filepath.Join(f.Root, ".git"), filepath.Join(temporary, "user", "retained.txt")} {
			before := gitReadStateImage(t, f, temporary)
			raw, e := os.ReadFile(p)
			if e != nil {
				t.Fatal(e)
			}
			info, e := os.Stat(p)
			if e != nil {
				t.Fatal(e)
			}
			if e = os.WriteFile(p, append(append([]byte(nil), raw...), []byte("mutation control\n")...), info.Mode().Perm()); e != nil {
				t.Fatal(e)
			}
			after := gitReadStateImage(t, f, temporary)
			if reflect.DeepEqual(before, after) {
				t.Fatalf("state observer missed file change %s", p)
			}
			if e = os.WriteFile(p, raw, info.Mode().Perm()); e != nil {
				t.Fatal(e)
			}
			if !reflect.DeepEqual(before, gitReadStateImage(t, f, temporary)) {
				t.Fatal("state observer control did not restore original")
			}
		}
		for _, parent := range []string{f.Common, f.GitDir} {
			before := gitReadStateImage(t, f, temporary)
			p := filepath.Join(parent, "owned-empty-directory")
			if e := os.Mkdir(p, 0700); e != nil {
				t.Fatal(e)
			}
			if len(gitReadNewScratch(before, gitReadStateImage(t, f, temporary))) == 0 {
				t.Fatal("cleanup observer missed repository-local owned scratch")
			}
			if e := os.Remove(p); e != nil {
				t.Fatal(e)
			}
			if len(gitReadNewScratch(before, gitReadStateImage(t, f, temporary))) != 0 {
				t.Fatal("cleanup observer failed to restore repository-local control")
			}
		}
		// Real immutable object additions remain allowed, but native temporary
		// object/pack files and locks must be cleaned like other owned scratch.
		beforeCache := gitReadStateImage(t, f, temporary)
		object := gitReadCommand(t, tools, f.Root, []byte("independent local object layout control\n"), "hash-object", "-w", "--stdin")
		if object.Exit != 0 {
			t.Fatal(string(object.Err))
		}
		oid := strings.TrimSpace(string(object.Out))
		if len(oid) != 40 {
			t.Fatal("native object layout control missing SHA1", oid)
		}
		if _, e := os.Stat(filepath.Join(f.Common, "objects", oid[:2], oid[2:])); e != nil {
			t.Fatal(e)
		}
		if len(gitReadNewScratch(beforeCache, gitReadStateImage(t, f, temporary))) != 0 {
			t.Fatal("cleanup observer overconstrains actual native object layout")
		}
		beforeShared := gitReadStateImage(t, f, temporary)
		gitReadOK(t, tools, f.Root, "update-index", "--split-index")
		foundShared := false
		for path := range gitReadTreeImage(t, f.Common, false) {
			if gitReadSharedIndex(path) {
				foundShared = true
			}
		}
		if !foundShared {
			t.Fatal("native shared-index cache control did not create a shared index")
		}
		if scratch := gitReadNewScratch(beforeShared, gitReadStateImage(t, f, temporary)); len(scratch) != 0 {
			t.Fatalf("cleanup observer overconstrains native linked shared-index storage: %v", scratch)
		}
		for _, relative := range []string{"objects/pack/tmp_pack_control", "objects/ab/tmp_obj_control", "objects/info/commit-graph.lock"} {
			p := filepath.Join(f.Common, filepath.FromSlash(relative))
			if e := os.MkdirAll(filepath.Dir(p), 0755); e != nil {
				t.Fatal(e)
			}
			before := gitReadStateImage(t, f, temporary)
			if e := os.WriteFile(p, []byte("owned temporary object control"), 0600); e != nil {
				t.Fatal(e)
			}
			if len(gitReadNewScratch(before, gitReadStateImage(t, f, temporary))) == 0 {
				t.Fatalf("cleanup observer missed native temporary object %s", relative)
			}
			if e := os.Remove(p); e != nil {
				t.Fatal(e)
			}
			if len(gitReadNewScratch(before, gitReadStateImage(t, f, temporary))) != 0 {
				t.Fatal("native temporary control failed to restore")
			}
		}
		before := gitReadStateImage(t, f, temporary)
		p := filepath.Join(temporary, "owned-empty-directory")
		if e := os.Mkdir(p, 0700); e != nil {
			t.Fatal(e)
		}
		if reflect.DeepEqual(before, gitReadStateImage(t, f, temporary)) {
			t.Fatal("state observer missed empty leaked directory")
		}
		if e := os.Remove(p); e != nil {
			t.Fatal(e)
		}
	})
}
