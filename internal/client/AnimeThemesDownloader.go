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

const ffMpeg = "bin/ffmpeg"
const upScayl = "bin/UpScayl/upscayl-bin"

func FetchAniLody(aniLody models.AniLody, userSettings models.UserSettings) string {
	musicName := calcMusicName(aniLody, userSettings)
	musicPathOgg := filepath.Join(userSettings.OutPath, musicName+".ogg")
	musicPathMp3 := filepath.Join(userSettings.OutPath, musicName+".mp3")

	_, err := os.Stat(musicPathMp3)
	if err == nil {
		return ""
	}

	dlOgg(aniLody, musicPathOgg)

	convertOggToMp3(musicPathOgg, musicPathMp3)
	err = os.Remove(musicPathOgg)
	if err != nil {
		panic(fmt.Sprintf("Failed Removing File %s\n%v", musicPathOgg, err))
	}

	imgPath := musicPathMp3 + ".jpg"
	dlImage(aniLody.ImageURL, imgPath)
	if userSettings.UpScaleImageWidth != "0" {
		upScaleImage(imgPath, userSettings.UpScaleImageWidth)
	}
	applyImage(musicPathMp3, imgPath)
	err = os.Remove(imgPath)
	if err != nil {
		panic(fmt.Sprintf("Failed Removing File %s\n%v", imgPath, err))
	}

	return musicName
}

func dlOgg(aniLody models.AniLody, musicPathOgg string) {
	audioFile, err := http.Get(aniLody.AudioURL)
	if err != nil {
		panic(fmt.Sprintf("Failed Downloading %s\n%v", aniLody.AudioURL, err))
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(fmt.Sprintf("Failed Closing Connection %s\n%v", aniLody.AudioURL, err))
		}
	}(audioFile.Body)

	if audioFile.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("Failed Downloading %s (%d)", aniLody.AudioURL, audioFile.StatusCode))
	}

	outFile, err := os.Create(musicPathOgg)
	if err != nil {
		panic(fmt.Sprintf("Failed Creating File %s\n%v", musicPathOgg, err))
	}
	defer func(outFile *os.File) {
		err := outFile.Close()
		if err != nil {
			panic(fmt.Sprintf("Failed Closing File %s\n%v", musicPathOgg, err))
		}
	}(outFile)

	_, err = io.Copy(outFile, audioFile.Body)
	if err != nil {
		panic(fmt.Sprintf("Failed Saving File %s\n%v", musicPathOgg, err))
	}
}

func calcMusicName(aniLody models.AniLody, userSettings models.UserSettings) string {
	musicName := userSettings.MusicNameFormat
	musicName = strings.Replace(musicName, "#AnimeTitle", aniLody.AnimeTitle, -1)
	musicName = strings.Replace(musicName, "#Slug", aniLody.Slug, -1)
	musicName = strings.Replace(musicName, "#SongTitle", aniLody.SongTitle, -1)
	musicName = cleanMusicName(musicName)
	musicName = handleSpaces(musicName)
	if userSettings.FmtNums {
		musicName = fmtNums(musicName)
	}
	if userSettings.CapWords {
		musicName = capWords(musicName, userSettings.LowWords)
	}
	return musicName
}

func cleanMusicName(musicName string) string {
	musicName = strings.Replace(musicName, "<", "", -1)
	musicName = strings.Replace(musicName, ">", "", -1)
	musicName = strings.Replace(musicName, ":", "", -1)
	musicName = strings.Replace(musicName, "\"", "", -1)
	musicName = strings.Replace(musicName, "/", "", -1)
	musicName = strings.Replace(musicName, "\\", "", -1)
	musicName = strings.Replace(musicName, "|", "", -1)
	musicName = strings.Replace(musicName, "?", "", -1)
	musicName = strings.Replace(musicName, "*", "", -1)
	return musicName
}

func handleSpaces(musicName string) string {
	musicName = strings.Trim(musicName, " ")
	musicName = regexp.MustCompile(`\s+`).ReplaceAllString(musicName, " ")
	return musicName
}

func fmtNums(musicName string) string {
	musicName = regexp.MustCompile(`\b(\d{1,2}(st|nd|rd|th))\b`).ReplaceAllStringFunc(musicName, func(m string) string {
		repl, in := utils.OrdMap[m]
		if in {
			return repl
		}
		return m
	})
	return musicName
}

func capWords(musicName string, lowWords bool) string {
	var musicNameBuilder strings.Builder
	var prevChar byte = ' '
	for i := 0; i < len(musicName); i++ {
		currChar := musicName[i]
		if !utils.IsLetter(prevChar) && prevChar != '\'' && utils.IsLowerCaseLetter(currChar) {
			currChar -= 32
		}
		if lowWords && utils.IsLetter(prevChar) && utils.IsUpperCaseLetter(currChar) {
			currChar += 32
		}
		musicNameBuilder.WriteByte(currChar)
		prevChar = currChar
	}
	return musicNameBuilder.String()
}

func convertOggToMp3(musicPathOgg string, musicPathMp3 string) {
	cmd := exec.Command(ffMpeg,
		"-i",
		musicPathOgg,
		"-vn",
		"-ar",
		"44100",
		"-ac",
		"2",
		"-b:a",
		"192k",
		musicPathMp3)
	err := cmd.Run()
	if err != nil {
		panic(fmt.Sprintf("Failed Converting File %s\n%v", musicPathOgg, err))
	}
}

func dlImage(imgURL string, imgPath string) {
	imgFile, err := http.Get(imgURL)
	if err != nil {
		panic(fmt.Sprintf("Failed Downloading %s\n%v", imgURL, err))
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(fmt.Sprintf("Failed Closing Connection %s\n%v", imgURL, err))
		}
	}(imgFile.Body)

	if imgFile.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("Failed Downloading %s (%d)", imgURL, imgFile.StatusCode))
	}

	outFile, err := os.Create(imgPath)
	if err != nil {
		panic(fmt.Sprintf("Failed Creating File %s\n%v", imgPath, err))
	}
	defer func(outFile *os.File) {
		err := outFile.Close()
		if err != nil {
			panic(fmt.Sprintf("Failed Closing File %s\n%v", imgPath, err))
		}
	}(outFile)

	_, err = io.Copy(outFile, imgFile.Body)
	if err != nil {
		panic(fmt.Sprintf("Failed Saving File %s\n%v", imgPath, err))
	}
}

func upScaleImage(imgPath string, imgWidth string) {
	pngImgPath := imgPath + ".png"

	cmd := exec.Command(upScayl,
		"-i", imgPath,
		"-o", pngImgPath,
		"-n", "realesr-animevideov3-x4",
		"-w", imgWidth)
	err := cmd.Run()
	if err != nil {
		panic(fmt.Sprintf("Failed UpScaling Image\n%v", err))
	}

	err = os.Remove(imgPath)
	if err != nil {
		panic(fmt.Sprintf("Failed Removing File %s\n%v", imgPath, err))
	}

	cmd = exec.Command(ffMpeg,
		"-i", pngImgPath,
		"-q:v", "2",
		imgPath)
	err = cmd.Run()
	if err != nil {
		panic(fmt.Sprintf("Failed Converting Image\n%v", err))
	}

	err = os.Remove(pngImgPath)
	if err != nil {
		panic(fmt.Sprintf("Failed Removing File %s\n%v", pngImgPath, err))
	}
}

func applyImage(musicPathMp3 string, imgPath string) {
	tmpOutputPath := musicPathMp3 + ".tmp.mp3"

	cmd := exec.Command(ffMpeg,
		"-i",
		musicPathMp3,
		"-i",
		imgPath,
		"-map",
		"0:0",
		"-map",
		"1:0",
		"-c",
		"copy",
		"-id3v2_version",
		"3",
		"-metadata:s:v",
		"title='KeyVisual'",
		"-metadata:s:v",
		"comment='KeyVisual'",
		tmpOutputPath,
	)
	err := cmd.Run()
	if err != nil {
		panic(fmt.Sprintf("Failed Adding Image\n%v", err))
	}

	err = os.Rename(tmpOutputPath, musicPathMp3)
	if err != nil {
		panic(fmt.Sprintf("Failed Replacing File %s\n%v", tmpOutputPath, err))
	}
}
