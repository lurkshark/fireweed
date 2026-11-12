package evals_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/project"
)

func Test_PRJ_070_CheckWarnsForTrackedSpecsOutsideModuleGlobs(t *testing.T) {
	root := fixture(t)
	revisionGit(t, root, "init", "-q")
	put(t, root, "fireweed.yaml", strings.Replace(fixtureConfig, `"packages/*"`, `"packages/alpha", "packages/bet*"`, 1))
	put(t, root, "packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\n---\n# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST work.\n")
	put(t, root, "packages/alpha/recipes/nested/SPEC.md", "Module template.\n")
	put(t, root, ".fireweed/templates/module/SPEC.md", "Project template.\n")
	put(t, root, "extras/OTHER.md", "Tracked file with a different name.\n")
	put(t, root, ".gitignore", "node_modules/\n")
	revisionCommit(t, root, "record configured modules and templates")

	put(t, root, "node_modules/ignored/SPEC.md", "Ignored document.\n")
	put(t, root, "extras/untracked/SPEC.md", "Untracked document.\n")
	const indexOnly = "extras/space module/SPEC.md"
	wantPaths := []string{"packages/alpha-moved/SPEC.md", indexOnly, ".fireweed-extra/SPEC.md"}
	for _, path := range wantPaths {
		put(t, root, path, "Document outside configured module directories.\n")
	}
	revisionGit(t, root, append([]string{"add", "--"}, wantPaths...)...)
	if err := os.Remove(filepath.Join(root, filepath.FromSlash(indexOnly))); err != nil {
		t.Fatal(err)
	}

	counts := make(map[string]int)
	for _, diagnostic := range project.Check(root) {
		if diagnostic.Code != "C20" {
			continue
		}
		named := 0
		for _, path := range wantPaths {
			if strings.Contains(diagnostic.Message, path) {
				counts[path]++
				named++
			}
		}
		if !diagnostic.Warning || named != 1 {
			t.Errorf("unexpected C20 diagnostic %#v; want a warning naming exactly one of %q", diagnostic, wantPaths)
		}
	}
	for _, path := range wantPaths {
		if counts[path] != 1 {
			t.Errorf("C20 diagnostics naming %q = %d; want 1", path, counts[path])
		}
	}
}
