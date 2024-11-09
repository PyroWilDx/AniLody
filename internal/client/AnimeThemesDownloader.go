package client

import (
	"anilody/internal/models"
	"anilody/internal/utils"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func DownloadFile(aniLody models.AniLody, userSettings models.UserSettings) string {
	audioFile, err := http.Get(aniLody.AudioURL)
	if err != nil {
		panic(fmt.Sprintf("Failed Downloading %s", aniLody.AudioURL))
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(fmt.Sprintf("Failed Closing Connection"))
		}
	}(audioFile.Body)

	if audioFile.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("Failed Downloading %s (%d)", aniLody.AudioURL, audioFile.StatusCode))
	}

	musicName := calcMusicName(aniLody, userSettings)
	musicPath := filepath.Join(userSettings.OutPath, musicName)

	_, err = os.Stat(musicPath)
	if err == nil {
		return ""
	}

	outFile, err := os.Create(musicPath)
	if err != nil {
		panic(fmt.Sprintf("Failed Creating File %s", musicPath))
	}
	defer func(outFile *os.File) {
		err := outFile.Close()
		if err != nil {
			panic(fmt.Sprintf("Failed Closing File %s", musicPath))
		}
	}(outFile)

	_, err = io.Copy(outFile, audioFile.Body)
	if err != nil {
		panic(fmt.Sprintf("Failed Saving File %s", musicPath))
	}

	return musicName
}

func calcMusicName(aniLody models.AniLody, userSettings models.UserSettings) string {
	musicName := userSettings.MusicNameFormat
	musicName = strings.Replace(musicName, "#Title", aniLody.Title, -1)
	musicName = strings.Replace(musicName, "#Slug", aniLody.Slug, -1)
	musicName = regexp.MustCompile(`[^a-zA-Z0-9\- ]`).ReplaceAllString(musicName, "")
	musicName = handleSpaces(musicName)
	if userSettings.CapWords {
		musicName = capWords(musicName)
	}
	musicName += ".ogg"
	return musicName
}

func handleSpaces(musicName string) string {
	musicName = strings.Trim(musicName, " ")
	musicName = regexp.MustCompile(`\s+`).ReplaceAllString(musicName, " ")
	return musicName
}

func capWords(musicName string) string {
	var musicNameBuilder strings.Builder
	var lastChar byte = ' '
	for i := 0; i < len(musicName); i++ {
		currChar := musicName[i]
		if !utils.IsLetter(lastChar) && utils.IsLowerCaseLetter(currChar) {
			currChar -= 32
		}
		musicNameBuilder.WriteByte(currChar)
		lastChar = currChar
	}
	return musicNameBuilder.String()
}
