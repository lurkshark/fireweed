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

func approvalGitObserverControls(t *testing.T, root, dir string, p approvalGitPolicy, private string, current bool) {
	t.Helper()
	proxy := filepath.Join(dir, "git")
	run := func(binary, input string, args ...string) ([]byte, int) {
		t.Helper()
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		c := exec.CommandContext(ctx, binary, args...)
		c.Dir = root
		c.Stdin = strings.NewReader(input)
		out, e := c.CombinedOutput()
		if ctx.Err() != nil {
			t.Fatal(ctx.Err())
		}
		code := 0
		if e != nil {
			x, ok := e.(*exec.ExitError)
			if !ok {
				t.Fatal(e)
			}
			code = x.ExitCode()
		}
		return out, code
	}
	clear := func() {
		t.Helper()
		if e := os.Remove(p.Marker); e != nil && !os.IsNotExist(e) {
			t.Fatal(e)
		}
	}
	if !p.RejectAll {
		allowed := [][]string{{"cat-file", "commit", p.Base}, {"-c", "example.value=commit", "cat-file", "commit", p.Base}, {"cat-file", "--batch-all-objects", "--batch-check=%(objectname)"}, {"rev-list", "--no-walk", "--missing=print", "--end-of-options", p.Base}, {"show", p.Base + ":fireweed.yaml"}, {"cat-file", "blob", p.Base + ":fireweed.yaml"}}
		ids := []string{}
		for id := range p.Trees {
			ids = append(ids, id)
		}
		sort.Strings(ids)
		if len(ids) < 2 {
			t.Fatal("nonempty root/subtree fixture premise")
		}
		for _, id := range ids {
			allowed = append(allowed, []string{"ls-tree", id})
		}
		if !current {
			allowed = append(allowed, []string{"cat-file", "blob", p.Selected}, []string{"show", p.Selected}, []string{"cat-file", "-t", p.Selected})
		}
		for _, args := range allowed {
			want, code := run(p.Real, "", args...)
			if code != 0 {
				t.Fatalf("real allowed source premise %v: %d %s", args, code, want)
			}
			got, code := run(proxy, "", args...)
			if code != 0 || string(got) != string(want) {
				t.Fatalf("allowed source %v: %d %s/%s", args, code, got, want)
			}
			if _, e := os.Stat(p.Marker); !os.IsNotExist(e) {
				t.Fatal("allowed metadata/body read flagged", args, e)
			}
		}

		queries := []string{p.Base, p.Base + ":fireweed.yaml"}
		if !current {
			queries = append(queries, p.Selected)
		}
		for _, query := range queries {
			args := []string{"cat-file", "--batch-check=%(objectname)"}
			input := query + "\n"
			want, code := run(p.Real, input, args...)
			if code != 0 || len(want) == 0 {
				t.Fatal("selected metadata batch premise", code, string(want))
			}
			got, code := run(proxy, input, args...)
			if code != 0 || string(got) != string(want) {
				t.Fatal("selected metadata batch rejected", code, string(got))
			}
			if _, e := os.Stat(p.Marker); !os.IsNotExist(e) {
				t.Fatal(e)
			}
		}
		// --batch-command is newer than the installed Git on some platforms. Only
		// claim its positive/negative calibration when real Git supports the mode.
		commandArgs := []string{"cat-file", "--batch-command"}
		input := "info " + p.Base + "\n"
		want, code := run(p.Real, input, commandArgs...)
		if code == 0 {
			got, status := run(proxy, input, commandArgs...)
			if status != 0 || string(got) != string(want) {
				t.Fatal("selected command metadata rejected", status, string(got))
			}
			if _, e := os.Stat(p.Marker); !os.IsNotExist(e) {
				t.Fatal(e)
			}
			deniedInput := "info refs/fireweed/claims/unrelated-malformed\n"
			readable, status := run(p.Real, deniedInput, commandArgs...)
			if status != 0 || len(readable) == 0 {
				t.Fatal("unrelated info premise", status, string(readable))
			}
			out, status := run(proxy, deniedInput, commandArgs...)
			if status != 93 || len(out) != 0 {
				t.Fatal("unrelated info metadata accepted", status, string(out))
			}
			if _, e := os.Stat(p.Marker); e != nil {
				t.Fatal(e)
			}
			clear()
		} else if !strings.Contains(string(want), "unknown option") && !strings.Contains(string(want), "unknown switch") {
			t.Fatalf("unexpected batch-command capability failure %d %s", code, want)
		}

	}
	// Private/mutable/other-claim controls must be real readable fixture objects;
	// otherwise an underlying Git failure could masquerade as a working observer.
	for _, q := range []struct {
		Args  []string
		Input string
	}{
		{[]string{"cat-file", "blob", private}, ""},
		{[]string{"cat-file", "--batch"}, private + "\n"},
		{[]string{"cat-file", "--batch-check=%(objectname)"}, "refs/fireweed/claims/unrelated-malformed\n"},
		{[]string{"cat-file", "--batch-all-objects", "--batch"}, ""},
		{[]string{"show", "HEAD:fireweed.yaml"}, ""},
		{[]string{"cat-file", "blob", "refs/fireweed/claims/unrelated-malformed"}, ""},
		{[]string{"for-each-ref", "--format=%(refname)", "refs/fireweed/claims/"}, ""},
	} {
		clear()
		want, code := run(p.Real, q.Input, q.Args...)
		if code != 0 || len(want) == 0 {
			t.Fatalf("forbidden source readable premise %v: %d %s", q.Args, code, want)
		}
		out, code := run(proxy, q.Input, q.Args...)
		if code != 93 || len(out) != 0 {
			t.Fatalf("forbidden source observer %v: %d %s", q.Args, code, out)
		}
		if _, e := os.Stat(p.Marker); e != nil {
			t.Fatal("forbidden source marker absent", q.Args, e)
		}
		clear()
	}
	for _, args := range [][]string{{"commit", "--allow-empty", "-m", "forbidden"}, {"fetch", "invalid"}, {"-c", "example.value=cat-file", "push", "invalid"}, {"symbolic-ref", "HEAD", "refs/heads/forbidden"}} {
		out, code := run(proxy, "", args...)
		if code != 93 || len(out) != 0 {
			t.Fatalf("actual mutation/network source control %v: %d %s", args, code, out)
		}
		if _, e := os.Stat(p.Marker); e != nil {
			t.Fatal(e)
		}
		clear()
	}
}
