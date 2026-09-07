package menus

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/rqpt/media-launcher/internal/player"
	"github.com/rqpt/picker"
)

type seriesMenuState int

const (
	stateSelectShow seriesMenuState = iota
	stateSelectSeason
	stateSelectEpisode
)

func OpenSeriesSubMenu() error {
	seriesPath := os.Getenv("SERIES_DIR")
	if seriesPath == "" {
		return errors.New("environment variable $SERIES_DIR is not set")
	}

	state := stateSelectShow

	var showPath, seasonPath string

	for {
		switch state {

		case stateSelectShow:
			selectedShow, err := picker.RunSubdirs(seriesPath)
			if err != nil {
				return err
			}
			if selectedShow == "" {
				return nil
			}
			showPath = filepath.Join(seriesPath, selectedShow)
			state = stateSelectSeason

		case stateSelectSeason:
			selectedSeason, err := picker.RunSubdirs(showPath)
			if err != nil {
				return err
			}
			if selectedSeason == "" {
				state = stateSelectShow
				continue
			}
			seasonPath = filepath.Join(showPath, selectedSeason)
			state = stateSelectEpisode

		case stateSelectEpisode:
			episodes, err := picker.ListFiles(
				seasonPath,
				[]string{".mkv", ".mp4"},
			)
			if err != nil {
				return err
			}

			selectedEpisodes, err := picker.RunMulti(episodes)
			if err != nil {
				return err
			}
			if len(selectedEpisodes) == 0 {
				state = stateSelectSeason
				continue
			}

			var episodeFiles []string
			for _, episode := range selectedEpisodes {
				episodeFiles = append(episodeFiles, filepath.Join(seasonPath, episode))
			}

			return player.Play(episodeFiles)
		}
	}
}
