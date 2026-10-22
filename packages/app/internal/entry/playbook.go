package entry

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/lurkshark/fireweed/packages/assets"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project"
)

// Only metadata is needed to decide whether a standalone instruction read binds
// to a project. Stop at the nearest checkout even when it has no configuration.
func playbookProjectRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	dir, err = filepath.EvalSymlinks(dir)
	if err != nil {
		return "", err
	}
	for {
		marker, err := os.Lstat(filepath.Join(dir, ".git"))
		if err == nil {
			if !marker.IsDir() && !marker.Mode().IsRegular() && marker.Mode()&os.ModeSymlink == 0 {
				return "", fmt.Errorf("restore a regular .git marker or Git directory")
			}
			_, err := os.Lstat(filepath.Join(dir, "fireweed.yaml"))
			if os.IsNotExist(err) {
				return "", nil
			}
			if err != nil {
				return "", err
			}
			return dir, nil
		}
		if !os.IsNotExist(err) {
			return "", err
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", nil
		}
		dir = parent
	}
}

func playbook(out, errOut io.Writer, option options, fail func(string) int) int {
	name := option.operands[0]
	embedded, err := assets.Playbook(name)
	if err != nil {
		return fail(fmt.Sprintf("unknown playbook %q; valid choices: %s", name, strings.Join(assets.Playbooks(), ", ")))
	}
	selected := project.Playbook{Name: name, Text: embedded.Text, Hash: embedded.Hash, EmbeddedHash: embedded.Hash}
	root, err := playbookProjectRoot()
	if err != nil {
		return loadFailure(out, errOut, option, err)
	}
	if root != "" {
		repo, err := gitstore.Open(root)
		if err != nil {
			return loadFailure(out, errOut, option, err)
		}
		if repo.Root != root {
			return loadFailure(out, errOut, option, fmt.Errorf("Git checkout does not match the enclosing project marker"))
		}
		full, err := gitstore.SafePath(repo.Root, "fireweed.yaml")
		if err != nil {
			return loadFailure(out, errOut, option, err)
		}
		info, err := os.Stat(full)
		if err != nil || !info.Mode().IsRegular() {
			return loadFailure(out, errOut, option, fmt.Errorf("restore fireweed.yaml as a readable contained regular configuration file"))
		}
		raw, err := os.ReadFile(full)
		if err != nil {
			return loadFailure(out, errOut, option, err)
		}
		if _, err := contracts.ParseConfig(string(raw), assets.Info().Version); err != nil {
			return loadFailure(out, errOut, option, err)
		}
		selected, err = project.ResolvePlaybook(repo.Root, name)
		if err != nil {
			if !strings.HasPrefix(err.Error(), "D05:") {
				return loadFailure(out, errOut, option, err)
			}
			return emit(out, errOut, option, nil, []diagnostic{{Code: "D05", Severity: "error", Path: ".fireweed/playbooks/" + name + ".md", Message: err.Error()}}, 2)
		}
	}
	if option.json {
		return emit(out, errOut, option, map[string]any{"name": selected.Name, "text": selected.Text, "hash": selected.Hash, "embedded_hash": selected.EmbeddedHash, "overridden": selected.Overridden}, nil, 0)
	}
	return emit(out, errOut, option, selected.Text, nil, 0)
}
