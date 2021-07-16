package main

type Entry struct {
	Id     string
	Title  string
	Album  string
	Artist string
	Year   int
	Path   string
}

type Playlist struct {
	SubsonicResponse struct {
		Playlist struct {
			Entry []Entry `json:"entry"`
		} `json:"playlist"`
	} `json:"subsonic-response"`
}
