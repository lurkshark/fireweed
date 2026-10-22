package evals

import (
	c "github.com/lurkshark/fireweed/packages/contracts"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func Test_CON008ConfigDefaults(t *testing.T) {
	got, err := c.ParseConfig("{}", "0.4.0")
	if err != nil {
		t.Fatal(err)
	}
	want := c.Config{Fireweed: ">=0.4, <0.5", Version: 1, Remote: "origin", Modules: []string{"packages/*"}, ModuleDefaults: c.Zones{Interface: []string{"src/index.ts"}, Evals: []string{"evals/**"}, Disposable: []string{"src/**"}}, Audit: c.Audit{Enabled: true}, Attempts: 2, Timeout: "90m", VerifyDependents: "all", LimitPatterns: []string{"usage limit", "rate limit"}, LimitWait: "30m", LimitBudget: "4h"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("defaults %+v, want %+v", got, want)
	}
	empty, err := c.ParseConfig("modules: []\nmodule_defaults:\n  interface: []\n  evals: []\n  disposable: []\n  preserve: []\n  append_only: []\nprotected: []\nworker_writable: []\nlimit_patterns: []\n", "0.4.1")
	if err != nil {
		t.Fatal(err)
	}
	for name, list := range map[string][]string{"modules": empty.Modules, "interface": empty.ModuleDefaults.Interface, "evals": empty.ModuleDefaults.Evals, "disposable": empty.ModuleDefaults.Disposable, "preserve": empty.ModuleDefaults.Preserve, "append_only": empty.ModuleDefaults.AppendOnly, "protected": empty.Protected, "worker_writable": empty.WorkerWritable, "limit_patterns": empty.LimitPatterns} {
		if list == nil || len(list) != 0 {
			t.Errorf("explicit %s [] not retained: %#v", name, list)
		}
	}
}

func Test_CON009ConfigFieldsAndIsolation(t *testing.T) {
	marker := filepath.Join(t.TempDir(), "must-not-execute")
	text := `fireweed: '>=0.4.1, <=0.4.9'
version: 0
base: trunk
remote: upstream
modules: ['components/**']
module_defaults:
  interface: [api.go]
  evals: ['*_eval_test.go']
  disposable: ['internal/**']
  preserve: ['kept/**']
  append_only: ['migrations/**']
hooks:
  eval: 'echo {module_dir} {report} && touch ` + marker + `'
  boundaries: 'checker {module}'
  setup: 'echo $HOME'
protected: ['ci/**']
worker_writable: [go.sum]
agent:
  command: [worker, '-p', '{prompt}']
  wrap: [sandbox, '{git_common_dir}']
audit:
  enabled: false
  command: [auditor, '{prompt}']
attempts: 5
timeout: 2m
verify_dependents: direct
limit_patterns: [over-quota]
limit_wait: 1ms
limit_budget: 3s
`
	got, err := c.ParseConfig(text, "0.4.5")
	if err != nil {
		t.Fatal(err)
	}
	want := c.Config{Fireweed: ">=0.4.1, <=0.4.9", Version: 0, Base: "trunk", Remote: "upstream", Modules: []string{"components/**"}, ModuleDefaults: c.Zones{Interface: []string{"api.go"}, Evals: []string{"*_eval_test.go"}, Disposable: []string{"internal/**"}, Preserve: []string{"kept/**"}, AppendOnly: []string{"migrations/**"}}, Hooks: c.Hooks{Eval: "echo {module_dir} {report} && touch " + marker, Boundaries: "checker {module}", Setup: "echo $HOME"}, Protected: []string{"ci/**"}, WorkerWritable: []string{"go.sum"}, Agent: c.Agent{Command: []string{"worker", "-p", "{prompt}"}, Wrap: []string{"sandbox", "{git_common_dir}"}}, Audit: c.Audit{Enabled: false, Command: []string{"auditor", "{prompt}"}}, Attempts: 5, Timeout: "2m", VerifyDependents: "direct", LimitPatterns: []string{"over-quota"}, LimitWait: "1ms", LimitBudget: "3s"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("fields %+v, want %+v", got, want)
	}
	if _, err := os.Stat(marker); !os.IsNotExist(err) {
		t.Errorf("parse executed hook: %v", err)
	}
	first, err := c.ParseConfig("{}", "0.4.0")
	if err != nil {
		t.Fatal(err)
	}
	first.Modules[0] = "mutated"
	first.ModuleDefaults.Interface[0] = "mutated"
	first.ModuleDefaults.Evals[0] = "mutated"
	first.ModuleDefaults.Disposable[0] = "mutated"
	first.LimitPatterns[0] = "mutated"
	second, err := c.ParseConfig("{}", "0.4.0")
	if err != nil {
		t.Fatal(err)
	}
	if second.Modules[0] != "packages/*" || second.ModuleDefaults.Interface[0] != "src/index.ts" || second.ModuleDefaults.Evals[0] != "evals/**" || second.ModuleDefaults.Disposable[0] != "src/**" || second.LimitPatterns[0] != "usage limit" {
		t.Errorf("parse results share slices: %+v", second)
	}
	for _, bad := range []string{"", "version: [", "version: 1\nversion: 0", "{}\n---\n{}", "[a,b]", "scalar", "attempts: many", "audit: false", "modules: packages/*", "module_defaults: {evals: nope}", "hooks: {eval: 123}", "agent: {command: [123]}", "unknown_field: true"} {
		if _, err := c.ParseConfig(bad, "0.4.0"); err == nil {
			t.Errorf("invalid YAML/type accepted %q", bad)
		}
	}
}

func Test_CON010ConfigVersionGate(t *testing.T) {
	for _, version := range []string{"0", "1"} {
		if _, err := c.ParseConfig("version: "+version, "0.4.0"); err != nil {
			t.Errorf("supported format %s: %v", version, err)
		}
	}
	for _, version := range []string{"-1", "2", "99"} {
		if _, err := c.ParseConfig("version: "+version, "0.4.0"); err == nil {
			t.Errorf("unsupported format %s", version)
		}
	}
	for _, tc := range []struct {
		version, pin string
		ok           bool
	}{
		{"0.4.0", ">=0.4, <0.5", true}, {"v0.4.7", ">=0.4 <0.5", true},
		{"0.3.9", ">=0.4, <0.5", false}, {"0.5.0", ">=0.4, <0.5", false},
		{"0.4.0", ">0.4.0", false}, {"0.4.1", ">0.4.0", true},
		{"0.4.1", "<=0.4.1", true}, {"0.4.2", "<=0.4.1", false},
		{"0.4.0", "=0.4", true}, {"0.4.1", "=0.4", false}, {"0.4.1", "0.4.1", true},
		{"0.4.1", "", false}, {"0.4.1", "banana", false}, {"nope", ">=0.4", false},
		{"0.4.1", ">=0.4 || <0.5", false}, {"0.4.1", "^0.4", false},
		{"0.4.1-alpha", ">=0.4", false}, {"0.4.1", ">=0.4-beta", false},
	} {
		err := c.CheckVersion(tc.version, tc.pin)
		if (err == nil) != tc.ok {
			t.Errorf("%s vs %q: %v, want success %v", tc.version, tc.pin, err, tc.ok)
		}
	}
	for _, tc := range []struct{ text, version string }{
		{"fireweed: '>=0.4, <0.5'", "0.5.0"}, {"fireweed: ''", "0.4.0"}, {"{}", "malformed"}, {"fireweed: 'nonsense'", "0.4.0"},
	} {
		if _, err := c.ParseConfig(tc.text, tc.version); err == nil {
			t.Errorf("ParseConfig bypassed version check: %+v", tc)
		}
	}
}

func Test_CON011ConfigValidation(t *testing.T) {
	for _, bad := range []string{"attempts: 0", "attempts: -1", "timeout: 0s", "timeout: -1s", "timeout: nonsense", "limit_wait: -1s", "limit_wait: nonsense", "limit_budget: -1s", "limit_budget: nonsense", "verify_dependents: invalid"} {
		if _, err := c.ParseConfig(bad, "0.4.0"); err == nil {
			t.Errorf("invalid config accepted %q", bad)
		}
	}
	for _, mode := range []string{"all", "direct", "none"} {
		got, err := c.ParseConfig("verify_dependents: "+mode+"\nlimit_wait: 0s\nlimit_budget: 0s\naudit: {enabled: false}", "0.4.0")
		if err != nil || got.VerifyDependents != mode || got.Audit.Enabled {
			t.Errorf("valid mode/zero limits/disabled audit: %+v %v", got, err)
		}
	}
	if got, err := c.ParseConfig("audit: {}", "0.4.0"); err != nil || !got.Audit.Enabled {
		t.Errorf("omitted audit enabled: %+v %v", got, err)
	}
}

func Test_CON008EmptyCommands(t *testing.T) {
	got, err := c.ParseConfig("agent: {command: [], wrap: []}\naudit: {command: []}", "0.4.0")
	if err != nil {
		t.Fatal(err)
	}
	for name, list := range map[string][]string{"agent.command": got.Agent.Command, "agent.wrap": got.Agent.Wrap, "audit.command": got.Audit.Command} {
		if list == nil || len(list) != 0 {
			t.Errorf("explicit empty %s not retained: %#v", name, list)
		}
	}
}

func Test_CON009CustomIndependenceAndNull(t *testing.T) {
	text := "module_defaults: {preserve: [kept], append_only: [added]}\nprotected: [protected]\nworker_writable: [writable]\nagent: {command: [worker], wrap: [wrapper]}\naudit: {command: [auditor]}"
	first, err := c.ParseConfig(text, "0.4.0")
	if err != nil {
		t.Fatal(err)
	}
	second, err := c.ParseConfig(text, "0.4.0")
	if err != nil {
		t.Fatal(err)
	}
	first.ModuleDefaults.Preserve[0] = "mutated"
	first.ModuleDefaults.AppendOnly[0] = "mutated"
	first.Protected[0] = "mutated"
	first.WorkerWritable[0] = "mutated"
	first.Agent.Command[0] = "mutated"
	first.Agent.Wrap[0] = "mutated"
	first.Audit.Command[0] = "mutated"
	third, err := c.ParseConfig(text, "0.4.0")
	if err != nil {
		t.Fatal(err)
	}
	for _, got := range []c.Config{second, third} {
		if got.ModuleDefaults.Preserve[0] != "kept" || got.ModuleDefaults.AppendOnly[0] != "added" || got.Protected[0] != "protected" || got.WorkerWritable[0] != "writable" || got.Agent.Command[0] != "worker" || got.Agent.Wrap[0] != "wrapper" || got.Audit.Command[0] != "auditor" {
			t.Errorf("shared parse values: %+v", got)
		}
	}
	fresh, err := c.ParseConfig("{}", "0.4.0")
	if err != nil {
		t.Fatal(err)
	}
	if len(fresh.ModuleDefaults.Preserve)+len(fresh.ModuleDefaults.AppendOnly)+len(fresh.Protected)+len(fresh.WorkerWritable)+len(fresh.Agent.Command)+len(fresh.Agent.Wrap)+len(fresh.Audit.Command) != 0 {
		t.Errorf("custom values leaked to defaults: %+v", fresh)
	}
	for _, bad := range []string{"null", "~", "---\n"} {
		if _, err := c.ParseConfig(bad, "0.4.0"); err == nil {
			t.Errorf("null root accepted %q", bad)
		}
	}
}

func Test_CON010VersionRegressions(t *testing.T) {
	for _, tc := range []struct {
		version, pin string
		ok           bool
	}{
		{"0.3.9", ">=0.4 <0.5", false}, {"0.5.0", ">=0.4 <0.5", false},
		{"0.4.0", "0.4", true}, {"0.4.1", "0.4", false}, {"0.4.2", "0.4.1", false},
		{"0.4.10", ">0.4.9", true}, {"0.4.9", ">=0.4.10", false},
		{"0.4.0+build", ">=0.4", false}, {"0.4.0", ">=0.4.0+build", false},
		{"0.-4.0", ">=0.4", false}, {"0.4.0", ">=0.-4.0", false},
	} {
		if err := c.CheckVersion(tc.version, tc.pin); (err == nil) != tc.ok {
			t.Errorf("%s vs %q = %v, want success %v", tc.version, tc.pin, err, tc.ok)
		}
	}
}
