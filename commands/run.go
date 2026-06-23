package commands

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func RunCommand(command string, args []string, stdout, stderr io.Writer) {
	switch command {
	case "exit":
		os.Exit(0)
	case "echo":
		Echo(stdout, args)
	case "type":
		if len(args) != 0 {
			Type(stdout, args[0])
		}
	case "cd":
		if len(args) == 0 {
			Cd("~")
		} else {
			Cd(args[0])
		}
	default:
		cmd := exec.Command(command, args...)
		cmd.Stdout = stdout
		cmd.Stderr = stderr
		if err := cmd.Run(); err != nil {
			fmt.Fprintln(stderr, err)
		}
	}
}
