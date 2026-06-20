package commands

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/artorias305/ash/commands/builtin"
)

func Type(command string) {
	if builtin.IsBuiltin(command) {
		fmt.Printf("%s is a shell builtin\n", command)
	} else {
		path, found := scanPathForCommand(command)
		if found {
			fmt.Printf("%s is %s\n", command, path)
		} else {
			fmt.Printf("%s: not found\n", command)
		}
	}
}

func scanPathForCommand(command string) (string, bool) {
	pathEnv := os.Getenv("PATH")
	if pathEnv == "" {
		fmt.Println("The PATH environment variable is empty.")
		return "", false
	}

	paths := filepath.SplitList(pathEnv)

	for _, dir := range paths {
		// skip empty entries
		if dir == "" {
			continue
		}

		entries, err := os.ReadDir(dir)
		if err != nil {
			// fmt.Printf("could not read directory: %v\n\n", err)
			continue
		}

		if len(entries) == 0 {
			continue
		}
		for _, entry := range entries {
			if !entry.IsDir() {
				if entry.Name() == command {
					return dir + "/" + entry.Name(), true
				}
			}
		}
	}
	return "", false
}
