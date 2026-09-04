package menus

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/rqpt/media-launcher/internal/player"
	"github.com/rqpt/picker"
)

func OpenRecordingsSubMenu() error {
	recordingsPath := os.Getenv("RECORDINGS_DIR")
	if recordingsPath == "" {
		return errors.New("Environment variable $RECORDINGS_DIR is not set.")
	}

	recordings, err := picker.ListFiles(
		recordingsPath,
		[]string{".mkv", ".mp4"},
	)
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
