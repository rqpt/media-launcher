package main

import (
	"fmt"
	"log"
	"os"

	"rqpt/media-launcher/internal/menus"
	"rqpt/media-launcher/internal/picker"
)

func main() {
	menuItems := []string{"movies", "series", "music", "recordings"}

	for {
		selectedMenuItem, err := picker.Run(menuItems)
		if err != nil {
			log.Fatalf("Error running picker: %v", err)
		}
		if selectedMenuItem == "" {
			return
		}

		var menuErr error

		switch selectedMenuItem {
		case "movies":
			menuErr = menus.OpenMoviesSubMenu()
		case "series":
			menuErr = menus.OpenSeriesSubMenu()
		case "recordings":
			menuErr = menus.OpenRecordingsSubMenu()
		case "music":
			menuErr = menus.OpenMusicSubMenu()
		}

		if menuErr != nil {
			showErrorAndPause(menuErr.Error())
		}
	}
}

func showErrorAndPause(msg string) {
	fmt.Fprintf(os.Stderr, "Error: %s\n\nPress Enter to return...", msg)

	var dummy string

	fmt.Scanln(&dummy)
}
