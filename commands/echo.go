package commands

import (
	"fmt"
	"io"
)

func Echo(w io.Writer, message []string) {
	for _, m := range message {
		fmt.Fprintf(w, "%s ", m)
	}
	fmt.Fprintln(w)
}
