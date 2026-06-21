package commands

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func RunCommand(command string, args []string, w io.Writer) {
	switch command {
	case "exit":
		os.Exit(0)
	case "echo":
		Echo(w, args)
	case "type":
		if len(args) != 0 {
			Type(w, args[0])
		}
	case "cd":
		if len(args) == 0 {
			Cd("~")
		} else {
			Cd(args[0])
		}
	default:
		cmd := exec.Command(command, args...)
		cmd.Stdout = w
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
