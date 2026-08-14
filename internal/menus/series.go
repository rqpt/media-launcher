package menus

import (
	"errors"
	"os"
	"path/filepath"

	"rqpt/media-launcher/internal/finder"
	"rqpt/media-launcher/internal/picker"
	"rqpt/media-launcher/internal/player"
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
			selectedShow, err := picker.SelectFrom(seriesPath)
			if err != nil {
				return err
			}
			if selectedShow == "" {
				return nil
			}
			showPath = filepath.Join(seriesPath, selectedShow)
			state = stateSelectSeason

		case stateSelectSeason:
			selectedSeason, err := picker.SelectFrom(showPath)
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
			episodes, err := finder.ListMediaItems(seasonPath, []string{".mkv", ".mp4"})
			if err != nil {
				return err
			}

			selectedEpisode, err := picker.Run(episodes)
			if err != nil {
				return err
			}
			if selectedEpisode == "" {
				state = stateSelectSeason
				continue
			}

			episodeFile := filepath.Join(seasonPath, selectedEpisode)
			return player.Play([]string{episodeFile})
		}
	}
}
