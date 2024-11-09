package api

import (
	"anilody/internal/data"
	"anilody/internal/models"
	"strings"
)

func GetAniLody(userSettings models.UserSettings) []models.AniLody {
	var aniLodies []models.AniLody

	userIds := data.GetPublicAniList(userSettings.UserName)

	userAnimes, perPage := data.GetAnimeThemes(userIds, userSettings.UserSite, 0, 1)
	for i := 0; i < len(userIds); i += perPage {
		userAnimes, perPage = data.GetAnimeThemes(userIds, userSettings.UserSite, i, perPage)

		for _, anime := range userAnimes {
			for _, animeTheme := range anime.AnimeThemes {
				if (isOp(animeTheme.Slug) && !userSettings.IncOp) ||
					(isEd(animeTheme.Slug) && !userSettings.IncEd) {
					continue
				}

				for _, animeThemeEntry := range animeTheme.AnimeThemeEntries {
					if animeThemeEntry.Version >= 2 {
						break
					}

					for _, video := range animeThemeEntry.Videos {
						aniLodies = append(aniLodies, models.AniLody{
							ImageURL: anime.Images[0].Link,
							Title:    anime.Name,
							Slug:     getSlug(animeTheme.Slug),
							AudioURL: video.Audio.Link,
						})

						break
					}
				}
			}
		}
	}

	return aniLodies
}

func isOp(aSlug string) bool {
	return aSlug[0] == 'O'
}

func isEd(aSlug string) bool {
	return aSlug[0] == 'E'
}

func getSlug(aSlug string) string {
	return string(aSlug[0]) + strings.ToLower(aSlug[1:])
}
