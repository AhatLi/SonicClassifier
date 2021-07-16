package main

type Playlist struct {
	SubsonicResponse struct {
		Playlist struct {
			Entry []Entry `json:"entry"`
		} `json:"playlist"`
	} `json:"subsonic-response"`
}
