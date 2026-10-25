package process

import "os/exec"

func newTree(command *exec.Cmd) (treeControl, error) { return newGroupTree(command) }
