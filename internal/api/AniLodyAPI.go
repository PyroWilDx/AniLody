package api

import (
	"anilody/internal/client"
	"anilody/internal/models"
	"strings"
)

func GetAniLodies(userSettings models.UserSettings) []models.AniLody {
	var aniLodies []models.AniLody

	var userIds []int
	switch userSettings.UserSite {
	case "AniList":
		userIds = client.GetPublicAniList(userSettings.UserName, userSettings)
	}

	userAnimes, perPage := client.GetAnimeThemes(userIds, userSettings.UserSite, 0, 1)
	for i := 0; i < len(userIds); i += perPage {
		userAnimes, perPage = client.GetAnimeThemes(userIds, userSettings.UserSite, i, perPage)

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
							ImageURL:   anime.Images[0].Link,
							AnimeTitle: anime.Name,
							Slug:       getSlug(animeTheme.Slug),
							SongTitle:  animeTheme.Song.Title,
							AudioURL:   video.Audio.Link,
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
