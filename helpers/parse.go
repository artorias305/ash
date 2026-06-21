package helpers

import (
	"fmt"
	"strings"
)

type Redirect struct {
	File string
}

func ParseCliInput(input string) []string {
	var args []string
	var buf strings.Builder
	inQuote := false
	inDoubleQuote := false

	for i := 0; i < len(input); i++ {
		c := input[i]
		if inQuote {

			if c == '\'' {
				inQuote = false
			} else {
				buf.WriteByte(c)
			}
		} else if inDoubleQuote {
			if c == '"' {
				inDoubleQuote = false
			} else if c == '\\' {
				if input[i+1] == '"' || input[i+1] == '$' || input[i+1] == '\\' || input[i+1] == '`' {
					i++
					buf.WriteByte(input[i])
				}
			} else {
				buf.WriteByte(c)
			}
		} else {
			switch c {
			case '\'':
				inQuote = true
			case '"':
				inDoubleQuote = true
			case '\\':
				if i+1 < len(input) {
					i++
					buf.WriteByte(input[i])
				}
			case ' ', '\t', '\n', '\r':
				if buf.Len() > 0 {
					args = append(args, buf.String())
					buf.Reset()
				}
			default:
				buf.WriteByte(c)
			}
		}
	}
	if buf.Len() > 0 {
		args = append(args, buf.String())
	}
	return args
}

func ExtractRedirect(args []string) ([]string, string, error) {
	var cleanArgs []string

	for i := 0; i < len(args); i++ {
		arg := args[i]

		if arg == ">" {
			if i+1 >= len(args) {
				return nil, "", fmt.Errorf("syntax error: expected filename after >")
			}
			return cleanArgs, args[i+1], nil
		}

		if strings.HasPrefix(arg, ">") {
			filename := arg[1:]
			if filename == "" {
				if i+1 >= len(args) {
					return nil, "", fmt.Errorf("syntax error: expected filename after >")
				}
				return cleanArgs, args[i+1], nil
			}
			return cleanArgs, filename, nil
		}

		if idx := strings.Index(arg, ">"); idx != -1 {
			filename := arg[idx+1:]
			cleanArgs = append(cleanArgs, arg[:idx])
			if filename == "" {
				if i+1 >= len(args) {
					return nil, "", fmt.Errorf("syntax error: expected filename after >")
				}
				return cleanArgs, args[i+1], nil
			}
			return cleanArgs, filename, nil

		}
		cleanArgs = append(cleanArgs, arg)

	}
	return cleanArgs, "", nil
}
