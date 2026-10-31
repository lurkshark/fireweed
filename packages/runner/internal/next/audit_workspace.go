package next

import (
	"bytes"
	"context"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/workflow"
)

// auditWorkspace holds the exact inputs supplied to one auditor invocation.
type auditWorkspace struct {
	root     string
	change   string
	original map[string][]byte
}

func newAuditWorkspace(state workflow.State) (_ *auditWorkspace, err error) {
	root, err := os.MkdirTemp("", "fireweed-audit-*")
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			os.RemoveAll(root)
		}
	}()
	w := &auditWorkspace{root: root, change: state.Change.Path, original: make(map[string][]byte)}
	if err = w.copyFile(state.Snapshot.Root, state.Change.Path); err != nil {
		return nil, err
	}
	var moduleDir string
	for _, module := range state.Snapshot.Modules {
		if module.Name != state.Change.Module {
			continue
		}
		moduleDir = module.Dir
		var spec string
		spec, err = contracts.RenderDocument(module.Document)
		if err != nil {
			return nil, fmt.Errorf("render projected spec: %w", err)
		}
		if err = w.writeFile(filepath.ToSlash(filepath.Join(module.Dir, "SPEC.md")), []byte(spec)); err != nil {
			return nil, err
		}
		moduleRoot := filepath.Join(state.Snapshot.Root, filepath.FromSlash(module.Dir))
		err = filepath.WalkDir(moduleRoot, func(path string, entry fs.DirEntry, walkErr error) error {
			if walkErr != nil {
				return walkErr
			}
			if entry.IsDir() {
				return nil
			}
			rel, relErr := filepath.Rel(moduleRoot, path)
			if relErr != nil {
				return relErr
			}
			for _, pattern := range module.Zones.Evals {
				if contracts.MatchGlob(pattern, filepath.ToSlash(rel)) {
					return w.copyFile(state.Snapshot.Root, filepath.ToSlash(filepath.Join(module.Dir, rel)))
				}
			}
			return nil
		})
		if err != nil {
			return nil, fmt.Errorf("copy audit evals: %w", err)
		}
		break
	}
	if moduleDir == "" {
		return nil, fmt.Errorf("held module %q is absent", state.Change.Module)
	}
	if err = auditGit(root, "init", "-q", "--template="); err != nil {
		return nil, err
	}
	if err = auditGit(root, "add", "-A"); err != nil {
		return nil, err
	}
	if err = auditGit(root, "-c", "user.name=Fireweed Audit", "-c", "user.email=audit@localhost", "commit", "-q", "-m", "audit inputs"); err != nil {
		return nil, err
	}
	return w, nil
}

func (w *auditWorkspace) copyFile(sourceRoot, rel string) error {
	data, err := os.ReadFile(filepath.Join(sourceRoot, filepath.FromSlash(rel)))
	if err != nil {
		return fmt.Errorf("read audit input %s: %w", rel, err)
	}
	return w.writeFile(rel, data)
}

func (w *auditWorkspace) writeFile(rel string, data []byte) error {
	path := filepath.Join(w.root, filepath.FromSlash(rel))
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	if err := os.WriteFile(path, data, 0644); err != nil {
		return err
	}
	w.original[rel] = data
	return nil
}

func auditGit(root string, args ...string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, "git", args...)
	cmd.Dir = root
	// Do not inherit repository selectors, alternates, or global Git settings.
	for _, value := range os.Environ() {
		if !strings.HasPrefix(value, "GIT_") {
			cmd.Env = append(cmd.Env, value)
		}
	}
	cmd.Env = append(cmd.Env, "GIT_CONFIG_NOSYSTEM=1", "GIT_CONFIG_GLOBAL=/dev/null", "GIT_NO_LAZY_FETCH=1", "GIT_TERMINAL_PROMPT=0")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("audit git %v: %w: %s", args, err, strings.TrimSpace(string(output)))
	}
	return nil
}

func (w *auditWorkspace) importAudit(heldRoot, changePath string) error {
	seen := make(map[string]bool)
	err := filepath.WalkDir(w.root, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if path == filepath.Join(w.root, ".git") {
			return filepath.SkipDir
		}
		if path == w.root || entry.IsDir() {
			return nil
		}
		rel, err := filepath.Rel(w.root, path)
		if err != nil {
			return err
		}
		rel = filepath.ToSlash(rel)
		original, allowed := w.original[rel]
		if !allowed || !entry.Type().IsRegular() {
			return fmt.Errorf("unexpected audit workspace file %s", rel)
		}
		seen[rel] = true
		if rel == changePath {
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		if !bytes.Equal(data, original) {
			return fmt.Errorf("audit workspace modified %s", rel)
		}
		return nil
	})
	if err != nil {
		return err
	}
	var missing []string
	for rel := range w.original {
		if !seen[rel] {
			missing = append(missing, rel)
		}
	}
	if len(missing) != 0 {
		sort.Strings(missing)
		return fmt.Errorf("audit workspace deleted %s", strings.Join(missing, ", "))
	}
	candidate, err := os.ReadFile(filepath.Join(w.root, filepath.FromSlash(changePath)))
	if err != nil {
		return err
	}
	original := string(w.original[changePath])
	audit, ok := contracts.Section(string(candidate), "Audit")
	if !ok {
		if string(candidate) != original {
			return fmt.Errorf("audit workspace modified change outside Audit")
		}
		return nil
	}
	if _, existed := contracts.Section(original, "Audit"); existed {
		origAudit, _ := contracts.Section(original, "Audit")
		if contracts.ReplaceSection(string(candidate), "Audit", origAudit) != original {
			return fmt.Errorf("audit workspace modified change outside Audit")
		}
	} else {
		heading := "## Audit\n"
		idx := strings.Index(string(candidate), heading)
		if idx < 0 || string(candidate[:idx]) != original && string(candidate[:idx]) != original+"\n" || idx+len(heading)+len(audit) != len(candidate) {
			return fmt.Errorf("audit workspace modified change outside Audit")
		}
	}
	heldPath := filepath.Join(heldRoot, filepath.FromSlash(changePath))
	held, err := os.ReadFile(heldPath)
	if err != nil {
		return err
	}
	updated := contracts.ReplaceSection(string(held), "Audit", audit)
	if updated == string(held) {
		return nil
	}
	return os.WriteFile(heldPath, []byte(updated), 0644)
}
