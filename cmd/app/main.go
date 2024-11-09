package main

import (
	"anilody/internal/api"
	"anilody/internal/client"
	"anilody/internal/config"
	"fmt"
	"os"
)

func main() {
	userSettings := config.ReadUserSettings("config/Settings.txt")
	err := os.MkdirAll(userSettings.OutPath, os.ModePerm)
	if err != nil {
		panic(fmt.Sprintf("Failed Creating Folder %s\n", userSettings.OutPath))
	}

	aniLodies := api.GetAniLodies(userSettings)
	for i, aniLody := range aniLodies {
		musicName := client.DownloadFile(aniLody, userSettings)
		if musicName != "" {
			fmt.Printf("%d/%d - %s (%s)\n", i+1, len(aniLodies), aniLody.AudioURL, musicName)
		} else {
			fmt.Printf("%d/%d - %s (Already Downloaded)\n", i+1, len(aniLodies), aniLody.AudioURL)
		}
	}
}
