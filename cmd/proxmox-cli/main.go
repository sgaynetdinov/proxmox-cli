package main

import (
	"proxmox-cli/internal/cli"
)

var version = "dev"
var gitCommit = "none"

func main() {
	cli.Execute(version, gitCommit)
}
