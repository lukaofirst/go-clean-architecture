package utils

import (
	"os"
	"path/filepath"
)

func FindGoRootDirectory(cwd string) string {
	// Traverse upwards from the current working directory to find the go.mod file
	for {
		// Check if the go.mod file exists in the current directory
		if _, err := os.Stat(filepath.Join(cwd, "go.mod")); err == nil {
			return cwd
		}
		// If not, move up one directory
		parentDir := filepath.Dir(cwd)
		if parentDir == cwd {
			// Reached the root of the file system, stop
			break
		}
		cwd = parentDir
	}
	return ""
}
