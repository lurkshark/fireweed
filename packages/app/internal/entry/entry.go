package entry

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/lurkshark/fireweed/packages/assets"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
)

type response struct {
	status      int
	plain       string
	data        any
	diagnostics []project.Diagnostic
}

func output(plain string, data any) response {
	return response{plain: plain, data: data}
}

func message(s string) response {
	return output(s+"\n", s)
}

func verbatim(s string) response { return output(s, s) }

func failure(status int, code, path, message string) response {
	return response{status: status, diagnostics: []project.Diagnostic{{Code: code, Path: path, Message: message}}}
}

func coded(code, path string, err error) error {
	return contracts.CodedError{Code: code, Subject: path, Cause: err}
}

func cause(err error, fallback string) string {
	if code, _, ok := contracts.ErrorCode(err); ok {
		return code
	}
	return fallback
}

func failed(status int, err error, fallback, action string) response {
	_, subject, _ := contracts.ErrorCode(err)
	return failure(status, cause(err, fallback), subject, err.Error()+". "+action)
}

func configFailure(err error) response {
	return failed(2, err, "C01", "Correct the project configuration or module declarations and retry.")
}

func operationFailure(err error, action string) response {
	return failed(1, err, "P01", action)
}

func configuredOperationFailure(err error, action string) response {
	if cause(err, "") == "C01" {
		return configFailure(err)
	}
	return operationFailure(err, action)
}

func hasErrors(diagnostics []project.Diagnostic) bool {
	return slices.ContainsFunc(diagnostics, func(d project.Diagnostic) bool { return !d.Warning })
}

func checked(result response, diagnostics []project.Diagnostic, action string) response {
	result.diagnostics = make([]project.Diagnostic, 0, len(diagnostics))
	for _, d := range diagnostics {
		if !d.Warning {
			result.status, result.data = 1, nil
			if action != "" {
				result.plain = ""
				d.Message += ". " + action
			}
		}
		result.diagnostics = append(result.diagnostics, d)
	}
	return result
}

func array[T any](values []T) []T {
	if values == nil {
		return []T{}
	}
	return values
}

type diagnostic struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Path    string `json:"path,omitempty"`
	Warning bool   `json:"warning"`
}

func (r response) write(out, errOut io.Writer, jsonMode bool) int {
	if jsonMode {
		ds := make([]diagnostic, 0, len(r.diagnostics))
		for _, d := range r.diagnostics {
			ds = append(ds, diagnostic(d))
		}
		_ = json.NewEncoder(out).Encode(struct {
			Schema      int          `json:"schema"`
			Data        any          `json:"data"`
			Diagnostics []diagnostic `json:"diagnostics"`
		}{1, r.data, ds})
		return r.status
	}
	fmt.Fprint(out, r.plain)
	for _, d := range r.diagnostics {
		fmt.Fprintf(errOut, "%s: ", d.Code)
		if d.Warning {
			fmt.Fprint(errOut, "warning: ")
		}
		if d.Path != "" {
			fmt.Fprintf(errOut, "%s: ", d.Path)
		}
		fmt.Fprintln(errOut, d.Message)
	}
	return r.status
}

type command struct {
	name     string
	args     []string
	options  map[string]string
	max      int
	attempt  int
	duration time.Duration
}

func (c command) has(option string) bool {
	_, ok := c.options[option]
	return ok
}

func (c command) arg(index int) string {
	if index >= len(c.args) {
		return ""
	}
	return c.args[index]
}

type syntax struct {
	min, max      int
	flags, values string
}

var commands = map[string]syntax{
	"version":       {},
	"explain":       {min: 1, max: 1},
	"playbook":      {min: 1, max: 1},
	"init":          {flags: "--no-app", values: "--recipe --agent --audit-agent"},
	"doctor":        {flags: "--no-agent"},
	"agents-md":     {flags: "--write"},
	"modules":       {},
	"new-module":    {min: 1, max: 1, values: "--prefix"},
	"new-change":    {min: 2, max: 2},
	"stamp":         {min: 1, max: 1},
	"list":          {},
	"status":        {},
	"show":          {min: 1, max: 2, flags: "--projected --history --stale"},
	"context":       {min: 1, max: 1, flags: "--full"},
	"check":         {flags: "--fix"},
	"coverage":      {min: 1, max: 1, values: "--with-change"},
	"check-commits": {max: 1},
	"ci":            {max: 1, flags: "--no-doctor"},
	"pr-body":       {min: 1, max: 1},
	"start":         {min: 1, max: 1, values: "--from-attempt"},
	"verify":        {max: 1, flags: "--quick"},
	"done":          {max: 1},
	"next":          {flags: "--allow-unmerged"},
	"run":           {flags: "--allow-unmerged", values: "--max --for"},
	"release":       {min: 1, max: 1},
	"commits":       {min: 1, max: 1},
	"abandon":       {min: 1, max: 1},
	"drill":         {min: 1, max: 1},
	"gc":            {},
	"migrate":       {},
}

const helpUsage = "Use fireweed help or fireweed help <command>."
const initUsage = "Use fireweed init --recipe <recipe> --agent <agent>. Recipes: go, typescript-pnpm, python-uv, none, or a local directory. Agents: codex, claude, or a shell template with {prompt} or {prompt_file}."

func help(args []string) (response, bool) {
	if len(args) == 0 {
		return verbatim(assets.Help()), true
	}
	var target string
	switch {
	case args[0] == "help" && len(args) == 1,
		(args[0] == "--help" || args[0] == "-h") && len(args) == 1:
		return verbatim(assets.Help()), true
	case args[0] == "help" && len(args) == 2 && !strings.HasPrefix(args[1], "-"):
		target = args[1]
	case len(args) == 2 && !strings.HasPrefix(args[0], "-") &&
		(args[1] == "help" || args[1] == "--help" || args[1] == "-h"):
		target = args[0]
	case args[0] == "help" || slices.Contains(args, "--help") || slices.Contains(args, "-h"):
		return failure(2, "U01", "", fmt.Sprintf("Invalid help request %q. %s", args, helpUsage)), true
	default:
		return response{}, false
	}
	page, err := assets.CommandHelp(target)
	if err != nil {
		return failed(2, err, "U02", helpUsage), true
	}
	return verbatim(page), true
}

func parse(args []string) (command, string) {
	c := command{name: args[0], options: make(map[string]string)}
	rules, known := commands[c.name]
	if !known {
		return c, fmt.Sprintf("Unknown command %q", c.name)
	}
	for i := 1; i < len(args); i++ {
		arg := args[i]
		switch {
		case slices.Contains(strings.Fields(rules.flags), arg):
			if c.has(arg) {
				return c, "Repeated option " + arg
			}
			c.options[arg] = ""
		case slices.Contains(strings.Fields(rules.values), arg):
			if c.has(arg) || i+1 >= len(args) || args[i+1] == "" || strings.HasPrefix(args[i+1], "--") {
				return c, "Provide one value for " + arg
			}
			i++
			c.options[arg] = args[i]
		case strings.HasPrefix(arg, "-"):
			return c, "Unsupported option " + arg
		default:
			c.args = append(c.args, arg)
		}
	}
	if len(c.args) < rules.min || len(c.args) > rules.max || (c.has("--stale") && len(c.args) != 1) || slices.Contains(c.args, "") {
		return c, fmt.Sprintf("Invalid arguments for %s: %q", c.name, c.args)
	}
	switch {
	case c.has("--stale") && (c.has("--history") || c.has("--projected")):
		return c, "--stale cannot be combined with --history or --projected"
	case c.has("--history") && c.has("--projected"):
		return c, "--history cannot be combined with --projected"
	case c.name == "agents-md" && !c.has("--write"):
		return c, "agents-md requires --write"
	case c.name == "init" && (!c.has("--recipe") || !c.has("--agent")):
		return c, "init requires --recipe and --agent"
	case c.name == "new-module" && (!moduleName.MatchString(c.arg(0)) || !modulePrefix.MatchString(c.options["--prefix"])):
		return c, "Provide a valid lowercase module name and uppercase ID prefix with --prefix"
	}
	for _, setting := range []struct {
		flag   string
		target *int
	}{{"--max", &c.max}, {"--from-attempt", &c.attempt}} {
		if c.has(setting.flag) {
			value, err := strconv.Atoi(c.options[setting.flag])
			if err != nil || value <= 0 {
				return c, setting.flag + " requires a positive integer"
			}
			*setting.target = value
		}
	}
	if c.has("--for") {
		value, err := time.ParseDuration(c.options["--for"])
		if err != nil || value <= 0 {
			return c, "--for requires a positive Go duration"
		}
		c.duration = value
	}
	return c, ""
}

func usage(args []string) string {
	switch {
	case slices.Contains(args, "--stale"):
		return "Use fireweed show --stale <change>."
	case slices.Contains(args, "--fix"):
		return "Use fireweed check --fix."
	case slices.Contains(args, "--for"):
		return "Use fireweed run --for <duration>."
	case args[0] == "start":
		return "Use fireweed start <change> [--from-attempt <N>]."
	case args[0] == "init":
		return initUsage
	}
	if _, known := commands[args[0]]; known {
		return fmt.Sprintf("Use fireweed help %s for usage.", args[0])
	}
	return helpUsage
}

func Run(args []string, out, errOut io.Writer) int {
	var tokens []string
	var jsonMode bool
	for _, arg := range args {
		if arg == "--json" {
			jsonMode = true
		} else {
			tokens = append(tokens, arg)
		}
	}
	if result, handled := help(tokens); handled {
		return result.write(out, errOut, jsonMode)
	}
	c, problem := parse(tokens)
	if problem != "" {
		return failure(2, "U01", "", problem+". "+usage(tokens)).write(out, errOut, jsonMode)
	}
	return c.execute().write(out, errOut, jsonMode)
}

func (c command) execute() response {
	switch c.name {
	case "version":
		return verbatim(assets.Version() + "\n")
	case "explain", "playbook":
		lookup := assets.Explain
		if c.name == "playbook" {
			lookup = assets.Playbook
		}
		body, err := lookup(c.arg(0))
		if err != nil {
			return failed(2, err, "U02", "Use fireweed help "+c.name+" for valid inputs.")
		}
		return verbatim(body)
	}
	dir, err := os.Getwd()
	if err != nil {
		return configFailure(err)
	}
	if c.name == "init" {
		return c.initProject(dir)
	}
	if c.name == "modules" || c.name == "coverage" {
		root, approved, err := discoverApproved(dir)
		if err != nil {
			return configFailure(err)
		}
		if approved != nil {
			if err := project.CheckCompatibility(approved.Snapshot.Config); err != nil {
				return approvedConfigFailure(err, approved.Slug)
			}
		}
		if c.name == "modules" {
			return modules(root, approved)
		}
		return c.coverage(root, approved)
	}
	root, err := discover(dir)
	if err != nil {
		return configFailure(err)
	}
	switch c.name {
	case "new-module", "new-change", "stamp", "show", "list", "agents-md", "migrate":
		return c.projectCommand(root)
	case "check", "check-commits", "ci", "doctor":
		return c.checkCommand(root)
	case "next", "run", "drill":
		return c.runCommand(root)
	default:
		return c.workflowCommand(root)
	}
}

func invokingBin() string {
	bin, _ := os.Executable()
	return bin
}

func discover(dir string) (string, error) {
	for root := dir; ; root = filepath.Dir(root) {
		if present, err := exists(filepath.Join(root, "fireweed.yaml")); err != nil || present {
			return root, err
		}
		if filepath.Dir(root) == root {
			return "", fmt.Errorf("fireweed.yaml not found from %s", dir)
		}
	}
}

func exists(name string) (bool, error) {
	_, err := os.Lstat(name)
	if os.IsNotExist(err) {
		return false, nil
	}
	return err == nil, err
}

func background() context.Context { return context.Background() }
