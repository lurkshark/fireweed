package entry

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/lurkshark/fireweed/packages/assets"
	"github.com/lurkshark/fireweed/packages/project"
)

const helpText = `fireweed: checked spec-driven development and regeneration

Commands:
  help
  version
  init --recipe <go|typescript-pnpm|python-uv|none|path> --agent <agent> [--audit-agent <agent>] [--no-app]
  doctor
  new-module <name> --prefix <P>
  new-change <module> <slug>
  stamp <change>
  list
  status
  modules
  show <module> [ID] [--projected|--history]
  context <change>
  explain <code>
  playbook <name>
  check
  coverage <module> [--with-change <change>]
  check-commits <range>
  ci [<range>] [--no-doctor]
  pr-body <range>
  start <change> [--from-attempt <n>]
  verify [<change>]
  done [<change>]
  next [--allow-unmerged] [--for <duration>]
  run [--max N] [--for <duration>] [--allow-unmerged]
  release <change>
  commits <change>
  drill <module>
  gc
  migrate
  agents-md --write

Use --json before or after a command for a versioned JSON result.
Use -- to stop option parsing. Commands never prompt.
Exit statuses:
  0 success or empty queue
  1 failed checks or blocked changes
  2 usage or configuration errors
  3 time or usage limit
`

type options struct {
	command            string
	operands           []string
	json, help         bool
	projected, history bool
	problem            string
	withChange         string
	withChangeSet      bool
}

func parse(args []string) options {
	result := options{command: "help"}
	positional := []string{}
	stopped := false
	for i := 0; i < len(args); i++ {
		arg := args[i]
		if stopped {
			positional = append(positional, arg)
			continue
		}
		switch arg {
		case "--":
			stopped = true
		case "--with-change":
			result.withChangeSet = true
			if i+1 >= len(args) || strings.TrimSpace(args[i+1]) == "" || strings.HasPrefix(args[i+1], "-") {
				if result.problem == "" {
					result.problem = "--with-change requires a nonempty queued change slug"
				}
			} else {
				i++
				result.withChange = args[i]
			}
		case "--projected":
			result.projected = true
		case "--history":
			result.history = true
		case "--json":
			result.json = true
		case "--help", "-h":
			result.help = true
		case "--version", "-v":
			positional = append(positional, "version")
		default:
			if value, matched := strings.CutPrefix(arg, "--with-change="); matched {
				result.withChangeSet = true
				result.withChange = value
				if strings.TrimSpace(value) == "" && result.problem == "" {
					result.problem = "--with-change requires a nonempty queued change slug"
				}
			} else if strings.HasPrefix(arg, "-") {
				if result.problem == "" {
					result.problem = fmt.Sprintf("unknown or valued option %q; valid options: --json, --help/-h, --version/-v; show also accepts --projected or --history; coverage accepts --with-change <slug>", arg)
				}
			} else {
				positional = append(positional, arg)
			}
		}
	}
	if len(positional) > 0 {
		result.command = positional[0]
		result.operands = positional[1:]
	}
	return result
}

type diagnostic struct {
	Code     string `json:"code"`
	Severity string `json:"severity"`
	Message  string `json:"message"`
	Path     string `json:"path,omitempty"`
}
type envelope struct {
	SchemaVersion int          `json:"schema_version"`
	Command       string       `json:"command"`
	Data          any          `json:"data"`
	Diagnostics   []diagnostic `json:"diagnostics"`
}

func write(out io.Writer, text string) bool {
	n, err := io.WriteString(out, text)
	return err == nil && n == len(text)
}
func emit(out, errOut io.Writer, option options, data any, diagnostics []diagnostic, code int) int {
	if diagnostics == nil {
		diagnostics = []diagnostic{}
	}
	if option.json {
		data, err := json.MarshalIndent(envelope{1, option.command, data, diagnostics}, "", "  ")
		if err != nil || !write(out, string(data)+"\n") {
			return 2
		}
		return code
	}
	if data != nil {
		text, ok := data.(string)
		if !ok {
			encoded, err := json.MarshalIndent(data, "", "  ")
			if err != nil {
				return 2
			}
			text = string(encoded) + "\n"
		}
		if !write(out, text) {
			return 2
		}
	}
	for _, d := range diagnostics {
		where := ""
		if d.Path != "" {
			where = " " + d.Path
		}
		if !write(errOut, fmt.Sprintf("%s %s%s: %s\n", d.Severity, d.Code, where, d.Message)) {
			return 2
		}
	}
	return code
}
func Run(args []string, out, errOut io.Writer) int {
	option := parse(args)
	fail := func(problem string) int {
		message := problem + ". Run fireweed help. Valid commands: help, version, explain, playbook, modules, check, show, coverage, context, start, verify."
		return emit(out, errOut, option, nil, []diagnostic{{Code: "U01", Severity: "error", Message: message}}, 2)
	}
	if option.problem != "" {
		return fail(option.problem)
	}
	switch option.command {
	case "help", "version", "explain", "playbook", "modules", "check", "show", "coverage", "context", "start", "verify":
	default:
		return fail("unknown command " + option.command)
	}
	if option.withChangeSet && option.command != "coverage" {
		return fail("--with-change is only valid with coverage")
	}
	if (option.projected || option.history) && option.command != "show" {
		return fail("--projected and --history are only valid with show")
	}
	if option.projected && option.history {
		return fail("choose only one of --projected or --history")
	}
	if option.help && option.command != "help" {
		option.command = "help"
		option.operands = nil
	}
	required := 0
	if option.command == "explain" || option.command == "playbook" || option.command == "coverage" || option.command == "context" || option.command == "start" {
		required = 1
	}
	if option.command == "verify" {
		if len(option.operands) > 1 {
			return fail("verify accepts zero or one exact change slug")
		}
	} else if option.command == "show" {
		if len(option.operands) < 1 || len(option.operands) > 2 {
			return fail("show requires a module and optional exact spec ID")
		}
	} else if len(option.operands) != required {
		return fail(fmt.Sprintf("%s requires %d positional arguments", option.command, required))
	}
	switch option.command {
	case "verify":
		return verify(out, errOut, option)
	case "start":
		return start(out, errOut, option)
	case "context":
		return workerContext(out, errOut, option)
	case "coverage":
		return coverage(out, errOut, option, fail)
	case "show":
		return show(out, errOut, option, fail)
	case "modules", "check":
		snapshot, err := project.Load(".")
		if err != nil {
			return loadFailure(out, errOut, option, err)
		}
		diagnostics, code := checkedDiagnostics(snapshot, option.command == "check")
		var data any = snapshot.Modules
		if option.command == "check" {
			data = map[string]bool{"valid": code == 0}
		}
		return emit(out, errOut, option, data, diagnostics, code)
	case "help":
		return emit(out, errOut, option, helpText, nil, 0)
	case "version":
		return emit(out, errOut, option, assets.Info().Version+plainNewline(option.json), nil, 0)
	case "explain":
		code := option.operands[0]
		text, err := assets.Explain(code)
		if err != nil {
			return fail(fmt.Sprintf("unknown diagnostic %q; valid choices: %s", code, strings.Join(assets.Codes(), ", ")))
		}
		return emit(out, errOut, option, map[string]string{"code": code, "explanation": text}, nil, 0)
	case "playbook":
		return playbook(out, errOut, option, fail)
	}
	return fail("unhandled command")
}
func plainNewline(json bool) string {
	if json {
		return ""
	}
	return "\n"
}
