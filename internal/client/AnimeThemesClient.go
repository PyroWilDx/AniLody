package client

import (
	"anilody/internal/models"
	"anilody/internal/utils"
	"encoding/json"
	"fmt"
	"io"
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

	fmt.Println("Query AnimeThemes:", queryURL)

	queryResp, err := http.Get(queryURL)
	if err != nil {
		panic(fmt.Sprintf("Error Executing Query\n%v", err))
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(fmt.Sprintf("Error Closing Body\n%v", err))
		}
	}(queryResp.Body)

	if queryResp.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("Error (Received Status Code %d)", queryResp.StatusCode))
	}

	var animeThemesResponse models.AnimeThemesResponse
	err = json.NewDecoder(queryResp.Body).Decode(&animeThemesResponse)
	if err != nil {
		panic(fmt.Sprintf("Error Decoding Response\n%v", err))
	}
	return &animeThemesResponse
}
