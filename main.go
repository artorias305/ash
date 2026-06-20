package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

		if line == "exit" {
			break
		}

		fmt.Printf("%s: command not found\n", line)
	}
}
