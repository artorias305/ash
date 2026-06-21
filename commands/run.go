package commands

import (
	"fmt"
	"os"
	"os/exec"
)

func RunCommand(command string, args []string) {
	switch command {
	case "exit":
		os.Exit(0)
	case "echo":
		Echo(args)
	case "type":
		if len(args) != 0 {
			Type(args[0])
		}
	case "cd":
		Cd(args[0])
	default:
		cmd := exec.Command(command, args...)
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		fmt.Print(string(output))
	}
}
