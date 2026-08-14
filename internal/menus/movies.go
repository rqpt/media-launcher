package menus

import (
	"errors"
	"path/filepath"

	"rqpt/media-launcher/internal/finder"
	"rqpt/media-launcher/internal/player"
)

func OpenMoviesSubMenu() error {
	moviesPath := "/home/user/Videos/movies" //os.Getenv("MOVIES_DIR")
	if moviesPath == "" {
		return errors.New("Environment variable $MOVIES_DIR is not set.")
	}

	selectedMovie, err := finder.SelectSubDir(moviesPath)
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
