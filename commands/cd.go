package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Cd(path string) {
	expanded := expandPath(path)
	if err := os.Chdir(expanded); err != nil {
		fmt.Printf("%v\n", err)
	}
}

func expandPath(path string) string {
	if strings.HasPrefix(path, "~") {
		home, err := os.UserHomeDir()
		if err != nil {
			return path
		}
		if path == "~" {
			return home
		}
		path = filepath.Join(home, path[2:])
	}
	return filepath.Clean(path)
}
