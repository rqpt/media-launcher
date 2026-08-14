package menus

import (
	"errors"
	"path/filepath"

	"rqpt/media-launcher/internal/finder"
	"rqpt/media-launcher/internal/picker"
	"rqpt/media-launcher/internal/player"
)

func OpenSeriesSubMenu() error {
	seriesPath := "/home/user/Videos/series" //os.Getenv("SERIES_DIR")
	if seriesPath == "" {
		return errors.New("Environment variable $SERIES_DIR is not set.")
	}

	for {
		selectedShow, err := finder.SelectSubDir(seriesPath)
		if err != nil {
			return err
		}
		if selectedShow == "" {
			return nil
		}

		showPath := filepath.Join(seriesPath, selectedShow)

		for {
			selectedSeason, err := finder.SelectSubDir(showPath)
			if err != nil {
				return err
			}
			if selectedSeason == "" {
				break
			}

			seasonPath := filepath.Join(showPath, selectedSeason)

			for {
				episodes, err := finder.ListMediaItems(seasonPath, []string{".mkv", ".mp4"})
				if err != nil {
					return err
				}

				selectedEpisode, err := picker.Run(episodes)
				if err != nil {
					return err
				}
				if selectedEpisode == "" {
					break
				}

				episodeFile := filepath.Join(seasonPath, selectedEpisode)

				return player.Play([]string{episodeFile})
			}
		}
	}
}
