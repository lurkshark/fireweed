package lifecycle

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project"
)

const (
	claimsPrefix   = "refs/fireweed/claims/"
	attemptsPrefix = "refs/fireweed/attempts/"
	drillsPrefix   = "refs/fireweed/drills/"
	runsDir        = ".fireweed/runs"
	excludeRuns    = ":(exclude,top).fireweed/runs"
)

func coded(code, subject string, err error) error {
	return contracts.CodedError{Code: code, Subject: subject, Cause: err}
}

func failure(code, subject, format string, args ...any) error {
	return coded(code, subject, fmt.Errorf(format, args...))
}

func diagnosticError(d project.Diagnostic) error {
	return failure(d.Code, d.Path, "%s", d.Message)
}

func warning(code, path, format string, args ...any) project.Diagnostic {
	return project.Diagnostic{Code: code, Path: path, Message: fmt.Sprintf(format, args...), Warning: true}
}

func gitInput(root string, input []byte, env map[string]string, args ...string) (string, error) {
	if env == nil {
		env = make(map[string]string)
	}
	env["GIT_NO_LAZY_FETCH"] = "1"
	env["GIT_TERMINAL_PROMPT"] = "0"
	result, err := gitstore.Run(context.Background(), gitstore.Request{
		Dir: root, Args: append([]string{"git", "-c", "core.hooksPath=/dev/null"}, args...),
		Input: input, Env: env, Timeout: 30 * time.Second,
	})
	if err != nil {
		return result.Stdout, fmt.Errorf("git %s: %w: %s", args[0], err, strings.TrimSpace(result.Stderr))
	}
	if result.ExitCode != 0 {
		return result.Stdout, fmt.Errorf("git %s: %s", args[0], strings.TrimSpace(result.Stderr))
	}
	return result.Stdout, nil
}

func git(root string, args ...string) (string, error) { return gitInput(root, nil, nil, args...) }

func gitFields(out string) []string {
	return strings.Split(strings.TrimSuffix(out, "\x00"), "\x00")
}

func refs(root, prefix string) ([]ref, error) {
	out, err := git(root, "for-each-ref", "--format=%(refname) %(objectname)", prefix)
	if err != nil {
		return nil, err
	}
	var result []ref
	for _, line := range strings.Split(strings.TrimSpace(out), "\n") {
		if name, oid, ok := strings.Cut(line, " "); ok {
			result = append(result, ref{name, oid})
		}
	}
	return result, nil
}

type ref struct{ name, oid string }

func disk(root, name string) string { return filepath.Join(root, filepath.FromSlash(name)) }

func readDocument(root, file string) (contracts.Document, error) {
	data, err := os.ReadFile(disk(root, file))
	if err != nil {
		return contracts.Document{}, err
	}
	return contracts.ParseDocument(string(data))
}

func writeDocument(root, file string, doc contracts.Document) error {
	content, err := contracts.RenderDocument(doc)
	if err != nil {
		return err
	}
	return writeFile(disk(root, file), []byte(content))
}

func writeFile(file string, content []byte) error {
	if err := os.MkdirAll(filepath.Dir(file), 0755); err != nil {
		return err
	}
	return os.WriteFile(file, content, 0644)
}

func section(doc contracts.Document, name string) string {
	body, _ := contracts.Section(doc.Body, name)
	return body
}

func items(doc contracts.Document, name string) ([]contracts.Item, error) {
	return contracts.ParseItems(section(doc, name))
}

func str(value any) string             { s, _ := value.(string); return s }
func mapping(value any) map[string]any { m, _ := value.(map[string]any); return m }

func integer(value any, fallback int) int {
	if n, ok := value.(int); ok {
		return n
	}
	return fallback
}

func stringsOf(value any) []string {
	if list, ok := value.([]string); ok {
		return list
	}
	list, _ := value.([]any)
	var result []string
	for _, item := range list {
		if s, ok := item.(string); ok {
			result = append(result, s)
		}
	}
	return result
}

func matches(value any, file string) bool {
	for _, pattern := range stringsOf(value) {
		if contracts.MatchGlob(pattern, file) {
			return true
		}
	}
	return false
}

func moduleNamed(snapshot project.Snapshot, name string) project.Module {
	for _, module := range snapshot.Modules {
		if module.Name == name {
			return module
		}
	}
	return project.Module{}
}

func ownerOf(snapshot project.Snapshot, file string) project.Module {
	var owner project.Module
	for _, module := range snapshot.Modules {
		if (module.Dir == "." || strings.HasPrefix(file, module.Dir+"/")) && len(module.Dir) > len(owner.Dir) {
			owner = module
		}
	}
	return owner
}

func fileZone(snapshot project.Snapshot, module project.Module, file string) string {
	if file == "SYSTEM.md" || file == "fireweed.yaml" || path.Base(file) == "AGENTS.md" ||
		path.Base(file) == "CLAUDE.md" || strings.HasPrefix(file, ".fireweed/") ||
		matches(snapshot.Config["protected"], file) {
		return "protected"
	}
	owner := ownerOf(snapshot, file)
	if owner.Name == "" && matches(snapshot.Config["worker_writable"], file) {
		return "worker_writable"
	}
	if owner.Name != module.Name || owner.Name == "" {
		return ""
	}
	relative := file
	if module.Dir != "." {
		relative = strings.TrimPrefix(file, module.Dir+"/")
	}
	return contracts.ResolveZone(relative, module.Zones)
}

func dependsOn(snapshot project.Snapshot, module, dependency string, transitive bool) bool {
	seen := make(map[string]bool)
	var search func(string) bool
	search = func(name string) bool {
		if seen[name] {
			return false
		}
		seen[name] = true
		for _, direct := range moduleNamed(snapshot, name).DependsOn {
			if direct == dependency || transitive && search(direct) {
				return true
			}
		}
		return false
	}
	return search(module)
}

const pendingDir = "changes"
const archiveDir = "changes/archive"

var changeName = regexp.MustCompile(`^([0-9]{4,})-([a-z0-9-]+)\.md$`)

func changeIdentity(file string) project.Change {
	if match := changeName.FindStringSubmatch(path.Base(file)); match != nil {
		if number, err := strconv.Atoi(match[1]); err == nil {
			return project.Change{Path: file, Slug: match[2], Number: number}
		}
	}
	return project.Change{}
}

func changePaths(root, revision, dir string) ([]project.Change, error) {
	var files []string
	if revision == "" {
		entries, err := os.ReadDir(disk(root, dir))
		if errors.Is(err, os.ErrNotExist) {
			return nil, nil
		}
		if err != nil {
			return nil, err
		}
		for _, entry := range entries {
			if !entry.IsDir() {
				files = append(files, path.Join(dir, entry.Name()))
			}
		}
	} else {
		out, err := git(root, "ls-tree", "-r", "--name-only", "-z", revision, "--", dir+"/")
		if err != nil {
			return nil, err
		}
		files = gitFields(out)
	}
	var result []project.Change
	for _, file := range files {
		if path.Dir(file) == dir {
			if change := changeIdentity(file); change.Slug != "" {
				result = append(result, change)
			}
		}
	}
	return result, nil
}

func selectChange(changes []project.Change, selector string) (project.Change, error) {
	var selected []project.Change
	var slugs []string
	for _, change := range changes {
		if !slices.Contains(slugs, change.Slug) {
			slugs = append(slugs, change.Slug)
		}
		if selector == change.Path || selector == path.Base(change.Path) || selector == change.Slug {
			selected = append(selected, change)
		}
	}
	switch len(selected) {
	case 0:
		return project.Change{}, failure("U02", selector, "change %q not found (valid slugs: %s)", selector, strings.Join(slugs, ", "))
	case 1:
		return selected[0], nil
	default:
		return project.Change{}, failure("U02", selector, "ambiguous change %q (valid slugs: %s)", selector, strings.Join(slugs, ", "))
	}
}

func parseChange(change project.Change, doc contracts.Document) (project.Change, error) {
	change.Document = doc
	change.Module = str(doc.Frontmatter["module"])
	for _, part := range []struct {
		name string
		dest *[]contracts.Item
	}{{"Added", &change.Added}, {"Modified", &change.Modified}, {"Removed", &change.Removed}} {
		list, err := items(doc, part.name)
		if err != nil {
			return change, err
		}
		*part.dest = list
	}
	return change, nil
}

func readChange(root, revision string, change project.Change) (project.Change, error) {
	var doc contracts.Document
	var err error
	if revision == "" {
		doc, err = readDocument(root, change.Path)
	} else {
		var data []byte
		data, err = gitstore.ReadAt(root, revision, change.Path)
		if err == nil {
			doc, err = contracts.ParseDocument(string(data))
		}
	}
	if err != nil {
		return change, err
	}
	return parseChange(change, doc)
}

func archivePath(change project.Change) string { return path.Join(archiveDir, path.Base(change.Path)) }
func pendingPath(change project.Change) string { return path.Join(pendingDir, path.Base(change.Path)) }
func numberedName(change project.Change) string {
	return strings.TrimSuffix(path.Base(change.Path), ".md")
}

func clean(root, operation string) error {
	out, err := git(root, "--no-optional-locks", "status", "--porcelain=v1", "--untracked-files=all", "-z")
	if err != nil {
		return err
	}
	if out != "" {
		return failure("P01", "", "%s requires clean worktree", operation)
	}
	return nil
}

func isRunArtifact(file string) bool { return strings.HasPrefix(file, runsDir+"/") }

func staticFailure(diagnostics []project.Diagnostic) error {
	for _, d := range diagnostics {
		if !d.Warning {
			err := diagnosticError(d)
			if d.Code == "C08" && strings.Contains(d.Message, "stale base for ") {
				return coded("P05", d.Path, err)
			}
			return err
		}
	}
	return nil
}

func tempIndex(root, base string, use func(map[string]string) error) error {
	dir, err := os.MkdirTemp("", "fireweed-index-")
	if err != nil {
		return err
	}
	defer os.RemoveAll(dir)
	env := map[string]string{"GIT_INDEX_FILE": filepath.Join(dir, "index")}
	if _, err := gitInput(root, nil, env, "read-tree", base); err != nil {
		return err
	}
	return use(env)
}

func baseRefs(snapshot project.Snapshot) (local, shared string, err error) {
	shared, err = gitstore.BaseRef(snapshot.Root, str(snapshot.Config["remote"]), str(snapshot.Config["base"]))
	if err != nil {
		return "", "", err
	}
	if tail, ok := strings.CutPrefix(shared, "refs/remotes/"); ok {
		_, branch, _ := strings.Cut(tail, "/")
		return "refs/heads/" + branch, shared, nil
	}
	return shared, shared, nil
}

func runPath(root, run, phase string) string { return disk(root, path.Join(runsDir, run, phase)) }

func provenanceDate(doc contracts.Document) time.Time {
	var latest time.Time
	for _, line := range strings.Split(section(doc, "Provenance"), "\n") {
		fields := strings.Fields(line)
		if len(fields) >= 3 && fields[0] == "-" && fields[1] == "impl" {
			date, err := time.Parse(time.RFC3339Nano, fields[2])
			if err == nil && date.After(latest) {
				latest = date
			}
		}
	}
	return latest
}
