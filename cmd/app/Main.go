package main

import (
	"anilody/internal/api"
	"anilody/internal/config"
	"fmt"
)

func main() {
	userSettings := config.ReadUserSettings("config/Settings.txt")
	fmt.Println(api.GetAniLody(userSettings))
}
