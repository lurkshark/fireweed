// Package assets exposes embedded protocol data; it has no runtime source dependency.
package assets

import "github.com/lurkshark/fireweed/packages/assets/internal/content"

type BuildInfo struct {
	Version                       string
	CurrentFormat, PreviousFormat int
}
type TextAsset struct{ Name, Text, Hash string }

func Info() BuildInfo {
	version, current, previous := content.Info()
	return BuildInfo{version, current, previous}
}
func Playbooks() []string { return content.Playbooks() }
func Playbook(name string) (TextAsset, error) {
	text, hash, err := content.Playbook(name)
	return TextAsset{Name: name, Text: text, Hash: hash}, err
}
func Codes() []string                     { return content.Codes() }
func Explain(code string) (string, error) { return content.Explain(code) }
