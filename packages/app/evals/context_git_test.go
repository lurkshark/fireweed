package evals

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

// The child Git proxy makes ignored broad/private reads observable independently
// of the native interposer in the calling CLI process. Fixture preparation always
// uses the real binary; only the public API/CLI child receives the proxy PATH.
func cliContextGitProxy(t *testing.T, root string) (directory, marker, loose string) {
	t.Helper()
	real, err := exec.LookPath("git")
	if err != nil {
		t.Fatal(err)
	}
	run := func(binary string, args ...string) (string, int) {
		t.Helper()
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		cmd := exec.CommandContext(ctx, binary, args...)
		cmd.Dir = root
		cmd.Env = append(os.Environ(), "GIT_NO_REPLACE_OBJECTS=1", "GIT_NO_LAZY_FETCH=1")
		out, err := cmd.CombinedOutput()
		if ctx.Err() != nil {
			t.Fatal(ctx.Err())
		}
		code := 0
		if err != nil {
			exit, ok := err.(*exec.ExitError)
			if !ok {
				t.Fatal(err)
			}
			code = exit.ExitCode()
		}
		return string(out), code
	}
	git := func(args ...string) string {
		t.Helper()
		out, code := run(real, args...)
		if code != 0 {
			t.Fatalf("real Git proxy fixture %v exit%d: %s", args, code, out)
		}
		return out
	}
	base := strings.TrimSpace(git("rev-parse", "HEAD"))
	tree := git("ls-tree", "-r", "-z", base)
	allowed := []string{base}
	configObject := ""
	for _, line := range strings.Split(strings.TrimSuffix(tree, "\x00"), "\x00") {
		header, path, ok := strings.Cut(line, "\t")
		fields := strings.Fields(header)
		if !ok || len(fields) != 3 {
			t.Fatalf("tree fixture record %q", line)
		}
		if path == "fireweed.yaml" || (strings.HasPrefix(path, "units/") && strings.Count(path, "/") == 2 && strings.HasSuffix(path, "/SPEC.md")) || (strings.HasPrefix(path, "changes/") && strings.Count(path, "/") == 1) || (strings.HasPrefix(path, "changes/archive/") && strings.Count(path, "/") == 2) {
			allowed = append(allowed, fields[2])
		}
		if path == "fireweed.yaml" {
			configObject = fields[2]
		}
	}
	selected := "refs/fireweed/claims/current"
	claim, claimExit := run(real, "rev-parse", "--verify", selected)
	if claimExit == 0 {
		allowed = append(allowed, strings.TrimSpace(claim))
	} else if claimExit != 128 {
		t.Fatalf("claim fixture status%d %s", claimExit, claim)
	}
	private := strings.TrimSpace(git("rev-parse", base+":units/a/internal/private.txt"))
	common := strings.TrimSpace(git("rev-parse", "--git-common-dir"))
	if !filepath.IsAbs(common) {
		common = filepath.Join(root, common)
	}
	common, err = filepath.EvalSymlinks(common)
	if err != nil {
		t.Fatal(err)
	}
	loose = filepath.Join(common, "objects", private[:2], private[2:])
	if raw, err := os.ReadFile(loose); err != nil || len(raw) == 0 {
		t.Fatalf("private loose-object control %v", err)
	}
	for _, id := range allowed {
		if id == private {
			t.Fatal("private object aliases approved metadata fixture")
		}
	}
	directory = t.TempDir()
	marker = filepath.Join(directory, "forbidden")
	reject := "printf '%s\\n' \"$@\" >> " + coverageQuote(marker) + ";exit 93"
	script := "#!/bin/sh\nop='';skip=0;last='';for a in \"$@\";do last=$a;case \"$a\" in refs/fireweed/claims/current|refs/fireweed/claims/current'^{object}'|refs/fireweed/claims/current'^{}') ;; refs/fireweed/claims*) " + reject + ";;esac;if [ -z \"$op\" ];then if [ \"$skip\" = 1 ];then skip=0;continue;fi;case \"$a\" in -c|--git-dir|--work-tree) skip=1;; -C) " + reject + ";; -*) ;; *) op=$a;;esac;fi;done\ncase \"$op\" in rev-parse|symbolic-ref|show-ref|check-ref-format|ls-files) ;; cat-file) case \"$last\" in " + strings.Join(allowed, "|") + ") ;; *) " + reject + ";;esac;; ls-tree) [ \"$last\" = " + coverageQuote(base) + " ] || { " + reject + "; };; *) " + reject + ";;esac\nexec " + coverageQuote(real) + " \"$@\"\n"
	projectPut(t, directory, "git", script)
	binary := filepath.Join(directory, "git")
	if err := os.Chmod(binary, 0755); err != nil {
		t.Fatal(err)
	}
	forbidden := [][]string{
		{"cat-file", "blob", private},
		{"show", base + ":units/a/internal/private.txt"},
		{"ls-tree", "-r", "HEAD"},
		{"for-each-ref", "--format=%(refname)", "refs/fireweed/claims/"},
		{"for-each-ref", "--format=%(refname)"},
		{"rev-parse", "--verify", "refs/fireweed/claims/current-neighbor"},
		{"log", "-p", "--all"},
	}
	for i, args := range forbidden {
		t.Run("git-forbidden-control-"+string(rune('a'+i)), func(t *testing.T) {
			out, code := run(binary, args...)
			if code != 93 || out != "" {
				t.Fatalf("Git denial control %v exit%d %s", args, code, out)
			}
			if _, err := os.Stat(marker); err != nil {
				t.Fatal("unobservable denial", err)
			}
			if err := os.Remove(marker); err != nil {
				t.Fatal(err)
			}
		})
	}
	// Allowed real metadata/object/ref operations must remain byte-exact.
	for _, args := range [][]string{{"cat-file", "blob", configObject}, {"ls-tree", "-r", "-z", base}, {"show-ref", "--verify", "--quiet", selected}} {
		t.Run("git-allowed-control-"+args[0], func(t *testing.T) {
			want, wantExit := run(real, args...)
			out, code := run(binary, args...)
			if code != wantExit || out != want {
				t.Fatalf("allowed metadata differs %v exit%d/%d %q/%q", args, code, wantExit, out, want)
			}
			if _, err := os.Stat(marker); !os.IsNotExist(err) {
				t.Fatal("allowed operation flagged", err)
			}
		})
	}
	return directory, marker, loose
}
