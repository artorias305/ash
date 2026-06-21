package commands

import (
	"fmt"
	"os"
)

func Cd(path string) {
	if err := os.Chdir(path); err != nil {
		fmt.Printf("%v\n", err)
		return
	}
}
