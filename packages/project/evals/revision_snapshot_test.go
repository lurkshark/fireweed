package evals_test

import (
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/project"
)

func revisionGit(t *testing.T, root string, args ...string) string {
	t.Helper()
	cmd := exec.Command("git", append([]string{"-C", root}, args...)...)
	cmd.Env = append(os.Environ(), "GIT_CONFIG_GLOBAL=/dev/null", "GIT_CONFIG_NOSYSTEM=1")
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("git %v: %v: %s", args, err, out)
	}
	return strings.TrimSpace(string(out))
}

func revisionCommit(t *testing.T, root, subject string) string {
	t.Helper()
	revisionGit(t, root, "add", "-A")
	revisionGit(t, root, "-c", "user.name=Fireweed Fixture", "-c", "user.email=fixture@example.invalid", "commit", "-qm", subject)
	return revisionGit(t, root, "rev-parse", "HEAD")
}

func historicalProject(t *testing.T) (string, string) {
	t.Helper()
	root := t.TempDir()
	revisionGit(t, root, "init", "-q")
	put(t, root, "fireweed.yaml", `version: 1
modules: ["packages/*"]
module_defaults:
  protected: [guard/**]
  evals: [evals/**]
  interface: [api.go]
  append_only: [ledger/**]
  preserve: [data/**]
  disposable: [internal/**]
`)
	put(t, root, "packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\nevals: []\n---\n# Alpha\n\n## Specs\n\n### ALP-001\nHistorical alpha behavior.\n")
	put(t, root, "packages/beta/SPEC.md", "---\nname: beta\nid_prefix: BET\ndepends_on: [alpha]\n---\n# Beta\n\n## Specs\n\n### BET-001\nHistorical beta behavior.\n")
	recorded := revisionCommit(t, root, "record original project")

	put(t, root, "fireweed.yaml", "version: 2\nmodules: [\"components/*\"]\nmodule_defaults:\n  evals: [later/**]\n")
	put(t, root, "packages/alpha/SPEC.md", "---\nname: later-alpha\nid_prefix: LAL\n---\n# Later\n\n## Specs\n")
	put(t, root, "components/next/SPEC.md", "---\nname: next\nid_prefix: NXT\n---\n# Next\n\n## Specs\n")
	revisionCommit(t, root, "change config glob and modules")

	put(t, root, "fireweed.yaml", "version: 3\nmodules: [\"dirty/*\"]\n")
	put(t, root, "packages/alpha/SPEC.md", "---\nname: dirty-alpha\nid_prefix: DAL\n---\n# Dirty\n\n## Specs\n")
	put(t, root, "dirty/only/SPEC.md", "---\nname: dirty\nid_prefix: DRT\n---\n# Dirty\n\n## Specs\n")
	return root, recorded
}

func Test_PRJ_047_LoadAtRecordedConfigGlobsAndDocuments(t *testing.T) {
	root, recorded := historicalProject(t)
	got, err := project.LoadAt(root, recorded)
	if err != nil {
		t.Fatalf("LoadAt(%s): %v", recorded, err)
	}
	if got.Config["version"] != 1 || !reflect.DeepEqual(got.Config["modules"], []any{"packages/*"}) {
		t.Errorf("recorded configuration = %#v", got.Config)
	}
	byName := make(map[string]project.Module)
	for _, module := range got.Modules {
		byName[module.Name] = module
	}
	if len(got.Modules) != 2 || len(byName) != 2 {
		t.Fatalf("recorded modules = %#v; want alpha and beta only", got.Modules)
	}
	alpha, beta := byName["alpha"], byName["beta"]
	if alpha.IDPrefix != "ALP" || alpha.Document.Frontmatter["name"] != "alpha" || !strings.Contains(alpha.Document.Body, "Historical alpha behavior.") {
		t.Errorf("recorded alpha document = %#v", alpha)
	}
	if beta.IDPrefix != "BET" || !reflect.DeepEqual(beta.DependsOn, []string{"alpha"}) || !strings.Contains(beta.Document.Body, "Historical beta behavior.") {
		t.Errorf("recorded beta document = %#v", beta)
	}
}

func Test_PRJ_048_LoadAtRecordedDefaultsAndExplicitEmptyZones(t *testing.T) {
	root, recorded := historicalProject(t)
	got, err := project.LoadAt(root, recorded)
	if err != nil {
		t.Fatalf("LoadAt(%s): %v", recorded, err)
	}
	byName := make(map[string]project.Module)
	for _, module := range got.Modules {
		byName[module.Name] = module
	}
	if len(got.Modules) != 2 {
		t.Fatalf("recorded modules = %#v", got.Modules)
	}
	alpha, beta := byName["alpha"], byName["beta"]
	if !reflect.DeepEqual(alpha.Zones.Protected, []string{"guard/**"}) ||
		!reflect.DeepEqual(alpha.Zones.Interface, []string{"api.go"}) ||
		!reflect.DeepEqual(alpha.Zones.AppendOnly, []string{"ledger/**"}) ||
		!reflect.DeepEqual(alpha.Zones.Preserve, []string{"data/**"}) ||
		!reflect.DeepEqual(alpha.Zones.Disposable, []string{"internal/**"}) {
		t.Errorf("recorded inherited zones = %#v", alpha.Zones)
	}
	if alpha.Zones.Evals == nil || len(alpha.Zones.Evals) != 0 {
		t.Errorf("explicit empty evals zone = %#v", alpha.Zones.Evals)
	}
	if !reflect.DeepEqual(beta.Zones.Evals, []string{"evals/**"}) {
		t.Errorf("recorded default evals zone = %#v", beta.Zones.Evals)
	}
}

func Test_PRJ_049_LoadAtCanonicalRootAndRelativeModuleDirs(t *testing.T) {
	root, recorded := historicalProject(t)
	canonical, err := filepath.EvalSymlinks(root)
	if err != nil {
		t.Fatal(err)
	}
	got, err := project.LoadAt(root+string(filepath.Separator)+".", recorded)
	if err != nil {
		t.Fatalf("LoadAt(%s): %v", recorded, err)
	}
	if got.Root != canonical {
		t.Errorf("Root = %q; want canonical %q", got.Root, canonical)
	}
	dirs := make(map[string]string)
	for _, module := range got.Modules {
		dirs[module.Name] = module.Dir
	}
	if !reflect.DeepEqual(dirs, map[string]string{"alpha": "packages/alpha", "beta": "packages/beta"}) {
		t.Errorf("module directories = %#v", dirs)
	}
}

func Test_PRJ_050_LoadAtRejectsInvalidRecordedInputs(t *testing.T) {
	root := t.TempDir()
	revisionGit(t, root, "init", "-q")
	put(t, root, "README.md", "repository without fireweed.yaml\n")
	missingConfig := revisionCommit(t, root, "missing config")
	put(t, root, "fireweed.yaml", "modules: [broken\n")
	malformedConfig := revisionCommit(t, root, "malformed config")
	put(t, root, "fireweed.yaml", "modules: [\"modules/*\"]\n")
	put(t, root, "modules/ghost/README.md", "directory without SPEC.md\n")
	missingSpec := revisionCommit(t, root, "missing module document")
	for _, tc := range []struct{ name, revision string }{
		{"invalid_revision", "not-a-revision"},
		{"missing_recorded_config", missingConfig},
		{"malformed_recorded_config", malformedConfig},
		{"missing_recorded_SPEC", missingSpec},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if _, err := project.LoadAt(root, tc.revision); err == nil {
				t.Errorf("LoadAt accepted %s (%s)", tc.name, tc.revision)
			}
		})
	}
}
