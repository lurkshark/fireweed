package entry

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/lurkshark/fireweed/packages/assets"
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
	command    string
	operands   []string
	json, help bool
	problem    string
}

func parse(args []string) options {
	result := options{command: "help"}
	positional := []string{}
	stopped := false
	for _, arg := range args {
		if stopped {
			positional = append(positional, arg)
			continue
		}
		switch arg {
		case "--":
			stopped = true
		case "--json":
			result.json = true
		case "--help", "-h":
			result.help = true
		case "--version", "-v":
			positional = append(positional, "version")
		default:
			if strings.HasPrefix(arg, "-") {
				if result.problem == "" {
					result.problem = fmt.Sprintf("unknown or valued option %q; valid options: --json, --help/-h, --version/-v", arg)
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
		if !write(errOut, fmt.Sprintf("%s %s: %s\n", d.Severity, d.Code, d.Message)) {
			return 2
		}
	}
	return code
}
func Run(args []string, out, errOut io.Writer) int {
	option := parse(args)
	fail := func(problem string) int {
		message := problem + ". Run fireweed help. Valid commands: help, version, explain, playbook."
		return emit(out, errOut, option, nil, []diagnostic{{Code: "U01", Severity: "error", Message: message}}, 2)
	}
	if option.problem != "" {
		return fail(option.problem)
	}
	switch option.command {
	case "help", "version", "explain", "playbook":
	default:
		return fail("unknown command " + option.command)
	}
	if option.help && option.command != "help" {
		option.command = "help"
		option.operands = nil
	}
	required := 0
	if option.command == "explain" || option.command == "playbook" {
		required = 1
	}
	if len(option.operands) != required {
		return fail(fmt.Sprintf("%s requires %d positional arguments", option.command, required))
	}
	switch option.command {
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
		name := option.operands[0]
		asset, err := assets.Playbook(name)
		if err != nil {
			return fail(fmt.Sprintf("unknown playbook %q; valid choices: %s", name, strings.Join(assets.Playbooks(), ", ")))
		}
		if option.json {
			return emit(out, errOut, option, map[string]any{"name": name, "text": asset.Text, "hash": asset.Hash, "embedded_hash": asset.Hash, "overridden": false}, nil, 0)
		}
		return emit(out, errOut, option, asset.Text, nil, 0)
	}
	return fail("unhandled command")
}
func plainNewline(json bool) string {
	if json {
		return ""
	}
	return "\n"
}
