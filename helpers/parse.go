package helpers

import "strings"

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
			} else {
				buf.WriteByte(c)
			}
		} else {
			switch c {
			case '\'':
				inQuote = true
			case '"':
				inDoubleQuote = true
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
