package menus

import (
	"errors"
	"os"
	"path/filepath"

	"rqpt/media-launcher/internal/finder"
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

	movieFile, err := finder.FindMediaFile(fullPath, selectedMovie, movieExtensions)
	if err != nil {
		return err
	}

	return player.Play([]string{movieFile})
}
