package main

import (
	"anilody/internal/api"
	"anilody/internal/client"
	"anilody/internal/config"
	"fmt"
	"os"
	"sync"
)

func main() {
	userSettings := config.ReadUserSettings("config/Settings.txt")
	err := os.MkdirAll(userSettings.OutPath, os.ModePerm)
	if err != nil {
		panic(fmt.Sprintf("Failed Creating Folder %s\n%v", userSettings.OutPath, err))
	}

	aniLodies := api.GetAniLodies(userSettings)

	fmt.Println("Starting Download...")

	var wg sync.WaitGroup
	semaphore := make(chan struct{}, userSettings.ThreadsCount)

	i := 0
	for _, aniLody := range aniLodies {
		wg.Add(1)
		semaphore <- struct{}{}

		go func() {
			defer wg.Done()
			defer func() { <-semaphore }()

			musicName := client.FetchAniLody(aniLody, userSettings)
			if musicName != "" {
				fmt.Printf("%d/%d - %s (%s)\n", i+1, len(aniLodies), aniLody.AudioURL, musicName)
			} else {
				fmt.Printf("%d/%d - %s (Already Downloaded)\n", i+1, len(aniLodies), aniLody.AudioURL)
			}
			i++
		}()
	}

	wg.Wait()

	fmt.Println("Finished Download.")
}
