package fixture

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"slices"
	"strings"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project/internal/source"
	"github.com/lurkshark/fireweed/packages/project/model"
	"gopkg.in/yaml.v3"
)

type projectSource struct {
	read func(string) ([]byte, error)
	dirs func([]string) ([]string, error)
}

func Load(root string) (model.Snapshot, error) {
	return load(root, projectSource{
		read: func(file string) ([]byte, error) { return os.ReadFile(source.Path(root, file)) },
		dirs: func(patterns []string) ([]string, error) { return workingDirs(root, patterns) },
	})
}

func LoadAt(root, revision string) (model.Snapshot, error) {
	canonical, err := source.Git(root, "rev-parse", "--show-toplevel")
	if err != nil {
		return model.Snapshot{}, err
	}
	root, err = filepath.EvalSymlinks(strings.TrimSuffix(canonical, "\n"))
	if err != nil {
		return model.Snapshot{}, err
	}
	commit, err := source.Git(root, "rev-parse", "--verify", "--end-of-options", revision+"^{commit}")
	if err != nil {
		return model.Snapshot{}, err
	}
	commit = strings.TrimSpace(commit)
	return load(root, projectSource{
		read: func(file string) ([]byte, error) { return gitstore.ReadAt(root, commit, file) },
		dirs: func(patterns []string) ([]string, error) {
			if len(patterns) == 0 {
				return nil, nil
			}
			listing, err := source.Git(root, "ls-tree", "-r", "-d", "--name-only", "-z", commit)
			if err != nil {
				return nil, err
			}
			var dirs []string
			for _, dir := range append([]string{"."}, strings.Split(strings.TrimSuffix(listing, "\x00"), "\x00")...) {
				if dir != "" && slices.ContainsFunc(patterns, func(pattern string) bool {
					return contracts.MatchGlob(pattern, dir)
				}) {
					dirs = append(dirs, dir)
				}
			}
			return dirs, nil
		},
	})
}

func load(root string, files projectSource) (model.Snapshot, error) {
	const configPath = "fireweed.yaml"
	data, err := files.read(configPath)
	if err != nil {
		return model.Snapshot{}, source.FileError("C01", configPath, err)
	}
	config, patterns, defaults, err := decodeConfig(data)
	if err != nil {
		return model.Snapshot{}, source.FileError("C01", configPath, err)
	}
	dirs, err := files.dirs(patterns)
	if err != nil {
		return model.Snapshot{}, source.FileError("C01", configPath, err)
	}
	slices.Sort(dirs)
	snapshot := model.Snapshot{Root: root, Config: config}
	for _, dir := range slices.Compact(dirs) {
		file := path.Join(dir, "SPEC.md")
		data, err := files.read(file)
		if err != nil {
			return model.Snapshot{}, source.FileError("C03", file, err)
		}
		doc, err := contracts.ParseDocument(string(data))
		if err != nil {
			return model.Snapshot{}, source.DocumentError(file, err)
		}
		if _, err := source.Items(doc, "Specs"); err != nil {
			return model.Snapshot{}, source.DocumentError(file, err)
		}
		module, err := ReadModule(dir, doc, defaults)
		if err != nil {
			return model.Snapshot{}, source.FileError("C12", file, err)
		}
		snapshot.Modules = append(snapshot.Modules, module)
	}
	return snapshot, nil
}

func decodeConfig(data []byte) (map[string]any, []string, contracts.Zones, error) {
	var config map[string]any
	var defaults contracts.Zones
	decoder := yaml.NewDecoder(strings.NewReader(string(data)))
	if err := decoder.Decode(&config); err != nil {
		return nil, nil, defaults, err
	}
	if config == nil {
		return nil, nil, defaults, errors.New("configuration must be a mapping")
	}
	var extra any
	if err := decoder.Decode(&extra); !errors.Is(err, io.EOF) {
		if err == nil {
			err = errors.New("configuration must contain one YAML document")
		}
		return nil, nil, defaults, err
	}

	patterns, err := stringList("modules", config["modules"])
	if err != nil {
		return nil, nil, defaults, err
	}
	for i, pattern := range patterns {
		cleaned := path.Clean(pattern)
		if pattern == "" || path.IsAbs(cleaned) || cleaned == ".." || strings.HasPrefix(cleaned, "../") {
			return nil, nil, defaults, fmt.Errorf("invalid module glob %q", pattern)
		}
		patterns[i] = cleaned
	}
	if raw, present := config["module_defaults"]; present {
		fields, ok := raw.(map[string]any)
		if !ok {
			return nil, nil, defaults, errors.New("module_defaults must be a mapping")
		}
		defaults, err = readZones(fields, defaults)
		if err != nil {
			return nil, nil, defaults, fmt.Errorf("module_defaults: %w", err)
		}
	}
	if err := setSpecLimits(config); err != nil {
		return nil, nil, defaults, err
	}
	return config, patterns, defaults, nil
}

func setSpecLimits(config map[string]any) error {
	limits := make(map[string]any)
	if raw, present := config["spec_limits"]; present {
		var ok bool
		limits, ok = raw.(map[string]any)
		if !ok {
			return errors.New("spec_limits must be a mapping")
		}
	}
	for _, limit := range []struct {
		key   string
		value int
	}{{"max_clauses", 3}, {"max_words", 80}, {"max_decision_words", 300}} {
		raw, present := limits[limit.key]
		if !present {
			limits[limit.key] = limit.value
			continue
		}
		if value, ok := raw.(int); !ok || value <= 0 {
			return fmt.Errorf("spec_limits.%s must be a positive integer", limit.key)
		}
	}
	config["spec_limits"] = limits
	return nil
}

func workingDirs(root string, patterns []string) ([]string, error) {
	var dirs []string
	for _, pattern := range patterns {
		if pattern == "." {
			dirs = append(dirs, ".")
			continue
		}
		matches, err := matchDirs(root, ".", strings.Split(pattern, "/"))
		if err != nil {
			return nil, err
		}
		dirs = append(dirs, matches...)
	}
	return dirs, nil
}

func matchDirs(root, dir string, parts []string) ([]string, error) {
	if len(parts) == 0 {
		return []string{dir}, nil
	}
	component := parts[0]
	if !strings.Contains(component, "*") {
		child := path.Join(dir, component)
		info, err := os.Stat(source.Path(root, child))
		if errors.Is(err, os.ErrNotExist) || (err == nil && !info.IsDir()) {
			return nil, nil
		}
		if err != nil {
			return nil, err
		}
		return matchDirs(root, child, parts[1:])
	}

	var dirs []string
	if component == "**" {
		matches, err := matchDirs(root, dir, parts[1:])
		if err != nil {
			return nil, err
		}
		dirs = matches
	}
	entries, err := os.ReadDir(source.Path(root, dir))
	if err != nil {
		return nil, err
	}
	for _, entry := range entries {
		if !entry.IsDir() || !contracts.MatchGlob(component, entry.Name()) {
			continue
		}
		remaining := parts[1:]
		if component == "**" {
			remaining = parts
		}
		matches, err := matchDirs(root, path.Join(dir, entry.Name()), remaining)
		if err != nil {
			return nil, err
		}
		dirs = append(dirs, matches...)
	}
	return dirs, nil
}
