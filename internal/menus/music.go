package menus

import (
	"errors"
	"os"
	"path/filepath"

	"rqpt/media-launcher/internal/finder"
	"rqpt/media-launcher/internal/picker"
	"rqpt/media-launcher/internal/player"
)

func OpenMusicSubMenu() error {
	musicPath := os.Getenv("MUSIC_DIR")
	if musicPath == "" {
		return errors.New("Environment variable $MUSIC_DIR is not set.")
	}

	for {
		selectedArtist, err := finder.SelectSubDir(musicPath)
		if err != nil {
			return err
		}
		if selectedArtist == "" {
			return nil
		}

		artistPath := filepath.Join(musicPath, selectedArtist)

		for {
			selectedAlbum, err := finder.SelectSubDir(artistPath)
			if err != nil {
				return err
			}
			if selectedAlbum == "" {
				break
			}

			albumPath := filepath.Join(artistPath, selectedAlbum)

			for {
				tracks, err := finder.ListMediaItems(albumPath, []string{".mp3", ".flac"})
				if err != nil {
					return err
				}

				selectedTracks, err := picker.RunMulti(tracks)
				if err != nil {
					return err
				}
				if len(selectedTracks) == 0 {
					break
				}

				var trackFiles []string
				for _, track := range selectedTracks {
					trackFiles = append(trackFiles, filepath.Join(albumPath, track))
				}

				return player.Play(trackFiles, "--wayland-app-id=music")
			}
		}
	}
}
