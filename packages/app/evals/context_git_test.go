package evals

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
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
	runInput := func(binary, input string, args ...string) (string, int) {
		t.Helper()
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		cmd := exec.CommandContext(ctx, binary, args...)
		cmd.Dir = root
		cmd.Stdin = strings.NewReader(input)
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
	run := func(binary string, args ...string) (string, int) { return runInput(binary, "", args...) }
	git := func(args ...string) string {
		t.Helper()
		out, code := run(real, args...)
		if code != 0 {
			t.Fatalf("real Git proxy fixture %v exit%d: %s", args, code, out)
		}
		return out
	}
	base := strings.TrimSpace(git("rev-parse", "HEAD"))
	tree := git("ls-tree", "-r", "-t", "-z", base)
	allowed := []string{base}
	treeIDs := map[string]bool{strings.TrimSpace(git("rev-parse", base+"^{tree}")): true}
	paths := map[string]bool{}
	configObject := ""
	for _, line := range strings.Split(strings.TrimSuffix(tree, "\x00"), "\x00") {
		header, path, ok := strings.Cut(line, "\t")
		fields := strings.Fields(header)
		if !ok || len(fields) != 3 {
			t.Fatalf("tree fixture record %q", line)
		}
		if fields[1] == "tree" {
			treeIDs[fields[2]] = true
		}
		if path == "fireweed.yaml" || (strings.HasPrefix(path, "units/") && strings.Count(path, "/") == 2 && strings.HasSuffix(path, "/SPEC.md")) || (strings.HasPrefix(path, "changes/") && strings.Count(path, "/") == 1) || (strings.HasPrefix(path, "changes/archive/") && strings.Count(path, "/") == 2) {
			allowed = append(allowed, fields[2])
			paths[path] = true
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
	content := map[string]bool{}
	for _, id := range allowed {
		content[id] = true
	}
	installContextGitProxy(t, directory, contextGitPolicy{Real: real, Marker: marker, Base: base, Selected: selected, Root: root, Content: content, Trees: treeIDs, Paths: paths})
	binary := filepath.Join(directory, "git")
	forbidden := [][]string{
		{"cat-file", "blob", private},
		{"show", base + ":units/a/internal/private.txt"},
		{"ls-tree", "-r", "HEAD"},
		{"for-each-ref", "--format=%(refname)", "refs/fireweed/claims/"},
		{"for-each-ref", "--format=%(refname)"},
		{"rev-parse", "--verify", "refs/fireweed/claims/current-neighbor"},
		{"log", "-p", "--all"},
		{"show-ref"},
		{"show-ref", "--head"},
		{"rev-parse", "--all"},
		{"rev-parse", "--glob=refs/*"},
		{"rev-parse", "--glob", "refs/*"},
		{"rev-parse", "--branches"},
		{"cat-file", "--batch-all-objects", "--batch"},
		{"rev-parse", "--exclude=refs/heads/*", "--all"},
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
	rootTree := strings.TrimSpace(git("rev-parse", base+"^{tree}"))
	queries := [][]string{{"show", base + ":fireweed.yaml"}, {"cat-file", "blob", rootTree + ":fireweed.yaml"}, {"ls-tree", "-z", base + ":units"}, {"ls-tree", "-z", rootTree + ":units"}, {"ls-tree", "-z", base + "^{tree}"}, {"cat-file", "blob", configObject}, {"ls-tree", "-r", "-z", base}, {"show-ref", "--verify", "--quiet", selected}, {"cat-file", "--batch-all-objects", "--batch-check=%(objectname)"}, {"rev-list", "--no-walk", "--missing=print", "--end-of-options", base}}
	if claimExit == 0 {
		queries = append(queries, []string{"cat-file", "blob", selected}, []string{"cat-file", "blob", selected + "^{object}"}, []string{"show", selected})
	}
	ids := []string{}
	for id := range treeIDs {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	for _, id := range ids {
		queries = append(queries, []string{"ls-tree", "-z", id})
	}
	queries = append(queries, []string{"-C", root, "cat-file", "blob", configObject}, []string{"rev-list", "--no-walk", base})
	for _, args := range queries {
		t.Run("git-allowed-control-"+args[0], func(t *testing.T) {
			want, wantExit := run(real, args...)
			if wantExit != 0 && !(wantExit == 1 && args[0] == "show-ref" && claimExit == 128) {
				t.Fatalf("allowed real operation did not execute successfully: %v exit%d %s", args, wantExit, want)
			}
			out, code := run(binary, args...)
			if code != wantExit || out != want {
				t.Fatalf("allowed metadata differs %v exit%d/%d %q/%q", args, code, wantExit, out, want)
			}
			if _, err := os.Stat(marker); !os.IsNotExist(err) {
				t.Fatal("allowed operation flagged", err)
			}
		})
	}
	t.Run("git-allowed-selected-batch-control", func(t *testing.T) {
		args := []string{"cat-file", "--batch"}
		want, wantExit := runInput(real, configObject+"\n", args...)
		if wantExit != 0 || want == "" {
			t.Fatalf("selected batch premise exit%d %s", wantExit, want)
		}
		got, code := runInput(binary, configObject+"\n", args...)
		if code != wantExit || got != want {
			t.Fatalf("selected batch differs exit%d/%d %q/%q", code, wantExit, got, want)
		}
		if _, e := os.Stat(marker); !os.IsNotExist(e) {
			t.Fatal("selected batch flagged", e)
		}
	})
	t.Run("git-forbidden-private-batch-control", func(t *testing.T) {
		got, code := runInput(binary, private+"\n", "cat-file", "--batch")
		if code != 93 || got != "" {
			t.Fatalf("private batch denial exit%d %q", code, got)
		}
		if _, e := os.Stat(marker); e != nil {
			t.Fatal("private batch unobserved", e)
		}
		if e := os.Remove(marker); e != nil {
			t.Fatal(e)
		}
	})

	return directory, marker, loose
}
