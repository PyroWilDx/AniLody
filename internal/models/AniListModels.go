package models

type AniListResponse struct {
	MediaListCollection MediaListCollection `json:"MediaListCollection"`
}

type MediaListCollection struct {
	Lists []List `json:"lists"`
}

type List struct {
	Entries []Entry `json:"entries"`
}

type Entry struct {
	Status string  `json:"status"`
	Score  float32 `json:"score"`
	Media  Media   `json:"media"`
}

type Media struct {
	Id int `json:"id"`
}
