// Package model holds retained wire data shared by disposable implementations.
package model

type Zones struct {
	Interface  []string `yaml:"interface" json:"interface"`
	Evals      []string `yaml:"evals" json:"evals"`
	Disposable []string `yaml:"disposable" json:"disposable"`
	Preserve   []string `yaml:"preserve" json:"preserve"`
	AppendOnly []string `yaml:"append_only" json:"append_only"`
}
type Hooks struct {
	Eval       string `yaml:"eval" json:"eval"`
	Boundaries string `yaml:"boundaries" json:"boundaries"`
	Setup      string `yaml:"setup" json:"setup"`
}
type Agent struct {
	Command []string `yaml:"command" json:"command"`
	Wrap    []string `yaml:"wrap" json:"wrap"`
}
type Audit struct {
	Enabled bool     `yaml:"enabled" json:"enabled"`
	Command []string `yaml:"command" json:"command"`
}
type Config struct {
	Fireweed         string   `yaml:"fireweed" json:"fireweed"`
	Version          int      `yaml:"version" json:"version"`
	Base             string   `yaml:"base" json:"base"`
	Remote           string   `yaml:"remote" json:"remote"`
	Modules          []string `yaml:"modules" json:"modules"`
	ModuleDefaults   Zones    `yaml:"module_defaults" json:"module_defaults"`
	Hooks            Hooks    `yaml:"hooks" json:"hooks"`
	Protected        []string `yaml:"protected" json:"protected"`
	WorkerWritable   []string `yaml:"worker_writable" json:"worker_writable"`
	Agent            Agent    `yaml:"agent" json:"agent"`
	Audit            Audit    `yaml:"audit" json:"audit"`
	Attempts         int      `yaml:"attempts" json:"attempts"`
	Timeout          string   `yaml:"timeout" json:"timeout"`
	VerifyDependents string   `yaml:"verify_dependents" json:"verify_dependents"`
	LimitPatterns    []string `yaml:"limit_patterns" json:"limit_patterns"`
	LimitWait        string   `yaml:"limit_wait" json:"limit_wait"`
	LimitBudget      string   `yaml:"limit_budget" json:"limit_budget"`
}
