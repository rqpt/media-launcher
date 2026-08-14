package menus

import (
	"errors"
	"path/filepath"

	"rqpt/media-launcher/internal/finder"
	"rqpt/media-launcher/internal/picker"
	"rqpt/media-launcher/internal/player"
)

func OpenRecordingsSubMenu() error {
	recordingsPath := "/home/user/Videos/recordings" //os.Getenv("RECORDINGS_DIR")
	if recordingsPath == "" {
		return errors.New("Environment variable $RECORDINGS_DIR is not set.")
	}

	recordings, err := finder.ListMediaItems(recordingsPath, []string{".mkv", ".mp4"})
	if err != nil {
		return err
	}

	selectedRecording, err := picker.Run(recordings)
	if err != nil {
		return err
	}
	if selectedRecording == "" {
		return nil
	}

	recordingFile := filepath.Join(recordingsPath, selectedRecording)

	return player.Play([]string{recordingFile})
}
