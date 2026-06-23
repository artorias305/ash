package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/artorias305/ash/commands"
	"github.com/artorias305/ash/helpers"

	"golang.org/x/term"
)

func autoComplete(input string) string {
	candidates := []string{"echo", "exit"}
	for _, c := range candidates {
		if strings.HasPrefix(c, input) {
			return c[len(input):]
		}
	}
	return ""
}

func readLine(prompt string) (string, error) {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return "", err
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	fmt.Print(prompt)

	var buf strings.Builder
	for {
		var b [1]byte
		n, err := os.Stdin.Read(b[:])
		if n > 0 {
			switch b[0] {
			case 9: // TAB
				suffix := autoComplete(buf.String())
				if suffix != "" {
					os.Stdout.Write([]byte(suffix))
					buf.WriteString(suffix)
					os.Stdout.Write([]byte{' '})
					buf.WriteByte(' ')

				}

			case 13: // Enter (CR in raw mode)
				os.Stdout.Write([]byte("\r\n"))
				return buf.String(), nil

			case 127: // Backspace
				if buf.Len() > 0 {
					s := buf.String()
					buf.Reset()
					buf.WriteString(s[:len(s)-1])
					os.Stdout.Write([]byte("\b \b"))
				}

			case 3: // Ctrl-C
				os.Stdout.Write([]byte("^C\r\n"))
				return "", nil

			case 4: // Ctrl-D (EOF)
				os.Stdout.Write([]byte("\r\n"))
				return "", io.EOF

			default:
				if b[0] >= 32 { // printable characters
					buf.WriteByte(b[0])
					os.Stdout.Write(b[:])
				}
			}
		}
		if err != nil {
			return buf.String(), err
		}
	}
}

func main() {
	for {
		line, err := readLine("$ ")

		if err == io.EOF {
			break
		}

		line = strings.TrimSpace(line)

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
