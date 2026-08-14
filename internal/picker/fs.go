package picker

import (
	"fmt"
	"os"
)

func SelectFrom(parentPath string) (string, error) {
	dirs, err := getSubdirs(parentPath)
	if err != nil {
		return "", err
	}
	return Run(dirs)
}

func getSubdirs(parentPath string) ([]string, error) {
	entries, err := os.ReadDir(parentPath)
	if err != nil {
		return nil, fmt.Errorf("Could not read directory %s: %w", parentPath, err)
	}

	var dirs []string
	for _, entry := range entries {
		if entry.IsDir() {
			dirs = append(dirs, entry.Name())
		}
	}

	if len(dirs) == 0 {
		return nil, fmt.Errorf("No subdirectories found in %s", parentPath)
	}

	return dirs, nil
}
