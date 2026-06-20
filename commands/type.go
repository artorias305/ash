package commands

import (
	"fmt"
	"slices"
)

var builtin_commands = []string{"echo", "type", "exit"}

func Type(command string) {
	if slices.Contains(builtin_commands, command) {
		fmt.Printf("%s is a shell builtin\n", command)
	} else {
		fmt.Printf("%s: command not found\n", command)
	}
}
