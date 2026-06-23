package helpers

import (
	"fmt"
	"strconv"
	"strings"
)

type Redirect struct {
	File   string
	Append bool
	FD     int
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

func ExtractRedirect(args []string) ([]string, *Redirect, error) {
	var cleanArgs []string

	for i := 0; i < len(args); i++ {
		arg := args[i]

		opIdx, opLen := findRedirectOp(arg)
		if opIdx != -1 {
			prefix := arg[:opIdx]
			suffix := arg[opIdx+opLen:]

			fd := 1
			if prefix != "" {
				n, err := strconv.Atoi(prefix)
				if err != nil {
					cleanArgs = append(cleanArgs, arg)
					continue
				}
				fd = n
			}

			isAppend := opLen == 2

			if suffix != "" {
				return cleanArgs, &Redirect{FD: fd, File: suffix, Append: isAppend}, nil
			}
			if i+1 >= len(args) {
				return nil, nil, fmt.Errorf("syntax error: expected filename after %s", arg)
			}
			return cleanArgs, &Redirect{FD: fd, File: args[i+1], Append: isAppend}, nil
		}
		cleanArgs = append(cleanArgs, arg)
	}
	return cleanArgs, nil, nil
}

func findRedirectOp(s string) (int, int) {
	idx := strings.Index(s, ">")
	if idx == -1 {
		return -1, 0
	}
	if idx+1 < len(s) && s[idx+1] == '>' {
		return idx, 2
	}
	return idx, 1
}
