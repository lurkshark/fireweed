package evals

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"runtime"
	"sort"
	"strings"
	"sync"
	"testing"
	"time"

	w "github.com/lurkshark/fireweed/packages/workflow"
)

func claimCommand(t *testing.T, root, input string, args ...string) (string, error) {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, "git", args...)
	cmd.Dir = root
	cmd.Stdin = strings.NewReader(input)
	cmd.Env = append(os.Environ(), "GIT_CONFIG_NOSYSTEM=1", "GIT_CONFIG_GLOBAL="+os.DevNull, "GIT_OPTIONAL_LOCKS=0")
	out, err := cmd.CombinedOutput()
	return strings.TrimSpace(string(out)), err
}
func claimGit(t *testing.T, root string, args ...string) string {
	t.Helper()
	out, err := claimCommand(t, root, "", args...)
	if err != nil {
		t.Fatalf("fixture git %q: %v %s", args, err, out)
	}
	return out
}
func claimRepo(t *testing.T, format string) string {
	t.Helper()
	root := filepath.Join(t.TempDir(), "repository with ' spaces")
	if err := os.MkdirAll(root, 0755); err != nil {
		t.Fatal(err)
	}
	out, err := claimCommand(t, root, "", "init", "-b", "main", "--object-format="+format)
	if err != nil {
		if format == "sha256" {
			t.Skipf("installed Git lacks SHA-256 repositories: %v %s", err, out)
		}
		t.Fatal(err, out)
	}
	if err := os.WriteFile(filepath.Join(root, "retained"), []byte("tracked\n"), 0644); err != nil {
		t.Fatal(err)
	}
	claimGit(t, root, "add", ".")
	claimGit(t, root, "-c", "user.name=Claim Fixture", "-c", "user.email=claim@example.invalid", "commit", "-m", "fixture")
	return stateCanonicalDirectory(t, root)
}
func claimBase(t *testing.T, root string) string {
	t.Helper()
	return claimGit(t, root, "rev-parse", "HEAD")
}
func claimPayload(t *testing.T, c w.Claim) string {
	t.Helper()
	raw, err := json.Marshal(c)
	if err != nil {
		t.Fatal(err)
	}
	return string(raw)
}
func claimRaw(t *testing.T, root, slug, raw string) string {
	t.Helper()
	id, err := claimCommand(t, root, raw, "hash-object", "-w", "--stdin")
	if err != nil {
		t.Fatal(err, id)
	}
	claimGit(t, root, "update-ref", "refs/fireweed/claims/"+slug, id)
	return id
}
func claimSeed(t *testing.T, root, slug, state string) w.Claim {
	t.Helper()
	c := w.Claim{Worktree: root, Change: slug, Run: "fixture-run", Base: claimBase(t, root), State: state}
	c.Object = claimRaw(t, root, slug, claimPayload(t, c))
	return c
}
func claimError(t *testing.T, got w.Claim, err error) {
	t.Helper()
	if got != (w.Claim{}) {
		t.Fatalf("partial Claim on error: %+v", got)
	}
	if err == nil || !strings.Contains(err.Error(), "P02") || !policyCorrective(err.Error()) {
		t.Fatalf("missing P02 corrective error: %v", err)
	}
}
func claimReleaseError(t *testing.T, err error) {
	t.Helper()
	if err == nil || !strings.Contains(err.Error(), "P02") || !policyCorrective(err.Error()) {
		t.Fatalf("missing corrective P02 release error: %v", err)
	}
}
func claimAssert(t *testing.T, root string, got w.Claim, slug, run, base, state string) {
	t.Helper()
	want := w.Claim{Worktree: root, Change: slug, Run: run, Base: base, State: state, Object: got.Object}
	if got != want || got.Object == "" {
		t.Fatalf("claim %+v; want %+v with a blob token", got, want)
	}
	if claimGit(t, root, "cat-file", "-t", got.Object) != "blob" || claimGit(t, root, "rev-parse", "refs/fireweed/claims/"+slug) != got.Object {
		t.Fatal("token is not the stored blob")
	}
	raw := claimGit(t, root, "cat-file", "blob", got.Object)
	var values map[string]any
	if err := json.Unmarshal([]byte(raw), &values); err != nil {
		t.Fatal(err)
	}
	if len(values) != 5 || values["worktree"] != root || values["change"] != slug || values["run"] != run || values["base"] != base || values["state"] != state {
		t.Fatalf("stored wire fields %+v", values)
	}
}
func claimPeer(t *testing.T, root string) string {
	t.Helper()
	peer := filepath.Join(t.TempDir(), "linked peer ' spaces")
	claimGit(t, root, "worktree", "add", "-b", "peer", peer)
	return stateCanonicalDirectory(t, peer)
}
func Test_FLOW008CreateClaims(t *testing.T) {
	for _, format := range []string{"sha1", "sha256"} {
		t.Run(format, func(t *testing.T) {
			root := claimRepo(t, format)
			base := claimBase(t, root)
			before := claimGit(t, root, "show-ref")
			run, state := " run ' unicode λ ", "line one\n---\nopaque \x00 state"
			nested := filepath.Join(root, "nested")
			if err := os.Mkdir(nested, 0755); err != nil {
				t.Fatal(err)
			}
			got, err := w.CreateClaim(nested, "alpha-0--", run, "HEAD", state)
			if err != nil {
				t.Fatal(err)
			}
			claimAssert(t, root, got, "alpha-0--", run, base, state)
			loaded, err := w.LoadClaim(root, "alpha-0--")
			if err != nil || loaded != got {
				t.Fatalf("load %+v %v", loaded, err)
			}
			if claimBase(t, root) != base || !strings.Contains(claimGit(t, root, "show-ref"), before) {
				t.Fatal("HEAD or original branch changed")
			}
			empty, err := w.CreateClaim(root, "reservation", "second", base, "")
			if err != nil {
				t.Fatal(err)
			}
			claimAssert(t, root, empty, "reservation", "second", base, "")
		})
	}
}
func Test_FLOW008InvalidCreateAndConflict(t *testing.T) {
	root := claimRepo(t, "sha1")
	base := claimBase(t, root)
	for _, slug := range []string{"", "-leading", "Upper", "../escape", "a/b", "a..b", "a.lock", "a b", "a\x00b", "a\\b"} {
		t.Run(fmt.Sprintf("slug_%q", slug), func(t *testing.T) {
			before := claimGit(t, root, "show-ref")
			got, err := w.CreateClaim(root, slug, "run", base, "state")
			claimError(t, got, err)
			if claimGit(t, root, "show-ref") != before {
				t.Fatal("invalid slug wrote a ref")
			}
		})
	}
	for _, run := range []string{"", " \t\n"} {
		t.Run(fmt.Sprintf("run_%q", run), func(t *testing.T) {
			before := claimGit(t, root, "show-ref")
			got, err := w.CreateClaim(root, "valid", run, base, "state")
			claimError(t, got, err)
			if claimGit(t, root, "show-ref") != before {
				t.Fatal("invalid run created a claim ref")
			}
		})
	}
	for _, base := range []string{"", "missing-commit", "--help"} {
		t.Run("base_"+base, func(t *testing.T) {
			got, err := w.CreateClaim(root, "valid", "run", base, "state")
			claimError(t, got, err)
		})
	}
	got, err := w.CreateClaim(t.TempDir(), "valid", "run", base, "state")
	claimError(t, got, err)
	existing := claimSeed(t, root, "taken", "approved")
	got, err = w.CreateClaim(root, "taken", "new-run", base, "new")
	claimError(t, got, err)
	if !strings.Contains(err.Error(), root) || !strings.Contains(strings.ToLower(err.Error()), "release") {
		t.Fatalf("conflict omits owner/release guidance: %v", err)
	}
	if claimGit(t, root, "rev-parse", "refs/fireweed/claims/taken") != existing.Object {
		t.Fatal("conflict replaced owner")
	}
}
func Test_FLOW008ConcurrentWorktreeClaim(t *testing.T) {
	root := claimRepo(t, "sha1")
	peer := claimPeer(t, root)
	base := claimBase(t, root)
	start := make(chan struct{})
	type result struct {
		c   w.Claim
		err error
	}
	results := make(chan result, 12)
	var wg sync.WaitGroup
	for i := 0; i < 12; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-start
			dir := root
			if i%2 != 0 {
				dir = peer
			}
			c, err := w.CreateClaim(dir, "contended", fmt.Sprintf("run-%d", i), base, "reservation")
			results <- result{c, err}
		}(i)
	}
	close(start)
	wg.Wait()
	close(results)
	winners := []w.Claim{}
	for r := range results {
		if r.err == nil {
			winners = append(winners, r.c)
		} else {
			claimError(t, r.c, r.err)
		}
	}
	if len(winners) != 1 {
		t.Fatalf("atomic create winners: %+v", winners)
	}
	for _, dir := range []string{root, peer} {
		loaded, err := w.LoadClaim(dir, "contended")
		if err != nil || loaded != winners[0] {
			t.Fatalf("shared visibility: %+v %v", loaded, err)
		}
	}
	other, err := w.CreateClaim(peer, "independent", "other", base, "")
	if err != nil {
		t.Fatal(err)
	}
	claimAssert(t, peer, other, "independent", "other", base, "")
}
func Test_FLOW009UpdateAndRelease(t *testing.T) {
	for _, format := range []string{"sha1", "sha256"} {
		t.Run(format, func(t *testing.T) {
			root := claimRepo(t, format)
			previous := claimSeed(t, root, "current", "old")
			saved := previous
			claimGit(t, root, "-c", "user.name=Claim Fixture", "-c", "user.email=claim@example.invalid", "commit", "--allow-empty", "-m", "next controller phase")
			base := claimBase(t, root)
			next, err := w.UpdateClaim(root, previous, "HEAD", "new\nphase\x00")
			if err != nil {
				t.Fatal(err)
			}
			claimAssert(t, root, next, "current", previous.Run, base, "new\nphase\x00")
			if previous != saved || next.Object == previous.Object {
				t.Fatal("input mutation or stale token retained after content change")
			}
			stale, err := w.UpdateClaim(root, previous, base, "stale overwrite")
			claimError(t, stale, err)
			claimReleaseError(t, w.ReleaseClaim(root, "current", previous.Object))
			if claimGit(t, root, "rev-parse", "refs/fireweed/claims/current") != next.Object {
				t.Fatal("stale operation replaced new claim")
			}
			unchanged, err := w.UpdateClaim(root, next, base, next.State)
			if err != nil || unchanged.Worktree != next.Worktree || unchanged.Change != next.Change || unchanged.Run != next.Run || unchanged.Base != base || unchanged.State != next.State || unchanged.Object == "" {
				t.Fatalf("unchanged update %+v %v", unchanged, err)
			}
			if err = w.ReleaseClaim(root, "current", unchanged.Object); err != nil {
				t.Fatal(err)
			}
			if _, err = claimCommand(t, root, "", "show-ref", "--verify", "refs/fireweed/claims/current"); err == nil {
				t.Fatal("release kept claim")
			}
			claimReleaseError(t, w.ReleaseClaim(root, "current", unchanged.Object))
		})
	}
}
func Test_FLOW009TamperedAndForeignUpdates(t *testing.T) {
	root := claimRepo(t, "sha1")
	original := claimSeed(t, root, "current", "old")
	peer := claimPeer(t, root)
	for _, field := range []string{"Worktree", "Change", "Run", "Base", "State", "Object"} {
		t.Run(field, func(t *testing.T) {
			changed := original
			switch field {
			case "Worktree":
				changed.Worktree = peer
			case "Change":
				changed.Change = "different"
			case "Run":
				changed.Run = "different"
			case "Base":
				changed.Base = "HEAD"
			case "State":
				changed.State = "different"
			case "Object":
				changed.Object = strings.Repeat("0", len(original.Object))
			}
			saved := changed
			got, err := w.UpdateClaim(root, changed, original.Base, "new")
			claimError(t, got, err)
			if changed != saved || claimGit(t, root, "rev-parse", "refs/fireweed/claims/current") != original.Object {
				t.Fatal("tampered update mutated input or claim")
			}
		})
	}
	got, err := w.UpdateClaim(peer, original, original.Base, "new")
	claimError(t, got, err)
	got, err = w.UpdateClaim(root, original, "missing-commit", "new")
	claimError(t, got, err)
	for _, token := range []string{"", "not-an-object", "--help", strings.Repeat("0", len(original.Object))} {
		claimReleaseError(t, w.ReleaseClaim(root, "current", token))
	}
	if err = w.ReleaseClaim(peer, "current", original.Object); err != nil {
		t.Fatalf("explicit cross-worktree release: %v", err)
	}
}
func Test_FLOW009RacingUpdatesAndRelease(t *testing.T) {
	root := claimRepo(t, "sha1")
	for i := 0; i < 6; i++ {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			previous := claimSeed(t, root, "race", "old")
			gate := make(chan struct{})
			errs := make(chan error, 2)
			go func() { <-gate; _, err := w.UpdateClaim(root, previous, previous.Base, "new"); errs <- err }()
			go func() { <-gate; errs <- w.ReleaseClaim(root, "race", previous.Object) }()
			close(gate)
			a, b := <-errs, <-errs
			if (a == nil) == (b == nil) {
				t.Fatalf("update/release must have one winner: %v / %v", a, b)
			}
			loser := a
			if loser == nil {
				loser = b
			}
			claimReleaseError(t, loser)
		})
	}
	previous := claimSeed(t, root, "updates", "old")
	gate := make(chan struct{})
	type result struct {
		c   w.Claim
		err error
	}
	results := make(chan result, 2)
	for _, state := range []string{"new-one", "new-two"} {
		go func(state string) {
			<-gate
			c, err := w.UpdateClaim(root, previous, previous.Base, state)
			results <- result{c, err}
		}(state)
	}
	close(gate)
	a, b := <-results, <-results
	if (a.err == nil) == (b.err == nil) {
		t.Fatalf("two updates must have one winner: %+v %+v", a, b)
	}
	winner, loser := a, b
	if winner.err != nil {
		winner, loser = loser, winner
	}
	claimError(t, loser.c, loser.err)
	if claimGit(t, root, "rev-parse", "refs/fireweed/claims/updates") != winner.c.Object {
		t.Fatal("winner token not stored")
	}
}
func Test_FLOW010ReadClaimValidation(t *testing.T) {
	root := claimRepo(t, "sha1")
	valid := w.Claim{Worktree: root, Change: "read", Run: "run", Base: claimBase(t, root), State: ""}
	good := claimPayload(t, valid)
	t.Run("missing", func(t *testing.T) { got, err := w.LoadClaim(root, "missing"); claimError(t, got, err) })
	tests := map[string]string{"broken": "{", "trailing": good + " {}"}
	var fields map[string]any
	if err := json.Unmarshal([]byte(good), &fields); err != nil {
		t.Fatal(err)
	}
	for _, key := range []string{"worktree", "change", "run", "base", "state"} {
		for _, mode := range []string{"absent", "null", "number", "bool", "list", "map"} {
			copy := map[string]any{}
			for k, v := range fields {
				copy[k] = v
			}
			switch mode {
			case "absent":
				delete(copy, key)
			case "null":
				copy[key] = nil
			case "number":
				copy[key] = 3
			case "bool":
				copy[key] = true
			case "list":
				copy[key] = []string{}
			case "map":
				copy[key] = map[string]string{}
			}
			raw, _ := json.Marshal(copy)
			tests[key+"_"+mode] = string(raw)
		}
	}
	for _, row := range []struct{ name, key, value string }{{"target", "change", "different"}, {"slug", "change", "../read"}, {"run", "run", " \t"}, {"relative_owner", "worktree", "relative"}, {"unclean_owner", "worktree", root + "/nested/.."}, {"empty_owner", "worktree", ""}, {"missing_base", "base", strings.Repeat("f", len(valid.Base))}, {"named_base", "base", "HEAD"}, {"blob_base", "base", claimRaw(t, root, "blob-seed", "not a commit")}} {
		copy := map[string]any{}
		for k, v := range fields {
			copy[k] = v
		}
		copy[row.key] = row.value
		raw, _ := json.Marshal(copy)
		tests[row.name] = string(raw)
	}
	for name, raw := range tests {
		t.Run(name, func(t *testing.T) {
			claimRaw(t, root, "read", raw)
			got, err := w.LoadClaim(root, "read")
			claimError(t, got, err)
		})
	}
	claimGit(t, root, "update-ref", "refs/fireweed/claims/read", valid.Base)
	t.Run("nonblob", func(t *testing.T) { got, err := w.LoadClaim(root, "read"); claimError(t, got, err) })
	for _, slug := range []string{"../read", "read/child", "Upper", ""} {
		t.Run("invalid_slug_"+slug, func(t *testing.T) { got, err := w.LoadClaim(root, slug); claimError(t, got, err) })
	}
	extra := map[string]any{}
	for k, v := range fields {
		extra[k] = v
	}
	extra["future"] = map[string]any{"nested": true}
	extra["worktree"] = filepath.Join(root, "owner-no-longer-exists")
	raw, _ := json.Marshal(extra)
	object := claimRaw(t, root, "read", string(raw))
	got, err := w.LoadClaim(root, "read")
	if err != nil || got.Worktree != extra["worktree"] || got.State != "" || got.Run != valid.Run || got.Base != valid.Base || got.Change != "read" || got.Object != object {
		t.Fatalf("stale owner/unknown fields: %+v %v", got, err)
	}
}
func Test_FLOW010ClaimsInventoryAndOwnership(t *testing.T) {
	root := claimRepo(t, "sha1")
	t.Run("empty_inventory", func(t *testing.T) {
		emptyA, err := w.Claims(root)
		if err != nil || emptyA == nil || len(emptyA) != 0 {
			t.Fatalf("empty inventory %+v %v", emptyA, err)
		}
		emptyB, err := w.Claims(root)
		if err != nil || emptyB == nil {
			t.Fatal(err)
		}
		emptyA = append(emptyA, w.Claim{Run: "a"})
		emptyB = append(emptyB, w.Claim{Run: "b"})
		if emptyA[0].Run != "a" || emptyB[0].Run != "b" {
			t.Fatal("empty slices alias")
		}
	})
	expected := []w.Claim{claimSeed(t, root, "zeta-0", "z"), claimSeed(t, root, "alpha", "a"), claimSeed(t, root, "middle", "m")}
	sort.Slice(expected, func(i, j int) bool { return expected[i].Change < expected[j].Change })
	claimGit(t, root, "update-ref", "refs/fireweed/claims-unrelated/ignored", expected[0].Object)
	t.Run("sorted_and_owned", func(t *testing.T) {
		a, err := w.Claims(root)
		if err != nil || !reflect.DeepEqual(a, expected) {
			t.Fatalf("inventory %+v %v; want %+v", a, err, expected)
		}
		b, err := w.Claims(root)
		if err != nil {
			t.Fatal(err)
		}
		a[0].State = "mutate result"
		a = append(a, w.Claim{Change: "new"})
		if !reflect.DeepEqual(b, expected) {
			t.Fatal("inventories share backing storage")
		}
		one, err := w.LoadClaim(root, expected[0].Change)
		if err != nil || one != expected[0] {
			t.Fatalf("mutated persisted payload: %+v %v", one, err)
		}
	})
	for _, slot := range []string{"bad", "Upper", "nested/child"} {
		t.Run(slot, func(t *testing.T) {
			payload := claimPayload(t, expected[0])
			if slot == "bad" {
				payload = "{malformed"
			}
			claimRaw(t, root, slot, payload)
			defer claimGit(t, root, "update-ref", "-d", "refs/fireweed/claims/"+slot)
			got, err := w.Claims(root)
			if err == nil || !strings.Contains(err.Error(), "P02") || !policyCorrective(err.Error()) || len(got) != 0 {
				t.Fatalf("partial or silent malformed inventory %+v %v", got, err)
			}
		})
	}
}

func claimTree(t *testing.T, root string, skip func(string) bool) map[string]string {
	t.Helper()
	result := map[string]string{}
	err := filepath.WalkDir(root, func(path string, e os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		rel = filepath.ToSlash(rel)
		if skip(rel) {
			if e.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}
		info, err := e.Info()
		if err != nil {
			return err
		}
		text := info.Mode().String()
		if !e.IsDir() {
			raw, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			text += "\n" + string(raw)
		}
		result[rel] = text
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	return result
}
func claimCommon(t *testing.T, root string) string {
	t.Helper()
	dir := claimGit(t, root, "rev-parse", "--git-common-dir")
	if !filepath.IsAbs(dir) {
		dir = filepath.Join(root, dir)
	}
	return stateCanonicalDirectory(t, dir)
}
func claimMetadata(t *testing.T, common, slug string) map[string]string {
	t.Helper()
	return claimTree(t, common, func(path string) bool {
		if path == "objects" || strings.HasPrefix(path, "objects/") {
			return true
		}
		for _, prefix := range []string{"refs/fireweed/claims", "logs/refs/fireweed/claims"} {
			if path == prefix+"/"+slug {
				return true
			}
			if path == prefix || strings.HasPrefix(prefix, path+"/") {
				return true
			}
		}
		return false
	})
}
func claimQuote(s string) string { return "'" + strings.ReplaceAll(s, "'", `'"'"'`) + "'" }
func Test_FLOW011ClaimSafetyAndNativeHookControl(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("POSIX executable-trap/native-hook fixture; other claim fixtures remain portable")
	}
	root := claimRepo(t, "sha1")
	base := claimBase(t, root)
	common := claimCommon(t, root)
	marker := filepath.Join(t.TempDir(), "unexpected command")
	nativeMarker := filepath.Join(t.TempDir(), "native hook")
	if err := os.WriteFile(filepath.Join(root, ".gitignore"), []byte("ignored\n"), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, "ignored"), []byte("ignored bytes"), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, "untracked"), []byte("user bytes"), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, "retained"), []byte("staged bytes"), 0644); err != nil {
		t.Fatal(err)
	}
	claimGit(t, root, "add", "retained")
	if err := os.WriteFile(filepath.Join(root, "retained"), []byte("unstaged bytes"), 0644); err != nil {
		t.Fatal(err)
	}
	// Independent fixture control proves native ref hooks really execute and can reject writes.
	hook := filepath.Join(common, "hooks", "reference-transaction")
	script := "#!/bin/sh\nprintf called >> " + claimQuote(nativeMarker) + "\nif [ \"$1\" = prepared ]; then exit 1; fi\nexit 0\n"
	if err := os.WriteFile(hook, []byte(script), 0755); err != nil {
		t.Fatal(err)
	}
	if out, err := claimCommand(t, root, "", "update-ref", "refs/fireweed/control", base); err == nil {
		t.Fatalf("native hook control must reject: %s", out)
	}
	if _, err := os.Stat(nativeMarker); err != nil {
		t.Fatal("native hook control did not execute", err)
	}
	if err := os.Remove(nativeMarker); err != nil {
		t.Fatal(err)
	}
	traps := t.TempDir()
	realGit, err := exec.LookPath("git")
	if err != nil {
		t.Fatal(err)
	}
	wrapper := "#!/bin/sh\nfor a in \"$@\"; do case \"$a\" in fetch|push|clone|ls-remote|commit|branch|checkout|switch|reset|restore|add|worktree) printf forbidden >> " + claimQuote(marker) + "; exit 97;; esac; done\nexec " + claimQuote(realGit) + " \"$@\"\n"
	if err := os.WriteFile(filepath.Join(traps, "git"), []byte(wrapper), 0755); err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"hook-trap", "agent-trap", "codex", "claude", "fireweed", "fireweed-bootstrap", "curl", "wget", "ssh", "nc"} {
		if err := os.WriteFile(filepath.Join(traps, name), []byte("#!/bin/sh\nprintf forbidden >> "+claimQuote(marker)+"\nexit 97\n"), 0755); err != nil {
			t.Fatal(err)
		}
	}
	t.Setenv("PATH", traps+string(os.PathListSeparator)+os.Getenv("PATH"))
	t.Setenv("FIREWEED_BIN", filepath.Join(traps, "fireweed-bootstrap"))
	t.Setenv("GIT_CONFIG_NOSYSTEM", "1")
	t.Setenv("GIT_CONFIG_GLOBAL", os.DevNull)
	for _, name := range []string{"GIT_AUTHOR_NAME", "GIT_AUTHOR_EMAIL", "GIT_COMMITTER_NAME", "GIT_COMMITTER_EMAIL"} {
		t.Setenv(name, "")
	}
	if err := exec.Command("git", "fetch", "forbidden.invalid").Run(); err == nil {
		t.Fatal("network command control unexpectedly succeeded")
	}
	if _, err := os.Stat(marker); err != nil {
		t.Fatal("external-command control did not execute")
	}
	if err := os.Remove(marker); err != nil {
		t.Fatal(err)
	}
	isolatedTemp := stateCanonicalDirectory(t, t.TempDir())
	t.Setenv("TMPDIR", isolatedTemp)
	t.Setenv("TMP", isolatedTemp)
	t.Setenv("TEMP", isolatedTemp)
	claimGit(t, root, "--version") // Establish native Git's own cache before measuring owned temporary cleanup.
	beforeTemp := claimTree(t, isolatedTemp, func(string) bool { return false })
	beforeObjects := claimTree(t, filepath.Join(common, "objects"), func(string) bool { return false })
	beforeFiles := claimTree(t, root, func(path string) bool { return path == ".git" })
	beforeMeta := claimMetadata(t, common, "safe")
	env := os.Environ()
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	assertSafe := func() {
		t.Helper()
		after, err := os.Getwd()
		if err != nil || after != cwd || !reflect.DeepEqual(env, os.Environ()) || !reflect.DeepEqual(beforeFiles, claimTree(t, root, func(path string) bool { return path == ".git" })) || !reflect.DeepEqual(beforeMeta, claimMetadata(t, common, "safe")) {
			t.Fatal("mutated user files, index/HEAD/other refs, leaked temporary storage, or parent state")
		}
		if !reflect.DeepEqual(beforeTemp, claimTree(t, isolatedTemp, func(string) bool { return false })) {
			t.Fatal("leaked or changed dedicated temporary files")
		}
		afterObjects := claimTree(t, filepath.Join(common, "objects"), func(string) bool { return false })
		for path, value := range beforeObjects {
			if afterObjects[path] != value {
				t.Fatalf("removed or changed existing object storage %s", path)
			}
		}
		for path, value := range afterObjects {
			if old, ok := beforeObjects[path]; ok {
				if old != value {
					t.Fatalf("changed existing Git object storage %s", path)
				}
				continue
			}
			parts := strings.Split(path, "/")
			if len(parts) == 1 && len(parts[0]) == 2 && strings.Trim(parts[0], "0123456789abcdef") == "" && strings.HasPrefix(value, "d") {
				populated := false
				for child := range afterObjects {
					if strings.HasPrefix(child, path+"/") {
						populated = true
						break
					}
				}
				if !populated {
					t.Fatalf("leaked empty object-directory temporary: %s", path)
				}
				continue
			}
			if len(parts) != 2 || len(parts[0]) != 2 || len(parts[0]+parts[1]) != len(base) || strings.Trim(parts[0]+parts[1], "0123456789abcdef") != "" {
				t.Fatalf("leaked unexpected temporary storage under objects: %s", path)
			}
			if claimGit(t, root, "cat-file", "-t", parts[0]+parts[1]) != "blob" {
				t.Fatalf("unexpected new non-claim object %s", path)
			}
		}
		for _, path := range []string{marker, nativeMarker} {
			if _, err := os.Stat(path); !os.IsNotExist(err) {
				t.Fatalf("ran forbidden hook/agent/Git/network command: %s", path)
			}
		}
	}
	created, err := w.CreateClaim(root, "safe", "new-run", base, "approved")
	if err != nil {
		t.Fatal(err)
	}
	claimAssert(t, root, created, "safe", "new-run", base, "approved")
	assertSafe()
	loaded, err := w.LoadClaim(root, "safe")
	if err != nil || loaded != created {
		t.Fatalf("load %+v %v", loaded, err)
	}
	inventory, err := w.Claims(root)
	if err != nil || !reflect.DeepEqual(inventory, []w.Claim{created}) {
		t.Fatalf("inventory %+v %v", inventory, err)
	}
	assertSafe()
	got, err := w.CreateClaim(root, "safe", "duplicate", base, "not approved")
	claimError(t, got, err)
	assertSafe()
	got, err = w.UpdateClaim(root, created, base, "next phase")
	if err != nil {
		t.Fatal(err)
	}
	claimAssert(t, root, got, "safe", created.Run, base, "next phase")
	assertSafe()
	bad, err := w.UpdateClaim(root, created, base, "stale write")
	claimError(t, bad, err)
	claimReleaseError(t, w.ReleaseClaim(root, "safe", created.Object))
	assertSafe()
	if err = w.ReleaseClaim(root, "safe", got.Object); err != nil {
		t.Fatal(err)
	}
	assertSafe()
}
func Test_FLOW011ReadAndRejectedCallsPreserveRefs(t *testing.T) {
	root := claimRepo(t, "sha1")
	previous := claimSeed(t, root, "owned", "state")
	common := claimCommon(t, root)
	before := claimTree(t, common, func(path string) bool { return false })
	files := claimTree(t, root, func(path string) bool { return path == ".git" })
	input := previous
	for _, call := range []struct {
		name string
		run  func(*testing.T)
	}{
		{"read", func(t *testing.T) {
			got, err := w.LoadClaim(root, "owned")
			if err != nil || got != previous {
				t.Fatalf("load %+v %v", got, err)
			}
		}},
		{"list", func(t *testing.T) {
			got, err := w.Claims(root)
			if err != nil || !reflect.DeepEqual(got, []w.Claim{previous}) {
				t.Fatalf("list %+v %v", got, err)
			}
		}},
		{"bad_create", func(t *testing.T) {
			got, err := w.CreateClaim(root, "../invalid", "run", previous.Base, "state")
			claimError(t, got, err)
		}},
		{"bad_base_create", func(t *testing.T) {
			got, err := w.CreateClaim(root, "valid", "run", "missing", "state")
			claimError(t, got, err)
		}},
		{"literal_base", func(t *testing.T) {
			got, err := w.CreateClaim(root, "valid", "run", "HEAD; printf injected > "+claimQuote(filepath.Join(root, "injected")), "state")
			claimError(t, got, err)
		}},
		{"missing_read", func(t *testing.T) { got, err := w.LoadClaim(root, "missing"); claimError(t, got, err) }},
		{"bad_token_update", func(t *testing.T) {
			changed := previous
			changed.Object = "bad-token"
			got, err := w.UpdateClaim(root, changed, previous.Base, "new")
			claimError(t, got, err)
		}},
		{"bad_base_update", func(t *testing.T) {
			got, err := w.UpdateClaim(root, previous, "missing", "new")
			claimError(t, got, err)
		}},
		{"bad_token_release", func(t *testing.T) { claimReleaseError(t, w.ReleaseClaim(root, "owned", "bad-token")) }},
	} {
		t.Run(call.name, func(t *testing.T) {
			call.run(t)
			if previous != input || !reflect.DeepEqual(before, claimTree(t, common, func(path string) bool { return false })) || !reflect.DeepEqual(files, claimTree(t, root, func(path string) bool { return path == ".git" })) {
				t.Fatal("read or failed precondition mutated inputs, files, refs or Git storage")
			}
		})
	}
}
func Test_FLOW008_FLOW009_FLOW010_FLOW011ClaimFixtureControls(t *testing.T) {
	for _, format := range []string{"sha1", "sha256"} {
		t.Run(format, func(t *testing.T) {
			root := claimRepo(t, format)
			peer := claimPeer(t, root)
			c := claimSeed(t, root, "fixture-control", "\nopaque\x00state")
			claimAssert(t, root, c, c.Change, c.Run, c.Base, c.State)
			var parsed w.Claim
			if err := json.Unmarshal([]byte(claimGit(t, peer, "cat-file", "blob", c.Object)), &parsed); err != nil {
				t.Fatal(err)
			}
			parsed.Object = c.Object
			if parsed != c {
				t.Fatalf("independent wire/linked Git fixture: %+v != %+v", parsed, c)
			}
			before := claimGit(t, root, "rev-parse", "refs/fireweed/claims/fixture-control")
			if out, err := claimCommand(t, peer, "", "update-ref", "refs/fireweed/claims/fixture-control", c.Base, ""); err == nil {
				t.Fatalf("Git create-only premise failed: %s", out)
			}
			if claimGit(t, root, "rev-parse", "refs/fireweed/claims/fixture-control") != before {
				t.Fatal("failed create modified ref")
			}
			if out, err := claimCommand(t, peer, "", "update-ref", "-d", "refs/fireweed/claims/fixture-control", c.Base); err == nil {
				t.Fatalf("Git stale-delete premise failed: %s", out)
			}
			if len(c.Object) != len(c.Base) || (format == "sha256" && len(c.Object) != 64) || (format == "sha1" && len(c.Object) != 40) {
				t.Fatal("fixture hash format mismatch")
			}
		})
	}
}

func Test_FLOW008_FLOW009NonCommitBases(t *testing.T) {
	root := claimRepo(t, "sha1")
	original := claimSeed(t, root, "owned", "old")
	blob := claimRaw(t, root, "blob-source", "opaque blob")
	refs := claimGit(t, root, "show-ref")
	t.Run("create", func(t *testing.T) {
		got, err := w.CreateClaim(root, "new", "run", blob, "state")
		claimError(t, got, err)
		if claimGit(t, root, "show-ref") != refs {
			t.Fatal("non-commit create changed refs")
		}
	})
	t.Run("update", func(t *testing.T) {
		got, err := w.UpdateClaim(root, original, blob, "state")
		claimError(t, got, err)
		if claimGit(t, root, "show-ref") != refs {
			t.Fatal("non-commit update changed refs")
		}
	})
}
func Test_FLOW009FullReleaseTokenRequired(t *testing.T) {
	root := claimRepo(t, "sha1")
	original := claimSeed(t, root, "owned", "state")
	short := claimGit(t, root, "rev-parse", "--short=12", original.Object)
	for _, token := range []string{short, "refs/fireweed/claims/owned"} {
		t.Run(token, func(t *testing.T) {
			claimReleaseError(t, w.ReleaseClaim(root, "owned", token))
			if claimGit(t, root, "rev-parse", "refs/fireweed/claims/owned") != original.Object {
				t.Fatal("non-full token released claim")
			}
		})
	}
}
func Test_FLOW011LinkedUpdateAndInventory(t *testing.T) {
	for _, format := range []string{"sha1", "sha256"} {
		t.Run(format, func(t *testing.T) {
			root := claimRepo(t, format)
			peer := claimPeer(t, root)
			previous := claimSeed(t, peer, "peer-owned", "old")
			updated, err := w.UpdateClaim(peer, previous, previous.Base, "new")
			if err != nil {
				t.Fatal(err)
			}
			claimAssert(t, peer, updated, previous.Change, previous.Run, previous.Base, "new")
			for _, dir := range []string{root, peer} {
				inventory, err := w.Claims(dir)
				if err != nil || !reflect.DeepEqual(inventory, []w.Claim{updated}) {
					t.Fatalf("shared %s inventory %+v %v", format, inventory, err)
				}
			}
		})
	}
}

// Observe Git's public compare-and-swap protocol. This accepts either argv or
// documented text/NUL stdin transactions, without reading implementation code.
func claimCASProtocol(raw, input []byte, kind, ref, newObject, oldObject string) bool {
	args := strings.Split(string(raw), "\x00")
	if len(args) > 0 && args[len(args)-1] == "" {
		args = args[:len(args)-1]
	}
	index := -1
	for i, arg := range args {
		if arg == "update-ref" {
			index = i
			break
		}
	}
	if index < 0 {
		return false
	}
	args = args[index+1:]
	stdin := false
	for _, arg := range args {
		if arg == "--stdin" {
			stdin = true
		}
	}
	validNew := func(value string) bool { return value == newObject && value != "" }
	if stdin {
		// A transaction may contain start/prepare/commit lines around its single ref operation.
		chunks := strings.Split(strings.ReplaceAll(string(input), "\x00", "\n"), "\n")
		for i, line := range chunks {
			fields := strings.Fields(line)
			if len(fields) < 2 || fields[1] != ref {
				continue
			}
			values := append([]string{}, fields[2:]...)
			for j := i + 1; len(values) < 2 && j < len(chunks); j++ {
				if strings.Contains(chunks[j], " ") {
					break
				}
				values = append(values, chunks[j])
			}
			switch kind {
			case "create":
				if fields[0] == "create" && len(values) >= 1 && validNew(values[0]) {
					return true
				}
				if fields[0] == "update" && len(values) >= 2 && validNew(values[0]) && len(values[1]) == len(newObject) && strings.Trim(values[1], "0") == "" {
					return true
				}
			case "update":
				if fields[0] == "update" && len(values) >= 2 && validNew(values[0]) && values[1] == oldObject {
					return true
				}
			case "delete":
				if fields[0] == "delete" && len(values) >= 1 && values[0] == oldObject {
					return true
				}
			}
		}
		return false
	}
	for i, arg := range args {
		if arg != ref {
			continue
		}
		switch kind {
		case "create":
			return i+2 < len(args) && validNew(args[i+1]) && (args[i+2] == "" || strings.Trim(args[i+2], "0") == "")
		case "update":
			return i+2 < len(args) && validNew(args[i+1]) && args[i+2] == oldObject
		case "delete":
			deleted := false
			for _, flag := range args[:i] {
				deleted = deleted || flag == "-d"
			}
			return deleted && i+1 < len(args) && args[i+1] == oldObject
		}
	}
	return false
}
func Test_FLOW008_FLOW009CASProtocolFixtureControl(t *testing.T) {
	ref, newOID, oldOID := "refs/fireweed/claims/control", strings.Repeat("a", 40), strings.Repeat("b", 40)
	encode := func(args ...string) []byte { return []byte(strings.Join(args, "\x00") + "\x00") }
	for _, row := range []struct {
		kind        string
		args, input []byte
		valid       bool
	}{
		{"create", encode("update-ref", ref, newOID, ""), nil, true}, {"create", encode("update-ref", ref, newOID), nil, false},
		{"update", encode("update-ref", ref, newOID, oldOID), nil, true}, {"update", encode("update-ref", ref, newOID), nil, false}, {"update", encode("update-ref", ref, newOID, strings.Repeat("0", 40)), nil, false},
		{"delete", encode("update-ref", "-d", ref, oldOID), nil, true}, {"delete", encode("update-ref", "-d", ref), nil, false},
		{"create", encode("update-ref", "--stdin"), []byte("create " + ref + " " + newOID + "\n"), true},
		{"create", encode("update-ref", "--stdin"), []byte("update " + ref + " " + newOID + "\n"), false},
		{"create", encode("update-ref", "--stdin"), []byte("update " + ref + " " + newOID + " " + strings.Repeat("0", 40) + "\n"), true},
		{"create", encode("update-ref", "-z", "--stdin"), []byte("update " + ref + "\x00" + newOID + "\x00\x00"), false},
		{"create", encode("update-ref", "-z", "--stdin"), []byte("update " + ref + "\x00" + newOID + "\x00" + strings.Repeat("0", 40) + "\x00"), true},
		{"update", encode("update-ref", "--stdin"), []byte("start\nupdate " + ref + " " + newOID + " " + oldOID + "\nprepare\ncommit\n"), true},
		{"delete", encode("update-ref", "-z", "--stdin"), []byte("delete " + ref + "\x00" + oldOID + "\x00"), true},
	} {
		if got := claimCASProtocol(row.args, row.input, row.kind, ref, newOID, oldOID); got != row.valid {
			t.Fatalf("CAS observer premise %s %q %q: %v", row.kind, row.args, row.input, got)
		}
	}
}
func Test_FLOW008_FLOW009NativeCASUpdates(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("POSIX Git argument observer; portable race/control fixtures still execute")
	}
	for _, format := range []string{"sha1", "sha256"} {
		t.Run(format, func(t *testing.T) { claimNativeCASUpdates(t, format) })
	}
}
func claimNativeCASUpdates(t *testing.T, format string) {
	root := claimRepo(t, format)
	base := claimBase(t, root)
	realGit, err := exec.LookPath("git")
	if err != nil {
		t.Fatal(err)
	}
	traps := t.TempDir()
	trace := filepath.Join(t.TempDir(), "git-cas-arguments")
	script := "#!/bin/sh\nfound=\nstdin=\nfor a in \"$@\"; do [ \"$a\" = update-ref ] && found=1; [ \"$a\" = --stdin ] && stdin=1; done\nif [ -n \"$found\" ]; then\n printf '%s\\000' \"$@\" > " + claimQuote(trace) + "\n if [ -n \"$stdin\" ]; then cat > " + claimQuote(trace+".stdin") + "; exec " + claimQuote(realGit) + " \"$@\" < " + claimQuote(trace+".stdin") + "; fi\nfi\nexec " + claimQuote(realGit) + " \"$@\"\n"
	if err := os.WriteFile(filepath.Join(traps, "git"), []byte(script), 0755); err != nil {
		t.Fatal(err)
	}
	t.Setenv("PATH", traps+string(os.PathListSeparator)+os.Getenv("PATH"))
	observe := func(t *testing.T, kind, slug, newOID, oldOID string) {
		t.Helper()
		raw, err := os.ReadFile(trace)
		if err != nil {
			t.Fatal("operation did not reach an atomic Git ref update", err)
		}
		input, err := os.ReadFile(trace + ".stdin")
		if err != nil && !os.IsNotExist(err) {
			t.Fatal(err)
		}
		if !claimCASProtocol(raw, input, kind, "refs/fireweed/claims/"+slug, newOID, oldOID) {
			t.Fatalf("missing native Git %s compare-and-swap: %q stdin %q", kind, raw, input)
		}
	}
	clear := func() {
		t.Helper()
		for _, path := range []string{trace, trace + ".stdin"} {
			if err := os.Remove(path); err != nil && !os.IsNotExist(err) {
				t.Fatal(err)
			}
		}
	}
	// Positive control verifies the observer against an actual Git create-only write.
	claimGit(t, root, "update-ref", "refs/fireweed/claims/observer", base, "")
	observe(t, "create", "observer", base, "")
	claimGit(t, root, "update-ref", "-d", "refs/fireweed/claims/observer", base)
	clear()
	t.Run("create", func(t *testing.T) {
		got, err := w.CreateClaim(root, "created", "run", base, "")
		if err != nil {
			t.Fatal(err)
		}
		observe(t, "create", "created", got.Object, "")
	})
	clear()
	previous := claimSeed(t, root, "updated", "old")
	clear()
	t.Run("update", func(t *testing.T) {
		got, err := w.UpdateClaim(root, previous, base, "changed")
		if err != nil {
			t.Fatal(err)
		}
		observe(t, "update", "updated", got.Object, previous.Object)
	})
	clear()
	released := claimSeed(t, root, "released", "old")
	clear()
	t.Run("release", func(t *testing.T) {
		if err := w.ReleaseClaim(root, "released", released.Object); err != nil {
			t.Fatal(err)
		}
		observe(t, "delete", "released", "", released.Object)
	})
}

func Test_FLOW008GitStdinCASFixtureControl(t *testing.T) {
	for _, format := range []string{"sha1", "sha256"} {
		t.Run(format, func(t *testing.T) {
			root := claimRepo(t, format)
			old := claimRaw(t, root, "old-source", "old")
			newOID := claimRaw(t, root, "new-source", "new")
			ref := "refs/fireweed/claims/stdin-control"
			zeros := strings.Repeat("0", len(old))
			for _, row := range []struct {
				name        string
				args        []string
				input       string
				conditional bool
			}{
				{"text_missing_old", []string{"update-ref", "--stdin"}, "update " + ref + " " + newOID + "\n", false},
				{"text_full_zero", []string{"update-ref", "--stdin"}, "update " + ref + " " + newOID + " " + zeros + "\n", true},
				{"nul_empty_old", []string{"update-ref", "--stdin", "-z"}, "update " + ref + "\x00" + newOID + "\x00\x00", false},
				{"nul_full_zero", []string{"update-ref", "--stdin", "-z"}, "update " + ref + "\x00" + newOID + "\x00" + zeros + "\x00", true},
				{"text_create", []string{"update-ref", "--stdin"}, "create " + ref + " " + newOID + "\n", true},
				{"nul_create", []string{"update-ref", "--stdin", "-z"}, "create " + ref + "\x00" + newOID + "\x00", true},
			} {
				t.Run(row.name, func(t *testing.T) {
					claimGit(t, root, "update-ref", ref, old)
					out, err := claimCommand(t, root, row.input, row.args...)
					if (err != nil) != row.conditional {
						t.Fatalf("native Git CAS premise: conditional %v, err %v, output %s", row.conditional, err, out)
					}
					expected := newOID
					if row.conditional {
						expected = old
					}
					if claimGit(t, root, "rev-parse", ref) != expected {
						t.Fatal("unexpected native Git ref result")
					}
					raw := []byte(strings.Join(row.args, "\x00") + "\x00")
					if claimCASProtocol(raw, []byte(row.input), "create", ref, newOID, "") != row.conditional {
						t.Fatalf("observer disagrees with actual Git create-only semantics: %s", row.name)
					}
				})
			}
		})
	}
}
func Test_FLOW008_FLOW009_FLOW011SymbolicClaimRefSafety(t *testing.T) {
	for _, operation := range []string{"create", "update", "release"} {
		t.Run(operation, func(t *testing.T) {
			root := claimRepo(t, "sha1")
			base := claimBase(t, root)
			previous := w.Claim{Worktree: root, Change: "alias", Run: "run", Base: base, State: "old"}
			previous.Object = claimRaw(t, root, "source", claimPayload(t, previous))
			claimGit(t, root, "update-ref", "-d", "refs/fireweed/claims/source", previous.Object)
			target := "refs/fireweed/unrelated-target"
			ref := "refs/fireweed/claims/alias"
			claimGit(t, root, "update-ref", target, previous.Object)
			claimGit(t, root, "symbolic-ref", ref, target)
			common := claimCommon(t, root)
			before := claimMetadata(t, common, "alias")
			files := claimTree(t, root, func(path string) bool { return path == ".git" })
			rawPath := filepath.Join(common, "refs/fireweed/claims/alias")
			original, err := os.ReadFile(rawPath)
			if err != nil {
				t.Fatal(err)
			}
			rejected := false
			switch operation {
			case "create":
				got, err := w.CreateClaim(root, "alias", "new-run", base, "new")
				claimError(t, got, err)
				rejected = true
			case "update":
				got, err := w.UpdateClaim(root, previous, base, "new")
				if err != nil {
					claimError(t, got, err)
					rejected = true
				} else {
					claimAssert(t, root, got, "alias", "run", base, "new")
				}
			case "release":
				err := w.ReleaseClaim(root, "alias", previous.Object)
				if err != nil {
					claimReleaseError(t, err)
					rejected = true
				} else if _, err := os.Lstat(rawPath); !os.IsNotExist(err) {
					t.Fatalf("release did not delete selected alias: %v", err)
				}
			}
			if rejected {
				after, err := os.ReadFile(rawPath)
				if err != nil || string(after) != string(original) {
					t.Fatal("rejection changed selected alias")
				}
			}
			if claimGit(t, root, "rev-parse", target) != previous.Object || !reflect.DeepEqual(before, claimMetadata(t, common, "alias")) || !reflect.DeepEqual(files, claimTree(t, root, func(path string) bool { return path == ".git" })) {
				t.Fatal("claim operation dereferenced an alias and changed unrelated refs or files")
			}
		})
	}
}
