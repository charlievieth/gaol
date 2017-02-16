// +build !windows

package main

import "github.com/contraband/gaol/commands"

func init() {
	cmds = append(cmds, command{"shell", "open a shell in the container", &commands.Shell{}})
}
