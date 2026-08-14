package finder

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

func FindMediaFile(dirPath, baseName string, extensions []string) (string, error) {
	for _, ext := range extensions {
		candidate := filepath.Join(dirPath, baseName+ext)
		if _, err := os.Stat(candidate); err == nil {
			return candidate, nil
		}
	}

	return "", fmt.Errorf("No matching media files found in %s", dirPath)
}
