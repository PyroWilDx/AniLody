package client

import (
	"anilody/internal/models"
	"context"
	"github.com/machinebox/graphql"
	"log"
)

var baseUrlAniList = "https://graphql.anilist.co"

func GetPublicAniList(userName string) []int {
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
			userIds = append(userIds, entry.Media.Id)
		}
	}
	return userIds
}
