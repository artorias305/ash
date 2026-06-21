package commands

import (
	"fmt"

	"github.com/artorias305/ash/commands/builtin"
)

func Type(command string) {
	if builtin.IsBuiltin(command) {
		fmt.Printf("%s is a shell builtin\n", command)
	} else {
		path, found := scanPathForCommand(command)
		if found {
			fmt.Printf("%s is %s\n", command, path)
		} else {
			fmt.Printf("%s: not found\n", command)
		}
	}
}
