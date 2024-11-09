package client

import (
	"anilody/internal/models"
	"anilody/internal/utils"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

const imgName = "Tmp.jpg"

func FetchAniLody(aniLody models.AniLody, userSettings models.UserSettings) string {
	musicName, musicPathOgg, musicPathMp3 := fetchAniLodyOgg(aniLody, userSettings)
	if musicName == "" {
		return ""
	}

	convertOggToMp3(musicPathOgg, musicPathMp3)
	err := os.Remove(musicPathOgg)
	if err != nil {
		panic(fmt.Sprintf("Failed Removing File %s", musicPathOgg))
	}

	imgPath := fetchImage(aniLody.ImageURL, userSettings)
	addImage(musicPathMp3, imgPath)
	//err = os.Remove(imgPath)
	//if err != nil {
	//	panic(fmt.Sprintf("Failed Removing File %s", imgPath))
	//}

	return musicName
}

func fetchAniLodyOgg(aniLody models.AniLody, userSettings models.UserSettings) (string, string, string) {
	audioFile, err := http.Get(aniLody.AudioURL)
	if err != nil {
		panic(fmt.Sprintf("Failed Downloading %s", aniLody.AudioURL))
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(fmt.Sprintf("Failed Closing Connection %s", aniLody.AudioURL))
		}
	}(audioFile.Body)

	if audioFile.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("Failed Downloading %s (%d)", aniLody.AudioURL, audioFile.StatusCode))
	}

	musicName := calcMusicName(aniLody, userSettings)
	musicPathOgg := filepath.Join(userSettings.OutPath, musicName+".ogg")
	musicPathMp3 := filepath.Join(userSettings.OutPath, musicName+".mp3")

	_, err = os.Stat(musicPathMp3)
	if err == nil {
		return "", "", ""
	}

	outFile, err := os.Create(musicPathOgg)
	if err != nil {
		panic(fmt.Sprintf("Failed Creating File %s", musicPathOgg))
	}
	defer func(outFile *os.File) {
		err := outFile.Close()
		if err != nil {
			panic(fmt.Sprintf("Failed Closing File %s", musicPathOgg))
		}
	}(outFile)

	_, err = io.Copy(outFile, audioFile.Body)
	if err != nil {
		panic(fmt.Sprintf("Failed Saving File %s", musicPathOgg))
	}

	return musicName, musicPathOgg, musicPathMp3
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

func convertOggToMp3(musicPathOgg string, musicPathMp3 string) {
	cmd := exec.Command("bin/ffmpeg", "-i", musicPathOgg, "-vn", "-ar", "44100", "-ac", "2", "-b:a", "192k", musicPathMp3)
	err := cmd.Run()
	if err != nil {
		panic(fmt.Sprintf("Failed Converting File %s", musicPathOgg))
	}
}

func fetchImage(imgURL string, userSettings models.UserSettings) string {
	imgFile, err := http.Get(imgURL)
	if err != nil {
		panic(fmt.Sprintf("Failed Downloading %s", imgURL))
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(fmt.Sprintf("Failed Closing Connection %s", imgURL))
		}
	}(imgFile.Body)

	if imgFile.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("Failed Downloading %s (%d)", imgURL, imgFile.StatusCode))
	}

	imgPath := filepath.Join(userSettings.OutPath, imgName)

	outFile, err := os.Create(imgPath)
	if err != nil {
		panic(fmt.Sprintf("Failed Creating File %s", imgPath))
	}
	defer func(outFile *os.File) {
		err := outFile.Close()
		if err != nil {
			panic(fmt.Sprintf("Failed Closing File %s", imgPath))
		}
	}(outFile)

	_, err = io.Copy(outFile, imgFile.Body)
	if err != nil {
		panic(fmt.Sprintf("Failed Saving File %s", imgPath))
	}

	return imgPath
}

func addImage(musicPathMp3 string, imgPath string) {

}
