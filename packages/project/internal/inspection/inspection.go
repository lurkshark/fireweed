package inspection

import (
	"fmt"
	"github.com/lurkshark/fireweed/packages/contracts"
)

type Diagnostic struct{ Code, Severity, Message, Path string }
type Data struct {
	Root, BaseRef    string
	Config           contracts.Config
	Modules          []contracts.Module
	Changes, Archive []contracts.Change
	Diagnostics      []Diagnostic
}

func Load(start string) (Data, error) { return Data{}, fmt.Errorf("discovery not implemented") }
func Check(s Data) []Diagnostic       { return nil }
