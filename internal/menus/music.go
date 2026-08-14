package menus

import (
	"errors"
	"os"
	"path/filepath"

	"rqpt/media-launcher/internal/picker"
	"rqpt/media-launcher/internal/player"
)

type musicMenuState int

const (
	stateSelectArtist musicMenuState = iota
	stateSelectAlbum
	stateSelectTracks
)

func OpenMusicSubMenu() error {
	musicPath := os.Getenv("MUSIC_DIR")
	if musicPath == "" {
		return errors.New("Environment variable $MUSIC_DIR is not set.")
	}

	state := stateSelectArtist

	var artistPath, albumPath string

	for {
		switch state {
		case stateSelectArtist:
			selectedArtist, err := picker.RunFrom(musicPath)
			if err != nil {
				return err
			}
			if selectedArtist == "" {
				return nil
			}

			artistPath = filepath.Join(musicPath, selectedArtist)
			state = stateSelectAlbum

		case stateSelectAlbum:
			selectedAlbum, err := picker.RunFrom(artistPath)
			if err != nil {
				return err
			}
			if selectedAlbum == "" {
				state = stateSelectArtist
				continue
			}

			albumPath = filepath.Join(artistPath, selectedAlbum)
			state = stateSelectTracks

		case stateSelectTracks:
			tracks, err := picker.ListMediaItems(
				albumPath,
				[]string{".mp3", ".flac"},
			)
			if err != nil {
				return err
			}

			selectedTracks, err := picker.RunMulti(tracks)
			if err != nil {
				return err
			}
			if len(selectedTracks) == 0 {
				state = stateSelectAlbum
				continue
			}

			var trackFiles []string
			for _, track := range selectedTracks {
				trackFiles = append(trackFiles, filepath.Join(albumPath, track))
			}

			return player.Play(trackFiles, "--wayland-app-id=music")
		}
	}
}
