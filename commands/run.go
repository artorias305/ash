package commands

import (
	"fmt"
	"os"
)

func RunCommand(command string, args []string) {
	switch command {
	case "exit":
	os.Exit(1)
	case "echo":
	Echo(args)
	case "type":
	Type(args[0])
	default:
		fmt.Printf("%s: command not found\n", command)
	}
}
