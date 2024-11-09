package models

type AnimeThemesResponse struct {
	Anime []Anime `json:"anime"`
	Links Links   `json:"links"`
	Meta  Meta    `json:"meta"`
}

type Anime struct {
	Id          int          `json:"id"`
	Name        string       `json:"name"`
	MediaFormat string       `json:"media_format"`
	Season      string       `json:"season"`
	Slug        string       `json:"slug"`
	Synopsis    string       `json:"synopsis"`
	Year        int          `json:"year"`
	Images      []Image      `json:"images"`
	AnimeThemes []AnimeTheme `json:"animethemes"`
}

type Image struct {
	Id    int    `json:"id"`
	Facet string `json:"facet"`
	Path  string `json:"path"`
	Link  string `json:"link"`
}

type AnimeTheme struct {
	Id                int               `json:"id"`
	Sequence          int               `json:"sequence"`
	Slug              string            `json:"slug"`
	Type              string            `json:"type"`
	AnimeThemeEntries []AnimeThemeEntry `json:"animethemeentries"`
	Song              Song              `json:"song"`
}

type AnimeThemeEntry struct {
	Id       int     `json:"id"`
	Episodes string  `json:"episodes"`
	Notes    string  `json:"notes"`
	NSFW     bool    `json:"nsfw"`
	Spoiler  bool    `json:"spoiler"`
	Version  int     `json:"version"`
	Videos   []Video `json:"videos"`
}

type Video struct {
	Id         int    `json:"id"`
	BaseName   string `json:"basename"`
	FileName   string `json:"filename"`
	Lyrics     bool   `json:"lyrics"`
	Nc         bool   `json:"nc"`
	Overlap    string `json:"overlap"`
	Path       string `json:"path"`
	Resolution int    `json:"resolution"`
	Size       int    `json:"size"`
	Source     string `json:"source"`
	Subbed     bool   `json:"subbed"`
	Uncen      bool   `json:"uncen"`
	Tags       string `json:"tags"`
	Link       string `json:"link"`
	Audio      Audio  `json:"audio"`
}

type Audio struct {
	Id       int    `json:"id"`
	BaseName string `json:"basename"`
	FileName string `json:"filename"`
	Path     string `json:"path"`
	Size     int    `json:"size"`
	Link     string `json:"link"`
}

type Song struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

type Links struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Prev  string `json:"prev"`
	Next  string `json:"next"`
}

type Meta struct {
	CurrentPage int    `json:"current_page"`
	From        int    `json:"from"`
	Path        string `json:"path"`
	PerPage     int    `json:"per_page"`
	To          int    `json:"to"`
}
