package evals_test

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	a "github.com/lurkshark/fireweed/packages/assets"
	c "github.com/lurkshark/fireweed/packages/contracts"
	g "github.com/lurkshark/fireweed/packages/gitstore"
	p "github.com/lurkshark/fireweed/packages/project"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"testing"
)

func playbookRoot(t *testing.T) string {
	t.Helper()
	root := fixture(t)
	root, err := filepath.EvalSymlinks(root)
	if err != nil {
		t.Fatal(err)
	}
	put(t, root, ".gitignore", "ignored\n")
	put(t, root, "ignored", "untouched ignored bytes")
	put(t, root, "fireweed.yaml", "deliberately invalid configuration: [\n")
	git(t, root, "add", ".")
	git(t, root, "commit", "-m", "playbook fixture")
	return root
}
func playbookHash(text string) string { return fmt.Sprintf("%x", sha256.Sum256([]byte(text))) }
func playbookWant(t *testing.T, name, text, based string, overridden, stale bool) p.Playbook {
	t.Helper()
	embedded, err := a.Playbook(name)
	if err != nil {
		t.Fatal(err)
	}
	return p.Playbook{Name: name, Text: text, Hash: playbookHash(text), EmbeddedHash: embedded.Hash, BasedOn: based, Overridden: overridden, Stale: stale}
}
func playbookAssert(t *testing.T, start, name string, want p.Playbook) {
	t.Helper()
	got, err := p.ResolvePlaybook(start, name)
	if err != nil || got != want {
		t.Fatalf("resolve(%q,%q) %+v %v; want %+v", start, name, got, err, want)
	}
	raw, err := json.Marshal(got)
	if err != nil {
		t.Fatal(err)
	}
	var fields map[string]any
	if err := json.Unmarshal(raw, &fields); err != nil {
		t.Fatal(err)
	}
	expected := map[string]any{"name": want.Name, "text": want.Text, "hash": want.Hash, "embedded_hash": want.EmbeddedHash, "based_on": want.BasedOn, "overridden": want.Overridden, "stale": want.Stale}
	if !reflect.DeepEqual(fields, expected) {
		t.Fatalf("public JSON fields %+v; want %+v", fields, expected)
	}
}
func playbookCorrective(text string) bool {
	text = strings.ToLower(text)
	for _, word := range []string{"restore", "open", "use ", "choose", "repair", "start", "select", "check", "create"} {
		if strings.Contains(text, word) {
			return true
		}
	}
	return false
}
func playbookError(t *testing.T, start, name, code, path string) {
	t.Helper()
	got, err := p.ResolvePlaybook(start, name)
	if err == nil || got != (p.Playbook{}) || !playbookCorrective(err.Error()) || (code != "" && !strings.Contains(err.Error(), code)) || (path != "" && !strings.Contains(err.Error(), path)) {
		t.Fatalf("expected zero result and corrective %s at %s: %+v %v", code, path, got, err)
	}
}
func Test_PRJ021EmbeddedResolution(t *testing.T) {
	root := playbookRoot(t)
	nested := filepath.Join(root, "nested 'quoted space'")
	if err := os.Mkdir(nested, 0755); err != nil {
		t.Fatal(err)
	}
	for _, name := range a.Playbooks() {
		for _, start := range []string{root, nested} {
			t.Run(name+"_"+filepath.Base(start), func(t *testing.T) {
				asset, err := a.Playbook(name)
				if err != nil {
					t.Fatal(err)
				}
				playbookAssert(t, start, name, playbookWant(t, name, asset.Text, "", false, false))
			})
		}
	}
}
func Test_PRJ021LinkedCheckoutOverrides(t *testing.T) {
	root := playbookRoot(t)
	peer := filepath.Join(t.TempDir(), "peer")
	git(t, root, "worktree", "add", "--detach", peer, "HEAD")
	name := "impl"
	asset, _ := a.Playbook(name)
	text := "---\nbased_on: '" + asset.Hash + "'\n---\n\nLinked checkout instructions.\n"
	put(t, root, ".fireweed/playbooks/impl.md", "different main checkout instructions")
	put(t, peer, ".fireweed/playbooks/impl.md", text)
	playbookAssert(t, peer, name, playbookWant(t, name, text, asset.Hash, true, false))
}
func Test_PRJ021_PRJ022ExactOverrideBytesAndStaleness(t *testing.T) {
	root := playbookRoot(t)
	for _, name := range a.Playbooks() {
		asset, _ := a.Playbook(name)
		otherName := "evals"
		if name == otherName {
			otherName = "impl"
		}
		other, _ := a.Playbook(otherName)
		if other.Hash == asset.Hash {
			t.Fatal("distinct hash premise")
		}
		rows := []struct {
			name, text, based string
			stale             bool
		}{
			{"current", "---\nbased_on: '" + asset.Hash + "'\n---\n\nExact complete instructions.\n", asset.Hash, false},
			{"crlf", "---\r\nbased_on: \"" + asset.Hash + "\"\r\n---\r\n\r\nCRLF retained.\r\n", asset.Hash, false},
			{"missing", "plain override instructions\n", "", true},
			{"empty", "", "", true},
			{"wrong", "---\nbased_on: wrong\n---\nInstructions.\n", "wrong", true},
			{"other_valid_hash", "---\nbased_on: '" + other.Hash + "'\n---\nInstructions.\n", other.Hash, true},
			{"duplicate", "---\nbased_on: '" + asset.Hash + "'\nbased_on: wrong\n---\nInstructions.\n", "", true},
			{"body_decoy", "---\nother: value\n---\nbased_on: '" + asset.Hash + "'\n", "", true},
			{"case_key", "---\nBased_On: '" + asset.Hash + "'\n---\nInstructions.\n", "", true},

			{"case", "---\nbased_on: " + strings.ToUpper(asset.Hash) + "\n---\nInstructions.\n", strings.ToUpper(asset.Hash), true},
			{"quoted_whitespace", "---\nbased_on: ' " + asset.Hash + " '\n---\nInstructions.\n", " " + asset.Hash + " ", true},
			{"numeric", "---\nbased_on: 42\n---\nInstructions.\n", "", true},
			{"null", "---\nbased_on: null\n---\nInstructions.\n", "", true},
			{"list", "---\nbased_on: [hash]\n---\nInstructions.\n", "", true},
			{"map", "---\nbased_on: {hash: nope}\n---\nInstructions.\n", "", true},
			{"malformed", "---\nbased_on: [\n---\nInstructions.\n", "", true},
			{"unclosed", "---\nbased_on: " + asset.Hash + "\n", "", true},
		}
		for _, row := range rows {
			t.Run(name+"_"+row.name, func(t *testing.T) {
				put(t, root, ".fireweed/playbooks/"+name+".md", row.text)
				playbookAssert(t, root, name, playbookWant(t, name, row.text, row.based, true, row.stale))
			})
		}
	}
}
func Test_PRJ021_PRJ022PublicFixtureControls(t *testing.T) {
	names := a.Playbooks()
	sort.Strings(names)
	if !reflect.DeepEqual(names, []string{"audit", "author-change", "evals", "impl", "regenerate"}) {
		t.Fatalf("public inventory %+v", names)
	}
	for _, name := range names {
		asset, err := a.Playbook(name)
		if err != nil || asset.Hash != playbookHash(asset.Text) {
			t.Fatal("embedded hash premise")
		}
		text := "---\nbased_on: '" + asset.Hash + "'\n---\nBody.\n"
		doc, err := c.ParseMarkdown(text)
		if err != nil || doc.Metadata["based_on"] != asset.Hash {
			t.Fatalf("public frontmatter premise %+v %v", doc, err)
		}
	}
}
func Test_PRJ021UnknownNameShortCircuitAndRepositoryErrors(t *testing.T) {
	if os.PathSeparator == '\\' {
		t.Skip("POSIX command traps")
	}
	traps := t.TempDir()
	marker := filepath.Join(traps, "attempted")
	put(t, traps, "git", "#!/bin/sh\nprintf attempted >> "+shellQuote(marker)+"\nexit 97\n")
	if err := os.Chmod(filepath.Join(traps, "git"), 0755); err != nil {
		t.Fatal(err)
	}
	t.Setenv("PATH", traps+string(os.PathListSeparator)+os.Getenv("PATH"))
	if err := exec.Command("git", "--version").Run(); err == nil {
		t.Fatal("trap positive control passed")
	}
	if _, err := os.Stat(marker); err != nil {
		t.Fatal("trap was not reached")
	}
	if err := os.Remove(marker); err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"", "unknown", "../impl", "/impl", "impl.md", "audit/../impl", "IMPL", "evals\\escape", "impl\x00ignored"} {
		t.Run(fmt.Sprintf("%q", name), func(t *testing.T) {
			got, err := p.ResolvePlaybook(filepath.Join(t.TempDir(), "not-even-present"), name)
			if err == nil || got != (p.Playbook{}) || !playbookCorrective(err.Error()) {
				t.Fatalf("invalid name accepted: %+v %v", got, err)
			}
			for _, valid := range a.Playbooks() {
				if !strings.Contains(err.Error(), valid) {
					t.Fatalf("valid-name guidance lacks %s: %v", valid, err)
				}
			}
			if _, err := os.Stat(marker); !os.IsNotExist(err) {
				t.Fatal("invalid name executed Git before validation")
			}
		})
	}
}
func Test_PRJ021KnownNameOutsideRepository(t *testing.T) {
	root := t.TempDir()
	if _, err := g.Open(root); err == nil {
		t.Fatal("fixture is a Git checkout")
	}
	got, err := p.ResolvePlaybook(root, "impl")
	if err == nil || got != (p.Playbook{}) || !playbookCorrective(err.Error()) {
		t.Fatalf("nonrepository result %+v %v", got, err)
	}
	lower := strings.ToLower(err.Error())
	if !strings.Contains(lower, "git") && !strings.Contains(lower, "repository") {
		t.Fatalf("missing repository-specific guidance: %v", err)
	}
}
func Test_PRJ023OverrideReadFailuresAndContainedLinks(t *testing.T) {
	for _, mode := range []string{"directory", "unreadable", "outside_file", "dangling_file", "outside_directory", "missing_outside_descendant", "inside_link"} {
		t.Run(mode, func(t *testing.T) {
			root := playbookRoot(t)
			path := ".fireweed/playbooks/evals.md"
			full := filepath.Join(root, filepath.FromSlash(path))
			outside := t.TempDir()
			put(t, outside, "private", "private instructions")
			put(t, root, path, "old override")
			switch mode {
			case "directory":
				if err := os.Remove(full); err != nil {
					t.Fatal(err)
				}
				if err := os.Mkdir(full, 0755); err != nil {
					t.Fatal(err)
				}
			case "unreadable":
				if err := os.Chmod(full, 0000); err != nil {
					t.Fatal(err)
				}
				defer os.Chmod(full, 0644)
				if _, err := os.ReadFile(full); !os.IsPermission(err) {
					t.Skipf("host does not establish unreadability: %v", err)
				}
			case "outside_file", "dangling_file", "inside_link":
				if err := os.Remove(full); err != nil {
					t.Fatal(err)
				}
				target := filepath.Join(outside, "private")
				if mode == "dangling_file" {
					target = filepath.Join(root, "missing-target")
				}
				if mode == "inside_link" {
					target = filepath.Join(root, "inside")
					put(t, root, "inside", "inside linked instructions")
				}
				if err := os.Symlink(target, full); err != nil {
					t.Skipf("symlink unavailable: %v", err)
				}
			case "outside_directory", "missing_outside_descendant":
				if mode == "outside_directory" {
					put(t, outside, "evals.md", "private instructions")
				}
				if err := os.RemoveAll(filepath.Dir(full)); err != nil {
					t.Fatal(err)
				}
				if err := os.Symlink(outside, filepath.Dir(full)); err != nil {
					t.Skipf("symlink unavailable: %v", err)
				}
			}
			if mode == "inside_link" {
				playbookAssert(t, root, "evals", playbookWant(t, "evals", "inside linked instructions", "", true, true))
			} else {
				playbookError(t, root, "evals", "D05", path)
			}
		})
	}
}

func playbookTree(t *testing.T, root string) map[string]string {
	t.Helper()
	result := map[string]string{}
	err := filepath.WalkDir(root, func(path string, entry os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		info, err := os.Lstat(path)
		if err != nil {
			return err
		}
		value := info.Mode().String()
		if info.Mode()&os.ModeSymlink != 0 {
			target, err := os.Readlink(path)
			if err != nil {
				return err
			}
			value += " link:" + target
		} else if info.Mode().IsRegular() {
			raw, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			value += " bytes:" + playbookHash(string(raw))
		}
		result[filepath.ToSlash(rel)] = value
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	return result
}
func playbookTraps(t *testing.T) string {
	t.Helper()
	traps := t.TempDir()
	marker := filepath.Join(traps, "attempted")
	realGit, err := exec.LookPath("git")
	if err != nil {
		t.Fatal(err)
	}
	put(t, traps, "git", "#!/bin/sh\nfor a in \"$@\";do case \"$a\" in fetch|push|clone|ls-remote|commit|branch|checkout|switch|reset|restore|add|worktree|update-ref|hash-object) printf attempted >> "+shellQuote(marker)+";exit 97;;esac;done\nexec "+shellQuote(realGit)+" \"$@\"\n")
	names := []string{"git", "hook-trap", "agent-trap", "codex", "claude", "fireweed", "fireweed-bootstrap", "curl", "wget", "ssh", "nc"}
	for _, name := range names {
		if name != "git" {
			put(t, traps, name, "#!/bin/sh\nprintf attempted >> "+shellQuote(marker)+"\nexit 97\n")
		}
		if err := os.Chmod(filepath.Join(traps, name), 0755); err != nil {
			t.Fatal(err)
		}
	}
	t.Setenv("PATH", traps+string(os.PathListSeparator)+os.Getenv("PATH"))
	t.Setenv("FIREWEED_BIN", filepath.Join(traps, "fireweed-bootstrap"))
	if err := exec.Command("git", "fetch", "forbidden.invalid").Run(); err == nil {
		t.Fatal("trap control accepted fetch")
	}
	if _, err := os.Stat(marker); err != nil {
		t.Fatal("trap positive control absent")
	}
	if err := os.Remove(marker); err != nil {
		t.Fatal(err)
	}
	return marker
}
func Test_PRJ021_PRJ022_PRJ023ValueAndInspectionSafety(t *testing.T) {
	if os.PathSeparator == '\\' {
		t.Skip("POSIX trap/safety fixture")
	}
	for _, mode := range []string{"embedded", "override", "stale", "read_error", "unknown"} {
		for _, linked := range []bool{false, true} {
			t.Run(fmt.Sprintf("%s_linked_%t", mode, linked), func(t *testing.T) {
				root := playbookRoot(t)
				base := strings.TrimSpace(git(t, root, "rev-parse", "HEAD"))
				put(t, root, ".git/claim-fixture.json", fmt.Sprintf(`{"worktree":%q,"change":"guard","run":"fixture","base":%q,"state":""}`, root, base))
				blob := strings.TrimSpace(git(t, root, "hash-object", "-w", ".git/claim-fixture.json"))
				git(t, root, "update-ref", "refs/fireweed/claims/guard", blob)
				git(t, root, "update-ref", "refs/custom/guard", base)
				if linked {
					peer := filepath.Join(t.TempDir(), "peer")
					git(t, root, "worktree", "add", "--detach", peer, "HEAD")
					root = peer
				}
				common := strings.TrimSpace(git(t, root, "rev-parse", "--git-common-dir"))
				if !filepath.IsAbs(common) {
					common = filepath.Join(root, common)
				}
				common, err := filepath.EvalSymlinks(common)
				if err != nil {
					t.Fatal(err)
				}
				asset, _ := a.Playbook("impl")
				name := "impl"
				path := ".fireweed/playbooks/impl.md"
				text := "---\nbased_on: '" + asset.Hash + "'\n---\nProject instructions.\n"
				outside := t.TempDir()
				put(t, outside, "private", "retained outside bytes")
				if mode == "override" || mode == "stale" {
					if mode == "stale" {
						text = "stale owned instructions\n"
					}
					put(t, root, path, text)
				}
				if mode == "read_error" {
					put(t, root, ".fireweed/playbooks/.keep", "")
					if err := os.Symlink(filepath.Join(outside, "private"), filepath.Join(root, path)); err != nil {
						t.Skipf("symlink unavailable: %v", err)
					}
				}
				if mode == "unknown" {
					name = "../impl"
				}
				marker := playbookTraps(t)
				temporary := t.TempDir()
				temporary, err = filepath.EvalSymlinks(temporary)
				if err != nil {
					t.Fatal(err)
				}
				t.Setenv("TMPDIR", temporary)
				t.Setenv("TMP", temporary)
				t.Setenv("TEMP", temporary)
				git(t, root, "--version")
				beforeFiles, beforeOutside, beforeTemp := playbookTree(t, root), playbookTree(t, outside), playbookTree(t, temporary)
				beforeCommon := playbookTree(t, common)
				beforeEnv := os.Environ()
				cwd, err := os.Getwd()
				if err != nil {
					t.Fatal(err)
				}
				savedRoot, savedName := root, name
				first, e1 := p.ResolvePlaybook(root, name)
				second, e2 := p.ResolvePlaybook(root, name)
				if mode == "unknown" || mode == "read_error" {
					if e1 == nil || e2 == nil || first != (p.Playbook{}) || second != (p.Playbook{}) || e1.Error() != e2.Error() {
						t.Fatalf("error result %+v %v %+v %v", first, e1, second, e2)
					}
				} else {
					want := playbookWant(t, name, asset.Text, "", false, false)
					if mode == "override" {
						want = playbookWant(t, name, text, asset.Hash, true, false)
					}
					if mode == "stale" {
						want = playbookWant(t, name, text, "", true, true)
					}
					if e1 != nil || e2 != nil || first != want || second != want {
						t.Fatalf("resolution %+v %v %+v %v; want %+v", first, e1, second, e2, want)
					}
					first.Name = "mutated"
					first.Text = "mutated"
					first.Hash = "mutated"
					first.BasedOn = "mutated"
					first.EmbeddedHash = "mutated"
					first.Overridden = !first.Overridden
					first.Stale = !first.Stale
					if second != want {
						t.Fatal("results alias")
					}
					playbookAssert(t, root, name, want)
				}
				after, err := os.Getwd()
				if err != nil || after != cwd || !reflect.DeepEqual(beforeEnv, os.Environ()) || root != savedRoot || name != savedName || !reflect.DeepEqual(beforeFiles, playbookTree(t, root)) || !reflect.DeepEqual(beforeOutside, playbookTree(t, outside)) || !reflect.DeepEqual(beforeTemp, playbookTree(t, temporary)) || !reflect.DeepEqual(beforeCommon, playbookTree(t, common)) {
					t.Fatal("resolution changed files, Git/index/refs, private/temp storage, cwd/environment or inputs")
				}
				if _, err := os.Stat(marker); !os.IsNotExist(err) {
					t.Fatal("resolution ran a forbidden command")
				}
			})
		}
	}
}

func Test_PRJ021NestedOverrideCanonicalLookup(t *testing.T) {
	for _, linked := range []bool{false, true} {
		t.Run(fmt.Sprint(linked), func(t *testing.T) {
			root := playbookRoot(t)
			if linked {
				peer := filepath.Join(t.TempDir(), "peer")
				git(t, root, "worktree", "add", "--detach", peer, "HEAD")
				root = peer
			}
			asset, _ := a.Playbook("impl")
			text := "---\nbased_on: '" + asset.Hash + "'\n---\nRoot instructions.\n"
			put(t, root, ".fireweed/playbooks/impl.md", text)
			nested := filepath.Join(root, "nested")
			put(t, nested, ".fireweed/playbooks/impl.md", "conflicting nested decoy")
			playbookAssert(t, nested, "impl", playbookWant(t, "impl", text, asset.Hash, true, false))
		})
	}
}
func Test_PRJ022FrontmatterFixtureControls(t *testing.T) {
	asset, _ := a.Playbook("evals")
	duplicate := "---\nbased_on: '" + asset.Hash + "'\nbased_on: wrong\n---\nBody.\n"
	if _, err := c.ParseMarkdown(duplicate); err == nil {
		t.Fatal("public parser must reject duplicate-key premise")
	}
	for _, text := range []string{"---\nother: value\n---\nbased_on: '" + asset.Hash + "'\n", "---\nBased_On: '" + asset.Hash + "'\n---\nBody.\n"} {
		doc, err := c.ParseMarkdown(text)
		if err != nil {
			t.Fatal(err)
		}
		if _, ok := doc.Metadata["based_on"]; ok {
			t.Fatal("decoy is actually based_on metadata")
		}
	}
}
