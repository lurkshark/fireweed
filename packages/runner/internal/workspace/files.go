package workspace

import (
	"context"
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project"
)

func Git(ctx context.Context, root string, args ...string) (string, error) {
	result, err := gitstore.Run(ctx, gitstore.Request{
		Dir: root, Args: append([]string{"git", "-c", "core.hooksPath=/dev/null"}, args...),
		Timeout: 30 * time.Second,
		Env:     map[string]string{"GIT_NO_LAZY_FETCH": "1", "GIT_TERMINAL_PROMPT": "0"},
	})
	if err != nil {
		return "", fmt.Errorf("git %s: %w: %s", args[0], err, strings.TrimSpace(result.Stderr))
	}
	return result.Stdout, nil
}

func Path(root, relative string) string {
	return filepath.Join(root, filepath.FromSlash(relative))
}

func Write(root, relative string, data []byte, mode fs.FileMode) error {
	file := Path(root, relative)
	if err := os.MkdirAll(filepath.Dir(file), 0755); err != nil {
		return err
	}
	return os.WriteFile(file, data, mode)
}

type File struct {
	Data string
	Mode fs.FileMode
}

func Read(root, relative string) (File, error) {
	file := Path(root, relative)
	info, err := os.Lstat(file)
	if err != nil {
		return File{}, err
	}
	switch {
	case info.Mode().IsRegular():
		data, err := os.ReadFile(file)
		return File{Data: string(data), Mode: info.Mode()}, err
	case info.Mode()&os.ModeSymlink != 0:
		target, err := os.Readlink(file)
		return File{Data: target, Mode: info.Mode()}, err
	default:
		return File{Mode: info.Mode()}, nil
	}
}

func (file File) Put(root, relative string) error {
	if file.Mode.IsRegular() {
		return Write(root, relative, []byte(file.Data), file.Mode.Perm())
	}
	if file.Mode&os.ModeSymlink != 0 {
		dst := Path(root, relative)
		if err := os.MkdirAll(filepath.Dir(dst), 0755); err != nil {
			return err
		}
		return os.Symlink(file.Data, dst)
	}
	return fmt.Errorf("unsupported file type at %s", relative)
}

// Walk presents all files outside the workspace's own Git directory.
func Walk(root string, visit func(string, fs.DirEntry) error) error {
	return filepath.WalkDir(root, func(full string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		name, err := filepath.Rel(root, full)
		if err != nil {
			return err
		}
		name = filepath.ToSlash(name)
		if name == ".git" {
			if entry.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}
		return visit(name, entry)
	})
}

func Owner(snapshot project.Snapshot, file string) project.Module {
	var owner project.Module
	for _, module := range snapshot.Modules {
		if (module.Dir == "." || strings.HasPrefix(file, module.Dir+"/")) && len(module.Dir) > len(owner.Dir) {
			owner = module
		}
	}
	return owner
}

func Zone(snapshot project.Snapshot, module project.Module, file string) string {
	if file == "SYSTEM.md" || file == "fireweed.yaml" || strings.HasPrefix(file, ".fireweed/") ||
		path.Base(file) == "AGENTS.md" || path.Base(file) == "CLAUDE.md" {
		return "protected"
	}
	match := func(pattern string) bool { return contracts.MatchGlob(pattern, file) }
	switch patterns := snapshot.Config["protected"].(type) {
	case []any:
		for _, pattern := range patterns {
			if value, ok := pattern.(string); ok && match(value) {
				return "protected"
			}
		}
	case []string:
		for _, pattern := range patterns {
			if match(pattern) {
				return "protected"
			}
		}
	}
	if module.Dir != "." {
		file = strings.TrimPrefix(file, module.Dir+"/")
	}
	return contracts.ResolveZone(file, module.Zones)
}

func Init(ctx context.Context, root string) error {
	_, err := Git(ctx, root, "init", "--quiet", "--template=")
	return err
}

// Seed records only the files explicitly copied into a new repository.
func Seed(ctx context.Context, root string) error {
	if _, err := Git(ctx, root, "add", "--force", "--all"); err != nil {
		return err
	}
	_, err := Git(ctx, root, "-c", "user.name=Fireweed", "-c", "user.email=fireweed@localhost",
		"-c", "commit.gpgSign=false", "commit", "--quiet", "--allow-empty", "-m", "Fireweed input")
	return err
}
