package picker

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ListMediaItems(parentPath string, extensions []string) ([]string, error) {
	entries, err := os.ReadDir(parentPath)
	if err != nil {
		return nil, fmt.Errorf("Could not read directory %s: %w", parentPath, err)
	}

	var items []string
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		ext := filepath.Ext(entry.Name())
		for _, validExt := range extensions {
			if strings.EqualFold(ext, validExt) {
				items = append(items, entry.Name())
				break
			}
		}
	}

	if len(items) == 0 {
		return nil, fmt.Errorf("No media files found in %s", parentPath)
	}

	return items, nil
}

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
