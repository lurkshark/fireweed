package main

import (
	"github.com/lurkshark/fireweed/packages/gitstore/internal/supervision"
	"os"
)

func main() { os.Exit(supervision.Run()) }
