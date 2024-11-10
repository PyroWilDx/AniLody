package models

type AniLodyResponse struct {
	AniLodies []AniLody `json:"aniLodies"`
}

type AniLody struct {
	ImageURL   string `json:"imageURL"`
	AnimeTitle string `json:"animeTitle"`
	Slug       string `json:"slug"`
	SongTitle  string `json:"songTitle"`
	AudioURL   string `json:"audioURL"`
}
