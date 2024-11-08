package data

import (
	"anilody/internal/models"
	"anilody/internal/utils"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

var baseUrlAnimeThemes = "https://api.animethemes.moe/anime"

func GetAnimeThemes(userIds []int, userSite string, currIndex int, perPage int) ([]models.Anime, int) {
	queryParams := url.Values{}
	queryParams.Add("filter[has]", "resources")
	queryParams.Add("filter[site]", userSite)
	queryParams.Add("filter[external_id]", strconv.Itoa(userIds[0]))
	queryParams.Add("include",
		strings.Join([]string{
			"images",
			"animethemes.song",
			"animethemes.animethemeentries.videos",
			"animethemes.animethemeentries.videos.audio",
		}, ","))

	currUserIds := userIds[currIndex:min(currIndex+perPage, len(userIds))]
	queryParams.Set("filter[external_id]",
		strings.Join(utils.IntSliceToStrSlice(currUserIds), ","))
	animeThemesResponse := execAnimeThemesQuery(queryParams)

	return animeThemesResponse.Anime, animeThemesResponse.Meta.PerPage
}

func execAnimeThemesQuery(queryParams url.Values) *models.AnimeThemesResponse {
	queryURL := fmt.Sprintf("%s?%s", baseUrlAnimeThemes, queryParams.Encode())

	log.Println("Query AnimeThemes:", queryURL)

	queryResp, err := http.Get(queryURL)
	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal("Error Closing Body:", err)
		}
	}(queryResp.Body)

	if queryResp.StatusCode != http.StatusOK {
		log.Fatalf("Error: Received Status Code %d", queryResp.StatusCode)
	}

	var animeThemesResponse models.AnimeThemesResponse
	err = json.NewDecoder(queryResp.Body).Decode(&animeThemesResponse)
	if err != nil {
		log.Fatal("Error Decoding:", err)
	}
	return &animeThemesResponse
}
