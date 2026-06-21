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

		cleanArgs, redirectFile, err := helpers.ExtractRedirect(args)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		cmd := cleanArgs[0]
		cleanArgs = cleanArgs[1:]

		w := io.Writer(os.Stdout)
		if redirectFile != "" {
			file, err := os.Create(redirectFile)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %v\n", err)
				continue
			}
			w = file
			defer file.Close()
		}

		commands.RunCommand(cmd, cleanArgs, w)
	}
}
