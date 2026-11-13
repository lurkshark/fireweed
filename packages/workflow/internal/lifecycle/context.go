package lifecycle

import (
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
)

func Context(root, selector string, full bool) (string, error) {
	state, _, err := loadHeld(root, selector)
	if err != nil {
		return "", err
	}
	root = state.Snapshot.Root
	module := moduleNamed(state.Snapshot, state.Change.Module)
	spec, err := contracts.RenderDocument(module.Document)
	if err != nil {
		return "", err
	}
	change, err := os.ReadFile(disk(root, state.Change.Path))
	if err != nil {
		return "", err
	}
	book, err := playbook(root, playbookName(state))
	if err != nil {
		return "", err
	}

	var body strings.Builder
	fmt.Fprintf(&body, "# Change\n%s\n\n# Projected spec\n%s\n", change, spec)
	var other []contextFile
	body.WriteString("\n# Interface\n")
	err = walkModule(state.Snapshot, module, func(file string, entry fs.DirEntry) error {
		info, err := entry.Info()
		if err != nil {
			return err
		}
		switch fileZone(state.Snapshot, module, file) {
		case "interface":
			data, err := os.ReadFile(disk(root, file))
			if err != nil {
				return err
			}
			fmt.Fprintf(&body, "\n## %s\n%s\n", file, data)
		case "spec", "protected":
		default:
			other = append(other, contextFile{file, info.Size(), info.Mode().IsRegular()})
		}
		return nil
	})
	if err != nil {
		return "", err
	}
	system, err := os.ReadFile(disk(root, "SYSTEM.md"))
	if err != nil && !os.IsNotExist(err) {
		return "", err
	}
	fmt.Fprintf(&body, "\n# SYSTEM.md\n%s\n# Playbook\n%s\n", system, book)
	if state.ClaimOID != "" && (state.Phase == "evals" || state.Phase == "impl") {
		fmt.Fprintf(&body, "To declare a block, set `blocked` to a nonempty reason in the YAML frontmatter of `%s` and write your questions under `## Questions`. The controller will save the attempt and record the block.\n", state.Change.Path)
	}
	body.WriteString("\n# Other relevant files\n")
	for _, file := range other {
		fmt.Fprintf(&body, "- %s (%d bytes)\n", file.path, file.size)
		if full && file.regular {
			data, err := os.ReadFile(disk(root, file.path))
			if err != nil {
				return "", err
			}
			fmt.Fprintf(&body, "%s\n", data)
		}
	}
	return body.String(), nil
}

type contextFile struct {
	path    string
	size    int64
	regular bool
}

func walkModule(snapshot project.Snapshot, module project.Module, visit func(string, fs.DirEntry) error) error {
	if module.Name == "" {
		return fmt.Errorf("module is absent from the project")
	}
	return filepath.WalkDir(disk(snapshot.Root, module.Dir), func(file string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(snapshot.Root, file)
		if err != nil {
			return err
		}
		name := filepath.ToSlash(rel)
		if entry.IsDir() {
			if name != module.Dir && (path.Base(name) == ".git" || ownerOf(snapshot, name+"/").Name != module.Name) {
				return filepath.SkipDir
			}
			return nil
		}
		return visit(name, entry)
	})
}
