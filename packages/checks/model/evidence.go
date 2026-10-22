// Package model contains retained evidence data shared by checks implementations.
package model

type Message struct {
	Message string `json:"message,omitempty"`
	Text    string `json:"text,omitempty"`
}
type TestCase struct {
	Name      string    `json:"name"`
	Classname string    `json:"classname"`
	Failures  []Message `json:"failures,omitempty"`
	Errors    []Message `json:"errors,omitempty"`
	Skipped   *Message  `json:"skipped,omitempty"`
	IDs       []string  `json:"spec_ids"`
}
type Diagnostic struct {
	Code     string `json:"code"`
	Severity string `json:"severity"`
	Message  string `json:"message"`
	Path     string `json:"path,omitempty"`
}
type Coverage struct {
	Module      string           `json:"module"`
	Tests       []TestCase       `json:"tests"`
	BySpec      map[string][]int `json:"by_spec"`
	Diagnostics []Diagnostic     `json:"diagnostics"`
	Output      string           `json:"output,omitempty"`
	HookExit    int              `json:"hook_exit"`
}
