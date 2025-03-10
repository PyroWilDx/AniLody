package config

import (
	"anilody/internal/models"
	"anilody/internal/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadUserSettings(sPath string) models.UserSettings {
	sFile, err := os.Open(sPath)
	if err != nil {
		panic(fmt.Sprintf("Error Opening %s\n%v", sPath, err))
	}
	defer func(sFile *os.File) {
		err := sFile.Close()
		if err != nil {
			panic(fmt.Sprintf("Error Closing %s\n%v", sPath, err))
		}
	}(sFile)

	var userSettings models.UserSettings
	fileScanner := bufio.NewScanner(sFile)
	for fileScanner.Scan() {
		currLine := fileScanner.Text()
		if currLine == "" || currLine[0] == '#' {
			continue
		}

		keyValue := strings.SplitN(currLine, "=", 2)
		if len(keyValue) != 2 {
			panic(fmt.Sprintf("Error Reading %s", sPath))
		}

		key := keyValue[0]
		value := keyValue[1]
		updateUserSettings(key, value, &userSettings)
	}

	return userSettings
}

func updateUserSettings(key string, value string, userSettings *models.UserSettings) {
	switch key {
	case "userName":
		userSettings.UserName = value
	case "userSite":
		userSettings.UserSite = value

	case "outPath":
		userSettings.OutPath = value
	case "threadsCount":
		userSettings.ThreadsCount = utils.ParseInt(value)

	case "musicNameFormat":
		userSettings.MusicNameFormat = value
	case "capWords":
		userSettings.CapWords = value != "0"
	case "lowWords":
		userSettings.LowWords = value != "0"
	case "fmtNums":
		userSettings.FmtNums = value != "0"

	case "applyImage":
		userSettings.ApplyImage = value != "0"
	case "upScaleImageWidth":
		userSettings.UpScaleImageWidth = value

	case "incOp":
		userSettings.IncOp = value != "0"
	case "incEd":
		userSettings.IncEd = value != "0"

	case "minScore":
		userSettings.MinScore = utils.ParseFloat32(value)
	case "maxScore":
		userSettings.MaxScore = utils.ParseFloat32(value)
	case "statusList":
		userSettings.StatusList = strings.Split(value, "|")
		for i, status := range userSettings.StatusList {
			userSettings.StatusList[i] = strings.ToUpper(status)
		}
	}
}
