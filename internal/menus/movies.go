package menus

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"rqpt/media-launcher/internal/picker"
	"rqpt/media-launcher/internal/player"
)

func OpenMoviesSubMenu() error {
	moviesPath := os.Getenv("MOVIES_DIR")
	if moviesPath == "" {
		return errors.New("Environment variable $MOVIES_DIR is not set.")
	}

	selectedMovie, err := picker.SelectFrom(moviesPath)
	if err != nil {
		return err
	}
	if selectedMovie == "" {
		return nil
	}

	fullPath := filepath.Join(moviesPath, selectedMovie)

	movieExtensions := []string{".mkv", ".mp4"}

	movieFile, err := findMovieFile(fullPath, selectedMovie, movieExtensions)
	if err != nil {
		return err
	}

	return player.Play([]string{movieFile})
}

func findMovieFile(fullPath, selectedMovie string, extensions []string) (string, error) {
	for _, ext := range extensions {
		candidate := filepath.Join(fullPath, selectedMovie+ext)
		if _, err := os.Stat(candidate); err == nil {
			return candidate, nil
		}
	}

	return "", fmt.Errorf("No matching movie file found in %s", fullPath)
}
