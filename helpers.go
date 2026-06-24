package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/artorias305/ash/commands"
	"github.com/artorias305/ash/commands/builtin"
	"golang.org/x/term"
)


func AutoComplete(input string) string {
	candidates := builtin.BuiltinCommands
	candidates = append(candidates, commands.ListPathExecutables()...)
	for _, c := range candidates {
		if strings.HasPrefix(c, input) {
			return c[len(input):]
		}
	}
	return ""
}

func ReadLine(prompt string) (string, error) {
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
				suffix := AutoComplete(buf.String())
				if suffix != "" {
					os.Stdout.Write([]byte(suffix))
					buf.WriteString(suffix)
					os.Stdout.Write([]byte{' '})
					buf.WriteByte(' ')

				} else {
					os.Stdout.Write([]byte{'\x07'})
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