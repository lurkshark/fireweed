package inspection

import (
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"

	"github.com/lurkshark/fireweed/packages/assets"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project/internal/changeview"
)

type Diagnostic struct{ Code, Severity, Message, Path string }
type Data struct {
	Root, BaseRef    string
	Config           contracts.Config
	Modules          []contracts.Module
	Changes, Archive []contracts.Change
	Diagnostics      []Diagnostic
}

func Load(start string) (Data, error) {
	d := Data{Modules: []contracts.Module{}, Changes: []contracts.Change{}, Archive: []contracts.Change{}, Diagnostics: []Diagnostic{}}
	repo, err := gitstore.Open(start)
	if err != nil {
		return d, err
	}
	d.Root = repo.Root
	text, err := read(d.Root, "fireweed.yaml")
	if err != nil {
		return d, fmt.Errorf("C01: read fireweed.yaml: %w; restore a valid project configuration", err)
	}
	d.Config, err = contracts.ParseConfig(text, assets.Info().Version)
	if err != nil {
		return d, fmt.Errorf("C01: fix fireweed.yaml or use a compatible binary: %w", err)
	}
	patterns := append([]string{}, d.Config.Modules...)
	patterns = append(patterns, zonePatterns(d.Config.ModuleDefaults)...)
	patterns = append(patterns, d.Config.Protected...)
	patterns = append(patterns, d.Config.WorkerWritable...)
	for _, pattern := range patterns {
		if _, err = contracts.MatchGlob(pattern, "future-path"); err != nil {
			return d, fmt.Errorf("C01: fix invalid configured glob %q: %w", pattern, err)
		}
	}
	d.BaseRef, err = gitstore.BaseRef(d.Root, d.Config.Base, d.Config.Remote)
	if err != nil {
		return d, fmt.Errorf("C01: resolve configured local base: %w", err)
	}
	err = filepath.WalkDir(d.Root, func(full string, entry fs.DirEntry, walkErr error) error {
		relative, relErr := filepath.Rel(d.Root, full)
		if relErr != nil {
			return relErr
		}
		relative = filepath.ToSlash(relative)
		if walkErr != nil {
			d.add("C03", "error", relative, "Cannot inspect path; correct its permissions: "+walkErr.Error())
			return nil
		}
		if entry.IsDir() {
			if relative != "." && (entry.Name() == ".git" || entry.Name() == ".fireweed") {
				return filepath.SkipDir
			}
			return nil
		}
		if entry.Name() != "SPEC.md" {
			return nil
		}
		dir := path.Dir(relative)
		matched := false
		for _, pattern := range d.Config.Modules {
			ok, _ := contracts.MatchGlob(pattern, dir)
			if ok {
				matched = true
				break
			}
		}
		if !matched {
			d.add("C02", "warning", relative, "Move SPEC.md under a configured modules glob or update fireweed.yaml.")
			return nil
		}
		text, err := read(d.Root, relative)
		if err != nil {
			d.add("C03", "error", relative, "Read the module through a safe repository path; fix its symlink or permissions: "+err.Error())
			return nil
		}
		m, code, err := parseModule(text, dir, d.Config)
		if err != nil {
			d.add(code, "error", relative, "Correct the module document: "+err.Error())
			return nil
		}
		d.Modules = append(d.Modules, m)
		return nil
	})
	if err != nil {
		return d, err
	}
	sort.Slice(d.Modules, func(i, j int) bool {
		a, b := d.Modules[i], d.Modules[j]
		if a.Name != b.Name {
			return a.Name < b.Name
		}
		return a.Dir < b.Dir
	})
	prefixes := map[string]string{}
	for _, m := range d.Modules {
		prefixes[m.Name] = m.IDPrefix
	}
	d.Changes = d.loadChanges("changes", prefixes)
	d.Archive = d.loadChanges("changes/archive", prefixes)
	seen := map[string]string{}
	for _, list := range [][]contracts.Change{d.Changes, d.Archive} {
		for _, ch := range list {
			if earlier, ok := seen[ch.Slug]; ok {
				d.add("C05", "error", ch.Path, fmt.Sprintf("Slug %q also appears in %s; choose unique change slugs.", ch.Slug, earlier))
			} else {
				seen[ch.Slug] = ch.Path
			}
		}
	}
	sortDiagnostics(d.Diagnostics)
	return d, nil
}
func read(root, relative string) (string, error) {
	full, err := gitstore.SafePath(root, relative)
	if err != nil {
		return "", err
	}
	bytes, err := os.ReadFile(full)
	return string(bytes), err
}
func (d *Data) add(code, severity, path, message string) {
	d.Diagnostics = append(d.Diagnostics, Diagnostic{Code: code, Severity: severity, Path: path, Message: message})
}
func zonePatterns(z contracts.Zones) []string {
	out := []string{}
	for _, group := range [][]string{z.Interface, z.Evals, z.Disposable, z.Preserve, z.AppendOnly} {
		out = append(out, group...)
	}
	return out
}

// Validate metadata without letting a malformed normative body obscure its category.
// The original document is retained, never replaced by this temporary view.
func parseModule(text, dir string, cfg contracts.Config) (contracts.Module, string, error) {
	doc, err := contracts.ParseMarkdown(text)
	if err != nil {
		code := "C03"
		if strings.Contains(err.Error(), "Specs") {
			code = "C04"
		}
		return contracts.Module{}, code, err
	}
	metadataDoc := doc
	metadataDoc.Sections = append([]contracts.Section{}, doc.Sections...)
	found := false
	var body string
	for i, s := range metadataDoc.Sections {
		if s.Name == "Specs" {
			found = true
			body = s.Body
			metadataDoc.Sections[i].Body = ""
		}
	}
	if !found {
		metadataDoc.Sections = append(metadataDoc.Sections, contracts.Section{Name: "Specs"})
	}
	rendered, err := contracts.RenderMarkdown(metadataDoc)
	if err != nil {
		return contracts.Module{}, "C03", err
	}
	m, err := contracts.ParseModule(rendered, dir, cfg.ModuleDefaults, cfg.Hooks)
	if err != nil {
		return m, "C03", err
	}
	if !found {
		return m, "C04", fmt.Errorf("exactly one Specs section is required")
	}
	m.Specs, err = contracts.ParseItems(body, m.IDPrefix, false)
	if err != nil {
		return m, "C04", err
	}
	m.Document = doc
	return m, "", nil
}
func (d *Data) loadChanges(dir string, prefixes map[string]string) []contracts.Change {
	out := []contracts.Change{}
	full, err := gitstore.SafePath(d.Root, dir)
	if err != nil {
		d.add("C05", "error", dir, "Fix the change directory symlink: "+err.Error())
		return out
	}
	entries, err := os.ReadDir(full)
	if os.IsNotExist(err) {
		return out
	}
	if err != nil {
		d.add("C05", "error", dir, "Make the change directory readable: "+err.Error())
		return out
	}
	for _, entry := range entries {
		if entry.IsDir() || path.Ext(entry.Name()) != ".md" {
			continue
		}
		relative := dir + "/" + entry.Name()
		text, err := read(d.Root, relative)
		if err != nil {
			d.add("C05", "error", relative, "Read changes through safe repository paths; fix the symlink or permissions: "+err.Error())
			continue
		}
		doc, err := contracts.ParseMarkdown(text)
		if err != nil {
			d.add("C05", "error", relative, "Correct the change document: "+err.Error())
			continue
		}
		target, _ := doc.Metadata["module"].(string)
		prefix, ok := prefixes[target]
		if !ok {
			d.add("C05", "error", relative, fmt.Sprintf("Target module %q is unknown; select an existing valid module.", target))
			continue
		}
		validationText := text
		if dir == "changes/archive" {
			validationText, err = changeview.ArchiveText(doc)
			if err != nil {
				d.add("C05", "error", relative, "Correct the archived change document: "+err.Error())
				continue
			}
		}
		ch, err := contracts.ParseChange(validationText, relative, prefix)
		if err != nil {
			d.add(changeview.ErrorCode(err), "error", relative, "Correct the change document for its target module; declare each live handle locally: "+err.Error())
			continue
		}
		ch.Document = doc
		out = append(out, ch)
	}
	sort.Slice(out, func(i, j int) bool {
		a, b := out[i], out[j]
		if a.Number != b.Number {
			return a.Number < b.Number
		}
		return a.Slug < b.Slug
	})
	return out
}

func Check(s Data) []Diagnostic {
	ds := append([]Diagnostic{}, s.Diagnostics...)
	add := func(code, severity, where, message string) {
		ds = append(ds, Diagnostic{Code: code, Severity: severity, Path: where, Message: message})
	}
	specPath := func(m contracts.Module) string { return path.Join(m.Dir, "SPEC.md") }
	names, prefixes := map[string]int{}, map[string]int{}
	for i, m := range s.Modules {
		where := specPath(m)
		if previous, ok := names[m.Name]; ok {
			add("C03", "error", where, fmt.Sprintf("Module name %q also occurs in %s; use unique names.", m.Name, specPath(s.Modules[previous])))
		} else {
			names[m.Name] = i
		}
		if previous, ok := prefixes[m.IDPrefix]; ok {
			add("C03", "error", where, fmt.Sprintf("ID prefix %q also occurs in %s; use unique prefixes.", m.IDPrefix, specPath(s.Modules[previous])))
		} else {
			prefixes[m.IDPrefix] = i
		}
		for j := 0; j < i; j++ {
			if overlapsDir(m.Dir, s.Modules[j].Dir) {
				add("C06", "error", where, fmt.Sprintf("Module directory overlaps %s; give each path one module owner.", s.Modules[j].Dir))
			}
		}
		for _, pattern := range zonePatterns(m.Zones) {
			if _, err := contracts.MatchGlob(pattern, "future-path"); err != nil {
				add("C03", "error", where, fmt.Sprintf("Correct invalid relative zone pattern %q: %v", pattern, err))
			}
		}
		for _, e := range m.Evals {
			for _, d := range m.Disposable {
				overlap, err := contracts.GlobsOverlap(e, d)
				if err == nil && overlap {
					add("C06", "error", where, fmt.Sprintf("Evals pattern %q intersects disposable %q; separate these zones.", e, d))
				}
			}
		}
		if strings.TrimSpace(m.Hooks.Eval) == "" {
			add("C07", "error", where, "Configure a nonblank eval hook for this module.")
		}
		reason := strings.TrimSpace(m.BoundariesOptOut)
		if reason != "" {
			add("D04", "warning", where, "Review explicit boundary opt-out: "+m.BoundariesOptOut)
		} else if strings.TrimSpace(m.Hooks.Boundaries) == "" {
			if len(m.Disposable) > 0 {
				add("C07", "error", where, "Disposable modules require a boundaries hook or an explicit reasoned boundaries_opt_out.")
			} else {
				add("D04", "warning", where, "Boundary enforcement is absent; configure a boundaries hook before making code disposable.")
			}
		}
	}
	for _, m := range s.Modules {
		for _, dep := range m.DependsOn {
			if dep == m.Name {
				add("C06", "error", specPath(m), "Remove the module's dependency on itself.")
			} else if _, ok := names[dep]; !ok {
				add("C06", "error", specPath(m), fmt.Sprintf("Dependency %q is missing; declare an existing module.", dep))
			}
		}
	}
	// Three colors distinguish a back edge from a shared already-completed dependency.
	colors := map[string]int{}
	var visit func(string)
	visit = func(name string) {
		colors[name] = 1
		m := s.Modules[names[name]]
		for _, dep := range m.DependsOn {
			if _, ok := names[dep]; !ok || dep == name {
				continue
			}
			if colors[dep] == 1 {
				add("C06", "error", specPath(m), fmt.Sprintf("Dependency on %q closes a cycle; make the module graph acyclic.", dep))
			} else if colors[dep] == 0 {
				visit(dep)
			}
		}
		colors[name] = 2
	}
	for _, m := range s.Modules {
		if colors[m.Name] == 0 {
			visit(m.Name)
		}
	}
	removed := map[string]string{}
	for _, ch := range s.Archive {
		for _, item := range ch.Removed {
			removed[item.ID] = ch.Path
		}
	}
	for _, m := range s.Modules {
		for _, item := range m.Specs {
			if archive, ok := removed[item.ID]; ok {
				add("C08", "error", specPath(m), fmt.Sprintf("ID %s was removed in %s; allocate a new ID and retain the reservation.", item.ID, archive))
			}
		}
	}
	sortDiagnostics(ds)
	return ds
}
func overlapsDir(a, b string) bool {
	return a == "." || b == "." || a == b || strings.HasPrefix(a, b+"/") || strings.HasPrefix(b, a+"/")
}
func sortDiagnostics(ds []Diagnostic) {
	sort.SliceStable(ds, func(i, j int) bool {
		a, b := ds[i], ds[j]
		if a.Path != b.Path {
			return a.Path < b.Path
		}
		if a.Code != b.Code {
			return a.Code < b.Code
		}
		if a.Severity != b.Severity {
			return a.Severity < b.Severity
		}
		return a.Message < b.Message
	})
}
