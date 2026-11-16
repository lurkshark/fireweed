// This executable supplies the module graph to the Go recipe's boundaries hook.
package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/lurkshark/fireweed/packages/project"
)

type zones struct {
	Interface []string `json:"interface"`
	Evals     []string `json:"evals"`
}

type module struct {
	Name      string   `json:"name"`
	Dir       string   `json:"dir"`
	DependsOn []string `json:"depends_on"`
	Zones     zones    `json:"zones"`
}

func main() {
	if len(os.Args) != 3 || os.Args[1] != "modules" || os.Args[2] != "--json" {
		fmt.Fprintln(os.Stderr, "usage: doctor-modules modules --json")
		os.Exit(2)
	}
	if err := writeModules(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func writeModules() error {
	// The hook runs from the scratch project, including its canary modules.
	snapshot, err := project.Load(".")
	if err != nil {
		return err
	}
	modules := make([]module, 0, len(snapshot.Modules))
	for _, m := range snapshot.Modules {
		modules = append(modules, module{
			Name:      m.Name,
			Dir:       m.Dir,
			DependsOn: m.DependsOn,
			Zones:     zones{Interface: m.Zones.Interface, Evals: m.Zones.Evals},
		})
	}
	return json.NewEncoder(os.Stdout).Encode(struct {
		Data []module `json:"data"`
	}{Data: modules})
}
