package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/artorias305/ash/commands"
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
		args := strings.Fields(line)

		cmd := args[0]
		args = args[1:]

		if cmd == "exit" {
			break
		} else if cmd == "echo" {
			commands.Echo(args)
		} else {
			fmt.Printf("%s: command not found\n", line)
		}
	}
}
