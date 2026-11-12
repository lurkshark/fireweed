package lifecycle

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
)

// removeDisposable uses the approved projection captured before Start changes
// the worktree. Only files owned by the target module are candidates.
func removeDisposable(root string, snapshot project.Snapshot, name string) error {
	m := module(snapshot, name)
	if m.Dir == "" {
		return fmt.Errorf("regenerate: module %q is missing", name)
	}
	protected := []string{".git", ".git/**", ".fireweed", ".fireweed/**", "SYSTEM.md", "fireweed.yaml"}
	protected = append(protected, list(snapshot.Config["protected"])...)
	var disposable []string
	err := filepath.WalkDir(file(root, m.Dir), func(full string, entry os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		rel, err := filepath.Rel(root, full)
		if err != nil {
			return err
		}
		rel = filepath.ToSlash(rel)
		for _, other := range snapshot.Modules {
			if other.Name != name && strings.HasPrefix(other.Dir, m.Dir+"/") && (rel == other.Dir || strings.HasPrefix(rel, other.Dir+"/")) {
				if entry.IsDir() {
					return filepath.SkipDir
				}
				return nil
			}
		}
		if path.Base(rel) == ".git" || path.Base(rel) == ".fireweed" {
			if entry.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}
		if entry.IsDir() {
			return nil
		}
		if path.Base(rel) == "AGENTS.md" || path.Base(rel) == "CLAUDE.md" {
			return nil
		}
		for _, pattern := range protected {
			if contracts.MatchGlob(pattern, rel) {
				return nil
			}
		}
		moduleRel := strings.TrimPrefix(rel, m.Dir+"/")
		if moduleRel != rel && contracts.ResolveZone(moduleRel, m.Zones) == "disposable" {
			disposable = append(disposable, full)
		}
		return nil
	})
	if err != nil {
		return err
	}
	for _, name := range disposable {
		if err := os.Remove(name); err != nil {
			return err
		}
	}
	return nil
}

func documentDate(value any) time.Time {
	switch date := value.(type) {
	case time.Time:
		return date
	case string:
		for _, layout := range []string{time.RFC3339Nano, "2006-01-02"} {
			if parsed, err := time.Parse(layout, date); err == nil {
				return parsed
			}
		}
	}
	return time.Time{}
}

func latestImplementation(root, name string) time.Time {
	entries, err := os.ReadDir(file(root, "changes/archive"))
	if err != nil {
		return time.Time{}
	}
	var latest time.Time
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".md") {
			continue
		}
		data, err := read(root, "changes/archive/"+entry.Name())
		if err != nil {
			continue
		}
		doc, err := contracts.ParseDocument(string(data))
		if err != nil || str(doc.Frontmatter["module"]) != name || str(doc.Frontmatter["phase"]) != "applied" {
			continue
		}
		provenance, _ := contracts.Section(doc.Body, "Provenance")
		for _, line := range strings.Split(provenance, "\n") {
			fields := strings.Fields(line)
			if len(fields) < 3 || fields[0] != "-" || fields[1] != "impl" {
				continue
			}
			date := documentDate(fields[2])
			if date.After(latest) {
				latest = date
			}
		}
	}
	return latest
}

func recentPassingDrill(root, name string) bool {
	data, err := read(root, ".fireweed/drills/"+name+".md")
	if err != nil {
		return false
	}
	doc, err := contracts.ParseDocument(string(data))
	if err != nil || str(doc.Frontmatter["result"]) != "pass" {
		return false
	}
	date := documentDate(doc.Frontmatter["date"])
	return !date.IsZero() && date.After(latestImplementation(root, name))
}
