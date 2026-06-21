package commands

import (
	"fmt"
	"io"

	"github.com/artorias305/ash/commands/builtin"
)

func Type(w io.Writer, command string) {
	if builtin.IsBuiltin(command) {
		fmt.Fprintf(w, "%s is a shell builtin\n", command)
	} else {
		path, found := scanPathForCommand(command)
		if found {
			fmt.Fprintf(w, "%s is %s\n", command, path)
		} else {
			fmt.Fprintf(w, "%s: not found\n", command)
		}
	}
}
