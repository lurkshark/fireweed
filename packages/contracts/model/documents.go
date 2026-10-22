package model

type Spec struct {
	ID   string `yaml:"id" json:"id"`
	Tag  string `yaml:"tag,omitempty" json:"tag,omitempty"`
	Body string `yaml:"body" json:"body"`
}
type Section struct {
	Name string `json:"name"`
	Body string `json:"body"`
}
type Document struct {
	Metadata map[string]any `json:"metadata"`
	Preamble string         `json:"preamble"`
	Sections []Section      `json:"sections"`
}
type Module struct {
	Name             string   `yaml:"name" json:"name"`
	IDPrefix         string   `yaml:"id_prefix" json:"id_prefix"`
	Dir              string   `yaml:"-" json:"dir"`
	DependsOn        []string `yaml:"depends_on" json:"depends_on"`
	Zones            `yaml:",inline"`
	Hooks            Hooks    `yaml:"hooks" json:"hooks"`
	BoundariesOptOut string   `yaml:"boundaries_opt_out,omitempty" json:"boundaries_opt_out,omitempty"`
	Specs            []Spec   `yaml:"-" json:"specs"`
	Decisions        string   `yaml:"-" json:"decisions,omitempty"`
	Document         Document `yaml:"-" json:"document"`
}
type Change struct {
	Module           string            `yaml:"module" json:"module"`
	Set              map[string]any    `yaml:"set,omitempty" json:"set,omitempty"`
	Bases            map[string]string `yaml:"bases,omitempty" json:"bases,omitempty"`
	Fixes            []string          `yaml:"fixes,omitempty" json:"fixes,omitempty"`
	AlreadySatisfied []string          `yaml:"already_satisfied,omitempty" json:"already_satisfied,omitempty"`
	Regenerate       bool              `yaml:"regenerate,omitempty" json:"regenerate,omitempty"`
	Run              string            `yaml:"run,omitempty" json:"run,omitempty"`
	Phase            string            `yaml:"phase,omitempty" json:"phase,omitempty"`
	Blocked          string            `yaml:"blocked,omitempty" json:"blocked,omitempty"`
	Path             string            `yaml:"-" json:"path"`
	Slug             string            `yaml:"-" json:"slug"`
	Number           int               `yaml:"-" json:"number"`
	Kind             string            `yaml:"-" json:"kind"`
	Added            []Spec            `yaml:"-" json:"added"`
	Modified         []Spec            `yaml:"-" json:"modified"`
	Removed          []Spec            `yaml:"-" json:"removed"`
	Document         Document          `yaml:"-" json:"document"`
}
