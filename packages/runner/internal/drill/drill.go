package drill

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/lurkshark/fireweed/packages/checks"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/runner/internal/invoke"
	"github.com/lurkshark/fireweed/packages/runner/internal/next"
	"github.com/lurkshark/fireweed/packages/runner/model"
)

// Run rebuilds a module in an independent input repository, then tests the
// resulting files in a full, disposable checkout of the invoking commit.
func Run(ctx context.Context, request model.DrillRequest) (result model.DrillResult, err error) {
	result.Module = request.Module
	result.Dependents = make(map[string]checks.Report)
	snapshot, err := project.Load(request.Root)
	if err != nil {
		return result, err
	}
	target := findModule(snapshot, request.Module)
	if target == nil {
		return result, fmt.Errorf("U02: unknown module %q", request.Module)
	}
	repo, err := gitstore.Inspect(request.Root)
	if err != nil {
		return result, err
	}
	result.Base = repo.Head
	status, err := git(repo.Root, "--no-optional-locks", "status", "--porcelain", "-uall")
	if err != nil {
		return result, err
	}
	if status != "" {
		return result, fmt.Errorf("drill requires a clean invoking checkout: dirty worktree")
	}
	claims, err := git(repo.Root, "for-each-ref", "--format=%(refname)", "refs/fireweed/claims")
	if err != nil {
		return result, err
	}
	if claims != "" {
		return result, fmt.Errorf("drill requires no live shared change claim: %s", claims)
	}

	result.Date = time.Now().UTC().Format("20060102T150405.000000000")
	result.Result = "fail"
	result.ReportPath = filepath.ToSlash(filepath.Join(".fireweed", "drills", request.Module+".md"))
	result.Agent = agentName(snapshot.Config)
	result.Model = modelName(snapshot.Config)
	parent := filepath.Join(repo.CommonDir, "fireweed", "worktrees")
	if err = os.MkdirAll(parent, 0755); err != nil {
		return result, err
	}
	scratch, err := os.MkdirTemp(parent, "drill-")
	if err != nil {
		return result, err
	}
	if err = os.Remove(scratch); err != nil {
		return result, err
	}
	worktreeAdded := false
	input := ""
	defer func() {
		var cleanup error
		if input != "" {
			cleanup = errors.Join(cleanup, os.RemoveAll(input))
		}
		if worktreeAdded {
			parentInfo, statErr := os.Stat(parent)
			if statErr != nil {
				cleanup = errors.Join(cleanup, statErr)
			} else if parentInfo.Mode().Perm()&0200 == 0 {
				cleanup = errors.Join(cleanup, fmt.Errorf("scratch parent %s is not writable", parent))
			} else {
				_, removeErr := git(repo.Root, "worktree", "remove", "--force", scratch)
				cleanup = errors.Join(cleanup, removeErr)
			}
		} else {
			cleanup = errors.Join(cleanup, os.RemoveAll(scratch))
		}
		if cleanup != nil {
			err = errors.Join(err, fmt.Errorf("drill cleanup: %w", cleanup))
			result.Result = "fail"
			if len(result.Steps) > 0 {
				err = errors.Join(err, markReportFailure(repo.Root, result.ReportPath, cleanup))
			}
		}
	}()
	if _, err = git(repo.Root, "worktree", "add", "--detach", scratch, repo.Head); err != nil {
		return result, err
	}
	worktreeAdded = true
	if err = setup(ctx, scratch, snapshot.Config, request); err != nil {
		return result, err
	}

	// Delete only the target's disposable files. The verifier retains all other
	// project files, including dependent implementation and hook configuration.
	if err = removeDisposable(scratch, snapshot, *target); err != nil {
		return result, err
	}
	input, err = makeInput(scratch, repo.Head, snapshot, *target)
	if err != nil {
		return result, err
	}
	if err = setup(ctx, input, snapshot.Config, request); err != nil {
		return result, err
	}
	original, err := inventory(input)
	if err != nil {
		return result, err
	}
	eval, err := evalCommand(input, snapshot.Config, *target)
	if err != nil {
		return result, err
	}
	prompt, err := reconstructionPrompt(*target, original, eval)
	if err != nil {
		return result, err
	}
	attempts := configuredAttempts(snapshot.Config)
	wait, budget, err := usagePolicy(snapshot.Config)
	if err != nil {
		return result, err
	}
	var waited time.Duration
	feedback := ""
	accepted := false
	var failure error
	logRun, err := invoke.NewRunID()
	if err != nil {
		return result, err
	}
	for attempt := 0; attempt < attempts; {
		if err = ctx.Err(); err != nil {
			failure = err
			break
		}
		agentSnapshot := snapshot
		agentSnapshot.Root = input
		step, invokeErr := invoke.Agent(ctx, model.AgentRequest{
			Snapshot: agentSnapshot, Role: "worker", Prompt: prompt + feedback,
			Bin: request.Bin, Timeout: request.Timeout,
			LogRoot: repo.Root, Run: logRun, Phase: "drill", Attempt: attempt + 1,
		})
		result.Steps = append(result.Steps, step)
		if validationErr := validateInput(input, original); validationErr != nil {
			failure = fmt.Errorf("invalid reconstruction edits: %w", validationErr)
			break
		}
		if ctx.Err() != nil {
			failure = ctx.Err()
			break
		}
		if errors.Is(invokeErr, context.DeadlineExceeded) || errors.Is(invokeErr, context.Canceled) {
			failure = invokeErr
			break
		}
		if invoke.IsLogError(invokeErr) {
			failure = invokeErr
			break
		}
		if step.UsageLimit {
			if waited > budget || wait > budget-waited {
				failure = &next.UsageStopError{Completed: waited, Next: wait, Budget: budget}
				result.Result = "stopped"
				break
			}
			timer := time.NewTimer(wait)
			select {
			case <-ctx.Done():
				failure = ctx.Err()
			case <-timer.C:
				waited += wait
			}
			timer.Stop()
			if failure != nil {
				break
			}
			continue
		}
		attempt++
		if invokeErr != nil {
			failure = fmt.Errorf("reconstruction agent: %w", invokeErr)
			feedback = "\n\n# Previous reconstruction failure\n" + failure.Error() + "\n" + step.Stdout + "\n" + step.Stderr
			continue
		}
		if err = importDisposable(input, scratch, snapshot, *target, original); err != nil {
			failure = err
			break
		}
		accepted = true
		verification := verifyAll(ctx, snapshot, scratch, *target, request, &result)
		if verification == nil {
			failure = nil
			break
		}
		failure = verification
		feedback = "\n\n# Previous verification failed\n" + verification.Error() + failureDetails(result)
		if ctx.Err() != nil {
			failure = ctx.Err()
			break
		}
	}
	if failure == nil && !accepted {
		failure = fmt.Errorf("reconstruction attempts exhausted")
	}
	if accepted {
		result.Ref = "refs/fireweed/drills/" + request.Module + "/" + result.Date
		if retainErr := retain(scratch, repo.Root, result.Ref, *target); retainErr != nil {
			failure = errors.Join(failure, retainErr)
			result.Ref = ""
		}
	}
	if failure == nil {
		result.Result = "pass"
	}
	if len(result.Steps) > 0 {
		if reportErr := writeReport(repo.Root, input, result, failure); reportErr != nil {
			failure = errors.Join(failure, reportErr)
			result.Result = "fail"
		}
	}
	if failure != nil {
		return result, failure
	}
	return result, nil
}

func section(config map[string]any, key string) map[string]any {
	value, _ := config[key].(map[string]any)
	return value
}
func findModule(snapshot project.Snapshot, name string) *project.Module {
	for i := range snapshot.Modules {
		if snapshot.Modules[i].Name == name {
			return &snapshot.Modules[i]
		}
	}
	return nil
}
func agentName(config map[string]any) string {
	command := section(config, "agent")["command"]
	switch v := command.(type) {
	case string:
		if fields := strings.Fields(v); len(fields) > 0 {
			return fields[0]
		}
	case []any:
		if len(v) > 0 {
			return fmt.Sprint(v[0])
		}
	case []string:
		if len(v) > 0 {
			return v[0]
		}
	}
	return "unknown"
}

func modelName(config map[string]any) string {
	agent := section(config, "agent")
	var args []string
	switch command := agent["command"].(type) {
	case string:
		args = strings.Fields(command)
	case []string:
		args = command
	case []any:
		for _, part := range command {
			args = append(args, fmt.Sprint(part))
		}
	}
	for i, arg := range args {
		if (arg == "-m" || arg == "--model") && i+1 < len(args) {
			return strings.Trim(args[i+1], "'\"")
		}
		if strings.HasPrefix(arg, "--model=") {
			return strings.Trim(strings.TrimPrefix(arg, "--model="), "'\"")
		}
	}
	model, _ := agent["model"].(string)
	return model
}

func git(root string, args ...string) (string, error) {
	out, err := gitOutput(root, args...)
	return strings.TrimSpace(out), err
}

func gitOutput(root string, args ...string) (string, error) {
	argv := append([]string{"git", "-C", root, "-c", "core.hooksPath=/dev/null"}, args...)
	r, err := gitstore.Run(context.Background(), gitstore.Request{Args: argv, Timeout: 30 * time.Second, Env: map[string]string{"GIT_NO_LAZY_FETCH": "1", "GIT_TERMINAL_PROMPT": "0"}})
	if err != nil {
		return "", fmt.Errorf("git %s: %w: %s", strings.Join(args, " "), err, strings.TrimSpace(r.Stderr))
	}
	return r.Stdout, nil
}

func setup(ctx context.Context, root string, config map[string]any, request model.DrillRequest) error {
	command := section(config, "hooks")["setup"]
	if command == nil {
		return nil
	}
	var args []string
	switch v := command.(type) {
	case string:
		args = []string{"sh", "-c", v}
	case []any:
		for _, part := range v {
			s, ok := part.(string)
			if !ok {
				return fmt.Errorf("invalid setup argv")
			}
			args = append(args, s)
		}
	case []string:
		args = v
	default:
		return fmt.Errorf("invalid setup hook")
	}
	if len(args) == 0 {
		return fmt.Errorf("empty setup hook")
	}
	timeout := request.Timeout
	if timeout <= 0 {
		timeout = configuredTimeout(config)
	}
	r, err := gitstore.Run(ctx, gitstore.Request{Args: args, Dir: root, Env: map[string]string{"FIREWEED_BIN": request.Bin}, Timeout: timeout})
	if err != nil {
		return fmt.Errorf("setup failed (exit %d): %w: %s %s", r.ExitCode, err, r.Stdout, r.Stderr)
	}
	return nil
}

func walkFiles(root string, visit func(rel string, entry fs.DirEntry) error) error {
	return filepath.WalkDir(root, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() && path == filepath.Join(root, ".git") {
			return filepath.SkipDir
		}
		if path == root || entry.IsDir() {
			return nil
		}
		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		return visit(filepath.ToSlash(rel), entry)
	})
}

// disposable applies module precedence and the approved project protection to
// a target path.
func disposable(snapshot project.Snapshot, module project.Module, rel string) bool {
	if contracts.ResolveZone(rel, module.Zones) != "disposable" {
		return false
	}
	full := filepath.ToSlash(filepath.Join(module.Dir, filepath.FromSlash(rel)))
	for _, segment := range strings.Split(rel, "/") {
		if segment == ".git" || segment == ".fireweed" {
			return false
		}
	}
	if base := filepath.Base(rel); base == "AGENTS.md" || base == "CLAUDE.md" {
		return false
	}
	var protected []string
	switch configured := snapshot.Config["protected"].(type) {
	case []string:
		protected = configured
	case []any:
		for _, value := range configured {
			if pattern, ok := value.(string); ok {
				protected = append(protected, pattern)
			}
		}
	}
	for _, pattern := range protected {
		if contracts.MatchGlob(pattern, full) {
			return false
		}
	}
	return true
}

func removeDisposable(root string, snapshot project.Snapshot, module project.Module) error {
	moduleRoot := filepath.Join(root, filepath.FromSlash(module.Dir))
	return walkFiles(moduleRoot, func(rel string, _ fs.DirEntry) error {
		if disposable(snapshot, module, rel) {
			return os.Remove(filepath.Join(moduleRoot, filepath.FromSlash(rel)))
		}
		return nil
	})
}

func dependencies(snapshot project.Snapshot, target project.Module) map[string]bool {
	seen := map[string]bool{}
	var visit func(string)
	visit = func(name string) {
		m := findModule(snapshot, name)
		if m == nil {
			return
		}
		for _, dep := range m.DependsOn {
			if !seen[dep] {
				seen[dep] = true
				visit(dep)
			}
		}
	}
	visit(target.Name)
	delete(seen, target.Name)
	return seen
}
func dependents(snapshot project.Snapshot, target project.Module) []string {
	var names []string
	for _, module := range snapshot.Modules {
		if module.Name != target.Name && dependencies(snapshot, module)[target.Name] {
			names = append(names, module.Name)
		}
	}
	sort.Strings(names)
	return names
}
func makeInput(scratch, base string, snapshot project.Snapshot, target project.Module) (string, error) {
	root, err := os.MkdirTemp("", "fireweed-drill-input-")
	if err != nil {
		return "", err
	}
	// Return the path even on failure so Run can clean it up.
	tracked, err := gitOutput(scratch, "ls-tree", "-r", "--name-only", "-z", base)
	if err != nil {
		return root, err
	}
	excluded := []string{"changes", ".fireweed/drills"}
	for _, name := range dependents(snapshot, target) {
		excluded = append(excluded, findModule(snapshot, name).Dir)
	}
	for _, path := range strings.Split(strings.TrimSuffix(tracked, "\x00"), "\x00") {
		if path == "" || excludedInput(path, excluded) {
			continue
		}
		if rel, ok := strings.CutPrefix(path, target.Dir+"/"); ok && disposable(snapshot, target, rel) {
			continue
		}
		file, err := readInputFile(scratch, path)
		if err != nil {
			return root, fmt.Errorf("read reconstruction input %s: %w", path, err)
		}
		if err := file.write(root, path); err != nil {
			return root, fmt.Errorf("copy reconstruction input %s: %w", path, err)
		}
	}
	return root, independentGit(root)
}

func excludedInput(path string, excluded []string) bool {
	for _, dir := range excluded {
		if path == dir || strings.HasPrefix(path, dir+"/") {
			return true
		}
	}
	return false
}

func independentGit(root string) error {
	for _, args := range [][]string{{"init", "-q", "--template="}, {"add", "-f", "-A"}, {"-c", "user.name=Fireweed Drill", "-c", "user.email=drill@localhost", "commit", "-q", "-m", "reconstruction inputs"}} {
		cmd := exec.Command("git", args...)
		cmd.Dir = root
		for _, v := range os.Environ() {
			if !strings.HasPrefix(v, "GIT_") {
				cmd.Env = append(cmd.Env, v)
			}
		}
		cmd.Env = append(cmd.Env, "GIT_CONFIG_NOSYSTEM=1", "GIT_CONFIG_GLOBAL=/dev/null", "GIT_NO_LAZY_FETCH=1")
		if out, err := cmd.CombinedOutput(); err != nil {
			return fmt.Errorf("input git %v: %w: %s", args, err, out)
		}
	}
	return nil
}

type inputFile struct {
	mode    fs.FileMode
	content string
}

func readInputFile(root, rel string) (inputFile, error) {
	path := filepath.Join(root, filepath.FromSlash(rel))
	info, err := os.Lstat(path)
	if err != nil {
		return inputFile{}, err
	}
	mode := info.Mode()
	switch {
	case mode.IsRegular():
		data, err := os.ReadFile(path)
		return inputFile{mode: mode, content: string(data)}, err
	case mode.Type() == fs.ModeSymlink:
		link, err := os.Readlink(path)
		return inputFile{mode: mode, content: link}, err
	default:
		return inputFile{}, fmt.Errorf("unsupported file type at %s", rel)
	}
}

func (file inputFile) write(root, rel string) error {
	path := filepath.Join(root, filepath.FromSlash(rel))
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	if file.mode.Type() == fs.ModeSymlink {
		return os.Symlink(file.content, path)
	}
	return os.WriteFile(path, []byte(file.content), file.mode.Perm())
}

func inventory(root string) (map[string]inputFile, error) {
	files := map[string]inputFile{}
	err := walkFiles(root, func(rel string, _ fs.DirEntry) error {
		file, err := readInputFile(root, rel)
		if err == nil {
			files[rel] = file
		}
		return err
	})
	return files, err
}

func evalCommand(root string, config map[string]any, target project.Module) (string, error) {
	hook := section(config, "hooks")["eval"]
	if override, ok := section(target.Document.Frontmatter, "hooks")["eval"]; ok {
		hook = override
	}
	const reportDir = ".fireweed/runs/drill"
	const prefix = "mkdir -p " + reportDir + " && "
	// Like checks.Evaluate, pass a path that remains valid if the hook changes directories.
	reportPath, err := filepath.Abs(filepath.Join(root, reportDir, "report.xml"))
	if err != nil {
		return "", err
	}
	quote := func(value string) string {
		return "'" + strings.ReplaceAll(value, "'", "'\"'\"'") + "'"
	}
	expansion := func(escape func(string) string) *strings.Replacer {
		return strings.NewReplacer(
			"{module}", escape(target.Name),
			"{module_dir}", escape(target.Dir),
			"{report}", escape(reportPath),
		)
	}
	var args []string
	switch command := hook.(type) {
	case string:
		if strings.TrimSpace(command) != "" {
			return prefix + "sh -c " + quote(expansion(quote).Replace(command)), nil
		}
	case []string:
		args = command
	case []any:
		for _, arg := range command {
			value, ok := arg.(string)
			if !ok {
				return "", fmt.Errorf("C11: eval hook argv for %s contains non-string", target.Name)
			}
			args = append(args, value)
		}
	}
	if len(args) == 0 || args[0] == "" {
		return "", fmt.Errorf("C11: missing or invalid eval hook for %s", target.Name)
	}
	expand := expansion(func(value string) string { return value })
	words := make([]string, len(args))
	for i, arg := range args {
		words[i] = quote(expand.Replace(arg))
	}
	return prefix + strings.Join(words, " "), nil
}

func reconstructionPrompt(target project.Module, original map[string]inputFile, eval string) (string, error) {
	spec, err := contracts.RenderDocument(target.Document)
	if err != nil {
		return "", err
	}
	var retained, evals []string
	for path := range original {
		if strings.HasPrefix(path, target.Dir+"/") {
			rel := strings.TrimPrefix(path, target.Dir+"/")
			if contracts.ResolveZone(rel, target.Zones) == "evals" {
				evals = append(evals, path)
			} else {
				retained = append(retained, path)
			}
		}
	}
	sort.Strings(retained)
	sort.Strings(evals)
	return fmt.Sprintf(`# Rebuild %s

## Target spec and Decisions

%s

## Retained interfaces and files

%s

## Retained evals

%s

Rebuild only the target's disposable paths. Keep all retained inputs intact and do not read outside the input workspace. Do not commit.

Run the target's eval hook from the workspace root as often as needed.

~~~sh
%s
~~~

Write DRILL_NOTES.md with separate ## Contradicted Decisions and ## Gaps sections. Every gap must name the nearest spec ID or "no spec" and a target module.
`, target.Name, spec, strings.Join(retained, "\n"), strings.Join(evals, "\n"), eval), nil
}
func validateInput(root string, original map[string]inputFile) error {
	info, err := os.Lstat(root)
	if err != nil || !info.IsDir() {
		return fmt.Errorf("missing or replaced reconstruction input workspace")
	}
	checkedDirs := make(map[string]bool)
	for path, before := range original {
		// A substituted directory symlink must not redirect retained reads.
		for dir := filepath.Dir(path); dir != "." && !checkedDirs[dir]; dir = filepath.Dir(dir) {
			info, err := os.Lstat(filepath.Join(root, dir))
			if err != nil || !info.IsDir() {
				return fmt.Errorf("missing or replaced retained directory %s", dir)
			}
			checkedDirs[dir] = true
		}
		after, err := readInputFile(root, path)
		if errors.Is(err, fs.ErrNotExist) {
			return fmt.Errorf("deleted retained file %s", path)
		}
		if err != nil || after != before {
			return fmt.Errorf("modified retained file %s", path)
		}
	}
	return nil
}
func importDisposable(input, scratch string, snapshot project.Snapshot, target project.Module, original map[string]inputFile) error {
	inputModule := filepath.Join(input, filepath.FromSlash(target.Dir))
	rebuilt := make(map[string]inputFile)
	err := walkFiles(inputModule, func(rel string, entry fs.DirEntry) error {
		path := target.Dir + "/" + rel
		if _, retained := original[path]; retained || !disposable(snapshot, target, rel) {
			return nil
		}
		if !entry.Type().IsRegular() {
			return fmt.Errorf("nonregular reconstruction output %s", path)
		}
		file, err := readInputFile(inputModule, rel)
		if err != nil {
			return err
		}
		rebuilt[path] = file
		return nil
	})
	if err != nil {
		return err
	}
	if err := removeDisposable(scratch, snapshot, target); err != nil {
		return err
	}
	for path, file := range rebuilt {
		if err := file.write(scratch, path); err != nil {
			return err
		}
	}
	return nil
}
func specs(module project.Module) ([]contracts.Item, error) {
	section, _ := contracts.Section(module.Document.Body, "Specs")
	return contracts.ParseItems(section)
}
func verifyAll(ctx context.Context, snapshot project.Snapshot, scratch string, target project.Module, request model.DrillRequest, result *model.DrillResult) error {
	snapshot.Root = scratch
	var failures []error
	names := append([]string{target.Name}, dependents(snapshot, target)...)
	for _, name := range names {
		if err := ctx.Err(); err != nil {
			return err
		}
		report, err := checks.Evaluate(ctx, checks.HookRequest{Snapshot: snapshot, Module: name, Bin: request.Bin, Timeout: request.Timeout})
		if name == target.Name {
			result.Own = report
		} else {
			result.Dependents[name] = report
		}
		module := findModule(snapshot, name)
		items, specErr := specs(*module)
		err = errors.Join(err, specErr)
		if specErr == nil {
			err = errors.Join(err, checks.RequirePass(report, items), checks.RequireCoverage(report, items))
		}
		if err != nil {
			failures = append(failures, fmt.Errorf("%s eval: %w", name, err))
		}
	}
	return errors.Join(failures...)
}
func failureDetails(result model.DrillResult) string {
	var b strings.Builder
	write := func(name string, report checks.Report) {
		fmt.Fprintf(&b, "\n%s stdout: %s\n%s stderr: %s\n", name, report.Stdout, name, report.Stderr)
		for _, c := range report.Cases {
			if c.Result == "fail" {
				fmt.Fprintf(&b, "%s %s %s: %s\n", name, c.Classname, c.Name, c.Failure)
			}
		}
	}
	write(result.Module, result.Own)
	names := make([]string, 0, len(result.Dependents))
	for name := range result.Dependents {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		write(name, result.Dependents[name])
	}
	return b.String()
}
func retain(scratch, root, ref string, target project.Module) error {
	if _, err := git(scratch, "add", "-A", "--", target.Dir); err != nil {
		return err
	}
	if _, err := git(scratch, "-c", "user.name=Fireweed Drill", "-c", "user.email=drill@localhost", "commit", "-q", "--allow-empty", "-m", "rebuilt drill code"); err != nil {
		return err
	}
	oid, err := git(scratch, "rev-parse", "HEAD")
	if err != nil {
		return err
	}
	_, err = git(root, "update-ref", ref, oid)
	return err
}

func configuredTimeout(config map[string]any) time.Duration {
	if raw, ok := config["timeout"].(string); ok {
		if duration, err := time.ParseDuration(raw); err == nil && duration > 0 {
			return duration
		}
	}
	return 90 * time.Minute
}
func writeReport(root, input string, result model.DrillResult, failure error) error {
	date, err := time.Parse("20060102T150405.000000000", result.Date)
	if err != nil {
		return fmt.Errorf("drill report date: %w", err)
	}
	var b strings.Builder
	fmt.Fprintf(&b, "---\ndate: %s\nbase_commit: %s\nagent: %s\nmodel: %s\nresult: %s\n---\n\n## Own evals\n", date.UTC().Format(time.RFC3339Nano), result.Base, result.Agent, result.Model, result.Result)
	writeCases := func(name string, report checks.Report) {
		for _, c := range report.Cases {
			fmt.Fprintf(&b, "- %s | %s | %s | %s\n", name, c.Name, strings.Join(c.IDs, ", "), c.Result)
			if c.Failure != "" {
				fmt.Fprintf(&b, "  Failure: %s\n", c.Failure)
			}
		}
		if report.Stdout != "" {
			fmt.Fprintf(&b, "  Stdout: %s\n", report.Stdout)
		}
		if report.Stderr != "" {
			fmt.Fprintf(&b, "  Stderr: %s\n", report.Stderr)
		}
	}
	writeCases(result.Module, result.Own)
	b.WriteString("\n## Dependent evals\n")
	names := make([]string, 0, len(result.Dependents))
	for name := range result.Dependents {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		writeCases(name, result.Dependents[name])
	}
	var contradicted, gaps string
	if input != "" {
		if notes, err := readInputFile(input, "DRILL_NOTES.md"); err == nil && notes.mode.IsRegular() {
			contradicted, _ = contracts.Section(notes.content, "Contradicted Decisions")
			gaps, _ = contracts.Section(notes.content, "Gaps")
		}
	}
	fmt.Fprintf(&b, "\n## Contradicted Decisions\n%s\n## Gaps\n%s\n## Agent observations\n", strings.TrimSpace(contradicted), strings.TrimSpace(gaps))
	for i, step := range result.Steps {
		fmt.Fprintf(&b, "\n### Attempt %d (exit %d)\n%s\n%s\n", i+1, step.ExitCode, step.Stdout, step.Stderr)
	}
	if failure != nil {
		fmt.Fprintf(&b, "\n## Verification failures\n%s\n", failure)
	}
	path := filepath.Join(root, filepath.FromSlash(result.ReportPath))
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	return os.WriteFile(path, []byte(b.String()), 0644)
}

func markReportFailure(root, relative string, cause error) error {
	path := filepath.Join(root, filepath.FromSlash(relative))
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	report := strings.Replace(string(data), "result: pass\n", "result: fail\n", 1)
	report += fmt.Sprintf("\n## Cleanup failure\n%s\n", cause)
	return os.WriteFile(path, []byte(report), 0644)
}

func configuredAttempts(config map[string]any) int {
	switch v := config["attempts"].(type) {
	case int:
		if v > 0 {
			return v
		}
	case float64:
		if v > 0 && v == float64(int(v)) {
			return int(v)
		}
	}
	return 2
}
func usagePolicy(config map[string]any) (time.Duration, time.Duration, error) {
	parse := func(key string, fallback time.Duration) (time.Duration, error) {
		value := config[key]
		if value == nil || value == "" {
			return fallback, nil
		}
		s, ok := value.(string)
		if !ok {
			return 0, fmt.Errorf("C01: %s must be a duration", key)
		}
		d, err := time.ParseDuration(s)
		if err != nil {
			return 0, fmt.Errorf("C01: invalid %s: %w", key, err)
		}
		return d, nil
	}
	w, err := parse("limit_wait", 30*time.Minute)
	if err != nil {
		return 0, 0, err
	}
	b, err := parse("limit_budget", 4*time.Hour)
	if err != nil {
		return 0, 0, err
	}
	if w <= 0 || b < 0 {
		return 0, 0, fmt.Errorf("C01: invalid usage wait or budget")
	}
	return w, b, nil
}
