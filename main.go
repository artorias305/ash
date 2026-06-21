package main

import (
	"bufio"
	"fmt"
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

		cmd := args[0]
		args = args[1:]

		commands.RunCommand(cmd, args)
	}
}
