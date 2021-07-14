package main

type Playlists struct {
	SubsonicResponse struct {
		status    string
		Playlists struct {
			Playlist []struct {
				Id   string `json:"id"`
				Name string `json:"name"`
			} `json:"playlist"`
		} `json:"playlists"`
	} `json:"subsonic-response"`
}
