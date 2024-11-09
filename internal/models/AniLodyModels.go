package models

type AniLodyResponse struct {
	AniLodies []AniLody `json:"aniLodies"`
}

type AniLody struct {
	ImageURL string `json:"imageURL"`
	Title    string `json:"title"`
	Slug     string `json:"slug"`
	AudioURL string `json:"audioURL"`
}
