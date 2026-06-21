package commands

import (
	"fmt"
	"io"
	"os"
)

func Pwd(w io.Writer) {
	pwd := os.Getenv("PWD")
	if pwd == "" {
		fmt.Fprintf(os.Stderr, "$PWD environment variable not set\n")
	} else {
		fmt.Fprintln(w, pwd)
	}
}
