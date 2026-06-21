package commands

import (
	"fmt"
	"os"
)

func Pwd() {
	pwd := os.Getenv("PWD")
	if pwd == "" {
		fmt.Fprintf(os.Stderr, "$PWD environment variable not set\n")
	} else {
		fmt.Println(pwd)
	}
}
