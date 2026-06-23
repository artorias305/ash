package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

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
				fullPath := dir + "/" + entry.Name()
				executable, err := checkIfCommandIsExecutable(fullPath)
				if err != nil {
					// log.Fatal(err)
					continue
				}
				if entry.Name() == command && executable {
					return fullPath, true
				}
			}
		}
	}
	return "", false
}

func checkIfCommandIsExecutable(command string) (bool, error) {
	info, err := os.Stat(command)
	if err != nil {
		return false, err
	}

	return info.Mode().Perm()&0111 != 0, nil
}

var pathExecutables []string
var pathExecutablesOnce sync.Once

func ListPathExecutables() []string {
	pathExecutablesOnce.Do(func() {
		pathEnv := os.Getenv("PATH")
		if pathEnv == "" {
			return
		}

		paths := filepath.SplitList(pathEnv)
		seen := make(map[string]bool)

		for _, dir := range paths {
			if dir == "" {
				continue
			}
			entries, err := os.ReadDir(dir)
			if err != nil {
				continue
			}

			if len(entries) == 0 {
				continue
			}
			for _, entry := range entries {
				if entry.IsDir() {
					continue
				}
				name := entry.Name()
				if seen[name] {
					continue
				}
				fullPath := filepath.Join(dir, name)
				executable, err := checkIfCommandIsExecutable(fullPath)
				if err != nil || !executable {
					continue
				}
				seen[name] = true
				pathExecutables = append(pathExecutables, name)
			}
		}
	})
	return pathExecutables
}
