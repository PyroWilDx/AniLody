package client

import (
	"anilody/internal/models"
	"context"
	"github.com/machinebox/graphql"
	"log"
	"slices"
)

var baseUrlAniList = "https://graphql.anilist.co"

func GetPublicAniList(userName string, userSettings models.UserSettings) []int {
	client := graphql.NewClient(baseUrlAniList)

	req := graphql.NewRequest(`
        query ($type: MediaType, $userName: String) {
            MediaListCollection(type: $type, userName: $userName) {
                lists {
                    entries {
                        status
                        score
                        media {
                            id
                        }
                    }
                }
            }
        }
    `)

	req.Var("type", "ANIME")
	req.Var("userName", userName)

	var userList models.AniListResponse
	err := client.Run(context.Background(), req, &userList)
	if err != nil {
		log.Fatal(err)
	}

	var userIds []int
	for _, list := range userList.MediaListCollection.Lists {
		for _, entry := range list.Entries {
			if entry.Score < userSettings.MinScore ||
				entry.Score > userSettings.MaxScore {
				continue
			}
			if !slices.Contains(userSettings.StatusList, entry.Status) {
				continue
			}
			userIds = append(userIds, entry.Media.Id)
		}
	}
	return userIds
}
