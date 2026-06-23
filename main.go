package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/artorias305/ash/commands"
	"github.com/artorias305/ash/helpers"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("$ ")

		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Fprintf(os.Stderr, "error: %v\n", err)
			}
			break
		}

		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}

		args := helpers.ParseCliInput(line)
		if len(args) == 0 {
			continue
		}

		cleanArgs, redirect, err := helpers.ExtractRedirect(args)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		cmd := cleanArgs[0]
		cleanArgs = cleanArgs[1:]

		stdout := io.Writer(os.Stdout)
		stderr := io.Writer(os.Stderr)

		if redirect != nil {
			var file *os.File
			if redirect.Append {
				file, err = os.OpenFile(redirect.File, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			} else {
				file, err = os.Create(redirect.File)
			}
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %v\n", err)
				continue
			}
			defer file.Close()

			switch redirect.FD {
			case 1:
				stdout = file
			case 2:
				stderr = file
			default:
				fmt.Fprintf(os.Stderr, "error: unsupported file descriptor %d\n", redirect.FD)
				continue
			}
		}

		commands.RunCommand(cmd, cleanArgs, stdout, stderr)
	}
}
