package engine

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/lurkshark/fireweed/packages/assets"
	"github.com/lurkshark/fireweed/packages/checks/model"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project"
	"gopkg.in/yaml.v3"
)

const doctorReply = "FIREWEED_DOCTOR_OK"

var canaryCaseToken = regexp.MustCompile(`Test_[A-Za-z0-9_]+|[A-Za-z][A-Za-z0-9]*-[0-9]+`)

const doctorBlock = "<!-- fireweed:start -->\nThis project uses fireweed. Specs and change files are the source of truth.\nRun `fireweed playbook author-change` before authoring changes and\n`fireweed context <change>` before each step.\n<!-- fireweed:end -->"

type doctorCanaries struct {
	Version      int               `yaml:"version"`
	Baseline     map[string]string `yaml:"baseline"`
	Undeclared   map[string]string `yaml:"undeclared"`
	InternalEval map[string]string `yaml:"internal_eval"`
	FailingEval  map[string]string `yaml:"failing_eval"`
	MissingID    map[string]string `yaml:"missing_id"`
}

type doctorState struct {
	ctx      context.Context
	req      model.DoctorRequest
	report   model.DoctorReport
	root     string
	snapshot project.Snapshot
	canaries doctorCanaries
	names    [2]string
	dirs     [2]string
	prefixes [2]string
	goModule string
}

func (d *doctorState) issue(code, msg, path string, warning bool) {
	d.report.Diagnostics = append(d.report.Diagnostics, project.Diagnostic{Code: code, Message: msg, Path: path, Warning: warning})
}
func (d *doctorState) check(name, module string, ok bool, detail string) {
	d.report.Checks = append(d.report.Checks, model.DoctorCheck{Name: name, Module: module, Passed: ok, Detail: detail})
}
func (d *doctorState) fail(name, module, hook, detail string) {
	d.check(name, module, false, detail)
	if hook == "boundaries" {
		hook = "boundary hook"
	}
	d.issue("C14", name+" "+hook+": "+detail, "", false)
}

func Doctor(ctx context.Context, request model.DoctorRequest) model.DoctorReport {
	if ctx == nil {
		ctx = context.Background()
	}
	d := &doctorState{ctx: ctx, req: request, root: request.Root}
	d.report.Diagnostics = append(d.report.Diagnostics, project.Check(request.Root)...)
	for _, v := range d.report.Diagnostics {
		if !v.Warning {
			return d.report
		}
	}
	d.guidance()
	repo, err := gitstore.Inspect(request.Root)
	if err != nil {
		d.issue("C14", "scratch worktree: "+err.Error(), "", false)
		return d.report
	}
	snapshot, err := project.Load(request.Root)
	if err != nil {
		d.issue("C14", "project: "+err.Error(), "", false)
		return d.report
	}
	d.snapshot = snapshot
	d.boundaryWarnings(snapshot)
	if err = d.readRecipe(request.Root); err != nil {
		d.issue("C14", err.Error(), "", false)
		return d.report
	}
	base := filepath.Join(repo.CommonDir, "fireweed", "worktrees")
	if err = os.MkdirAll(base, 0755); err != nil {
		d.issue("C14", "scratch worktree: "+err.Error(), "", false)
		return d.report
	}
	id, err := randomHex(8)
	if err != nil {
		d.issue("C14", "scratch worktree: "+err.Error(), "", false)
		return d.report
	}
	scratch := filepath.Join(base, "doctor-"+id)
	if _, err = d.git(request.Root, "worktree", "add", "--detach", scratch, repo.Head); err != nil {
		d.issue("C14", "scratch worktree: "+err.Error(), "", false)
		return d.report
	}
	defer func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		if _, e := d.gitWithContext(cleanupCtx, request.Root, "worktree", "remove", "--force", scratch); e != nil {
			d.issue("C14", "scratch cleanup: "+e.Error(), "", false)
		}
		_ = os.RemoveAll(scratch)
	}()
	if err = d.copyCurrent(request.Root, scratch); err != nil {
		d.issue("C14", "scratch copy: "+err.Error(), "", false)
		return d.report
	}
	d.root = scratch
	setup := mapValue(snapshot.Config, "hooks")["setup"]
	if setup != nil {
		p, e := runHook(ctx, scratch, setup, nil, request.Bin, request.Timeout)
		if e != nil || p.ExitCode != 0 {
			d.fail("setup", "", "setup", fmt.Sprintf("hook %v: %v; exit %d; %s %s", setup, e, p.ExitCode, p.Stdout, p.Stderr))
			return d.report
		}
	}
	if err = d.renderModules(); err != nil {
		d.issue("C14", err.Error(), "", false)
		return d.report
	}
	d.probeHooks()
	if !request.NoAgent {
		d.probeAgents(repo.CommonDir)
	}
	return d.report
}

func randomHex(n int) (string, error) {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}
func (d *doctorState) git(root string, args ...string) (gitstore.Result, error) {
	return d.gitWithContext(d.ctx, root, args...)
}
func (d *doctorState) gitWithContext(ctx context.Context, root string, args ...string) (gitstore.Result, error) {
	argv := append([]string{"git", "-C", root, "-c", "core.hooksPath=/dev/null"}, args...)
	p, e := gitstore.Run(ctx, gitstore.Request{Args: argv, Timeout: 30 * time.Second, Env: map[string]string{"GIT_NO_LAZY_FETCH": "1", "GIT_TERMINAL_PROMPT": "0"}})
	if e != nil {
		return p, fmt.Errorf("git %s: %w: %s", strings.Join(args, " "), e, p.Stderr)
	}
	return p, nil
}
func (d *doctorState) copyCurrent(src, dst string) error {
	p, err := d.git(src, "ls-files", "-co", "--exclude-standard", "-z")
	if err != nil {
		return err
	}
	seen := map[string]bool{}
	for _, rel := range strings.Split(p.Stdout, "\x00") {
		if rel == "" || seen[rel] {
			continue
		}
		seen[rel] = true
		from := filepath.Join(src, filepath.FromSlash(rel))
		to := filepath.Join(dst, filepath.FromSlash(rel))
		info, e := os.Lstat(from)
		if os.IsNotExist(e) {
			_ = os.RemoveAll(to)
			continue
		}
		if e != nil {
			return e
		}
		if e = os.MkdirAll(filepath.Dir(to), 0755); e != nil {
			return e
		}
		_ = os.RemoveAll(to)
		if info.Mode()&os.ModeSymlink != 0 {
			target, e := os.Readlink(from)
			if e != nil {
				return e
			}
			if e = os.Symlink(target, to); e != nil {
				return e
			}
			continue
		}
		if !info.Mode().IsRegular() {
			continue
		}
		in, e := os.Open(from)
		if e != nil {
			return e
		}
		out, e := os.OpenFile(to, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, info.Mode().Perm())
		if e != nil {
			in.Close()
			return e
		}
		_, e = io.Copy(out, in)
		ce := out.Close()
		in.Close()
		if e != nil {
			return e
		}
		if ce != nil {
			return ce
		}
	}
	return nil
}
func (d *doctorState) readRecipe(root string) error {
	templateDir := filepath.Join(root, ".fireweed", "templates", "module")
	entries, err := os.ReadDir(templateDir)
	if err != nil {
		return fmt.Errorf("module template %s: %w", templateDir, err)
	}
	hasSpec := false
	for _, entry := range entries {
		if entry.Name() == "SPEC.md" {
			hasSpec = true
		}
		if entry.IsDir() {
			return fmt.Errorf("module template %s has directory", entry.Name())
		}
	}
	if !hasSpec {
		return fmt.Errorf("module template SPEC.md missing")
	}
	data, err := os.ReadFile(filepath.Join(root, ".fireweed", "canaries", "canaries.yaml"))
	if err != nil {
		return fmt.Errorf("canary canaries.yaml: %w", err)
	}
	if err = yaml.Unmarshal(data, &d.canaries); err != nil {
		return fmt.Errorf("canary canaries.yaml: %w", err)
	}
	if d.canaries.Version != 1 {
		return fmt.Errorf("canary canaries.yaml version must be 1")
	}
	for _, v := range []struct {
		name  string
		files map[string]string
	}{{"baseline", d.canaries.Baseline}, {"undeclared", d.canaries.Undeclared}, {"internal_eval", d.canaries.InternalEval}, {"failing_eval", d.canaries.FailingEval}, {"missing_id", d.canaries.MissingID}} {
		if len(v.files) == 0 {
			return fmt.Errorf("canary %s map missing or empty in canaries.yaml", v.name)
		}
		for name := range v.files {
			if !safeCanaryPath(name) {
				return fmt.Errorf("canary %s path %q invalid", v.name, name)
			}
		}
	}
	data, err = os.ReadFile(filepath.Join(root, "go.mod"))
	if err == nil {
		for _, line := range strings.Split(string(data), "\n") {
			fields := strings.Fields(line)
			if len(fields) >= 2 && fields[0] == "module" {
				d.goModule = strings.Trim(fields[1], "\"")
				break
			}
		}
	}
	return nil
}
func safeCanaryPath(name string) bool {
	return name != "" && !filepath.IsAbs(name) && filepath.Clean(name) != ".." && !strings.HasPrefix(filepath.Clean(name), ".."+string(filepath.Separator))
}
func renderDoctorText(s string, v map[string]string) string {
	for k, x := range v {
		s = strings.ReplaceAll(s, "{{"+k+"}}", x)
	}
	return s
}
func (d *doctorState) writeMap(dir string, files map[string]string, values map[string]string) ([]string, error) {
	var paths []string
	keys := make([]string, 0, len(files))
	for k := range files {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, name := range keys {
		if !safeCanaryPath(name) {
			return nil, fmt.Errorf("invalid canary path %s", name)
		}
		path := filepath.Join(d.root, dir, filepath.FromSlash(strings.TrimSuffix(name, ".tmpl")))
		if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			return nil, err
		}
		if err := os.WriteFile(path, []byte(renderDoctorText(files[name], values)), 0644); err != nil {
			return nil, err
		}
		paths = append(paths, path)
	}
	return paths, nil
}
func (d *doctorState) renderModules() error {
	entries, _ := os.ReadDir(filepath.Join(d.root, ".fireweed", "templates", "module"))
	templates := map[string]string{}
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		b, err := os.ReadFile(filepath.Join(d.root, ".fireweed", "templates", "module", e.Name()))
		if err != nil {
			return fmt.Errorf("module template %s: %w", e.Name(), err)
		}
		templates[e.Name()] = string(b)
	}
	for i := 0; i < 2; i++ {
		id, err := randomHex(5)
		if err != nil {
			return err
		}
		d.names[i] = "doctor" + id
		d.prefixes[i] = "DOC" + strings.ToUpper(id)
		d.dirs[i] = filepath.ToSlash(filepath.Join("packages", d.names[i]))
	}
	for i := 0; i < 2; i++ {
		v := map[string]string{"name": d.names[i], "prefix": d.prefixes[i], "package": d.names[i], "go_module": d.goModule, "module_dir": d.dirs[i], "other_dir": d.dirs[1-i]}
		if _, err := d.writeMap(d.dirs[i], templates, v); err != nil {
			return fmt.Errorf("module template: %w", err)
		}
		original, err := os.ReadFile(filepath.Join(d.root, d.dirs[i], "SPEC.md"))
		if err != nil {
			return fmt.Errorf("module template SPEC.md: %w", err)
		}
		if _, err := contracts.ParseDocument(string(original)); err != nil {
			return fmt.Errorf("module template SPEC.md malformed: %w", err)
		}
		if _, err := d.writeMap(d.dirs[i], d.canaries.Baseline, v); err != nil {
			return fmt.Errorf("canary baseline: %w", err)
		}
		b, err := os.ReadFile(filepath.Join(d.root, d.dirs[i], "SPEC.md"))
		if err != nil {
			return fmt.Errorf("module template SPEC.md: %w", err)
		}
		doc, err := contracts.ParseDocument(string(b))
		if err != nil {
			return fmt.Errorf("module template SPEC.md malformed: %w", err)
		}
		if doc.Frontmatter["name"] != d.names[i] || doc.Frontmatter["id_prefix"] != d.prefixes[i] {
			return fmt.Errorf("module template SPEC.md malformed identity")
		}
	}
	// Explicit paths keep the generated graph visible even when the original pattern is narrow.
	config := d.snapshot.Config
	paths := make([]any, 0, len(d.snapshot.Modules)+2)
	for _, m := range d.snapshot.Modules {
		paths = append(paths, m.Dir)
	}
	for _, dir := range d.dirs {
		paths = append(paths, dir)
	}
	config["modules"] = paths
	b, err := yaml.Marshal(config)
	if err != nil {
		return err
	}
	if err = os.WriteFile(filepath.Join(d.root, "fireweed.yaml"), b, 0644); err != nil {
		return err
	}
	d.snapshot, err = project.Load(d.root)
	if err != nil {
		return fmt.Errorf("module template graph: %w", err)
	}
	return nil
}
func (d *doctorState) values(i int) map[string]string {
	return map[string]string{"name": d.names[i], "prefix": d.prefixes[i], "package": d.names[i], "go_module": d.goModule, "module_dir": d.dirs[i], "other_dir": d.dirs[1-i]}
}
func (d *doctorState) boundaryWarnings(s project.Snapshot) {
	global := mapValue(s.Config, "hooks")["boundaries"]
	configured := global != nil
	required := false
	for _, m := range s.Modules {
		if mapValue(m.Document.Frontmatter, "hooks")["boundaries"] != nil {
			configured = true
		}
		if reason := boundaryReason(m); reason != "" {
			d.issue("X03", m.Name+": boundaries opt-out: "+reason, "", true)
		} else if len(m.Zones.Disposable) > 0 {
			required = true
		}
	}
	if !configured && !required {
		d.issue("X04", "no boundary enforcement required; boundaries hook absent", "", true)
	}
}
func (d *doctorState) guidance() {
	b, err := os.ReadFile(filepath.Join(d.root, "AGENTS.md"))
	if err != nil || !strings.Contains(string(b), doctorBlock) {
		d.issue("C16", "AGENTS.md managed block is missing or stale", "AGENTS.md", true)
	}
	dir := filepath.Join(d.root, ".fireweed", "playbooks")
	entries, err := os.ReadDir(dir)
	if err != nil {
		return
	}
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".md") {
			continue
		}
		name := strings.TrimSuffix(e.Name(), ".md")
		base, err := assets.Playbook(name)
		if err != nil {
			continue
		}
		b, err := os.ReadFile(filepath.Join(dir, e.Name()))
		if err != nil {
			continue
		}
		doc, err := contracts.ParseDocument(string(b))
		sum := sha256.Sum256([]byte(base))
		want := hex.EncodeToString(sum[:])
		if err != nil || doc.Frontmatter["embedded_base"] != want {
			d.issue("C16", name+" playbook embedded_base is missing or stale", filepath.Join(".fireweed", "playbooks", e.Name()), true)
		}
	}
}

type doctorVariant struct {
	name    string
	command any
	module  int
	kind    string
}

func (d *doctorState) variants(kind string) []doctorVariant {
	var out []doctorVariant
	if command := mapValue(d.snapshot.Config, "hooks")[kind]; command != nil {
		out = append(out, doctorVariant{name: "global " + kind, command: command, module: 0, kind: kind})
	}
	for _, m := range d.snapshot.Modules {
		if m.Name == d.names[0] || m.Name == d.names[1] {
			continue
		}
		if command := mapValue(m.Document.Frontmatter, "hooks")[kind]; command != nil {
			out = append(out, doctorVariant{name: m.Name + " " + kind + " override", command: command, module: 0, kind: kind})
		}
	}
	return out
}
func (d *doctorState) variantSnapshot(v doctorVariant) project.Snapshot {
	s := d.snapshot
	s.Modules = append([]project.Module(nil), s.Modules...)
	s.Config = make(map[string]any, len(d.snapshot.Config))
	for k, x := range d.snapshot.Config {
		s.Config[k] = x
	}
	hooks := map[string]any{}
	for k, x := range mapValue(d.snapshot.Config, "hooks") {
		hooks[k] = x
	}
	hooks[v.kind] = nil
	if v.kind == "boundaries" {
		hooks[v.kind] = v.command
	} else if v.kind == "eval" {
		hooks[v.kind] = mapValue(d.snapshot.Config, "hooks")[v.kind]
		if hooks[v.kind] == nil {
			hooks[v.kind] = v.command
		}
	}
	s.Config["hooks"] = hooks
	for i := range s.Modules {
		if s.Modules[i].Name == d.names[v.module] {
			front := map[string]any{}
			for k, x := range s.Modules[i].Document.Frontmatter {
				front[k] = x
			}
			front["hooks"] = map[string]any{v.kind: v.command}
			s.Modules[i].Document.Frontmatter = front
			break
		}
	}
	// Other real overrides must not execute during this variant's evidence.
	for i := range s.Modules {
		if s.Modules[i].Name == d.names[0] || s.Modules[i].Name == d.names[1] {
			continue
		}
		front := map[string]any{}
		for k, x := range s.Modules[i].Document.Frontmatter {
			front[k] = x
		}
		delete(front, "hooks")
		s.Modules[i].Document.Frontmatter = front
	}
	return s
}
func (d *doctorState) evalClean(s project.Snapshot, module string) error {
	report, err := Evaluate(d.ctx, model.HookRequest{Snapshot: s, Module: module, Bin: d.req.Bin, Timeout: d.req.Timeout})
	if err != nil {
		return fmt.Errorf("eval: %w", err)
	}
	if report.HookExit != 0 {
		return fmt.Errorf("eval exit %d: %s %s", report.HookExit, report.Stdout, report.Stderr)
	}
	m := project.Module{}
	for _, x := range s.Modules {
		if x.Name == module {
			m = x
			break
		}
	}
	specs, err := moduleSpecs(m)
	if err != nil {
		return fmt.Errorf("eval specs: %w", err)
	}
	if err = RequireCoverage(report, specs); err != nil {
		return fmt.Errorf("eval coverage: %w", err)
	}
	if err = RequirePass(report, specs); err != nil {
		return fmt.Errorf("eval: %w", err)
	}
	return nil
}
func (d *doctorState) boundaryClean(s project.Snapshot) error {
	ds, err := Boundaries(d.ctx, model.BoundaryRequest{Snapshot: s, Bin: d.req.Bin, Timeout: d.req.Timeout})
	if err != nil {
		return fmt.Errorf("boundaries: %w: %v", err, ds)
	}
	return nil
}
func (d *doctorState) crossClean(canary string, evals, bounds []doctorVariant) error {
	if len(evals) > 0 {
		s := d.variantSnapshot(evals[0])
		for _, module := range d.names {
			if err := d.evalClean(s, module); err != nil {
				return fmt.Errorf("%s eval: %w", canary, err)
			}
		}
	}
	if len(bounds) > 0 {
		if err := d.boundaryClean(d.variantSnapshot(bounds[0])); err != nil {
			return fmt.Errorf("%s boundaries: %w", canary, err)
		}
	}
	return nil
}
func (d *doctorState) probeHooks() {
	evals := d.variants("eval")
	bounds := d.variants("boundaries")
	// Probe variants supply temporary hooks, so check required coverage on the
	// real modules before those hooks can mask a missing configuration.
	real := d.snapshot
	real.Modules = nil
	globalBoundary := mapValue(real.Config, "hooks")["boundaries"]
	missingBoundary := false
	for _, m := range d.snapshot.Modules {
		if m.Name == d.names[0] || m.Name == d.names[1] {
			continue
		}
		real.Modules = append(real.Modules, m)
		if globalBoundary == nil && mapValue(m.Document.Frontmatter, "hooks")["boundaries"] == nil && len(m.Zones.Disposable) > 0 && boundaryReason(m) == "" {
			missingBoundary = true
		}
	}
	if missingBoundary {
		ds, err := Boundaries(d.ctx, model.BoundaryRequest{Snapshot: real, Bin: d.req.Bin, Timeout: d.req.Timeout})
		for _, diagnostic := range ds {
			if !diagnostic.Warning {
				d.report.Diagnostics = append(d.report.Diagnostics, diagnostic)
			}
		}
		d.fail("baseline boundaries", "", "boundary hook", fmt.Sprintf("required hook missing: %v", err))
		for _, name := range []string{"undeclared", "internal_eval"} {
			d.fail(name, "", "boundary hook", "skipped after missing required boundary hook")
		}
		return
	}
	for _, v := range evals {
		s := d.variantSnapshot(v)
		label := "baseline " + v.name
		if err := d.evalClean(s, d.names[0]); err != nil {
			d.fail(label, d.names[0], "eval", err.Error())
			for _, name := range []string{"failing_eval", "missing_id"} {
				d.fail(name+" "+v.name, d.names[0], "eval", "skipped after failed baseline")
			}
			continue
		}
		d.check(label, d.names[0], true, "eval and coverage pass")
		for _, c := range []struct {
			name  string
			files map[string]string
		}{{"failing_eval", d.canaries.FailingEval}, {"missing_id", d.canaries.MissingID}} {
			if err := d.crossClean("baseline "+c.name, evals, bounds); err != nil {
				d.fail("baseline "+c.name, d.names[0], "eval", err.Error())
				continue
			}
			d.probeEval(s, v, c.name, c.files)
			if err := d.crossClean("recovery "+c.name, evals, bounds); err != nil {
				d.fail("recovery "+c.name, d.names[0], "eval", err.Error())
			}
		}
	}
	for _, v := range bounds {
		s := d.variantSnapshot(v)
		label := "baseline " + v.name
		if err := d.boundaryClean(s); err != nil {
			d.fail(label, d.names[0], "boundaries", err.Error())
			for _, c := range []struct{ name, module string }{{"undeclared", d.names[0]}, {"internal_eval", d.names[1]}} {
				d.fail(c.name+" "+v.name, c.module, "boundaries", "skipped after failed baseline")
			}
			continue
		}
		d.check(label, d.names[0], true, "boundaries pass")
		for _, c := range []struct {
			name   string
			files  map[string]string
			module int
		}{{"undeclared", d.canaries.Undeclared, 0}, {"internal_eval", d.canaries.InternalEval, 1}} {
			if err := d.crossClean("baseline "+c.name, evals, bounds); err != nil {
				d.fail("baseline "+c.name, d.names[c.module], "boundary hook", err.Error())
				continue
			}
			d.probeBoundary(s, v, c.name, c.files, c.module)
			if err := d.crossClean("recovery "+c.name, evals, bounds); err != nil {
				d.fail("recovery "+c.name, d.names[c.module], "boundary hook", err.Error())
			}
		}
	}
	if len(evals) == 0 {
		d.issue("C14", "eval hook missing for canaries", "", false)
	}
}
func (d *doctorState) plant(name string, i int, files map[string]string) ([]string, error) {
	paths, err := d.writeMap(d.dirs[i], files, d.values(i))
	if err != nil {
		return nil, fmt.Errorf("%s canary: %w", name, err)
	}
	return paths, nil
}
func removeCanary(paths []string) {
	for _, p := range paths {
		_ = os.Remove(p)
	}
}
func (d *doctorState) probeEval(s project.Snapshot, v doctorVariant, name string, files map[string]string) {
	module := d.names[0]
	paths, err := d.plant(name, 0, files)
	if err != nil {
		d.fail(name, module, "eval", err.Error())
		return
	}
	report, e := Evaluate(d.ctx, model.HookRequest{Snapshot: s, Module: module, Bin: d.req.Bin, Timeout: d.req.Timeout})
	detected := false
	switch name {
	case "failing_eval":
		if e == nil && report.HookExit != 0 {
			for _, c := range report.Cases {
				if c.Result == "fail" {
					for _, id := range c.IDs {
						if id == d.prefixes[0]+"-001" {
							detected = true
						}
					}
				}
			}
		}
	case "missing_id":
		code, _, coded := contracts.ErrorCode(e)
		var unmapped unmappedCaseError
		if coded && code == "C11" && errors.As(e, &unmapped) {
			for _, body := range files {
				for _, token := range canaryCaseToken.FindAllString(renderDoctorText(body, d.values(0)), -1) {
					if strings.Contains(unmapped.name, token) {
						detected = true
					}
				}
			}
		}
	}
	if !detected {
		d.fail(name+" "+v.name, module, "eval", fmt.Sprintf("missed %s canary: report=%+v error=%v", name, report, e))
	} else {
		d.check(name+" "+v.name, module, true, "eval detected "+name)
	}
	removeCanary(paths)
	if err := d.evalClean(s, module); err != nil {
		d.fail("recovery "+name+" "+v.name, module, "eval", err.Error())
	} else {
		d.check("recovery "+v.name, module, true, name+": eval and coverage pass")
	}
}
func (d *doctorState) probeBoundary(s project.Snapshot, v doctorVariant, name string, files map[string]string, i int) {
	module := d.names[i]
	paths, err := d.plant(name, i, files)
	if err != nil {
		d.fail(name, module, "boundaries", err.Error())
		return
	}
	ds, e := Boundaries(d.ctx, model.BoundaryRequest{Snapshot: s, Bin: d.req.Bin, Timeout: d.req.Timeout})
	detected := false
	for _, diag := range ds {
		if diag.Code != "X01" || !strings.HasPrefix(diag.Message, module+":") {
			continue
		}
		lower := strings.ToLower(diag.Message)
		if name == "undeclared" || strings.Contains(lower, "interface") || strings.Contains(lower, "retained interface") {
			detected = true
		}
	}
	if !detected {
		d.fail(name+" "+v.name, module, "boundaries", fmt.Sprintf("missed %s canary: diagnostics=%v error=%v", name, ds, e))
	} else {
		d.check(name+" "+v.name, module, true, "boundaries detected "+name)
	}
	removeCanary(paths)
	if err := d.boundaryClean(s); err != nil {
		d.fail("recovery "+name+" "+v.name, module, "boundaries", err.Error())
	} else {
		d.check("recovery "+v.name, module, true, name+": boundaries pass")
	}
}
func agentArgs(command any, values map[string]string) ([]string, error) {
	replace := func(s string, quote bool) string {
		for _, k := range []string{"prompt", "git_common_dir"} {
			v := values[k]
			if quote {
				v = shellQuote(v)
			}
			s = strings.ReplaceAll(s, "{"+k+"}", v)
		}
		return s
	}
	switch v := command.(type) {
	case string:
		if strings.TrimSpace(v) == "" {
			return nil, errors.New("empty agent command")
		}
		return []string{"sh", "-c", replace(v, true)}, nil
	case []any:
		args := make([]string, len(v))
		for i, x := range v {
			s, ok := x.(string)
			if !ok {
				return nil, errors.New("agent argv contains non-string")
			}
			args[i] = replace(s, false)
		}
		if len(args) == 0 || args[0] == "" {
			return nil, errors.New("empty agent command")
		}
		return args, nil
	case []string:
		args := make([]string, len(v))
		for i, s := range v {
			args[i] = replace(s, false)
		}
		if len(args) == 0 || args[0] == "" {
			return nil, errors.New("empty agent command")
		}
		return args, nil
	default:
		return nil, fmt.Errorf("invalid agent command %T", command)
	}
}
func (d *doctorState) probeAgents(common string) {
	agent := mapValue(d.snapshot.Config, "agent")
	audit := mapValue(d.snapshot.Config, "audit")
	worker := agent["command"]
	auditCommand := audit["command"]
	if auditCommand == nil {
		auditCommand = worker
		d.issue("C15", "audit uses shared agent command; shared-command blind spot", "", true)
	}
	for _, v := range []struct {
		name    string
		command any
	}{{"worker", worker}, {"audit", auditCommand}} {
		d.probeAgent(v.name, v.command, agent["wrap"], common)
	}
}
func (d *doctorState) probeAgent(name string, command, wrap any, common string) {
	values := map[string]string{"prompt": "Reply with exactly " + doctorReply, "git_common_dir": common}
	args, err := agentArgs(command, values)
	display := fmt.Sprint(command)
	if err == nil && wrap != nil {
		var prefix []string
		prefix, err = agentArgs(wrap, values)
		if err == nil {
			args = append(prefix, args...)
		}
	}
	if err != nil {
		d.check(name, "", false, err.Error())
		d.issue("C15", name+" command "+display+": "+err.Error(), "", false)
		return
	}
	timeout := d.req.Timeout
	if timeout <= 0 {
		timeout = 3 * time.Minute
	}
	p, e := gitstore.Run(d.ctx, gitstore.Request{Args: args, Dir: d.root, Timeout: timeout})
	tail := finalLines(p.Stdout+"\n"+p.Stderr, 20)
	ok := e == nil && p.ExitCode == 0 && strings.Contains(p.Stdout+p.Stderr, doctorReply)
	if ok {
		d.check(name, "", true, "fixed reply received")
		return
	}
	d.check(name, "", false, fmt.Sprintf("exit %d: %v", p.ExitCode, e))
	kind := "command failure"
	for _, pattern := range doctorStringList(d.snapshot.Config["limit_patterns"]) {
		if pattern != "" && strings.Contains(strings.ToLower(tail), strings.ToLower(pattern)) {
			kind = "usage limit"
			break
		}
	}
	d.issue("C15", fmt.Sprintf("%s %s: %s; exit %d; %v\n%s", name, kind, display, p.ExitCode, e, tail), "", false)
}
func finalLines(s string, n int) string {
	lines := strings.Split(strings.TrimRight(s, "\n"), "\n")
	if len(lines) > n {
		lines = lines[len(lines)-n:]
	}
	return strings.Join(lines, "\n")
}
func doctorStringList(v any) []string {
	var out []string
	switch x := v.(type) {
	case []any:
		for _, a := range x {
			if s, ok := a.(string); ok {
				out = append(out, s)
			}
		}
	case []string:
		out = append(out, x...)
	}
	return out
}
