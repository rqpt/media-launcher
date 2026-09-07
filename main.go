package main

import (
	"fmt"
	"log"
	"maps"
	"os"
	"slices"

	"github.com/rqpt/media-launcher/internal/menus"
	"github.com/rqpt/picker"
)

func main() {
	menus := map[string]func() error{
		"movies":     menus.OpenMoviesSubMenu,
		"series":     menus.OpenSeriesSubMenu,
		"recordings": menus.OpenRecordingsSubMenu,
		"music":      menus.OpenMusicSubMenu,
	}

	for {
		selectedMenuItem, err := picker.Run(
			slices.Sorted(maps.Keys(menus)),
		)
		if err != nil {
			log.Fatalf("Error running picker: %v", err)
		}
		if selectedMenuItem == "" {
			return
		}

		menuErr := menus[selectedMenuItem]()
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
