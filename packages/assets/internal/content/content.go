package content

import "errors"

func Info() (string, int, int)                     { return "", -1, -1 }
func Playbooks() []string                          { return nil }
func Playbook(name string) (string, string, error) { return "", "", errors.New("not implemented") }
func Codes() []string                              { return nil }
func Explain(code string) (string, error)          { return "", errors.New("not implemented") }
